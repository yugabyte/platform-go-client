# StorageConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**StorageConfigSpec**](StorageConfigSpec.md) |  | 
**Info** | [**StorageConfigInfo**](StorageConfigInfo.md) |  | 

## Methods

### NewStorageConfig

`func NewStorageConfig(spec StorageConfigSpec, info StorageConfigInfo, ) *StorageConfig`

NewStorageConfig instantiates a new StorageConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageConfigWithDefaults

`func NewStorageConfigWithDefaults() *StorageConfig`

NewStorageConfigWithDefaults instantiates a new StorageConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *StorageConfig) GetSpec() StorageConfigSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *StorageConfig) GetSpecOk() (*StorageConfigSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *StorageConfig) SetSpec(v StorageConfigSpec)`

SetSpec sets Spec field to given value.


### GetInfo

`func (o *StorageConfig) GetInfo() StorageConfigInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *StorageConfig) GetInfoOk() (*StorageConfigInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *StorageConfig) SetInfo(v StorageConfigInfo)`

SetInfo sets Info field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


