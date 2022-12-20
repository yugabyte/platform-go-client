# AlertChannelEmailParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultRecipients** | Pointer to **bool** | Use health check notification recipients | [optional] 
**DefaultSmtpSettings** | Pointer to **bool** | Use health check notification SMTP settings | [optional] 
**Recipients** | Pointer to **[]string** | List of recipients | [optional] 
**SmtpData** | Pointer to [**SmtpData**](SmtpData.md) |  | [optional] 

## Methods

### NewAlertChannelEmailParams

`func NewAlertChannelEmailParams() *AlertChannelEmailParams`

NewAlertChannelEmailParams instantiates a new AlertChannelEmailParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertChannelEmailParamsWithDefaults

`func NewAlertChannelEmailParamsWithDefaults() *AlertChannelEmailParams`

NewAlertChannelEmailParamsWithDefaults instantiates a new AlertChannelEmailParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultRecipients

`func (o *AlertChannelEmailParams) GetDefaultRecipients() bool`

GetDefaultRecipients returns the DefaultRecipients field if non-nil, zero value otherwise.

### GetDefaultRecipientsOk

`func (o *AlertChannelEmailParams) GetDefaultRecipientsOk() (*bool, bool)`

GetDefaultRecipientsOk returns a tuple with the DefaultRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRecipients

`func (o *AlertChannelEmailParams) SetDefaultRecipients(v bool)`

SetDefaultRecipients sets DefaultRecipients field to given value.

### HasDefaultRecipients

`func (o *AlertChannelEmailParams) HasDefaultRecipients() bool`

HasDefaultRecipients returns a boolean if a field has been set.

### GetDefaultSmtpSettings

`func (o *AlertChannelEmailParams) GetDefaultSmtpSettings() bool`

GetDefaultSmtpSettings returns the DefaultSmtpSettings field if non-nil, zero value otherwise.

### GetDefaultSmtpSettingsOk

`func (o *AlertChannelEmailParams) GetDefaultSmtpSettingsOk() (*bool, bool)`

GetDefaultSmtpSettingsOk returns a tuple with the DefaultSmtpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSmtpSettings

`func (o *AlertChannelEmailParams) SetDefaultSmtpSettings(v bool)`

SetDefaultSmtpSettings sets DefaultSmtpSettings field to given value.

### HasDefaultSmtpSettings

`func (o *AlertChannelEmailParams) HasDefaultSmtpSettings() bool`

HasDefaultSmtpSettings returns a boolean if a field has been set.

### GetRecipients

`func (o *AlertChannelEmailParams) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *AlertChannelEmailParams) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *AlertChannelEmailParams) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *AlertChannelEmailParams) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetSmtpData

`func (o *AlertChannelEmailParams) GetSmtpData() SmtpData`

GetSmtpData returns the SmtpData field if non-nil, zero value otherwise.

### GetSmtpDataOk

`func (o *AlertChannelEmailParams) GetSmtpDataOk() (*SmtpData, bool)`

GetSmtpDataOk returns a tuple with the SmtpData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpData

`func (o *AlertChannelEmailParams) SetSmtpData(v SmtpData)`

SetSmtpData sets SmtpData field to given value.

### HasSmtpData

`func (o *AlertChannelEmailParams) HasSmtpData() bool`

HasSmtpData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


