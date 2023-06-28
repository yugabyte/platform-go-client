# AzureCloudInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzuClientId** | Pointer to **string** |  | [optional] 
**AzuClientSecret** | Pointer to **string** |  | [optional] 
**AzuHostedZoneId** | Pointer to **string** |  | [optional] 
**AzuRG** | Pointer to **string** |  | [optional] 
**AzuSubscriptionId** | Pointer to **string** |  | [optional] 
**AzuTenantId** | Pointer to **string** |  | [optional] 
**VpcType** | Pointer to **string** | New/Existing VPC for provider creation | [optional] [readonly] 

## Methods

### NewAzureCloudInfo

`func NewAzureCloudInfo() *AzureCloudInfo`

NewAzureCloudInfo instantiates a new AzureCloudInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCloudInfoWithDefaults

`func NewAzureCloudInfoWithDefaults() *AzureCloudInfo`

NewAzureCloudInfoWithDefaults instantiates a new AzureCloudInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzuClientId

`func (o *AzureCloudInfo) GetAzuClientId() string`

GetAzuClientId returns the AzuClientId field if non-nil, zero value otherwise.

### GetAzuClientIdOk

`func (o *AzureCloudInfo) GetAzuClientIdOk() (*string, bool)`

GetAzuClientIdOk returns a tuple with the AzuClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzuClientId

`func (o *AzureCloudInfo) SetAzuClientId(v string)`

SetAzuClientId sets AzuClientId field to given value.

### HasAzuClientId

`func (o *AzureCloudInfo) HasAzuClientId() bool`

HasAzuClientId returns a boolean if a field has been set.

### GetAzuClientSecret

`func (o *AzureCloudInfo) GetAzuClientSecret() string`

GetAzuClientSecret returns the AzuClientSecret field if non-nil, zero value otherwise.

### GetAzuClientSecretOk

`func (o *AzureCloudInfo) GetAzuClientSecretOk() (*string, bool)`

GetAzuClientSecretOk returns a tuple with the AzuClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzuClientSecret

`func (o *AzureCloudInfo) SetAzuClientSecret(v string)`

SetAzuClientSecret sets AzuClientSecret field to given value.

### HasAzuClientSecret

`func (o *AzureCloudInfo) HasAzuClientSecret() bool`

HasAzuClientSecret returns a boolean if a field has been set.

### GetAzuHostedZoneId

`func (o *AzureCloudInfo) GetAzuHostedZoneId() string`

GetAzuHostedZoneId returns the AzuHostedZoneId field if non-nil, zero value otherwise.

### GetAzuHostedZoneIdOk

`func (o *AzureCloudInfo) GetAzuHostedZoneIdOk() (*string, bool)`

GetAzuHostedZoneIdOk returns a tuple with the AzuHostedZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzuHostedZoneId

`func (o *AzureCloudInfo) SetAzuHostedZoneId(v string)`

SetAzuHostedZoneId sets AzuHostedZoneId field to given value.

### HasAzuHostedZoneId

`func (o *AzureCloudInfo) HasAzuHostedZoneId() bool`

HasAzuHostedZoneId returns a boolean if a field has been set.

### GetAzuRG

`func (o *AzureCloudInfo) GetAzuRG() string`

GetAzuRG returns the AzuRG field if non-nil, zero value otherwise.

### GetAzuRGOk

`func (o *AzureCloudInfo) GetAzuRGOk() (*string, bool)`

GetAzuRGOk returns a tuple with the AzuRG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzuRG

`func (o *AzureCloudInfo) SetAzuRG(v string)`

SetAzuRG sets AzuRG field to given value.

### HasAzuRG

`func (o *AzureCloudInfo) HasAzuRG() bool`

HasAzuRG returns a boolean if a field has been set.

### GetAzuSubscriptionId

`func (o *AzureCloudInfo) GetAzuSubscriptionId() string`

GetAzuSubscriptionId returns the AzuSubscriptionId field if non-nil, zero value otherwise.

### GetAzuSubscriptionIdOk

`func (o *AzureCloudInfo) GetAzuSubscriptionIdOk() (*string, bool)`

GetAzuSubscriptionIdOk returns a tuple with the AzuSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzuSubscriptionId

`func (o *AzureCloudInfo) SetAzuSubscriptionId(v string)`

SetAzuSubscriptionId sets AzuSubscriptionId field to given value.

### HasAzuSubscriptionId

`func (o *AzureCloudInfo) HasAzuSubscriptionId() bool`

HasAzuSubscriptionId returns a boolean if a field has been set.

### GetAzuTenantId

`func (o *AzureCloudInfo) GetAzuTenantId() string`

GetAzuTenantId returns the AzuTenantId field if non-nil, zero value otherwise.

### GetAzuTenantIdOk

`func (o *AzureCloudInfo) GetAzuTenantIdOk() (*string, bool)`

GetAzuTenantIdOk returns a tuple with the AzuTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzuTenantId

`func (o *AzureCloudInfo) SetAzuTenantId(v string)`

SetAzuTenantId sets AzuTenantId field to given value.

### HasAzuTenantId

`func (o *AzureCloudInfo) HasAzuTenantId() bool`

HasAzuTenantId returns a boolean if a field has been set.

### GetVpcType

`func (o *AzureCloudInfo) GetVpcType() string`

GetVpcType returns the VpcType field if non-nil, zero value otherwise.

### GetVpcTypeOk

`func (o *AzureCloudInfo) GetVpcTypeOk() (*string, bool)`

GetVpcTypeOk returns a tuple with the VpcType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcType

`func (o *AzureCloudInfo) SetVpcType(v string)`

SetVpcType sets VpcType field to given value.

### HasVpcType

`func (o *AzureCloudInfo) HasVpcType() bool`

HasVpcType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


