# ScopedConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigEntries** | Pointer to [**[]ConfigEntry**](ConfigEntry.md) | List of configurations | [optional] 
**MutableScope** | Pointer to **bool** | Mutability of the scope. Only super admin users can change global scope; other scopes are customer-mutable. | [optional] 
**Type** | Pointer to **string** | Scope type | [optional] 
**Uuid** | Pointer to **string** | Scope UIID | [optional] 

## Methods

### NewScopedConfig

`func NewScopedConfig() *ScopedConfig`

NewScopedConfig instantiates a new ScopedConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopedConfigWithDefaults

`func NewScopedConfigWithDefaults() *ScopedConfig`

NewScopedConfigWithDefaults instantiates a new ScopedConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigEntries

`func (o *ScopedConfig) GetConfigEntries() []ConfigEntry`

GetConfigEntries returns the ConfigEntries field if non-nil, zero value otherwise.

### GetConfigEntriesOk

`func (o *ScopedConfig) GetConfigEntriesOk() (*[]ConfigEntry, bool)`

GetConfigEntriesOk returns a tuple with the ConfigEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigEntries

`func (o *ScopedConfig) SetConfigEntries(v []ConfigEntry)`

SetConfigEntries sets ConfigEntries field to given value.

### HasConfigEntries

`func (o *ScopedConfig) HasConfigEntries() bool`

HasConfigEntries returns a boolean if a field has been set.

### GetMutableScope

`func (o *ScopedConfig) GetMutableScope() bool`

GetMutableScope returns the MutableScope field if non-nil, zero value otherwise.

### GetMutableScopeOk

`func (o *ScopedConfig) GetMutableScopeOk() (*bool, bool)`

GetMutableScopeOk returns a tuple with the MutableScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutableScope

`func (o *ScopedConfig) SetMutableScope(v bool)`

SetMutableScope sets MutableScope field to given value.

### HasMutableScope

`func (o *ScopedConfig) HasMutableScope() bool`

HasMutableScope returns a boolean if a field has been set.

### GetType

`func (o *ScopedConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScopedConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScopedConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ScopedConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *ScopedConfig) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ScopedConfig) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ScopedConfig) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ScopedConfig) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


