# \TableManagementAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlterTable**](TableManagementAPI.md#AlterTable) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/tables/{tableUUID} | Alter a YugabyteDB table
[**BulkImportData**](TableManagementAPI.md#BulkImportData) | **Put** /api/v1/customers/{cUUID}/universes/{uniUUID}/tables/{tableUUID}/bulk_import | Bulk import data
[**CreateTable**](TableManagementAPI.md#CreateTable) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/tables | Create a YugabyteDB table
[**CreateTableSpaces**](TableManagementAPI.md#CreateTableSpaces) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/tablespaces | Create tableSpaces
[**DescribeTable**](TableManagementAPI.md#DescribeTable) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/tables/{tableUUID} | Describe a table
[**DropTable**](TableManagementAPI.md#DropTable) | **Delete** /api/v1/customers/{cUUID}/universes/{uniUUID}/tables/{tableUUID} | Drop a YugabyteDB table
[**GetAllNamespaces**](TableManagementAPI.md#GetAllNamespaces) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/namespaces | Get a list of all namespaces in the specified universe.
[**GetAllTableSpaces**](TableManagementAPI.md#GetAllTableSpaces) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/tablespaces | Get a list of all tablespaces of a given universe.
[**GetAllTables**](TableManagementAPI.md#GetAllTables) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/tables | Get a list of all tables in the specified universe
[**GetYQLDataTypes**](TableManagementAPI.md#GetYQLDataTypes) | **Get** /api/v1/metadata/yql_data_types | List column types



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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tableUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TableManagementAPI.AlterTable(context.Background(), cUUID, uniUUID, tableUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TableManagementAPI.AlterTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlterTable`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TableManagementAPI.AlterTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 
**tableUUID** | **string** |  | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tableUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	bulkImport := *openapiclient.NewBulkImportParams(*openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), "PlatformUrl_example", "S3Bucket_example", int32(123), int32(123)) // BulkImportParams | Bulk data to be imported
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TableManagementAPI.BulkImportData(context.Background(), cUUID, uniUUID, tableUUID).BulkImport(bulkImport).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TableManagementAPI.BulkImportData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkImportData`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `TableManagementAPI.BulkImportData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 
**tableUUID** | **string** |  | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	table := *openapiclient.NewTableDefinitionTaskParams(*openapiclient.NewUsers("username1@example.com", []string{"GroupMemberships_example"}, false), "PlatformUrl_example", int32(123), int32(123), *openapiclient.NewTableDetails(), "TableType_example", "TableUUID_example") // TableDefinitionTaskParams | Table definition to be created
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TableManagementAPI.CreateTable(context.Background(), cUUID, uniUUID).Table(table).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TableManagementAPI.CreateTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTable`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `TableManagementAPI.CreateTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	createTableSpacesRequest := *openapiclient.NewCreateTablespaceParams([]openapiclient.TableSpaceInfo{*openapiclient.NewTableSpaceInfo("Name_example", []openapiclient.PlacementBlock{*openapiclient.NewPlacementBlock("Cloud_example", "Region_example", "Zone_example")})}) // CreateTablespaceParams | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TableManagementAPI.CreateTableSpaces(context.Background(), cUUID, uniUUID).CreateTableSpacesRequest(createTableSpacesRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TableManagementAPI.CreateTableSpaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTableSpaces`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `TableManagementAPI.CreateTableSpaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tableUUID := "tableUUID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TableManagementAPI.DescribeTable(context.Background(), cUUID, uniUUID, tableUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TableManagementAPI.DescribeTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeTable`: TableDefinitionTaskParams
	fmt.Fprintf(os.Stdout, "Response from `TableManagementAPI.DescribeTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 
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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tableUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TableManagementAPI.DropTable(context.Background(), cUUID, uniUUID, tableUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TableManagementAPI.DropTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DropTable`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `TableManagementAPI.DropTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 
**tableUUID** | **string** |  | 

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

Get a list of all namespaces in the specified universe.



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
	includeSystemNamespaces := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TableManagementAPI.GetAllNamespaces(context.Background(), cUUID, uniUUID).IncludeSystemNamespaces(includeSystemNamespaces).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TableManagementAPI.GetAllNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllNamespaces`: []NamespaceInfoResp
	fmt.Fprintf(os.Stdout, "Response from `TableManagementAPI.GetAllNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TableManagementAPI.GetAllTableSpaces(context.Background(), cUUID, uniUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TableManagementAPI.GetAllTableSpaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllTableSpaces`: []TableSpaceInfo
	fmt.Fprintf(os.Stdout, "Response from `TableManagementAPI.GetAllTableSpaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	includeParentTableInfo := true // bool |  (optional) (default to false)
	excludeColocatedTables := true // bool |  (optional) (default to false)
	includeColocatedParentTables := true // bool |  (optional) (default to true)
	xClusterSupportedOnly := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TableManagementAPI.GetAllTables(context.Background(), cUUID, uniUUID).IncludeParentTableInfo(includeParentTableInfo).ExcludeColocatedTables(excludeColocatedTables).IncludeColocatedParentTables(includeColocatedParentTables).XClusterSupportedOnly(xClusterSupportedOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TableManagementAPI.GetAllTables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllTables`: []TableInfoResp
	fmt.Fprintf(os.Stdout, "Response from `TableManagementAPI.GetAllTables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**uniUUID** | **string** |  | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TableManagementAPI.GetYQLDataTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TableManagementAPI.GetYQLDataTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetYQLDataTypes`: []string
	fmt.Fprintf(os.Stdout, "Response from `TableManagementAPI.GetYQLDataTypes`: %v\n", resp)
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

