# \BackupsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdvancedRestorePreflight**](BackupsApi.md#AdvancedRestorePreflight) | **Post** /api/v1/customers/{cUUID}/restore/advanced_restore_preflight | Advanced Restore Preflight checks
[**CreateBackupSchedule**](BackupsApi.md#CreateBackupSchedule) | **Post** /api/v1/customers/{cUUID}/create_backup_schedule | Create Backup Schedule - deprecated
[**CreateBackupScheduleAsync**](BackupsApi.md#CreateBackupScheduleAsync) | **Post** /api/v1/customers/{cUUID}/create_backup_schedule_async | Create Backup Schedule Async
[**CreateMultiTableBackup**](BackupsApi.md#CreateMultiTableBackup) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/multi_table_backup | Create a multi-table backup - deprecated
[**CreateSingleTableBackup**](BackupsApi.md#CreateSingleTableBackup) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/tables/{tableUUID}/create_backup | Create a single-table backup - deprecated
[**Createbackup**](BackupsApi.md#Createbackup) | **Post** /api/v1/customers/{cUUID}/backups | Create a backup V2
[**DeleteBackups**](BackupsApi.md#DeleteBackups) | **Delete** /api/v1/customers/{cUUID}/backups | Delete Backups - deprecated
[**DeleteBackupsV2**](BackupsApi.md#DeleteBackupsV2) | **Post** /api/v1/customers/{cUUID}/backups/delete | Delete backups V2
[**EditBackupV2**](BackupsApi.md#EditBackupV2) | **Put** /api/v1/customers/{cUUID}/backups/{backupUUID} | Edit a backup V2
[**FetchBackupsByTaskUUID**](BackupsApi.md#FetchBackupsByTaskUUID) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/backups/tasks/{tUUID} | List backups associated with a task
[**GetBackupV2**](BackupsApi.md#GetBackupV2) | **Get** /api/v1/customers/{cUUID}/backups/{backupUUID} | Get Backup V2
[**GetThrottleParams**](BackupsApi.md#GetThrottleParams) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/ybc_throttle_params | Get throttle params from YB-Controller
[**ListBackupRestoresV2**](BackupsApi.md#ListBackupRestoresV2) | **Post** /api/v1/customers/{cUUID}/restore/page | List Backup Restores (paginated)
[**ListBackupsV2**](BackupsApi.md#ListBackupsV2) | **Post** /api/v1/customers/{cUUID}/backups/page | List Backups (paginated) V2
[**ListIncrementalBackups**](BackupsApi.md#ListIncrementalBackups) | **Get** /api/v1/customers/{cUUID}/backups/{backupUUID}/list_increments | List Incremental backups
[**ListOfBackups**](BackupsApi.md#ListOfBackups) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/backups | List a customer&#39;s backups - deprecated
[**ListRestorableKeyspaceTables**](BackupsApi.md#ListRestorableKeyspaceTables) | **Get** /api/v1/customers/{cUUID}/backups/{baseBackupUUID}/restorable_keyspace_tables | List of all restorable entities in the incremental backup chain
[**Restore**](BackupsApi.md#Restore) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/backups/restore | Restore from a backup - deprecated
[**RestoreBackupV2**](BackupsApi.md#RestoreBackupV2) | **Post** /api/v1/customers/{cUUID}/restore | Restore from a backup V2
[**RestorePreflight**](BackupsApi.md#RestorePreflight) | **Post** /api/v1/customers/{cUUID}/restore/preflight | Restore preflight checks
[**SetThrottleParams**](BackupsApi.md#SetThrottleParams) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/ybc_throttle_params | Set throttle params in YB-Controller
[**SetUniverseBackupFlag**](BackupsApi.md#SetUniverseBackupFlag) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/update_backup_state | Set a universe&#39;s backup flag
[**StopBackup**](BackupsApi.md#StopBackup) | **Post** /api/v1/customers/{cUUID}/backups/{backupUUID}/stop | Stop a backup
[**UniverseBackup**](BackupsApi.md#UniverseBackup) | **Post** /api/v1/customers/{customerUUID}/universes/{universeUUID}/universe_backup | Create a Universe Backup
[**ValidateKeyspaceTablesToRestore**](BackupsApi.md#ValidateKeyspaceTablesToRestore) | **Post** /api/v1/customers/{cUUID}/restore/validate_restorable_keyspace_tables | Validate keyspace and tables to Restore against Backup



## AdvancedRestorePreflight

> RestorePreflightResponse AdvancedRestorePreflight(ctx, cUUID).AdvancedRestorePreflightParams(advancedRestorePreflightParams).Request(request).Execute()

Advanced Restore Preflight checks



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
    advancedRestorePreflightParams := *openapiclient.NewParametersForAdvancedRestorePreflightChecks([]string{"BackupLocations_example"}, "StorageConfigUUID_example", "UniverseUUID_example") // ParametersForAdvancedRestorePreflightChecks | Parameters for advanced restore preflight checks
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.AdvancedRestorePreflight(context.Background(), cUUID).AdvancedRestorePreflightParams(advancedRestorePreflightParams).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.AdvancedRestorePreflight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdvancedRestorePreflight`: RestorePreflightResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.AdvancedRestorePreflight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdvancedRestorePreflightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **advancedRestorePreflightParams** | [**ParametersForAdvancedRestorePreflightChecks**](ParametersForAdvancedRestorePreflightChecks.md) | Parameters for advanced restore preflight checks | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**RestorePreflightResponse**](RestorePreflightResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBackupSchedule

> Schedule CreateBackupSchedule(ctx, cUUID).Backup(backup).Request(request).Execute()

Create Backup Schedule - deprecated



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
    backup := *openapiclient.NewBackupRequestParams("BackupUUID_example", *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), "PlatformUrl_example", int32(123), int32(123), "StorageConfigUUID_example", "UniverseUUID_example") // BackupRequestParams | Parameters of the backup to be restored
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.CreateBackupSchedule(context.Background(), cUUID).Backup(backup).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.CreateBackupSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBackupSchedule`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.CreateBackupSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backup** | [**BackupRequestParams**](BackupRequestParams.md) | Parameters of the backup to be restored | 
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


## CreateBackupScheduleAsync

> YBPTask CreateBackupScheduleAsync(ctx, cUUID).Backup(backup).Request(request).Execute()

Create Backup Schedule Async

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
    backup := *openapiclient.NewBackupRequestParams("BackupUUID_example", *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), "PlatformUrl_example", int32(123), int32(123), "StorageConfigUUID_example", "UniverseUUID_example") // BackupRequestParams | Parameters of the backup to be restored
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.CreateBackupScheduleAsync(context.Background(), cUUID).Backup(backup).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.CreateBackupScheduleAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBackupScheduleAsync`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.CreateBackupScheduleAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackupScheduleAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backup** | [**BackupRequestParams**](BackupRequestParams.md) | Parameters of the backup to be restored | 
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


## CreateMultiTableBackup

> Schedule CreateMultiTableBackup(ctx, cUUID, uniUUID).TableBackup(tableBackup).Request(request).Execute()

Create a multi-table backup - deprecated



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
    tableBackup := *openapiclient.NewMultiTableBackupRequestParams(*openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), false, "PlatformUrl_example", int32(123), int32(123), "StorageConfigUUID_example", int64(123)) // MultiTableBackupRequestParams | Table backup data to be created
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.CreateMultiTableBackup(context.Background(), cUUID, uniUUID).TableBackup(tableBackup).Request(request).Execute()
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

Create a single-table backup - deprecated



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
    backup := *openapiclient.NewBackupTableParams(*openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), false, "PlatformUrl_example", int32(123), int32(123), "StorageConfigUUID_example", int64(123)) // BackupTableParams | Backup data to be created
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.CreateSingleTableBackup(context.Background(), cUUID, uniUUID, tableUUID).Backup(backup).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.CreateSingleTableBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSingleTableBackup`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.CreateSingleTableBackup`: %v\n", resp)
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


## Createbackup

> YBPTask Createbackup(ctx, cUUID).Backup(backup).Request(request).Execute()

Create a backup V2

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
    backup := *openapiclient.NewBackupRequestParams("BackupUUID_example", *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), "PlatformUrl_example", int32(123), int32(123), "StorageConfigUUID_example", "UniverseUUID_example") // BackupRequestParams | Backup data to be created
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.Createbackup(context.Background(), cUUID).Backup(backup).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.Createbackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Createbackup`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.Createbackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatebackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backup** | [**BackupRequestParams**](BackupRequestParams.md) | Backup data to be created | 
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


## DeleteBackups

> map[string]interface{} DeleteBackups(ctx, cUUID).Request(request).Execute()

Delete Backups - deprecated



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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.DeleteBackups(context.Background(), cUUID).Request(request).Execute()
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

 **request** | [**interface{}**](interface{}.md) |  | 

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


## DeleteBackupsV2

> map[string]interface{} DeleteBackupsV2(ctx, cUUID).DeleteBackup(deleteBackup).Request(request).Execute()

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
    deleteBackup := *openapiclient.NewDeleteBackupParams([]openapiclient.DeleteBackupInfo{*openapiclient.NewDeleteBackupInfo("BackupUUID_example")}) // DeleteBackupParams | Parameters of the backup to be deleted
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.DeleteBackupsV2(context.Background(), cUUID).DeleteBackup(deleteBackup).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.DeleteBackupsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBackupsV2`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.DeleteBackupsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackupsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteBackup** | [**DeleteBackupParams**](DeleteBackupParams.md) | Parameters of the backup to be deleted | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditBackupV2

> Backup EditBackupV2(ctx, cUUID, backupUUID).Backup(backup).Request(request).Execute()

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
    backup := *openapiclient.NewEditBackupParams() // EditBackupParams | Parameters of the backup to be edited
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.EditBackupV2(context.Background(), cUUID, backupUUID).Backup(backup).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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

List backups associated with a task

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


## GetThrottleParams

> YbcThrottleParametersResponse GetThrottleParams(ctx, cUUID, uniUUID).Execute()

Get throttle params from YB-Controller

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
    resp, r, err := api_client.BackupsApi.GetThrottleParams(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.GetThrottleParams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetThrottleParams`: YbcThrottleParametersResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.GetThrottleParams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThrottleParamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**YbcThrottleParametersResponse**](YbcThrottleParametersResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackupRestoresV2

> RestorePagedApiResponse ListBackupRestoresV2(ctx, cUUID).PageRestoresRequest(pageRestoresRequest).Request(request).Execute()

List Backup Restores (paginated)

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
    pageRestoresRequest := *openapiclient.NewRestorePagedApiQuery("Direction_example", *openapiclient.NewRestoreApiFilter(false, []string{"RestoreUUIDList_example"}, false, []string{"SourceUniverseNameList_example"}, []string{"States_example"}, []string{"StorageConfigUUIDList_example"}, []string{"UniverseNameList_example"}, []string{"UniverseUUIDList_example"}), int32(123), false, int32(123), "SortBy_example") // RestorePagedApiQuery | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.ListBackupRestoresV2(context.Background(), cUUID).PageRestoresRequest(pageRestoresRequest).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.ListBackupRestoresV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBackupRestoresV2`: RestorePagedApiResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.ListBackupRestoresV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupRestoresV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageRestoresRequest** | [**RestorePagedApiQuery**](RestorePagedApiQuery.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**RestorePagedApiResponse**](RestorePagedApiResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackupsV2

> BackupPagedApiResponse ListBackupsV2(ctx, cUUID).PageBackupsRequest(pageBackupsRequest).Request(request).Execute()

List Backups (paginated) V2

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
    pageBackupsRequest := *openapiclient.NewBackupPagedApiQuery("Direction_example", *openapiclient.NewBackupApiFilter([]string{"BackupUUIDList_example"}, []string{"KeyspaceList_example"}, false, false, []string{"ScheduleUUIDList_example"}, false, []string{"States_example"}, []string{"StorageConfigUUIDList_example"}, []string{"UniverseNameList_example"}, []string{"UniverseUUIDList_example"}), int32(123), false, int32(123), "SortBy_example") // BackupPagedApiQuery | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.ListBackupsV2(context.Background(), cUUID).PageBackupsRequest(pageBackupsRequest).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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


## ListIncrementalBackups

> []CommonBackupInfo ListIncrementalBackups(ctx, cUUID, backupUUID).Execute()

List Incremental backups

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
    resp, r, err := api_client.BackupsApi.ListIncrementalBackups(context.Background(), cUUID, backupUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.ListIncrementalBackups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIncrementalBackups`: []CommonBackupInfo
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.ListIncrementalBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**backupUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIncrementalBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]CommonBackupInfo**](CommonBackupInfo.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOfBackups

> []Backup ListOfBackups(ctx, cUUID, uniUUID).Execute()

List a customer's backups - deprecated



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


## ListRestorableKeyspaceTables

> []KeyspaceTables ListRestorableKeyspaceTables(ctx, cUUID, baseBackupUUID).Execute()

List of all restorable entities in the incremental backup chain



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
    baseBackupUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.ListRestorableKeyspaceTables(context.Background(), cUUID, baseBackupUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.ListRestorableKeyspaceTables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRestorableKeyspaceTables`: []KeyspaceTables
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.ListRestorableKeyspaceTables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**baseBackupUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRestorableKeyspaceTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]KeyspaceTables**](KeyspaceTables.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Restore

> YBPTask Restore(ctx, cUUID, uniUUID).Backup(backup).Request(request).Execute()

Restore from a backup - deprecated



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
    backup := *openapiclient.NewBackupTableParams(*openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), false, "PlatformUrl_example", int32(123), int32(123), "StorageConfigUUID_example", int64(123)) // BackupTableParams | Parameters of the backup to be restored
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.Restore(context.Background(), cUUID, uniUUID).Backup(backup).Request(request).Execute()
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


## RestoreBackupV2

> YBPTask RestoreBackupV2(ctx, cUUID).Backup(backup).Request(request).Execute()

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
    backup := *openapiclient.NewRestoreBackupParams(*openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), "PlatformUrl_example", int32(123), int32(123), "UniverseUUID_example") // RestoreBackupParams | Parameters of the backup to be restored
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.RestoreBackupV2(context.Background(), cUUID).Backup(backup).Request(request).Execute()
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


## RestorePreflight

> RestorePreflightResponse RestorePreflight(ctx, cUUID).RestorePreflightParams(restorePreflightParams).Request(request).Execute()

Restore preflight checks



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
    restorePreflightParams := *openapiclient.NewParametersForRestorePreflightChecks("BackupUUID_example", "UniverseUUID_example") // ParametersForRestorePreflightChecks | Parameters for restore preflight check
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.RestorePreflight(context.Background(), cUUID).RestorePreflightParams(restorePreflightParams).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.RestorePreflight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestorePreflight`: RestorePreflightResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.RestorePreflight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestorePreflightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restorePreflightParams** | [**ParametersForRestorePreflightChecks**](ParametersForRestorePreflightChecks.md) | Parameters for restore preflight check | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**RestorePreflightResponse**](RestorePreflightResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetThrottleParams

> YBPSuccess SetThrottleParams(ctx, cUUID, uniUUID).ThrottleParams(throttleParams).Request(request).Execute()

Set throttle params in YB-Controller

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
    throttleParams := *openapiclient.NewYbcThrottleParameters() // YbcThrottleParameters | Parameters for YB-Controller throttling
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.SetThrottleParams(context.Background(), cUUID, uniUUID).ThrottleParams(throttleParams).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.SetThrottleParams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetThrottleParams`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.SetThrottleParams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetThrottleParamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **throttleParams** | [**YbcThrottleParameters**](YbcThrottleParameters.md) | Parameters for YB-Controller throttling | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPSuccess**](YBPSuccess.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetUniverseBackupFlag

> YBPSuccess SetUniverseBackupFlag(ctx, cUUID, uniUUID).MarkActive(markActive).Request(request).Execute()

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.SetUniverseBackupFlag(context.Background(), cUUID, uniUUID).MarkActive(markActive).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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

> YBPSuccess StopBackup(ctx, cUUID, backupUUID).Request(request).Execute()

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.StopBackup(context.Background(), cUUID, backupUUID).Request(request).Execute()
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


 **request** | [**interface{}**](interface{}.md) |  | 

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


## UniverseBackup

> YBPTask UniverseBackup(ctx, customerUUID, universeUUID).BackupUniverse(backupUniverse).Request(request).Execute()

Create a Universe Backup



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
    customerUUID := TODO // string | 
    universeUUID := TODO // string | 
    backupUniverse := *openapiclient.NewUniverseBackupRequestFormData("StorageConfigUUID_example") // UniverseBackupRequestFormData | Universe Backup data to be created
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.UniverseBackup(context.Background(), customerUUID, universeUUID).BackupUniverse(backupUniverse).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.UniverseBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UniverseBackup`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.UniverseBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerUUID** | [**string**](.md) |  | 
**universeUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUniverseBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **backupUniverse** | [**UniverseBackupRequestFormData**](UniverseBackupRequestFormData.md) | Universe Backup data to be created | 
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


## ValidateKeyspaceTablesToRestore

> YBPSuccess ValidateKeyspaceTablesToRestore(ctx, cUUID).RestoreItemsValidationParams(restoreItemsValidationParams).Request(request).Execute()

Validate keyspace and tables to Restore against Backup



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
    restoreItemsValidationParams := *openapiclient.NewParametersForValidatingRestorableKeyspaceAndTablesInBackup("BackupUUID_example") // ParametersForValidatingRestorableKeyspaceAndTablesInBackup | Parameters for validating Restorable keyspace and tables
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.ValidateKeyspaceTablesToRestore(context.Background(), cUUID).RestoreItemsValidationParams(restoreItemsValidationParams).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.ValidateKeyspaceTablesToRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateKeyspaceTablesToRestore`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.ValidateKeyspaceTablesToRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateKeyspaceTablesToRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restoreItemsValidationParams** | [**ParametersForValidatingRestorableKeyspaceAndTablesInBackup**](ParametersForValidatingRestorableKeyspaceAndTablesInBackup.md) | Parameters for validating Restorable keyspace and tables | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPSuccess**](YBPSuccess.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

