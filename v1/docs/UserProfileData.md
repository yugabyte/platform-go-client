# UserProfileData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfirmPassword** | Pointer to **string** | Password confirmation | [optional] 
**Password** | Pointer to **string** | Password | [optional] 
**Role** | **string** | User role | 
**Timezone** | Pointer to **string** | User timezone | [optional] 

## Methods

### NewUserProfileData

`func NewUserProfileData(role string, ) *UserProfileData`

NewUserProfileData instantiates a new UserProfileData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfileDataWithDefaults

`func NewUserProfileDataWithDefaults() *UserProfileData`

NewUserProfileDataWithDefaults instantiates a new UserProfileData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmPassword

`func (o *UserProfileData) GetConfirmPassword() string`

GetConfirmPassword returns the ConfirmPassword field if non-nil, zero value otherwise.

### GetConfirmPasswordOk

`func (o *UserProfileData) GetConfirmPasswordOk() (*string, bool)`

GetConfirmPasswordOk returns a tuple with the ConfirmPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmPassword

`func (o *UserProfileData) SetConfirmPassword(v string)`

SetConfirmPassword sets ConfirmPassword field to given value.

### HasConfirmPassword

`func (o *UserProfileData) HasConfirmPassword() bool`

HasConfirmPassword returns a boolean if a field has been set.

### GetPassword

`func (o *UserProfileData) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserProfileData) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserProfileData) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserProfileData) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRole

`func (o *UserProfileData) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserProfileData) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserProfileData) SetRole(v string)`

SetRole sets Role field to given value.


### GetTimezone

`func (o *UserProfileData) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UserProfileData) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UserProfileData) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UserProfileData) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


