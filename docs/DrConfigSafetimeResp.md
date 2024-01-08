# DrConfigSafetimeResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Safetimes** | Pointer to [**[]NamespaceSafetime**](NamespaceSafetime.md) | The list of current safetime for each database | [optional] 

## Methods

### NewDrConfigSafetimeResp

`func NewDrConfigSafetimeResp() *DrConfigSafetimeResp`

NewDrConfigSafetimeResp instantiates a new DrConfigSafetimeResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrConfigSafetimeRespWithDefaults

`func NewDrConfigSafetimeRespWithDefaults() *DrConfigSafetimeResp`

NewDrConfigSafetimeRespWithDefaults instantiates a new DrConfigSafetimeResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSafetimes

`func (o *DrConfigSafetimeResp) GetSafetimes() []NamespaceSafetime`

GetSafetimes returns the Safetimes field if non-nil, zero value otherwise.

### GetSafetimesOk

`func (o *DrConfigSafetimeResp) GetSafetimesOk() (*[]NamespaceSafetime, bool)`

GetSafetimesOk returns a tuple with the Safetimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetimes

`func (o *DrConfigSafetimeResp) SetSafetimes(v []NamespaceSafetime)`

SetSafetimes sets Safetimes field to given value.

### HasSafetimes

`func (o *DrConfigSafetimeResp) HasSafetimes() bool`

HasSafetimes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


