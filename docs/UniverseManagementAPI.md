# \UniverseManagementAPI

All URIs are relative to *http://localhost*

| Method                                                                                | HTTP request                                                                | Description                                       |
| ------------------------------------------------------------------------------------- | --------------------------------------------------------------------------- | ------------------------------------------------- |
| [**ConfigureUniverseAlerts**](UniverseManagementAPI.md#ConfigureUniverseAlerts)       | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/config_alerts        | Configure alerts for a universe                   |
| [**DeleteUniverse**](UniverseManagementAPI.md#DeleteUniverse)                         | **Delete** /api/v1/customers/{cUUID}/universes/{uniUUID}                    | Delete a universe                                 |
| [**GetUniverse**](UniverseManagementAPI.md#GetUniverse)                               | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}                       | Get a universe                                    |
| [**ListUniverses**](UniverseManagementAPI.md#ListUniverses)                           | **Get** /api/v1/customers/{cUUID}/universes                                 | List universes                                    |
| [**PauseUniverse**](UniverseManagementAPI.md#PauseUniverse)                           | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/pause                | Pause a universe                                  |
| [**ResetUniverseVersion**](UniverseManagementAPI.md#ResetUniverseVersion)             | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/setup_universe_2dc    | Reset universe version                            |
| [**ResumeUniverse**](UniverseManagementAPI.md#ResumeUniverse)                         | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/resume               | Resume a paused universe                          |
| [**SetUniverseBackupFlag**](UniverseManagementAPI.md#SetUniverseBackupFlag)           | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/update_backup_state   | Set a universe&#39;s backup flag                  |
| [**SetUniverseHelm3Compatible**](UniverseManagementAPI.md#SetUniverseHelm3Compatible) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/mark_helm3_compatible | Flag a universe as Helm 3-compatible - deprecated |
| [**SetUniverseKey**](UniverseManagementAPI.md#SetUniverseKey)                         | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/set_key              | Set a universe&#39;s key                          |
| [**UpdateLoadBalancerConfig**](UniverseManagementAPI.md#UpdateLoadBalancerConfig)     | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/update_lb_config      | Update load balancer config                       |



## ConfigureUniverseAlerts

> YBPSuccess ConfigureUniverseAlerts(ctx, cUUID, uniUUID).Request(request).Execute()

Configure alerts for a universe



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
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseManagementAPI.ConfigureUniverseAlerts(context.Background(), cUUID, uniUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementAPI.ConfigureUniverseAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigureUniverseAlerts`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `UniverseManagementAPI.ConfigureUniverseAlerts`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiConfigureUniverseAlertsRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


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


## DeleteUniverse

> YBPTask DeleteUniverse(ctx, cUUID, uniUUID).IsForceDelete(isForceDelete).IsDeleteBackups(isDeleteBackups).IsDeleteAssociatedCerts(isDeleteAssociatedCerts).Request(request).Execute()

Delete a universe



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
	isForceDelete := true // bool |  (optional) (default to false)
	isDeleteBackups := true // bool |  (optional) (default to false)
	isDeleteAssociatedCerts := true // bool |  (optional) (default to false)
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseManagementAPI.DeleteUniverse(context.Background(), cUUID, uniUUID).IsForceDelete(isForceDelete).IsDeleteBackups(isDeleteBackups).IsDeleteAssociatedCerts(isDeleteAssociatedCerts).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementAPI.DeleteUniverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUniverse`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `UniverseManagementAPI.DeleteUniverse`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUniverseRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **isForceDelete** | **bool** |  | [default to false]
 **isDeleteBackups** | **bool** |  | [default to false]
 **isDeleteAssociatedCerts** | **bool** |  | [default to false]
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


## GetUniverse

> UniverseResp GetUniverse(ctx, cUUID, uniUUID).Execute()

Get a universe



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
	resp, r, err := apiClient.UniverseManagementAPI.GetUniverse(context.Background(), cUUID, uniUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementAPI.GetUniverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverse`: UniverseResp
	fmt.Fprintf(os.Stdout, "Response from `UniverseManagementAPI.GetUniverse`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

[**UniverseResp**](UniverseResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUniverses

> []UniverseResp ListUniverses(ctx, cUUID).Name(name).Execute()

List universes



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
	name := "name_example" // string |  (optional) (default to "null")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseManagementAPI.ListUniverses(context.Background(), cUUID).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementAPI.ListUniverses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUniverses`: []UniverseResp
	fmt.Fprintf(os.Stdout, "Response from `UniverseManagementAPI.ListUniverses`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiListUniversesRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **name** | **string** |  | [default to &quot;null&quot;]

### Return type

[**[]UniverseResp**](UniverseResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseUniverse

> YBPTask PauseUniverse(ctx, cUUID, uniUUID).Request(request).Execute()

Pause a universe



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
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseManagementAPI.PauseUniverse(context.Background(), cUUID, uniUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementAPI.PauseUniverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PauseUniverse`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `UniverseManagementAPI.PauseUniverse`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiPauseUniverseRequest struct via the builder pattern


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


## ResetUniverseVersion

> YBPSuccess ResetUniverseVersion(ctx, cUUID, uniUUID).Request(request).Execute()

Reset universe version



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
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseManagementAPI.ResetUniverseVersion(context.Background(), cUUID, uniUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementAPI.ResetUniverseVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetUniverseVersion`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `UniverseManagementAPI.ResetUniverseVersion`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiResetUniverseVersionRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


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


## ResumeUniverse

> YBPTask ResumeUniverse(ctx, cUUID, uniUUID).Request(request).Execute()

Resume a paused universe



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
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseManagementAPI.ResumeUniverse(context.Background(), cUUID, uniUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementAPI.ResumeUniverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResumeUniverse`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `UniverseManagementAPI.ResumeUniverse`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiResumeUniverseRequest struct via the builder pattern


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
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	markActive := true // bool |  (optional)
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseManagementAPI.SetUniverseBackupFlag(context.Background(), cUUID, uniUUID).MarkActive(markActive).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementAPI.SetUniverseBackupFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetUniverseBackupFlag`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `UniverseManagementAPI.SetUniverseBackupFlag`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiSetUniverseBackupFlagRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


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


## SetUniverseHelm3Compatible

> YBPSuccess SetUniverseHelm3Compatible(ctx, cUUID, uniUUID).Request(request).Execute()

Flag a universe as Helm 3-compatible - deprecated



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
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseManagementAPI.SetUniverseHelm3Compatible(context.Background(), cUUID, uniUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementAPI.SetUniverseHelm3Compatible``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetUniverseHelm3Compatible`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `UniverseManagementAPI.SetUniverseHelm3Compatible`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiSetUniverseHelm3CompatibleRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


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


## SetUniverseKey

> UniverseResp SetUniverseKey(ctx, cUUID, uniUUID).SetUniverseKeyRequest(setUniverseKeyRequest).Request(request).Execute()

Set a universe's key



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
	setUniverseKeyRequest := *openapiclient.NewEncryptionAtRestConfig() // EncryptionAtRestConfig | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseManagementAPI.SetUniverseKey(context.Background(), cUUID, uniUUID).SetUniverseKeyRequest(setUniverseKeyRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementAPI.SetUniverseKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetUniverseKey`: UniverseResp
	fmt.Fprintf(os.Stdout, "Response from `UniverseManagementAPI.SetUniverseKey`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiSetUniverseKeyRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **setUniverseKeyRequest** | [**EncryptionAtRestConfig**](EncryptionAtRestConfig.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**UniverseResp**](UniverseResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoadBalancerConfig

> UpdateLoadBalancerConfig UpdateLoadBalancerConfig(ctx, cUUID, uniUUID).Request(request).Execute()

Update load balancer config



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
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseManagementAPI.UpdateLoadBalancerConfig(context.Background(), cUUID, uniUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementAPI.UpdateLoadBalancerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLoadBalancerConfig`: UpdateLoadBalancerConfig
	fmt.Fprintf(os.Stdout, "Response from `UniverseManagementAPI.UpdateLoadBalancerConfig`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoadBalancerConfigRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**UpdateLoadBalancerConfig**](UpdateLoadBalancerConfig.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

