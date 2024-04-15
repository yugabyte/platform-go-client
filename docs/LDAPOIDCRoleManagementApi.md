# \LDAPOIDCRoleManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOidcGroupMapping**](LDAPOIDCRoleManagementApi.md#DeleteOidcGroupMapping) | **Delete** /api/v1/customers/{cUUID}/oidc_mappings/{groupName} | Delete a OIDC group mapping
[**ListLdapDnToYbaRoles**](LDAPOIDCRoleManagementApi.md#ListLdapDnToYbaRoles) | **Get** /api/v1/customers/{cUUID}/ldap_mappings | List LDAP Mappings
[**ListOidcGroupToYbaRoles**](LDAPOIDCRoleManagementApi.md#ListOidcGroupToYbaRoles) | **Get** /api/v1/customers/{cUUID}/oidc_mappings | List OIDC Group Mappings
[**MapOidcGroupToYbaRoles**](LDAPOIDCRoleManagementApi.md#MapOidcGroupToYbaRoles) | **Put** /api/v1/customers/{cUUID}/oidc_mappings | Set OIDC Mappings
[**SetLdapDnToYbaRoles**](LDAPOIDCRoleManagementApi.md#SetLdapDnToYbaRoles) | **Put** /api/v1/customers/{cUUID}/ldap_mappings | Set LDAP Mappings



## DeleteOidcGroupMapping

> YBPSuccess DeleteOidcGroupMapping(ctx, cUUID, groupName).Request(request).Execute()

Delete a OIDC group mapping



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
    groupName := "groupName_example" // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LDAPOIDCRoleManagementApi.DeleteOidcGroupMapping(context.Background(), cUUID, groupName).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPOIDCRoleManagementApi.DeleteOidcGroupMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOidcGroupMapping`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `LDAPOIDCRoleManagementApi.DeleteOidcGroupMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**groupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOidcGroupMappingRequest struct via the builder pattern


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


## ListLdapDnToYbaRoles

> LdapDnToYbaRoleData ListLdapDnToYbaRoles(ctx, cUUID).Execute()

List LDAP Mappings



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
    resp, r, err := api_client.LDAPOIDCRoleManagementApi.ListLdapDnToYbaRoles(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPOIDCRoleManagementApi.ListLdapDnToYbaRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLdapDnToYbaRoles`: LdapDnToYbaRoleData
    fmt.Fprintf(os.Stdout, "Response from `LDAPOIDCRoleManagementApi.ListLdapDnToYbaRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLdapDnToYbaRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LdapDnToYbaRoleData**](LdapDnToYbaRoleData.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOidcGroupToYbaRoles

> OidcGroupToYbaRolesData ListOidcGroupToYbaRoles(ctx, cUUID).Execute()

List OIDC Group Mappings



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
    resp, r, err := api_client.LDAPOIDCRoleManagementApi.ListOidcGroupToYbaRoles(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPOIDCRoleManagementApi.ListOidcGroupToYbaRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOidcGroupToYbaRoles`: OidcGroupToYbaRolesData
    fmt.Fprintf(os.Stdout, "Response from `LDAPOIDCRoleManagementApi.ListOidcGroupToYbaRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOidcGroupToYbaRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OidcGroupToYbaRolesData**](OidcGroupToYbaRolesData.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MapOidcGroupToYbaRoles

> MapOidcGroupToYbaRoles(ctx, cUUID).OidcMappings(oidcMappings).Request(request).Execute()

Set OIDC Mappings



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
    oidcMappings := *openapiclient.NewOidcGroupToYbaRolesData() // OidcGroupToYbaRolesData | New OIDC Mappings to be set
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LDAPOIDCRoleManagementApi.MapOidcGroupToYbaRoles(context.Background(), cUUID).OidcMappings(oidcMappings).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPOIDCRoleManagementApi.MapOidcGroupToYbaRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMapOidcGroupToYbaRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oidcMappings** | [**OidcGroupToYbaRolesData**](OidcGroupToYbaRolesData.md) | New OIDC Mappings to be set | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetLdapDnToYbaRoles

> SetLdapDnToYbaRoles(ctx, cUUID).LdapMappings(ldapMappings).Request(request).Execute()

Set LDAP Mappings



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
    ldapMappings := *openapiclient.NewLdapDnToYbaRoleData() // LdapDnToYbaRoleData | New LDAP Mappings to be set
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LDAPOIDCRoleManagementApi.SetLdapDnToYbaRoles(context.Background(), cUUID).LdapMappings(ldapMappings).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPOIDCRoleManagementApi.SetLdapDnToYbaRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLdapDnToYbaRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ldapMappings** | [**LdapDnToYbaRoleData**](LdapDnToYbaRoleData.md) | New LDAP Mappings to be set | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

