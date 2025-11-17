# \DisasterRecoveryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDrConfig**](DisasterRecoveryAPI.md#CreateDrConfig) | **Post** /api/v1/customers/{cUUID}/dr_configs | Create disaster recovery config
[**DeleteXClusterConfig**](DisasterRecoveryAPI.md#DeleteXClusterConfig) | **Delete** /api/v1/customers/{cUUID}/dr_configs/{drUUID} | Delete xcluster config
[**EditDrConfig**](DisasterRecoveryAPI.md#EditDrConfig) | **Post** /api/v1/customers/{cUUID}/dr_configs/{drUUID}/edit | Edit disaster recovery config
[**FailoverDrConfig**](DisasterRecoveryAPI.md#FailoverDrConfig) | **Post** /api/v1/customers/{cUUID}/dr_configs/{drUUID}/failover | Failover a disaster recovery config
[**GetDrConfig**](DisasterRecoveryAPI.md#GetDrConfig) | **Get** /api/v1/customers/{cUUID}/dr_configs/{drUUID} | Get disaster recovery config
[**GetDrConfigSafetime**](DisasterRecoveryAPI.md#GetDrConfigSafetime) | **Get** /api/v1/customers/{cUUID}/dr_configs/{drUUID}/safetime | Get disaster recovery config safetime
[**PauseDrConfig**](DisasterRecoveryAPI.md#PauseDrConfig) | **Post** /api/v1/customers/{cUUID}/dr_configs/{drUUID}/pause | Pause DR config
[**PauseDrUniverses**](DisasterRecoveryAPI.md#PauseDrUniverses) | **Post** /api/v1/customers/{cUUID}/dr_configs/{drUUID}/pause_universes | Pause DR config and universes associated with DR
[**ReplaceReplicaDrConfig**](DisasterRecoveryAPI.md#ReplaceReplicaDrConfig) | **Post** /api/v1/customers/{cUUID}/dr_configs/{drUUID}/replace_replica | Replace Replica in a disaster recovery config
[**RestartDrConfig**](DisasterRecoveryAPI.md#RestartDrConfig) | **Post** /api/v1/customers/{cUUID}/dr_configs/{drUUID}/restart | Restart disaster recovery config
[**ResumeDrConfig**](DisasterRecoveryAPI.md#ResumeDrConfig) | **Post** /api/v1/customers/{cUUID}/dr_configs/{drUUID}/resume | Resume DR config
[**ResumeDrUniverses**](DisasterRecoveryAPI.md#ResumeDrUniverses) | **Post** /api/v1/customers/{cUUID}/dr_configs/{drUUID}/resume_universes | Resume DR config and universes associated with DR
[**SetDatabasesDrConfig**](DisasterRecoveryAPI.md#SetDatabasesDrConfig) | **Put** /api/v1/customers/{cUUID}/dr_configs/{drUUID}/set_dbs | Set databases in disaster recovery config
[**SetTablesDrConfig**](DisasterRecoveryAPI.md#SetTablesDrConfig) | **Post** /api/v1/customers/{cUUID}/dr_configs/{drUUID}/set_tables | Set tables in disaster recovery config
[**SwitchoverDrConfig**](DisasterRecoveryAPI.md#SwitchoverDrConfig) | **Post** /api/v1/customers/{cUUID}/dr_configs/{drUUID}/switchover | Switchover a disaster recovery config
[**SyncDrConfig**](DisasterRecoveryAPI.md#SyncDrConfig) | **Post** /api/v1/customers/{cUUID}/dr_configs/{drUUID}/sync | Sync disaster recovery config



## CreateDrConfig

> YBPTask CreateDrConfig(ctx, cUUID).DisasterRecoveryCreateFormData(disasterRecoveryCreateFormData).Request(request).Execute()

Create disaster recovery config



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
	disasterRecoveryCreateFormData := *openapiclient.NewDrConfigCreateForm([]string{"Dbs_example"}, "Dr-config1", "SourceUniverseUUID_example", "TargetUniverseUUID_example") // DrConfigCreateForm | Disaster Recovery Create Form Data
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.CreateDrConfig(context.Background(), cUUID).DisasterRecoveryCreateFormData(disasterRecoveryCreateFormData).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.CreateDrConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDrConfig`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.CreateDrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDrConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disasterRecoveryCreateFormData** | [**DrConfigCreateForm**](DrConfigCreateForm.md) | Disaster Recovery Create Form Data | 
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


## DeleteXClusterConfig

> YBPTask DeleteXClusterConfig(ctx, cUUID, drUUID).IsForceDelete(isForceDelete).Request(request).Execute()

Delete xcluster config



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
	drUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	isForceDelete := true // bool |  (optional) (default to false)
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.DeleteXClusterConfig(context.Background(), cUUID, drUUID).IsForceDelete(isForceDelete).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.DeleteXClusterConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteXClusterConfig`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.DeleteXClusterConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**drUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteXClusterConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isForceDelete** | **bool** |  | [default to false]
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditDrConfig

> YBPTask EditDrConfig(ctx, cUUID, drUUID).DisasterRecoveryEditFormData(disasterRecoveryEditFormData).Request(request).Execute()

Edit disaster recovery config



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
	drUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	disasterRecoveryEditFormData := *openapiclient.NewDrConfigEditForm() // DrConfigEditForm | Disaster Recovery Edit Form Data
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.EditDrConfig(context.Background(), cUUID, drUUID).DisasterRecoveryEditFormData(disasterRecoveryEditFormData).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.EditDrConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditDrConfig`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.EditDrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**drUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditDrConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disasterRecoveryEditFormData** | [**DrConfigEditForm**](DrConfigEditForm.md) | Disaster Recovery Edit Form Data | 
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


## FailoverDrConfig

> YBPTask FailoverDrConfig(ctx, cUUID, drUUID).DisasterRecoveryFailoverFormData(disasterRecoveryFailoverFormData).Request(request).Execute()

Failover a disaster recovery config



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
	drUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	disasterRecoveryFailoverFormData := *openapiclient.NewDrConfigFailoverForm() // DrConfigFailoverForm | Disaster Recovery Failover Form Data
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.FailoverDrConfig(context.Background(), cUUID, drUUID).DisasterRecoveryFailoverFormData(disasterRecoveryFailoverFormData).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.FailoverDrConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FailoverDrConfig`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.FailoverDrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**drUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFailoverDrConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disasterRecoveryFailoverFormData** | [**DrConfigFailoverForm**](DrConfigFailoverForm.md) | Disaster Recovery Failover Form Data | 
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


## GetDrConfig

> DrConfigGetResp GetDrConfig(ctx, cUUID, drUUID).SyncWithDB(syncWithDB).Execute()

Get disaster recovery config



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
	drUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	syncWithDB := true // bool |  (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.GetDrConfig(context.Background(), cUUID, drUUID).SyncWithDB(syncWithDB).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.GetDrConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDrConfig`: DrConfigGetResp
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.GetDrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**drUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDrConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **syncWithDB** | **bool** |  | [default to true]

### Return type

[**DrConfigGetResp**](DrConfigGetResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDrConfigSafetime

> DrConfigSafetimeResp GetDrConfigSafetime(ctx, cUUID, drUUID).Execute()

Get disaster recovery config safetime



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
	drUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.GetDrConfigSafetime(context.Background(), cUUID, drUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.GetDrConfigSafetime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDrConfigSafetime`: DrConfigSafetimeResp
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.GetDrConfigSafetime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**drUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDrConfigSafetimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DrConfigSafetimeResp**](DrConfigSafetimeResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseDrConfig

> YBPTask PauseDrConfig(ctx, cUUID, drUUID).Request(request).Execute()

Pause DR config



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
	drUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.PauseDrConfig(context.Background(), cUUID, drUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.PauseDrConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PauseDrConfig`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.PauseDrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**drUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseDrConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseDrUniverses

> YBPTask PauseDrUniverses(ctx, cUUID, drUUID).Request(request).Execute()

Pause DR config and universes associated with DR



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
	drUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.PauseDrUniverses(context.Background(), cUUID, drUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.PauseDrUniverses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PauseDrUniverses`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.PauseDrUniverses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**drUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseDrUniversesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceReplicaDrConfig

> YBPTask ReplaceReplicaDrConfig(ctx, cUUID, drUUID).DisasterRecoveryReplaceReplicaFormData(disasterRecoveryReplaceReplicaFormData).Request(request).Execute()

Replace Replica in a disaster recovery config



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
	drUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	disasterRecoveryReplaceReplicaFormData := *openapiclient.NewDrConfigReplaceReplicaForm() // DrConfigReplaceReplicaForm | Disaster Recovery Replace Replica Form Data
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.ReplaceReplicaDrConfig(context.Background(), cUUID, drUUID).DisasterRecoveryReplaceReplicaFormData(disasterRecoveryReplaceReplicaFormData).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.ReplaceReplicaDrConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceReplicaDrConfig`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.ReplaceReplicaDrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**drUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceReplicaDrConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disasterRecoveryReplaceReplicaFormData** | [**DrConfigReplaceReplicaForm**](DrConfigReplaceReplicaForm.md) | Disaster Recovery Replace Replica Form Data | 
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


## RestartDrConfig

> YBPTask RestartDrConfig(ctx, cUUID, drUUID).DisasterRecoveryRestartFormData(disasterRecoveryRestartFormData).IsForceDelete(isForceDelete).Request(request).Execute()

Restart disaster recovery config



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
	drUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	disasterRecoveryRestartFormData := *openapiclient.NewDrConfigRestartForm([]string{"Dbs_example"}) // DrConfigRestartForm | Disaster Recovery Restart Form Data
	isForceDelete := true // bool |  (optional) (default to false)
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.RestartDrConfig(context.Background(), cUUID, drUUID).DisasterRecoveryRestartFormData(disasterRecoveryRestartFormData).IsForceDelete(isForceDelete).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.RestartDrConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestartDrConfig`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.RestartDrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**drUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartDrConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disasterRecoveryRestartFormData** | [**DrConfigRestartForm**](DrConfigRestartForm.md) | Disaster Recovery Restart Form Data | 
 **isForceDelete** | **bool** |  | [default to false]
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


## ResumeDrConfig

> YBPTask ResumeDrConfig(ctx, cUUID, drUUID).Request(request).Execute()

Resume DR config



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
	drUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.ResumeDrConfig(context.Background(), cUUID, drUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.ResumeDrConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResumeDrConfig`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.ResumeDrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**drUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeDrConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeDrUniverses

> YBPTask ResumeDrUniverses(ctx, cUUID, drUUID).Request(request).Execute()

Resume DR config and universes associated with DR



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
	drUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.ResumeDrUniverses(context.Background(), cUUID, drUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.ResumeDrUniverses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResumeDrUniverses`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.ResumeDrUniverses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**drUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeDrUniversesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDatabasesDrConfig

> YBPTask SetDatabasesDrConfig(ctx, cUUID, drUUID).DisasterRecoverySetDatabasesFormData(disasterRecoverySetDatabasesFormData).Request(request).Execute()

Set databases in disaster recovery config



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
	drUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	disasterRecoverySetDatabasesFormData := *openapiclient.NewDrConfigSetDatabasesForm() // DrConfigSetDatabasesForm | Disaster Recovery Set Databases Form Data
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.SetDatabasesDrConfig(context.Background(), cUUID, drUUID).DisasterRecoverySetDatabasesFormData(disasterRecoverySetDatabasesFormData).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.SetDatabasesDrConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDatabasesDrConfig`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.SetDatabasesDrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**drUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDatabasesDrConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disasterRecoverySetDatabasesFormData** | [**DrConfigSetDatabasesForm**](DrConfigSetDatabasesForm.md) | Disaster Recovery Set Databases Form Data | 
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


## SetTablesDrConfig

> YBPTask SetTablesDrConfig(ctx, cUUID, drUUID).DisasterRecoverySetTablesFormData(disasterRecoverySetTablesFormData).Request(request).Execute()

Set tables in disaster recovery config



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
	drUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	disasterRecoverySetTablesFormData := *openapiclient.NewDrConfigSetTablesForm() // DrConfigSetTablesForm | Disaster Recovery Set Tables Form Data
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.SetTablesDrConfig(context.Background(), cUUID, drUUID).DisasterRecoverySetTablesFormData(disasterRecoverySetTablesFormData).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.SetTablesDrConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetTablesDrConfig`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.SetTablesDrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**drUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetTablesDrConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disasterRecoverySetTablesFormData** | [**DrConfigSetTablesForm**](DrConfigSetTablesForm.md) | Disaster Recovery Set Tables Form Data | 
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


## SwitchoverDrConfig

> YBPTask SwitchoverDrConfig(ctx, cUUID, drUUID).DisasterRecoverySwitchoverFormData(disasterRecoverySwitchoverFormData).Request(request).Execute()

Switchover a disaster recovery config



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
	drUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	disasterRecoverySwitchoverFormData := *openapiclient.NewDrConfigSwitchoverForm() // DrConfigSwitchoverForm | Disaster Recovery Switchover Form Data
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.SwitchoverDrConfig(context.Background(), cUUID, drUUID).DisasterRecoverySwitchoverFormData(disasterRecoverySwitchoverFormData).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.SwitchoverDrConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SwitchoverDrConfig`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.SwitchoverDrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**drUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwitchoverDrConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disasterRecoverySwitchoverFormData** | [**DrConfigSwitchoverForm**](DrConfigSwitchoverForm.md) | Disaster Recovery Switchover Form Data | 
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


## SyncDrConfig

> YBPTask SyncDrConfig(ctx, cUUID, drUUID).Request(request).Execute()

Sync disaster recovery config



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
	drUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.SyncDrConfig(context.Background(), cUUID, drUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.SyncDrConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncDrConfig`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.SyncDrConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**drUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncDrConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

