# \NodeInstancesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNodeInstance**](NodeInstancesAPI.md#CreateNodeInstance) | **Post** /api/v1/customers/{cUUID}/zones/{azUUID}/nodes | Create a node instance
[**DeleteInstance**](NodeInstancesAPI.md#DeleteInstance) | **Delete** /api/v1/customers/{cUUID}/providers/{pUUID}/instances/{instanceIP} | Delete a node instance
[**DetachedNodeAction**](NodeInstancesAPI.md#DetachedNodeAction) | **Post** /api/v1/customers/{cUUID}/providers/{pUUID}/instances/{instanceIP} | Detached node action
[**GetNodeDetails**](NodeInstancesAPI.md#GetNodeDetails) | **Get** /api/v1/customers/{cUUID}/universes/{universeUUID}/nodes/{nodeName}/details | Get node details
[**GetNodeInstance**](NodeInstancesAPI.md#GetNodeInstance) | **Get** /api/v1/customers/{cUUID}/nodes/{nodeUUID}/list | Get a node instance
[**ListByProvider**](NodeInstancesAPI.md#ListByProvider) | **Get** /api/v1/customers/{cUUID}/providers/{pUUID}/nodes/list | List all of a provider&#39;s node instances
[**ListByZone**](NodeInstancesAPI.md#ListByZone) | **Get** /api/v1/customers/{cUUID}/zones/{azUUID}/nodes/list | List all of a zone&#39;s node instances
[**NodeAction**](NodeInstancesAPI.md#NodeAction) | **Put** /api/v1/customers/{cUUID}/universes/{universeUUID}/nodes/{nodeName} | Perform the specified action on the universe node
[**UpdateState**](NodeInstancesAPI.md#UpdateState) | **Put** /api/v1/customers/{cUUID}/providers/{pUUID}/instances/{instanceIP}/state | Update node instance state
[**ValidateNodeInstance**](NodeInstancesAPI.md#ValidateNodeInstance) | **Post** /api/v1/customers/{cUUID}/zones/{azUUID}/nodes/validate | Validate a node instance



## CreateNodeInstance

> map[string]NodeInstance CreateNodeInstance(ctx, cUUID, azUUID).NodeInstance(nodeInstance).Request(request).Execute()

Create a node instance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	azUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	nodeInstance := *openapiclient.NewNodeInstanceFormData([]openapiclient.NodeInstanceData{*openapiclient.NewNodeInstanceData("Mumbai instance", "c5large", "1.1.1.1", "south-east", "centos", "south-east")}) // NodeInstanceFormData | Node instance data to be created
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeInstancesAPI.CreateNodeInstance(context.Background(), cUUID, azUUID).NodeInstance(nodeInstance).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeInstancesAPI.CreateNodeInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNodeInstance`: map[string]NodeInstance
	fmt.Fprintf(os.Stdout, "Response from `NodeInstancesAPI.CreateNodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**azUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nodeInstance** | [**NodeInstanceFormData**](NodeInstanceFormData.md) | Node instance data to be created | 
 **request** | [**interface{}**](interface{}.md) |  | 

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

> YBPSuccess DeleteInstance(ctx, cUUID, pUUID, instanceIP).Request(request).Execute()

Delete a node instance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	instanceIP := "instanceIP_example" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeInstancesAPI.DeleteInstance(context.Background(), cUUID, pUUID, instanceIP).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeInstancesAPI.DeleteInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInstance`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `NodeInstancesAPI.DeleteInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**pUUID** | **string** |  | 
**instanceIP** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceRequest struct via the builder pattern


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


## DetachedNodeAction

> YBPTask DetachedNodeAction(ctx, cUUID, pUUID, instanceIP).NodeAction(nodeAction).Request(request).Execute()

Detached node action

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	instanceIP := "instanceIP_example" // string | 
	nodeAction := *openapiclient.NewNodeActionFormData("NodeAction_example") // NodeActionFormData | Node action data to be updated
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeInstancesAPI.DetachedNodeAction(context.Background(), cUUID, pUUID, instanceIP).NodeAction(nodeAction).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeInstancesAPI.DetachedNodeAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachedNodeAction`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `NodeInstancesAPI.DetachedNodeAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**pUUID** | **string** |  | 
**instanceIP** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachedNodeActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **nodeAction** | [**NodeActionFormData**](NodeActionFormData.md) | Node action data to be updated | 
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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	universeUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	nodeName := "nodeName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeInstancesAPI.GetNodeDetails(context.Background(), cUUID, universeUUID, nodeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeInstancesAPI.GetNodeDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodeDetails`: NodeDetailsResp
	fmt.Fprintf(os.Stdout, "Response from `NodeInstancesAPI.GetNodeDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**universeUUID** | **string** |  | 
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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	nodeUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeInstancesAPI.GetNodeInstance(context.Background(), cUUID, nodeUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeInstancesAPI.GetNodeInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodeInstance`: NodeInstance
	fmt.Fprintf(os.Stdout, "Response from `NodeInstancesAPI.GetNodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**nodeUUID** | **string** |  | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeInstancesAPI.ListByProvider(context.Background(), cUUID, pUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeInstancesAPI.ListByProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListByProvider`: []NodeInstance
	fmt.Fprintf(os.Stdout, "Response from `NodeInstancesAPI.ListByProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**pUUID** | **string** |  | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	azUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeInstancesAPI.ListByZone(context.Background(), cUUID, azUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeInstancesAPI.ListByZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListByZone`: []NodeInstance
	fmt.Fprintf(os.Stdout, "Response from `NodeInstancesAPI.ListByZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**azUUID** | **string** |  | 

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

> YBPTask NodeAction(ctx, cUUID, universeUUID, nodeName).NodeAction(nodeAction).Request(request).Execute()

Perform the specified action on the universe node

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	universeUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	nodeName := "nodeName_example" // string | 
	nodeAction := *openapiclient.NewNodeActionFormData("NodeAction_example") // NodeActionFormData | Details of the node action to be performed
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeInstancesAPI.NodeAction(context.Background(), cUUID, universeUUID, nodeName).NodeAction(nodeAction).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeInstancesAPI.NodeAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NodeAction`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `NodeInstancesAPI.NodeAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**universeUUID** | **string** |  | 
**nodeName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNodeActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **nodeAction** | [**NodeActionFormData**](NodeActionFormData.md) | Details of the node action to be performed | 
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


## UpdateState

> YBPTask UpdateState(ctx, cUUID, pUUID, instanceIP).NodeInstanceStateFormData(nodeInstanceStateFormData).Request(request).Execute()

Update node instance state



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	instanceIP := "instanceIP_example" // string | 
	nodeInstanceStateFormData := *openapiclient.NewNodeInstanceStateFormData() // NodeInstanceStateFormData | Resultant node instance state to transition to
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeInstancesAPI.UpdateState(context.Background(), cUUID, pUUID, instanceIP).NodeInstanceStateFormData(nodeInstanceStateFormData).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeInstancesAPI.UpdateState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateState`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `NodeInstancesAPI.UpdateState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**pUUID** | **string** |  | 
**instanceIP** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **nodeInstanceStateFormData** | [**NodeInstanceStateFormData**](NodeInstanceStateFormData.md) | Resultant node instance state to transition to | 
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


## ValidateNodeInstance

> map[string]ValidationResult ValidateNodeInstance(ctx, cUUID, azUUID).Request(request).Execute()

Validate a node instance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	azUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeInstancesAPI.ValidateNodeInstance(context.Background(), cUUID, azUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeInstancesAPI.ValidateNodeInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateNodeInstance`: map[string]ValidationResult
	fmt.Fprintf(os.Stdout, "Response from `NodeInstancesAPI.ValidateNodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**azUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateNodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

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

