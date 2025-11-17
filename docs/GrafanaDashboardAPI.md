# \GrafanaDashboardAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GrafanaDashboard**](GrafanaDashboardAPI.md#GrafanaDashboard) | **Get** /api/v1/grafana_dashboard | Get Grafana Dashboard



## GrafanaDashboard

> string GrafanaDashboard(ctx).Execute()

Get Grafana Dashboard



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GrafanaDashboardAPI.GrafanaDashboard(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GrafanaDashboardAPI.GrafanaDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GrafanaDashboard`: string
	fmt.Fprintf(os.Stdout, "Response from `GrafanaDashboardAPI.GrafanaDashboard`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGrafanaDashboardRequest struct via the builder pattern


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

