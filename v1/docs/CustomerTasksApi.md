# \CustomerTasksApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AbortTask**](CustomerTasksApi.md#AbortTask) | **Post** /api/v1/customers/{cUUID}/tasks/{tUUID}/abort | Abort a task
[**FailedSubtasks**](CustomerTasksApi.md#FailedSubtasks) | **Get** /api/v1/customers/{cUUID}/tasks/{tUUID}/failed | Fetch failed subtasks - deprecated
[**ListFailedSubtasks**](CustomerTasksApi.md#ListFailedSubtasks) | **Get** /api/v1/customers/{cUUID}/tasks/{tUUID}/failed_subtasks | Get a list of task&#39;s failed subtasks
[**RetryTask**](CustomerTasksApi.md#RetryTask) | **Post** /api/v1/customers/{cUUID}/tasks/{tUUID}/retry | Retry a Universe or Provider task
[**TaskStatus**](CustomerTasksApi.md#TaskStatus) | **Get** /api/v1/customers/{cUUID}/tasks/{tUUID} | Get a task&#39;s status
[**TasksList**](CustomerTasksApi.md#TasksList) | **Get** /api/v1/customers/{cUUID}/tasks_list | List task



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
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    tUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerTasksApi.AbortTask(context.Background(), cUUID, tUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerTasksApi.AbortTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AbortTask`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `CustomerTasksApi.AbortTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**tUUID** | [**string**](.md) |  | 

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
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    tUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerTasksApi.FailedSubtasks(context.Background(), cUUID, tUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerTasksApi.FailedSubtasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FailedSubtasks`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CustomerTasksApi.FailedSubtasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**tUUID** | [**string**](.md) |  | 

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
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    tUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerTasksApi.ListFailedSubtasks(context.Background(), cUUID, tUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerTasksApi.ListFailedSubtasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFailedSubtasks`: FailedSubtasks
    fmt.Fprintf(os.Stdout, "Response from `CustomerTasksApi.ListFailedSubtasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**tUUID** | [**string**](.md) |  | 

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
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    tUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerTasksApi.RetryTask(context.Background(), cUUID, tUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerTasksApi.RetryTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetryTask`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `CustomerTasksApi.RetryTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**tUUID** | [**string**](.md) |  | 

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
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    tUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerTasksApi.TaskStatus(context.Background(), cUUID, tUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerTasksApi.TaskStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CustomerTasksApi.TaskStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**tUUID** | [**string**](.md) |  | 

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
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    uUUID := TODO // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerTasksApi.TasksList(context.Background(), cUUID).UUUID(uUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerTasksApi.TasksList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksList`: []CustomerTaskData
    fmt.Fprintf(os.Stdout, "Response from `CustomerTasksApi.TasksList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uUUID** | [**string**](string.md) |  | 

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

