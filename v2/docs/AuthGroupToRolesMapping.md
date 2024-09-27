# AuthGroupToRolesMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupIdentifier** | **string** | Group name incase of OIDC. Group DN incase of LDAP. | 
**Uuid** | Pointer to **string** | System generated UUID for this group mapping. | [optional] [readonly] 
**Type** | **string** | The type of group. Can be either LDAP/OIDC. | 
**CreationDate** | Pointer to **time.Time** | Group mapping creation date. | [optional] [readonly] 
**RoleResourceDefinitions** | [**[]RoleResourceDefinition**](RoleResourceDefinition.md) |  | 

## Methods

### NewAuthGroupToRolesMapping

`func NewAuthGroupToRolesMapping(groupIdentifier string, type_ string, roleResourceDefinitions []RoleResourceDefinition, ) *AuthGroupToRolesMapping`

NewAuthGroupToRolesMapping instantiates a new AuthGroupToRolesMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthGroupToRolesMappingWithDefaults

`func NewAuthGroupToRolesMappingWithDefaults() *AuthGroupToRolesMapping`

NewAuthGroupToRolesMappingWithDefaults instantiates a new AuthGroupToRolesMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupIdentifier

`func (o *AuthGroupToRolesMapping) GetGroupIdentifier() string`

GetGroupIdentifier returns the GroupIdentifier field if non-nil, zero value otherwise.

### GetGroupIdentifierOk

`func (o *AuthGroupToRolesMapping) GetGroupIdentifierOk() (*string, bool)`

GetGroupIdentifierOk returns a tuple with the GroupIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIdentifier

`func (o *AuthGroupToRolesMapping) SetGroupIdentifier(v string)`

SetGroupIdentifier sets GroupIdentifier field to given value.


### GetUuid

`func (o *AuthGroupToRolesMapping) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AuthGroupToRolesMapping) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AuthGroupToRolesMapping) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AuthGroupToRolesMapping) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *AuthGroupToRolesMapping) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthGroupToRolesMapping) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthGroupToRolesMapping) SetType(v string)`

SetType sets Type field to given value.


### GetCreationDate

`func (o *AuthGroupToRolesMapping) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *AuthGroupToRolesMapping) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *AuthGroupToRolesMapping) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *AuthGroupToRolesMapping) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetRoleResourceDefinitions

`func (o *AuthGroupToRolesMapping) GetRoleResourceDefinitions() []RoleResourceDefinition`

GetRoleResourceDefinitions returns the RoleResourceDefinitions field if non-nil, zero value otherwise.

### GetRoleResourceDefinitionsOk

`func (o *AuthGroupToRolesMapping) GetRoleResourceDefinitionsOk() (*[]RoleResourceDefinition, bool)`

GetRoleResourceDefinitionsOk returns a tuple with the RoleResourceDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleResourceDefinitions

`func (o *AuthGroupToRolesMapping) SetRoleResourceDefinitions(v []RoleResourceDefinition)`

SetRoleResourceDefinitions sets RoleResourceDefinitions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


