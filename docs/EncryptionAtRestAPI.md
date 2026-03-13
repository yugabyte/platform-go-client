# \EncryptionAtRestAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKMSConfig**](EncryptionAtRestAPI.md#CreateKMSConfig) | **Post** /api/v1/customers/{cUUID}/kms_configs/{kmsProvider} | Create a KMS configuration
[**DeleteKMSConfig**](EncryptionAtRestAPI.md#DeleteKMSConfig) | **Delete** /api/v1/customers/{cUUID}/kms_configs/{configUUID} | Delete a KMS configuration
[**EditKMSConfig**](EncryptionAtRestAPI.md#EditKMSConfig) | **Post** /api/v1/customers/{cUUID}/kms_configs/{configUUID}/edit | Edit a KMS configuration
[**GetCurrentKeyRef**](EncryptionAtRestAPI.md#GetCurrentKeyRef) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/kms/key_ref | Get a universe&#39;s key reference
[**GetKMSConfig**](EncryptionAtRestAPI.md#GetKMSConfig) | **Get** /api/v1/customers/{cUUID}/kms_configs/{configUUID} | Get details of a KMS configuration
[**GetKeyRefHistory**](EncryptionAtRestAPI.md#GetKeyRefHistory) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/kms | Get a universe&#39;s key reference history
[**ListKMSConfigs**](EncryptionAtRestAPI.md#ListKMSConfigs) | **Get** /api/v1/customers/{cUUID}/kms_configs | List KMS configurations
[**RefreshKMSConfig**](EncryptionAtRestAPI.md#RefreshKMSConfig) | **Put** /api/v1/customers/{cUUID}/kms_configs/{configUUID}/refresh | Refresh KMS Config
[**RemoveKeyRefHistory**](EncryptionAtRestAPI.md#RemoveKeyRefHistory) | **Delete** /api/v1/customers/{cUUID}/universes/{uniUUID}/kms | This API removes a universe&#39;s key reference history - deprecated
[**RetrieveKey**](EncryptionAtRestAPI.md#RetrieveKey) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/kms | Retrive a universe&#39;s KMS key



## CreateKMSConfig

> YBPTask CreateKMSConfig(ctx, cUUID, kmsProvider).KMSConfig(kMSConfig).Request(request).Execute()

Create a KMS configuration

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
	kmsProvider := "kmsProvider_example" // string | 
	kMSConfig := map[string]interface{}{ ... } // map[string]interface{} | KMS config to be created
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EncryptionAtRestAPI.CreateKMSConfig(context.Background(), cUUID, kmsProvider).KMSConfig(kMSConfig).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestAPI.CreateKMSConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKMSConfig`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestAPI.CreateKMSConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**kmsProvider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKMSConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kMSConfig** | **map[string]interface{}** | KMS config to be created | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKMSConfig

> YBPTask DeleteKMSConfig(ctx, cUUID, configUUID).Request(request).Execute()

Delete a KMS configuration

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
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EncryptionAtRestAPI.DeleteKMSConfig(context.Background(), cUUID, configUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestAPI.DeleteKMSConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteKMSConfig`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestAPI.DeleteKMSConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**configUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKMSConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## EditKMSConfig

> YBPTask EditKMSConfig(ctx, cUUID, configUUID).KMSConfig(kMSConfig).Request(request).Execute()

Edit a KMS configuration

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
	kMSConfig := map[string]interface{}{ ... } // map[string]interface{} | KMS config to be edited
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EncryptionAtRestAPI.EditKMSConfig(context.Background(), cUUID, configUUID).KMSConfig(kMSConfig).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestAPI.EditKMSConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditKMSConfig`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestAPI.EditKMSConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**configUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditKMSConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kMSConfig** | **map[string]interface{}** | KMS config to be edited | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentKeyRef

> map[string]map[string]interface{} GetCurrentKeyRef(ctx, cUUID, uniUUID).Execute()

Get a universe's key reference



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
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EncryptionAtRestAPI.GetCurrentKeyRef(context.Background(), cUUID, uniUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestAPI.GetCurrentKeyRef``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentKeyRef`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestAPI.GetCurrentKeyRef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentKeyRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKMSConfig

> map[string]map[string]interface{} GetKMSConfig(ctx, cUUID, configUUID).Execute()

Get details of a KMS configuration

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
	resp, r, err := apiClient.EncryptionAtRestAPI.GetKMSConfig(context.Background(), cUUID, configUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestAPI.GetKMSConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKMSConfig`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestAPI.GetKMSConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**configUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKMSConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeyRefHistory

> []map[string]interface{} GetKeyRefHistory(ctx, cUUID, uniUUID).Execute()

Get a universe's key reference history



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
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EncryptionAtRestAPI.GetKeyRefHistory(context.Background(), cUUID, uniUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestAPI.GetKeyRefHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKeyRefHistory`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestAPI.GetKeyRefHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyRefHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**[]map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKMSConfigs

> []map[string]interface{} ListKMSConfigs(ctx, cUUID).Execute()

List KMS configurations

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
	resp, r, err := apiClient.EncryptionAtRestAPI.ListKMSConfigs(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestAPI.ListKMSConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKMSConfigs`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestAPI.ListKMSConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKMSConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshKMSConfig

> YBPSuccess RefreshKMSConfig(ctx, cUUID, configUUID).Request(request).Execute()

Refresh KMS Config



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
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EncryptionAtRestAPI.RefreshKMSConfig(context.Background(), cUUID, configUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestAPI.RefreshKMSConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshKMSConfig`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestAPI.RefreshKMSConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**configUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshKMSConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPSuccess**](YBPSuccess.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveKeyRefHistory

> YBPSuccess RemoveKeyRefHistory(ctx, cUUID, uniUUID).Request(request).Execute()

This API removes a universe's key reference history - deprecated



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
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EncryptionAtRestAPI.RemoveKeyRefHistory(context.Background(), cUUID, uniUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestAPI.RemoveKeyRefHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveKeyRefHistory`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestAPI.RemoveKeyRefHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveKeyRefHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPSuccess**](YBPSuccess.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveKey

> map[string]map[string]interface{} RetrieveKey(ctx, cUUID, uniUUID).Request(request).Execute()

Retrive a universe's KMS key



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
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EncryptionAtRestAPI.RetrieveKey(context.Background(), cUUID, uniUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestAPI.RetrieveKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveKey`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestAPI.RetrieveKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

**map[string]map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

