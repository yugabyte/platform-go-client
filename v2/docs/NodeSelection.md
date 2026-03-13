# NodeSelection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeNames** | Pointer to **[]string** | List of specific node names to run the script on. | [optional] 
**MastersOnly** | Pointer to **bool** | Run the script only on master nodes from the specified node_names (or all masters if node_names is not provided). | [optional] [default to false]
**TserversOnly** | Pointer to **bool** | Run the script only on tserver nodes from the specified node_names (or all tservers if node_names is not provided). | [optional] [default to false]
**ClusterUuid** | Pointer to **string** | Run the script only on nodes belonging to a specific cluster (e.g., primary or read replica). If not provided, runs on nodes across all clusters in the universe.  | [optional] 
**MaxParallelNodes** | Pointer to **int32** | Maximum number of nodes to execute the script on in parallel. Higher values complete faster but use more resources. Set to 1 for sequential execution.  | [optional] [default to 50]

## Methods

### NewNodeSelection

`func NewNodeSelection() *NodeSelection`

NewNodeSelection instantiates a new NodeSelection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeSelectionWithDefaults

`func NewNodeSelectionWithDefaults() *NodeSelection`

NewNodeSelectionWithDefaults instantiates a new NodeSelection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeNames

`func (o *NodeSelection) GetNodeNames() []string`

GetNodeNames returns the NodeNames field if non-nil, zero value otherwise.

### GetNodeNamesOk

`func (o *NodeSelection) GetNodeNamesOk() (*[]string, bool)`

GetNodeNamesOk returns a tuple with the NodeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeNames

`func (o *NodeSelection) SetNodeNames(v []string)`

SetNodeNames sets NodeNames field to given value.

### HasNodeNames

`func (o *NodeSelection) HasNodeNames() bool`

HasNodeNames returns a boolean if a field has been set.

### GetMastersOnly

`func (o *NodeSelection) GetMastersOnly() bool`

GetMastersOnly returns the MastersOnly field if non-nil, zero value otherwise.

### GetMastersOnlyOk

`func (o *NodeSelection) GetMastersOnlyOk() (*bool, bool)`

GetMastersOnlyOk returns a tuple with the MastersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMastersOnly

`func (o *NodeSelection) SetMastersOnly(v bool)`

SetMastersOnly sets MastersOnly field to given value.

### HasMastersOnly

`func (o *NodeSelection) HasMastersOnly() bool`

HasMastersOnly returns a boolean if a field has been set.

### GetTserversOnly

`func (o *NodeSelection) GetTserversOnly() bool`

GetTserversOnly returns the TserversOnly field if non-nil, zero value otherwise.

### GetTserversOnlyOk

`func (o *NodeSelection) GetTserversOnlyOk() (*bool, bool)`

GetTserversOnlyOk returns a tuple with the TserversOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserversOnly

`func (o *NodeSelection) SetTserversOnly(v bool)`

SetTserversOnly sets TserversOnly field to given value.

### HasTserversOnly

`func (o *NodeSelection) HasTserversOnly() bool`

HasTserversOnly returns a boolean if a field has been set.

### GetClusterUuid

`func (o *NodeSelection) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *NodeSelection) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *NodeSelection) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *NodeSelection) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetMaxParallelNodes

`func (o *NodeSelection) GetMaxParallelNodes() int32`

GetMaxParallelNodes returns the MaxParallelNodes field if non-nil, zero value otherwise.

### GetMaxParallelNodesOk

`func (o *NodeSelection) GetMaxParallelNodesOk() (*int32, bool)`

GetMaxParallelNodesOk returns a tuple with the MaxParallelNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParallelNodes

`func (o *NodeSelection) SetMaxParallelNodes(v int32)`

SetMaxParallelNodes sets MaxParallelNodes field to given value.

### HasMaxParallelNodes

`func (o *NodeSelection) HasMaxParallelNodes() bool`

HasMaxParallelNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


