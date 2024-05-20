# AlertLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**AlertLabelKey**](AlertLabelKey.md) |  | 
**Name** | **string** |  | 
**Value** | **string** |  | 

## Methods

### NewAlertLabel

`func NewAlertLabel(key AlertLabelKey, name string, value string, ) *AlertLabel`

NewAlertLabel instantiates a new AlertLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertLabelWithDefaults

`func NewAlertLabelWithDefaults() *AlertLabel`

NewAlertLabelWithDefaults instantiates a new AlertLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AlertLabel) GetKey() AlertLabelKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AlertLabel) GetKeyOk() (*AlertLabelKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AlertLabel) SetKey(v AlertLabelKey)`

SetKey sets Key field to given value.


### GetName

`func (o *AlertLabel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertLabel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertLabel) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *AlertLabel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AlertLabel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AlertLabel) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


