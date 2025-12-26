# \GFlagsValidationAPIsAPI

All URIs are relative to *http://localhost*

| Method                                                              | HTTP request                                                | Description                   |
| ------------------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------- |
| [**GetGFlagMetadata**](GFlagsValidationAPIsAPI.md#GetGFlagMetadata) | **Get** /api/v1/metadata/version/{version}/gflag            | Get gflag metadata            |
| [**ListGFlags**](GFlagsValidationAPIsAPI.md#ListGFlags)             | **Get** /api/v1/metadata/version/{version}/list_gflags      | List all gflags for a release |
| [**ValidateGFlags**](GFlagsValidationAPIsAPI.md#ValidateGFlags)     | **Post** /api/v1/metadata/version/{version}/validate_gflags | Validate gflags               |



## GetGFlagMetadata

> GFlagDetails GetGFlagMetadata(ctx, version).Name(name).Server(server).Execute()

Get gflag metadata



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
	version := "version_example" // string | 
	name := "name_example" // string |  (optional)
	server := "server_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GFlagsValidationAPIsAPI.GetGFlagMetadata(context.Background(), version).Name(name).Server(server).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GFlagsValidationAPIsAPI.GetGFlagMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGFlagMetadata`: GFlagDetails
	fmt.Fprintf(os.Stdout, "Response from `GFlagsValidationAPIsAPI.GetGFlagMetadata`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **version** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetGFlagMetadataRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **name** | **string** |  | 
 **server** | **string** |  | 

### Return type

[**GFlagDetails**](GFlagDetails.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGFlags

> []GFlagDetails ListGFlags(ctx, version).Name(name).Server(server).MostUsedGFlags(mostUsedGFlags).ShowExperimental(showExperimental).Execute()

List all gflags for a release



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
	version := "version_example" // string | 
	name := "name_example" // string |  (optional)
	server := "server_example" // string |  (optional)
	mostUsedGFlags := true // bool |  (optional) (default to false)
	showExperimental := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GFlagsValidationAPIsAPI.ListGFlags(context.Background(), version).Name(name).Server(server).MostUsedGFlags(mostUsedGFlags).ShowExperimental(showExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GFlagsValidationAPIsAPI.ListGFlags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGFlags`: []GFlagDetails
	fmt.Fprintf(os.Stdout, "Response from `GFlagsValidationAPIsAPI.ListGFlags`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **version** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiListGFlagsRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **name** | **string** |  | 
 **server** | **string** |  | 
 **mostUsedGFlags** | **bool** |  | [default to false]
 **showExperimental** | **bool** |  | [default to false]

### Return type

[**[]GFlagDetails**](GFlagDetails.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateGFlags

> []GFlagsValidationResponse ValidateGFlags(ctx, version).GflagValidationFormData(gflagValidationFormData).Request(request).Execute()

Validate gflags



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
	version := "version_example" // string | 
	gflagValidationFormData := *openapiclient.NewGFlagsValidationFormData() // GFlagsValidationFormData | GFlag validation form data
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GFlagsValidationAPIsAPI.ValidateGFlags(context.Background(), version).GflagValidationFormData(gflagValidationFormData).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GFlagsValidationAPIsAPI.ValidateGFlags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateGFlags`: []GFlagsValidationResponse
	fmt.Fprintf(os.Stdout, "Response from `GFlagsValidationAPIsAPI.ValidateGFlags`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **version** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiValidateGFlagsRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **gflagValidationFormData** | [**GFlagsValidationFormData**](GFlagsValidationFormData.md) | GFlag validation form data | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**[]GFlagsValidationResponse**](GFlagsValidationResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

