# \AuthenticationApi

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteGroupMappings**](AuthenticationApi.md#DeleteGroupMappings) | **Delete** /customers/{cUUID}/auth/group-mappings/{gUUID} | Delete Group Mappings
[**ListMappings**](AuthenticationApi.md#ListMappings) | **Get** /customers/{cUUID}/auth/group-mappings | List Group Mappings
[**UpdateGroupMappings**](AuthenticationApi.md#UpdateGroupMappings) | **Put** /customers/{cUUID}/auth/group-mappings | Create Group Mappings



## DeleteGroupMappings

> DeleteGroupMappings(ctx, cUUID, gUUID).Execute()

Delete Group Mappings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | Customer UUID
    gUUID := TODO // string | Group UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationApi.DeleteGroupMappings(context.Background(), cUUID, gUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.DeleteGroupMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**gUUID** | [**string**](.md) | Group UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMappings

> []AuthGroupToRolesMapping ListMappings(ctx, cUUID).Execute()

List Group Mappings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | Customer UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationApi.ListMappings(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.ListMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMappings`: []AuthGroupToRolesMapping
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.ListMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AuthGroupToRolesMapping**](AuthGroupToRolesMapping.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupMappings

> UpdateGroupMappings(ctx, cUUID).AuthGroupToRolesMapping(authGroupToRolesMapping).Execute()

Create Group Mappings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | Customer UUID
    authGroupToRolesMapping := []openapiclient.AuthGroupToRolesMapping{*openapiclient.NewAuthGroupToRolesMapping("GroupIdentifier_example", "Type_example", []openapiclient.RoleResourceDefinition{*openapiclient.NewRoleResourceDefinition("RoleUuid_example")})} // []AuthGroupToRolesMapping | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationApi.UpdateGroupMappings(context.Background(), cUUID).AuthGroupToRolesMapping(authGroupToRolesMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.UpdateGroupMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authGroupToRolesMapping** | [**[]AuthGroupToRolesMapping**](AuthGroupToRolesMapping.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

