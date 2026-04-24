# ClusterResizeStorageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeSize** | Pointer to **int32** | The size of each volume in each instance. Could be modified in payload for /resize_node API call | [optional] 
**DiskIops** | Pointer to **int32** | Desired IOPS for the volumes mounted on this aws, gcp or azu instance | [optional] 
**Throughput** | Pointer to **int32** | Desired throughput for the volumes mounted on this aws, gcp or azu instance | [optional] 

## Methods

### NewClusterResizeStorageSpec

`func NewClusterResizeStorageSpec() *ClusterResizeStorageSpec`

NewClusterResizeStorageSpec instantiates a new ClusterResizeStorageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterResizeStorageSpecWithDefaults

`func NewClusterResizeStorageSpecWithDefaults() *ClusterResizeStorageSpec`

NewClusterResizeStorageSpecWithDefaults instantiates a new ClusterResizeStorageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeSize

`func (o *ClusterResizeStorageSpec) GetVolumeSize() int32`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *ClusterResizeStorageSpec) GetVolumeSizeOk() (*int32, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *ClusterResizeStorageSpec) SetVolumeSize(v int32)`

SetVolumeSize sets VolumeSize field to given value.

### HasVolumeSize

`func (o *ClusterResizeStorageSpec) HasVolumeSize() bool`

HasVolumeSize returns a boolean if a field has been set.

### GetDiskIops

`func (o *ClusterResizeStorageSpec) GetDiskIops() int32`

GetDiskIops returns the DiskIops field if non-nil, zero value otherwise.

### GetDiskIopsOk

`func (o *ClusterResizeStorageSpec) GetDiskIopsOk() (*int32, bool)`

GetDiskIopsOk returns a tuple with the DiskIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskIops

`func (o *ClusterResizeStorageSpec) SetDiskIops(v int32)`

SetDiskIops sets DiskIops field to given value.

### HasDiskIops

`func (o *ClusterResizeStorageSpec) HasDiskIops() bool`

HasDiskIops returns a boolean if a field has been set.

### GetThroughput

`func (o *ClusterResizeStorageSpec) GetThroughput() int32`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *ClusterResizeStorageSpec) GetThroughputOk() (*int32, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *ClusterResizeStorageSpec) SetThroughput(v int32)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *ClusterResizeStorageSpec) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


