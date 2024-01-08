# \UniverseInformationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadNodeLogs**](UniverseInformationApi.md#DownloadNodeLogs) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/{nodeName}/download_logs | Deprecated since YBA version 2.20.0.0 (use support bundle). Download a node&#39;s logs.
[**GetLiveQueries**](UniverseInformationApi.md#GetLiveQueries) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/live_queries | Available since YBA version 2.4.0.0. Get live queries for a universe
[**GetMasterLeaderIP**](UniverseInformationApi.md#GetMasterLeaderIP) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/leader | Available since YBA version 2.2.0.0. Get IP address of a universe&#39;s master leader
[**GetSlowQueries**](UniverseInformationApi.md#GetSlowQueries) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/slow_queries | Available since YBA version 2.6.0.0. Get slow queries for a universe
[**GetUniverseCost**](UniverseInformationApi.md#GetUniverseCost) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/cost | Deprecated since YBA version 2.20.0.0(use /universe_resources). Get a cost estimate for a universe.
[**GetUniverseResources**](UniverseInformationApi.md#GetUniverseResources) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/universe_resources | Available since YBA version 2.20.0.0. Get a resource usage estimate for a universe
[**HealthCheckUniverse**](UniverseInformationApi.md#HealthCheckUniverse) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/health_check | Available since YBA version 2.2.0.0. Return results for the last health check
[**ResetSlowQueries**](UniverseInformationApi.md#ResetSlowQueries) | **Delete** /api/v1/customers/{cUUID}/universes/{uniUUID}/slow_queries | Available since YBA version 2.6.0.0. Reset slow queries for a universe
[**UniverseStatus**](UniverseInformationApi.md#UniverseStatus) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/status | Available since YBA version 2.2.0.0. Get a universe&#39;s status



## DownloadNodeLogs

> string DownloadNodeLogs(ctx, cUUID, uniUUID, nodeName).Execute()

Deprecated since YBA version 2.20.0.0 (use support bundle). Download a node's logs.



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
    nodeName := "nodeName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseInformationApi.DownloadNodeLogs(context.Background(), cUUID, uniUUID, nodeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseInformationApi.DownloadNodeLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadNodeLogs`: string
    fmt.Fprintf(os.Stdout, "Response from `UniverseInformationApi.DownloadNodeLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 
**nodeName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadNodeLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## GetLiveQueries

> map[string]interface{} GetLiveQueries(ctx, cUUID, uniUUID).Execute()

Available since YBA version 2.4.0.0. Get live queries for a universe

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
    resp, r, err := api_client.UniverseInformationApi.GetLiveQueries(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseInformationApi.GetLiveQueries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLiveQueries`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UniverseInformationApi.GetLiveQueries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLiveQueriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetMasterLeaderIP

> map[string]interface{} GetMasterLeaderIP(ctx, cUUID, uniUUID).Execute()

Available since YBA version 2.2.0.0. Get IP address of a universe's master leader

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
    resp, r, err := api_client.UniverseInformationApi.GetMasterLeaderIP(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseInformationApi.GetMasterLeaderIP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMasterLeaderIP`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UniverseInformationApi.GetMasterLeaderIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMasterLeaderIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetSlowQueries

> map[string]interface{} GetSlowQueries(ctx, cUUID, uniUUID).Execute()

Available since YBA version 2.6.0.0. Get slow queries for a universe

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
    resp, r, err := api_client.UniverseInformationApi.GetSlowQueries(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseInformationApi.GetSlowQueries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSlowQueries`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UniverseInformationApi.GetSlowQueries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSlowQueriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetUniverseCost

> UniverseResourceDetails GetUniverseCost(ctx, cUUID, uniUUID).Execute()

Deprecated since YBA version 2.20.0.0(use /universe_resources). Get a cost estimate for a universe.

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
    resp, r, err := api_client.UniverseInformationApi.GetUniverseCost(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseInformationApi.GetUniverseCost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUniverseCost`: UniverseResourceDetails
    fmt.Fprintf(os.Stdout, "Response from `UniverseInformationApi.GetUniverseCost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UniverseResourceDetails**](UniverseResourceDetails.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniverseResources

> UniverseResourceDetails GetUniverseResources(ctx, cUUID, uniUUID).Execute()

Available since YBA version 2.20.0.0. Get a resource usage estimate for a universe



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
    resp, r, err := api_client.UniverseInformationApi.GetUniverseResources(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseInformationApi.GetUniverseResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUniverseResources`: UniverseResourceDetails
    fmt.Fprintf(os.Stdout, "Response from `UniverseInformationApi.GetUniverseResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UniverseResourceDetails**](UniverseResourceDetails.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealthCheckUniverse

> []Details HealthCheckUniverse(ctx, cUUID, uniUUID).Execute()

Available since YBA version 2.2.0.0. Return results for the last health check



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
    resp, r, err := api_client.UniverseInformationApi.HealthCheckUniverse(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseInformationApi.HealthCheckUniverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HealthCheckUniverse`: []Details
    fmt.Fprintf(os.Stdout, "Response from `UniverseInformationApi.HealthCheckUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHealthCheckUniverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Details**](Details.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetSlowQueries

> map[string]interface{} ResetSlowQueries(ctx, cUUID, uniUUID).Request(request).Execute()

Available since YBA version 2.6.0.0. Reset slow queries for a universe

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
    resp, r, err := api_client.UniverseInformationApi.ResetSlowQueries(context.Background(), cUUID, uniUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseInformationApi.ResetSlowQueries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetSlowQueries`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UniverseInformationApi.ResetSlowQueries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetSlowQueriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UniverseStatus

> map[string]map[string]interface{} UniverseStatus(ctx, cUUID, uniUUID).Execute()

Available since YBA version 2.2.0.0. Get a universe's status



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
    resp, r, err := api_client.UniverseInformationApi.UniverseStatus(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseInformationApi.UniverseStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UniverseStatus`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UniverseInformationApi.UniverseStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUniverseStatusRequest struct via the builder pattern


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

