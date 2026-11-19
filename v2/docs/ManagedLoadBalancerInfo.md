# ManagedLoadBalancerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scheme** | Pointer to [**ManagedLoadBalancerScheme**](ManagedLoadBalancerScheme.md) | Who can reach the load balancer. Today always PRIVATE. | [optional] [readonly] 
**RegionCode** | Pointer to **string** | The region of the load balancer. | [optional] [readonly] 
**AzUuids** | Pointer to **[]string** | The availability zones that the load balancer serves, as the UUIDs of the provider zones. Today every zone of the region. A later release can give a region several load balancers that each serve some of its zones. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the load balancer in the cloud. | [optional] [readonly] 
**Address** | Pointer to **string** | The address that the cloud gives the load balancer. A DNS name on AWS, an IP address on GCP and Azure. Empty until the load balancer is created. | [optional] [readonly] 

## Methods

### NewManagedLoadBalancerInfo

`func NewManagedLoadBalancerInfo() *ManagedLoadBalancerInfo`

NewManagedLoadBalancerInfo instantiates a new ManagedLoadBalancerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedLoadBalancerInfoWithDefaults

`func NewManagedLoadBalancerInfoWithDefaults() *ManagedLoadBalancerInfo`

NewManagedLoadBalancerInfoWithDefaults instantiates a new ManagedLoadBalancerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheme

`func (o *ManagedLoadBalancerInfo) GetScheme() ManagedLoadBalancerScheme`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *ManagedLoadBalancerInfo) GetSchemeOk() (*ManagedLoadBalancerScheme, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *ManagedLoadBalancerInfo) SetScheme(v ManagedLoadBalancerScheme)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *ManagedLoadBalancerInfo) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetRegionCode

`func (o *ManagedLoadBalancerInfo) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *ManagedLoadBalancerInfo) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *ManagedLoadBalancerInfo) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *ManagedLoadBalancerInfo) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetAzUuids

`func (o *ManagedLoadBalancerInfo) GetAzUuids() []string`

GetAzUuids returns the AzUuids field if non-nil, zero value otherwise.

### GetAzUuidsOk

`func (o *ManagedLoadBalancerInfo) GetAzUuidsOk() (*[]string, bool)`

GetAzUuidsOk returns a tuple with the AzUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzUuids

`func (o *ManagedLoadBalancerInfo) SetAzUuids(v []string)`

SetAzUuids sets AzUuids field to given value.

### HasAzUuids

`func (o *ManagedLoadBalancerInfo) HasAzUuids() bool`

HasAzUuids returns a boolean if a field has been set.

### GetName

`func (o *ManagedLoadBalancerInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedLoadBalancerInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedLoadBalancerInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManagedLoadBalancerInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *ManagedLoadBalancerInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ManagedLoadBalancerInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ManagedLoadBalancerInfo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ManagedLoadBalancerInfo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


