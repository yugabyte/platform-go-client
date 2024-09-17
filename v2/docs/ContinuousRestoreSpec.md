# ContinuousRestoreSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageConfigUuid** | **string** | UUID of the storage config to use | 
**BackupDir** | **string** | The name of the directory to restore the most recent back from. | 

## Methods

### NewContinuousRestoreSpec

`func NewContinuousRestoreSpec(storageConfigUuid string, backupDir string, ) *ContinuousRestoreSpec`

NewContinuousRestoreSpec instantiates a new ContinuousRestoreSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContinuousRestoreSpecWithDefaults

`func NewContinuousRestoreSpecWithDefaults() *ContinuousRestoreSpec`

NewContinuousRestoreSpecWithDefaults instantiates a new ContinuousRestoreSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageConfigUuid

`func (o *ContinuousRestoreSpec) GetStorageConfigUuid() string`

GetStorageConfigUuid returns the StorageConfigUuid field if non-nil, zero value otherwise.

### GetStorageConfigUuidOk

`func (o *ContinuousRestoreSpec) GetStorageConfigUuidOk() (*string, bool)`

GetStorageConfigUuidOk returns a tuple with the StorageConfigUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUuid

`func (o *ContinuousRestoreSpec) SetStorageConfigUuid(v string)`

SetStorageConfigUuid sets StorageConfigUuid field to given value.


### GetBackupDir

`func (o *ContinuousRestoreSpec) GetBackupDir() string`

GetBackupDir returns the BackupDir field if non-nil, zero value otherwise.

### GetBackupDirOk

`func (o *ContinuousRestoreSpec) GetBackupDirOk() (*string, bool)`

GetBackupDirOk returns a tuple with the BackupDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDir

`func (o *ContinuousRestoreSpec) SetBackupDir(v string)`

SetBackupDir sets BackupDir field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


