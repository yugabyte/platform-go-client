# \UniverseManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigureUniverseAlerts**](UniverseManagementApi.md#ConfigureUniverseAlerts) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/config_alerts | Configure alerts for a universe
[**DeleteUniverse**](UniverseManagementApi.md#DeleteUniverse) | **Delete** /api/v1/customers/{cUUID}/universes/{uniUUID} | Delete a universe
[**GetUniverse**](UniverseManagementApi.md#GetUniverse) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID} | Get a universe
[**ListUniverses**](UniverseManagementApi.md#ListUniverses) | **Get** /api/v1/customers/{cUUID}/universes | List universes
[**PauseUniverse**](UniverseManagementApi.md#PauseUniverse) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/pause | Pause a universe
[**ResetUniverseVersion**](UniverseManagementApi.md#ResetUniverseVersion) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/setup_universe_2dc | Reset universe version
[**ResumeUniverse**](UniverseManagementApi.md#ResumeUniverse) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/resume | Resume a paused universe
[**SetUniverseBackupFlag**](UniverseManagementApi.md#SetUniverseBackupFlag) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/update_backup_state | Set a universe&#39;s backup flag
[**SetUniverseHelm3Compatible**](UniverseManagementApi.md#SetUniverseHelm3Compatible) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/mark_helm3_compatible | Flag a universe as Helm 3-compatible
[**SetUniverseKey**](UniverseManagementApi.md#SetUniverseKey) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/set_key | Set a universe&#39;s key
[**UpdateLoadBalancerConfig**](UniverseManagementApi.md#UpdateLoadBalancerConfig) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/update_lb_config | Update load balancer config



## ConfigureUniverseAlerts

> YBPSuccess ConfigureUniverseAlerts(ctx, cUUID, uniUUID).Request(request).Execute()

Configure alerts for a universe

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
    resp, r, err := api_client.UniverseManagementApi.ConfigureUniverseAlerts(context.Background(), cUUID, uniUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementApi.ConfigureUniverseAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfigureUniverseAlerts`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `UniverseManagementApi.ConfigureUniverseAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigureUniverseAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteUniverse

> YBPTask DeleteUniverse(ctx, cUUID, uniUUID).IsForceDelete(isForceDelete).IsDeleteBackups(isDeleteBackups).IsDeleteAssociatedCerts(isDeleteAssociatedCerts).Request(request).Execute()

Delete a universe

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
    isForceDelete := true // bool |  (optional) (default to false)
    isDeleteBackups := true // bool |  (optional) (default to false)
    isDeleteAssociatedCerts := true // bool |  (optional) (default to false)
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseManagementApi.DeleteUniverse(context.Background(), cUUID, uniUUID).IsForceDelete(isForceDelete).IsDeleteBackups(isDeleteBackups).IsDeleteAssociatedCerts(isDeleteAssociatedCerts).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementApi.DeleteUniverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUniverse`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseManagementApi.DeleteUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUniverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isForceDelete** | **bool** |  | [default to false]
 **isDeleteBackups** | **bool** |  | [default to false]
 **isDeleteAssociatedCerts** | **bool** |  | [default to false]
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUniverse

> UniverseResp GetUniverse(ctx, cUUID, uniUUID).Execute()

Get a universe

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
    resp, r, err := api_client.UniverseManagementApi.GetUniverse(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementApi.GetUniverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUniverse`: UniverseResp
    fmt.Fprintf(os.Stdout, "Response from `UniverseManagementApi.GetUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UniverseResp**](UniverseResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUniverses

> []UniverseResp ListUniverses(ctx, cUUID).Name(name).Execute()

List universes

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
    name := "name_example" // string |  (optional) (default to "null")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseManagementApi.ListUniverses(context.Background(), cUUID).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementApi.ListUniverses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUniverses`: []UniverseResp
    fmt.Fprintf(os.Stdout, "Response from `UniverseManagementApi.ListUniverses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUniversesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | [default to &quot;null&quot;]

### Return type

[**[]UniverseResp**](UniverseResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseUniverse

> YBPTask PauseUniverse(ctx, cUUID, uniUUID).Request(request).Execute()

Pause a universe

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
    resp, r, err := api_client.UniverseManagementApi.PauseUniverse(context.Background(), cUUID, uniUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementApi.PauseUniverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PauseUniverse`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseManagementApi.PauseUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseUniverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetUniverseVersion

> YBPSuccess ResetUniverseVersion(ctx, cUUID, uniUUID).Request(request).Execute()

Reset universe version

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
    resp, r, err := api_client.UniverseManagementApi.ResetUniverseVersion(context.Background(), cUUID, uniUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementApi.ResetUniverseVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetUniverseVersion`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `UniverseManagementApi.ResetUniverseVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetUniverseVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ResumeUniverse

> YBPTask ResumeUniverse(ctx, cUUID, uniUUID).Request(request).Execute()

Resume a paused universe

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
    resp, r, err := api_client.UniverseManagementApi.ResumeUniverse(context.Background(), cUUID, uniUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementApi.ResumeUniverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeUniverse`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseManagementApi.ResumeUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeUniverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetUniverseBackupFlag

> YBPSuccess SetUniverseBackupFlag(ctx, cUUID, uniUUID).MarkActive(markActive).Request(request).Execute()

Set a universe's backup flag

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
    markActive := true // bool |  (optional)
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseManagementApi.SetUniverseBackupFlag(context.Background(), cUUID, uniUUID).MarkActive(markActive).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementApi.SetUniverseBackupFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetUniverseBackupFlag`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `UniverseManagementApi.SetUniverseBackupFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetUniverseBackupFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **markActive** | **bool** |  | 
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


## SetUniverseHelm3Compatible

> YBPSuccess SetUniverseHelm3Compatible(ctx, cUUID, uniUUID).Request(request).Execute()

Flag a universe as Helm 3-compatible

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
    resp, r, err := api_client.UniverseManagementApi.SetUniverseHelm3Compatible(context.Background(), cUUID, uniUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementApi.SetUniverseHelm3Compatible``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetUniverseHelm3Compatible`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `UniverseManagementApi.SetUniverseHelm3Compatible`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetUniverseHelm3CompatibleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## SetUniverseKey

> UniverseResp SetUniverseKey(ctx, cUUID, uniUUID).Request(request).Execute()

Set a universe's key

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
    resp, r, err := api_client.UniverseManagementApi.SetUniverseKey(context.Background(), cUUID, uniUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementApi.SetUniverseKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetUniverseKey`: UniverseResp
    fmt.Fprintf(os.Stdout, "Response from `UniverseManagementApi.SetUniverseKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetUniverseKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**UniverseResp**](UniverseResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoadBalancerConfig

> UpdateLoadBalancerConfig UpdateLoadBalancerConfig(ctx, cUUID, uniUUID).Request(request).Execute()

Update load balancer config

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
    resp, r, err := api_client.UniverseManagementApi.UpdateLoadBalancerConfig(context.Background(), cUUID, uniUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseManagementApi.UpdateLoadBalancerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLoadBalancerConfig`: UpdateLoadBalancerConfig
    fmt.Fprintf(os.Stdout, "Response from `UniverseManagementApi.UpdateLoadBalancerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoadBalancerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**UpdateLoadBalancerConfig**](UpdateLoadBalancerConfig.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

