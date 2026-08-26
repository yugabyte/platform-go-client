# PerProviderState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**EarlyoomConfig**](EarlyoomConfig.md) |  | [optional] 
**EarlyoomEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewPerProviderState

`func NewPerProviderState() *PerProviderState`

NewPerProviderState instantiates a new PerProviderState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerProviderStateWithDefaults

`func NewPerProviderStateWithDefaults() *PerProviderState`

NewPerProviderStateWithDefaults instantiates a new PerProviderState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *PerProviderState) GetConfig() EarlyoomConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PerProviderState) GetConfigOk() (*EarlyoomConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PerProviderState) SetConfig(v EarlyoomConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *PerProviderState) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEarlyoomEnabled

`func (o *PerProviderState) GetEarlyoomEnabled() bool`

GetEarlyoomEnabled returns the EarlyoomEnabled field if non-nil, zero value otherwise.

### GetEarlyoomEnabledOk

`func (o *PerProviderState) GetEarlyoomEnabledOk() (*bool, bool)`

GetEarlyoomEnabledOk returns a tuple with the EarlyoomEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyoomEnabled

`func (o *PerProviderState) SetEarlyoomEnabled(v bool)`

SetEarlyoomEnabled sets EarlyoomEnabled field to given value.

### HasEarlyoomEnabled

`func (o *PerProviderState) HasEarlyoomEnabled() bool`

HasEarlyoomEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


