# AlertChannelEmailParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultRecipients** | Pointer to **bool** | Use health check notification recipients | [optional] 
**DefaultSmtpSettings** | Pointer to **bool** | Use health check notification SMTP settings | [optional] 
**Recipients** | Pointer to **[]string** | List of recipients | [optional] 
**SmtpData** | Pointer to [**SmtpData**](SmtpData.md) |  | [optional] 

## Methods

### NewAlertChannelEmailParamsAllOf

`func NewAlertChannelEmailParamsAllOf() *AlertChannelEmailParamsAllOf`

NewAlertChannelEmailParamsAllOf instantiates a new AlertChannelEmailParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertChannelEmailParamsAllOfWithDefaults

`func NewAlertChannelEmailParamsAllOfWithDefaults() *AlertChannelEmailParamsAllOf`

NewAlertChannelEmailParamsAllOfWithDefaults instantiates a new AlertChannelEmailParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultRecipients

`func (o *AlertChannelEmailParamsAllOf) GetDefaultRecipients() bool`

GetDefaultRecipients returns the DefaultRecipients field if non-nil, zero value otherwise.

### GetDefaultRecipientsOk

`func (o *AlertChannelEmailParamsAllOf) GetDefaultRecipientsOk() (*bool, bool)`

GetDefaultRecipientsOk returns a tuple with the DefaultRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRecipients

`func (o *AlertChannelEmailParamsAllOf) SetDefaultRecipients(v bool)`

SetDefaultRecipients sets DefaultRecipients field to given value.

### HasDefaultRecipients

`func (o *AlertChannelEmailParamsAllOf) HasDefaultRecipients() bool`

HasDefaultRecipients returns a boolean if a field has been set.

### GetDefaultSmtpSettings

`func (o *AlertChannelEmailParamsAllOf) GetDefaultSmtpSettings() bool`

GetDefaultSmtpSettings returns the DefaultSmtpSettings field if non-nil, zero value otherwise.

### GetDefaultSmtpSettingsOk

`func (o *AlertChannelEmailParamsAllOf) GetDefaultSmtpSettingsOk() (*bool, bool)`

GetDefaultSmtpSettingsOk returns a tuple with the DefaultSmtpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSmtpSettings

`func (o *AlertChannelEmailParamsAllOf) SetDefaultSmtpSettings(v bool)`

SetDefaultSmtpSettings sets DefaultSmtpSettings field to given value.

### HasDefaultSmtpSettings

`func (o *AlertChannelEmailParamsAllOf) HasDefaultSmtpSettings() bool`

HasDefaultSmtpSettings returns a boolean if a field has been set.

### GetRecipients

`func (o *AlertChannelEmailParamsAllOf) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *AlertChannelEmailParamsAllOf) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *AlertChannelEmailParamsAllOf) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *AlertChannelEmailParamsAllOf) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetSmtpData

`func (o *AlertChannelEmailParamsAllOf) GetSmtpData() SmtpData`

GetSmtpData returns the SmtpData field if non-nil, zero value otherwise.

### GetSmtpDataOk

`func (o *AlertChannelEmailParamsAllOf) GetSmtpDataOk() (*SmtpData, bool)`

GetSmtpDataOk returns a tuple with the SmtpData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpData

`func (o *AlertChannelEmailParamsAllOf) SetSmtpData(v SmtpData)`

SetSmtpData sets SmtpData field to given value.

### HasSmtpData

`func (o *AlertChannelEmailParamsAllOf) HasSmtpData() bool`

HasSmtpData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


