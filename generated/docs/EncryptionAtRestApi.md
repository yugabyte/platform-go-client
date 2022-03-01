# \EncryptionAtRestApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKMSConfig**](EncryptionAtRestApi.md#CreateKMSConfig) | **Post** /api/v1/customers/{cUUID}/kms_configs/{kmsProvider} | Create a KMS configuration
[**DeleteKMSConfig**](EncryptionAtRestApi.md#DeleteKMSConfig) | **Delete** /api/v1/customers/{cUUID}/kms_configs/{configUUID} | Delete a KMS configuration
[**EditKMSConfig**](EncryptionAtRestApi.md#EditKMSConfig) | **Post** /api/v1/customers/{cUUID}/kms_configs/{configUUID}/edit | Edit a KMS configuration
[**GetCurrentKeyRef**](EncryptionAtRestApi.md#GetCurrentKeyRef) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/kms/key_ref | Get a universe&#39;s key reference
[**GetKMSConfig**](EncryptionAtRestApi.md#GetKMSConfig) | **Get** /api/v1/customers/{cUUID}/kms_configs/{configUUID} | Get details of a KMS configuration
[**GetKeyRefHistory**](EncryptionAtRestApi.md#GetKeyRefHistory) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/kms | Get a universe&#39;s key reference history
[**ListKMSConfigs**](EncryptionAtRestApi.md#ListKMSConfigs) | **Get** /api/v1/customers/{cUUID}/kms_configs | List KMS configurations
[**RemoveKeyRefHistory**](EncryptionAtRestApi.md#RemoveKeyRefHistory) | **Delete** /api/v1/customers/{cUUID}/universes/{uniUUID}/kms | Remove a universe&#39;s key reference history
[**RetrieveKey**](EncryptionAtRestApi.md#RetrieveKey) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/kms | Retrive a universe&#39;s KMS key



## CreateKMSConfig

> YBPTask CreateKMSConfig(ctx, cUUID, kmsProvider).KMSConfig(kMSConfig).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EncryptionAtRestApi.CreateKMSConfig(context.Background(), cUUID, kmsProvider).KMSConfig(kMSConfig).Execute()
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

> YBPTask DeleteKMSConfig(ctx, cUUID, configUUID).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EncryptionAtRestApi.DeleteKMSConfig(context.Background(), cUUID, configUUID).Execute()
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

> YBPTask EditKMSConfig(ctx, cUUID, configUUID).KMSConfig(kMSConfig).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EncryptionAtRestApi.EditKMSConfig(context.Background(), cUUID, configUUID).KMSConfig(kMSConfig).Execute()
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


## GetCurrentKeyRef

> map[string]map[string]interface{} GetCurrentKeyRef(ctx, cUUID, uniUUID).Execute()

Get a universe's key reference

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EncryptionAtRestApi.GetCurrentKeyRef(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestApi.GetCurrentKeyRef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentKeyRef`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestApi.GetCurrentKeyRef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentKeyRefRequest struct via the builder pattern


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


## GetKeyRefHistory

> []map[string]interface{} GetKeyRefHistory(ctx, cUUID, uniUUID).Execute()

Get a universe's key reference history

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EncryptionAtRestApi.GetKeyRefHistory(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestApi.GetKeyRefHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeyRefHistory`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestApi.GetKeyRefHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyRefHistoryRequest struct via the builder pattern


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


## RemoveKeyRefHistory

> YBPSuccess RemoveKeyRefHistory(ctx, cUUID, uniUUID).Execute()

Remove a universe's key reference history

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EncryptionAtRestApi.RemoveKeyRefHistory(context.Background(), cUUID, uniUUID).Execute()
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


## RetrieveKey

> map[string]map[string]interface{} RetrieveKey(ctx, cUUID, uniUUID).Execute()

Retrive a universe's KMS key

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EncryptionAtRestApi.RetrieveKey(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestApi.RetrieveKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveKey`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestApi.RetrieveKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveKeyRequest struct via the builder pattern


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

