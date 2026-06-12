# Backup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | Pointer to [**BackupSpec**](BackupSpec.md) |  | [optional] 
**Info** | Pointer to [**BackupInfo**](BackupInfo.md) |  | [optional] 

## Methods

### NewBackup

`func NewBackup() *Backup`

NewBackup instantiates a new Backup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupWithDefaults

`func NewBackupWithDefaults() *Backup`

NewBackupWithDefaults instantiates a new Backup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *Backup) GetSpec() BackupSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *Backup) GetSpecOk() (*BackupSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *Backup) SetSpec(v BackupSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *Backup) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetInfo

`func (o *Backup) GetInfo() BackupInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Backup) GetInfoOk() (*BackupInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Backup) SetInfo(v BackupInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Backup) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


