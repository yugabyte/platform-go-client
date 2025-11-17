# CustomerConfigUI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigName** | **string** | Config name | 
**ConfigUUID** | Pointer to **string** | Config UUID | [optional] [readonly] 
**CustomerUUID** | **string** | Customer UUID | [readonly] 
**Data** | **map[string]interface{}** | Configuration data | 
**InUse** | Pointer to **bool** | True if there is an in use reference to the object | [optional] [readonly] 
**KubernetesOperatorControlled** | **bool** |  | 
**Name** | **string** | Name | 
**State** | Pointer to **string** | state of the customerConfig. Possible values are Active, QueuedForDeletion. | [optional] [readonly] 
**Type** | **string** | Config type | 
**UniverseDetails** | Pointer to [**[]UniverseDetailSubset**](UniverseDetailSubset.md) | Universe details | [optional] 

## Methods

### NewCustomerConfigUI

`func NewCustomerConfigUI(configName string, customerUUID string, data map[string]interface{}, kubernetesOperatorControlled bool, name string, type_ string, ) *CustomerConfigUI`

NewCustomerConfigUI instantiates a new CustomerConfigUI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerConfigUIWithDefaults

`func NewCustomerConfigUIWithDefaults() *CustomerConfigUI`

NewCustomerConfigUIWithDefaults instantiates a new CustomerConfigUI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigName

`func (o *CustomerConfigUI) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *CustomerConfigUI) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *CustomerConfigUI) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.


### GetConfigUUID

`func (o *CustomerConfigUI) GetConfigUUID() string`

GetConfigUUID returns the ConfigUUID field if non-nil, zero value otherwise.

### GetConfigUUIDOk

`func (o *CustomerConfigUI) GetConfigUUIDOk() (*string, bool)`

GetConfigUUIDOk returns a tuple with the ConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigUUID

`func (o *CustomerConfigUI) SetConfigUUID(v string)`

SetConfigUUID sets ConfigUUID field to given value.

### HasConfigUUID

`func (o *CustomerConfigUI) HasConfigUUID() bool`

HasConfigUUID returns a boolean if a field has been set.

### GetCustomerUUID

`func (o *CustomerConfigUI) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *CustomerConfigUI) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *CustomerConfigUI) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetData

`func (o *CustomerConfigUI) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CustomerConfigUI) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CustomerConfigUI) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetInUse

`func (o *CustomerConfigUI) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *CustomerConfigUI) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *CustomerConfigUI) SetInUse(v bool)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *CustomerConfigUI) HasInUse() bool`

HasInUse returns a boolean if a field has been set.

### GetKubernetesOperatorControlled

`func (o *CustomerConfigUI) GetKubernetesOperatorControlled() bool`

GetKubernetesOperatorControlled returns the KubernetesOperatorControlled field if non-nil, zero value otherwise.

### GetKubernetesOperatorControlledOk

`func (o *CustomerConfigUI) GetKubernetesOperatorControlledOk() (*bool, bool)`

GetKubernetesOperatorControlledOk returns a tuple with the KubernetesOperatorControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesOperatorControlled

`func (o *CustomerConfigUI) SetKubernetesOperatorControlled(v bool)`

SetKubernetesOperatorControlled sets KubernetesOperatorControlled field to given value.


### GetName

`func (o *CustomerConfigUI) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerConfigUI) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerConfigUI) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *CustomerConfigUI) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CustomerConfigUI) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CustomerConfigUI) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CustomerConfigUI) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *CustomerConfigUI) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerConfigUI) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerConfigUI) SetType(v string)`

SetType sets Type field to given value.


### GetUniverseDetails

`func (o *CustomerConfigUI) GetUniverseDetails() []UniverseDetailSubset`

GetUniverseDetails returns the UniverseDetails field if non-nil, zero value otherwise.

### GetUniverseDetailsOk

`func (o *CustomerConfigUI) GetUniverseDetailsOk() (*[]UniverseDetailSubset, bool)`

GetUniverseDetailsOk returns a tuple with the UniverseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseDetails

`func (o *CustomerConfigUI) SetUniverseDetails(v []UniverseDetailSubset)`

SetUniverseDetails sets UniverseDetails field to given value.

### HasUniverseDetails

`func (o *CustomerConfigUI) HasUniverseDetails() bool`

HasUniverseDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


