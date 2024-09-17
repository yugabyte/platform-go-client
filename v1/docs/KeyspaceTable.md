# KeyspaceTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keyspace** | Pointer to **string** | keyspace | [optional] 
**TableNameList** | Pointer to **[]string** | Tables | [optional] 
**TableUUIDList** | Pointer to **[]string** | Table UUIDs | [optional] 

## Methods

### NewKeyspaceTable

`func NewKeyspaceTable() *KeyspaceTable`

NewKeyspaceTable instantiates a new KeyspaceTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyspaceTableWithDefaults

`func NewKeyspaceTableWithDefaults() *KeyspaceTable`

NewKeyspaceTableWithDefaults instantiates a new KeyspaceTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyspace

`func (o *KeyspaceTable) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *KeyspaceTable) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *KeyspaceTable) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.

### HasKeyspace

`func (o *KeyspaceTable) HasKeyspace() bool`

HasKeyspace returns a boolean if a field has been set.

### GetTableNameList

`func (o *KeyspaceTable) GetTableNameList() []string`

GetTableNameList returns the TableNameList field if non-nil, zero value otherwise.

### GetTableNameListOk

`func (o *KeyspaceTable) GetTableNameListOk() (*[]string, bool)`

GetTableNameListOk returns a tuple with the TableNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableNameList

`func (o *KeyspaceTable) SetTableNameList(v []string)`

SetTableNameList sets TableNameList field to given value.

### HasTableNameList

`func (o *KeyspaceTable) HasTableNameList() bool`

HasTableNameList returns a boolean if a field has been set.

### GetTableUUIDList

`func (o *KeyspaceTable) GetTableUUIDList() []string`

GetTableUUIDList returns the TableUUIDList field if non-nil, zero value otherwise.

### GetTableUUIDListOk

`func (o *KeyspaceTable) GetTableUUIDListOk() (*[]string, bool)`

GetTableUUIDListOk returns a tuple with the TableUUIDList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableUUIDList

`func (o *KeyspaceTable) SetTableUUIDList(v []string)`

SetTableUUIDList sets TableUUIDList field to given value.

### HasTableUUIDList

`func (o *KeyspaceTable) HasTableUUIDList() bool`

HasTableUUIDList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


