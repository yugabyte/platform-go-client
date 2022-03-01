# EditBackupScheduleParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpression** | Pointer to **string** | Cron expression for scheduling | [optional] 
**Frequency** | Pointer to **int64** | Frequency of the schedule | [optional] 
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


