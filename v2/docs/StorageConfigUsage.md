# StorageConfigUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InUse** | **bool** | Whether in-progress or completed backups are stored with the configuration, or a backup schedule (other than one being deleted or failed) or a DR config uses it | [readonly] 
**UniverseUuids** | Pointer to **[]string** | Existing universes whose in-progress or completed backups or backup schedules use the configuration, and the universes of DR configs that use it | [optional] [readonly] 
**ScheduleUuids** | Pointer to **[]string** | Backup schedules that use the configuration, other than those being deleted or failed | [optional] [readonly] 
**DrConfigUuids** | Pointer to **[]string** | DR configs that use the configuration | [optional] [readonly] 

## Methods

### NewStorageConfigUsage

`func NewStorageConfigUsage(inUse bool, ) *StorageConfigUsage`

NewStorageConfigUsage instantiates a new StorageConfigUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageConfigUsageWithDefaults

`func NewStorageConfigUsageWithDefaults() *StorageConfigUsage`

NewStorageConfigUsageWithDefaults instantiates a new StorageConfigUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInUse

`func (o *StorageConfigUsage) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *StorageConfigUsage) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *StorageConfigUsage) SetInUse(v bool)`

SetInUse sets InUse field to given value.


### GetUniverseUuids

`func (o *StorageConfigUsage) GetUniverseUuids() []string`

GetUniverseUuids returns the UniverseUuids field if non-nil, zero value otherwise.

### GetUniverseUuidsOk

`func (o *StorageConfigUsage) GetUniverseUuidsOk() (*[]string, bool)`

GetUniverseUuidsOk returns a tuple with the UniverseUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUuids

`func (o *StorageConfigUsage) SetUniverseUuids(v []string)`

SetUniverseUuids sets UniverseUuids field to given value.

### HasUniverseUuids

`func (o *StorageConfigUsage) HasUniverseUuids() bool`

HasUniverseUuids returns a boolean if a field has been set.

### GetScheduleUuids

`func (o *StorageConfigUsage) GetScheduleUuids() []string`

GetScheduleUuids returns the ScheduleUuids field if non-nil, zero value otherwise.

### GetScheduleUuidsOk

`func (o *StorageConfigUsage) GetScheduleUuidsOk() (*[]string, bool)`

GetScheduleUuidsOk returns a tuple with the ScheduleUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleUuids

`func (o *StorageConfigUsage) SetScheduleUuids(v []string)`

SetScheduleUuids sets ScheduleUuids field to given value.

### HasScheduleUuids

`func (o *StorageConfigUsage) HasScheduleUuids() bool`

HasScheduleUuids returns a boolean if a field has been set.

### GetDrConfigUuids

`func (o *StorageConfigUsage) GetDrConfigUuids() []string`

GetDrConfigUuids returns the DrConfigUuids field if non-nil, zero value otherwise.

### GetDrConfigUuidsOk

`func (o *StorageConfigUsage) GetDrConfigUuidsOk() (*[]string, bool)`

GetDrConfigUuidsOk returns a tuple with the DrConfigUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrConfigUuids

`func (o *StorageConfigUsage) SetDrConfigUuids(v []string)`

SetDrConfigUuids sets DrConfigUuids field to given value.

### HasDrConfigUuids

`func (o *StorageConfigUsage) HasDrConfigUuids() bool`

HasDrConfigUuids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


