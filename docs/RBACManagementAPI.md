# \RBACManagementAPI

All URIs are relative to *http://localhost*

| Method                                                      | HTTP request                                                    | Description                         |
| ----------------------------------------------------------- | --------------------------------------------------------------- | ----------------------------------- |
| [**CreateRole**](RBACManagementAPI.md#CreateRole)           | **Post** /api/v1/customers/{cUUID}/rbac/role                    | Create a custom role                |
| [**DeleteRole**](RBACManagementAPI.md#DeleteRole)           | **Delete** /api/v1/customers/{cUUID}/rbac/role/{rUUID}          | Delete a custom role                |
| [**EditRole**](RBACManagementAPI.md#EditRole)               | **Put** /api/v1/customers/{cUUID}/rbac/role/{rUUID}             | Edit a custom role                  |
| [**GetRole**](RBACManagementAPI.md#GetRole)                 | **Get** /api/v1/customers/{cUUID}/rbac/role/{rUUID}             | Get a role&#39;s information        |
| [**GetRoleBindings**](RBACManagementAPI.md#GetRoleBindings) | **Get** /api/v1/customers/{cUUID}/rbac/role_binding             | Get all the role bindings available |
| [**ListPermissions**](RBACManagementAPI.md#ListPermissions) | **Get** /api/v1/customers/{cUUID}/rbac/permissions              | List all the permissions available  |
| [**ListRoles**](RBACManagementAPI.md#ListRoles)             | **Get** /api/v1/customers/{cUUID}/rbac/role                     | List all the roles available        |
| [**SetRoleBinding**](RBACManagementAPI.md#SetRoleBinding)   | **Post** /api/v1/customers/{cUUID}/rbac/role_binding/{userUUID} | Set the role bindings of a user     |



## CreateRole

> Role CreateRole(ctx, cUUID).RoleFormData(roleFormData).Request(request).Execute()

Create a custom role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleFormData := *openapiclient.NewRoleFormData("Description_example", "Name_example", []openapiclient.Permission{*openapiclient.NewPermission()}) // RoleFormData | create role form data
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RBACManagementAPI.CreateRole(context.Background(), cUUID).RoleFormData(roleFormData).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RBACManagementAPI.CreateRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRole`: Role
	fmt.Fprintf(os.Stdout, "Response from `RBACManagementAPI.CreateRole`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **roleFormData** | [**RoleFormData**](RoleFormData.md) | create role form data | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**Role**](Role.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	rUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RBACManagementAPI.DeleteRole(context.Background(), cUUID, rUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RBACManagementAPI.DeleteRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRole`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `RBACManagementAPI.DeleteRole`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **rUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


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

> Role EditRole(ctx, cUUID, rUUID).RoleFormData(roleFormData).Request(request).Execute()

Edit a custom role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	rUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleFormData := *openapiclient.NewRoleFormData("Description_example", "Name_example", []openapiclient.Permission{*openapiclient.NewPermission()}) // RoleFormData | edit role form data
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RBACManagementAPI.EditRole(context.Background(), cUUID, rUUID).RoleFormData(roleFormData).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RBACManagementAPI.EditRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditRole`: Role
	fmt.Fprintf(os.Stdout, "Response from `RBACManagementAPI.EditRole`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **rUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiEditRoleRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **roleFormData** | [**RoleFormData**](RoleFormData.md) | edit role form data | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**Role**](Role.md)

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
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	rUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RBACManagementAPI.GetRole(context.Background(), cUUID, rUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RBACManagementAPI.GetRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRole`: Role
	fmt.Fprintf(os.Stdout, "Response from `RBACManagementAPI.GetRole`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **rUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



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
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	userUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional user UUID to filter role binding map (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RBACManagementAPI.GetRoleBindings(context.Background(), cUUID).UserUUID(userUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RBACManagementAPI.GetRoleBindings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoleBindings`: map[string]RoleBinding
	fmt.Fprintf(os.Stdout, "Response from `RBACManagementAPI.GetRoleBindings`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleBindingsRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **userUUID** | **string** | Optional user UUID to filter role binding map | 

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
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	resourceType := "resourceType_example" // string | Optional resource type to filter permission list (optional) (default to "null")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RBACManagementAPI.ListPermissions(context.Background(), cUUID).ResourceType(resourceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RBACManagementAPI.ListPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPermissions`: []PermissionInfo
	fmt.Fprintf(os.Stdout, "Response from `RBACManagementAPI.ListPermissions`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiListPermissionsRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

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
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleType := "roleType_example" // string | Optional role type to filter roles list (optional) (default to "null")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RBACManagementAPI.ListRoles(context.Background(), cUUID).RoleType(roleType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RBACManagementAPI.ListRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRoles`: []Role
	fmt.Fprintf(os.Stdout, "Response from `RBACManagementAPI.ListRoles`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiListRolesRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

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


## SetRoleBinding

> []RoleBinding SetRoleBinding(ctx, cUUID, userUUID).RoleBindingFormData(roleBindingFormData).Request(request).Execute()

Set the role bindings of a user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	userUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roleBindingFormData := *openapiclient.NewRoleBindingFormData([]openapiclient.RoleResourceDefinition{*openapiclient.NewRoleResourceDefinition("RoleUUID_example")}) // RoleBindingFormData | set role bindings form data
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RBACManagementAPI.SetRoleBinding(context.Background(), cUUID, userUUID).RoleBindingFormData(roleBindingFormData).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RBACManagementAPI.SetRoleBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetRoleBinding`: []RoleBinding
	fmt.Fprintf(os.Stdout, "Response from `RBACManagementAPI.SetRoleBinding`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**    | **string**          |                                                                             |
| **userUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiSetRoleBindingRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **roleBindingFormData** | [**RoleBindingFormData**](RoleBindingFormData.md) | set role bindings form data | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**[]RoleBinding**](RoleBinding.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

