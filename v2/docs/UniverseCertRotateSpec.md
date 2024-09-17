# UniverseCertRotateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SleepAfterMasterRestartMillis** | Pointer to **int32** | Applicable for rolling restarts. Time to wait between master restarts. Defaults to 180000. | [optional] [default to 180000]
**SleepAfterTserverRestartMillis** | Pointer to **int32** | Applicable for rolling restarts. Time to wait between tserver restarts. Defaults to 180000. | [optional] [default to 180000]
**RollingUpgrade** | Pointer to **bool** | Perform a rolling upgrade where only one node is upgraded at a time. This is the default behavior. False will perform a non-rolling upgrade where all nodes are upgraded at the same  | [optional] [default to true]
**RollMaxBatchSize** | Pointer to [**RollMaxBatchSize**](RollMaxBatchSize.md) |  | [optional] 
**SelfSignedServerCertRotate** | Pointer to **bool** |  | [optional] 
**SelfSignedClientCertRotate** | Pointer to **bool** |  | [optional] 
**RootCa** | Pointer to **string** |  | [optional] 
**ClientRootCa** | Pointer to **string** |  | [optional] 

## Methods

### NewUniverseCertRotateSpec

`func NewUniverseCertRotateSpec() *UniverseCertRotateSpec`

NewUniverseCertRotateSpec instantiates a new UniverseCertRotateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseCertRotateSpecWithDefaults

`func NewUniverseCertRotateSpecWithDefaults() *UniverseCertRotateSpec`

NewUniverseCertRotateSpecWithDefaults instantiates a new UniverseCertRotateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSleepAfterMasterRestartMillis

`func (o *UniverseCertRotateSpec) GetSleepAfterMasterRestartMillis() int32`

GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterMasterRestartMillisOk

`func (o *UniverseCertRotateSpec) GetSleepAfterMasterRestartMillisOk() (*int32, bool)`

GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterMasterRestartMillis

`func (o *UniverseCertRotateSpec) SetSleepAfterMasterRestartMillis(v int32)`

SetSleepAfterMasterRestartMillis sets SleepAfterMasterRestartMillis field to given value.

### HasSleepAfterMasterRestartMillis

`func (o *UniverseCertRotateSpec) HasSleepAfterMasterRestartMillis() bool`

HasSleepAfterMasterRestartMillis returns a boolean if a field has been set.

### GetSleepAfterTserverRestartMillis

`func (o *UniverseCertRotateSpec) GetSleepAfterTserverRestartMillis() int32`

GetSleepAfterTserverRestartMillis returns the SleepAfterTserverRestartMillis field if non-nil, zero value otherwise.

### GetSleepAfterTserverRestartMillisOk

`func (o *UniverseCertRotateSpec) GetSleepAfterTserverRestartMillisOk() (*int32, bool)`

GetSleepAfterTserverRestartMillisOk returns a tuple with the SleepAfterTserverRestartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepAfterTserverRestartMillis

`func (o *UniverseCertRotateSpec) SetSleepAfterTserverRestartMillis(v int32)`

SetSleepAfterTserverRestartMillis sets SleepAfterTserverRestartMillis field to given value.

### HasSleepAfterTserverRestartMillis

`func (o *UniverseCertRotateSpec) HasSleepAfterTserverRestartMillis() bool`

HasSleepAfterTserverRestartMillis returns a boolean if a field has been set.

### GetRollingUpgrade

`func (o *UniverseCertRotateSpec) GetRollingUpgrade() bool`

GetRollingUpgrade returns the RollingUpgrade field if non-nil, zero value otherwise.

### GetRollingUpgradeOk

`func (o *UniverseCertRotateSpec) GetRollingUpgradeOk() (*bool, bool)`

GetRollingUpgradeOk returns a tuple with the RollingUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingUpgrade

`func (o *UniverseCertRotateSpec) SetRollingUpgrade(v bool)`

SetRollingUpgrade sets RollingUpgrade field to given value.

### HasRollingUpgrade

`func (o *UniverseCertRotateSpec) HasRollingUpgrade() bool`

HasRollingUpgrade returns a boolean if a field has been set.

### GetRollMaxBatchSize

`func (o *UniverseCertRotateSpec) GetRollMaxBatchSize() RollMaxBatchSize`

GetRollMaxBatchSize returns the RollMaxBatchSize field if non-nil, zero value otherwise.

### GetRollMaxBatchSizeOk

`func (o *UniverseCertRotateSpec) GetRollMaxBatchSizeOk() (*RollMaxBatchSize, bool)`

GetRollMaxBatchSizeOk returns a tuple with the RollMaxBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollMaxBatchSize

`func (o *UniverseCertRotateSpec) SetRollMaxBatchSize(v RollMaxBatchSize)`

SetRollMaxBatchSize sets RollMaxBatchSize field to given value.

### HasRollMaxBatchSize

`func (o *UniverseCertRotateSpec) HasRollMaxBatchSize() bool`

HasRollMaxBatchSize returns a boolean if a field has been set.

### GetSelfSignedServerCertRotate

`func (o *UniverseCertRotateSpec) GetSelfSignedServerCertRotate() bool`

GetSelfSignedServerCertRotate returns the SelfSignedServerCertRotate field if non-nil, zero value otherwise.

### GetSelfSignedServerCertRotateOk

`func (o *UniverseCertRotateSpec) GetSelfSignedServerCertRotateOk() (*bool, bool)`

GetSelfSignedServerCertRotateOk returns a tuple with the SelfSignedServerCertRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfSignedServerCertRotate

`func (o *UniverseCertRotateSpec) SetSelfSignedServerCertRotate(v bool)`

SetSelfSignedServerCertRotate sets SelfSignedServerCertRotate field to given value.

### HasSelfSignedServerCertRotate

`func (o *UniverseCertRotateSpec) HasSelfSignedServerCertRotate() bool`

HasSelfSignedServerCertRotate returns a boolean if a field has been set.

### GetSelfSignedClientCertRotate

`func (o *UniverseCertRotateSpec) GetSelfSignedClientCertRotate() bool`

GetSelfSignedClientCertRotate returns the SelfSignedClientCertRotate field if non-nil, zero value otherwise.

### GetSelfSignedClientCertRotateOk

`func (o *UniverseCertRotateSpec) GetSelfSignedClientCertRotateOk() (*bool, bool)`

GetSelfSignedClientCertRotateOk returns a tuple with the SelfSignedClientCertRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfSignedClientCertRotate

`func (o *UniverseCertRotateSpec) SetSelfSignedClientCertRotate(v bool)`

SetSelfSignedClientCertRotate sets SelfSignedClientCertRotate field to given value.

### HasSelfSignedClientCertRotate

`func (o *UniverseCertRotateSpec) HasSelfSignedClientCertRotate() bool`

HasSelfSignedClientCertRotate returns a boolean if a field has been set.

### GetRootCa

`func (o *UniverseCertRotateSpec) GetRootCa() string`

GetRootCa returns the RootCa field if non-nil, zero value otherwise.

### GetRootCaOk

`func (o *UniverseCertRotateSpec) GetRootCaOk() (*string, bool)`

GetRootCaOk returns a tuple with the RootCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCa

`func (o *UniverseCertRotateSpec) SetRootCa(v string)`

SetRootCa sets RootCa field to given value.

### HasRootCa

`func (o *UniverseCertRotateSpec) HasRootCa() bool`

HasRootCa returns a boolean if a field has been set.

### GetClientRootCa

`func (o *UniverseCertRotateSpec) GetClientRootCa() string`

GetClientRootCa returns the ClientRootCa field if non-nil, zero value otherwise.

### GetClientRootCaOk

`func (o *UniverseCertRotateSpec) GetClientRootCaOk() (*string, bool)`

GetClientRootCaOk returns a tuple with the ClientRootCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRootCa

`func (o *UniverseCertRotateSpec) SetClientRootCa(v string)`

SetClientRootCa sets ClientRootCa field to given value.

### HasClientRootCa

`func (o *UniverseCertRotateSpec) HasClientRootCa() bool`

HasClientRootCa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


