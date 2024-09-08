# PerLocationBackupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupLocation** | Pointer to **string** | Backup location | [optional] 
**IsSelectiveRestoreSupported** | Pointer to **bool** | Whether selective table restore is supported for this backup | [optional] 
**IsYSQLBackup** | Pointer to **bool** | Whether backup type is YSQL | [optional] 
**PerBackupLocationKeyspaceTables** | Pointer to [**PerBackupLocationKeyspaceTables**](PerBackupLocationKeyspaceTables.md) |  | [optional] 
**PointInTimeRestoreWindow** | Pointer to [**BackupPointInTimeRestoreWindow**](BackupPointInTimeRestoreWindow.md) |  | [optional] 
**TablespaceResponse** | Pointer to [**TablespaceResponse**](TablespaceResponse.md) |  | [optional] 

## Methods

### NewPerLocationBackupInfo

`func NewPerLocationBackupInfo() *PerLocationBackupInfo`

NewPerLocationBackupInfo instantiates a new PerLocationBackupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerLocationBackupInfoWithDefaults

`func NewPerLocationBackupInfoWithDefaults() *PerLocationBackupInfo`

NewPerLocationBackupInfoWithDefaults instantiates a new PerLocationBackupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupLocation

`func (o *PerLocationBackupInfo) GetBackupLocation() string`

GetBackupLocation returns the BackupLocation field if non-nil, zero value otherwise.

### GetBackupLocationOk

`func (o *PerLocationBackupInfo) GetBackupLocationOk() (*string, bool)`

GetBackupLocationOk returns a tuple with the BackupLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLocation

`func (o *PerLocationBackupInfo) SetBackupLocation(v string)`

SetBackupLocation sets BackupLocation field to given value.

### HasBackupLocation

`func (o *PerLocationBackupInfo) HasBackupLocation() bool`

HasBackupLocation returns a boolean if a field has been set.

### GetIsSelectiveRestoreSupported

`func (o *PerLocationBackupInfo) GetIsSelectiveRestoreSupported() bool`

GetIsSelectiveRestoreSupported returns the IsSelectiveRestoreSupported field if non-nil, zero value otherwise.

### GetIsSelectiveRestoreSupportedOk

`func (o *PerLocationBackupInfo) GetIsSelectiveRestoreSupportedOk() (*bool, bool)`

GetIsSelectiveRestoreSupportedOk returns a tuple with the IsSelectiveRestoreSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelectiveRestoreSupported

`func (o *PerLocationBackupInfo) SetIsSelectiveRestoreSupported(v bool)`

SetIsSelectiveRestoreSupported sets IsSelectiveRestoreSupported field to given value.

### HasIsSelectiveRestoreSupported

`func (o *PerLocationBackupInfo) HasIsSelectiveRestoreSupported() bool`

HasIsSelectiveRestoreSupported returns a boolean if a field has been set.

### GetIsYSQLBackup

`func (o *PerLocationBackupInfo) GetIsYSQLBackup() bool`

GetIsYSQLBackup returns the IsYSQLBackup field if non-nil, zero value otherwise.

### GetIsYSQLBackupOk

`func (o *PerLocationBackupInfo) GetIsYSQLBackupOk() (*bool, bool)`

GetIsYSQLBackupOk returns a tuple with the IsYSQLBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYSQLBackup

`func (o *PerLocationBackupInfo) SetIsYSQLBackup(v bool)`

SetIsYSQLBackup sets IsYSQLBackup field to given value.

### HasIsYSQLBackup

`func (o *PerLocationBackupInfo) HasIsYSQLBackup() bool`

HasIsYSQLBackup returns a boolean if a field has been set.

### GetPerBackupLocationKeyspaceTables

`func (o *PerLocationBackupInfo) GetPerBackupLocationKeyspaceTables() PerBackupLocationKeyspaceTables`

GetPerBackupLocationKeyspaceTables returns the PerBackupLocationKeyspaceTables field if non-nil, zero value otherwise.

### GetPerBackupLocationKeyspaceTablesOk

`func (o *PerLocationBackupInfo) GetPerBackupLocationKeyspaceTablesOk() (*PerBackupLocationKeyspaceTables, bool)`

GetPerBackupLocationKeyspaceTablesOk returns a tuple with the PerBackupLocationKeyspaceTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerBackupLocationKeyspaceTables

`func (o *PerLocationBackupInfo) SetPerBackupLocationKeyspaceTables(v PerBackupLocationKeyspaceTables)`

SetPerBackupLocationKeyspaceTables sets PerBackupLocationKeyspaceTables field to given value.

### HasPerBackupLocationKeyspaceTables

`func (o *PerLocationBackupInfo) HasPerBackupLocationKeyspaceTables() bool`

HasPerBackupLocationKeyspaceTables returns a boolean if a field has been set.

### GetPointInTimeRestoreWindow

`func (o *PerLocationBackupInfo) GetPointInTimeRestoreWindow() BackupPointInTimeRestoreWindow`

GetPointInTimeRestoreWindow returns the PointInTimeRestoreWindow field if non-nil, zero value otherwise.

### GetPointInTimeRestoreWindowOk

`func (o *PerLocationBackupInfo) GetPointInTimeRestoreWindowOk() (*BackupPointInTimeRestoreWindow, bool)`

GetPointInTimeRestoreWindowOk returns a tuple with the PointInTimeRestoreWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeRestoreWindow

`func (o *PerLocationBackupInfo) SetPointInTimeRestoreWindow(v BackupPointInTimeRestoreWindow)`

SetPointInTimeRestoreWindow sets PointInTimeRestoreWindow field to given value.

### HasPointInTimeRestoreWindow

`func (o *PerLocationBackupInfo) HasPointInTimeRestoreWindow() bool`

HasPointInTimeRestoreWindow returns a boolean if a field has been set.

### GetTablespaceResponse

`func (o *PerLocationBackupInfo) GetTablespaceResponse() TablespaceResponse`

GetTablespaceResponse returns the TablespaceResponse field if non-nil, zero value otherwise.

### GetTablespaceResponseOk

`func (o *PerLocationBackupInfo) GetTablespaceResponseOk() (*TablespaceResponse, bool)`

GetTablespaceResponseOk returns a tuple with the TablespaceResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTablespaceResponse

`func (o *PerLocationBackupInfo) SetTablespaceResponse(v TablespaceResponse)`

SetTablespaceResponse sets TablespaceResponse field to given value.

### HasTablespaceResponse

`func (o *PerLocationBackupInfo) HasTablespaceResponse() bool`

HasTablespaceResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


