# ThrottleParamValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentValue** | **int64** |  | 
**PresetValues** | [**PresetThrottleValues**](PresetThrottleValues.md) |  | 

## Methods

### NewThrottleParamValue

`func NewThrottleParamValue(currentValue int64, presetValues PresetThrottleValues, ) *ThrottleParamValue`

NewThrottleParamValue instantiates a new ThrottleParamValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThrottleParamValueWithDefaults

`func NewThrottleParamValueWithDefaults() *ThrottleParamValue`

NewThrottleParamValueWithDefaults instantiates a new ThrottleParamValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentValue

`func (o *ThrottleParamValue) GetCurrentValue() int64`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *ThrottleParamValue) GetCurrentValueOk() (*int64, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *ThrottleParamValue) SetCurrentValue(v int64)`

SetCurrentValue sets CurrentValue field to given value.


### GetPresetValues

`func (o *ThrottleParamValue) GetPresetValues() PresetThrottleValues`

GetPresetValues returns the PresetValues field if non-nil, zero value otherwise.

### GetPresetValuesOk

`func (o *ThrottleParamValue) GetPresetValuesOk() (*PresetThrottleValues, bool)`

GetPresetValuesOk returns a tuple with the PresetValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetValues

`func (o *ThrottleParamValue) SetPresetValues(v PresetThrottleValues)`

SetPresetValues sets PresetValues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


