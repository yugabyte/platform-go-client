# AzureRegionCloudInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroupId** | Pointer to **string** |  | [optional] 
**Vnet** | Pointer to **string** |  | [optional] 
**YbImage** | Pointer to **string** | Deprecated since YBA version 2.20.0. Use provider.imageBundle instead | [optional] 

## Methods

### NewAzureRegionCloudInfo

`func NewAzureRegionCloudInfo() *AzureRegionCloudInfo`

NewAzureRegionCloudInfo instantiates a new AzureRegionCloudInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureRegionCloudInfoWithDefaults

`func NewAzureRegionCloudInfoWithDefaults() *AzureRegionCloudInfo`

NewAzureRegionCloudInfoWithDefaults instantiates a new AzureRegionCloudInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroupId

`func (o *AzureRegionCloudInfo) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *AzureRegionCloudInfo) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *AzureRegionCloudInfo) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.

### HasSecurityGroupId

`func (o *AzureRegionCloudInfo) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### GetVnet

`func (o *AzureRegionCloudInfo) GetVnet() string`

GetVnet returns the Vnet field if non-nil, zero value otherwise.

### GetVnetOk

`func (o *AzureRegionCloudInfo) GetVnetOk() (*string, bool)`

GetVnetOk returns a tuple with the Vnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnet

`func (o *AzureRegionCloudInfo) SetVnet(v string)`

SetVnet sets Vnet field to given value.

### HasVnet

`func (o *AzureRegionCloudInfo) HasVnet() bool`

HasVnet returns a boolean if a field has been set.

### GetYbImage

`func (o *AzureRegionCloudInfo) GetYbImage() string`

GetYbImage returns the YbImage field if non-nil, zero value otherwise.

### GetYbImageOk

`func (o *AzureRegionCloudInfo) GetYbImageOk() (*string, bool)`

GetYbImageOk returns a tuple with the YbImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbImage

`func (o *AzureRegionCloudInfo) SetYbImage(v string)`

SetYbImage sets YbImage field to given value.

### HasYbImage

`func (o *AzureRegionCloudInfo) HasYbImage() bool`

HasYbImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


