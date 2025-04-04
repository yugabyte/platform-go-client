# \IsolatedBackupApi

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateYbaBackup**](IsolatedBackupApi.md#CreateYbaBackup) | **Post** /customers/{cUUID}/yba-backups | Create YBA Backup
[**RestoreYbaBackup**](IsolatedBackupApi.md#RestoreYbaBackup) | **Post** /customers/{cUUID}/yba-backups/restore | Restore YBA Backup



## CreateYbaBackup

> YBATask CreateYbaBackup(ctx, cUUID).IsolatedBackupCreateSpec(isolatedBackupCreateSpec).Execute()

Create YBA Backup



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
    isolatedBackupCreateSpec := *openapiclient.NewIsolatedBackupCreateSpec("/opt/yba_backups", []openapiclient.YbaComponent{openapiclient.YbaComponent("YBA")}) // IsolatedBackupCreateSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IsolatedBackupApi.CreateYbaBackup(context.Background(), cUUID).IsolatedBackupCreateSpec(isolatedBackupCreateSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IsolatedBackupApi.CreateYbaBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateYbaBackup`: YBATask
    fmt.Fprintf(os.Stdout, "Response from `IsolatedBackupApi.CreateYbaBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateYbaBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isolatedBackupCreateSpec** | [**IsolatedBackupCreateSpec**](IsolatedBackupCreateSpec.md) |  | 

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


## RestoreYbaBackup

> YBATask RestoreYbaBackup(ctx, cUUID).IsolatedBackupRestoreSpec(isolatedBackupRestoreSpec).Execute()

Restore YBA Backup



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
    isolatedBackupRestoreSpec := *openapiclient.NewIsolatedBackupRestoreSpec("/opt/yba_backups/yba_backup.tar.gz") // IsolatedBackupRestoreSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IsolatedBackupApi.RestoreYbaBackup(context.Background(), cUUID).IsolatedBackupRestoreSpec(isolatedBackupRestoreSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IsolatedBackupApi.RestoreYbaBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestoreYbaBackup`: YBATask
    fmt.Fprintf(os.Stdout, "Response from `IsolatedBackupApi.RestoreYbaBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreYbaBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isolatedBackupRestoreSpec** | [**IsolatedBackupRestoreSpec**](IsolatedBackupRestoreSpec.md) |  | 

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

