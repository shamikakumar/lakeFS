"""
    lakeFS API

    lakeFS HTTP API  # noqa: E501

    The version of the OpenAPI document: 0.1.0
    Contact: services@treeverse.io
    Generated by: https://openapi-generator.tech
"""


import re  # noqa: F401
import sys  # noqa: F401

from lakefs_client.api_client import ApiClient, Endpoint as _Endpoint
from lakefs_client.model_utils import (  # noqa: F401
    check_allowed_values,
    check_validations,
    date,
    datetime,
    file_type,
    none_type,
    validate_and_convert_types
)
from lakefs_client.model.error import Error
from lakefs_client.model.ref import Ref
from lakefs_client.model.ref_list import RefList
from lakefs_client.model.tag_creation import TagCreation


class TagsApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient()
        self.api_client = api_client

        def __create_tag(
            self,
            repository,
            tag_creation,
            **kwargs
        ):
            """create tag  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.create_tag(repository, tag_creation, async_req=True)
            >>> result = thread.get()

            Args:
                repository (str):
                tag_creation (TagCreation):

            Keyword Args:
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (float/tuple): timeout setting for this request. If one
                    number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                Ref
                    If the method is called asynchronously, returns the request
                    thread.
            """
            kwargs['async_req'] = kwargs.get(
                'async_req', False
            )
            kwargs['_return_http_data_only'] = kwargs.get(
                '_return_http_data_only', True
            )
            kwargs['_preload_content'] = kwargs.get(
                '_preload_content', True
            )
            kwargs['_request_timeout'] = kwargs.get(
                '_request_timeout', None
            )
            kwargs['_check_input_type'] = kwargs.get(
                '_check_input_type', True
            )
            kwargs['_check_return_type'] = kwargs.get(
                '_check_return_type', True
            )
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['repository'] = \
                repository
            kwargs['tag_creation'] = \
                tag_creation
            return self.call_with_http_info(**kwargs)

        self.create_tag = _Endpoint(
            settings={
                'response_type': (Ref,),
                'auth': [
                    'basic_auth',
                    'cookie_auth',
                    'jwt_token'
                ],
                'endpoint_path': '/repositories/{repository}/tags',
                'operation_id': 'create_tag',
                'http_method': 'POST',
                'servers': None,
            },
            params_map={
                'all': [
                    'repository',
                    'tag_creation',
                ],
                'required': [
                    'repository',
                    'tag_creation',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'repository':
                        (str,),
                    'tag_creation':
                        (TagCreation,),
                },
                'attribute_map': {
                    'repository': 'repository',
                },
                'location_map': {
                    'repository': 'path',
                    'tag_creation': 'body',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [
                    'application/json'
                ]
            },
            api_client=api_client,
            callable=__create_tag
        )

        def __delete_tag(
            self,
            repository,
            tag,
            **kwargs
        ):
            """delete tag  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.delete_tag(repository, tag, async_req=True)
            >>> result = thread.get()

            Args:
                repository (str):
                tag (str):

            Keyword Args:
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (float/tuple): timeout setting for this request. If one
                    number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                None
                    If the method is called asynchronously, returns the request
                    thread.
            """
            kwargs['async_req'] = kwargs.get(
                'async_req', False
            )
            kwargs['_return_http_data_only'] = kwargs.get(
                '_return_http_data_only', True
            )
            kwargs['_preload_content'] = kwargs.get(
                '_preload_content', True
            )
            kwargs['_request_timeout'] = kwargs.get(
                '_request_timeout', None
            )
            kwargs['_check_input_type'] = kwargs.get(
                '_check_input_type', True
            )
            kwargs['_check_return_type'] = kwargs.get(
                '_check_return_type', True
            )
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['repository'] = \
                repository
            kwargs['tag'] = \
                tag
            return self.call_with_http_info(**kwargs)

        self.delete_tag = _Endpoint(
            settings={
                'response_type': None,
                'auth': [
                    'basic_auth',
                    'cookie_auth',
                    'jwt_token'
                ],
                'endpoint_path': '/repositories/{repository}/tags/{tag}',
                'operation_id': 'delete_tag',
                'http_method': 'DELETE',
                'servers': None,
            },
            params_map={
                'all': [
                    'repository',
                    'tag',
                ],
                'required': [
                    'repository',
                    'tag',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'repository':
                        (str,),
                    'tag':
                        (str,),
                },
                'attribute_map': {
                    'repository': 'repository',
                    'tag': 'tag',
                },
                'location_map': {
                    'repository': 'path',
                    'tag': 'path',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client,
            callable=__delete_tag
        )

        def __get_tag(
            self,
            repository,
            tag,
            **kwargs
        ):
            """get tag  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.get_tag(repository, tag, async_req=True)
            >>> result = thread.get()

            Args:
                repository (str):
                tag (str):

            Keyword Args:
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (float/tuple): timeout setting for this request. If one
                    number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                Ref
                    If the method is called asynchronously, returns the request
                    thread.
            """
            kwargs['async_req'] = kwargs.get(
                'async_req', False
            )
            kwargs['_return_http_data_only'] = kwargs.get(
                '_return_http_data_only', True
            )
            kwargs['_preload_content'] = kwargs.get(
                '_preload_content', True
            )
            kwargs['_request_timeout'] = kwargs.get(
                '_request_timeout', None
            )
            kwargs['_check_input_type'] = kwargs.get(
                '_check_input_type', True
            )
            kwargs['_check_return_type'] = kwargs.get(
                '_check_return_type', True
            )
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['repository'] = \
                repository
            kwargs['tag'] = \
                tag
            return self.call_with_http_info(**kwargs)

        self.get_tag = _Endpoint(
            settings={
                'response_type': (Ref,),
                'auth': [
                    'basic_auth',
                    'cookie_auth',
                    'jwt_token'
                ],
                'endpoint_path': '/repositories/{repository}/tags/{tag}',
                'operation_id': 'get_tag',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'repository',
                    'tag',
                ],
                'required': [
                    'repository',
                    'tag',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                ]
            },
            root_map={
                'validations': {
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'repository':
                        (str,),
                    'tag':
                        (str,),
                },
                'attribute_map': {
                    'repository': 'repository',
                    'tag': 'tag',
                },
                'location_map': {
                    'repository': 'path',
                    'tag': 'path',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client,
            callable=__get_tag
        )

        def __list_tags(
            self,
            repository,
            **kwargs
        ):
            """list tags  # noqa: E501

            This method makes a synchronous HTTP request by default. To make an
            asynchronous HTTP request, please pass async_req=True

            >>> thread = api.list_tags(repository, async_req=True)
            >>> result = thread.get()

            Args:
                repository (str):

            Keyword Args:
                after (str): return items after this value. [optional]
                amount (int): how many items to return. [optional] if omitted the server will use the default value of 100
                _return_http_data_only (bool): response data without head status
                    code and headers. Default is True.
                _preload_content (bool): if False, the urllib3.HTTPResponse object
                    will be returned without reading/decoding response data.
                    Default is True.
                _request_timeout (float/tuple): timeout setting for this request. If one
                    number provided, it will be total request timeout. It can also
                    be a pair (tuple) of (connection, read) timeouts.
                    Default is None.
                _check_input_type (bool): specifies if type checking
                    should be done one the data sent to the server.
                    Default is True.
                _check_return_type (bool): specifies if type checking
                    should be done one the data received from the server.
                    Default is True.
                _host_index (int/None): specifies the index of the server
                    that we want to use.
                    Default is read from the configuration.
                async_req (bool): execute request asynchronously

            Returns:
                RefList
                    If the method is called asynchronously, returns the request
                    thread.
            """
            kwargs['async_req'] = kwargs.get(
                'async_req', False
            )
            kwargs['_return_http_data_only'] = kwargs.get(
                '_return_http_data_only', True
            )
            kwargs['_preload_content'] = kwargs.get(
                '_preload_content', True
            )
            kwargs['_request_timeout'] = kwargs.get(
                '_request_timeout', None
            )
            kwargs['_check_input_type'] = kwargs.get(
                '_check_input_type', True
            )
            kwargs['_check_return_type'] = kwargs.get(
                '_check_return_type', True
            )
            kwargs['_host_index'] = kwargs.get('_host_index')
            kwargs['repository'] = \
                repository
            return self.call_with_http_info(**kwargs)

        self.list_tags = _Endpoint(
            settings={
                'response_type': (RefList,),
                'auth': [
                    'basic_auth',
                    'cookie_auth',
                    'jwt_token'
                ],
                'endpoint_path': '/repositories/{repository}/tags',
                'operation_id': 'list_tags',
                'http_method': 'GET',
                'servers': None,
            },
            params_map={
                'all': [
                    'repository',
                    'after',
                    'amount',
                ],
                'required': [
                    'repository',
                ],
                'nullable': [
                ],
                'enum': [
                ],
                'validation': [
                    'amount',
                ]
            },
            root_map={
                'validations': {
                    ('amount',): {

                        'inclusive_maximum': 1000,
                        'inclusive_minimum': -1,
                    },
                },
                'allowed_values': {
                },
                'openapi_types': {
                    'repository':
                        (str,),
                    'after':
                        (str,),
                    'amount':
                        (int,),
                },
                'attribute_map': {
                    'repository': 'repository',
                    'after': 'after',
                    'amount': 'amount',
                },
                'location_map': {
                    'repository': 'path',
                    'after': 'query',
                    'amount': 'query',
                },
                'collection_format_map': {
                }
            },
            headers_map={
                'accept': [
                    'application/json'
                ],
                'content_type': [],
            },
            api_client=api_client,
            callable=__list_tags
        )
