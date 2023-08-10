# PermissionInfoIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 

## Methods

### NewPermissionInfoIdentifier

`func NewPermissionInfoIdentifier() *PermissionInfoIdentifier`

NewPermissionInfoIdentifier instantiates a new PermissionInfoIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionInfoIdentifierWithDefaults

`func NewPermissionInfoIdentifierWithDefaults() *PermissionInfoIdentifier`

NewPermissionInfoIdentifierWithDefaults instantiates a new PermissionInfoIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *PermissionInfoIdentifier) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *PermissionInfoIdentifier) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *PermissionInfoIdentifier) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *PermissionInfoIdentifier) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetResourceType

`func (o *PermissionInfoIdentifier) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *PermissionInfoIdentifier) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *PermissionInfoIdentifier) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *PermissionInfoIdentifier) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


