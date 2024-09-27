# UniverseNetworkingSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommunicationPorts** | Pointer to [**CommunicationPortsSpec**](CommunicationPortsSpec.md) |  | [optional] 
**AssignPublicIp** | Pointer to **bool** | Whether to assign a public IP for nodes in this Universe. | [optional] [default to true]
**AssignStaticPublicIp** | Pointer to **bool** | Whether to assign a static public IP for nodes in this Universe. | [optional] [default to false]
**EnableIpv6** | Pointer to **bool** | Whether to enable IPv6 on nodes in this cluster. Defaults to false. | [optional] 

## Methods

### NewUniverseNetworkingSpec

`func NewUniverseNetworkingSpec() *UniverseNetworkingSpec`

NewUniverseNetworkingSpec instantiates a new UniverseNetworkingSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseNetworkingSpecWithDefaults

`func NewUniverseNetworkingSpecWithDefaults() *UniverseNetworkingSpec`

NewUniverseNetworkingSpecWithDefaults instantiates a new UniverseNetworkingSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommunicationPorts

`func (o *UniverseNetworkingSpec) GetCommunicationPorts() CommunicationPortsSpec`

GetCommunicationPorts returns the CommunicationPorts field if non-nil, zero value otherwise.

### GetCommunicationPortsOk

`func (o *UniverseNetworkingSpec) GetCommunicationPortsOk() (*CommunicationPortsSpec, bool)`

GetCommunicationPortsOk returns a tuple with the CommunicationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationPorts

`func (o *UniverseNetworkingSpec) SetCommunicationPorts(v CommunicationPortsSpec)`

SetCommunicationPorts sets CommunicationPorts field to given value.

### HasCommunicationPorts

`func (o *UniverseNetworkingSpec) HasCommunicationPorts() bool`

HasCommunicationPorts returns a boolean if a field has been set.

### GetAssignPublicIp

`func (o *UniverseNetworkingSpec) GetAssignPublicIp() bool`

GetAssignPublicIp returns the AssignPublicIp field if non-nil, zero value otherwise.

### GetAssignPublicIpOk

`func (o *UniverseNetworkingSpec) GetAssignPublicIpOk() (*bool, bool)`

GetAssignPublicIpOk returns a tuple with the AssignPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignPublicIp

`func (o *UniverseNetworkingSpec) SetAssignPublicIp(v bool)`

SetAssignPublicIp sets AssignPublicIp field to given value.

### HasAssignPublicIp

`func (o *UniverseNetworkingSpec) HasAssignPublicIp() bool`

HasAssignPublicIp returns a boolean if a field has been set.

### GetAssignStaticPublicIp

`func (o *UniverseNetworkingSpec) GetAssignStaticPublicIp() bool`

GetAssignStaticPublicIp returns the AssignStaticPublicIp field if non-nil, zero value otherwise.

### GetAssignStaticPublicIpOk

`func (o *UniverseNetworkingSpec) GetAssignStaticPublicIpOk() (*bool, bool)`

GetAssignStaticPublicIpOk returns a tuple with the AssignStaticPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignStaticPublicIp

`func (o *UniverseNetworkingSpec) SetAssignStaticPublicIp(v bool)`

SetAssignStaticPublicIp sets AssignStaticPublicIp field to given value.

### HasAssignStaticPublicIp

`func (o *UniverseNetworkingSpec) HasAssignStaticPublicIp() bool`

HasAssignStaticPublicIp returns a boolean if a field has been set.

### GetEnableIpv6

`func (o *UniverseNetworkingSpec) GetEnableIpv6() bool`

GetEnableIpv6 returns the EnableIpv6 field if non-nil, zero value otherwise.

### GetEnableIpv6Ok

`func (o *UniverseNetworkingSpec) GetEnableIpv6Ok() (*bool, bool)`

GetEnableIpv6Ok returns a tuple with the EnableIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIpv6

`func (o *UniverseNetworkingSpec) SetEnableIpv6(v bool)`

SetEnableIpv6 sets EnableIpv6 field to given value.

### HasEnableIpv6

`func (o *UniverseNetworkingSpec) HasEnableIpv6() bool`

HasEnableIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


