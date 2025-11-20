# \RuntimeConfigurationAPI

All URIs are relative to *http://localhost*

| Method                                                                    | HTTP request                                                          | Description                                                     |
| ------------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------- |
| [**DeleteKey**](RuntimeConfigurationAPI.md#DeleteKey)                     | **Delete** /api/v1/customers/{cUUID}/runtime_config/{scope}/key/{key} | Delete a configuration key                                      |
| [**GetConfig**](RuntimeConfigurationAPI.md#GetConfig)                     | **Get** /api/v1/customers/{cUUID}/runtime_config/{scope}              | List configuration entries for a scope                          |
| [**GetConfigurationKey**](RuntimeConfigurationAPI.md#GetConfigurationKey) | **Get** /api/v1/customers/{cUUID}/runtime_config/{scope}/key/{key}    | Get a configuration key                                         |
| [**ListFeatureFlags**](RuntimeConfigurationAPI.md#ListFeatureFlags)       | **Get** /api/v1/runtime_config/feature_flags                          | List all the feature flag runtime config keys and their values. |
| [**ListKeyInfo**](RuntimeConfigurationAPI.md#ListKeyInfo)                 | **Get** /api/v1/runtime_config/mutable_key_info                       | List mutable keys                                               |
| [**ListKeys**](RuntimeConfigurationAPI.md#ListKeys)                       | **Get** /api/v1/runtime_config/mutable_keys                           | List mutable keys                                               |
| [**ListScopes**](RuntimeConfigurationAPI.md#ListScopes)                   | **Get** /api/v1/customers/{cUUID}/runtime_config/scopes               | List configuration scopes                                       |
| [**SetKey**](RuntimeConfigurationAPI.md#SetKey)                           | **Put** /api/v1/customers/{cUUID}/runtime_config/{scope}/key/{key}    | Update a configuration key                                      |



## DeleteKey

> YBPSuccess DeleteKey(ctx, cUUID, scope, key).Request(request).Execute()

Delete a configuration key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	scope := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	key := "key_example" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RuntimeConfigurationAPI.DeleteKey(context.Background(), cUUID, scope, key).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RuntimeConfigurationAPI.DeleteKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteKey`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `RuntimeConfigurationAPI.DeleteKey`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **scope** | **string**          |                                                                             |
| **key**   | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



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


## GetConfig

> ScopedConfig GetConfig(ctx, cUUID, scope).IncludeInherited(includeInherited).Request(request).Execute()

List configuration entries for a scope



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	scope := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	includeInherited := true // bool |  (optional) (default to false)
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RuntimeConfigurationAPI.GetConfig(context.Background(), cUUID, scope).IncludeInherited(includeInherited).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RuntimeConfigurationAPI.GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfig`: ScopedConfig
	fmt.Fprintf(os.Stdout, "Response from `RuntimeConfigurationAPI.GetConfig`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **scope** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **includeInherited** | **bool** |  | [default to false]
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**ScopedConfig**](ScopedConfig.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurationKey

> string GetConfigurationKey(ctx, cUUID, scope, key).Request(request).Execute()

Get a configuration key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	scope := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	key := "key_example" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RuntimeConfigurationAPI.GetConfigurationKey(context.Background(), cUUID, scope, key).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RuntimeConfigurationAPI.GetConfigurationKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigurationKey`: string
	fmt.Fprintf(os.Stdout, "Response from `RuntimeConfigurationAPI.GetConfigurationKey`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **scope** | **string**          |                                                                             |
| **key**   | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationKeyRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

**string**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFeatureFlags

> []ConfigEntry ListFeatureFlags(ctx).Execute()

List all the feature flag runtime config keys and their values.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RuntimeConfigurationAPI.ListFeatureFlags(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RuntimeConfigurationAPI.ListFeatureFlags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFeatureFlags`: []ConfigEntry
	fmt.Fprintf(os.Stdout, "Response from `RuntimeConfigurationAPI.ListFeatureFlags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFeatureFlagsRequest struct via the builder pattern


### Return type

[**[]ConfigEntry**](ConfigEntry.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKeyInfo

> []ConfKeyInfo ListKeyInfo(ctx).Execute()

List mutable keys



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RuntimeConfigurationAPI.ListKeyInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RuntimeConfigurationAPI.ListKeyInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKeyInfo`: []ConfKeyInfo
	fmt.Fprintf(os.Stdout, "Response from `RuntimeConfigurationAPI.ListKeyInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListKeyInfoRequest struct via the builder pattern


### Return type

[**[]ConfKeyInfo**](ConfKeyInfo.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKeys

> []string ListKeys(ctx).Execute()

List mutable keys



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RuntimeConfigurationAPI.ListKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RuntimeConfigurationAPI.ListKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKeys`: []string
	fmt.Fprintf(os.Stdout, "Response from `RuntimeConfigurationAPI.ListKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListKeysRequest struct via the builder pattern


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


## ListScopes

> RuntimeConfigData ListScopes(ctx, cUUID).Request(request).Execute()

List configuration scopes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RuntimeConfigurationAPI.ListScopes(context.Background(), cUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RuntimeConfigurationAPI.ListScopes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListScopes`: RuntimeConfigData
	fmt.Fprintf(os.Stdout, "Response from `RuntimeConfigurationAPI.ListScopes`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiListScopesRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**RuntimeConfigData**](RuntimeConfigData.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetKey

> YBPSuccess SetKey(ctx, cUUID, scope, key).NewValue(newValue).Request(request).Execute()

Update a configuration key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	scope := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	key := "key_example" // string | 
	newValue := "newValue_example" // string | New value for config key
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RuntimeConfigurationAPI.SetKey(context.Background(), cUUID, scope, key).NewValue(newValue).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RuntimeConfigurationAPI.SetKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetKey`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `RuntimeConfigurationAPI.SetKey`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **scope** | **string**          |                                                                             |
| **key**   | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiSetKeyRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



 **newValue** | **string** | New value for config key | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPSuccess**](YBPSuccess.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

