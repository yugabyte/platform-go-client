# Alert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcknowledgedTime** | Pointer to **time.Time** | Timestamp at which the alert was acknowledged | [optional] [readonly] 
**ConfigurationType** | **string** | Alert configuration type | [readonly] 
**ConfigurationUuid** | **string** | Alert configuration UUID | [readonly] 
**CreateTime** | **time.Time** | Alert creation timestamp | [readonly] 
**CustomerUUID** | **string** | Customer UUID | [readonly] 
**DefinitionUuid** | **string** | Alert definition UUID | [readonly] 
**Labels** | [**[]AlertLabel**](AlertLabel.md) |  | 
**Message** | **string** | The alert&#39;s message text | [readonly] 
**Name** | **string** | The alert&#39;s name | [readonly] 
**NextNotificationTime** | Pointer to **time.Time** | Time of the next notification attempt | [optional] [readonly] 
**NotificationAttemptTime** | Pointer to **time.Time** | Time of the last notification attempt | [optional] [readonly] 
**NotificationsFailed** | Pointer to **int32** | Count of failures to send a notification | [optional] [readonly] 
**NotifiedState** | Pointer to **string** | Alert state in the last-sent notification | [optional] [readonly] 
**ResolvedTime** | Pointer to **time.Time** | Timestamp at which the alert was resolved | [optional] [readonly] 
**Severity** | **string** | Alert configuration severity | [readonly] 
**SeverityIndex** | **int32** |  | 
**SourceName** | **string** | The source of the alert | [readonly] 
**SourceUUID** | **string** | The sourceUUID of the alert | [readonly] 
**State** | **string** | The alert&#39;s state | [readonly] 
**StateIndex** | **int32** |  | 
**Uuid** | Pointer to **string** | Alert UUID | [optional] [readonly] 

## Methods

### NewAlert

`func NewAlert(configurationType string, configurationUuid string, createTime time.Time, customerUUID string, definitionUuid string, labels []AlertLabel, message string, name string, severity string, severityIndex int32, sourceName string, sourceUUID string, state string, stateIndex int32, ) *Alert`

NewAlert instantiates a new Alert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithDefaults

`func NewAlertWithDefaults() *Alert`

NewAlertWithDefaults instantiates a new Alert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledgedTime

`func (o *Alert) GetAcknowledgedTime() time.Time`

GetAcknowledgedTime returns the AcknowledgedTime field if non-nil, zero value otherwise.

### GetAcknowledgedTimeOk

`func (o *Alert) GetAcknowledgedTimeOk() (*time.Time, bool)`

GetAcknowledgedTimeOk returns a tuple with the AcknowledgedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedTime

`func (o *Alert) SetAcknowledgedTime(v time.Time)`

SetAcknowledgedTime sets AcknowledgedTime field to given value.

### HasAcknowledgedTime

`func (o *Alert) HasAcknowledgedTime() bool`

HasAcknowledgedTime returns a boolean if a field has been set.

### GetConfigurationType

`func (o *Alert) GetConfigurationType() string`

GetConfigurationType returns the ConfigurationType field if non-nil, zero value otherwise.

### GetConfigurationTypeOk

`func (o *Alert) GetConfigurationTypeOk() (*string, bool)`

GetConfigurationTypeOk returns a tuple with the ConfigurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationType

`func (o *Alert) SetConfigurationType(v string)`

SetConfigurationType sets ConfigurationType field to given value.


### GetConfigurationUuid

`func (o *Alert) GetConfigurationUuid() string`

GetConfigurationUuid returns the ConfigurationUuid field if non-nil, zero value otherwise.

### GetConfigurationUuidOk

`func (o *Alert) GetConfigurationUuidOk() (*string, bool)`

GetConfigurationUuidOk returns a tuple with the ConfigurationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationUuid

`func (o *Alert) SetConfigurationUuid(v string)`

SetConfigurationUuid sets ConfigurationUuid field to given value.


### GetCreateTime

`func (o *Alert) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Alert) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Alert) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.


### GetCustomerUUID

`func (o *Alert) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *Alert) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *Alert) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetDefinitionUuid

`func (o *Alert) GetDefinitionUuid() string`

GetDefinitionUuid returns the DefinitionUuid field if non-nil, zero value otherwise.

### GetDefinitionUuidOk

`func (o *Alert) GetDefinitionUuidOk() (*string, bool)`

GetDefinitionUuidOk returns a tuple with the DefinitionUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionUuid

`func (o *Alert) SetDefinitionUuid(v string)`

SetDefinitionUuid sets DefinitionUuid field to given value.


### GetLabels

`func (o *Alert) GetLabels() []AlertLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Alert) GetLabelsOk() (*[]AlertLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Alert) SetLabels(v []AlertLabel)`

SetLabels sets Labels field to given value.


### GetMessage

`func (o *Alert) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Alert) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Alert) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetName

`func (o *Alert) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Alert) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Alert) SetName(v string)`

SetName sets Name field to given value.


### GetNextNotificationTime

`func (o *Alert) GetNextNotificationTime() time.Time`

GetNextNotificationTime returns the NextNotificationTime field if non-nil, zero value otherwise.

### GetNextNotificationTimeOk

`func (o *Alert) GetNextNotificationTimeOk() (*time.Time, bool)`

GetNextNotificationTimeOk returns a tuple with the NextNotificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextNotificationTime

`func (o *Alert) SetNextNotificationTime(v time.Time)`

SetNextNotificationTime sets NextNotificationTime field to given value.

### HasNextNotificationTime

`func (o *Alert) HasNextNotificationTime() bool`

HasNextNotificationTime returns a boolean if a field has been set.

### GetNotificationAttemptTime

`func (o *Alert) GetNotificationAttemptTime() time.Time`

GetNotificationAttemptTime returns the NotificationAttemptTime field if non-nil, zero value otherwise.

### GetNotificationAttemptTimeOk

`func (o *Alert) GetNotificationAttemptTimeOk() (*time.Time, bool)`

GetNotificationAttemptTimeOk returns a tuple with the NotificationAttemptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationAttemptTime

`func (o *Alert) SetNotificationAttemptTime(v time.Time)`

SetNotificationAttemptTime sets NotificationAttemptTime field to given value.

### HasNotificationAttemptTime

`func (o *Alert) HasNotificationAttemptTime() bool`

HasNotificationAttemptTime returns a boolean if a field has been set.

### GetNotificationsFailed

`func (o *Alert) GetNotificationsFailed() int32`

GetNotificationsFailed returns the NotificationsFailed field if non-nil, zero value otherwise.

### GetNotificationsFailedOk

`func (o *Alert) GetNotificationsFailedOk() (*int32, bool)`

GetNotificationsFailedOk returns a tuple with the NotificationsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsFailed

`func (o *Alert) SetNotificationsFailed(v int32)`

SetNotificationsFailed sets NotificationsFailed field to given value.

### HasNotificationsFailed

`func (o *Alert) HasNotificationsFailed() bool`

HasNotificationsFailed returns a boolean if a field has been set.

### GetNotifiedState

`func (o *Alert) GetNotifiedState() string`

GetNotifiedState returns the NotifiedState field if non-nil, zero value otherwise.

### GetNotifiedStateOk

`func (o *Alert) GetNotifiedStateOk() (*string, bool)`

GetNotifiedStateOk returns a tuple with the NotifiedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifiedState

`func (o *Alert) SetNotifiedState(v string)`

SetNotifiedState sets NotifiedState field to given value.

### HasNotifiedState

`func (o *Alert) HasNotifiedState() bool`

HasNotifiedState returns a boolean if a field has been set.

### GetResolvedTime

`func (o *Alert) GetResolvedTime() time.Time`

GetResolvedTime returns the ResolvedTime field if non-nil, zero value otherwise.

### GetResolvedTimeOk

`func (o *Alert) GetResolvedTimeOk() (*time.Time, bool)`

GetResolvedTimeOk returns a tuple with the ResolvedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedTime

`func (o *Alert) SetResolvedTime(v time.Time)`

SetResolvedTime sets ResolvedTime field to given value.

### HasResolvedTime

`func (o *Alert) HasResolvedTime() bool`

HasResolvedTime returns a boolean if a field has been set.

### GetSeverity

`func (o *Alert) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Alert) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Alert) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetSeverityIndex

`func (o *Alert) GetSeverityIndex() int32`

GetSeverityIndex returns the SeverityIndex field if non-nil, zero value otherwise.

### GetSeverityIndexOk

`func (o *Alert) GetSeverityIndexOk() (*int32, bool)`

GetSeverityIndexOk returns a tuple with the SeverityIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityIndex

`func (o *Alert) SetSeverityIndex(v int32)`

SetSeverityIndex sets SeverityIndex field to given value.


### GetSourceName

`func (o *Alert) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *Alert) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *Alert) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.


### GetSourceUUID

`func (o *Alert) GetSourceUUID() string`

GetSourceUUID returns the SourceUUID field if non-nil, zero value otherwise.

### GetSourceUUIDOk

`func (o *Alert) GetSourceUUIDOk() (*string, bool)`

GetSourceUUIDOk returns a tuple with the SourceUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUUID

`func (o *Alert) SetSourceUUID(v string)`

SetSourceUUID sets SourceUUID field to given value.


### GetState

`func (o *Alert) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Alert) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Alert) SetState(v string)`

SetState sets State field to given value.


### GetStateIndex

`func (o *Alert) GetStateIndex() int32`

GetStateIndex returns the StateIndex field if non-nil, zero value otherwise.

### GetStateIndexOk

`func (o *Alert) GetStateIndexOk() (*int32, bool)`

GetStateIndexOk returns a tuple with the StateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateIndex

`func (o *Alert) SetStateIndex(v int32)`

SetStateIndex sets StateIndex field to given value.


### GetUuid

`func (o *Alert) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Alert) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Alert) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Alert) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


