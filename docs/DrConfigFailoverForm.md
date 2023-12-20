# DrConfigFailoverForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DrReplicaUniverseUuid** | Pointer to **string** | New dr replica universe UUID | [optional] 
**NamespaceIdSafetimeEpochUsMap** | Pointer to **map[string]int64** | A map from database ID to its safetime since epoch in micro-seconds to use during unplanned failover | [optional] 
**PrimaryUniverseUuid** | Pointer to **string** | New primary universe UUID | [optional] 

## Methods

### NewDrConfigFailoverForm

`func NewDrConfigFailoverForm() *DrConfigFailoverForm`

NewDrConfigFailoverForm instantiates a new DrConfigFailoverForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrConfigFailoverFormWithDefaults

`func NewDrConfigFailoverFormWithDefaults() *DrConfigFailoverForm`

NewDrConfigFailoverFormWithDefaults instantiates a new DrConfigFailoverForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrReplicaUniverseUuid

`func (o *DrConfigFailoverForm) GetDrReplicaUniverseUuid() string`

GetDrReplicaUniverseUuid returns the DrReplicaUniverseUuid field if non-nil, zero value otherwise.

### GetDrReplicaUniverseUuidOk

`func (o *DrConfigFailoverForm) GetDrReplicaUniverseUuidOk() (*string, bool)`

GetDrReplicaUniverseUuidOk returns a tuple with the DrReplicaUniverseUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrReplicaUniverseUuid

`func (o *DrConfigFailoverForm) SetDrReplicaUniverseUuid(v string)`

SetDrReplicaUniverseUuid sets DrReplicaUniverseUuid field to given value.

### HasDrReplicaUniverseUuid

`func (o *DrConfigFailoverForm) HasDrReplicaUniverseUuid() bool`

HasDrReplicaUniverseUuid returns a boolean if a field has been set.

### GetNamespaceIdSafetimeEpochUsMap

`func (o *DrConfigFailoverForm) GetNamespaceIdSafetimeEpochUsMap() map[string]int64`

GetNamespaceIdSafetimeEpochUsMap returns the NamespaceIdSafetimeEpochUsMap field if non-nil, zero value otherwise.

### GetNamespaceIdSafetimeEpochUsMapOk

`func (o *DrConfigFailoverForm) GetNamespaceIdSafetimeEpochUsMapOk() (*map[string]int64, bool)`

GetNamespaceIdSafetimeEpochUsMapOk returns a tuple with the NamespaceIdSafetimeEpochUsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceIdSafetimeEpochUsMap

`func (o *DrConfigFailoverForm) SetNamespaceIdSafetimeEpochUsMap(v map[string]int64)`

SetNamespaceIdSafetimeEpochUsMap sets NamespaceIdSafetimeEpochUsMap field to given value.

### HasNamespaceIdSafetimeEpochUsMap

`func (o *DrConfigFailoverForm) HasNamespaceIdSafetimeEpochUsMap() bool`

HasNamespaceIdSafetimeEpochUsMap returns a boolean if a field has been set.

### GetPrimaryUniverseUuid

`func (o *DrConfigFailoverForm) GetPrimaryUniverseUuid() string`

GetPrimaryUniverseUuid returns the PrimaryUniverseUuid field if non-nil, zero value otherwise.

### GetPrimaryUniverseUuidOk

`func (o *DrConfigFailoverForm) GetPrimaryUniverseUuidOk() (*string, bool)`

GetPrimaryUniverseUuidOk returns a tuple with the PrimaryUniverseUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryUniverseUuid

`func (o *DrConfigFailoverForm) SetPrimaryUniverseUuid(v string)`

SetPrimaryUniverseUuid sets PrimaryUniverseUuid field to given value.

### HasPrimaryUniverseUuid

`func (o *DrConfigFailoverForm) HasPrimaryUniverseUuid() bool`

HasPrimaryUniverseUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


