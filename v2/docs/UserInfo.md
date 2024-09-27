# UserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | User UUID | [optional] [readonly] 
**UserType** | Pointer to **string** | User Type | [optional] 
**CreationDate** | Pointer to **time.Time** | User creation date | [optional] [readonly] 
**CustomerUuid** | Pointer to **string** | Customer account to which this User belongs to | [optional] [readonly] 
**IsPrimary** | Pointer to **bool** | True if the user is the primary local user in this YBA | [optional] [readonly] 
**LdapSpecifiedRole** | Pointer to **bool** | Whether the user&#39;s role is inherited from LDAP | [optional] 
**Timezone** | Pointer to **string** | User timezone | [optional] 

## Methods

### NewUserInfo

`func NewUserInfo() *UserInfo`

NewUserInfo instantiates a new UserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInfoWithDefaults

`func NewUserInfoWithDefaults() *UserInfo`

NewUserInfoWithDefaults instantiates a new UserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *UserInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UserInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UserInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UserInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetUserType

`func (o *UserInfo) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UserInfo) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UserInfo) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *UserInfo) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetCreationDate

`func (o *UserInfo) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *UserInfo) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *UserInfo) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *UserInfo) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetCustomerUuid

`func (o *UserInfo) GetCustomerUuid() string`

GetCustomerUuid returns the CustomerUuid field if non-nil, zero value otherwise.

### GetCustomerUuidOk

`func (o *UserInfo) GetCustomerUuidOk() (*string, bool)`

GetCustomerUuidOk returns a tuple with the CustomerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUuid

`func (o *UserInfo) SetCustomerUuid(v string)`

SetCustomerUuid sets CustomerUuid field to given value.

### HasCustomerUuid

`func (o *UserInfo) HasCustomerUuid() bool`

HasCustomerUuid returns a boolean if a field has been set.

### GetIsPrimary

`func (o *UserInfo) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *UserInfo) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *UserInfo) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *UserInfo) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetLdapSpecifiedRole

`func (o *UserInfo) GetLdapSpecifiedRole() bool`

GetLdapSpecifiedRole returns the LdapSpecifiedRole field if non-nil, zero value otherwise.

### GetLdapSpecifiedRoleOk

`func (o *UserInfo) GetLdapSpecifiedRoleOk() (*bool, bool)`

GetLdapSpecifiedRoleOk returns a tuple with the LdapSpecifiedRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapSpecifiedRole

`func (o *UserInfo) SetLdapSpecifiedRole(v bool)`

SetLdapSpecifiedRole sets LdapSpecifiedRole field to given value.

### HasLdapSpecifiedRole

`func (o *UserInfo) HasLdapSpecifiedRole() bool`

HasLdapSpecifiedRole returns a boolean if a field has been set.

### GetTimezone

`func (o *UserInfo) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UserInfo) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UserInfo) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UserInfo) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


