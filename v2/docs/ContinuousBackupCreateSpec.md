# ContinuousBackupCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageConfigUuid** | **string** | UUID of the storage config to use | 
**Frequency** | **int64** | Interval between two backups. | 
**FrequencyTimeUnit** | [**TimeUnitType**](TimeUnitType.md) |  | 
**NumBackups** | Pointer to **int32** | The number of historical backups to retain in the storage bucket. | [optional] [default to 5]
**BackupDir** | **string** | The name of the directory in the storage config to use for YBA backups. | 

## Methods

### NewContinuousBackupCreateSpec

`func NewContinuousBackupCreateSpec(storageConfigUuid string, frequency int64, frequencyTimeUnit TimeUnitType, backupDir string, ) *ContinuousBackupCreateSpec`

NewContinuousBackupCreateSpec instantiates a new ContinuousBackupCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContinuousBackupCreateSpecWithDefaults

`func NewContinuousBackupCreateSpecWithDefaults() *ContinuousBackupCreateSpec`

NewContinuousBackupCreateSpecWithDefaults instantiates a new ContinuousBackupCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageConfigUuid

`func (o *ContinuousBackupCreateSpec) GetStorageConfigUuid() string`

GetStorageConfigUuid returns the StorageConfigUuid field if non-nil, zero value otherwise.

### GetStorageConfigUuidOk

`func (o *ContinuousBackupCreateSpec) GetStorageConfigUuidOk() (*string, bool)`

GetStorageConfigUuidOk returns a tuple with the StorageConfigUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUuid

`func (o *ContinuousBackupCreateSpec) SetStorageConfigUuid(v string)`

SetStorageConfigUuid sets StorageConfigUuid field to given value.


### GetFrequency

`func (o *ContinuousBackupCreateSpec) GetFrequency() int64`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *ContinuousBackupCreateSpec) GetFrequencyOk() (*int64, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *ContinuousBackupCreateSpec) SetFrequency(v int64)`

SetFrequency sets Frequency field to given value.


### GetFrequencyTimeUnit

`func (o *ContinuousBackupCreateSpec) GetFrequencyTimeUnit() TimeUnitType`

GetFrequencyTimeUnit returns the FrequencyTimeUnit field if non-nil, zero value otherwise.

### GetFrequencyTimeUnitOk

`func (o *ContinuousBackupCreateSpec) GetFrequencyTimeUnitOk() (*TimeUnitType, bool)`

GetFrequencyTimeUnitOk returns a tuple with the FrequencyTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyTimeUnit

`func (o *ContinuousBackupCreateSpec) SetFrequencyTimeUnit(v TimeUnitType)`

SetFrequencyTimeUnit sets FrequencyTimeUnit field to given value.


### GetNumBackups

`func (o *ContinuousBackupCreateSpec) GetNumBackups() int32`

GetNumBackups returns the NumBackups field if non-nil, zero value otherwise.

### GetNumBackupsOk

`func (o *ContinuousBackupCreateSpec) GetNumBackupsOk() (*int32, bool)`

GetNumBackupsOk returns a tuple with the NumBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumBackups

`func (o *ContinuousBackupCreateSpec) SetNumBackups(v int32)`

SetNumBackups sets NumBackups field to given value.

### HasNumBackups

`func (o *ContinuousBackupCreateSpec) HasNumBackups() bool`

HasNumBackups returns a boolean if a field has been set.

### GetBackupDir

`func (o *ContinuousBackupCreateSpec) GetBackupDir() string`

GetBackupDir returns the BackupDir field if non-nil, zero value otherwise.

### GetBackupDirOk

`func (o *ContinuousBackupCreateSpec) GetBackupDirOk() (*string, bool)`

GetBackupDirOk returns a tuple with the BackupDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDir

`func (o *ContinuousBackupCreateSpec) SetBackupDir(v string)`

SetBackupDir sets BackupDir field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


