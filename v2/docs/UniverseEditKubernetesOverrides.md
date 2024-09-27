# UniverseEditKubernetesOverrides

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overrides** | Pointer to **string** | Global kubernetes overrides to apply across the entire universe. | [optional] 
**AzOverrides** | Pointer to **map[string]string** | Granular kubernetes overrides per Availability Zone identified by AZ uuid. | [optional] 

## Methods

### NewUniverseEditKubernetesOverrides

`func NewUniverseEditKubernetesOverrides() *UniverseEditKubernetesOverrides`

NewUniverseEditKubernetesOverrides instantiates a new UniverseEditKubernetesOverrides object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseEditKubernetesOverridesWithDefaults

`func NewUniverseEditKubernetesOverridesWithDefaults() *UniverseEditKubernetesOverrides`

NewUniverseEditKubernetesOverridesWithDefaults instantiates a new UniverseEditKubernetesOverrides object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverrides

`func (o *UniverseEditKubernetesOverrides) GetOverrides() string`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *UniverseEditKubernetesOverrides) GetOverridesOk() (*string, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *UniverseEditKubernetesOverrides) SetOverrides(v string)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *UniverseEditKubernetesOverrides) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### GetAzOverrides

`func (o *UniverseEditKubernetesOverrides) GetAzOverrides() map[string]string`

GetAzOverrides returns the AzOverrides field if non-nil, zero value otherwise.

### GetAzOverridesOk

`func (o *UniverseEditKubernetesOverrides) GetAzOverridesOk() (*map[string]string, bool)`

GetAzOverridesOk returns a tuple with the AzOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzOverrides

`func (o *UniverseEditKubernetesOverrides) SetAzOverrides(v map[string]string)`

SetAzOverrides sets AzOverrides field to given value.

### HasAzOverrides

`func (o *UniverseEditKubernetesOverrides) HasAzOverrides() bool`

HasAzOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


