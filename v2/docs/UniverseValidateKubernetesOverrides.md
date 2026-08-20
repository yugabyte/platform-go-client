# UniverseValidateKubernetesOverrides

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**YbSoftwareVersion** | Pointer to **string** | YugabyteDB Software version installed in DB nodes of this Universe | [optional] 
**NodePrefix** | Pointer to **string** | A globally unique name generated as a combination of the customer id and the universe name. This is used as the prefix of node names in the universe. | [optional] 
**IsReadonlyCluster** | Pointer to **bool** | Whether it is readonly cluster or not. | [optional] 
**PlacementSpec** | Pointer to [**ClusterPlacementSpec**](ClusterPlacementSpec.md) |  | [optional] 
**Overrides** | Pointer to **string** | Global kubernetes overrides to apply across the entire universe. | [optional] 
**AzOverrides** | Pointer to **map[string]string** | Granular kubernetes overrides per Availability Zone identified by AZ uuid. | [optional] 

## Methods

### NewUniverseValidateKubernetesOverrides

`func NewUniverseValidateKubernetesOverrides() *UniverseValidateKubernetesOverrides`

NewUniverseValidateKubernetesOverrides instantiates a new UniverseValidateKubernetesOverrides object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseValidateKubernetesOverridesWithDefaults

`func NewUniverseValidateKubernetesOverridesWithDefaults() *UniverseValidateKubernetesOverrides`

NewUniverseValidateKubernetesOverridesWithDefaults instantiates a new UniverseValidateKubernetesOverrides object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetYbSoftwareVersion

`func (o *UniverseValidateKubernetesOverrides) GetYbSoftwareVersion() string`

GetYbSoftwareVersion returns the YbSoftwareVersion field if non-nil, zero value otherwise.

### GetYbSoftwareVersionOk

`func (o *UniverseValidateKubernetesOverrides) GetYbSoftwareVersionOk() (*string, bool)`

GetYbSoftwareVersionOk returns a tuple with the YbSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbSoftwareVersion

`func (o *UniverseValidateKubernetesOverrides) SetYbSoftwareVersion(v string)`

SetYbSoftwareVersion sets YbSoftwareVersion field to given value.

### HasYbSoftwareVersion

`func (o *UniverseValidateKubernetesOverrides) HasYbSoftwareVersion() bool`

HasYbSoftwareVersion returns a boolean if a field has been set.

### GetNodePrefix

`func (o *UniverseValidateKubernetesOverrides) GetNodePrefix() string`

GetNodePrefix returns the NodePrefix field if non-nil, zero value otherwise.

### GetNodePrefixOk

`func (o *UniverseValidateKubernetesOverrides) GetNodePrefixOk() (*string, bool)`

GetNodePrefixOk returns a tuple with the NodePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePrefix

`func (o *UniverseValidateKubernetesOverrides) SetNodePrefix(v string)`

SetNodePrefix sets NodePrefix field to given value.

### HasNodePrefix

`func (o *UniverseValidateKubernetesOverrides) HasNodePrefix() bool`

HasNodePrefix returns a boolean if a field has been set.

### GetIsReadonlyCluster

`func (o *UniverseValidateKubernetesOverrides) GetIsReadonlyCluster() bool`

GetIsReadonlyCluster returns the IsReadonlyCluster field if non-nil, zero value otherwise.

### GetIsReadonlyClusterOk

`func (o *UniverseValidateKubernetesOverrides) GetIsReadonlyClusterOk() (*bool, bool)`

GetIsReadonlyClusterOk returns a tuple with the IsReadonlyCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadonlyCluster

`func (o *UniverseValidateKubernetesOverrides) SetIsReadonlyCluster(v bool)`

SetIsReadonlyCluster sets IsReadonlyCluster field to given value.

### HasIsReadonlyCluster

`func (o *UniverseValidateKubernetesOverrides) HasIsReadonlyCluster() bool`

HasIsReadonlyCluster returns a boolean if a field has been set.

### GetPlacementSpec

`func (o *UniverseValidateKubernetesOverrides) GetPlacementSpec() ClusterPlacementSpec`

GetPlacementSpec returns the PlacementSpec field if non-nil, zero value otherwise.

### GetPlacementSpecOk

`func (o *UniverseValidateKubernetesOverrides) GetPlacementSpecOk() (*ClusterPlacementSpec, bool)`

GetPlacementSpecOk returns a tuple with the PlacementSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementSpec

`func (o *UniverseValidateKubernetesOverrides) SetPlacementSpec(v ClusterPlacementSpec)`

SetPlacementSpec sets PlacementSpec field to given value.

### HasPlacementSpec

`func (o *UniverseValidateKubernetesOverrides) HasPlacementSpec() bool`

HasPlacementSpec returns a boolean if a field has been set.

### GetOverrides

`func (o *UniverseValidateKubernetesOverrides) GetOverrides() string`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *UniverseValidateKubernetesOverrides) GetOverridesOk() (*string, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *UniverseValidateKubernetesOverrides) SetOverrides(v string)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *UniverseValidateKubernetesOverrides) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### GetAzOverrides

`func (o *UniverseValidateKubernetesOverrides) GetAzOverrides() map[string]string`

GetAzOverrides returns the AzOverrides field if non-nil, zero value otherwise.

### GetAzOverridesOk

`func (o *UniverseValidateKubernetesOverrides) GetAzOverridesOk() (*map[string]string, bool)`

GetAzOverridesOk returns a tuple with the AzOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzOverrides

`func (o *UniverseValidateKubernetesOverrides) SetAzOverrides(v map[string]string)`

SetAzOverrides sets AzOverrides field to given value.

### HasAzOverrides

`func (o *UniverseValidateKubernetesOverrides) HasAzOverrides() bool`

HasAzOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


