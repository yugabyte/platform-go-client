# ClusterGFlags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tserver** | Pointer to **map[string]string** | GFlags applied on TServer process | [optional] 
**Master** | Pointer to **map[string]string** | GFlags applied on Master process | [optional] 
**AzGflags** | Pointer to [**map[string]AvailabilityZoneGFlags**](AvailabilityZoneGFlags.md) | GFlags per availability zone uuid | [optional] 
**GflagGroups** | Pointer to **[]string** | GFlag groups to be applied | [optional] 

## Methods

### NewClusterGFlags

`func NewClusterGFlags() *ClusterGFlags`

NewClusterGFlags instantiates a new ClusterGFlags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterGFlagsWithDefaults

`func NewClusterGFlagsWithDefaults() *ClusterGFlags`

NewClusterGFlagsWithDefaults instantiates a new ClusterGFlags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTserver

`func (o *ClusterGFlags) GetTserver() map[string]string`

GetTserver returns the Tserver field if non-nil, zero value otherwise.

### GetTserverOk

`func (o *ClusterGFlags) GetTserverOk() (*map[string]string, bool)`

GetTserverOk returns a tuple with the Tserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserver

`func (o *ClusterGFlags) SetTserver(v map[string]string)`

SetTserver sets Tserver field to given value.

### HasTserver

`func (o *ClusterGFlags) HasTserver() bool`

HasTserver returns a boolean if a field has been set.

### GetMaster

`func (o *ClusterGFlags) GetMaster() map[string]string`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *ClusterGFlags) GetMasterOk() (*map[string]string, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *ClusterGFlags) SetMaster(v map[string]string)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *ClusterGFlags) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetAzGflags

`func (o *ClusterGFlags) GetAzGflags() map[string]AvailabilityZoneGFlags`

GetAzGflags returns the AzGflags field if non-nil, zero value otherwise.

### GetAzGflagsOk

`func (o *ClusterGFlags) GetAzGflagsOk() (*map[string]AvailabilityZoneGFlags, bool)`

GetAzGflagsOk returns a tuple with the AzGflags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzGflags

`func (o *ClusterGFlags) SetAzGflags(v map[string]AvailabilityZoneGFlags)`

SetAzGflags sets AzGflags field to given value.

### HasAzGflags

`func (o *ClusterGFlags) HasAzGflags() bool`

HasAzGflags returns a boolean if a field has been set.

### GetGflagGroups

`func (o *ClusterGFlags) GetGflagGroups() []string`

GetGflagGroups returns the GflagGroups field if non-nil, zero value otherwise.

### GetGflagGroupsOk

`func (o *ClusterGFlags) GetGflagGroupsOk() (*[]string, bool)`

GetGflagGroupsOk returns a tuple with the GflagGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGflagGroups

`func (o *ClusterGFlags) SetGflagGroups(v []string)`

SetGflagGroups sets GflagGroups field to given value.

### HasGflagGroups

`func (o *ClusterGFlags) HasGflagGroups() bool`

HasGflagGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


