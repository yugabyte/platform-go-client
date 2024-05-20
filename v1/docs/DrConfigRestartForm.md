# DrConfigRestartForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootstrapParams** | Pointer to [**RestartBootstrapParams**](RestartBootstrapParams.md) |  | [optional] 
**Dbs** | **[]string** | Primary Universe DB IDs | 

## Methods

### NewDrConfigRestartForm

`func NewDrConfigRestartForm(dbs []string, ) *DrConfigRestartForm`

NewDrConfigRestartForm instantiates a new DrConfigRestartForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrConfigRestartFormWithDefaults

`func NewDrConfigRestartFormWithDefaults() *DrConfigRestartForm`

NewDrConfigRestartFormWithDefaults instantiates a new DrConfigRestartForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootstrapParams

`func (o *DrConfigRestartForm) GetBootstrapParams() RestartBootstrapParams`

GetBootstrapParams returns the BootstrapParams field if non-nil, zero value otherwise.

### GetBootstrapParamsOk

`func (o *DrConfigRestartForm) GetBootstrapParamsOk() (*RestartBootstrapParams, bool)`

GetBootstrapParamsOk returns a tuple with the BootstrapParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapParams

`func (o *DrConfigRestartForm) SetBootstrapParams(v RestartBootstrapParams)`

SetBootstrapParams sets BootstrapParams field to given value.

### HasBootstrapParams

`func (o *DrConfigRestartForm) HasBootstrapParams() bool`

HasBootstrapParams returns a boolean if a field has been set.

### GetDbs

`func (o *DrConfigRestartForm) GetDbs() []string`

GetDbs returns the Dbs field if non-nil, zero value otherwise.

### GetDbsOk

`func (o *DrConfigRestartForm) GetDbsOk() (*[]string, bool)`

GetDbsOk returns a tuple with the Dbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbs

`func (o *DrConfigRestartForm) SetDbs(v []string)`

SetDbs sets Dbs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


