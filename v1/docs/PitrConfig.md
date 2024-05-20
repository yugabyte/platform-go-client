# PitrConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **time.Time** | Create time of the PITR config | [optional] [readonly] 
**CreatedForDr** | Pointer to **bool** | Created for DR | [optional] [readonly] 
**CustomerUUID** | Pointer to **string** | Customer UUID of this config | [optional] 
**DbName** | Pointer to **string** | DB Name | [optional] 
**MaxRecoverTimeInMillis** | **int64** |  | 
**MinRecoverTimeInMillis** | **int64** |  | 
**Name** | Pointer to **string** | PITR config name | [optional] 
**RetentionPeriod** | Pointer to **int64** | Retention Period in seconds | [optional] 
**ScheduleInterval** | Pointer to **int64** | Interval between snasphots in seconds | [optional] 
**State** | **string** |  | 
**TableType** | Pointer to **string** | Table Type | [optional] 
**UpdateTime** | Pointer to **time.Time** | Update time of the PITR con | [optional] 
**UsedForXCluster** | Pointer to **bool** |  | [optional] [readonly] 
**Uuid** | Pointer to **string** | PITR config UUID | [optional] 

## Methods

### NewPitrConfig

`func NewPitrConfig(maxRecoverTimeInMillis int64, minRecoverTimeInMillis int64, state string, ) *PitrConfig`

NewPitrConfig instantiates a new PitrConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPitrConfigWithDefaults

`func NewPitrConfigWithDefaults() *PitrConfig`

NewPitrConfigWithDefaults instantiates a new PitrConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *PitrConfig) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *PitrConfig) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *PitrConfig) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *PitrConfig) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreatedForDr

`func (o *PitrConfig) GetCreatedForDr() bool`

GetCreatedForDr returns the CreatedForDr field if non-nil, zero value otherwise.

### GetCreatedForDrOk

`func (o *PitrConfig) GetCreatedForDrOk() (*bool, bool)`

GetCreatedForDrOk returns a tuple with the CreatedForDr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedForDr

`func (o *PitrConfig) SetCreatedForDr(v bool)`

SetCreatedForDr sets CreatedForDr field to given value.

### HasCreatedForDr

`func (o *PitrConfig) HasCreatedForDr() bool`

HasCreatedForDr returns a boolean if a field has been set.

### GetCustomerUUID

`func (o *PitrConfig) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *PitrConfig) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *PitrConfig) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.

### HasCustomerUUID

`func (o *PitrConfig) HasCustomerUUID() bool`

HasCustomerUUID returns a boolean if a field has been set.

### GetDbName

`func (o *PitrConfig) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *PitrConfig) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *PitrConfig) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *PitrConfig) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetMaxRecoverTimeInMillis

`func (o *PitrConfig) GetMaxRecoverTimeInMillis() int64`

GetMaxRecoverTimeInMillis returns the MaxRecoverTimeInMillis field if non-nil, zero value otherwise.

### GetMaxRecoverTimeInMillisOk

`func (o *PitrConfig) GetMaxRecoverTimeInMillisOk() (*int64, bool)`

GetMaxRecoverTimeInMillisOk returns a tuple with the MaxRecoverTimeInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRecoverTimeInMillis

`func (o *PitrConfig) SetMaxRecoverTimeInMillis(v int64)`

SetMaxRecoverTimeInMillis sets MaxRecoverTimeInMillis field to given value.


### GetMinRecoverTimeInMillis

`func (o *PitrConfig) GetMinRecoverTimeInMillis() int64`

GetMinRecoverTimeInMillis returns the MinRecoverTimeInMillis field if non-nil, zero value otherwise.

### GetMinRecoverTimeInMillisOk

`func (o *PitrConfig) GetMinRecoverTimeInMillisOk() (*int64, bool)`

GetMinRecoverTimeInMillisOk returns a tuple with the MinRecoverTimeInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRecoverTimeInMillis

`func (o *PitrConfig) SetMinRecoverTimeInMillis(v int64)`

SetMinRecoverTimeInMillis sets MinRecoverTimeInMillis field to given value.


### GetName

`func (o *PitrConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PitrConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PitrConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PitrConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetentionPeriod

`func (o *PitrConfig) GetRetentionPeriod() int64`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *PitrConfig) GetRetentionPeriodOk() (*int64, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *PitrConfig) SetRetentionPeriod(v int64)`

SetRetentionPeriod sets RetentionPeriod field to given value.

### HasRetentionPeriod

`func (o *PitrConfig) HasRetentionPeriod() bool`

HasRetentionPeriod returns a boolean if a field has been set.

### GetScheduleInterval

`func (o *PitrConfig) GetScheduleInterval() int64`

GetScheduleInterval returns the ScheduleInterval field if non-nil, zero value otherwise.

### GetScheduleIntervalOk

`func (o *PitrConfig) GetScheduleIntervalOk() (*int64, bool)`

GetScheduleIntervalOk returns a tuple with the ScheduleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleInterval

`func (o *PitrConfig) SetScheduleInterval(v int64)`

SetScheduleInterval sets ScheduleInterval field to given value.

### HasScheduleInterval

`func (o *PitrConfig) HasScheduleInterval() bool`

HasScheduleInterval returns a boolean if a field has been set.

### GetState

`func (o *PitrConfig) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PitrConfig) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PitrConfig) SetState(v string)`

SetState sets State field to given value.


### GetTableType

`func (o *PitrConfig) GetTableType() string`

GetTableType returns the TableType field if non-nil, zero value otherwise.

### GetTableTypeOk

`func (o *PitrConfig) GetTableTypeOk() (*string, bool)`

GetTableTypeOk returns a tuple with the TableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableType

`func (o *PitrConfig) SetTableType(v string)`

SetTableType sets TableType field to given value.

### HasTableType

`func (o *PitrConfig) HasTableType() bool`

HasTableType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *PitrConfig) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *PitrConfig) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *PitrConfig) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *PitrConfig) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetUsedForXCluster

`func (o *PitrConfig) GetUsedForXCluster() bool`

GetUsedForXCluster returns the UsedForXCluster field if non-nil, zero value otherwise.

### GetUsedForXClusterOk

`func (o *PitrConfig) GetUsedForXClusterOk() (*bool, bool)`

GetUsedForXClusterOk returns a tuple with the UsedForXCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedForXCluster

`func (o *PitrConfig) SetUsedForXCluster(v bool)`

SetUsedForXCluster sets UsedForXCluster field to given value.

### HasUsedForXCluster

`func (o *PitrConfig) HasUsedForXCluster() bool`

HasUsedForXCluster returns a boolean if a field has been set.

### GetUuid

`func (o *PitrConfig) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PitrConfig) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PitrConfig) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PitrConfig) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


