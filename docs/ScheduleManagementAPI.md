# \ScheduleManagementAPI

All URIs are relative to *http://localhost*

| Method                                                                              | HTTP request                                                                                            | Description                            |
| ----------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | -------------------------------------- |
| [**DeleteBackupScheduleAsync**](ScheduleManagementAPI.md#DeleteBackupScheduleAsync) | **Delete** /api/v1/customers/{cUUID}/universes/{uniUUID}/schedules/{sUUID}/delete_backup_schedule_async | Delete a backup schedule async         |
| [**DeleteSchedule**](ScheduleManagementAPI.md#DeleteSchedule)                       | **Delete** /api/v1/customers/{cUUID}/schedules/{sUUID}                                                  | Delete a schedule  - deprecated        |
| [**DeleteScheduleV2**](ScheduleManagementAPI.md#DeleteScheduleV2)                   | **Delete** /api/v1/customers/{cUUID}/schedules/{sUUID}/delete                                           | Delete a schedule V2 - deprecated      |
| [**EditBackupScheduleAsync**](ScheduleManagementAPI.md#EditBackupScheduleAsync)     | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/schedules/{sUUID}/edit_backup_schedule_async      | Edit a backup schedule async           |
| [**EditBackupScheduleV2**](ScheduleManagementAPI.md#EditBackupScheduleV2)           | **Put** /api/v1/customers/{cUUID}/schedules/{sUUID}                                                     | Edit a backup schedule V2 - deprecated |
| [**GetSchedule**](ScheduleManagementAPI.md#GetSchedule)                             | **Get** /api/v1/customers/{cUUID}/schedules/{sUUID}                                                     | Get Schedule                           |
| [**ListSchedules**](ScheduleManagementAPI.md#ListSchedules)                         | **Get** /api/v1/customers/{cUUID}/schedules                                                             | List schedules - deprecated            |
| [**ListSchedulesV2**](ScheduleManagementAPI.md#ListSchedulesV2)                     | **Post** /api/v1/customers/{cUUID}/schedules/page                                                       | List schedules V2                      |
| [**ToggleBackupSchedule**](ScheduleManagementAPI.md#ToggleBackupSchedule)           | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/schedules/{sUUID}/pause_resume                    | Toggle a backup schedule               |



## DeleteBackupScheduleAsync

> Schedule DeleteBackupScheduleAsync(ctx, cUUID, uniUUID, sUUID).Request(request).Execute()

Delete a backup schedule async



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
	sUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleManagementAPI.DeleteBackupScheduleAsync(context.Background(), cUUID, uniUUID, sUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleManagementAPI.DeleteBackupScheduleAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBackupScheduleAsync`: Schedule
	fmt.Fprintf(os.Stdout, "Response from `ScheduleManagementAPI.DeleteBackupScheduleAsync`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |
| **sUUID**   | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackupScheduleAsyncRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSchedule

> YBPSuccess DeleteSchedule(ctx, cUUID, sUUID).Request(request).Execute()

Delete a schedule  - deprecated



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
	sUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleManagementAPI.DeleteSchedule(context.Background(), cUUID, sUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleManagementAPI.DeleteSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSchedule`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `ScheduleManagementAPI.DeleteSchedule`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **sUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScheduleRequest struct via the builder pattern


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


## DeleteScheduleV2

> YBPSuccess DeleteScheduleV2(ctx, cUUID, sUUID).Request(request).Execute()

Delete a schedule V2 - deprecated



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
	sUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleManagementAPI.DeleteScheduleV2(context.Background(), cUUID, sUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleManagementAPI.DeleteScheduleV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteScheduleV2`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `ScheduleManagementAPI.DeleteScheduleV2`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **sUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScheduleV2Request struct via the builder pattern


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


## EditBackupScheduleAsync

> Schedule EditBackupScheduleAsync(ctx, cUUID, uniUUID, sUUID).Body(body).Request(request).Execute()

Edit a backup schedule async



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
	sUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	body := *openapiclient.NewBackupScheduleEditParams() // BackupScheduleEditParams | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleManagementAPI.EditBackupScheduleAsync(context.Background(), cUUID, uniUUID, sUUID).Body(body).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleManagementAPI.EditBackupScheduleAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditBackupScheduleAsync`: Schedule
	fmt.Fprintf(os.Stdout, "Response from `ScheduleManagementAPI.EditBackupScheduleAsync`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |
| **sUUID**   | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiEditBackupScheduleAsyncRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



 **body** | [**BackupScheduleEditParams**](BackupScheduleEditParams.md) |  | 
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


## EditBackupScheduleV2

> Schedule EditBackupScheduleV2(ctx, cUUID, sUUID).Body(body).Request(request).Execute()

Edit a backup schedule V2 - deprecated



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
	sUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	body := *openapiclient.NewEditBackupScheduleParams() // EditBackupScheduleParams | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleManagementAPI.EditBackupScheduleV2(context.Background(), cUUID, sUUID).Body(body).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleManagementAPI.EditBackupScheduleV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditBackupScheduleV2`: Schedule
	fmt.Fprintf(os.Stdout, "Response from `ScheduleManagementAPI.EditBackupScheduleV2`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **sUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiEditBackupScheduleV2Request struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **body** | [**EditBackupScheduleParams**](EditBackupScheduleParams.md) |  | 
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


## GetSchedule

> Schedule GetSchedule(ctx, cUUID, sUUID).Execute()

Get Schedule

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
	sUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleManagementAPI.GetSchedule(context.Background(), cUUID, sUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleManagementAPI.GetSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSchedule`: Schedule
	fmt.Fprintf(os.Stdout, "Response from `ScheduleManagementAPI.GetSchedule`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **sUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduleRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

[**Schedule**](Schedule.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSchedules

> []Schedule ListSchedules(ctx, cUUID).Execute()

List schedules - deprecated



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleManagementAPI.ListSchedules(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleManagementAPI.ListSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSchedules`: []Schedule
	fmt.Fprintf(os.Stdout, "Response from `ScheduleManagementAPI.ListSchedules`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiListSchedulesRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


### Return type

[**[]Schedule**](Schedule.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSchedulesV2

> SchedulePagedResponse ListSchedulesV2(ctx, cUUID).PageScheduleRequest(pageScheduleRequest).Request(request).Execute()

List schedules V2

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
	pageScheduleRequest := *openapiclient.NewSchedulePagedApiQuery("Direction_example", *openapiclient.NewScheduleApiFilter([]string{"Status_example"}, []string{"TaskTypes_example"}, []string{"UniverseUUIDList_example"}), int32(123), false, int32(123), "SortBy_example") // SchedulePagedApiQuery | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleManagementAPI.ListSchedulesV2(context.Background(), cUUID).PageScheduleRequest(pageScheduleRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleManagementAPI.ListSchedulesV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSchedulesV2`: SchedulePagedResponse
	fmt.Fprintf(os.Stdout, "Response from `ScheduleManagementAPI.ListSchedulesV2`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiListSchedulesV2Request struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **pageScheduleRequest** | [**SchedulePagedApiQuery**](SchedulePagedApiQuery.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**SchedulePagedResponse**](SchedulePagedResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleBackupSchedule

> Schedule ToggleBackupSchedule(ctx, cUUID, uniUUID, sUUID).Body(body).Request(request).Execute()

Toggle a backup schedule



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
	sUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	body := *openapiclient.NewBackupScheduleToggleParams("Status_example") // BackupScheduleToggleParams | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduleManagementAPI.ToggleBackupSchedule(context.Background(), cUUID, uniUUID, sUUID).Body(body).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduleManagementAPI.ToggleBackupSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToggleBackupSchedule`: Schedule
	fmt.Fprintf(os.Stdout, "Response from `ScheduleManagementAPI.ToggleBackupSchedule`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |
| **sUUID**   | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiToggleBackupScheduleRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



 **body** | [**BackupScheduleToggleParams**](BackupScheduleToggleParams.md) |  | 
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

