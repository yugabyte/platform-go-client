# \UniverseDatabaseManagementAPI

All URIs are relative to *http://localhost*

| Method                                                                                | HTTP request                                                                 | Description                               |
| ------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ----------------------------------------- |
| [**ConfigureYCQL**](UniverseDatabaseManagementAPI.md#ConfigureYCQL)                   | **Post** /api/v1/customers/{cUUID}/universes/{univUUID}/configure/ycql       | Configure YCQL                            |
| [**ConfigureYSQL**](UniverseDatabaseManagementAPI.md#ConfigureYSQL)                   | **Post** /api/v1/customers/{cUUID}/universes/{univUUID}/configure/ysql       | Configure YSQL                            |
| [**CreateUserInDB**](UniverseDatabaseManagementAPI.md#CreateUserInDB)                 | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/create_db_credentials | Create a database user for a universe     |
| [**RunYsqlQueryUniverse**](UniverseDatabaseManagementAPI.md#RunYsqlQueryUniverse)     | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/run_query             | Run a YSQL query in a universe            |
| [**SetDatabaseCredentials**](UniverseDatabaseManagementAPI.md#SetDatabaseCredentials) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/update_db_credentials | Set a universe&#39;s database credentials |



## ConfigureYCQL

> YBPTask ConfigureYCQL(ctx, cUUID, univUUID).ConfigureYcqlFormData(configureYcqlFormData).Request(request).Execute()

Configure YCQL



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	univUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	configureYcqlFormData := *openapiclient.NewConfigureYCQLFormData() // ConfigureYCQLFormData | Configure YCQL Form Data
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseDatabaseManagementAPI.ConfigureYCQL(context.Background(), cUUID, univUUID).ConfigureYcqlFormData(configureYcqlFormData).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseDatabaseManagementAPI.ConfigureYCQL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigureYCQL`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `UniverseDatabaseManagementAPI.ConfigureYCQL`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**    | **string**          |                                                                             |
| **univUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiConfigureYCQLRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


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

Configure YSQL



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	univUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	configureYsqlFormData := *openapiclient.NewConfigureYSQLFormData() // ConfigureYSQLFormData | Configure YSQL Form Data
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseDatabaseManagementAPI.ConfigureYSQL(context.Background(), cUUID, univUUID).ConfigureYsqlFormData(configureYsqlFormData).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseDatabaseManagementAPI.ConfigureYSQL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigureYSQL`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `UniverseDatabaseManagementAPI.ConfigureYSQL`: %v\n", resp)
}
```

### Path Parameters


| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**    | **string**          |                                                                             |
| **univUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiConfigureYSQLRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


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
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	databaseUserFormData := *openapiclient.NewDatabaseUserFormData("DbName_example", "Password_example", "Username_example", "YcqlAdminPassword_example", "YcqlAdminUsername_example", "YsqlAdminPassword_example", "YsqlAdminUsername_example") // DatabaseUserFormData | The database user to create
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseDatabaseManagementAPI.CreateUserInDB(context.Background(), cUUID, uniUUID).DatabaseUserFormData(databaseUserFormData).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseDatabaseManagementAPI.CreateUserInDB``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserInDB`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `UniverseDatabaseManagementAPI.CreateUserInDB`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserInDBRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


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
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	runQueryFormData := *openapiclient.NewRunQueryFormData("DbName_example", "NodeName_example", "Query_example", "TableType_example") // RunQueryFormData | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseDatabaseManagementAPI.RunYsqlQueryUniverse(context.Background(), cUUID, uniUUID).RunQueryFormData(runQueryFormData).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseDatabaseManagementAPI.RunYsqlQueryUniverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunYsqlQueryUniverse`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UniverseDatabaseManagementAPI.RunYsqlQueryUniverse`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiRunYsqlQueryUniverseRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


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
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	uniUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	databaseSecurityFormData := *openapiclient.NewDatabaseSecurityFormData() // DatabaseSecurityFormData | The database credentials
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniverseDatabaseManagementAPI.SetDatabaseCredentials(context.Background(), cUUID, uniUUID).DatabaseSecurityFormData(databaseSecurityFormData).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniverseDatabaseManagementAPI.SetDatabaseCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDatabaseCredentials`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `UniverseDatabaseManagementAPI.SetDatabaseCredentials`: %v\n", resp)
}
```

### Path Parameters


| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID**   | **string**          |                                                                             |
| **uniUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiSetDatabaseCredentialsRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


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

