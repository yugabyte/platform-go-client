# \UniverseAPI

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCluster**](UniverseAPI.md#AddCluster) | **Post** /customers/{cUUID}/universes/{uniUUID}/clusters | Add a cluster to a YugabyteDB Universe
[**AttachUniverse**](UniverseAPI.md#AttachUniverse) | **Post** /customers/{cUUID}/universes/{uniUUID}/attach | Attach universe
[**ConfigureMetricsExport**](UniverseAPI.md#ConfigureMetricsExport) | **Post** /customers/{cUUID}/universes/{uniUUID}/metrics-export-config | Configure metrics export
[**ConfigureQueryLogging**](UniverseAPI.md#ConfigureQueryLogging) | **Post** /customers/{cUUID}/universes/{uniUUID}/query-log-config | Configure Query Log for YugabyteDB Universe
[**CreateUniverse**](UniverseAPI.md#CreateUniverse) | **Post** /customers/{cUUID}/universes | Create a YugabyteDB Universe
[**DeleteAttachDetachMetadata**](UniverseAPI.md#DeleteAttachDetachMetadata) | **Delete** /customers/{cUUID}/universes/{uniUUID}/attach-detach-metadata | Delete attach/detach metadata
[**DeleteCluster**](UniverseAPI.md#DeleteCluster) | **Delete** /customers/{cUUID}/universes/{uniUUID}/clusters/{clsUUID} | Delete an additional cluster(s) of a YugabyteDB Universe
[**DeleteUniverse**](UniverseAPI.md#DeleteUniverse) | **Delete** /customers/{cUUID}/universes/{uniUUID} | Delete a universe
[**DetachUniverse**](UniverseAPI.md#DetachUniverse) | **Post** /customers/{cUUID}/universes/{uniUUID}/detach | Detach universe
[**EditGFlags**](UniverseAPI.md#EditGFlags) | **Post** /customers/{cUUID}/universes/{uniUUID}/gflags | Edit GFlags
[**EditKubernetesOverrides**](UniverseAPI.md#EditKubernetesOverrides) | **Post** /customers/{cUUID}/universes/{uniUUID}/kubernetes-overrides | Edit Kubernetes Helm Overrides
[**EditUniverse**](UniverseAPI.md#EditUniverse) | **Put** /customers/{cUUID}/universes/{uniUUID} | Edit a YugabyteDB Universe
[**EncryptionInTransitCertRotate**](UniverseAPI.md#EncryptionInTransitCertRotate) | **Post** /customers/{cUUID}/universes/{uniUUID}/encryption/in-transit/rotate | Rotate TLS Certs
[**EncryptionInTransitToggle**](UniverseAPI.md#EncryptionInTransitToggle) | **Post** /customers/{cUUID}/universes/{uniUUID}/encryption/in-transit | Enable or disable encryption in transit
[**FinalizeSoftwareUpgrade**](UniverseAPI.md#FinalizeSoftwareUpgrade) | **Post** /customers/{cUUID}/universes/{uniUUID}/upgrade/software/finalize | Finalize the Upgrade YugabyteDB
[**GetFinalizeSoftwareUpgradeInfo**](UniverseAPI.md#GetFinalizeSoftwareUpgradeInfo) | **Get** /customers/{cUUID}/universes/{uniUUID}/upgrade/software/finalize | Get finalize information on the YugabyteDB upgrade
[**GetUniverse**](UniverseAPI.md#GetUniverse) | **Get** /customers/{cUUID}/universes/{uniUUID} | Get a YugabyteDB Universe
[**GetUniverseResources**](UniverseAPI.md#GetUniverseResources) | **Post** /customers/{cUUID}/fetch-universe-resources | Get resource utilisation of a YugabyteDB Universe
[**OperatorImportUniverse**](UniverseAPI.md#OperatorImportUniverse) | **Post** /customers/{cUUID}/universes/{uniUUID}/operator-import | Import universe to operator
[**OperatorImportUniversePrecheck**](UniverseAPI.md#OperatorImportUniversePrecheck) | **Post** /customers/{cUUID}/universes/{uniUUID}/operator-import/precheck | Precheck universe import to operator
[**PrecheckSoftwareUpgrade**](UniverseAPI.md#PrecheckSoftwareUpgrade) | **Post** /customers/{cUUID}/universes/{uniUUID}/upgrade/software/precheck | Precheck YugabyteDB version upgrade
[**RestartUniverse**](UniverseAPI.md#RestartUniverse) | **Post** /customers/{cUUID}/universes/{uniUUID}/restart | Restart a YugabyteDB Universe
[**RollbackDetachUniverse**](UniverseAPI.md#RollbackDetachUniverse) | **Delete** /customers/{cUUID}/universes/{uniUUID}/detach | Rollback detach universe
[**RollbackSoftwareUpgrade**](UniverseAPI.md#RollbackSoftwareUpgrade) | **Post** /customers/{cUUID}/universes/{uniUUID}/upgrade/software/rollback | Rollback YugabyteDB version
[**StartSoftwareUpgrade**](UniverseAPI.md#StartSoftwareUpgrade) | **Post** /customers/{cUUID}/universes/{uniUUID}/upgrade/software | Upgrade YugabyteDB version
[**StartThirdPartySoftwareUpgrade**](UniverseAPI.md#StartThirdPartySoftwareUpgrade) | **Post** /customers/{cUUID}/universes/{uniUUID}/upgrade/third-party-software | Upgrade third party software
[**SystemdEnable**](UniverseAPI.md#SystemdEnable) | **Post** /customers/{cUUID}/universes/{uniUUID}/systemd | Migrate to Systemd controlled services



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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	clusterAddSpec := *openapiclient.NewClusterAddSpec("READ_REPLICA", int32(3), *openapiclient.NewClusterNodeSpec(), *openapiclient.NewClusterProviderEditSpec([]string{"RegionList_example"})) // ClusterAddSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.AddCluster(context.Background(), cUUID, uniUUID).ClusterAddSpec(clusterAddSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.AddCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddCluster`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.AddCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	attachUniverseSpec := *openapiclient.NewAttachUniverseSpec() // AttachUniverseSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UniverseAPI.AttachUniverse(context.Background(), cUUID, uniUUID).AttachUniverseSpec(attachUniverseSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.AttachUniverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

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

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigureMetricsExport

> YBATask ConfigureMetricsExport(ctx, cUUID, uniUUID).ConfigureMetricsExportSpec(configureMetricsExportSpec).Execute()

Configure metrics export



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	configureMetricsExportSpec := *openapiclient.NewConfigureMetricsExportSpec() // ConfigureMetricsExportSpec | ConfigureMetricsExportReq  Payload to configure metrics export. Part of ConfigureMetricsExportReq  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.ConfigureMetricsExport(context.Background(), cUUID, uniUUID).ConfigureMetricsExportSpec(configureMetricsExportSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.ConfigureMetricsExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigureMetricsExport`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.ConfigureMetricsExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigureMetricsExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **configureMetricsExportSpec** | [**ConfigureMetricsExportSpec**](ConfigureMetricsExportSpec.md) | ConfigureMetricsExportReq  Payload to configure metrics export. Part of ConfigureMetricsExportReq  | 

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


## ConfigureQueryLogging

> YBATask ConfigureQueryLogging(ctx, cUUID, uniUUID).UniverseQueryLogsExport(universeQueryLogsExport).Execute()

Configure Query Log for YugabyteDB Universe



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	universeQueryLogsExport := *openapiclient.NewUniverseQueryLogsExport(false, *openapiclient.NewQueryLogConfig([]openapiclient.UniverseQueryLogsExporterConfig{*openapiclient.NewUniverseQueryLogsExporterConfig("ExporterUuid_example")})) // UniverseQueryLogsExport | UniverseQueryLogsExportReq  Payload to configure export of query logs. Part of UniverseQueryLogsExportReq  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.ConfigureQueryLogging(context.Background(), cUUID, uniUUID).UniverseQueryLogsExport(universeQueryLogsExport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.ConfigureQueryLogging``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigureQueryLogging`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.ConfigureQueryLogging`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigureQueryLoggingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **universeQueryLogsExport** | [**UniverseQueryLogsExport**](UniverseQueryLogsExport.md) | UniverseQueryLogsExportReq  Payload to configure export of query logs. Part of UniverseQueryLogsExportReq  | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	universeCreateSpec := *openapiclient.NewUniverseCreateSpec(*openapiclient.NewUniverseSpec("my-yb-universe", "2024.2.0.0-b600", []openapiclient.ClusterSpec{*openapiclient.NewClusterSpec("PRIMARY", int32(3), *openapiclient.NewClusterNodeSpec(), *openapiclient.NewClusterProviderSpec("89a46d52-4edd-4736-922a-35177a0b990c"))}), "x86_64") // UniverseCreateSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.CreateUniverse(context.Background(), cUUID).UniverseCreateSpec(universeCreateSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.CreateUniverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUniverse`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.CreateUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UniverseAPI.DeleteAttachDetachMetadata(context.Background(), cUUID, uniUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.DeleteAttachDetachMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	clsUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cluster UUID
	isForceDelete := true // bool | Whether to force delete the cluster (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.DeleteCluster(context.Background(), cUUID, uniUUID, clsUUID).IsForceDelete(isForceDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.DeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCluster`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.DeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 
**clsUUID** | **string** | Cluster UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	universeDeleteSpec := *openapiclient.NewUniverseDeleteSpec() // UniverseDeleteSpec |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.DeleteUniverse(context.Background(), cUUID, uniUUID).UniverseDeleteSpec(universeDeleteSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.DeleteUniverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUniverse`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.DeleteUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	detachUniverseSpec := *openapiclient.NewDetachUniverseSpec() // DetachUniverseSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.DetachUniverse(context.Background(), cUUID, uniUUID).DetachUniverseSpec(detachUniverseSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.DetachUniverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachUniverse`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.DetachUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	universeEditGFlags := *openapiclient.NewUniverseEditGFlags() // UniverseEditGFlags | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.EditGFlags(context.Background(), cUUID, uniUUID).UniverseEditGFlags(universeEditGFlags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.EditGFlags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditGFlags`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.EditGFlags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	universeEditKubernetesOverrides := *openapiclient.NewUniverseEditKubernetesOverrides() // UniverseEditKubernetesOverrides | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.EditKubernetesOverrides(context.Background(), cUUID, uniUUID).UniverseEditKubernetesOverrides(universeEditKubernetesOverrides).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.EditKubernetesOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditKubernetesOverrides`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.EditKubernetesOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	universeEditSpec := *openapiclient.NewUniverseEditSpec([]openapiclient.ClusterEditSpec{*openapiclient.NewClusterEditSpec("19ebde21-d537-47dc-8fab-3edc243c6f68")}, int32(123)) // UniverseEditSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.EditUniverse(context.Background(), cUUID, uniUUID).UniverseEditSpec(universeEditSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.EditUniverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditUniverse`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.EditUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	universeCertRotateSpec := *openapiclient.NewUniverseCertRotateSpec() // UniverseCertRotateSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.EncryptionInTransitCertRotate(context.Background(), cUUID, uniUUID).UniverseCertRotateSpec(universeCertRotateSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.EncryptionInTransitCertRotate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EncryptionInTransitCertRotate`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.EncryptionInTransitCertRotate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	universeEditEncryptionInTransit := *openapiclient.NewUniverseEditEncryptionInTransit() // UniverseEditEncryptionInTransit | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.EncryptionInTransitToggle(context.Background(), cUUID, uniUUID).UniverseEditEncryptionInTransit(universeEditEncryptionInTransit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.EncryptionInTransitToggle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EncryptionInTransitToggle`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.EncryptionInTransitToggle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	universeSoftwareUpgradeFinalize := *openapiclient.NewUniverseSoftwareUpgradeFinalize() // UniverseSoftwareUpgradeFinalize |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.FinalizeSoftwareUpgrade(context.Background(), cUUID, uniUUID).UniverseSoftwareUpgradeFinalize(universeSoftwareUpgradeFinalize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.FinalizeSoftwareUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FinalizeSoftwareUpgrade`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.FinalizeSoftwareUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.GetFinalizeSoftwareUpgradeInfo(context.Background(), cUUID, uniUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetFinalizeSoftwareUpgradeInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinalizeSoftwareUpgradeInfo`: UniverseSoftwareUpgradeFinalizeInfo
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetFinalizeSoftwareUpgradeInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.GetUniverse(context.Background(), cUUID, uniUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverse`: Universe
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	universeCreateSpec := *openapiclient.NewUniverseCreateSpec(*openapiclient.NewUniverseSpec("my-yb-universe", "2024.2.0.0-b600", []openapiclient.ClusterSpec{*openapiclient.NewClusterSpec("PRIMARY", int32(3), *openapiclient.NewClusterNodeSpec(), *openapiclient.NewClusterProviderSpec("89a46d52-4edd-4736-922a-35177a0b990c"))}), "x86_64") // UniverseCreateSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.GetUniverseResources(context.Background(), cUUID).UniverseCreateSpec(universeCreateSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.GetUniverseResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseResources`: UniverseResourceDetails
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.GetUniverseResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

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


## OperatorImportUniverse

> YBATask OperatorImportUniverse(ctx, cUUID, uniUUID).UniverseOperatorImportReq(universeOperatorImportReq).Execute()

Import universe to operator



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	universeOperatorImportReq := *openapiclient.NewUniverseOperatorImportReq() // UniverseOperatorImportReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.OperatorImportUniverse(context.Background(), cUUID, uniUUID).UniverseOperatorImportReq(universeOperatorImportReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.OperatorImportUniverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperatorImportUniverse`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.OperatorImportUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorImportUniverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **universeOperatorImportReq** | [**UniverseOperatorImportReq**](UniverseOperatorImportReq.md) |  | 

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


## OperatorImportUniversePrecheck

> OperatorImportUniversePrecheck(ctx, cUUID, uniUUID).UniverseOperatorImportReq(universeOperatorImportReq).Execute()

Precheck universe import to operator



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	universeOperatorImportReq := *openapiclient.NewUniverseOperatorImportReq() // UniverseOperatorImportReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UniverseAPI.OperatorImportUniversePrecheck(context.Background(), cUUID, uniUUID).UniverseOperatorImportReq(universeOperatorImportReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.OperatorImportUniversePrecheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperatorImportUniversePrecheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **universeOperatorImportReq** | [**UniverseOperatorImportReq**](UniverseOperatorImportReq.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	universeSoftwareUpgradePrecheckReq := *openapiclient.NewUniverseSoftwareUpgradePrecheckReq("YbSoftwareVersion_example") // UniverseSoftwareUpgradePrecheckReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.PrecheckSoftwareUpgrade(context.Background(), cUUID, uniUUID).UniverseSoftwareUpgradePrecheckReq(universeSoftwareUpgradePrecheckReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.PrecheckSoftwareUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrecheckSoftwareUpgrade`: UniverseSoftwareUpgradePrecheckResp
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.PrecheckSoftwareUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	universeRestart := *openapiclient.NewUniverseRestart() // UniverseRestart |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.RestartUniverse(context.Background(), cUUID, uniUUID).UniverseRestart(universeRestart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.RestartUniverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestartUniverse`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.RestartUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

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


## RollbackDetachUniverse

> RollbackDetachUniverse(ctx, cUUID, uniUUID).IsForceRollback(isForceRollback).Execute()

Rollback detach universe



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	isForceRollback := true // bool | Force rollback without checking current owner or detached state. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UniverseAPI.RollbackDetachUniverse(context.Background(), cUUID, uniUUID).IsForceRollback(isForceRollback).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.RollbackDetachUniverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackDetachUniverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isForceRollback** | **bool** | Force rollback without checking current owner or detached state. | [default to false]

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	universeRollbackUpgradeReq := *openapiclient.NewUniverseRollbackUpgradeReq() // UniverseRollbackUpgradeReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.RollbackSoftwareUpgrade(context.Background(), cUUID, uniUUID).UniverseRollbackUpgradeReq(universeRollbackUpgradeReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.RollbackSoftwareUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RollbackSoftwareUpgrade`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.RollbackSoftwareUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	universeSoftwareUpgradeStart := *openapiclient.NewUniverseSoftwareUpgradeStart("Version_example") // UniverseSoftwareUpgradeStart | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.StartSoftwareUpgrade(context.Background(), cUUID, uniUUID).UniverseSoftwareUpgradeStart(universeSoftwareUpgradeStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.StartSoftwareUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartSoftwareUpgrade`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.StartSoftwareUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	universeThirdPartySoftwareUpgradeStart := *openapiclient.NewUniverseThirdPartySoftwareUpgradeStart() // UniverseThirdPartySoftwareUpgradeStart |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.StartThirdPartySoftwareUpgrade(context.Background(), cUUID, uniUUID).UniverseThirdPartySoftwareUpgradeStart(universeThirdPartySoftwareUpgradeStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.StartThirdPartySoftwareUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartThirdPartySoftwareUpgrade`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.StartThirdPartySoftwareUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	universeSystemdEnableStart := *openapiclient.NewUniverseSystemdEnableStart() // UniverseSystemdEnableStart |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseAPI.SystemdEnable(context.Background(), cUUID, uniUUID).UniverseSystemdEnableStart(universeSystemdEnableStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseAPI.SystemdEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemdEnable`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `UniverseAPI.SystemdEnable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

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

