# \NodeAgentAPI

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpgradeNodeAgent**](NodeAgentAPI.md#UpgradeNodeAgent) | **Post** /customers/{cUUID}/universes/{uniUUID}/upgrade/node-agent | Upgrade node agents



## UpgradeNodeAgent

> YBATask UpgradeNodeAgent(ctx, cUUID, uniUUID).NodeAgentUpgradeSpec(nodeAgentUpgradeSpec).Execute()

Upgrade node agents



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Universe UUID
	nodeAgentUpgradeSpec := *openapiclient.NewNodeAgentUpgradeSpec() // NodeAgentUpgradeSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeAgentAPI.UpgradeNodeAgent(context.Background(), cUUID, uniUUID).NodeAgentUpgradeSpec(nodeAgentUpgradeSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeAgentAPI.UpgradeNodeAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeNodeAgent`: YBATask
	fmt.Fprintf(os.Stdout, "Response from `NodeAgentAPI.UpgradeNodeAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeNodeAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nodeAgentUpgradeSpec** | [**NodeAgentUpgradeSpec**](NodeAgentUpgradeSpec.md) |  | 

### Return type

[**YBATask**](YBATask.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

