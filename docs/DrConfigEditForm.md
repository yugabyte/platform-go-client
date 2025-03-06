# DrConfigEditForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootstrapParams** | Pointer to [**RestartBootstrapParams**](RestartBootstrapParams.md) |  | [optional] 
**PitrParams** | Pointer to [**PitrParams**](PitrParams.md) |  | [optional] 
**WebhookUrls** | Pointer to **[]string** | List of urls for webhook | [optional] 

## Methods

### NewDrConfigEditForm

`func NewDrConfigEditForm() *DrConfigEditForm`

NewDrConfigEditForm instantiates a new DrConfigEditForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrConfigEditFormWithDefaults

`func NewDrConfigEditFormWithDefaults() *DrConfigEditForm`

NewDrConfigEditFormWithDefaults instantiates a new DrConfigEditForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootstrapParams

`func (o *DrConfigEditForm) GetBootstrapParams() RestartBootstrapParams`

GetBootstrapParams returns the BootstrapParams field if non-nil, zero value otherwise.

### GetBootstrapParamsOk

`func (o *DrConfigEditForm) GetBootstrapParamsOk() (*RestartBootstrapParams, bool)`

GetBootstrapParamsOk returns a tuple with the BootstrapParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapParams

`func (o *DrConfigEditForm) SetBootstrapParams(v RestartBootstrapParams)`

SetBootstrapParams sets BootstrapParams field to given value.

### HasBootstrapParams

`func (o *DrConfigEditForm) HasBootstrapParams() bool`

HasBootstrapParams returns a boolean if a field has been set.

### GetPitrParams

`func (o *DrConfigEditForm) GetPitrParams() PitrParams`

GetPitrParams returns the PitrParams field if non-nil, zero value otherwise.

### GetPitrParamsOk

`func (o *DrConfigEditForm) GetPitrParamsOk() (*PitrParams, bool)`

GetPitrParamsOk returns a tuple with the PitrParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitrParams

`func (o *DrConfigEditForm) SetPitrParams(v PitrParams)`

SetPitrParams sets PitrParams field to given value.

### HasPitrParams

`func (o *DrConfigEditForm) HasPitrParams() bool`

HasPitrParams returns a boolean if a field has been set.

### GetWebhookUrls

`func (o *DrConfigEditForm) GetWebhookUrls() []string`

GetWebhookUrls returns the WebhookUrls field if non-nil, zero value otherwise.

### GetWebhookUrlsOk

`func (o *DrConfigEditForm) GetWebhookUrlsOk() (*[]string, bool)`

GetWebhookUrlsOk returns a tuple with the WebhookUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrls

`func (o *DrConfigEditForm) SetWebhookUrls(v []string)`

SetWebhookUrls sets WebhookUrls field to given value.

### HasWebhookUrls

`func (o *DrConfigEditForm) HasWebhookUrls() bool`

HasWebhookUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


