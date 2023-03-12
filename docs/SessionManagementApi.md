# \SessionManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiLogin**](SessionManagementApi.md#ApiLogin) | **Post** /api/v1/api_login | Authenticate user and return api token
[**AppVersion**](SessionManagementApi.md#AppVersion) | **Get** /api/v1/app_version | appVersion
[**CustomerCount**](SessionManagementApi.md#CustomerCount) | **Get** /api/v1/customer_count | customerCount
[**GetAdminNotifications**](SessionManagementApi.md#GetAdminNotifications) | **Get** /api/v1/customers/{cUUID}/admin_notifications | getAdminNotifications
[**GetFilteredLogs**](SessionManagementApi.md#GetFilteredLogs) | **Get** /api/v1/logs | getFilteredLogs
[**GetLogs**](SessionManagementApi.md#GetLogs) | **Get** /api/v1/logs/{maxLines} | getLogs
[**GetSessionInfo**](SessionManagementApi.md#GetSessionInfo) | **Get** /api/v1/session_info | Get current user/customer uuid auth/api token
[**RegisterCustomer**](SessionManagementApi.md#RegisterCustomer) | **Post** /api/v1/register | Register a customer



## ApiLogin

> SessionInfo ApiLogin(ctx).CustomerLoginFormData(customerLoginFormData).Execute()

Authenticate user and return api token

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
    customerLoginFormData := *openapiclient.NewCustomerLoginFormData("Email_example", "Password_example") // CustomerLoginFormData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SessionManagementApi.ApiLogin(context.Background()).CustomerLoginFormData(customerLoginFormData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionManagementApi.ApiLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiLogin`: SessionInfo
    fmt.Fprintf(os.Stdout, "Response from `SessionManagementApi.ApiLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerLoginFormData** | [**CustomerLoginFormData**](CustomerLoginFormData.md) |  | 

### Return type

[**SessionInfo**](SessionInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppVersion

> map[string]string AppVersion(ctx).Execute()

appVersion

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
    resp, r, err := api_client.SessionManagementApi.AppVersion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionManagementApi.AppVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVersion`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `SessionManagementApi.AppVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppVersionRequest struct via the builder pattern


### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerCount

> CustomerCountResp CustomerCount(ctx).Execute()

customerCount

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
    resp, r, err := api_client.SessionManagementApi.CustomerCount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionManagementApi.CustomerCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerCount`: CustomerCountResp
    fmt.Fprintf(os.Stdout, "Response from `SessionManagementApi.CustomerCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerCountRequest struct via the builder pattern


### Return type

[**CustomerCountResp**](CustomerCountResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminNotifications

> CurrentAdminNotificationMessages GetAdminNotifications(ctx, cUUID).Execute()

getAdminNotifications

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
    cUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SessionManagementApi.GetAdminNotifications(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionManagementApi.GetAdminNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdminNotifications`: CurrentAdminNotificationMessages
    fmt.Fprintf(os.Stdout, "Response from `SessionManagementApi.GetAdminNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CurrentAdminNotificationMessages**](Current admin notification messages.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilteredLogs

> string GetFilteredLogs(ctx).MaxLines(maxLines).UniverseName(universeName).QueryRegex(queryRegex).StartDate(startDate).EndDate(endDate).Execute()

getFilteredLogs

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
    maxLines := int32(56) // int32 |  (optional) (default to 10000)
    universeName := "universeName_example" // string |  (optional) (default to "null")
    queryRegex := "queryRegex_example" // string |  (optional) (default to "null")
    startDate := "startDate_example" // string |  (optional) (default to "null")
    endDate := "endDate_example" // string |  (optional) (default to "null")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SessionManagementApi.GetFilteredLogs(context.Background()).MaxLines(maxLines).UniverseName(universeName).QueryRegex(queryRegex).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionManagementApi.GetFilteredLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFilteredLogs`: string
    fmt.Fprintf(os.Stdout, "Response from `SessionManagementApi.GetFilteredLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFilteredLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxLines** | **int32** |  | [default to 10000]
 **universeName** | **string** |  | [default to &quot;null&quot;]
 **queryRegex** | **string** |  | [default to &quot;null&quot;]
 **startDate** | **string** |  | [default to &quot;null&quot;]
 **endDate** | **string** |  | [default to &quot;null&quot;]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogs

> LogData GetLogs(ctx, maxLines).Execute()

getLogs

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
    maxLines := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SessionManagementApi.GetLogs(context.Background(), maxLines).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionManagementApi.GetLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogs`: LogData
    fmt.Fprintf(os.Stdout, "Response from `SessionManagementApi.GetLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**maxLines** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogData**](LogData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionInfo

> SessionInfo GetSessionInfo(ctx).Execute()

Get current user/customer uuid auth/api token

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
    resp, r, err := api_client.SessionManagementApi.GetSessionInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionManagementApi.GetSessionInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionInfo`: SessionInfo
    fmt.Fprintf(os.Stdout, "Response from `SessionManagementApi.GetSessionInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionInfoRequest struct via the builder pattern


### Return type

[**SessionInfo**](SessionInfo.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterCustomer

> SessionInfo RegisterCustomer(ctx).CustomerRegisterFormData(customerRegisterFormData).GenerateApiToken(generateApiToken).Execute()

Register a customer



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
    customerRegisterFormData := *openapiclient.NewCustomerRegisterFormData("Code_example", "Email_example", "Name_example", "Password_example") // CustomerRegisterFormData | 
    generateApiToken := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SessionManagementApi.RegisterCustomer(context.Background()).CustomerRegisterFormData(customerRegisterFormData).GenerateApiToken(generateApiToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionManagementApi.RegisterCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterCustomer`: SessionInfo
    fmt.Fprintf(os.Stdout, "Response from `SessionManagementApi.RegisterCustomer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerRegisterFormData** | [**CustomerRegisterFormData**](CustomerRegisterFormData.md) |  | 
 **generateApiToken** | **bool** |  | [default to false]

### Return type

[**SessionInfo**](SessionInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

