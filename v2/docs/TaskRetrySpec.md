# TaskRetrySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Optional human-readable reason for the retry, recorded in the audit log. | [optional] 

## Methods

### NewTaskRetrySpec

`func NewTaskRetrySpec() *TaskRetrySpec`

NewTaskRetrySpec instantiates a new TaskRetrySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskRetrySpecWithDefaults

`func NewTaskRetrySpecWithDefaults() *TaskRetrySpec`

NewTaskRetrySpecWithDefaults instantiates a new TaskRetrySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *TaskRetrySpec) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TaskRetrySpec) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TaskRetrySpec) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TaskRetrySpec) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


