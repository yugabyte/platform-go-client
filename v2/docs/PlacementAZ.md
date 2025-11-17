# PlacementAZ

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The AZ id | [optional] 
**Name** | Pointer to **string** | The AZ name | [optional] 
**ReplicationFactor** | Pointer to **int32** | The minimum number of copies of data we should place into this AZ. | [optional] 
**Subnet** | Pointer to **string** | The subnet in the AZ. | [optional] 
**SecondarySubnet** | Pointer to **string** | The secondary subnet in the AZ. | [optional] 
**NumNodesInAz** | Pointer to **int32** | Number of nodes in each AZ. | [optional] 
**LeaderAffinity** | Pointer to **bool** | Affinitizes raft leaders to this AZ. | [optional] 
**LeaderPreference** | Pointer to **int32** | Priority of zone (for leaders placement). Values have to be contiguous non-zero integers. Multiple zones can have the same value. A lower value indicates higher zone priority. | [optional] 
**LbName** | Pointer to **string** | The Load Balancer id. | [optional] 

## Methods

### NewPlacementAZ

`func NewPlacementAZ() *PlacementAZ`

NewPlacementAZ instantiates a new PlacementAZ object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlacementAZWithDefaults

`func NewPlacementAZWithDefaults() *PlacementAZ`

NewPlacementAZWithDefaults instantiates a new PlacementAZ object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PlacementAZ) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PlacementAZ) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PlacementAZ) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PlacementAZ) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *PlacementAZ) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlacementAZ) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlacementAZ) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlacementAZ) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReplicationFactor

`func (o *PlacementAZ) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *PlacementAZ) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *PlacementAZ) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *PlacementAZ) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.

### GetSubnet

`func (o *PlacementAZ) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *PlacementAZ) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *PlacementAZ) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *PlacementAZ) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetSecondarySubnet

`func (o *PlacementAZ) GetSecondarySubnet() string`

GetSecondarySubnet returns the SecondarySubnet field if non-nil, zero value otherwise.

### GetSecondarySubnetOk

`func (o *PlacementAZ) GetSecondarySubnetOk() (*string, bool)`

GetSecondarySubnetOk returns a tuple with the SecondarySubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySubnet

`func (o *PlacementAZ) SetSecondarySubnet(v string)`

SetSecondarySubnet sets SecondarySubnet field to given value.

### HasSecondarySubnet

`func (o *PlacementAZ) HasSecondarySubnet() bool`

HasSecondarySubnet returns a boolean if a field has been set.

### GetNumNodesInAz

`func (o *PlacementAZ) GetNumNodesInAz() int32`

GetNumNodesInAz returns the NumNodesInAz field if non-nil, zero value otherwise.

### GetNumNodesInAzOk

`func (o *PlacementAZ) GetNumNodesInAzOk() (*int32, bool)`

GetNumNodesInAzOk returns a tuple with the NumNodesInAz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNodesInAz

`func (o *PlacementAZ) SetNumNodesInAz(v int32)`

SetNumNodesInAz sets NumNodesInAz field to given value.

### HasNumNodesInAz

`func (o *PlacementAZ) HasNumNodesInAz() bool`

HasNumNodesInAz returns a boolean if a field has been set.

### GetLeaderAffinity

`func (o *PlacementAZ) GetLeaderAffinity() bool`

GetLeaderAffinity returns the LeaderAffinity field if non-nil, zero value otherwise.

### GetLeaderAffinityOk

`func (o *PlacementAZ) GetLeaderAffinityOk() (*bool, bool)`

GetLeaderAffinityOk returns a tuple with the LeaderAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaderAffinity

`func (o *PlacementAZ) SetLeaderAffinity(v bool)`

SetLeaderAffinity sets LeaderAffinity field to given value.

### HasLeaderAffinity

`func (o *PlacementAZ) HasLeaderAffinity() bool`

HasLeaderAffinity returns a boolean if a field has been set.

### GetLeaderPreference

`func (o *PlacementAZ) GetLeaderPreference() int32`

GetLeaderPreference returns the LeaderPreference field if non-nil, zero value otherwise.

### GetLeaderPreferenceOk

`func (o *PlacementAZ) GetLeaderPreferenceOk() (*int32, bool)`

GetLeaderPreferenceOk returns a tuple with the LeaderPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaderPreference

`func (o *PlacementAZ) SetLeaderPreference(v int32)`

SetLeaderPreference sets LeaderPreference field to given value.

### HasLeaderPreference

`func (o *PlacementAZ) HasLeaderPreference() bool`

HasLeaderPreference returns a boolean if a field has been set.

### GetLbName

`func (o *PlacementAZ) GetLbName() string`

GetLbName returns the LbName field if non-nil, zero value otherwise.

### GetLbNameOk

`func (o *PlacementAZ) GetLbNameOk() (*string, bool)`

GetLbNameOk returns a tuple with the LbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbName

`func (o *PlacementAZ) SetLbName(v string)`

SetLbName sets LbName field to given value.

### HasLbName

`func (o *PlacementAZ) HasLbName() bool`

HasLbName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


