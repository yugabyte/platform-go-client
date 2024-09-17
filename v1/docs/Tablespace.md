# Tablespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplicaPlacement** | [**ReplicaPlacement**](ReplicaPlacement.md) |  | 
**TablespaceName** | **string** |  | 

## Methods

### NewTablespace

`func NewTablespace(replicaPlacement ReplicaPlacement, tablespaceName string, ) *Tablespace`

NewTablespace instantiates a new Tablespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTablespaceWithDefaults

`func NewTablespaceWithDefaults() *Tablespace`

NewTablespaceWithDefaults instantiates a new Tablespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicaPlacement

`func (o *Tablespace) GetReplicaPlacement() ReplicaPlacement`

GetReplicaPlacement returns the ReplicaPlacement field if non-nil, zero value otherwise.

### GetReplicaPlacementOk

`func (o *Tablespace) GetReplicaPlacementOk() (*ReplicaPlacement, bool)`

GetReplicaPlacementOk returns a tuple with the ReplicaPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaPlacement

`func (o *Tablespace) SetReplicaPlacement(v ReplicaPlacement)`

SetReplicaPlacement sets ReplicaPlacement field to given value.


### GetTablespaceName

`func (o *Tablespace) GetTablespaceName() string`

GetTablespaceName returns the TablespaceName field if non-nil, zero value otherwise.

### GetTablespaceNameOk

`func (o *Tablespace) GetTablespaceNameOk() (*string, bool)`

GetTablespaceNameOk returns a tuple with the TablespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTablespaceName

`func (o *Tablespace) SetTablespaceName(v string)`

SetTablespaceName sets TablespaceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


