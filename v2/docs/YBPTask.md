# YBPTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceUuid** | Pointer to **string** | UUID of the resource being modified by the task | [optional] [readonly] 
**TaskUuid** | Pointer to **string** | Task UUID | [optional] [readonly] 

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

### GetResourceUuid

`func (o *YBPTask) GetResourceUuid() string`

GetResourceUuid returns the ResourceUuid field if non-nil, zero value otherwise.

### GetResourceUuidOk

`func (o *YBPTask) GetResourceUuidOk() (*string, bool)`

GetResourceUuidOk returns a tuple with the ResourceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUuid

`func (o *YBPTask) SetResourceUuid(v string)`

SetResourceUuid sets ResourceUuid field to given value.

### HasResourceUuid

`func (o *YBPTask) HasResourceUuid() bool`

HasResourceUuid returns a boolean if a field has been set.

### GetTaskUuid

`func (o *YBPTask) GetTaskUuid() string`

GetTaskUuid returns the TaskUuid field if non-nil, zero value otherwise.

### GetTaskUuidOk

`func (o *YBPTask) GetTaskUuidOk() (*string, bool)`

GetTaskUuidOk returns a tuple with the TaskUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUuid

`func (o *YBPTask) SetTaskUuid(v string)`

SetTaskUuid sets TaskUuid field to given value.

### HasTaskUuid

`func (o *YBPTask) HasTaskUuid() bool`

HasTaskUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


