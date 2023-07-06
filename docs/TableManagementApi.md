# \TableManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlterTable**](TableManagementApi.md#AlterTable) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/tables/{tableUUID} | Alter a YugabyteDB table
[**BulkImportData**](TableManagementApi.md#BulkImportData) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/tables/{tableUUID}/bulk_import | Bulk import data
[**CreateMultiTableBackup**](TableManagementApi.md#CreateMultiTableBackup) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/multi_table_backup | Create a multi-table backup
[**CreateSingleTableBackup**](TableManagementApi.md#CreateSingleTableBackup) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/tables/{tableUUID}/create_backup | Create a single-table backup
[**CreateTable**](TableManagementApi.md#CreateTable) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/tables | Create a YugabyteDB table
[**CreateTableSpaces**](TableManagementApi.md#CreateTableSpaces) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/tablespaces | Create tableSpaces
[**DescribeTable**](TableManagementApi.md#DescribeTable) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/tables/{tableUUID} | Describe a table
[**DropTable**](TableManagementApi.md#DropTable) | **Delete** /api/v1/customers/{cUUID}/universes/{uniUUID}/tables/{tableUUID} | Drop a YugabyteDB table
[**GetAllNamespaces**](TableManagementApi.md#GetAllNamespaces) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/namespaces | List all namespaces
[**GetAllTableSpaces**](TableManagementApi.md#GetAllTableSpaces) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/tablespaces | List all tablespaces
[**GetAllTables**](TableManagementApi.md#GetAllTables) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/tables | List all tables
[**GetYQLDataTypes**](TableManagementApi.md#GetYQLDataTypes) | **Get** /api/v1/metadata/yql_data_types | List column types



## AlterTable

> map[string]map[string]interface{} AlterTable(ctx, cUUID, uniUUID, tableUUID).Request(request).Execute()

Alter a YugabyteDB table

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
    tableUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TableManagementApi.AlterTable(context.Background(), cUUID, uniUUID, tableUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableManagementApi.AlterTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlterTable`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TableManagementApi.AlterTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 
**tableUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlterTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

**map[string]map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkImportData

> YBPTask BulkImportData(ctx, cUUID, uniUUID, tableUUID).BulkImport(bulkImport).Request(request).Execute()

Bulk import data



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
    tableUUID := TODO // string | 
    bulkImport := *openapiclient.NewBulkImportParams(*openapiclient.NewUsers("username1@example.com"), "PlatformUrl_example", "PlatformVersion_example", "S3Bucket_example", int32(123), int32(123)) // BulkImportParams | Bulk data to be imported
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TableManagementApi.BulkImportData(context.Background(), cUUID, uniUUID, tableUUID).BulkImport(bulkImport).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableManagementApi.BulkImportData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkImportData`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `TableManagementApi.BulkImportData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 
**tableUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkImportDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bulkImport** | [**BulkImportParams**](BulkImportParams.md) | Bulk data to be imported | 
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


## CreateMultiTableBackup

> Schedule CreateMultiTableBackup(ctx, cUUID, uniUUID).TableBackup(tableBackup).Request(request).Execute()

Create a multi-table backup

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
    tableBackup := *openapiclient.NewMultiTableBackupRequestParams(*openapiclient.NewUsers("username1@example.com"), "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123), "StorageConfigUUID_example", int64(123)) // MultiTableBackupRequestParams | Table backup data to be created
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TableManagementApi.CreateMultiTableBackup(context.Background(), cUUID, uniUUID).TableBackup(tableBackup).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableManagementApi.CreateMultiTableBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMultiTableBackup`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `TableManagementApi.CreateMultiTableBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMultiTableBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tableBackup** | [**MultiTableBackupRequestParams**](MultiTableBackupRequestParams.md) | Table backup data to be created | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSingleTableBackup

> YBPTask CreateSingleTableBackup(ctx, cUUID, uniUUID, tableUUID).Backup(backup).Request(request).Execute()

Create a single-table backup

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
    tableUUID := TODO // string | 
    backup := *openapiclient.NewBackupTableParams(*openapiclient.NewUsers("username1@example.com"), "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123), "StorageConfigUUID_example", int64(123)) // BackupTableParams | Backup data to be created
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TableManagementApi.CreateSingleTableBackup(context.Background(), cUUID, uniUUID, tableUUID).Backup(backup).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableManagementApi.CreateSingleTableBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSingleTableBackup`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `TableManagementApi.CreateSingleTableBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 
**tableUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSingleTableBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **backup** | [**BackupTableParams**](BackupTableParams.md) | Backup data to be created | 
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


## CreateTable

> YBPTask CreateTable(ctx, cUUID, uniUUID).Table(table).Request(request).Execute()

Create a YugabyteDB table

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
    table := *openapiclient.NewTableDefinitionTaskParams(*openapiclient.NewUsers("username1@example.com"), "PlatformUrl_example", "PlatformVersion_example", int32(123), int32(123), *openapiclient.NewTableDetails(), "TableType_example", "TableUUID_example") // TableDefinitionTaskParams | Table definition to be created
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TableManagementApi.CreateTable(context.Background(), cUUID, uniUUID).Table(table).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableManagementApi.CreateTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTable`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `TableManagementApi.CreateTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **table** | [**TableDefinitionTaskParams**](TableDefinitionTaskParams.md) | Table definition to be created | 
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


## CreateTableSpaces

> YBPTask CreateTableSpaces(ctx, cUUID, uniUUID).CreateTableSpacesRequest(createTableSpacesRequest).Request(request).Execute()

Create tableSpaces

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
    createTableSpacesRequest := *openapiclient.NewCreateTablespaceParams([]openapiclient.TableSpaceInfo{*openapiclient.NewTableSpaceInfo("Name_example", []openapiclient.PlacementBlock{*openapiclient.NewPlacementBlock("Cloud_example", "Region_example", "Zone_example")})}) // CreateTablespaceParams | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TableManagementApi.CreateTableSpaces(context.Background(), cUUID, uniUUID).CreateTableSpacesRequest(createTableSpacesRequest).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableManagementApi.CreateTableSpaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTableSpaces`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `TableManagementApi.CreateTableSpaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTableSpacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createTableSpacesRequest** | [**CreateTablespaceParams**](CreateTablespaceParams.md) |  | 
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
    tableUUID := TODO // string | 

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
**tableUUID** | [**string**](.md) |  | 

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


## DropTable

> YBPTask DropTable(ctx, cUUID, uniUUID, tableUUID).Request(request).Execute()

Drop a YugabyteDB table

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
    tableUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TableManagementApi.DropTable(context.Background(), cUUID, uniUUID, tableUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableManagementApi.DropTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DropTable`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `TableManagementApi.DropTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 
**tableUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDropTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**interface{}**](interface{}.md) |  | 

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


## GetAllNamespaces

> []NamespaceInfoResp GetAllNamespaces(ctx, cUUID, uniUUID).IncludeSystemNamespaces(includeSystemNamespaces).Execute()

List all namespaces



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

List all tablespaces



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

> []TableInfoResp GetAllTables(ctx, cUUID, uniUUID).IncludeParentTableInfo(includeParentTableInfo).ExcludeColocatedTables(excludeColocatedTables).IncludeColocatedParentTables(includeColocatedParentTables).Execute()

List all tables



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
    includeColocatedParentTables := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TableManagementApi.GetAllTables(context.Background(), cUUID, uniUUID).IncludeParentTableInfo(includeParentTableInfo).ExcludeColocatedTables(excludeColocatedTables).IncludeColocatedParentTables(includeColocatedParentTables).Execute()
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
 **includeColocatedParentTables** | **bool** |  | [default to false]

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


## GetYQLDataTypes

> []string GetYQLDataTypes(ctx).Execute()

List column types



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TableManagementApi.GetYQLDataTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TableManagementApi.GetYQLDataTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetYQLDataTypes`: []string
    fmt.Fprintf(os.Stdout, "Response from `TableManagementApi.GetYQLDataTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetYQLDataTypesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

