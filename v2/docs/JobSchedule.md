# JobSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | Pointer to [**JobScheduleSpec**](JobScheduleSpec.md) |  | [optional] 
**Info** | Pointer to [**JobScheduleInfo**](JobScheduleInfo.md) |  | [optional] 

## Methods

### NewJobSchedule

`func NewJobSchedule() *JobSchedule`

NewJobSchedule instantiates a new JobSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobScheduleWithDefaults

`func NewJobScheduleWithDefaults() *JobSchedule`

NewJobScheduleWithDefaults instantiates a new JobSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *JobSchedule) GetSpec() JobScheduleSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *JobSchedule) GetSpecOk() (*JobScheduleSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *JobSchedule) SetSpec(v JobScheduleSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *JobSchedule) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetInfo

`func (o *JobSchedule) GetInfo() JobScheduleInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *JobSchedule) GetInfoOk() (*JobScheduleInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *JobSchedule) SetInfo(v JobScheduleInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *JobSchedule) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


