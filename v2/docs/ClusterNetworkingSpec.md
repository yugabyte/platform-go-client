# ClusterNetworkingSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tserver** | Pointer to **map[string]interface{}** | (Place holder for) Network settings that can be overridden per tserver or master process for nodes in the cluster. The node instances can be onprem nodes, VMs in GCP/AWS/Azure, or pods in k8s. Part of AvailabilityZoneNetworking. | [optional] 
**Master** | Pointer to **map[string]interface{}** | (Place holder for) Network settings that can be overridden per tserver or master process for nodes in the cluster. The node instances can be onprem nodes, VMs in GCP/AWS/Azure, or pods in k8s. Part of AvailabilityZoneNetworking. | [optional] 
**ProxyConfig** | Pointer to [**NodeProxyConfig**](NodeProxyConfig.md) |  | [optional] 
**EnableExposingService** | Pointer to **string** | Whether to create a load balancer service for this cluster. Defaults to NONE. | [optional] [default to "NONE"]
**EnableLb** | Pointer to **bool** | Create target groups if enabled. Used by YBM. | [optional] 
**AzNetworking** | Pointer to [**map[string]AvailabilityZoneNetworking**](AvailabilityZoneNetworking.md) | Granular network settings overridden per Availability Zone identified by AZ uuid. | [optional] 

## Methods

### NewClusterNetworkingSpec

`func NewClusterNetworkingSpec() *ClusterNetworkingSpec`

NewClusterNetworkingSpec instantiates a new ClusterNetworkingSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNetworkingSpecWithDefaults

`func NewClusterNetworkingSpecWithDefaults() *ClusterNetworkingSpec`

NewClusterNetworkingSpecWithDefaults instantiates a new ClusterNetworkingSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTserver

`func (o *ClusterNetworkingSpec) GetTserver() map[string]interface{}`

GetTserver returns the Tserver field if non-nil, zero value otherwise.

### GetTserverOk

`func (o *ClusterNetworkingSpec) GetTserverOk() (*map[string]interface{}, bool)`

GetTserverOk returns a tuple with the Tserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserver

`func (o *ClusterNetworkingSpec) SetTserver(v map[string]interface{})`

SetTserver sets Tserver field to given value.

### HasTserver

`func (o *ClusterNetworkingSpec) HasTserver() bool`

HasTserver returns a boolean if a field has been set.

### GetMaster

`func (o *ClusterNetworkingSpec) GetMaster() map[string]interface{}`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *ClusterNetworkingSpec) GetMasterOk() (*map[string]interface{}, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *ClusterNetworkingSpec) SetMaster(v map[string]interface{})`

SetMaster sets Master field to given value.

### HasMaster

`func (o *ClusterNetworkingSpec) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetProxyConfig

`func (o *ClusterNetworkingSpec) GetProxyConfig() NodeProxyConfig`

GetProxyConfig returns the ProxyConfig field if non-nil, zero value otherwise.

### GetProxyConfigOk

`func (o *ClusterNetworkingSpec) GetProxyConfigOk() (*NodeProxyConfig, bool)`

GetProxyConfigOk returns a tuple with the ProxyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyConfig

`func (o *ClusterNetworkingSpec) SetProxyConfig(v NodeProxyConfig)`

SetProxyConfig sets ProxyConfig field to given value.

### HasProxyConfig

`func (o *ClusterNetworkingSpec) HasProxyConfig() bool`

HasProxyConfig returns a boolean if a field has been set.

### GetEnableExposingService

`func (o *ClusterNetworkingSpec) GetEnableExposingService() string`

GetEnableExposingService returns the EnableExposingService field if non-nil, zero value otherwise.

### GetEnableExposingServiceOk

`func (o *ClusterNetworkingSpec) GetEnableExposingServiceOk() (*string, bool)`

GetEnableExposingServiceOk returns a tuple with the EnableExposingService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExposingService

`func (o *ClusterNetworkingSpec) SetEnableExposingService(v string)`

SetEnableExposingService sets EnableExposingService field to given value.

### HasEnableExposingService

`func (o *ClusterNetworkingSpec) HasEnableExposingService() bool`

HasEnableExposingService returns a boolean if a field has been set.

### GetEnableLb

`func (o *ClusterNetworkingSpec) GetEnableLb() bool`

GetEnableLb returns the EnableLb field if non-nil, zero value otherwise.

### GetEnableLbOk

`func (o *ClusterNetworkingSpec) GetEnableLbOk() (*bool, bool)`

GetEnableLbOk returns a tuple with the EnableLb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLb

`func (o *ClusterNetworkingSpec) SetEnableLb(v bool)`

SetEnableLb sets EnableLb field to given value.

### HasEnableLb

`func (o *ClusterNetworkingSpec) HasEnableLb() bool`

HasEnableLb returns a boolean if a field has been set.

### GetAzNetworking

`func (o *ClusterNetworkingSpec) GetAzNetworking() map[string]AvailabilityZoneNetworking`

GetAzNetworking returns the AzNetworking field if non-nil, zero value otherwise.

### GetAzNetworkingOk

`func (o *ClusterNetworkingSpec) GetAzNetworkingOk() (*map[string]AvailabilityZoneNetworking, bool)`

GetAzNetworkingOk returns a tuple with the AzNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzNetworking

`func (o *ClusterNetworkingSpec) SetAzNetworking(v map[string]AvailabilityZoneNetworking)`

SetAzNetworking sets AzNetworking field to given value.

### HasAzNetworking

`func (o *ClusterNetworkingSpec) HasAzNetworking() bool`

HasAzNetworking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


