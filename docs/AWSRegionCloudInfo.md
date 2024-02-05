# AWSRegionCloudInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** | &lt;b style&#x3D;\&quot;color:#ff0000\&quot;&gt;Deprecated since YBA version 2.20.0.&lt;/b&gt; Use provider.imageBundle instead | [optional] 
**SecurityGroupId** | Pointer to **string** |  | [optional] 
**Vnet** | Pointer to **string** |  | [optional] 
**YbImage** | Pointer to **string** | &lt;b style&#x3D;\&quot;color:#ff0000\&quot;&gt;Deprecated since YBA version 2.20.0.&lt;/b&gt; Use provider.imageBundle instead | [optional] 

## Methods

### NewAWSRegionCloudInfo

`func NewAWSRegionCloudInfo() *AWSRegionCloudInfo`

NewAWSRegionCloudInfo instantiates a new AWSRegionCloudInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSRegionCloudInfoWithDefaults

`func NewAWSRegionCloudInfoWithDefaults() *AWSRegionCloudInfo`

NewAWSRegionCloudInfoWithDefaults instantiates a new AWSRegionCloudInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *AWSRegionCloudInfo) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *AWSRegionCloudInfo) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *AWSRegionCloudInfo) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *AWSRegionCloudInfo) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetSecurityGroupId

`func (o *AWSRegionCloudInfo) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *AWSRegionCloudInfo) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *AWSRegionCloudInfo) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.

### HasSecurityGroupId

`func (o *AWSRegionCloudInfo) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### GetVnet

`func (o *AWSRegionCloudInfo) GetVnet() string`

GetVnet returns the Vnet field if non-nil, zero value otherwise.

### GetVnetOk

`func (o *AWSRegionCloudInfo) GetVnetOk() (*string, bool)`

GetVnetOk returns a tuple with the Vnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnet

`func (o *AWSRegionCloudInfo) SetVnet(v string)`

SetVnet sets Vnet field to given value.

### HasVnet

`func (o *AWSRegionCloudInfo) HasVnet() bool`

HasVnet returns a boolean if a field has been set.

### GetYbImage

`func (o *AWSRegionCloudInfo) GetYbImage() string`

GetYbImage returns the YbImage field if non-nil, zero value otherwise.

### GetYbImageOk

`func (o *AWSRegionCloudInfo) GetYbImageOk() (*string, bool)`

GetYbImageOk returns a tuple with the YbImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbImage

`func (o *AWSRegionCloudInfo) SetYbImage(v string)`

SetYbImage sets YbImage field to given value.

### HasYbImage

`func (o *AWSRegionCloudInfo) HasYbImage() bool`

HasYbImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


