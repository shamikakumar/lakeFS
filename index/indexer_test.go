package index_test

import (
	"testing"
	"time"

	"github.com/treeverse/lakefs/index/path"

	"github.com/treeverse/lakefs/ident"

	"github.com/treeverse/lakefs/db"
	"golang.org/x/xerrors"

	"github.com/treeverse/lakefs/index/model"
	"github.com/treeverse/lakefs/testutil"

	"github.com/treeverse/lakefs/index"
)

const testBranch = "testBranch"

type Command int

const (
	write Command = iota
	commit
	revertTree
	revertObj
	deleteEntry
)

func TestKVIndex_GetCommit(t *testing.T) {
	kv, repo, closer := testutil.GetIndexStoreWithRepo(t, 1)
	defer closer()

	kvIndex := index.NewKVIndex(kv)

	commit, err := kvIndex.Commit(repo.GetRepoId(), repo.GetDefaultBranch(), "test msg", "committer", nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("get commit", func(t *testing.T) {
		got, err := kvIndex.GetCommit(repo.GetRepoId(), commit.GetAddress())
		if err != nil {
			t.Fatal(err)
		}
		//compare commitrer
		if commit.GetCommitter() != got.GetCommitter() {
			t.Errorf("got wrong committer. got:%s, expected:%s", got.GetCommitter(), commit.GetCommitter())
		}
		//compare message
		if commit.GetMessage() != got.GetMessage() {
			t.Errorf("got wrong message. got:%s, expected:%s", got.GetMessage(), commit.GetMessage())
		}
	})

	t.Run("get non existing commit - expect error", func(t *testing.T) {
		_, err := kvIndex.GetCommit(repo.RepoId, "nonexistingcommitid")
		if !xerrors.Is(err, db.ErrNotFound) {
			t.Errorf("expected to get not found error for non existing commit")
		}
	})

}

func TestKVIndex_RevertCommit(t *testing.T) {
	kv, repo, closer := testutil.GetIndexStoreWithRepo(t, 1)
	defer closer()

	kvIndex := index.NewKVIndex(kv)
	firstEntry := &model.Entry{
		Name: "bar",
		Type: model.Entry_OBJECT,
	}
	err := kvIndex.WriteEntry(repo.RepoId, repo.DefaultBranch, "", firstEntry)
	if err != nil {
		t.Fatal(err)
	}

	commit, err := kvIndex.Commit(repo.RepoId, repo.DefaultBranch, "test msg", "committer", nil)
	if err != nil {
		t.Fatal(err)
	}
	commitId := ident.Hash(commit)
	_, err = kvIndex.CreateBranch(repo.RepoId, testBranch, commitId)
	if err != nil {
		t.Fatal(err)
	}
	secondEntry := &model.Entry{
		Name: "foo",
		Type: model.Entry_OBJECT,
	}
	// commit second entry to default branch
	err = kvIndex.WriteEntry(repo.RepoId, repo.DefaultBranch, "", secondEntry)
	if err != nil {
		t.Fatal(err)
	}
	_, err = kvIndex.Commit(repo.RepoId, repo.DefaultBranch, "test msg", "committer", nil)
	if err != nil {
		t.Fatal(err)
	}
	//commit second entry to test branch
	err = kvIndex.WriteEntry(repo.RepoId, testBranch, "", secondEntry)
	if err != nil {
		t.Fatal(err)
	}
	_, err = kvIndex.Commit(repo.RepoId, testBranch, "test msg", "committer", nil)
	if err != nil {
		t.Fatal(err)
	}

	err = kvIndex.RevertCommit(repo.RepoId, repo.DefaultBranch, commitId)
	if err != nil {
		t.Fatal(err)
	}

	// test entry1 exists
	te, err := kvIndex.ReadEntryObject(repo.RepoId, repo.DefaultBranch, "bar")
	if err != nil {
		t.Fatal(err)
	}
	if te.Name != firstEntry.Name {
		t.Fatalf("missing data from requested commit")
	}
	// test secondEntry does not exist
	_, err = kvIndex.ReadEntryObject(repo.RepoId, repo.DefaultBranch, "foo")
	if !xerrors.Is(err, db.ErrNotFound) {
		t.Fatalf("missing data from requested commit")
	}

	// test secondEntry exists on test branch
	_, err = kvIndex.ReadEntryObject(repo.RepoId, testBranch, "foo")
	if err != nil {
		if xerrors.Is(err, db.ErrNotFound) {
			t.Fatalf("errased data from test branch after revert from defult branch")
		} else {
			t.Fatal(err)
		}
	}

}

func TestKVIndex_RevertPath(t *testing.T) {

	type Action struct {
		command Command
		path    string
	}
	testData := []struct {
		Name           string
		Actions        []Action
		ExpectExisting []string
		ExpectMissing  []string
		ExpectedError  error
	}{
		{
			"commit and revert",
			[]Action{
				{write, "a/foo"},
				{commit, ""},
				{write, "a/bar"},
				{revertTree, "a/"},
			},

			[]string{"a/foo"},
			[]string{"a/bar"},
			nil,
		},
		{
			"reset - commit and revert on root",
			[]Action{
				{write, "foo"},
				{commit, ""},
				{write, "bar"},
				{revertTree, ""},
			},
			[]string{"foo"},
			[]string{"bar"},
			nil,
		},
		{
			"only revert",
			[]Action{
				{write, "foo"},
				{write, "a/foo"},
				{write, "a/bar"},
				{revertTree, "a/"},
			},
			[]string{"foo"},
			[]string{"a/bar", "a/foo"},
			nil,
		},
		{
			"only revert different path",
			[]Action{
				{write, "a/foo"},
				{write, "b/bar"},
				{revertTree, "a/"},
			},
			[]string{"b/bar"},
			[]string{"a/bar", "a/foo"},
			nil,
		},
		{
			"revert on Object",
			[]Action{
				{write, "a/foo"},
				{write, "a/bar"},
				{revertObj, "a/foo"},
			},
			[]string{"a/bar"},
			[]string{"a/foo"},
			nil,
		},
		{
			"revert non existing object",
			[]Action{
				{write, "a/foo"},
				{revertObj, "a/bar"},
			},
			nil,
			nil,
			db.ErrNotFound,
		},
		{
			"revert non existing tree",
			[]Action{
				{write, "a/foo"},
				{revertTree, "b/"},
			},
			nil,
			nil,
			db.ErrNotFound,
		},
	}

	for _, tc := range testData {
		t.Run(tc.Name, func(t *testing.T) {
			kv, repo, closer := testutil.GetIndexStoreWithRepo(t, 1)
			defer closer()

			kvIndex := index.NewKVIndex(kv)
			var err error
			for _, action := range tc.Actions {
				err = runCommand(kvIndex, repo, action.command, action.path)
				if err != nil {
					if xerrors.Is(err, tc.ExpectedError) {
						return
					}
					t.Fatal(err)
				}
			}
			if tc.ExpectedError != nil {
				t.Fatalf("expected to get error but did not get any")
			}
			for _, entryPath := range tc.ExpectExisting {
				_, err := kvIndex.ReadEntryObject(repo.RepoId, repo.DefaultBranch, entryPath)
				if err != nil {
					if xerrors.Is(err, db.ErrNotFound) {
						t.Fatalf("files added before commit should be available after revert")
					} else {
						t.Fatal(err)
					}
				}
			}
			for _, entryPath := range tc.ExpectMissing {
				_, err := kvIndex.ReadEntryObject(repo.RepoId, repo.DefaultBranch, entryPath)
				if !xerrors.Is(err, db.ErrNotFound) {
					t.Fatalf("files added after commit should be removed after revert")
				}
			}
		})
	}
}

func TestKVIndex_DeleteObject(t *testing.T) {
	type Action struct {
		command Command
		path    string
	}
	testData := []struct {
		Name           string
		Actions        []Action
		ExpectExisting []string
		ExpectMissing  []string
		ExpectedError  error
	}{
		{
			"add and delete",
			[]Action{
				{write, "a/foo"},
				{deleteEntry, "a/foo"},
			},

			nil,
			[]string{"a/foo"},
			nil,
		},
		{
			"delete non existing",
			[]Action{
				{write, "a/bar"},
				{deleteEntry, "a/foo"},
			},

			[]string{"a/bar"},
			[]string{"a/foo"},
			db.ErrNotFound,
		},
		{
			"rewrite deleted",
			[]Action{
				{write, "a/foo"},
				{deleteEntry, "a/foo"},
				{write, "a/foo"},
			},

			[]string{"a/foo"},
			nil,
			nil,
		},
		{
			"included",
			[]Action{
				{write, "a/foo/bar"},
				{write, "a/foo"},
				{write, "a/foo/bar/one"},
				{deleteEntry, "a/foo"},
			},

			[]string{"a/foo/bar", "a/foo/bar/one"},
			[]string{"a/foo"},
			nil,
		},
	}

	for _, tc := range testData {
		t.Run(tc.Name, func(t *testing.T) {
			kv, repo, closer := testutil.GetIndexStoreWithRepo(t, 1)
			defer closer()

			kvIndex := index.NewKVIndex(kv)
			var err error
			for _, action := range tc.Actions {
				err = runCommand(kvIndex, repo, action.command, action.path)
				if err != nil {
					if xerrors.Is(err, tc.ExpectedError) {
						return
					}
					t.Fatal(err)
				}
			}
			if tc.ExpectedError != nil {
				t.Fatalf("expected to get error but did not get any")
			}
			for _, entryPath := range tc.ExpectExisting {
				_, err := kvIndex.ReadEntryObject(repo.RepoId, repo.DefaultBranch, entryPath)
				if err != nil {
					if xerrors.Is(err, db.ErrNotFound) {
						t.Fatalf("files added before commit should be available after revert")
					} else {
						t.Fatal(err)
					}
				}
			}
			for _, entryPath := range tc.ExpectMissing {
				_, err := kvIndex.ReadEntryObject(repo.RepoId, repo.DefaultBranch, entryPath)
				if !xerrors.Is(err, db.ErrNotFound) {
					t.Fatalf("files added after commit should be removed after revert")
				}
			}
		})
	}
}

func TestSizeConsistency(t *testing.T) {
	type Object struct {
		name string
		path string
		size int64
	}
	type Tree struct {
		name       string
		wantedSize int64
	}
	testData := []struct {
		name           string
		objectList     []Object
		wantedTrees    []Tree
		wantedRootSize int64
	}{
		{
			name: "simple case",
			objectList: []Object{
				{"file1", "a/", 100},
				{"file2", "a/", 100},
				{"file3", "a/", 100},
				{"file4", "a/", 100},
				{"file5", "a/", 100},
			},
			wantedTrees: []Tree{
				{"a/", 500},
			},
			wantedRootSize: 500,
		},
		{
			name: "two separate trees",
			objectList: []Object{
				{"file1", "a/", 100},
				{"file2", "a/", 100},
				{"file3", "b/", 100},
				{"file4", "b/", 100},
				{"file5", "b/", 100},
			},
			wantedTrees: []Tree{
				{"a/", 200},
				{"b/", 300},
			},
			wantedRootSize: 500,
		},
		{
			name: "two sub trees",
			objectList: []Object{
				{"file1", "parent/a/", 100},
				{"file2", "parent/a/", 100},
				{"file3", "parent/b/", 100},
				{"file4", "parent/b/", 100},
				{"file5", "parent/b/", 100},
			},
			wantedTrees: []Tree{
				{"parent/a/", 200},
				{"parent/b/", 300},
				{"parent/", 500},
			},
			wantedRootSize: 500,
		},
		{
			name: "deep sub trees",
			objectList: []Object{
				{"file1", "a/", 100},
				{"file2", "a/b/", 100},
				{"file3", "a/b/c/", 100},
				{"file4", "a/b/c/d/", 100},
				{"file5", "a/b/c/d/e/", 100},
			},
			wantedTrees: []Tree{
				{"a/", 500},
				{"a/b/", 400},
				{"a/b/c/", 300},
				{"a/b/c/d/", 200},
				{"a/b/c/d/e/", 100},
			},
			wantedRootSize: 500,
		},
	}

	for _, tc := range testData {
		t.Run(tc.name, func(t *testing.T) {
			//TODO: add tests with different partial commit ratio
			kv, repo, closer := testutil.GetIndexStoreWithRepo(t, 1)
			defer closer()

			kvIndex := index.NewKVIndex(kv)

			for _, object := range tc.objectList {
				err := kvIndex.WriteEntry(repo.GetRepoId(), repo.GetDefaultBranch(), object.path, &model.Entry{
					Name: object.name,
					Size: object.size,
					Type: model.Entry_OBJECT,
				})
				if err != nil {
					t.Fatal(err)
				}
			}

			for _, tree := range tc.wantedTrees {
				entry, err := kvIndex.ReadEntryTree(repo.GetRepoId(), repo.GetDefaultBranch(), tree.name)
				if err != nil {
					t.Fatal(err)
				}

				// test entry size
				if entry.GetSize() != tree.wantedSize {
					t.Errorf("did not get the expected size for directory %s want: %d, got: %d", tree.name, tree.wantedSize, entry.GetSize())
				}
			}

			rootObject, err := kvIndex.ReadRootObject(repo.GetRepoId(), repo.DefaultBranch)
			if err != nil {
				t.Fatal(err)
			}
			if rootObject.GetSize() != tc.wantedRootSize {
				t.Errorf("did not get the expected size for root want: %d, got: %d", tc.wantedRootSize, rootObject.GetSize())
			}
		})
	}
}

func TestTimeStampConsistency(t *testing.T) {
	type timedObject struct {
		path    string
		name    string
		seconds time.Duration
	}
	type expectedTree struct {
		path    string
		seconds time.Duration
	}
	testData := []struct {
		name           string
		timedObjects   []timedObject
		deleteObjects  []timedObject
		expectedTrees  []expectedTree
		expectedRootTS time.Duration
	}{
		{
			timedObjects:   []timedObject{{"a/", "foo", 10}, {"a/", "bar", 20}},
			expectedTrees:  []expectedTree{{"a/", 20}},
			expectedRootTS: 20,
		},
		{
			timedObjects:   []timedObject{{"a/", "wow", 5}, {"a/c/", "bar", 10}, {"a/b/", "bar", 20}},
			expectedTrees:  []expectedTree{{"a/", 20}, {"a/c/", 10}, {"a/b/", 20}},
			expectedRootTS: 20,
		},
		{
			name:           "delete file",
			timedObjects:   []timedObject{{"a/", "wow", 5}, {"a/c/", "bar", 10}, {"a/c/", "foo", 15}, {"a/b/", "bar", 20}},
			deleteObjects:  []timedObject{{"a/c/", "foo", 25}},
			expectedTrees:  []expectedTree{{"a/", 25}, {"a/c/", 25}, {"a/b/", 20}},
			expectedRootTS: 25,
		},
	}

	for _, tc := range testData {
		t.Run(tc.name, func(t *testing.T) {
			kv, repo, closer := testutil.GetIndexStoreWithRepo(t, 1)
			defer closer()
			now := time.Now()
			currentTime := now
			mockTime := func() int64 {
				return currentTime.Unix()
			}
			kvIndex := index.NewKVIndex(kv, index.WithTimeGenerator(mockTime))
			for _, obj := range tc.timedObjects {
				ts := now.Add(obj.seconds * time.Second).Unix()
				err := kvIndex.WriteEntry(repo.GetRepoId(), repo.GetDefaultBranch(), obj.path, &model.Entry{
					Name:      obj.name,
					Address:   "12345678",
					Type:      model.Entry_OBJECT,
					Timestamp: ts,
				})
				if err != nil {
					t.Fatal(err)
				}
			}
			for _, obj := range tc.deleteObjects {
				currentTime = now.Add(obj.seconds * time.Second)
				err := kvIndex.DeleteObject(repo.GetRepoId(), repo.DefaultBranch, obj.path+obj.name)
				if err != nil {
					t.Fatal(err)
				}
			}
			for _, tree := range tc.expectedTrees {
				entry, err := kvIndex.ReadEntryTree(repo.GetRepoId(), repo.GetDefaultBranch(), tree.path)
				if err != nil {
					t.Fatal(err)
				}
				expectedTS := now.Add(tree.seconds * time.Second).Unix()
				if entry.GetTimestamp() != expectedTS {
					t.Errorf("unexpected times stamp for tree, expected: %v , got: %v", expectedTS, entry.GetTimestamp())
				}
			}

			rootObject, err := kvIndex.ReadRootObject(repo.GetRepoId(), repo.DefaultBranch)
			if err != nil {
				t.Fatal(err)
			}
			expectedTS := now.Add(tc.expectedRootTS * time.Second).Unix()
			if rootObject.GetTimestamp() != expectedTS {
				t.Errorf("unexpected times stamp for tree, expected: %v , got: %v", expectedTS, rootObject.GetTimestamp())
			}

		})
	}
}
func runCommand(kvIndex *index.KVIndex, repo *model.Repo, command Command, actionPath string) error {
	var err error
	switch command {
	case write:
		err = kvIndex.WriteEntry(repo.RepoId, repo.DefaultBranch, actionPath, &model.Entry{
			Name:      path.New(actionPath).Basename(),
			Address:   "123456789",
			Timestamp: time.Now().Unix(),
			Type:      model.Entry_OBJECT,
		})

	case commit:
		_, err = kvIndex.Commit(repo.RepoId, repo.DefaultBranch, "test msg", "committer", nil)

	case revertTree:
		err = kvIndex.RevertPath(repo.RepoId, repo.DefaultBranch, actionPath)

	case revertObj:
		err = kvIndex.RevertObject(repo.RepoId, repo.DefaultBranch, actionPath)

	case deleteEntry:
		err = kvIndex.DeleteObject(repo.RepoId, repo.DefaultBranch, actionPath)

	default:
		err = xerrors.Errorf("unknown command")
	}
	return err
}