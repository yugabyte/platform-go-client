# KeyspaceTables

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keyspace** | Pointer to **string** | Keyspace | [optional] 
**TableNames** | Pointer to **[]string** | Tables | [optional] 

## Methods

### NewKeyspaceTables

`func NewKeyspaceTables() *KeyspaceTables`

NewKeyspaceTables instantiates a new KeyspaceTables object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyspaceTablesWithDefaults

`func NewKeyspaceTablesWithDefaults() *KeyspaceTables`

NewKeyspaceTablesWithDefaults instantiates a new KeyspaceTables object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyspace

`func (o *KeyspaceTables) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *KeyspaceTables) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *KeyspaceTables) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.

### HasKeyspace

`func (o *KeyspaceTables) HasKeyspace() bool`

HasKeyspace returns a boolean if a field has been set.

### GetTableNames

`func (o *KeyspaceTables) GetTableNames() []string`

GetTableNames returns the TableNames field if non-nil, zero value otherwise.

### GetTableNamesOk

`func (o *KeyspaceTables) GetTableNamesOk() (*[]string, bool)`

GetTableNamesOk returns a tuple with the TableNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableNames

`func (o *KeyspaceTables) SetTableNames(v []string)`

SetTableNames sets TableNames field to given value.

### HasTableNames

`func (o *KeyspaceTables) HasTableNames() bool`

HasTableNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


