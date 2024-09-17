# RoleResourceDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleUuid** | **string** | UUID of the role to attach resource group to. | 
**ResourceGroup** | Pointer to [**ResourceGroup**](ResourceGroup.md) |  | [optional] 

## Methods

### NewRoleResourceDefinition

`func NewRoleResourceDefinition(roleUuid string, ) *RoleResourceDefinition`

NewRoleResourceDefinition instantiates a new RoleResourceDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleResourceDefinitionWithDefaults

`func NewRoleResourceDefinitionWithDefaults() *RoleResourceDefinition`

NewRoleResourceDefinitionWithDefaults instantiates a new RoleResourceDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleUuid

`func (o *RoleResourceDefinition) GetRoleUuid() string`

GetRoleUuid returns the RoleUuid field if non-nil, zero value otherwise.

### GetRoleUuidOk

`func (o *RoleResourceDefinition) GetRoleUuidOk() (*string, bool)`

GetRoleUuidOk returns a tuple with the RoleUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleUuid

`func (o *RoleResourceDefinition) SetRoleUuid(v string)`

SetRoleUuid sets RoleUuid field to given value.


### GetResourceGroup

`func (o *RoleResourceDefinition) GetResourceGroup() ResourceGroup`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *RoleResourceDefinition) GetResourceGroupOk() (*ResourceGroup, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *RoleResourceDefinition) SetResourceGroup(v ResourceGroup)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *RoleResourceDefinition) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


