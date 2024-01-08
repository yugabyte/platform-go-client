# \UserManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UserManagementApi.md#CreateUser) | **Post** /api/v1/customers/{cUUID}/users | Create a user
[**DeleteUser**](UserManagementApi.md#DeleteUser) | **Delete** /api/v1/customers/{cUUID}/users/{uUUID} | Delete a user
[**GetUserDetails**](UserManagementApi.md#GetUserDetails) | **Get** /api/v1/customers/{cUUID}/users/{uUUID} | Get a user&#39;s details
[**ListUsers**](UserManagementApi.md#ListUsers) | **Get** /api/v1/customers/{cUUID}/users | List all users
[**RetrieveOIDCAuthToken**](UserManagementApi.md#RetrieveOIDCAuthToken) | **Get** /api/v1/customers/{cUUID}/users/{uUUID}/oidc_auth_token | Retrieve OIDC auth token
[**UpdateUserPassword**](UserManagementApi.md#UpdateUserPassword) | **Put** /api/v1/customers/{cUUID}/users/{uUUID}/change_password | Change a user&#39;s password
[**UpdateUserProfile**](UserManagementApi.md#UpdateUserProfile) | **Put** /api/v1/customers/{cUUID}/users/{uUUID}/update_profile | Update a user&#39;s profile
[**UpdateUserRole**](UserManagementApi.md#UpdateUserRole) | **Put** /api/v1/customers/{cUUID}/users/{uUUID} | Change a user&#39;s role



## CreateUser

> UserWithFeatures CreateUser(ctx, cUUID).User(user).Request(request).Execute()

Create a user

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
    user := *openapiclient.NewUserRegistrationData("test@example.com") // UserRegistrationData | Details of the new user
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.CreateUser(context.Background(), cUUID).User(user).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: UserWithFeatures
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**UserRegistrationData**](UserRegistrationData.md) | Details of the new user | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**UserWithFeatures**](UserWithFeatures.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> YBPSuccess DeleteUser(ctx, cUUID, uUUID).Request(request).Execute()

Delete a user



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
    uUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.DeleteUser(context.Background(), cUUID, uUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUser`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPSuccess**](YBPSuccess.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserDetails

> UserWithFeatures GetUserDetails(ctx, cUUID, uUUID).Execute()

Get a user's details

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
    uUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.GetUserDetails(context.Background(), cUUID, uUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.GetUserDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserDetails`: UserWithFeatures
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.GetUserDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserWithFeatures**](UserWithFeatures.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> []UserWithFeatures ListUsers(ctx, cUUID).Execute()

List all users

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.ListUsers(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: []UserWithFeatures
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.ListUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UserWithFeatures**](UserWithFeatures.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveOIDCAuthToken

> UserOIDCAuthToken RetrieveOIDCAuthToken(ctx, cUUID, uUUID).Request(request).Execute()

Retrieve OIDC auth token

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
    uUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.RetrieveOIDCAuthToken(context.Background(), cUUID, uUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.RetrieveOIDCAuthToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveOIDCAuthToken`: UserOIDCAuthToken
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.RetrieveOIDCAuthToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveOIDCAuthTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**UserOIDCAuthToken**](UserOIDCAuthToken.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPassword

> YBPSuccess UpdateUserPassword(ctx, cUUID, uUUID).Users(users).Request(request).Execute()

Change a user's password

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
    uUUID := TODO // string | 
    users := *openapiclient.NewUserRegistrationData("test@example.com") // UserRegistrationData | User data containing the new password
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.UpdateUserPassword(context.Background(), cUUID, uUUID).Users(users).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.UpdateUserPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserPassword`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.UpdateUserPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **users** | [**UserRegistrationData**](UserRegistrationData.md) | User data containing the new password | 
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


## UpdateUserProfile

> Users UpdateUserProfile(ctx, cUUID, uUUID).Users(users).Request(request).Execute()

Update a user's profile

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
    uUUID := TODO // string | 
    users := *openapiclient.NewUserProfileData("Admin") // UserProfileData | User data in profile to be updated
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.UpdateUserProfile(context.Background(), cUUID, uUUID).Users(users).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.UpdateUserProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserProfile`: Users
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.UpdateUserProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **users** | [**UserProfileData**](UserProfileData.md) | User data in profile to be updated | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**Users**](Users.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserRole

> YBPSuccess UpdateUserRole(ctx, cUUID, uUUID).Role(role).Request(request).Execute()

Change a user's role



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
    uUUID := TODO // string | 
    role := "role_example" // string |  (optional)
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserManagementApi.UpdateUserRole(context.Background(), cUUID, uUUID).Role(role).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementApi.UpdateUserRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserRole`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `UserManagementApi.UpdateUserRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **role** | **string** |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPSuccess**](YBPSuccess.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

