# \UniverseApi

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCluster**](UniverseApi.md#AddCluster) | **Post** /customers/{cUUID}/universes/{uniUUID}/clusters | Add a cluster to a YugabyteDB Universe
[**AttachUniverse**](UniverseApi.md#AttachUniverse) | **Post** /customers/{cUUID}/universes/{uniUUID}/attach | Attach universe
[**CreateUniverse**](UniverseApi.md#CreateUniverse) | **Post** /customers/{cUUID}/universes | Create a YugabyteDB Universe
[**DeleteAttachDetachMetadata**](UniverseApi.md#DeleteAttachDetachMetadata) | **Delete** /customers/{cUUID}/universes/{uniUUID}/attach-detach-metadata | Delete attach/detach metadata
[**DeleteCluster**](UniverseApi.md#DeleteCluster) | **Delete** /customers/{cUUID}/universes/{uniUUID}/clusters/{clsUUID} | Delete an additional cluster(s) of a YugabyteDB Universe
[**DeleteUniverse**](UniverseApi.md#DeleteUniverse) | **Delete** /customers/{cUUID}/universes/{uniUUID} | Delete a universe
[**DetachUniverse**](UniverseApi.md#DetachUniverse) | **Post** /customers/{cUUID}/universes/{uniUUID}/detach | Detach universe
[**EditGFlags**](UniverseApi.md#EditGFlags) | **Post** /customers/{cUUID}/universes/{uniUUID}/gflags | Edit GFlags
[**EditKubernetesOverrides**](UniverseApi.md#EditKubernetesOverrides) | **Post** /customers/{cUUID}/universes/{uniUUID}/kubernetes-overrides | Edit Kubernetes Helm Overrides
[**EditUniverse**](UniverseApi.md#EditUniverse) | **Put** /customers/{cUUID}/universes/{uniUUID} | Edit a YugabyteDB Universe
[**EncryptionInTransitCertRotate**](UniverseApi.md#EncryptionInTransitCertRotate) | **Post** /customers/{cUUID}/universes/{uniUUID}/encryption/in-transit/rotate | Rotate TLS Certs
[**EncryptionInTransitToggle**](UniverseApi.md#EncryptionInTransitToggle) | **Post** /customers/{cUUID}/universes/{uniUUID}/encryption/in-transit | Enable or disable encryption in transit
[**FinalizeSoftwareUpgrade**](UniverseApi.md#FinalizeSoftwareUpgrade) | **Post** /customers/{cUUID}/universes/{uniUUID}/upgrade/software/finalize | Finalize the Upgrade YugabyteDB
[**GetFinalizeSoftwareUpgradeInfo**](UniverseApi.md#GetFinalizeSoftwareUpgradeInfo) | **Get** /customers/{cUUID}/universes/{uniUUID}/upgrade/software/finalize | Get finalize information on the YugabyteDB upgrade
[**GetUniverse**](UniverseApi.md#GetUniverse) | **Get** /customers/{cUUID}/universes/{uniUUID} | Get a YugabyteDB Universe
[**GetUniverseResources**](UniverseApi.md#GetUniverseResources) | **Post** /customers/{cUUID}/fetch-universe-resources | Get resource utilisation of a YugabyteDB Universe
[**PrecheckSoftwareUpgrade**](UniverseApi.md#PrecheckSoftwareUpgrade) | **Post** /customers/{cUUID}/universes/{uniUUID}/upgrade/software/precheck | Precheck YugabyteDB version upgrade
[**RestartUniverse**](UniverseApi.md#RestartUniverse) | **Post** /customers/{cUUID}/universes/{uniUUID}/restart | Restart a YugabyteDB Universe
[**RollbackSoftwareUpgrade**](UniverseApi.md#RollbackSoftwareUpgrade) | **Post** /customers/{cUUID}/universes/{uniUUID}/upgrade/software/rollback | Rollback YugabyteDB version
[**StartSoftwareUpgrade**](UniverseApi.md#StartSoftwareUpgrade) | **Post** /customers/{cUUID}/universes/{uniUUID}/upgrade/software | Upgrade YugabyteDB version
[**StartThirdPartySoftwareUpgrade**](UniverseApi.md#StartThirdPartySoftwareUpgrade) | **Post** /customers/{cUUID}/universes/{uniUUID}/upgrade/third-party-software | Upgrade third party software
[**SystemdEnable**](UniverseApi.md#SystemdEnable) | **Post** /customers/{cUUID}/universes/{uniUUID}/systemd | Migrate to Systemd controlled services



## AddCluster

> YBATask AddCluster(ctx, cUUID, uniUUID).ClusterAddSpec(clusterAddSpec).Execute()

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
    clusterAddSpec := *openapiclient.NewClusterAddSpec("READ_REPLICA", int32(3), *openapiclient.NewClusterNodeSpec(), *openapiclient.NewClusterProviderEditSpec([]string{"RegionList_example"})) // ClusterAddSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.AddCluster(context.Background(), cUUID, uniUUID).ClusterAddSpec(clusterAddSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.AddCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCluster`: YBATask
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

[**YBATask**](YBATask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachUniverse

> AttachUniverse(ctx, cUUID, uniUUID).AttachUniverseSpec(attachUniverseSpec).Execute()

Attach universe



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
    attachUniverseSpec := *openapiclient.NewAttachUniverseSpec() // AttachUniverseSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.AttachUniverse(context.Background(), cUUID, uniUUID).AttachUniverseSpec(attachUniverseSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.AttachUniverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**uniUUID** | [**string**](.md) | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachUniverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **attachUniverseSpec** | [**AttachUniverseSpec**](AttachUniverseSpec.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUniverse

> YBATask CreateUniverse(ctx, cUUID).UniverseCreateSpec(universeCreateSpec).Execute()

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
    universeCreateSpec := *openapiclient.NewUniverseCreateSpec(*openapiclient.NewUniverseSpec("my-yb-universe", "2024.2.0.0-b600", []openapiclient.ClusterSpec{*openapiclient.NewClusterSpec("PRIMARY", int32(3), *openapiclient.NewClusterNodeSpec(), *openapiclient.NewClusterProviderSpec("89a46d52-4edd-4736-922a-35177a0b990c"))}), "x86_64") // UniverseCreateSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.CreateUniverse(context.Background(), cUUID).UniverseCreateSpec(universeCreateSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.CreateUniverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUniverse`: YBATask
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

[**YBATask**](YBATask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAttachDetachMetadata

> DeleteAttachDetachMetadata(ctx, cUUID, uniUUID).Execute()

Delete attach/detach metadata



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
    resp, r, err := api_client.UniverseApi.DeleteAttachDetachMetadata(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.DeleteAttachDetachMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**uniUUID** | [**string**](.md) | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttachDetachMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCluster

> YBATask DeleteCluster(ctx, cUUID, uniUUID, clsUUID).IsForceDelete(isForceDelete).Execute()

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
    // response from `DeleteCluster`: YBATask
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

[**YBATask**](YBATask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUniverse

> YBATask DeleteUniverse(ctx, cUUID, uniUUID).UniverseDeleteSpec(universeDeleteSpec).Execute()

Delete a universe



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
    universeDeleteSpec := *openapiclient.NewUniverseDeleteSpec() // UniverseDeleteSpec |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.DeleteUniverse(context.Background(), cUUID, uniUUID).UniverseDeleteSpec(universeDeleteSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.DeleteUniverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUniverse`: YBATask
    fmt.Fprintf(os.Stdout, "Response from `UniverseApi.DeleteUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**uniUUID** | [**string**](.md) | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUniverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **universeDeleteSpec** | [**UniverseDeleteSpec**](UniverseDeleteSpec.md) |  | 

### Return type

[**YBATask**](YBATask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachUniverse

> *os.File DetachUniverse(ctx, cUUID, uniUUID).DetachUniverseSpec(detachUniverseSpec).Execute()

Detach universe



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
    detachUniverseSpec := *openapiclient.NewDetachUniverseSpec() // DetachUniverseSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.DetachUniverse(context.Background(), cUUID, uniUUID).DetachUniverseSpec(detachUniverseSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.DetachUniverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachUniverse`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `UniverseApi.DetachUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**uniUUID** | [**string**](.md) | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachUniverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **detachUniverseSpec** | [**DetachUniverseSpec**](DetachUniverseSpec.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditGFlags

> YBATask EditGFlags(ctx, cUUID, uniUUID).UniverseEditGFlags(universeEditGFlags).Execute()

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
    // response from `EditGFlags`: YBATask
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

[**YBATask**](YBATask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditKubernetesOverrides

> YBATask EditKubernetesOverrides(ctx, cUUID, uniUUID).UniverseEditKubernetesOverrides(universeEditKubernetesOverrides).Execute()

Edit Kubernetes Helm Overrides



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
    universeEditKubernetesOverrides := *openapiclient.NewUniverseEditKubernetesOverrides() // UniverseEditKubernetesOverrides | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.EditKubernetesOverrides(context.Background(), cUUID, uniUUID).UniverseEditKubernetesOverrides(universeEditKubernetesOverrides).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.EditKubernetesOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditKubernetesOverrides`: YBATask
    fmt.Fprintf(os.Stdout, "Response from `UniverseApi.EditKubernetesOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**uniUUID** | [**string**](.md) | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditKubernetesOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **universeEditKubernetesOverrides** | [**UniverseEditKubernetesOverrides**](UniverseEditKubernetesOverrides.md) |  | 

### Return type

[**YBATask**](YBATask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditUniverse

> YBATask EditUniverse(ctx, cUUID, uniUUID).UniverseEditSpec(universeEditSpec).Execute()

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
    // response from `EditUniverse`: YBATask
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

[**YBATask**](YBATask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EncryptionInTransitCertRotate

> YBATask EncryptionInTransitCertRotate(ctx, cUUID, uniUUID).UniverseCertRotateSpec(universeCertRotateSpec).Execute()

Rotate TLS Certs



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
    universeCertRotateSpec := *openapiclient.NewUniverseCertRotateSpec() // UniverseCertRotateSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.EncryptionInTransitCertRotate(context.Background(), cUUID, uniUUID).UniverseCertRotateSpec(universeCertRotateSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.EncryptionInTransitCertRotate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EncryptionInTransitCertRotate`: YBATask
    fmt.Fprintf(os.Stdout, "Response from `UniverseApi.EncryptionInTransitCertRotate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**uniUUID** | [**string**](.md) | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEncryptionInTransitCertRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **universeCertRotateSpec** | [**UniverseCertRotateSpec**](UniverseCertRotateSpec.md) |  | 

### Return type

[**YBATask**](YBATask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EncryptionInTransitToggle

> YBATask EncryptionInTransitToggle(ctx, cUUID, uniUUID).UniverseEditEncryptionInTransit(universeEditEncryptionInTransit).Execute()

Enable or disable encryption in transit



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
    universeEditEncryptionInTransit := *openapiclient.NewUniverseEditEncryptionInTransit() // UniverseEditEncryptionInTransit | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.EncryptionInTransitToggle(context.Background(), cUUID, uniUUID).UniverseEditEncryptionInTransit(universeEditEncryptionInTransit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.EncryptionInTransitToggle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EncryptionInTransitToggle`: YBATask
    fmt.Fprintf(os.Stdout, "Response from `UniverseApi.EncryptionInTransitToggle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**uniUUID** | [**string**](.md) | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEncryptionInTransitToggleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **universeEditEncryptionInTransit** | [**UniverseEditEncryptionInTransit**](UniverseEditEncryptionInTransit.md) |  | 

### Return type

[**YBATask**](YBATask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinalizeSoftwareUpgrade

> YBATask FinalizeSoftwareUpgrade(ctx, cUUID, uniUUID).UniverseSoftwareUpgradeFinalize(universeSoftwareUpgradeFinalize).Execute()

Finalize the Upgrade YugabyteDB



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
    universeSoftwareUpgradeFinalize := *openapiclient.NewUniverseSoftwareUpgradeFinalize() // UniverseSoftwareUpgradeFinalize |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.FinalizeSoftwareUpgrade(context.Background(), cUUID, uniUUID).UniverseSoftwareUpgradeFinalize(universeSoftwareUpgradeFinalize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.FinalizeSoftwareUpgrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FinalizeSoftwareUpgrade`: YBATask
    fmt.Fprintf(os.Stdout, "Response from `UniverseApi.FinalizeSoftwareUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**uniUUID** | [**string**](.md) | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFinalizeSoftwareUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **universeSoftwareUpgradeFinalize** | [**UniverseSoftwareUpgradeFinalize**](UniverseSoftwareUpgradeFinalize.md) |  | 

### Return type

[**YBATask**](YBATask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinalizeSoftwareUpgradeInfo

> UniverseSoftwareUpgradeFinalizeInfo GetFinalizeSoftwareUpgradeInfo(ctx, cUUID, uniUUID).Execute()

Get finalize information on the YugabyteDB upgrade



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
    resp, r, err := api_client.UniverseApi.GetFinalizeSoftwareUpgradeInfo(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.GetFinalizeSoftwareUpgradeInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFinalizeSoftwareUpgradeInfo`: UniverseSoftwareUpgradeFinalizeInfo
    fmt.Fprintf(os.Stdout, "Response from `UniverseApi.GetFinalizeSoftwareUpgradeInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**uniUUID** | [**string**](.md) | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinalizeSoftwareUpgradeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UniverseSoftwareUpgradeFinalizeInfo**](UniverseSoftwareUpgradeFinalizeInfo.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
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


## GetUniverseResources

> UniverseResourceDetails GetUniverseResources(ctx, cUUID).UniverseCreateSpec(universeCreateSpec).Execute()

Get resource utilisation of a YugabyteDB Universe



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
    universeCreateSpec := *openapiclient.NewUniverseCreateSpec(*openapiclient.NewUniverseSpec("my-yb-universe", "2024.2.0.0-b600", []openapiclient.ClusterSpec{*openapiclient.NewClusterSpec("PRIMARY", int32(3), *openapiclient.NewClusterNodeSpec(), *openapiclient.NewClusterProviderSpec("89a46d52-4edd-4736-922a-35177a0b990c"))}), "x86_64") // UniverseCreateSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.GetUniverseResources(context.Background(), cUUID).UniverseCreateSpec(universeCreateSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.GetUniverseResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUniverseResources`: UniverseResourceDetails
    fmt.Fprintf(os.Stdout, "Response from `UniverseApi.GetUniverseResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **universeCreateSpec** | [**UniverseCreateSpec**](UniverseCreateSpec.md) |  | 

### Return type

[**UniverseResourceDetails**](UniverseResourceDetails.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrecheckSoftwareUpgrade

> UniverseSoftwareUpgradePrecheckResp PrecheckSoftwareUpgrade(ctx, cUUID, uniUUID).UniverseSoftwareUpgradePrecheckReq(universeSoftwareUpgradePrecheckReq).Execute()

Precheck YugabyteDB version upgrade



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
    universeSoftwareUpgradePrecheckReq := *openapiclient.NewUniverseSoftwareUpgradePrecheckReq("YbSoftwareVersion_example") // UniverseSoftwareUpgradePrecheckReq | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.PrecheckSoftwareUpgrade(context.Background(), cUUID, uniUUID).UniverseSoftwareUpgradePrecheckReq(universeSoftwareUpgradePrecheckReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.PrecheckSoftwareUpgrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrecheckSoftwareUpgrade`: UniverseSoftwareUpgradePrecheckResp
    fmt.Fprintf(os.Stdout, "Response from `UniverseApi.PrecheckSoftwareUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**uniUUID** | [**string**](.md) | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrecheckSoftwareUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **universeSoftwareUpgradePrecheckReq** | [**UniverseSoftwareUpgradePrecheckReq**](UniverseSoftwareUpgradePrecheckReq.md) |  | 

### Return type

[**UniverseSoftwareUpgradePrecheckResp**](UniverseSoftwareUpgradePrecheckResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartUniverse

> YBATask RestartUniverse(ctx, cUUID, uniUUID).UniverseRestart(universeRestart).Execute()

Restart a YugabyteDB Universe



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
    universeRestart := *openapiclient.NewUniverseRestart() // UniverseRestart |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.RestartUniverse(context.Background(), cUUID, uniUUID).UniverseRestart(universeRestart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.RestartUniverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestartUniverse`: YBATask
    fmt.Fprintf(os.Stdout, "Response from `UniverseApi.RestartUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**uniUUID** | [**string**](.md) | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartUniverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **universeRestart** | [**UniverseRestart**](UniverseRestart.md) |  | 

### Return type

[**YBATask**](YBATask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbackSoftwareUpgrade

> YBATask RollbackSoftwareUpgrade(ctx, cUUID, uniUUID).UniverseRollbackUpgradeReq(universeRollbackUpgradeReq).Execute()

Rollback YugabyteDB version



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
    universeRollbackUpgradeReq := *openapiclient.NewUniverseRollbackUpgradeReq() // UniverseRollbackUpgradeReq |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.RollbackSoftwareUpgrade(context.Background(), cUUID, uniUUID).UniverseRollbackUpgradeReq(universeRollbackUpgradeReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.RollbackSoftwareUpgrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RollbackSoftwareUpgrade`: YBATask
    fmt.Fprintf(os.Stdout, "Response from `UniverseApi.RollbackSoftwareUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**uniUUID** | [**string**](.md) | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackSoftwareUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **universeRollbackUpgradeReq** | [**UniverseRollbackUpgradeReq**](UniverseRollbackUpgradeReq.md) |  | 

### Return type

[**YBATask**](YBATask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartSoftwareUpgrade

> YBATask StartSoftwareUpgrade(ctx, cUUID, uniUUID).UniverseSoftwareUpgradeStart(universeSoftwareUpgradeStart).Execute()

Upgrade YugabyteDB version



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
    universeSoftwareUpgradeStart := *openapiclient.NewUniverseSoftwareUpgradeStart("Version_example") // UniverseSoftwareUpgradeStart | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.StartSoftwareUpgrade(context.Background(), cUUID, uniUUID).UniverseSoftwareUpgradeStart(universeSoftwareUpgradeStart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.StartSoftwareUpgrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartSoftwareUpgrade`: YBATask
    fmt.Fprintf(os.Stdout, "Response from `UniverseApi.StartSoftwareUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**uniUUID** | [**string**](.md) | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartSoftwareUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **universeSoftwareUpgradeStart** | [**UniverseSoftwareUpgradeStart**](UniverseSoftwareUpgradeStart.md) |  | 

### Return type

[**YBATask**](YBATask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartThirdPartySoftwareUpgrade

> YBATask StartThirdPartySoftwareUpgrade(ctx, cUUID, uniUUID).UniverseThirdPartySoftwareUpgradeStart(universeThirdPartySoftwareUpgradeStart).Execute()

Upgrade third party software



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
    universeThirdPartySoftwareUpgradeStart := *openapiclient.NewUniverseThirdPartySoftwareUpgradeStart() // UniverseThirdPartySoftwareUpgradeStart |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.StartThirdPartySoftwareUpgrade(context.Background(), cUUID, uniUUID).UniverseThirdPartySoftwareUpgradeStart(universeThirdPartySoftwareUpgradeStart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.StartThirdPartySoftwareUpgrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartThirdPartySoftwareUpgrade`: YBATask
    fmt.Fprintf(os.Stdout, "Response from `UniverseApi.StartThirdPartySoftwareUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**uniUUID** | [**string**](.md) | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartThirdPartySoftwareUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **universeThirdPartySoftwareUpgradeStart** | [**UniverseThirdPartySoftwareUpgradeStart**](UniverseThirdPartySoftwareUpgradeStart.md) |  | 

### Return type

[**YBATask**](YBATask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemdEnable

> YBATask SystemdEnable(ctx, cUUID, uniUUID).UniverseSystemdEnableStart(universeSystemdEnableStart).Execute()

Migrate to Systemd controlled services



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
    universeSystemdEnableStart := *openapiclient.NewUniverseSystemdEnableStart() // UniverseSystemdEnableStart |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseApi.SystemdEnable(context.Background(), cUUID, uniUUID).UniverseSystemdEnableStart(universeSystemdEnableStart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseApi.SystemdEnable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SystemdEnable`: YBATask
    fmt.Fprintf(os.Stdout, "Response from `UniverseApi.SystemdEnable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 
**uniUUID** | [**string**](.md) | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemdEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **universeSystemdEnableStart** | [**UniverseSystemdEnableStart**](UniverseSystemdEnableStart.md) |  | 

### Return type

[**YBATask**](YBATask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

