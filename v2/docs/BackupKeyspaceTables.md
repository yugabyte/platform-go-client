# BackupKeyspaceTables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keyspace** | Pointer to **string** | Keyspace or namespace name. | [optional] [readonly] 
**AllTables** | Pointer to **bool** | Whether all tables in the keyspace are included. | [optional] [readonly] 
**TablesList** | Pointer to **[]string** | Table names included in the backup. | [optional] [readonly] 
**TableUuidList** | Pointer to **[]string** | Table UUIDs included in the backup. | [optional] [readonly] 
**BackupSizeInBytes** | Pointer to **int64** | Backup size for this keyspace, in bytes. | [optional] [readonly] 
**DefaultLocation** | Pointer to **string** | Default storage location for this keyspace backup. | [optional] [readonly] 
**PerRegionLocations** | Pointer to [**[]BackupRegionLocation**](BackupRegionLocation.md) | Per-region storage locations for this keyspace backup. | [optional] [readonly] 
**BackupPointInTimeRestoreWindow** | Pointer to [**BackupPointInTimeRestoreWindow**](BackupPointInTimeRestoreWindow.md) |  | [optional] 

## Methods

### NewBackupKeyspaceTables

`func NewBackupKeyspaceTables() *BackupKeyspaceTables`

NewBackupKeyspaceTables instantiates a new BackupKeyspaceTables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupKeyspaceTablesWithDefaults

`func NewBackupKeyspaceTablesWithDefaults() *BackupKeyspaceTables`

NewBackupKeyspaceTablesWithDefaults instantiates a new BackupKeyspaceTables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyspace

`func (o *BackupKeyspaceTables) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *BackupKeyspaceTables) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *BackupKeyspaceTables) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.

### HasKeyspace

`func (o *BackupKeyspaceTables) HasKeyspace() bool`

HasKeyspace returns a boolean if a field has been set.

### GetAllTables

`func (o *BackupKeyspaceTables) GetAllTables() bool`

GetAllTables returns the AllTables field if non-nil, zero value otherwise.

### GetAllTablesOk

`func (o *BackupKeyspaceTables) GetAllTablesOk() (*bool, bool)`

GetAllTablesOk returns a tuple with the AllTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTables

`func (o *BackupKeyspaceTables) SetAllTables(v bool)`

SetAllTables sets AllTables field to given value.

### HasAllTables

`func (o *BackupKeyspaceTables) HasAllTables() bool`

HasAllTables returns a boolean if a field has been set.

### GetTablesList

`func (o *BackupKeyspaceTables) GetTablesList() []string`

GetTablesList returns the TablesList field if non-nil, zero value otherwise.

### GetTablesListOk

`func (o *BackupKeyspaceTables) GetTablesListOk() (*[]string, bool)`

GetTablesListOk returns a tuple with the TablesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTablesList

`func (o *BackupKeyspaceTables) SetTablesList(v []string)`

SetTablesList sets TablesList field to given value.

### HasTablesList

`func (o *BackupKeyspaceTables) HasTablesList() bool`

HasTablesList returns a boolean if a field has been set.

### GetTableUuidList

`func (o *BackupKeyspaceTables) GetTableUuidList() []string`

GetTableUuidList returns the TableUuidList field if non-nil, zero value otherwise.

### GetTableUuidListOk

`func (o *BackupKeyspaceTables) GetTableUuidListOk() (*[]string, bool)`

GetTableUuidListOk returns a tuple with the TableUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableUuidList

`func (o *BackupKeyspaceTables) SetTableUuidList(v []string)`

SetTableUuidList sets TableUuidList field to given value.

### HasTableUuidList

`func (o *BackupKeyspaceTables) HasTableUuidList() bool`

HasTableUuidList returns a boolean if a field has been set.

### GetBackupSizeInBytes

`func (o *BackupKeyspaceTables) GetBackupSizeInBytes() int64`

GetBackupSizeInBytes returns the BackupSizeInBytes field if non-nil, zero value otherwise.

### GetBackupSizeInBytesOk

`func (o *BackupKeyspaceTables) GetBackupSizeInBytesOk() (*int64, bool)`

GetBackupSizeInBytesOk returns a tuple with the BackupSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSizeInBytes

`func (o *BackupKeyspaceTables) SetBackupSizeInBytes(v int64)`

SetBackupSizeInBytes sets BackupSizeInBytes field to given value.

### HasBackupSizeInBytes

`func (o *BackupKeyspaceTables) HasBackupSizeInBytes() bool`

HasBackupSizeInBytes returns a boolean if a field has been set.

### GetDefaultLocation

`func (o *BackupKeyspaceTables) GetDefaultLocation() string`

GetDefaultLocation returns the DefaultLocation field if non-nil, zero value otherwise.

### GetDefaultLocationOk

`func (o *BackupKeyspaceTables) GetDefaultLocationOk() (*string, bool)`

GetDefaultLocationOk returns a tuple with the DefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLocation

`func (o *BackupKeyspaceTables) SetDefaultLocation(v string)`

SetDefaultLocation sets DefaultLocation field to given value.

### HasDefaultLocation

`func (o *BackupKeyspaceTables) HasDefaultLocation() bool`

HasDefaultLocation returns a boolean if a field has been set.

### GetPerRegionLocations

`func (o *BackupKeyspaceTables) GetPerRegionLocations() []BackupRegionLocation`

GetPerRegionLocations returns the PerRegionLocations field if non-nil, zero value otherwise.

### GetPerRegionLocationsOk

`func (o *BackupKeyspaceTables) GetPerRegionLocationsOk() (*[]BackupRegionLocation, bool)`

GetPerRegionLocationsOk returns a tuple with the PerRegionLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerRegionLocations

`func (o *BackupKeyspaceTables) SetPerRegionLocations(v []BackupRegionLocation)`

SetPerRegionLocations sets PerRegionLocations field to given value.

### HasPerRegionLocations

`func (o *BackupKeyspaceTables) HasPerRegionLocations() bool`

HasPerRegionLocations returns a boolean if a field has been set.

### GetBackupPointInTimeRestoreWindow

`func (o *BackupKeyspaceTables) GetBackupPointInTimeRestoreWindow() BackupPointInTimeRestoreWindow`

GetBackupPointInTimeRestoreWindow returns the BackupPointInTimeRestoreWindow field if non-nil, zero value otherwise.

### GetBackupPointInTimeRestoreWindowOk

`func (o *BackupKeyspaceTables) GetBackupPointInTimeRestoreWindowOk() (*BackupPointInTimeRestoreWindow, bool)`

GetBackupPointInTimeRestoreWindowOk returns a tuple with the BackupPointInTimeRestoreWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPointInTimeRestoreWindow

`func (o *BackupKeyspaceTables) SetBackupPointInTimeRestoreWindow(v BackupPointInTimeRestoreWindow)`

SetBackupPointInTimeRestoreWindow sets BackupPointInTimeRestoreWindow field to given value.

### HasBackupPointInTimeRestoreWindow

`func (o *BackupKeyspaceTables) HasBackupPointInTimeRestoreWindow() bool`

HasBackupPointInTimeRestoreWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


