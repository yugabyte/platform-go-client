# \PITRManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneNamespace**](PITRManagementApi.md#CloneNamespace) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/clone | Clone namespace via PITR on a universe
[**CreatePitrConfig**](PITRManagementApi.md#CreatePitrConfig) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/keyspaces/{tableType}/{keyspaceName}/pitr_config | Create pitr config for a keyspace in a universe
[**DeletePitrConfig**](PITRManagementApi.md#DeletePitrConfig) | **Delete** /api/v1/customers/{cUUID}/universes/{uniUUID}/pitr_config/{pUUID} | Delete pitr config on a universe
[**ListOfPitrConfigs**](PITRManagementApi.md#ListOfPitrConfigs) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/pitr_config | List the PITR configs of a universe
[**PerformPitr**](PITRManagementApi.md#PerformPitr) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/pitr | Perform PITR on a universe
[**UpdatePitrConfig**](PITRManagementApi.md#UpdatePitrConfig) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/pitr_config/{pUUID} | Update pitr config for a keyspace in a universe



## CloneNamespace

> YBPTask CloneNamespace(ctx, cUUID, uniUUID).NamespaceClone(namespaceClone).Request(request).Execute()

Clone namespace via PITR on a universe



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
    namespaceClone := *openapiclient.NewCloneNamespaceParams(*openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), "PlatformUrl_example", int32(123), int32(123)) // CloneNamespaceParams | perform clone via PITR
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PITRManagementApi.CloneNamespace(context.Background(), cUUID, uniUUID).NamespaceClone(namespaceClone).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PITRManagementApi.CloneNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneNamespace`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `PITRManagementApi.CloneNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **namespaceClone** | [**CloneNamespaceParams**](CloneNamespaceParams.md) | perform clone via PITR | 
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


## CreatePitrConfig

> YBPTask CreatePitrConfig(ctx, cUUID, uniUUID, tableType, keyspaceName).PitrConfig(pitrConfig).Request(request).Execute()

Create pitr config for a keyspace in a universe

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
    tableType := "tableType_example" // string | 
    keyspaceName := "keyspaceName_example" // string | 
    pitrConfig := *openapiclient.NewCreatePitrConfigParams(*openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), "PlatformUrl_example", int32(123), int32(123)) // CreatePitrConfigParams | post pitr config
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PITRManagementApi.CreatePitrConfig(context.Background(), cUUID, uniUUID, tableType, keyspaceName).PitrConfig(pitrConfig).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PITRManagementApi.CreatePitrConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePitrConfig`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `PITRManagementApi.CreatePitrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 
**tableType** | **string** |  | 
**keyspaceName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePitrConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pitrConfig** | [**CreatePitrConfigParams**](CreatePitrConfigParams.md) | post pitr config | 
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


## DeletePitrConfig

> YBPTask DeletePitrConfig(ctx, cUUID, uniUUID, pUUID).Request(request).Execute()

Delete pitr config on a universe

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
    pUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PITRManagementApi.DeletePitrConfig(context.Background(), cUUID, uniUUID, pUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PITRManagementApi.DeletePitrConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePitrConfig`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `PITRManagementApi.DeletePitrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePitrConfigRequest struct via the builder pattern


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


## ListOfPitrConfigs

> []PitrConfig ListOfPitrConfigs(ctx, cUUID, uniUUID).Execute()

List the PITR configs of a universe

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PITRManagementApi.ListOfPitrConfigs(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PITRManagementApi.ListOfPitrConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOfPitrConfigs`: []PitrConfig
    fmt.Fprintf(os.Stdout, "Response from `PITRManagementApi.ListOfPitrConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOfPitrConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]PitrConfig**](PitrConfig.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PerformPitr

> YBPTask PerformPitr(ctx, cUUID, uniUUID).PerformPitr(performPitr).Request(request).Execute()

Perform PITR on a universe

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
    performPitr := *openapiclient.NewRestoreSnapshotScheduleParams(*openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), "PlatformUrl_example", int32(123), int32(123)) // RestoreSnapshotScheduleParams | perform PITR
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PITRManagementApi.PerformPitr(context.Background(), cUUID, uniUUID).PerformPitr(performPitr).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PITRManagementApi.PerformPitr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PerformPitr`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `PITRManagementApi.PerformPitr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPerformPitrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **performPitr** | [**RestoreSnapshotScheduleParams**](RestoreSnapshotScheduleParams.md) | perform PITR | 
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


## UpdatePitrConfig

> YBPTask UpdatePitrConfig(ctx, cUUID, uniUUID, pUUID).PitrConfig(pitrConfig).Request(request).Execute()

Update pitr config for a keyspace in a universe



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
    pUUID := TODO // string | 
    pitrConfig := *openapiclient.NewUpdatePitrConfigParams(*openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), "PlatformUrl_example", int32(123), int32(123)) // UpdatePitrConfigParams | put pitr config
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PITRManagementApi.UpdatePitrConfig(context.Background(), cUUID, uniUUID, pUUID).PitrConfig(pitrConfig).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PITRManagementApi.UpdatePitrConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePitrConfig`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `PITRManagementApi.UpdatePitrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePitrConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pitrConfig** | [**UpdatePitrConfigParams**](UpdatePitrConfigParams.md) | put pitr config | 
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

