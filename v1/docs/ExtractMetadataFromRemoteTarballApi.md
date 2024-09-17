# \ExtractMetadataFromRemoteTarballApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExtractMetadata**](ExtractMetadataFromRemoteTarballApi.md#ExtractMetadata) | **Post** /api/v1/customers/{cUUID}/ybdb_release/extract_metadata | helper to extract release metadata from a remote tarball
[**ExtractMetadata_0**](ExtractMetadataFromRemoteTarballApi.md#ExtractMetadata_0) | **Get** /api/v1/customers/{cUUID}/ybdb_release/extract_metadata/{rUUID} | get the extract release metadata from a remote tarball



## ExtractMetadata

> YBPSuccess ExtractMetadata(ctx, cUUID).ReleaseURL(releaseURL).Request(request).Execute()

helper to extract release metadata from a remote tarball



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
    releaseURL := *openapiclient.NewExtractMetadata("Url_example", "Uuid_example") // ExtractMetadata | Release URL to extract metadata from
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExtractMetadataFromRemoteTarballApi.ExtractMetadata(context.Background(), cUUID).ReleaseURL(releaseURL).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtractMetadataFromRemoteTarballApi.ExtractMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtractMetadata`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `ExtractMetadataFromRemoteTarballApi.ExtractMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtractMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **releaseURL** | [**ExtractMetadata**](ExtractMetadata.md) | Release URL to extract metadata from | 
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


## ExtractMetadata_0

> ResponseExtractMetadata ExtractMetadata_0(ctx, cUUID, rUUID).Request(request).Execute()

get the extract release metadata from a remote tarball



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
    resp, r, err := api_client.ExtractMetadataFromRemoteTarballApi.ExtractMetadata_0(context.Background(), cUUID, rUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtractMetadataFromRemoteTarballApi.ExtractMetadata_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtractMetadata_0`: ResponseExtractMetadata
    fmt.Fprintf(os.Stdout, "Response from `ExtractMetadataFromRemoteTarballApi.ExtractMetadata_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**rUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtractMetadata_1Request struct via the builder pattern


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

