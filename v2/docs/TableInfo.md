# TableInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableId** | Pointer to **string** | Table ID. | [optional] [readonly] 
**TableUuid** | Pointer to **string** | Table UUID. | [optional] [readonly] 
**Keyspace** | Pointer to **string** | Keyspace name. | [optional] [readonly] 
**TableType** | Pointer to [**XClusterTableType**](XClusterTableType.md) |  | [optional] 
**TableName** | Pointer to **string** | Table name. | [optional] [readonly] 
**RelationType** | Pointer to [**TableRelationType**](TableRelationType.md) |  | [optional] 
**SizeBytes** | Pointer to **float64** | SST size in bytes. | [optional] [readonly] 
**WalSizeBytes** | Pointer to **float64** | WAL size in bytes. | [optional] [readonly] 
**IndexTable** | Pointer to **bool** | Whether this row describes an index table. | [optional] [readonly] 
**IndexTableIds** | Pointer to **[]string** | Index table IDs when this is a main table. | [optional] [readonly] 
**PgSchemaName** | Pointer to **string** | PostgreSQL schema name for YSQL tables. | [optional] [readonly] 
**Colocated** | Pointer to **bool** | Whether the table is colocated. | [optional] [readonly] 

## Methods

### NewTableInfo

`func NewTableInfo() *TableInfo`

NewTableInfo instantiates a new TableInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableInfoWithDefaults

`func NewTableInfoWithDefaults() *TableInfo`

NewTableInfoWithDefaults instantiates a new TableInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableId

`func (o *TableInfo) GetTableId() string`

GetTableId returns the TableId field if non-nil, zero value otherwise.

### GetTableIdOk

`func (o *TableInfo) GetTableIdOk() (*string, bool)`

GetTableIdOk returns a tuple with the TableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableId

`func (o *TableInfo) SetTableId(v string)`

SetTableId sets TableId field to given value.

### HasTableId

`func (o *TableInfo) HasTableId() bool`

HasTableId returns a boolean if a field has been set.

### GetTableUuid

`func (o *TableInfo) GetTableUuid() string`

GetTableUuid returns the TableUuid field if non-nil, zero value otherwise.

### GetTableUuidOk

`func (o *TableInfo) GetTableUuidOk() (*string, bool)`

GetTableUuidOk returns a tuple with the TableUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableUuid

`func (o *TableInfo) SetTableUuid(v string)`

SetTableUuid sets TableUuid field to given value.

### HasTableUuid

`func (o *TableInfo) HasTableUuid() bool`

HasTableUuid returns a boolean if a field has been set.

### GetKeyspace

`func (o *TableInfo) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *TableInfo) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *TableInfo) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.

### HasKeyspace

`func (o *TableInfo) HasKeyspace() bool`

HasKeyspace returns a boolean if a field has been set.

### GetTableType

`func (o *TableInfo) GetTableType() XClusterTableType`

GetTableType returns the TableType field if non-nil, zero value otherwise.

### GetTableTypeOk

`func (o *TableInfo) GetTableTypeOk() (*XClusterTableType, bool)`

GetTableTypeOk returns a tuple with the TableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableType

`func (o *TableInfo) SetTableType(v XClusterTableType)`

SetTableType sets TableType field to given value.

### HasTableType

`func (o *TableInfo) HasTableType() bool`

HasTableType returns a boolean if a field has been set.

### GetTableName

`func (o *TableInfo) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *TableInfo) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *TableInfo) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *TableInfo) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetRelationType

`func (o *TableInfo) GetRelationType() TableRelationType`

GetRelationType returns the RelationType field if non-nil, zero value otherwise.

### GetRelationTypeOk

`func (o *TableInfo) GetRelationTypeOk() (*TableRelationType, bool)`

GetRelationTypeOk returns a tuple with the RelationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationType

`func (o *TableInfo) SetRelationType(v TableRelationType)`

SetRelationType sets RelationType field to given value.

### HasRelationType

`func (o *TableInfo) HasRelationType() bool`

HasRelationType returns a boolean if a field has been set.

### GetSizeBytes

`func (o *TableInfo) GetSizeBytes() float64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *TableInfo) GetSizeBytesOk() (*float64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *TableInfo) SetSizeBytes(v float64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *TableInfo) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetWalSizeBytes

`func (o *TableInfo) GetWalSizeBytes() float64`

GetWalSizeBytes returns the WalSizeBytes field if non-nil, zero value otherwise.

### GetWalSizeBytesOk

`func (o *TableInfo) GetWalSizeBytesOk() (*float64, bool)`

GetWalSizeBytesOk returns a tuple with the WalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalSizeBytes

`func (o *TableInfo) SetWalSizeBytes(v float64)`

SetWalSizeBytes sets WalSizeBytes field to given value.

### HasWalSizeBytes

`func (o *TableInfo) HasWalSizeBytes() bool`

HasWalSizeBytes returns a boolean if a field has been set.

### GetIndexTable

`func (o *TableInfo) GetIndexTable() bool`

GetIndexTable returns the IndexTable field if non-nil, zero value otherwise.

### GetIndexTableOk

`func (o *TableInfo) GetIndexTableOk() (*bool, bool)`

GetIndexTableOk returns a tuple with the IndexTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexTable

`func (o *TableInfo) SetIndexTable(v bool)`

SetIndexTable sets IndexTable field to given value.

### HasIndexTable

`func (o *TableInfo) HasIndexTable() bool`

HasIndexTable returns a boolean if a field has been set.

### GetIndexTableIds

`func (o *TableInfo) GetIndexTableIds() []string`

GetIndexTableIds returns the IndexTableIds field if non-nil, zero value otherwise.

### GetIndexTableIdsOk

`func (o *TableInfo) GetIndexTableIdsOk() (*[]string, bool)`

GetIndexTableIdsOk returns a tuple with the IndexTableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexTableIds

`func (o *TableInfo) SetIndexTableIds(v []string)`

SetIndexTableIds sets IndexTableIds field to given value.

### HasIndexTableIds

`func (o *TableInfo) HasIndexTableIds() bool`

HasIndexTableIds returns a boolean if a field has been set.

### GetPgSchemaName

`func (o *TableInfo) GetPgSchemaName() string`

GetPgSchemaName returns the PgSchemaName field if non-nil, zero value otherwise.

### GetPgSchemaNameOk

`func (o *TableInfo) GetPgSchemaNameOk() (*string, bool)`

GetPgSchemaNameOk returns a tuple with the PgSchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgSchemaName

`func (o *TableInfo) SetPgSchemaName(v string)`

SetPgSchemaName sets PgSchemaName field to given value.

### HasPgSchemaName

`func (o *TableInfo) HasPgSchemaName() bool`

HasPgSchemaName returns a boolean if a field has been set.

### GetColocated

`func (o *TableInfo) GetColocated() bool`

GetColocated returns the Colocated field if non-nil, zero value otherwise.

### GetColocatedOk

`func (o *TableInfo) GetColocatedOk() (*bool, bool)`

GetColocatedOk returns a tuple with the Colocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColocated

`func (o *TableInfo) SetColocated(v bool)`

SetColocated sets Colocated field to given value.

### HasColocated

`func (o *TableInfo) HasColocated() bool`

HasColocated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


