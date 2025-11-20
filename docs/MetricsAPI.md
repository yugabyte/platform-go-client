# \MetricsAPI

All URIs are relative to *http://localhost*

| Method                                           | HTTP request                       | Description            |
| ------------------------------------------------ | ---------------------------------- | ---------------------- |
| [**MetricsDetail**](MetricsAPI.md#MetricsDetail) | **Get** /api/v1/prometheus_metrics | Get Prometheus metrics |



## MetricsDetail

> string MetricsDetail(ctx).Request(request).Execute()

Get Prometheus metrics



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
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.MetricsDetail(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.MetricsDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricsDetail`: string
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.MetricsDetail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricsDetailRequest struct via the builder pattern


| Name        | Type                              | Description | Notes |
| ----------- | --------------------------------- | ----------- | ----- |
| **request** | [**interface{}**](interface{}.md) |             |

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

