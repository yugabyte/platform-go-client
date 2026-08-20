# ClusterProviderEditSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionList** | **[]string** | Edit the list of regions in the cloud provider to place data replicas | 
**AwsInstanceProfile** | Pointer to **string** | The AWS IAM instance profile ARN to use for the nodes in this cluster. Applicable only for nodes on AWS Cloud Provider. If specified, YugabyteDB Anywhere will use this instance profile instead of the access key. | [optional] 
**ImageBundleUuid** | Pointer to **string** | Image bundle UUID to use for node VM image. Refers to one of the image bundles defined in the cloud provider. | [optional] 

## Methods

### NewClusterProviderEditSpec

`func NewClusterProviderEditSpec(regionList []string, ) *ClusterProviderEditSpec`

NewClusterProviderEditSpec instantiates a new ClusterProviderEditSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterProviderEditSpecWithDefaults

`func NewClusterProviderEditSpecWithDefaults() *ClusterProviderEditSpec`

NewClusterProviderEditSpecWithDefaults instantiates a new ClusterProviderEditSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionList

`func (o *ClusterProviderEditSpec) GetRegionList() []string`

GetRegionList returns the RegionList field if non-nil, zero value otherwise.

### GetRegionListOk

`func (o *ClusterProviderEditSpec) GetRegionListOk() (*[]string, bool)`

GetRegionListOk returns a tuple with the RegionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionList

`func (o *ClusterProviderEditSpec) SetRegionList(v []string)`

SetRegionList sets RegionList field to given value.


### GetAwsInstanceProfile

`func (o *ClusterProviderEditSpec) GetAwsInstanceProfile() string`

GetAwsInstanceProfile returns the AwsInstanceProfile field if non-nil, zero value otherwise.

### GetAwsInstanceProfileOk

`func (o *ClusterProviderEditSpec) GetAwsInstanceProfileOk() (*string, bool)`

GetAwsInstanceProfileOk returns a tuple with the AwsInstanceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsInstanceProfile

`func (o *ClusterProviderEditSpec) SetAwsInstanceProfile(v string)`

SetAwsInstanceProfile sets AwsInstanceProfile field to given value.

### HasAwsInstanceProfile

`func (o *ClusterProviderEditSpec) HasAwsInstanceProfile() bool`

HasAwsInstanceProfile returns a boolean if a field has been set.

### GetImageBundleUuid

`func (o *ClusterProviderEditSpec) GetImageBundleUuid() string`

GetImageBundleUuid returns the ImageBundleUuid field if non-nil, zero value otherwise.

### GetImageBundleUuidOk

`func (o *ClusterProviderEditSpec) GetImageBundleUuidOk() (*string, bool)`

GetImageBundleUuidOk returns a tuple with the ImageBundleUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBundleUuid

`func (o *ClusterProviderEditSpec) SetImageBundleUuid(v string)`

SetImageBundleUuid sets ImageBundleUuid field to given value.

### HasImageBundleUuid

`func (o *ClusterProviderEditSpec) HasImageBundleUuid() bool`

HasImageBundleUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


