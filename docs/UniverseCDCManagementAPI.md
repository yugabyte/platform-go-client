# \UniverseCDCManagementAPI

All URIs are relative to *http://localhost*

| Method                                                                     | HTTP request                                                                    | Description                             |
| -------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | --------------------------------------- |
| [**CreateCdcStream**](UniverseCDCManagementAPI.md#CreateCdcStream)         | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/cdc_streams              | Create CDC Stream for a cluster         |
| [**DeleteCdcStream**](UniverseCDCManagementAPI.md#DeleteCdcStream)         | **Delete** /api/v1/customers/{cUUID}/universes/{uniUUID}/cdc_streams/{streamId} | Delete a CDC stream for a cluster       |
| [**ListCdcStreams**](UniverseCDCManagementAPI.md#ListCdcStreams)           | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/cdc_streams               | List CDC Streams for a cluster          |
| [**ListReplicationSlot**](UniverseCDCManagementAPI.md#ListReplicationSlot) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/cdc_replication_slots     | List CDC Replication slot for a cluster |



## CreateCdcStream

> CreateCdcStream(ctx, cUUID, uniUUID).Request(request).Execute()

Create CDC Stream for a cluster



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
	r, err := apiClient.UniverseCDCManagementAPI.CreateCdcStream(context.Background(), cUUID, uniUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseCDCManagementAPI.CreateCdcStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCdcStreamRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


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


## DeleteCdcStream

> DeleteCdcStream(ctx, cUUID, uniUUID, streamId).Execute()

Delete a CDC stream for a cluster



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
	streamId := "streamId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UniverseCDCManagementAPI.DeleteCdcStream(context.Background(), cUUID, uniUUID, streamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseCDCManagementAPI.DeleteCdcStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**    | **string**          |                                                                             |
| **uniUUID**  | **string**          |                                                                             |
| **streamId** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCdcStreamRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |




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


## ListCdcStreams

> ListCdcStreams(ctx, cUUID, uniUUID).Execute()

List CDC Streams for a cluster



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
	r, err := apiClient.UniverseCDCManagementAPI.ListCdcStreams(context.Background(), cUUID, uniUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseCDCManagementAPI.ListCdcStreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiListCdcStreamsRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



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


## ListReplicationSlot

> ListReplicationSlot(ctx, cUUID, uniUUID).Execute()

List CDC Replication slot for a cluster



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
	r, err := apiClient.UniverseCDCManagementAPI.ListReplicationSlot(context.Background(), cUUID, uniUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseCDCManagementAPI.ListReplicationSlot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiListReplicationSlotRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



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

