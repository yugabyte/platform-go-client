# \PITRAPI

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PageListPitrConfigs**](PITRAPI.md#PageListPitrConfigs) | **Post** /customers/{cUUID}/universes/{uniUUID}/pitr-configs/page | List PITR configurations (paged)



## PageListPitrConfigs

> PitrConfigPagedResp PageListPitrConfigs(ctx, cUUID, uniUUID).PitrConfigPagedQuerySpec(pitrConfigPagedQuerySpec).Execute()

List PITR configurations (paged)



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
	pitrConfigPagedQuerySpec := *openapiclient.NewPitrConfigPagedQuerySpec() // PitrConfigPagedQuerySpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PITRAPI.PageListPitrConfigs(context.Background(), cUUID, uniUUID).PitrConfigPagedQuerySpec(pitrConfigPagedQuerySpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PITRAPI.PageListPitrConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageListPitrConfigs`: PitrConfigPagedResp
	fmt.Fprintf(os.Stdout, "Response from `PITRAPI.PageListPitrConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageListPitrConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pitrConfigPagedQuerySpec** | [**PitrConfigPagedQuerySpec**](PitrConfigPagedQuerySpec.md) |  | 

### Return type

[**PitrConfigPagedResp**](PitrConfigPagedResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

