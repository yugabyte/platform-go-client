# UniverseSoftwareUpgradeStart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SleepAfterMasterRestartMillis** | Pointer to **int32** | Applicable for rolling restarts. Time to wait between master restarts. Defaults to 180000. | [optional] [default to 180000]
**SleepAfterTserverRestartMillis** | Pointer to **int32** | Applicable for rolling restarts. Time to wait between tserver restarts. Defaults to 180000. | [optional] [default to 180000]
**RollingUpgrade** | Pointer to **bool** | Perform a rolling upgrade where only one node is upgraded at a time. This is the default behavior. False will perform a non-rolling upgrade where all nodes are upgraded at the same  | [optional] [default to true]
**RollMaxBatchSize** | Pointer to [**RollMaxBatchSize**](RollMaxBatchSize.md) |  | [optional] 
**AllowRollback** | Pointer to **bool** | perform an upgrade where rollback is allowed | [optional] [default to true]
**UpgradeSystemCatalog** | Pointer to **bool** | Upgrade the YugabyteDB Catalog | [optional] [default to true]
**Version** | **string** | The target release version to upgrade to. | 

## Methods

### NewUniverseSoftwareUpgradeStart

`func NewUniverseSoftwareUpgradeStart(version string, ) *UniverseSoftwareUpgradeStart`

NewUniverseSoftwareUpgradeStart instantiates a new UniverseSoftwareUpgradeStart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseSoftwareUpgradeStartWithDefaults

`func NewUniverseSoftwareUpgradeStartWithDefaults() *UniverseSoftwareUpgradeStart`

NewUniverseSoftwareUpgradeStartWithDefaults instantiates a new UniverseSoftwareUpgradeStart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSleepAfterMasterRestartMillis

`func (o *UniverseSoftwareUpgradeStart) GetSleepAfterMasterRestartMillis() int32`

GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterMasterRestartMillisOk

`func (o *UniverseSoftwareUpgradeStart) GetSleepAfterMasterRestartMillisOk() (*int32, bool)`

GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterMasterRestartMillis

`func (o *UniverseSoftwareUpgradeStart) SetSleepAfterMasterRestartMillis(v int32)`

SetSleepAfterMasterRestartMillis sets SleepAfterMasterRestartMillis field to given value.

### HasSleepAfterMasterRestartMillis

`func (o *UniverseSoftwareUpgradeStart) HasSleepAfterMasterRestartMillis() bool`

HasSleepAfterMasterRestartMillis returns a boolean if a field has been set.

### GetSleepAfterTserverRestartMillis

`func (o *UniverseSoftwareUpgradeStart) GetSleepAfterTserverRestartMillis() int32`

GetSleepAfterTserverRestartMillis returns the SleepAfterTserverRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterTserverRestartMillisOk

`func (o *UniverseSoftwareUpgradeStart) GetSleepAfterTserverRestartMillisOk() (*int32, bool)`

GetSleepAfterTserverRestartMillisOk returns a tuple with the SleepAfterTserverRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterTserverRestartMillis

`func (o *UniverseSoftwareUpgradeStart) SetSleepAfterTserverRestartMillis(v int32)`

SetSleepAfterTserverRestartMillis sets SleepAfterTserverRestartMillis field to given value.

### HasSleepAfterTserverRestartMillis

`func (o *UniverseSoftwareUpgradeStart) HasSleepAfterTserverRestartMillis() bool`

HasSleepAfterTserverRestartMillis returns a boolean if a field has been set.

### GetRollingUpgrade

`func (o *UniverseSoftwareUpgradeStart) GetRollingUpgrade() bool`

GetRollingUpgrade returns the RollingUpgrade field if non-nil, zero value otherwise.

### GetRollingUpgradeOk

`func (o *UniverseSoftwareUpgradeStart) GetRollingUpgradeOk() (*bool, bool)`

GetRollingUpgradeOk returns a tuple with the RollingUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingUpgrade

`func (o *UniverseSoftwareUpgradeStart) SetRollingUpgrade(v bool)`

SetRollingUpgrade sets RollingUpgrade field to given value.

### HasRollingUpgrade

`func (o *UniverseSoftwareUpgradeStart) HasRollingUpgrade() bool`

HasRollingUpgrade returns a boolean if a field has been set.

### GetRollMaxBatchSize

`func (o *UniverseSoftwareUpgradeStart) GetRollMaxBatchSize() RollMaxBatchSize`

GetRollMaxBatchSize returns the RollMaxBatchSize field if non-nil, zero value otherwise.

### GetRollMaxBatchSizeOk

`func (o *UniverseSoftwareUpgradeStart) GetRollMaxBatchSizeOk() (*RollMaxBatchSize, bool)`

GetRollMaxBatchSizeOk returns a tuple with the RollMaxBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollMaxBatchSize

`func (o *UniverseSoftwareUpgradeStart) SetRollMaxBatchSize(v RollMaxBatchSize)`

SetRollMaxBatchSize sets RollMaxBatchSize field to given value.

### HasRollMaxBatchSize

`func (o *UniverseSoftwareUpgradeStart) HasRollMaxBatchSize() bool`

HasRollMaxBatchSize returns a boolean if a field has been set.

### GetAllowRollback

`func (o *UniverseSoftwareUpgradeStart) GetAllowRollback() bool`

GetAllowRollback returns the AllowRollback field if non-nil, zero value otherwise.

### GetAllowRollbackOk

`func (o *UniverseSoftwareUpgradeStart) GetAllowRollbackOk() (*bool, bool)`

GetAllowRollbackOk returns a tuple with the AllowRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRollback

`func (o *UniverseSoftwareUpgradeStart) SetAllowRollback(v bool)`

SetAllowRollback sets AllowRollback field to given value.

### HasAllowRollback

`func (o *UniverseSoftwareUpgradeStart) HasAllowRollback() bool`

HasAllowRollback returns a boolean if a field has been set.

### GetUpgradeSystemCatalog

`func (o *UniverseSoftwareUpgradeStart) GetUpgradeSystemCatalog() bool`

GetUpgradeSystemCatalog returns the UpgradeSystemCatalog field if non-nil, zero value otherwise.

### GetUpgradeSystemCatalogOk

`func (o *UniverseSoftwareUpgradeStart) GetUpgradeSystemCatalogOk() (*bool, bool)`

GetUpgradeSystemCatalogOk returns a tuple with the UpgradeSystemCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeSystemCatalog

`func (o *UniverseSoftwareUpgradeStart) SetUpgradeSystemCatalog(v bool)`

SetUpgradeSystemCatalog sets UpgradeSystemCatalog field to given value.

### HasUpgradeSystemCatalog

`func (o *UniverseSoftwareUpgradeStart) HasUpgradeSystemCatalog() bool`

HasUpgradeSystemCatalog returns a boolean if a field has been set.

### GetVersion

`func (o *UniverseSoftwareUpgradeStart) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UniverseSoftwareUpgradeStart) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UniverseSoftwareUpgradeStart) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


