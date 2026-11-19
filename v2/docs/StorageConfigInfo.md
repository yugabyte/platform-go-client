# StorageConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | UUID of the storage configuration | [readonly] 
**State** | [**StorageConfigState**](StorageConfigState.md) |  | 
**IsKubernetesOperatorControlled** | Pointer to **bool** | Whether the YBA Kubernetes operator manages this configuration through a StorageConfig resource. Change such a configuration through its resource instead. | [optional] [readonly] 

## Methods

### NewStorageConfigInfo

`func NewStorageConfigInfo(uuid string, state StorageConfigState, ) *StorageConfigInfo`

NewStorageConfigInfo instantiates a new StorageConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageConfigInfoWithDefaults

`func NewStorageConfigInfoWithDefaults() *StorageConfigInfo`

NewStorageConfigInfoWithDefaults instantiates a new StorageConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *StorageConfigInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageConfigInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageConfigInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetState

`func (o *StorageConfigInfo) GetState() StorageConfigState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageConfigInfo) GetStateOk() (*StorageConfigState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageConfigInfo) SetState(v StorageConfigState)`

SetState sets State field to given value.


### GetIsKubernetesOperatorControlled

`func (o *StorageConfigInfo) GetIsKubernetesOperatorControlled() bool`

GetIsKubernetesOperatorControlled returns the IsKubernetesOperatorControlled field if non-nil, zero value otherwise.

### GetIsKubernetesOperatorControlledOk

`func (o *StorageConfigInfo) GetIsKubernetesOperatorControlledOk() (*bool, bool)`

GetIsKubernetesOperatorControlledOk returns a tuple with the IsKubernetesOperatorControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKubernetesOperatorControlled

`func (o *StorageConfigInfo) SetIsKubernetesOperatorControlled(v bool)`

SetIsKubernetesOperatorControlled sets IsKubernetesOperatorControlled field to given value.

### HasIsKubernetesOperatorControlled

`func (o *StorageConfigInfo) HasIsKubernetesOperatorControlled() bool`

HasIsKubernetesOperatorControlled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


