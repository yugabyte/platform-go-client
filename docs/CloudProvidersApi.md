# \CloudProvidersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccessKeyRotation**](CloudProvidersApi.md#AccessKeyRotation) | **Post** /api/v1/customers/{cUUID}/providers/{pUUID}/access_key_rotation | Rotate access key for a provider
[**CreateProviders**](CloudProvidersApi.md#CreateProviders) | **Post** /api/v1/customers/{cUUID}/providers | Create a provider
[**EditAccessKeyRotationSchedule**](CloudProvidersApi.md#EditAccessKeyRotationSchedule) | **Put** /api/v1/customers/{cUUID}/providers/{pUUID}/access_key_rotation/schedule/{sUUID} | Edit a access key rotation schedule
[**EditProvider**](CloudProvidersApi.md#EditProvider) | **Put** /api/v1/customers/{cUUID}/providers/{pUUID}/edit | Update a provider
[**GetListOfProviders**](CloudProvidersApi.md#GetListOfProviders) | **Get** /api/v1/customers/{cUUID}/providers | List cloud providers
[**GetProvider**](CloudProvidersApi.md#GetProvider) | **Get** /api/v1/customers/{cUUID}/providers/{pUUID} | Get a cloud provider
[**ListSchedules**](CloudProvidersApi.md#ListSchedules) | **Get** /api/v1/customers/{cUUID}/providers/{pUUID}/access_key_rotation/schedule | List all schedules for a provider&#39;s access key rotation
[**PatchProvider**](CloudProvidersApi.md#PatchProvider) | **Patch** /api/v1/customers/{cUUID}/providers/{pUUID} | Patch a provider
[**RefreshPricing**](CloudProvidersApi.md#RefreshPricing) | **Put** /api/v1/customers/{cUUID}/providers/{pUUID}/refresh_pricing | Refresh pricing
[**ScheduledAccessKeyRotation**](CloudProvidersApi.md#ScheduledAccessKeyRotation) | **Post** /api/v1/customers/{cUUID}/providers/{pUUID}/access_key_rotation/schedule | Rotate access key for a provider - Scheduled



## AccessKeyRotation

> YBPTask AccessKeyRotation(ctx, cUUID, pUUID).Execute()

Rotate access key for a provider

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
    resp, r, err := api_client.CloudProvidersApi.AccessKeyRotation(context.Background(), cUUID, pUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.AccessKeyRotation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccessKeyRotation`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.AccessKeyRotation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessKeyRotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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
    createProviderRequest := *openapiclient.NewProvider(*openapiclient.NewProviderDetails(), []openapiclient.Region{*openapiclient.NewRegion([]openapiclient.AvailabilityZone{*openapiclient.NewAvailabilityZone("us-west1-a")})}) // Provider | 

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


## EditAccessKeyRotationSchedule

> Schedule EditAccessKeyRotationSchedule(ctx, cUUID, pUUID, sUUID).Body(body).Execute()

Edit a access key rotation schedule

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
    sUUID := TODO // string | 
    body := *openapiclient.NewEditAccessKeyRotationScheduleParams() // EditAccessKeyRotationScheduleParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProvidersApi.EditAccessKeyRotationSchedule(context.Background(), cUUID, pUUID, sUUID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.EditAccessKeyRotationSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditAccessKeyRotationSchedule`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.EditAccessKeyRotationSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 
**sUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAccessKeyRotationScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**EditAccessKeyRotationScheduleParams**](EditAccessKeyRotationScheduleParams.md) |  | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditProvider

> YBPTask EditProvider(ctx, cUUID, pUUID).EditProviderRequest(editProviderRequest).Execute()

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
    editProviderRequest := *openapiclient.NewProvider(*openapiclient.NewProviderDetails(), []openapiclient.Region{*openapiclient.NewRegion([]openapiclient.AvailabilityZone{*openapiclient.NewAvailabilityZone("us-west1-a")})}) // Provider | edit provider form data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProvidersApi.EditProvider(context.Background(), cUUID, pUUID).EditProviderRequest(editProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.EditProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditProvider`: YBPTask
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


 **editProviderRequest** | [**Provider**](Provider.md) | edit provider form data | 

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


## GetProvider

> Provider GetProvider(ctx, cUUID, pUUID).Execute()

Get a cloud provider

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
    resp, r, err := api_client.CloudProvidersApi.GetProvider(context.Background(), cUUID, pUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.GetProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProvider`: Provider
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.GetProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Provider**](Provider.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSchedules

> []Schedule ListSchedules(ctx, cUUID, pUUID).Execute()

List all schedules for a provider's access key rotation

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
    resp, r, err := api_client.CloudProvidersApi.ListSchedules(context.Background(), cUUID, pUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.ListSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSchedules`: []Schedule
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.ListSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Schedule**](Schedule.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchProvider

> YBPTask PatchProvider(ctx, cUUID, pUUID).PatchProviderRequest(patchProviderRequest).Execute()

Patch a provider

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
    patchProviderRequest := *openapiclient.NewProvider(*openapiclient.NewProviderDetails(), []openapiclient.Region{*openapiclient.NewRegion([]openapiclient.AvailabilityZone{*openapiclient.NewAvailabilityZone("us-west1-a")})}) // Provider | patch provider form data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CloudProvidersApi.PatchProvider(context.Background(), cUUID, pUUID).PatchProviderRequest(patchProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.PatchProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchProvider`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.PatchProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchProviderRequest** | [**Provider**](Provider.md) | patch provider form data | 

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


## ScheduledAccessKeyRotation

> Schedule ScheduledAccessKeyRotation(ctx, cUUID, pUUID).Execute()

Rotate access key for a provider - Scheduled

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
    resp, r, err := api_client.CloudProvidersApi.ScheduledAccessKeyRotation(context.Background(), cUUID, pUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudProvidersApi.ScheduledAccessKeyRotation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduledAccessKeyRotation`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `CloudProvidersApi.ScheduledAccessKeyRotation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiScheduledAccessKeyRotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Schedule**](Schedule.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

