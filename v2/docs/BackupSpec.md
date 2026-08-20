# BackupSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerUuid** | Pointer to **string** | Customer UUID that owns the backup. | [optional] 
**UniverseUuid** | Pointer to **string** | Universe UUID that was backed up. | [optional] 
**UniverseName** | Pointer to **string** | Universe name at backup time. | [optional] 
**ScheduleUuid** | Pointer to **string** | Backup schedule UUID, when created by a schedule. | [optional] 
**ScheduleName** | Pointer to **string** | Backup schedule name, when created by a schedule. | [optional] 
**BackupType** | Pointer to [**TableType**](TableType.md) |  | [optional] 
**Category** | Pointer to **string** | Backup implementation category. | [optional] 
**StorageConfigType** | Pointer to **string** | Storage provider type for the backup. | [optional] 
**IsFullBackup** | Pointer to **bool** | Whether this is a full backup rather than incremental. | [optional] 
**OnDemand** | Pointer to **bool** | True when the backup was not created by a schedule. | [optional] 
**UseTablespaces** | Pointer to **bool** | Whether tablespaces were included in the backup. | [optional] 
**UseRoles** | Pointer to **bool** | Whether database roles were included in the backup. | [optional] 
**ExpiryTime** | Pointer to **time.Time** | Time when the backup expires. | [optional] 
**ExpiryTimeUnit** | Pointer to **string** | Unit used for backup expiry configuration. | [optional] 
**KeyspaceTables** | Pointer to [**[]BackupKeyspaceTables**](BackupKeyspaceTables.md) | Keyspaces and tables included in the backup. | [optional] 

## Methods

### NewBackupSpec

`func NewBackupSpec() *BackupSpec`

NewBackupSpec instantiates a new BackupSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupSpecWithDefaults

`func NewBackupSpecWithDefaults() *BackupSpec`

NewBackupSpecWithDefaults instantiates a new BackupSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerUuid

`func (o *BackupSpec) GetCustomerUuid() string`

GetCustomerUuid returns the CustomerUuid field if non-nil, zero value otherwise.

### GetCustomerUuidOk

`func (o *BackupSpec) GetCustomerUuidOk() (*string, bool)`

GetCustomerUuidOk returns a tuple with the CustomerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUuid

`func (o *BackupSpec) SetCustomerUuid(v string)`

SetCustomerUuid sets CustomerUuid field to given value.

### HasCustomerUuid

`func (o *BackupSpec) HasCustomerUuid() bool`

HasCustomerUuid returns a boolean if a field has been set.

### GetUniverseUuid

`func (o *BackupSpec) GetUniverseUuid() string`

GetUniverseUuid returns the UniverseUuid field if non-nil, zero value otherwise.

### GetUniverseUuidOk

`func (o *BackupSpec) GetUniverseUuidOk() (*string, bool)`

GetUniverseUuidOk returns a tuple with the UniverseUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUuid

`func (o *BackupSpec) SetUniverseUuid(v string)`

SetUniverseUuid sets UniverseUuid field to given value.

### HasUniverseUuid

`func (o *BackupSpec) HasUniverseUuid() bool`

HasUniverseUuid returns a boolean if a field has been set.

### GetUniverseName

`func (o *BackupSpec) GetUniverseName() string`

GetUniverseName returns the UniverseName field if non-nil, zero value otherwise.

### GetUniverseNameOk

`func (o *BackupSpec) GetUniverseNameOk() (*string, bool)`

GetUniverseNameOk returns a tuple with the UniverseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseName

`func (o *BackupSpec) SetUniverseName(v string)`

SetUniverseName sets UniverseName field to given value.

### HasUniverseName

`func (o *BackupSpec) HasUniverseName() bool`

HasUniverseName returns a boolean if a field has been set.

### GetScheduleUuid

`func (o *BackupSpec) GetScheduleUuid() string`

GetScheduleUuid returns the ScheduleUuid field if non-nil, zero value otherwise.

### GetScheduleUuidOk

`func (o *BackupSpec) GetScheduleUuidOk() (*string, bool)`

GetScheduleUuidOk returns a tuple with the ScheduleUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUuid

`func (o *BackupSpec) SetScheduleUuid(v string)`

SetScheduleUuid sets ScheduleUuid field to given value.

### HasScheduleUuid

`func (o *BackupSpec) HasScheduleUuid() bool`

HasScheduleUuid returns a boolean if a field has been set.

### GetScheduleName

`func (o *BackupSpec) GetScheduleName() string`

GetScheduleName returns the ScheduleName field if non-nil, zero value otherwise.

### GetScheduleNameOk

`func (o *BackupSpec) GetScheduleNameOk() (*string, bool)`

GetScheduleNameOk returns a tuple with the ScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleName

`func (o *BackupSpec) SetScheduleName(v string)`

SetScheduleName sets ScheduleName field to given value.

### HasScheduleName

`func (o *BackupSpec) HasScheduleName() bool`

HasScheduleName returns a boolean if a field has been set.

### GetBackupType

`func (o *BackupSpec) GetBackupType() TableType`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *BackupSpec) GetBackupTypeOk() (*TableType, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *BackupSpec) SetBackupType(v TableType)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *BackupSpec) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetCategory

`func (o *BackupSpec) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *BackupSpec) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *BackupSpec) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *BackupSpec) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetStorageConfigType

`func (o *BackupSpec) GetStorageConfigType() string`

GetStorageConfigType returns the StorageConfigType field if non-nil, zero value otherwise.

### GetStorageConfigTypeOk

`func (o *BackupSpec) GetStorageConfigTypeOk() (*string, bool)`

GetStorageConfigTypeOk returns a tuple with the StorageConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigType

`func (o *BackupSpec) SetStorageConfigType(v string)`

SetStorageConfigType sets StorageConfigType field to given value.

### HasStorageConfigType

`func (o *BackupSpec) HasStorageConfigType() bool`

HasStorageConfigType returns a boolean if a field has been set.

### GetIsFullBackup

`func (o *BackupSpec) GetIsFullBackup() bool`

GetIsFullBackup returns the IsFullBackup field if non-nil, zero value otherwise.

### GetIsFullBackupOk

`func (o *BackupSpec) GetIsFullBackupOk() (*bool, bool)`

GetIsFullBackupOk returns a tuple with the IsFullBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullBackup

`func (o *BackupSpec) SetIsFullBackup(v bool)`

SetIsFullBackup sets IsFullBackup field to given value.

### HasIsFullBackup

`func (o *BackupSpec) HasIsFullBackup() bool`

HasIsFullBackup returns a boolean if a field has been set.

### GetOnDemand

`func (o *BackupSpec) GetOnDemand() bool`

GetOnDemand returns the OnDemand field if non-nil, zero value otherwise.

### GetOnDemandOk

`func (o *BackupSpec) GetOnDemandOk() (*bool, bool)`

GetOnDemandOk returns a tuple with the OnDemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDemand

`func (o *BackupSpec) SetOnDemand(v bool)`

SetOnDemand sets OnDemand field to given value.

### HasOnDemand

`func (o *BackupSpec) HasOnDemand() bool`

HasOnDemand returns a boolean if a field has been set.

### GetUseTablespaces

`func (o *BackupSpec) GetUseTablespaces() bool`

GetUseTablespaces returns the UseTablespaces field if non-nil, zero value otherwise.

### GetUseTablespacesOk

`func (o *BackupSpec) GetUseTablespacesOk() (*bool, bool)`

GetUseTablespacesOk returns a tuple with the UseTablespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTablespaces

`func (o *BackupSpec) SetUseTablespaces(v bool)`

SetUseTablespaces sets UseTablespaces field to given value.

### HasUseTablespaces

`func (o *BackupSpec) HasUseTablespaces() bool`

HasUseTablespaces returns a boolean if a field has been set.

### GetUseRoles

`func (o *BackupSpec) GetUseRoles() bool`

GetUseRoles returns the UseRoles field if non-nil, zero value otherwise.

### GetUseRolesOk

`func (o *BackupSpec) GetUseRolesOk() (*bool, bool)`

GetUseRolesOk returns a tuple with the UseRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRoles

`func (o *BackupSpec) SetUseRoles(v bool)`

SetUseRoles sets UseRoles field to given value.

### HasUseRoles

`func (o *BackupSpec) HasUseRoles() bool`

HasUseRoles returns a boolean if a field has been set.

### GetExpiryTime

`func (o *BackupSpec) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *BackupSpec) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *BackupSpec) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *BackupSpec) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetExpiryTimeUnit

`func (o *BackupSpec) GetExpiryTimeUnit() string`

GetExpiryTimeUnit returns the ExpiryTimeUnit field if non-nil, zero value otherwise.

### GetExpiryTimeUnitOk

`func (o *BackupSpec) GetExpiryTimeUnitOk() (*string, bool)`

GetExpiryTimeUnitOk returns a tuple with the ExpiryTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUnit

`func (o *BackupSpec) SetExpiryTimeUnit(v string)`

SetExpiryTimeUnit sets ExpiryTimeUnit field to given value.

### HasExpiryTimeUnit

`func (o *BackupSpec) HasExpiryTimeUnit() bool`

HasExpiryTimeUnit returns a boolean if a field has been set.

### GetKeyspaceTables

`func (o *BackupSpec) GetKeyspaceTables() []BackupKeyspaceTables`

GetKeyspaceTables returns the KeyspaceTables field if non-nil, zero value otherwise.

### GetKeyspaceTablesOk

`func (o *BackupSpec) GetKeyspaceTablesOk() (*[]BackupKeyspaceTables, bool)`

GetKeyspaceTablesOk returns a tuple with the KeyspaceTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspaceTables

`func (o *BackupSpec) SetKeyspaceTables(v []BackupKeyspaceTables)`

SetKeyspaceTables sets KeyspaceTables field to given value.

### HasKeyspaceTables

`func (o *BackupSpec) HasKeyspaceTables() bool`

HasKeyspaceTables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


