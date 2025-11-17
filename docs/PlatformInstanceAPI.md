# \PlatformInstanceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInstance**](PlatformInstanceAPI.md#CreateInstance) | **Post** /api/v1/settings/ha/config/{cUUID}/instance | Create platform instance
[**DeleteInstance**](PlatformInstanceAPI.md#DeleteInstance) | **Delete** /api/v1/settings/ha/config/{cUUID}/instance/{iUUID} | 
[**GetLocal**](PlatformInstanceAPI.md#GetLocal) | **Get** /api/v1/settings/ha/config/{cUUID}/instance/local | 
[**PromoteInstance**](PlatformInstanceAPI.md#PromoteInstance) | **Post** /api/v1/settings/ha/config/{cUUID}/instance/{iUUID}/promote | Promote platform instance



## CreateInstance

> PlatformInstance CreateInstance(ctx, cUUID).PlatformInstanceFormRequest(platformInstanceFormRequest).Request(request).Execute()

Create platform instance



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
	platformInstanceFormRequest := *openapiclient.NewPlatformInstanceFormData("Address_example", "CleanAddress_example", false, false) // PlatformInstanceFormData | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformInstanceAPI.CreateInstance(context.Background(), cUUID).PlatformInstanceFormRequest(platformInstanceFormRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformInstanceAPI.CreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInstance`: PlatformInstance
	fmt.Fprintf(os.Stdout, "Response from `PlatformInstanceAPI.CreateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **platformInstanceFormRequest** | [**PlatformInstanceFormData**](PlatformInstanceFormData.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**PlatformInstance**](PlatformInstance.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstance

> DeleteInstance(ctx, cUUID, iUUID).Request(request).Execute()



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
	iUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformInstanceAPI.DeleteInstance(context.Background(), cUUID, iUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformInstanceAPI.DeleteInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**iUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceRequest struct via the builder pattern


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


## GetLocal

> GetLocal(ctx, cUUID).Execute()



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
	r, err := apiClient.PlatformInstanceAPI.GetLocal(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformInstanceAPI.GetLocal``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetLocalRequest struct via the builder pattern


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


## PromoteInstance

> PromoteInstance(ctx, cUUID, iUUID).PlatformBackupRestoreRequest(platformBackupRestoreRequest).CurLeader(curLeader).IsForcePromote(isForcePromote).Request(request).Execute()

Promote platform instance



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
	iUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	platformBackupRestoreRequest := *openapiclient.NewRestorePlatformBackupFormData("BackupFile_example") // RestorePlatformBackupFormData | 
	curLeader := "curLeader_example" // string |  (optional) (default to "null")
	isForcePromote := true // bool |  (optional) (default to true)
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformInstanceAPI.PromoteInstance(context.Background(), cUUID, iUUID).PlatformBackupRestoreRequest(platformBackupRestoreRequest).CurLeader(curLeader).IsForcePromote(isForcePromote).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformInstanceAPI.PromoteInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**iUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromoteInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **platformBackupRestoreRequest** | [**RestorePlatformBackupFormData**](RestorePlatformBackupFormData.md) |  | 
 **curLeader** | **string** |  | [default to &quot;null&quot;]
 **isForcePromote** | **bool** |  | [default to true]
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

