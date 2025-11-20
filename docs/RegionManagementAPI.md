# \RegionManagementAPI

All URIs are relative to *http://localhost*

| Method                                                                  | HTTP request                                                                 | Description                    |
| ----------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ------------------------------ |
| [**CreateProviderRegion**](RegionManagementAPI.md#CreateProviderRegion) | **Post** /api/v1/customers/{cUUID}/providers/{pUUID}/provider_regions        | Create a new region            |
| [**CreateRegion**](RegionManagementAPI.md#CreateRegion)                 | **Post** /api/v1/customers/{cUUID}/providers/{pUUID}/regions                 | Create Region - deprecated     |
| [**DeleteRegion**](RegionManagementAPI.md#DeleteRegion)                 | **Delete** /api/v1/customers/{cUUID}/providers/{pUUID}/regions/{rUUID}       | Delete a region                |
| [**EditProviderRegion**](RegionManagementAPI.md#EditProviderRegion)     | **Put** /api/v1/customers/{cUUID}/providers/{pUUID}/provider_regions/{rUUID} | Modify a region                |
| [**EditRegion**](RegionManagementAPI.md#EditRegion)                     | **Put** /api/v1/customers/{cUUID}/providers/{pUUID}/regions/{rUUID}          | Edit regions - deprecated      |
| [**GetRegion**](RegionManagementAPI.md#GetRegion)                       | **Get** /api/v1/customers/{cUUID}/providers/{pUUID}/regions                  | List a provider&#39;s regions  |
| [**ListAllRegions**](RegionManagementAPI.md#ListAllRegions)             | **Get** /api/v1/customers/{cUUID}/regions                                    | List regions for all providers |



## CreateProviderRegion

> Region CreateProviderRegion(ctx, cUUID, pUUID).Region(region).Request(request).Execute()

Create a new region



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
	region := *openapiclient.NewRegion([]openapiclient.AvailabilityZone{*openapiclient.NewAvailabilityZone("us-west1-a")}) // Region | Specification of Region to be created
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionManagementAPI.CreateProviderRegion(context.Background(), cUUID, pUUID).Region(region).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionManagementAPI.CreateProviderRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProviderRegion`: Region
	fmt.Fprintf(os.Stdout, "Response from `RegionManagementAPI.CreateProviderRegion`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **pUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProviderRegionRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **region** | [**Region**](Region.md) | Specification of Region to be created | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**Region**](Region.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRegion

> Region CreateRegion(ctx, cUUID, pUUID).Region(region).Request(request).Execute()

Create Region - deprecated



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
	region := *openapiclient.NewRegionFormData("Code_example", "DestVpcId_example", "HostVpcId_example", "HostVpcRegion_example", float64(123), float64(123), "Name_example", "SecurityGroupId_example", "VnetName_example", "YbImage_example") // RegionFormData | region form data for new region to be created
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionManagementAPI.CreateRegion(context.Background(), cUUID, pUUID).Region(region).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionManagementAPI.CreateRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRegion`: Region
	fmt.Fprintf(os.Stdout, "Response from `RegionManagementAPI.CreateRegion`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **pUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRegionRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **region** | [**RegionFormData**](RegionFormData.md) | region form data for new region to be created | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**Region**](Region.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRegion

> map[string]interface{} DeleteRegion(ctx, cUUID, pUUID, rUUID).Request(request).Execute()

Delete a region

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
	rUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionManagementAPI.DeleteRegion(context.Background(), cUUID, pUUID, rUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionManagementAPI.DeleteRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRegion`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RegionManagementAPI.DeleteRegion`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **pUUID** | **string**          |                                                                             |
| **rUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRegionRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditProviderRegion

> Region EditProviderRegion(ctx, cUUID, pUUID, rUUID).Region(region).Request(request).Execute()

Modify a region



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
	rUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	region := *openapiclient.NewRegion([]openapiclient.AvailabilityZone{*openapiclient.NewAvailabilityZone("us-west1-a")}) // Region | Specification of Region to be edited
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionManagementAPI.EditProviderRegion(context.Background(), cUUID, pUUID, rUUID).Region(region).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionManagementAPI.EditProviderRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditProviderRegion`: Region
	fmt.Fprintf(os.Stdout, "Response from `RegionManagementAPI.EditProviderRegion`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **pUUID** | **string**          |                                                                             |
| **rUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiEditProviderRegionRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



 **region** | [**Region**](Region.md) | Specification of Region to be edited | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**Region**](Region.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditRegion

> map[string]interface{} EditRegion(ctx, cUUID, pUUID, rUUID).Region(region).Request(request).Execute()

Edit regions - deprecated



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
	rUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	region := *openapiclient.NewRegionFormData("Code_example", "DestVpcId_example", "HostVpcId_example", "HostVpcRegion_example", float64(123), float64(123), "Name_example", "SecurityGroupId_example", "VnetName_example", "YbImage_example") // RegionFormData | region edit form data
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionManagementAPI.EditRegion(context.Background(), cUUID, pUUID, rUUID).Region(region).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionManagementAPI.EditRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditRegion`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RegionManagementAPI.EditRegion`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **pUUID** | **string**          |                                                                             |
| **rUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiEditRegionRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



 **region** | [**RegionFormData**](RegionFormData.md) | region edit form data | 
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


## GetRegion

> []Region GetRegion(ctx, cUUID, pUUID).Execute()

List a provider's regions

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionManagementAPI.GetRegion(context.Background(), cUUID, pUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionManagementAPI.GetRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegion`: []Region
	fmt.Fprintf(os.Stdout, "Response from `RegionManagementAPI.GetRegion`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **pUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

[**[]Region**](Region.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllRegions

> []Region ListAllRegions(ctx, cUUID).Execute()

List regions for all providers

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
	resp, r, err := apiClient.RegionManagementAPI.ListAllRegions(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionManagementAPI.ListAllRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllRegions`: []Region
	fmt.Fprintf(os.Stdout, "Response from `RegionManagementAPI.ListAllRegions`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiListAllRegionsRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


### Return type

[**[]Region**](Region.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

