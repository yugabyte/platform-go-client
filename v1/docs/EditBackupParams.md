# EditBackupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiryTimeUnit** | Pointer to **string** | Time unit for backup expiry | [optional] 
**StorageConfigUUID** | Pointer to **string** | New backup Storage config | [optional] 
**TimeBeforeDeleteFromPresentInMillis** | Pointer to **int64** | Time before deleting the backup from storage, in milliseconds | [optional] 

## Methods

### NewEditBackupParams

`func NewEditBackupParams() *EditBackupParams`

NewEditBackupParams instantiates a new EditBackupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditBackupParamsWithDefaults

`func NewEditBackupParamsWithDefaults() *EditBackupParams`

NewEditBackupParamsWithDefaults instantiates a new EditBackupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiryTimeUnit

`func (o *EditBackupParams) GetExpiryTimeUnit() string`

GetExpiryTimeUnit returns the ExpiryTimeUnit field if non-nil, zero value otherwise.

### GetExpiryTimeUnitOk

`func (o *EditBackupParams) GetExpiryTimeUnitOk() (*string, bool)`

GetExpiryTimeUnitOk returns a tuple with the ExpiryTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTimeUnit

`func (o *EditBackupParams) SetExpiryTimeUnit(v string)`

SetExpiryTimeUnit sets ExpiryTimeUnit field to given value.

### HasExpiryTimeUnit

`func (o *EditBackupParams) HasExpiryTimeUnit() bool`

HasExpiryTimeUnit returns a boolean if a field has been set.

### GetStorageConfigUUID

`func (o *EditBackupParams) GetStorageConfigUUID() string`

GetStorageConfigUUID returns the StorageConfigUUID field if non-nil, zero value otherwise.

### GetStorageConfigUUIDOk

`func (o *EditBackupParams) GetStorageConfigUUIDOk() (*string, bool)`

GetStorageConfigUUIDOk returns a tuple with the StorageConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUUID

`func (o *EditBackupParams) SetStorageConfigUUID(v string)`

SetStorageConfigUUID sets StorageConfigUUID field to given value.

### HasStorageConfigUUID

`func (o *EditBackupParams) HasStorageConfigUUID() bool`

HasStorageConfigUUID returns a boolean if a field has been set.

### GetTimeBeforeDeleteFromPresentInMillis

`func (o *EditBackupParams) GetTimeBeforeDeleteFromPresentInMillis() int64`

GetTimeBeforeDeleteFromPresentInMillis returns the TimeBeforeDeleteFromPresentInMillis field if non-nil, zero value otherwise.

### GetTimeBeforeDeleteFromPresentInMillisOk

`func (o *EditBackupParams) GetTimeBeforeDeleteFromPresentInMillisOk() (*int64, bool)`

GetTimeBeforeDeleteFromPresentInMillisOk returns a tuple with the TimeBeforeDeleteFromPresentInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBeforeDeleteFromPresentInMillis

`func (o *EditBackupParams) SetTimeBeforeDeleteFromPresentInMillis(v int64)`

SetTimeBeforeDeleteFromPresentInMillis sets TimeBeforeDeleteFromPresentInMillis field to given value.

### HasTimeBeforeDeleteFromPresentInMillis

`func (o *EditBackupParams) HasTimeBeforeDeleteFromPresentInMillis() bool`

HasTimeBeforeDeleteFromPresentInMillis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


