# ClusterProviderEditSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionList** | **[]string** | Edit the list of regions in the cloud provider to place data replicas | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


