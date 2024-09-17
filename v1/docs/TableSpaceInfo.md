# TableSpaceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Tablespace Name | 
**NumReplicas** | Pointer to **int32** | numReplicas | [optional] 
**PlacementBlocks** | [**[]PlacementBlock**](PlacementBlock.md) | placements | 

## Methods

### NewTableSpaceInfo

`func NewTableSpaceInfo(name string, placementBlocks []PlacementBlock, ) *TableSpaceInfo`

NewTableSpaceInfo instantiates a new TableSpaceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableSpaceInfoWithDefaults

`func NewTableSpaceInfoWithDefaults() *TableSpaceInfo`

NewTableSpaceInfoWithDefaults instantiates a new TableSpaceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TableSpaceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TableSpaceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TableSpaceInfo) SetName(v string)`

SetName sets Name field to given value.


### GetNumReplicas

`func (o *TableSpaceInfo) GetNumReplicas() int32`

GetNumReplicas returns the NumReplicas field if non-nil, zero value otherwise.

### GetNumReplicasOk

`func (o *TableSpaceInfo) GetNumReplicasOk() (*int32, bool)`

GetNumReplicasOk returns a tuple with the NumReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumReplicas

`func (o *TableSpaceInfo) SetNumReplicas(v int32)`

SetNumReplicas sets NumReplicas field to given value.

### HasNumReplicas

`func (o *TableSpaceInfo) HasNumReplicas() bool`

HasNumReplicas returns a boolean if a field has been set.

### GetPlacementBlocks

`func (o *TableSpaceInfo) GetPlacementBlocks() []PlacementBlock`

GetPlacementBlocks returns the PlacementBlocks field if non-nil, zero value otherwise.

### GetPlacementBlocksOk

`func (o *TableSpaceInfo) GetPlacementBlocksOk() (*[]PlacementBlock, bool)`

GetPlacementBlocksOk returns a tuple with the PlacementBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementBlocks

`func (o *TableSpaceInfo) SetPlacementBlocks(v []PlacementBlock)`

SetPlacementBlocks sets PlacementBlocks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


