# \MetricsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetricsDetail**](MetricsApi.md#MetricsDetail) | **Get** /api/v1/prometheus_metrics | Available since YBA version 2.8.0.0. Get Prometheus metrics



## MetricsDetail

> string MetricsDetail(ctx).Execute()

Available since YBA version 2.8.0.0. Get Prometheus metrics

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.MetricsDetail(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.MetricsDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetricsDetail`: string
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.MetricsDetail`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMetricsDetailRequest struct via the builder pattern


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

