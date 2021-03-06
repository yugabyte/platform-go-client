# XClusterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **time.Time** | Create time | [optional] 
**ModifyTime** | Pointer to **time.Time** | Modify time | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**SourceUniverseUUID** | Pointer to **string** | Source Universe UUID | [optional] 
**Status** | Pointer to **string** | Status | [optional] 
**Tables** | Pointer to **[]string** | Source Universe table IDs | [optional] 
**TargetUniverseUUID** | Pointer to **string** | Target Universe UUID | [optional] 
**Uuid** | Pointer to **string** | UUID | [optional] 

## Methods

### NewXClusterConfig

`func NewXClusterConfig() *XClusterConfig`

NewXClusterConfig instantiates a new XClusterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXClusterConfigWithDefaults

`func NewXClusterConfigWithDefaults() *XClusterConfig`

NewXClusterConfigWithDefaults instantiates a new XClusterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *XClusterConfig) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *XClusterConfig) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *XClusterConfig) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *XClusterConfig) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetModifyTime

`func (o *XClusterConfig) GetModifyTime() time.Time`

GetModifyTime returns the ModifyTime field if non-nil, zero value otherwise.

### GetModifyTimeOk

`func (o *XClusterConfig) GetModifyTimeOk() (*time.Time, bool)`

GetModifyTimeOk returns a tuple with the ModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyTime

`func (o *XClusterConfig) SetModifyTime(v time.Time)`

SetModifyTime sets ModifyTime field to given value.

### HasModifyTime

`func (o *XClusterConfig) HasModifyTime() bool`

HasModifyTime returns a boolean if a field has been set.

### GetName

`func (o *XClusterConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *XClusterConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *XClusterConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *XClusterConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceUniverseUUID

`func (o *XClusterConfig) GetSourceUniverseUUID() string`

GetSourceUniverseUUID returns the SourceUniverseUUID field if non-nil, zero value otherwise.

### GetSourceUniverseUUIDOk

`func (o *XClusterConfig) GetSourceUniverseUUIDOk() (*string, bool)`

GetSourceUniverseUUIDOk returns a tuple with the SourceUniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUniverseUUID

`func (o *XClusterConfig) SetSourceUniverseUUID(v string)`

SetSourceUniverseUUID sets SourceUniverseUUID field to given value.

### HasSourceUniverseUUID

`func (o *XClusterConfig) HasSourceUniverseUUID() bool`

HasSourceUniverseUUID returns a boolean if a field has been set.

### GetStatus

`func (o *XClusterConfig) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *XClusterConfig) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *XClusterConfig) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *XClusterConfig) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTables

`func (o *XClusterConfig) GetTables() []string`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *XClusterConfig) GetTablesOk() (*[]string, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *XClusterConfig) SetTables(v []string)`

SetTables sets Tables field to given value.

### HasTables

`func (o *XClusterConfig) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetTargetUniverseUUID

`func (o *XClusterConfig) GetTargetUniverseUUID() string`

GetTargetUniverseUUID returns the TargetUniverseUUID field if non-nil, zero value otherwise.

### GetTargetUniverseUUIDOk

`func (o *XClusterConfig) GetTargetUniverseUUIDOk() (*string, bool)`

GetTargetUniverseUUIDOk returns a tuple with the TargetUniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUniverseUUID

`func (o *XClusterConfig) SetTargetUniverseUUID(v string)`

SetTargetUniverseUUID sets TargetUniverseUUID field to given value.

### HasTargetUniverseUUID

`func (o *XClusterConfig) HasTargetUniverseUUID() bool`

HasTargetUniverseUUID returns a boolean if a field has been set.

### GetUuid

`func (o *XClusterConfig) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *XClusterConfig) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *XClusterConfig) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *XClusterConfig) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


