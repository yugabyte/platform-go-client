# ClusterPerProviderSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | Cloud provider UUID | 
**RegionList** | Pointer to **[]string** | Deprecated: specify the placement explicitly. The list of regions in the cloud provider to place data replicas | [optional] 
**PreferredRegion** | Pointer to **string** | Deprecated: use default partition. The region to nominate as the preferred region in a geo-partitioned multi-region cluster | [optional] 
**AccessKeyCode** | Pointer to **string** | One of the SSH access keys defined in Cloud Provider to be configured on nodes VMs. Required for AWS, Azure and GCP Cloud Providers. | [optional] 
**AwsInstanceProfile** | Pointer to **string** | The AWS IAM instance profile ARN to use for the nodes in this cluster. Applicable only for nodes on AWS Cloud Provider. If specified, YugabyteDB Anywhere will use this instance profile instead of the access key. | [optional] 
**ImageBundleUuid** | Pointer to **string** | Image bundle UUID to use for node VM image. Refers to one of the image bundles defined in the cloud provider. | [optional] 
**HelmOverrides** | Pointer to **string** | Helm overrides for this cluster. Applicable only for a k8s cloud provider. Refer https://github.com/yugabyte/charts/blob/master/stable/yugabyte/values.yaml for the list of supported overrides. | [optional] 
**AzHelmOverrides** | Pointer to **map[string]string** | Helm overrides per availability zone of this cluster. Applicable only if this is a k8s cloud provider. Refer https://github.com/yugabyte/charts/blob/master/stable/yugabyte/values.yaml for the list of supported overrides. | [optional] 
**NodesSpecs** | Pointer to [**ProviderRootNodesSpec**](ProviderRootNodesSpec.md) |  | [optional] 
**NetworkingSpec** | Pointer to [**ClusterNetworkingSpec**](ClusterNetworkingSpec.md) |  | [optional] 
**InstanceTags** | Pointer to **map[string]string** | A map of strings representing a set of Tags and Values to apply on nodes in the aws/gcp/azu cloud. See https://docs.yugabyte.com/preview/yugabyte-platform/manage-deployments/instance-tags/. | [optional] 

## Methods

### NewClusterPerProviderSpec

`func NewClusterPerProviderSpec(provider string, ) *ClusterPerProviderSpec`

NewClusterPerProviderSpec instantiates a new ClusterPerProviderSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPerProviderSpecWithDefaults

`func NewClusterPerProviderSpecWithDefaults() *ClusterPerProviderSpec`

NewClusterPerProviderSpecWithDefaults instantiates a new ClusterPerProviderSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *ClusterPerProviderSpec) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ClusterPerProviderSpec) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ClusterPerProviderSpec) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetRegionList

`func (o *ClusterPerProviderSpec) GetRegionList() []string`

GetRegionList returns the RegionList field if non-nil, zero value otherwise.

### GetRegionListOk

`func (o *ClusterPerProviderSpec) GetRegionListOk() (*[]string, bool)`

GetRegionListOk returns a tuple with the RegionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionList

`func (o *ClusterPerProviderSpec) SetRegionList(v []string)`

SetRegionList sets RegionList field to given value.

### HasRegionList

`func (o *ClusterPerProviderSpec) HasRegionList() bool`

HasRegionList returns a boolean if a field has been set.

### GetPreferredRegion

`func (o *ClusterPerProviderSpec) GetPreferredRegion() string`

GetPreferredRegion returns the PreferredRegion field if non-nil, zero value otherwise.

### GetPreferredRegionOk

`func (o *ClusterPerProviderSpec) GetPreferredRegionOk() (*string, bool)`

GetPreferredRegionOk returns a tuple with the PreferredRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredRegion

`func (o *ClusterPerProviderSpec) SetPreferredRegion(v string)`

SetPreferredRegion sets PreferredRegion field to given value.

### HasPreferredRegion

`func (o *ClusterPerProviderSpec) HasPreferredRegion() bool`

HasPreferredRegion returns a boolean if a field has been set.

### GetAccessKeyCode

`func (o *ClusterPerProviderSpec) GetAccessKeyCode() string`

GetAccessKeyCode returns the AccessKeyCode field if non-nil, zero value otherwise.

### GetAccessKeyCodeOk

`func (o *ClusterPerProviderSpec) GetAccessKeyCodeOk() (*string, bool)`

GetAccessKeyCodeOk returns a tuple with the AccessKeyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyCode

`func (o *ClusterPerProviderSpec) SetAccessKeyCode(v string)`

SetAccessKeyCode sets AccessKeyCode field to given value.

### HasAccessKeyCode

`func (o *ClusterPerProviderSpec) HasAccessKeyCode() bool`

HasAccessKeyCode returns a boolean if a field has been set.

### GetAwsInstanceProfile

`func (o *ClusterPerProviderSpec) GetAwsInstanceProfile() string`

GetAwsInstanceProfile returns the AwsInstanceProfile field if non-nil, zero value otherwise.

### GetAwsInstanceProfileOk

`func (o *ClusterPerProviderSpec) GetAwsInstanceProfileOk() (*string, bool)`

GetAwsInstanceProfileOk returns a tuple with the AwsInstanceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsInstanceProfile

`func (o *ClusterPerProviderSpec) SetAwsInstanceProfile(v string)`

SetAwsInstanceProfile sets AwsInstanceProfile field to given value.

### HasAwsInstanceProfile

`func (o *ClusterPerProviderSpec) HasAwsInstanceProfile() bool`

HasAwsInstanceProfile returns a boolean if a field has been set.

### GetImageBundleUuid

`func (o *ClusterPerProviderSpec) GetImageBundleUuid() string`

GetImageBundleUuid returns the ImageBundleUuid field if non-nil, zero value otherwise.

### GetImageBundleUuidOk

`func (o *ClusterPerProviderSpec) GetImageBundleUuidOk() (*string, bool)`

GetImageBundleUuidOk returns a tuple with the ImageBundleUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBundleUuid

`func (o *ClusterPerProviderSpec) SetImageBundleUuid(v string)`

SetImageBundleUuid sets ImageBundleUuid field to given value.

### HasImageBundleUuid

`func (o *ClusterPerProviderSpec) HasImageBundleUuid() bool`

HasImageBundleUuid returns a boolean if a field has been set.

### GetHelmOverrides

`func (o *ClusterPerProviderSpec) GetHelmOverrides() string`

GetHelmOverrides returns the HelmOverrides field if non-nil, zero value otherwise.

### GetHelmOverridesOk

`func (o *ClusterPerProviderSpec) GetHelmOverridesOk() (*string, bool)`

GetHelmOverridesOk returns a tuple with the HelmOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmOverrides

`func (o *ClusterPerProviderSpec) SetHelmOverrides(v string)`

SetHelmOverrides sets HelmOverrides field to given value.

### HasHelmOverrides

`func (o *ClusterPerProviderSpec) HasHelmOverrides() bool`

HasHelmOverrides returns a boolean if a field has been set.

### GetAzHelmOverrides

`func (o *ClusterPerProviderSpec) GetAzHelmOverrides() map[string]string`

GetAzHelmOverrides returns the AzHelmOverrides field if non-nil, zero value otherwise.

### GetAzHelmOverridesOk

`func (o *ClusterPerProviderSpec) GetAzHelmOverridesOk() (*map[string]string, bool)`

GetAzHelmOverridesOk returns a tuple with the AzHelmOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzHelmOverrides

`func (o *ClusterPerProviderSpec) SetAzHelmOverrides(v map[string]string)`

SetAzHelmOverrides sets AzHelmOverrides field to given value.

### HasAzHelmOverrides

`func (o *ClusterPerProviderSpec) HasAzHelmOverrides() bool`

HasAzHelmOverrides returns a boolean if a field has been set.

### GetNodesSpecs

`func (o *ClusterPerProviderSpec) GetNodesSpecs() ProviderRootNodesSpec`

GetNodesSpecs returns the NodesSpecs field if non-nil, zero value otherwise.

### GetNodesSpecsOk

`func (o *ClusterPerProviderSpec) GetNodesSpecsOk() (*ProviderRootNodesSpec, bool)`

GetNodesSpecsOk returns a tuple with the NodesSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesSpecs

`func (o *ClusterPerProviderSpec) SetNodesSpecs(v ProviderRootNodesSpec)`

SetNodesSpecs sets NodesSpecs field to given value.

### HasNodesSpecs

`func (o *ClusterPerProviderSpec) HasNodesSpecs() bool`

HasNodesSpecs returns a boolean if a field has been set.

### GetNetworkingSpec

`func (o *ClusterPerProviderSpec) GetNetworkingSpec() ClusterNetworkingSpec`

GetNetworkingSpec returns the NetworkingSpec field if non-nil, zero value otherwise.

### GetNetworkingSpecOk

`func (o *ClusterPerProviderSpec) GetNetworkingSpecOk() (*ClusterNetworkingSpec, bool)`

GetNetworkingSpecOk returns a tuple with the NetworkingSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkingSpec

`func (o *ClusterPerProviderSpec) SetNetworkingSpec(v ClusterNetworkingSpec)`

SetNetworkingSpec sets NetworkingSpec field to given value.

### HasNetworkingSpec

`func (o *ClusterPerProviderSpec) HasNetworkingSpec() bool`

HasNetworkingSpec returns a boolean if a field has been set.

### GetInstanceTags

`func (o *ClusterPerProviderSpec) GetInstanceTags() map[string]string`

GetInstanceTags returns the InstanceTags field if non-nil, zero value otherwise.

### GetInstanceTagsOk

`func (o *ClusterPerProviderSpec) GetInstanceTagsOk() (*map[string]string, bool)`

GetInstanceTagsOk returns a tuple with the InstanceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTags

`func (o *ClusterPerProviderSpec) SetInstanceTags(v map[string]string)`

SetInstanceTags sets InstanceTags field to given value.

### HasInstanceTags

`func (o *ClusterPerProviderSpec) HasInstanceTags() bool`

HasInstanceTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


