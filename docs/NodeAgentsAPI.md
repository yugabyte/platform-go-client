# \NodeAgentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadNodeAgentInstaller**](NodeAgentsAPI.md#DownloadNodeAgentInstaller) | **Get** /api/v1/node_agents/download | Download Node Agent Installer or Package
[**GetNodeAgent**](NodeAgentsAPI.md#GetNodeAgent) | **Get** /api/v1/customers/{cUUID}/node_agents/{nUUID} | Get Node Agent
[**ListNodeAgents**](NodeAgentsAPI.md#ListNodeAgents) | **Get** /api/v1/customers/{cUUID}/node_agents | List Node Agents
[**PageListNodeAgents**](NodeAgentsAPI.md#PageListNodeAgents) | **Post** /api/v1/customers/{cUUID}/node_agents/page | List Node Agents (paginated)
[**ReinstallNodeAgent**](NodeAgentsAPI.md#ReinstallNodeAgent) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/node_agents | Reinstall Node Agent



## DownloadNodeAgentInstaller

> string DownloadNodeAgentInstaller(ctx).DownloadType(downloadType).Os(os).Arch(arch).Execute()

Download Node Agent Installer or Package

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
	downloadType := "downloadType_example" // string |  (optional) (default to "INSTALLER")
	os := "os_example" // string |  (optional)
	arch := "arch_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeAgentsAPI.DownloadNodeAgentInstaller(context.Background()).DownloadType(downloadType).Os(os).Arch(arch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAgentsAPI.DownloadNodeAgentInstaller``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadNodeAgentInstaller`: string
	fmt.Fprintf(os.Stdout, "Response from `NodeAgentsAPI.DownloadNodeAgentInstaller`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadNodeAgentInstallerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadType** | **string** |  | [default to &quot;INSTALLER&quot;]
 **os** | **string** |  | 
 **arch** | **string** |  | 

### Return type

**string**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/gzip, application/x-sh

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeAgent

> NodeAgentResp GetNodeAgent(ctx, cUUID, nUUID).Execute()

Get Node Agent

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
	nUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeAgentsAPI.GetNodeAgent(context.Background(), cUUID, nUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAgentsAPI.GetNodeAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodeAgent`: NodeAgentResp
	fmt.Fprintf(os.Stdout, "Response from `NodeAgentsAPI.GetNodeAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**nUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NodeAgentResp**](NodeAgentResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNodeAgents

> []NodeAgentResp ListNodeAgents(ctx, cUUID).NodeIp(nodeIp).Execute()

List Node Agents

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
	nodeIp := "nodeIp_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeAgentsAPI.ListNodeAgents(context.Background(), cUUID).NodeIp(nodeIp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAgentsAPI.ListNodeAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNodeAgents`: []NodeAgentResp
	fmt.Fprintf(os.Stdout, "Response from `NodeAgentsAPI.ListNodeAgents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNodeAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeIp** | **string** |  | 

### Return type

[**[]NodeAgentResp**](NodeAgentResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageListNodeAgents

> NodeAgentPagedApiResponse PageListNodeAgents(ctx, cUUID).PageNodeAgentRequest(pageNodeAgentRequest).Request(request).Execute()

List Node Agents (paginated)

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
	pageNodeAgentRequest := *openapiclient.NewNodeAgentPagedApiQuery("Direction_example", *openapiclient.NewNodeAgentApiFilter("CloudType_example", []string{"NodeIps_example"}, "ProviderUuid_example", "RegionUuid_example", "UniverseUuid_example", "ZoneUuid_example"), int32(123), false, int32(123), "SortBy_example") // NodeAgentPagedApiQuery | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeAgentsAPI.PageListNodeAgents(context.Background(), cUUID).PageNodeAgentRequest(pageNodeAgentRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAgentsAPI.PageListNodeAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageListNodeAgents`: NodeAgentPagedApiResponse
	fmt.Fprintf(os.Stdout, "Response from `NodeAgentsAPI.PageListNodeAgents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageListNodeAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNodeAgentRequest** | [**NodeAgentPagedApiQuery**](NodeAgentPagedApiQuery.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**NodeAgentPagedApiResponse**](NodeAgentPagedApiResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReinstallNodeAgent

> NodeAgent ReinstallNodeAgent(ctx, cUUID, uniUUID).ReinstallNodeAgentForm(reinstallNodeAgentForm).Request(request).Execute()

Reinstall Node Agent

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
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	reinstallNodeAgentForm := *openapiclient.NewReinstallNodeAgentForm() // ReinstallNodeAgentForm | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeAgentsAPI.ReinstallNodeAgent(context.Background(), cUUID, uniUUID).ReinstallNodeAgentForm(reinstallNodeAgentForm).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAgentsAPI.ReinstallNodeAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReinstallNodeAgent`: NodeAgent
	fmt.Fprintf(os.Stdout, "Response from `NodeAgentsAPI.ReinstallNodeAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReinstallNodeAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reinstallNodeAgentForm** | [**ReinstallNodeAgentForm**](ReinstallNodeAgentForm.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**NodeAgent**](NodeAgent.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

