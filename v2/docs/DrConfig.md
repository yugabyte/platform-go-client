# DrConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | Pointer to [**DrConfigSpec**](DrConfigSpec.md) |  | [optional] 
**Info** | Pointer to [**DrConfigInfo**](DrConfigInfo.md) |  | [optional] 

## Methods

### NewDrConfig

`func NewDrConfig() *DrConfig`

NewDrConfig instantiates a new DrConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrConfigWithDefaults

`func NewDrConfigWithDefaults() *DrConfig`

NewDrConfigWithDefaults instantiates a new DrConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *DrConfig) GetSpec() DrConfigSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *DrConfig) GetSpecOk() (*DrConfigSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *DrConfig) SetSpec(v DrConfigSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *DrConfig) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetInfo

`func (o *DrConfig) GetInfo() DrConfigInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *DrConfig) GetInfoOk() (*DrConfigInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *DrConfig) SetInfo(v DrConfigInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *DrConfig) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


