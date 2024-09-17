# RoleResourceDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceGroup** | Pointer to [**ResourceGroup**](ResourceGroup.md) |  | [optional] 
**RoleUUID** | **string** | UUID of the role to attach resource group to. | 

## Methods

### NewRoleResourceDefinition

`func NewRoleResourceDefinition(roleUUID string, ) *RoleResourceDefinition`

NewRoleResourceDefinition instantiates a new RoleResourceDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleResourceDefinitionWithDefaults

`func NewRoleResourceDefinitionWithDefaults() *RoleResourceDefinition`

NewRoleResourceDefinitionWithDefaults instantiates a new RoleResourceDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### GetRoleUUID

`func (o *RoleResourceDefinition) GetRoleUUID() string`

GetRoleUUID returns the RoleUUID field if non-nil, zero value otherwise.

### GetRoleUUIDOk

`func (o *RoleResourceDefinition) GetRoleUUIDOk() (*string, bool)`

GetRoleUUIDOk returns a tuple with the RoleUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleUUID

`func (o *RoleResourceDefinition) SetRoleUUID(v string)`

SetRoleUUID sets RoleUUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


