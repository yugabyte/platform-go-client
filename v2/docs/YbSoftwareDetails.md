# YbSoftwareDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**YbSoftwareVersion** | Pointer to **string** | YugabyteDB Software version installed in DB nodes of this Universe | [optional] 
**AutoFlagConfigVersion** | Pointer to **int32** | Auto flag version installed in DB nodes of this Universe | [optional] 

## Methods

### NewYbSoftwareDetails

`func NewYbSoftwareDetails() *YbSoftwareDetails`

NewYbSoftwareDetails instantiates a new YbSoftwareDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYbSoftwareDetailsWithDefaults

`func NewYbSoftwareDetailsWithDefaults() *YbSoftwareDetails`

NewYbSoftwareDetailsWithDefaults instantiates a new YbSoftwareDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetYbSoftwareVersion

`func (o *YbSoftwareDetails) GetYbSoftwareVersion() string`

GetYbSoftwareVersion returns the YbSoftwareVersion field if non-nil, zero value otherwise.

### GetYbSoftwareVersionOk

`func (o *YbSoftwareDetails) GetYbSoftwareVersionOk() (*string, bool)`

GetYbSoftwareVersionOk returns a tuple with the YbSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbSoftwareVersion

`func (o *YbSoftwareDetails) SetYbSoftwareVersion(v string)`

SetYbSoftwareVersion sets YbSoftwareVersion field to given value.

### HasYbSoftwareVersion

`func (o *YbSoftwareDetails) HasYbSoftwareVersion() bool`

HasYbSoftwareVersion returns a boolean if a field has been set.

### GetAutoFlagConfigVersion

`func (o *YbSoftwareDetails) GetAutoFlagConfigVersion() int32`

GetAutoFlagConfigVersion returns the AutoFlagConfigVersion field if non-nil, zero value otherwise.

### GetAutoFlagConfigVersionOk

`func (o *YbSoftwareDetails) GetAutoFlagConfigVersionOk() (*int32, bool)`

GetAutoFlagConfigVersionOk returns a tuple with the AutoFlagConfigVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoFlagConfigVersion

`func (o *YbSoftwareDetails) SetAutoFlagConfigVersion(v int32)`

SetAutoFlagConfigVersion sets AutoFlagConfigVersion field to given value.

### HasAutoFlagConfigVersion

`func (o *YbSoftwareDetails) HasAutoFlagConfigVersion() bool`

HasAutoFlagConfigVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


