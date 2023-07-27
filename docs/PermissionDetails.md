# PermissionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionList** | Pointer to [**[]PermissionInfoIdentifier**](PermissionInfoIdentifier.md) | Set of permissions | [optional] 

## Methods

### NewPermissionDetails

`func NewPermissionDetails() *PermissionDetails`

NewPermissionDetails instantiates a new PermissionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionDetailsWithDefaults

`func NewPermissionDetailsWithDefaults() *PermissionDetails`

NewPermissionDetailsWithDefaults instantiates a new PermissionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionList

`func (o *PermissionDetails) GetPermissionList() []PermissionInfoIdentifier`

GetPermissionList returns the PermissionList field if non-nil, zero value otherwise.

### GetPermissionListOk

`func (o *PermissionDetails) GetPermissionListOk() (*[]PermissionInfoIdentifier, bool)`

GetPermissionListOk returns a tuple with the PermissionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionList

`func (o *PermissionDetails) SetPermissionList(v []PermissionInfoIdentifier)`

SetPermissionList sets PermissionList field to given value.

### HasPermissionList

`func (o *PermissionDetails) HasPermissionList() bool`

HasPermissionList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


