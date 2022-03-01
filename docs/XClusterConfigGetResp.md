# XClusterConfigGetResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **time.Time** | Create time | [optional] 
**Lag** | **map[string]interface{}** | Lag metric data | 
**ModifyTime** | Pointer to **time.Time** | Modify time | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**SourceUniverseUUID** | Pointer to **string** | Source Universe UUID | [optional] 
**Status** | Pointer to **string** | Status | [optional] 
**Tables** | Pointer to **[]string** | Source Universe table IDs | [optional] 
**TargetUniverseUUID** | Pointer to **string** | Target Universe UUID | [optional] 
**Uuid** | Pointer to **string** | UUID | [optional] 
**XclusterConfig** | [**XClusterConfig**](XClusterConfig.md) |  | 

## Methods

### NewXClusterConfigGetResp

`func NewXClusterConfigGetResp(lag map[string]interface{}, xclusterConfig XClusterConfig, ) *XClusterConfigGetResp`

NewXClusterConfigGetResp instantiates a new XClusterConfigGetResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXClusterConfigGetRespWithDefaults

`func NewXClusterConfigGetRespWithDefaults() *XClusterConfigGetResp`

NewXClusterConfigGetRespWithDefaults instantiates a new XClusterConfigGetResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *XClusterConfigGetResp) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *XClusterConfigGetResp) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *XClusterConfigGetResp) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *XClusterConfigGetResp) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetLag

`func (o *XClusterConfigGetResp) GetLag() map[string]interface{}`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *XClusterConfigGetResp) GetLagOk() (*map[string]interface{}, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *XClusterConfigGetResp) SetLag(v map[string]interface{})`

SetLag sets Lag field to given value.


### GetModifyTime

`func (o *XClusterConfigGetResp) GetModifyTime() time.Time`

GetModifyTime returns the ModifyTime field if non-nil, zero value otherwise.

### GetModifyTimeOk

`func (o *XClusterConfigGetResp) GetModifyTimeOk() (*time.Time, bool)`

GetModifyTimeOk returns a tuple with the ModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyTime

`func (o *XClusterConfigGetResp) SetModifyTime(v time.Time)`

SetModifyTime sets ModifyTime field to given value.

### HasModifyTime

`func (o *XClusterConfigGetResp) HasModifyTime() bool`

HasModifyTime returns a boolean if a field has been set.

### GetName

`func (o *XClusterConfigGetResp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *XClusterConfigGetResp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *XClusterConfigGetResp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *XClusterConfigGetResp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceUniverseUUID

`func (o *XClusterConfigGetResp) GetSourceUniverseUUID() string`

GetSourceUniverseUUID returns the SourceUniverseUUID field if non-nil, zero value otherwise.

### GetSourceUniverseUUIDOk

`func (o *XClusterConfigGetResp) GetSourceUniverseUUIDOk() (*string, bool)`

GetSourceUniverseUUIDOk returns a tuple with the SourceUniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUniverseUUID

`func (o *XClusterConfigGetResp) SetSourceUniverseUUID(v string)`

SetSourceUniverseUUID sets SourceUniverseUUID field to given value.

### HasSourceUniverseUUID

`func (o *XClusterConfigGetResp) HasSourceUniverseUUID() bool`

HasSourceUniverseUUID returns a boolean if a field has been set.

### GetStatus

`func (o *XClusterConfigGetResp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *XClusterConfigGetResp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *XClusterConfigGetResp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *XClusterConfigGetResp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTables

`func (o *XClusterConfigGetResp) GetTables() []string`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *XClusterConfigGetResp) GetTablesOk() (*[]string, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *XClusterConfigGetResp) SetTables(v []string)`

SetTables sets Tables field to given value.

### HasTables

`func (o *XClusterConfigGetResp) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetTargetUniverseUUID

`func (o *XClusterConfigGetResp) GetTargetUniverseUUID() string`

GetTargetUniverseUUID returns the TargetUniverseUUID field if non-nil, zero value otherwise.

### GetTargetUniverseUUIDOk

`func (o *XClusterConfigGetResp) GetTargetUniverseUUIDOk() (*string, bool)`

GetTargetUniverseUUIDOk returns a tuple with the TargetUniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUniverseUUID

`func (o *XClusterConfigGetResp) SetTargetUniverseUUID(v string)`

SetTargetUniverseUUID sets TargetUniverseUUID field to given value.

### HasTargetUniverseUUID

`func (o *XClusterConfigGetResp) HasTargetUniverseUUID() bool`

HasTargetUniverseUUID returns a boolean if a field has been set.

### GetUuid

`func (o *XClusterConfigGetResp) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *XClusterConfigGetResp) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *XClusterConfigGetResp) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *XClusterConfigGetResp) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetXclusterConfig

`func (o *XClusterConfigGetResp) GetXclusterConfig() XClusterConfig`

GetXclusterConfig returns the XclusterConfig field if non-nil, zero value otherwise.

### GetXclusterConfigOk

`func (o *XClusterConfigGetResp) GetXclusterConfigOk() (*XClusterConfig, bool)`

GetXclusterConfigOk returns a tuple with the XclusterConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXclusterConfig

`func (o *XClusterConfigGetResp) SetXclusterConfig(v XClusterConfig)`

SetXclusterConfig sets XclusterConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


