# BatchedExporterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SendBatchMaxSize** | Pointer to **int32** | Send batch max size | [optional] [default to 1000]
**SendBatchSize** | Pointer to **int32** | Send batch size | [optional] [default to 100]
**SendBatchTimeoutSeconds** | Pointer to **int32** | Send batch timeout in seconds | [optional] [default to 10]

## Methods

### NewBatchedExporterConfig

`func NewBatchedExporterConfig() *BatchedExporterConfig`

NewBatchedExporterConfig instantiates a new BatchedExporterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchedExporterConfigWithDefaults

`func NewBatchedExporterConfigWithDefaults() *BatchedExporterConfig`

NewBatchedExporterConfigWithDefaults instantiates a new BatchedExporterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSendBatchMaxSize

`func (o *BatchedExporterConfig) GetSendBatchMaxSize() int32`

GetSendBatchMaxSize returns the SendBatchMaxSize field if non-nil, zero value otherwise.

### GetSendBatchMaxSizeOk

`func (o *BatchedExporterConfig) GetSendBatchMaxSizeOk() (*int32, bool)`

GetSendBatchMaxSizeOk returns a tuple with the SendBatchMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBatchMaxSize

`func (o *BatchedExporterConfig) SetSendBatchMaxSize(v int32)`

SetSendBatchMaxSize sets SendBatchMaxSize field to given value.

### HasSendBatchMaxSize

`func (o *BatchedExporterConfig) HasSendBatchMaxSize() bool`

HasSendBatchMaxSize returns a boolean if a field has been set.

### GetSendBatchSize

`func (o *BatchedExporterConfig) GetSendBatchSize() int32`

GetSendBatchSize returns the SendBatchSize field if non-nil, zero value otherwise.

### GetSendBatchSizeOk

`func (o *BatchedExporterConfig) GetSendBatchSizeOk() (*int32, bool)`

GetSendBatchSizeOk returns a tuple with the SendBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBatchSize

`func (o *BatchedExporterConfig) SetSendBatchSize(v int32)`

SetSendBatchSize sets SendBatchSize field to given value.

### HasSendBatchSize

`func (o *BatchedExporterConfig) HasSendBatchSize() bool`

HasSendBatchSize returns a boolean if a field has been set.

### GetSendBatchTimeoutSeconds

`func (o *BatchedExporterConfig) GetSendBatchTimeoutSeconds() int32`

GetSendBatchTimeoutSeconds returns the SendBatchTimeoutSeconds field if non-nil, zero value otherwise.

### GetSendBatchTimeoutSecondsOk

`func (o *BatchedExporterConfig) GetSendBatchTimeoutSecondsOk() (*int32, bool)`

GetSendBatchTimeoutSecondsOk returns a tuple with the SendBatchTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBatchTimeoutSeconds

`func (o *BatchedExporterConfig) SetSendBatchTimeoutSeconds(v int32)`

SetSendBatchTimeoutSeconds sets SendBatchTimeoutSeconds field to given value.

### HasSendBatchTimeoutSeconds

`func (o *BatchedExporterConfig) HasSendBatchTimeoutSeconds() bool`

HasSendBatchTimeoutSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


