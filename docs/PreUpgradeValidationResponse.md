# PreUpgradeValidationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **[]string** | WARNING: This is a preview API that could change. List of errors that occurred during validation | [optional] [readonly] 
**Success** | Pointer to **bool** | WARNING: This is a preview API that could change. Indicates whether all the checks passed | [optional] [readonly] 

## Methods

### NewPreUpgradeValidationResponse

`func NewPreUpgradeValidationResponse() *PreUpgradeValidationResponse`

NewPreUpgradeValidationResponse instantiates a new PreUpgradeValidationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreUpgradeValidationResponseWithDefaults

`func NewPreUpgradeValidationResponseWithDefaults() *PreUpgradeValidationResponse`

NewPreUpgradeValidationResponseWithDefaults instantiates a new PreUpgradeValidationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *PreUpgradeValidationResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *PreUpgradeValidationResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *PreUpgradeValidationResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *PreUpgradeValidationResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetSuccess

`func (o *PreUpgradeValidationResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PreUpgradeValidationResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PreUpgradeValidationResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *PreUpgradeValidationResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


