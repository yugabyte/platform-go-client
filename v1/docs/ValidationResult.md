# ValidationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**IsRequired** | **bool** |  | 
**IsValid** | **bool** |  | 
**Type** | **string** |  | 
**Value** | **string** |  | 

## Methods

### NewValidationResult

`func NewValidationResult(description string, isRequired bool, isValid bool, type_ string, value string, ) *ValidationResult`

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


### GetIsRequired

`func (o *ValidationResult) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *ValidationResult) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *ValidationResult) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.


### GetIsValid

`func (o *ValidationResult) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *ValidationResult) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *ValidationResult) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.


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


