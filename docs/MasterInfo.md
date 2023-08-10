# MasterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeerRole** | **string** |  | 
**PrivateIp** | **string** |  | 
**UptimeSeconds** | **int64** |  | 

## Methods

### NewMasterInfo

`func NewMasterInfo(peerRole string, privateIp string, uptimeSeconds int64, ) *MasterInfo`

NewMasterInfo instantiates a new MasterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMasterInfoWithDefaults

`func NewMasterInfoWithDefaults() *MasterInfo`

NewMasterInfoWithDefaults instantiates a new MasterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeerRole

`func (o *MasterInfo) GetPeerRole() string`

GetPeerRole returns the PeerRole field if non-nil, zero value otherwise.

### GetPeerRoleOk

`func (o *MasterInfo) GetPeerRoleOk() (*string, bool)`

GetPeerRoleOk returns a tuple with the PeerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerRole

`func (o *MasterInfo) SetPeerRole(v string)`

SetPeerRole sets PeerRole field to given value.


### GetPrivateIp

`func (o *MasterInfo) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *MasterInfo) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *MasterInfo) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.


### GetUptimeSeconds

`func (o *MasterInfo) GetUptimeSeconds() int64`

GetUptimeSeconds returns the UptimeSeconds field if non-nil, zero value otherwise.

### GetUptimeSecondsOk

`func (o *MasterInfo) GetUptimeSecondsOk() (*int64, bool)`

GetUptimeSecondsOk returns a tuple with the UptimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeSeconds

`func (o *MasterInfo) SetUptimeSeconds(v int64)`

SetUptimeSeconds sets UptimeSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


