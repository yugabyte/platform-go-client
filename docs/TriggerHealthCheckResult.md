# TriggerHealthCheckResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | The ISO-8601 timestamp when the health check was triggered. | [optional] 

## Methods

### NewTriggerHealthCheckResult

`func NewTriggerHealthCheckResult() *TriggerHealthCheckResult`

NewTriggerHealthCheckResult instantiates a new TriggerHealthCheckResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerHealthCheckResultWithDefaults

`func NewTriggerHealthCheckResultWithDefaults() *TriggerHealthCheckResult`

NewTriggerHealthCheckResultWithDefaults instantiates a new TriggerHealthCheckResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *TriggerHealthCheckResult) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TriggerHealthCheckResult) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TriggerHealthCheckResult) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TriggerHealthCheckResult) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


