# UserIntentOverrides

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzOverrides** | Pointer to [**map[string]AZOverrides**](AZOverrides.md) |  | [optional] 
**PerProcess** | Pointer to [**map[string]PerProcessDetails**](PerProcessDetails.md) |  | [optional] 

## Methods

### NewUserIntentOverrides

`func NewUserIntentOverrides() *UserIntentOverrides`

NewUserIntentOverrides instantiates a new UserIntentOverrides object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIntentOverridesWithDefaults

`func NewUserIntentOverridesWithDefaults() *UserIntentOverrides`

NewUserIntentOverridesWithDefaults instantiates a new UserIntentOverrides object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzOverrides

`func (o *UserIntentOverrides) GetAzOverrides() map[string]AZOverrides`

GetAzOverrides returns the AzOverrides field if non-nil, zero value otherwise.

### GetAzOverridesOk

`func (o *UserIntentOverrides) GetAzOverridesOk() (*map[string]AZOverrides, bool)`

GetAzOverridesOk returns a tuple with the AzOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzOverrides

`func (o *UserIntentOverrides) SetAzOverrides(v map[string]AZOverrides)`

SetAzOverrides sets AzOverrides field to given value.

### HasAzOverrides

`func (o *UserIntentOverrides) HasAzOverrides() bool`

HasAzOverrides returns a boolean if a field has been set.

### GetPerProcess

`func (o *UserIntentOverrides) GetPerProcess() map[string]PerProcessDetails`

GetPerProcess returns the PerProcess field if non-nil, zero value otherwise.

### GetPerProcessOk

`func (o *UserIntentOverrides) GetPerProcessOk() (*map[string]PerProcessDetails, bool)`

GetPerProcessOk returns a tuple with the PerProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerProcess

`func (o *UserIntentOverrides) SetPerProcess(v map[string]PerProcessDetails)`

SetPerProcess sets PerProcess field to given value.

### HasPerProcess

`func (o *UserIntentOverrides) HasPerProcess() bool`

HasPerProcess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


