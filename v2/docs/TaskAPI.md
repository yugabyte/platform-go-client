# \TaskAPI

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PageListTasks**](TaskAPI.md#PageListTasks) | **Post** /customers/{cUUID}/tasks/page | List customer tasks (paged)
[**RollbackTask**](TaskAPI.md#RollbackTask) | **Post** /customers/{cUUID}/tasks/{tUUID}/rollback | Rollback a failed task



## PageListTasks

> TaskPagedResp PageListTasks(ctx, cUUID).TaskPagedQuerySpec(taskPagedQuerySpec).Execute()

List customer tasks (paged)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	taskPagedQuerySpec := *openapiclient.NewTaskPagedQuerySpec() // TaskPagedQuerySpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.PageListTasks(context.Background(), cUUID).TaskPagedQuerySpec(taskPagedQuerySpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.PageListTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageListTasks`: TaskPagedResp
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.PageListTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageListTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskPagedQuerySpec** | [**TaskPagedQuerySpec**](TaskPagedQuerySpec.md) |  | 

### Return type

[**TaskPagedResp**](TaskPagedResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbackTask

> YBATask RollbackTask(ctx, cUUID, tUUID).TaskRollbackSpec(taskRollbackSpec).Execute()

Rollback a failed task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	tUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Task UUID
	taskRollbackSpec := *openapiclient.NewTaskRollbackSpec() // TaskRollbackSpec |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.RollbackTask(context.Background(), cUUID, tUUID).TaskRollbackSpec(taskRollbackSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.RollbackTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RollbackTask`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.RollbackTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**tUUID** | **string** | Task UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **taskRollbackSpec** | [**TaskRollbackSpec**](TaskRollbackSpec.md) |  | 

### Return type

[**YBATask**](YBATask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

