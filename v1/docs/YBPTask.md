# YBPTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceUUID** | Pointer to **string** | UUID of the resource being modified by the task | [optional] [readonly] 
**TaskUUID** | Pointer to **string** | Task UUID | [optional] [readonly] 

## Methods

### NewYBPTask

`func NewYBPTask() *YBPTask`

NewYBPTask instantiates a new YBPTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYBPTaskWithDefaults

`func NewYBPTaskWithDefaults() *YBPTask`

NewYBPTaskWithDefaults instantiates a new YBPTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceUUID

`func (o *YBPTask) GetResourceUUID() string`

GetResourceUUID returns the ResourceUUID field if non-nil, zero value otherwise.

### GetResourceUUIDOk

`func (o *YBPTask) GetResourceUUIDOk() (*string, bool)`

GetResourceUUIDOk returns a tuple with the ResourceUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUUID

`func (o *YBPTask) SetResourceUUID(v string)`

SetResourceUUID sets ResourceUUID field to given value.

### HasResourceUUID

`func (o *YBPTask) HasResourceUUID() bool`

HasResourceUUID returns a boolean if a field has been set.

### GetTaskUUID

`func (o *YBPTask) GetTaskUUID() string`

GetTaskUUID returns the TaskUUID field if non-nil, zero value otherwise.

### GetTaskUUIDOk

`func (o *YBPTask) GetTaskUUIDOk() (*string, bool)`

GetTaskUUIDOk returns a tuple with the TaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUUID

`func (o *YBPTask) SetTaskUUID(v string)`

SetTaskUUID sets TaskUUID field to given value.

### HasTaskUUID

`func (o *YBPTask) HasTaskUUID() bool`

HasTaskUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


