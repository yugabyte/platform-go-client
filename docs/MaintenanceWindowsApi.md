# \MaintenanceWindowsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](MaintenanceWindowsApi.md#Create) | **Post** /api/v1/customers/{cUUID}/maintenance_windows | Create maintenance window
[**Delete**](MaintenanceWindowsApi.md#Delete) | **Delete** /api/v1/customers/{cUUID}/maintenance_windows/{windowUUID} | Delete maintenance window
[**Get**](MaintenanceWindowsApi.md#Get) | **Get** /api/v1/customers/{cUUID}/maintenance_windows/{windowUUID} | Get details of a maintenance window
[**ListOfMaintenanceWindows**](MaintenanceWindowsApi.md#ListOfMaintenanceWindows) | **Post** /api/v1/customers/{cUUID}/maintenance_windows/list | List maintenance windows
[**Page**](MaintenanceWindowsApi.md#Page) | **Post** /api/v1/customers/{cUUID}/maintenance_windows/page | List maintenance windows (paginated)
[**Update**](MaintenanceWindowsApi.md#Update) | **Put** /api/v1/customers/{cUUID}/maintenance_windows/{windowUUID} | Update maintenance window



## Create

> MaintenanceWindow Create(ctx, cUUID).CreateMaintenanceWindowRequest(createMaintenanceWindowRequest).Request(request).Execute()

Create maintenance window

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    createMaintenanceWindowRequest := *openapiclient.NewMaintenanceWindow(*openapiclient.NewAlertConfigurationApiFilter(false, "DestinationType_example", "DestinationUuid_example", "Name_example", "Severity_example", *openapiclient.NewAlertConfigurationTarget(), "TargetType_example", "Template_example", []string{"Uuids_example"}), time.Now(), "CustomerUUID_example", "Description_example", time.Now(), "Name_example", time.Now()) // MaintenanceWindow | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaintenanceWindowsApi.Create(context.Background(), cUUID).CreateMaintenanceWindowRequest(createMaintenanceWindowRequest).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: MaintenanceWindow
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsApi.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createMaintenanceWindowRequest** | [**MaintenanceWindow**](MaintenanceWindow.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> YBPSuccess Delete(ctx, cUUID, windowUUID).Request(request).Execute()

Delete maintenance window

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
    windowUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaintenanceWindowsApi.Delete(context.Background(), cUUID, windowUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsApi.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**windowUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


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


## Get

> MaintenanceWindow Get(ctx, cUUID, windowUUID).Execute()

Get details of a maintenance window

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
    windowUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaintenanceWindowsApi.Get(context.Background(), cUUID, windowUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: MaintenanceWindow
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**windowUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOfMaintenanceWindows

> []MaintenanceWindow ListOfMaintenanceWindows(ctx, cUUID).Request(request).Execute()

List maintenance windows

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
    resp, r, err := api_client.MaintenanceWindowsApi.ListOfMaintenanceWindows(context.Background(), cUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.ListOfMaintenanceWindows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOfMaintenanceWindows`: []MaintenanceWindow
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsApi.ListOfMaintenanceWindows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOfMaintenanceWindowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**[]MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Page

> MaintenanceWindowPagedResponse Page(ctx, cUUID).PageMaintenanceWindowsRequest(pageMaintenanceWindowsRequest).Request(request).Execute()

List maintenance windows (paginated)

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
    pageMaintenanceWindowsRequest := *openapiclient.NewMaintenanceWindowPagedApiQuery("Direction_example", *openapiclient.NewMaintenanceWindowApiFilter([]string{"States_example"}, []string{"Uuids_example"}), int32(123), false, int32(123), "SortBy_example") // MaintenanceWindowPagedApiQuery | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaintenanceWindowsApi.Page(context.Background(), cUUID).PageMaintenanceWindowsRequest(pageMaintenanceWindowsRequest).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.Page``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Page`: MaintenanceWindowPagedResponse
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsApi.Page`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageMaintenanceWindowsRequest** | [**MaintenanceWindowPagedApiQuery**](MaintenanceWindowPagedApiQuery.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**MaintenanceWindowPagedResponse**](MaintenanceWindowPagedResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> MaintenanceWindow Update(ctx, cUUID, windowUUID).UpdateMaintenanceWindowRequest(updateMaintenanceWindowRequest).Request(request).Execute()

Update maintenance window

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    windowUUID := TODO // string | 
    updateMaintenanceWindowRequest := *openapiclient.NewMaintenanceWindow(*openapiclient.NewAlertConfigurationApiFilter(false, "DestinationType_example", "DestinationUuid_example", "Name_example", "Severity_example", *openapiclient.NewAlertConfigurationTarget(), "TargetType_example", "Template_example", []string{"Uuids_example"}), time.Now(), "CustomerUUID_example", "Description_example", time.Now(), "Name_example", time.Now()) // MaintenanceWindow | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaintenanceWindowsApi.Update(context.Background(), cUUID, windowUUID).UpdateMaintenanceWindowRequest(updateMaintenanceWindowRequest).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: MaintenanceWindow
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**windowUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateMaintenanceWindowRequest** | [**MaintenanceWindow**](MaintenanceWindow.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

