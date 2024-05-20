# ClusterProviderSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | Cloud provider UUID | 
**RegionList** | Pointer to **[]string** | The list of regions in the cloud provider to place data replicas | [optional] 
**PreferredRegion** | Pointer to **string** | The region to nominate as the preferred region in a geo-partitioned multi-region cluster | [optional] 
**AccessKeyCode** | **string** | One of the SSH access keys defined in Cloud Provider to be configured on nodes VMs | 
**ImageBundleUuid** | Pointer to **string** | Image bundle UUID to use for node VM image. Refers to one of the image bundles defined in the cloud provider. | [optional] 

## Methods

### NewClusterProviderSpec

`func NewClusterProviderSpec(provider string, accessKeyCode string, ) *ClusterProviderSpec`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


