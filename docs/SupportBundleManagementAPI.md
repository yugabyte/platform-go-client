# \SupportBundleManagementAPI

All URIs are relative to *http://localhost*

| Method                                                                                       | HTTP request                                                                           | Description                                        |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------- |
| [**CreateSupportBundle**](SupportBundleManagementAPI.md#CreateSupportBundle)                 | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/support_bundle                  | Create support bundle for specific universe        |
| [**DeleteSupportBundle**](SupportBundleManagementAPI.md#DeleteSupportBundle)                 | **Delete** /api/v1/customers/{cUUID}/universes/{uniUUID}/support_bundle/{sbUUID}       | Delete a support bundle                            |
| [**DownloadSupportBundle**](SupportBundleManagementAPI.md#DownloadSupportBundle)             | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/support_bundle/{sbUUID}/download | Download support bundle                            |
| [**EstimateSupportBundleSize**](SupportBundleManagementAPI.md#EstimateSupportBundleSize)     | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/support_bundle/estimate_size    | Estimate support bundle size for specific universe |
| [**GetSupportBundle**](SupportBundleManagementAPI.md#GetSupportBundle)                       | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/support_bundle/{sbUUID}          | Get a support bundle from a universe               |
| [**ListSupportBundle**](SupportBundleManagementAPI.md#ListSupportBundle)                     | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/support_bundle                   | List all support bundles from a universe           |
| [**ListSupportBundleComponents**](SupportBundleManagementAPI.md#ListSupportBundleComponents) | **Get** /api/v1/customers/{cUUID}/support_bundle/components                            | List all components available in support bundle    |



## CreateSupportBundle

> YBPTask CreateSupportBundle(ctx, cUUID, uniUUID).SupportBundle(supportBundle).Request(request).Execute()

Create support bundle for specific universe

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	supportBundle := *openapiclient.NewSupportBundleFormData([]string{"Components_example"}, time.Now(), time.Now()) // SupportBundleFormData | post support bundle info
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportBundleManagementAPI.CreateSupportBundle(context.Background(), cUUID, uniUUID).SupportBundle(supportBundle).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleManagementAPI.CreateSupportBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSupportBundle`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `SupportBundleManagementAPI.CreateSupportBundle`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSupportBundleRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **supportBundle** | [**SupportBundleFormData**](SupportBundleFormData.md) | post support bundle info | 
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


## DeleteSupportBundle

> YBPSuccess DeleteSupportBundle(ctx, cUUID, uniUUID, sbUUID).Request(request).Execute()

Delete a support bundle

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
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	sbUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportBundleManagementAPI.DeleteSupportBundle(context.Background(), cUUID, uniUUID, sbUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleManagementAPI.DeleteSupportBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSupportBundle`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `SupportBundleManagementAPI.DeleteSupportBundle`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |
| **sbUUID**  | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSupportBundleRequest struct via the builder pattern


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


## DownloadSupportBundle

> string DownloadSupportBundle(ctx, cUUID, uniUUID, sbUUID).Execute()

Download support bundle

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
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	sbUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportBundleManagementAPI.DownloadSupportBundle(context.Background(), cUUID, uniUUID, sbUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleManagementAPI.DownloadSupportBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadSupportBundle`: string
	fmt.Fprintf(os.Stdout, "Response from `SupportBundleManagementAPI.DownloadSupportBundle`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |
| **sbUUID**  | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadSupportBundleRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |




### Return type

**string**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-compressed

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EstimateSupportBundleSize

> SupportBundleSizeEstimateResponse EstimateSupportBundleSize(ctx, cUUID, uniUUID).SupportBundle(supportBundle).Request(request).Execute()

Estimate support bundle size for specific universe



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	supportBundle := *openapiclient.NewSupportBundleFormData([]string{"Components_example"}, time.Now(), time.Now()) // SupportBundleFormData | support bundle info
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportBundleManagementAPI.EstimateSupportBundleSize(context.Background(), cUUID, uniUUID).SupportBundle(supportBundle).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleManagementAPI.EstimateSupportBundleSize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EstimateSupportBundleSize`: SupportBundleSizeEstimateResponse
	fmt.Fprintf(os.Stdout, "Response from `SupportBundleManagementAPI.EstimateSupportBundleSize`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiEstimateSupportBundleSizeRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **supportBundle** | [**SupportBundleFormData**](SupportBundleFormData.md) | support bundle info | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**SupportBundleSizeEstimateResponse**](SupportBundleSizeEstimateResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportBundle

> SupportBundle GetSupportBundle(ctx, cUUID, uniUUID, sbUUID).Execute()

Get a support bundle from a universe

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
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	sbUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportBundleManagementAPI.GetSupportBundle(context.Background(), cUUID, uniUUID, sbUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleManagementAPI.GetSupportBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportBundle`: SupportBundle
	fmt.Fprintf(os.Stdout, "Response from `SupportBundleManagementAPI.GetSupportBundle`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |
| **sbUUID**  | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportBundleRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |




### Return type

[**SupportBundle**](SupportBundle.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportBundle

> []SupportBundle ListSupportBundle(ctx, cUUID, uniUUID).Execute()

List all support bundles from a universe

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
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportBundleManagementAPI.ListSupportBundle(context.Background(), cUUID, uniUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleManagementAPI.ListSupportBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSupportBundle`: []SupportBundle
	fmt.Fprintf(os.Stdout, "Response from `SupportBundleManagementAPI.ListSupportBundle`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiListSupportBundleRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

[**[]SupportBundle**](SupportBundle.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportBundleComponents

> []string ListSupportBundleComponents(ctx, cUUID).Execute()

List all components available in support bundle

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportBundleManagementAPI.ListSupportBundleComponents(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportBundleManagementAPI.ListSupportBundleComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSupportBundleComponents`: []string
	fmt.Fprintf(os.Stdout, "Response from `SupportBundleManagementAPI.ListSupportBundleComponents`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiListSupportBundleComponentsRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


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

