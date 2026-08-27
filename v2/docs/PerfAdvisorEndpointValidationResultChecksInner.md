# PerfAdvisorEndpointValidationResultChecksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | The property the operator has to fix, so a caller can attach the message to the right input.  | [optional] 
**Ok** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** | What went wrong, written for the operator. Unset when the check passed. | [optional] 

## Methods

### NewPerfAdvisorEndpointValidationResultChecksInner

`func NewPerfAdvisorEndpointValidationResultChecksInner() *PerfAdvisorEndpointValidationResultChecksInner`

NewPerfAdvisorEndpointValidationResultChecksInner instantiates a new PerfAdvisorEndpointValidationResultChecksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerfAdvisorEndpointValidationResultChecksInnerWithDefaults

`func NewPerfAdvisorEndpointValidationResultChecksInnerWithDefaults() *PerfAdvisorEndpointValidationResultChecksInner`

NewPerfAdvisorEndpointValidationResultChecksInnerWithDefaults instantiates a new PerfAdvisorEndpointValidationResultChecksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *PerfAdvisorEndpointValidationResultChecksInner) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *PerfAdvisorEndpointValidationResultChecksInner) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *PerfAdvisorEndpointValidationResultChecksInner) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *PerfAdvisorEndpointValidationResultChecksInner) HasField() bool`

HasField returns a boolean if a field has been set.

### GetOk

`func (o *PerfAdvisorEndpointValidationResultChecksInner) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *PerfAdvisorEndpointValidationResultChecksInner) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *PerfAdvisorEndpointValidationResultChecksInner) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *PerfAdvisorEndpointValidationResultChecksInner) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetMessage

`func (o *PerfAdvisorEndpointValidationResultChecksInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PerfAdvisorEndpointValidationResultChecksInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PerfAdvisorEndpointValidationResultChecksInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PerfAdvisorEndpointValidationResultChecksInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


