# \UniverseUpgradesManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FinalizeUpgrade**](UniverseUpgradesManagementApi.md#FinalizeUpgrade) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/finalize | Finalize Upgrade
[**PreFinalizeSoftwareUpgradeInfo**](UniverseUpgradesManagementApi.md#PreFinalizeSoftwareUpgradeInfo) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/finalize/info | Finalize Software Upgrade info
[**PreUpgradeValidation**](UniverseUpgradesManagementApi.md#PreUpgradeValidation) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/pre_upgrade_validation | Run validation for upgrade
[**RebootUniverse**](UniverseUpgradesManagementApi.md#RebootUniverse) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/reboot | Reboot universe
[**ResizeNode**](UniverseUpgradesManagementApi.md#ResizeNode) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/resize_node | Resize Node
[**RestartUniverse**](UniverseUpgradesManagementApi.md#RestartUniverse) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/restart | Restart Universe
[**RollbackUpgrade**](UniverseUpgradesManagementApi.md#RollbackUpgrade) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/rollback | Rollback Upgrade
[**SoftwareUpgradePreCheck**](UniverseUpgradesManagementApi.md#SoftwareUpgradePreCheck) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/software/precheck | Software Upgrade universe pre-check
[**UpdateProxyConfig**](UniverseUpgradesManagementApi.md#UpdateProxyConfig) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/proxy_config | Update Proxy Config
[**UpgradeCerts**](UniverseUpgradesManagementApi.md#UpgradeCerts) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/certs | Upgrade Certs
[**UpgradeDBVersion**](UniverseUpgradesManagementApi.md#UpgradeDBVersion) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/db_version | Upgrade DB version
[**UpgradeGFlags**](UniverseUpgradesManagementApi.md#UpgradeGFlags) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/gflags | Upgrade GFlags
[**UpgradeKubernetesOverrides**](UniverseUpgradesManagementApi.md#UpgradeKubernetesOverrides) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/kubernetes_overrides | Upgrade KubernetesOverrides
[**UpgradeSoftware**](UniverseUpgradesManagementApi.md#UpgradeSoftware) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/software | Upgrade Software
[**UpgradeSystemd**](UniverseUpgradesManagementApi.md#UpgradeSystemd) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/systemd | Upgrade Systemd
[**UpgradeThirdpartySoftware**](UniverseUpgradesManagementApi.md#UpgradeThirdpartySoftware) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/thirdparty_software | Upgrade third-party software
[**UpgradeTls**](UniverseUpgradesManagementApi.md#UpgradeTls) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/tls | Upgrade TLS
[**UpgradeVMImage**](UniverseUpgradesManagementApi.md#UpgradeVMImage) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/upgrade/vm | Upgrade VM Image



## FinalizeUpgrade

> YBPTask FinalizeUpgrade(ctx, cUUID, uniUUID).FinalizeUpgradeParams(finalizeUpgradeParams).Request(request).Execute()

Finalize Upgrade



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
    finalizeUpgradeParams := *openapiclient.NewFinalizeUpgradeParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), false, "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123), "UpgradeOption_example", false) // FinalizeUpgradeParams | Finalize Upgrade Params
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.FinalizeUpgrade(context.Background(), cUUID, uniUUID).FinalizeUpgradeParams(finalizeUpgradeParams).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseUpgradesManagementApi.FinalizeUpgrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FinalizeUpgrade`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseUpgradesManagementApi.FinalizeUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFinalizeUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **finalizeUpgradeParams** | [**FinalizeUpgradeParams**](FinalizeUpgradeParams.md) | Finalize Upgrade Params | 
 **request** | [**interface{}**](interface{}.md) |  | 

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


## PreFinalizeSoftwareUpgradeInfo

> FinalizeUpgradeInfoResponse PreFinalizeSoftwareUpgradeInfo(ctx, cUUID, uniUUID).Execute()

Finalize Software Upgrade info



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
    resp, r, err := api_client.UniverseUpgradesManagementApi.PreFinalizeSoftwareUpgradeInfo(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseUpgradesManagementApi.PreFinalizeSoftwareUpgradeInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreFinalizeSoftwareUpgradeInfo`: FinalizeUpgradeInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `UniverseUpgradesManagementApi.PreFinalizeSoftwareUpgradeInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreFinalizeSoftwareUpgradeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FinalizeUpgradeInfoResponse**](FinalizeUpgradeInfoResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreUpgradeValidation

> PreUpgradeValidationResponse PreUpgradeValidation(ctx, cUUID, uniUUID).UpgradeParams(upgradeParams).UpgradeType(upgradeType).Request(request).Execute()

Run validation for upgrade



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
    upgradeParams := *openapiclient.NewUpgradeTaskParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), false, "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123), "UpgradeOption_example") // UpgradeTaskParams | Upgrade Params
    upgradeType := "upgradeType_example" // string |  (optional)
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.PreUpgradeValidation(context.Background(), cUUID, uniUUID).UpgradeParams(upgradeParams).UpgradeType(upgradeType).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseUpgradesManagementApi.PreUpgradeValidation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreUpgradeValidation`: PreUpgradeValidationResponse
    fmt.Fprintf(os.Stdout, "Response from `UniverseUpgradesManagementApi.PreUpgradeValidation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreUpgradeValidationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upgradeParams** | [**UpgradeTaskParams**](UpgradeTaskParams.md) | Upgrade Params | 
 **upgradeType** | **string** |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**PreUpgradeValidationResponse**](PreUpgradeValidationResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootUniverse

> YBPTask RebootUniverse(ctx, cUUID, uniUUID).UpgradeTaskParams(upgradeTaskParams).Request(request).Execute()

Reboot universe



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
    upgradeTaskParams := *openapiclient.NewUpgradeTaskParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), false, "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123), "UpgradeOption_example") // UpgradeTaskParams | Upgrade Task Params
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.RebootUniverse(context.Background(), cUUID, uniUUID).UpgradeTaskParams(upgradeTaskParams).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseUpgradesManagementApi.RebootUniverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RebootUniverse`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseUpgradesManagementApi.RebootUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootUniverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upgradeTaskParams** | [**UpgradeTaskParams**](UpgradeTaskParams.md) | Upgrade Task Params | 
 **request** | [**interface{}**](interface{}.md) |  | 

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


## ResizeNode

> YBPTask ResizeNode(ctx, cUUID, uniUUID).ResizeNodeParams(resizeNodeParams).Request(request).Execute()

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
    resizeNodeParams := *openapiclient.NewResizeNodeParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), false, false, map[string]string{"key": "Inner_example"}, "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123), map[string]string{"key": "Inner_example"}, "UpgradeOption_example") // ResizeNodeParams | Resize Node Params
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.ResizeNode(context.Background(), cUUID, uniUUID).ResizeNodeParams(resizeNodeParams).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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

> YBPTask RestartUniverse(ctx, cUUID, uniUUID).RestartTaskParams(restartTaskParams).Request(request).Execute()

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
    restartTaskParams := *openapiclient.NewRestartTaskParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), false, "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123), "UpgradeOption_example") // RestartTaskParams | Restart Task Params
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.RestartUniverse(context.Background(), cUUID, uniUUID).RestartTaskParams(restartTaskParams).Request(request).Execute()
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


 **restartTaskParams** | [**RestartTaskParams**](RestartTaskParams.md) | Restart Task Params | 
 **request** | [**interface{}**](interface{}.md) |  | 

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


## RollbackUpgrade

> YBPTask RollbackUpgrade(ctx, cUUID, uniUUID).RollbackUpgradeParams(rollbackUpgradeParams).Request(request).Execute()

Rollback Upgrade



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
    rollbackUpgradeParams := *openapiclient.NewRollbackUpgradeParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), false, "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123), "UpgradeOption_example") // RollbackUpgradeParams | RollBack Upgrade Params
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.RollbackUpgrade(context.Background(), cUUID, uniUUID).RollbackUpgradeParams(rollbackUpgradeParams).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseUpgradesManagementApi.RollbackUpgrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RollbackUpgrade`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseUpgradesManagementApi.RollbackUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rollbackUpgradeParams** | [**RollbackUpgradeParams**](RollbackUpgradeParams.md) | RollBack Upgrade Params | 
 **request** | [**interface{}**](interface{}.md) |  | 

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


## SoftwareUpgradePreCheck

> SoftwareUpgradeInfoResponse SoftwareUpgradePreCheck(ctx, cUUID, uniUUID).SoftwareUpgradeInfoRequest(softwareUpgradeInfoRequest).Request(request).Execute()

Software Upgrade universe pre-check



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
    softwareUpgradeInfoRequest := *openapiclient.NewSoftwareUpgradeInfoRequest("YbSoftwareVersion_example") // SoftwareUpgradeInfoRequest | Software Upgrade Info Request
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.SoftwareUpgradePreCheck(context.Background(), cUUID, uniUUID).SoftwareUpgradeInfoRequest(softwareUpgradeInfoRequest).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseUpgradesManagementApi.SoftwareUpgradePreCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SoftwareUpgradePreCheck`: SoftwareUpgradeInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `UniverseUpgradesManagementApi.SoftwareUpgradePreCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSoftwareUpgradePreCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **softwareUpgradeInfoRequest** | [**SoftwareUpgradeInfoRequest**](SoftwareUpgradeInfoRequest.md) | Software Upgrade Info Request | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**SoftwareUpgradeInfoResponse**](SoftwareUpgradeInfoResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProxyConfig

> YBPTask UpdateProxyConfig(ctx, cUUID, uniUUID).UpdateProxyConfigParams(updateProxyConfigParams).Request(request).Execute()

Update Proxy Config



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
    updateProxyConfigParams := *openapiclient.NewProxyConfigUpdateParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), false, "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123), "UpgradeOption_example") // ProxyConfigUpdateParams | Update Proxy Config Params
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.UpdateProxyConfig(context.Background(), cUUID, uniUUID).UpdateProxyConfigParams(updateProxyConfigParams).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseUpgradesManagementApi.UpdateProxyConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProxyConfig`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseUpgradesManagementApi.UpdateProxyConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProxyConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateProxyConfigParams** | [**ProxyConfigUpdateParams**](ProxyConfigUpdateParams.md) | Update Proxy Config Params | 
 **request** | [**interface{}**](interface{}.md) |  | 

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

> YBPTask UpgradeCerts(ctx, cUUID, uniUUID).CertsRotateParams(certsRotateParams).Request(request).Execute()

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
    certsRotateParams := *openapiclient.NewCertsRotateParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), false, "PlatformUrl_example", "PlatformVersion_example", false, false, int32(123), int32(123), "UpgradeOption_example") // CertsRotateParams | Certs Rotate Params
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.UpgradeCerts(context.Background(), cUUID, uniUUID).CertsRotateParams(certsRotateParams).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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


## UpgradeDBVersion

> YBPTask UpgradeDBVersion(ctx, cUUID, uniUUID).SoftwareUpgradeParams(softwareUpgradeParams).Request(request).Execute()

Upgrade DB version



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
    softwareUpgradeParams := *openapiclient.NewSoftwareUpgradeParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), false, "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123), "UpgradeOption_example", false, "YbSoftwareVersion_example") // SoftwareUpgradeParams | Software Upgrade Params
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.UpgradeDBVersion(context.Background(), cUUID, uniUUID).SoftwareUpgradeParams(softwareUpgradeParams).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseUpgradesManagementApi.UpgradeDBVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeDBVersion`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseUpgradesManagementApi.UpgradeDBVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeDBVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **softwareUpgradeParams** | [**SoftwareUpgradeParams**](SoftwareUpgradeParams.md) | Software Upgrade Params | 
 **request** | [**interface{}**](interface{}.md) |  | 

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

> YBPTask UpgradeGFlags(ctx, cUUID, uniUUID).GflagsUpgradeParams(gflagsUpgradeParams).Request(request).Execute()

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
    gflagsUpgradeParams := *openapiclient.NewGFlagsUpgradeParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), false, map[string]string{"key": "Inner_example"}, "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123), map[string]string{"key": "Inner_example"}, "UpgradeOption_example") // GFlagsUpgradeParams | GFlags Upgrade Params
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.UpgradeGFlags(context.Background(), cUUID, uniUUID).GflagsUpgradeParams(gflagsUpgradeParams).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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


## UpgradeKubernetesOverrides

> YBPTask UpgradeKubernetesOverrides(ctx, cUUID, uniUUID).KubernetesOverridesUpgradeParams(kubernetesOverridesUpgradeParams).Request(request).Execute()

Upgrade KubernetesOverrides



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
    kubernetesOverridesUpgradeParams := *openapiclient.NewKubernetesOverridesUpgradeParams(map[string]string{"key": "Inner_example"}, []openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), false, "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123), "UniverseOverrides_example", "UpgradeOption_example") // KubernetesOverridesUpgradeParams | Kubernetes Override Upgrade Params
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.UpgradeKubernetesOverrides(context.Background(), cUUID, uniUUID).KubernetesOverridesUpgradeParams(kubernetesOverridesUpgradeParams).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseUpgradesManagementApi.UpgradeKubernetesOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeKubernetesOverrides`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseUpgradesManagementApi.UpgradeKubernetesOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeKubernetesOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kubernetesOverridesUpgradeParams** | [**KubernetesOverridesUpgradeParams**](KubernetesOverridesUpgradeParams.md) | Kubernetes Override Upgrade Params | 
 **request** | [**interface{}**](interface{}.md) |  | 

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

> YBPTask UpgradeSoftware(ctx, cUUID, uniUUID).SoftwareUpgradeParams(softwareUpgradeParams).Request(request).Execute()

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
    softwareUpgradeParams := *openapiclient.NewSoftwareUpgradeParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), false, "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123), "UpgradeOption_example", false, "YbSoftwareVersion_example") // SoftwareUpgradeParams | Software Upgrade Params
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.UpgradeSoftware(context.Background(), cUUID, uniUUID).SoftwareUpgradeParams(softwareUpgradeParams).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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

> YBPTask UpgradeSystemd(ctx, cUUID, uniUUID).SystemdUpgradeParams(systemdUpgradeParams).Request(request).Execute()

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
    systemdUpgradeParams := *openapiclient.NewSystemdUpgradeParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), false, "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123), "UpgradeOption_example") // SystemdUpgradeParams | Systemd Upgrade Params
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.UpgradeSystemd(context.Background(), cUUID, uniUUID).SystemdUpgradeParams(systemdUpgradeParams).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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


## UpgradeThirdpartySoftware

> YBPTask UpgradeThirdpartySoftware(ctx, cUUID, uniUUID).ThirdpartySoftwareUpgradeParams(thirdpartySoftwareUpgradeParams).Request(request).Execute()

Upgrade third-party software



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
    thirdpartySoftwareUpgradeParams := *openapiclient.NewThirdpartySoftwareUpgradeParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), false, false, "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123), "UpgradeOption_example") // ThirdpartySoftwareUpgradeParams | Thirdparty Software Upgrade Params
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.UpgradeThirdpartySoftware(context.Background(), cUUID, uniUUID).ThirdpartySoftwareUpgradeParams(thirdpartySoftwareUpgradeParams).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseUpgradesManagementApi.UpgradeThirdpartySoftware``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeThirdpartySoftware`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseUpgradesManagementApi.UpgradeThirdpartySoftware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeThirdpartySoftwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **thirdpartySoftwareUpgradeParams** | [**ThirdpartySoftwareUpgradeParams**](ThirdpartySoftwareUpgradeParams.md) | Thirdparty Software Upgrade Params | 
 **request** | [**interface{}**](interface{}.md) |  | 

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

> YBPTask UpgradeTls(ctx, cUUID, uniUUID).TlsToggleParams(tlsToggleParams).Request(request).Execute()

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
    tlsToggleParams := *openapiclient.NewTlsToggleParams(false, []openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), false, false, false, "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123), "UpgradeOption_example") // TlsToggleParams | TLS Toggle Params
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.UpgradeTls(context.Background(), cUUID, uniUUID).TlsToggleParams(tlsToggleParams).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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

> YBPTask UpgradeVMImage(ctx, cUUID, uniUUID).VmimageUpgradeParams(vmimageUpgradeParams).Request(request).Execute()

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
    vmimageUpgradeParams := *openapiclient.NewVMImageUpgradeParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, *openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), false, false, "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123), "UpgradeOption_example", "YbSoftwareVersion_example") // VMImageUpgradeParams | VM Image Upgrade Params
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseUpgradesManagementApi.UpgradeVMImage(context.Background(), cUUID, uniUUID).VmimageUpgradeParams(vmimageUpgradeParams).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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

