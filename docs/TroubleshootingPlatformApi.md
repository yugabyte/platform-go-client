# \TroubleshootingPlatformApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckRegistered**](TroubleshootingPlatformApi.md#CheckRegistered) | **Get** /api/v1/customers/{cUUID}/troubleshooting_platform/{tpUUID}/universes/{uUUID} | Check if universe is registered with Troubleshooting Platform
[**CreateTroubleshootingPlatform**](TroubleshootingPlatformApi.md#CreateTroubleshootingPlatform) | **Post** /api/v1/customers/{cUUID}/troubleshooting_platform | Create Troubleshooting Platform
[**DeleteTroubleshootingPlatform**](TroubleshootingPlatformApi.md#DeleteTroubleshootingPlatform) | **Delete** /api/v1/customers/{cUUID}/troubleshooting_platform/{tpUUID} | Delete Troubleshooting Platform
[**EditTroubleshootingPlatform**](TroubleshootingPlatformApi.md#EditTroubleshootingPlatform) | **Put** /api/v1/customers/{cUUID}/troubleshooting_platform/{tpUUID} | Edit Troubleshooting Platform
[**GetTroubleshootingPlatform**](TroubleshootingPlatformApi.md#GetTroubleshootingPlatform) | **Get** /api/v1/customers/{cUUID}/troubleshooting_platform/{tpUUID} | Get Troubleshooting Platform
[**ListAllTroubleshootingPlatforms**](TroubleshootingPlatformApi.md#ListAllTroubleshootingPlatforms) | **Get** /api/v1/customers/{cUUID}/troubleshooting_platform | List All Troubleshooting Platforms
[**RegisterUniverse**](TroubleshootingPlatformApi.md#RegisterUniverse) | **Put** /api/v1/customers/{cUUID}/troubleshooting_platform/{tpUUID}/universes/{uUUID} | Register universe with Troubleshooting Platform
[**UnregisterUniverse**](TroubleshootingPlatformApi.md#UnregisterUniverse) | **Delete** /api/v1/customers/{cUUID}/troubleshooting_platform/{tpUUID}/universes/{uUUID} | Unregister universe from Troubleshooting Platform



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
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    tpUUID := TODO // string | 
    uUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TroubleshootingPlatformApi.CheckRegistered(context.Background(), cUUID, tpUUID, uUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingPlatformApi.CheckRegistered``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckRegistered`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingPlatformApi.CheckRegistered`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**tpUUID** | [**string**](.md) |  | 
**uUUID** | [**string**](.md) |  | 

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
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    platformData := *openapiclient.NewTroubleshootingPlatform("ApiToken_example", "CustomerUUID_example", int64(123), "MetricsUrl_example", "TpUrl_example", "Uuid_example", "YbaUrl_example") // TroubleshootingPlatform | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TroubleshootingPlatformApi.CreateTroubleshootingPlatform(context.Background(), cUUID).PlatformData(platformData).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingPlatformApi.CreateTroubleshootingPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTroubleshootingPlatform`: TroubleshootingPlatform
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingPlatformApi.CreateTroubleshootingPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

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
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    tpUUID := TODO // string | 
    force := true // bool |  (optional) (default to false)
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TroubleshootingPlatformApi.DeleteTroubleshootingPlatform(context.Background(), cUUID, tpUUID).Force(force).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingPlatformApi.DeleteTroubleshootingPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTroubleshootingPlatform`: bool
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingPlatformApi.DeleteTroubleshootingPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**tpUUID** | [**string**](.md) |  | 

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
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    tpUUID := TODO // string | 
    platformData := *openapiclient.NewTroubleshootingPlatform("ApiToken_example", "CustomerUUID_example", int64(123), "MetricsUrl_example", "TpUrl_example", "Uuid_example", "YbaUrl_example") // TroubleshootingPlatform | 
    force := true // bool |  (optional) (default to false)
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TroubleshootingPlatformApi.EditTroubleshootingPlatform(context.Background(), cUUID, tpUUID).PlatformData(platformData).Force(force).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingPlatformApi.EditTroubleshootingPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditTroubleshootingPlatform`: TroubleshootingPlatform
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingPlatformApi.EditTroubleshootingPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**tpUUID** | [**string**](.md) |  | 

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
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    tpUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TroubleshootingPlatformApi.GetTroubleshootingPlatform(context.Background(), cUUID, tpUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingPlatformApi.GetTroubleshootingPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTroubleshootingPlatform`: TroubleshootingPlatformDetailsModel
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingPlatformApi.GetTroubleshootingPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**tpUUID** | [**string**](.md) |  | 

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
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TroubleshootingPlatformApi.ListAllTroubleshootingPlatforms(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingPlatformApi.ListAllTroubleshootingPlatforms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllTroubleshootingPlatforms`: []TroubleshootingPlatformDetailsModel
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingPlatformApi.ListAllTroubleshootingPlatforms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

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
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    tpUUID := TODO // string | 
    uUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TroubleshootingPlatformApi.RegisterUniverse(context.Background(), cUUID, tpUUID, uUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingPlatformApi.RegisterUniverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterUniverse`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingPlatformApi.RegisterUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**tpUUID** | [**string**](.md) |  | 
**uUUID** | [**string**](.md) |  | 

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
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    tpUUID := TODO // string | 
    uUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TroubleshootingPlatformApi.UnregisterUniverse(context.Background(), cUUID, tpUUID, uUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TroubleshootingPlatformApi.UnregisterUniverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnregisterUniverse`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `TroubleshootingPlatformApi.UnregisterUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**tpUUID** | [**string**](.md) |  | 
**uUUID** | [**string**](.md) |  | 

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

