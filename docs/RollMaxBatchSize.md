# RollMaxBatchSize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryBatchSize** | Pointer to **int32** | WARNING: This is a preview API that could change.Suggested number nodes for primary cluster | [optional] 
**ReadReplicaBatchSize** | Pointer to **int32** | WARNING: This is a preview API that could change.Suggested number nodes for readonly cluster | [optional] 

## Methods

### NewRollMaxBatchSize

`func NewRollMaxBatchSize() *RollMaxBatchSize`

NewRollMaxBatchSize instantiates a new RollMaxBatchSize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollMaxBatchSizeWithDefaults

`func NewRollMaxBatchSizeWithDefaults() *RollMaxBatchSize`

NewRollMaxBatchSizeWithDefaults instantiates a new RollMaxBatchSize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryBatchSize

`func (o *RollMaxBatchSize) GetPrimaryBatchSize() int32`

GetPrimaryBatchSize returns the PrimaryBatchSize field if non-nil, zero value otherwise.

### GetPrimaryBatchSizeOk

`func (o *RollMaxBatchSize) GetPrimaryBatchSizeOk() (*int32, bool)`

GetPrimaryBatchSizeOk returns a tuple with the PrimaryBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryBatchSize

`func (o *RollMaxBatchSize) SetPrimaryBatchSize(v int32)`

SetPrimaryBatchSize sets PrimaryBatchSize field to given value.

### HasPrimaryBatchSize

`func (o *RollMaxBatchSize) HasPrimaryBatchSize() bool`

HasPrimaryBatchSize returns a boolean if a field has been set.

### GetReadReplicaBatchSize

`func (o *RollMaxBatchSize) GetReadReplicaBatchSize() int32`

GetReadReplicaBatchSize returns the ReadReplicaBatchSize field if non-nil, zero value otherwise.

### GetReadReplicaBatchSizeOk

`func (o *RollMaxBatchSize) GetReadReplicaBatchSizeOk() (*int32, bool)`

GetReadReplicaBatchSizeOk returns a tuple with the ReadReplicaBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadReplicaBatchSize

`func (o *RollMaxBatchSize) SetReadReplicaBatchSize(v int32)`

SetReadReplicaBatchSize sets ReadReplicaBatchSize field to given value.

### HasReadReplicaBatchSize

`func (o *RollMaxBatchSize) HasReadReplicaBatchSize() bool`

HasReadReplicaBatchSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


