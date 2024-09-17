# AlertChannelParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelType** | Pointer to **string** | Channel type | [optional] 
**TextTemplate** | Pointer to **string** | Notification text template | [optional] 
**TitleTemplate** | Pointer to **string** | Notification title template | [optional] 

## Methods

### NewAlertChannelParams

`func NewAlertChannelParams() *AlertChannelParams`

NewAlertChannelParams instantiates a new AlertChannelParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertChannelParamsWithDefaults

`func NewAlertChannelParamsWithDefaults() *AlertChannelParams`

NewAlertChannelParamsWithDefaults instantiates a new AlertChannelParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelType

`func (o *AlertChannelParams) GetChannelType() string`

GetChannelType returns the ChannelType field if non-nil, zero value otherwise.

### GetChannelTypeOk

`func (o *AlertChannelParams) GetChannelTypeOk() (*string, bool)`

GetChannelTypeOk returns a tuple with the ChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelType

`func (o *AlertChannelParams) SetChannelType(v string)`

SetChannelType sets ChannelType field to given value.

### HasChannelType

`func (o *AlertChannelParams) HasChannelType() bool`

HasChannelType returns a boolean if a field has been set.

### GetTextTemplate

`func (o *AlertChannelParams) GetTextTemplate() string`

GetTextTemplate returns the TextTemplate field if non-nil, zero value otherwise.

### GetTextTemplateOk

`func (o *AlertChannelParams) GetTextTemplateOk() (*string, bool)`

GetTextTemplateOk returns a tuple with the TextTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextTemplate

`func (o *AlertChannelParams) SetTextTemplate(v string)`

SetTextTemplate sets TextTemplate field to given value.

### HasTextTemplate

`func (o *AlertChannelParams) HasTextTemplate() bool`

HasTextTemplate returns a boolean if a field has been set.

### GetTitleTemplate

`func (o *AlertChannelParams) GetTitleTemplate() string`

GetTitleTemplate returns the TitleTemplate field if non-nil, zero value otherwise.

### GetTitleTemplateOk

`func (o *AlertChannelParams) GetTitleTemplateOk() (*string, bool)`

GetTitleTemplateOk returns a tuple with the TitleTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleTemplate

`func (o *AlertChannelParams) SetTitleTemplate(v string)`

SetTitleTemplate sets TitleTemplate field to given value.

### HasTitleTemplate

`func (o *AlertChannelParams) HasTitleTemplate() bool`

HasTitleTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


