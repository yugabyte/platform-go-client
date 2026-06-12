# CustomerConfigSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigName** | Pointer to **string** | Configuration name. | [optional] 
**Type** | Pointer to **string** | Configuration category (storage, alerts, call-home, password policy, etc.). | [optional] 
**Name** | Pointer to **string** | Configuration key or subtype within the category. | [optional] 
**Data** | Pointer to **map[string]interface{}** | Masked configuration payload (opaque structure depends on type/name). | [optional] 

## Methods

### NewCustomerConfigSpec

`func NewCustomerConfigSpec() *CustomerConfigSpec`

NewCustomerConfigSpec instantiates a new CustomerConfigSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerConfigSpecWithDefaults

`func NewCustomerConfigSpecWithDefaults() *CustomerConfigSpec`

NewCustomerConfigSpecWithDefaults instantiates a new CustomerConfigSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigName

`func (o *CustomerConfigSpec) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *CustomerConfigSpec) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *CustomerConfigSpec) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.

### HasConfigName

`func (o *CustomerConfigSpec) HasConfigName() bool`

HasConfigName returns a boolean if a field has been set.

### GetType

`func (o *CustomerConfigSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerConfigSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerConfigSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomerConfigSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *CustomerConfigSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerConfigSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerConfigSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomerConfigSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetData

`func (o *CustomerConfigSpec) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CustomerConfigSpec) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CustomerConfigSpec) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *CustomerConfigSpec) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


