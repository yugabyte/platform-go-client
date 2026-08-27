# \PerfAdvisorEndpointAPI

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePerfAdvisorEndpoint**](PerfAdvisorEndpointAPI.md#CreatePerfAdvisorEndpoint) | **Post** /customers/{cUUID}/perf-advisor-endpoints | Create Perf Advisor Endpoint
[**DeletePerfAdvisorEndpoint**](PerfAdvisorEndpointAPI.md#DeletePerfAdvisorEndpoint) | **Delete** /customers/{cUUID}/perf-advisor-endpoints/{peUUID} | Delete a Perf Advisor Endpoint
[**EditPerfAdvisorEndpoint**](PerfAdvisorEndpointAPI.md#EditPerfAdvisorEndpoint) | **Put** /customers/{cUUID}/perf-advisor-endpoints/{peUUID} | Edit a Perf Advisor Endpoint
[**GetPerfAdvisorEndpoint**](PerfAdvisorEndpointAPI.md#GetPerfAdvisorEndpoint) | **Get** /customers/{cUUID}/perf-advisor-endpoints/{peUUID} | Get a Perf Advisor Endpoint
[**ListPerfAdvisorEndpoints**](PerfAdvisorEndpointAPI.md#ListPerfAdvisorEndpoints) | **Get** /customers/{cUUID}/perf-advisor-endpoints | List Perf Advisor Endpoints
[**ValidatePerfAdvisorEndpoint**](PerfAdvisorEndpointAPI.md#ValidatePerfAdvisorEndpoint) | **Post** /customers/{cUUID}/perf-advisor-endpoints/validate | Validate a Perf Advisor Endpoint



## CreatePerfAdvisorEndpoint

> PerfAdvisorEndpoint CreatePerfAdvisorEndpoint(ctx, cUUID).PerfAdvisorEndpointSpec(perfAdvisorEndpointSpec).Execute()

Create Perf Advisor Endpoint



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
	perfAdvisorEndpointSpec := *openapiclient.NewPerfAdvisorEndpointSpec("byoc-prod", openapiclient.PerfAdvisorEndpointType("BYOC"), "https://byoc.cloud.yugabyte.com/api/v1/otlp/metrics", openapiclient.PerfAdvisorEndpointMetricsType("otlphttp"), "https://byoc.cloud.yugabyte.com") // PerfAdvisorEndpointSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PerfAdvisorEndpointAPI.CreatePerfAdvisorEndpoint(context.Background(), cUUID).PerfAdvisorEndpointSpec(perfAdvisorEndpointSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PerfAdvisorEndpointAPI.CreatePerfAdvisorEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePerfAdvisorEndpoint`: PerfAdvisorEndpoint
	fmt.Fprintf(os.Stdout, "Response from `PerfAdvisorEndpointAPI.CreatePerfAdvisorEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePerfAdvisorEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perfAdvisorEndpointSpec** | [**PerfAdvisorEndpointSpec**](PerfAdvisorEndpointSpec.md) |  | 

### Return type

[**PerfAdvisorEndpoint**](PerfAdvisorEndpoint.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePerfAdvisorEndpoint

> DeletePerfAdvisorEndpoint(ctx, cUUID, peUUID).Execute()

Delete a Perf Advisor Endpoint



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
	peUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Perf Advisor Endpoint UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PerfAdvisorEndpointAPI.DeletePerfAdvisorEndpoint(context.Background(), cUUID, peUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PerfAdvisorEndpointAPI.DeletePerfAdvisorEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**peUUID** | **string** | Perf Advisor Endpoint UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePerfAdvisorEndpointRequest struct via the builder pattern


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


## EditPerfAdvisorEndpoint

> PerfAdvisorEndpoint EditPerfAdvisorEndpoint(ctx, cUUID, peUUID).PerfAdvisorEndpointSpec(perfAdvisorEndpointSpec).Execute()

Edit a Perf Advisor Endpoint



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
	peUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Perf Advisor Endpoint UUID
	perfAdvisorEndpointSpec := *openapiclient.NewPerfAdvisorEndpointSpec("byoc-prod", openapiclient.PerfAdvisorEndpointType("BYOC"), "https://byoc.cloud.yugabyte.com/api/v1/otlp/metrics", openapiclient.PerfAdvisorEndpointMetricsType("otlphttp"), "https://byoc.cloud.yugabyte.com") // PerfAdvisorEndpointSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PerfAdvisorEndpointAPI.EditPerfAdvisorEndpoint(context.Background(), cUUID, peUUID).PerfAdvisorEndpointSpec(perfAdvisorEndpointSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PerfAdvisorEndpointAPI.EditPerfAdvisorEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditPerfAdvisorEndpoint`: PerfAdvisorEndpoint
	fmt.Fprintf(os.Stdout, "Response from `PerfAdvisorEndpointAPI.EditPerfAdvisorEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**peUUID** | **string** | Perf Advisor Endpoint UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditPerfAdvisorEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perfAdvisorEndpointSpec** | [**PerfAdvisorEndpointSpec**](PerfAdvisorEndpointSpec.md) |  | 

### Return type

[**PerfAdvisorEndpoint**](PerfAdvisorEndpoint.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPerfAdvisorEndpoint

> PerfAdvisorEndpoint GetPerfAdvisorEndpoint(ctx, cUUID, peUUID).Execute()

Get a Perf Advisor Endpoint



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
	peUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Perf Advisor Endpoint UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PerfAdvisorEndpointAPI.GetPerfAdvisorEndpoint(context.Background(), cUUID, peUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PerfAdvisorEndpointAPI.GetPerfAdvisorEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPerfAdvisorEndpoint`: PerfAdvisorEndpoint
	fmt.Fprintf(os.Stdout, "Response from `PerfAdvisorEndpointAPI.GetPerfAdvisorEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**peUUID** | **string** | Perf Advisor Endpoint UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPerfAdvisorEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PerfAdvisorEndpoint**](PerfAdvisorEndpoint.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPerfAdvisorEndpoints

> []PerfAdvisorEndpoint ListPerfAdvisorEndpoints(ctx, cUUID).Execute()

List Perf Advisor Endpoints



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
	resp, r, err := apiClient.PerfAdvisorEndpointAPI.ListPerfAdvisorEndpoints(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PerfAdvisorEndpointAPI.ListPerfAdvisorEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPerfAdvisorEndpoints`: []PerfAdvisorEndpoint
	fmt.Fprintf(os.Stdout, "Response from `PerfAdvisorEndpointAPI.ListPerfAdvisorEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPerfAdvisorEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PerfAdvisorEndpoint**](PerfAdvisorEndpoint.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatePerfAdvisorEndpoint

> PerfAdvisorEndpointValidationResult ValidatePerfAdvisorEndpoint(ctx, cUUID).PerfAdvisorEndpointSpec(perfAdvisorEndpointSpec).Execute()

Validate a Perf Advisor Endpoint



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
	perfAdvisorEndpointSpec := *openapiclient.NewPerfAdvisorEndpointSpec("byoc-prod", openapiclient.PerfAdvisorEndpointType("BYOC"), "https://byoc.cloud.yugabyte.com/api/v1/otlp/metrics", openapiclient.PerfAdvisorEndpointMetricsType("otlphttp"), "https://byoc.cloud.yugabyte.com") // PerfAdvisorEndpointSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PerfAdvisorEndpointAPI.ValidatePerfAdvisorEndpoint(context.Background(), cUUID).PerfAdvisorEndpointSpec(perfAdvisorEndpointSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PerfAdvisorEndpointAPI.ValidatePerfAdvisorEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidatePerfAdvisorEndpoint`: PerfAdvisorEndpointValidationResult
	fmt.Fprintf(os.Stdout, "Response from `PerfAdvisorEndpointAPI.ValidatePerfAdvisorEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidatePerfAdvisorEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perfAdvisorEndpointSpec** | [**PerfAdvisorEndpointSpec**](PerfAdvisorEndpointSpec.md) |  | 

### Return type

[**PerfAdvisorEndpointValidationResult**](PerfAdvisorEndpointValidationResult.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

