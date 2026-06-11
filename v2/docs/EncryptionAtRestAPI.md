# \EncryptionAtRestAPI

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PageListKmsConfigs**](EncryptionAtRestAPI.md#PageListKmsConfigs) | **Post** /customers/{cUUID}/kms-configs/page | List KMS configurations (paged)



## PageListKmsConfigs

> KmsConfigPagedResp PageListKmsConfigs(ctx, cUUID).KmsConfigPagedQuerySpec(kmsConfigPagedQuerySpec).Execute()

List KMS configurations (paged)



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
	kmsConfigPagedQuerySpec := *openapiclient.NewKmsConfigPagedQuerySpec() // KmsConfigPagedQuerySpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EncryptionAtRestAPI.PageListKmsConfigs(context.Background(), cUUID).KmsConfigPagedQuerySpec(kmsConfigPagedQuerySpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncryptionAtRestAPI.PageListKmsConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageListKmsConfigs`: KmsConfigPagedResp
	fmt.Fprintf(os.Stdout, "Response from `EncryptionAtRestAPI.PageListKmsConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageListKmsConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsConfigPagedQuerySpec** | [**KmsConfigPagedQuerySpec**](KmsConfigPagedQuerySpec.md) |  | 

### Return type

[**KmsConfigPagedResp**](KmsConfigPagedResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

