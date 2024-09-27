# ResourceDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | Resource Type | 
**AllowAll** | **bool** | Select all resources (including future resources) | 
**ResourceUuidSet** | Pointer to **[]string** | Set of resource UUIDs | [optional] [default to []]

## Methods

### NewResourceDefinition

`func NewResourceDefinition(resourceType string, allowAll bool, ) *ResourceDefinition`

NewResourceDefinition instantiates a new ResourceDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceDefinitionWithDefaults

`func NewResourceDefinitionWithDefaults() *ResourceDefinition`

NewResourceDefinitionWithDefaults instantiates a new ResourceDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetResourceUuidSet

`func (o *ResourceDefinition) GetResourceUuidSet() []string`

GetResourceUuidSet returns the ResourceUuidSet field if non-nil, zero value otherwise.

### GetResourceUuidSetOk

`func (o *ResourceDefinition) GetResourceUuidSetOk() (*[]string, bool)`

GetResourceUuidSetOk returns a tuple with the ResourceUuidSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUuidSet

`func (o *ResourceDefinition) SetResourceUuidSet(v []string)`

SetResourceUuidSet sets ResourceUuidSet field to given value.

### HasResourceUuidSet

`func (o *ResourceDefinition) HasResourceUuidSet() bool`

HasResourceUuidSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


