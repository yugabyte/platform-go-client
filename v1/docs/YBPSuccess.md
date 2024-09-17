# YBPSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | API response message. | [optional] [readonly] 
**Success** | Pointer to **bool** | API operation status. A value of true indicates the operation was successful. | [optional] [readonly] 

## Methods

### NewYBPSuccess

`func NewYBPSuccess() *YBPSuccess`

NewYBPSuccess instantiates a new YBPSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYBPSuccessWithDefaults

`func NewYBPSuccessWithDefaults() *YBPSuccess`

NewYBPSuccessWithDefaults instantiates a new YBPSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *YBPSuccess) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *YBPSuccess) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *YBPSuccess) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *YBPSuccess) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSuccess

`func (o *YBPSuccess) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *YBPSuccess) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *YBPSuccess) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *YBPSuccess) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


