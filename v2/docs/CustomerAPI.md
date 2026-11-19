# \CustomerAPI

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateApiToken**](CustomerAPI.md#GenerateApiToken) | **Put** /customers/{cUUID}/api-token | Generate an API token
[**GetHostCloudInfo**](CustomerAPI.md#GetHostCloudInfo) | **Get** /customers/{cUUID}/host-cloud-info | Get the YBA host&#39;s cloud info
[**RegisterCustomer**](CustomerAPI.md#RegisterCustomer) | **Post** /customers | Register a customer



## GenerateApiToken

> UserApiToken GenerateApiToken(ctx, cUUID).ApiTokenSpec(apiTokenSpec).Execute()

Generate an API token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	apiTokenSpec := *openapiclient.NewApiTokenSpec() // ApiTokenSpec |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.GenerateApiToken(context.Background(), cUUID).ApiTokenSpec(apiTokenSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.GenerateApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateApiToken`: UserApiToken
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.GenerateApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiTokenSpec** | [**ApiTokenSpec**](ApiTokenSpec.md) |  | 

### Return type

[**UserApiToken**](UserApiToken.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [authTokenAuth](../README.md#authTokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostCloudInfo

> HostCloudInfo GetHostCloudInfo(ctx, cUUID).Execute()

Get the YBA host's cloud info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.GetHostCloudInfo(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.GetHostCloudInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHostCloudInfo`: HostCloudInfo
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.GetHostCloudInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostCloudInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostCloudInfo**](HostCloudInfo.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [authTokenAuth](../README.md#authTokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterCustomer

> CustomerRegistration RegisterCustomer(ctx).CustomerRegistrationSpec(customerRegistrationSpec).Execute()

Register a customer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	customerRegistrationSpec := *openapiclient.NewCustomerRegistrationSpec("admin", "Example Customer", "admin@example.com", "Password_example") // CustomerRegistrationSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.RegisterCustomer(context.Background()).CustomerRegistrationSpec(customerRegistrationSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.RegisterCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterCustomer`: CustomerRegistration
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.RegisterCustomer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerRegistrationSpec** | [**CustomerRegistrationSpec**](CustomerRegistrationSpec.md) |  | 

### Return type

[**CustomerRegistration**](CustomerRegistration.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

