# DrConfigBootstrapBackupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageConfigUuid** | **string** | Storage configuration UUID for bootstrap backup. | 
**Parallelism** | Pointer to **int32** | Number of concurrent commands used by yb_backup (not YBC) on nodes over SSH.  | [optional] [default to 8]

## Methods

### NewDrConfigBootstrapBackupParams

`func NewDrConfigBootstrapBackupParams(storageConfigUuid string, ) *DrConfigBootstrapBackupParams`

NewDrConfigBootstrapBackupParams instantiates a new DrConfigBootstrapBackupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrConfigBootstrapBackupParamsWithDefaults

`func NewDrConfigBootstrapBackupParamsWithDefaults() *DrConfigBootstrapBackupParams`

NewDrConfigBootstrapBackupParamsWithDefaults instantiates a new DrConfigBootstrapBackupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageConfigUuid

`func (o *DrConfigBootstrapBackupParams) GetStorageConfigUuid() string`

GetStorageConfigUuid returns the StorageConfigUuid field if non-nil, zero value otherwise.

### GetStorageConfigUuidOk

`func (o *DrConfigBootstrapBackupParams) GetStorageConfigUuidOk() (*string, bool)`

GetStorageConfigUuidOk returns a tuple with the StorageConfigUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUuid

`func (o *DrConfigBootstrapBackupParams) SetStorageConfigUuid(v string)`

SetStorageConfigUuid sets StorageConfigUuid field to given value.


### GetParallelism

`func (o *DrConfigBootstrapBackupParams) GetParallelism() int32`

GetParallelism returns the Parallelism field if non-nil, zero value otherwise.

### GetParallelismOk

`func (o *DrConfigBootstrapBackupParams) GetParallelismOk() (*int32, bool)`

GetParallelismOk returns a tuple with the Parallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelism

`func (o *DrConfigBootstrapBackupParams) SetParallelism(v int32)`

SetParallelism sets Parallelism field to given value.

### HasParallelism

`func (o *DrConfigBootstrapBackupParams) HasParallelism() bool`

HasParallelism returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


