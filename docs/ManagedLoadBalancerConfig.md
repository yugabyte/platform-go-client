# ManagedLoadBalancerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnablePrivate** | Pointer to **bool** | Create a private L4 load balancer in each region of the cluster | [optional] 
**EnablePublic** | Pointer to **bool** | Create a public L4 load balancer in each region of the cluster. Not supported yet. Must be false. | [optional] 

## Methods

### NewManagedLoadBalancerConfig

`func NewManagedLoadBalancerConfig() *ManagedLoadBalancerConfig`

NewManagedLoadBalancerConfig instantiates a new ManagedLoadBalancerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedLoadBalancerConfigWithDefaults

`func NewManagedLoadBalancerConfigWithDefaults() *ManagedLoadBalancerConfig`

NewManagedLoadBalancerConfigWithDefaults instantiates a new ManagedLoadBalancerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnablePrivate

`func (o *ManagedLoadBalancerConfig) GetEnablePrivate() bool`

GetEnablePrivate returns the EnablePrivate field if non-nil, zero value otherwise.

### GetEnablePrivateOk

`func (o *ManagedLoadBalancerConfig) GetEnablePrivateOk() (*bool, bool)`

GetEnablePrivateOk returns a tuple with the EnablePrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePrivate

`func (o *ManagedLoadBalancerConfig) SetEnablePrivate(v bool)`

SetEnablePrivate sets EnablePrivate field to given value.

### HasEnablePrivate

`func (o *ManagedLoadBalancerConfig) HasEnablePrivate() bool`

HasEnablePrivate returns a boolean if a field has been set.

### GetEnablePublic

`func (o *ManagedLoadBalancerConfig) GetEnablePublic() bool`

GetEnablePublic returns the EnablePublic field if non-nil, zero value otherwise.

### GetEnablePublicOk

`func (o *ManagedLoadBalancerConfig) GetEnablePublicOk() (*bool, bool)`

GetEnablePublicOk returns a tuple with the EnablePublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePublic

`func (o *ManagedLoadBalancerConfig) SetEnablePublic(v bool)`

SetEnablePublic sets EnablePublic field to given value.

### HasEnablePublic

`func (o *ManagedLoadBalancerConfig) HasEnablePublic() bool`

HasEnablePublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


