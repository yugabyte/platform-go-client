# \NodeInstancesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNodeInstance**](NodeInstancesApi.md#CreateNodeInstance) | **Post** /api/v1/customers/{cUUID}/zones/{azUUID}/nodes | Create a node instance
[**DeleteInstance**](NodeInstancesApi.md#DeleteInstance) | **Delete** /api/v1/customers/{cUUID}/providers/{pUUID}/instances/{instanceIP} | Delete a node instance
[**DetachedNodeAction**](NodeInstancesApi.md#DetachedNodeAction) | **Post** /api/v1/customers/{cUUID}/providers/{pUUID}/instances/{instanceIP} | Detached node action
[**GetNodeDetails**](NodeInstancesApi.md#GetNodeDetails) | **Get** /api/v1/customers/{cUUID}/universes/{universeUUID}/nodes/{nodeName}/details | Get node details
[**GetNodeInstance**](NodeInstancesApi.md#GetNodeInstance) | **Get** /api/v1/customers/{cUUID}/nodes/{nodeUUID}/list | Get a node instance
[**ListByProvider**](NodeInstancesApi.md#ListByProvider) | **Get** /api/v1/customers/{cUUID}/providers/{pUUID}/nodes/list | List all of a provider&#39;s node instances
[**ListByZone**](NodeInstancesApi.md#ListByZone) | **Get** /api/v1/customers/{cUUID}/zones/{azUUID}/nodes/list | List all of a zone&#39;s node instances
[**NodeAction**](NodeInstancesApi.md#NodeAction) | **Put** /api/v1/customers/{cUUID}/universes/{universeUUID}/nodes/{nodeName} | Update a node
[**ValidateNodeInstance**](NodeInstancesApi.md#ValidateNodeInstance) | **Post** /api/v1/customers/{cUUID}/zones/{azUUID}/nodes/validate | Validate a node instance



## CreateNodeInstance

> map[string]NodeInstance CreateNodeInstance(ctx, cUUID, azUUID).NodeInstance(nodeInstance).Execute()

Create a node instance

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
    azUUID := TODO // string | 
    nodeInstance := *openapiclient.NewNodeInstanceFormData([]openapiclient.NodeInstanceData{*openapiclient.NewNodeInstanceData("Mumbai instance", "c5large", "1.1.1.1", "south-east", "centos", "south-east")}) // NodeInstanceFormData | Node instance data to be created

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeInstancesApi.CreateNodeInstance(context.Background(), cUUID, azUUID).NodeInstance(nodeInstance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeInstancesApi.CreateNodeInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNodeInstance`: map[string]NodeInstance
    fmt.Fprintf(os.Stdout, "Response from `NodeInstancesApi.CreateNodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**azUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nodeInstance** | [**NodeInstanceFormData**](NodeInstanceFormData.md) | Node instance data to be created | 

### Return type

[**map[string]NodeInstance**](NodeInstance.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstance

> YBPSuccess DeleteInstance(ctx, cUUID, pUUID, instanceIP).Execute()

Delete a node instance

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
    pUUID := TODO // string | 
    instanceIP := "instanceIP_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeInstancesApi.DeleteInstance(context.Background(), cUUID, pUUID, instanceIP).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeInstancesApi.DeleteInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInstance`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `NodeInstancesApi.DeleteInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 
**instanceIP** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceRequest struct via the builder pattern


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


## DetachedNodeAction

> YBPTask DetachedNodeAction(ctx, cUUID, pUUID, instanceIP).Execute()

Detached node action

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
    pUUID := TODO // string | 
    instanceIP := "instanceIP_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeInstancesApi.DetachedNodeAction(context.Background(), cUUID, pUUID, instanceIP).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeInstancesApi.DetachedNodeAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachedNodeAction`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `NodeInstancesApi.DetachedNodeAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 
**instanceIP** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachedNodeActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## GetNodeDetails

> NodeDetailsResp GetNodeDetails(ctx, cUUID, universeUUID, nodeName).Execute()

Get node details

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
    universeUUID := TODO // string | 
    nodeName := "nodeName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeInstancesApi.GetNodeDetails(context.Background(), cUUID, universeUUID, nodeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeInstancesApi.GetNodeDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeDetails`: NodeDetailsResp
    fmt.Fprintf(os.Stdout, "Response from `NodeInstancesApi.GetNodeDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**universeUUID** | [**string**](.md) |  | 
**nodeName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**NodeDetailsResp**](NodeDetailsResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeInstance

> NodeInstance GetNodeInstance(ctx, cUUID, nodeUUID).Execute()

Get a node instance

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
    nodeUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeInstancesApi.GetNodeInstance(context.Background(), cUUID, nodeUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeInstancesApi.GetNodeInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeInstance`: NodeInstance
    fmt.Fprintf(os.Stdout, "Response from `NodeInstancesApi.GetNodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**nodeUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NodeInstance**](NodeInstance.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListByProvider

> []NodeInstance ListByProvider(ctx, cUUID, pUUID).Execute()

List all of a provider's node instances

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
    pUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeInstancesApi.ListByProvider(context.Background(), cUUID, pUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeInstancesApi.ListByProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListByProvider`: []NodeInstance
    fmt.Fprintf(os.Stdout, "Response from `NodeInstancesApi.ListByProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListByProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]NodeInstance**](NodeInstance.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListByZone

> []NodeInstance ListByZone(ctx, cUUID, azUUID).Execute()

List all of a zone's node instances

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
    azUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeInstancesApi.ListByZone(context.Background(), cUUID, azUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeInstancesApi.ListByZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListByZone`: []NodeInstance
    fmt.Fprintf(os.Stdout, "Response from `NodeInstancesApi.ListByZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**azUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListByZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]NodeInstance**](NodeInstance.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NodeAction

> YBPTask NodeAction(ctx, cUUID, universeUUID, nodeName).NodeAction(nodeAction).Execute()

Update a node

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
    universeUUID := TODO // string | 
    nodeName := "nodeName_example" // string | 
    nodeAction := *openapiclient.NewNodeActionFormData("NodeAction_example") // NodeActionFormData | Node action data to be updated

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeInstancesApi.NodeAction(context.Background(), cUUID, universeUUID, nodeName).NodeAction(nodeAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeInstancesApi.NodeAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NodeAction`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `NodeInstancesApi.NodeAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**universeUUID** | [**string**](.md) |  | 
**nodeName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNodeActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **nodeAction** | [**NodeActionFormData**](NodeActionFormData.md) | Node action data to be updated | 

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


## ValidateNodeInstance

> map[string]ValidationResult ValidateNodeInstance(ctx, cUUID, azUUID).Execute()

Validate a node instance

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
    azUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NodeInstancesApi.ValidateNodeInstance(context.Background(), cUUID, azUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodeInstancesApi.ValidateNodeInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateNodeInstance`: map[string]ValidationResult
    fmt.Fprintf(os.Stdout, "Response from `NodeInstancesApi.ValidateNodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**azUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateNodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**map[string]ValidationResult**](ValidationResult.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

