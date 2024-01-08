# \UniverseDatabaseManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigureYCQL**](UniverseDatabaseManagementApi.md#ConfigureYCQL) | **Post** /api/v1/customers/{cUUID}/universes/{univUUID}/configure/ycql | WARNING: This is a preview API that could change. Configure YCQL
[**ConfigureYSQL**](UniverseDatabaseManagementApi.md#ConfigureYSQL) | **Post** /api/v1/customers/{cUUID}/universes/{univUUID}/configure/ysql | WARNING: This is a preview API that could change. Configure YSQL



## ConfigureYCQL

> YBPTask ConfigureYCQL(ctx, cUUID, univUUID).ConfigureYcqlFormData(configureYcqlFormData).Request(request).Execute()

WARNING: This is a preview API that could change. Configure YCQL



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
    univUUID := TODO // string | 
    configureYcqlFormData := *openapiclient.NewConfigureYCQLFormData() // ConfigureYCQLFormData | Configure YCQL Form Data
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseDatabaseManagementApi.ConfigureYCQL(context.Background(), cUUID, univUUID).ConfigureYcqlFormData(configureYcqlFormData).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseDatabaseManagementApi.ConfigureYCQL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfigureYCQL`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseDatabaseManagementApi.ConfigureYCQL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**univUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigureYCQLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **configureYcqlFormData** | [**ConfigureYCQLFormData**](ConfigureYCQLFormData.md) | Configure YCQL Form Data | 
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


## ConfigureYSQL

> YBPTask ConfigureYSQL(ctx, cUUID, univUUID).ConfigureYsqlFormData(configureYsqlFormData).Request(request).Execute()

WARNING: This is a preview API that could change. Configure YSQL



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
    univUUID := TODO // string | 
    configureYsqlFormData := *openapiclient.NewConfigureYSQLFormData() // ConfigureYSQLFormData | Configure YSQL Form Data
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseDatabaseManagementApi.ConfigureYSQL(context.Background(), cUUID, univUUID).ConfigureYsqlFormData(configureYsqlFormData).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseDatabaseManagementApi.ConfigureYSQL``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfigureYSQL`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `UniverseDatabaseManagementApi.ConfigureYSQL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**univUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigureYSQLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **configureYsqlFormData** | [**ConfigureYSQLFormData**](ConfigureYSQLFormData.md) | Configure YSQL Form Data | 
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

