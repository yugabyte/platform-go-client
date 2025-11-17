# \CustomerConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerConfig**](CustomerConfigurationAPI.md#CreateCustomerConfig) | **Post** /api/v1/customers/{cUUID}/configs | Create a customer configuration
[**DeleteCustomerConfig**](CustomerConfigurationAPI.md#DeleteCustomerConfig) | **Delete** /api/v1/customers/{cUUID}/configs/{configUUID} | Delete a customer configuration
[**DeleteCustomerConfigV2**](CustomerConfigurationAPI.md#DeleteCustomerConfigV2) | **Delete** /api/v1/customers/{cUUID}/configs/{configUUID}/delete | Delete a customer configuration V2
[**EditCustomerConfig**](CustomerConfigurationAPI.md#EditCustomerConfig) | **Put** /api/v1/customers/{cUUID}/configs/{configUUID} | Update a customer configuration
[**EditCustomerConfig_0**](CustomerConfigurationAPI.md#EditCustomerConfig_0) | **Put** /api/v1/customers/{cUUID}/configs/{configUUID}/edit | Update a customer configuration V2
[**GetListOfCustomerConfig**](CustomerConfigurationAPI.md#GetListOfCustomerConfig) | **Get** /api/v1/customers/{cUUID}/configs | List all customer configurations
[**GetListOfYbaBackupDirsCustomerConfig**](CustomerConfigurationAPI.md#GetListOfYbaBackupDirsCustomerConfig) | **Get** /api/v1/customers/{cUUID}/configs/{configUUID}/backup_dirs | List all backup dirs within a customer configurations
[**ListBuckets**](CustomerConfigurationAPI.md#ListBuckets) | **Post** /api/v1/customers/{cUUID}/cloud/{cloud}/buckets | List buckets with provided credentials



## CreateCustomerConfig

> CustomerConfig CreateCustomerConfig(ctx, cUUID).Config(config).Request(request).Execute()

Create a customer configuration

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
	config := *openapiclient.NewCustomerConfig("backup20-01-2021", "CustomerUUID_example", map[string]interface{}("{\"AWS_ACCESS_KEY_ID\": \"AK****************ZD\"}"), false, "S3", "STORAGE") // CustomerConfig | Configuration data to be created
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerConfigurationAPI.CreateCustomerConfig(context.Background(), cUUID).Config(config).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerConfigurationAPI.CreateCustomerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomerConfig`: CustomerConfig
	fmt.Fprintf(os.Stdout, "Response from `CustomerConfigurationAPI.CreateCustomerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **config** | [**CustomerConfig**](CustomerConfig.md) | Configuration data to be created | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**CustomerConfig**](CustomerConfig.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomerConfig

> YBPTask DeleteCustomerConfig(ctx, cUUID, configUUID).IsDeleteBackups(isDeleteBackups).Request(request).Execute()

Delete a customer configuration



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
	configUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	isDeleteBackups := true // bool |  (optional) (default to false)
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerConfigurationAPI.DeleteCustomerConfig(context.Background(), cUUID, configUUID).IsDeleteBackups(isDeleteBackups).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerConfigurationAPI.DeleteCustomerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCustomerConfig`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `CustomerConfigurationAPI.DeleteCustomerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**configUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isDeleteBackups** | **bool** |  | [default to false]
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomerConfigV2

> YBPTask DeleteCustomerConfigV2(ctx, cUUID, configUUID).IsDeleteBackups(isDeleteBackups).Request(request).Execute()

Delete a customer configuration V2

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
	configUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	isDeleteBackups := true // bool |  (optional) (default to false)
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerConfigurationAPI.DeleteCustomerConfigV2(context.Background(), cUUID, configUUID).IsDeleteBackups(isDeleteBackups).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerConfigurationAPI.DeleteCustomerConfigV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCustomerConfigV2`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `CustomerConfigurationAPI.DeleteCustomerConfigV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**configUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerConfigV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isDeleteBackups** | **bool** |  | [default to false]
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditCustomerConfig

> CustomerConfig EditCustomerConfig(ctx, cUUID, configUUID).Config(config).Request(request).Execute()

Update a customer configuration

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
	configUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	config := *openapiclient.NewCustomerConfig("backup20-01-2021", "CustomerUUID_example", map[string]interface{}("{\"AWS_ACCESS_KEY_ID\": \"AK****************ZD\"}"), false, "S3", "STORAGE") // CustomerConfig | Configuration data to be updated
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerConfigurationAPI.EditCustomerConfig(context.Background(), cUUID, configUUID).Config(config).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerConfigurationAPI.EditCustomerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditCustomerConfig`: CustomerConfig
	fmt.Fprintf(os.Stdout, "Response from `CustomerConfigurationAPI.EditCustomerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**configUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditCustomerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **config** | [**CustomerConfig**](CustomerConfig.md) | Configuration data to be updated | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**CustomerConfig**](CustomerConfig.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditCustomerConfig_0

> CustomerConfig EditCustomerConfig_0(ctx, cUUID, configUUID).Config(config).Request(request).Execute()

Update a customer configuration V2

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
	configUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	config := *openapiclient.NewCustomerConfig("backup20-01-2021", "CustomerUUID_example", map[string]interface{}("{\"AWS_ACCESS_KEY_ID\": \"AK****************ZD\"}"), false, "S3", "STORAGE") // CustomerConfig | Configuration data to be updated
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerConfigurationAPI.EditCustomerConfig_0(context.Background(), cUUID, configUUID).Config(config).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerConfigurationAPI.EditCustomerConfig_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditCustomerConfig_0`: CustomerConfig
	fmt.Fprintf(os.Stdout, "Response from `CustomerConfigurationAPI.EditCustomerConfig_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**configUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditCustomerConfig_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **config** | [**CustomerConfig**](CustomerConfig.md) | Configuration data to be updated | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**CustomerConfig**](CustomerConfig.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListOfCustomerConfig

> []CustomerConfigUI GetListOfCustomerConfig(ctx, cUUID).Execute()

List all customer configurations

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
	resp, r, err := apiClient.CustomerConfigurationAPI.GetListOfCustomerConfig(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerConfigurationAPI.GetListOfCustomerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListOfCustomerConfig`: []CustomerConfigUI
	fmt.Fprintf(os.Stdout, "Response from `CustomerConfigurationAPI.GetListOfCustomerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListOfCustomerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CustomerConfigUI**](CustomerConfigUI.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListOfYbaBackupDirsCustomerConfig

> []string GetListOfYbaBackupDirsCustomerConfig(ctx, cUUID, configUUID).Execute()

List all backup dirs within a customer configurations

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
	configUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerConfigurationAPI.GetListOfYbaBackupDirsCustomerConfig(context.Background(), cUUID, configUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerConfigurationAPI.GetListOfYbaBackupDirsCustomerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListOfYbaBackupDirsCustomerConfig`: []string
	fmt.Fprintf(os.Stdout, "Response from `CustomerConfigurationAPI.GetListOfYbaBackupDirsCustomerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**configUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListOfYbaBackupDirsCustomerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**[]string**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuckets

> map[string]interface{} ListBuckets(ctx, cUUID, cloud).Credentials(credentials).Request(request).Execute()

List buckets with provided credentials

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
	cloud := "cloud_example" // string | 
	credentials := map[string]interface{}{ ... } // map[string]interface{} | Credentials to list buckets
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerConfigurationAPI.ListBuckets(context.Background(), cUUID, cloud).Credentials(credentials).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerConfigurationAPI.ListBuckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBuckets`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomerConfigurationAPI.ListBuckets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**cloud** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **credentials** | **map[string]interface{}** | Credentials to list buckets | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

