# ExporterRetryConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether retry on failure is enabled | [optional] 
**InitialInterval** | Pointer to **string** | Initial interval between retries. Duration string (e.g. \&quot;30s\&quot;, \&quot;1m\&quot;). Used as initial_interval in Otel collector. | [optional] 
**MaxElapsedTime** | Pointer to **string** | Maximum elapsed time for all retries. Duration string (e.g. \&quot;60m\&quot;, \&quot;1800m\&quot;). Used as max_elapsed_time in Otel collector. | [optional] 
**MaxInterval** | Pointer to **string** | Maximum interval between retries. Duration string (e.g. \&quot;10m\&quot;, \&quot;1800m\&quot;). Used as max_interval in Otel collector. | [optional] 

## Methods

### NewExporterRetryConfig

`func NewExporterRetryConfig() *ExporterRetryConfig`

NewExporterRetryConfig instantiates a new ExporterRetryConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExporterRetryConfigWithDefaults

`func NewExporterRetryConfigWithDefaults() *ExporterRetryConfig`

NewExporterRetryConfigWithDefaults instantiates a new ExporterRetryConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ExporterRetryConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExporterRetryConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExporterRetryConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ExporterRetryConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInitialInterval

`func (o *ExporterRetryConfig) GetInitialInterval() string`

GetInitialInterval returns the InitialInterval field if non-nil, zero value otherwise.

### GetInitialIntervalOk

`func (o *ExporterRetryConfig) GetInitialIntervalOk() (*string, bool)`

GetInitialIntervalOk returns a tuple with the InitialInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialInterval

`func (o *ExporterRetryConfig) SetInitialInterval(v string)`

SetInitialInterval sets InitialInterval field to given value.

### HasInitialInterval

`func (o *ExporterRetryConfig) HasInitialInterval() bool`

HasInitialInterval returns a boolean if a field has been set.

### GetMaxElapsedTime

`func (o *ExporterRetryConfig) GetMaxElapsedTime() string`

GetMaxElapsedTime returns the MaxElapsedTime field if non-nil, zero value otherwise.

### GetMaxElapsedTimeOk

`func (o *ExporterRetryConfig) GetMaxElapsedTimeOk() (*string, bool)`

GetMaxElapsedTimeOk returns a tuple with the MaxElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxElapsedTime

`func (o *ExporterRetryConfig) SetMaxElapsedTime(v string)`

SetMaxElapsedTime sets MaxElapsedTime field to given value.

### HasMaxElapsedTime

`func (o *ExporterRetryConfig) HasMaxElapsedTime() bool`

HasMaxElapsedTime returns a boolean if a field has been set.

### GetMaxInterval

`func (o *ExporterRetryConfig) GetMaxInterval() string`

GetMaxInterval returns the MaxInterval field if non-nil, zero value otherwise.

### GetMaxIntervalOk

`func (o *ExporterRetryConfig) GetMaxIntervalOk() (*string, bool)`

GetMaxIntervalOk returns a tuple with the MaxInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInterval

`func (o *ExporterRetryConfig) SetMaxInterval(v string)`

SetMaxInterval sets MaxInterval field to given value.

### HasMaxInterval

`func (o *ExporterRetryConfig) HasMaxInterval() bool`

HasMaxInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


