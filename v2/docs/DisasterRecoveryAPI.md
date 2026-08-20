# \DisasterRecoveryAPI

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PageListDrConfigDatabases**](DisasterRecoveryAPI.md#PageListDrConfigDatabases) | **Post** /customers/{cUUID}/dr-configs/{drUUID}/databases/page | List DR config database replication details.
[**PageListDrConfigTables**](DisasterRecoveryAPI.md#PageListDrConfigTables) | **Post** /customers/{cUUID}/dr-configs/{drUUID}/tables/page | List DR config table replication details.
[**PageListDrConfigs**](DisasterRecoveryAPI.md#PageListDrConfigs) | **Post** /customers/{cUUID}/universes/{uniUUID}/dr-configs/page | List disaster recovery configurations.



## PageListDrConfigDatabases

> DrConfigDbDetailPagedResp PageListDrConfigDatabases(ctx, cUUID, drUUID).DrConfigDbDetailPagedQuerySpec(drConfigDbDetailPagedQuerySpec).Execute()

List DR config database replication details.



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
	drUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | DR config UUID
	drConfigDbDetailPagedQuerySpec := *openapiclient.NewDrConfigDbDetailPagedQuerySpec() // DrConfigDbDetailPagedQuerySpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.PageListDrConfigDatabases(context.Background(), cUUID, drUUID).DrConfigDbDetailPagedQuerySpec(drConfigDbDetailPagedQuerySpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.PageListDrConfigDatabases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageListDrConfigDatabases`: DrConfigDbDetailPagedResp
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.PageListDrConfigDatabases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**drUUID** | **string** | DR config UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageListDrConfigDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **drConfigDbDetailPagedQuerySpec** | [**DrConfigDbDetailPagedQuerySpec**](DrConfigDbDetailPagedQuerySpec.md) |  | 

### Return type

[**DrConfigDbDetailPagedResp**](DrConfigDbDetailPagedResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageListDrConfigTables

> DrConfigTableDetailPagedResp PageListDrConfigTables(ctx, cUUID, drUUID).DrConfigTableDetailPagedQuerySpec(drConfigTableDetailPagedQuerySpec).Execute()

List DR config table replication details.



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
	drUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | DR config UUID
	drConfigTableDetailPagedQuerySpec := *openapiclient.NewDrConfigTableDetailPagedQuerySpec() // DrConfigTableDetailPagedQuerySpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.PageListDrConfigTables(context.Background(), cUUID, drUUID).DrConfigTableDetailPagedQuerySpec(drConfigTableDetailPagedQuerySpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.PageListDrConfigTables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageListDrConfigTables`: DrConfigTableDetailPagedResp
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.PageListDrConfigTables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**drUUID** | **string** | DR config UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageListDrConfigTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **drConfigTableDetailPagedQuerySpec** | [**DrConfigTableDetailPagedQuerySpec**](DrConfigTableDetailPagedQuerySpec.md) |  | 

### Return type

[**DrConfigTableDetailPagedResp**](DrConfigTableDetailPagedResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageListDrConfigs

> DrConfigPagedResp PageListDrConfigs(ctx, cUUID, uniUUID).DrConfigPagedQuerySpec(drConfigPagedQuerySpec).Execute()

List disaster recovery configurations.



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
	drConfigPagedQuerySpec := *openapiclient.NewDrConfigPagedQuerySpec() // DrConfigPagedQuerySpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisasterRecoveryAPI.PageListDrConfigs(context.Background(), cUUID, uniUUID).DrConfigPagedQuerySpec(drConfigPagedQuerySpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisasterRecoveryAPI.PageListDrConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageListDrConfigs`: DrConfigPagedResp
	fmt.Fprintf(os.Stdout, "Response from `DisasterRecoveryAPI.PageListDrConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** | Customer UUID | 
**uniUUID** | **string** | Universe UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageListDrConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **drConfigPagedQuerySpec** | [**DrConfigPagedQuerySpec**](DrConfigPagedQuerySpec.md) |  | 

### Return type

[**DrConfigPagedResp**](DrConfigPagedResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

