# NodeQueryDistributionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | **string** |  | 
**NumDelete** | **int32** |  | 
**NumInsert** | **int32** |  | 
**NumSelect** | **int32** |  | 
**NumUpdate** | **int32** |  | 

## Methods

### NewNodeQueryDistributionDetails

`func NewNodeQueryDistributionDetails(node string, numDelete int32, numInsert int32, numSelect int32, numUpdate int32, ) *NodeQueryDistributionDetails`

NewNodeQueryDistributionDetails instantiates a new NodeQueryDistributionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeQueryDistributionDetailsWithDefaults

`func NewNodeQueryDistributionDetailsWithDefaults() *NodeQueryDistributionDetails`

NewNodeQueryDistributionDetailsWithDefaults instantiates a new NodeQueryDistributionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNode

`func (o *NodeQueryDistributionDetails) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *NodeQueryDistributionDetails) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *NodeQueryDistributionDetails) SetNode(v string)`

SetNode sets Node field to given value.


### GetNumDelete

`func (o *NodeQueryDistributionDetails) GetNumDelete() int32`

GetNumDelete returns the NumDelete field if non-nil, zero value otherwise.

### GetNumDeleteOk

`func (o *NodeQueryDistributionDetails) GetNumDeleteOk() (*int32, bool)`

GetNumDeleteOk returns a tuple with the NumDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDelete

`func (o *NodeQueryDistributionDetails) SetNumDelete(v int32)`

SetNumDelete sets NumDelete field to given value.


### GetNumInsert

`func (o *NodeQueryDistributionDetails) GetNumInsert() int32`

GetNumInsert returns the NumInsert field if non-nil, zero value otherwise.

### GetNumInsertOk

`func (o *NodeQueryDistributionDetails) GetNumInsertOk() (*int32, bool)`

GetNumInsertOk returns a tuple with the NumInsert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInsert

`func (o *NodeQueryDistributionDetails) SetNumInsert(v int32)`

SetNumInsert sets NumInsert field to given value.


### GetNumSelect

`func (o *NodeQueryDistributionDetails) GetNumSelect() int32`

GetNumSelect returns the NumSelect field if non-nil, zero value otherwise.

### GetNumSelectOk

`func (o *NodeQueryDistributionDetails) GetNumSelectOk() (*int32, bool)`

GetNumSelectOk returns a tuple with the NumSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSelect

`func (o *NodeQueryDistributionDetails) SetNumSelect(v int32)`

SetNumSelect sets NumSelect field to given value.


### GetNumUpdate

`func (o *NodeQueryDistributionDetails) GetNumUpdate() int32`

GetNumUpdate returns the NumUpdate field if non-nil, zero value otherwise.

### GetNumUpdateOk

`func (o *NodeQueryDistributionDetails) GetNumUpdateOk() (*int32, bool)`

GetNumUpdateOk returns a tuple with the NumUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumUpdate

`func (o *NodeQueryDistributionDetails) SetNumUpdate(v int32)`

SetNumUpdate sets NumUpdate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


