# AwsHostInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | Pointer to **string** | EC2 instance ID | [optional] [readonly] 
**VpcId** | Pointer to **string** | VPC of the instance&#39;s primary network interface | [optional] [readonly] 
**PrivateIp** | Pointer to **string** | Private IP address of the instance | [optional] [readonly] 
**Region** | Pointer to **string** | AWS region | [optional] [readonly] 

## Methods

### NewAwsHostInfo

`func NewAwsHostInfo() *AwsHostInfo`

NewAwsHostInfo instantiates a new AwsHostInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsHostInfoWithDefaults

`func NewAwsHostInfoWithDefaults() *AwsHostInfo`

NewAwsHostInfoWithDefaults instantiates a new AwsHostInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *AwsHostInfo) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AwsHostInfo) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AwsHostInfo) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *AwsHostInfo) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetVpcId

`func (o *AwsHostInfo) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AwsHostInfo) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AwsHostInfo) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *AwsHostInfo) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetPrivateIp

`func (o *AwsHostInfo) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *AwsHostInfo) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *AwsHostInfo) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *AwsHostInfo) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetRegion

`func (o *AwsHostInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AwsHostInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AwsHostInfo) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AwsHostInfo) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


