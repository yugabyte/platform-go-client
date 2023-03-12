# CloudInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to [**AWSCloudInfo**](AWSCloudInfo.md) |  | [optional] 
**Azu** | Pointer to [**AzureCloudInfo**](AzureCloudInfo.md) |  | [optional] 
**Gcp** | Pointer to [**GCPCloudInfo**](GCPCloudInfo.md) |  | [optional] 
**Kubernetes** | Pointer to [**KubernetesInfo**](KubernetesInfo.md) |  | [optional] 
**Onprem** | Pointer to [**OnPremCloudInfo**](OnPremCloudInfo.md) |  | [optional] 

## Methods

### NewCloudInfo

`func NewCloudInfo() *CloudInfo`

NewCloudInfo instantiates a new CloudInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInfoWithDefaults

`func NewCloudInfoWithDefaults() *CloudInfo`

NewCloudInfoWithDefaults instantiates a new CloudInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *CloudInfo) GetAws() AWSCloudInfo`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *CloudInfo) GetAwsOk() (*AWSCloudInfo, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *CloudInfo) SetAws(v AWSCloudInfo)`

SetAws sets Aws field to given value.

### HasAws

`func (o *CloudInfo) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetAzu

`func (o *CloudInfo) GetAzu() AzureCloudInfo`

GetAzu returns the Azu field if non-nil, zero value otherwise.

### GetAzuOk

`func (o *CloudInfo) GetAzuOk() (*AzureCloudInfo, bool)`

GetAzuOk returns a tuple with the Azu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzu

`func (o *CloudInfo) SetAzu(v AzureCloudInfo)`

SetAzu sets Azu field to given value.

### HasAzu

`func (o *CloudInfo) HasAzu() bool`

HasAzu returns a boolean if a field has been set.

### GetGcp

`func (o *CloudInfo) GetGcp() GCPCloudInfo`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *CloudInfo) GetGcpOk() (*GCPCloudInfo, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *CloudInfo) SetGcp(v GCPCloudInfo)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *CloudInfo) HasGcp() bool`

HasGcp returns a boolean if a field has been set.

### GetKubernetes

`func (o *CloudInfo) GetKubernetes() KubernetesInfo`

GetKubernetes returns the Kubernetes field if non-nil, zero value otherwise.

### GetKubernetesOk

`func (o *CloudInfo) GetKubernetesOk() (*KubernetesInfo, bool)`

GetKubernetesOk returns a tuple with the Kubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetes

`func (o *CloudInfo) SetKubernetes(v KubernetesInfo)`

SetKubernetes sets Kubernetes field to given value.

### HasKubernetes

`func (o *CloudInfo) HasKubernetes() bool`

HasKubernetes returns a boolean if a field has been set.

### GetOnprem

`func (o *CloudInfo) GetOnprem() OnPremCloudInfo`

GetOnprem returns the Onprem field if non-nil, zero value otherwise.

### GetOnpremOk

`func (o *CloudInfo) GetOnpremOk() (*OnPremCloudInfo, bool)`

GetOnpremOk returns a tuple with the Onprem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnprem

`func (o *CloudInfo) SetOnprem(v OnPremCloudInfo)`

SetOnprem sets Onprem field to given value.

### HasOnprem

`func (o *CloudInfo) HasOnprem() bool`

HasOnprem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


