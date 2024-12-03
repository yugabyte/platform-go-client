# ContinuousBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to [**ContinuousBackupInfo**](ContinuousBackupInfo.md) |  | [optional] 
**Spec** | Pointer to [**ContinuousBackupSpec**](ContinuousBackupSpec.md) |  | [optional] 

## Methods

### NewContinuousBackup

`func NewContinuousBackup() *ContinuousBackup`

NewContinuousBackup instantiates a new ContinuousBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContinuousBackupWithDefaults

`func NewContinuousBackupWithDefaults() *ContinuousBackup`

NewContinuousBackupWithDefaults instantiates a new ContinuousBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *ContinuousBackup) GetInfo() ContinuousBackupInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ContinuousBackup) GetInfoOk() (*ContinuousBackupInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ContinuousBackup) SetInfo(v ContinuousBackupInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ContinuousBackup) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetSpec

`func (o *ContinuousBackup) GetSpec() ContinuousBackupSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ContinuousBackup) GetSpecOk() (*ContinuousBackupSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ContinuousBackup) SetSpec(v ContinuousBackupSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ContinuousBackup) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

