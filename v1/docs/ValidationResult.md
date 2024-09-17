# ValidationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Required** | **bool** |  | 
**Type** | **string** |  | 
**Valid** | **bool** |  | 
**Value** | **string** |  | 

## Methods

### NewValidationResult

`func NewValidationResult(description string, required bool, type_ string, valid bool, value string, ) *ValidationResult`

NewValidationResult instantiates a new ValidationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationResultWithDefaults

`func NewValidationResultWithDefaults() *ValidationResult`

NewValidationResultWithDefaults instantiates a new ValidationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ValidationResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ValidationResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ValidationResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRequired

`func (o *ValidationResult) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ValidationResult) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ValidationResult) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetType

`func (o *ValidationResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ValidationResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ValidationResult) SetType(v string)`

SetType sets Type field to given value.


### GetValid

`func (o *ValidationResult) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ValidationResult) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ValidationResult) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetValue

`func (o *ValidationResult) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ValidationResult) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ValidationResult) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


