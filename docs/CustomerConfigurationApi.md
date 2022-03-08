# \CustomerConfigurationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerConfig**](CustomerConfigurationApi.md#CreateCustomerConfig) | **Post** /api/v1/customers/{cUUID}/configs | Create a customer configuration
[**DeleteCustomerConfig**](CustomerConfigurationApi.md#DeleteCustomerConfig) | **Delete** /api/v1/customers/{cUUID}/configs/{configUUID} | Delete a customer configuration
[**DeleteCustomerConfigV2**](CustomerConfigurationApi.md#DeleteCustomerConfigV2) | **Delete** /api/v1/customers/{cUUID}/configs/{configUUID}/delete | Delete a customer configuration V2
[**EditCustomerConfig**](CustomerConfigurationApi.md#EditCustomerConfig) | **Put** /api/v1/customers/{cUUID}/configs/{configUUID} | Update a customer configuration
[**GetCustomerConfig**](CustomerConfigurationApi.md#GetCustomerConfig) | **Put** /api/v1/customers/{cUUID}/configs/{configUUID}/edit | Update a customer configuration V2
[**GetListOfCustomerConfig**](CustomerConfigurationApi.md#GetListOfCustomerConfig) | **Get** /api/v1/customers/{cUUID}/configs | List all customer configurations



## CreateCustomerConfig

> CustomerConfig CreateCustomerConfig(ctx, cUUID).Config(config).Execute()

Create a customer configuration

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
    config := *openapiclient.NewCustomerConfig("backup20-01-2021", "CustomerUUID_example", map[string]interface{}("{\"AWS_ACCESS_KEY_ID\": \"AK****************ZD\"}"), "S3", "STORAGE") // CustomerConfig | Configuration data to be created

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerConfigurationApi.CreateCustomerConfig(context.Background(), cUUID).Config(config).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerConfigurationApi.CreateCustomerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomerConfig`: CustomerConfig
    fmt.Fprintf(os.Stdout, "Response from `CustomerConfigurationApi.CreateCustomerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **config** | [**CustomerConfig**](CustomerConfig.md) | Configuration data to be created | 

### Return type

[**CustomerConfig**](CustomerConfig.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomerConfig

> YBPTask DeleteCustomerConfig(ctx, cUUID, configUUID).IsDeleteBackups(isDeleteBackups).Execute()

Delete a customer configuration

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
    isDeleteBackups := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerConfigurationApi.DeleteCustomerConfig(context.Background(), cUUID, configUUID).IsDeleteBackups(isDeleteBackups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerConfigurationApi.DeleteCustomerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCustomerConfig`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `CustomerConfigurationApi.DeleteCustomerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**configUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isDeleteBackups** | **bool** |  | [default to false]

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


## DeleteCustomerConfigV2

> YBPTask DeleteCustomerConfigV2(ctx, cUUID, configUUID).IsDeleteBackups(isDeleteBackups).Execute()

Delete a customer configuration V2

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
    isDeleteBackups := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerConfigurationApi.DeleteCustomerConfigV2(context.Background(), cUUID, configUUID).IsDeleteBackups(isDeleteBackups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerConfigurationApi.DeleteCustomerConfigV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCustomerConfigV2`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `CustomerConfigurationApi.DeleteCustomerConfigV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**configUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerConfigV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isDeleteBackups** | **bool** |  | [default to false]

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


## EditCustomerConfig

> CustomerConfig EditCustomerConfig(ctx, cUUID, configUUID).Config(config).Execute()

Update a customer configuration

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
    config := *openapiclient.NewCustomerConfig("backup20-01-2021", "CustomerUUID_example", map[string]interface{}("{\"AWS_ACCESS_KEY_ID\": \"AK****************ZD\"}"), "S3", "STORAGE") // CustomerConfig | Configuration data to be updated

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerConfigurationApi.EditCustomerConfig(context.Background(), cUUID, configUUID).Config(config).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerConfigurationApi.EditCustomerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditCustomerConfig`: CustomerConfig
    fmt.Fprintf(os.Stdout, "Response from `CustomerConfigurationApi.EditCustomerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**configUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditCustomerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **config** | [**CustomerConfig**](CustomerConfig.md) | Configuration data to be updated | 

### Return type

[**CustomerConfig**](CustomerConfig.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerConfig

> CustomerConfig GetCustomerConfig(ctx, cUUID, configUUID).Config(config).Execute()

Update a customer configuration V2

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
    config := *openapiclient.NewCustomerConfig("backup20-01-2021", "CustomerUUID_example", map[string]interface{}("{\"AWS_ACCESS_KEY_ID\": \"AK****************ZD\"}"), "S3", "STORAGE") // CustomerConfig | Configuration data to be updated

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerConfigurationApi.GetCustomerConfig(context.Background(), cUUID, configUUID).Config(config).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerConfigurationApi.GetCustomerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerConfig`: CustomerConfig
    fmt.Fprintf(os.Stdout, "Response from `CustomerConfigurationApi.GetCustomerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**configUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **config** | [**CustomerConfig**](CustomerConfig.md) | Configuration data to be updated | 

### Return type

[**CustomerConfig**](CustomerConfig.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListOfCustomerConfig

> []CustomerConfigUI GetListOfCustomerConfig(ctx, cUUID).Execute()

List all customer configurations

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
    resp, r, err := api_client.CustomerConfigurationApi.GetListOfCustomerConfig(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerConfigurationApi.GetListOfCustomerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListOfCustomerConfig`: []CustomerConfigUI
    fmt.Fprintf(os.Stdout, "Response from `CustomerConfigurationApi.GetListOfCustomerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListOfCustomerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CustomerConfigUI**](CustomerConfigUI.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

