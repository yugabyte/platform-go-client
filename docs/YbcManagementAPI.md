# \YbcManagementAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableYbc**](YbcManagementAPI.md#DisableYbc) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/ybc/disable | Disable YBC on the universe nodes
[**InstallYbc**](YbcManagementAPI.md#InstallYbc) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/ybc/install | Install YBC on the universe nodes
[**SetThrottleParams**](YbcManagementAPI.md#SetThrottleParams) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/ybc_throttle_params_async | Set throttle params in YB-Controller( async )
[**UpgradeYbc**](YbcManagementAPI.md#UpgradeYbc) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/ybc/upgrade | Upgrade YBC on the universe nodes
[**UpgradeYbcGflags**](YbcManagementAPI.md#UpgradeYbcGflags) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/ybc/upgrade/gflags | Upgrade YBC gflags on the universe nodes



## DisableYbc

> YBPTask DisableYbc(ctx, cUUID, uniUUID).Request(request).Execute()

Disable YBC on the universe nodes



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
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.YbcManagementAPI.DisableYbc(context.Background(), cUUID, uniUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `YbcManagementAPI.DisableYbc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableYbc`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `YbcManagementAPI.DisableYbc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableYbcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallYbc

> YBPTask InstallYbc(ctx, cUUID, uniUUID).YbcVersion(ybcVersion).Request(request).Execute()

Install YBC on the universe nodes



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
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	ybcVersion := "ybcVersion_example" // string |  (optional)
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.YbcManagementAPI.InstallYbc(context.Background(), cUUID, uniUUID).YbcVersion(ybcVersion).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `YbcManagementAPI.InstallYbc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstallYbc`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `YbcManagementAPI.InstallYbc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallYbcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ybcVersion** | **string** |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetThrottleParams

> YBPSuccess SetThrottleParams(ctx, cUUID, uniUUID).ThrottleParams(throttleParams).Request(request).Execute()

Set throttle params in YB-Controller( async )

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
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	throttleParams := *openapiclient.NewYbcThrottleParameters() // YbcThrottleParameters | Parameters for YB-Controller throttling
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.YbcManagementAPI.SetThrottleParams(context.Background(), cUUID, uniUUID).ThrottleParams(throttleParams).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `YbcManagementAPI.SetThrottleParams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetThrottleParams`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `YbcManagementAPI.SetThrottleParams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetThrottleParamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **throttleParams** | [**YbcThrottleParameters**](YbcThrottleParameters.md) | Parameters for YB-Controller throttling | 
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


## UpgradeYbc

> YBPTask UpgradeYbc(ctx, cUUID, uniUUID).YbcVersion(ybcVersion).Request(request).Execute()

Upgrade YBC on the universe nodes



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
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	ybcVersion := "ybcVersion_example" // string |  (optional)
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.YbcManagementAPI.UpgradeYbc(context.Background(), cUUID, uniUUID).YbcVersion(ybcVersion).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `YbcManagementAPI.UpgradeYbc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeYbc`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `YbcManagementAPI.UpgradeYbc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeYbcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ybcVersion** | **string** |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeYbcGflags

> YBPTask UpgradeYbcGflags(ctx, cUUID, uniUUID).Request(request).Execute()

Upgrade YBC gflags on the universe nodes



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
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.YbcManagementAPI.UpgradeYbcGflags(context.Background(), cUUID, uniUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `YbcManagementAPI.UpgradeYbcGflags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeYbcGflags`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `YbcManagementAPI.UpgradeYbcGflags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeYbcGflagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

