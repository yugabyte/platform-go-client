# SpecificGFlags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InheritFromPrimary** | Pointer to **bool** |  | [optional] 
**PerAZ** | Pointer to [**map[string]PerProcessFlags**](PerProcessFlags.md) | Overrides for gflags per availability zone | [optional] 
**PerProcessFlags** | Pointer to [**PerProcessFlags**](PerProcessFlags.md) |  | [optional] 

## Methods

### NewSpecificGFlags

`func NewSpecificGFlags() *SpecificGFlags`

NewSpecificGFlags instantiates a new SpecificGFlags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecificGFlagsWithDefaults

`func NewSpecificGFlagsWithDefaults() *SpecificGFlags`

NewSpecificGFlagsWithDefaults instantiates a new SpecificGFlags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInheritFromPrimary

`func (o *SpecificGFlags) GetInheritFromPrimary() bool`

GetInheritFromPrimary returns the InheritFromPrimary field if non-nil, zero value otherwise.

### GetInheritFromPrimaryOk

`func (o *SpecificGFlags) GetInheritFromPrimaryOk() (*bool, bool)`

GetInheritFromPrimaryOk returns a tuple with the InheritFromPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritFromPrimary

`func (o *SpecificGFlags) SetInheritFromPrimary(v bool)`

SetInheritFromPrimary sets InheritFromPrimary field to given value.

### HasInheritFromPrimary

`func (o *SpecificGFlags) HasInheritFromPrimary() bool`

HasInheritFromPrimary returns a boolean if a field has been set.

### GetPerAZ

`func (o *SpecificGFlags) GetPerAZ() map[string]PerProcessFlags`

GetPerAZ returns the PerAZ field if non-nil, zero value otherwise.

### GetPerAZOk

`func (o *SpecificGFlags) GetPerAZOk() (*map[string]PerProcessFlags, bool)`

GetPerAZOk returns a tuple with the PerAZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerAZ

`func (o *SpecificGFlags) SetPerAZ(v map[string]PerProcessFlags)`

SetPerAZ sets PerAZ field to given value.

### HasPerAZ

`func (o *SpecificGFlags) HasPerAZ() bool`

HasPerAZ returns a boolean if a field has been set.

### GetPerProcessFlags

`func (o *SpecificGFlags) GetPerProcessFlags() PerProcessFlags`

GetPerProcessFlags returns the PerProcessFlags field if non-nil, zero value otherwise.

### GetPerProcessFlagsOk

`func (o *SpecificGFlags) GetPerProcessFlagsOk() (*PerProcessFlags, bool)`

GetPerProcessFlagsOk returns a tuple with the PerProcessFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerProcessFlags

`func (o *SpecificGFlags) SetPerProcessFlags(v PerProcessFlags)`

SetPerProcessFlags sets PerProcessFlags field to given value.

### HasPerProcessFlags

`func (o *SpecificGFlags) HasPerProcessFlags() bool`

HasPerProcessFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


