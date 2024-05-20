# Universe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | Pointer to [**UniverseSpec**](UniverseSpec.md) |  | [optional] 
**Info** | Pointer to [**UniverseInfo**](UniverseInfo.md) |  | [optional] 

## Methods

### NewUniverse

`func NewUniverse() *Universe`

NewUniverse instantiates a new Universe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseWithDefaults

`func NewUniverseWithDefaults() *Universe`

NewUniverseWithDefaults instantiates a new Universe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *Universe) GetSpec() UniverseSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *Universe) GetSpecOk() (*UniverseSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *Universe) SetSpec(v UniverseSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *Universe) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetInfo

`func (o *Universe) GetInfo() UniverseInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Universe) GetInfoOk() (*UniverseInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Universe) SetInfo(v UniverseInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Universe) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


