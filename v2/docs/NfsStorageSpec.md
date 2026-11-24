# NfsStorageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageType** | [**StorageType**](StorageType.md) |  | 
**BackupLocation** | **string** | Absolute path of the NFS mount on the database nodes. Can&#39;t be changed after creation. | 
**Bucket** | Pointer to **string** | Directory under backup_location that holds the backups. Defaults to yugabyte_backup. | [optional] 
**Volumes** | Pointer to **[]string** | NFS volumes to check, in the backup prechecks, for being mounted on the nodes | [optional] 
**RegionLocations** | Pointer to [**[]NfsRegionLocation**](NfsRegionLocation.md) | Per-region backup locations, for universes that store each region&#39;s backup in that region | [optional] 

## Methods

### NewNfsStorageSpec

`func NewNfsStorageSpec(storageType StorageType, backupLocation string, ) *NfsStorageSpec`

NewNfsStorageSpec instantiates a new NfsStorageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsStorageSpecWithDefaults

`func NewNfsStorageSpecWithDefaults() *NfsStorageSpec`

NewNfsStorageSpecWithDefaults instantiates a new NfsStorageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageType

`func (o *NfsStorageSpec) GetStorageType() StorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *NfsStorageSpec) GetStorageTypeOk() (*StorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *NfsStorageSpec) SetStorageType(v StorageType)`

SetStorageType sets StorageType field to given value.


### GetBackupLocation

`func (o *NfsStorageSpec) GetBackupLocation() string`

GetBackupLocation returns the BackupLocation field if non-nil, zero value otherwise.

### GetBackupLocationOk

`func (o *NfsStorageSpec) GetBackupLocationOk() (*string, bool)`

GetBackupLocationOk returns a tuple with the BackupLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLocation

`func (o *NfsStorageSpec) SetBackupLocation(v string)`

SetBackupLocation sets BackupLocation field to given value.


### GetBucket

`func (o *NfsStorageSpec) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *NfsStorageSpec) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *NfsStorageSpec) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *NfsStorageSpec) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetVolumes

`func (o *NfsStorageSpec) GetVolumes() []string`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *NfsStorageSpec) GetVolumesOk() (*[]string, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *NfsStorageSpec) SetVolumes(v []string)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *NfsStorageSpec) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetRegionLocations

`func (o *NfsStorageSpec) GetRegionLocations() []NfsRegionLocation`

GetRegionLocations returns the RegionLocations field if non-nil, zero value otherwise.

### GetRegionLocationsOk

`func (o *NfsStorageSpec) GetRegionLocationsOk() (*[]NfsRegionLocation, bool)`

GetRegionLocationsOk returns a tuple with the RegionLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionLocations

`func (o *NfsStorageSpec) SetRegionLocations(v []NfsRegionLocation)`

SetRegionLocations sets RegionLocations field to given value.

### HasRegionLocations

`func (o *NfsStorageSpec) HasRegionLocations() bool`

HasRegionLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


