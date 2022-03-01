# \LoggingConfigApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SetLoggingSettings**](LoggingConfigApi.md#SetLoggingSettings) | **Post** /api/v1/logging_config | Set Logging Level



## SetLoggingSettings

> PlatformLoggingConfig SetLoggingSettings(ctx).LoggingConfig(loggingConfig).Execute()

Set Logging Level

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
    loggingConfig := *openapiclient.NewPlatformLoggingConfig("Level_example", int32(123), "RolloverPattern_example") // PlatformLoggingConfig | Logging config to be updated

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoggingConfigApi.SetLoggingSettings(context.Background()).LoggingConfig(loggingConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingConfigApi.SetLoggingSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetLoggingSettings`: PlatformLoggingConfig
    fmt.Fprintf(os.Stdout, "Response from `LoggingConfigApi.SetLoggingSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetLoggingSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loggingConfig** | [**PlatformLoggingConfig**](PlatformLoggingConfig.md) | Logging config to be updated | 

### Return type

[**PlatformLoggingConfig**](PlatformLoggingConfig.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

