# UniverseBackupRequestFormData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorString** | Pointer to **string** | Error message | [optional] 
**PlatformVersion** | **string** |  | 
**PreviousTaskUUID** | Pointer to **string** | Previous task UUID of a retry | [optional] 
**StorageConfigUUID** | **string** | WARNING: This is a preview API that could change.Storage configuration UUID | 
**TimeBeforeDelete** | Pointer to **int64** | WARNING: This is a preview API that could change.Time before deleting the backup from storage, in milliseconds | [optional] 

## Methods

### NewUniverseBackupRequestFormData

`func NewUniverseBackupRequestFormData(platformVersion string, storageConfigUUID string, ) *UniverseBackupRequestFormData`

NewUniverseBackupRequestFormData instantiates a new UniverseBackupRequestFormData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseBackupRequestFormDataWithDefaults

`func NewUniverseBackupRequestFormDataWithDefaults() *UniverseBackupRequestFormData`

NewUniverseBackupRequestFormDataWithDefaults instantiates a new UniverseBackupRequestFormData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorString

`func (o *UniverseBackupRequestFormData) GetErrorString() string`

GetErrorString returns the ErrorString field if non-nil, zero value otherwise.

### GetErrorStringOk

`func (o *UniverseBackupRequestFormData) GetErrorStringOk() (*string, bool)`

GetErrorStringOk returns a tuple with the ErrorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorString

`func (o *UniverseBackupRequestFormData) SetErrorString(v string)`

SetErrorString sets ErrorString field to given value.

### HasErrorString

`func (o *UniverseBackupRequestFormData) HasErrorString() bool`

HasErrorString returns a boolean if a field has been set.

### GetPlatformVersion

`func (o *UniverseBackupRequestFormData) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *UniverseBackupRequestFormData) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *UniverseBackupRequestFormData) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.


### GetPreviousTaskUUID

`func (o *UniverseBackupRequestFormData) GetPreviousTaskUUID() string`

GetPreviousTaskUUID returns the PreviousTaskUUID field if non-nil, zero value otherwise.

### GetPreviousTaskUUIDOk

`func (o *UniverseBackupRequestFormData) GetPreviousTaskUUIDOk() (*string, bool)`

GetPreviousTaskUUIDOk returns a tuple with the PreviousTaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTaskUUID

`func (o *UniverseBackupRequestFormData) SetPreviousTaskUUID(v string)`

SetPreviousTaskUUID sets PreviousTaskUUID field to given value.

### HasPreviousTaskUUID

`func (o *UniverseBackupRequestFormData) HasPreviousTaskUUID() bool`

HasPreviousTaskUUID returns a boolean if a field has been set.

### GetStorageConfigUUID

`func (o *UniverseBackupRequestFormData) GetStorageConfigUUID() string`

GetStorageConfigUUID returns the StorageConfigUUID field if non-nil, zero value otherwise.

### GetStorageConfigUUIDOk

`func (o *UniverseBackupRequestFormData) GetStorageConfigUUIDOk() (*string, bool)`

GetStorageConfigUUIDOk returns a tuple with the StorageConfigUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageConfigUUID

`func (o *UniverseBackupRequestFormData) SetStorageConfigUUID(v string)`

SetStorageConfigUUID sets StorageConfigUUID field to given value.


### GetTimeBeforeDelete

`func (o *UniverseBackupRequestFormData) GetTimeBeforeDelete() int64`

GetTimeBeforeDelete returns the TimeBeforeDelete field if non-nil, zero value otherwise.

### GetTimeBeforeDeleteOk

`func (o *UniverseBackupRequestFormData) GetTimeBeforeDeleteOk() (*int64, bool)`

GetTimeBeforeDeleteOk returns a tuple with the TimeBeforeDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBeforeDelete

`func (o *UniverseBackupRequestFormData) SetTimeBeforeDelete(v int64)`

SetTimeBeforeDelete sets TimeBeforeDelete field to given value.

### HasTimeBeforeDelete

`func (o *UniverseBackupRequestFormData) HasTimeBeforeDelete() bool`

HasTimeBeforeDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


