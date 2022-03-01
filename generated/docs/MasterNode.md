# MasterNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudInfo** | [**CloudSpecificInfo**](CloudSpecificInfo.md) |  | 
**MasterRpcPort** | **int32** |  | 

## Methods

### NewMasterNode

`func NewMasterNode(cloudInfo CloudSpecificInfo, masterRpcPort int32, ) *MasterNode`

NewMasterNode instantiates a new MasterNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMasterNodeWithDefaults

`func NewMasterNodeWithDefaults() *MasterNode`

NewMasterNodeWithDefaults instantiates a new MasterNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudInfo

`func (o *MasterNode) GetCloudInfo() CloudSpecificInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *MasterNode) GetCloudInfoOk() (*CloudSpecificInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *MasterNode) SetCloudInfo(v CloudSpecificInfo)`

SetCloudInfo sets CloudInfo field to given value.


### GetMasterRpcPort

`func (o *MasterNode) GetMasterRpcPort() int32`

GetMasterRpcPort returns the MasterRpcPort field if non-nil, zero value otherwise.

### GetMasterRpcPortOk

`func (o *MasterNode) GetMasterRpcPortOk() (*int32, bool)`

GetMasterRpcPortOk returns a tuple with the MasterRpcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterRpcPort

`func (o *MasterNode) SetMasterRpcPort(v int32)`

SetMasterRpcPort sets MasterRpcPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


