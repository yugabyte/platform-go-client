# RestoreResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | **time.Time** |  | 
**CustomerUUID** | **string** |  | 
**IsSourceUniversePresent** | **bool** |  | 
**RestoreKeyspaceList** | [**[]RestoreKeyspace**](RestoreKeyspace.md) |  | 
**RestoreSizeInBytes** | **int64** |  | 
**RestoreUUID** | **string** |  | 
**SourceUniverseName** | **string** |  | 
**State** | **string** |  | 
**TargetUniverseName** | **string** |  | 
**UniverseUUID** | **string** |  | 
**UpdateTime** | **time.Time** |  | 

## Methods

### NewRestoreResp

`func NewRestoreResp(creationTime time.Time, customerUUID string, isSourceUniversePresent bool, restoreKeyspaceList []RestoreKeyspace, restoreSizeInBytes int64, restoreUUID string, sourceUniverseName string, state string, targetUniverseName string, universeUUID string, updateTime time.Time, ) *RestoreResp`

NewRestoreResp instantiates a new RestoreResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreRespWithDefaults

`func NewRestoreRespWithDefaults() *RestoreResp`

NewRestoreRespWithDefaults instantiates a new RestoreResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *RestoreResp) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RestoreResp) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RestoreResp) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetCustomerUUID

`func (o *RestoreResp) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *RestoreResp) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *RestoreResp) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetIsSourceUniversePresent

`func (o *RestoreResp) GetIsSourceUniversePresent() bool`

GetIsSourceUniversePresent returns the IsSourceUniversePresent field if non-nil, zero value otherwise.

### GetIsSourceUniversePresentOk

`func (o *RestoreResp) GetIsSourceUniversePresentOk() (*bool, bool)`

GetIsSourceUniversePresentOk returns a tuple with the IsSourceUniversePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSourceUniversePresent

`func (o *RestoreResp) SetIsSourceUniversePresent(v bool)`

SetIsSourceUniversePresent sets IsSourceUniversePresent field to given value.


### GetRestoreKeyspaceList

`func (o *RestoreResp) GetRestoreKeyspaceList() []RestoreKeyspace`

GetRestoreKeyspaceList returns the RestoreKeyspaceList field if non-nil, zero value otherwise.

### GetRestoreKeyspaceListOk

`func (o *RestoreResp) GetRestoreKeyspaceListOk() (*[]RestoreKeyspace, bool)`

GetRestoreKeyspaceListOk returns a tuple with the RestoreKeyspaceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreKeyspaceList

`func (o *RestoreResp) SetRestoreKeyspaceList(v []RestoreKeyspace)`

SetRestoreKeyspaceList sets RestoreKeyspaceList field to given value.


### GetRestoreSizeInBytes

`func (o *RestoreResp) GetRestoreSizeInBytes() int64`

GetRestoreSizeInBytes returns the RestoreSizeInBytes field if non-nil, zero value otherwise.

### GetRestoreSizeInBytesOk

`func (o *RestoreResp) GetRestoreSizeInBytesOk() (*int64, bool)`

GetRestoreSizeInBytesOk returns a tuple with the RestoreSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreSizeInBytes

`func (o *RestoreResp) SetRestoreSizeInBytes(v int64)`

SetRestoreSizeInBytes sets RestoreSizeInBytes field to given value.


### GetRestoreUUID

`func (o *RestoreResp) GetRestoreUUID() string`

GetRestoreUUID returns the RestoreUUID field if non-nil, zero value otherwise.

### GetRestoreUUIDOk

`func (o *RestoreResp) GetRestoreUUIDOk() (*string, bool)`

GetRestoreUUIDOk returns a tuple with the RestoreUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreUUID

`func (o *RestoreResp) SetRestoreUUID(v string)`

SetRestoreUUID sets RestoreUUID field to given value.


### GetSourceUniverseName

`func (o *RestoreResp) GetSourceUniverseName() string`

GetSourceUniverseName returns the SourceUniverseName field if non-nil, zero value otherwise.

### GetSourceUniverseNameOk

`func (o *RestoreResp) GetSourceUniverseNameOk() (*string, bool)`

GetSourceUniverseNameOk returns a tuple with the SourceUniverseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUniverseName

`func (o *RestoreResp) SetSourceUniverseName(v string)`

SetSourceUniverseName sets SourceUniverseName field to given value.


### GetState

`func (o *RestoreResp) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RestoreResp) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RestoreResp) SetState(v string)`

SetState sets State field to given value.


### GetTargetUniverseName

`func (o *RestoreResp) GetTargetUniverseName() string`

GetTargetUniverseName returns the TargetUniverseName field if non-nil, zero value otherwise.

### GetTargetUniverseNameOk

`func (o *RestoreResp) GetTargetUniverseNameOk() (*string, bool)`

GetTargetUniverseNameOk returns a tuple with the TargetUniverseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUniverseName

`func (o *RestoreResp) SetTargetUniverseName(v string)`

SetTargetUniverseName sets TargetUniverseName field to given value.


### GetUniverseUUID

`func (o *RestoreResp) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *RestoreResp) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *RestoreResp) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.


### GetUpdateTime

`func (o *RestoreResp) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *RestoreResp) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *RestoreResp) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


