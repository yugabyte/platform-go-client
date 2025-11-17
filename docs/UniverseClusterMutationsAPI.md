# \UniverseClusterMutationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAllClusters**](UniverseClusterMutationsAPI.md#CreateAllClusters) | **Post** /api/v1/customers/{cUUID}/universes/clusters | Create Universe Clusters
[**CreateReadOnlyCluster**](UniverseClusterMutationsAPI.md#CreateReadOnlyCluster) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/clusters/read_only | Create ReadOnly Cluster
[**DeleteReadonlyCluster**](UniverseClusterMutationsAPI.md#DeleteReadonlyCluster) | **Delete** /api/v1/customers/{cUUID}/universes/{uniUUID}/clusters/read_only/{clustUUID} | Delete Readonly Cluster
[**UpdatePrimaryCluster**](UniverseClusterMutationsAPI.md#UpdatePrimaryCluster) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/clusters/primary | Update Primary Cluster
[**UpdateReadOnlyCluster**](UniverseClusterMutationsAPI.md#UpdateReadOnlyCluster) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/clusters/read_only | Update Readonly Cluster



## CreateAllClusters

> YBPTask CreateAllClusters(ctx, cUUID).UniverseConfigureTaskParams(universeConfigureTaskParams).Request(request).Execute()

Create Universe Clusters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	universeConfigureTaskParams := *openapiclient.NewUniverseConfigureTaskParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), "PlatformUrl_example", "RawClientRootCA_example", int32(123), int32(123)) // UniverseConfigureTaskParams | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseClusterMutationsAPI.CreateAllClusters(context.Background(), cUUID).UniverseConfigureTaskParams(universeConfigureTaskParams).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseClusterMutationsAPI.CreateAllClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAllClusters`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `UniverseClusterMutationsAPI.CreateAllClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAllClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **universeConfigureTaskParams** | [**UniverseConfigureTaskParams**](UniverseConfigureTaskParams.md) |  | 
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


## CreateReadOnlyCluster

> YBPTask CreateReadOnlyCluster(ctx, cUUID, uniUUID).UniverseConfigureTaskParams(universeConfigureTaskParams).Request(request).Execute()

Create ReadOnly Cluster



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	universeConfigureTaskParams := *openapiclient.NewUniverseDefinitionTaskParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), "PlatformUrl_example", "RawClientRootCA_example", int32(123), int32(123)) // UniverseDefinitionTaskParams | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseClusterMutationsAPI.CreateReadOnlyCluster(context.Background(), cUUID, uniUUID).UniverseConfigureTaskParams(universeConfigureTaskParams).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseClusterMutationsAPI.CreateReadOnlyCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReadOnlyCluster`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `UniverseClusterMutationsAPI.CreateReadOnlyCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReadOnlyClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **universeConfigureTaskParams** | [**UniverseDefinitionTaskParams**](UniverseDefinitionTaskParams.md) |  | 
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


## DeleteReadonlyCluster

> YBPTask DeleteReadonlyCluster(ctx, cUUID, uniUUID, clustUUID).IsForceDelete(isForceDelete).Request(request).Execute()

Delete Readonly Cluster



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	clustUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	isForceDelete := true // bool |  (optional) (default to false)
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseClusterMutationsAPI.DeleteReadonlyCluster(context.Background(), cUUID, uniUUID, clustUUID).IsForceDelete(isForceDelete).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseClusterMutationsAPI.DeleteReadonlyCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteReadonlyCluster`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `UniverseClusterMutationsAPI.DeleteReadonlyCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 
**clustUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReadonlyClusterRequest struct via the builder pattern


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


## UpdatePrimaryCluster

> YBPTask UpdatePrimaryCluster(ctx, cUUID, uniUUID).UniverseConfigureTaskParams(universeConfigureTaskParams).Request(request).Execute()

Update Primary Cluster



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	universeConfigureTaskParams := *openapiclient.NewUniverseConfigureTaskParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), "PlatformUrl_example", "RawClientRootCA_example", int32(123), int32(123)) // UniverseConfigureTaskParams | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseClusterMutationsAPI.UpdatePrimaryCluster(context.Background(), cUUID, uniUUID).UniverseConfigureTaskParams(universeConfigureTaskParams).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseClusterMutationsAPI.UpdatePrimaryCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePrimaryCluster`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `UniverseClusterMutationsAPI.UpdatePrimaryCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrimaryClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **universeConfigureTaskParams** | [**UniverseConfigureTaskParams**](UniverseConfigureTaskParams.md) |  | 
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


## UpdateReadOnlyCluster

> YBPTask UpdateReadOnlyCluster(ctx, cUUID, uniUUID).UniverseConfigureTaskParams(universeConfigureTaskParams).Request(request).Execute()

Update Readonly Cluster



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	universeConfigureTaskParams := *openapiclient.NewUniverseConfigureTaskParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), "PlatformUrl_example", "RawClientRootCA_example", int32(123), int32(123)) // UniverseConfigureTaskParams | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseClusterMutationsAPI.UpdateReadOnlyCluster(context.Background(), cUUID, uniUUID).UniverseConfigureTaskParams(universeConfigureTaskParams).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseClusterMutationsAPI.UpdateReadOnlyCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateReadOnlyCluster`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `UniverseClusterMutationsAPI.UpdateReadOnlyCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReadOnlyClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **universeConfigureTaskParams** | [**UniverseConfigureTaskParams**](UniverseConfigureTaskParams.md) |  | 
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

