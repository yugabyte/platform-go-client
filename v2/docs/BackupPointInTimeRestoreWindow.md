# BackupPointInTimeRestoreWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimestampRetentionWindowStartMillis** | Pointer to **int64** | Start of the PITR retention window, in milliseconds since epoch. | [optional] [readonly] 
**TimestampRetentionWindowEndMillis** | Pointer to **int64** | End of the PITR retention window, in milliseconds since epoch. | [optional] [readonly] 

## Methods

### NewBackupPointInTimeRestoreWindow

`func NewBackupPointInTimeRestoreWindow() *BackupPointInTimeRestoreWindow`

NewBackupPointInTimeRestoreWindow instantiates a new BackupPointInTimeRestoreWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPointInTimeRestoreWindowWithDefaults

`func NewBackupPointInTimeRestoreWindowWithDefaults() *BackupPointInTimeRestoreWindow`

NewBackupPointInTimeRestoreWindowWithDefaults instantiates a new BackupPointInTimeRestoreWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestampRetentionWindowStartMillis

`func (o *BackupPointInTimeRestoreWindow) GetTimestampRetentionWindowStartMillis() int64`

GetTimestampRetentionWindowStartMillis returns the TimestampRetentionWindowStartMillis field if non-nil, zero value otherwise.

### GetTimestampRetentionWindowStartMillisOk

`func (o *BackupPointInTimeRestoreWindow) GetTimestampRetentionWindowStartMillisOk() (*int64, bool)`

GetTimestampRetentionWindowStartMillisOk returns a tuple with the TimestampRetentionWindowStartMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampRetentionWindowStartMillis

`func (o *BackupPointInTimeRestoreWindow) SetTimestampRetentionWindowStartMillis(v int64)`

SetTimestampRetentionWindowStartMillis sets TimestampRetentionWindowStartMillis field to given value.

### HasTimestampRetentionWindowStartMillis

`func (o *BackupPointInTimeRestoreWindow) HasTimestampRetentionWindowStartMillis() bool`

HasTimestampRetentionWindowStartMillis returns a boolean if a field has been set.

### GetTimestampRetentionWindowEndMillis

`func (o *BackupPointInTimeRestoreWindow) GetTimestampRetentionWindowEndMillis() int64`

GetTimestampRetentionWindowEndMillis returns the TimestampRetentionWindowEndMillis field if non-nil, zero value otherwise.

### GetTimestampRetentionWindowEndMillisOk

`func (o *BackupPointInTimeRestoreWindow) GetTimestampRetentionWindowEndMillisOk() (*int64, bool)`

GetTimestampRetentionWindowEndMillisOk returns a tuple with the TimestampRetentionWindowEndMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampRetentionWindowEndMillis

`func (o *BackupPointInTimeRestoreWindow) SetTimestampRetentionWindowEndMillis(v int64)`

SetTimestampRetentionWindowEndMillis sets TimestampRetentionWindowEndMillis field to given value.

### HasTimestampRetentionWindowEndMillis

`func (o *BackupPointInTimeRestoreWindow) HasTimestampRetentionWindowEndMillis() bool`

HasTimestampRetentionWindowEndMillis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


