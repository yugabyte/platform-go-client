# PlacementBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloud** | **string** | Cloud | 
**LeaderPreference** | Pointer to **int32** | Leader preference | [optional] 
**MinNumReplicas** | Pointer to **int32** | Minimum replicas | [optional] 
**Region** | **string** | Region | 
**Zone** | **string** | Zone | 

## Methods

### NewPlacementBlock

`func NewPlacementBlock(cloud string, region string, zone string, ) *PlacementBlock`

NewPlacementBlock instantiates a new PlacementBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlacementBlockWithDefaults

`func NewPlacementBlockWithDefaults() *PlacementBlock`

NewPlacementBlockWithDefaults instantiates a new PlacementBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloud

`func (o *PlacementBlock) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *PlacementBlock) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *PlacementBlock) SetCloud(v string)`

SetCloud sets Cloud field to given value.


### GetLeaderPreference

`func (o *PlacementBlock) GetLeaderPreference() int32`

GetLeaderPreference returns the LeaderPreference field if non-nil, zero value otherwise.

### GetLeaderPreferenceOk

`func (o *PlacementBlock) GetLeaderPreferenceOk() (*int32, bool)`

GetLeaderPreferenceOk returns a tuple with the LeaderPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaderPreference

`func (o *PlacementBlock) SetLeaderPreference(v int32)`

SetLeaderPreference sets LeaderPreference field to given value.

### HasLeaderPreference

`func (o *PlacementBlock) HasLeaderPreference() bool`

HasLeaderPreference returns a boolean if a field has been set.

### GetMinNumReplicas

`func (o *PlacementBlock) GetMinNumReplicas() int32`

GetMinNumReplicas returns the MinNumReplicas field if non-nil, zero value otherwise.

### GetMinNumReplicasOk

`func (o *PlacementBlock) GetMinNumReplicasOk() (*int32, bool)`

GetMinNumReplicasOk returns a tuple with the MinNumReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNumReplicas

`func (o *PlacementBlock) SetMinNumReplicas(v int32)`

SetMinNumReplicas sets MinNumReplicas field to given value.

### HasMinNumReplicas

`func (o *PlacementBlock) HasMinNumReplicas() bool`

HasMinNumReplicas returns a boolean if a field has been set.

### GetRegion

`func (o *PlacementBlock) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PlacementBlock) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PlacementBlock) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetZone

`func (o *PlacementBlock) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *PlacementBlock) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *PlacementBlock) SetZone(v string)`

SetZone sets Zone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


