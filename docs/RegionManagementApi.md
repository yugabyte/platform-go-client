# \RegionManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRegion**](RegionManagementApi.md#CreateRegion) | **Post** /api/v1/customers/{cUUID}/providers/{pUUID}/regions | Create a new region
[**DeleteRegion**](RegionManagementApi.md#DeleteRegion) | **Delete** /api/v1/customers/{cUUID}/providers/{pUUID}/regions/{rUUID} | Delete a region
[**EditRegion**](RegionManagementApi.md#EditRegion) | **Put** /api/v1/customers/{cUUID}/providers/{pUUID}/regions/{rUUID} | Modify a region
[**GetRegion**](RegionManagementApi.md#GetRegion) | **Get** /api/v1/customers/{cUUID}/providers/{pUUID}/regions | List a provider&#39;s regions
[**ListAllRegions**](RegionManagementApi.md#ListAllRegions) | **Get** /api/v1/customers/{cUUID}/regions | List regions for all providers



## CreateRegion

> Region CreateRegion(ctx, cUUID, pUUID).Region(region).Execute()

Create a new region

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
    region := *openapiclient.NewRegionFormData("Code_example", "DestVpcId_example", "HostVpcId_example", "HostVpcRegion_example", float64(123), float64(123), "Name_example", "SecurityGroupId_example", "VnetName_example", "YbImage_example") // RegionFormData | region form data for new region to be created

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegionManagementApi.CreateRegion(context.Background(), cUUID, pUUID).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionManagementApi.CreateRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRegion`: Region
    fmt.Fprintf(os.Stdout, "Response from `RegionManagementApi.CreateRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **region** | [**RegionFormData**](RegionFormData.md) | region form data for new region to be created | 

### Return type

[**Region**](Region.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRegion

> map[string]interface{} DeleteRegion(ctx, cUUID, pUUID, rUUID).Execute()

Delete a region

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
    resp, r, err := api_client.RegionManagementApi.DeleteRegion(context.Background(), cUUID, pUUID, rUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionManagementApi.DeleteRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRegion`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RegionManagementApi.DeleteRegion`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteRegionRequest struct via the builder pattern


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


## EditRegion

> map[string]interface{} EditRegion(ctx, cUUID, pUUID, rUUID).Region(region).Execute()

Modify a region

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
    region := *openapiclient.NewRegionEditFormData("SecurityGroupId_example", "VnetName_example", "YbImage_example") // RegionEditFormData | region edit form data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegionManagementApi.EditRegion(context.Background(), cUUID, pUUID, rUUID).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionManagementApi.EditRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditRegion`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RegionManagementApi.EditRegion`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **region** | [**RegionEditFormData**](RegionEditFormData.md) | region edit form data | 

### Return type

**map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegion

> []Region GetRegion(ctx, cUUID, pUUID).Execute()

List a provider's regions

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegionManagementApi.GetRegion(context.Background(), cUUID, pUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionManagementApi.GetRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegion`: []Region
    fmt.Fprintf(os.Stdout, "Response from `RegionManagementApi.GetRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Region**](Region.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllRegions

> []Region ListAllRegions(ctx, cUUID).Execute()

List regions for all providers

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
    resp, r, err := api_client.RegionManagementApi.ListAllRegions(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionManagementApi.ListAllRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllRegions`: []Region
    fmt.Fprintf(os.Stdout, "Response from `RegionManagementApi.ListAllRegions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Region**](Region.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

