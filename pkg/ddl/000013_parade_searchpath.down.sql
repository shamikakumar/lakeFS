-- Do nothing on DOWN.  "UP" exists only to fix a bug in a function defined in
-- 000011_parade.up.sql, which looks up `gen_random_uuid' with no schema.  Now
-- that version is gone, so always uses the schema.
