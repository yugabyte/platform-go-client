# UniverseEditEncryptionInTransitAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeToNode** | Pointer to **bool** | Control encryption in transit between YugabyteDB nodes | [optional] 
**ClientToNode** | Pointer to **bool** | Control encryption in transit between YugabyteDB nodes and clients | [optional] 
**RootCa** | Pointer to **string** | Root CA cert for node to node encryption. Required if node_to_node is true | [optional] 
**ClientRootCa** | Pointer to **string** | Root CA used for node to client encryption. Required if client_to_node is true | [optional] 

## Methods

### NewUniverseEditEncryptionInTransitAllOf

`func NewUniverseEditEncryptionInTransitAllOf() *UniverseEditEncryptionInTransitAllOf`

NewUniverseEditEncryptionInTransitAllOf instantiates a new UniverseEditEncryptionInTransitAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseEditEncryptionInTransitAllOfWithDefaults

`func NewUniverseEditEncryptionInTransitAllOfWithDefaults() *UniverseEditEncryptionInTransitAllOf`

NewUniverseEditEncryptionInTransitAllOfWithDefaults instantiates a new UniverseEditEncryptionInTransitAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeToNode

`func (o *UniverseEditEncryptionInTransitAllOf) GetNodeToNode() bool`

GetNodeToNode returns the NodeToNode field if non-nil, zero value otherwise.

### GetNodeToNodeOk

`func (o *UniverseEditEncryptionInTransitAllOf) GetNodeToNodeOk() (*bool, bool)`

GetNodeToNodeOk returns a tuple with the NodeToNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeToNode

`func (o *UniverseEditEncryptionInTransitAllOf) SetNodeToNode(v bool)`

SetNodeToNode sets NodeToNode field to given value.

### HasNodeToNode

`func (o *UniverseEditEncryptionInTransitAllOf) HasNodeToNode() bool`

HasNodeToNode returns a boolean if a field has been set.

### GetClientToNode

`func (o *UniverseEditEncryptionInTransitAllOf) GetClientToNode() bool`

GetClientToNode returns the ClientToNode field if non-nil, zero value otherwise.

### GetClientToNodeOk

`func (o *UniverseEditEncryptionInTransitAllOf) GetClientToNodeOk() (*bool, bool)`

GetClientToNodeOk returns a tuple with the ClientToNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientToNode

`func (o *UniverseEditEncryptionInTransitAllOf) SetClientToNode(v bool)`

SetClientToNode sets ClientToNode field to given value.

### HasClientToNode

`func (o *UniverseEditEncryptionInTransitAllOf) HasClientToNode() bool`

HasClientToNode returns a boolean if a field has been set.

### GetRootCa

`func (o *UniverseEditEncryptionInTransitAllOf) GetRootCa() string`

GetRootCa returns the RootCa field if non-nil, zero value otherwise.

### GetRootCaOk

`func (o *UniverseEditEncryptionInTransitAllOf) GetRootCaOk() (*string, bool)`

GetRootCaOk returns a tuple with the RootCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCa

`func (o *UniverseEditEncryptionInTransitAllOf) SetRootCa(v string)`

SetRootCa sets RootCa field to given value.

### HasRootCa

`func (o *UniverseEditEncryptionInTransitAllOf) HasRootCa() bool`

HasRootCa returns a boolean if a field has been set.

### GetClientRootCa

`func (o *UniverseEditEncryptionInTransitAllOf) GetClientRootCa() string`

GetClientRootCa returns the ClientRootCa field if non-nil, zero value otherwise.

### GetClientRootCaOk

`func (o *UniverseEditEncryptionInTransitAllOf) GetClientRootCaOk() (*string, bool)`

GetClientRootCaOk returns a tuple with the ClientRootCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRootCa

`func (o *UniverseEditEncryptionInTransitAllOf) SetClientRootCa(v string)`

SetClientRootCa sets ClientRootCa field to given value.

### HasClientRootCa

`func (o *UniverseEditEncryptionInTransitAllOf) HasClientRootCa() bool`

HasClientRootCa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


