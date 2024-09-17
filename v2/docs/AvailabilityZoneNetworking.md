# AvailabilityZoneNetworking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tserver** | Pointer to **map[string]interface{}** | (Place holder for) Network settings that can be overridden per tserver or master process for nodes in the cluster. The node instances can be onprem nodes, VMs in GCP/AWS/Azure, or pods in k8s. Part of AvailabilityZoneNetworking. | [optional] 
**Master** | Pointer to **map[string]interface{}** | (Place holder for) Network settings that can be overridden per tserver or master process for nodes in the cluster. The node instances can be onprem nodes, VMs in GCP/AWS/Azure, or pods in k8s. Part of AvailabilityZoneNetworking. | [optional] 
**ProxyConfig** | Pointer to [**NodeProxyConfig**](NodeProxyConfig.md) |  | [optional] 

## Methods

### NewAvailabilityZoneNetworking

`func NewAvailabilityZoneNetworking() *AvailabilityZoneNetworking`

NewAvailabilityZoneNetworking instantiates a new AvailabilityZoneNetworking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityZoneNetworkingWithDefaults

`func NewAvailabilityZoneNetworkingWithDefaults() *AvailabilityZoneNetworking`

NewAvailabilityZoneNetworkingWithDefaults instantiates a new AvailabilityZoneNetworking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTserver

`func (o *AvailabilityZoneNetworking) GetTserver() map[string]interface{}`

GetTserver returns the Tserver field if non-nil, zero value otherwise.

### GetTserverOk

`func (o *AvailabilityZoneNetworking) GetTserverOk() (*map[string]interface{}, bool)`

GetTserverOk returns a tuple with the Tserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserver

`func (o *AvailabilityZoneNetworking) SetTserver(v map[string]interface{})`

SetTserver sets Tserver field to given value.

### HasTserver

`func (o *AvailabilityZoneNetworking) HasTserver() bool`

HasTserver returns a boolean if a field has been set.

### GetMaster

`func (o *AvailabilityZoneNetworking) GetMaster() map[string]interface{}`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *AvailabilityZoneNetworking) GetMasterOk() (*map[string]interface{}, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *AvailabilityZoneNetworking) SetMaster(v map[string]interface{})`

SetMaster sets Master field to given value.

### HasMaster

`func (o *AvailabilityZoneNetworking) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetProxyConfig

`func (o *AvailabilityZoneNetworking) GetProxyConfig() NodeProxyConfig`

GetProxyConfig returns the ProxyConfig field if non-nil, zero value otherwise.

### GetProxyConfigOk

`func (o *AvailabilityZoneNetworking) GetProxyConfigOk() (*NodeProxyConfig, bool)`

GetProxyConfigOk returns a tuple with the ProxyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyConfig

`func (o *AvailabilityZoneNetworking) SetProxyConfig(v NodeProxyConfig)`

SetProxyConfig sets ProxyConfig field to given value.

### HasProxyConfig

`func (o *AvailabilityZoneNetworking) HasProxyConfig() bool`

HasProxyConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


