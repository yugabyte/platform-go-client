# DeviceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudVolumeEncryption** | Pointer to [**CloudVolumeEncryption**](CloudVolumeEncryption.md) |  | [optional] 
**DiskIops** | Pointer to **int32** | Desired IOPS for the volumes mounted on this instance | [optional] 
**MountPoints** | Pointer to **string** | Comma-separated list of mount points for the devices in each instance | [optional] 
**NumVolumes** | Pointer to **int32** | Number of volumes to be mounted on this instance at the default path | [optional] 
**StorageClass** | Pointer to **string** | Name of the storage class | [optional] 
**StorageType** | Pointer to **string** | Storage type used for this instance | [optional] 
**Throughput** | Pointer to **int32** | Desired throughput for the volumes mounted on this instance | [optional] 
**VolumeSize** | Pointer to **int32** | The size of each volume in each instance. Could be modified in payload for /resize_node API call | [optional] 

## Methods

### NewDeviceInfo

`func NewDeviceInfo() *DeviceInfo`

NewDeviceInfo instantiates a new DeviceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceInfoWithDefaults

`func NewDeviceInfoWithDefaults() *DeviceInfo`

NewDeviceInfoWithDefaults instantiates a new DeviceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudVolumeEncryption

`func (o *DeviceInfo) GetCloudVolumeEncryption() CloudVolumeEncryption`

GetCloudVolumeEncryption returns the CloudVolumeEncryption field if non-nil, zero value otherwise.

### GetCloudVolumeEncryptionOk

`func (o *DeviceInfo) GetCloudVolumeEncryptionOk() (*CloudVolumeEncryption, bool)`

GetCloudVolumeEncryptionOk returns a tuple with the CloudVolumeEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudVolumeEncryption

`func (o *DeviceInfo) SetCloudVolumeEncryption(v CloudVolumeEncryption)`

SetCloudVolumeEncryption sets CloudVolumeEncryption field to given value.

### HasCloudVolumeEncryption

`func (o *DeviceInfo) HasCloudVolumeEncryption() bool`

HasCloudVolumeEncryption returns a boolean if a field has been set.

### GetDiskIops

`func (o *DeviceInfo) GetDiskIops() int32`

GetDiskIops returns the DiskIops field if non-nil, zero value otherwise.

### GetDiskIopsOk

`func (o *DeviceInfo) GetDiskIopsOk() (*int32, bool)`

GetDiskIopsOk returns a tuple with the DiskIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskIops

`func (o *DeviceInfo) SetDiskIops(v int32)`

SetDiskIops sets DiskIops field to given value.

### HasDiskIops

`func (o *DeviceInfo) HasDiskIops() bool`

HasDiskIops returns a boolean if a field has been set.

### GetMountPoints

`func (o *DeviceInfo) GetMountPoints() string`

GetMountPoints returns the MountPoints field if non-nil, zero value otherwise.

### GetMountPointsOk

`func (o *DeviceInfo) GetMountPointsOk() (*string, bool)`

GetMountPointsOk returns a tuple with the MountPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoints

`func (o *DeviceInfo) SetMountPoints(v string)`

SetMountPoints sets MountPoints field to given value.

### HasMountPoints

`func (o *DeviceInfo) HasMountPoints() bool`

HasMountPoints returns a boolean if a field has been set.

### GetNumVolumes

`func (o *DeviceInfo) GetNumVolumes() int32`

GetNumVolumes returns the NumVolumes field if non-nil, zero value otherwise.

### GetNumVolumesOk

`func (o *DeviceInfo) GetNumVolumesOk() (*int32, bool)`

GetNumVolumesOk returns a tuple with the NumVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumVolumes

`func (o *DeviceInfo) SetNumVolumes(v int32)`

SetNumVolumes sets NumVolumes field to given value.

### HasNumVolumes

`func (o *DeviceInfo) HasNumVolumes() bool`

HasNumVolumes returns a boolean if a field has been set.

### GetStorageClass

`func (o *DeviceInfo) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *DeviceInfo) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *DeviceInfo) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *DeviceInfo) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### GetStorageType

`func (o *DeviceInfo) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *DeviceInfo) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *DeviceInfo) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *DeviceInfo) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetThroughput

`func (o *DeviceInfo) GetThroughput() int32`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *DeviceInfo) GetThroughputOk() (*int32, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *DeviceInfo) SetThroughput(v int32)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *DeviceInfo) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.

### GetVolumeSize

`func (o *DeviceInfo) GetVolumeSize() int32`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *DeviceInfo) GetVolumeSizeOk() (*int32, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *DeviceInfo) SetVolumeSize(v int32)`

SetVolumeSize sets VolumeSize field to given value.

### HasVolumeSize

`func (o *DeviceInfo) HasVolumeSize() bool`

HasVolumeSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


