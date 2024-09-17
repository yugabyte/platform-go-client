# AlertTemplateVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerUUID** | **string** | Customer UUID | [readonly] 
**DefaultValue** | **string** | Default value | 
**Name** | **string** | Name | 
**PossibleValues** | **[]string** | Possible values | 
**Uuid** | Pointer to **string** | Variable UUID | [optional] [readonly] 

## Methods

### NewAlertTemplateVariable

`func NewAlertTemplateVariable(customerUUID string, defaultValue string, name string, possibleValues []string, ) *AlertTemplateVariable`

NewAlertTemplateVariable instantiates a new AlertTemplateVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertTemplateVariableWithDefaults

`func NewAlertTemplateVariableWithDefaults() *AlertTemplateVariable`

NewAlertTemplateVariableWithDefaults instantiates a new AlertTemplateVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerUUID

`func (o *AlertTemplateVariable) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *AlertTemplateVariable) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *AlertTemplateVariable) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetDefaultValue

`func (o *AlertTemplateVariable) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *AlertTemplateVariable) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *AlertTemplateVariable) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.


### GetName

`func (o *AlertTemplateVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertTemplateVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertTemplateVariable) SetName(v string)`

SetName sets Name field to given value.


### GetPossibleValues

`func (o *AlertTemplateVariable) GetPossibleValues() []string`

GetPossibleValues returns the PossibleValues field if non-nil, zero value otherwise.

### GetPossibleValuesOk

`func (o *AlertTemplateVariable) GetPossibleValuesOk() (*[]string, bool)`

GetPossibleValuesOk returns a tuple with the PossibleValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleValues

`func (o *AlertTemplateVariable) SetPossibleValues(v []string)`

SetPossibleValues sets PossibleValues field to given value.


### GetUuid

`func (o *AlertTemplateVariable) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AlertTemplateVariable) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AlertTemplateVariable) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AlertTemplateVariable) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


