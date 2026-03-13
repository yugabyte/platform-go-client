# \SessionManagementAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiLogin**](SessionManagementAPI.md#ApiLogin) | **Post** /api/v1/api_login | Authenticate user using email and password
[**ApiToken**](SessionManagementAPI.md#ApiToken) | **Put** /api/v1/customers/{cUUID}/api_token | Regenerate and fetch API token
[**AppVersion**](SessionManagementAPI.md#AppVersion) | **Get** /api/v1/app_version | appVersion
[**CustomerCount**](SessionManagementAPI.md#CustomerCount) | **Get** /api/v1/customer_count | customerCount
[**GetAdminNotifications**](SessionManagementAPI.md#GetAdminNotifications) | **Get** /api/v1/customers/{cUUID}/admin_notifications | Current list of notifications for admin
[**GetFilteredLogs**](SessionManagementAPI.md#GetFilteredLogs) | **Get** /api/v1/logs | getFilteredLogs
[**GetLogs**](SessionManagementAPI.md#GetLogs) | **Get** /api/v1/logs/{maxLines} | getLogs
[**GetSessionInfo**](SessionManagementAPI.md#GetSessionInfo) | **Get** /api/v1/session_info | Get current user and customer uuid. This will not generate or return the API token, use /api_token API for that.
[**RegisterCustomer**](SessionManagementAPI.md#RegisterCustomer) | **Post** /api/v1/register | Register a customer



## ApiLogin

> SessionInfo ApiLogin(ctx).CustomerLoginFormData(customerLoginFormData).Request(request).Execute()

Authenticate user using email and password

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
	customerLoginFormData := *openapiclient.NewCustomerLoginFormData("Email_example", "Password_example") // CustomerLoginFormData | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionManagementAPI.ApiLogin(context.Background()).CustomerLoginFormData(customerLoginFormData).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionManagementAPI.ApiLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiLogin`: SessionInfo
	fmt.Fprintf(os.Stdout, "Response from `SessionManagementAPI.ApiLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerLoginFormData** | [**CustomerLoginFormData**](CustomerLoginFormData.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

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


## ApiToken

> SessionInfo ApiToken(ctx, cUUID).ApiTokenVersion(apiTokenVersion).Request(request).Execute()

Regenerate and fetch API token

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
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiTokenVersion := int64(789) // int64 |  (optional) (default to -1)
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionManagementAPI.ApiToken(context.Background(), cUUID).ApiTokenVersion(apiTokenVersion).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionManagementAPI.ApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiToken`: SessionInfo
	fmt.Fprintf(os.Stdout, "Response from `SessionManagementAPI.ApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiTokenVersion** | **int64** |  | [default to -1]
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**SessionInfo**](SessionInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionManagementAPI.AppVersion(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionManagementAPI.AppVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppVersion`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `SessionManagementAPI.AppVersion`: %v\n", resp)
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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionManagementAPI.CustomerCount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionManagementAPI.CustomerCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerCount`: CustomerCountResp
	fmt.Fprintf(os.Stdout, "Response from `SessionManagementAPI.CustomerCount`: %v\n", resp)
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

Current list of notifications for admin



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
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionManagementAPI.GetAdminNotifications(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionManagementAPI.GetAdminNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminNotifications`: CurrentAdminNotificationMessages
	fmt.Fprintf(os.Stdout, "Response from `SessionManagementAPI.GetAdminNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CurrentAdminNotificationMessages**](CurrentAdminNotificationMessages.md)

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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	maxLines := int32(56) // int32 |  (optional) (default to 10000)
	universeName := "universeName_example" // string |  (optional) (default to "null")
	queryRegex := "queryRegex_example" // string |  (optional) (default to "null")
	startDate := "startDate_example" // string |  (optional) (default to "null")
	endDate := "endDate_example" // string |  (optional) (default to "null")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionManagementAPI.GetFilteredLogs(context.Background()).MaxLines(maxLines).UniverseName(universeName).QueryRegex(queryRegex).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionManagementAPI.GetFilteredLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilteredLogs`: string
	fmt.Fprintf(os.Stdout, "Response from `SessionManagementAPI.GetFilteredLogs`: %v\n", resp)
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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	maxLines := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionManagementAPI.GetLogs(context.Background(), maxLines).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionManagementAPI.GetLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogs`: LogData
	fmt.Fprintf(os.Stdout, "Response from `SessionManagementAPI.GetLogs`: %v\n", resp)
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

> SessionInfo GetSessionInfo(ctx).Request(request).Execute()

Get current user and customer uuid. This will not generate or return the API token, use /api_token API for that.



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
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionManagementAPI.GetSessionInfo(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionManagementAPI.GetSessionInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSessionInfo`: SessionInfo
	fmt.Fprintf(os.Stdout, "Response from `SessionManagementAPI.GetSessionInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**interface{}**](interface{}.md) |  | 

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

> SessionInfo RegisterCustomer(ctx).CustomerRegisterFormData(customerRegisterFormData).GenerateApiToken(generateApiToken).Request(request).Execute()

Register a customer



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
	customerRegisterFormData := *openapiclient.NewCustomerRegisterFormData("Code_example", "Email_example", "Name_example", "Password_example") // CustomerRegisterFormData | 
	generateApiToken := true // bool |  (optional) (default to false)
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionManagementAPI.RegisterCustomer(context.Background()).CustomerRegisterFormData(customerRegisterFormData).GenerateApiToken(generateApiToken).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionManagementAPI.RegisterCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterCustomer`: SessionInfo
	fmt.Fprintf(os.Stdout, "Response from `SessionManagementAPI.RegisterCustomer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerRegisterFormData** | [**CustomerRegisterFormData**](CustomerRegisterFormData.md) |  | 
 **generateApiToken** | **bool** |  | [default to false]
 **request** | [**interface{}**](interface{}.md) |  | 

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

