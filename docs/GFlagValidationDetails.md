# GFlagValidationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | WARNING: This is a preview API that could change. Error message if gflag is invalid | [optional] 
**Exist** | Pointer to **bool** | WARNING: This is a preview API that could change. Flag to indicate if gflag exists | [optional] 

## Methods

### NewGFlagValidationDetails

`func NewGFlagValidationDetails() *GFlagValidationDetails`

NewGFlagValidationDetails instantiates a new GFlagValidationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGFlagValidationDetailsWithDefaults

`func NewGFlagValidationDetailsWithDefaults() *GFlagValidationDetails`

NewGFlagValidationDetailsWithDefaults instantiates a new GFlagValidationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GFlagValidationDetails) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GFlagValidationDetails) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GFlagValidationDetails) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GFlagValidationDetails) HasError() bool`

HasError returns a boolean if a field has been set.

### GetExist

`func (o *GFlagValidationDetails) GetExist() bool`

GetExist returns the Exist field if non-nil, zero value otherwise.

### GetExistOk

`func (o *GFlagValidationDetails) GetExistOk() (*bool, bool)`

GetExistOk returns a tuple with the Exist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExist

`func (o *GFlagValidationDetails) SetExist(v bool)`

SetExist sets Exist field to given value.

### HasExist

`func (o *GFlagValidationDetails) HasExist() bool`

HasExist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


