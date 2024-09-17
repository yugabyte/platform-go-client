# AlertTemplateVariablesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomVariables** | [**[]AlertTemplateVariable**](AlertTemplateVariable.md) |  | 
**SystemVariables** | **[]string** |  | 

## Methods

### NewAlertTemplateVariablesList

`func NewAlertTemplateVariablesList(customVariables []AlertTemplateVariable, systemVariables []string, ) *AlertTemplateVariablesList`

NewAlertTemplateVariablesList instantiates a new AlertTemplateVariablesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertTemplateVariablesListWithDefaults

`func NewAlertTemplateVariablesListWithDefaults() *AlertTemplateVariablesList`

NewAlertTemplateVariablesListWithDefaults instantiates a new AlertTemplateVariablesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomVariables

`func (o *AlertTemplateVariablesList) GetCustomVariables() []AlertTemplateVariable`

GetCustomVariables returns the CustomVariables field if non-nil, zero value otherwise.

### GetCustomVariablesOk

`func (o *AlertTemplateVariablesList) GetCustomVariablesOk() (*[]AlertTemplateVariable, bool)`

GetCustomVariablesOk returns a tuple with the CustomVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomVariables

`func (o *AlertTemplateVariablesList) SetCustomVariables(v []AlertTemplateVariable)`

SetCustomVariables sets CustomVariables field to given value.


### GetSystemVariables

`func (o *AlertTemplateVariablesList) GetSystemVariables() []string`

GetSystemVariables returns the SystemVariables field if non-nil, zero value otherwise.

### GetSystemVariablesOk

`func (o *AlertTemplateVariablesList) GetSystemVariablesOk() (*[]string, bool)`

GetSystemVariablesOk returns a tuple with the SystemVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemVariables

`func (o *AlertTemplateVariablesList) SetSystemVariables(v []string)`

SetSystemVariables sets SystemVariables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


