# \PITRManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePitrConfig**](PITRManagementApi.md#CreatePitrConfig) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/keyspaces/{tableType}/{keyspaceName}/pitr_config | Create pitr config for a keyspace in a universe
[**DeletePitrConfig**](PITRManagementApi.md#DeletePitrConfig) | **Delete** /api/v1/customers/{cUUID}/universes/{uniUUID}/pitr_config/{pUUID} | Delete pitr config on a universe
[**ListOfPitrConfigs**](PITRManagementApi.md#ListOfPitrConfigs) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/pitr_config | List the PITR configs of a universe
[**PerformPitr**](PITRManagementApi.md#PerformPitr) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/pitr | Perform PITR on a universe



## CreatePitrConfig

> YBPTask CreatePitrConfig(ctx, cUUID, uniUUID, tableType, keyspaceName).PitrConfig(pitrConfig).Execute()

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
    pitrConfig := *openapiclient.NewCreatePitrConfigParams(*openapiclient.NewUsers("username1@example.com"), "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123)) // CreatePitrConfigParams | post pitr config

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PITRManagementApi.CreatePitrConfig(context.Background(), cUUID, uniUUID, tableType, keyspaceName).PitrConfig(pitrConfig).Execute()
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

> YBPSuccess DeletePitrConfig(ctx, cUUID, uniUUID, pUUID).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PITRManagementApi.DeletePitrConfig(context.Background(), cUUID, uniUUID, pUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PITRManagementApi.DeletePitrConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePitrConfig`: YBPSuccess
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




### Return type

[**YBPSuccess**](YBPSuccess.md)

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

> YBPTask PerformPitr(ctx, cUUID, uniUUID).PerformPitr(performPitr).Execute()

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
    performPitr := *openapiclient.NewRestoreSnapshotScheduleParams(*openapiclient.NewUsers("username1@example.com"), "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123)) // RestoreSnapshotScheduleParams | perform PITR

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PITRManagementApi.PerformPitr(context.Background(), cUUID, uniUUID).PerformPitr(performPitr).Execute()
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

