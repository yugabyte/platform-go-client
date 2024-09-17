# ReplicaPlacement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumReplicas** | **int32** |  | 
**PlacementBlocks** | [**[]PlacementBlock**](PlacementBlock.md) |  | 

## Methods

### NewReplicaPlacement

`func NewReplicaPlacement(numReplicas int32, placementBlocks []PlacementBlock, ) *ReplicaPlacement`

NewReplicaPlacement instantiates a new ReplicaPlacement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicaPlacementWithDefaults

`func NewReplicaPlacementWithDefaults() *ReplicaPlacement`

NewReplicaPlacementWithDefaults instantiates a new ReplicaPlacement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumReplicas

`func (o *ReplicaPlacement) GetNumReplicas() int32`

GetNumReplicas returns the NumReplicas field if non-nil, zero value otherwise.

### GetNumReplicasOk

`func (o *ReplicaPlacement) GetNumReplicasOk() (*int32, bool)`

GetNumReplicasOk returns a tuple with the NumReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumReplicas

`func (o *ReplicaPlacement) SetNumReplicas(v int32)`

SetNumReplicas sets NumReplicas field to given value.


### GetPlacementBlocks

`func (o *ReplicaPlacement) GetPlacementBlocks() []PlacementBlock`

GetPlacementBlocks returns the PlacementBlocks field if non-nil, zero value otherwise.

### GetPlacementBlocksOk

`func (o *ReplicaPlacement) GetPlacementBlocksOk() (*[]PlacementBlock, bool)`

GetPlacementBlocksOk returns a tuple with the PlacementBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementBlocks

`func (o *ReplicaPlacement) SetPlacementBlocks(v []PlacementBlock)`

SetPlacementBlocks sets PlacementBlocks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


