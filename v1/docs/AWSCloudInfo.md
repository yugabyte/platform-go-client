# AWSCloudInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccessKeyID** | Pointer to **string** |  | [optional] 
**AwsAccessKeySecret** | Pointer to **string** |  | [optional] 
**AwsHostedZoneId** | Pointer to **string** |  | [optional] 
**AwsHostedZoneName** | Pointer to **string** |  | [optional] 
**HostVpcId** | Pointer to **string** |  | [optional] [readonly] 
**HostVpcRegion** | Pointer to **string** |  | [optional] [readonly] 
**UseIMDSv2** | Pointer to **bool** |  | [optional] 
**VpcType** | Pointer to **string** | New/Existing VPC for provider creation | [optional] [readonly] 

## Methods

### NewAWSCloudInfo

`func NewAWSCloudInfo() *AWSCloudInfo`

NewAWSCloudInfo instantiates a new AWSCloudInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSCloudInfoWithDefaults

`func NewAWSCloudInfoWithDefaults() *AWSCloudInfo`

NewAWSCloudInfoWithDefaults instantiates a new AWSCloudInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccessKeyID

`func (o *AWSCloudInfo) GetAwsAccessKeyID() string`

GetAwsAccessKeyID returns the AwsAccessKeyID field if non-nil, zero value otherwise.

### GetAwsAccessKeyIDOk

`func (o *AWSCloudInfo) GetAwsAccessKeyIDOk() (*string, bool)`

GetAwsAccessKeyIDOk returns a tuple with the AwsAccessKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyID

`func (o *AWSCloudInfo) SetAwsAccessKeyID(v string)`

SetAwsAccessKeyID sets AwsAccessKeyID field to given value.

### HasAwsAccessKeyID

`func (o *AWSCloudInfo) HasAwsAccessKeyID() bool`

HasAwsAccessKeyID returns a boolean if a field has been set.

### GetAwsAccessKeySecret

`func (o *AWSCloudInfo) GetAwsAccessKeySecret() string`

GetAwsAccessKeySecret returns the AwsAccessKeySecret field if non-nil, zero value otherwise.

### GetAwsAccessKeySecretOk

`func (o *AWSCloudInfo) GetAwsAccessKeySecretOk() (*string, bool)`

GetAwsAccessKeySecretOk returns a tuple with the AwsAccessKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeySecret

`func (o *AWSCloudInfo) SetAwsAccessKeySecret(v string)`

SetAwsAccessKeySecret sets AwsAccessKeySecret field to given value.

### HasAwsAccessKeySecret

`func (o *AWSCloudInfo) HasAwsAccessKeySecret() bool`

HasAwsAccessKeySecret returns a boolean if a field has been set.

### GetAwsHostedZoneId

`func (o *AWSCloudInfo) GetAwsHostedZoneId() string`

GetAwsHostedZoneId returns the AwsHostedZoneId field if non-nil, zero value otherwise.

### GetAwsHostedZoneIdOk

`func (o *AWSCloudInfo) GetAwsHostedZoneIdOk() (*string, bool)`

GetAwsHostedZoneIdOk returns a tuple with the AwsHostedZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsHostedZoneId

`func (o *AWSCloudInfo) SetAwsHostedZoneId(v string)`

SetAwsHostedZoneId sets AwsHostedZoneId field to given value.

### HasAwsHostedZoneId

`func (o *AWSCloudInfo) HasAwsHostedZoneId() bool`

HasAwsHostedZoneId returns a boolean if a field has been set.

### GetAwsHostedZoneName

`func (o *AWSCloudInfo) GetAwsHostedZoneName() string`

GetAwsHostedZoneName returns the AwsHostedZoneName field if non-nil, zero value otherwise.

### GetAwsHostedZoneNameOk

`func (o *AWSCloudInfo) GetAwsHostedZoneNameOk() (*string, bool)`

GetAwsHostedZoneNameOk returns a tuple with the AwsHostedZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsHostedZoneName

`func (o *AWSCloudInfo) SetAwsHostedZoneName(v string)`

SetAwsHostedZoneName sets AwsHostedZoneName field to given value.

### HasAwsHostedZoneName

`func (o *AWSCloudInfo) HasAwsHostedZoneName() bool`

HasAwsHostedZoneName returns a boolean if a field has been set.

### GetHostVpcId

`func (o *AWSCloudInfo) GetHostVpcId() string`

GetHostVpcId returns the HostVpcId field if non-nil, zero value otherwise.

### GetHostVpcIdOk

`func (o *AWSCloudInfo) GetHostVpcIdOk() (*string, bool)`

GetHostVpcIdOk returns a tuple with the HostVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostVpcId

`func (o *AWSCloudInfo) SetHostVpcId(v string)`

SetHostVpcId sets HostVpcId field to given value.

### HasHostVpcId

`func (o *AWSCloudInfo) HasHostVpcId() bool`

HasHostVpcId returns a boolean if a field has been set.

### GetHostVpcRegion

`func (o *AWSCloudInfo) GetHostVpcRegion() string`

GetHostVpcRegion returns the HostVpcRegion field if non-nil, zero value otherwise.

### GetHostVpcRegionOk

`func (o *AWSCloudInfo) GetHostVpcRegionOk() (*string, bool)`

GetHostVpcRegionOk returns a tuple with the HostVpcRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostVpcRegion

`func (o *AWSCloudInfo) SetHostVpcRegion(v string)`

SetHostVpcRegion sets HostVpcRegion field to given value.

### HasHostVpcRegion

`func (o *AWSCloudInfo) HasHostVpcRegion() bool`

HasHostVpcRegion returns a boolean if a field has been set.

### GetUseIMDSv2

`func (o *AWSCloudInfo) GetUseIMDSv2() bool`

GetUseIMDSv2 returns the UseIMDSv2 field if non-nil, zero value otherwise.

### GetUseIMDSv2Ok

`func (o *AWSCloudInfo) GetUseIMDSv2Ok() (*bool, bool)`

GetUseIMDSv2Ok returns a tuple with the UseIMDSv2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIMDSv2

`func (o *AWSCloudInfo) SetUseIMDSv2(v bool)`

SetUseIMDSv2 sets UseIMDSv2 field to given value.

### HasUseIMDSv2

`func (o *AWSCloudInfo) HasUseIMDSv2() bool`

HasUseIMDSv2 returns a boolean if a field has been set.

### GetVpcType

`func (o *AWSCloudInfo) GetVpcType() string`

GetVpcType returns the VpcType field if non-nil, zero value otherwise.

### GetVpcTypeOk

`func (o *AWSCloudInfo) GetVpcTypeOk() (*string, bool)`

GetVpcTypeOk returns a tuple with the VpcType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcType

`func (o *AWSCloudInfo) SetVpcType(v string)`

SetVpcType sets VpcType field to given value.

### HasVpcType

`func (o *AWSCloudInfo) HasVpcType() bool`

HasVpcType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


