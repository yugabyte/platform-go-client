# DeleteBackupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteBackupInfos** | [**[]DeleteBackupInfo**](DeleteBackupInfo.md) | Backups to be deleted | 
**DeleteForcefully** | Pointer to **bool** | Delete Backups forcefully | [optional] 

## Methods

### NewDeleteBackupParams

`func NewDeleteBackupParams(deleteBackupInfos []DeleteBackupInfo, ) *DeleteBackupParams`

NewDeleteBackupParams instantiates a new DeleteBackupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteBackupParamsWithDefaults

`func NewDeleteBackupParamsWithDefaults() *DeleteBackupParams`

NewDeleteBackupParamsWithDefaults instantiates a new DeleteBackupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteBackupInfos

`func (o *DeleteBackupParams) GetDeleteBackupInfos() []DeleteBackupInfo`

GetDeleteBackupInfos returns the DeleteBackupInfos field if non-nil, zero value otherwise.

### GetDeleteBackupInfosOk

`func (o *DeleteBackupParams) GetDeleteBackupInfosOk() (*[]DeleteBackupInfo, bool)`

GetDeleteBackupInfosOk returns a tuple with the DeleteBackupInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteBackupInfos

`func (o *DeleteBackupParams) SetDeleteBackupInfos(v []DeleteBackupInfo)`

SetDeleteBackupInfos sets DeleteBackupInfos field to given value.


### GetDeleteForcefully

`func (o *DeleteBackupParams) GetDeleteForcefully() bool`

GetDeleteForcefully returns the DeleteForcefully field if non-nil, zero value otherwise.

### GetDeleteForcefullyOk

`func (o *DeleteBackupParams) GetDeleteForcefullyOk() (*bool, bool)`

GetDeleteForcefullyOk returns a tuple with the DeleteForcefully field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteForcefully

`func (o *DeleteBackupParams) SetDeleteForcefully(v bool)`

SetDeleteForcefully sets DeleteForcefully field to given value.

### HasDeleteForcefully

`func (o *DeleteBackupParams) HasDeleteForcefully() bool`

HasDeleteForcefully returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


