# NodeInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | [**NodeInstanceData**](NodeInstanceData.md) |  | 
**DetailsJson** | Pointer to **string** | Node details (as a JSON object) | [optional] 
**InstanceName** | Pointer to **string** | The node instance&#39;s name | [optional] 
**InstanceTypeCode** | Pointer to **string** | The node&#39;s type code | [optional] 
**NodeName** | Pointer to **string** | The node&#39;s name | [optional] [readonly] 
**NodeUuid** | Pointer to **string** | The node&#39;s UUID | [optional] [readonly] 
**State** | Pointer to **string** | State of on-prem node | [optional] [readonly] 
**ZoneUuid** | Pointer to **string** | The availability zone&#39;s UUID | [optional] 

## Methods

### NewNodeInstance

`func NewNodeInstance(details NodeInstanceData, ) *NodeInstance`

NewNodeInstance instantiates a new NodeInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeInstanceWithDefaults

`func NewNodeInstanceWithDefaults() *NodeInstance`

NewNodeInstanceWithDefaults instantiates a new NodeInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *NodeInstance) GetDetails() NodeInstanceData`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *NodeInstance) GetDetailsOk() (*NodeInstanceData, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *NodeInstance) SetDetails(v NodeInstanceData)`

SetDetails sets Details field to given value.


### GetDetailsJson

`func (o *NodeInstance) GetDetailsJson() string`

GetDetailsJson returns the DetailsJson field if non-nil, zero value otherwise.

### GetDetailsJsonOk

`func (o *NodeInstance) GetDetailsJsonOk() (*string, bool)`

GetDetailsJsonOk returns a tuple with the DetailsJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsJson

`func (o *NodeInstance) SetDetailsJson(v string)`

SetDetailsJson sets DetailsJson field to given value.

### HasDetailsJson

`func (o *NodeInstance) HasDetailsJson() bool`

HasDetailsJson returns a boolean if a field has been set.

### GetInstanceName

`func (o *NodeInstance) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *NodeInstance) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *NodeInstance) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *NodeInstance) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetInstanceTypeCode

`func (o *NodeInstance) GetInstanceTypeCode() string`

GetInstanceTypeCode returns the InstanceTypeCode field if non-nil, zero value otherwise.

### GetInstanceTypeCodeOk

`func (o *NodeInstance) GetInstanceTypeCodeOk() (*string, bool)`

GetInstanceTypeCodeOk returns a tuple with the InstanceTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeCode

`func (o *NodeInstance) SetInstanceTypeCode(v string)`

SetInstanceTypeCode sets InstanceTypeCode field to given value.

### HasInstanceTypeCode

`func (o *NodeInstance) HasInstanceTypeCode() bool`

HasInstanceTypeCode returns a boolean if a field has been set.

### GetNodeName

`func (o *NodeInstance) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *NodeInstance) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *NodeInstance) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *NodeInstance) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetNodeUuid

`func (o *NodeInstance) GetNodeUuid() string`

GetNodeUuid returns the NodeUuid field if non-nil, zero value otherwise.

### GetNodeUuidOk

`func (o *NodeInstance) GetNodeUuidOk() (*string, bool)`

GetNodeUuidOk returns a tuple with the NodeUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeUuid

`func (o *NodeInstance) SetNodeUuid(v string)`

SetNodeUuid sets NodeUuid field to given value.

### HasNodeUuid

`func (o *NodeInstance) HasNodeUuid() bool`

HasNodeUuid returns a boolean if a field has been set.

### GetState

`func (o *NodeInstance) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NodeInstance) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NodeInstance) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NodeInstance) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZoneUuid

`func (o *NodeInstance) GetZoneUuid() string`

GetZoneUuid returns the ZoneUuid field if non-nil, zero value otherwise.

### GetZoneUuidOk

`func (o *NodeInstance) GetZoneUuidOk() (*string, bool)`

GetZoneUuidOk returns a tuple with the ZoneUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneUuid

`func (o *NodeInstance) SetZoneUuid(v string)`

SetZoneUuid sets ZoneUuid field to given value.

### HasZoneUuid

`func (o *NodeInstance) HasZoneUuid() bool`

HasZoneUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


