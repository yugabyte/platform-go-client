# ColumnDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnOrder** | Pointer to **int32** | Relative position (column order) for this column, in the table and in CQL commands | [optional] 
**IsClusteringKey** | Pointer to **bool** | True if this column is a clustering key | [optional] 
**IsPartitionKey** | Pointer to **bool** | True if this column is a partition key | [optional] 
**KeyType** | Pointer to **string** | Column key type | [optional] 
**Name** | Pointer to **string** | Column name | [optional] 
**SortOrder** | Pointer to **string** | Sort order for this column. Valid only for clustering columns. | [optional] 
**Type** | Pointer to **string** | The column&#39;s data type | [optional] 
**ValueType** | Pointer to **string** | Column value name | [optional] 

## Methods

### NewColumnDetails

`func NewColumnDetails() *ColumnDetails`

NewColumnDetails instantiates a new ColumnDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColumnDetailsWithDefaults

`func NewColumnDetailsWithDefaults() *ColumnDetails`

NewColumnDetailsWithDefaults instantiates a new ColumnDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnOrder

`func (o *ColumnDetails) GetColumnOrder() int32`

GetColumnOrder returns the ColumnOrder field if non-nil, zero value otherwise.

### GetColumnOrderOk

`func (o *ColumnDetails) GetColumnOrderOk() (*int32, bool)`

GetColumnOrderOk returns a tuple with the ColumnOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnOrder

`func (o *ColumnDetails) SetColumnOrder(v int32)`

SetColumnOrder sets ColumnOrder field to given value.

### HasColumnOrder

`func (o *ColumnDetails) HasColumnOrder() bool`

HasColumnOrder returns a boolean if a field has been set.

### GetIsClusteringKey

`func (o *ColumnDetails) GetIsClusteringKey() bool`

GetIsClusteringKey returns the IsClusteringKey field if non-nil, zero value otherwise.

### GetIsClusteringKeyOk

`func (o *ColumnDetails) GetIsClusteringKeyOk() (*bool, bool)`

GetIsClusteringKeyOk returns a tuple with the IsClusteringKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClusteringKey

`func (o *ColumnDetails) SetIsClusteringKey(v bool)`

SetIsClusteringKey sets IsClusteringKey field to given value.

### HasIsClusteringKey

`func (o *ColumnDetails) HasIsClusteringKey() bool`

HasIsClusteringKey returns a boolean if a field has been set.

### GetIsPartitionKey

`func (o *ColumnDetails) GetIsPartitionKey() bool`

GetIsPartitionKey returns the IsPartitionKey field if non-nil, zero value otherwise.

### GetIsPartitionKeyOk

`func (o *ColumnDetails) GetIsPartitionKeyOk() (*bool, bool)`

GetIsPartitionKeyOk returns a tuple with the IsPartitionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartitionKey

`func (o *ColumnDetails) SetIsPartitionKey(v bool)`

SetIsPartitionKey sets IsPartitionKey field to given value.

### HasIsPartitionKey

`func (o *ColumnDetails) HasIsPartitionKey() bool`

HasIsPartitionKey returns a boolean if a field has been set.

### GetKeyType

`func (o *ColumnDetails) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *ColumnDetails) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *ColumnDetails) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *ColumnDetails) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetName

`func (o *ColumnDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ColumnDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ColumnDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ColumnDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSortOrder

`func (o *ColumnDetails) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ColumnDetails) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ColumnDetails) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ColumnDetails) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetType

`func (o *ColumnDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ColumnDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ColumnDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ColumnDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValueType

`func (o *ColumnDetails) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *ColumnDetails) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *ColumnDetails) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *ColumnDetails) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


