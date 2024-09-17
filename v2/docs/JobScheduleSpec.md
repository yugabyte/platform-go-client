# JobScheduleSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Job schedule name | [readonly] 
**JobConfig** | [**JobConfigSpec**](JobConfigSpec.md) |  | 
**ScheduleConfig** | [**JobScheduleConfigSpec**](JobScheduleConfigSpec.md) |  | 

## Methods

### NewJobScheduleSpec

`func NewJobScheduleSpec(name string, jobConfig JobConfigSpec, scheduleConfig JobScheduleConfigSpec, ) *JobScheduleSpec`

NewJobScheduleSpec instantiates a new JobScheduleSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobScheduleSpecWithDefaults

`func NewJobScheduleSpecWithDefaults() *JobScheduleSpec`

NewJobScheduleSpecWithDefaults instantiates a new JobScheduleSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JobScheduleSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobScheduleSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobScheduleSpec) SetName(v string)`

SetName sets Name field to given value.


### GetJobConfig

`func (o *JobScheduleSpec) GetJobConfig() JobConfigSpec`

GetJobConfig returns the JobConfig field if non-nil, zero value otherwise.

### GetJobConfigOk

`func (o *JobScheduleSpec) GetJobConfigOk() (*JobConfigSpec, bool)`

GetJobConfigOk returns a tuple with the JobConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobConfig

`func (o *JobScheduleSpec) SetJobConfig(v JobConfigSpec)`

SetJobConfig sets JobConfig field to given value.


### GetScheduleConfig

`func (o *JobScheduleSpec) GetScheduleConfig() JobScheduleConfigSpec`

GetScheduleConfig returns the ScheduleConfig field if non-nil, zero value otherwise.

### GetScheduleConfigOk

`func (o *JobScheduleSpec) GetScheduleConfigOk() (*JobScheduleConfigSpec, bool)`

GetScheduleConfigOk returns a tuple with the ScheduleConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleConfig

`func (o *JobScheduleSpec) SetScheduleConfig(v JobScheduleConfigSpec)`

SetScheduleConfig sets ScheduleConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


