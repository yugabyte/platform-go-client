# \DisasterRecoveryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDrConfig**](DisasterRecoveryApi.md#CreateDrConfig) | **Post** /api/v1/customers/{cUUID}/dr_configs | Create disaster recovery config
[**DeleteXClusterConfig**](DisasterRecoveryApi.md#DeleteXClusterConfig) | **Delete** /api/v1/customers/{cUUID}/dr_configs/{drUUID} | Delete xcluster config
[**FailoverDrConfig**](DisasterRecoveryApi.md#FailoverDrConfig) | **Post** /api/v1/customers/{cUUID}/dr_configs/{drUUID}/failover | Failover a disaster recovery config
[**GetDrConfig**](DisasterRecoveryApi.md#GetDrConfig) | **Get** /api/v1/customers/{cUUID}/dr_configs/{drUUID} | Get disaster recovery config
[**GetDrConfigSafetime**](DisasterRecoveryApi.md#GetDrConfigSafetime) | **Get** /api/v1/customers/{cUUID}/dr_configs/{drUUID}/safetime | Get disaster recovery config safetime
[**ReplaceReplicaDrConfig**](DisasterRecoveryApi.md#ReplaceReplicaDrConfig) | **Post** /api/v1/customers/{cUUID}/dr_configs/{drUUID}/replace_replica | Replace Replica in a disaster recovery config
[**RestartDrConfig**](DisasterRecoveryApi.md#RestartDrConfig) | **Post** /api/v1/customers/{cUUID}/dr_configs/{drUUID}/restart | Restart disaster recovery config
[**SetTablesDrConfig**](DisasterRecoveryApi.md#SetTablesDrConfig) | **Post** /api/v1/customers/{cUUID}/dr_configs/{drUUID}/set_tables | Set tables in disaster recovery config
[**SwitchoverDrConfig**](DisasterRecoveryApi.md#SwitchoverDrConfig) | **Post** /api/v1/customers/{cUUID}/dr_configs/{drUUID}/switchover | Switchover a disaster recovery config
[**SyncDrConfig**](DisasterRecoveryApi.md#SyncDrConfig) | **Post** /api/v1/customers/{cUUID}/dr_configs/{drUUID}/sync | Sync disaster recovery config



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


## FailoverDrConfig

> YBPTask FailoverDrConfig(ctx, cUUID, drUUID).DisasterRecoveryFailoverFormData(disasterRecoveryFailoverFormData).Request(request).Execute()

Failover a disaster recovery config

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
    disasterRecoveryFailoverFormData := *openapiclient.NewDrConfigFailoverForm() // DrConfigFailoverForm | Disaster Recovery Failover Form Data
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DisasterRecoveryApi.FailoverDrConfig(context.Background(), cUUID, drUUID).DisasterRecoveryFailoverFormData(disasterRecoveryFailoverFormData).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryApi.FailoverDrConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FailoverDrConfig`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryApi.FailoverDrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**drUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFailoverDrConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disasterRecoveryFailoverFormData** | [**DrConfigFailoverForm**](DrConfigFailoverForm.md) | Disaster Recovery Failover Form Data | 
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


## GetDrConfig

> DrConfig GetDrConfig(ctx, cUUID, drUUID).Execute()

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
    // response from `GetDrConfig`: DrConfig
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

[**DrConfig**](DrConfig.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDrConfigSafetime

> DrConfigSafetimeResp GetDrConfigSafetime(ctx, cUUID, drUUID).Execute()

Get disaster recovery config safetime

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
    resp, r, err := api_client.DisasterRecoveryApi.GetDrConfigSafetime(context.Background(), cUUID, drUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryApi.GetDrConfigSafetime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDrConfigSafetime`: DrConfigSafetimeResp
    fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryApi.GetDrConfigSafetime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**drUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDrConfigSafetimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DrConfigSafetimeResp**](DrConfigSafetimeResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceReplicaDrConfig

> YBPTask ReplaceReplicaDrConfig(ctx, cUUID, drUUID).DisasterRecoveryReplaceReplicaFormData(disasterRecoveryReplaceReplicaFormData).Request(request).Execute()

Replace Replica in a disaster recovery config

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
    disasterRecoveryReplaceReplicaFormData := *openapiclient.NewDrConfigReplaceReplicaForm() // DrConfigReplaceReplicaForm | Disaster Recovery Replace Replica Form Data
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DisasterRecoveryApi.ReplaceReplicaDrConfig(context.Background(), cUUID, drUUID).DisasterRecoveryReplaceReplicaFormData(disasterRecoveryReplaceReplicaFormData).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryApi.ReplaceReplicaDrConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceReplicaDrConfig`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryApi.ReplaceReplicaDrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**drUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceReplicaDrConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disasterRecoveryReplaceReplicaFormData** | [**DrConfigReplaceReplicaForm**](DrConfigReplaceReplicaForm.md) | Disaster Recovery Replace Replica Form Data | 
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


## RestartDrConfig

> YBPTask RestartDrConfig(ctx, cUUID, drUUID).DisasterRecoveryRestartFormData(disasterRecoveryRestartFormData).IsForceDelete(isForceDelete).Request(request).Execute()

Restart disaster recovery config

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
    disasterRecoveryRestartFormData := *openapiclient.NewDrConfigRestartForm([]string{"Dbs_example"}) // DrConfigRestartForm | Disaster Recovery Restart Form Data
    isForceDelete := true // bool |  (optional) (default to false)
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DisasterRecoveryApi.RestartDrConfig(context.Background(), cUUID, drUUID).DisasterRecoveryRestartFormData(disasterRecoveryRestartFormData).IsForceDelete(isForceDelete).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryApi.RestartDrConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestartDrConfig`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryApi.RestartDrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**drUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartDrConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disasterRecoveryRestartFormData** | [**DrConfigRestartForm**](DrConfigRestartForm.md) | Disaster Recovery Restart Form Data | 
 **isForceDelete** | **bool** |  | [default to false]
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


## SetTablesDrConfig

> YBPTask SetTablesDrConfig(ctx, cUUID, drUUID).DisasterRecoverySetTablesFormData(disasterRecoverySetTablesFormData).Request(request).Execute()

Set tables in disaster recovery config

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
    disasterRecoverySetTablesFormData := *openapiclient.NewDrConfigSetTablesForm() // DrConfigSetTablesForm | Disaster Recovery Set Tables Form Data
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DisasterRecoveryApi.SetTablesDrConfig(context.Background(), cUUID, drUUID).DisasterRecoverySetTablesFormData(disasterRecoverySetTablesFormData).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryApi.SetTablesDrConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetTablesDrConfig`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryApi.SetTablesDrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**drUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetTablesDrConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disasterRecoverySetTablesFormData** | [**DrConfigSetTablesForm**](DrConfigSetTablesForm.md) | Disaster Recovery Set Tables Form Data | 
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


## SwitchoverDrConfig

> YBPTask SwitchoverDrConfig(ctx, cUUID, drUUID).DisasterRecoverySwitchoverFormData(disasterRecoverySwitchoverFormData).Request(request).Execute()

Switchover a disaster recovery config

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
    disasterRecoverySwitchoverFormData := *openapiclient.NewDrConfigSwitchoverForm() // DrConfigSwitchoverForm | Disaster Recovery Switchover Form Data
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DisasterRecoveryApi.SwitchoverDrConfig(context.Background(), cUUID, drUUID).DisasterRecoverySwitchoverFormData(disasterRecoverySwitchoverFormData).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryApi.SwitchoverDrConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SwitchoverDrConfig`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryApi.SwitchoverDrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**drUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwitchoverDrConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disasterRecoverySwitchoverFormData** | [**DrConfigSwitchoverForm**](DrConfigSwitchoverForm.md) | Disaster Recovery Switchover Form Data | 
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


## SyncDrConfig

> YBPTask SyncDrConfig(ctx, cUUID, drUUID).Request(request).Execute()

Sync disaster recovery config

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DisasterRecoveryApi.SyncDrConfig(context.Background(), cUUID, drUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryApi.SyncDrConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncDrConfig`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryApi.SyncDrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**drUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncDrConfigRequest struct via the builder pattern


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

