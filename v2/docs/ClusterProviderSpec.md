# ClusterProviderSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | Cloud provider UUID | 
**RegionList** | Pointer to **[]string** | The list of regions in the cloud provider to place data replicas | [optional] 
**PreferredRegion** | Pointer to **string** | The region to nominate as the preferred region in a geo-partitioned multi-region cluster | [optional] 
**AccessKeyCode** | Pointer to **string** | One of the SSH access keys defined in Cloud Provider to be configured on nodes VMs. Required for AWS, Azure and GCP Cloud Providers. | [optional] 
**AwsInstanceProfile** | Pointer to **string** | The AWS IAM instance profile ARN to use for the nodes in this cluster. Applicable only for nodes on AWS Cloud Provider. If specified, YugabyteDB Anywhere will use this instance profile instead of the access key. | [optional] 
**ImageBundleUuid** | Pointer to **string** | Image bundle UUID to use for node VM image. Refers to one of the image bundles defined in the cloud provider. | [optional] 
**HelmOverrides** | Pointer to **string** | Helm overrides for this cluster. Applicable only for a k8s cloud provider. Refer https://github.com/yugabyte/charts/blob/master/stable/yugabyte/values.yaml for the list of supported overrides. | [optional] 
**AzHelmOverrides** | Pointer to **map[string]string** | Helm overrides per availability zone of this cluster. Applicable only if this is a k8s cloud provider. Refer https://github.com/yugabyte/charts/blob/master/stable/yugabyte/values.yaml for the list of supported overrides. | [optional] 

## Methods

### NewClusterProviderSpec

`func NewClusterProviderSpec(provider string, ) *ClusterProviderSpec`

NewClusterProviderSpec instantiates a new ClusterProviderSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterProviderSpecWithDefaults

`func NewClusterProviderSpecWithDefaults() *ClusterProviderSpec`

NewClusterProviderSpecWithDefaults instantiates a new ClusterProviderSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *ClusterProviderSpec) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ClusterProviderSpec) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ClusterProviderSpec) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetRegionList

`func (o *ClusterProviderSpec) GetRegionList() []string`

GetRegionList returns the RegionList field if non-nil, zero value otherwise.

### GetRegionListOk

`func (o *ClusterProviderSpec) GetRegionListOk() (*[]string, bool)`

GetRegionListOk returns a tuple with the RegionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionList

`func (o *ClusterProviderSpec) SetRegionList(v []string)`

SetRegionList sets RegionList field to given value.

### HasRegionList

`func (o *ClusterProviderSpec) HasRegionList() bool`

HasRegionList returns a boolean if a field has been set.

### GetPreferredRegion

`func (o *ClusterProviderSpec) GetPreferredRegion() string`

GetPreferredRegion returns the PreferredRegion field if non-nil, zero value otherwise.

### GetPreferredRegionOk

`func (o *ClusterProviderSpec) GetPreferredRegionOk() (*string, bool)`

GetPreferredRegionOk returns a tuple with the PreferredRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredRegion

`func (o *ClusterProviderSpec) SetPreferredRegion(v string)`

SetPreferredRegion sets PreferredRegion field to given value.

### HasPreferredRegion

`func (o *ClusterProviderSpec) HasPreferredRegion() bool`

HasPreferredRegion returns a boolean if a field has been set.

### GetAccessKeyCode

`func (o *ClusterProviderSpec) GetAccessKeyCode() string`

GetAccessKeyCode returns the AccessKeyCode field if non-nil, zero value otherwise.

### GetAccessKeyCodeOk

`func (o *ClusterProviderSpec) GetAccessKeyCodeOk() (*string, bool)`

GetAccessKeyCodeOk returns a tuple with the AccessKeyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyCode

`func (o *ClusterProviderSpec) SetAccessKeyCode(v string)`

SetAccessKeyCode sets AccessKeyCode field to given value.

### HasAccessKeyCode

`func (o *ClusterProviderSpec) HasAccessKeyCode() bool`

HasAccessKeyCode returns a boolean if a field has been set.

### GetAwsInstanceProfile

`func (o *ClusterProviderSpec) GetAwsInstanceProfile() string`

GetAwsInstanceProfile returns the AwsInstanceProfile field if non-nil, zero value otherwise.

### GetAwsInstanceProfileOk

`func (o *ClusterProviderSpec) GetAwsInstanceProfileOk() (*string, bool)`

GetAwsInstanceProfileOk returns a tuple with the AwsInstanceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsInstanceProfile

`func (o *ClusterProviderSpec) SetAwsInstanceProfile(v string)`

SetAwsInstanceProfile sets AwsInstanceProfile field to given value.

### HasAwsInstanceProfile

`func (o *ClusterProviderSpec) HasAwsInstanceProfile() bool`

HasAwsInstanceProfile returns a boolean if a field has been set.

### GetImageBundleUuid

`func (o *ClusterProviderSpec) GetImageBundleUuid() string`

GetImageBundleUuid returns the ImageBundleUuid field if non-nil, zero value otherwise.

### GetImageBundleUuidOk

`func (o *ClusterProviderSpec) GetImageBundleUuidOk() (*string, bool)`

GetImageBundleUuidOk returns a tuple with the ImageBundleUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBundleUuid

`func (o *ClusterProviderSpec) SetImageBundleUuid(v string)`

SetImageBundleUuid sets ImageBundleUuid field to given value.

### HasImageBundleUuid

`func (o *ClusterProviderSpec) HasImageBundleUuid() bool`

HasImageBundleUuid returns a boolean if a field has been set.

### GetHelmOverrides

`func (o *ClusterProviderSpec) GetHelmOverrides() string`

GetHelmOverrides returns the HelmOverrides field if non-nil, zero value otherwise.

### GetHelmOverridesOk

`func (o *ClusterProviderSpec) GetHelmOverridesOk() (*string, bool)`

GetHelmOverridesOk returns a tuple with the HelmOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmOverrides

`func (o *ClusterProviderSpec) SetHelmOverrides(v string)`

SetHelmOverrides sets HelmOverrides field to given value.

### HasHelmOverrides

`func (o *ClusterProviderSpec) HasHelmOverrides() bool`

HasHelmOverrides returns a boolean if a field has been set.

### GetAzHelmOverrides

`func (o *ClusterProviderSpec) GetAzHelmOverrides() map[string]string`

GetAzHelmOverrides returns the AzHelmOverrides field if non-nil, zero value otherwise.

### GetAzHelmOverridesOk

`func (o *ClusterProviderSpec) GetAzHelmOverridesOk() (*map[string]string, bool)`

GetAzHelmOverridesOk returns a tuple with the AzHelmOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzHelmOverrides

`func (o *ClusterProviderSpec) SetAzHelmOverrides(v map[string]string)`

SetAzHelmOverrides sets AzHelmOverrides field to given value.

### HasAzHelmOverrides

`func (o *ClusterProviderSpec) HasAzHelmOverrides() bool`

HasAzHelmOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


