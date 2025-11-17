# \TroubleshootingPlatformAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckRegistered**](TroubleshootingPlatformAPI.md#CheckRegistered) | **Get** /api/v1/customers/{cUUID}/troubleshooting_platform/{tpUUID}/universes/{uUUID} | Check if universe is registered with Troubleshooting Platform
[**CreateTroubleshootingPlatform**](TroubleshootingPlatformAPI.md#CreateTroubleshootingPlatform) | **Post** /api/v1/customers/{cUUID}/troubleshooting_platform | Create Troubleshooting Platform
[**DeleteTroubleshootingPlatform**](TroubleshootingPlatformAPI.md#DeleteTroubleshootingPlatform) | **Delete** /api/v1/customers/{cUUID}/troubleshooting_platform/{tpUUID} | Delete Troubleshooting Platform
[**EditTroubleshootingPlatform**](TroubleshootingPlatformAPI.md#EditTroubleshootingPlatform) | **Put** /api/v1/customers/{cUUID}/troubleshooting_platform/{tpUUID} | Edit Troubleshooting Platform
[**GetTroubleshootingPlatform**](TroubleshootingPlatformAPI.md#GetTroubleshootingPlatform) | **Get** /api/v1/customers/{cUUID}/troubleshooting_platform/{tpUUID} | Get Troubleshooting Platform
[**ListAllTroubleshootingPlatforms**](TroubleshootingPlatformAPI.md#ListAllTroubleshootingPlatforms) | **Get** /api/v1/customers/{cUUID}/troubleshooting_platform | List All Troubleshooting Platforms
[**RegisterUniverse**](TroubleshootingPlatformAPI.md#RegisterUniverse) | **Put** /api/v1/customers/{cUUID}/troubleshooting_platform/{tpUUID}/universes/{uUUID} | Register universe with Troubleshooting Platform
[**UnregisterUniverse**](TroubleshootingPlatformAPI.md#UnregisterUniverse) | **Delete** /api/v1/customers/{cUUID}/troubleshooting_platform/{tpUUID}/universes/{uUUID} | Unregister universe from Troubleshooting Platform



## CheckRegistered

> YBPSuccess CheckRegistered(ctx, cUUID, tpUUID, uUUID).Request(request).Execute()

Check if universe is registered with Troubleshooting Platform



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
	tpUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TroubleshootingPlatformAPI.CheckRegistered(context.Background(), cUUID, tpUUID, uUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingPlatformAPI.CheckRegistered``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckRegistered`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `TroubleshootingPlatformAPI.CheckRegistered`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**tpUUID** | **string** |  | 
**uUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckRegisteredRequest struct via the builder pattern


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


## CreateTroubleshootingPlatform

> TroubleshootingPlatform CreateTroubleshootingPlatform(ctx, cUUID).PlatformData(platformData).Request(request).Execute()

Create Troubleshooting Platform



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
	platformData := *openapiclient.NewTroubleshootingPlatform("ApiToken_example", "CustomerUUID_example", int64(123), "MetricsUrl_example", "TpUrl_example", "Uuid_example", "YbaUrl_example") // TroubleshootingPlatform | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TroubleshootingPlatformAPI.CreateTroubleshootingPlatform(context.Background(), cUUID).PlatformData(platformData).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingPlatformAPI.CreateTroubleshootingPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTroubleshootingPlatform`: TroubleshootingPlatform
	fmt.Fprintf(os.Stdout, "Response from `TroubleshootingPlatformAPI.CreateTroubleshootingPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTroubleshootingPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **platformData** | [**TroubleshootingPlatform**](TroubleshootingPlatform.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**TroubleshootingPlatform**](TroubleshootingPlatform.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTroubleshootingPlatform

> bool DeleteTroubleshootingPlatform(ctx, cUUID, tpUUID).Force(force).Request(request).Execute()

Delete Troubleshooting Platform



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
	tpUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	force := true // bool |  (optional) (default to false)
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TroubleshootingPlatformAPI.DeleteTroubleshootingPlatform(context.Background(), cUUID, tpUUID).Force(force).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingPlatformAPI.DeleteTroubleshootingPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTroubleshootingPlatform`: bool
	fmt.Fprintf(os.Stdout, "Response from `TroubleshootingPlatformAPI.DeleteTroubleshootingPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**tpUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTroubleshootingPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** |  | [default to false]
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

**bool**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditTroubleshootingPlatform

> TroubleshootingPlatform EditTroubleshootingPlatform(ctx, cUUID, tpUUID).PlatformData(platformData).Force(force).Request(request).Execute()

Edit Troubleshooting Platform



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
	tpUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	platformData := *openapiclient.NewTroubleshootingPlatform("ApiToken_example", "CustomerUUID_example", int64(123), "MetricsUrl_example", "TpUrl_example", "Uuid_example", "YbaUrl_example") // TroubleshootingPlatform | 
	force := true // bool |  (optional) (default to false)
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TroubleshootingPlatformAPI.EditTroubleshootingPlatform(context.Background(), cUUID, tpUUID).PlatformData(platformData).Force(force).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingPlatformAPI.EditTroubleshootingPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditTroubleshootingPlatform`: TroubleshootingPlatform
	fmt.Fprintf(os.Stdout, "Response from `TroubleshootingPlatformAPI.EditTroubleshootingPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**tpUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditTroubleshootingPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **platformData** | [**TroubleshootingPlatform**](TroubleshootingPlatform.md) |  | 
 **force** | **bool** |  | [default to false]
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**TroubleshootingPlatform**](TroubleshootingPlatform.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTroubleshootingPlatform

> TroubleshootingPlatformDetailsModel GetTroubleshootingPlatform(ctx, cUUID, tpUUID).Execute()

Get Troubleshooting Platform



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
	tpUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TroubleshootingPlatformAPI.GetTroubleshootingPlatform(context.Background(), cUUID, tpUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingPlatformAPI.GetTroubleshootingPlatform``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTroubleshootingPlatform`: TroubleshootingPlatformDetailsModel
	fmt.Fprintf(os.Stdout, "Response from `TroubleshootingPlatformAPI.GetTroubleshootingPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**tpUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTroubleshootingPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TroubleshootingPlatformDetailsModel**](TroubleshootingPlatformDetailsModel.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllTroubleshootingPlatforms

> []TroubleshootingPlatformDetailsModel ListAllTroubleshootingPlatforms(ctx, cUUID).Execute()

List All Troubleshooting Platforms



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
	resp, r, err := apiClient.TroubleshootingPlatformAPI.ListAllTroubleshootingPlatforms(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingPlatformAPI.ListAllTroubleshootingPlatforms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllTroubleshootingPlatforms`: []TroubleshootingPlatformDetailsModel
	fmt.Fprintf(os.Stdout, "Response from `TroubleshootingPlatformAPI.ListAllTroubleshootingPlatforms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllTroubleshootingPlatformsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TroubleshootingPlatformDetailsModel**](TroubleshootingPlatformDetailsModel.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterUniverse

> YBPSuccess RegisterUniverse(ctx, cUUID, tpUUID, uUUID).Request(request).Execute()

Register universe with Troubleshooting Platform



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
	tpUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TroubleshootingPlatformAPI.RegisterUniverse(context.Background(), cUUID, tpUUID, uUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingPlatformAPI.RegisterUniverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterUniverse`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `TroubleshootingPlatformAPI.RegisterUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**tpUUID** | **string** |  | 
**uUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterUniverseRequest struct via the builder pattern


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


## UnregisterUniverse

> YBPSuccess UnregisterUniverse(ctx, cUUID, tpUUID, uUUID).Request(request).Execute()

Unregister universe from Troubleshooting Platform



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
	tpUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TroubleshootingPlatformAPI.UnregisterUniverse(context.Background(), cUUID, tpUUID, uUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingPlatformAPI.UnregisterUniverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnregisterUniverse`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `TroubleshootingPlatformAPI.UnregisterUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**tpUUID** | **string** |  | 
**uUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterUniverseRequest struct via the builder pattern


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

