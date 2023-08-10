# RestorePreflightResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupCategory** | Pointer to **string** | Backup Category | [optional] 
**HasKMSHistory** | Pointer to **bool** | Whether backup was KMS encrypted | [optional] 
**PerLocationBackupInfoMap** | Pointer to [**map[string]PerLocationBackupInfo**](PerLocationBackupInfo.md) | Map of backup location and backup-info object | [optional] 

## Methods

### NewRestorePreflightResponse

`func NewRestorePreflightResponse() *RestorePreflightResponse`

NewRestorePreflightResponse instantiates a new RestorePreflightResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestorePreflightResponseWithDefaults

`func NewRestorePreflightResponseWithDefaults() *RestorePreflightResponse`

NewRestorePreflightResponseWithDefaults instantiates a new RestorePreflightResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupCategory

`func (o *RestorePreflightResponse) GetBackupCategory() string`

GetBackupCategory returns the BackupCategory field if non-nil, zero value otherwise.

### GetBackupCategoryOk

`func (o *RestorePreflightResponse) GetBackupCategoryOk() (*string, bool)`

GetBackupCategoryOk returns a tuple with the BackupCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCategory

`func (o *RestorePreflightResponse) SetBackupCategory(v string)`

SetBackupCategory sets BackupCategory field to given value.

### HasBackupCategory

`func (o *RestorePreflightResponse) HasBackupCategory() bool`

HasBackupCategory returns a boolean if a field has been set.

### GetHasKMSHistory

`func (o *RestorePreflightResponse) GetHasKMSHistory() bool`

GetHasKMSHistory returns the HasKMSHistory field if non-nil, zero value otherwise.

### GetHasKMSHistoryOk

`func (o *RestorePreflightResponse) GetHasKMSHistoryOk() (*bool, bool)`

GetHasKMSHistoryOk returns a tuple with the HasKMSHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasKMSHistory

`func (o *RestorePreflightResponse) SetHasKMSHistory(v bool)`

SetHasKMSHistory sets HasKMSHistory field to given value.

### HasHasKMSHistory

`func (o *RestorePreflightResponse) HasHasKMSHistory() bool`

HasHasKMSHistory returns a boolean if a field has been set.

### GetPerLocationBackupInfoMap

`func (o *RestorePreflightResponse) GetPerLocationBackupInfoMap() map[string]PerLocationBackupInfo`

GetPerLocationBackupInfoMap returns the PerLocationBackupInfoMap field if non-nil, zero value otherwise.

### GetPerLocationBackupInfoMapOk

`func (o *RestorePreflightResponse) GetPerLocationBackupInfoMapOk() (*map[string]PerLocationBackupInfo, bool)`

GetPerLocationBackupInfoMapOk returns a tuple with the PerLocationBackupInfoMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerLocationBackupInfoMap

`func (o *RestorePreflightResponse) SetPerLocationBackupInfoMap(v map[string]PerLocationBackupInfo)`

SetPerLocationBackupInfoMap sets PerLocationBackupInfoMap field to given value.

### HasPerLocationBackupInfoMap

`func (o *RestorePreflightResponse) HasPerLocationBackupInfoMap() bool`

HasPerLocationBackupInfoMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


