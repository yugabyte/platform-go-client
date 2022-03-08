# \BackupsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMultiTableBackup**](BackupsApi.md#CreateMultiTableBackup) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/multi_table_backup | Create a multi-table backup
[**DeleteBackups**](BackupsApi.md#DeleteBackups) | **Delete** /api/v1/customers/{cUUID}/backups | Delete backups
[**DeleteBackupsv2**](BackupsApi.md#DeleteBackupsv2) | **Delete** /api/v1/customers/{cUUID}/delete_backups | Delete backups V2
[**EditBackupV2**](BackupsApi.md#EditBackupV2) | **Put** /api/v1/customers/{cUUID}/backups/{backupUUID} | Edit a backup V2
[**FetchBackupsByTaskUUID**](BackupsApi.md#FetchBackupsByTaskUUID) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/backups/tasks/{tUUID} | List a task&#39;s backups
[**GetBackupV2**](BackupsApi.md#GetBackupV2) | **Get** /api/v1/customers/{cUUID}/backups/{backupUUID} | Get Backup V2
[**ListBackupsV2**](BackupsApi.md#ListBackupsV2) | **Post** /api/v1/customers/{cUUID}/backups/page | List Backups (paginated) V2
[**ListOfBackups**](BackupsApi.md#ListOfBackups) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/backups | List a customer&#39;s backups
[**Restore**](BackupsApi.md#Restore) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/backups/restore | Restore from a backup
[**RestoreBackupV2**](BackupsApi.md#RestoreBackupV2) | **Post** /api/v1/customers/{cUUID}/restore | Restore from a backup V2
[**SetUniverseBackupFlag**](BackupsApi.md#SetUniverseBackupFlag) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/update_backup_state | Set a universe&#39;s backup flag
[**StopBackup**](BackupsApi.md#StopBackup) | **Post** /api/v1/customers/{cUUID}/backups/{backupUUID}/stop | Stop a backup



## CreateMultiTableBackup

> Schedule CreateMultiTableBackup(ctx, cUUID, uniUUID).TableBackup(tableBackup).Execute()

Create a multi-table backup

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
    tableBackup := *openapiclient.NewMultiTableBackupRequestParams("StorageConfigUUID_example") // MultiTableBackupRequestParams | Table backup data to be created

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.CreateMultiTableBackup(context.Background(), cUUID, uniUUID).TableBackup(tableBackup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.CreateMultiTableBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMultiTableBackup`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.CreateMultiTableBackup`: %v\n", resp)
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


## DeleteBackups

> map[string]interface{} DeleteBackups(ctx, cUUID).Execute()

Delete backups

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.DeleteBackups(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.DeleteBackups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBackups`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.DeleteBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBackupsv2

> map[string]interface{} DeleteBackupsv2(ctx, cUUID).Execute()

Delete backups V2

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.DeleteBackupsv2(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.DeleteBackupsv2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBackupsv2`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.DeleteBackupsv2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackupsv2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditBackupV2

> Backup EditBackupV2(ctx, cUUID, backupUUID).Backup(backup).Execute()

Edit a backup V2



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
    backupUUID := TODO // string | 
    backup := *openapiclient.NewEditBackupParams(int64(123)) // EditBackupParams | Parameters of the backup to be edited

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.EditBackupV2(context.Background(), cUUID, backupUUID).Backup(backup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.EditBackupV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditBackupV2`: Backup
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.EditBackupV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**backupUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditBackupV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **backup** | [**EditBackupParams**](EditBackupParams.md) | Parameters of the backup to be edited | 

### Return type

[**Backup**](Backup.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBackupsByTaskUUID

> []Backup FetchBackupsByTaskUUID(ctx, cUUID, uniUUID, tUUID).Execute()

List a task's backups

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
    tUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.FetchBackupsByTaskUUID(context.Background(), cUUID, uniUUID, tUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.FetchBackupsByTaskUUID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchBackupsByTaskUUID`: []Backup
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.FetchBackupsByTaskUUID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 
**tUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchBackupsByTaskUUIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]Backup**](Backup.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupV2

> Backup GetBackupV2(ctx, cUUID, backupUUID).Execute()

Get Backup V2

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
    backupUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.GetBackupV2(context.Background(), cUUID, backupUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.GetBackupV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBackupV2`: Backup
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.GetBackupV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**backupUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Backup**](Backup.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackupsV2

> BackupPagedApiResponse ListBackupsV2(ctx, cUUID).PageBackupsRequest(pageBackupsRequest).Execute()

List Backups (paginated) V2

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    pageBackupsRequest := *openapiclient.NewBackupPagedApiQuery("Direction_example", *openapiclient.NewBackupApiFilter(time.Now(), time.Now(), []string{"KeyspaceList_example"}, []string{"ScheduleUUIDList_example"}, []string{"States_example"}, []string{"StorageConfigUUIDList_example"}, []string{"UniverseNameList_example"}, []string{"UniverseUUIDList_example"}), int32(123), false, int32(123), "SortBy_example") // BackupPagedApiQuery | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.ListBackupsV2(context.Background(), cUUID).PageBackupsRequest(pageBackupsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.ListBackupsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBackupsV2`: BackupPagedApiResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.ListBackupsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageBackupsRequest** | [**BackupPagedApiQuery**](BackupPagedApiQuery.md) |  | 

### Return type

[**BackupPagedApiResponse**](BackupPagedApiResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOfBackups

> []Backup ListOfBackups(ctx, cUUID, uniUUID).Execute()

List a customer's backups

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
    resp, r, err := api_client.BackupsApi.ListOfBackups(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.ListOfBackups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOfBackups`: []Backup
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.ListOfBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOfBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Backup**](Backup.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Restore

> YBPTask Restore(ctx, cUUID, uniUUID).Backup(backup).Execute()

Restore from a backup

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
    backup := *openapiclient.NewBackupTableParams("StorageConfigUUID_example") // BackupTableParams | Parameters of the backup to be restored

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.Restore(context.Background(), cUUID, uniUUID).Backup(backup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.Restore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Restore`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.Restore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **backup** | [**BackupTableParams**](BackupTableParams.md) | Parameters of the backup to be restored | 

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


## RestoreBackupV2

> YBPTask RestoreBackupV2(ctx, cUUID).Backup(backup).Execute()

Restore from a backup V2

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
    backup := *openapiclient.NewRestoreBackupParams("UniverseUUID_example") // RestoreBackupParams | Parameters of the backup to be restored

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.RestoreBackupV2(context.Background(), cUUID).Backup(backup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.RestoreBackupV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestoreBackupV2`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.RestoreBackupV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreBackupV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backup** | [**RestoreBackupParams**](RestoreBackupParams.md) | Parameters of the backup to be restored | 

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


## SetUniverseBackupFlag

> YBPSuccess SetUniverseBackupFlag(ctx, cUUID, uniUUID).MarkActive(markActive).Execute()

Set a universe's backup flag

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
    markActive := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.SetUniverseBackupFlag(context.Background(), cUUID, uniUUID).MarkActive(markActive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.SetUniverseBackupFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetUniverseBackupFlag`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.SetUniverseBackupFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetUniverseBackupFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **markActive** | **bool** |  | 

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


## StopBackup

> YBPSuccess StopBackup(ctx, cUUID, backupUUID).Execute()

Stop a backup



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
    backupUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.StopBackup(context.Background(), cUUID, backupUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.StopBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopBackup`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.StopBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**backupUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopBackupRequest struct via the builder pattern


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

