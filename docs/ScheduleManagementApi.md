# \ScheduleManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBackupScheduleAsync**](ScheduleManagementApi.md#DeleteBackupScheduleAsync) | **Delete** /api/v1/customers/{cUUID}/universes/{uniUUID}/schedules/{sUUID}/delete_backup_schedule_async | Delete a backup schedule async
[**DeleteSchedule**](ScheduleManagementApi.md#DeleteSchedule) | **Delete** /api/v1/customers/{cUUID}/schedules/{sUUID} | Delete a schedule  - deprecated
[**DeleteScheduleV2**](ScheduleManagementApi.md#DeleteScheduleV2) | **Delete** /api/v1/customers/{cUUID}/schedules/{sUUID}/delete | Delete a schedule V2 - deprecated
[**EditBackupScheduleAsync**](ScheduleManagementApi.md#EditBackupScheduleAsync) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/schedules/{sUUID}/edit_backup_schedule_async | Edit a backup schedule async
[**EditBackupScheduleV2**](ScheduleManagementApi.md#EditBackupScheduleV2) | **Put** /api/v1/customers/{cUUID}/schedules/{sUUID} | Edit a backup schedule V2 - deprecated
[**GetSchedule**](ScheduleManagementApi.md#GetSchedule) | **Get** /api/v1/customers/{cUUID}/schedules/{sUUID} | Get Schedule
[**ListSchedules**](ScheduleManagementApi.md#ListSchedules) | **Get** /api/v1/customers/{cUUID}/schedules | List schedules - deprecated
[**ListSchedulesV2**](ScheduleManagementApi.md#ListSchedulesV2) | **Post** /api/v1/customers/{cUUID}/schedules/page | List schedules V2
[**ToggleBackupSchedule**](ScheduleManagementApi.md#ToggleBackupSchedule) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/schedules/{sUUID}/pause_resume | Toggle a backup schedule



## DeleteBackupScheduleAsync

> Schedule DeleteBackupScheduleAsync(ctx, cUUID, uniUUID, sUUID).Request(request).Execute()

Delete a backup schedule async



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
    uniUUID := TODO // string | 
    sUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduleManagementApi.DeleteBackupScheduleAsync(context.Background(), cUUID, uniUUID, sUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleManagementApi.DeleteBackupScheduleAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBackupScheduleAsync`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `ScheduleManagementApi.DeleteBackupScheduleAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 
**sUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackupScheduleAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSchedule

> YBPSuccess DeleteSchedule(ctx, cUUID, sUUID).Request(request).Execute()

Delete a schedule  - deprecated



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
    sUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduleManagementApi.DeleteSchedule(context.Background(), cUUID, sUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleManagementApi.DeleteSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSchedule`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `ScheduleManagementApi.DeleteSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**sUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScheduleRequest struct via the builder pattern


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


## DeleteScheduleV2

> YBPSuccess DeleteScheduleV2(ctx, cUUID, sUUID).Request(request).Execute()

Delete a schedule V2 - deprecated



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
    sUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduleManagementApi.DeleteScheduleV2(context.Background(), cUUID, sUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleManagementApi.DeleteScheduleV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteScheduleV2`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `ScheduleManagementApi.DeleteScheduleV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**sUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScheduleV2Request struct via the builder pattern


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


## EditBackupScheduleAsync

> Schedule EditBackupScheduleAsync(ctx, cUUID, uniUUID, sUUID).Body(body).Request(request).Execute()

Edit a backup schedule async



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
    uniUUID := TODO // string | 
    sUUID := TODO // string | 
    body := *openapiclient.NewBackupScheduleEditParams() // BackupScheduleEditParams | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduleManagementApi.EditBackupScheduleAsync(context.Background(), cUUID, uniUUID, sUUID).Body(body).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleManagementApi.EditBackupScheduleAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditBackupScheduleAsync`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `ScheduleManagementApi.EditBackupScheduleAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 
**sUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditBackupScheduleAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**BackupScheduleEditParams**](BackupScheduleEditParams.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditBackupScheduleV2

> Schedule EditBackupScheduleV2(ctx, cUUID, sUUID).Body(body).Request(request).Execute()

Edit a backup schedule V2 - deprecated



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
    sUUID := TODO // string | 
    body := *openapiclient.NewEditBackupScheduleParams() // EditBackupScheduleParams | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduleManagementApi.EditBackupScheduleV2(context.Background(), cUUID, sUUID).Body(body).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleManagementApi.EditBackupScheduleV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditBackupScheduleV2`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `ScheduleManagementApi.EditBackupScheduleV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**sUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditBackupScheduleV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**EditBackupScheduleParams**](EditBackupScheduleParams.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchedule

> Schedule GetSchedule(ctx, cUUID, sUUID).Execute()

Get Schedule

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
    sUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduleManagementApi.GetSchedule(context.Background(), cUUID, sUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleManagementApi.GetSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchedule`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `ScheduleManagementApi.GetSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**sUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Schedule**](Schedule.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSchedules

> []Schedule ListSchedules(ctx, cUUID).Execute()

List schedules - deprecated



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduleManagementApi.ListSchedules(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleManagementApi.ListSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSchedules`: []Schedule
    fmt.Fprintf(os.Stdout, "Response from `ScheduleManagementApi.ListSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Schedule**](Schedule.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSchedulesV2

> SchedulePagedResponse ListSchedulesV2(ctx, cUUID).PageScheduleRequest(pageScheduleRequest).Request(request).Execute()

List schedules V2

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
    pageScheduleRequest := *openapiclient.NewSchedulePagedApiQuery("Direction_example", *openapiclient.NewScheduleApiFilter([]string{"Status_example"}, []string{"TaskTypes_example"}, []string{"UniverseUUIDList_example"}), int32(123), false, int32(123), "SortBy_example") // SchedulePagedApiQuery | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduleManagementApi.ListSchedulesV2(context.Background(), cUUID).PageScheduleRequest(pageScheduleRequest).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleManagementApi.ListSchedulesV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSchedulesV2`: SchedulePagedResponse
    fmt.Fprintf(os.Stdout, "Response from `ScheduleManagementApi.ListSchedulesV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSchedulesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageScheduleRequest** | [**SchedulePagedApiQuery**](SchedulePagedApiQuery.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**SchedulePagedResponse**](SchedulePagedResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleBackupSchedule

> Schedule ToggleBackupSchedule(ctx, cUUID, uniUUID, sUUID).Body(body).Request(request).Execute()

Toggle a backup schedule



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
    uniUUID := TODO // string | 
    sUUID := TODO // string | 
    body := *openapiclient.NewBackupScheduleToggleParams("Status_example") // BackupScheduleToggleParams | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduleManagementApi.ToggleBackupSchedule(context.Background(), cUUID, uniUUID, sUUID).Body(body).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleManagementApi.ToggleBackupSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ToggleBackupSchedule`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `ScheduleManagementApi.ToggleBackupSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 
**sUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**BackupScheduleToggleParams**](BackupScheduleToggleParams.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

