# \ReleaseManagementAPI

All URIs are relative to *http://localhost*

| Method                                                                         | HTTP request                                                 | Description                                                                               |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------ | ----------------------------------------------------------------------------------------- |
| [**CreateRelease**](ReleaseManagementAPI.md#CreateRelease)                     | **Post** /api/v1/customers/{cUUID}/releases                  | Deprecated: sinceVersion 2024.1. Use ReleasesController.create instead. Create a release  |
| [**DeleteRelease**](ReleaseManagementAPI.md#DeleteRelease)                     | **Delete** /api/v1/customers/{cUUID}/releases/{name}         | Deprecated: sinceVersion: 2024.1. Use ReleasesController.delete instead. Delete a release |
| [**GetListOfRegionReleases**](ReleaseManagementAPI.md#GetListOfRegionReleases) | **Get** /api/v1/customers/{cUUID}/providers/{pUUID}/releases | List releases by provider - deprecated                                                    |
| [**GetListOfReleases**](ReleaseManagementAPI.md#GetListOfReleases)             | **Get** /api/v1/customers/{cUUID}/releases                   | Deprecated: sinceVersion: 2024.1. Use ReleasesController.list instead. List all releases  |
| [**Refresh**](ReleaseManagementAPI.md#Refresh)                                 | **Put** /api/v1/customers/{cUUID}/releases                   | Refresh a release                                                                         |
| [**UpdateRelease**](ReleaseManagementAPI.md#UpdateRelease)                     | **Put** /api/v1/customers/{cUUID}/releases/{name}            | Deprecated: sinceVersion: 2024.1. Use ReleasesController.update instead. Update a release |



## CreateRelease

> YBPSuccess CreateRelease(ctx, cUUID).Release(release).Request(request).Execute()

Deprecated: sinceVersion 2024.1. Use ReleasesController.create instead. Create a release

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
	release := *openapiclient.NewReleaseFormData(*openapiclient.NewGCSLocation(), *openapiclient.NewHttpLocation(), *openapiclient.NewS3Location(), "Version_example") // ReleaseFormData | Release data for remote downloading to be created
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseManagementAPI.CreateRelease(context.Background(), cUUID).Release(release).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementAPI.CreateRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRelease`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementAPI.CreateRelease`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReleaseRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **release** | [**ReleaseFormData**](ReleaseFormData.md) | Release data for remote downloading to be created | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPSuccess**](YBPSuccess.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRelease

> ReleaseMetadata DeleteRelease(ctx, cUUID, name).Request(request).Execute()

Deprecated: sinceVersion: 2024.1. Use ReleasesController.delete instead. Delete a release

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
	name := "name_example" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseManagementAPI.DeleteRelease(context.Background(), cUUID, name).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementAPI.DeleteRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRelease`: ReleaseMetadata
	fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementAPI.DeleteRelease`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **name**  | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReleaseRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**ReleaseMetadata**](ReleaseMetadata.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListOfRegionReleases

> map[string]map[string]interface{} GetListOfRegionReleases(ctx, cUUID, pUUID).IncludeMetadata(includeMetadata).Execute()

List releases by provider - deprecated



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
	pUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	includeMetadata := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseManagementAPI.GetListOfRegionReleases(context.Background(), cUUID, pUUID).IncludeMetadata(includeMetadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementAPI.GetListOfRegionReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListOfRegionReleases`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementAPI.GetListOfRegionReleases`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **pUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetListOfRegionReleasesRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **includeMetadata** | **bool** |  | [default to false]

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


## GetListOfReleases

> map[string]map[string]interface{} GetListOfReleases(ctx, cUUID).IncludeMetadata(includeMetadata).Arch(arch).Execute()

Deprecated: sinceVersion: 2024.1. Use ReleasesController.list instead. List all releases

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
	includeMetadata := true // bool |  (optional) (default to false)
	arch := "arch_example" // string |  (optional) (default to "null")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseManagementAPI.GetListOfReleases(context.Background(), cUUID).IncludeMetadata(includeMetadata).Arch(arch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementAPI.GetListOfReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListOfReleases`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementAPI.GetListOfReleases`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetListOfReleasesRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **includeMetadata** | **bool** |  | [default to false]
 **arch** | **string** |  | [default to &quot;null&quot;]

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


## Refresh

> YBPSuccess Refresh(ctx, cUUID).Request(request).Execute()

Refresh a release

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
	resp, r, err := apiClient.ReleaseManagementAPI.Refresh(context.Background(), cUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementAPI.Refresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Refresh`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementAPI.Refresh`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshRequest struct via the builder pattern


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


## UpdateRelease

> ReleaseMetadata UpdateRelease(ctx, cUUID, name).Release(release).Request(request).Execute()

Deprecated: sinceVersion: 2024.1. Use ReleasesController.update instead. Update a release

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
	name := "name_example" // string | 
	release := map[string]interface{}{ ... } // map[string]interface{} | Release data to be updated
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseManagementAPI.UpdateRelease(context.Background(), cUUID, name).Release(release).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementAPI.UpdateRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRelease`: ReleaseMetadata
	fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementAPI.UpdateRelease`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **name**  | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReleaseRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **release** | **map[string]interface{}** | Release data to be updated | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**ReleaseMetadata**](ReleaseMetadata.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

