# \SupportBundleManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSupportBundle**](SupportBundleManagementApi.md#CreateSupportBundle) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/support_bundle | Create support bundle for specific universe
[**DeleteSupportBundle**](SupportBundleManagementApi.md#DeleteSupportBundle) | **Delete** /api/v1/customers/{cUUID}/universes/{uniUUID}/support_bundle/{sbUUID} | Delete a support bundle
[**DownloadSupportBundle**](SupportBundleManagementApi.md#DownloadSupportBundle) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/support_bundle/{sbUUID}/download | Download support bundle
[**GetSupportBundle**](SupportBundleManagementApi.md#GetSupportBundle) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/support_bundle/{sbUUID} | Get a support bundle from a universe
[**ListSupportBundle**](SupportBundleManagementApi.md#ListSupportBundle) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/support_bundle | List all support bundles from a universe
[**ListSupportBundleComponents**](SupportBundleManagementApi.md#ListSupportBundleComponents) | **Get** /api/v1/customers/{cUUID}/support_bundle/components | List all components available in support bundle



## CreateSupportBundle

> YBPTask CreateSupportBundle(ctx, cUUID, uniUUID).SupportBundle(supportBundle).Execute()

Create support bundle for specific universe

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
    uniUUID := TODO // string | 
    supportBundle := *openapiclient.NewSupportBundleFormData([]string{"Components_example"}, time.Now(), time.Now()) // SupportBundleFormData | post support bundle info

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupportBundleManagementApi.CreateSupportBundle(context.Background(), cUUID, uniUUID).SupportBundle(supportBundle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleManagementApi.CreateSupportBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSupportBundle`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `SupportBundleManagementApi.CreateSupportBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSupportBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **supportBundle** | [**SupportBundleFormData**](SupportBundleFormData.md) | post support bundle info | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSupportBundle

> YBPSuccess DeleteSupportBundle(ctx, cUUID, uniUUID, sbUUID).Execute()

Delete a support bundle

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
    sbUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupportBundleManagementApi.DeleteSupportBundle(context.Background(), cUUID, uniUUID, sbUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleManagementApi.DeleteSupportBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSupportBundle`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `SupportBundleManagementApi.DeleteSupportBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 
**sbUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSupportBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## DownloadSupportBundle

> string DownloadSupportBundle(ctx, cUUID, uniUUID, sbUUID).Execute()

Download support bundle

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
    sbUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupportBundleManagementApi.DownloadSupportBundle(context.Background(), cUUID, uniUUID, sbUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleManagementApi.DownloadSupportBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadSupportBundle`: string
    fmt.Fprintf(os.Stdout, "Response from `SupportBundleManagementApi.DownloadSupportBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 
**sbUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadSupportBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-compressed

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportBundle

> SupportBundle GetSupportBundle(ctx, cUUID, uniUUID, sbUUID).Execute()

Get a support bundle from a universe

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
    sbUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupportBundleManagementApi.GetSupportBundle(context.Background(), cUUID, uniUUID, sbUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleManagementApi.GetSupportBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupportBundle`: SupportBundle
    fmt.Fprintf(os.Stdout, "Response from `SupportBundleManagementApi.GetSupportBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 
**sbUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SupportBundle**](SupportBundle.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportBundle

> []SupportBundle ListSupportBundle(ctx, cUUID, uniUUID).Execute()

List all support bundles from a universe

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupportBundleManagementApi.ListSupportBundle(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleManagementApi.ListSupportBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSupportBundle`: []SupportBundle
    fmt.Fprintf(os.Stdout, "Response from `SupportBundleManagementApi.ListSupportBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSupportBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]SupportBundle**](SupportBundle.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportBundleComponents

> []string ListSupportBundleComponents(ctx, cUUID).Execute()

List all components available in support bundle

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
    resp, r, err := api_client.SupportBundleManagementApi.ListSupportBundleComponents(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleManagementApi.ListSupportBundleComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSupportBundleComponents`: []string
    fmt.Fprintf(os.Stdout, "Response from `SupportBundleManagementApi.ListSupportBundleComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSupportBundleComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

