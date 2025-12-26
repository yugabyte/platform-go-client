# \PITRManagementAPI

All URIs are relative to *http://localhost*

| Method                                                          | HTTP request                                                                                            | Description                                     |
| --------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ----------------------------------------------- |
| [**CloneNamespace**](PITRManagementAPI.md#CloneNamespace)       | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/clone                                            | Clone namespace via PITR on a universe          |
| [**CreatePitrConfig**](PITRManagementAPI.md#CreatePitrConfig)   | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/keyspaces/{tableType}/{keyspaceName}/pitr_config | Create pitr config for a keyspace in a universe |
| [**DeletePitrConfig**](PITRManagementAPI.md#DeletePitrConfig)   | **Delete** /api/v1/customers/{cUUID}/universes/{uniUUID}/pitr_config/{pUUID}                            | Delete pitr config on a universe                |
| [**ListOfPitrConfigs**](PITRManagementAPI.md#ListOfPitrConfigs) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/pitr_config                                       | List the PITR configs of a universe             |
| [**PerformPitr**](PITRManagementAPI.md#PerformPitr)             | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/pitr                                             | Perform PITR on a universe                      |
| [**UpdatePitrConfig**](PITRManagementAPI.md#UpdatePitrConfig)   | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/pitr_config/{pUUID}                               | Update pitr config for a keyspace in a universe |



## CloneNamespace

> YBPTask CloneNamespace(ctx, cUUID, uniUUID).NamespaceClone(namespaceClone).Request(request).Execute()

Clone namespace via PITR on a universe



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	namespaceClone := *openapiclient.NewCloneNamespaceParams(*openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), "PlatformUrl_example", int32(123), int32(123)) // CloneNamespaceParams | perform clone via PITR
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PITRManagementAPI.CloneNamespace(context.Background(), cUUID, uniUUID).NamespaceClone(namespaceClone).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PITRManagementAPI.CloneNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloneNamespace`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `PITRManagementAPI.CloneNamespace`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiCloneNamespaceRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **namespaceClone** | [**CloneNamespaceParams**](CloneNamespaceParams.md) | perform clone via PITR | 
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


## CreatePitrConfig

> YBPTask CreatePitrConfig(ctx, cUUID, uniUUID, tableType, keyspaceName).PitrConfig(pitrConfig).Request(request).Execute()

Create pitr config for a keyspace in a universe

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tableType := "tableType_example" // string | 
	keyspaceName := "keyspaceName_example" // string | 
	pitrConfig := *openapiclient.NewCreatePitrConfigParams(*openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), "PlatformUrl_example", int32(123), int32(123)) // CreatePitrConfigParams | post pitr config
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PITRManagementAPI.CreatePitrConfig(context.Background(), cUUID, uniUUID, tableType, keyspaceName).PitrConfig(pitrConfig).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PITRManagementAPI.CreatePitrConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePitrConfig`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `PITRManagementAPI.CreatePitrConfig`: %v\n", resp)
}
```

### Path Parameters


| Name             | Type                | Description                                                                 | Notes |
| ---------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**          | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**        | **string**          |                                                                             |
| **uniUUID**      | **string**          |                                                                             |
| **tableType**    | **string**          |                                                                             |
| **keyspaceName** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePitrConfigRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |




 **pitrConfig** | [**CreatePitrConfigParams**](CreatePitrConfigParams.md) | post pitr config | 
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


## DeletePitrConfig

> YBPTask DeletePitrConfig(ctx, cUUID, uniUUID, pUUID).Request(request).Execute()

Delete pitr config on a universe

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PITRManagementAPI.DeletePitrConfig(context.Background(), cUUID, uniUUID, pUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PITRManagementAPI.DeletePitrConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePitrConfig`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `PITRManagementAPI.DeletePitrConfig`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |
| **pUUID**   | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePitrConfigRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



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


## ListOfPitrConfigs

> []PitrConfig ListOfPitrConfigs(ctx, cUUID, uniUUID).Execute()

List the PITR configs of a universe

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PITRManagementAPI.ListOfPitrConfigs(context.Background(), cUUID, uniUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PITRManagementAPI.ListOfPitrConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOfPitrConfigs`: []PitrConfig
	fmt.Fprintf(os.Stdout, "Response from `PITRManagementAPI.ListOfPitrConfigs`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiListOfPitrConfigsRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

[**[]PitrConfig**](PitrConfig.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PerformPitr

> YBPTask PerformPitr(ctx, cUUID, uniUUID).PerformPitr(performPitr).Request(request).Execute()

Perform PITR on a universe

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	performPitr := *openapiclient.NewRestoreSnapshotScheduleParams(*openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), "PlatformUrl_example", int32(123), int32(123)) // RestoreSnapshotScheduleParams | perform PITR
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PITRManagementAPI.PerformPitr(context.Background(), cUUID, uniUUID).PerformPitr(performPitr).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PITRManagementAPI.PerformPitr``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PerformPitr`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `PITRManagementAPI.PerformPitr`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiPerformPitrRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **performPitr** | [**RestoreSnapshotScheduleParams**](RestoreSnapshotScheduleParams.md) | perform PITR | 
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


## UpdatePitrConfig

> YBPTask UpdatePitrConfig(ctx, cUUID, uniUUID, pUUID).PitrConfig(pitrConfig).Request(request).Execute()

Update pitr config for a keyspace in a universe



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pitrConfig := *openapiclient.NewUpdatePitrConfigParams(*openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), "PlatformUrl_example", int32(123), int32(123)) // UpdatePitrConfigParams | put pitr config
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PITRManagementAPI.UpdatePitrConfig(context.Background(), cUUID, uniUUID, pUUID).PitrConfig(pitrConfig).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PITRManagementAPI.UpdatePitrConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePitrConfig`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `PITRManagementAPI.UpdatePitrConfig`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |
| **pUUID**   | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePitrConfigRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



 **pitrConfig** | [**UpdatePitrConfigParams**](UpdatePitrConfigParams.md) | put pitr config | 
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

