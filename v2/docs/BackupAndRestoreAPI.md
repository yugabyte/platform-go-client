# \BackupAndRestoreAPI

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListYbcGflagsMetadata**](BackupAndRestoreAPI.md#ListYbcGflagsMetadata) | **Get** /ybc/gflags-metadata | List YBC Gflags metadata
[**PageListBackups**](BackupAndRestoreAPI.md#PageListBackups) | **Post** /customers/{cUUID}/backups/page | List universe backups.
[**PageListIncrementalBackups**](BackupAndRestoreAPI.md#PageListIncrementalBackups) | **Post** /customers/{cUUID}/backups/{bUUID}/increments/page | List incremental backups for a backup chain.
[**PageListRestoreKeyspaces**](BackupAndRestoreAPI.md#PageListRestoreKeyspaces) | **Post** /customers/{cUUID}/restores/{rUUID}/keyspaces/page | List keyspaces for a universe restore.
[**PageListRestores**](BackupAndRestoreAPI.md#PageListRestores) | **Post** /customers/{cUUID}/restores/page | List universe restores.



## ListYbcGflagsMetadata

> []GflagMetadata ListYbcGflagsMetadata(ctx).Execute()

List YBC Gflags metadata



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAndRestoreAPI.ListYbcGflagsMetadata(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAndRestoreAPI.ListYbcGflagsMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListYbcGflagsMetadata`: []GflagMetadata
	fmt.Fprintf(os.Stdout, "Response from `BackupAndRestoreAPI.ListYbcGflagsMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListYbcGflagsMetadataRequest struct via the builder pattern


### Return type

[**[]GflagMetadata**](GflagMetadata.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageListBackups

> BackupPagedResp PageListBackups(ctx, cUUID).BackupPagedQuerySpec(backupPagedQuerySpec).Execute()

List universe backups.



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
	backupPagedQuerySpec := *openapiclient.NewBackupPagedQuerySpec() // BackupPagedQuerySpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAndRestoreAPI.PageListBackups(context.Background(), cUUID).BackupPagedQuerySpec(backupPagedQuerySpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAndRestoreAPI.PageListBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageListBackups`: BackupPagedResp
	fmt.Fprintf(os.Stdout, "Response from `BackupAndRestoreAPI.PageListBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageListBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backupPagedQuerySpec** | [**BackupPagedQuerySpec**](BackupPagedQuerySpec.md) |  | 

### Return type

[**BackupPagedResp**](BackupPagedResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageListIncrementalBackups

> IncrementalBackupPagedResp PageListIncrementalBackups(ctx, cUUID, bUUID).IncrementalBackupPagedQuerySpec(incrementalBackupPagedQuerySpec).Execute()

List incremental backups for a backup chain.



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
	bUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Backup UUID of any backup in the incremental chain
	incrementalBackupPagedQuerySpec := *openapiclient.NewIncrementalBackupPagedQuerySpec() // IncrementalBackupPagedQuerySpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAndRestoreAPI.PageListIncrementalBackups(context.Background(), cUUID, bUUID).IncrementalBackupPagedQuerySpec(incrementalBackupPagedQuerySpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAndRestoreAPI.PageListIncrementalBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageListIncrementalBackups`: IncrementalBackupPagedResp
	fmt.Fprintf(os.Stdout, "Response from `BackupAndRestoreAPI.PageListIncrementalBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**bUUID** | **string** | Backup UUID of any backup in the incremental chain | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageListIncrementalBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **incrementalBackupPagedQuerySpec** | [**IncrementalBackupPagedQuerySpec**](IncrementalBackupPagedQuerySpec.md) |  | 

### Return type

[**IncrementalBackupPagedResp**](IncrementalBackupPagedResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageListRestoreKeyspaces

> RestoreKeyspacePagedResp PageListRestoreKeyspaces(ctx, cUUID, rUUID).RestoreKeyspacePagedQuerySpec(restoreKeyspacePagedQuerySpec).Execute()

List keyspaces for a universe restore.



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
	rUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Restore UUID
	restoreKeyspacePagedQuerySpec := *openapiclient.NewRestoreKeyspacePagedQuerySpec() // RestoreKeyspacePagedQuerySpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAndRestoreAPI.PageListRestoreKeyspaces(context.Background(), cUUID, rUUID).RestoreKeyspacePagedQuerySpec(restoreKeyspacePagedQuerySpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAndRestoreAPI.PageListRestoreKeyspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageListRestoreKeyspaces`: RestoreKeyspacePagedResp
	fmt.Fprintf(os.Stdout, "Response from `BackupAndRestoreAPI.PageListRestoreKeyspaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**rUUID** | **string** | Restore UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageListRestoreKeyspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **restoreKeyspacePagedQuerySpec** | [**RestoreKeyspacePagedQuerySpec**](RestoreKeyspacePagedQuerySpec.md) |  | 

### Return type

[**RestoreKeyspacePagedResp**](RestoreKeyspacePagedResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageListRestores

> RestorePagedResp PageListRestores(ctx, cUUID).RestorePagedQuerySpec(restorePagedQuerySpec).Execute()

List universe restores.



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
	restorePagedQuerySpec := *openapiclient.NewRestorePagedQuerySpec() // RestorePagedQuerySpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAndRestoreAPI.PageListRestores(context.Background(), cUUID).RestorePagedQuerySpec(restorePagedQuerySpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAndRestoreAPI.PageListRestores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageListRestores`: RestorePagedResp
	fmt.Fprintf(os.Stdout, "Response from `BackupAndRestoreAPI.PageListRestores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageListRestoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restorePagedQuerySpec** | [**RestorePagedQuerySpec**](RestorePagedQuerySpec.md) |  | 

### Return type

[**RestorePagedResp**](RestorePagedResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

