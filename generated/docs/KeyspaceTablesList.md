# KeyspaceTablesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keyspace** | **string** |  | 
**StorageLocation** | **string** |  | 
**TablesList** | **[]string** |  | 

## Methods

### NewKeyspaceTablesList

`func NewKeyspaceTablesList(keyspace string, storageLocation string, tablesList []string, ) *KeyspaceTablesList`

NewKeyspaceTablesList instantiates a new KeyspaceTablesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyspaceTablesListWithDefaults

`func NewKeyspaceTablesListWithDefaults() *KeyspaceTablesList`

NewKeyspaceTablesListWithDefaults instantiates a new KeyspaceTablesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetStorageLocation

`func (o *KeyspaceTablesList) GetStorageLocation() string`

GetStorageLocation returns the StorageLocation field if non-nil, zero value otherwise.

### GetStorageLocationOk

`func (o *KeyspaceTablesList) GetStorageLocationOk() (*string, bool)`

GetStorageLocationOk returns a tuple with the StorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLocation

`func (o *KeyspaceTablesList) SetStorageLocation(v string)`

SetStorageLocation sets StorageLocation field to given value.


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


