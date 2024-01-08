# \TableManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMultiTableBackup**](TableManagementApi.md#CreateMultiTableBackup) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/multi_table_backup | Deprecated since YBA version 2.20.0.0 (Use BackupsController). Create a multi-table backup
[**CreateSingleTableBackup**](TableManagementApi.md#CreateSingleTableBackup) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/tables/{tableUUID}/create_backup | Deprecated since YBA version 2.20.0.0 (Use BackupsController). Create a single-table backup



## CreateMultiTableBackup

> Schedule CreateMultiTableBackup(ctx, cUUID, uniUUID).TableBackup(tableBackup).Request(request).Execute()

Deprecated since YBA version 2.20.0.0 (Use BackupsController). Create a multi-table backup

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
    tableBackup := *openapiclient.NewMultiTableBackupRequestParams(*openapiclient.NewUsers("username1@example.com"), "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123), "StorageConfigUUID_example", int64(123)) // MultiTableBackupRequestParams | Table backup data to be created
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TableManagementApi.CreateMultiTableBackup(context.Background(), cUUID, uniUUID).TableBackup(tableBackup).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableManagementApi.CreateMultiTableBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMultiTableBackup`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `TableManagementApi.CreateMultiTableBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMultiTableBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tableBackup** | [**MultiTableBackupRequestParams**](MultiTableBackupRequestParams.md) | Table backup data to be created | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSingleTableBackup

> YBPTask CreateSingleTableBackup(ctx, cUUID, uniUUID, tableUUID).Backup(backup).Request(request).Execute()

Deprecated since YBA version 2.20.0.0 (Use BackupsController). Create a single-table backup

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
    tableUUID := TODO // string | 
    backup := *openapiclient.NewBackupTableParams(*openapiclient.NewUsers("username1@example.com"), "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123), "StorageConfigUUID_example", int64(123)) // BackupTableParams | Backup data to be created
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TableManagementApi.CreateSingleTableBackup(context.Background(), cUUID, uniUUID, tableUUID).Backup(backup).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableManagementApi.CreateSingleTableBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSingleTableBackup`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `TableManagementApi.CreateSingleTableBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 
**tableUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSingleTableBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **backup** | [**BackupTableParams**](BackupTableParams.md) | Backup data to be created | 
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

