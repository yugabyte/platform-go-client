# UserRegistrationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfirmPassword** | Pointer to **string** | Password confirmation | [optional] 
**Email** | **string** | Email address | 
**Features** | Pointer to **map[string]map[string]interface{}** | User features | [optional] 
**Password** | Pointer to **string** | Password | [optional] 
**Role** | Pointer to **string** | &lt;b style&#x3D;\&quot;color:#ff0000\&quot;&gt;Deprecated since YBA version 2.19.3.0.&lt;/b&gt; Use field roleResourceDefinitions instead. | [optional] 
**RoleResourceDefinitions** | Pointer to [**[]RoleResourceDefinition**](RoleResourceDefinition.md) | List of roles and resource groups defined for user. | [optional] 
**Timezone** | Pointer to **string** | User timezone | [optional] 

## Methods

### NewUserRegistrationData

`func NewUserRegistrationData(email string, ) *UserRegistrationData`

NewUserRegistrationData instantiates a new UserRegistrationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRegistrationDataWithDefaults

`func NewUserRegistrationDataWithDefaults() *UserRegistrationData`

NewUserRegistrationDataWithDefaults instantiates a new UserRegistrationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmPassword

`func (o *UserRegistrationData) GetConfirmPassword() string`

GetConfirmPassword returns the ConfirmPassword field if non-nil, zero value otherwise.

### GetConfirmPasswordOk

`func (o *UserRegistrationData) GetConfirmPasswordOk() (*string, bool)`

GetConfirmPasswordOk returns a tuple with the ConfirmPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmPassword

`func (o *UserRegistrationData) SetConfirmPassword(v string)`

SetConfirmPassword sets ConfirmPassword field to given value.

### HasConfirmPassword

`func (o *UserRegistrationData) HasConfirmPassword() bool`

HasConfirmPassword returns a boolean if a field has been set.

### GetEmail

`func (o *UserRegistrationData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserRegistrationData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserRegistrationData) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFeatures

`func (o *UserRegistrationData) GetFeatures() map[string]map[string]interface{}`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *UserRegistrationData) GetFeaturesOk() (*map[string]map[string]interface{}, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *UserRegistrationData) SetFeatures(v map[string]map[string]interface{})`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *UserRegistrationData) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetPassword

`func (o *UserRegistrationData) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserRegistrationData) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserRegistrationData) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserRegistrationData) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRole

`func (o *UserRegistrationData) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserRegistrationData) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserRegistrationData) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserRegistrationData) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRoleResourceDefinitions

`func (o *UserRegistrationData) GetRoleResourceDefinitions() []RoleResourceDefinition`

GetRoleResourceDefinitions returns the RoleResourceDefinitions field if non-nil, zero value otherwise.

### GetRoleResourceDefinitionsOk

`func (o *UserRegistrationData) GetRoleResourceDefinitionsOk() (*[]RoleResourceDefinition, bool)`

GetRoleResourceDefinitionsOk returns a tuple with the RoleResourceDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleResourceDefinitions

`func (o *UserRegistrationData) SetRoleResourceDefinitions(v []RoleResourceDefinition)`

SetRoleResourceDefinitions sets RoleResourceDefinitions field to given value.

### HasRoleResourceDefinitions

`func (o *UserRegistrationData) HasRoleResourceDefinitions() bool`

HasRoleResourceDefinitions returns a boolean if a field has been set.

### GetTimezone

`func (o *UserRegistrationData) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UserRegistrationData) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UserRegistrationData) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UserRegistrationData) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


