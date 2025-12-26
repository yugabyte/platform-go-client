# \IsolatedBackupAPI

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateYbaBackup**](IsolatedBackupAPI.md#CreateYbaBackup) | **Post** /customers/{cUUID}/yba-backups | Create YBA Backup
[**RestoreYbaBackup**](IsolatedBackupAPI.md#RestoreYbaBackup) | **Post** /customers/{cUUID}/yba-backups/restore | Restore YBA Backup



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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	isolatedBackupCreateSpec := *openapiclient.NewIsolatedBackupCreateSpec("/opt/yba_backups", []openapiclient.YbaComponent{openapiclient.YbaComponent("YBA")}) // IsolatedBackupCreateSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IsolatedBackupAPI.CreateYbaBackup(context.Background(), cUUID).IsolatedBackupCreateSpec(isolatedBackupCreateSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IsolatedBackupAPI.CreateYbaBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateYbaBackup`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `IsolatedBackupAPI.CreateYbaBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	isolatedBackupRestoreSpec := *openapiclient.NewIsolatedBackupRestoreSpec("/opt/yba_backups/yba_backup.tar.gz") // IsolatedBackupRestoreSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IsolatedBackupAPI.RestoreYbaBackup(context.Background(), cUUID).IsolatedBackupRestoreSpec(isolatedBackupRestoreSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IsolatedBackupAPI.RestoreYbaBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreYbaBackup`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `IsolatedBackupAPI.RestoreYbaBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

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

