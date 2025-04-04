# \MetricsApi

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPrometheusHostInfo**](MetricsApi.md#GetPrometheusHostInfo) | **Get** /host-info | Get Prometheus host info



## GetPrometheusHostInfo

> PrometheusHostInfo GetPrometheusHostInfo(ctx).Execute()

Get Prometheus host info



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
    resp, r, err := api_client.MetricsApi.GetPrometheusHostInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.GetPrometheusHostInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrometheusHostInfo`: PrometheusHostInfo
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.GetPrometheusHostInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrometheusHostInfoRequest struct via the builder pattern


### Return type

[**PrometheusHostInfo**](PrometheusHostInfo.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

