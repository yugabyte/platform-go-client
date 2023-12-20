# \AsynchronousReplicationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateXClusterConfig**](AsynchronousReplicationApi.md#CreateXClusterConfig) | **Post** /api/v1/customers/{cUUID}/xcluster_configs | Available since YBA version 2.16.0.0. Create xcluster config
[**DeleteXClusterConfig**](AsynchronousReplicationApi.md#DeleteXClusterConfig) | **Delete** /api/v1/customers/{cUUID}/xcluster_configs/{xccUUID} | Available since YBA version 2.16.0.0. Delete xcluster config
[**EditXClusterConfig**](AsynchronousReplicationApi.md#EditXClusterConfig) | **Put** /api/v1/customers/{cUUID}/xcluster_configs/{xccUUID} | Available since YBA version 2.16.0.0. Edit xcluster config
[**GetXClusterConfig**](AsynchronousReplicationApi.md#GetXClusterConfig) | **Get** /api/v1/customers/{cUUID}/xcluster_configs/{xccUUID} | Available since YBA version 2.16.0.0. Get xcluster config
[**NeedBootstrapTable**](AsynchronousReplicationApi.md#NeedBootstrapTable) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/need_bootstrap | WARNING: This is a preview API that could change. Whether tables need bootstrap before setting up cross cluster replication
[**RestartXClusterConfig**](AsynchronousReplicationApi.md#RestartXClusterConfig) | **Post** /api/v1/customers/{cUUID}/xcluster_configs/{xccUUID} | Available since YBA version 2.16.0.0. Restart xcluster config
[**SyncXClusterConfig**](AsynchronousReplicationApi.md#SyncXClusterConfig) | **Post** /api/v1/customers/{cUUID}/xcluster_configs/sync | Available since YBA version 2.16.0.0. Sync xcluster config



## CreateXClusterConfig

> YBPTask CreateXClusterConfig(ctx, cUUID).XclusterReplicationCreateFormData(xclusterReplicationCreateFormData).Request(request).Execute()

Available since YBA version 2.16.0.0. Create xcluster config

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

Available since YBA version 2.16.0.0. Delete xcluster config

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

Available since YBA version 2.16.0.0. Edit xcluster config

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

> XClusterConfigGetResp GetXClusterConfig(ctx, cUUID, xccUUID).Execute()

Available since YBA version 2.16.0.0. Get xcluster config

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AsynchronousReplicationApi.GetXClusterConfig(context.Background(), cUUID, xccUUID).Execute()
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

> map[string]map[string]interface{} NeedBootstrapTable(ctx, cUUID, uniUUID).XclusterNeedBootstrapFormData(xclusterNeedBootstrapFormData).ConfigType(configType).Request(request).Execute()

WARNING: This is a preview API that could change. Whether tables need bootstrap before setting up cross cluster replication

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AsynchronousReplicationApi.NeedBootstrapTable(context.Background(), cUUID, uniUUID).XclusterNeedBootstrapFormData(xclusterNeedBootstrapFormData).ConfigType(configType).Request(request).Execute()
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

Available since YBA version 2.16.0.0. Restart xcluster config

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

Available since YBA version 2.16.0.0. Sync xcluster config

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

