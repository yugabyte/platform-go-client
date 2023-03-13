# AlertChannelTemplates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerUUID** | **string** | Customer UUID | [readonly] 
**TextTemplate** | **string** | Notification text template | 
**TitleTemplate** | Pointer to **string** | Notification title template | [optional] 
**Type** | **string** | Channel type | 

## Methods

### NewAlertChannelTemplates

`func NewAlertChannelTemplates(customerUUID string, textTemplate string, type_ string, ) *AlertChannelTemplates`

NewAlertChannelTemplates instantiates a new AlertChannelTemplates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertChannelTemplatesWithDefaults

`func NewAlertChannelTemplatesWithDefaults() *AlertChannelTemplates`

NewAlertChannelTemplatesWithDefaults instantiates a new AlertChannelTemplates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerUUID

`func (o *AlertChannelTemplates) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *AlertChannelTemplates) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *AlertChannelTemplates) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetTextTemplate

`func (o *AlertChannelTemplates) GetTextTemplate() string`

GetTextTemplate returns the TextTemplate field if non-nil, zero value otherwise.

### GetTextTemplateOk

`func (o *AlertChannelTemplates) GetTextTemplateOk() (*string, bool)`

GetTextTemplateOk returns a tuple with the TextTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextTemplate

`func (o *AlertChannelTemplates) SetTextTemplate(v string)`

SetTextTemplate sets TextTemplate field to given value.


### GetTitleTemplate

`func (o *AlertChannelTemplates) GetTitleTemplate() string`

GetTitleTemplate returns the TitleTemplate field if non-nil, zero value otherwise.

### GetTitleTemplateOk

`func (o *AlertChannelTemplates) GetTitleTemplateOk() (*string, bool)`

GetTitleTemplateOk returns a tuple with the TitleTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleTemplate

`func (o *AlertChannelTemplates) SetTitleTemplate(v string)`

SetTitleTemplate sets TitleTemplate field to given value.

### HasTitleTemplate

`func (o *AlertChannelTemplates) HasTitleTemplate() bool`

HasTitleTemplate returns a boolean if a field has been set.

### GetType

`func (o *AlertChannelTemplates) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertChannelTemplates) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertChannelTemplates) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


