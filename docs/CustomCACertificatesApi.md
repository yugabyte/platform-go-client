# \CustomCACertificatesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCA**](CustomCACertificatesApi.md#AddCA) | **Post** /api/v1/customers/{cUUID}/customCAStore | WARNING: This is a preview API that could change. Add a named custom CA certificate
[**DeleteCustomCACertificate**](CustomCACertificatesApi.md#DeleteCustomCACertificate) | **Delete** /api/v1/customers/{cUUID}/customCAStoreCertificates/{certUUID} | WARNING: This is a preview API that could change. Delete a named custom CA certificate
[**GetAllCustomCaCertificates**](CustomCACertificatesApi.md#GetAllCustomCaCertificates) | **Get** /api/v1/customers/{cUUID}/customCAStoreCertificates/{certUUID} | WARNING: This is a preview API that could change. Download a custom CA certificates of a customer
[**ListAllCustomCaCertificates**](CustomCACertificatesApi.md#ListAllCustomCaCertificates) | **Get** /api/v1/customers/{cUUID}/customCAStoreCertificates | WARNING: This is a preview API that could change. List all custom CA certificates of a customer
[**UpdateCA**](CustomCACertificatesApi.md#UpdateCA) | **Post** /api/v1/customers/{cUUID}/customCAStore/{certUUID} | WARNING: This is a preview API that could change. Update a named custom CA certificate



## AddCA

> string AddCA(ctx, cUUID).X509CACertificate(x509CACertificate).Request(request).Execute()

WARNING: This is a preview API that could change. Add a named custom CA certificate

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
    x509CACertificate := *openapiclient.NewCustomCACertParams("Contents_example", "Name_example") // CustomCACertParams | CA certificate contents in 'X509' format
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomCACertificatesApi.AddCA(context.Background(), cUUID).X509CACertificate(x509CACertificate).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomCACertificatesApi.AddCA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCA`: string
    fmt.Fprintf(os.Stdout, "Response from `CustomCACertificatesApi.AddCA`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **x509CACertificate** | [**CustomCACertParams**](CustomCACertParams.md) | CA certificate contents in &#39;X509&#39; format | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

**string**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomCACertificate

> YBPSuccess DeleteCustomCACertificate(ctx, cUUID, certUUID).Request(request).Execute()

WARNING: This is a preview API that could change. Delete a named custom CA certificate

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
    certUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomCACertificatesApi.DeleteCustomCACertificate(context.Background(), cUUID, certUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomCACertificatesApi.DeleteCustomCACertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCustomCACertificate`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `CustomCACertificatesApi.DeleteCustomCACertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**certUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomCACertificateRequest struct via the builder pattern


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


## GetAllCustomCaCertificates

> CustomCaCertificateInfo GetAllCustomCaCertificates(ctx, cUUID, certUUID).Request(request).Execute()

WARNING: This is a preview API that could change. Download a custom CA certificates of a customer

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
    certUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomCACertificatesApi.GetAllCustomCaCertificates(context.Background(), cUUID, certUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomCACertificatesApi.GetAllCustomCaCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllCustomCaCertificates`: CustomCaCertificateInfo
    fmt.Fprintf(os.Stdout, "Response from `CustomCACertificatesApi.GetAllCustomCaCertificates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**certUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCustomCaCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**CustomCaCertificateInfo**](CustomCaCertificateInfo.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllCustomCaCertificates

> []CustomCaCertificateInfo ListAllCustomCaCertificates(ctx, cUUID).Execute()

WARNING: This is a preview API that could change. List all custom CA certificates of a customer

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
    resp, r, err := api_client.CustomCACertificatesApi.ListAllCustomCaCertificates(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomCACertificatesApi.ListAllCustomCaCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllCustomCaCertificates`: []CustomCaCertificateInfo
    fmt.Fprintf(os.Stdout, "Response from `CustomCACertificatesApi.ListAllCustomCaCertificates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllCustomCaCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CustomCaCertificateInfo**](CustomCaCertificateInfo.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCA

> string UpdateCA(ctx, cUUID, certUUID).X509CACertificate(x509CACertificate).Request(request).Execute()

WARNING: This is a preview API that could change. Update a named custom CA certificate

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
    certUUID := TODO // string | 
    x509CACertificate := *openapiclient.NewCustomCACertParams("Contents_example", "Name_example") // CustomCACertParams | CA certificate contents in 'X509' format
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomCACertificatesApi.UpdateCA(context.Background(), cUUID, certUUID).X509CACertificate(x509CACertificate).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomCACertificatesApi.UpdateCA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCA`: string
    fmt.Fprintf(os.Stdout, "Response from `CustomCACertificatesApi.UpdateCA`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**certUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **x509CACertificate** | [**CustomCACertParams**](CustomCACertParams.md) | CA certificate contents in &#39;X509&#39; format | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

**string**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

