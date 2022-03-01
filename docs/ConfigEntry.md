# ConfigEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inherited** | Pointer to **bool** | Is this configuration inherited? | [optional] 
**Key** | Pointer to **string** | Configuration key | [optional] 
**Value** | Pointer to **string** | Configuration value | [optional] 

## Methods

### NewConfigEntry

`func NewConfigEntry() *ConfigEntry`

NewConfigEntry instantiates a new ConfigEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigEntryWithDefaults

`func NewConfigEntryWithDefaults() *ConfigEntry`

NewConfigEntryWithDefaults instantiates a new ConfigEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInherited

`func (o *ConfigEntry) GetInherited() bool`

GetInherited returns the Inherited field if non-nil, zero value otherwise.

### GetInheritedOk

`func (o *ConfigEntry) GetInheritedOk() (*bool, bool)`

GetInheritedOk returns a tuple with the Inherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInherited

`func (o *ConfigEntry) SetInherited(v bool)`

SetInherited sets Inherited field to given value.

### HasInherited

`func (o *ConfigEntry) HasInherited() bool`

HasInherited returns a boolean if a field has been set.

### GetKey

`func (o *ConfigEntry) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ConfigEntry) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ConfigEntry) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ConfigEntry) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *ConfigEntry) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConfigEntry) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConfigEntry) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ConfigEntry) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


