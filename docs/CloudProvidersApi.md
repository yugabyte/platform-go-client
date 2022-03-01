# \CloudProvidersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProviders**](CloudProvidersApi.md#CreateProviders) | **Post** /api/v1/customers/{cUUID}/providers | Create a provider
[**EditProvider**](CloudProvidersApi.md#EditProvider) | **Put** /api/v1/customers/{cUUID}/providers/{pUUID}/edit | Update a provider
[**GetListOfProviders**](CloudProvidersApi.md#GetListOfProviders) | **Get** /api/v1/customers/{cUUID}/providers | List cloud providers
[**RefreshPricing**](CloudProvidersApi.md#RefreshPricing) | **Put** /api/v1/customers/{cUUID}/providers/{pUUID}/refresh_pricing | Refresh pricing



## CreateProviders

> YBPTask CreateProviders(ctx, cUUID).CreateProviderRequest(createProviderRequest).Execute()

Create a provider

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
    createProviderRequest := *openapiclient.NewProvider([]openapiclient.Region{*openapiclient.NewRegion([]openapiclient.AvailabilityZone{*openapiclient.NewAvailabilityZone("south-east-1")})}) // Provider | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProvidersApi.CreateProviders(context.Background(), cUUID).CreateProviderRequest(createProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.CreateProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProviders`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.CreateProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createProviderRequest** | [**Provider**](Provider.md) |  | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditProvider

> Provider EditProvider(ctx, cUUID, pUUID).EditProviderFormData(editProviderFormData).Execute()

Update a provider

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
    pUUID := TODO // string | 
    editProviderFormData := *openapiclient.NewEditProviderRequest(map[string]string{"key": "Inner_example"}, "HostedZoneId_example") // EditProviderRequest | edit provider form data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProvidersApi.EditProvider(context.Background(), cUUID, pUUID).EditProviderFormData(editProviderFormData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.EditProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditProvider`: Provider
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.EditProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **editProviderFormData** | [**EditProviderRequest**](EditProviderRequest.md) | edit provider form data | 

### Return type

[**Provider**](Provider.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListOfProviders

> []Provider GetListOfProviders(ctx, cUUID).Name(name).ProviderCode(providerCode).Execute()

List cloud providers

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
    name := "name_example" // string |  (optional) (default to "null")
    providerCode := "providerCode_example" // string |  (optional) (default to "null")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProvidersApi.GetListOfProviders(context.Background(), cUUID).Name(name).ProviderCode(providerCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.GetListOfProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListOfProviders`: []Provider
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.GetListOfProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListOfProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | [default to &quot;null&quot;]
 **providerCode** | **string** |  | [default to &quot;null&quot;]

### Return type

[**[]Provider**](Provider.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshPricing

> YBPSuccess RefreshPricing(ctx, cUUID, pUUID).Execute()

Refresh pricing



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
    pUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProvidersApi.RefreshPricing(context.Background(), cUUID, pUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.RefreshPricing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshPricing`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.RefreshPricing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshPricingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

