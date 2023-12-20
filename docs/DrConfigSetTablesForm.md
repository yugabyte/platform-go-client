# DrConfigSetTablesForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootstrapParams** | Pointer to [**RestartBootstrapParams**](RestartBootstrapParams.md) |  | [optional] 
**Tables** | Pointer to **[]string** | Source universe table IDs | [optional] 

## Methods

### NewDrConfigSetTablesForm

`func NewDrConfigSetTablesForm() *DrConfigSetTablesForm`

NewDrConfigSetTablesForm instantiates a new DrConfigSetTablesForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrConfigSetTablesFormWithDefaults

`func NewDrConfigSetTablesFormWithDefaults() *DrConfigSetTablesForm`

NewDrConfigSetTablesFormWithDefaults instantiates a new DrConfigSetTablesForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootstrapParams

`func (o *DrConfigSetTablesForm) GetBootstrapParams() RestartBootstrapParams`

GetBootstrapParams returns the BootstrapParams field if non-nil, zero value otherwise.

### GetBootstrapParamsOk

`func (o *DrConfigSetTablesForm) GetBootstrapParamsOk() (*RestartBootstrapParams, bool)`

GetBootstrapParamsOk returns a tuple with the BootstrapParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapParams

`func (o *DrConfigSetTablesForm) SetBootstrapParams(v RestartBootstrapParams)`

SetBootstrapParams sets BootstrapParams field to given value.

### HasBootstrapParams

`func (o *DrConfigSetTablesForm) HasBootstrapParams() bool`

HasBootstrapParams returns a boolean if a field has been set.

### GetTables

`func (o *DrConfigSetTablesForm) GetTables() []string`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *DrConfigSetTablesForm) GetTablesOk() (*[]string, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *DrConfigSetTablesForm) SetTables(v []string)`

SetTables sets Tables field to given value.

### HasTables

`func (o *DrConfigSetTablesForm) HasTables() bool`

HasTables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


