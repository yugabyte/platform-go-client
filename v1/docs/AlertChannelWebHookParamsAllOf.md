# AlertChannelWebHookParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpAuth** | Pointer to [**HTTPAuthInformation**](HTTPAuthInformation.md) |  | [optional] 
**SendResolved** | Pointer to **bool** | WARNING: This is a preview API that could change. Send resolved alert notification | [optional] 
**WebhookUrl** | **string** | Webhook URL | 

## Methods

### NewAlertChannelWebHookParamsAllOf

`func NewAlertChannelWebHookParamsAllOf(webhookUrl string, ) *AlertChannelWebHookParamsAllOf`

NewAlertChannelWebHookParamsAllOf instantiates a new AlertChannelWebHookParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertChannelWebHookParamsAllOfWithDefaults

`func NewAlertChannelWebHookParamsAllOfWithDefaults() *AlertChannelWebHookParamsAllOf`

NewAlertChannelWebHookParamsAllOfWithDefaults instantiates a new AlertChannelWebHookParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpAuth

`func (o *AlertChannelWebHookParamsAllOf) GetHttpAuth() HTTPAuthInformation`

GetHttpAuth returns the HttpAuth field if non-nil, zero value otherwise.

### GetHttpAuthOk

`func (o *AlertChannelWebHookParamsAllOf) GetHttpAuthOk() (*HTTPAuthInformation, bool)`

GetHttpAuthOk returns a tuple with the HttpAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpAuth

`func (o *AlertChannelWebHookParamsAllOf) SetHttpAuth(v HTTPAuthInformation)`

SetHttpAuth sets HttpAuth field to given value.

### HasHttpAuth

`func (o *AlertChannelWebHookParamsAllOf) HasHttpAuth() bool`

HasHttpAuth returns a boolean if a field has been set.

### GetSendResolved

`func (o *AlertChannelWebHookParamsAllOf) GetSendResolved() bool`

GetSendResolved returns the SendResolved field if non-nil, zero value otherwise.

### GetSendResolvedOk

`func (o *AlertChannelWebHookParamsAllOf) GetSendResolvedOk() (*bool, bool)`

GetSendResolvedOk returns a tuple with the SendResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendResolved

`func (o *AlertChannelWebHookParamsAllOf) SetSendResolved(v bool)`

SetSendResolved sets SendResolved field to given value.

### HasSendResolved

`func (o *AlertChannelWebHookParamsAllOf) HasSendResolved() bool`

HasSendResolved returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *AlertChannelWebHookParamsAllOf) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *AlertChannelWebHookParamsAllOf) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *AlertChannelWebHookParamsAllOf) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


