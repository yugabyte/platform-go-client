# \RuntimeConfigurationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKey**](RuntimeConfigurationApi.md#DeleteKey) | **Delete** /api/v1/customers/{cUUID}/runtime_config/{scope}/key/{key} | Delete a configuration key
[**GetConfig**](RuntimeConfigurationApi.md#GetConfig) | **Get** /api/v1/customers/{cUUID}/runtime_config/{scope} | List configuration entries for a scope
[**GetConfigurationKey**](RuntimeConfigurationApi.md#GetConfigurationKey) | **Get** /api/v1/customers/{cUUID}/runtime_config/{scope}/key/{key} | Get a configuration key
[**ListKeyInfo**](RuntimeConfigurationApi.md#ListKeyInfo) | **Get** /api/v1/runtime_config/mutable_key_info | List mutable keys
[**ListKeys**](RuntimeConfigurationApi.md#ListKeys) | **Get** /api/v1/runtime_config/mutable_keys | List mutable keys
[**ListScopes**](RuntimeConfigurationApi.md#ListScopes) | **Get** /api/v1/customers/{cUUID}/runtime_config/scopes | List configuration scopes
[**SetKey**](RuntimeConfigurationApi.md#SetKey) | **Put** /api/v1/customers/{cUUID}/runtime_config/{scope}/key/{key} | Update a configuration key



## DeleteKey

> YBPSuccess DeleteKey(ctx, cUUID, scope, key).Execute()

Delete a configuration key

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
    scope := TODO // string | 
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RuntimeConfigurationApi.DeleteKey(context.Background(), cUUID, scope, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuntimeConfigurationApi.DeleteKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteKey`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `RuntimeConfigurationApi.DeleteKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**scope** | [**string**](.md) |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyRequest struct via the builder pattern


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


## GetConfig

> ScopedConfig GetConfig(ctx, cUUID, scope).IncludeInherited(includeInherited).Execute()

List configuration entries for a scope



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
    scope := TODO // string | 
    includeInherited := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RuntimeConfigurationApi.GetConfig(context.Background(), cUUID, scope).IncludeInherited(includeInherited).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuntimeConfigurationApi.GetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfig`: ScopedConfig
    fmt.Fprintf(os.Stdout, "Response from `RuntimeConfigurationApi.GetConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**scope** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeInherited** | **bool** |  | [default to false]

### Return type

[**ScopedConfig**](ScopedConfig.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurationKey

> string GetConfigurationKey(ctx, cUUID, scope, key).Execute()

Get a configuration key

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
    scope := TODO // string | 
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RuntimeConfigurationApi.GetConfigurationKey(context.Background(), cUUID, scope, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuntimeConfigurationApi.GetConfigurationKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigurationKey`: string
    fmt.Fprintf(os.Stdout, "Response from `RuntimeConfigurationApi.GetConfigurationKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**scope** | [**string**](.md) |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKeyInfo

> []ConfKeyInfo ListKeyInfo(ctx).Execute()

List mutable keys



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RuntimeConfigurationApi.ListKeyInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuntimeConfigurationApi.ListKeyInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKeyInfo`: []ConfKeyInfo
    fmt.Fprintf(os.Stdout, "Response from `RuntimeConfigurationApi.ListKeyInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListKeyInfoRequest struct via the builder pattern


### Return type

[**[]ConfKeyInfo**](ConfKeyInfo.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKeys

> []string ListKeys(ctx).Execute()

List mutable keys



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RuntimeConfigurationApi.ListKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuntimeConfigurationApi.ListKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKeys`: []string
    fmt.Fprintf(os.Stdout, "Response from `RuntimeConfigurationApi.ListKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListKeysRequest struct via the builder pattern


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


## ListScopes

> RuntimeConfigData ListScopes(ctx, cUUID).Execute()

List configuration scopes



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
    resp, r, err := api_client.RuntimeConfigurationApi.ListScopes(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuntimeConfigurationApi.ListScopes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListScopes`: RuntimeConfigData
    fmt.Fprintf(os.Stdout, "Response from `RuntimeConfigurationApi.ListScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RuntimeConfigData**](RuntimeConfigData.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetKey

> YBPSuccess SetKey(ctx, cUUID, scope, key).NewValue(newValue).Execute()

Update a configuration key

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
    scope := TODO // string | 
    key := "key_example" // string | 
    newValue := "newValue_example" // string | New value for config key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RuntimeConfigurationApi.SetKey(context.Background(), cUUID, scope, key).NewValue(newValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuntimeConfigurationApi.SetKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetKey`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `RuntimeConfigurationApi.SetKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**scope** | [**string**](.md) |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **newValue** | **string** | New value for config key | 

### Return type

[**YBPSuccess**](YBPSuccess.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

