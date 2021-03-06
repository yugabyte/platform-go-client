# BackupStorageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupType** | Pointer to **string** | Backup type | [optional] 
**Keyspace** | Pointer to **string** | Keyspace name | [optional] 
**Sse** | Pointer to **bool** | Is SSE | [optional] 
**StorageLocation** | Pointer to **string** | Storage location | [optional] 
**TableNameList** | Pointer to **[]string** | Tables | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


