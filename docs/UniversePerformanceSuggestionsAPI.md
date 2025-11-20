# \UniversePerformanceSuggestionsAPI

All URIs are relative to *http://localhost*

| Method                                                                                                      | HTTP request                                                                         | Description                                                  |
| ----------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------ | ------------------------------------------------------------ |
| [**GetQueryDistributionSuggestions**](UniversePerformanceSuggestionsAPI.md#GetQueryDistributionSuggestions) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/query_distribution_suggestions | Get query distribution improvement suggestion for a universe |
| [**GetRangeHash**](UniversePerformanceSuggestionsAPI.md#GetRangeHash)                                       | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/range_hash                     | Return list of hash indexes                                  |
| [**GetUnusedIndexes**](UniversePerformanceSuggestionsAPI.md#GetUnusedIndexes)                               | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/unused_indexes                 | Return list of each unused index across the universe         |



## GetQueryDistributionSuggestions

> QueryDistributionSuggestionResponse GetQueryDistributionSuggestions(ctx, cUUID, uniUUID).Execute()

Get query distribution improvement suggestion for a universe



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
	resp, r, err := apiClient.UniversePerformanceSuggestionsAPI.GetQueryDistributionSuggestions(context.Background(), cUUID, uniUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniversePerformanceSuggestionsAPI.GetQueryDistributionSuggestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQueryDistributionSuggestions`: QueryDistributionSuggestionResponse
	fmt.Fprintf(os.Stdout, "Response from `UniversePerformanceSuggestionsAPI.GetQueryDistributionSuggestions`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueryDistributionSuggestionsRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

[**QueryDistributionSuggestionResponse**](QueryDistributionSuggestionResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRangeHash

> HashedTimestampColumnFinderResponse GetRangeHash(ctx, cUUID, uniUUID).Execute()

Return list of hash indexes



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
	resp, r, err := apiClient.UniversePerformanceSuggestionsAPI.GetRangeHash(context.Background(), cUUID, uniUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniversePerformanceSuggestionsAPI.GetRangeHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRangeHash`: HashedTimestampColumnFinderResponse
	fmt.Fprintf(os.Stdout, "Response from `UniversePerformanceSuggestionsAPI.GetRangeHash`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetRangeHashRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

[**HashedTimestampColumnFinderResponse**](HashedTimestampColumnFinderResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnusedIndexes

> UnusedIndexFinderResponse GetUnusedIndexes(ctx, cUUID, uniUUID).Execute()

Return list of each unused index across the universe



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
	resp, r, err := apiClient.UniversePerformanceSuggestionsAPI.GetUnusedIndexes(context.Background(), cUUID, uniUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniversePerformanceSuggestionsAPI.GetUnusedIndexes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUnusedIndexes`: UnusedIndexFinderResponse
	fmt.Fprintf(os.Stdout, "Response from `UniversePerformanceSuggestionsAPI.GetUnusedIndexes`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnusedIndexesRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

[**UnusedIndexFinderResponse**](UnusedIndexFinderResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

