# \DisasterRecoveryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDrConfig**](DisasterRecoveryApi.md#CreateDrConfig) | **Post** /api/v1/customers/{cUUID}/dr_configs | Create disaster recovery config
[**DeleteXClusterConfig**](DisasterRecoveryApi.md#DeleteXClusterConfig) | **Delete** /api/v1/customers/{cUUID}/dr_configs/{drUUID} | Delete xcluster config
[**GetDrConfig**](DisasterRecoveryApi.md#GetDrConfig) | **Get** /api/v1/customers/{cUUID}/dr_configs/{drUUID} | Get disaster recovery config



## CreateDrConfig

> YBPTask CreateDrConfig(ctx, cUUID).DisasterRecoveryCreateFormData(disasterRecoveryCreateFormData).Request(request).Execute()

Create disaster recovery config

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
    disasterRecoveryCreateFormData := *openapiclient.NewDrConfigCreateForm([]string{"Dbs_example"}, "Dr-config1", "SourceUniverseUUID_example", "TargetUniverseUUID_example") // DrConfigCreateForm | Disaster Recovery Create Form Data
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DisasterRecoveryApi.CreateDrConfig(context.Background(), cUUID).DisasterRecoveryCreateFormData(disasterRecoveryCreateFormData).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryApi.CreateDrConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDrConfig`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryApi.CreateDrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDrConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disasterRecoveryCreateFormData** | [**DrConfigCreateForm**](DrConfigCreateForm.md) | Disaster Recovery Create Form Data | 
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


## DeleteXClusterConfig

> YBPTask DeleteXClusterConfig(ctx, cUUID, drUUID).IsForceDelete(isForceDelete).Request(request).Execute()

Delete xcluster config

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
    drUUID := TODO // string | 
    isForceDelete := true // bool |  (optional) (default to false)
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DisasterRecoveryApi.DeleteXClusterConfig(context.Background(), cUUID, drUUID).IsForceDelete(isForceDelete).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryApi.DeleteXClusterConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteXClusterConfig`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryApi.DeleteXClusterConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**drUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteXClusterConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isForceDelete** | **bool** |  | [default to false]
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


## GetDrConfig

> XClusterConfigGetResp GetDrConfig(ctx, cUUID, drUUID).Execute()

Get disaster recovery config

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
    drUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DisasterRecoveryApi.GetDrConfig(context.Background(), cUUID, drUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryApi.GetDrConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDrConfig`: XClusterConfigGetResp
    fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryApi.GetDrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**drUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDrConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**XClusterConfigGetResp**](XClusterConfigGetResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

