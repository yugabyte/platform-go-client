# \YbcManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableYbc**](YbcManagementApi.md#DisableYbc) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/ybc/disable | Disable YBC on the universe nodes
[**InstallYbc**](YbcManagementApi.md#InstallYbc) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/ybc/install | Install YBC on the universe nodes
[**UpgradeYbc**](YbcManagementApi.md#UpgradeYbc) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/ybc/upgrade | Upgrade YBC on the universe nodes
[**UpgradeYbcGflags**](YbcManagementApi.md#UpgradeYbcGflags) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/ybc/upgrade/gflags | Upgrade YBC gflags on the universe nodes



## DisableYbc

> YBPTask DisableYbc(ctx, cUUID, uniUUID).Request(request).Execute()

Disable YBC on the universe nodes



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
    resp, r, err := api_client.YbcManagementApi.DisableYbc(context.Background(), cUUID, uniUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `YbcManagementApi.DisableYbc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableYbc`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `YbcManagementApi.DisableYbc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableYbcRequest struct via the builder pattern


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


## InstallYbc

> YBPTask InstallYbc(ctx, cUUID, uniUUID).YbcVersion(ybcVersion).Request(request).Execute()

Install YBC on the universe nodes



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
    ybcVersion := "ybcVersion_example" // string |  (optional)
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.YbcManagementApi.InstallYbc(context.Background(), cUUID, uniUUID).YbcVersion(ybcVersion).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `YbcManagementApi.InstallYbc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstallYbc`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `YbcManagementApi.InstallYbc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallYbcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ybcVersion** | **string** |  | 
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


## UpgradeYbc

> YBPTask UpgradeYbc(ctx, cUUID, uniUUID).YbcVersion(ybcVersion).Request(request).Execute()

Upgrade YBC on the universe nodes



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
    ybcVersion := "ybcVersion_example" // string |  (optional)
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.YbcManagementApi.UpgradeYbc(context.Background(), cUUID, uniUUID).YbcVersion(ybcVersion).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `YbcManagementApi.UpgradeYbc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeYbc`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `YbcManagementApi.UpgradeYbc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeYbcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ybcVersion** | **string** |  | 
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


## UpgradeYbcGflags

> YBPTask UpgradeYbcGflags(ctx, cUUID, uniUUID).Request(request).Execute()

Upgrade YBC gflags on the universe nodes



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
    resp, r, err := api_client.YbcManagementApi.UpgradeYbcGflags(context.Background(), cUUID, uniUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `YbcManagementApi.UpgradeYbcGflags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeYbcGflags`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `YbcManagementApi.UpgradeYbcGflags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeYbcGflagsRequest struct via the builder pattern


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

