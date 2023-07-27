# NodeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | **[]string** |  | 
**HasError** | **bool** |  | 
**HasWarning** | **bool** |  | 
**Message** | **string** |  | 
**Metrics** | [**[]Metric**](Metric.md) |  | 
**MetricsOnly** | **bool** |  | 
**Node** | **string** |  | 
**NodeIdentifier** | **string** |  | 
**NodeName** | **string** |  | 
**Process** | **string** |  | 
**TimestampIso** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewNodeData

`func NewNodeData(details []string, hasError bool, hasWarning bool, message string, metrics []Metric, metricsOnly bool, node string, nodeIdentifier string, nodeName string, process string, ) *NodeData`

NewNodeData instantiates a new NodeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeDataWithDefaults

`func NewNodeDataWithDefaults() *NodeData`

NewNodeDataWithDefaults instantiates a new NodeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *NodeData) GetDetails() []string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *NodeData) GetDetailsOk() (*[]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *NodeData) SetDetails(v []string)`

SetDetails sets Details field to given value.


### GetHasError

`func (o *NodeData) GetHasError() bool`

GetHasError returns the HasError field if non-nil, zero value otherwise.

### GetHasErrorOk

`func (o *NodeData) GetHasErrorOk() (*bool, bool)`

GetHasErrorOk returns a tuple with the HasError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasError

`func (o *NodeData) SetHasError(v bool)`

SetHasError sets HasError field to given value.


### GetHasWarning

`func (o *NodeData) GetHasWarning() bool`

GetHasWarning returns the HasWarning field if non-nil, zero value otherwise.

### GetHasWarningOk

`func (o *NodeData) GetHasWarningOk() (*bool, bool)`

GetHasWarningOk returns a tuple with the HasWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWarning

`func (o *NodeData) SetHasWarning(v bool)`

SetHasWarning sets HasWarning field to given value.


### GetMessage

`func (o *NodeData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NodeData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NodeData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMetrics

`func (o *NodeData) GetMetrics() []Metric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *NodeData) GetMetricsOk() (*[]Metric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *NodeData) SetMetrics(v []Metric)`

SetMetrics sets Metrics field to given value.


### GetMetricsOnly

`func (o *NodeData) GetMetricsOnly() bool`

GetMetricsOnly returns the MetricsOnly field if non-nil, zero value otherwise.

### GetMetricsOnlyOk

`func (o *NodeData) GetMetricsOnlyOk() (*bool, bool)`

GetMetricsOnlyOk returns a tuple with the MetricsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsOnly

`func (o *NodeData) SetMetricsOnly(v bool)`

SetMetricsOnly sets MetricsOnly field to given value.


### GetNode

`func (o *NodeData) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *NodeData) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *NodeData) SetNode(v string)`

SetNode sets Node field to given value.


### GetNodeIdentifier

`func (o *NodeData) GetNodeIdentifier() string`

GetNodeIdentifier returns the NodeIdentifier field if non-nil, zero value otherwise.

### GetNodeIdentifierOk

`func (o *NodeData) GetNodeIdentifierOk() (*string, bool)`

GetNodeIdentifierOk returns a tuple with the NodeIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIdentifier

`func (o *NodeData) SetNodeIdentifier(v string)`

SetNodeIdentifier sets NodeIdentifier field to given value.


### GetNodeName

`func (o *NodeData) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *NodeData) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *NodeData) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.


### GetProcess

`func (o *NodeData) GetProcess() string`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *NodeData) GetProcessOk() (*string, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *NodeData) SetProcess(v string)`

SetProcess sets Process field to given value.


### GetTimestampIso

`func (o *NodeData) GetTimestampIso() time.Time`

GetTimestampIso returns the TimestampIso field if non-nil, zero value otherwise.

### GetTimestampIsoOk

`func (o *NodeData) GetTimestampIsoOk() (*time.Time, bool)`

GetTimestampIsoOk returns a tuple with the TimestampIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampIso

`func (o *NodeData) SetTimestampIso(v time.Time)`

SetTimestampIso sets TimestampIso field to given value.

### HasTimestampIso

`func (o *NodeData) HasTimestampIso() bool`

HasTimestampIso returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


