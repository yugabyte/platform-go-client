# StorageConfigSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the storage configuration, unique within the customer | 
**Storage** | [**StorageSpec**](StorageSpec.md) |  | 

## Methods

### NewStorageConfigSpec

`func NewStorageConfigSpec(name string, storage StorageSpec, ) *StorageConfigSpec`

NewStorageConfigSpec instantiates a new StorageConfigSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageConfigSpecWithDefaults

`func NewStorageConfigSpecWithDefaults() *StorageConfigSpec`

NewStorageConfigSpecWithDefaults instantiates a new StorageConfigSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StorageConfigSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageConfigSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageConfigSpec) SetName(v string)`

SetName sets Name field to given value.


### GetStorage

`func (o *StorageConfigSpec) GetStorage() StorageSpec`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *StorageConfigSpec) GetStorageOk() (*StorageSpec, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *StorageConfigSpec) SetStorage(v StorageSpec)`

SetStorage sets Storage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


