# CustomerConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Unique identifier of the customer configuration row. | [optional] [readonly] 
**CustomerUuid** | Pointer to **string** | Customer that owns this configuration. | [optional] [readonly] 
**State** | Pointer to **string** | Lifecycle state of the configuration. | [optional] [readonly] 
**IsKubernetesOperatorControlled** | Pointer to **bool** | Whether this configuration is managed by the Kubernetes operator. | [optional] [readonly] 

## Methods

### NewCustomerConfigInfo

`func NewCustomerConfigInfo() *CustomerConfigInfo`

NewCustomerConfigInfo instantiates a new CustomerConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerConfigInfoWithDefaults

`func NewCustomerConfigInfoWithDefaults() *CustomerConfigInfo`

NewCustomerConfigInfoWithDefaults instantiates a new CustomerConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *CustomerConfigInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CustomerConfigInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CustomerConfigInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CustomerConfigInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCustomerUuid

`func (o *CustomerConfigInfo) GetCustomerUuid() string`

GetCustomerUuid returns the CustomerUuid field if non-nil, zero value otherwise.

### GetCustomerUuidOk

`func (o *CustomerConfigInfo) GetCustomerUuidOk() (*string, bool)`

GetCustomerUuidOk returns a tuple with the CustomerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUuid

`func (o *CustomerConfigInfo) SetCustomerUuid(v string)`

SetCustomerUuid sets CustomerUuid field to given value.

### HasCustomerUuid

`func (o *CustomerConfigInfo) HasCustomerUuid() bool`

HasCustomerUuid returns a boolean if a field has been set.

### GetState

`func (o *CustomerConfigInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CustomerConfigInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CustomerConfigInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CustomerConfigInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetIsKubernetesOperatorControlled

`func (o *CustomerConfigInfo) GetIsKubernetesOperatorControlled() bool`

GetIsKubernetesOperatorControlled returns the IsKubernetesOperatorControlled field if non-nil, zero value otherwise.

### GetIsKubernetesOperatorControlledOk

`func (o *CustomerConfigInfo) GetIsKubernetesOperatorControlledOk() (*bool, bool)`

GetIsKubernetesOperatorControlledOk returns a tuple with the IsKubernetesOperatorControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKubernetesOperatorControlled

`func (o *CustomerConfigInfo) SetIsKubernetesOperatorControlled(v bool)`

SetIsKubernetesOperatorControlled sets IsKubernetesOperatorControlled field to given value.

### HasIsKubernetesOperatorControlled

`func (o *CustomerConfigInfo) HasIsKubernetesOperatorControlled() bool`

HasIsKubernetesOperatorControlled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


