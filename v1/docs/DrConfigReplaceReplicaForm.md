# DrConfigReplaceReplicaForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootstrapParams** | Pointer to [**RestartBootstrapParams**](RestartBootstrapParams.md) |  | [optional] 
**DrReplicaUniverseUuid** | Pointer to **string** | New dr replica universe UUID | [optional] 
**PrimaryUniverseUuid** | Pointer to **string** | The current primary universe UUID | [optional] 

## Methods

### NewDrConfigReplaceReplicaForm

`func NewDrConfigReplaceReplicaForm() *DrConfigReplaceReplicaForm`

NewDrConfigReplaceReplicaForm instantiates a new DrConfigReplaceReplicaForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrConfigReplaceReplicaFormWithDefaults

`func NewDrConfigReplaceReplicaFormWithDefaults() *DrConfigReplaceReplicaForm`

NewDrConfigReplaceReplicaFormWithDefaults instantiates a new DrConfigReplaceReplicaForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootstrapParams

`func (o *DrConfigReplaceReplicaForm) GetBootstrapParams() RestartBootstrapParams`

GetBootstrapParams returns the BootstrapParams field if non-nil, zero value otherwise.

### GetBootstrapParamsOk

`func (o *DrConfigReplaceReplicaForm) GetBootstrapParamsOk() (*RestartBootstrapParams, bool)`

GetBootstrapParamsOk returns a tuple with the BootstrapParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapParams

`func (o *DrConfigReplaceReplicaForm) SetBootstrapParams(v RestartBootstrapParams)`

SetBootstrapParams sets BootstrapParams field to given value.

### HasBootstrapParams

`func (o *DrConfigReplaceReplicaForm) HasBootstrapParams() bool`

HasBootstrapParams returns a boolean if a field has been set.

### GetDrReplicaUniverseUuid

`func (o *DrConfigReplaceReplicaForm) GetDrReplicaUniverseUuid() string`

GetDrReplicaUniverseUuid returns the DrReplicaUniverseUuid field if non-nil, zero value otherwise.

### GetDrReplicaUniverseUuidOk

`func (o *DrConfigReplaceReplicaForm) GetDrReplicaUniverseUuidOk() (*string, bool)`

GetDrReplicaUniverseUuidOk returns a tuple with the DrReplicaUniverseUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrReplicaUniverseUuid

`func (o *DrConfigReplaceReplicaForm) SetDrReplicaUniverseUuid(v string)`

SetDrReplicaUniverseUuid sets DrReplicaUniverseUuid field to given value.

### HasDrReplicaUniverseUuid

`func (o *DrConfigReplaceReplicaForm) HasDrReplicaUniverseUuid() bool`

HasDrReplicaUniverseUuid returns a boolean if a field has been set.

### GetPrimaryUniverseUuid

`func (o *DrConfigReplaceReplicaForm) GetPrimaryUniverseUuid() string`

GetPrimaryUniverseUuid returns the PrimaryUniverseUuid field if non-nil, zero value otherwise.

### GetPrimaryUniverseUuidOk

`func (o *DrConfigReplaceReplicaForm) GetPrimaryUniverseUuidOk() (*string, bool)`

GetPrimaryUniverseUuidOk returns a tuple with the PrimaryUniverseUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryUniverseUuid

`func (o *DrConfigReplaceReplicaForm) SetPrimaryUniverseUuid(v string)`

SetPrimaryUniverseUuid sets PrimaryUniverseUuid field to given value.

### HasPrimaryUniverseUuid

`func (o *DrConfigReplaceReplicaForm) HasPrimaryUniverseUuid() bool`

HasPrimaryUniverseUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


