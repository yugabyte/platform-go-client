# TableInfoResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeySpace** | Pointer to **string** | Keyspace | [optional] 
**RelationType** | Pointer to **string** | Relation type | [optional] 
**SizeBytes** | Pointer to **float64** | Size in bytes | [optional] [readonly] 
**TableName** | Pointer to **string** | Table name | [optional] 
**TableType** | Pointer to **string** | Table type | [optional] 
**TableUUID** | Pointer to **string** | Table UUID | [optional] [readonly] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


