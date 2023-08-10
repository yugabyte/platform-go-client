# PermissionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Permission** | Pointer to **string** |  | [optional] 
**PrerequisitePermissions** | Pointer to [**[]PermissionInfoIdentifier**](PermissionInfoIdentifier.md) |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 

## Methods

### NewPermissionInfo

`func NewPermissionInfo() *PermissionInfo`

NewPermissionInfo instantiates a new PermissionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionInfoWithDefaults

`func NewPermissionInfoWithDefaults() *PermissionInfo`

NewPermissionInfoWithDefaults instantiates a new PermissionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PermissionInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PermissionInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PermissionInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PermissionInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPermission

`func (o *PermissionInfo) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *PermissionInfo) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *PermissionInfo) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *PermissionInfo) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetPrerequisitePermissions

`func (o *PermissionInfo) GetPrerequisitePermissions() []PermissionInfoIdentifier`

GetPrerequisitePermissions returns the PrerequisitePermissions field if non-nil, zero value otherwise.

### GetPrerequisitePermissionsOk

`func (o *PermissionInfo) GetPrerequisitePermissionsOk() (*[]PermissionInfoIdentifier, bool)`

GetPrerequisitePermissionsOk returns a tuple with the PrerequisitePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisitePermissions

`func (o *PermissionInfo) SetPrerequisitePermissions(v []PermissionInfoIdentifier)`

SetPrerequisitePermissions sets PrerequisitePermissions field to given value.

### HasPrerequisitePermissions

`func (o *PermissionInfo) HasPrerequisitePermissions() bool`

HasPrerequisitePermissions returns a boolean if a field has been set.

### GetResourceType

`func (o *PermissionInfo) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *PermissionInfo) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *PermissionInfo) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *PermissionInfo) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


