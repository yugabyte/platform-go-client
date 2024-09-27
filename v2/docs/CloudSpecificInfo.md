# CloudSpecificInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignPublicIp** | Pointer to **bool** | True if the node has a public IP address assigned | [optional] 
**Az** | Pointer to **string** | The node&#39;s availability zone | [optional] 
**Cloud** | Pointer to **string** | The node&#39;s Cloud Provider | [optional] 
**InstanceType** | Pointer to **string** | The node&#39;s instance type | [optional] 
**KubernetesNamespace** | Pointer to **string** | Kubernetes namespace | [optional] 
**KubernetesPodName** | Pointer to **string** | Pod name in Kubernetes | [optional] 
**LunIndexes** | Pointer to **[]int32** | Mounted disks LUN indexes | [optional] 
**MountRoots** | Pointer to **string** | Mount roots | [optional] 
**PrivateDns** | Pointer to **string** | The node&#39;s private DNS | [optional] 
**PrivateIp** | Pointer to **string** | The node&#39;s private IP address | [optional] 
**PublicDns** | Pointer to **string** | The node&#39;s public DNS name | [optional] 
**PublicIp** | Pointer to **string** | The node&#39;s public IP address | [optional] 
**Region** | Pointer to **string** | The node&#39;s region | [optional] 
**RootVolume** | Pointer to **string** | Root volume ID or name | [optional] 
**SecondaryPrivateIp** | Pointer to **string** | Secondary Private IP | [optional] 
**SecondarySubnetId** | Pointer to **string** | Secondary Subnet IP | [optional] 
**SubnetId** | Pointer to **string** | ID of the subnet on which this node is deployed | [optional] 
**UseTimeSync** | Pointer to **bool** | True if &#x60;use time sync&#x60; is enabled | [optional] 

## Methods

### NewCloudSpecificInfo

`func NewCloudSpecificInfo() *CloudSpecificInfo`

NewCloudSpecificInfo instantiates a new CloudSpecificInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSpecificInfoWithDefaults

`func NewCloudSpecificInfoWithDefaults() *CloudSpecificInfo`

NewCloudSpecificInfoWithDefaults instantiates a new CloudSpecificInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignPublicIp

`func (o *CloudSpecificInfo) GetAssignPublicIp() bool`

GetAssignPublicIp returns the AssignPublicIp field if non-nil, zero value otherwise.

### GetAssignPublicIpOk

`func (o *CloudSpecificInfo) GetAssignPublicIpOk() (*bool, bool)`

GetAssignPublicIpOk returns a tuple with the AssignPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignPublicIp

`func (o *CloudSpecificInfo) SetAssignPublicIp(v bool)`

SetAssignPublicIp sets AssignPublicIp field to given value.

### HasAssignPublicIp

`func (o *CloudSpecificInfo) HasAssignPublicIp() bool`

HasAssignPublicIp returns a boolean if a field has been set.

### GetAz

`func (o *CloudSpecificInfo) GetAz() string`

GetAz returns the Az field if non-nil, zero value otherwise.

### GetAzOk

`func (o *CloudSpecificInfo) GetAzOk() (*string, bool)`

GetAzOk returns a tuple with the Az field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAz

`func (o *CloudSpecificInfo) SetAz(v string)`

SetAz sets Az field to given value.

### HasAz

`func (o *CloudSpecificInfo) HasAz() bool`

HasAz returns a boolean if a field has been set.

### GetCloud

`func (o *CloudSpecificInfo) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *CloudSpecificInfo) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *CloudSpecificInfo) SetCloud(v string)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *CloudSpecificInfo) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetInstanceType

`func (o *CloudSpecificInfo) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *CloudSpecificInfo) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *CloudSpecificInfo) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *CloudSpecificInfo) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetKubernetesNamespace

`func (o *CloudSpecificInfo) GetKubernetesNamespace() string`

GetKubernetesNamespace returns the KubernetesNamespace field if non-nil, zero value otherwise.

### GetKubernetesNamespaceOk

`func (o *CloudSpecificInfo) GetKubernetesNamespaceOk() (*string, bool)`

GetKubernetesNamespaceOk returns a tuple with the KubernetesNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesNamespace

`func (o *CloudSpecificInfo) SetKubernetesNamespace(v string)`

SetKubernetesNamespace sets KubernetesNamespace field to given value.

### HasKubernetesNamespace

`func (o *CloudSpecificInfo) HasKubernetesNamespace() bool`

HasKubernetesNamespace returns a boolean if a field has been set.

### GetKubernetesPodName

`func (o *CloudSpecificInfo) GetKubernetesPodName() string`

GetKubernetesPodName returns the KubernetesPodName field if non-nil, zero value otherwise.

### GetKubernetesPodNameOk

`func (o *CloudSpecificInfo) GetKubernetesPodNameOk() (*string, bool)`

GetKubernetesPodNameOk returns a tuple with the KubernetesPodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesPodName

`func (o *CloudSpecificInfo) SetKubernetesPodName(v string)`

SetKubernetesPodName sets KubernetesPodName field to given value.

### HasKubernetesPodName

`func (o *CloudSpecificInfo) HasKubernetesPodName() bool`

HasKubernetesPodName returns a boolean if a field has been set.

### GetLunIndexes

`func (o *CloudSpecificInfo) GetLunIndexes() []int32`

GetLunIndexes returns the LunIndexes field if non-nil, zero value otherwise.

### GetLunIndexesOk

`func (o *CloudSpecificInfo) GetLunIndexesOk() (*[]int32, bool)`

GetLunIndexesOk returns a tuple with the LunIndexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunIndexes

`func (o *CloudSpecificInfo) SetLunIndexes(v []int32)`

SetLunIndexes sets LunIndexes field to given value.

### HasLunIndexes

`func (o *CloudSpecificInfo) HasLunIndexes() bool`

HasLunIndexes returns a boolean if a field has been set.

### GetMountRoots

`func (o *CloudSpecificInfo) GetMountRoots() string`

GetMountRoots returns the MountRoots field if non-nil, zero value otherwise.

### GetMountRootsOk

`func (o *CloudSpecificInfo) GetMountRootsOk() (*string, bool)`

GetMountRootsOk returns a tuple with the MountRoots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountRoots

`func (o *CloudSpecificInfo) SetMountRoots(v string)`

SetMountRoots sets MountRoots field to given value.

### HasMountRoots

`func (o *CloudSpecificInfo) HasMountRoots() bool`

HasMountRoots returns a boolean if a field has been set.

### GetPrivateDns

`func (o *CloudSpecificInfo) GetPrivateDns() string`

GetPrivateDns returns the PrivateDns field if non-nil, zero value otherwise.

### GetPrivateDnsOk

`func (o *CloudSpecificInfo) GetPrivateDnsOk() (*string, bool)`

GetPrivateDnsOk returns a tuple with the PrivateDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDns

`func (o *CloudSpecificInfo) SetPrivateDns(v string)`

SetPrivateDns sets PrivateDns field to given value.

### HasPrivateDns

`func (o *CloudSpecificInfo) HasPrivateDns() bool`

HasPrivateDns returns a boolean if a field has been set.

### GetPrivateIp

`func (o *CloudSpecificInfo) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *CloudSpecificInfo) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *CloudSpecificInfo) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *CloudSpecificInfo) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetPublicDns

`func (o *CloudSpecificInfo) GetPublicDns() string`

GetPublicDns returns the PublicDns field if non-nil, zero value otherwise.

### GetPublicDnsOk

`func (o *CloudSpecificInfo) GetPublicDnsOk() (*string, bool)`

GetPublicDnsOk returns a tuple with the PublicDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicDns

`func (o *CloudSpecificInfo) SetPublicDns(v string)`

SetPublicDns sets PublicDns field to given value.

### HasPublicDns

`func (o *CloudSpecificInfo) HasPublicDns() bool`

HasPublicDns returns a boolean if a field has been set.

### GetPublicIp

`func (o *CloudSpecificInfo) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *CloudSpecificInfo) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *CloudSpecificInfo) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *CloudSpecificInfo) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRegion

`func (o *CloudSpecificInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudSpecificInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudSpecificInfo) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudSpecificInfo) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRootVolume

`func (o *CloudSpecificInfo) GetRootVolume() string`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *CloudSpecificInfo) GetRootVolumeOk() (*string, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *CloudSpecificInfo) SetRootVolume(v string)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *CloudSpecificInfo) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetSecondaryPrivateIp

`func (o *CloudSpecificInfo) GetSecondaryPrivateIp() string`

GetSecondaryPrivateIp returns the SecondaryPrivateIp field if non-nil, zero value otherwise.

### GetSecondaryPrivateIpOk

`func (o *CloudSpecificInfo) GetSecondaryPrivateIpOk() (*string, bool)`

GetSecondaryPrivateIpOk returns a tuple with the SecondaryPrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryPrivateIp

`func (o *CloudSpecificInfo) SetSecondaryPrivateIp(v string)`

SetSecondaryPrivateIp sets SecondaryPrivateIp field to given value.

### HasSecondaryPrivateIp

`func (o *CloudSpecificInfo) HasSecondaryPrivateIp() bool`

HasSecondaryPrivateIp returns a boolean if a field has been set.

### GetSecondarySubnetId

`func (o *CloudSpecificInfo) GetSecondarySubnetId() string`

GetSecondarySubnetId returns the SecondarySubnetId field if non-nil, zero value otherwise.

### GetSecondarySubnetIdOk

`func (o *CloudSpecificInfo) GetSecondarySubnetIdOk() (*string, bool)`

GetSecondarySubnetIdOk returns a tuple with the SecondarySubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySubnetId

`func (o *CloudSpecificInfo) SetSecondarySubnetId(v string)`

SetSecondarySubnetId sets SecondarySubnetId field to given value.

### HasSecondarySubnetId

`func (o *CloudSpecificInfo) HasSecondarySubnetId() bool`

HasSecondarySubnetId returns a boolean if a field has been set.

### GetSubnetId

`func (o *CloudSpecificInfo) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CloudSpecificInfo) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *CloudSpecificInfo) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *CloudSpecificInfo) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### GetUseTimeSync

`func (o *CloudSpecificInfo) GetUseTimeSync() bool`

GetUseTimeSync returns the UseTimeSync field if non-nil, zero value otherwise.

### GetUseTimeSyncOk

`func (o *CloudSpecificInfo) GetUseTimeSyncOk() (*bool, bool)`

GetUseTimeSyncOk returns a tuple with the UseTimeSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTimeSync

`func (o *CloudSpecificInfo) SetUseTimeSync(v bool)`

SetUseTimeSync sets UseTimeSync field to given value.

### HasUseTimeSync

`func (o *CloudSpecificInfo) HasUseTimeSync() bool`

HasUseTimeSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


