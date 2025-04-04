# \JobSchedulerApi

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteJobSchedule**](JobSchedulerApi.md#DeleteJobSchedule) | **Delete** /customers/{cUUID}/job-schedules/{jUUID} | Delete Job Schedule
[**GetJobSchedule**](JobSchedulerApi.md#GetJobSchedule) | **Get** /customers/{cUUID}/job-schedules/{jUUID} | Get Job Schedule
[**PageListJobInstances**](JobSchedulerApi.md#PageListJobInstances) | **Post** /customers/{cUUID}/job-schedules/{jUUID}/job-instances/page | List Job Instances (paginated)
[**PageListJobSchedules**](JobSchedulerApi.md#PageListJobSchedules) | **Post** /customers/{cUUID}/job-schedules/page | List Job Schedules (paginated)
[**SnoozeJobSchedule**](JobSchedulerApi.md#SnoozeJobSchedule) | **Post** /customers/{cUUID}/job-schedules/{jUUID}/snooze | Snooze Job Schedule
[**UpdateJobSchedule**](JobSchedulerApi.md#UpdateJobSchedule) | **Put** /customers/{cUUID}/job-schedules/{jUUID} | Update Job Schedule



## DeleteJobSchedule

> JobSchedule DeleteJobSchedule(ctx, cUUID, jUUID).Execute()

Delete Job Schedule



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
    jUUID := TODO // string | Job Schedule UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.JobSchedulerApi.DeleteJobSchedule(context.Background(), cUUID, jUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobSchedulerApi.DeleteJobSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteJobSchedule`: JobSchedule
    fmt.Fprintf(os.Stdout, "Response from `JobSchedulerApi.DeleteJobSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**jUUID** | [**string**](.md) | Job Schedule UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJobScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**JobSchedule**](JobSchedule.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobSchedule

> JobSchedule GetJobSchedule(ctx, cUUID, jUUID).Execute()

Get Job Schedule



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
    jUUID := TODO // string | Job Schedule UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.JobSchedulerApi.GetJobSchedule(context.Background(), cUUID, jUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobSchedulerApi.GetJobSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobSchedule`: JobSchedule
    fmt.Fprintf(os.Stdout, "Response from `JobSchedulerApi.GetJobSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**jUUID** | [**string**](.md) | Job Schedule UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**JobSchedule**](JobSchedule.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageListJobInstances

> JobInstancePagedResp PageListJobInstances(ctx, cUUID, jUUID).JobInstancePagedQuerySpec(jobInstancePagedQuerySpec).Execute()

List Job Instances (paginated)



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
    jUUID := TODO // string | Job Schedule UUID
    jobInstancePagedQuerySpec := *openapiclient.NewJobInstancePagedQuerySpec() // JobInstancePagedQuerySpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.JobSchedulerApi.PageListJobInstances(context.Background(), cUUID, jUUID).JobInstancePagedQuerySpec(jobInstancePagedQuerySpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobSchedulerApi.PageListJobInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PageListJobInstances`: JobInstancePagedResp
    fmt.Fprintf(os.Stdout, "Response from `JobSchedulerApi.PageListJobInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**jUUID** | [**string**](.md) | Job Schedule UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageListJobInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **jobInstancePagedQuerySpec** | [**JobInstancePagedQuerySpec**](JobInstancePagedQuerySpec.md) |  | 

### Return type

[**JobInstancePagedResp**](JobInstancePagedResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageListJobSchedules

> JobSchedulePagedResp PageListJobSchedules(ctx, cUUID).JobSchedulePagedQuerySpec(jobSchedulePagedQuerySpec).Execute()

List Job Schedules (paginated)



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
    jobSchedulePagedQuerySpec := *openapiclient.NewJobSchedulePagedQuerySpec() // JobSchedulePagedQuerySpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.JobSchedulerApi.PageListJobSchedules(context.Background(), cUUID).JobSchedulePagedQuerySpec(jobSchedulePagedQuerySpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobSchedulerApi.PageListJobSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PageListJobSchedules`: JobSchedulePagedResp
    fmt.Fprintf(os.Stdout, "Response from `JobSchedulerApi.PageListJobSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageListJobSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jobSchedulePagedQuerySpec** | [**JobSchedulePagedQuerySpec**](JobSchedulePagedQuerySpec.md) |  | 

### Return type

[**JobSchedulePagedResp**](JobSchedulePagedResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SnoozeJobSchedule

> JobSchedule SnoozeJobSchedule(ctx, cUUID, jUUID).JobScheduleSnoozeSpec(jobScheduleSnoozeSpec).Execute()

Snooze Job Schedule



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
    jUUID := TODO // string | Job Schedule UUID
    jobScheduleSnoozeSpec := *openapiclient.NewJobScheduleSnoozeSpec(int64(123)) // JobScheduleSnoozeSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.JobSchedulerApi.SnoozeJobSchedule(context.Background(), cUUID, jUUID).JobScheduleSnoozeSpec(jobScheduleSnoozeSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobSchedulerApi.SnoozeJobSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SnoozeJobSchedule`: JobSchedule
    fmt.Fprintf(os.Stdout, "Response from `JobSchedulerApi.SnoozeJobSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**jUUID** | [**string**](.md) | Job Schedule UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSnoozeJobScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **jobScheduleSnoozeSpec** | [**JobScheduleSnoozeSpec**](JobScheduleSnoozeSpec.md) |  | 

### Return type

[**JobSchedule**](JobSchedule.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateJobSchedule

> JobSchedule UpdateJobSchedule(ctx, cUUID, jUUID).JobScheduleUpdateSpec(jobScheduleUpdateSpec).Execute()

Update Job Schedule



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
    jUUID := TODO // string | Job Schedule UUID
    jobScheduleUpdateSpec := *openapiclient.NewJobScheduleUpdateSpec(false, int64(123), openapiclient.JobScheduleType("FIXED_DELAY")) // JobScheduleUpdateSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.JobSchedulerApi.UpdateJobSchedule(context.Background(), cUUID, jUUID).JobScheduleUpdateSpec(jobScheduleUpdateSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobSchedulerApi.UpdateJobSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateJobSchedule`: JobSchedule
    fmt.Fprintf(os.Stdout, "Response from `JobSchedulerApi.UpdateJobSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**jUUID** | [**string**](.md) | Job Schedule UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateJobScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **jobScheduleUpdateSpec** | [**JobScheduleUpdateSpec**](JobScheduleUpdateSpec.md) |  | 

### Return type

[**JobSchedule**](JobSchedule.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

