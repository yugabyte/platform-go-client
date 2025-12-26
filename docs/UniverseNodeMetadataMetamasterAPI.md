# \UniverseNodeMetadataMetamasterAPI

All URIs are relative to *http://localhost*

| Method                                                                                      | HTTP request                                                          | Description                                |
| ------------------------------------------------------------------------------------------- | --------------------------------------------------------------------- | ------------------------------------------ |
| [**GetMasterAddresses**](UniverseNodeMetadataMetamasterAPI.md#GetMasterAddresses)           | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/masters         | List a master node&#39;s addresses         |
| [**GetMasterLBState**](UniverseNodeMetadataMetamasterAPI.md#GetMasterLBState)               | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/master_lb_state | Get the state of master load balancing ops |
| [**GetRedisServerAddresses**](UniverseNodeMetadataMetamasterAPI.md#GetRedisServerAddresses) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/redisservers    | List a REDIS server&#39;s addresses        |
| [**GetUniverseMasterNodes**](UniverseNodeMetadataMetamasterAPI.md#GetUniverseMasterNodes)   | **Get** /metamaster/universe/{universeUUID}                           | List a universe&#39;s master nodes         |
| [**GetYQLServerAddresses**](UniverseNodeMetadataMetamasterAPI.md#GetYQLServerAddresses)     | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/yqlservers      | List a YQL server&#39;s addresses          |
| [**GetYSQLServerAddresses**](UniverseNodeMetadataMetamasterAPI.md#GetYSQLServerAddresses)   | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/ysqlservers     | List a YSQL server&#39;s addresses         |



## GetMasterAddresses

> string GetMasterAddresses(ctx, cUUID, uniUUID).Execute()

List a master node's addresses



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
	resp, r, err := apiClient.UniverseNodeMetadataMetamasterAPI.GetMasterAddresses(context.Background(), cUUID, uniUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseNodeMetadataMetamasterAPI.GetMasterAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMasterAddresses`: string
	fmt.Fprintf(os.Stdout, "Response from `UniverseNodeMetadataMetamasterAPI.GetMasterAddresses`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetMasterAddressesRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

**string**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMasterLBState

> MasterLBStateResponse GetMasterLBState(ctx, cUUID, uniUUID).Execute()

Get the state of master load balancing ops



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
	resp, r, err := apiClient.UniverseNodeMetadataMetamasterAPI.GetMasterLBState(context.Background(), cUUID, uniUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseNodeMetadataMetamasterAPI.GetMasterLBState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMasterLBState`: MasterLBStateResponse
	fmt.Fprintf(os.Stdout, "Response from `UniverseNodeMetadataMetamasterAPI.GetMasterLBState`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetMasterLBStateRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

[**MasterLBStateResponse**](MasterLBStateResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRedisServerAddresses

> string GetRedisServerAddresses(ctx, cUUID, uniUUID).Execute()

List a REDIS server's addresses



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
	resp, r, err := apiClient.UniverseNodeMetadataMetamasterAPI.GetRedisServerAddresses(context.Background(), cUUID, uniUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseNodeMetadataMetamasterAPI.GetRedisServerAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedisServerAddresses`: string
	fmt.Fprintf(os.Stdout, "Response from `UniverseNodeMetadataMetamasterAPI.GetRedisServerAddresses`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedisServerAddressesRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

**string**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniverseMasterNodes

> MastersList GetUniverseMasterNodes(ctx, universeUUID).Execute()

List a universe's master nodes

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
	universeUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseNodeMetadataMetamasterAPI.GetUniverseMasterNodes(context.Background(), universeUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseNodeMetadataMetamasterAPI.GetUniverseMasterNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniverseMasterNodes`: MastersList
	fmt.Fprintf(os.Stdout, "Response from `UniverseNodeMetadataMetamasterAPI.GetUniverseMasterNodes`: %v\n", resp)
}
```

### Path Parameters


| Name             | Type                | Description                                                                 | Notes |
| ---------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**          | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **universeUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseMasterNodesRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


### Return type

[**MastersList**](MastersList.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetYQLServerAddresses

> string GetYQLServerAddresses(ctx, cUUID, uniUUID).Execute()

List a YQL server's addresses



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
	resp, r, err := apiClient.UniverseNodeMetadataMetamasterAPI.GetYQLServerAddresses(context.Background(), cUUID, uniUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseNodeMetadataMetamasterAPI.GetYQLServerAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetYQLServerAddresses`: string
	fmt.Fprintf(os.Stdout, "Response from `UniverseNodeMetadataMetamasterAPI.GetYQLServerAddresses`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetYQLServerAddressesRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

**string**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetYSQLServerAddresses

> string GetYSQLServerAddresses(ctx, cUUID, uniUUID).Execute()

List a YSQL server's addresses



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
	resp, r, err := apiClient.UniverseNodeMetadataMetamasterAPI.GetYSQLServerAddresses(context.Background(), cUUID, uniUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseNodeMetadataMetamasterAPI.GetYSQLServerAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetYSQLServerAddresses`: string
	fmt.Fprintf(os.Stdout, "Response from `UniverseNodeMetadataMetamasterAPI.GetYSQLServerAddresses`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetYSQLServerAddressesRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

**string**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

