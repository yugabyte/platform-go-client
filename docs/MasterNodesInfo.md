# MasterNodesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** |  | 
**IsLeader** | **bool** |  | 
**MasterUUID** | **string** |  | 
**Port** | **int32** |  | 

## Methods

### NewMasterNodesInfo

`func NewMasterNodesInfo(host string, isLeader bool, masterUUID string, port int32, ) *MasterNodesInfo`

NewMasterNodesInfo instantiates a new MasterNodesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMasterNodesInfoWithDefaults

`func NewMasterNodesInfoWithDefaults() *MasterNodesInfo`

NewMasterNodesInfoWithDefaults instantiates a new MasterNodesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *MasterNodesInfo) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MasterNodesInfo) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MasterNodesInfo) SetHost(v string)`

SetHost sets Host field to given value.


### GetIsLeader

`func (o *MasterNodesInfo) GetIsLeader() bool`

GetIsLeader returns the IsLeader field if non-nil, zero value otherwise.

### GetIsLeaderOk

`func (o *MasterNodesInfo) GetIsLeaderOk() (*bool, bool)`

GetIsLeaderOk returns a tuple with the IsLeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLeader

`func (o *MasterNodesInfo) SetIsLeader(v bool)`

SetIsLeader sets IsLeader field to given value.


### GetMasterUUID

`func (o *MasterNodesInfo) GetMasterUUID() string`

GetMasterUUID returns the MasterUUID field if non-nil, zero value otherwise.

### GetMasterUUIDOk

`func (o *MasterNodesInfo) GetMasterUUIDOk() (*string, bool)`

GetMasterUUIDOk returns a tuple with the MasterUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterUUID

`func (o *MasterNodesInfo) SetMasterUUID(v string)`

SetMasterUUID sets MasterUUID field to given value.


### GetPort

`func (o *MasterNodesInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MasterNodesInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MasterNodesInfo) SetPort(v int32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


