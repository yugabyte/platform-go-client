# \GrafanaDashboardApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GrafanaDashboard**](GrafanaDashboardApi.md#GrafanaDashboard) | **Get** /api/v1/grafana_dashboard | WARNING: This is a preview API that could change. Get Grafana Dashboard



## GrafanaDashboard

> string GrafanaDashboard(ctx).Execute()

WARNING: This is a preview API that could change. Get Grafana Dashboard

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
    resp, r, err := api_client.GrafanaDashboardApi.GrafanaDashboard(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GrafanaDashboardApi.GrafanaDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GrafanaDashboard`: string
    fmt.Fprintf(os.Stdout, "Response from `GrafanaDashboardApi.GrafanaDashboard`: %v\n", resp)
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

