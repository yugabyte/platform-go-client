# ParametersForValidatingRestorableKeyspaceAndTablesInBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupUUID** | **string** | UUID of the backup being restored | 
**KeyspaceTables** | Pointer to [**[]KeyspaceTables**](KeyspaceTables.md) | List of keyspace(s) and tables to be restored | [optional] 
**RestoreToPointInTimeMillis** | Pointer to **int64** | Point in restore timestamp in millis | [optional] 

## Methods

### NewParametersForValidatingRestorableKeyspaceAndTablesInBackup

`func NewParametersForValidatingRestorableKeyspaceAndTablesInBackup(backupUUID string, ) *ParametersForValidatingRestorableKeyspaceAndTablesInBackup`

NewParametersForValidatingRestorableKeyspaceAndTablesInBackup instantiates a new ParametersForValidatingRestorableKeyspaceAndTablesInBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParametersForValidatingRestorableKeyspaceAndTablesInBackupWithDefaults

`func NewParametersForValidatingRestorableKeyspaceAndTablesInBackupWithDefaults() *ParametersForValidatingRestorableKeyspaceAndTablesInBackup`

NewParametersForValidatingRestorableKeyspaceAndTablesInBackupWithDefaults instantiates a new ParametersForValidatingRestorableKeyspaceAndTablesInBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupUUID

`func (o *ParametersForValidatingRestorableKeyspaceAndTablesInBackup) GetBackupUUID() string`

GetBackupUUID returns the BackupUUID field if non-nil, zero value otherwise.

### GetBackupUUIDOk

`func (o *ParametersForValidatingRestorableKeyspaceAndTablesInBackup) GetBackupUUIDOk() (*string, bool)`

GetBackupUUIDOk returns a tuple with the BackupUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupUUID

`func (o *ParametersForValidatingRestorableKeyspaceAndTablesInBackup) SetBackupUUID(v string)`

SetBackupUUID sets BackupUUID field to given value.


### GetKeyspaceTables

`func (o *ParametersForValidatingRestorableKeyspaceAndTablesInBackup) GetKeyspaceTables() []KeyspaceTables`

GetKeyspaceTables returns the KeyspaceTables field if non-nil, zero value otherwise.

### GetKeyspaceTablesOk

`func (o *ParametersForValidatingRestorableKeyspaceAndTablesInBackup) GetKeyspaceTablesOk() (*[]KeyspaceTables, bool)`

GetKeyspaceTablesOk returns a tuple with the KeyspaceTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspaceTables

`func (o *ParametersForValidatingRestorableKeyspaceAndTablesInBackup) SetKeyspaceTables(v []KeyspaceTables)`

SetKeyspaceTables sets KeyspaceTables field to given value.

### HasKeyspaceTables

`func (o *ParametersForValidatingRestorableKeyspaceAndTablesInBackup) HasKeyspaceTables() bool`

HasKeyspaceTables returns a boolean if a field has been set.

### GetRestoreToPointInTimeMillis

`func (o *ParametersForValidatingRestorableKeyspaceAndTablesInBackup) GetRestoreToPointInTimeMillis() int64`

GetRestoreToPointInTimeMillis returns the RestoreToPointInTimeMillis field if non-nil, zero value otherwise.

### GetRestoreToPointInTimeMillisOk

`func (o *ParametersForValidatingRestorableKeyspaceAndTablesInBackup) GetRestoreToPointInTimeMillisOk() (*int64, bool)`

GetRestoreToPointInTimeMillisOk returns a tuple with the RestoreToPointInTimeMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToPointInTimeMillis

`func (o *ParametersForValidatingRestorableKeyspaceAndTablesInBackup) SetRestoreToPointInTimeMillis(v int64)`

SetRestoreToPointInTimeMillis sets RestoreToPointInTimeMillis field to given value.

### HasRestoreToPointInTimeMillis

`func (o *ParametersForValidatingRestorableKeyspaceAndTablesInBackup) HasRestoreToPointInTimeMillis() bool`

HasRestoreToPointInTimeMillis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


