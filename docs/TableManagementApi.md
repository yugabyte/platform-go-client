# \TableManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DescribeTable**](TableManagementApi.md#DescribeTable) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/tables/{tableUUID} | Describe a table
[**GetAllNamespaces**](TableManagementApi.md#GetAllNamespaces) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/namespaces | Get a list of all namespaces in the specified universe.
[**GetAllTableSpaces**](TableManagementApi.md#GetAllTableSpaces) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/tablespaces | Get a list of all tablespaces of a given universe.
[**GetAllTables**](TableManagementApi.md#GetAllTables) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/tables | Get a list of all tables in the specified universe



## DescribeTable

> TableDefinitionTaskParams DescribeTable(ctx, cUUID, uniUUID, tableUUID).Execute()

Describe a table



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
    tableUUID := "tableUUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TableManagementApi.DescribeTable(context.Background(), cUUID, uniUUID, tableUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableManagementApi.DescribeTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DescribeTable`: TableDefinitionTaskParams
    fmt.Fprintf(os.Stdout, "Response from `TableManagementApi.DescribeTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 
**tableUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDescribeTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TableDefinitionTaskParams**](TableDefinitionTaskParams.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllNamespaces

> []NamespaceInfoResp GetAllNamespaces(ctx, cUUID, uniUUID).IncludeSystemNamespaces(includeSystemNamespaces).Execute()

Get a list of all namespaces in the specified universe.



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
    includeSystemNamespaces := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TableManagementApi.GetAllNamespaces(context.Background(), cUUID, uniUUID).IncludeSystemNamespaces(includeSystemNamespaces).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableManagementApi.GetAllNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllNamespaces`: []NamespaceInfoResp
    fmt.Fprintf(os.Stdout, "Response from `TableManagementApi.GetAllNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeSystemNamespaces** | **bool** |  | [default to false]

### Return type

[**[]NamespaceInfoResp**](NamespaceInfoResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllTableSpaces

> []TableSpaceInfo GetAllTableSpaces(ctx, cUUID, uniUUID).Execute()

Get a list of all tablespaces of a given universe.



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
    resp, r, err := api_client.TableManagementApi.GetAllTableSpaces(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableManagementApi.GetAllTableSpaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllTableSpaces`: []TableSpaceInfo
    fmt.Fprintf(os.Stdout, "Response from `TableManagementApi.GetAllTableSpaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTableSpacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]TableSpaceInfo**](TableSpaceInfo.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllTables

> []TableInfoResp GetAllTables(ctx, cUUID, uniUUID).IncludeParentTableInfo(includeParentTableInfo).ExcludeColocatedTables(excludeColocatedTables).IncludeColocatedParentTables(includeColocatedParentTables).XClusterSupportedOnly(xClusterSupportedOnly).Execute()

Get a list of all tables in the specified universe



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
    includeParentTableInfo := true // bool |  (optional) (default to false)
    excludeColocatedTables := true // bool |  (optional) (default to false)
    includeColocatedParentTables := true // bool |  (optional) (default to true)
    xClusterSupportedOnly := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TableManagementApi.GetAllTables(context.Background(), cUUID, uniUUID).IncludeParentTableInfo(includeParentTableInfo).ExcludeColocatedTables(excludeColocatedTables).IncludeColocatedParentTables(includeColocatedParentTables).XClusterSupportedOnly(xClusterSupportedOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableManagementApi.GetAllTables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllTables`: []TableInfoResp
    fmt.Fprintf(os.Stdout, "Response from `TableManagementApi.GetAllTables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeParentTableInfo** | **bool** |  | [default to false]
 **excludeColocatedTables** | **bool** |  | [default to false]
 **includeColocatedParentTables** | **bool** |  | [default to true]
 **xClusterSupportedOnly** | **bool** |  | [default to false]

### Return type

[**[]TableInfoResp**](TableInfoResp.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

