# AlertChannelWebHookParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpAuth** | Pointer to [**HTTPAuthInformation**](HTTP Auth information.md) |  | [optional] 
**WebhookUrl** | **string** | Webhook URL | 

## Methods

### NewAlertChannelWebHookParams

`func NewAlertChannelWebHookParams(webhookUrl string, ) *AlertChannelWebHookParams`

NewAlertChannelWebHookParams instantiates a new AlertChannelWebHookParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertChannelWebHookParamsWithDefaults

`func NewAlertChannelWebHookParamsWithDefaults() *AlertChannelWebHookParams`

NewAlertChannelWebHookParamsWithDefaults instantiates a new AlertChannelWebHookParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpAuth

`func (o *AlertChannelWebHookParams) GetHttpAuth() HTTPAuthInformation`

GetHttpAuth returns the HttpAuth field if non-nil, zero value otherwise.

### GetHttpAuthOk

`func (o *AlertChannelWebHookParams) GetHttpAuthOk() (*HTTPAuthInformation, bool)`

GetHttpAuthOk returns a tuple with the HttpAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpAuth

`func (o *AlertChannelWebHookParams) SetHttpAuth(v HTTPAuthInformation)`

SetHttpAuth sets HttpAuth field to given value.

### HasHttpAuth

`func (o *AlertChannelWebHookParams) HasHttpAuth() bool`

HasHttpAuth returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *AlertChannelWebHookParams) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *AlertChannelWebHookParams) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *AlertChannelWebHookParams) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


