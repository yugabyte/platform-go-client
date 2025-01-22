# \TelemetryProviderApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTelemetry**](TelemetryProviderApi.md#CreateTelemetry) | **Post** /api/v1/customers/{cUUID}/telemetry_provider | Create Telemetry Provider
[**DeleteTelemetryProvider**](TelemetryProviderApi.md#DeleteTelemetryProvider) | **Delete** /api/v1/customers/{cUUID}/telemetry_provider/{intUUID} | Delete a telemetry provider
[**GetTelemetryProvider**](TelemetryProviderApi.md#GetTelemetryProvider) | **Get** /api/v1/customers/{cUUID}/telemetry_provider/{intUUID} | Get Telemetry Provider
[**ListAllTelemetryProviders**](TelemetryProviderApi.md#ListAllTelemetryProviders) | **Get** /api/v1/customers/{cUUID}/telemetry_provider | List All Telemetry Providers



## CreateTelemetry

> TelemetryProvider CreateTelemetry(ctx, cUUID).ProviderData(providerData).Request(request).Execute()

Create Telemetry Provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    providerData := *openapiclient.NewTelemetryProvider(*openapiclient.NewTelemetryProviderConfig(), "CustomerUUID_example", "Name_example", map[string]string{"key": "Inner_example"}, "Uuid_example") // TelemetryProvider | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryProviderApi.CreateTelemetry(context.Background(), cUUID).ProviderData(providerData).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryProviderApi.CreateTelemetry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTelemetry`: TelemetryProvider
    fmt.Fprintf(os.Stdout, "Response from `TelemetryProviderApi.CreateTelemetry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTelemetryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerData** | [**TelemetryProvider**](TelemetryProvider.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**TelemetryProvider**](TelemetryProvider.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTelemetryProvider

> YBPSuccess DeleteTelemetryProvider(ctx, cUUID, intUUID).Request(request).Execute()

Delete a telemetry provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    intUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryProviderApi.DeleteTelemetryProvider(context.Background(), cUUID, intUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryProviderApi.DeleteTelemetryProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTelemetryProvider`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `TelemetryProviderApi.DeleteTelemetryProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**intUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTelemetryProviderRequest struct via the builder pattern


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


## GetTelemetryProvider

> TelemetryProvider GetTelemetryProvider(ctx, cUUID, intUUID).Execute()

Get Telemetry Provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    intUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryProviderApi.GetTelemetryProvider(context.Background(), cUUID, intUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryProviderApi.GetTelemetryProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTelemetryProvider`: TelemetryProvider
    fmt.Fprintf(os.Stdout, "Response from `TelemetryProviderApi.GetTelemetryProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**intUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelemetryProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TelemetryProvider**](TelemetryProvider.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllTelemetryProviders

> []TelemetryProvider ListAllTelemetryProviders(ctx, cUUID).Execute()

List All Telemetry Providers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryProviderApi.ListAllTelemetryProviders(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryProviderApi.ListAllTelemetryProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllTelemetryProviders`: []TelemetryProvider
    fmt.Fprintf(os.Stdout, "Response from `TelemetryProviderApi.ListAllTelemetryProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllTelemetryProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TelemetryProvider**](TelemetryProvider.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

