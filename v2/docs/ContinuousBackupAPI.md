# \ContinuousBackupAPI

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContinuousBackup**](ContinuousBackupAPI.md#CreateContinuousBackup) | **Post** /customers/{cUUID}/auto-yba-backups | Create Continous Backup
[**DeleteContinuousBackup**](ContinuousBackupAPI.md#DeleteContinuousBackup) | **Delete** /customers/{cUUID}/auto-yba-backups/{bUUID} | Delete Continuous Backup config
[**EditContinuousBackup**](ContinuousBackupAPI.md#EditContinuousBackup) | **Put** /customers/{cUUID}/auto-yba-backups/{bUUID} | Edit Continuous Backup config
[**GetContinuousBackup**](ContinuousBackupAPI.md#GetContinuousBackup) | **Get** /customers/{cUUID}/auto-yba-backups | Get Continuous Backup
[**RestoreContinuousBackup**](ContinuousBackupAPI.md#RestoreContinuousBackup) | **Post** /customers/{cUUID}/auto-yba-backups/restore | Restore Continuous Backup



## CreateContinuousBackup

> ContinuousBackup CreateContinuousBackup(ctx, cUUID).ContinuousBackupSpec(continuousBackupSpec).Execute()

Create Continous Backup



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
	continuousBackupSpec := *openapiclient.NewContinuousBackupSpec("f33e3c9b-75ab-4c30-80ad-cba85646ea39", int64(123), openapiclient.TimeUnitType("NANOSECONDS"), "yba_backups") // ContinuousBackupSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContinuousBackupAPI.CreateContinuousBackup(context.Background(), cUUID).ContinuousBackupSpec(continuousBackupSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContinuousBackupAPI.CreateContinuousBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContinuousBackup`: ContinuousBackup
	fmt.Fprintf(os.Stdout, "Response from `ContinuousBackupAPI.CreateContinuousBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContinuousBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **continuousBackupSpec** | [**ContinuousBackupSpec**](ContinuousBackupSpec.md) |  | 

### Return type

[**ContinuousBackup**](ContinuousBackup.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContinuousBackup

> ContinuousBackup DeleteContinuousBackup(ctx, cUUID, bUUID).Execute()

Delete Continuous Backup config



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
	bUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Continuous Backup UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContinuousBackupAPI.DeleteContinuousBackup(context.Background(), cUUID, bUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContinuousBackupAPI.DeleteContinuousBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteContinuousBackup`: ContinuousBackup
	fmt.Fprintf(os.Stdout, "Response from `ContinuousBackupAPI.DeleteContinuousBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**bUUID** | **string** | Continuous Backup UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContinuousBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ContinuousBackup**](ContinuousBackup.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditContinuousBackup

> ContinuousBackup EditContinuousBackup(ctx, cUUID, bUUID).ContinuousBackupSpec(continuousBackupSpec).Execute()

Edit Continuous Backup config



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
	bUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Continuous Backup UUID
	continuousBackupSpec := *openapiclient.NewContinuousBackupSpec("f33e3c9b-75ab-4c30-80ad-cba85646ea39", int64(123), openapiclient.TimeUnitType("NANOSECONDS"), "yba_backups") // ContinuousBackupSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContinuousBackupAPI.EditContinuousBackup(context.Background(), cUUID, bUUID).ContinuousBackupSpec(continuousBackupSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContinuousBackupAPI.EditContinuousBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditContinuousBackup`: ContinuousBackup
	fmt.Fprintf(os.Stdout, "Response from `ContinuousBackupAPI.EditContinuousBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**bUUID** | **string** | Continuous Backup UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditContinuousBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **continuousBackupSpec** | [**ContinuousBackupSpec**](ContinuousBackupSpec.md) |  | 

### Return type

[**ContinuousBackup**](ContinuousBackup.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContinuousBackup

> ContinuousBackup GetContinuousBackup(ctx, cUUID).Execute()

Get Continuous Backup



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContinuousBackupAPI.GetContinuousBackup(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContinuousBackupAPI.GetContinuousBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContinuousBackup`: ContinuousBackup
	fmt.Fprintf(os.Stdout, "Response from `ContinuousBackupAPI.GetContinuousBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContinuousBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContinuousBackup**](ContinuousBackup.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreContinuousBackup

> YBATask RestoreContinuousBackup(ctx, cUUID).ContinuousRestoreSpec(continuousRestoreSpec).Execute()

Restore Continuous Backup



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
	continuousRestoreSpec := *openapiclient.NewContinuousRestoreSpec("f33e3c9b-75ab-4c30-80ad-cba85646ea39", "YBA.1.2.3.4") // ContinuousRestoreSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContinuousBackupAPI.RestoreContinuousBackup(context.Background(), cUUID).ContinuousRestoreSpec(continuousRestoreSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContinuousBackupAPI.RestoreContinuousBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreContinuousBackup`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `ContinuousBackupAPI.RestoreContinuousBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreContinuousBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **continuousRestoreSpec** | [**ContinuousRestoreSpec**](ContinuousRestoreSpec.md) |  | 

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

