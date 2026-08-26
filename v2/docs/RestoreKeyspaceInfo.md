# RestoreKeyspaceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Restore keyspace UUID. | [optional] [readonly] 
**RestoreUuid** | Pointer to **string** | Parent (universe-level) restore UUID. | [optional] [readonly] 
**SourceKeyspace** | Pointer to **string** | Source keyspace name. | [optional] [readonly] 
**TargetKeyspace** | Pointer to **string** | Target keyspace name. | [optional] [readonly] 
**StorageLocation** | Pointer to **string** | Storage location for the keyspace backup that was restored. | [optional] [readonly] 
**TableNameList** | Pointer to **[]string** | Tables restored into the keyspace. | [optional] [readonly] 
**State** | Pointer to [**RestoreState**](RestoreState.md) |  | [optional] 
**CreateTime** | Pointer to **time.Time** | Time when the keyspace restore was created. | [optional] [readonly] 
**CompleteTime** | Pointer to **time.Time** | Time when the keyspace restore completed. | [optional] [readonly] 

## Methods

### NewRestoreKeyspaceInfo

`func NewRestoreKeyspaceInfo() *RestoreKeyspaceInfo`

NewRestoreKeyspaceInfo instantiates a new RestoreKeyspaceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreKeyspaceInfoWithDefaults

`func NewRestoreKeyspaceInfoWithDefaults() *RestoreKeyspaceInfo`

NewRestoreKeyspaceInfoWithDefaults instantiates a new RestoreKeyspaceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *RestoreKeyspaceInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RestoreKeyspaceInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RestoreKeyspaceInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RestoreKeyspaceInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetRestoreUuid

`func (o *RestoreKeyspaceInfo) GetRestoreUuid() string`

GetRestoreUuid returns the RestoreUuid field if non-nil, zero value otherwise.

### GetRestoreUuidOk

`func (o *RestoreKeyspaceInfo) GetRestoreUuidOk() (*string, bool)`

GetRestoreUuidOk returns a tuple with the RestoreUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreUuid

`func (o *RestoreKeyspaceInfo) SetRestoreUuid(v string)`

SetRestoreUuid sets RestoreUuid field to given value.

### HasRestoreUuid

`func (o *RestoreKeyspaceInfo) HasRestoreUuid() bool`

HasRestoreUuid returns a boolean if a field has been set.

### GetSourceKeyspace

`func (o *RestoreKeyspaceInfo) GetSourceKeyspace() string`

GetSourceKeyspace returns the SourceKeyspace field if non-nil, zero value otherwise.

### GetSourceKeyspaceOk

`func (o *RestoreKeyspaceInfo) GetSourceKeyspaceOk() (*string, bool)`

GetSourceKeyspaceOk returns a tuple with the SourceKeyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceKeyspace

`func (o *RestoreKeyspaceInfo) SetSourceKeyspace(v string)`

SetSourceKeyspace sets SourceKeyspace field to given value.

### HasSourceKeyspace

`func (o *RestoreKeyspaceInfo) HasSourceKeyspace() bool`

HasSourceKeyspace returns a boolean if a field has been set.

### GetTargetKeyspace

`func (o *RestoreKeyspaceInfo) GetTargetKeyspace() string`

GetTargetKeyspace returns the TargetKeyspace field if non-nil, zero value otherwise.

### GetTargetKeyspaceOk

`func (o *RestoreKeyspaceInfo) GetTargetKeyspaceOk() (*string, bool)`

GetTargetKeyspaceOk returns a tuple with the TargetKeyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetKeyspace

`func (o *RestoreKeyspaceInfo) SetTargetKeyspace(v string)`

SetTargetKeyspace sets TargetKeyspace field to given value.

### HasTargetKeyspace

`func (o *RestoreKeyspaceInfo) HasTargetKeyspace() bool`

HasTargetKeyspace returns a boolean if a field has been set.

### GetStorageLocation

`func (o *RestoreKeyspaceInfo) GetStorageLocation() string`

GetStorageLocation returns the StorageLocation field if non-nil, zero value otherwise.

### GetStorageLocationOk

`func (o *RestoreKeyspaceInfo) GetStorageLocationOk() (*string, bool)`

GetStorageLocationOk returns a tuple with the StorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLocation

`func (o *RestoreKeyspaceInfo) SetStorageLocation(v string)`

SetStorageLocation sets StorageLocation field to given value.

### HasStorageLocation

`func (o *RestoreKeyspaceInfo) HasStorageLocation() bool`

HasStorageLocation returns a boolean if a field has been set.

### GetTableNameList

`func (o *RestoreKeyspaceInfo) GetTableNameList() []string`

GetTableNameList returns the TableNameList field if non-nil, zero value otherwise.

### GetTableNameListOk

`func (o *RestoreKeyspaceInfo) GetTableNameListOk() (*[]string, bool)`

GetTableNameListOk returns a tuple with the TableNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableNameList

`func (o *RestoreKeyspaceInfo) SetTableNameList(v []string)`

SetTableNameList sets TableNameList field to given value.

### HasTableNameList

`func (o *RestoreKeyspaceInfo) HasTableNameList() bool`

HasTableNameList returns a boolean if a field has been set.

### GetState

`func (o *RestoreKeyspaceInfo) GetState() RestoreState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RestoreKeyspaceInfo) GetStateOk() (*RestoreState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RestoreKeyspaceInfo) SetState(v RestoreState)`

SetState sets State field to given value.

### HasState

`func (o *RestoreKeyspaceInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCreateTime

`func (o *RestoreKeyspaceInfo) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *RestoreKeyspaceInfo) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *RestoreKeyspaceInfo) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *RestoreKeyspaceInfo) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCompleteTime

`func (o *RestoreKeyspaceInfo) GetCompleteTime() time.Time`

GetCompleteTime returns the CompleteTime field if non-nil, zero value otherwise.

### GetCompleteTimeOk

`func (o *RestoreKeyspaceInfo) GetCompleteTimeOk() (*time.Time, bool)`

GetCompleteTimeOk returns a tuple with the CompleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteTime

`func (o *RestoreKeyspaceInfo) SetCompleteTime(v time.Time)`

SetCompleteTime sets CompleteTime field to given value.

### HasCompleteTime

`func (o *RestoreKeyspaceInfo) HasCompleteTime() bool`

HasCompleteTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


