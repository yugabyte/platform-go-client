# AlertingData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveAlertNotificationIntervalMs** | Pointer to **int64** | Period, which is used to send active alert notifications | [optional] 
**AlertingEmail** | Pointer to **string** | Alert email address | [optional] 
**CheckIntervalMs** | Pointer to **int64** | Alert interval, in milliseconds | [optional] 
**ReportOnlyErrors** | Pointer to **bool** | Trigger an alert only for errors | [optional] 
**SendAlertsToYb** | Pointer to **bool** | Send alerts to YB as well as to customer | [optional] 
**StatusUpdateIntervalMs** | Pointer to **int64** | Status update of alert interval, in milliseconds | [optional] 

## Methods

### NewAlertingData

`func NewAlertingData() *AlertingData`

NewAlertingData instantiates a new AlertingData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertingDataWithDefaults

`func NewAlertingDataWithDefaults() *AlertingData`

NewAlertingDataWithDefaults instantiates a new AlertingData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveAlertNotificationIntervalMs

`func (o *AlertingData) GetActiveAlertNotificationIntervalMs() int64`

GetActiveAlertNotificationIntervalMs returns the ActiveAlertNotificationIntervalMs field if non-nil, zero value otherwise.

### GetActiveAlertNotificationIntervalMsOk

`func (o *AlertingData) GetActiveAlertNotificationIntervalMsOk() (*int64, bool)`

GetActiveAlertNotificationIntervalMsOk returns a tuple with the ActiveAlertNotificationIntervalMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAlertNotificationIntervalMs

`func (o *AlertingData) SetActiveAlertNotificationIntervalMs(v int64)`

SetActiveAlertNotificationIntervalMs sets ActiveAlertNotificationIntervalMs field to given value.

### HasActiveAlertNotificationIntervalMs

`func (o *AlertingData) HasActiveAlertNotificationIntervalMs() bool`

HasActiveAlertNotificationIntervalMs returns a boolean if a field has been set.

### GetAlertingEmail

`func (o *AlertingData) GetAlertingEmail() string`

GetAlertingEmail returns the AlertingEmail field if non-nil, zero value otherwise.

### GetAlertingEmailOk

`func (o *AlertingData) GetAlertingEmailOk() (*string, bool)`

GetAlertingEmailOk returns a tuple with the AlertingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingEmail

`func (o *AlertingData) SetAlertingEmail(v string)`

SetAlertingEmail sets AlertingEmail field to given value.

### HasAlertingEmail

`func (o *AlertingData) HasAlertingEmail() bool`

HasAlertingEmail returns a boolean if a field has been set.

### GetCheckIntervalMs

`func (o *AlertingData) GetCheckIntervalMs() int64`

GetCheckIntervalMs returns the CheckIntervalMs field if non-nil, zero value otherwise.

### GetCheckIntervalMsOk

`func (o *AlertingData) GetCheckIntervalMsOk() (*int64, bool)`

GetCheckIntervalMsOk returns a tuple with the CheckIntervalMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckIntervalMs

`func (o *AlertingData) SetCheckIntervalMs(v int64)`

SetCheckIntervalMs sets CheckIntervalMs field to given value.

### HasCheckIntervalMs

`func (o *AlertingData) HasCheckIntervalMs() bool`

HasCheckIntervalMs returns a boolean if a field has been set.

### GetReportOnlyErrors

`func (o *AlertingData) GetReportOnlyErrors() bool`

GetReportOnlyErrors returns the ReportOnlyErrors field if non-nil, zero value otherwise.

### GetReportOnlyErrorsOk

`func (o *AlertingData) GetReportOnlyErrorsOk() (*bool, bool)`

GetReportOnlyErrorsOk returns a tuple with the ReportOnlyErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportOnlyErrors

`func (o *AlertingData) SetReportOnlyErrors(v bool)`

SetReportOnlyErrors sets ReportOnlyErrors field to given value.

### HasReportOnlyErrors

`func (o *AlertingData) HasReportOnlyErrors() bool`

HasReportOnlyErrors returns a boolean if a field has been set.

### GetSendAlertsToYb

`func (o *AlertingData) GetSendAlertsToYb() bool`

GetSendAlertsToYb returns the SendAlertsToYb field if non-nil, zero value otherwise.

### GetSendAlertsToYbOk

`func (o *AlertingData) GetSendAlertsToYbOk() (*bool, bool)`

GetSendAlertsToYbOk returns a tuple with the SendAlertsToYb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAlertsToYb

`func (o *AlertingData) SetSendAlertsToYb(v bool)`

SetSendAlertsToYb sets SendAlertsToYb field to given value.

### HasSendAlertsToYb

`func (o *AlertingData) HasSendAlertsToYb() bool`

HasSendAlertsToYb returns a boolean if a field has been set.

### GetStatusUpdateIntervalMs

`func (o *AlertingData) GetStatusUpdateIntervalMs() int64`

GetStatusUpdateIntervalMs returns the StatusUpdateIntervalMs field if non-nil, zero value otherwise.

### GetStatusUpdateIntervalMsOk

`func (o *AlertingData) GetStatusUpdateIntervalMsOk() (*int64, bool)`

GetStatusUpdateIntervalMsOk returns a tuple with the StatusUpdateIntervalMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUpdateIntervalMs

`func (o *AlertingData) SetStatusUpdateIntervalMs(v int64)`

SetStatusUpdateIntervalMs sets StatusUpdateIntervalMs field to given value.

### HasStatusUpdateIntervalMs

`func (o *AlertingData) HasStatusUpdateIntervalMs() bool`

HasStatusUpdateIntervalMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


