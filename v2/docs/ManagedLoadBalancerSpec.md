# ManagedLoadBalancerSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnablePrivate** | Pointer to **bool** | Whether YBA creates a private load balancer in each region. Clients in the VPC of the universe, or in a network peered with it, can reach it. | [optional] [default to false]
**EnablePublic** | Pointer to **bool** | Whether YBA creates a public load balancer in each region, which clients on the internet can reach. Not supported yet. YBA rejects true. | [optional] [default to false]

## Methods

### NewManagedLoadBalancerSpec

`func NewManagedLoadBalancerSpec() *ManagedLoadBalancerSpec`

NewManagedLoadBalancerSpec instantiates a new ManagedLoadBalancerSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedLoadBalancerSpecWithDefaults

`func NewManagedLoadBalancerSpecWithDefaults() *ManagedLoadBalancerSpec`

NewManagedLoadBalancerSpecWithDefaults instantiates a new ManagedLoadBalancerSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnablePrivate

`func (o *ManagedLoadBalancerSpec) GetEnablePrivate() bool`

GetEnablePrivate returns the EnablePrivate field if non-nil, zero value otherwise.

### GetEnablePrivateOk

`func (o *ManagedLoadBalancerSpec) GetEnablePrivateOk() (*bool, bool)`

GetEnablePrivateOk returns a tuple with the EnablePrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePrivate

`func (o *ManagedLoadBalancerSpec) SetEnablePrivate(v bool)`

SetEnablePrivate sets EnablePrivate field to given value.

### HasEnablePrivate

`func (o *ManagedLoadBalancerSpec) HasEnablePrivate() bool`

HasEnablePrivate returns a boolean if a field has been set.

### GetEnablePublic

`func (o *ManagedLoadBalancerSpec) GetEnablePublic() bool`

GetEnablePublic returns the EnablePublic field if non-nil, zero value otherwise.

### GetEnablePublicOk

`func (o *ManagedLoadBalancerSpec) GetEnablePublicOk() (*bool, bool)`

GetEnablePublicOk returns a tuple with the EnablePublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePublic

`func (o *ManagedLoadBalancerSpec) SetEnablePublic(v bool)`

SetEnablePublic sets EnablePublic field to given value.

### HasEnablePublic

`func (o *ManagedLoadBalancerSpec) HasEnablePublic() bool`

HasEnablePublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


