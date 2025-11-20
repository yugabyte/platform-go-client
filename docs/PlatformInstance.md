# PlatformInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**ConfigUuid** | [**HighAvailabilityConfig**](HighAvailabilityConfig.md) |  | 
**InstanceState** | **string** |  | 
**IsLeader** | **bool** |  | 
**IsLocal** | **bool** |  | 
**LastBackup** | Pointer to **time.Time** | Last backup time | [optional] 
**Uuid** | **string** |  | 
**YbaVersion** | **string** |  | 

## Methods

### NewPlatformInstance

`func NewPlatformInstance(address string, configUuid HighAvailabilityConfig, instanceState string, isLeader bool, isLocal bool, uuid string, ybaVersion string, ) *PlatformInstance`

NewPlatformInstance instantiates a new PlatformInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformInstanceWithDefaults

`func NewPlatformInstanceWithDefaults() *PlatformInstance`

NewPlatformInstanceWithDefaults instantiates a new PlatformInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PlatformInstance) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PlatformInstance) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PlatformInstance) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetConfigUuid

`func (o *PlatformInstance) GetConfigUuid() HighAvailabilityConfig`

GetConfigUuid returns the ConfigUuid field if non-nil, zero value otherwise.

### GetConfigUuidOk

`func (o *PlatformInstance) GetConfigUuidOk() (*HighAvailabilityConfig, bool)`

GetConfigUuidOk returns a tuple with the ConfigUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigUuid

`func (o *PlatformInstance) SetConfigUuid(v HighAvailabilityConfig)`

SetConfigUuid sets ConfigUuid field to given value.


### GetInstanceState

`func (o *PlatformInstance) GetInstanceState() string`

GetInstanceState returns the InstanceState field if non-nil, zero value otherwise.

### GetInstanceStateOk

`func (o *PlatformInstance) GetInstanceStateOk() (*string, bool)`

GetInstanceStateOk returns a tuple with the InstanceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceState

`func (o *PlatformInstance) SetInstanceState(v string)`

SetInstanceState sets InstanceState field to given value.


### GetIsLeader

`func (o *PlatformInstance) GetIsLeader() bool`

GetIsLeader returns the IsLeader field if non-nil, zero value otherwise.

### GetIsLeaderOk

`func (o *PlatformInstance) GetIsLeaderOk() (*bool, bool)`

GetIsLeaderOk returns a tuple with the IsLeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLeader

`func (o *PlatformInstance) SetIsLeader(v bool)`

SetIsLeader sets IsLeader field to given value.


### GetIsLocal

`func (o *PlatformInstance) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *PlatformInstance) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *PlatformInstance) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.


### GetLastBackup

`func (o *PlatformInstance) GetLastBackup() time.Time`

GetLastBackup returns the LastBackup field if non-nil, zero value otherwise.

### GetLastBackupOk

`func (o *PlatformInstance) GetLastBackupOk() (*time.Time, bool)`

GetLastBackupOk returns a tuple with the LastBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackup

`func (o *PlatformInstance) SetLastBackup(v time.Time)`

SetLastBackup sets LastBackup field to given value.

### HasLastBackup

`func (o *PlatformInstance) HasLastBackup() bool`

HasLastBackup returns a boolean if a field has been set.

### GetUuid

`func (o *PlatformInstance) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PlatformInstance) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PlatformInstance) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetYbaVersion

`func (o *PlatformInstance) GetYbaVersion() string`

GetYbaVersion returns the YbaVersion field if non-nil, zero value otherwise.

### GetYbaVersionOk

`func (o *PlatformInstance) GetYbaVersionOk() (*string, bool)`

GetYbaVersionOk returns a tuple with the YbaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbaVersion

`func (o *PlatformInstance) SetYbaVersion(v string)`

SetYbaVersion sets YbaVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


