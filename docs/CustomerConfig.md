# CustomerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigName** | **string** | Config name | 
**ConfigUUID** | Pointer to **string** | Config UUID | [optional] [readonly] 
**CustomerUUID** | **string** | Customer UUID | [readonly] 
**Data** | **map[string]interface{}** | Configuration data | 
**KubernetesOperatorControlled** | **bool** |  | 
**Name** | **string** | Name | 
**State** | Pointer to **string** | state of the customerConfig. Possible values are Active, QueuedForDeletion. | [optional] [readonly] 
**Type** | **string** | Config type | 

## Methods

### NewCustomerConfig

`func NewCustomerConfig(configName string, customerUUID string, data map[string]interface{}, kubernetesOperatorControlled bool, name string, type_ string, ) *CustomerConfig`

NewCustomerConfig instantiates a new CustomerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerConfigWithDefaults

`func NewCustomerConfigWithDefaults() *CustomerConfig`

NewCustomerConfigWithDefaults instantiates a new CustomerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigName

`func (o *CustomerConfig) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *CustomerConfig) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *CustomerConfig) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.


### GetConfigUUID

`func (o *CustomerConfig) GetConfigUUID() string`

GetConfigUUID returns the ConfigUUID field if non-nil, zero value otherwise.

### GetConfigUUIDOk

`func (o *CustomerConfig) GetConfigUUIDOk() (*string, bool)`

GetConfigUUIDOk returns a tuple with the ConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigUUID

`func (o *CustomerConfig) SetConfigUUID(v string)`

SetConfigUUID sets ConfigUUID field to given value.

### HasConfigUUID

`func (o *CustomerConfig) HasConfigUUID() bool`

HasConfigUUID returns a boolean if a field has been set.

### GetCustomerUUID

`func (o *CustomerConfig) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *CustomerConfig) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *CustomerConfig) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetData

`func (o *CustomerConfig) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CustomerConfig) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CustomerConfig) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetKubernetesOperatorControlled

`func (o *CustomerConfig) GetKubernetesOperatorControlled() bool`

GetKubernetesOperatorControlled returns the KubernetesOperatorControlled field if non-nil, zero value otherwise.

### GetKubernetesOperatorControlledOk

`func (o *CustomerConfig) GetKubernetesOperatorControlledOk() (*bool, bool)`

GetKubernetesOperatorControlledOk returns a tuple with the KubernetesOperatorControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesOperatorControlled

`func (o *CustomerConfig) SetKubernetesOperatorControlled(v bool)`

SetKubernetesOperatorControlled sets KubernetesOperatorControlled field to given value.


### GetName

`func (o *CustomerConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerConfig) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *CustomerConfig) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CustomerConfig) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CustomerConfig) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CustomerConfig) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *CustomerConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerConfig) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


