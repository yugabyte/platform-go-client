# \LDAPRoleManagementAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SyncLdapUniverse**](LDAPRoleManagementAPI.md#SyncLdapUniverse) | **Post** /api/v1/customers/{cUUID}/universes/{univUUID}/ldap_roles_sync | Perform an LDAP users sync on the universe



## SyncLdapUniverse

> YBPTask SyncLdapUniverse(ctx, cUUID, univUUID).SyncLdapUniverse(syncLdapUniverse).Request(request).Execute()

Perform an LDAP users sync on the universe



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
	univUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	syncLdapUniverse := *openapiclient.NewLdapUnivSyncFormData("DbuserPassword_example", "cn", "cn, sAMAccountName", "TargetApi_example") // LdapUnivSyncFormData | config to sync universe roles with ldap users
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LDAPRoleManagementAPI.SyncLdapUniverse(context.Background(), cUUID, univUUID).SyncLdapUniverse(syncLdapUniverse).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LDAPRoleManagementAPI.SyncLdapUniverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncLdapUniverse`: YBPTask
	fmt.Fprintf(os.Stdout, "Response from `LDAPRoleManagementAPI.SyncLdapUniverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**univUUID** | **string** |  | 

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

