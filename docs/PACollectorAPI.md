# \PACollectorAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckRegistered**](PACollectorAPI.md#CheckRegistered) | **Get** /api/v1/customers/{cUUID}/universes/{uUUID}/pa_collector | Check if universe is registered with PA Collector
[**CreatePACollector**](PACollectorAPI.md#CreatePACollector) | **Post** /api/v1/customers/{cUUID}/pa_collector | Create PA Collector
[**DeletePACollector**](PACollectorAPI.md#DeletePACollector) | **Delete** /api/v1/customers/{cUUID}/pa_collector/{paUUID} | Delete PA Collector
[**EditPACollector**](PACollectorAPI.md#EditPACollector) | **Put** /api/v1/customers/{cUUID}/pa_collector/{paUUID} | Edit PA Collector
[**GetPACollector**](PACollectorAPI.md#GetPACollector) | **Get** /api/v1/customers/{cUUID}/pa_collector/{paUUID} | Get PA Collector
[**ListAllPACollectors**](PACollectorAPI.md#ListAllPACollectors) | **Get** /api/v1/customers/{cUUID}/pa_collector | List All PA Collectors
[**RegisterUniverse**](PACollectorAPI.md#RegisterUniverse) | **Put** /api/v1/customers/{cUUID}/universes/{uUUID}/pa_collector/{paUUID} | Register universe with PA Collector
[**UnregisterUniverse**](PACollectorAPI.md#UnregisterUniverse) | **Delete** /api/v1/customers/{cUUID}/universes/{uUUID}/pa_collector | Unregister universe from PA Collector



## CheckRegistered

> PACollectorUniverseRegistrationStatus CheckRegistered(ctx, cUUID, uUUID).Execute()

Check if universe is registered with PA Collector



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
	uUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PACollectorAPI.CheckRegistered(context.Background(), cUUID, uUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PACollectorAPI.CheckRegistered``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckRegistered`: PACollectorUniverseRegistrationStatus
	fmt.Fprintf(os.Stdout, "Response from `PACollectorAPI.CheckRegistered`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckRegisteredRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PACollectorUniverseRegistrationStatus**](PACollectorUniverseRegistrationStatus.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePACollector

> PACollector CreatePACollector(ctx, cUUID).CollectorData(collectorData).Request(request).Execute()

Create PA Collector



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
	collectorData := *openapiclient.NewPACollector("ApiToken_example", "CustomerUUID_example", int64(123), "MetricsUrl_example", "PaUrl_example", "Uuid_example", "YbaUrl_example") // PACollector | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PACollectorAPI.CreatePACollector(context.Background(), cUUID).CollectorData(collectorData).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PACollectorAPI.CreatePACollector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePACollector`: PACollector
	fmt.Fprintf(os.Stdout, "Response from `PACollectorAPI.CreatePACollector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePACollectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **collectorData** | [**PACollector**](PACollector.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**PACollector**](PACollector.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePACollector

> bool DeletePACollector(ctx, cUUID, paUUID).Force(force).Request(request).Execute()

Delete PA Collector



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
	paUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	force := true // bool |  (optional) (default to false)
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PACollectorAPI.DeletePACollector(context.Background(), cUUID, paUUID).Force(force).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PACollectorAPI.DeletePACollector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePACollector`: bool
	fmt.Fprintf(os.Stdout, "Response from `PACollectorAPI.DeletePACollector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**paUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePACollectorRequest struct via the builder pattern


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


## EditPACollector

> PACollector EditPACollector(ctx, cUUID, paUUID).CollectorData(collectorData).Force(force).Request(request).Execute()

Edit PA Collector



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
	paUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	collectorData := *openapiclient.NewPACollector("ApiToken_example", "CustomerUUID_example", int64(123), "MetricsUrl_example", "PaUrl_example", "Uuid_example", "YbaUrl_example") // PACollector | 
	force := true // bool |  (optional) (default to false)
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PACollectorAPI.EditPACollector(context.Background(), cUUID, paUUID).CollectorData(collectorData).Force(force).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PACollectorAPI.EditPACollector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditPACollector`: PACollector
	fmt.Fprintf(os.Stdout, "Response from `PACollectorAPI.EditPACollector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**paUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditPACollectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **collectorData** | [**PACollector**](PACollector.md) |  | 
 **force** | **bool** |  | [default to false]
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**PACollector**](PACollector.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPACollector

> PACollectorDetailsModel GetPACollector(ctx, cUUID, paUUID).Execute()

Get PA Collector



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
	paUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PACollectorAPI.GetPACollector(context.Background(), cUUID, paUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PACollectorAPI.GetPACollector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPACollector`: PACollectorDetailsModel
	fmt.Fprintf(os.Stdout, "Response from `PACollectorAPI.GetPACollector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**paUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPACollectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PACollectorDetailsModel**](PACollectorDetailsModel.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllPACollectors

> []PACollectorDetailsModel ListAllPACollectors(ctx, cUUID).Execute()

List All PA Collectors



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
	resp, r, err := apiClient.PACollectorAPI.ListAllPACollectors(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PACollectorAPI.ListAllPACollectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllPACollectors`: []PACollectorDetailsModel
	fmt.Fprintf(os.Stdout, "Response from `PACollectorAPI.ListAllPACollectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllPACollectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PACollectorDetailsModel**](PACollectorDetailsModel.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterUniverse

> YBPSuccess RegisterUniverse(ctx, cUUID, uUUID, paUUID).AdvancedObservability(advancedObservability).Request(request).Execute()

Register universe with PA Collector



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
	uUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	paUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	advancedObservability := true // bool |  (optional) (default to false)
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PACollectorAPI.RegisterUniverse(context.Background(), cUUID, uUUID, paUUID).AdvancedObservability(advancedObservability).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PACollectorAPI.RegisterUniverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterUniverse`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `PACollectorAPI.RegisterUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uUUID** | **string** |  | 
**paUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterUniverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **advancedObservability** | **bool** |  | [default to false]
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

> YBPSuccess UnregisterUniverse(ctx, cUUID, uUUID).Request(request).Execute()

Unregister universe from PA Collector



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
	uUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PACollectorAPI.UnregisterUniverse(context.Background(), cUUID, uUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PACollectorAPI.UnregisterUniverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnregisterUniverse`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `PACollectorAPI.UnregisterUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
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

