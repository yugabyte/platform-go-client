# \UploadReleasePackagesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUploadRelease**](UploadReleasePackagesApi.md#GetUploadRelease) | **Get** /api/v1/customers/{cUUID}/ybdb_release/upload/{rUUID} | get an uploaded release metadata
[**UploadRelease**](UploadReleasePackagesApi.md#UploadRelease) | **Post** /api/v1/customers/{cUUID}/ybdb_release/upload | upload a release tgz



## GetUploadRelease

> ResponseExtractMetadata GetUploadRelease(ctx, cUUID, rUUID).Request(request).Execute()

get an uploaded release metadata



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
    rUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UploadReleasePackagesApi.GetUploadRelease(context.Background(), cUUID, rUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadReleasePackagesApi.GetUploadRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUploadRelease`: ResponseExtractMetadata
    fmt.Fprintf(os.Stdout, "Response from `UploadReleasePackagesApi.GetUploadRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**rUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUploadReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**ResponseExtractMetadata**](ResponseExtractMetadata.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadRelease

> YBPCreateSuccess UploadRelease(ctx, cUUID).File(file).Request(request).Execute()

upload a release tgz



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
    file := "file_example" // string | Release tgz file to be uploaded
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UploadReleasePackagesApi.UploadRelease(context.Background(), cUUID).File(file).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadReleasePackagesApi.UploadRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadRelease`: YBPCreateSuccess
    fmt.Fprintf(os.Stdout, "Response from `UploadReleasePackagesApi.UploadRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **string** | Release tgz file to be uploaded | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPCreateSuccess**](YBPCreateSuccess.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

