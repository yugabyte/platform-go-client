# AdminNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Notification code | [optional] [readonly] 
**HtmlMessage** | Pointer to **string** | Notification message with HTML markup | [optional] [readonly] 

## Methods

### NewAdminNotification

`func NewAdminNotification() *AdminNotification`

NewAdminNotification instantiates a new AdminNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminNotificationWithDefaults

`func NewAdminNotificationWithDefaults() *AdminNotification`

NewAdminNotificationWithDefaults instantiates a new AdminNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AdminNotification) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AdminNotification) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AdminNotification) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AdminNotification) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetHtmlMessage

`func (o *AdminNotification) GetHtmlMessage() string`

GetHtmlMessage returns the HtmlMessage field if non-nil, zero value otherwise.

### GetHtmlMessageOk

`func (o *AdminNotification) GetHtmlMessageOk() (*string, bool)`

GetHtmlMessageOk returns a tuple with the HtmlMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlMessage

`func (o *AdminNotification) SetHtmlMessage(v string)`

SetHtmlMessage sets HtmlMessage field to given value.

### HasHtmlMessage

`func (o *AdminNotification) HasHtmlMessage() bool`

HasHtmlMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


