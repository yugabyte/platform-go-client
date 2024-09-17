# PermissionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PermissionValidOnResource** | Pointer to **bool** |  | [optional] 
**PrerequisitePermissions** | Pointer to [**[]Permission**](Permission.md) |  | [optional] 
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

### GetAction

`func (o *PermissionInfo) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PermissionInfo) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PermissionInfo) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PermissionInfo) HasAction() bool`

HasAction returns a boolean if a field has been set.

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

### GetName

`func (o *PermissionInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PermissionInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissionValidOnResource

`func (o *PermissionInfo) GetPermissionValidOnResource() bool`

GetPermissionValidOnResource returns the PermissionValidOnResource field if non-nil, zero value otherwise.

### GetPermissionValidOnResourceOk

`func (o *PermissionInfo) GetPermissionValidOnResourceOk() (*bool, bool)`

GetPermissionValidOnResourceOk returns a tuple with the PermissionValidOnResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionValidOnResource

`func (o *PermissionInfo) SetPermissionValidOnResource(v bool)`

SetPermissionValidOnResource sets PermissionValidOnResource field to given value.

### HasPermissionValidOnResource

`func (o *PermissionInfo) HasPermissionValidOnResource() bool`

HasPermissionValidOnResource returns a boolean if a field has been set.

### GetPrerequisitePermissions

`func (o *PermissionInfo) GetPrerequisitePermissions() []Permission`

GetPrerequisitePermissions returns the PrerequisitePermissions field if non-nil, zero value otherwise.

### GetPrerequisitePermissionsOk

`func (o *PermissionInfo) GetPrerequisitePermissionsOk() (*[]Permission, bool)`

GetPrerequisitePermissionsOk returns a tuple with the PrerequisitePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisitePermissions

`func (o *PermissionInfo) SetPrerequisitePermissions(v []Permission)`

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


