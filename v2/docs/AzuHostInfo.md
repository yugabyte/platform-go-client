# AzuHostInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | Pointer to **string** | Azure subscription of the VM | [optional] [readonly] 
**ResourceGroup** | Pointer to **string** | Resource group of the VM | [optional] [readonly] 
**VmId** | Pointer to **string** | Azure VM ID | [optional] [readonly] 
**PrivateIp** | Pointer to **string** | Private IP address of the VM&#39;s first network interface | [optional] [readonly] 
**Region** | Pointer to **string** | Azure region (location) of the VM | [optional] [readonly] 

## Methods

### NewAzuHostInfo

`func NewAzuHostInfo() *AzuHostInfo`

NewAzuHostInfo instantiates a new AzuHostInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzuHostInfoWithDefaults

`func NewAzuHostInfoWithDefaults() *AzuHostInfo`

NewAzuHostInfoWithDefaults instantiates a new AzuHostInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *AzuHostInfo) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzuHostInfo) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzuHostInfo) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AzuHostInfo) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetResourceGroup

`func (o *AzuHostInfo) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AzuHostInfo) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AzuHostInfo) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *AzuHostInfo) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetVmId

`func (o *AzuHostInfo) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *AzuHostInfo) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *AzuHostInfo) SetVmId(v string)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *AzuHostInfo) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### GetPrivateIp

`func (o *AzuHostInfo) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *AzuHostInfo) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *AzuHostInfo) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *AzuHostInfo) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetRegion

`func (o *AzuHostInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AzuHostInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AzuHostInfo) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AzuHostInfo) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


