# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | Pointer to [**UserSpec**](UserSpec.md) |  | [optional] 
**Info** | Pointer to [**UserInfo**](UserInfo.md) |  | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *User) GetSpec() UserSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *User) GetSpecOk() (*UserSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *User) SetSpec(v UserSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *User) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetInfo

`func (o *User) GetInfo() UserInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *User) GetInfoOk() (*UserInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *User) SetInfo(v UserInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *User) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


