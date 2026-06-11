# PitrConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | PITR config UUID | [optional] [readonly] 
**CustomerUuid** | Pointer to **string** | Customer UUID of this config | [optional] [readonly] 
**MinRecoverTimeInMillis** | **int64** | Earliest point in time (UTC epoch millis) that can be restored to. | [readonly] 
**MaxRecoverTimeInMillis** | **int64** | Latest point in time (UTC epoch millis) that can be restored to. | [readonly] 
**State** | **string** | Snapshot schedule state from the universe master. | [readonly] 
**CreateTime** | Pointer to **time.Time** | Time when the PITR config was created. | [optional] [readonly] 
**UpdateTime** | Pointer to **time.Time** | Time when the PITR config was last updated. | [optional] [readonly] 
**CreatedForDr** | Pointer to **bool** | Whether this PITR config was created for disaster recovery. | [optional] [readonly] 
**UsedForXCluster** | Pointer to **bool** | Whether this PITR config is used by transactional xCluster replication. | [optional] [readonly] 
**Disabled** | Pointer to **bool** | Whether the PITR config is disabled. | [optional] [readonly] 

## Methods

### NewPitrConfigInfo

`func NewPitrConfigInfo(minRecoverTimeInMillis int64, maxRecoverTimeInMillis int64, state string, ) *PitrConfigInfo`

NewPitrConfigInfo instantiates a new PitrConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPitrConfigInfoWithDefaults

`func NewPitrConfigInfoWithDefaults() *PitrConfigInfo`

NewPitrConfigInfoWithDefaults instantiates a new PitrConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PitrConfigInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PitrConfigInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PitrConfigInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PitrConfigInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCustomerUuid

`func (o *PitrConfigInfo) GetCustomerUuid() string`

GetCustomerUuid returns the CustomerUuid field if non-nil, zero value otherwise.

### GetCustomerUuidOk

`func (o *PitrConfigInfo) GetCustomerUuidOk() (*string, bool)`

GetCustomerUuidOk returns a tuple with the CustomerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUuid

`func (o *PitrConfigInfo) SetCustomerUuid(v string)`

SetCustomerUuid sets CustomerUuid field to given value.

### HasCustomerUuid

`func (o *PitrConfigInfo) HasCustomerUuid() bool`

HasCustomerUuid returns a boolean if a field has been set.

### GetMinRecoverTimeInMillis

`func (o *PitrConfigInfo) GetMinRecoverTimeInMillis() int64`

GetMinRecoverTimeInMillis returns the MinRecoverTimeInMillis field if non-nil, zero value otherwise.

### GetMinRecoverTimeInMillisOk

`func (o *PitrConfigInfo) GetMinRecoverTimeInMillisOk() (*int64, bool)`

GetMinRecoverTimeInMillisOk returns a tuple with the MinRecoverTimeInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRecoverTimeInMillis

`func (o *PitrConfigInfo) SetMinRecoverTimeInMillis(v int64)`

SetMinRecoverTimeInMillis sets MinRecoverTimeInMillis field to given value.


### GetMaxRecoverTimeInMillis

`func (o *PitrConfigInfo) GetMaxRecoverTimeInMillis() int64`

GetMaxRecoverTimeInMillis returns the MaxRecoverTimeInMillis field if non-nil, zero value otherwise.

### GetMaxRecoverTimeInMillisOk

`func (o *PitrConfigInfo) GetMaxRecoverTimeInMillisOk() (*int64, bool)`

GetMaxRecoverTimeInMillisOk returns a tuple with the MaxRecoverTimeInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRecoverTimeInMillis

`func (o *PitrConfigInfo) SetMaxRecoverTimeInMillis(v int64)`

SetMaxRecoverTimeInMillis sets MaxRecoverTimeInMillis field to given value.


### GetState

`func (o *PitrConfigInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PitrConfigInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PitrConfigInfo) SetState(v string)`

SetState sets State field to given value.


### GetCreateTime

`func (o *PitrConfigInfo) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *PitrConfigInfo) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *PitrConfigInfo) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *PitrConfigInfo) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *PitrConfigInfo) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *PitrConfigInfo) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *PitrConfigInfo) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *PitrConfigInfo) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetCreatedForDr

`func (o *PitrConfigInfo) GetCreatedForDr() bool`

GetCreatedForDr returns the CreatedForDr field if non-nil, zero value otherwise.

### GetCreatedForDrOk

`func (o *PitrConfigInfo) GetCreatedForDrOk() (*bool, bool)`

GetCreatedForDrOk returns a tuple with the CreatedForDr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedForDr

`func (o *PitrConfigInfo) SetCreatedForDr(v bool)`

SetCreatedForDr sets CreatedForDr field to given value.

### HasCreatedForDr

`func (o *PitrConfigInfo) HasCreatedForDr() bool`

HasCreatedForDr returns a boolean if a field has been set.

### GetUsedForXCluster

`func (o *PitrConfigInfo) GetUsedForXCluster() bool`

GetUsedForXCluster returns the UsedForXCluster field if non-nil, zero value otherwise.

### GetUsedForXClusterOk

`func (o *PitrConfigInfo) GetUsedForXClusterOk() (*bool, bool)`

GetUsedForXClusterOk returns a tuple with the UsedForXCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedForXCluster

`func (o *PitrConfigInfo) SetUsedForXCluster(v bool)`

SetUsedForXCluster sets UsedForXCluster field to given value.

### HasUsedForXCluster

`func (o *PitrConfigInfo) HasUsedForXCluster() bool`

HasUsedForXCluster returns a boolean if a field has been set.

### GetDisabled

`func (o *PitrConfigInfo) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *PitrConfigInfo) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *PitrConfigInfo) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *PitrConfigInfo) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


