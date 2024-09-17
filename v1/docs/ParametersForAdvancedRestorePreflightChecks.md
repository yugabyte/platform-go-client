# ParametersForAdvancedRestorePreflightChecks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupLocations** | **[]string** | List of backup locations to restore from | 
**RestoreToPointInTimeMillis** | Pointer to **int64** | Point in restore timestamp in millis | [optional] 
**StorageConfigUUID** | **string** | Storage config UUID | 
**UniverseUUID** | **string** | Target universe UUID | 

## Methods

### NewParametersForAdvancedRestorePreflightChecks

`func NewParametersForAdvancedRestorePreflightChecks(backupLocations []string, storageConfigUUID string, universeUUID string, ) *ParametersForAdvancedRestorePreflightChecks`

NewParametersForAdvancedRestorePreflightChecks instantiates a new ParametersForAdvancedRestorePreflightChecks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParametersForAdvancedRestorePreflightChecksWithDefaults

`func NewParametersForAdvancedRestorePreflightChecksWithDefaults() *ParametersForAdvancedRestorePreflightChecks`

NewParametersForAdvancedRestorePreflightChecksWithDefaults instantiates a new ParametersForAdvancedRestorePreflightChecks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupLocations

`func (o *ParametersForAdvancedRestorePreflightChecks) GetBackupLocations() []string`

GetBackupLocations returns the BackupLocations field if non-nil, zero value otherwise.

### GetBackupLocationsOk

`func (o *ParametersForAdvancedRestorePreflightChecks) GetBackupLocationsOk() (*[]string, bool)`

GetBackupLocationsOk returns a tuple with the BackupLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLocations

`func (o *ParametersForAdvancedRestorePreflightChecks) SetBackupLocations(v []string)`

SetBackupLocations sets BackupLocations field to given value.


### GetRestoreToPointInTimeMillis

`func (o *ParametersForAdvancedRestorePreflightChecks) GetRestoreToPointInTimeMillis() int64`

GetRestoreToPointInTimeMillis returns the RestoreToPointInTimeMillis field if non-nil, zero value otherwise.

### GetRestoreToPointInTimeMillisOk

`func (o *ParametersForAdvancedRestorePreflightChecks) GetRestoreToPointInTimeMillisOk() (*int64, bool)`

GetRestoreToPointInTimeMillisOk returns a tuple with the RestoreToPointInTimeMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToPointInTimeMillis

`func (o *ParametersForAdvancedRestorePreflightChecks) SetRestoreToPointInTimeMillis(v int64)`

SetRestoreToPointInTimeMillis sets RestoreToPointInTimeMillis field to given value.

### HasRestoreToPointInTimeMillis

`func (o *ParametersForAdvancedRestorePreflightChecks) HasRestoreToPointInTimeMillis() bool`

HasRestoreToPointInTimeMillis returns a boolean if a field has been set.

### GetStorageConfigUUID

`func (o *ParametersForAdvancedRestorePreflightChecks) GetStorageConfigUUID() string`

GetStorageConfigUUID returns the StorageConfigUUID field if non-nil, zero value otherwise.

### GetStorageConfigUUIDOk

`func (o *ParametersForAdvancedRestorePreflightChecks) GetStorageConfigUUIDOk() (*string, bool)`

GetStorageConfigUUIDOk returns a tuple with the StorageConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUUID

`func (o *ParametersForAdvancedRestorePreflightChecks) SetStorageConfigUUID(v string)`

SetStorageConfigUUID sets StorageConfigUUID field to given value.


### GetUniverseUUID

`func (o *ParametersForAdvancedRestorePreflightChecks) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *ParametersForAdvancedRestorePreflightChecks) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *ParametersForAdvancedRestorePreflightChecks) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


