# UniverseCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**UniverseSpec**](UniverseSpec.md) |  | 
**Arch** | **string** | CPU Arch of DB nodes. | 

## Methods

### NewUniverseCreateSpec

`func NewUniverseCreateSpec(spec UniverseSpec, arch string, ) *UniverseCreateSpec`

NewUniverseCreateSpec instantiates a new UniverseCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseCreateSpecWithDefaults

`func NewUniverseCreateSpecWithDefaults() *UniverseCreateSpec`

NewUniverseCreateSpecWithDefaults instantiates a new UniverseCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *UniverseCreateSpec) GetSpec() UniverseSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *UniverseCreateSpec) GetSpecOk() (*UniverseSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *UniverseCreateSpec) SetSpec(v UniverseSpec)`

SetSpec sets Spec field to given value.


### GetArch

`func (o *UniverseCreateSpec) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *UniverseCreateSpec) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *UniverseCreateSpec) SetArch(v string)`

SetArch sets Arch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


