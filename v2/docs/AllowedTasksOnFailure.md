# AllowedTasksOnFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Restricted** | Pointer to **bool** | Whether the Universe is restricted from performing other operations | [optional] [readonly] 
**TaskTypes** | Pointer to **[]string** | If restricted, this is the list of task types that are allowed to be retried on this Universe. This  is in the format &lt;taskType&gt;_&lt;targetResource&gt;. Eg. Create_Backup, Delete_Universe, etc. | [optional] [readonly] 

## Methods

### NewAllowedTasksOnFailure

`func NewAllowedTasksOnFailure() *AllowedTasksOnFailure`

NewAllowedTasksOnFailure instantiates a new AllowedTasksOnFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedTasksOnFailureWithDefaults

`func NewAllowedTasksOnFailureWithDefaults() *AllowedTasksOnFailure`

NewAllowedTasksOnFailureWithDefaults instantiates a new AllowedTasksOnFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestricted

`func (o *AllowedTasksOnFailure) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *AllowedTasksOnFailure) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *AllowedTasksOnFailure) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *AllowedTasksOnFailure) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### GetTaskTypes

`func (o *AllowedTasksOnFailure) GetTaskTypes() []string`

GetTaskTypes returns the TaskTypes field if non-nil, zero value otherwise.

### GetTaskTypesOk

`func (o *AllowedTasksOnFailure) GetTaskTypesOk() (*[]string, bool)`

GetTaskTypesOk returns a tuple with the TaskTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskTypes

`func (o *AllowedTasksOnFailure) SetTaskTypes(v []string)`

SetTaskTypes sets TaskTypes field to given value.

### HasTaskTypes

`func (o *AllowedTasksOnFailure) HasTaskTypes() bool`

HasTaskTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


