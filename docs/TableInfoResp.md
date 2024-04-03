# TableInfoResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Colocated** | Pointer to **bool** | Flag, indicating colocated table | [optional] 
**ColocationParentId** | Pointer to **string** | Colocation parent id | [optional] 
**IndexTableIDs** | Pointer to **[]string** | Index Table IDs of main table | [optional] 
**KeySpace** | Pointer to **string** | Keyspace | [optional] 
**MainTableUUID** | Pointer to **string** | Main Table UUID of index tables | [optional] 
**NameSpace** | Pointer to **string** | Namespace or Schema | [optional] 
**ParentTableUUID** | Pointer to **string** | Parent Table UUID | [optional] 
**PgSchemaName** | Pointer to **string** | Postgres schema name of the table | [optional] 
**RelationType** | Pointer to **string** | Relation type | [optional] 
**SizeBytes** | Pointer to **float64** | SST size in bytes | [optional] [readonly] 
**TableID** | Pointer to **string** | Table ID | [optional] [readonly] 
**TableName** | Pointer to **string** | Table name | [optional] 
**TableSpace** | Pointer to **string** | Table space | [optional] 
**TableType** | Pointer to **string** | Table type | [optional] 
**TableUUID** | Pointer to **string** | Table UUID | [optional] [readonly] 
**WalSizeBytes** | Pointer to **float64** | WAL size in bytes | [optional] [readonly] 

## Methods

### NewTableInfoResp

`func NewTableInfoResp() *TableInfoResp`

NewTableInfoResp instantiates a new TableInfoResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableInfoRespWithDefaults

`func NewTableInfoRespWithDefaults() *TableInfoResp`

NewTableInfoRespWithDefaults instantiates a new TableInfoResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColocated

`func (o *TableInfoResp) GetColocated() bool`

GetColocated returns the Colocated field if non-nil, zero value otherwise.

### GetColocatedOk

`func (o *TableInfoResp) GetColocatedOk() (*bool, bool)`

GetColocatedOk returns a tuple with the Colocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColocated

`func (o *TableInfoResp) SetColocated(v bool)`

SetColocated sets Colocated field to given value.

### HasColocated

`func (o *TableInfoResp) HasColocated() bool`

HasColocated returns a boolean if a field has been set.

### GetColocationParentId

`func (o *TableInfoResp) GetColocationParentId() string`

GetColocationParentId returns the ColocationParentId field if non-nil, zero value otherwise.

### GetColocationParentIdOk

`func (o *TableInfoResp) GetColocationParentIdOk() (*string, bool)`

GetColocationParentIdOk returns a tuple with the ColocationParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColocationParentId

`func (o *TableInfoResp) SetColocationParentId(v string)`

SetColocationParentId sets ColocationParentId field to given value.

### HasColocationParentId

`func (o *TableInfoResp) HasColocationParentId() bool`

HasColocationParentId returns a boolean if a field has been set.

### GetIndexTableIDs

`func (o *TableInfoResp) GetIndexTableIDs() []string`

GetIndexTableIDs returns the IndexTableIDs field if non-nil, zero value otherwise.

### GetIndexTableIDsOk

`func (o *TableInfoResp) GetIndexTableIDsOk() (*[]string, bool)`

GetIndexTableIDsOk returns a tuple with the IndexTableIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexTableIDs

`func (o *TableInfoResp) SetIndexTableIDs(v []string)`

SetIndexTableIDs sets IndexTableIDs field to given value.

### HasIndexTableIDs

`func (o *TableInfoResp) HasIndexTableIDs() bool`

HasIndexTableIDs returns a boolean if a field has been set.

### GetKeySpace

`func (o *TableInfoResp) GetKeySpace() string`

GetKeySpace returns the KeySpace field if non-nil, zero value otherwise.

### GetKeySpaceOk

`func (o *TableInfoResp) GetKeySpaceOk() (*string, bool)`

GetKeySpaceOk returns a tuple with the KeySpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySpace

`func (o *TableInfoResp) SetKeySpace(v string)`

SetKeySpace sets KeySpace field to given value.

### HasKeySpace

`func (o *TableInfoResp) HasKeySpace() bool`

HasKeySpace returns a boolean if a field has been set.

### GetMainTableUUID

`func (o *TableInfoResp) GetMainTableUUID() string`

GetMainTableUUID returns the MainTableUUID field if non-nil, zero value otherwise.

### GetMainTableUUIDOk

`func (o *TableInfoResp) GetMainTableUUIDOk() (*string, bool)`

GetMainTableUUIDOk returns a tuple with the MainTableUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainTableUUID

`func (o *TableInfoResp) SetMainTableUUID(v string)`

SetMainTableUUID sets MainTableUUID field to given value.

### HasMainTableUUID

`func (o *TableInfoResp) HasMainTableUUID() bool`

HasMainTableUUID returns a boolean if a field has been set.

### GetNameSpace

`func (o *TableInfoResp) GetNameSpace() string`

GetNameSpace returns the NameSpace field if non-nil, zero value otherwise.

### GetNameSpaceOk

`func (o *TableInfoResp) GetNameSpaceOk() (*string, bool)`

GetNameSpaceOk returns a tuple with the NameSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSpace

`func (o *TableInfoResp) SetNameSpace(v string)`

SetNameSpace sets NameSpace field to given value.

### HasNameSpace

`func (o *TableInfoResp) HasNameSpace() bool`

HasNameSpace returns a boolean if a field has been set.

### GetParentTableUUID

`func (o *TableInfoResp) GetParentTableUUID() string`

GetParentTableUUID returns the ParentTableUUID field if non-nil, zero value otherwise.

### GetParentTableUUIDOk

`func (o *TableInfoResp) GetParentTableUUIDOk() (*string, bool)`

GetParentTableUUIDOk returns a tuple with the ParentTableUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTableUUID

`func (o *TableInfoResp) SetParentTableUUID(v string)`

SetParentTableUUID sets ParentTableUUID field to given value.

### HasParentTableUUID

`func (o *TableInfoResp) HasParentTableUUID() bool`

HasParentTableUUID returns a boolean if a field has been set.

### GetPgSchemaName

`func (o *TableInfoResp) GetPgSchemaName() string`

GetPgSchemaName returns the PgSchemaName field if non-nil, zero value otherwise.

### GetPgSchemaNameOk

`func (o *TableInfoResp) GetPgSchemaNameOk() (*string, bool)`

GetPgSchemaNameOk returns a tuple with the PgSchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgSchemaName

`func (o *TableInfoResp) SetPgSchemaName(v string)`

SetPgSchemaName sets PgSchemaName field to given value.

### HasPgSchemaName

`func (o *TableInfoResp) HasPgSchemaName() bool`

HasPgSchemaName returns a boolean if a field has been set.

### GetRelationType

`func (o *TableInfoResp) GetRelationType() string`

GetRelationType returns the RelationType field if non-nil, zero value otherwise.

### GetRelationTypeOk

`func (o *TableInfoResp) GetRelationTypeOk() (*string, bool)`

GetRelationTypeOk returns a tuple with the RelationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationType

`func (o *TableInfoResp) SetRelationType(v string)`

SetRelationType sets RelationType field to given value.

### HasRelationType

`func (o *TableInfoResp) HasRelationType() bool`

HasRelationType returns a boolean if a field has been set.

### GetSizeBytes

`func (o *TableInfoResp) GetSizeBytes() float64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *TableInfoResp) GetSizeBytesOk() (*float64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *TableInfoResp) SetSizeBytes(v float64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *TableInfoResp) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetTableID

`func (o *TableInfoResp) GetTableID() string`

GetTableID returns the TableID field if non-nil, zero value otherwise.

### GetTableIDOk

`func (o *TableInfoResp) GetTableIDOk() (*string, bool)`

GetTableIDOk returns a tuple with the TableID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableID

`func (o *TableInfoResp) SetTableID(v string)`

SetTableID sets TableID field to given value.

### HasTableID

`func (o *TableInfoResp) HasTableID() bool`

HasTableID returns a boolean if a field has been set.

### GetTableName

`func (o *TableInfoResp) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *TableInfoResp) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *TableInfoResp) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *TableInfoResp) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetTableSpace

`func (o *TableInfoResp) GetTableSpace() string`

GetTableSpace returns the TableSpace field if non-nil, zero value otherwise.

### GetTableSpaceOk

`func (o *TableInfoResp) GetTableSpaceOk() (*string, bool)`

GetTableSpaceOk returns a tuple with the TableSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableSpace

`func (o *TableInfoResp) SetTableSpace(v string)`

SetTableSpace sets TableSpace field to given value.

### HasTableSpace

`func (o *TableInfoResp) HasTableSpace() bool`

HasTableSpace returns a boolean if a field has been set.

### GetTableType

`func (o *TableInfoResp) GetTableType() string`

GetTableType returns the TableType field if non-nil, zero value otherwise.

### GetTableTypeOk

`func (o *TableInfoResp) GetTableTypeOk() (*string, bool)`

GetTableTypeOk returns a tuple with the TableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableType

`func (o *TableInfoResp) SetTableType(v string)`

SetTableType sets TableType field to given value.

### HasTableType

`func (o *TableInfoResp) HasTableType() bool`

HasTableType returns a boolean if a field has been set.

### GetTableUUID

`func (o *TableInfoResp) GetTableUUID() string`

GetTableUUID returns the TableUUID field if non-nil, zero value otherwise.

### GetTableUUIDOk

`func (o *TableInfoResp) GetTableUUIDOk() (*string, bool)`

GetTableUUIDOk returns a tuple with the TableUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableUUID

`func (o *TableInfoResp) SetTableUUID(v string)`

SetTableUUID sets TableUUID field to given value.

### HasTableUUID

`func (o *TableInfoResp) HasTableUUID() bool`

HasTableUUID returns a boolean if a field has been set.

### GetWalSizeBytes

`func (o *TableInfoResp) GetWalSizeBytes() float64`

GetWalSizeBytes returns the WalSizeBytes field if non-nil, zero value otherwise.

### GetWalSizeBytesOk

`func (o *TableInfoResp) GetWalSizeBytesOk() (*float64, bool)`

GetWalSizeBytesOk returns a tuple with the WalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalSizeBytes

`func (o *TableInfoResp) SetWalSizeBytes(v float64)`

SetWalSizeBytes sets WalSizeBytes field to given value.

### HasWalSizeBytes

`func (o *TableInfoResp) HasWalSizeBytes() bool`

HasWalSizeBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


