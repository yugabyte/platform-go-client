# \StorageConfigAPI

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStorageConfig**](StorageConfigAPI.md#CreateStorageConfig) | **Post** /customers/{cUUID}/storage-configs | Create a storage configuration
[**GetStorageConfig**](StorageConfigAPI.md#GetStorageConfig) | **Get** /customers/{cUUID}/storage-configs/{scUUID} | Get a storage configuration
[**ValidateStorageConfig**](StorageConfigAPI.md#ValidateStorageConfig) | **Post** /customers/{cUUID}/storage-configs/validate | Validate a storage configuration



## CreateStorageConfig

> StorageConfig CreateStorageConfig(ctx, cUUID).StorageConfigSpec(storageConfigSpec).IdempotencyKey(idempotencyKey).Execute()

Create a storage configuration



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
	storageConfigSpec := *openapiclient.NewStorageConfigSpec("Name_example", openapiclient.StorageSpec{AzureStorageSpec: openapiclient.NewAzureStorageSpec(openapiclient.StorageType("S3"), "https://example.blob.core.windows.net/backups", openapiclient.AzureCredentials{AzureManagedIdentityCredentials: openapiclient.NewAzureManagedIdentityCredentials(openapiclient.AzureCredentialType("SAS_TOKEN"))})}) // StorageConfigSpec | 
	idempotencyKey := "3fa85f64-5717-4562-b3fc-2c963f66afa6" // string | Makes the request safe to retry. A unique client-generated value, such as a UUID, for this request.  The first request with a key runs the operation. If it succeeds, its response is kept for 24 hours (yb.api.idempotency.retention), and a retry with the same key and the same request returns that response with an X-IDEMPOTENT-REPLAYED: true header, without running the operation again. A retry while the first request is still running gets 409; the same key with a different request gets 422. A failed request does not keep the key, so it can be retried with the same one.  Keys are scoped to the caller's customer.  If the operation starts a task, that task's UUID is derived from the key: the version 5 UUID of \"<customer UUID>:<key>\" in the namespace 3bf373dd-d374-4df5-bdea-c3cd942ae7f4. A client can follow the task even if it lost the response. If that task was started, for example before YBA restarted and lost the response, a retry gets 409 instead of starting the operation again, with the task's UUID in error.task_uuid, and in error.latest_task_uuid the UUID of its newest retry if YBA retried it (a restart fails or aborts running tasks, and may retry them under a new UUID). If the task failed while being submitted, the retry runs the operation again under a new task UUID.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageConfigAPI.CreateStorageConfig(context.Background(), cUUID).StorageConfigSpec(storageConfigSpec).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageConfigAPI.CreateStorageConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStorageConfig`: StorageConfig
	fmt.Fprintf(os.Stdout, "Response from `StorageConfigAPI.CreateStorageConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorageConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storageConfigSpec** | [**StorageConfigSpec**](StorageConfigSpec.md) |  | 
 **idempotencyKey** | **string** | Makes the request safe to retry. A unique client-generated value, such as a UUID, for this request.  The first request with a key runs the operation. If it succeeds, its response is kept for 24 hours (yb.api.idempotency.retention), and a retry with the same key and the same request returns that response with an X-IDEMPOTENT-REPLAYED: true header, without running the operation again. A retry while the first request is still running gets 409; the same key with a different request gets 422. A failed request does not keep the key, so it can be retried with the same one.  Keys are scoped to the caller&#39;s customer.  If the operation starts a task, that task&#39;s UUID is derived from the key: the version 5 UUID of \&quot;&lt;customer UUID&gt;:&lt;key&gt;\&quot; in the namespace 3bf373dd-d374-4df5-bdea-c3cd942ae7f4. A client can follow the task even if it lost the response. If that task was started, for example before YBA restarted and lost the response, a retry gets 409 instead of starting the operation again, with the task&#39;s UUID in error.task_uuid, and in error.latest_task_uuid the UUID of its newest retry if YBA retried it (a restart fails or aborts running tasks, and may retry them under a new UUID). If the task failed while being submitted, the retry runs the operation again under a new task UUID.  | 

### Return type

[**StorageConfig**](StorageConfig.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [authTokenAuth](../README.md#authTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageConfig

> StorageConfig GetStorageConfig(ctx, cUUID, scUUID).Execute()

Get a storage configuration



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
	scUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Storage configuration UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageConfigAPI.GetStorageConfig(context.Background(), cUUID, scUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageConfigAPI.GetStorageConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageConfig`: StorageConfig
	fmt.Fprintf(os.Stdout, "Response from `StorageConfigAPI.GetStorageConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**scUUID** | **string** | Storage configuration UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StorageConfig**](StorageConfig.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [authTokenAuth](../README.md#authTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateStorageConfig

> ValidateStorageConfig(ctx, cUUID).StorageConfigSpec(storageConfigSpec).Execute()

Validate a storage configuration



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
	storageConfigSpec := *openapiclient.NewStorageConfigSpec("Name_example", openapiclient.StorageSpec{AzureStorageSpec: openapiclient.NewAzureStorageSpec(openapiclient.StorageType("S3"), "https://example.blob.core.windows.net/backups", openapiclient.AzureCredentials{AzureManagedIdentityCredentials: openapiclient.NewAzureManagedIdentityCredentials(openapiclient.AzureCredentialType("SAS_TOKEN"))})}) // StorageConfigSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageConfigAPI.ValidateStorageConfig(context.Background(), cUUID).StorageConfigSpec(storageConfigSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageConfigAPI.ValidateStorageConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateStorageConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storageConfigSpec** | [**StorageConfigSpec**](StorageConfigSpec.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [authTokenAuth](../README.md#authTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

