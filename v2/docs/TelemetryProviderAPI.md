# \TelemetryProviderAPI

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListTelemetryProviderTypes**](TelemetryProviderAPI.md#ListTelemetryProviderTypes) | **Get** /customers/{cUUID}/telemetry-provider/types | List Available Telemetry Provider Types



## ListTelemetryProviderTypes

> []TelemetryProviderTypeInfo ListTelemetryProviderTypes(ctx, cUUID).ExportType(exportType).Execute()

List Available Telemetry Provider Types



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
	exportType := "exportType_example" // string | Filter by export type (logs or metrics) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TelemetryProviderAPI.ListTelemetryProviderTypes(context.Background(), cUUID).ExportType(exportType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TelemetryProviderAPI.ListTelemetryProviderTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTelemetryProviderTypes`: []TelemetryProviderTypeInfo
	fmt.Fprintf(os.Stdout, "Response from `TelemetryProviderAPI.ListTelemetryProviderTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTelemetryProviderTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportType** | **string** | Filter by export type (logs or metrics) | 

### Return type

[**[]TelemetryProviderTypeInfo**](TelemetryProviderTypeInfo.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

