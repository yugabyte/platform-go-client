# \LoggingConfigApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SetAuditLoggingSettings**](LoggingConfigApi.md#SetAuditLoggingSettings) | **Post** /api/v1/audit_logging_config | Set Audit Logging Level
[**SetLoggingSettings**](LoggingConfigApi.md#SetLoggingSettings) | **Post** /api/v1/logging_config | Set Logging Level



## SetAuditLoggingSettings

> AuditLoggingConfig SetAuditLoggingSettings(ctx).AuditLoggingConfig(auditLoggingConfig).Request(request).Execute()

Set Audit Logging Level

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
    auditLoggingConfig := *openapiclient.NewAuditLoggingConfig(false, false) // AuditLoggingConfig | Audit Logging config to be updated
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoggingConfigApi.SetAuditLoggingSettings(context.Background()).AuditLoggingConfig(auditLoggingConfig).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingConfigApi.SetAuditLoggingSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetAuditLoggingSettings`: AuditLoggingConfig
    fmt.Fprintf(os.Stdout, "Response from `LoggingConfigApi.SetAuditLoggingSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetAuditLoggingSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auditLoggingConfig** | [**AuditLoggingConfig**](AuditLoggingConfig.md) | Audit Logging config to be updated | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**AuditLoggingConfig**](AuditLoggingConfig.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetLoggingSettings

> PlatformLoggingConfig SetLoggingSettings(ctx).LoggingConfig(loggingConfig).Request(request).Execute()

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LoggingConfigApi.SetLoggingSettings(context.Background()).LoggingConfig(loggingConfig).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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

