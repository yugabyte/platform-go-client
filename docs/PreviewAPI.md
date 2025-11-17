# \PreviewAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImageBundle**](PreviewAPI.md#CreateImageBundle) | **Post** /api/v1/customers/{cUUID}/providers/{pUUID}/image_bundle | Create a image bundle
[**Delete**](PreviewAPI.md#Delete) | **Delete** /api/v1/customers/{cUUID}/providers/{pUUID}/image_bundle/{iBUUID} | Delete a image bundle
[**EditImageBundle**](PreviewAPI.md#EditImageBundle) | **Put** /api/v1/customers/{cUUID}/providers/{pUUID}/image_bundle/{iBUUID} | Update a image bundle
[**GetImageBundle**](PreviewAPI.md#GetImageBundle) | **Get** /api/v1/customers/{cUUID}/providers/{pUUID}/image_bundle/{iBUUID} | Get a image bundle
[**GetListOfImageBundles**](PreviewAPI.md#GetListOfImageBundles) | **Get** /api/v1/customers/{cUUID}/providers/{pUUID}/image_bundle | List image bundles



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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	body := *openapiclient.NewImageBundle() // ImageBundle | CreateImageBundleRequest
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PreviewAPI.CreateImageBundle(context.Background(), cUUID, pUUID).Body(body).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreviewAPI.CreateImageBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateImageBundle`: ImageBundle
	fmt.Fprintf(os.Stdout, "Response from `PreviewAPI.CreateImageBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**pUUID** | **string** |  | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	iBUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PreviewAPI.Delete(context.Background(), cUUID, pUUID, iBUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreviewAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Delete`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `PreviewAPI.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**pUUID** | **string** |  | 
**iBUUID** | **string** |  | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	iBUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	body := *openapiclient.NewImageBundle() // ImageBundle | EditImageBundleRequest
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PreviewAPI.EditImageBundle(context.Background(), cUUID, pUUID, iBUUID).Body(body).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreviewAPI.EditImageBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditImageBundle`: ImageBundle
	fmt.Fprintf(os.Stdout, "Response from `PreviewAPI.EditImageBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**pUUID** | **string** |  | 
**iBUUID** | **string** |  | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	iBUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PreviewAPI.GetImageBundle(context.Background(), cUUID, pUUID, iBUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreviewAPI.GetImageBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageBundle`: ImageBundle
	fmt.Fprintf(os.Stdout, "Response from `PreviewAPI.GetImageBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**pUUID** | **string** |  | 
**iBUUID** | **string** |  | 

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

> []ImageBundle GetListOfImageBundles(ctx, cUUID, pUUID).Arch(arch).Execute()

List image bundles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	arch := "arch_example" // string |  (optional) (default to "null")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PreviewAPI.GetListOfImageBundles(context.Background(), cUUID, pUUID).Arch(arch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreviewAPI.GetListOfImageBundles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListOfImageBundles`: []ImageBundle
	fmt.Fprintf(os.Stdout, "Response from `PreviewAPI.GetListOfImageBundles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**pUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListOfImageBundlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **arch** | **string** |  | [default to &quot;null&quot;]

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

