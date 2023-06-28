# \UniverseDatabaseManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserInDB**](UniverseDatabaseManagementApi.md#CreateUserInDB) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/create_db_credentials | Create a database user for a universe
[**RunYsqlQueryUniverse**](UniverseDatabaseManagementApi.md#RunYsqlQueryUniverse) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/run_query | Run a YSQL query in a universe
[**SetDatabaseCredentials**](UniverseDatabaseManagementApi.md#SetDatabaseCredentials) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/update_db_credentials | Set a universe&#39;s database credentials



## CreateUserInDB

> YBPSuccess CreateUserInDB(ctx, cUUID, uniUUID).DatabaseUserFormData(databaseUserFormData).Request(request).Execute()

Create a database user for a universe

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
    databaseUserFormData := *openapiclient.NewDatabaseUserFormData("DbName_example", "Password_example", "Username_example", "YcqlAdminPassword_example", "YcqlAdminUsername_example", "YsqlAdminPassword_example", "YsqlAdminUsername_example") // DatabaseUserFormData | The database user to create
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseDatabaseManagementApi.CreateUserInDB(context.Background(), cUUID, uniUUID).DatabaseUserFormData(databaseUserFormData).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseDatabaseManagementApi.CreateUserInDB``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserInDB`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `UniverseDatabaseManagementApi.CreateUserInDB`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserInDBRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **databaseUserFormData** | [**DatabaseUserFormData**](DatabaseUserFormData.md) | The database user to create | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPSuccess**](YBPSuccess.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunYsqlQueryUniverse

> map[string]interface{} RunYsqlQueryUniverse(ctx, cUUID, uniUUID).RunQueryFormData(runQueryFormData).Request(request).Execute()

Run a YSQL query in a universe



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
    runQueryFormData := *openapiclient.NewRunQueryFormData("DbName_example", "Query_example", "TableType_example") // RunQueryFormData | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseDatabaseManagementApi.RunYsqlQueryUniverse(context.Background(), cUUID, uniUUID).RunQueryFormData(runQueryFormData).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseDatabaseManagementApi.RunYsqlQueryUniverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunYsqlQueryUniverse`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UniverseDatabaseManagementApi.RunYsqlQueryUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunYsqlQueryUniverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **runQueryFormData** | [**RunQueryFormData**](RunQueryFormData.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDatabaseCredentials

> YBPSuccess SetDatabaseCredentials(ctx, cUUID, uniUUID).DatabaseSecurityFormData(databaseSecurityFormData).Request(request).Execute()

Set a universe's database credentials

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
    databaseSecurityFormData := *openapiclient.NewDatabaseSecurityFormData("DbName_example", "YcqlAdminPassword_example", "YcqlAdminUsername_example", "YcqlCurrAdminPassword_example", "YsqlAdminPassword_example", "YsqlAdminUsername_example", "YsqlCurrAdminPassword_example") // DatabaseSecurityFormData | The database credentials
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseDatabaseManagementApi.SetDatabaseCredentials(context.Background(), cUUID, uniUUID).DatabaseSecurityFormData(databaseSecurityFormData).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseDatabaseManagementApi.SetDatabaseCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetDatabaseCredentials`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `UniverseDatabaseManagementApi.SetDatabaseCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDatabaseCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **databaseSecurityFormData** | [**DatabaseSecurityFormData**](DatabaseSecurityFormData.md) | The database credentials | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPSuccess**](YBPSuccess.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

