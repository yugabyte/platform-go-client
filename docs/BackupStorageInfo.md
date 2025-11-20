# BackupStorageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupType** | Pointer to **string** | Backup type | [optional] 
**ErrorIfRolesExists** | Pointer to **bool** | WARNING: This is a preview API that could change. Fail if roles exist with same names.Only applicable with new roles behavior. | [optional] 
**ErrorIfTablespacesExists** | Pointer to **bool** | WARNING: This is a preview API that could change. Fail if tablespaces exist with same names. Only applicable with new roles behavior. | [optional] 
**IgnoreErrors** | Pointer to **bool** | WARNING: This is a preview API that could change. Ignore all restore errors | [optional] 
**Keyspace** | Pointer to **string** | Keyspace name | [optional] 
**NewOwner** | Pointer to **string** | User name of the new tables owner | [optional] 
**OldOwner** | Pointer to **string** | User name of the current tables owner | [optional] 
**SelectiveTableRestore** | Pointer to **bool** | Only restore selected tables instead of restoring all tables in backup | [optional] 
**Sse** | Pointer to **bool** | Is SSE | [optional] 
**StorageLocation** | Pointer to **string** | Storage location | [optional] 
**TableNameList** | Pointer to **[]string** | Tables | [optional] 
**UseRoles** | Pointer to **bool** | Restore global ysql roles | [optional] 
**UseTablespaces** | Pointer to **bool** | Use tablespaces during restore | [optional] 

## Methods

### NewBackupStorageInfo

`func NewBackupStorageInfo() *BackupStorageInfo`

NewBackupStorageInfo instantiates a new BackupStorageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupStorageInfoWithDefaults

`func NewBackupStorageInfoWithDefaults() *BackupStorageInfo`

NewBackupStorageInfoWithDefaults instantiates a new BackupStorageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupType

`func (o *BackupStorageInfo) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *BackupStorageInfo) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *BackupStorageInfo) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *BackupStorageInfo) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetErrorIfRolesExists

`func (o *BackupStorageInfo) GetErrorIfRolesExists() bool`

GetErrorIfRolesExists returns the ErrorIfRolesExists field if non-nil, zero value otherwise.

### GetErrorIfRolesExistsOk

`func (o *BackupStorageInfo) GetErrorIfRolesExistsOk() (*bool, bool)`

GetErrorIfRolesExistsOk returns a tuple with the ErrorIfRolesExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorIfRolesExists

`func (o *BackupStorageInfo) SetErrorIfRolesExists(v bool)`

SetErrorIfRolesExists sets ErrorIfRolesExists field to given value.

### HasErrorIfRolesExists

`func (o *BackupStorageInfo) HasErrorIfRolesExists() bool`

HasErrorIfRolesExists returns a boolean if a field has been set.

### GetErrorIfTablespacesExists

`func (o *BackupStorageInfo) GetErrorIfTablespacesExists() bool`

GetErrorIfTablespacesExists returns the ErrorIfTablespacesExists field if non-nil, zero value otherwise.

### GetErrorIfTablespacesExistsOk

`func (o *BackupStorageInfo) GetErrorIfTablespacesExistsOk() (*bool, bool)`

GetErrorIfTablespacesExistsOk returns a tuple with the ErrorIfTablespacesExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorIfTablespacesExists

`func (o *BackupStorageInfo) SetErrorIfTablespacesExists(v bool)`

SetErrorIfTablespacesExists sets ErrorIfTablespacesExists field to given value.

### HasErrorIfTablespacesExists

`func (o *BackupStorageInfo) HasErrorIfTablespacesExists() bool`

HasErrorIfTablespacesExists returns a boolean if a field has been set.

### GetIgnoreErrors

`func (o *BackupStorageInfo) GetIgnoreErrors() bool`

GetIgnoreErrors returns the IgnoreErrors field if non-nil, zero value otherwise.

### GetIgnoreErrorsOk

`func (o *BackupStorageInfo) GetIgnoreErrorsOk() (*bool, bool)`

GetIgnoreErrorsOk returns a tuple with the IgnoreErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreErrors

`func (o *BackupStorageInfo) SetIgnoreErrors(v bool)`

SetIgnoreErrors sets IgnoreErrors field to given value.

### HasIgnoreErrors

`func (o *BackupStorageInfo) HasIgnoreErrors() bool`

HasIgnoreErrors returns a boolean if a field has been set.

### GetKeyspace

`func (o *BackupStorageInfo) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *BackupStorageInfo) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *BackupStorageInfo) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.

### HasKeyspace

`func (o *BackupStorageInfo) HasKeyspace() bool`

HasKeyspace returns a boolean if a field has been set.

### GetNewOwner

`func (o *BackupStorageInfo) GetNewOwner() string`

GetNewOwner returns the NewOwner field if non-nil, zero value otherwise.

### GetNewOwnerOk

`func (o *BackupStorageInfo) GetNewOwnerOk() (*string, bool)`

GetNewOwnerOk returns a tuple with the NewOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOwner

`func (o *BackupStorageInfo) SetNewOwner(v string)`

SetNewOwner sets NewOwner field to given value.

### HasNewOwner

`func (o *BackupStorageInfo) HasNewOwner() bool`

HasNewOwner returns a boolean if a field has been set.

### GetOldOwner

`func (o *BackupStorageInfo) GetOldOwner() string`

GetOldOwner returns the OldOwner field if non-nil, zero value otherwise.

### GetOldOwnerOk

`func (o *BackupStorageInfo) GetOldOwnerOk() (*string, bool)`

GetOldOwnerOk returns a tuple with the OldOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldOwner

`func (o *BackupStorageInfo) SetOldOwner(v string)`

SetOldOwner sets OldOwner field to given value.

### HasOldOwner

`func (o *BackupStorageInfo) HasOldOwner() bool`

HasOldOwner returns a boolean if a field has been set.

### GetSelectiveTableRestore

`func (o *BackupStorageInfo) GetSelectiveTableRestore() bool`

GetSelectiveTableRestore returns the SelectiveTableRestore field if non-nil, zero value otherwise.

### GetSelectiveTableRestoreOk

`func (o *BackupStorageInfo) GetSelectiveTableRestoreOk() (*bool, bool)`

GetSelectiveTableRestoreOk returns a tuple with the SelectiveTableRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectiveTableRestore

`func (o *BackupStorageInfo) SetSelectiveTableRestore(v bool)`

SetSelectiveTableRestore sets SelectiveTableRestore field to given value.

### HasSelectiveTableRestore

`func (o *BackupStorageInfo) HasSelectiveTableRestore() bool`

HasSelectiveTableRestore returns a boolean if a field has been set.

### GetSse

`func (o *BackupStorageInfo) GetSse() bool`

GetSse returns the Sse field if non-nil, zero value otherwise.

### GetSseOk

`func (o *BackupStorageInfo) GetSseOk() (*bool, bool)`

GetSseOk returns a tuple with the Sse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSse

`func (o *BackupStorageInfo) SetSse(v bool)`

SetSse sets Sse field to given value.

### HasSse

`func (o *BackupStorageInfo) HasSse() bool`

HasSse returns a boolean if a field has been set.

### GetStorageLocation

`func (o *BackupStorageInfo) GetStorageLocation() string`

GetStorageLocation returns the StorageLocation field if non-nil, zero value otherwise.

### GetStorageLocationOk

`func (o *BackupStorageInfo) GetStorageLocationOk() (*string, bool)`

GetStorageLocationOk returns a tuple with the StorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLocation

`func (o *BackupStorageInfo) SetStorageLocation(v string)`

SetStorageLocation sets StorageLocation field to given value.

### HasStorageLocation

`func (o *BackupStorageInfo) HasStorageLocation() bool`

HasStorageLocation returns a boolean if a field has been set.

### GetTableNameList

`func (o *BackupStorageInfo) GetTableNameList() []string`

GetTableNameList returns the TableNameList field if non-nil, zero value otherwise.

### GetTableNameListOk

`func (o *BackupStorageInfo) GetTableNameListOk() (*[]string, bool)`

GetTableNameListOk returns a tuple with the TableNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableNameList

`func (o *BackupStorageInfo) SetTableNameList(v []string)`

SetTableNameList sets TableNameList field to given value.

### HasTableNameList

`func (o *BackupStorageInfo) HasTableNameList() bool`

HasTableNameList returns a boolean if a field has been set.

### GetUseRoles

`func (o *BackupStorageInfo) GetUseRoles() bool`

GetUseRoles returns the UseRoles field if non-nil, zero value otherwise.

### GetUseRolesOk

`func (o *BackupStorageInfo) GetUseRolesOk() (*bool, bool)`

GetUseRolesOk returns a tuple with the UseRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRoles

`func (o *BackupStorageInfo) SetUseRoles(v bool)`

SetUseRoles sets UseRoles field to given value.

### HasUseRoles

`func (o *BackupStorageInfo) HasUseRoles() bool`

HasUseRoles returns a boolean if a field has been set.

### GetUseTablespaces

`func (o *BackupStorageInfo) GetUseTablespaces() bool`

GetUseTablespaces returns the UseTablespaces field if non-nil, zero value otherwise.

### GetUseTablespacesOk

`func (o *BackupStorageInfo) GetUseTablespacesOk() (*bool, bool)`

GetUseTablespacesOk returns a tuple with the UseTablespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTablespaces

`func (o *BackupStorageInfo) SetUseTablespaces(v bool)`

SetUseTablespaces sets UseTablespaces field to given value.

### HasUseTablespaces

`func (o *BackupStorageInfo) HasUseTablespaces() bool`

HasUseTablespaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


