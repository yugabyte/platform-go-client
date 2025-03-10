# \AsynchronousReplicationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateXClusterConfig**](AsynchronousReplicationApi.md#CreateXClusterConfig) | **Post** /api/v1/customers/{cUUID}/xcluster_configs | Create xcluster config
[**DeleteXClusterConfig**](AsynchronousReplicationApi.md#DeleteXClusterConfig) | **Delete** /api/v1/customers/{cUUID}/xcluster_configs/{xccUUID} | Delete xcluster config
[**EditXClusterConfig**](AsynchronousReplicationApi.md#EditXClusterConfig) | **Put** /api/v1/customers/{cUUID}/xcluster_configs/{xccUUID} | Edit xcluster config
[**GetXClusterConfig**](AsynchronousReplicationApi.md#GetXClusterConfig) | **Get** /api/v1/customers/{cUUID}/xcluster_configs/{xccUUID} | Get xcluster config
[**NeedBootstrapTable**](AsynchronousReplicationApi.md#NeedBootstrapTable) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/need_bootstrap | Whether tables need bootstrap before setting up cross cluster replication
[**NeedBootstrapXClusterConfig**](AsynchronousReplicationApi.md#NeedBootstrapXClusterConfig) | **Post** /api/v1/customers/{cUUID}/xcluster_configs/{xccUUID}/need_bootstrap | Whether tables in an xCluster replication config have fallen far behind and need bootstrap
[**RestartXClusterConfig**](AsynchronousReplicationApi.md#RestartXClusterConfig) | **Post** /api/v1/customers/{cUUID}/xcluster_configs/{xccUUID} | Restart xcluster config
[**SyncXClusterConfig**](AsynchronousReplicationApi.md#SyncXClusterConfig) | **Post** /api/v1/customers/{cUUID}/xcluster_configs/sync | Sync xcluster config - deprecated
[**SyncXClusterConfigV2**](AsynchronousReplicationApi.md#SyncXClusterConfigV2) | **Post** /api/v1/customers/{cUUID}/xcluster_configs/{xccUUID}/sync | Sync xcluster config (V2)



## CreateXClusterConfig

> YBPTask CreateXClusterConfig(ctx, cUUID).XclusterReplicationCreateFormData(xclusterReplicationCreateFormData).Request(request).Execute()

Create xcluster config



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
    xclusterReplicationCreateFormData := *openapiclient.NewXClusterConfigCreateFormData("Repl-config1", "SourceUniverseUUID_example", []string{"Tables_example"}, "TargetUniverseUUID_example") // XClusterConfigCreateFormData | XCluster Replication Create Form Data
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AsynchronousReplicationApi.CreateXClusterConfig(context.Background(), cUUID).XclusterReplicationCreateFormData(xclusterReplicationCreateFormData).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AsynchronousReplicationApi.CreateXClusterConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateXClusterConfig`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `AsynchronousReplicationApi.CreateXClusterConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateXClusterConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xclusterReplicationCreateFormData** | [**XClusterConfigCreateFormData**](XClusterConfigCreateFormData.md) | XCluster Replication Create Form Data | 
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

> YBPTask DeleteXClusterConfig(ctx, cUUID, xccUUID).IsForceDelete(isForceDelete).Request(request).Execute()

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
    xccUUID := TODO // string | 
    isForceDelete := true // bool |  (optional) (default to false)
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AsynchronousReplicationApi.DeleteXClusterConfig(context.Background(), cUUID, xccUUID).IsForceDelete(isForceDelete).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AsynchronousReplicationApi.DeleteXClusterConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteXClusterConfig`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `AsynchronousReplicationApi.DeleteXClusterConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**xccUUID** | [**string**](.md) |  | 

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


## EditXClusterConfig

> YBPTask EditXClusterConfig(ctx, cUUID, xccUUID).XclusterReplicationEditFormData(xclusterReplicationEditFormData).Request(request).Execute()

Edit xcluster config



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
    xccUUID := TODO // string | 
    xclusterReplicationEditFormData := *openapiclient.NewXClusterConfigEditFormData() // XClusterConfigEditFormData | XCluster Replication Edit Form Data
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AsynchronousReplicationApi.EditXClusterConfig(context.Background(), cUUID, xccUUID).XclusterReplicationEditFormData(xclusterReplicationEditFormData).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AsynchronousReplicationApi.EditXClusterConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditXClusterConfig`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `AsynchronousReplicationApi.EditXClusterConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**xccUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditXClusterConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xclusterReplicationEditFormData** | [**XClusterConfigEditFormData**](XClusterConfigEditFormData.md) | XCluster Replication Edit Form Data | 
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


## GetXClusterConfig

> XClusterConfigGetResp GetXClusterConfig(ctx, cUUID, xccUUID).SyncWithDB(syncWithDB).Execute()

Get xcluster config



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
    xccUUID := TODO // string | 
    syncWithDB := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AsynchronousReplicationApi.GetXClusterConfig(context.Background(), cUUID, xccUUID).SyncWithDB(syncWithDB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AsynchronousReplicationApi.GetXClusterConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetXClusterConfig`: XClusterConfigGetResp
    fmt.Fprintf(os.Stdout, "Response from `AsynchronousReplicationApi.GetXClusterConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**xccUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetXClusterConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **syncWithDB** | **bool** |  | [default to true]

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


## NeedBootstrapTable

> map[string]map[string]interface{} NeedBootstrapTable(ctx, cUUID, uniUUID).XclusterNeedBootstrapFormData(xclusterNeedBootstrapFormData).ConfigType(configType).IncludeDetails(includeDetails).Request(request).Execute()

Whether tables need bootstrap before setting up cross cluster replication



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
    xclusterNeedBootstrapFormData := *openapiclient.NewXClusterConfigNeedBootstrapFormData([]string{"Tables_example"}) // XClusterConfigNeedBootstrapFormData | XCluster Need Bootstrap Form Data
    configType := "configType_example" // string |  (optional) (default to "null")
    includeDetails := true // bool |  (optional) (default to false)
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AsynchronousReplicationApi.NeedBootstrapTable(context.Background(), cUUID, uniUUID).XclusterNeedBootstrapFormData(xclusterNeedBootstrapFormData).ConfigType(configType).IncludeDetails(includeDetails).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AsynchronousReplicationApi.NeedBootstrapTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NeedBootstrapTable`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AsynchronousReplicationApi.NeedBootstrapTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNeedBootstrapTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xclusterNeedBootstrapFormData** | [**XClusterConfigNeedBootstrapFormData**](XClusterConfigNeedBootstrapFormData.md) | XCluster Need Bootstrap Form Data | 
 **configType** | **string** |  | [default to &quot;null&quot;]
 **includeDetails** | **bool** |  | [default to false]
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

**map[string]map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NeedBootstrapXClusterConfig

> map[string]map[string]interface{} NeedBootstrapXClusterConfig(ctx, cUUID, xccUUID).XclusterNeedBootstrapFormData(xclusterNeedBootstrapFormData).Request(request).Execute()

Whether tables in an xCluster replication config have fallen far behind and need bootstrap



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
    xccUUID := TODO // string | 
    xclusterNeedBootstrapFormData := *openapiclient.NewXClusterConfigNeedBootstrapFormData([]string{"Tables_example"}) // XClusterConfigNeedBootstrapFormData | XCluster Need Bootstrap Form Data
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AsynchronousReplicationApi.NeedBootstrapXClusterConfig(context.Background(), cUUID, xccUUID).XclusterNeedBootstrapFormData(xclusterNeedBootstrapFormData).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AsynchronousReplicationApi.NeedBootstrapXClusterConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NeedBootstrapXClusterConfig`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AsynchronousReplicationApi.NeedBootstrapXClusterConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**xccUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNeedBootstrapXClusterConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xclusterNeedBootstrapFormData** | [**XClusterConfigNeedBootstrapFormData**](XClusterConfigNeedBootstrapFormData.md) | XCluster Need Bootstrap Form Data | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

**map[string]map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartXClusterConfig

> YBPTask RestartXClusterConfig(ctx, cUUID, xccUUID).XclusterReplicationRestartFormData(xclusterReplicationRestartFormData).IsForceDelete(isForceDelete).Request(request).Execute()

Restart xcluster config



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
    xccUUID := TODO // string | 
    xclusterReplicationRestartFormData := *openapiclient.NewXClusterConfigRestartFormData([]string{"Tables_example"}) // XClusterConfigRestartFormData | XCluster Replication Restart Form Data
    isForceDelete := true // bool |  (optional) (default to false)
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AsynchronousReplicationApi.RestartXClusterConfig(context.Background(), cUUID, xccUUID).XclusterReplicationRestartFormData(xclusterReplicationRestartFormData).IsForceDelete(isForceDelete).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AsynchronousReplicationApi.RestartXClusterConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestartXClusterConfig`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `AsynchronousReplicationApi.RestartXClusterConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**xccUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartXClusterConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xclusterReplicationRestartFormData** | [**XClusterConfigRestartFormData**](XClusterConfigRestartFormData.md) | XCluster Replication Restart Form Data | 
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


## SyncXClusterConfig

> YBPTask SyncXClusterConfig(ctx, cUUID).TargetUniverseUUID(targetUniverseUUID).Request(request).Execute()

Sync xcluster config - deprecated



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
    targetUniverseUUID := TODO // string |  (optional)
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AsynchronousReplicationApi.SyncXClusterConfig(context.Background(), cUUID).TargetUniverseUUID(targetUniverseUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AsynchronousReplicationApi.SyncXClusterConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncXClusterConfig`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `AsynchronousReplicationApi.SyncXClusterConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncXClusterConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targetUniverseUUID** | [**string**](string.md) |  | 
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


## SyncXClusterConfigV2

> YBPTask SyncXClusterConfigV2(ctx, cUUID, xccUUID).Request(request).Execute()

Sync xcluster config (V2)



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
    xccUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AsynchronousReplicationApi.SyncXClusterConfigV2(context.Background(), cUUID, xccUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AsynchronousReplicationApi.SyncXClusterConfigV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncXClusterConfigV2`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `AsynchronousReplicationApi.SyncXClusterConfigV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**xccUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncXClusterConfigV2Request struct via the builder pattern


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

