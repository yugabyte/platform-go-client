# \MaintenanceWindowsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](MaintenanceWindowsAPI.md#Create) | **Post** /api/v1/customers/{cUUID}/maintenance_windows | Create maintenance window
[**Delete**](MaintenanceWindowsAPI.md#Delete) | **Delete** /api/v1/customers/{cUUID}/maintenance_windows/{windowUUID} | Delete maintenance window
[**Get**](MaintenanceWindowsAPI.md#Get) | **Get** /api/v1/customers/{cUUID}/maintenance_windows/{windowUUID} | Get details of a maintenance window
[**ListOfMaintenanceWindows**](MaintenanceWindowsAPI.md#ListOfMaintenanceWindows) | **Post** /api/v1/customers/{cUUID}/maintenance_windows/list | List maintenance windows
[**Page**](MaintenanceWindowsAPI.md#Page) | **Post** /api/v1/customers/{cUUID}/maintenance_windows/page | List maintenance windows (paginated)
[**Update**](MaintenanceWindowsAPI.md#Update) | **Put** /api/v1/customers/{cUUID}/maintenance_windows/{windowUUID} | Update maintenance window



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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	createMaintenanceWindowRequest := *openapiclient.NewMaintenanceWindow(*openapiclient.NewAlertConfigurationApiFilter(), time.Now(), "CustomerUUID_example", "Description_example", time.Now(), "Name_example", time.Now()) // MaintenanceWindow | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceWindowsAPI.Create(context.Background(), cUUID).CreateMaintenanceWindowRequest(createMaintenanceWindowRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: MaintenanceWindow
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsAPI.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	windowUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceWindowsAPI.Delete(context.Background(), cUUID, windowUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Delete`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsAPI.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**windowUUID** | **string** |  | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	windowUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceWindowsAPI.Get(context.Background(), cUUID, windowUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: MaintenanceWindow
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsAPI.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**windowUUID** | **string** |  | 

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

> []MaintenanceWindow ListOfMaintenanceWindows(ctx, cUUID).ListMaintenanceWindowsRequest(listMaintenanceWindowsRequest).Request(request).Execute()

List maintenance windows



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
	listMaintenanceWindowsRequest := *openapiclient.NewMaintenanceWindowApiFilter([]string{"States_example"}, []string{"Uuids_example"}) // MaintenanceWindowApiFilter | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceWindowsAPI.ListOfMaintenanceWindows(context.Background(), cUUID).ListMaintenanceWindowsRequest(listMaintenanceWindowsRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsAPI.ListOfMaintenanceWindows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOfMaintenanceWindows`: []MaintenanceWindow
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsAPI.ListOfMaintenanceWindows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOfMaintenanceWindowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **listMaintenanceWindowsRequest** | [**MaintenanceWindowApiFilter**](MaintenanceWindowApiFilter.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**[]MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pageMaintenanceWindowsRequest := *openapiclient.NewMaintenanceWindowPagedApiQuery("Direction_example", *openapiclient.NewMaintenanceWindowApiFilter([]string{"States_example"}, []string{"Uuids_example"}), int32(123), false, int32(123), "SortBy_example") // MaintenanceWindowPagedApiQuery | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceWindowsAPI.Page(context.Background(), cUUID).PageMaintenanceWindowsRequest(pageMaintenanceWindowsRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsAPI.Page``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Page`: MaintenanceWindowPagedResponse
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsAPI.Page`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	windowUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateMaintenanceWindowRequest := *openapiclient.NewMaintenanceWindow(*openapiclient.NewAlertConfigurationApiFilter(), time.Now(), "CustomerUUID_example", "Description_example", time.Now(), "Name_example", time.Now()) // MaintenanceWindow | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceWindowsAPI.Update(context.Background(), cUUID, windowUUID).UpdateMaintenanceWindowRequest(updateMaintenanceWindowRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowsAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: MaintenanceWindow
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowsAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**windowUUID** | **string** |  | 

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

