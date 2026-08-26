# UniverseEditSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | [**[]ClusterEditSpec**](ClusterEditSpec.md) |  | 
**NetworkingSpec** | Pointer to [**UniverseNetworkingSpec**](UniverseNetworkingSpec.md) |  | [optional] 
**ExpectedUniverseVersion** | **int32** | Expected universe version. Set to -1 to ignore version checking. | 
**UniverseSettings** | Pointer to [**UniverseSettings**](UniverseSettings.md) |  | [optional] 

## Methods

### NewUniverseEditSpec

`func NewUniverseEditSpec(clusters []ClusterEditSpec, expectedUniverseVersion int32, ) *UniverseEditSpec`

NewUniverseEditSpec instantiates a new UniverseEditSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseEditSpecWithDefaults

`func NewUniverseEditSpecWithDefaults() *UniverseEditSpec`

NewUniverseEditSpecWithDefaults instantiates a new UniverseEditSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *UniverseEditSpec) GetClusters() []ClusterEditSpec`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *UniverseEditSpec) GetClustersOk() (*[]ClusterEditSpec, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *UniverseEditSpec) SetClusters(v []ClusterEditSpec)`

SetClusters sets Clusters field to given value.


### GetNetworkingSpec

`func (o *UniverseEditSpec) GetNetworkingSpec() UniverseNetworkingSpec`

GetNetworkingSpec returns the NetworkingSpec field if non-nil, zero value otherwise.

### GetNetworkingSpecOk

`func (o *UniverseEditSpec) GetNetworkingSpecOk() (*UniverseNetworkingSpec, bool)`

GetNetworkingSpecOk returns a tuple with the NetworkingSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkingSpec

`func (o *UniverseEditSpec) SetNetworkingSpec(v UniverseNetworkingSpec)`

SetNetworkingSpec sets NetworkingSpec field to given value.

### HasNetworkingSpec

`func (o *UniverseEditSpec) HasNetworkingSpec() bool`

HasNetworkingSpec returns a boolean if a field has been set.

### GetExpectedUniverseVersion

`func (o *UniverseEditSpec) GetExpectedUniverseVersion() int32`

GetExpectedUniverseVersion returns the ExpectedUniverseVersion field if non-nil, zero value otherwise.

### GetExpectedUniverseVersionOk

`func (o *UniverseEditSpec) GetExpectedUniverseVersionOk() (*int32, bool)`

GetExpectedUniverseVersionOk returns a tuple with the ExpectedUniverseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedUniverseVersion

`func (o *UniverseEditSpec) SetExpectedUniverseVersion(v int32)`

SetExpectedUniverseVersion sets ExpectedUniverseVersion field to given value.


### GetUniverseSettings

`func (o *UniverseEditSpec) GetUniverseSettings() UniverseSettings`

GetUniverseSettings returns the UniverseSettings field if non-nil, zero value otherwise.

### GetUniverseSettingsOk

`func (o *UniverseEditSpec) GetUniverseSettingsOk() (*UniverseSettings, bool)`

GetUniverseSettingsOk returns a tuple with the UniverseSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseSettings

`func (o *UniverseEditSpec) SetUniverseSettings(v UniverseSettings)`

SetUniverseSettings sets UniverseSettings field to given value.

### HasUniverseSettings

`func (o *UniverseEditSpec) HasUniverseSettings() bool`

HasUniverseSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


