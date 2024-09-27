# JobScheduleApiFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigClass** | Pointer to **string** | Filter by config class of the job. | [optional] 
**EnabledOnly** | Pointer to **bool** | Filter out disabled job schedules if true. | [optional] 
**NameRegex** | Pointer to **string** | Filter by name regex. | [optional] 
**NextStartWindowSecs** | Pointer to **int64** | Filter by next start time window from now. | [optional] 
**Type** | Pointer to [**JobScheduleType**](JobScheduleType.md) |  | [optional] 

## Methods

### NewJobScheduleApiFilter

`func NewJobScheduleApiFilter() *JobScheduleApiFilter`

NewJobScheduleApiFilter instantiates a new JobScheduleApiFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobScheduleApiFilterWithDefaults

`func NewJobScheduleApiFilterWithDefaults() *JobScheduleApiFilter`

NewJobScheduleApiFilterWithDefaults instantiates a new JobScheduleApiFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigClass

`func (o *JobScheduleApiFilter) GetConfigClass() string`

GetConfigClass returns the ConfigClass field if non-nil, zero value otherwise.

### GetConfigClassOk

`func (o *JobScheduleApiFilter) GetConfigClassOk() (*string, bool)`

GetConfigClassOk returns a tuple with the ConfigClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigClass

`func (o *JobScheduleApiFilter) SetConfigClass(v string)`

SetConfigClass sets ConfigClass field to given value.

### HasConfigClass

`func (o *JobScheduleApiFilter) HasConfigClass() bool`

HasConfigClass returns a boolean if a field has been set.

### GetEnabledOnly

`func (o *JobScheduleApiFilter) GetEnabledOnly() bool`

GetEnabledOnly returns the EnabledOnly field if non-nil, zero value otherwise.

### GetEnabledOnlyOk

`func (o *JobScheduleApiFilter) GetEnabledOnlyOk() (*bool, bool)`

GetEnabledOnlyOk returns a tuple with the EnabledOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledOnly

`func (o *JobScheduleApiFilter) SetEnabledOnly(v bool)`

SetEnabledOnly sets EnabledOnly field to given value.

### HasEnabledOnly

`func (o *JobScheduleApiFilter) HasEnabledOnly() bool`

HasEnabledOnly returns a boolean if a field has been set.

### GetNameRegex

`func (o *JobScheduleApiFilter) GetNameRegex() string`

GetNameRegex returns the NameRegex field if non-nil, zero value otherwise.

### GetNameRegexOk

`func (o *JobScheduleApiFilter) GetNameRegexOk() (*string, bool)`

GetNameRegexOk returns a tuple with the NameRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameRegex

`func (o *JobScheduleApiFilter) SetNameRegex(v string)`

SetNameRegex sets NameRegex field to given value.

### HasNameRegex

`func (o *JobScheduleApiFilter) HasNameRegex() bool`

HasNameRegex returns a boolean if a field has been set.

### GetNextStartWindowSecs

`func (o *JobScheduleApiFilter) GetNextStartWindowSecs() int64`

GetNextStartWindowSecs returns the NextStartWindowSecs field if non-nil, zero value otherwise.

### GetNextStartWindowSecsOk

`func (o *JobScheduleApiFilter) GetNextStartWindowSecsOk() (*int64, bool)`

GetNextStartWindowSecsOk returns a tuple with the NextStartWindowSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStartWindowSecs

`func (o *JobScheduleApiFilter) SetNextStartWindowSecs(v int64)`

SetNextStartWindowSecs sets NextStartWindowSecs field to given value.

### HasNextStartWindowSecs

`func (o *JobScheduleApiFilter) HasNextStartWindowSecs() bool`

HasNextStartWindowSecs returns a boolean if a field has been set.

### GetType

`func (o *JobScheduleApiFilter) GetType() JobScheduleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobScheduleApiFilter) GetTypeOk() (*JobScheduleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobScheduleApiFilter) SetType(v JobScheduleType)`

SetType sets Type field to given value.

### HasType

`func (o *JobScheduleApiFilter) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


