# \AsynchronousReplicationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateXClusterConfig**](AsynchronousReplicationApi.md#CreateXClusterConfig) | **Post** /api/v1/customers/{cUUID}/xcluster_configs | Create xcluster config
[**DeleteXClusterConfig**](AsynchronousReplicationApi.md#DeleteXClusterConfig) | **Delete** /api/v1/customers/{cUUID}/xcluster_configs/{xccUUID} | Delete xcluster config
[**EditXClusterConfig**](AsynchronousReplicationApi.md#EditXClusterConfig) | **Put** /api/v1/customers/{cUUID}/xcluster_configs/{xccUUID} | Edit xcluster config
[**GetXClusterConfig**](AsynchronousReplicationApi.md#GetXClusterConfig) | **Get** /api/v1/customers/{cUUID}/xcluster_configs/{xccUUID} | Get xcluster config
[**SyncXClusterConfig**](AsynchronousReplicationApi.md#SyncXClusterConfig) | **Post** /api/v1/customers/{cUUID}/xcluster_configs/sync | Sync xcluster config



## CreateXClusterConfig

> YBPTask CreateXClusterConfig(ctx, cUUID).XclusterReplicationCreateFormData(xclusterReplicationCreateFormData).Execute()

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
    xclusterReplicationCreateFormData := *openapiclient.NewXClusterConfigCreateFormData("Name_example", "SourceUniverseUUID_example", []string{"Tables_example"}, "TargetUniverseUUID_example") // XClusterConfigCreateFormData | XCluster Replication Create Form Data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AsynchronousReplicationApi.CreateXClusterConfig(context.Background(), cUUID).XclusterReplicationCreateFormData(xclusterReplicationCreateFormData).Execute()
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

> YBPTask DeleteXClusterConfig(ctx, cUUID, xccUUID).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AsynchronousReplicationApi.DeleteXClusterConfig(context.Background(), cUUID, xccUUID).Execute()
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

> YBPTask EditXClusterConfig(ctx, cUUID, xccUUID).XclusterReplicationEditFormData(xclusterReplicationEditFormData).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AsynchronousReplicationApi.EditXClusterConfig(context.Background(), cUUID, xccUUID).XclusterReplicationEditFormData(xclusterReplicationEditFormData).Execute()
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


## SyncXClusterConfig

> YBPTask SyncXClusterConfig(ctx, cUUID).TargetUniverseUUID(targetUniverseUUID).Execute()

Sync xcluster config

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AsynchronousReplicationApi.SyncXClusterConfig(context.Background(), cUUID).TargetUniverseUUID(targetUniverseUUID).Execute()
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

