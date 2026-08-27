# PACollectorUniverseRegistrationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedObservability** | Pointer to **bool** | Whether advanced observability (metrics export to Prometheus) is enabled | [optional] 
**Mode** | Pointer to **string** | How the universe is registered with the collector | [optional] 
**PaEndpointName** | Pointer to **string** | Perf Advisor Endpoint name, set for ONLINE mode only | [optional] 
**PaEndpointUuid** | Pointer to **string** | Perf Advisor Endpoint UUID, set for ONLINE mode only | [optional] 
**Success** | Pointer to **bool** | Whether the universe is registered with PA Collector | [optional] 

## Methods

### NewPACollectorUniverseRegistrationStatus

`func NewPACollectorUniverseRegistrationStatus() *PACollectorUniverseRegistrationStatus`

NewPACollectorUniverseRegistrationStatus instantiates a new PACollectorUniverseRegistrationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPACollectorUniverseRegistrationStatusWithDefaults

`func NewPACollectorUniverseRegistrationStatusWithDefaults() *PACollectorUniverseRegistrationStatus`

NewPACollectorUniverseRegistrationStatusWithDefaults instantiates a new PACollectorUniverseRegistrationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedObservability

`func (o *PACollectorUniverseRegistrationStatus) GetAdvancedObservability() bool`

GetAdvancedObservability returns the AdvancedObservability field if non-nil, zero value otherwise.

### GetAdvancedObservabilityOk

`func (o *PACollectorUniverseRegistrationStatus) GetAdvancedObservabilityOk() (*bool, bool)`

GetAdvancedObservabilityOk returns a tuple with the AdvancedObservability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedObservability

`func (o *PACollectorUniverseRegistrationStatus) SetAdvancedObservability(v bool)`

SetAdvancedObservability sets AdvancedObservability field to given value.

### HasAdvancedObservability

`func (o *PACollectorUniverseRegistrationStatus) HasAdvancedObservability() bool`

HasAdvancedObservability returns a boolean if a field has been set.

### GetMode

`func (o *PACollectorUniverseRegistrationStatus) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PACollectorUniverseRegistrationStatus) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PACollectorUniverseRegistrationStatus) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PACollectorUniverseRegistrationStatus) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPaEndpointName

`func (o *PACollectorUniverseRegistrationStatus) GetPaEndpointName() string`

GetPaEndpointName returns the PaEndpointName field if non-nil, zero value otherwise.

### GetPaEndpointNameOk

`func (o *PACollectorUniverseRegistrationStatus) GetPaEndpointNameOk() (*string, bool)`

GetPaEndpointNameOk returns a tuple with the PaEndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaEndpointName

`func (o *PACollectorUniverseRegistrationStatus) SetPaEndpointName(v string)`

SetPaEndpointName sets PaEndpointName field to given value.

### HasPaEndpointName

`func (o *PACollectorUniverseRegistrationStatus) HasPaEndpointName() bool`

HasPaEndpointName returns a boolean if a field has been set.

### GetPaEndpointUuid

`func (o *PACollectorUniverseRegistrationStatus) GetPaEndpointUuid() string`

GetPaEndpointUuid returns the PaEndpointUuid field if non-nil, zero value otherwise.

### GetPaEndpointUuidOk

`func (o *PACollectorUniverseRegistrationStatus) GetPaEndpointUuidOk() (*string, bool)`

GetPaEndpointUuidOk returns a tuple with the PaEndpointUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaEndpointUuid

`func (o *PACollectorUniverseRegistrationStatus) SetPaEndpointUuid(v string)`

SetPaEndpointUuid sets PaEndpointUuid field to given value.

### HasPaEndpointUuid

`func (o *PACollectorUniverseRegistrationStatus) HasPaEndpointUuid() bool`

HasPaEndpointUuid returns a boolean if a field has been set.

### GetSuccess

`func (o *PACollectorUniverseRegistrationStatus) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PACollectorUniverseRegistrationStatus) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PACollectorUniverseRegistrationStatus) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *PACollectorUniverseRegistrationStatus) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


