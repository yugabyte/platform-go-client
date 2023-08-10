# ResourceDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAll** | Pointer to **bool** | Select all resources (including future resources) | [optional] 
**ResourceType** | Pointer to **string** | Resource Type | [optional] 
**ResourceUUIDSet** | Pointer to **[]string** | Set of resource uuids | [optional] 

## Methods

### NewResourceDefinition

`func NewResourceDefinition() *ResourceDefinition`

NewResourceDefinition instantiates a new ResourceDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceDefinitionWithDefaults

`func NewResourceDefinitionWithDefaults() *ResourceDefinition`

NewResourceDefinitionWithDefaults instantiates a new ResourceDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAll

`func (o *ResourceDefinition) GetAllowAll() bool`

GetAllowAll returns the AllowAll field if non-nil, zero value otherwise.

### GetAllowAllOk

`func (o *ResourceDefinition) GetAllowAllOk() (*bool, bool)`

GetAllowAllOk returns a tuple with the AllowAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAll

`func (o *ResourceDefinition) SetAllowAll(v bool)`

SetAllowAll sets AllowAll field to given value.

### HasAllowAll

`func (o *ResourceDefinition) HasAllowAll() bool`

HasAllowAll returns a boolean if a field has been set.

### GetResourceType

`func (o *ResourceDefinition) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceDefinition) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceDefinition) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourceDefinition) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResourceUUIDSet

`func (o *ResourceDefinition) GetResourceUUIDSet() []string`

GetResourceUUIDSet returns the ResourceUUIDSet field if non-nil, zero value otherwise.

### GetResourceUUIDSetOk

`func (o *ResourceDefinition) GetResourceUUIDSetOk() (*[]string, bool)`

GetResourceUUIDSetOk returns a tuple with the ResourceUUIDSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUUIDSet

`func (o *ResourceDefinition) SetResourceUUIDSet(v []string)`

SetResourceUUIDSet sets ResourceUUIDSet field to given value.

### HasResourceUUIDSet

`func (o *ResourceDefinition) HasResourceUUIDSet() bool`

HasResourceUUIDSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


