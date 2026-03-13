# \InstanceTypesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInstanceType**](InstanceTypesAPI.md#CreateInstanceType) | **Post** /api/v1/customers/{cUUID}/providers/{pUUID}/instance_types | Create an instance type
[**DeleteInstanceType**](InstanceTypesAPI.md#DeleteInstanceType) | **Delete** /api/v1/customers/{cUUID}/providers/{pUUID}/instance_types/{code} | Delete an instance type
[**GetAZUTypes**](InstanceTypesAPI.md#GetAZUTypes) | **Get** /api/v1/metadata/azu_types | List supported Azure disk types
[**GetEBSTypes**](InstanceTypesAPI.md#GetEBSTypes) | **Get** /api/v1/metadata/ebs_types | List supported EBS volume types
[**GetGCPTypes**](InstanceTypesAPI.md#GetGCPTypes) | **Get** /api/v1/metadata/gcp_types | List supported GCP disk types
[**InstanceTypeDetail**](InstanceTypesAPI.md#InstanceTypeDetail) | **Get** /api/v1/customers/{cUUID}/providers/{pUUID}/instance_types/{code} | Get details of an instance type
[**ListOfInstanceType**](InstanceTypesAPI.md#ListOfInstanceType) | **Get** /api/v1/customers/{cUUID}/providers/{pUUID}/instance_types | List a provider&#39;s instance types



## CreateInstanceType

> InstanceTypeResp CreateInstanceType(ctx, cUUID, pUUID).InstanceType(instanceType).Request(request).Execute()

Create an instance type

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
	instanceType := *openapiclient.NewInstanceType(*openapiclient.NewInstanceTypeKey("InstanceTypeCode_example", "ProviderUuid_example")) // InstanceType | Instance type data of the instance to be stored
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceTypesAPI.CreateInstanceType(context.Background(), cUUID, pUUID).InstanceType(instanceType).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceTypesAPI.CreateInstanceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInstanceType`: InstanceTypeResp
	fmt.Fprintf(os.Stdout, "Response from `InstanceTypesAPI.CreateInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**pUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **instanceType** | [**InstanceType**](InstanceType.md) | Instance type data of the instance to be stored | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**InstanceTypeResp**](InstanceTypeResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstanceType

> YBPSuccess DeleteInstanceType(ctx, cUUID, pUUID, code).Request(request).Execute()

Delete an instance type

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
	code := "code_example" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceTypesAPI.DeleteInstanceType(context.Background(), cUUID, pUUID, code).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceTypesAPI.DeleteInstanceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInstanceType`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `InstanceTypesAPI.DeleteInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**pUUID** | **string** |  | 
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceTypeRequest struct via the builder pattern


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


## GetAZUTypes

> []string GetAZUTypes(ctx).Execute()

List supported Azure disk types

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceTypesAPI.GetAZUTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceTypesAPI.GetAZUTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAZUTypes`: []string
	fmt.Fprintf(os.Stdout, "Response from `InstanceTypesAPI.GetAZUTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAZUTypesRequest struct via the builder pattern


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


## GetEBSTypes

> []string GetEBSTypes(ctx).Execute()

List supported EBS volume types

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceTypesAPI.GetEBSTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceTypesAPI.GetEBSTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEBSTypes`: []string
	fmt.Fprintf(os.Stdout, "Response from `InstanceTypesAPI.GetEBSTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEBSTypesRequest struct via the builder pattern


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


## GetGCPTypes

> []string GetGCPTypes(ctx).Execute()

List supported GCP disk types

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceTypesAPI.GetGCPTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceTypesAPI.GetGCPTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGCPTypes`: []string
	fmt.Fprintf(os.Stdout, "Response from `InstanceTypesAPI.GetGCPTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGCPTypesRequest struct via the builder pattern


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


## InstanceTypeDetail

> InstanceTypeResp InstanceTypeDetail(ctx, cUUID, pUUID, code).Execute()

Get details of an instance type

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
	code := "code_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceTypesAPI.InstanceTypeDetail(context.Background(), cUUID, pUUID, code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceTypesAPI.InstanceTypeDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstanceTypeDetail`: InstanceTypeResp
	fmt.Fprintf(os.Stdout, "Response from `InstanceTypesAPI.InstanceTypeDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**pUUID** | **string** |  | 
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstanceTypeDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**InstanceTypeResp**](InstanceTypeResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOfInstanceType

> []InstanceTypeResp ListOfInstanceType(ctx, cUUID, pUUID).Zone(zone).Arch(arch).Execute()

List a provider's instance types

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
	zone := []string{"Inner_example"} // []string |  (optional)
	arch := "arch_example" // string |  (optional) (default to "null")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceTypesAPI.ListOfInstanceType(context.Background(), cUUID, pUUID).Zone(zone).Arch(arch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceTypesAPI.ListOfInstanceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOfInstanceType`: []InstanceTypeResp
	fmt.Fprintf(os.Stdout, "Response from `InstanceTypesAPI.ListOfInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**pUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOfInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **zone** | **[]string** |  | 
 **arch** | **string** |  | [default to &quot;null&quot;]

### Return type

[**[]InstanceTypeResp**](InstanceTypeResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

