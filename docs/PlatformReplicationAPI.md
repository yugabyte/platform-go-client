# \PlatformReplicationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBackupInfo**](PlatformReplicationAPI.md#GetBackupInfo) | **Get** /api/v1/settings/ha/config/{cUUID}/replication_schedule | 
[**ListBackups**](PlatformReplicationAPI.md#ListBackups) | **Get** /api/v1/settings/ha/config/{cUUID}/backup/list | List backups
[**StartPeriodicBackup**](PlatformReplicationAPI.md#StartPeriodicBackup) | **Put** /api/v1/settings/ha/config/{cUUID}/replication_schedule/start | Start periodic backup
[**StopPeriodicBackup**](PlatformReplicationAPI.md#StopPeriodicBackup) | **Put** /api/v1/settings/ha/config/{cUUID}/replication_schedule/stop | 



## GetBackupInfo

> GetBackupInfo(ctx, cUUID).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformReplicationAPI.GetBackupInfo(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformReplicationAPI.GetBackupInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackups

> []string ListBackups(ctx, cUUID).Leader(leader).Execute()

List backups



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
	leader := "leader_example" // string |  (optional) (default to "null")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformReplicationAPI.ListBackups(context.Background(), cUUID).Leader(leader).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformReplicationAPI.ListBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackups`: []string
	fmt.Fprintf(os.Stdout, "Response from `PlatformReplicationAPI.ListBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **leader** | **string** |  | [default to &quot;null&quot;]

### Return type

**[]string**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartPeriodicBackup

> StartPeriodicBackup(ctx, cUUID).PlatformBackupFrequencyRequest(platformBackupFrequencyRequest).Request(request).Execute()

Start periodic backup



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
	platformBackupFrequencyRequest := *openapiclient.NewPlatformBackupFrequencyFormData(int64(123)) // PlatformBackupFrequencyFormData | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformReplicationAPI.StartPeriodicBackup(context.Background(), cUUID).PlatformBackupFrequencyRequest(platformBackupFrequencyRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformReplicationAPI.StartPeriodicBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartPeriodicBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **platformBackupFrequencyRequest** | [**PlatformBackupFrequencyFormData**](PlatformBackupFrequencyFormData.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopPeriodicBackup

> StopPeriodicBackup(ctx, cUUID).Request(request).Execute()



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
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformReplicationAPI.StopPeriodicBackup(context.Background(), cUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformReplicationAPI.StopPeriodicBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopPeriodicBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

