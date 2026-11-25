# StorageConfigPagedResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]StorageConfig**](StorageConfig.md) | Storage configurations for this page. | [optional] 

## Methods

### NewStorageConfigPagedResp

`func NewStorageConfigPagedResp() *StorageConfigPagedResp`

NewStorageConfigPagedResp instantiates a new StorageConfigPagedResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageConfigPagedRespWithDefaults

`func NewStorageConfigPagedRespWithDefaults() *StorageConfigPagedResp`

NewStorageConfigPagedRespWithDefaults instantiates a new StorageConfigPagedResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *StorageConfigPagedResp) GetEntities() []StorageConfig`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *StorageConfigPagedResp) GetEntitiesOk() (*[]StorageConfig, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *StorageConfigPagedResp) SetEntities(v []StorageConfig)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *StorageConfigPagedResp) HasEntities() bool`

HasEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


