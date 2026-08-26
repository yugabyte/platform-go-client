# ProviderSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyCode** | Pointer to **string** |  | [optional] 
**AwsInstanceProfile** | Pointer to **string** |  | [optional] 
**AzHelmOverrides** | Pointer to **map[string]string** |  | [optional] 
**EnableLoadBalancer** | Pointer to **bool** |  | [optional] 
**ExposingServiceState** | Pointer to **string** |  | [optional] 
**HelmOverrides** | Pointer to **string** |  | [optional] 
**ImageBundleUUID** | Pointer to **string** |  | [optional] 
**InstanceTags** | Pointer to **map[string]string** |  | [optional] 
**NodesSpecs** | [**RootNodesSpec**](RootNodesSpec.md) |  | 
**ProviderType** | **string** |  | 
**ProviderUUID** | **string** |  | 

## Methods

### NewProviderSpecification

`func NewProviderSpecification(nodesSpecs RootNodesSpec, providerType string, providerUUID string, ) *ProviderSpecification`

NewProviderSpecification instantiates a new ProviderSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSpecificationWithDefaults

`func NewProviderSpecificationWithDefaults() *ProviderSpecification`

NewProviderSpecificationWithDefaults instantiates a new ProviderSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyCode

`func (o *ProviderSpecification) GetAccessKeyCode() string`

GetAccessKeyCode returns the AccessKeyCode field if non-nil, zero value otherwise.

### GetAccessKeyCodeOk

`func (o *ProviderSpecification) GetAccessKeyCodeOk() (*string, bool)`

GetAccessKeyCodeOk returns a tuple with the AccessKeyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyCode

`func (o *ProviderSpecification) SetAccessKeyCode(v string)`

SetAccessKeyCode sets AccessKeyCode field to given value.

### HasAccessKeyCode

`func (o *ProviderSpecification) HasAccessKeyCode() bool`

HasAccessKeyCode returns a boolean if a field has been set.

### GetAwsInstanceProfile

`func (o *ProviderSpecification) GetAwsInstanceProfile() string`

GetAwsInstanceProfile returns the AwsInstanceProfile field if non-nil, zero value otherwise.

### GetAwsInstanceProfileOk

`func (o *ProviderSpecification) GetAwsInstanceProfileOk() (*string, bool)`

GetAwsInstanceProfileOk returns a tuple with the AwsInstanceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsInstanceProfile

`func (o *ProviderSpecification) SetAwsInstanceProfile(v string)`

SetAwsInstanceProfile sets AwsInstanceProfile field to given value.

### HasAwsInstanceProfile

`func (o *ProviderSpecification) HasAwsInstanceProfile() bool`

HasAwsInstanceProfile returns a boolean if a field has been set.

### GetAzHelmOverrides

`func (o *ProviderSpecification) GetAzHelmOverrides() map[string]string`

GetAzHelmOverrides returns the AzHelmOverrides field if non-nil, zero value otherwise.

### GetAzHelmOverridesOk

`func (o *ProviderSpecification) GetAzHelmOverridesOk() (*map[string]string, bool)`

GetAzHelmOverridesOk returns a tuple with the AzHelmOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzHelmOverrides

`func (o *ProviderSpecification) SetAzHelmOverrides(v map[string]string)`

SetAzHelmOverrides sets AzHelmOverrides field to given value.

### HasAzHelmOverrides

`func (o *ProviderSpecification) HasAzHelmOverrides() bool`

HasAzHelmOverrides returns a boolean if a field has been set.

### GetEnableLoadBalancer

`func (o *ProviderSpecification) GetEnableLoadBalancer() bool`

GetEnableLoadBalancer returns the EnableLoadBalancer field if non-nil, zero value otherwise.

### GetEnableLoadBalancerOk

`func (o *ProviderSpecification) GetEnableLoadBalancerOk() (*bool, bool)`

GetEnableLoadBalancerOk returns a tuple with the EnableLoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLoadBalancer

`func (o *ProviderSpecification) SetEnableLoadBalancer(v bool)`

SetEnableLoadBalancer sets EnableLoadBalancer field to given value.

### HasEnableLoadBalancer

`func (o *ProviderSpecification) HasEnableLoadBalancer() bool`

HasEnableLoadBalancer returns a boolean if a field has been set.

### GetExposingServiceState

`func (o *ProviderSpecification) GetExposingServiceState() string`

GetExposingServiceState returns the ExposingServiceState field if non-nil, zero value otherwise.

### GetExposingServiceStateOk

`func (o *ProviderSpecification) GetExposingServiceStateOk() (*string, bool)`

GetExposingServiceStateOk returns a tuple with the ExposingServiceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposingServiceState

`func (o *ProviderSpecification) SetExposingServiceState(v string)`

SetExposingServiceState sets ExposingServiceState field to given value.

### HasExposingServiceState

`func (o *ProviderSpecification) HasExposingServiceState() bool`

HasExposingServiceState returns a boolean if a field has been set.

### GetHelmOverrides

`func (o *ProviderSpecification) GetHelmOverrides() string`

GetHelmOverrides returns the HelmOverrides field if non-nil, zero value otherwise.

### GetHelmOverridesOk

`func (o *ProviderSpecification) GetHelmOverridesOk() (*string, bool)`

GetHelmOverridesOk returns a tuple with the HelmOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmOverrides

`func (o *ProviderSpecification) SetHelmOverrides(v string)`

SetHelmOverrides sets HelmOverrides field to given value.

### HasHelmOverrides

`func (o *ProviderSpecification) HasHelmOverrides() bool`

HasHelmOverrides returns a boolean if a field has been set.

### GetImageBundleUUID

`func (o *ProviderSpecification) GetImageBundleUUID() string`

GetImageBundleUUID returns the ImageBundleUUID field if non-nil, zero value otherwise.

### GetImageBundleUUIDOk

`func (o *ProviderSpecification) GetImageBundleUUIDOk() (*string, bool)`

GetImageBundleUUIDOk returns a tuple with the ImageBundleUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBundleUUID

`func (o *ProviderSpecification) SetImageBundleUUID(v string)`

SetImageBundleUUID sets ImageBundleUUID field to given value.

### HasImageBundleUUID

`func (o *ProviderSpecification) HasImageBundleUUID() bool`

HasImageBundleUUID returns a boolean if a field has been set.

### GetInstanceTags

`func (o *ProviderSpecification) GetInstanceTags() map[string]string`

GetInstanceTags returns the InstanceTags field if non-nil, zero value otherwise.

### GetInstanceTagsOk

`func (o *ProviderSpecification) GetInstanceTagsOk() (*map[string]string, bool)`

GetInstanceTagsOk returns a tuple with the InstanceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTags

`func (o *ProviderSpecification) SetInstanceTags(v map[string]string)`

SetInstanceTags sets InstanceTags field to given value.

### HasInstanceTags

`func (o *ProviderSpecification) HasInstanceTags() bool`

HasInstanceTags returns a boolean if a field has been set.

### GetNodesSpecs

`func (o *ProviderSpecification) GetNodesSpecs() RootNodesSpec`

GetNodesSpecs returns the NodesSpecs field if non-nil, zero value otherwise.

### GetNodesSpecsOk

`func (o *ProviderSpecification) GetNodesSpecsOk() (*RootNodesSpec, bool)`

GetNodesSpecsOk returns a tuple with the NodesSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesSpecs

`func (o *ProviderSpecification) SetNodesSpecs(v RootNodesSpec)`

SetNodesSpecs sets NodesSpecs field to given value.


### GetProviderType

`func (o *ProviderSpecification) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *ProviderSpecification) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *ProviderSpecification) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetProviderUUID

`func (o *ProviderSpecification) GetProviderUUID() string`

GetProviderUUID returns the ProviderUUID field if non-nil, zero value otherwise.

### GetProviderUUIDOk

`func (o *ProviderSpecification) GetProviderUUIDOk() (*string, bool)`

GetProviderUUIDOk returns a tuple with the ProviderUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUUID

`func (o *ProviderSpecification) SetProviderUUID(v string)`

SetProviderUUID sets ProviderUUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


