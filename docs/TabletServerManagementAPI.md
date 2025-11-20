# \TabletServerManagementAPI

All URIs are relative to *http://localhost*

| Method                                                                  | HTTP request                                                         | Description                         |
| ----------------------------------------------------------------------- | -------------------------------------------------------------------- | ----------------------------------- |
| [**ListTabletServers**](TabletServerManagementAPI.md#ListTabletServers) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/tablet-servers | List all tablet servers information |



## ListTabletServers

> map[string]map[string]interface{} ListTabletServers(ctx, cUUID, uniUUID).Execute()

List all tablet servers information



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
	resp, r, err := apiClient.TabletServerManagementAPI.ListTabletServers(context.Background(), cUUID, uniUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TabletServerManagementAPI.ListTabletServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTabletServers`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TabletServerManagementAPI.ListTabletServers`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiListTabletServersRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

**map[string]map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

