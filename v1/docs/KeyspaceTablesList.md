# KeyspaceTablesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllTables** | **bool** |  | 
**BackupSizeInBytes** | **int64** |  | 
**DefaultLocation** | **string** |  | 
**Keyspace** | **string** |  | 
**PerRegionLocations** | [**[]RegionLocations**](RegionLocations.md) |  | 
**TableUUIDList** | **[]string** |  | 
**TablesList** | **[]string** |  | 

## Methods

### NewKeyspaceTablesList

`func NewKeyspaceTablesList(allTables bool, backupSizeInBytes int64, defaultLocation string, keyspace string, perRegionLocations []RegionLocations, tableUUIDList []string, tablesList []string, ) *KeyspaceTablesList`

NewKeyspaceTablesList instantiates a new KeyspaceTablesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyspaceTablesListWithDefaults

`func NewKeyspaceTablesListWithDefaults() *KeyspaceTablesList`

NewKeyspaceTablesListWithDefaults instantiates a new KeyspaceTablesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllTables

`func (o *KeyspaceTablesList) GetAllTables() bool`

GetAllTables returns the AllTables field if non-nil, zero value otherwise.

### GetAllTablesOk

`func (o *KeyspaceTablesList) GetAllTablesOk() (*bool, bool)`

GetAllTablesOk returns a tuple with the AllTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTables

`func (o *KeyspaceTablesList) SetAllTables(v bool)`

SetAllTables sets AllTables field to given value.


### GetBackupSizeInBytes

`func (o *KeyspaceTablesList) GetBackupSizeInBytes() int64`

GetBackupSizeInBytes returns the BackupSizeInBytes field if non-nil, zero value otherwise.

### GetBackupSizeInBytesOk

`func (o *KeyspaceTablesList) GetBackupSizeInBytesOk() (*int64, bool)`

GetBackupSizeInBytesOk returns a tuple with the BackupSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSizeInBytes

`func (o *KeyspaceTablesList) SetBackupSizeInBytes(v int64)`

SetBackupSizeInBytes sets BackupSizeInBytes field to given value.


### GetDefaultLocation

`func (o *KeyspaceTablesList) GetDefaultLocation() string`

GetDefaultLocation returns the DefaultLocation field if non-nil, zero value otherwise.

### GetDefaultLocationOk

`func (o *KeyspaceTablesList) GetDefaultLocationOk() (*string, bool)`

GetDefaultLocationOk returns a tuple with the DefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLocation

`func (o *KeyspaceTablesList) SetDefaultLocation(v string)`

SetDefaultLocation sets DefaultLocation field to given value.


### GetKeyspace

`func (o *KeyspaceTablesList) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *KeyspaceTablesList) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *KeyspaceTablesList) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.


### GetPerRegionLocations

`func (o *KeyspaceTablesList) GetPerRegionLocations() []RegionLocations`

GetPerRegionLocations returns the PerRegionLocations field if non-nil, zero value otherwise.

### GetPerRegionLocationsOk

`func (o *KeyspaceTablesList) GetPerRegionLocationsOk() (*[]RegionLocations, bool)`

GetPerRegionLocationsOk returns a tuple with the PerRegionLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerRegionLocations

`func (o *KeyspaceTablesList) SetPerRegionLocations(v []RegionLocations)`

SetPerRegionLocations sets PerRegionLocations field to given value.


### GetTableUUIDList

`func (o *KeyspaceTablesList) GetTableUUIDList() []string`

GetTableUUIDList returns the TableUUIDList field if non-nil, zero value otherwise.

### GetTableUUIDListOk

`func (o *KeyspaceTablesList) GetTableUUIDListOk() (*[]string, bool)`

GetTableUUIDListOk returns a tuple with the TableUUIDList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableUUIDList

`func (o *KeyspaceTablesList) SetTableUUIDList(v []string)`

SetTableUUIDList sets TableUUIDList field to given value.


### GetTablesList

`func (o *KeyspaceTablesList) GetTablesList() []string`

GetTablesList returns the TablesList field if non-nil, zero value otherwise.

### GetTablesListOk

`func (o *KeyspaceTablesList) GetTablesListOk() (*[]string, bool)`

GetTablesListOk returns a tuple with the TablesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTablesList

`func (o *KeyspaceTablesList) SetTablesList(v []string)`

SetTablesList sets TablesList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


