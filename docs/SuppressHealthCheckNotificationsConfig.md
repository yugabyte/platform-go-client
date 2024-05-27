# SuppressHealthCheckNotificationsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuppressAllUniverses** | Pointer to **bool** | Suppress health check notifications on all the universes (including future universes) | [optional] 
**UniverseUUIDSet** | Pointer to **[]string** | Set of universe uuids to suppress health check notifications on | [optional] 

## Methods

### NewSuppressHealthCheckNotificationsConfig

`func NewSuppressHealthCheckNotificationsConfig() *SuppressHealthCheckNotificationsConfig`

NewSuppressHealthCheckNotificationsConfig instantiates a new SuppressHealthCheckNotificationsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuppressHealthCheckNotificationsConfigWithDefaults

`func NewSuppressHealthCheckNotificationsConfigWithDefaults() *SuppressHealthCheckNotificationsConfig`

NewSuppressHealthCheckNotificationsConfigWithDefaults instantiates a new SuppressHealthCheckNotificationsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuppressAllUniverses

`func (o *SuppressHealthCheckNotificationsConfig) GetSuppressAllUniverses() bool`

GetSuppressAllUniverses returns the SuppressAllUniverses field if non-nil, zero value otherwise.

### GetSuppressAllUniversesOk

`func (o *SuppressHealthCheckNotificationsConfig) GetSuppressAllUniversesOk() (*bool, bool)`

GetSuppressAllUniversesOk returns a tuple with the SuppressAllUniverses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressAllUniverses

`func (o *SuppressHealthCheckNotificationsConfig) SetSuppressAllUniverses(v bool)`

SetSuppressAllUniverses sets SuppressAllUniverses field to given value.

### HasSuppressAllUniverses

`func (o *SuppressHealthCheckNotificationsConfig) HasSuppressAllUniverses() bool`

HasSuppressAllUniverses returns a boolean if a field has been set.

### GetUniverseUUIDSet

`func (o *SuppressHealthCheckNotificationsConfig) GetUniverseUUIDSet() []string`

GetUniverseUUIDSet returns the UniverseUUIDSet field if non-nil, zero value otherwise.

### GetUniverseUUIDSetOk

`func (o *SuppressHealthCheckNotificationsConfig) GetUniverseUUIDSetOk() (*[]string, bool)`

GetUniverseUUIDSetOk returns a tuple with the UniverseUUIDSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUIDSet

`func (o *SuppressHealthCheckNotificationsConfig) SetUniverseUUIDSet(v []string)`

SetUniverseUUIDSet sets UniverseUUIDSet field to given value.

### HasUniverseUUIDSet

`func (o *SuppressHealthCheckNotificationsConfig) HasUniverseUUIDSet() bool`

HasUniverseUUIDSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


