# \PackagesControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchPackage**](PackagesControllerApi.md#FetchPackage) | **Post** /api/v1/fetch_package | Fetch a package



## FetchPackage

> string FetchPackage(ctx).Package_(package_).Request(request).Execute()

Fetch a package

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
    package_ := *openapiclient.NewPackagesRequestParams("BuildNumber_example") // PackagesRequestParams | Package to be imported
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PackagesControllerApi.FetchPackage(context.Background()).Package_(package_).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PackagesControllerApi.FetchPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchPackage`: string
    fmt.Fprintf(os.Stdout, "Response from `PackagesControllerApi.FetchPackage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **package_** | [**PackagesRequestParams**](PackagesRequestParams.md) | Package to be imported | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

