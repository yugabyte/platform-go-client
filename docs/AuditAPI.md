# \AuditAPI

All URIs are relative to *http://localhost*

| Method                                             | HTTP request                                                | Description                         |
| -------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------- |
| [**GetTaskAudit**](AuditAPI.md#GetTaskAudit)       | **Get** /api/v1/customers/{cUUID}/tasks/{tUUID}/audit_info  | Get audit info for a task           |
| [**GetUserFromTask**](AuditAPI.md#GetUserFromTask) | **Get** /api/v1/customers/{cUUID}/tasks/{tUUID}/audit_user  | Get the user associated with a task |
| [**ListOfAudit**](AuditAPI.md#ListOfAudit)         | **Get** /api/v1/customers/{cUUID}/users/{uUUID}/audit_trail | List a user&#39;s audit entries     |



## GetTaskAudit

> Audit GetTaskAudit(ctx, cUUID, tUUID).Execute()

Get audit info for a task

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
	tUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditAPI.GetTaskAudit(context.Background(), cUUID, tUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditAPI.GetTaskAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskAudit`: Audit
	fmt.Fprintf(os.Stdout, "Response from `AuditAPI.GetTaskAudit`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **tUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskAuditRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

[**Audit**](Audit.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserFromTask

> Audit GetUserFromTask(ctx, cUUID, tUUID).Execute()

Get the user associated with a task

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
	tUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditAPI.GetUserFromTask(context.Background(), cUUID, tUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditAPI.GetUserFromTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserFromTask`: Audit
	fmt.Fprintf(os.Stdout, "Response from `AuditAPI.GetUserFromTask`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **tUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserFromTaskRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

[**Audit**](Audit.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOfAudit

> []Audit ListOfAudit(ctx, cUUID, uUUID).Execute()

List a user's audit entries

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
	uUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditAPI.ListOfAudit(context.Background(), cUUID, uUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditAPI.ListOfAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOfAudit`: []Audit
	fmt.Fprintf(os.Stdout, "Response from `AuditAPI.ListOfAudit`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **uUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiListOfAuditRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

[**[]Audit**](Audit.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

