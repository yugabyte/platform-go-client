# RegionCloudInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to [**AWSRegionCloudInfo**](AWSRegionCloudInfo.md) |  | [optional] 
**Azu** | Pointer to [**AzureRegionCloudInfo**](AzureRegionCloudInfo.md) |  | [optional] 
**Gcp** | Pointer to [**GCPRegionCloudInfo**](GCPRegionCloudInfo.md) |  | [optional] 
**Kubernetes** | Pointer to [**KubernetesRegionInfo**](KubernetesRegionInfo.md) |  | [optional] 

## Methods

### NewRegionCloudInfo

`func NewRegionCloudInfo() *RegionCloudInfo`

NewRegionCloudInfo instantiates a new RegionCloudInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionCloudInfoWithDefaults

`func NewRegionCloudInfoWithDefaults() *RegionCloudInfo`

NewRegionCloudInfoWithDefaults instantiates a new RegionCloudInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *RegionCloudInfo) GetAws() AWSRegionCloudInfo`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *RegionCloudInfo) GetAwsOk() (*AWSRegionCloudInfo, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *RegionCloudInfo) SetAws(v AWSRegionCloudInfo)`

SetAws sets Aws field to given value.

### HasAws

`func (o *RegionCloudInfo) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetAzu

`func (o *RegionCloudInfo) GetAzu() AzureRegionCloudInfo`

GetAzu returns the Azu field if non-nil, zero value otherwise.

### GetAzuOk

`func (o *RegionCloudInfo) GetAzuOk() (*AzureRegionCloudInfo, bool)`

GetAzuOk returns a tuple with the Azu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzu

`func (o *RegionCloudInfo) SetAzu(v AzureRegionCloudInfo)`

SetAzu sets Azu field to given value.

### HasAzu

`func (o *RegionCloudInfo) HasAzu() bool`

HasAzu returns a boolean if a field has been set.

### GetGcp

`func (o *RegionCloudInfo) GetGcp() GCPRegionCloudInfo`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *RegionCloudInfo) GetGcpOk() (*GCPRegionCloudInfo, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *RegionCloudInfo) SetGcp(v GCPRegionCloudInfo)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *RegionCloudInfo) HasGcp() bool`

HasGcp returns a boolean if a field has been set.

### GetKubernetes

`func (o *RegionCloudInfo) GetKubernetes() KubernetesRegionInfo`

GetKubernetes returns the Kubernetes field if non-nil, zero value otherwise.

### GetKubernetesOk

`func (o *RegionCloudInfo) GetKubernetesOk() (*KubernetesRegionInfo, bool)`

GetKubernetesOk returns a tuple with the Kubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetes

`func (o *RegionCloudInfo) SetKubernetes(v KubernetesRegionInfo)`

SetKubernetes sets Kubernetes field to given value.

### HasKubernetes

`func (o *RegionCloudInfo) HasKubernetes() bool`

HasKubernetes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


