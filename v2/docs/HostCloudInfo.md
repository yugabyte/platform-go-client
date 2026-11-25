# HostCloudInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloud** | Pointer to [**HostCloud**](HostCloud.md) |  | [optional] 
**Region** | Pointer to **string** | Cloud region of the YBA host. Absent when the cloud is unknown or its metadata service is unreachable. | [optional] [readonly] 
**Aws** | Pointer to [**AwsHostInfo**](AwsHostInfo.md) |  | [optional] 
**Gcp** | Pointer to [**GcpHostInfo**](GcpHostInfo.md) |  | [optional] 
**Azu** | Pointer to [**AzuHostInfo**](AzuHostInfo.md) |  | [optional] 

## Methods

### NewHostCloudInfo

`func NewHostCloudInfo() *HostCloudInfo`

NewHostCloudInfo instantiates a new HostCloudInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostCloudInfoWithDefaults

`func NewHostCloudInfoWithDefaults() *HostCloudInfo`

NewHostCloudInfoWithDefaults instantiates a new HostCloudInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloud

`func (o *HostCloudInfo) GetCloud() HostCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *HostCloudInfo) GetCloudOk() (*HostCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *HostCloudInfo) SetCloud(v HostCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *HostCloudInfo) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetRegion

`func (o *HostCloudInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HostCloudInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HostCloudInfo) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *HostCloudInfo) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetAws

`func (o *HostCloudInfo) GetAws() AwsHostInfo`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *HostCloudInfo) GetAwsOk() (*AwsHostInfo, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *HostCloudInfo) SetAws(v AwsHostInfo)`

SetAws sets Aws field to given value.

### HasAws

`func (o *HostCloudInfo) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetGcp

`func (o *HostCloudInfo) GetGcp() GcpHostInfo`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *HostCloudInfo) GetGcpOk() (*GcpHostInfo, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *HostCloudInfo) SetGcp(v GcpHostInfo)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *HostCloudInfo) HasGcp() bool`

HasGcp returns a boolean if a field has been set.

### GetAzu

`func (o *HostCloudInfo) GetAzu() AzuHostInfo`

GetAzu returns the Azu field if non-nil, zero value otherwise.

### GetAzuOk

`func (o *HostCloudInfo) GetAzuOk() (*AzuHostInfo, bool)`

GetAzuOk returns a tuple with the Azu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzu

`func (o *HostCloudInfo) SetAzu(v AzuHostInfo)`

SetAzu sets Azu field to given value.

### HasAzu

`func (o *HostCloudInfo) HasAzu() bool`

HasAzu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


