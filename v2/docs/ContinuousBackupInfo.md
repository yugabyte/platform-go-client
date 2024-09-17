# ContinuousBackupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | UUID of the Continuous Backup Config | [optional] [readonly] 
**Frequency** | Pointer to **int64** | Interval between two backups. | [optional] 
**FrequencyTimeUnit** | Pointer to [**TimeUnitType**](TimeUnitType.md) |  | [optional] 
**StorageLocation** | Pointer to **string** | bucket or directory where backups are stored | [optional] [readonly] 
**LastBackup** | Pointer to **time.Time** | time of last backup stored | [optional] [readonly] 

## Methods

### NewContinuousBackupInfo

`func NewContinuousBackupInfo() *ContinuousBackupInfo`

NewContinuousBackupInfo instantiates a new ContinuousBackupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContinuousBackupInfoWithDefaults

`func NewContinuousBackupInfoWithDefaults() *ContinuousBackupInfo`

NewContinuousBackupInfoWithDefaults instantiates a new ContinuousBackupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ContinuousBackupInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ContinuousBackupInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ContinuousBackupInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ContinuousBackupInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetFrequency

`func (o *ContinuousBackupInfo) GetFrequency() int64`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *ContinuousBackupInfo) GetFrequencyOk() (*int64, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *ContinuousBackupInfo) SetFrequency(v int64)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *ContinuousBackupInfo) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetFrequencyTimeUnit

`func (o *ContinuousBackupInfo) GetFrequencyTimeUnit() TimeUnitType`

GetFrequencyTimeUnit returns the FrequencyTimeUnit field if non-nil, zero value otherwise.

### GetFrequencyTimeUnitOk

`func (o *ContinuousBackupInfo) GetFrequencyTimeUnitOk() (*TimeUnitType, bool)`

GetFrequencyTimeUnitOk returns a tuple with the FrequencyTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyTimeUnit

`func (o *ContinuousBackupInfo) SetFrequencyTimeUnit(v TimeUnitType)`

SetFrequencyTimeUnit sets FrequencyTimeUnit field to given value.

### HasFrequencyTimeUnit

`func (o *ContinuousBackupInfo) HasFrequencyTimeUnit() bool`

HasFrequencyTimeUnit returns a boolean if a field has been set.

### GetStorageLocation

`func (o *ContinuousBackupInfo) GetStorageLocation() string`

GetStorageLocation returns the StorageLocation field if non-nil, zero value otherwise.

### GetStorageLocationOk

`func (o *ContinuousBackupInfo) GetStorageLocationOk() (*string, bool)`

GetStorageLocationOk returns a tuple with the StorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLocation

`func (o *ContinuousBackupInfo) SetStorageLocation(v string)`

SetStorageLocation sets StorageLocation field to given value.

### HasStorageLocation

`func (o *ContinuousBackupInfo) HasStorageLocation() bool`

HasStorageLocation returns a boolean if a field has been set.

### GetLastBackup

`func (o *ContinuousBackupInfo) GetLastBackup() time.Time`

GetLastBackup returns the LastBackup field if non-nil, zero value otherwise.

### GetLastBackupOk

`func (o *ContinuousBackupInfo) GetLastBackupOk() (*time.Time, bool)`

GetLastBackupOk returns a tuple with the LastBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackup

`func (o *ContinuousBackupInfo) SetLastBackup(v time.Time)`

SetLastBackup sets LastBackup field to given value.

### HasLastBackup

`func (o *ContinuousBackupInfo) HasLastBackup() bool`

HasLastBackup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


