# UserWithFeatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthTokenIssueDate** | Pointer to **time.Time** | API token creation date | [optional] [readonly] 
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

### NewUserWithFeatures

`func NewUserWithFeatures(email string, ) *UserWithFeatures`

NewUserWithFeatures instantiates a new UserWithFeatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithFeaturesWithDefaults

`func NewUserWithFeaturesWithDefaults() *UserWithFeatures`

NewUserWithFeaturesWithDefaults instantiates a new UserWithFeatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthTokenIssueDate

`func (o *UserWithFeatures) GetAuthTokenIssueDate() time.Time`

GetAuthTokenIssueDate returns the AuthTokenIssueDate field if non-nil, zero value otherwise.

### GetAuthTokenIssueDateOk

`func (o *UserWithFeatures) GetAuthTokenIssueDateOk() (*time.Time, bool)`

GetAuthTokenIssueDateOk returns a tuple with the AuthTokenIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTokenIssueDate

`func (o *UserWithFeatures) SetAuthTokenIssueDate(v time.Time)`

SetAuthTokenIssueDate sets AuthTokenIssueDate field to given value.

### HasAuthTokenIssueDate

`func (o *UserWithFeatures) HasAuthTokenIssueDate() bool`

HasAuthTokenIssueDate returns a boolean if a field has been set.

### GetCreationDate

`func (o *UserWithFeatures) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *UserWithFeatures) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *UserWithFeatures) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *UserWithFeatures) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetCustomerUUID

`func (o *UserWithFeatures) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *UserWithFeatures) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *UserWithFeatures) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.

### HasCustomerUUID

`func (o *UserWithFeatures) HasCustomerUUID() bool`

HasCustomerUUID returns a boolean if a field has been set.

### GetEmail

`func (o *UserWithFeatures) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserWithFeatures) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserWithFeatures) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetIsPrimary

`func (o *UserWithFeatures) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *UserWithFeatures) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *UserWithFeatures) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *UserWithFeatures) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetLdapSpecifiedRole

`func (o *UserWithFeatures) GetLdapSpecifiedRole() bool`

GetLdapSpecifiedRole returns the LdapSpecifiedRole field if non-nil, zero value otherwise.

### GetLdapSpecifiedRoleOk

`func (o *UserWithFeatures) GetLdapSpecifiedRoleOk() (*bool, bool)`

GetLdapSpecifiedRoleOk returns a tuple with the LdapSpecifiedRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapSpecifiedRole

`func (o *UserWithFeatures) SetLdapSpecifiedRole(v bool)`

SetLdapSpecifiedRole sets LdapSpecifiedRole field to given value.

### HasLdapSpecifiedRole

`func (o *UserWithFeatures) HasLdapSpecifiedRole() bool`

HasLdapSpecifiedRole returns a boolean if a field has been set.

### GetRole

`func (o *UserWithFeatures) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserWithFeatures) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserWithFeatures) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserWithFeatures) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTimezone

`func (o *UserWithFeatures) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UserWithFeatures) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UserWithFeatures) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UserWithFeatures) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetUserType

`func (o *UserWithFeatures) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UserWithFeatures) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UserWithFeatures) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *UserWithFeatures) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetUuid

`func (o *UserWithFeatures) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UserWithFeatures) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UserWithFeatures) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UserWithFeatures) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


