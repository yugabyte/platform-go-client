# PerBackupLocationKeyspaceTables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalKeyspace** | Pointer to **string** | Original keyspace name | [optional] 
**TableNameList** | Pointer to **[]string** | List of parent tables associated with the keyspace | [optional] 

## Methods

### NewPerBackupLocationKeyspaceTables

`func NewPerBackupLocationKeyspaceTables() *PerBackupLocationKeyspaceTables`

NewPerBackupLocationKeyspaceTables instantiates a new PerBackupLocationKeyspaceTables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerBackupLocationKeyspaceTablesWithDefaults

`func NewPerBackupLocationKeyspaceTablesWithDefaults() *PerBackupLocationKeyspaceTables`

NewPerBackupLocationKeyspaceTablesWithDefaults instantiates a new PerBackupLocationKeyspaceTables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalKeyspace

`func (o *PerBackupLocationKeyspaceTables) GetOriginalKeyspace() string`

GetOriginalKeyspace returns the OriginalKeyspace field if non-nil, zero value otherwise.

### GetOriginalKeyspaceOk

`func (o *PerBackupLocationKeyspaceTables) GetOriginalKeyspaceOk() (*string, bool)`

GetOriginalKeyspaceOk returns a tuple with the OriginalKeyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalKeyspace

`func (o *PerBackupLocationKeyspaceTables) SetOriginalKeyspace(v string)`

SetOriginalKeyspace sets OriginalKeyspace field to given value.

### HasOriginalKeyspace

`func (o *PerBackupLocationKeyspaceTables) HasOriginalKeyspace() bool`

HasOriginalKeyspace returns a boolean if a field has been set.

### GetTableNameList

`func (o *PerBackupLocationKeyspaceTables) GetTableNameList() []string`

GetTableNameList returns the TableNameList field if non-nil, zero value otherwise.

### GetTableNameListOk

`func (o *PerBackupLocationKeyspaceTables) GetTableNameListOk() (*[]string, bool)`

GetTableNameListOk returns a tuple with the TableNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableNameList

`func (o *PerBackupLocationKeyspaceTables) SetTableNameList(v []string)`

SetTableNameList sets TableNameList field to given value.

### HasTableNameList

`func (o *PerBackupLocationKeyspaceTables) HasTableNameList() bool`

HasTableNameList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


