# \ReleaseManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRelease**](ReleaseManagementApi.md#CreateRelease) | **Post** /api/v1/customers/{cUUID}/releases | Create a release
[**DeleteRelease**](ReleaseManagementApi.md#DeleteRelease) | **Delete** /api/v1/customers/{cUUID}/releases/{name} | Delete a release
[**GetListOfRegionReleases**](ReleaseManagementApi.md#GetListOfRegionReleases) | **Get** /api/v1/customers/{cUUID}/providers/{pUUID}/releases | List all releases valid in region
[**GetListOfReleases**](ReleaseManagementApi.md#GetListOfReleases) | **Get** /api/v1/customers/{cUUID}/releases | List all releases
[**Refresh**](ReleaseManagementApi.md#Refresh) | **Put** /api/v1/customers/{cUUID}/releases | Refresh a release
[**UpdateRelease**](ReleaseManagementApi.md#UpdateRelease) | **Put** /api/v1/customers/{cUUID}/releases/{name} | Update a release



## CreateRelease

> YBPSuccess CreateRelease(ctx, cUUID).Release(release).Request(request).Execute()

Create a release

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
    release := *openapiclient.NewReleaseFormData(*openapiclient.NewGCSLocation(), *openapiclient.NewHttpLocation(), *openapiclient.NewS3Location(), "Version_example") // ReleaseFormData | Release data for remote downloading to be created
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReleaseManagementApi.CreateRelease(context.Background(), cUUID).Release(release).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementApi.CreateRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRelease`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementApi.CreateRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **release** | [**ReleaseFormData**](ReleaseFormData.md) | Release data for remote downloading to be created | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPSuccess**](YBPSuccess.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRelease

> ReleaseMetadata DeleteRelease(ctx, cUUID, name).Request(request).Execute()

Delete a release

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
    name := "name_example" // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReleaseManagementApi.DeleteRelease(context.Background(), cUUID, name).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementApi.DeleteRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRelease`: ReleaseMetadata
    fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementApi.DeleteRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**ReleaseMetadata**](ReleaseMetadata.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListOfRegionReleases

> map[string]map[string]interface{} GetListOfRegionReleases(ctx, cUUID, pUUID).IncludeMetadata(includeMetadata).Execute()

List all releases valid in region

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
    includeMetadata := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReleaseManagementApi.GetListOfRegionReleases(context.Background(), cUUID, pUUID).IncludeMetadata(includeMetadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementApi.GetListOfRegionReleases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListOfRegionReleases`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementApi.GetListOfRegionReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListOfRegionReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeMetadata** | **bool** |  | [default to false]

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


## GetListOfReleases

> map[string]map[string]interface{} GetListOfReleases(ctx, cUUID).IncludeMetadata(includeMetadata).Execute()

List all releases

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
    includeMetadata := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReleaseManagementApi.GetListOfReleases(context.Background(), cUUID).IncludeMetadata(includeMetadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementApi.GetListOfReleases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListOfReleases`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementApi.GetListOfReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListOfReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeMetadata** | **bool** |  | [default to false]

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


## Refresh

> YBPSuccess Refresh(ctx, cUUID).Request(request).Execute()

Refresh a release

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
    resp, r, err := api_client.ReleaseManagementApi.Refresh(context.Background(), cUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementApi.Refresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Refresh`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementApi.Refresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshRequest struct via the builder pattern


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


## UpdateRelease

> ReleaseMetadata UpdateRelease(ctx, cUUID, name).Release(release).Request(request).Execute()

Update a release

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
    name := "name_example" // string | 
    release := map[string]interface{}(Object) // map[string]interface{} | Release data to be updated
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReleaseManagementApi.UpdateRelease(context.Background(), cUUID, name).Release(release).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementApi.UpdateRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRelease`: ReleaseMetadata
    fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementApi.UpdateRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **release** | **map[string]interface{}** | Release data to be updated | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**ReleaseMetadata**](ReleaseMetadata.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

