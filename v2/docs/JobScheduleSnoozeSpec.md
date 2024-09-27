# JobScheduleSnoozeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnoozeSecs** | **int64** | Schedule the next job after the sum of snooze time and interval from now. | 

## Methods

### NewJobScheduleSnoozeSpec

`func NewJobScheduleSnoozeSpec(snoozeSecs int64, ) *JobScheduleSnoozeSpec`

NewJobScheduleSnoozeSpec instantiates a new JobScheduleSnoozeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobScheduleSnoozeSpecWithDefaults

`func NewJobScheduleSnoozeSpecWithDefaults() *JobScheduleSnoozeSpec`

NewJobScheduleSnoozeSpecWithDefaults instantiates a new JobScheduleSnoozeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnoozeSecs

`func (o *JobScheduleSnoozeSpec) GetSnoozeSecs() int64`

GetSnoozeSecs returns the SnoozeSecs field if non-nil, zero value otherwise.

### GetSnoozeSecsOk

`func (o *JobScheduleSnoozeSpec) GetSnoozeSecsOk() (*int64, bool)`

GetSnoozeSecsOk returns a tuple with the SnoozeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnoozeSecs

`func (o *JobScheduleSnoozeSpec) SetSnoozeSecs(v int64)`

SetSnoozeSecs sets SnoozeSecs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


