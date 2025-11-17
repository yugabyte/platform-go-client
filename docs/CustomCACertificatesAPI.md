# \CustomCACertificatesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCA**](CustomCACertificatesAPI.md#AddCA) | **Post** /api/v1/customers/{cUUID}/customCAStore | Add a named custom CA certificate
[**DeleteCustomCACertificate**](CustomCACertificatesAPI.md#DeleteCustomCACertificate) | **Delete** /api/v1/customers/{cUUID}/customCAStoreCertificates/{certUUID} | Delete a named custom CA certificate
[**GetAllCustomCaCertificates**](CustomCACertificatesAPI.md#GetAllCustomCaCertificates) | **Get** /api/v1/customers/{cUUID}/customCAStoreCertificates/{certUUID} | Download a custom CA certificates of a customer
[**ListAllCustomCaCertificates**](CustomCACertificatesAPI.md#ListAllCustomCaCertificates) | **Get** /api/v1/customers/{cUUID}/customCAStoreCertificates | List all custom CA certificates of a customer
[**UpdateCA**](CustomCACertificatesAPI.md#UpdateCA) | **Post** /api/v1/customers/{cUUID}/customCAStore/{certUUID} | Update a named custom CA certificate



## AddCA

> string AddCA(ctx, cUUID).X509CACertificate(x509CACertificate).Request(request).Execute()

Add a named custom CA certificate



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
	x509CACertificate := *openapiclient.NewCustomCACertParams("Contents_example", "Name_example") // CustomCACertParams | CA certificate contents in 'X509' format
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomCACertificatesAPI.AddCA(context.Background(), cUUID).X509CACertificate(x509CACertificate).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomCACertificatesAPI.AddCA``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddCA`: string
	fmt.Fprintf(os.Stdout, "Response from `CustomCACertificatesAPI.AddCA`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

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

Delete a named custom CA certificate



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
	certUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomCACertificatesAPI.DeleteCustomCACertificate(context.Background(), cUUID, certUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomCACertificatesAPI.DeleteCustomCACertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCustomCACertificate`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `CustomCACertificatesAPI.DeleteCustomCACertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**certUUID** | **string** |  | 

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

Download a custom CA certificates of a customer



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
	certUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomCACertificatesAPI.GetAllCustomCaCertificates(context.Background(), cUUID, certUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomCACertificatesAPI.GetAllCustomCaCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllCustomCaCertificates`: CustomCaCertificateInfo
	fmt.Fprintf(os.Stdout, "Response from `CustomCACertificatesAPI.GetAllCustomCaCertificates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**certUUID** | **string** |  | 

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

List all custom CA certificates of a customer



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomCACertificatesAPI.ListAllCustomCaCertificates(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomCACertificatesAPI.ListAllCustomCaCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllCustomCaCertificates`: []CustomCaCertificateInfo
	fmt.Fprintf(os.Stdout, "Response from `CustomCACertificatesAPI.ListAllCustomCaCertificates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

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

Update a named custom CA certificate



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
	certUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	x509CACertificate := *openapiclient.NewCustomCACertParams("Contents_example", "Name_example") // CustomCACertParams | CA certificate contents in 'X509' format
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomCACertificatesAPI.UpdateCA(context.Background(), cUUID, certUUID).X509CACertificate(x509CACertificate).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomCACertificatesAPI.UpdateCA``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCA`: string
	fmt.Fprintf(os.Stdout, "Response from `CustomCACertificatesAPI.UpdateCA`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**certUUID** | **string** |  | 

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

