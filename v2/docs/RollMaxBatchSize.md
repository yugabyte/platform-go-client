# RollMaxBatchSize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryBatchSize** | Pointer to **float32** | Suggested number nodes for primary cluster. | [optional] [default to 1]
**ReadReplicaBatchSize** | Pointer to **float32** | Suggested number nodes for readonly cluster. | [optional] [default to 1]

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

`func (o *RollMaxBatchSize) GetPrimaryBatchSize() float32`

GetPrimaryBatchSize returns the PrimaryBatchSize field if non-nil, zero value otherwise.

### GetPrimaryBatchSizeOk

`func (o *RollMaxBatchSize) GetPrimaryBatchSizeOk() (*float32, bool)`

GetPrimaryBatchSizeOk returns a tuple with the PrimaryBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryBatchSize

`func (o *RollMaxBatchSize) SetPrimaryBatchSize(v float32)`

SetPrimaryBatchSize sets PrimaryBatchSize field to given value.

### HasPrimaryBatchSize

`func (o *RollMaxBatchSize) HasPrimaryBatchSize() bool`

HasPrimaryBatchSize returns a boolean if a field has been set.

### GetReadReplicaBatchSize

`func (o *RollMaxBatchSize) GetReadReplicaBatchSize() float32`

GetReadReplicaBatchSize returns the ReadReplicaBatchSize field if non-nil, zero value otherwise.

### GetReadReplicaBatchSizeOk

`func (o *RollMaxBatchSize) GetReadReplicaBatchSizeOk() (*float32, bool)`

GetReadReplicaBatchSizeOk returns a tuple with the ReadReplicaBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadReplicaBatchSize

`func (o *RollMaxBatchSize) SetReadReplicaBatchSize(v float32)`

SetReadReplicaBatchSize sets ReadReplicaBatchSize field to given value.

### HasReadReplicaBatchSize

`func (o *RollMaxBatchSize) HasReadReplicaBatchSize() bool`

HasReadReplicaBatchSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


