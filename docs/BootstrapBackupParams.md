# BootstrapBackupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parallelism** | Pointer to **int32** | Number of concurrent commands used by yb_backup (not ybc) to run on nodes over SSH | [optional] 
**StorageConfigUUID** | **string** | Storage configuration UUID | 

## Methods

### NewBootstrapBackupParams

`func NewBootstrapBackupParams(storageConfigUUID string, ) *BootstrapBackupParams`

NewBootstrapBackupParams instantiates a new BootstrapBackupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootstrapBackupParamsWithDefaults

`func NewBootstrapBackupParamsWithDefaults() *BootstrapBackupParams`

NewBootstrapBackupParamsWithDefaults instantiates a new BootstrapBackupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParallelism

`func (o *BootstrapBackupParams) GetParallelism() int32`

GetParallelism returns the Parallelism field if non-nil, zero value otherwise.

### GetParallelismOk

`func (o *BootstrapBackupParams) GetParallelismOk() (*int32, bool)`

GetParallelismOk returns a tuple with the Parallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelism

`func (o *BootstrapBackupParams) SetParallelism(v int32)`

SetParallelism sets Parallelism field to given value.

### HasParallelism

`func (o *BootstrapBackupParams) HasParallelism() bool`

HasParallelism returns a boolean if a field has been set.

### GetStorageConfigUUID

`func (o *BootstrapBackupParams) GetStorageConfigUUID() string`

GetStorageConfigUUID returns the StorageConfigUUID field if non-nil, zero value otherwise.

### GetStorageConfigUUIDOk

`func (o *BootstrapBackupParams) GetStorageConfigUUIDOk() (*string, bool)`

GetStorageConfigUUIDOk returns a tuple with the StorageConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUUID

`func (o *BootstrapBackupParams) SetStorageConfigUUID(v string)`

SetStorageConfigUUID sets StorageConfigUUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


