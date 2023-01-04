# EditAccessKeyRotationScheduleParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchedulingFrequencyDays** | Pointer to **int32** | Frequency of the schedule in days | [optional] 
**Status** | Pointer to **string** | State of the schedule | [optional] 

## Methods

### NewEditAccessKeyRotationScheduleParams

`func NewEditAccessKeyRotationScheduleParams() *EditAccessKeyRotationScheduleParams`

NewEditAccessKeyRotationScheduleParams instantiates a new EditAccessKeyRotationScheduleParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditAccessKeyRotationScheduleParamsWithDefaults

`func NewEditAccessKeyRotationScheduleParamsWithDefaults() *EditAccessKeyRotationScheduleParams`

NewEditAccessKeyRotationScheduleParamsWithDefaults instantiates a new EditAccessKeyRotationScheduleParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedulingFrequencyDays

`func (o *EditAccessKeyRotationScheduleParams) GetSchedulingFrequencyDays() int32`

GetSchedulingFrequencyDays returns the SchedulingFrequencyDays field if non-nil, zero value otherwise.

### GetSchedulingFrequencyDaysOk

`func (o *EditAccessKeyRotationScheduleParams) GetSchedulingFrequencyDaysOk() (*int32, bool)`

GetSchedulingFrequencyDaysOk returns a tuple with the SchedulingFrequencyDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingFrequencyDays

`func (o *EditAccessKeyRotationScheduleParams) SetSchedulingFrequencyDays(v int32)`

SetSchedulingFrequencyDays sets SchedulingFrequencyDays field to given value.

### HasSchedulingFrequencyDays

`func (o *EditAccessKeyRotationScheduleParams) HasSchedulingFrequencyDays() bool`

HasSchedulingFrequencyDays returns a boolean if a field has been set.

### GetStatus

`func (o *EditAccessKeyRotationScheduleParams) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EditAccessKeyRotationScheduleParams) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EditAccessKeyRotationScheduleParams) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EditAccessKeyRotationScheduleParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


