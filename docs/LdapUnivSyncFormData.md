# LdapUnivSyncFormData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateGroups** | Pointer to **bool** | Allow the API to create the LDAP groups as DB superusers | [optional] 
**DbUser** | Pointer to **string** | Database user to connect: yugabyte for ysql, cassandra for ycql | [optional] 
**DbuserPassword** | **string** |  | 
**ExcludeUsers** | Pointer to **[]string** | List of users to exclude while revoking and dropping | [optional] 
**LdapBasedn** | Pointer to **string** | Dn of the search starting point. | [optional] 
**LdapBindDn** | Pointer to **string** | Dn of the user authenticating to LDAP. | [optional] 
**LdapBindPassword** | Pointer to **string** | Password of the user authenticating to LDAP. | [optional] 
**LdapGroupMemberOfAttribute** | Pointer to **string** | LDAP group dn attribute to which the user belongs | [optional] 
**LdapGroupfield** | **string** | LDAP field to get the group information | 
**LdapPort** | Pointer to **int32** | Port of the ldap server : 389 or 636(tls) | [optional] 
**LdapSearchFilter** | Pointer to **string** | LDAP search filter to get the user entries. This filter can also be used to search for the users based on their group memberships. | [optional] 
**LdapServer** | Pointer to **string** | IP address of the LDAP server | [optional] 
**LdapTlsProtocol** | Pointer to **string** | TLS versions for LDAPS : TLSv1, TLSv1_1, TLSv1_2 | [optional] 
**LdapUserfield** | **string** | LDAP field to get the user information | 
**TargetApi** | **string** |  | 
**UseLdapTls** | Pointer to **bool** | Use LDAP TLS | [optional] 

## Methods

### NewLdapUnivSyncFormData

`func NewLdapUnivSyncFormData(dbuserPassword string, ldapGroupfield string, ldapUserfield string, targetApi string, ) *LdapUnivSyncFormData`

NewLdapUnivSyncFormData instantiates a new LdapUnivSyncFormData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapUnivSyncFormDataWithDefaults

`func NewLdapUnivSyncFormDataWithDefaults() *LdapUnivSyncFormData`

NewLdapUnivSyncFormDataWithDefaults instantiates a new LdapUnivSyncFormData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateGroups

`func (o *LdapUnivSyncFormData) GetCreateGroups() bool`

GetCreateGroups returns the CreateGroups field if non-nil, zero value otherwise.

### GetCreateGroupsOk

`func (o *LdapUnivSyncFormData) GetCreateGroupsOk() (*bool, bool)`

GetCreateGroupsOk returns a tuple with the CreateGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateGroups

`func (o *LdapUnivSyncFormData) SetCreateGroups(v bool)`

SetCreateGroups sets CreateGroups field to given value.

### HasCreateGroups

`func (o *LdapUnivSyncFormData) HasCreateGroups() bool`

HasCreateGroups returns a boolean if a field has been set.

### GetDbUser

`func (o *LdapUnivSyncFormData) GetDbUser() string`

GetDbUser returns the DbUser field if non-nil, zero value otherwise.

### GetDbUserOk

`func (o *LdapUnivSyncFormData) GetDbUserOk() (*string, bool)`

GetDbUserOk returns a tuple with the DbUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUser

`func (o *LdapUnivSyncFormData) SetDbUser(v string)`

SetDbUser sets DbUser field to given value.

### HasDbUser

`func (o *LdapUnivSyncFormData) HasDbUser() bool`

HasDbUser returns a boolean if a field has been set.

### GetDbuserPassword

`func (o *LdapUnivSyncFormData) GetDbuserPassword() string`

GetDbuserPassword returns the DbuserPassword field if non-nil, zero value otherwise.

### GetDbuserPasswordOk

`func (o *LdapUnivSyncFormData) GetDbuserPasswordOk() (*string, bool)`

GetDbuserPasswordOk returns a tuple with the DbuserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbuserPassword

`func (o *LdapUnivSyncFormData) SetDbuserPassword(v string)`

SetDbuserPassword sets DbuserPassword field to given value.


### GetExcludeUsers

`func (o *LdapUnivSyncFormData) GetExcludeUsers() []string`

GetExcludeUsers returns the ExcludeUsers field if non-nil, zero value otherwise.

### GetExcludeUsersOk

`func (o *LdapUnivSyncFormData) GetExcludeUsersOk() (*[]string, bool)`

GetExcludeUsersOk returns a tuple with the ExcludeUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUsers

`func (o *LdapUnivSyncFormData) SetExcludeUsers(v []string)`

SetExcludeUsers sets ExcludeUsers field to given value.

### HasExcludeUsers

`func (o *LdapUnivSyncFormData) HasExcludeUsers() bool`

HasExcludeUsers returns a boolean if a field has been set.

### GetLdapBasedn

`func (o *LdapUnivSyncFormData) GetLdapBasedn() string`

GetLdapBasedn returns the LdapBasedn field if non-nil, zero value otherwise.

### GetLdapBasednOk

`func (o *LdapUnivSyncFormData) GetLdapBasednOk() (*string, bool)`

GetLdapBasednOk returns a tuple with the LdapBasedn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBasedn

`func (o *LdapUnivSyncFormData) SetLdapBasedn(v string)`

SetLdapBasedn sets LdapBasedn field to given value.

### HasLdapBasedn

`func (o *LdapUnivSyncFormData) HasLdapBasedn() bool`

HasLdapBasedn returns a boolean if a field has been set.

### GetLdapBindDn

`func (o *LdapUnivSyncFormData) GetLdapBindDn() string`

GetLdapBindDn returns the LdapBindDn field if non-nil, zero value otherwise.

### GetLdapBindDnOk

`func (o *LdapUnivSyncFormData) GetLdapBindDnOk() (*string, bool)`

GetLdapBindDnOk returns a tuple with the LdapBindDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBindDn

`func (o *LdapUnivSyncFormData) SetLdapBindDn(v string)`

SetLdapBindDn sets LdapBindDn field to given value.

### HasLdapBindDn

`func (o *LdapUnivSyncFormData) HasLdapBindDn() bool`

HasLdapBindDn returns a boolean if a field has been set.

### GetLdapBindPassword

`func (o *LdapUnivSyncFormData) GetLdapBindPassword() string`

GetLdapBindPassword returns the LdapBindPassword field if non-nil, zero value otherwise.

### GetLdapBindPasswordOk

`func (o *LdapUnivSyncFormData) GetLdapBindPasswordOk() (*string, bool)`

GetLdapBindPasswordOk returns a tuple with the LdapBindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBindPassword

`func (o *LdapUnivSyncFormData) SetLdapBindPassword(v string)`

SetLdapBindPassword sets LdapBindPassword field to given value.

### HasLdapBindPassword

`func (o *LdapUnivSyncFormData) HasLdapBindPassword() bool`

HasLdapBindPassword returns a boolean if a field has been set.

### GetLdapGroupMemberOfAttribute

`func (o *LdapUnivSyncFormData) GetLdapGroupMemberOfAttribute() string`

GetLdapGroupMemberOfAttribute returns the LdapGroupMemberOfAttribute field if non-nil, zero value otherwise.

### GetLdapGroupMemberOfAttributeOk

`func (o *LdapUnivSyncFormData) GetLdapGroupMemberOfAttributeOk() (*string, bool)`

GetLdapGroupMemberOfAttributeOk returns a tuple with the LdapGroupMemberOfAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupMemberOfAttribute

`func (o *LdapUnivSyncFormData) SetLdapGroupMemberOfAttribute(v string)`

SetLdapGroupMemberOfAttribute sets LdapGroupMemberOfAttribute field to given value.

### HasLdapGroupMemberOfAttribute

`func (o *LdapUnivSyncFormData) HasLdapGroupMemberOfAttribute() bool`

HasLdapGroupMemberOfAttribute returns a boolean if a field has been set.

### GetLdapGroupfield

`func (o *LdapUnivSyncFormData) GetLdapGroupfield() string`

GetLdapGroupfield returns the LdapGroupfield field if non-nil, zero value otherwise.

### GetLdapGroupfieldOk

`func (o *LdapUnivSyncFormData) GetLdapGroupfieldOk() (*string, bool)`

GetLdapGroupfieldOk returns a tuple with the LdapGroupfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupfield

`func (o *LdapUnivSyncFormData) SetLdapGroupfield(v string)`

SetLdapGroupfield sets LdapGroupfield field to given value.


### GetLdapPort

`func (o *LdapUnivSyncFormData) GetLdapPort() int32`

GetLdapPort returns the LdapPort field if non-nil, zero value otherwise.

### GetLdapPortOk

`func (o *LdapUnivSyncFormData) GetLdapPortOk() (*int32, bool)`

GetLdapPortOk returns a tuple with the LdapPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapPort

`func (o *LdapUnivSyncFormData) SetLdapPort(v int32)`

SetLdapPort sets LdapPort field to given value.

### HasLdapPort

`func (o *LdapUnivSyncFormData) HasLdapPort() bool`

HasLdapPort returns a boolean if a field has been set.

### GetLdapSearchFilter

`func (o *LdapUnivSyncFormData) GetLdapSearchFilter() string`

GetLdapSearchFilter returns the LdapSearchFilter field if non-nil, zero value otherwise.

### GetLdapSearchFilterOk

`func (o *LdapUnivSyncFormData) GetLdapSearchFilterOk() (*string, bool)`

GetLdapSearchFilterOk returns a tuple with the LdapSearchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapSearchFilter

`func (o *LdapUnivSyncFormData) SetLdapSearchFilter(v string)`

SetLdapSearchFilter sets LdapSearchFilter field to given value.

### HasLdapSearchFilter

`func (o *LdapUnivSyncFormData) HasLdapSearchFilter() bool`

HasLdapSearchFilter returns a boolean if a field has been set.

### GetLdapServer

`func (o *LdapUnivSyncFormData) GetLdapServer() string`

GetLdapServer returns the LdapServer field if non-nil, zero value otherwise.

### GetLdapServerOk

`func (o *LdapUnivSyncFormData) GetLdapServerOk() (*string, bool)`

GetLdapServerOk returns a tuple with the LdapServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServer

`func (o *LdapUnivSyncFormData) SetLdapServer(v string)`

SetLdapServer sets LdapServer field to given value.

### HasLdapServer

`func (o *LdapUnivSyncFormData) HasLdapServer() bool`

HasLdapServer returns a boolean if a field has been set.

### GetLdapTlsProtocol

`func (o *LdapUnivSyncFormData) GetLdapTlsProtocol() string`

GetLdapTlsProtocol returns the LdapTlsProtocol field if non-nil, zero value otherwise.

### GetLdapTlsProtocolOk

`func (o *LdapUnivSyncFormData) GetLdapTlsProtocolOk() (*string, bool)`

GetLdapTlsProtocolOk returns a tuple with the LdapTlsProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapTlsProtocol

`func (o *LdapUnivSyncFormData) SetLdapTlsProtocol(v string)`

SetLdapTlsProtocol sets LdapTlsProtocol field to given value.

### HasLdapTlsProtocol

`func (o *LdapUnivSyncFormData) HasLdapTlsProtocol() bool`

HasLdapTlsProtocol returns a boolean if a field has been set.

### GetLdapUserfield

`func (o *LdapUnivSyncFormData) GetLdapUserfield() string`

GetLdapUserfield returns the LdapUserfield field if non-nil, zero value otherwise.

### GetLdapUserfieldOk

`func (o *LdapUnivSyncFormData) GetLdapUserfieldOk() (*string, bool)`

GetLdapUserfieldOk returns a tuple with the LdapUserfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserfield

`func (o *LdapUnivSyncFormData) SetLdapUserfield(v string)`

SetLdapUserfield sets LdapUserfield field to given value.


### GetTargetApi

`func (o *LdapUnivSyncFormData) GetTargetApi() string`

GetTargetApi returns the TargetApi field if non-nil, zero value otherwise.

### GetTargetApiOk

`func (o *LdapUnivSyncFormData) GetTargetApiOk() (*string, bool)`

GetTargetApiOk returns a tuple with the TargetApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetApi

`func (o *LdapUnivSyncFormData) SetTargetApi(v string)`

SetTargetApi sets TargetApi field to given value.


### GetUseLdapTls

`func (o *LdapUnivSyncFormData) GetUseLdapTls() bool`

GetUseLdapTls returns the UseLdapTls field if non-nil, zero value otherwise.

### GetUseLdapTlsOk

`func (o *LdapUnivSyncFormData) GetUseLdapTlsOk() (*bool, bool)`

GetUseLdapTlsOk returns a tuple with the UseLdapTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLdapTls

`func (o *LdapUnivSyncFormData) SetUseLdapTls(v bool)`

SetUseLdapTls sets UseLdapTls field to given value.

### HasUseLdapTls

`func (o *LdapUnivSyncFormData) HasUseLdapTls() bool`

HasUseLdapTls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


