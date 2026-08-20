# RestoreInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Restore UUID. | [optional] [readonly] 
**CustomerUuid** | Pointer to **string** | Customer UUID that owns the restore. | [optional] [readonly] 
**UniverseUuid** | Pointer to **string** | Target universe UUID that was restored into. | [optional] [readonly] 
**State** | Pointer to [**RestoreState**](RestoreState.md) |  | [optional] 
**UniverseName** | Pointer to **string** | Target universe name. | [optional] [readonly] 
**SourceUniverseUuid** | Pointer to **string** | Source universe UUID the backup was taken from. | [optional] [readonly] 
**SourceUniverseName** | Pointer to **string** | Source universe name the backup was taken from. | [optional] [readonly] 
**BackupType** | Pointer to [**TableType**](TableType.md) |  | [optional] 
**BackupCreatedOnDate** | Pointer to **time.Time** | Creation time of the backup that was restored. | [optional] [readonly] 
**CreateTime** | Pointer to **time.Time** | Time when the restore was created. | [optional] [readonly] 
**UpdateTime** | Pointer to **time.Time** | Time when the restore was last updated. | [optional] [readonly] 
**RestoreSizeInBytes** | Pointer to **int64** | Total restored size, in bytes. | [optional] [readonly] 
**IsSourceUniversePresent** | Pointer to **bool** | Whether the source universe still exists. | [optional] [readonly] 

## Methods

### NewRestoreInfo

`func NewRestoreInfo() *RestoreInfo`

NewRestoreInfo instantiates a new RestoreInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreInfoWithDefaults

`func NewRestoreInfoWithDefaults() *RestoreInfo`

NewRestoreInfoWithDefaults instantiates a new RestoreInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *RestoreInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RestoreInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RestoreInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RestoreInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCustomerUuid

`func (o *RestoreInfo) GetCustomerUuid() string`

GetCustomerUuid returns the CustomerUuid field if non-nil, zero value otherwise.

### GetCustomerUuidOk

`func (o *RestoreInfo) GetCustomerUuidOk() (*string, bool)`

GetCustomerUuidOk returns a tuple with the CustomerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUuid

`func (o *RestoreInfo) SetCustomerUuid(v string)`

SetCustomerUuid sets CustomerUuid field to given value.

### HasCustomerUuid

`func (o *RestoreInfo) HasCustomerUuid() bool`

HasCustomerUuid returns a boolean if a field has been set.

### GetUniverseUuid

`func (o *RestoreInfo) GetUniverseUuid() string`

GetUniverseUuid returns the UniverseUuid field if non-nil, zero value otherwise.

### GetUniverseUuidOk

`func (o *RestoreInfo) GetUniverseUuidOk() (*string, bool)`

GetUniverseUuidOk returns a tuple with the UniverseUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUuid

`func (o *RestoreInfo) SetUniverseUuid(v string)`

SetUniverseUuid sets UniverseUuid field to given value.

### HasUniverseUuid

`func (o *RestoreInfo) HasUniverseUuid() bool`

HasUniverseUuid returns a boolean if a field has been set.

### GetState

`func (o *RestoreInfo) GetState() RestoreState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RestoreInfo) GetStateOk() (*RestoreState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RestoreInfo) SetState(v RestoreState)`

SetState sets State field to given value.

### HasState

`func (o *RestoreInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUniverseName

`func (o *RestoreInfo) GetUniverseName() string`

GetUniverseName returns the UniverseName field if non-nil, zero value otherwise.

### GetUniverseNameOk

`func (o *RestoreInfo) GetUniverseNameOk() (*string, bool)`

GetUniverseNameOk returns a tuple with the UniverseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseName

`func (o *RestoreInfo) SetUniverseName(v string)`

SetUniverseName sets UniverseName field to given value.

### HasUniverseName

`func (o *RestoreInfo) HasUniverseName() bool`

HasUniverseName returns a boolean if a field has been set.

### GetSourceUniverseUuid

`func (o *RestoreInfo) GetSourceUniverseUuid() string`

GetSourceUniverseUuid returns the SourceUniverseUuid field if non-nil, zero value otherwise.

### GetSourceUniverseUuidOk

`func (o *RestoreInfo) GetSourceUniverseUuidOk() (*string, bool)`

GetSourceUniverseUuidOk returns a tuple with the SourceUniverseUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUniverseUuid

`func (o *RestoreInfo) SetSourceUniverseUuid(v string)`

SetSourceUniverseUuid sets SourceUniverseUuid field to given value.

### HasSourceUniverseUuid

`func (o *RestoreInfo) HasSourceUniverseUuid() bool`

HasSourceUniverseUuid returns a boolean if a field has been set.

### GetSourceUniverseName

`func (o *RestoreInfo) GetSourceUniverseName() string`

GetSourceUniverseName returns the SourceUniverseName field if non-nil, zero value otherwise.

### GetSourceUniverseNameOk

`func (o *RestoreInfo) GetSourceUniverseNameOk() (*string, bool)`

GetSourceUniverseNameOk returns a tuple with the SourceUniverseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUniverseName

`func (o *RestoreInfo) SetSourceUniverseName(v string)`

SetSourceUniverseName sets SourceUniverseName field to given value.

### HasSourceUniverseName

`func (o *RestoreInfo) HasSourceUniverseName() bool`

HasSourceUniverseName returns a boolean if a field has been set.

### GetBackupType

`func (o *RestoreInfo) GetBackupType() TableType`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *RestoreInfo) GetBackupTypeOk() (*TableType, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *RestoreInfo) SetBackupType(v TableType)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *RestoreInfo) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetBackupCreatedOnDate

`func (o *RestoreInfo) GetBackupCreatedOnDate() time.Time`

GetBackupCreatedOnDate returns the BackupCreatedOnDate field if non-nil, zero value otherwise.

### GetBackupCreatedOnDateOk

`func (o *RestoreInfo) GetBackupCreatedOnDateOk() (*time.Time, bool)`

GetBackupCreatedOnDateOk returns a tuple with the BackupCreatedOnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCreatedOnDate

`func (o *RestoreInfo) SetBackupCreatedOnDate(v time.Time)`

SetBackupCreatedOnDate sets BackupCreatedOnDate field to given value.

### HasBackupCreatedOnDate

`func (o *RestoreInfo) HasBackupCreatedOnDate() bool`

HasBackupCreatedOnDate returns a boolean if a field has been set.

### GetCreateTime

`func (o *RestoreInfo) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *RestoreInfo) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *RestoreInfo) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *RestoreInfo) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *RestoreInfo) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *RestoreInfo) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *RestoreInfo) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *RestoreInfo) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetRestoreSizeInBytes

`func (o *RestoreInfo) GetRestoreSizeInBytes() int64`

GetRestoreSizeInBytes returns the RestoreSizeInBytes field if non-nil, zero value otherwise.

### GetRestoreSizeInBytesOk

`func (o *RestoreInfo) GetRestoreSizeInBytesOk() (*int64, bool)`

GetRestoreSizeInBytesOk returns a tuple with the RestoreSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreSizeInBytes

`func (o *RestoreInfo) SetRestoreSizeInBytes(v int64)`

SetRestoreSizeInBytes sets RestoreSizeInBytes field to given value.

### HasRestoreSizeInBytes

`func (o *RestoreInfo) HasRestoreSizeInBytes() bool`

HasRestoreSizeInBytes returns a boolean if a field has been set.

### GetIsSourceUniversePresent

`func (o *RestoreInfo) GetIsSourceUniversePresent() bool`

GetIsSourceUniversePresent returns the IsSourceUniversePresent field if non-nil, zero value otherwise.

### GetIsSourceUniversePresentOk

`func (o *RestoreInfo) GetIsSourceUniversePresentOk() (*bool, bool)`

GetIsSourceUniversePresentOk returns a tuple with the IsSourceUniversePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSourceUniversePresent

`func (o *RestoreInfo) SetIsSourceUniversePresent(v bool)`

SetIsSourceUniversePresent sets IsSourceUniversePresent field to given value.

### HasIsSourceUniversePresent

`func (o *RestoreInfo) HasIsSourceUniversePresent() bool`

HasIsSourceUniversePresent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


