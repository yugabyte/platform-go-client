# \AvailabilityZonesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAZ**](AvailabilityZonesApi.md#CreateAZ) | **Post** /api/v1/customers/{cUUID}/providers/{pUUID}/regions/{rUUID}/zones | Create an availability zone - deprecated
[**CreateZone**](AvailabilityZonesApi.md#CreateZone) | **Post** /api/v1/customers/{cUUID}/providers/{pUUID}/provider_regions/{rUUID}/region_zones | Create an availability zone
[**DeleteAZ**](AvailabilityZonesApi.md#DeleteAZ) | **Delete** /api/v1/customers/{cUUID}/providers/{pUUID}/regions/{rUUID}/zones/{azUUID} | Delete an availability zone
[**EditAZ**](AvailabilityZonesApi.md#EditAZ) | **Put** /api/v1/customers/{cUUID}/providers/{pUUID}/regions/{rUUID}/zones/{azUUID} | Edit an Availabilty Zone - deprecated
[**EditZone**](AvailabilityZonesApi.md#EditZone) | **Put** /api/v1/customers/{cUUID}/providers/{pUUID}/provider_regions/{rUUID}/region_zones/{azUUID} | Modify an availability zone
[**ListOfAZ**](AvailabilityZonesApi.md#ListOfAZ) | **Get** /api/v1/customers/{cUUID}/providers/{pUUID}/regions/{rUUID}/zones | List availability zones



## CreateAZ

> map[string]AvailabilityZone CreateAZ(ctx, cUUID, pUUID, rUUID).AzFormData(azFormData).Request(request).Execute()

Create an availability zone - deprecated



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
    pUUID := TODO // string | 
    rUUID := TODO // string | 
    azFormData := *openapiclient.NewAvailabilityZoneFormData([]openapiclient.AvailabilityZoneData{*openapiclient.NewAvailabilityZoneData("Code_example", "Name_example")}) // AvailabilityZoneFormData | Availability zone form data
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AvailabilityZonesApi.CreateAZ(context.Background(), cUUID, pUUID, rUUID).AzFormData(azFormData).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvailabilityZonesApi.CreateAZ``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAZ`: map[string]AvailabilityZone
    fmt.Fprintf(os.Stdout, "Response from `AvailabilityZonesApi.CreateAZ`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 
**rUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAZRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **azFormData** | [**AvailabilityZoneFormData**](AvailabilityZoneFormData.md) | Availability zone form data | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**map[string]AvailabilityZone**](AvailabilityZone.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateZone

> map[string]AvailabilityZone CreateZone(ctx, cUUID, pUUID, rUUID).AzData(azData).Request(request).Execute()

Create an availability zone



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
    pUUID := TODO // string | 
    rUUID := TODO // string | 
    azData := *openapiclient.NewAvailabilityZone("us-west1-a") // AvailabilityZone | Availability Zone to create
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AvailabilityZonesApi.CreateZone(context.Background(), cUUID, pUUID, rUUID).AzData(azData).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvailabilityZonesApi.CreateZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateZone`: map[string]AvailabilityZone
    fmt.Fprintf(os.Stdout, "Response from `AvailabilityZonesApi.CreateZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 
**rUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **azData** | [**AvailabilityZone**](AvailabilityZone.md) | Availability Zone to create | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**map[string]AvailabilityZone**](AvailabilityZone.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAZ

> YBPSuccess DeleteAZ(ctx, cUUID, pUUID, rUUID, azUUID).Request(request).Execute()

Delete an availability zone

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
    pUUID := TODO // string | 
    rUUID := TODO // string | 
    azUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AvailabilityZonesApi.DeleteAZ(context.Background(), cUUID, pUUID, rUUID, azUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvailabilityZonesApi.DeleteAZ``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAZ`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `AvailabilityZonesApi.DeleteAZ`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 
**rUUID** | [**string**](.md) |  | 
**azUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAZRequest struct via the builder pattern


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


## EditAZ

> AvailabilityZone EditAZ(ctx, cUUID, pUUID, rUUID, azUUID).AzFormData(azFormData).Request(request).Execute()

Edit an Availabilty Zone - deprecated



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
    pUUID := TODO // string | 
    rUUID := TODO // string | 
    azUUID := TODO // string | 
    azFormData := *openapiclient.NewAvailabilityZoneFormData([]openapiclient.AvailabilityZoneData{*openapiclient.NewAvailabilityZoneData("Code_example", "Name_example")}) // AvailabilityZoneFormData | Availability zone edit form data
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AvailabilityZonesApi.EditAZ(context.Background(), cUUID, pUUID, rUUID, azUUID).AzFormData(azFormData).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvailabilityZonesApi.EditAZ``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditAZ`: AvailabilityZone
    fmt.Fprintf(os.Stdout, "Response from `AvailabilityZonesApi.EditAZ`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 
**rUUID** | [**string**](.md) |  | 
**azUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAZRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **azFormData** | [**AvailabilityZoneFormData**](AvailabilityZoneFormData.md) | Availability zone edit form data | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**AvailabilityZone**](AvailabilityZone.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditZone

> AvailabilityZone EditZone(ctx, cUUID, pUUID, rUUID, azUUID).AzData(azData).Request(request).Execute()

Modify an availability zone



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
    pUUID := TODO // string | 
    rUUID := TODO // string | 
    azUUID := TODO // string | 
    azData := *openapiclient.NewAvailabilityZone("us-west1-a") // AvailabilityZone | Availability zone to edit
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AvailabilityZonesApi.EditZone(context.Background(), cUUID, pUUID, rUUID, azUUID).AzData(azData).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvailabilityZonesApi.EditZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditZone`: AvailabilityZone
    fmt.Fprintf(os.Stdout, "Response from `AvailabilityZonesApi.EditZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 
**rUUID** | [**string**](.md) |  | 
**azUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **azData** | [**AvailabilityZone**](AvailabilityZone.md) | Availability zone to edit | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**AvailabilityZone**](AvailabilityZone.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOfAZ

> []AvailabilityZone ListOfAZ(ctx, cUUID, pUUID, rUUID).Execute()

List availability zones

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
    pUUID := TODO // string | 
    rUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AvailabilityZonesApi.ListOfAZ(context.Background(), cUUID, pUUID, rUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvailabilityZonesApi.ListOfAZ``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOfAZ`: []AvailabilityZone
    fmt.Fprintf(os.Stdout, "Response from `AvailabilityZonesApi.ListOfAZ`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 
**rUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOfAZRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]AvailabilityZone**](AvailabilityZone.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

