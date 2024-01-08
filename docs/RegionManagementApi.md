# \RegionManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProviderRegion**](RegionManagementApi.md#CreateProviderRegion) | **Post** /api/v1/customers/{cUUID}/providers/{pUUID}/provider_regions | WARNING: This is a preview API that could change. Create a new region
[**CreateRegion**](RegionManagementApi.md#CreateRegion) | **Post** /api/v1/customers/{cUUID}/providers/{pUUID}/regions | Deprecated since YBA version 2.18.2.0, Use /api/v1/customers/{cUUID}/provider/{pUUID}/provider_regions instead
[**DeleteRegion**](RegionManagementApi.md#DeleteRegion) | **Delete** /api/v1/customers/{cUUID}/providers/{pUUID}/regions/{rUUID} | Delete a region
[**EditProviderRegion**](RegionManagementApi.md#EditProviderRegion) | **Put** /api/v1/customers/{cUUID}/providers/{pUUID}/provider_regions/{rUUID} | WARNING: This is a preview API that could change. Modify a region
[**EditRegion**](RegionManagementApi.md#EditRegion) | **Put** /api/v1/customers/{cUUID}/providers/{pUUID}/regions/{rUUID} | Deprecated since YBA version 2.18.2.0, Use /api/v1/customers/{cUUID}/provider/{pUUID}/provider_regions instead
[**GetRegion**](RegionManagementApi.md#GetRegion) | **Get** /api/v1/customers/{cUUID}/providers/{pUUID}/regions | List a provider&#39;s regions
[**ListAllRegions**](RegionManagementApi.md#ListAllRegions) | **Get** /api/v1/customers/{cUUID}/regions | List regions for all providers



## CreateProviderRegion

> Region CreateProviderRegion(ctx, cUUID, pUUID).Region(region).Request(request).Execute()

WARNING: This is a preview API that could change. Create a new region

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
    region := *openapiclient.NewRegion([]openapiclient.AvailabilityZone{*openapiclient.NewAvailabilityZone("us-west1-a")}) // Region | Specification of Region to be created
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegionManagementApi.CreateProviderRegion(context.Background(), cUUID, pUUID).Region(region).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionManagementApi.CreateProviderRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProviderRegion`: Region
    fmt.Fprintf(os.Stdout, "Response from `RegionManagementApi.CreateProviderRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProviderRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **region** | [**Region**](Region.md) | Specification of Region to be created | 
 **request** | [**interface{}**](interface{}.md) |  | 

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


## CreateRegion

> Region CreateRegion(ctx, cUUID, pUUID).Region(region).Request(request).Execute()

Deprecated since YBA version 2.18.2.0, Use /api/v1/customers/{cUUID}/provider/{pUUID}/provider_regions instead

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegionManagementApi.CreateRegion(context.Background(), cUUID, pUUID).Region(region).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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

> map[string]interface{} DeleteRegion(ctx, cUUID, pUUID, rUUID).Request(request).Execute()

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegionManagementApi.DeleteRegion(context.Background(), cUUID, pUUID, rUUID).Request(request).Execute()
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



 **request** | [**interface{}**](interface{}.md) |  | 

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


## EditProviderRegion

> Region EditProviderRegion(ctx, cUUID, pUUID, rUUID).Region(region).Request(request).Execute()

WARNING: This is a preview API that could change. Modify a region

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
    region := *openapiclient.NewRegion([]openapiclient.AvailabilityZone{*openapiclient.NewAvailabilityZone("us-west1-a")}) // Region | Specification of Region to be edited
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegionManagementApi.EditProviderRegion(context.Background(), cUUID, pUUID, rUUID).Region(region).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionManagementApi.EditProviderRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditProviderRegion`: Region
    fmt.Fprintf(os.Stdout, "Response from `RegionManagementApi.EditProviderRegion`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiEditProviderRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **region** | [**Region**](Region.md) | Specification of Region to be edited | 
 **request** | [**interface{}**](interface{}.md) |  | 

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


## EditRegion

> map[string]interface{} EditRegion(ctx, cUUID, pUUID, rUUID).Region(region).Request(request).Execute()

Deprecated since YBA version 2.18.2.0, Use /api/v1/customers/{cUUID}/provider/{pUUID}/provider_regions instead

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
    region := *openapiclient.NewRegionFormData("Code_example", "DestVpcId_example", "HostVpcId_example", "HostVpcRegion_example", float64(123), float64(123), "Name_example", "SecurityGroupId_example", "VnetName_example", "YbImage_example") // RegionFormData | region edit form data
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegionManagementApi.EditRegion(context.Background(), cUUID, pUUID, rUUID).Region(region).Request(request).Execute()
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



 **region** | [**RegionFormData**](RegionFormData.md) | region edit form data | 
 **request** | [**interface{}**](interface{}.md) |  | 

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

