# \YBAInstanceApi

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetYBAInstanceInfo**](YBAInstanceApi.md#GetYBAInstanceInfo) | **Get** /yba-info | Get YBAInstance info



## GetYBAInstanceInfo

> YBAInfo GetYBAInstanceInfo(ctx).Execute()

Get YBAInstance info



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
    resp, r, err := api_client.YBAInstanceApi.GetYBAInstanceInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `YBAInstanceApi.GetYBAInstanceInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetYBAInstanceInfo`: YBAInfo
    fmt.Fprintf(os.Stdout, "Response from `YBAInstanceApi.GetYBAInstanceInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetYBAInstanceInfoRequest struct via the builder pattern


### Return type

[**YBAInfo**](YBAInfo.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

