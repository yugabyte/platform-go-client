# \UniverseApi

All URIs are relative to *http://localhost:9000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCluster**](UniverseApi.md#AddCluster) | **Post** /customers/{cUUID}/universes/{uniUUID}/clusters | Add a cluster to a YugabyteDB Universe
[**CreateUniverse**](UniverseApi.md#CreateUniverse) | **Post** /customers/{cUUID}/universes | Create a YugabyteDB Universe
[**DeleteCluster**](UniverseApi.md#DeleteCluster) | **Delete** /customers/{cUUID}/universes/{uniUUID}/clusters/{clsUUID} | Delete an additional cluster(s) of a YugabyteDB Universe
[**EditGFlags**](UniverseApi.md#EditGFlags) | **Post** /customers/{cUUID}/universes/{uniUUID}/gflags | Edit GFlags
[**EditUniverse**](UniverseApi.md#EditUniverse) | **Put** /customers/{cUUID}/universes/{uniUUID} | Edit a YugabyteDB Universe
[**GetUniverse**](UniverseApi.md#GetUniverse) | **Get** /customers/{cUUID}/universes/{uniUUID} | Get a YugabyteDB Universe



## AddCluster

> YBPTask AddCluster(ctx, cUUID, uniUUID).ClusterAddSpec(clusterAddSpec).Execute()

Add a cluster to a YugabyteDB Universe



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
    cUUID := TODO // string | Customer UUID
    uniUUID := TODO // string | Universe UUID
    clusterAddSpec := *openapiclient.NewClusterAddSpec("READ_REPLICA", int32(3), "c5.xlarge", *openapiclient.NewClusterProviderEditSpec([]string{"RegionList_example"})) // ClusterAddSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.AddCluster(context.Background(), cUUID, uniUUID).ClusterAddSpec(clusterAddSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.AddCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCluster`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseApi.AddCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**uniUUID** | [**string**](.md) | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clusterAddSpec** | [**ClusterAddSpec**](ClusterAddSpec.md) |  | 

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


## CreateUniverse

> YBPTask CreateUniverse(ctx, cUUID).UniverseCreateSpec(universeCreateSpec).Execute()

Create a YugabyteDB Universe



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
    cUUID := TODO // string | Customer UUID
    universeCreateSpec := *openapiclient.NewUniverseCreateSpec(*openapiclient.NewUniverseSpec("my-yb-universe", "2024.2.0.0-b600", []openapiclient.ClusterSpec{*openapiclient.NewClusterSpec("PRIMARY", int32(3), "c5.xlarge", *openapiclient.NewClusterStorageSpec(), *openapiclient.NewClusterProviderSpec("89a46d52-4edd-4736-922a-35177a0b990c", "aws-ssh-key"))}), "x86_64") // UniverseCreateSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.CreateUniverse(context.Background(), cUUID).UniverseCreateSpec(universeCreateSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.CreateUniverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUniverse`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseApi.CreateUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUniverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **universeCreateSpec** | [**UniverseCreateSpec**](UniverseCreateSpec.md) |  | 

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


## DeleteCluster

> YBPTask DeleteCluster(ctx, cUUID, uniUUID, clsUUID).IsForceDelete(isForceDelete).Execute()

Delete an additional cluster(s) of a YugabyteDB Universe



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
    cUUID := TODO // string | Customer UUID
    uniUUID := TODO // string | Universe UUID
    clsUUID := TODO // string | Cluster UUID
    isForceDelete := true // bool | Whether to force delete the cluster (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.DeleteCluster(context.Background(), cUUID, uniUUID, clsUUID).IsForceDelete(isForceDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.DeleteCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCluster`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseApi.DeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**uniUUID** | [**string**](.md) | Universe UUID | 
**clsUUID** | [**string**](.md) | Cluster UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **isForceDelete** | **bool** | Whether to force delete the cluster | [default to false]

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


## EditGFlags

> YBPTask EditGFlags(ctx, cUUID, uniUUID).UniverseEditGFlags(universeEditGFlags).Execute()

Edit GFlags



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
    cUUID := TODO // string | Customer UUID
    uniUUID := TODO // string | Universe UUID
    universeEditGFlags := *openapiclient.NewUniverseEditGFlags() // UniverseEditGFlags | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.EditGFlags(context.Background(), cUUID, uniUUID).UniverseEditGFlags(universeEditGFlags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.EditGFlags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditGFlags`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseApi.EditGFlags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**uniUUID** | [**string**](.md) | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditGFlagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **universeEditGFlags** | [**UniverseEditGFlags**](UniverseEditGFlags.md) |  | 

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


## EditUniverse

> YBPTask EditUniverse(ctx, cUUID, uniUUID).UniverseEditSpec(universeEditSpec).Execute()

Edit a YugabyteDB Universe



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
    cUUID := TODO // string | Customer UUID
    uniUUID := TODO // string | Universe UUID
    universeEditSpec := *openapiclient.NewUniverseEditSpec([]openapiclient.ClusterEditSpec{*openapiclient.NewClusterEditSpec("19ebde21-d537-47dc-8fab-3edc243c6f68")}, int32(123)) // UniverseEditSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.EditUniverse(context.Background(), cUUID, uniUUID).UniverseEditSpec(universeEditSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.EditUniverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditUniverse`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseApi.EditUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**uniUUID** | [**string**](.md) | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditUniverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **universeEditSpec** | [**UniverseEditSpec**](UniverseEditSpec.md) |  | 

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


## GetUniverse

> Universe GetUniverse(ctx, cUUID, uniUUID).Execute()

Get a YugabyteDB Universe



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
    cUUID := TODO // string | Customer UUID
    uniUUID := TODO // string | Universe UUID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.GetUniverse(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.GetUniverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUniverse`: Universe
    fmt.Fprintf(os.Stdout, "Response from `UniverseApi.GetUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**uniUUID** | [**string**](.md) | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Universe**](Universe.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

