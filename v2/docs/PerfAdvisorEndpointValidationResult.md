# PerfAdvisorEndpointValidationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | Pointer to **bool** | True when every check passed. | [optional] [readonly] 
**Checks** | Pointer to [**[]PerfAdvisorEndpointValidationResultChecksInner**](PerfAdvisorEndpointValidationResultChecksInner.md) | One entry per probed endpoint. | [optional] [readonly] 

## Methods

### NewPerfAdvisorEndpointValidationResult

`func NewPerfAdvisorEndpointValidationResult() *PerfAdvisorEndpointValidationResult`

NewPerfAdvisorEndpointValidationResult instantiates a new PerfAdvisorEndpointValidationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerfAdvisorEndpointValidationResultWithDefaults

`func NewPerfAdvisorEndpointValidationResultWithDefaults() *PerfAdvisorEndpointValidationResult`

NewPerfAdvisorEndpointValidationResultWithDefaults instantiates a new PerfAdvisorEndpointValidationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *PerfAdvisorEndpointValidationResult) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *PerfAdvisorEndpointValidationResult) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *PerfAdvisorEndpointValidationResult) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *PerfAdvisorEndpointValidationResult) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetChecks

`func (o *PerfAdvisorEndpointValidationResult) GetChecks() []PerfAdvisorEndpointValidationResultChecksInner`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *PerfAdvisorEndpointValidationResult) GetChecksOk() (*[]PerfAdvisorEndpointValidationResultChecksInner, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *PerfAdvisorEndpointValidationResult) SetChecks(v []PerfAdvisorEndpointValidationResultChecksInner)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *PerfAdvisorEndpointValidationResult) HasChecks() bool`

HasChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


