# ClusterStorageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeSize** | **int32** | The size of each volume in each instance. Could be modified in payload for /resize_node API call | 
**NumVolumes** | **int32** | Number of volumes to be mounted on this instance at the default path | 
**MountPoints** | Pointer to **string** | Comma-separated list of mount points for the volumes in each instance. Required for an onprem cluster. | [optional] 
**StorageClass** | Pointer to **string** | Name of the storage class, if this is a kubernetes cluster | [optional] 
**StorageType** | Pointer to **string** | Storage type used for this instance, if this is a aws (IO1, GP2, GP3), gcp (Scratch, Persistent) or azu (StandardSSD_LRS, Premium_LRS, PremiumV2_LRS, UltraSSD_LRS) cluster. | [optional] 
**DiskIops** | Pointer to **int32** | Desired IOPS for the volumes mounted on this aws, gcp or azu instance | [optional] 
**Throughput** | Pointer to **int32** | Desired throughput for the volumes mounted on this aws, gcp or azu instance | [optional] 

## Methods

### NewClusterStorageSpec

`func NewClusterStorageSpec(volumeSize int32, numVolumes int32, ) *ClusterStorageSpec`

NewClusterStorageSpec instantiates a new ClusterStorageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStorageSpecWithDefaults

`func NewClusterStorageSpecWithDefaults() *ClusterStorageSpec`

NewClusterStorageSpecWithDefaults instantiates a new ClusterStorageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeSize

`func (o *ClusterStorageSpec) GetVolumeSize() int32`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *ClusterStorageSpec) GetVolumeSizeOk() (*int32, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *ClusterStorageSpec) SetVolumeSize(v int32)`

SetVolumeSize sets VolumeSize field to given value.


### GetNumVolumes

`func (o *ClusterStorageSpec) GetNumVolumes() int32`

GetNumVolumes returns the NumVolumes field if non-nil, zero value otherwise.

### GetNumVolumesOk

`func (o *ClusterStorageSpec) GetNumVolumesOk() (*int32, bool)`

GetNumVolumesOk returns a tuple with the NumVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumVolumes

`func (o *ClusterStorageSpec) SetNumVolumes(v int32)`

SetNumVolumes sets NumVolumes field to given value.


### GetMountPoints

`func (o *ClusterStorageSpec) GetMountPoints() string`

GetMountPoints returns the MountPoints field if non-nil, zero value otherwise.

### GetMountPointsOk

`func (o *ClusterStorageSpec) GetMountPointsOk() (*string, bool)`

GetMountPointsOk returns a tuple with the MountPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoints

`func (o *ClusterStorageSpec) SetMountPoints(v string)`

SetMountPoints sets MountPoints field to given value.

### HasMountPoints

`func (o *ClusterStorageSpec) HasMountPoints() bool`

HasMountPoints returns a boolean if a field has been set.

### GetStorageClass

`func (o *ClusterStorageSpec) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *ClusterStorageSpec) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *ClusterStorageSpec) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *ClusterStorageSpec) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### GetStorageType

`func (o *ClusterStorageSpec) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *ClusterStorageSpec) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *ClusterStorageSpec) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *ClusterStorageSpec) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetDiskIops

`func (o *ClusterStorageSpec) GetDiskIops() int32`

GetDiskIops returns the DiskIops field if non-nil, zero value otherwise.

### GetDiskIopsOk

`func (o *ClusterStorageSpec) GetDiskIopsOk() (*int32, bool)`

GetDiskIopsOk returns a tuple with the DiskIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskIops

`func (o *ClusterStorageSpec) SetDiskIops(v int32)`

SetDiskIops sets DiskIops field to given value.

### HasDiskIops

`func (o *ClusterStorageSpec) HasDiskIops() bool`

HasDiskIops returns a boolean if a field has been set.

### GetThroughput

`func (o *ClusterStorageSpec) GetThroughput() int32`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *ClusterStorageSpec) GetThroughputOk() (*int32, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *ClusterStorageSpec) SetThroughput(v int32)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *ClusterStorageSpec) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


