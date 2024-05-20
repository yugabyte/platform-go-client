# XClusterConfigRestartFormData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootstrapParams** | Pointer to [**RestartBootstrapParams**](RestartBootstrapParams.md) |  | [optional] 
**DryRun** | Pointer to **bool** | Run the pre-checks without actually running the subtasks | [optional] 
**Tables** | **[]string** | Source Universe table IDs; if empty, the whole config will restart | 

## Methods

### NewXClusterConfigRestartFormData

`func NewXClusterConfigRestartFormData(tables []string, ) *XClusterConfigRestartFormData`

NewXClusterConfigRestartFormData instantiates a new XClusterConfigRestartFormData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXClusterConfigRestartFormDataWithDefaults

`func NewXClusterConfigRestartFormDataWithDefaults() *XClusterConfigRestartFormData`

NewXClusterConfigRestartFormDataWithDefaults instantiates a new XClusterConfigRestartFormData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootstrapParams

`func (o *XClusterConfigRestartFormData) GetBootstrapParams() RestartBootstrapParams`

GetBootstrapParams returns the BootstrapParams field if non-nil, zero value otherwise.

### GetBootstrapParamsOk

`func (o *XClusterConfigRestartFormData) GetBootstrapParamsOk() (*RestartBootstrapParams, bool)`

GetBootstrapParamsOk returns a tuple with the BootstrapParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapParams

`func (o *XClusterConfigRestartFormData) SetBootstrapParams(v RestartBootstrapParams)`

SetBootstrapParams sets BootstrapParams field to given value.

### HasBootstrapParams

`func (o *XClusterConfigRestartFormData) HasBootstrapParams() bool`

HasBootstrapParams returns a boolean if a field has been set.

### GetDryRun

`func (o *XClusterConfigRestartFormData) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *XClusterConfigRestartFormData) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *XClusterConfigRestartFormData) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *XClusterConfigRestartFormData) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetTables

`func (o *XClusterConfigRestartFormData) GetTables() []string`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *XClusterConfigRestartFormData) GetTablesOk() (*[]string, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *XClusterConfigRestartFormData) SetTables(v []string)`

SetTables sets Tables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


