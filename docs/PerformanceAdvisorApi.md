# \PerformanceAdvisorApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](PerformanceAdvisorApi.md#Delete) | **Delete** /api/v1/customers/{cUUID}/performance_recommendations | Delete performance recommendations
[**Get**](PerformanceAdvisorApi.md#Get) | **Get** /api/v1/customers/{cUUID}/performance_recommendations/{rUUID} | Get performance recommendation details
[**GetLatestRun**](PerformanceAdvisorApi.md#GetLatestRun) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/last_run | Get last performance advisor run details
[**GetSettings**](PerformanceAdvisorApi.md#GetSettings) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/perf_advisor_settings | Get universe performance advisor settings
[**Hide**](PerformanceAdvisorApi.md#Hide) | **Post** /api/v1/customers/{cUUID}/performance_recommendations/hide | Hide performance recommendations
[**Page**](PerformanceAdvisorApi.md#Page) | **Post** /api/v1/customers/{cUUID}/performance_recommendations/page | List performance recommendations (paginated)
[**PageAuditInfo**](PerformanceAdvisorApi.md#PageAuditInfo) | **Post** /api/v1/customers/{cUUID}/performance_recommendation_state_change/page | List performance recommendations state change audit events (paginated)
[**Resolve**](PerformanceAdvisorApi.md#Resolve) | **Post** /api/v1/customers/{cUUID}/performance_recommendations/resolve | Resolve performance recommendations
[**RunPerfAdvisor**](PerformanceAdvisorApi.md#RunPerfAdvisor) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/start_manually | Start performance advisor run for universe
[**UpdateSettings**](PerformanceAdvisorApi.md#UpdateSettings) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/perf_advisor_settings | Update universe performance advisor settings



## Delete

> YBPSuccess Delete(ctx, cUUID).DeletePerformanceRecommendationsRequest(deletePerformanceRecommendationsRequest).Request(request).Execute()

Delete performance recommendations

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
    deletePerformanceRecommendationsRequest := *openapiclient.NewPerformanceRecommendationFilter(int64(123), "CustomerId_example", []string{"Ids_example"}, false, []string{"Priorities_example"}, []string{"States_example"}, []string{"Types_example"}, "UniverseId_example") // PerformanceRecommendationFilter | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PerformanceAdvisorApi.Delete(context.Background(), cUUID).DeletePerformanceRecommendationsRequest(deletePerformanceRecommendationsRequest).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deletePerformanceRecommendationsRequest** | [**PerformanceRecommendationFilter**](PerformanceRecommendationFilter.md) |  | 
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


## Get

> PerformanceRecommendation Get(ctx, cUUID, rUUID).Execute()

Get performance recommendation details

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
    rUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PerformanceAdvisorApi.Get(context.Background(), cUUID, rUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: PerformanceRecommendation
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**rUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PerformanceRecommendation**](PerformanceRecommendation.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestRun

> YBPSuccess GetLatestRun(ctx, cUUID, uniUUID).Execute()

Get last performance advisor run details

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
    resp, r, err := api_client.PerformanceAdvisorApi.GetLatestRun(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.GetLatestRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestRun`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.GetLatestRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetSettings

> PerfAdvisorSettingsWithDefaults GetSettings(ctx, cUUID, uniUUID).Request(request).Execute()

Get universe performance advisor settings

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PerformanceAdvisorApi.GetSettings(context.Background(), cUUID, uniUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.GetSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSettings`: PerfAdvisorSettingsWithDefaults
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.GetSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**PerfAdvisorSettingsWithDefaults**](PerfAdvisorSettingsWithDefaults.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Hide

> YBPSuccess Hide(ctx, cUUID).HidePerformanceRecommendationsRequest(hidePerformanceRecommendationsRequest).Request(request).Execute()

Hide performance recommendations

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
    hidePerformanceRecommendationsRequest := *openapiclient.NewPerformanceRecommendationFilter(int64(123), "CustomerId_example", []string{"Ids_example"}, false, []string{"Priorities_example"}, []string{"States_example"}, []string{"Types_example"}, "UniverseId_example") // PerformanceRecommendationFilter | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PerformanceAdvisorApi.Hide(context.Background(), cUUID).HidePerformanceRecommendationsRequest(hidePerformanceRecommendationsRequest).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.Hide``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Hide`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.Hide`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hidePerformanceRecommendationsRequest** | [**PerformanceRecommendationFilter**](PerformanceRecommendationFilter.md) |  | 
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


## Page

> PerformanceRecommendationPagedResponse Page(ctx, cUUID).PagePerformanceRecommendationRequest(pagePerformanceRecommendationRequest).Request(request).Execute()

List performance recommendations (paginated)

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
    pagePerformanceRecommendationRequest := *openapiclient.NewPerformanceRecommendationPagedQuery("Direction_example", *openapiclient.NewPerformanceRecommendationFilter(int64(123), "CustomerId_example", []string{"Ids_example"}, false, []string{"Priorities_example"}, []string{"States_example"}, []string{"Types_example"}, "UniverseId_example"), int32(123), false, int32(123), "SortBy_example") // PerformanceRecommendationPagedQuery | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PerformanceAdvisorApi.Page(context.Background(), cUUID).PagePerformanceRecommendationRequest(pagePerformanceRecommendationRequest).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.Page``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Page`: PerformanceRecommendationPagedResponse
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.Page`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pagePerformanceRecommendationRequest** | [**PerformanceRecommendationPagedQuery**](PerformanceRecommendationPagedQuery.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**PerformanceRecommendationPagedResponse**](PerformanceRecommendationPagedResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageAuditInfo

> StateChangeAuditInfoPagedResponse PageAuditInfo(ctx, cUUID).PageStateChangeAuditInfoRequest(pageStateChangeAuditInfoRequest).Request(request).Execute()

List performance recommendations state change audit events (paginated)

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
    pageStateChangeAuditInfoRequest := *openapiclient.NewStateChangeAuditInfoPagedQuery("Direction_example", *openapiclient.NewStateChangeAuditInfoFilter("CustomerId_example", []string{"Ids_example"}, "RecommendationId_example", "UserId_example"), int32(123), false, int32(123), "SortBy_example") // StateChangeAuditInfoPagedQuery | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PerformanceAdvisorApi.PageAuditInfo(context.Background(), cUUID).PageStateChangeAuditInfoRequest(pageStateChangeAuditInfoRequest).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.PageAuditInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PageAuditInfo`: StateChangeAuditInfoPagedResponse
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.PageAuditInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageAuditInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageStateChangeAuditInfoRequest** | [**StateChangeAuditInfoPagedQuery**](StateChangeAuditInfoPagedQuery.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**StateChangeAuditInfoPagedResponse**](StateChangeAuditInfoPagedResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Resolve

> YBPSuccess Resolve(ctx, cUUID).ResolvePerformanceRecommendationsRequest(resolvePerformanceRecommendationsRequest).Request(request).Execute()

Resolve performance recommendations

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
    resolvePerformanceRecommendationsRequest := *openapiclient.NewPerformanceRecommendationFilter(int64(123), "CustomerId_example", []string{"Ids_example"}, false, []string{"Priorities_example"}, []string{"States_example"}, []string{"Types_example"}, "UniverseId_example") // PerformanceRecommendationFilter | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PerformanceAdvisorApi.Resolve(context.Background(), cUUID).ResolvePerformanceRecommendationsRequest(resolvePerformanceRecommendationsRequest).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.Resolve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Resolve`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.Resolve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resolvePerformanceRecommendationsRequest** | [**PerformanceRecommendationFilter**](PerformanceRecommendationFilter.md) |  | 
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


## RunPerfAdvisor

> PerfAdvisorManualRunStatus RunPerfAdvisor(ctx, cUUID, uniUUID).Request(request).Execute()

Start performance advisor run for universe

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PerformanceAdvisorApi.RunPerfAdvisor(context.Background(), cUUID, uniUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.RunPerfAdvisor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunPerfAdvisor`: PerfAdvisorManualRunStatus
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.RunPerfAdvisor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunPerfAdvisorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**PerfAdvisorManualRunStatus**](PerfAdvisorManualRunStatus.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSettings

> YBPSuccess UpdateSettings(ctx, cUUID, uniUUID).PerformanceAdvisorSettingsRequest(performanceAdvisorSettingsRequest).Request(request).Execute()

Update universe performance advisor settings

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
    performanceAdvisorSettingsRequest := *openapiclient.NewPerfAdvisorSettingsFormData() // PerfAdvisorSettingsFormData | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PerformanceAdvisorApi.UpdateSettings(context.Background(), cUUID, uniUUID).PerformanceAdvisorSettingsRequest(performanceAdvisorSettingsRequest).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PerformanceAdvisorApi.UpdateSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSettings`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `PerformanceAdvisorApi.UpdateSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **performanceAdvisorSettingsRequest** | [**PerfAdvisorSettingsFormData**](PerfAdvisorSettingsFormData.md) |  | 
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

