# IsolatedBackupRestoreSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalPath** | **string** | full path to YBA backup file on host | 

## Methods

### NewIsolatedBackupRestoreSpec

`func NewIsolatedBackupRestoreSpec(localPath string, ) *IsolatedBackupRestoreSpec`

NewIsolatedBackupRestoreSpec instantiates a new IsolatedBackupRestoreSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsolatedBackupRestoreSpecWithDefaults

`func NewIsolatedBackupRestoreSpecWithDefaults() *IsolatedBackupRestoreSpec`

NewIsolatedBackupRestoreSpecWithDefaults instantiates a new IsolatedBackupRestoreSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalPath

`func (o *IsolatedBackupRestoreSpec) GetLocalPath() string`

GetLocalPath returns the LocalPath field if non-nil, zero value otherwise.

### GetLocalPathOk

`func (o *IsolatedBackupRestoreSpec) GetLocalPathOk() (*string, bool)`

GetLocalPathOk returns a tuple with the LocalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPath

`func (o *IsolatedBackupRestoreSpec) SetLocalPath(v string)`

SetLocalPath sets LocalPath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


