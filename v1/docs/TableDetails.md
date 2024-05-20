# TableDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to [**[]ColumnDetails**](ColumnDetails.md) | Details of all columns in the table | [optional] 
**Keyspace** | Pointer to **string** | Keyspace to which this table belongs | [optional] 
**SplitValues** | Pointer to **[]string** | Primary key values used to split table into tablets | [optional] 
**TableName** | Pointer to **string** | Table name | [optional] 
**TtlInSeconds** | Pointer to **int64** | The default table-level time to live, in seconds. A value of &#x60;-1&#x60; represents an infinite TTL. | [optional] 

## Methods

### NewTableDetails

`func NewTableDetails() *TableDetails`

NewTableDetails instantiates a new TableDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableDetailsWithDefaults

`func NewTableDetailsWithDefaults() *TableDetails`

NewTableDetailsWithDefaults instantiates a new TableDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *TableDetails) GetColumns() []ColumnDetails`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TableDetails) GetColumnsOk() (*[]ColumnDetails, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TableDetails) SetColumns(v []ColumnDetails)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *TableDetails) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetKeyspace

`func (o *TableDetails) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *TableDetails) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *TableDetails) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.

### HasKeyspace

`func (o *TableDetails) HasKeyspace() bool`

HasKeyspace returns a boolean if a field has been set.

### GetSplitValues

`func (o *TableDetails) GetSplitValues() []string`

GetSplitValues returns the SplitValues field if non-nil, zero value otherwise.

### GetSplitValuesOk

`func (o *TableDetails) GetSplitValuesOk() (*[]string, bool)`

GetSplitValuesOk returns a tuple with the SplitValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitValues

`func (o *TableDetails) SetSplitValues(v []string)`

SetSplitValues sets SplitValues field to given value.

### HasSplitValues

`func (o *TableDetails) HasSplitValues() bool`

HasSplitValues returns a boolean if a field has been set.

### GetTableName

`func (o *TableDetails) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *TableDetails) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *TableDetails) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *TableDetails) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetTtlInSeconds

`func (o *TableDetails) GetTtlInSeconds() int64`

GetTtlInSeconds returns the TtlInSeconds field if non-nil, zero value otherwise.

### GetTtlInSecondsOk

`func (o *TableDetails) GetTtlInSecondsOk() (*int64, bool)`

GetTtlInSecondsOk returns a tuple with the TtlInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlInSeconds

`func (o *TableDetails) SetTtlInSeconds(v int64)`

SetTtlInSeconds sets TtlInSeconds field to given value.

### HasTtlInSeconds

`func (o *TableDetails) HasTtlInSeconds() bool`

HasTtlInSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


