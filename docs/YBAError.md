# YBAError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Error code | [optional] [readonly] 
**Message** | Pointer to **string** | Error message | [optional] [readonly] 

## Methods

### NewYBAError

`func NewYBAError() *YBAError`

NewYBAError instantiates a new YBAError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYBAErrorWithDefaults

`func NewYBAErrorWithDefaults() *YBAError`

NewYBAErrorWithDefaults instantiates a new YBAError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *YBAError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *YBAError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *YBAError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *YBAError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *YBAError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *YBAError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *YBAError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *YBAError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


