# UserSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | User email address | 
**Role** | Pointer to **string** | User role | [optional] 

## Methods

### NewUserSpec

`func NewUserSpec(email string, ) *UserSpec`

NewUserSpec instantiates a new UserSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSpecWithDefaults

`func NewUserSpecWithDefaults() *UserSpec`

NewUserSpecWithDefaults instantiates a new UserSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserSpec) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserSpec) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserSpec) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *UserSpec) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserSpec) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserSpec) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserSpec) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


