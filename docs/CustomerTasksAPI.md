# \CustomerTasksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AbortTask**](CustomerTasksAPI.md#AbortTask) | **Post** /api/v1/customers/{cUUID}/tasks/{tUUID}/abort | Abort a task
[**FailedSubtasks**](CustomerTasksAPI.md#FailedSubtasks) | **Get** /api/v1/customers/{cUUID}/tasks/{tUUID}/failed | Fetch failed subtasks - deprecated
[**ListFailedSubtasks**](CustomerTasksAPI.md#ListFailedSubtasks) | **Get** /api/v1/customers/{cUUID}/tasks/{tUUID}/failed_subtasks | Get a list of task&#39;s failed subtasks
[**ListTasksV2**](CustomerTasksAPI.md#ListTasksV2) | **Post** /api/v1/customers/{cUUID}/tasks_list/page | List Tasks (paginated)
[**RetryTask**](CustomerTasksAPI.md#RetryTask) | **Post** /api/v1/customers/{cUUID}/tasks/{tUUID}/retry | Retry a Universe or Provider task
[**RollbackTask**](CustomerTasksAPI.md#RollbackTask) | **Post** /api/v1/customers/{cUUID}/tasks/{tUUID}/rollback | Rollback a Universe or Provider task
[**TaskStatus**](CustomerTasksAPI.md#TaskStatus) | **Get** /api/v1/customers/{cUUID}/tasks/{tUUID} | Get a task&#39;s status
[**TasksList**](CustomerTasksAPI.md#TasksList) | **Get** /api/v1/customers/{cUUID}/tasks_list | List task



## AbortTask

> YBPSuccess AbortTask(ctx, cUUID, tUUID).Request(request).Execute()

Abort a task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerTasksAPI.AbortTask(context.Background(), cUUID, tUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerTasksAPI.AbortTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AbortTask`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `CustomerTasksAPI.AbortTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**tUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAbortTaskRequest struct via the builder pattern


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


## FailedSubtasks

> map[string]map[string]interface{} FailedSubtasks(ctx, cUUID, tUUID).Execute()

Fetch failed subtasks - deprecated



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerTasksAPI.FailedSubtasks(context.Background(), cUUID, tUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerTasksAPI.FailedSubtasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FailedSubtasks`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomerTasksAPI.FailedSubtasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**tUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFailedSubtasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFailedSubtasks

> FailedSubtasks ListFailedSubtasks(ctx, cUUID, tUUID).Execute()

Get a list of task's failed subtasks

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerTasksAPI.ListFailedSubtasks(context.Background(), cUUID, tUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerTasksAPI.ListFailedSubtasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFailedSubtasks`: FailedSubtasks
	fmt.Fprintf(os.Stdout, "Response from `CustomerTasksAPI.ListFailedSubtasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**tUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFailedSubtasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FailedSubtasks**](FailedSubtasks.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTasksV2

> TaskPagedApiResponse ListTasksV2(ctx, cUUID).PageTasksRequest(pageTasksRequest).Request(request).Execute()

List Tasks (paginated)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pageTasksRequest := *openapiclient.NewTaskPagedApiQuery("Direction_example", *openapiclient.NewTaskApiFilter([]string{"Status_example"}, []string{"TargetList_example"}, []string{"TargetUUIDList_example"}, []string{"TypeList_example"}, []string{"TypeNameList_example"}), int32(123), false, int32(123), "SortBy_example") // TaskPagedApiQuery | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerTasksAPI.ListTasksV2(context.Background(), cUUID).PageTasksRequest(pageTasksRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerTasksAPI.ListTasksV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTasksV2`: TaskPagedApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomerTasksAPI.ListTasksV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTasksV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageTasksRequest** | [**TaskPagedApiQuery**](TaskPagedApiQuery.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**TaskPagedApiResponse**](TaskPagedApiResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryTask

> YBPTask RetryTask(ctx, cUUID, tUUID).Request(request).Execute()

Retry a Universe or Provider task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerTasksAPI.RetryTask(context.Background(), cUUID, tUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerTasksAPI.RetryTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryTask`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `CustomerTasksAPI.RetryTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**tUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbackTask

> YBPTask RollbackTask(ctx, cUUID, tUUID).Request(request).Execute()

Rollback a Universe or Provider task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerTasksAPI.RollbackTask(context.Background(), cUUID, tUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerTasksAPI.RollbackTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RollbackTask`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `CustomerTasksAPI.RollbackTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**tUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskStatus

> map[string]interface{} TaskStatus(ctx, cUUID, tUUID).Execute()

Get a task's status

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerTasksAPI.TaskStatus(context.Background(), cUUID, tUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerTasksAPI.TaskStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaskStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomerTasksAPI.TaskStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**tUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksList

> []CustomerTaskData TasksList(ctx, cUUID).UUUID(uUUID).Execute()

List task

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerTasksAPI.TasksList(context.Background(), cUUID).UUUID(uUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerTasksAPI.TasksList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksList`: []CustomerTaskData
	fmt.Fprintf(os.Stdout, "Response from `CustomerTasksAPI.TasksList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uUUID** | **string** |  | 

### Return type

[**[]CustomerTaskData**](CustomerTaskData.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

