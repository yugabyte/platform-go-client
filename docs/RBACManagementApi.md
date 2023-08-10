# \RBACManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRole**](RBACManagementApi.md#CreateRole) | **Post** /api/v1/customers/{cUUID}/rbac/role | Create a custom role
[**DeleteRole**](RBACManagementApi.md#DeleteRole) | **Delete** /api/v1/customers/{cUUID}/rbac/role/{rUUID} | Delete a custom role
[**EditRole**](RBACManagementApi.md#EditRole) | **Put** /api/v1/customers/{cUUID}/rbac/role/{rUUID} | Edit a custom role
[**EditRoleBinding**](RBACManagementApi.md#EditRoleBinding) | **Post** /api/v1/customers/{cUUID}/rbac/role_binding/{userUUID} | Edit the role bindings of a user
[**GetRole**](RBACManagementApi.md#GetRole) | **Get** /api/v1/customers/{cUUID}/rbac/role/{rUUID} | Get a role&#39;s information
[**GetRoleBindings**](RBACManagementApi.md#GetRoleBindings) | **Get** /api/v1/customers/{cUUID}/rbac/role_binding | Get all the role bindings available
[**ListPermissions**](RBACManagementApi.md#ListPermissions) | **Get** /api/v1/customers/{cUUID}/rbac/permissions | List all the permissions available
[**ListRoles**](RBACManagementApi.md#ListRoles) | **Get** /api/v1/customers/{cUUID}/rbac/role | List all the roles available



## CreateRole

> Role CreateRole(ctx, cUUID).Request(request).Execute()

Create a custom role

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
    cUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RBACManagementApi.CreateRole(context.Background(), cUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACManagementApi.CreateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `RBACManagementApi.CreateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> YBPSuccess DeleteRole(ctx, cUUID, rUUID).Request(request).Execute()

Delete a custom role

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
    cUUID := TODO // string | 
    rUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RBACManagementApi.DeleteRole(context.Background(), cUUID, rUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACManagementApi.DeleteRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRole`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `RBACManagementApi.DeleteRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**rUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPSuccess**](YBPSuccess.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditRole

> Role EditRole(ctx, cUUID, rUUID).Request(request).Execute()

Edit a custom role

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
    cUUID := TODO // string | 
    rUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RBACManagementApi.EditRole(context.Background(), cUUID, rUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACManagementApi.EditRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `RBACManagementApi.EditRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**rUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditRoleBinding

> RoleBinding EditRoleBinding(ctx, cUUID, userUUID).RoleBindingFormData(roleBindingFormData).Request(request).Execute()

Edit the role bindings of a user

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
    cUUID := TODO // string | 
    userUUID := TODO // string | 
    roleBindingFormData := *openapiclient.NewRoleBindingFormData([]openapiclient.RoleResourceDefinition{*openapiclient.NewRoleResourceDefinition(*openapiclient.NewResourceGroup([]openapiclient.ResourceDefinition{*openapiclient.NewResourceDefinition()}), "RoleUUID_example")}) // RoleBindingFormData | set role bindings form data
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RBACManagementApi.EditRoleBinding(context.Background(), cUUID, userUUID).RoleBindingFormData(roleBindingFormData).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACManagementApi.EditRoleBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditRoleBinding`: RoleBinding
    fmt.Fprintf(os.Stdout, "Response from `RBACManagementApi.EditRoleBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**userUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditRoleBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleBindingFormData** | [**RoleBindingFormData**](RoleBindingFormData.md) | set role bindings form data | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**RoleBinding**](RoleBinding.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> Role GetRole(ctx, cUUID, rUUID).Execute()

Get a role's information

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
    cUUID := TODO // string | 
    rUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RBACManagementApi.GetRole(context.Background(), cUUID, rUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACManagementApi.GetRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `RBACManagementApi.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**rUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Role**](Role.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleBindings

> map[string]RoleBinding GetRoleBindings(ctx, cUUID).UserUUID(userUUID).Execute()

Get all the role bindings available

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
    cUUID := TODO // string | 
    userUUID := TODO // string | Optional user UUID to filter role binding map (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RBACManagementApi.GetRoleBindings(context.Background(), cUUID).UserUUID(userUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACManagementApi.GetRoleBindings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleBindings`: map[string]RoleBinding
    fmt.Fprintf(os.Stdout, "Response from `RBACManagementApi.GetRoleBindings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleBindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userUUID** | [**string**](string.md) | Optional user UUID to filter role binding map | 

### Return type

[**map[string]RoleBinding**](RoleBinding.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPermissions

> []PermissionInfo ListPermissions(ctx, cUUID).ResourceType(resourceType).Execute()

List all the permissions available

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
    cUUID := TODO // string | 
    resourceType := "resourceType_example" // string | Optional resource type to filter permission list (optional) (default to "null")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RBACManagementApi.ListPermissions(context.Background(), cUUID).ResourceType(resourceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACManagementApi.ListPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPermissions`: []PermissionInfo
    fmt.Fprintf(os.Stdout, "Response from `RBACManagementApi.ListPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceType** | **string** | Optional resource type to filter permission list | [default to &quot;null&quot;]

### Return type

[**[]PermissionInfo**](PermissionInfo.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoles

> []Role ListRoles(ctx, cUUID).RoleType(roleType).Execute()

List all the roles available

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
    cUUID := TODO // string | 
    roleType := "roleType_example" // string | Optional role type to filter roles list (optional) (default to "null")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RBACManagementApi.ListRoles(context.Background(), cUUID).RoleType(roleType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RBACManagementApi.ListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoles`: []Role
    fmt.Fprintf(os.Stdout, "Response from `RBACManagementApi.ListRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleType** | **string** | Optional role type to filter roles list | [default to &quot;null&quot;]

### Return type

[**[]Role**](Role.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

