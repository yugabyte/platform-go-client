# PartitionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultPartition** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Placement** | Pointer to [**PlacementInfo**](PlacementInfo.md) |  | [optional] 
**ReplicationFactor** | Pointer to **int32** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewPartitionInfo

`func NewPartitionInfo() *PartitionInfo`

NewPartitionInfo instantiates a new PartitionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionInfoWithDefaults

`func NewPartitionInfoWithDefaults() *PartitionInfo`

NewPartitionInfoWithDefaults instantiates a new PartitionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultPartition

`func (o *PartitionInfo) GetDefaultPartition() bool`

GetDefaultPartition returns the DefaultPartition field if non-nil, zero value otherwise.

### GetDefaultPartitionOk

`func (o *PartitionInfo) GetDefaultPartitionOk() (*bool, bool)`

GetDefaultPartitionOk returns a tuple with the DefaultPartition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPartition

`func (o *PartitionInfo) SetDefaultPartition(v bool)`

SetDefaultPartition sets DefaultPartition field to given value.

### HasDefaultPartition

`func (o *PartitionInfo) HasDefaultPartition() bool`

HasDefaultPartition returns a boolean if a field has been set.

### GetName

`func (o *PartitionInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartitionInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartitionInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartitionInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacement

`func (o *PartitionInfo) GetPlacement() PlacementInfo`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *PartitionInfo) GetPlacementOk() (*PlacementInfo, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *PartitionInfo) SetPlacement(v PlacementInfo)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *PartitionInfo) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### GetReplicationFactor

`func (o *PartitionInfo) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *PartitionInfo) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *PartitionInfo) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *PartitionInfo) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.

### GetUuid

`func (o *PartitionInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PartitionInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PartitionInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PartitionInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


