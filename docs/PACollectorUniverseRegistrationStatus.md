# PACollectorUniverseRegistrationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedObservability** | Pointer to **bool** | Whether advanced observability (metrics export to Prometheus) is enabled | [optional] 
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


