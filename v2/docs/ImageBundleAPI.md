# \ImageBundleAPI

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PageListImageBundles**](ImageBundleAPI.md#PageListImageBundles) | **Post** /customers/{cUUID}/providers/{providerUUID}/image-bundles/page | List image bundles (paged)



## PageListImageBundles

> ImageBundlePagedResp PageListImageBundles(ctx, cUUID, providerUUID).ImageBundlePagedQuerySpec(imageBundlePagedQuerySpec).Execute()

List image bundles (paged)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	providerUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Provider UUID
	imageBundlePagedQuerySpec := *openapiclient.NewImageBundlePagedQuerySpec() // ImageBundlePagedQuerySpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageBundleAPI.PageListImageBundles(context.Background(), cUUID, providerUUID).ImageBundlePagedQuerySpec(imageBundlePagedQuerySpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageBundleAPI.PageListImageBundles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageListImageBundles`: ImageBundlePagedResp
	fmt.Fprintf(os.Stdout, "Response from `ImageBundleAPI.PageListImageBundles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**providerUUID** | **string** | Provider UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageListImageBundlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **imageBundlePagedQuerySpec** | [**ImageBundlePagedQuerySpec**](ImageBundlePagedQuerySpec.md) |  | 

### Return type

[**ImageBundlePagedResp**](ImageBundlePagedResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

