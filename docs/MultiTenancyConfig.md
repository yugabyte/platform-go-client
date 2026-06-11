# MultiTenancyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableQos** | Pointer to **bool** | Enable QoS-based multi-tenancy (sets enable_qos gflag) | [optional] 
**QosMaxDbCount** | Pointer to **int32** | Maximum number of databases allowed | [optional] 
**QosMaxDbCpuPercent** | Pointer to **float64** | Maximum per-database CPU percentage | [optional] 

## Methods

### NewMultiTenancyConfig

`func NewMultiTenancyConfig() *MultiTenancyConfig`

NewMultiTenancyConfig instantiates a new MultiTenancyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiTenancyConfigWithDefaults

`func NewMultiTenancyConfigWithDefaults() *MultiTenancyConfig`

NewMultiTenancyConfigWithDefaults instantiates a new MultiTenancyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableQos

`func (o *MultiTenancyConfig) GetEnableQos() bool`

GetEnableQos returns the EnableQos field if non-nil, zero value otherwise.

### GetEnableQosOk

`func (o *MultiTenancyConfig) GetEnableQosOk() (*bool, bool)`

GetEnableQosOk returns a tuple with the EnableQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQos

`func (o *MultiTenancyConfig) SetEnableQos(v bool)`

SetEnableQos sets EnableQos field to given value.

### HasEnableQos

`func (o *MultiTenancyConfig) HasEnableQos() bool`

HasEnableQos returns a boolean if a field has been set.

### GetQosMaxDbCount

`func (o *MultiTenancyConfig) GetQosMaxDbCount() int32`

GetQosMaxDbCount returns the QosMaxDbCount field if non-nil, zero value otherwise.

### GetQosMaxDbCountOk

`func (o *MultiTenancyConfig) GetQosMaxDbCountOk() (*int32, bool)`

GetQosMaxDbCountOk returns a tuple with the QosMaxDbCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosMaxDbCount

`func (o *MultiTenancyConfig) SetQosMaxDbCount(v int32)`

SetQosMaxDbCount sets QosMaxDbCount field to given value.

### HasQosMaxDbCount

`func (o *MultiTenancyConfig) HasQosMaxDbCount() bool`

HasQosMaxDbCount returns a boolean if a field has been set.

### GetQosMaxDbCpuPercent

`func (o *MultiTenancyConfig) GetQosMaxDbCpuPercent() float64`

GetQosMaxDbCpuPercent returns the QosMaxDbCpuPercent field if non-nil, zero value otherwise.

### GetQosMaxDbCpuPercentOk

`func (o *MultiTenancyConfig) GetQosMaxDbCpuPercentOk() (*float64, bool)`

GetQosMaxDbCpuPercentOk returns a tuple with the QosMaxDbCpuPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosMaxDbCpuPercent

`func (o *MultiTenancyConfig) SetQosMaxDbCpuPercent(v float64)`

SetQosMaxDbCpuPercent sets QosMaxDbCpuPercent field to given value.

### HasQosMaxDbCpuPercent

`func (o *MultiTenancyConfig) HasQosMaxDbCpuPercent() bool`

HasQosMaxDbCpuPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


