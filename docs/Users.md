# Users

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthTokenIssueDate** | Pointer to **time.Time** | UI session token creation date | [optional] [readonly] 
**CreationDate** | Pointer to **time.Time** | User creation date | [optional] [readonly] 
**CustomerUUID** | Pointer to **string** | Customer UUID | [optional] [readonly] 
**Email** | **string** | User email address | 
**IsPrimary** | Pointer to **bool** | True if the user is the primary user | [optional] 
**LdapSpecifiedRole** | Pointer to **bool** | LDAP Specified Role | [optional] 
**Role** | Pointer to **string** | User role | [optional] 
**Timezone** | Pointer to **string** | User timezone | [optional] 
**UserType** | Pointer to **string** | User Type | [optional] 
**Uuid** | Pointer to **string** | User UUID | [optional] [readonly] 

## Methods

### NewUsers

`func NewUsers(email string, ) *Users`

NewUsers instantiates a new Users object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersWithDefaults

`func NewUsersWithDefaults() *Users`

NewUsersWithDefaults instantiates a new Users object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthTokenIssueDate

`func (o *Users) GetAuthTokenIssueDate() time.Time`

GetAuthTokenIssueDate returns the AuthTokenIssueDate field if non-nil, zero value otherwise.

### GetAuthTokenIssueDateOk

`func (o *Users) GetAuthTokenIssueDateOk() (*time.Time, bool)`

GetAuthTokenIssueDateOk returns a tuple with the AuthTokenIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTokenIssueDate

`func (o *Users) SetAuthTokenIssueDate(v time.Time)`

SetAuthTokenIssueDate sets AuthTokenIssueDate field to given value.

### HasAuthTokenIssueDate

`func (o *Users) HasAuthTokenIssueDate() bool`

HasAuthTokenIssueDate returns a boolean if a field has been set.

### GetCreationDate

`func (o *Users) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Users) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Users) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Users) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetCustomerUUID

`func (o *Users) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *Users) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *Users) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.

### HasCustomerUUID

`func (o *Users) HasCustomerUUID() bool`

HasCustomerUUID returns a boolean if a field has been set.

### GetEmail

`func (o *Users) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Users) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Users) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetIsPrimary

`func (o *Users) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *Users) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *Users) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *Users) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetLdapSpecifiedRole

`func (o *Users) GetLdapSpecifiedRole() bool`

GetLdapSpecifiedRole returns the LdapSpecifiedRole field if non-nil, zero value otherwise.

### GetLdapSpecifiedRoleOk

`func (o *Users) GetLdapSpecifiedRoleOk() (*bool, bool)`

GetLdapSpecifiedRoleOk returns a tuple with the LdapSpecifiedRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapSpecifiedRole

`func (o *Users) SetLdapSpecifiedRole(v bool)`

SetLdapSpecifiedRole sets LdapSpecifiedRole field to given value.

### HasLdapSpecifiedRole

`func (o *Users) HasLdapSpecifiedRole() bool`

HasLdapSpecifiedRole returns a boolean if a field has been set.

### GetRole

`func (o *Users) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Users) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Users) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Users) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTimezone

`func (o *Users) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Users) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Users) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Users) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetUserType

`func (o *Users) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *Users) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *Users) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *Users) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetUuid

`func (o *Users) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Users) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Users) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Users) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


