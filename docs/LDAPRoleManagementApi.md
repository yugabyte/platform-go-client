# \LDAPRoleManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListLdapDnToYbaRoles**](LDAPRoleManagementApi.md#ListLdapDnToYbaRoles) | **Get** /api/v1/customers/{cUUID}/ldap_mappings | List LDAP Mappings
[**SetLdapDnToYbaRoles**](LDAPRoleManagementApi.md#SetLdapDnToYbaRoles) | **Put** /api/v1/customers/{cUUID}/ldap_mappings | Set LDAP Mappings
[**SyncLdapUniverse**](LDAPRoleManagementApi.md#SyncLdapUniverse) | **Post** /api/v1/customers/{cUUID}/universes/{univUUID}/ldap_roles_sync | WARNING: This is a preview API that could change. Perform an LDAP users sync on the universe



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
    resp, r, err := api_client.LDAPRoleManagementApi.ListLdapDnToYbaRoles(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPRoleManagementApi.ListLdapDnToYbaRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLdapDnToYbaRoles`: LdapDnToYbaRoleData
    fmt.Fprintf(os.Stdout, "Response from `LDAPRoleManagementApi.ListLdapDnToYbaRoles`: %v\n", resp)
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
    resp, r, err := api_client.LDAPRoleManagementApi.SetLdapDnToYbaRoles(context.Background(), cUUID).LdapMappings(ldapMappings).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPRoleManagementApi.SetLdapDnToYbaRoles``: %v\n", err)
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


## SyncLdapUniverse

> YBPTask SyncLdapUniverse(ctx, cUUID, univUUID).SyncLdapUniverse(syncLdapUniverse).Request(request).Execute()

WARNING: This is a preview API that could change. Perform an LDAP users sync on the universe



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
    syncLdapUniverse := *openapiclient.NewLdapUnivSyncFormData("DbuserPassword_example", "cn", "cn", "TargetApi_example") // LdapUnivSyncFormData | config to sync universe roles with ldap users
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LDAPRoleManagementApi.SyncLdapUniverse(context.Background(), cUUID, univUUID).SyncLdapUniverse(syncLdapUniverse).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LDAPRoleManagementApi.SyncLdapUniverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncLdapUniverse`: YBPTask
    fmt.Fprintf(os.Stdout, "Response from `LDAPRoleManagementApi.SyncLdapUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**univUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncLdapUniverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **syncLdapUniverse** | [**LdapUnivSyncFormData**](LdapUnivSyncFormData.md) | config to sync universe roles with ldap users | 
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

