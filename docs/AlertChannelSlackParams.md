# AlertChannelSlackParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IconUrl** | Pointer to **string** | Slack icon URL | [optional] 
**Username** | **string** | Slack username | 
**WebhookUrl** | **string** | Slack webhook URL | 

## Methods

### NewAlertChannelSlackParams

`func NewAlertChannelSlackParams(username string, webhookUrl string, ) *AlertChannelSlackParams`

NewAlertChannelSlackParams instantiates a new AlertChannelSlackParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertChannelSlackParamsWithDefaults

`func NewAlertChannelSlackParamsWithDefaults() *AlertChannelSlackParams`

NewAlertChannelSlackParamsWithDefaults instantiates a new AlertChannelSlackParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIconUrl

`func (o *AlertChannelSlackParams) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *AlertChannelSlackParams) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *AlertChannelSlackParams) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *AlertChannelSlackParams) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetUsername

`func (o *AlertChannelSlackParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AlertChannelSlackParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AlertChannelSlackParams) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetWebhookUrl

`func (o *AlertChannelSlackParams) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *AlertChannelSlackParams) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *AlertChannelSlackParams) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


