# \JobSchedulerAPI

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteJobSchedule**](JobSchedulerAPI.md#DeleteJobSchedule) | **Delete** /customers/{cUUID}/job-schedules/{jUUID} | Delete Job Schedule
[**GetJobSchedule**](JobSchedulerAPI.md#GetJobSchedule) | **Get** /customers/{cUUID}/job-schedules/{jUUID} | Get Job Schedule
[**PageListJobInstances**](JobSchedulerAPI.md#PageListJobInstances) | **Post** /customers/{cUUID}/job-schedules/{jUUID}/job-instances/page | List Job Instances (paginated)
[**PageListJobSchedules**](JobSchedulerAPI.md#PageListJobSchedules) | **Post** /customers/{cUUID}/job-schedules/page | List Job Schedules (paginated)
[**SnoozeJobSchedule**](JobSchedulerAPI.md#SnoozeJobSchedule) | **Post** /customers/{cUUID}/job-schedules/{jUUID}/snooze | Snooze Job Schedule
[**UpdateJobSchedule**](JobSchedulerAPI.md#UpdateJobSchedule) | **Put** /customers/{cUUID}/job-schedules/{jUUID} | Update Job Schedule



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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	jUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job Schedule UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobSchedulerAPI.DeleteJobSchedule(context.Background(), cUUID, jUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobSchedulerAPI.DeleteJobSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteJobSchedule`: JobSchedule
	fmt.Fprintf(os.Stdout, "Response from `JobSchedulerAPI.DeleteJobSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**jUUID** | **string** | Job Schedule UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	jUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job Schedule UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobSchedulerAPI.GetJobSchedule(context.Background(), cUUID, jUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobSchedulerAPI.GetJobSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobSchedule`: JobSchedule
	fmt.Fprintf(os.Stdout, "Response from `JobSchedulerAPI.GetJobSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**jUUID** | **string** | Job Schedule UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	jUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job Schedule UUID
	jobInstancePagedQuerySpec := *openapiclient.NewJobInstancePagedQuerySpec() // JobInstancePagedQuerySpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobSchedulerAPI.PageListJobInstances(context.Background(), cUUID, jUUID).JobInstancePagedQuerySpec(jobInstancePagedQuerySpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobSchedulerAPI.PageListJobInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageListJobInstances`: JobInstancePagedResp
	fmt.Fprintf(os.Stdout, "Response from `JobSchedulerAPI.PageListJobInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**jUUID** | **string** | Job Schedule UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	jobSchedulePagedQuerySpec := *openapiclient.NewJobSchedulePagedQuerySpec() // JobSchedulePagedQuerySpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobSchedulerAPI.PageListJobSchedules(context.Background(), cUUID).JobSchedulePagedQuerySpec(jobSchedulePagedQuerySpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobSchedulerAPI.PageListJobSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageListJobSchedules`: JobSchedulePagedResp
	fmt.Fprintf(os.Stdout, "Response from `JobSchedulerAPI.PageListJobSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	jUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job Schedule UUID
	jobScheduleSnoozeSpec := *openapiclient.NewJobScheduleSnoozeSpec(int64(123)) // JobScheduleSnoozeSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobSchedulerAPI.SnoozeJobSchedule(context.Background(), cUUID, jUUID).JobScheduleSnoozeSpec(jobScheduleSnoozeSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobSchedulerAPI.SnoozeJobSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SnoozeJobSchedule`: JobSchedule
	fmt.Fprintf(os.Stdout, "Response from `JobSchedulerAPI.SnoozeJobSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**jUUID** | **string** | Job Schedule UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	jUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Job Schedule UUID
	jobScheduleUpdateSpec := *openapiclient.NewJobScheduleUpdateSpec(false, int64(123), openapiclient.JobScheduleType("FIXED_DELAY")) // JobScheduleUpdateSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobSchedulerAPI.UpdateJobSchedule(context.Background(), cUUID, jUUID).JobScheduleUpdateSpec(jobScheduleUpdateSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobSchedulerAPI.UpdateJobSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateJobSchedule`: JobSchedule
	fmt.Fprintf(os.Stdout, "Response from `JobSchedulerAPI.UpdateJobSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**jUUID** | **string** | Job Schedule UUID | 

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

