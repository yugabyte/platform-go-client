# DrConfigDbDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceNamespaceId** | Pointer to **string** | Source namespace ID. | [optional] [readonly] 
**Status** | Pointer to [**DrConfigReplicationDetailStatus**](DrConfigReplicationDetailStatus.md) |  | [optional] 
**ReplicationSetupTime** | Pointer to **time.Time** | Time when replication was set up for this database. | [optional] [readonly] 
**BackupUuid** | Pointer to **string** | UUID of the backup used for bootstrapping this database. | [optional] [readonly] 
**RestoreUuid** | Pointer to **string** | UUID of the restore used for bootstrapping this database. | [optional] [readonly] 
**SourceNamespaceInfo** | Pointer to [**NamespaceInfo**](NamespaceInfo.md) |  | [optional] 
**TargetNamespaceInfo** | Pointer to [**NamespaceInfo**](NamespaceInfo.md) |  | [optional] 

## Methods

### NewDrConfigDbDetail

`func NewDrConfigDbDetail() *DrConfigDbDetail`

NewDrConfigDbDetail instantiates a new DrConfigDbDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrConfigDbDetailWithDefaults

`func NewDrConfigDbDetailWithDefaults() *DrConfigDbDetail`

NewDrConfigDbDetailWithDefaults instantiates a new DrConfigDbDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceNamespaceId

`func (o *DrConfigDbDetail) GetSourceNamespaceId() string`

GetSourceNamespaceId returns the SourceNamespaceId field if non-nil, zero value otherwise.

### GetSourceNamespaceIdOk

`func (o *DrConfigDbDetail) GetSourceNamespaceIdOk() (*string, bool)`

GetSourceNamespaceIdOk returns a tuple with the SourceNamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNamespaceId

`func (o *DrConfigDbDetail) SetSourceNamespaceId(v string)`

SetSourceNamespaceId sets SourceNamespaceId field to given value.

### HasSourceNamespaceId

`func (o *DrConfigDbDetail) HasSourceNamespaceId() bool`

HasSourceNamespaceId returns a boolean if a field has been set.

### GetStatus

`func (o *DrConfigDbDetail) GetStatus() DrConfigReplicationDetailStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DrConfigDbDetail) GetStatusOk() (*DrConfigReplicationDetailStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DrConfigDbDetail) SetStatus(v DrConfigReplicationDetailStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DrConfigDbDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReplicationSetupTime

`func (o *DrConfigDbDetail) GetReplicationSetupTime() time.Time`

GetReplicationSetupTime returns the ReplicationSetupTime field if non-nil, zero value otherwise.

### GetReplicationSetupTimeOk

`func (o *DrConfigDbDetail) GetReplicationSetupTimeOk() (*time.Time, bool)`

GetReplicationSetupTimeOk returns a tuple with the ReplicationSetupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSetupTime

`func (o *DrConfigDbDetail) SetReplicationSetupTime(v time.Time)`

SetReplicationSetupTime sets ReplicationSetupTime field to given value.

### HasReplicationSetupTime

`func (o *DrConfigDbDetail) HasReplicationSetupTime() bool`

HasReplicationSetupTime returns a boolean if a field has been set.

### GetBackupUuid

`func (o *DrConfigDbDetail) GetBackupUuid() string`

GetBackupUuid returns the BackupUuid field if non-nil, zero value otherwise.

### GetBackupUuidOk

`func (o *DrConfigDbDetail) GetBackupUuidOk() (*string, bool)`

GetBackupUuidOk returns a tuple with the BackupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupUuid

`func (o *DrConfigDbDetail) SetBackupUuid(v string)`

SetBackupUuid sets BackupUuid field to given value.

### HasBackupUuid

`func (o *DrConfigDbDetail) HasBackupUuid() bool`

HasBackupUuid returns a boolean if a field has been set.

### GetRestoreUuid

`func (o *DrConfigDbDetail) GetRestoreUuid() string`

GetRestoreUuid returns the RestoreUuid field if non-nil, zero value otherwise.

### GetRestoreUuidOk

`func (o *DrConfigDbDetail) GetRestoreUuidOk() (*string, bool)`

GetRestoreUuidOk returns a tuple with the RestoreUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreUuid

`func (o *DrConfigDbDetail) SetRestoreUuid(v string)`

SetRestoreUuid sets RestoreUuid field to given value.

### HasRestoreUuid

`func (o *DrConfigDbDetail) HasRestoreUuid() bool`

HasRestoreUuid returns a boolean if a field has been set.

### GetSourceNamespaceInfo

`func (o *DrConfigDbDetail) GetSourceNamespaceInfo() NamespaceInfo`

GetSourceNamespaceInfo returns the SourceNamespaceInfo field if non-nil, zero value otherwise.

### GetSourceNamespaceInfoOk

`func (o *DrConfigDbDetail) GetSourceNamespaceInfoOk() (*NamespaceInfo, bool)`

GetSourceNamespaceInfoOk returns a tuple with the SourceNamespaceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNamespaceInfo

`func (o *DrConfigDbDetail) SetSourceNamespaceInfo(v NamespaceInfo)`

SetSourceNamespaceInfo sets SourceNamespaceInfo field to given value.

### HasSourceNamespaceInfo

`func (o *DrConfigDbDetail) HasSourceNamespaceInfo() bool`

HasSourceNamespaceInfo returns a boolean if a field has been set.

### GetTargetNamespaceInfo

`func (o *DrConfigDbDetail) GetTargetNamespaceInfo() NamespaceInfo`

GetTargetNamespaceInfo returns the TargetNamespaceInfo field if non-nil, zero value otherwise.

### GetTargetNamespaceInfoOk

`func (o *DrConfigDbDetail) GetTargetNamespaceInfoOk() (*NamespaceInfo, bool)`

GetTargetNamespaceInfoOk returns a tuple with the TargetNamespaceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNamespaceInfo

`func (o *DrConfigDbDetail) SetTargetNamespaceInfo(v NamespaceInfo)`

SetTargetNamespaceInfo sets TargetNamespaceInfo field to given value.

### HasTargetNamespaceInfo

`func (o *DrConfigDbDetail) HasTargetNamespaceInfo() bool`

HasTargetNamespaceInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


