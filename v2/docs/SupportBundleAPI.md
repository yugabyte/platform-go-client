# \SupportBundleAPI

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSupportBundle**](SupportBundleAPI.md#CreateSupportBundle) | **Post** /customers/{cUUID}/universes/{uniUUID}/support-bundles | Create support bundle
[**CreateYbaSupportBundle**](SupportBundleAPI.md#CreateYbaSupportBundle) | **Post** /customers/{cUUID}/support-bundles/yba | Create YBA-only support bundle
[**DeleteSupportBundle**](SupportBundleAPI.md#DeleteSupportBundle) | **Delete** /customers/{cUUID}/universes/{uniUUID}/support-bundles/{sbUUID} | Delete support bundle
[**DeleteYbaSupportBundle**](SupportBundleAPI.md#DeleteYbaSupportBundle) | **Delete** /customers/{cUUID}/support-bundles/yba/{sbUUID} | Delete YBA-only support bundle
[**DownloadSupportBundle**](SupportBundleAPI.md#DownloadSupportBundle) | **Get** /customers/{cUUID}/universes/{uniUUID}/support-bundles/{sbUUID}/download | Download support bundle
[**DownloadYbaSupportBundle**](SupportBundleAPI.md#DownloadYbaSupportBundle) | **Get** /customers/{cUUID}/support-bundles/yba/{sbUUID}/download | Download YBA-only support bundle
[**EstimateSupportBundleSize**](SupportBundleAPI.md#EstimateSupportBundleSize) | **Post** /customers/{cUUID}/universes/{uniUUID}/support-bundles/estimate-size | Estimate support bundle size
[**EstimateYbaSupportBundleSize**](SupportBundleAPI.md#EstimateYbaSupportBundleSize) | **Post** /customers/{cUUID}/support-bundles/yba/estimate-size | Estimate YBA-only support bundle size
[**GetSupportBundle**](SupportBundleAPI.md#GetSupportBundle) | **Get** /customers/{cUUID}/universes/{uniUUID}/support-bundles/{sbUUID} | Get support bundle
[**GetYbaSupportBundle**](SupportBundleAPI.md#GetYbaSupportBundle) | **Get** /customers/{cUUID}/support-bundles/yba/{sbUUID} | Get YBA-only support bundle
[**ListSupportBundleComponents**](SupportBundleAPI.md#ListSupportBundleComponents) | **Get** /customers/{cUUID}/support-bundle/components | List support bundle components
[**PageListSupportBundles**](SupportBundleAPI.md#PageListSupportBundles) | **Post** /customers/{cUUID}/universes/{uniUUID}/support-bundles/page | List support bundles (paged)
[**PageListYbaSupportBundles**](SupportBundleAPI.md#PageListYbaSupportBundles) | **Post** /customers/{cUUID}/support-bundles/yba/page | List YBA-only support bundles (paged)



## CreateSupportBundle

> YBATask CreateSupportBundle(ctx, cUUID, uniUUID).SupportBundleCreateSpec(supportBundleCreateSpec).Execute()

Create support bundle



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
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	supportBundleCreateSpec := *openapiclient.NewSupportBundleCreateSpec([]openapiclient.SupportBundleComponentType{openapiclient.SupportBundleComponentType("UniverseLogs")}) // SupportBundleCreateSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportBundleAPI.CreateSupportBundle(context.Background(), cUUID, uniUUID).SupportBundleCreateSpec(supportBundleCreateSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleAPI.CreateSupportBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSupportBundle`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `SupportBundleAPI.CreateSupportBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSupportBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **supportBundleCreateSpec** | [**SupportBundleCreateSpec**](SupportBundleCreateSpec.md) |  | 

### Return type

[**YBATask**](YBATask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateYbaSupportBundle

> YBATask CreateYbaSupportBundle(ctx, cUUID).SupportBundleCreateSpec(supportBundleCreateSpec).Execute()

Create YBA-only support bundle



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
	supportBundleCreateSpec := *openapiclient.NewSupportBundleCreateSpec([]openapiclient.SupportBundleComponentType{openapiclient.SupportBundleComponentType("UniverseLogs")}) // SupportBundleCreateSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportBundleAPI.CreateYbaSupportBundle(context.Background(), cUUID).SupportBundleCreateSpec(supportBundleCreateSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleAPI.CreateYbaSupportBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateYbaSupportBundle`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `SupportBundleAPI.CreateYbaSupportBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateYbaSupportBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportBundleCreateSpec** | [**SupportBundleCreateSpec**](SupportBundleCreateSpec.md) |  | 

### Return type

[**YBATask**](YBATask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSupportBundle

> DeleteSupportBundle(ctx, cUUID, uniUUID, sbUUID).Execute()

Delete support bundle



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
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	sbUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Support bundle UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SupportBundleAPI.DeleteSupportBundle(context.Background(), cUUID, uniUUID, sbUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleAPI.DeleteSupportBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 
**sbUUID** | **string** | Support bundle UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSupportBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteYbaSupportBundle

> DeleteYbaSupportBundle(ctx, cUUID, sbUUID).Execute()

Delete YBA-only support bundle



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
	sbUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Support bundle UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SupportBundleAPI.DeleteYbaSupportBundle(context.Background(), cUUID, sbUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleAPI.DeleteYbaSupportBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**sbUUID** | **string** | Support bundle UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteYbaSupportBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadSupportBundle

> *os.File DownloadSupportBundle(ctx, cUUID, uniUUID, sbUUID).Execute()

Download support bundle



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
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	sbUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Support bundle UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportBundleAPI.DownloadSupportBundle(context.Background(), cUUID, uniUUID, sbUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleAPI.DownloadSupportBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadSupportBundle`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SupportBundleAPI.DownloadSupportBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 
**sbUUID** | **string** | Support bundle UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadSupportBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[***os.File**](*os.File.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-compressed

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadYbaSupportBundle

> *os.File DownloadYbaSupportBundle(ctx, cUUID, sbUUID).Execute()

Download YBA-only support bundle



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
	sbUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Support bundle UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportBundleAPI.DownloadYbaSupportBundle(context.Background(), cUUID, sbUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleAPI.DownloadYbaSupportBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadYbaSupportBundle`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SupportBundleAPI.DownloadYbaSupportBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**sbUUID** | **string** | Support bundle UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadYbaSupportBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-compressed

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EstimateSupportBundleSize

> SupportBundleSizeEstimateResponse EstimateSupportBundleSize(ctx, cUUID, uniUUID).SupportBundleCreateSpec(supportBundleCreateSpec).Execute()

Estimate support bundle size



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
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	supportBundleCreateSpec := *openapiclient.NewSupportBundleCreateSpec([]openapiclient.SupportBundleComponentType{openapiclient.SupportBundleComponentType("UniverseLogs")}) // SupportBundleCreateSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportBundleAPI.EstimateSupportBundleSize(context.Background(), cUUID, uniUUID).SupportBundleCreateSpec(supportBundleCreateSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleAPI.EstimateSupportBundleSize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EstimateSupportBundleSize`: SupportBundleSizeEstimateResponse
	fmt.Fprintf(os.Stdout, "Response from `SupportBundleAPI.EstimateSupportBundleSize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEstimateSupportBundleSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **supportBundleCreateSpec** | [**SupportBundleCreateSpec**](SupportBundleCreateSpec.md) |  | 

### Return type

[**SupportBundleSizeEstimateResponse**](SupportBundleSizeEstimateResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EstimateYbaSupportBundleSize

> SupportBundleSizeEstimateResponse EstimateYbaSupportBundleSize(ctx, cUUID).SupportBundleCreateSpec(supportBundleCreateSpec).Execute()

Estimate YBA-only support bundle size



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
	supportBundleCreateSpec := *openapiclient.NewSupportBundleCreateSpec([]openapiclient.SupportBundleComponentType{openapiclient.SupportBundleComponentType("UniverseLogs")}) // SupportBundleCreateSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportBundleAPI.EstimateYbaSupportBundleSize(context.Background(), cUUID).SupportBundleCreateSpec(supportBundleCreateSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleAPI.EstimateYbaSupportBundleSize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EstimateYbaSupportBundleSize`: SupportBundleSizeEstimateResponse
	fmt.Fprintf(os.Stdout, "Response from `SupportBundleAPI.EstimateYbaSupportBundleSize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEstimateYbaSupportBundleSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportBundleCreateSpec** | [**SupportBundleCreateSpec**](SupportBundleCreateSpec.md) |  | 

### Return type

[**SupportBundleSizeEstimateResponse**](SupportBundleSizeEstimateResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportBundle

> SupportBundle GetSupportBundle(ctx, cUUID, uniUUID, sbUUID).Execute()

Get support bundle



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
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	sbUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Support bundle UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportBundleAPI.GetSupportBundle(context.Background(), cUUID, uniUUID, sbUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleAPI.GetSupportBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportBundle`: SupportBundle
	fmt.Fprintf(os.Stdout, "Response from `SupportBundleAPI.GetSupportBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 
**sbUUID** | **string** | Support bundle UUID | 

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


## GetYbaSupportBundle

> SupportBundle GetYbaSupportBundle(ctx, cUUID, sbUUID).Execute()

Get YBA-only support bundle



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
	sbUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Support bundle UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportBundleAPI.GetYbaSupportBundle(context.Background(), cUUID, sbUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleAPI.GetYbaSupportBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetYbaSupportBundle`: SupportBundle
	fmt.Fprintf(os.Stdout, "Response from `SupportBundleAPI.GetYbaSupportBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**sbUUID** | **string** | Support bundle UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetYbaSupportBundleRequest struct via the builder pattern


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


## ListSupportBundleComponents

> []SupportBundleComponentType ListSupportBundleComponents(ctx, cUUID).Execute()

List support bundle components



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportBundleAPI.ListSupportBundleComponents(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleAPI.ListSupportBundleComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSupportBundleComponents`: []SupportBundleComponentType
	fmt.Fprintf(os.Stdout, "Response from `SupportBundleAPI.ListSupportBundleComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSupportBundleComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SupportBundleComponentType**](SupportBundleComponentType.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageListSupportBundles

> SupportBundlePagedResp PageListSupportBundles(ctx, cUUID, uniUUID).SupportBundlePagedQuerySpec(supportBundlePagedQuerySpec).Execute()

List support bundles (paged)



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
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	supportBundlePagedQuerySpec := *openapiclient.NewSupportBundlePagedQuerySpec() // SupportBundlePagedQuerySpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportBundleAPI.PageListSupportBundles(context.Background(), cUUID, uniUUID).SupportBundlePagedQuerySpec(supportBundlePagedQuerySpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleAPI.PageListSupportBundles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageListSupportBundles`: SupportBundlePagedResp
	fmt.Fprintf(os.Stdout, "Response from `SupportBundleAPI.PageListSupportBundles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageListSupportBundlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **supportBundlePagedQuerySpec** | [**SupportBundlePagedQuerySpec**](SupportBundlePagedQuerySpec.md) |  | 

### Return type

[**SupportBundlePagedResp**](SupportBundlePagedResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageListYbaSupportBundles

> SupportBundlePagedResp PageListYbaSupportBundles(ctx, cUUID).SupportBundlePagedQuerySpec(supportBundlePagedQuerySpec).Execute()

List YBA-only support bundles (paged)



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
	supportBundlePagedQuerySpec := *openapiclient.NewSupportBundlePagedQuerySpec() // SupportBundlePagedQuerySpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportBundleAPI.PageListYbaSupportBundles(context.Background(), cUUID).SupportBundlePagedQuerySpec(supportBundlePagedQuerySpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleAPI.PageListYbaSupportBundles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageListYbaSupportBundles`: SupportBundlePagedResp
	fmt.Fprintf(os.Stdout, "Response from `SupportBundleAPI.PageListYbaSupportBundles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageListYbaSupportBundlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportBundlePagedQuerySpec** | [**SupportBundlePagedQuerySpec**](SupportBundlePagedQuerySpec.md) |  | 

### Return type

[**SupportBundlePagedResp**](SupportBundlePagedResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

