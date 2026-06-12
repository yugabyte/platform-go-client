# PACollectorRegisteredUniverseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedObservability** | Pointer to **bool** | Whether advanced observability (metrics export) is enabled | [optional] 
**DataMountPoints** | Pointer to **[]string** | Data mount points | [optional] 
**OtherMountPoints** | Pointer to **[]string** | Other mount points | [optional] 
**UniverseName** | Pointer to **string** | Universe name (from YBA, or null if universe was deleted) | [optional] 
**UniverseUuid** | Pointer to **string** | Universe UUID | [optional] 

## Methods

### NewPACollectorRegisteredUniverseInfo

`func NewPACollectorRegisteredUniverseInfo() *PACollectorRegisteredUniverseInfo`

NewPACollectorRegisteredUniverseInfo instantiates a new PACollectorRegisteredUniverseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPACollectorRegisteredUniverseInfoWithDefaults

`func NewPACollectorRegisteredUniverseInfoWithDefaults() *PACollectorRegisteredUniverseInfo`

NewPACollectorRegisteredUniverseInfoWithDefaults instantiates a new PACollectorRegisteredUniverseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedObservability

`func (o *PACollectorRegisteredUniverseInfo) GetAdvancedObservability() bool`

GetAdvancedObservability returns the AdvancedObservability field if non-nil, zero value otherwise.

### GetAdvancedObservabilityOk

`func (o *PACollectorRegisteredUniverseInfo) GetAdvancedObservabilityOk() (*bool, bool)`

GetAdvancedObservabilityOk returns a tuple with the AdvancedObservability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedObservability

`func (o *PACollectorRegisteredUniverseInfo) SetAdvancedObservability(v bool)`

SetAdvancedObservability sets AdvancedObservability field to given value.

### HasAdvancedObservability

`func (o *PACollectorRegisteredUniverseInfo) HasAdvancedObservability() bool`

HasAdvancedObservability returns a boolean if a field has been set.

### GetDataMountPoints

`func (o *PACollectorRegisteredUniverseInfo) GetDataMountPoints() []string`

GetDataMountPoints returns the DataMountPoints field if non-nil, zero value otherwise.

### GetDataMountPointsOk

`func (o *PACollectorRegisteredUniverseInfo) GetDataMountPointsOk() (*[]string, bool)`

GetDataMountPointsOk returns a tuple with the DataMountPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataMountPoints

`func (o *PACollectorRegisteredUniverseInfo) SetDataMountPoints(v []string)`

SetDataMountPoints sets DataMountPoints field to given value.

### HasDataMountPoints

`func (o *PACollectorRegisteredUniverseInfo) HasDataMountPoints() bool`

HasDataMountPoints returns a boolean if a field has been set.

### GetOtherMountPoints

`func (o *PACollectorRegisteredUniverseInfo) GetOtherMountPoints() []string`

GetOtherMountPoints returns the OtherMountPoints field if non-nil, zero value otherwise.

### GetOtherMountPointsOk

`func (o *PACollectorRegisteredUniverseInfo) GetOtherMountPointsOk() (*[]string, bool)`

GetOtherMountPointsOk returns a tuple with the OtherMountPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherMountPoints

`func (o *PACollectorRegisteredUniverseInfo) SetOtherMountPoints(v []string)`

SetOtherMountPoints sets OtherMountPoints field to given value.

### HasOtherMountPoints

`func (o *PACollectorRegisteredUniverseInfo) HasOtherMountPoints() bool`

HasOtherMountPoints returns a boolean if a field has been set.

### GetUniverseName

`func (o *PACollectorRegisteredUniverseInfo) GetUniverseName() string`

GetUniverseName returns the UniverseName field if non-nil, zero value otherwise.

### GetUniverseNameOk

`func (o *PACollectorRegisteredUniverseInfo) GetUniverseNameOk() (*string, bool)`

GetUniverseNameOk returns a tuple with the UniverseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseName

`func (o *PACollectorRegisteredUniverseInfo) SetUniverseName(v string)`

SetUniverseName sets UniverseName field to given value.

### HasUniverseName

`func (o *PACollectorRegisteredUniverseInfo) HasUniverseName() bool`

HasUniverseName returns a boolean if a field has been set.

### GetUniverseUuid

`func (o *PACollectorRegisteredUniverseInfo) GetUniverseUuid() string`

GetUniverseUuid returns the UniverseUuid field if non-nil, zero value otherwise.

### GetUniverseUuidOk

`func (o *PACollectorRegisteredUniverseInfo) GetUniverseUuidOk() (*string, bool)`

GetUniverseUuidOk returns a tuple with the UniverseUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUuid

`func (o *PACollectorRegisteredUniverseInfo) SetUniverseUuid(v string)`

SetUniverseUuid sets UniverseUuid field to given value.

### HasUniverseUuid

`func (o *PACollectorRegisteredUniverseInfo) HasUniverseUuid() bool`

HasUniverseUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


