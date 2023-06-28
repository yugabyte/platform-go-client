# \PreviewApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImageBundle**](PreviewApi.md#CreateImageBundle) | **Post** /api/v1/customers/{cUUID}/providers/{pUUID}/image_bundle | Create a image bundle
[**Delete**](PreviewApi.md#Delete) | **Delete** /api/v1/customers/{cUUID}/providers/{pUUID}/image_bundle/{iBUUID} | Delete a image bundle
[**EditImageBundle**](PreviewApi.md#EditImageBundle) | **Put** /api/v1/customers/{cUUID}/providers/{pUUID}/image_bundle/{iBUUID} | Update a image bundle
[**GetImageBundle**](PreviewApi.md#GetImageBundle) | **Get** /api/v1/customers/{cUUID}/providers/{pUUID}/image_bundle/{iBUUID} | Get a image bundle
[**GetListOfImageBundles**](PreviewApi.md#GetListOfImageBundles) | **Get** /api/v1/customers/{cUUID}/providers/{pUUID}/image_bundle | List image bundles



## CreateImageBundle

> ImageBundle CreateImageBundle(ctx, cUUID, pUUID).Body(body).Request(request).Execute()

Create a image bundle

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
    body := *openapiclient.NewImageBundle() // ImageBundle | CreateImageBundleRequest
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PreviewApi.CreateImageBundle(context.Background(), cUUID, pUUID).Body(body).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreviewApi.CreateImageBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateImageBundle`: ImageBundle
    fmt.Fprintf(os.Stdout, "Response from `PreviewApi.CreateImageBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateImageBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ImageBundle**](ImageBundle.md) | CreateImageBundleRequest | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**ImageBundle**](ImageBundle.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> YBPSuccess Delete(ctx, cUUID, pUUID, iBUUID).Request(request).Execute()

Delete a image bundle

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
    iBUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PreviewApi.Delete(context.Background(), cUUID, pUUID, iBUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreviewApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `PreviewApi.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 
**iBUUID** | [**string**](.md) |  | 

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


## EditImageBundle

> ImageBundle EditImageBundle(ctx, cUUID, pUUID, iBUUID).Body(body).Request(request).Execute()

Update a image bundle

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
    iBUUID := TODO // string | 
    body := *openapiclient.NewImageBundle() // ImageBundle | EditImageBundleRequest
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PreviewApi.EditImageBundle(context.Background(), cUUID, pUUID, iBUUID).Body(body).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreviewApi.EditImageBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditImageBundle`: ImageBundle
    fmt.Fprintf(os.Stdout, "Response from `PreviewApi.EditImageBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 
**iBUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditImageBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**ImageBundle**](ImageBundle.md) | EditImageBundleRequest | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**ImageBundle**](ImageBundle.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImageBundle

> ImageBundle GetImageBundle(ctx, cUUID, pUUID, iBUUID).Execute()

Get a image bundle

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
    iBUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PreviewApi.GetImageBundle(context.Background(), cUUID, pUUID, iBUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreviewApi.GetImageBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImageBundle`: ImageBundle
    fmt.Fprintf(os.Stdout, "Response from `PreviewApi.GetImageBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 
**iBUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ImageBundle**](ImageBundle.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListOfImageBundles

> []ImageBundle GetListOfImageBundles(ctx, cUUID, pUUID).Execute()

List image bundles

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
    resp, r, err := api_client.PreviewApi.GetListOfImageBundles(context.Background(), cUUID, pUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreviewApi.GetListOfImageBundles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListOfImageBundles`: []ImageBundle
    fmt.Fprintf(os.Stdout, "Response from `PreviewApi.GetListOfImageBundles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListOfImageBundlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ImageBundle**](ImageBundle.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

