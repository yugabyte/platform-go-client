# PerfAdvisorSettingsFormData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionSkewIntervalMins** | Pointer to **int32** | Perf advisor connection skew check interval | [optional] 
**ConnectionSkewMinConnections** | Pointer to **int32** | Perf advisor connection skew min connections | [optional] 
**ConnectionSkewThresholdPct** | Pointer to **float64** | Perf advisor connection skew threshold | [optional] 
**CpuSkewIntervalMins** | Pointer to **int32** | Perf advisor cpu skew check interval | [optional] 
**CpuSkewMinUsagePct** | Pointer to **float64** | Perf advisor cpu skew min cpu usage | [optional] 
**CpuSkewThresholdPct** | Pointer to **float64** | Perf advisor cpu skew threshold | [optional] 
**CpuUsageIntervalMins** | Pointer to **int32** | Perf advisor CPU usage check interval | [optional] 
**CpuUsageThreshold** | Pointer to **float64** | Perf advisor CPU usage threshold | [optional] 
**Enabled** | Pointer to **bool** | Enable/disable perf advisor runs for the universe | [optional] 
**HotShardIntervalMins** | Pointer to **int32** | Perf Advisor hot shard check interval | [optional] 
**HotShardMinimalReads** | Pointer to **int32** | Perf advisor hot shard minimal reads | [optional] 
**HotShardMinimalWrites** | Pointer to **int32** | Perf advisor hot shard minimal writes | [optional] 
**HotShardReadSkewThresholdPct** | Pointer to **float64** | Perf Advisor hot shard read skew threshold | [optional] 
**HotShardWriteSkewThresholdPct** | Pointer to **float64** | Perf Advisor hot shard write skew threshold | [optional] 
**QuerySkewIntervalMins** | Pointer to **int32** | Perf advisor query skew check interval | [optional] 
**QuerySkewMinQueries** | Pointer to **int32** | Perf advisor query skew min queries | [optional] 
**QuerySkewThresholdPct** | Pointer to **float64** | Perf advisor query skew threshold | [optional] 
**RejectedConnIntervalMins** | Pointer to **int32** | Perf advisor rejected connections check interval | [optional] 
**RejectedConnThreshold** | Pointer to **int32** | Perf advisor rejected connections threshold | [optional] 
**UniverseFrequencyMins** | Pointer to **int32** | Perf advisor runs frequency, in minutes | [optional] 

## Methods

### NewPerfAdvisorSettingsFormData

`func NewPerfAdvisorSettingsFormData() *PerfAdvisorSettingsFormData`

NewPerfAdvisorSettingsFormData instantiates a new PerfAdvisorSettingsFormData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerfAdvisorSettingsFormDataWithDefaults

`func NewPerfAdvisorSettingsFormDataWithDefaults() *PerfAdvisorSettingsFormData`

NewPerfAdvisorSettingsFormDataWithDefaults instantiates a new PerfAdvisorSettingsFormData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionSkewIntervalMins

`func (o *PerfAdvisorSettingsFormData) GetConnectionSkewIntervalMins() int32`

GetConnectionSkewIntervalMins returns the ConnectionSkewIntervalMins field if non-nil, zero value otherwise.

### GetConnectionSkewIntervalMinsOk

`func (o *PerfAdvisorSettingsFormData) GetConnectionSkewIntervalMinsOk() (*int32, bool)`

GetConnectionSkewIntervalMinsOk returns a tuple with the ConnectionSkewIntervalMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSkewIntervalMins

`func (o *PerfAdvisorSettingsFormData) SetConnectionSkewIntervalMins(v int32)`

SetConnectionSkewIntervalMins sets ConnectionSkewIntervalMins field to given value.

### HasConnectionSkewIntervalMins

`func (o *PerfAdvisorSettingsFormData) HasConnectionSkewIntervalMins() bool`

HasConnectionSkewIntervalMins returns a boolean if a field has been set.

### GetConnectionSkewMinConnections

`func (o *PerfAdvisorSettingsFormData) GetConnectionSkewMinConnections() int32`

GetConnectionSkewMinConnections returns the ConnectionSkewMinConnections field if non-nil, zero value otherwise.

### GetConnectionSkewMinConnectionsOk

`func (o *PerfAdvisorSettingsFormData) GetConnectionSkewMinConnectionsOk() (*int32, bool)`

GetConnectionSkewMinConnectionsOk returns a tuple with the ConnectionSkewMinConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSkewMinConnections

`func (o *PerfAdvisorSettingsFormData) SetConnectionSkewMinConnections(v int32)`

SetConnectionSkewMinConnections sets ConnectionSkewMinConnections field to given value.

### HasConnectionSkewMinConnections

`func (o *PerfAdvisorSettingsFormData) HasConnectionSkewMinConnections() bool`

HasConnectionSkewMinConnections returns a boolean if a field has been set.

### GetConnectionSkewThresholdPct

`func (o *PerfAdvisorSettingsFormData) GetConnectionSkewThresholdPct() float64`

GetConnectionSkewThresholdPct returns the ConnectionSkewThresholdPct field if non-nil, zero value otherwise.

### GetConnectionSkewThresholdPctOk

`func (o *PerfAdvisorSettingsFormData) GetConnectionSkewThresholdPctOk() (*float64, bool)`

GetConnectionSkewThresholdPctOk returns a tuple with the ConnectionSkewThresholdPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSkewThresholdPct

`func (o *PerfAdvisorSettingsFormData) SetConnectionSkewThresholdPct(v float64)`

SetConnectionSkewThresholdPct sets ConnectionSkewThresholdPct field to given value.

### HasConnectionSkewThresholdPct

`func (o *PerfAdvisorSettingsFormData) HasConnectionSkewThresholdPct() bool`

HasConnectionSkewThresholdPct returns a boolean if a field has been set.

### GetCpuSkewIntervalMins

`func (o *PerfAdvisorSettingsFormData) GetCpuSkewIntervalMins() int32`

GetCpuSkewIntervalMins returns the CpuSkewIntervalMins field if non-nil, zero value otherwise.

### GetCpuSkewIntervalMinsOk

`func (o *PerfAdvisorSettingsFormData) GetCpuSkewIntervalMinsOk() (*int32, bool)`

GetCpuSkewIntervalMinsOk returns a tuple with the CpuSkewIntervalMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSkewIntervalMins

`func (o *PerfAdvisorSettingsFormData) SetCpuSkewIntervalMins(v int32)`

SetCpuSkewIntervalMins sets CpuSkewIntervalMins field to given value.

### HasCpuSkewIntervalMins

`func (o *PerfAdvisorSettingsFormData) HasCpuSkewIntervalMins() bool`

HasCpuSkewIntervalMins returns a boolean if a field has been set.

### GetCpuSkewMinUsagePct

`func (o *PerfAdvisorSettingsFormData) GetCpuSkewMinUsagePct() float64`

GetCpuSkewMinUsagePct returns the CpuSkewMinUsagePct field if non-nil, zero value otherwise.

### GetCpuSkewMinUsagePctOk

`func (o *PerfAdvisorSettingsFormData) GetCpuSkewMinUsagePctOk() (*float64, bool)`

GetCpuSkewMinUsagePctOk returns a tuple with the CpuSkewMinUsagePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSkewMinUsagePct

`func (o *PerfAdvisorSettingsFormData) SetCpuSkewMinUsagePct(v float64)`

SetCpuSkewMinUsagePct sets CpuSkewMinUsagePct field to given value.

### HasCpuSkewMinUsagePct

`func (o *PerfAdvisorSettingsFormData) HasCpuSkewMinUsagePct() bool`

HasCpuSkewMinUsagePct returns a boolean if a field has been set.

### GetCpuSkewThresholdPct

`func (o *PerfAdvisorSettingsFormData) GetCpuSkewThresholdPct() float64`

GetCpuSkewThresholdPct returns the CpuSkewThresholdPct field if non-nil, zero value otherwise.

### GetCpuSkewThresholdPctOk

`func (o *PerfAdvisorSettingsFormData) GetCpuSkewThresholdPctOk() (*float64, bool)`

GetCpuSkewThresholdPctOk returns a tuple with the CpuSkewThresholdPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSkewThresholdPct

`func (o *PerfAdvisorSettingsFormData) SetCpuSkewThresholdPct(v float64)`

SetCpuSkewThresholdPct sets CpuSkewThresholdPct field to given value.

### HasCpuSkewThresholdPct

`func (o *PerfAdvisorSettingsFormData) HasCpuSkewThresholdPct() bool`

HasCpuSkewThresholdPct returns a boolean if a field has been set.

### GetCpuUsageIntervalMins

`func (o *PerfAdvisorSettingsFormData) GetCpuUsageIntervalMins() int32`

GetCpuUsageIntervalMins returns the CpuUsageIntervalMins field if non-nil, zero value otherwise.

### GetCpuUsageIntervalMinsOk

`func (o *PerfAdvisorSettingsFormData) GetCpuUsageIntervalMinsOk() (*int32, bool)`

GetCpuUsageIntervalMinsOk returns a tuple with the CpuUsageIntervalMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageIntervalMins

`func (o *PerfAdvisorSettingsFormData) SetCpuUsageIntervalMins(v int32)`

SetCpuUsageIntervalMins sets CpuUsageIntervalMins field to given value.

### HasCpuUsageIntervalMins

`func (o *PerfAdvisorSettingsFormData) HasCpuUsageIntervalMins() bool`

HasCpuUsageIntervalMins returns a boolean if a field has been set.

### GetCpuUsageThreshold

`func (o *PerfAdvisorSettingsFormData) GetCpuUsageThreshold() float64`

GetCpuUsageThreshold returns the CpuUsageThreshold field if non-nil, zero value otherwise.

### GetCpuUsageThresholdOk

`func (o *PerfAdvisorSettingsFormData) GetCpuUsageThresholdOk() (*float64, bool)`

GetCpuUsageThresholdOk returns a tuple with the CpuUsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageThreshold

`func (o *PerfAdvisorSettingsFormData) SetCpuUsageThreshold(v float64)`

SetCpuUsageThreshold sets CpuUsageThreshold field to given value.

### HasCpuUsageThreshold

`func (o *PerfAdvisorSettingsFormData) HasCpuUsageThreshold() bool`

HasCpuUsageThreshold returns a boolean if a field has been set.

### GetEnabled

`func (o *PerfAdvisorSettingsFormData) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PerfAdvisorSettingsFormData) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PerfAdvisorSettingsFormData) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PerfAdvisorSettingsFormData) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHotShardIntervalMins

`func (o *PerfAdvisorSettingsFormData) GetHotShardIntervalMins() int32`

GetHotShardIntervalMins returns the HotShardIntervalMins field if non-nil, zero value otherwise.

### GetHotShardIntervalMinsOk

`func (o *PerfAdvisorSettingsFormData) GetHotShardIntervalMinsOk() (*int32, bool)`

GetHotShardIntervalMinsOk returns a tuple with the HotShardIntervalMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotShardIntervalMins

`func (o *PerfAdvisorSettingsFormData) SetHotShardIntervalMins(v int32)`

SetHotShardIntervalMins sets HotShardIntervalMins field to given value.

### HasHotShardIntervalMins

`func (o *PerfAdvisorSettingsFormData) HasHotShardIntervalMins() bool`

HasHotShardIntervalMins returns a boolean if a field has been set.

### GetHotShardMinimalReads

`func (o *PerfAdvisorSettingsFormData) GetHotShardMinimalReads() int32`

GetHotShardMinimalReads returns the HotShardMinimalReads field if non-nil, zero value otherwise.

### GetHotShardMinimalReadsOk

`func (o *PerfAdvisorSettingsFormData) GetHotShardMinimalReadsOk() (*int32, bool)`

GetHotShardMinimalReadsOk returns a tuple with the HotShardMinimalReads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotShardMinimalReads

`func (o *PerfAdvisorSettingsFormData) SetHotShardMinimalReads(v int32)`

SetHotShardMinimalReads sets HotShardMinimalReads field to given value.

### HasHotShardMinimalReads

`func (o *PerfAdvisorSettingsFormData) HasHotShardMinimalReads() bool`

HasHotShardMinimalReads returns a boolean if a field has been set.

### GetHotShardMinimalWrites

`func (o *PerfAdvisorSettingsFormData) GetHotShardMinimalWrites() int32`

GetHotShardMinimalWrites returns the HotShardMinimalWrites field if non-nil, zero value otherwise.

### GetHotShardMinimalWritesOk

`func (o *PerfAdvisorSettingsFormData) GetHotShardMinimalWritesOk() (*int32, bool)`

GetHotShardMinimalWritesOk returns a tuple with the HotShardMinimalWrites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotShardMinimalWrites

`func (o *PerfAdvisorSettingsFormData) SetHotShardMinimalWrites(v int32)`

SetHotShardMinimalWrites sets HotShardMinimalWrites field to given value.

### HasHotShardMinimalWrites

`func (o *PerfAdvisorSettingsFormData) HasHotShardMinimalWrites() bool`

HasHotShardMinimalWrites returns a boolean if a field has been set.

### GetHotShardReadSkewThresholdPct

`func (o *PerfAdvisorSettingsFormData) GetHotShardReadSkewThresholdPct() float64`

GetHotShardReadSkewThresholdPct returns the HotShardReadSkewThresholdPct field if non-nil, zero value otherwise.

### GetHotShardReadSkewThresholdPctOk

`func (o *PerfAdvisorSettingsFormData) GetHotShardReadSkewThresholdPctOk() (*float64, bool)`

GetHotShardReadSkewThresholdPctOk returns a tuple with the HotShardReadSkewThresholdPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotShardReadSkewThresholdPct

`func (o *PerfAdvisorSettingsFormData) SetHotShardReadSkewThresholdPct(v float64)`

SetHotShardReadSkewThresholdPct sets HotShardReadSkewThresholdPct field to given value.

### HasHotShardReadSkewThresholdPct

`func (o *PerfAdvisorSettingsFormData) HasHotShardReadSkewThresholdPct() bool`

HasHotShardReadSkewThresholdPct returns a boolean if a field has been set.

### GetHotShardWriteSkewThresholdPct

`func (o *PerfAdvisorSettingsFormData) GetHotShardWriteSkewThresholdPct() float64`

GetHotShardWriteSkewThresholdPct returns the HotShardWriteSkewThresholdPct field if non-nil, zero value otherwise.

### GetHotShardWriteSkewThresholdPctOk

`func (o *PerfAdvisorSettingsFormData) GetHotShardWriteSkewThresholdPctOk() (*float64, bool)`

GetHotShardWriteSkewThresholdPctOk returns a tuple with the HotShardWriteSkewThresholdPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotShardWriteSkewThresholdPct

`func (o *PerfAdvisorSettingsFormData) SetHotShardWriteSkewThresholdPct(v float64)`

SetHotShardWriteSkewThresholdPct sets HotShardWriteSkewThresholdPct field to given value.

### HasHotShardWriteSkewThresholdPct

`func (o *PerfAdvisorSettingsFormData) HasHotShardWriteSkewThresholdPct() bool`

HasHotShardWriteSkewThresholdPct returns a boolean if a field has been set.

### GetQuerySkewIntervalMins

`func (o *PerfAdvisorSettingsFormData) GetQuerySkewIntervalMins() int32`

GetQuerySkewIntervalMins returns the QuerySkewIntervalMins field if non-nil, zero value otherwise.

### GetQuerySkewIntervalMinsOk

`func (o *PerfAdvisorSettingsFormData) GetQuerySkewIntervalMinsOk() (*int32, bool)`

GetQuerySkewIntervalMinsOk returns a tuple with the QuerySkewIntervalMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerySkewIntervalMins

`func (o *PerfAdvisorSettingsFormData) SetQuerySkewIntervalMins(v int32)`

SetQuerySkewIntervalMins sets QuerySkewIntervalMins field to given value.

### HasQuerySkewIntervalMins

`func (o *PerfAdvisorSettingsFormData) HasQuerySkewIntervalMins() bool`

HasQuerySkewIntervalMins returns a boolean if a field has been set.

### GetQuerySkewMinQueries

`func (o *PerfAdvisorSettingsFormData) GetQuerySkewMinQueries() int32`

GetQuerySkewMinQueries returns the QuerySkewMinQueries field if non-nil, zero value otherwise.

### GetQuerySkewMinQueriesOk

`func (o *PerfAdvisorSettingsFormData) GetQuerySkewMinQueriesOk() (*int32, bool)`

GetQuerySkewMinQueriesOk returns a tuple with the QuerySkewMinQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerySkewMinQueries

`func (o *PerfAdvisorSettingsFormData) SetQuerySkewMinQueries(v int32)`

SetQuerySkewMinQueries sets QuerySkewMinQueries field to given value.

### HasQuerySkewMinQueries

`func (o *PerfAdvisorSettingsFormData) HasQuerySkewMinQueries() bool`

HasQuerySkewMinQueries returns a boolean if a field has been set.

### GetQuerySkewThresholdPct

`func (o *PerfAdvisorSettingsFormData) GetQuerySkewThresholdPct() float64`

GetQuerySkewThresholdPct returns the QuerySkewThresholdPct field if non-nil, zero value otherwise.

### GetQuerySkewThresholdPctOk

`func (o *PerfAdvisorSettingsFormData) GetQuerySkewThresholdPctOk() (*float64, bool)`

GetQuerySkewThresholdPctOk returns a tuple with the QuerySkewThresholdPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerySkewThresholdPct

`func (o *PerfAdvisorSettingsFormData) SetQuerySkewThresholdPct(v float64)`

SetQuerySkewThresholdPct sets QuerySkewThresholdPct field to given value.

### HasQuerySkewThresholdPct

`func (o *PerfAdvisorSettingsFormData) HasQuerySkewThresholdPct() bool`

HasQuerySkewThresholdPct returns a boolean if a field has been set.

### GetRejectedConnIntervalMins

`func (o *PerfAdvisorSettingsFormData) GetRejectedConnIntervalMins() int32`

GetRejectedConnIntervalMins returns the RejectedConnIntervalMins field if non-nil, zero value otherwise.

### GetRejectedConnIntervalMinsOk

`func (o *PerfAdvisorSettingsFormData) GetRejectedConnIntervalMinsOk() (*int32, bool)`

GetRejectedConnIntervalMinsOk returns a tuple with the RejectedConnIntervalMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedConnIntervalMins

`func (o *PerfAdvisorSettingsFormData) SetRejectedConnIntervalMins(v int32)`

SetRejectedConnIntervalMins sets RejectedConnIntervalMins field to given value.

### HasRejectedConnIntervalMins

`func (o *PerfAdvisorSettingsFormData) HasRejectedConnIntervalMins() bool`

HasRejectedConnIntervalMins returns a boolean if a field has been set.

### GetRejectedConnThreshold

`func (o *PerfAdvisorSettingsFormData) GetRejectedConnThreshold() int32`

GetRejectedConnThreshold returns the RejectedConnThreshold field if non-nil, zero value otherwise.

### GetRejectedConnThresholdOk

`func (o *PerfAdvisorSettingsFormData) GetRejectedConnThresholdOk() (*int32, bool)`

GetRejectedConnThresholdOk returns a tuple with the RejectedConnThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedConnThreshold

`func (o *PerfAdvisorSettingsFormData) SetRejectedConnThreshold(v int32)`

SetRejectedConnThreshold sets RejectedConnThreshold field to given value.

### HasRejectedConnThreshold

`func (o *PerfAdvisorSettingsFormData) HasRejectedConnThreshold() bool`

HasRejectedConnThreshold returns a boolean if a field has been set.

### GetUniverseFrequencyMins

`func (o *PerfAdvisorSettingsFormData) GetUniverseFrequencyMins() int32`

GetUniverseFrequencyMins returns the UniverseFrequencyMins field if non-nil, zero value otherwise.

### GetUniverseFrequencyMinsOk

`func (o *PerfAdvisorSettingsFormData) GetUniverseFrequencyMinsOk() (*int32, bool)`

GetUniverseFrequencyMinsOk returns a tuple with the UniverseFrequencyMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseFrequencyMins

`func (o *PerfAdvisorSettingsFormData) SetUniverseFrequencyMins(v int32)`

SetUniverseFrequencyMins sets UniverseFrequencyMins field to given value.

### HasUniverseFrequencyMins

`func (o *PerfAdvisorSettingsFormData) HasUniverseFrequencyMins() bool`

HasUniverseFrequencyMins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


