# ClusterNetworkingSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableExposingService** | Pointer to **string** | Whether to create a load balancer service for this cluster. Defaults to NONE. | [optional] [default to "NONE"]
**EnableLb** | Pointer to **bool** | Create target groups if enabled. Used by YBM. | [optional] 
**AzNetworking** | Pointer to [**map[string]AvailabilityZoneNetworking**](AvailabilityZoneNetworking.md) | Granular network settings overridden per Availability Zone identified by AZ uuid. | [optional] 

## Methods

### NewClusterNetworkingSpecAllOf

`func NewClusterNetworkingSpecAllOf() *ClusterNetworkingSpecAllOf`

NewClusterNetworkingSpecAllOf instantiates a new ClusterNetworkingSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNetworkingSpecAllOfWithDefaults

`func NewClusterNetworkingSpecAllOfWithDefaults() *ClusterNetworkingSpecAllOf`

NewClusterNetworkingSpecAllOfWithDefaults instantiates a new ClusterNetworkingSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableExposingService

`func (o *ClusterNetworkingSpecAllOf) GetEnableExposingService() string`

GetEnableExposingService returns the EnableExposingService field if non-nil, zero value otherwise.

### GetEnableExposingServiceOk

`func (o *ClusterNetworkingSpecAllOf) GetEnableExposingServiceOk() (*string, bool)`

GetEnableExposingServiceOk returns a tuple with the EnableExposingService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExposingService

`func (o *ClusterNetworkingSpecAllOf) SetEnableExposingService(v string)`

SetEnableExposingService sets EnableExposingService field to given value.

### HasEnableExposingService

`func (o *ClusterNetworkingSpecAllOf) HasEnableExposingService() bool`

HasEnableExposingService returns a boolean if a field has been set.

### GetEnableLb

`func (o *ClusterNetworkingSpecAllOf) GetEnableLb() bool`

GetEnableLb returns the EnableLb field if non-nil, zero value otherwise.

### GetEnableLbOk

`func (o *ClusterNetworkingSpecAllOf) GetEnableLbOk() (*bool, bool)`

GetEnableLbOk returns a tuple with the EnableLb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLb

`func (o *ClusterNetworkingSpecAllOf) SetEnableLb(v bool)`

SetEnableLb sets EnableLb field to given value.

### HasEnableLb

`func (o *ClusterNetworkingSpecAllOf) HasEnableLb() bool`

HasEnableLb returns a boolean if a field has been set.

### GetAzNetworking

`func (o *ClusterNetworkingSpecAllOf) GetAzNetworking() map[string]AvailabilityZoneNetworking`

GetAzNetworking returns the AzNetworking field if non-nil, zero value otherwise.

### GetAzNetworkingOk

`func (o *ClusterNetworkingSpecAllOf) GetAzNetworkingOk() (*map[string]AvailabilityZoneNetworking, bool)`

GetAzNetworkingOk returns a tuple with the AzNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzNetworking

`func (o *ClusterNetworkingSpecAllOf) SetAzNetworking(v map[string]AvailabilityZoneNetworking)`

SetAzNetworking sets AzNetworking field to given value.

### HasAzNetworking

`func (o *ClusterNetworkingSpecAllOf) HasAzNetworking() bool`

HasAzNetworking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


