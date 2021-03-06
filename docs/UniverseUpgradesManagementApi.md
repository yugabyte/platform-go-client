# \UniverseUpgradesManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResizeNode**](UniverseUpgradesManagementApi.md#ResizeNode) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/resize_node | Resize Node
[**RestartUniverse**](UniverseUpgradesManagementApi.md#RestartUniverse) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/restart | Restart Universe
[**UpgradeCerts**](UniverseUpgradesManagementApi.md#UpgradeCerts) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/certs | Upgrade Certs
[**UpgradeGFlags**](UniverseUpgradesManagementApi.md#UpgradeGFlags) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/gflags | Upgrade GFlags
[**UpgradeSoftware**](UniverseUpgradesManagementApi.md#UpgradeSoftware) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/software | Upgrade Software
[**UpgradeSystemd**](UniverseUpgradesManagementApi.md#UpgradeSystemd) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/systemd | Upgrade Systemd
[**UpgradeTls**](UniverseUpgradesManagementApi.md#UpgradeTls) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/tls | Upgrade TLS
[**UpgradeVMImage**](UniverseUpgradesManagementApi.md#UpgradeVMImage) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/vm | Upgrade VM Image



## ResizeNode

> YBPTask ResizeNode(ctx, cUUID, uniUUID).ResizeNodeParams(resizeNodeParams).Execute()

Resize Node



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
    resizeNodeParams := *openapiclient.NewResizeNodeParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, false, false, int32(123), int32(123), "UpgradeOption_example") // ResizeNodeParams | Resize Node Params

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.ResizeNode(context.Background(), cUUID, uniUUID).ResizeNodeParams(resizeNodeParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseUpgradesManagementApi.ResizeNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResizeNode`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseUpgradesManagementApi.ResizeNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResizeNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resizeNodeParams** | [**ResizeNodeParams**](ResizeNodeParams.md) | Resize Node Params | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartUniverse

> YBPTask RestartUniverse(ctx, cUUID, uniUUID).UpgradeTaskParams(upgradeTaskParams).Execute()

Restart Universe



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
    upgradeTaskParams := *openapiclient.NewUpgradeTaskParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, false, int32(123), int32(123), "UpgradeOption_example") // UpgradeTaskParams | Upgrade Task Params

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.RestartUniverse(context.Background(), cUUID, uniUUID).UpgradeTaskParams(upgradeTaskParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseUpgradesManagementApi.RestartUniverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestartUniverse`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseUpgradesManagementApi.RestartUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartUniverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upgradeTaskParams** | [**UpgradeTaskParams**](UpgradeTaskParams.md) | Upgrade Task Params | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeCerts

> YBPTask UpgradeCerts(ctx, cUUID, uniUUID).CertsRotateParams(certsRotateParams).Execute()

Upgrade Certs



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
    certsRotateParams := *openapiclient.NewCertsRotateParams("ClientRootCA_example", []openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, false, false, "RootCA_example", int32(123), int32(123), "UpgradeOption_example") // CertsRotateParams | Certs Rotate Params

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.UpgradeCerts(context.Background(), cUUID, uniUUID).CertsRotateParams(certsRotateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseUpgradesManagementApi.UpgradeCerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeCerts`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseUpgradesManagementApi.UpgradeCerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeCertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **certsRotateParams** | [**CertsRotateParams**](CertsRotateParams.md) | Certs Rotate Params | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeGFlags

> YBPTask UpgradeGFlags(ctx, cUUID, uniUUID).GflagsUpgradeParams(gflagsUpgradeParams).Execute()

Upgrade GFlags



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
    gflagsUpgradeParams := *openapiclient.NewGFlagsUpgradeParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, false, map[string]string{"key": "Inner_example"}, int32(123), int32(123), map[string]string{"key": "Inner_example"}, "UpgradeOption_example") // GFlagsUpgradeParams | GFlags Upgrade Params

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.UpgradeGFlags(context.Background(), cUUID, uniUUID).GflagsUpgradeParams(gflagsUpgradeParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseUpgradesManagementApi.UpgradeGFlags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeGFlags`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseUpgradesManagementApi.UpgradeGFlags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeGFlagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gflagsUpgradeParams** | [**GFlagsUpgradeParams**](GFlagsUpgradeParams.md) | GFlags Upgrade Params | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeSoftware

> YBPTask UpgradeSoftware(ctx, cUUID, uniUUID).SoftwareUpgradeParams(softwareUpgradeParams).Execute()

Upgrade Software



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
    softwareUpgradeParams := *openapiclient.NewSoftwareUpgradeParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, false, int32(123), int32(123), "UpgradeOption_example", "YbSoftwareVersion_example") // SoftwareUpgradeParams | Software Upgrade Params

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.UpgradeSoftware(context.Background(), cUUID, uniUUID).SoftwareUpgradeParams(softwareUpgradeParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseUpgradesManagementApi.UpgradeSoftware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeSoftware`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseUpgradesManagementApi.UpgradeSoftware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeSoftwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **softwareUpgradeParams** | [**SoftwareUpgradeParams**](SoftwareUpgradeParams.md) | Software Upgrade Params | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeSystemd

> YBPTask UpgradeSystemd(ctx, cUUID, uniUUID).SystemdUpgradeParams(systemdUpgradeParams).Execute()

Upgrade Systemd



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
    systemdUpgradeParams := *openapiclient.NewSystemdUpgradeParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, false, int32(123), int32(123), "UpgradeOption_example") // SystemdUpgradeParams | Systemd Upgrade Params

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.UpgradeSystemd(context.Background(), cUUID, uniUUID).SystemdUpgradeParams(systemdUpgradeParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseUpgradesManagementApi.UpgradeSystemd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeSystemd`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseUpgradesManagementApi.UpgradeSystemd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeSystemdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **systemdUpgradeParams** | [**SystemdUpgradeParams**](SystemdUpgradeParams.md) | Systemd Upgrade Params | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeTls

> YBPTask UpgradeTls(ctx, cUUID, uniUUID).TlsToggleParams(tlsToggleParams).Execute()

Upgrade TLS



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
    tlsToggleParams := *openapiclient.NewTlsToggleParams(false, "ClientRootCA_example", []openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, false, false, false, false, "RootCA_example", int32(123), int32(123), "UpgradeOption_example") // TlsToggleParams | TLS Toggle Params

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.UpgradeTls(context.Background(), cUUID, uniUUID).TlsToggleParams(tlsToggleParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseUpgradesManagementApi.UpgradeTls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeTls`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseUpgradesManagementApi.UpgradeTls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeTlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tlsToggleParams** | [**TlsToggleParams**](TlsToggleParams.md) | TLS Toggle Params | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeVMImage

> YBPTask UpgradeVMImage(ctx, cUUID, uniUUID).VmimageUpgradeParams(vmimageUpgradeParams).Execute()

Upgrade VM Image



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
    vmimageUpgradeParams := *openapiclient.NewVMImageUpgradeParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, false, false, map[string]string{"key": "Inner_example"}, int32(123), int32(123), "UpgradeOption_example") // VMImageUpgradeParams | VM Image Upgrade Params

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.UpgradeVMImage(context.Background(), cUUID, uniUUID).VmimageUpgradeParams(vmimageUpgradeParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseUpgradesManagementApi.UpgradeVMImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeVMImage`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseUpgradesManagementApi.UpgradeVMImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeVMImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vmimageUpgradeParams** | [**VMImageUpgradeParams**](VMImageUpgradeParams.md) | VM Image Upgrade Params | 

### Return type

[**YBPTask**](YBPTask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

