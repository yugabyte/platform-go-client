# \NewReleaseManagementAPI

All URIs are relative to *http://localhost*

| Method                                                              | HTTP request                                              | Description      |
| ------------------------------------------------------------------- | --------------------------------------------------------- | ---------------- |
| [**CreateNewRelease**](NewReleaseManagementAPI.md#CreateNewRelease) | **Post** /api/v1/customers/{cUUID}/ybdb_release           | Create a release |
| [**DeleteNewRelease**](NewReleaseManagementAPI.md#DeleteNewRelease) | **Delete** /api/v1/customers/{cUUID}/ybdb_release/{rUUID} | delete a release |
| [**GetNewRelease**](NewReleaseManagementAPI.md#GetNewRelease)       | **Get** /api/v1/customers/{cUUID}/ybdb_release/{rUUID}    | Get a release    |
| [**ListNewReleases**](NewReleaseManagementAPI.md#ListNewReleases)   | **Get** /api/v1/customers/{cUUID}/ybdb_release            | List releases    |
| [**UpdateNewRelease**](NewReleaseManagementAPI.md#UpdateNewRelease) | **Put** /api/v1/customers/{cUUID}/ybdb_release/{rUUID}    | Update a release |



## CreateNewRelease

> YBPCreateSuccess CreateNewRelease(ctx, cUUID).Release(release).Request(request).Execute()

Create a release



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
	release := *openapiclient.NewCreateRelease([]openapiclient.Artifact{*openapiclient.NewArtifact("Architecture_example", "PackageFileId_example", "PackageUrl_example", "Platform_example", "Sha256_example")}, int64(123), "ReleaseNotes_example", "ReleaseTag_example", "ReleaseType_example", "ReleaseUuid_example", "Version_example", "YbType_example") // CreateRelease | Release data to be created
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewReleaseManagementAPI.CreateNewRelease(context.Background(), cUUID).Release(release).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewReleaseManagementAPI.CreateNewRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewRelease`: YBPCreateSuccess
	fmt.Fprintf(os.Stdout, "Response from `NewReleaseManagementAPI.CreateNewRelease`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewReleaseRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **release** | [**CreateRelease**](CreateRelease.md) | Release data to be created | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPCreateSuccess**](YBPCreateSuccess.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNewRelease

> YBPSuccess DeleteNewRelease(ctx, cUUID, rUUID).Request(request).Execute()

delete a release



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
	rUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewReleaseManagementAPI.DeleteNewRelease(context.Background(), cUUID, rUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewReleaseManagementAPI.DeleteNewRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNewRelease`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `NewReleaseManagementAPI.DeleteNewRelease`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **rUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNewReleaseRequest struct via the builder pattern


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


## GetNewRelease

> ResponseRelease GetNewRelease(ctx, cUUID, rUUID).Request(request).Execute()

Get a release



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
	rUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewReleaseManagementAPI.GetNewRelease(context.Background(), cUUID, rUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewReleaseManagementAPI.GetNewRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNewRelease`: ResponseRelease
	fmt.Fprintf(os.Stdout, "Response from `NewReleaseManagementAPI.GetNewRelease`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **rUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetNewReleaseRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**ResponseRelease**](ResponseRelease.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNewReleases

> []ResponseRelease ListNewReleases(ctx, cUUID).DeploymentType(deploymentType).Request(request).Execute()

List releases



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
	deploymentType := "deploymentType_example" // string |  (optional) (default to "null")
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewReleaseManagementAPI.ListNewReleases(context.Background(), cUUID).DeploymentType(deploymentType).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewReleaseManagementAPI.ListNewReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNewReleases`: []ResponseRelease
	fmt.Fprintf(os.Stdout, "Response from `NewReleaseManagementAPI.ListNewReleases`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiListNewReleasesRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **deploymentType** | **string** |  | [default to &quot;null&quot;]
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**[]ResponseRelease**](ResponseRelease.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNewRelease

> YBPSuccess UpdateNewRelease(ctx, cUUID, rUUID).Release(release).Request(request).Execute()

Update a release



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
	rUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	release := *openapiclient.NewUpdateRelease([]openapiclient.Artifact{*openapiclient.NewArtifact("Architecture_example", "PackageFileId_example", "PackageUrl_example", "Platform_example", "Sha256_example")}, int64(123), "ReleaseNotes_example", "ReleaseTag_example", "State_example") // UpdateRelease | Release data to be updated
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewReleaseManagementAPI.UpdateNewRelease(context.Background(), cUUID, rUUID).Release(release).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewReleaseManagementAPI.UpdateNewRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNewRelease`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `NewReleaseManagementAPI.UpdateNewRelease`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **rUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNewReleaseRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **release** | [**UpdateRelease**](UpdateRelease.md) | Release data to be updated | 
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

