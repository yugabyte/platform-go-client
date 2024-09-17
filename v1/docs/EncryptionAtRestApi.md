# \EncryptionAtRestApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKMSConfig**](EncryptionAtRestApi.md#CreateKMSConfig) | **Post** /api/v1/customers/{cUUID}/kms_configs/{kmsProvider} | Create a KMS configuration
[**DeleteKMSConfig**](EncryptionAtRestApi.md#DeleteKMSConfig) | **Delete** /api/v1/customers/{cUUID}/kms_configs/{configUUID} | Delete a KMS configuration
[**EditKMSConfig**](EncryptionAtRestApi.md#EditKMSConfig) | **Post** /api/v1/customers/{cUUID}/kms_configs/{configUUID}/edit | Edit a KMS configuration
[**GetKMSConfig**](EncryptionAtRestApi.md#GetKMSConfig) | **Get** /api/v1/customers/{cUUID}/kms_configs/{configUUID} | Get details of a KMS configuration
[**ListKMSConfigs**](EncryptionAtRestApi.md#ListKMSConfigs) | **Get** /api/v1/customers/{cUUID}/kms_configs | List KMS configurations
[**RefreshKMSConfig**](EncryptionAtRestApi.md#RefreshKMSConfig) | **Put** /api/v1/customers/{cUUID}/kms_configs/{configUUID}/refresh | Refresh KMS Config
[**RemoveKeyRefHistory**](EncryptionAtRestApi.md#RemoveKeyRefHistory) | **Delete** /api/v1/customers/{cUUID}/universes/{uniUUID}/kms | This API removes a universe&#39;s key reference history - deprecated



## CreateKMSConfig

> YBPTask CreateKMSConfig(ctx, cUUID, kmsProvider).KMSConfig(kMSConfig).Request(request).Execute()

Create a KMS configuration

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
    kmsProvider := "kmsProvider_example" // string | 
    kMSConfig := map[string]interface{}(Object) // map[string]interface{} | KMS config to be created
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EncryptionAtRestApi.CreateKMSConfig(context.Background(), cUUID, kmsProvider).KMSConfig(kMSConfig).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestApi.CreateKMSConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKMSConfig`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestApi.CreateKMSConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**kmsProvider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKMSConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kMSConfig** | **map[string]interface{}** | KMS config to be created | 
 **request** | [**interface{}**](interface{}.md) |  | 

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


## DeleteKMSConfig

> YBPTask DeleteKMSConfig(ctx, cUUID, configUUID).Request(request).Execute()

Delete a KMS configuration

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
    configUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EncryptionAtRestApi.DeleteKMSConfig(context.Background(), cUUID, configUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestApi.DeleteKMSConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteKMSConfig`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestApi.DeleteKMSConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**configUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKMSConfigRequest struct via the builder pattern


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


## EditKMSConfig

> YBPTask EditKMSConfig(ctx, cUUID, configUUID).KMSConfig(kMSConfig).Request(request).Execute()

Edit a KMS configuration

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
    configUUID := TODO // string | 
    kMSConfig := map[string]interface{}(Object) // map[string]interface{} | KMS config to be edited
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EncryptionAtRestApi.EditKMSConfig(context.Background(), cUUID, configUUID).KMSConfig(kMSConfig).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestApi.EditKMSConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditKMSConfig`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestApi.EditKMSConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**configUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditKMSConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kMSConfig** | **map[string]interface{}** | KMS config to be edited | 
 **request** | [**interface{}**](interface{}.md) |  | 

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


## GetKMSConfig

> map[string]map[string]interface{} GetKMSConfig(ctx, cUUID, configUUID).Execute()

Get details of a KMS configuration

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
    configUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EncryptionAtRestApi.GetKMSConfig(context.Background(), cUUID, configUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestApi.GetKMSConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKMSConfig`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestApi.GetKMSConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**configUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKMSConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKMSConfigs

> []map[string]interface{} ListKMSConfigs(ctx, cUUID).Execute()

List KMS configurations

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
    resp, r, err := api_client.EncryptionAtRestApi.ListKMSConfigs(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestApi.ListKMSConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKMSConfigs`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestApi.ListKMSConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKMSConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshKMSConfig

> YBPSuccess RefreshKMSConfig(ctx, cUUID, configUUID).Request(request).Execute()

Refresh KMS Config



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
    configUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EncryptionAtRestApi.RefreshKMSConfig(context.Background(), cUUID, configUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestApi.RefreshKMSConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshKMSConfig`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestApi.RefreshKMSConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**configUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshKMSConfigRequest struct via the builder pattern


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


## RemoveKeyRefHistory

> YBPSuccess RemoveKeyRefHistory(ctx, cUUID, uniUUID).Request(request).Execute()

This API removes a universe's key reference history - deprecated



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
    uniUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EncryptionAtRestApi.RemoveKeyRefHistory(context.Background(), cUUID, uniUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestApi.RemoveKeyRefHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveKeyRefHistory`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestApi.RemoveKeyRefHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveKeyRefHistoryRequest struct via the builder pattern


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

