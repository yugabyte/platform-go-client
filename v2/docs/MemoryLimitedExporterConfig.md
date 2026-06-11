# MemoryLimitedExporterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemoryLimitMib** | Pointer to **int32** | Memory limit in MiB for the OpenTelemetry Collector process in the config file. | [optional] [default to 2048]
**MemoryLimitCheckIntervalSeconds** | Pointer to **int32** | Check interval in seconds for the MemoryLimiterProcessor. | [optional] [default to 10]

## Methods

### NewMemoryLimitedExporterConfig

`func NewMemoryLimitedExporterConfig() *MemoryLimitedExporterConfig`

NewMemoryLimitedExporterConfig instantiates a new MemoryLimitedExporterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryLimitedExporterConfigWithDefaults

`func NewMemoryLimitedExporterConfigWithDefaults() *MemoryLimitedExporterConfig`

NewMemoryLimitedExporterConfigWithDefaults instantiates a new MemoryLimitedExporterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemoryLimitMib

`func (o *MemoryLimitedExporterConfig) GetMemoryLimitMib() int32`

GetMemoryLimitMib returns the MemoryLimitMib field if non-nil, zero value otherwise.

### GetMemoryLimitMibOk

`func (o *MemoryLimitedExporterConfig) GetMemoryLimitMibOk() (*int32, bool)`

GetMemoryLimitMibOk returns a tuple with the MemoryLimitMib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimitMib

`func (o *MemoryLimitedExporterConfig) SetMemoryLimitMib(v int32)`

SetMemoryLimitMib sets MemoryLimitMib field to given value.

### HasMemoryLimitMib

`func (o *MemoryLimitedExporterConfig) HasMemoryLimitMib() bool`

HasMemoryLimitMib returns a boolean if a field has been set.

### GetMemoryLimitCheckIntervalSeconds

`func (o *MemoryLimitedExporterConfig) GetMemoryLimitCheckIntervalSeconds() int32`

GetMemoryLimitCheckIntervalSeconds returns the MemoryLimitCheckIntervalSeconds field if non-nil, zero value otherwise.

### GetMemoryLimitCheckIntervalSecondsOk

`func (o *MemoryLimitedExporterConfig) GetMemoryLimitCheckIntervalSecondsOk() (*int32, bool)`

GetMemoryLimitCheckIntervalSecondsOk returns a tuple with the MemoryLimitCheckIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimitCheckIntervalSeconds

`func (o *MemoryLimitedExporterConfig) SetMemoryLimitCheckIntervalSeconds(v int32)`

SetMemoryLimitCheckIntervalSeconds sets MemoryLimitCheckIntervalSeconds field to given value.

### HasMemoryLimitCheckIntervalSeconds

`func (o *MemoryLimitedExporterConfig) HasMemoryLimitCheckIntervalSeconds() bool`

HasMemoryLimitCheckIntervalSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


