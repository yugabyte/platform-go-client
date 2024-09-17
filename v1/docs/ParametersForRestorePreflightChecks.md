# ParametersForRestorePreflightChecks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupUUID** | **string** | UUID of the backup being restored | 
**KeyspaceTables** | Pointer to [**[]KeyspaceTables**](KeyspaceTables.md) | List of keyspace(s) and tables to be restored | [optional] 
**RestoreToPointInTimeMillis** | Pointer to **int64** | Point in restore timestamp in millis | [optional] 
**UniverseUUID** | **string** | Target universe UUID | 

## Methods

### NewParametersForRestorePreflightChecks

`func NewParametersForRestorePreflightChecks(backupUUID string, universeUUID string, ) *ParametersForRestorePreflightChecks`

NewParametersForRestorePreflightChecks instantiates a new ParametersForRestorePreflightChecks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParametersForRestorePreflightChecksWithDefaults

`func NewParametersForRestorePreflightChecksWithDefaults() *ParametersForRestorePreflightChecks`

NewParametersForRestorePreflightChecksWithDefaults instantiates a new ParametersForRestorePreflightChecks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupUUID

`func (o *ParametersForRestorePreflightChecks) GetBackupUUID() string`

GetBackupUUID returns the BackupUUID field if non-nil, zero value otherwise.

### GetBackupUUIDOk

`func (o *ParametersForRestorePreflightChecks) GetBackupUUIDOk() (*string, bool)`

GetBackupUUIDOk returns a tuple with the BackupUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupUUID

`func (o *ParametersForRestorePreflightChecks) SetBackupUUID(v string)`

SetBackupUUID sets BackupUUID field to given value.


### GetKeyspaceTables

`func (o *ParametersForRestorePreflightChecks) GetKeyspaceTables() []KeyspaceTables`

GetKeyspaceTables returns the KeyspaceTables field if non-nil, zero value otherwise.

### GetKeyspaceTablesOk

`func (o *ParametersForRestorePreflightChecks) GetKeyspaceTablesOk() (*[]KeyspaceTables, bool)`

GetKeyspaceTablesOk returns a tuple with the KeyspaceTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspaceTables

`func (o *ParametersForRestorePreflightChecks) SetKeyspaceTables(v []KeyspaceTables)`

SetKeyspaceTables sets KeyspaceTables field to given value.

### HasKeyspaceTables

`func (o *ParametersForRestorePreflightChecks) HasKeyspaceTables() bool`

HasKeyspaceTables returns a boolean if a field has been set.

### GetRestoreToPointInTimeMillis

`func (o *ParametersForRestorePreflightChecks) GetRestoreToPointInTimeMillis() int64`

GetRestoreToPointInTimeMillis returns the RestoreToPointInTimeMillis field if non-nil, zero value otherwise.

### GetRestoreToPointInTimeMillisOk

`func (o *ParametersForRestorePreflightChecks) GetRestoreToPointInTimeMillisOk() (*int64, bool)`

GetRestoreToPointInTimeMillisOk returns a tuple with the RestoreToPointInTimeMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToPointInTimeMillis

`func (o *ParametersForRestorePreflightChecks) SetRestoreToPointInTimeMillis(v int64)`

SetRestoreToPointInTimeMillis sets RestoreToPointInTimeMillis field to given value.

### HasRestoreToPointInTimeMillis

`func (o *ParametersForRestorePreflightChecks) HasRestoreToPointInTimeMillis() bool`

HasRestoreToPointInTimeMillis returns a boolean if a field has been set.

### GetUniverseUUID

`func (o *ParametersForRestorePreflightChecks) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *ParametersForRestorePreflightChecks) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *ParametersForRestorePreflightChecks) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


