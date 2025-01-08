# NodeActionFormData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Force** | Pointer to **bool** | Should ignore master unavailability and proceed with the node action | [optional] 
**NodeAction** | **string** | Action to perform on the node. | 

## Methods

### NewNodeActionFormData

`func NewNodeActionFormData(nodeAction string, ) *NodeActionFormData`

NewNodeActionFormData instantiates a new NodeActionFormData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeActionFormDataWithDefaults

`func NewNodeActionFormDataWithDefaults() *NodeActionFormData`

NewNodeActionFormDataWithDefaults instantiates a new NodeActionFormData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForce

`func (o *NodeActionFormData) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *NodeActionFormData) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *NodeActionFormData) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *NodeActionFormData) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetNodeAction

`func (o *NodeActionFormData) GetNodeAction() string`

GetNodeAction returns the NodeAction field if non-nil, zero value otherwise.

### GetNodeActionOk

`func (o *NodeActionFormData) GetNodeActionOk() (*string, bool)`

GetNodeActionOk returns a tuple with the NodeAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAction

`func (o *NodeActionFormData) SetNodeAction(v string)`

SetNodeAction sets NodeAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


