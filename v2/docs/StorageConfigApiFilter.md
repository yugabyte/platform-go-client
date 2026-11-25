# StorageConfigApiFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageTypes** | Pointer to [**[]StorageType**](StorageType.md) | Only configurations of these storage types | [optional] 
**Name** | Pointer to **string** | Only the configuration with this name | [optional] 
**States** | Pointer to [**[]StorageConfigState**](StorageConfigState.md) | Only configurations in these states | [optional] 
**InUse** | Pointer to **bool** | Only configurations in use (true) or not in use (false), as StorageConfigUsage.in_use reports it | [optional] 

## Methods

### NewStorageConfigApiFilter

`func NewStorageConfigApiFilter() *StorageConfigApiFilter`

NewStorageConfigApiFilter instantiates a new StorageConfigApiFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageConfigApiFilterWithDefaults

`func NewStorageConfigApiFilterWithDefaults() *StorageConfigApiFilter`

NewStorageConfigApiFilterWithDefaults instantiates a new StorageConfigApiFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageTypes

`func (o *StorageConfigApiFilter) GetStorageTypes() []StorageType`

GetStorageTypes returns the StorageTypes field if non-nil, zero value otherwise.

### GetStorageTypesOk

`func (o *StorageConfigApiFilter) GetStorageTypesOk() (*[]StorageType, bool)`

GetStorageTypesOk returns a tuple with the StorageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTypes

`func (o *StorageConfigApiFilter) SetStorageTypes(v []StorageType)`

SetStorageTypes sets StorageTypes field to given value.

### HasStorageTypes

`func (o *StorageConfigApiFilter) HasStorageTypes() bool`

HasStorageTypes returns a boolean if a field has been set.

### GetName

`func (o *StorageConfigApiFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageConfigApiFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageConfigApiFilter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageConfigApiFilter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStates

`func (o *StorageConfigApiFilter) GetStates() []StorageConfigState`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *StorageConfigApiFilter) GetStatesOk() (*[]StorageConfigState, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *StorageConfigApiFilter) SetStates(v []StorageConfigState)`

SetStates sets States field to given value.

### HasStates

`func (o *StorageConfigApiFilter) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetInUse

`func (o *StorageConfigApiFilter) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *StorageConfigApiFilter) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *StorageConfigApiFilter) SetInUse(v bool)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *StorageConfigApiFilter) HasInUse() bool`

HasInUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


