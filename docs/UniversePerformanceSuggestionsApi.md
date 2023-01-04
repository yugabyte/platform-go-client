# \UniversePerformanceSuggestionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQueryDistributionSuggestions**](UniversePerformanceSuggestionsApi.md#GetQueryDistributionSuggestions) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/query_distribution_suggestions | Get query distribution improvement suggestion for a universe
[**GetRangeHash**](UniversePerformanceSuggestionsApi.md#GetRangeHash) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/range_hash | Return list of hash indexes
[**GetUnusedIndexes**](UniversePerformanceSuggestionsApi.md#GetUnusedIndexes) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/unused_indexes | Return list of each unused index across the universe



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
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    uniUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniversePerformanceSuggestionsApi.GetQueryDistributionSuggestions(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniversePerformanceSuggestionsApi.GetQueryDistributionSuggestions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQueryDistributionSuggestions`: QueryDistributionSuggestionResponse
    fmt.Fprintf(os.Stdout, "Response from `UniversePerformanceSuggestionsApi.GetQueryDistributionSuggestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueryDistributionSuggestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    uniUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniversePerformanceSuggestionsApi.GetRangeHash(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniversePerformanceSuggestionsApi.GetRangeHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRangeHash`: HashedTimestampColumnFinderResponse
    fmt.Fprintf(os.Stdout, "Response from `UniversePerformanceSuggestionsApi.GetRangeHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRangeHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    uniUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniversePerformanceSuggestionsApi.GetUnusedIndexes(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniversePerformanceSuggestionsApi.GetUnusedIndexes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUnusedIndexes`: UnusedIndexFinderResponse
    fmt.Fprintf(os.Stdout, "Response from `UniversePerformanceSuggestionsApi.GetUnusedIndexes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnusedIndexesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

