# EditBackupScheduleParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpression** | Pointer to **string** | Cron expression for scheduling | [optional] 
**Frequency** | Pointer to **int64** | Frequency of the schedule | [optional] 
**FrequencyTimeUnit** | Pointer to **string** | Time Unit for frequency | [optional] 
**IncrementalBackupFrequency** | Pointer to **int64** | Frequency of incremental backup schedule | [optional] 
**IncrementalBackupFrequencyTimeUnit** | Pointer to **string** | TimeUnit for incremental Backup Schedule frequency | [optional] 
**Status** | Pointer to **string** | State of the schedule | [optional] 

## Methods

### NewEditBackupScheduleParams

`func NewEditBackupScheduleParams() *EditBackupScheduleParams`

NewEditBackupScheduleParams instantiates a new EditBackupScheduleParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditBackupScheduleParamsWithDefaults

`func NewEditBackupScheduleParamsWithDefaults() *EditBackupScheduleParams`

NewEditBackupScheduleParamsWithDefaults instantiates a new EditBackupScheduleParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpression

`func (o *EditBackupScheduleParams) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *EditBackupScheduleParams) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *EditBackupScheduleParams) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *EditBackupScheduleParams) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetFrequency

`func (o *EditBackupScheduleParams) GetFrequency() int64`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *EditBackupScheduleParams) GetFrequencyOk() (*int64, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *EditBackupScheduleParams) SetFrequency(v int64)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *EditBackupScheduleParams) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetFrequencyTimeUnit

`func (o *EditBackupScheduleParams) GetFrequencyTimeUnit() string`

GetFrequencyTimeUnit returns the FrequencyTimeUnit field if non-nil, zero value otherwise.

### GetFrequencyTimeUnitOk

`func (o *EditBackupScheduleParams) GetFrequencyTimeUnitOk() (*string, bool)`

GetFrequencyTimeUnitOk returns a tuple with the FrequencyTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyTimeUnit

`func (o *EditBackupScheduleParams) SetFrequencyTimeUnit(v string)`

SetFrequencyTimeUnit sets FrequencyTimeUnit field to given value.

### HasFrequencyTimeUnit

`func (o *EditBackupScheduleParams) HasFrequencyTimeUnit() bool`

HasFrequencyTimeUnit returns a boolean if a field has been set.

### GetIncrementalBackupFrequency

`func (o *EditBackupScheduleParams) GetIncrementalBackupFrequency() int64`

GetIncrementalBackupFrequency returns the IncrementalBackupFrequency field if non-nil, zero value otherwise.

### GetIncrementalBackupFrequencyOk

`func (o *EditBackupScheduleParams) GetIncrementalBackupFrequencyOk() (*int64, bool)`

GetIncrementalBackupFrequencyOk returns a tuple with the IncrementalBackupFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalBackupFrequency

`func (o *EditBackupScheduleParams) SetIncrementalBackupFrequency(v int64)`

SetIncrementalBackupFrequency sets IncrementalBackupFrequency field to given value.

### HasIncrementalBackupFrequency

`func (o *EditBackupScheduleParams) HasIncrementalBackupFrequency() bool`

HasIncrementalBackupFrequency returns a boolean if a field has been set.

### GetIncrementalBackupFrequencyTimeUnit

`func (o *EditBackupScheduleParams) GetIncrementalBackupFrequencyTimeUnit() string`

GetIncrementalBackupFrequencyTimeUnit returns the IncrementalBackupFrequencyTimeUnit field if non-nil, zero value otherwise.

### GetIncrementalBackupFrequencyTimeUnitOk

`func (o *EditBackupScheduleParams) GetIncrementalBackupFrequencyTimeUnitOk() (*string, bool)`

GetIncrementalBackupFrequencyTimeUnitOk returns a tuple with the IncrementalBackupFrequencyTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalBackupFrequencyTimeUnit

`func (o *EditBackupScheduleParams) SetIncrementalBackupFrequencyTimeUnit(v string)`

SetIncrementalBackupFrequencyTimeUnit sets IncrementalBackupFrequencyTimeUnit field to given value.

### HasIncrementalBackupFrequencyTimeUnit

`func (o *EditBackupScheduleParams) HasIncrementalBackupFrequencyTimeUnit() bool`

HasIncrementalBackupFrequencyTimeUnit returns a boolean if a field has been set.

### GetStatus

`func (o *EditBackupScheduleParams) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EditBackupScheduleParams) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EditBackupScheduleParams) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EditBackupScheduleParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


