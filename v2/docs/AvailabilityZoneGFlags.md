# AvailabilityZoneGFlags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tserver** | Pointer to **map[string]string** | GFlags applied on TServer process | [optional] 
**Master** | Pointer to **map[string]string** | GFlags applied on Master process | [optional] 

## Methods

### NewAvailabilityZoneGFlags

`func NewAvailabilityZoneGFlags() *AvailabilityZoneGFlags`

NewAvailabilityZoneGFlags instantiates a new AvailabilityZoneGFlags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityZoneGFlagsWithDefaults

`func NewAvailabilityZoneGFlagsWithDefaults() *AvailabilityZoneGFlags`

NewAvailabilityZoneGFlagsWithDefaults instantiates a new AvailabilityZoneGFlags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTserver

`func (o *AvailabilityZoneGFlags) GetTserver() map[string]string`

GetTserver returns the Tserver field if non-nil, zero value otherwise.

### GetTserverOk

`func (o *AvailabilityZoneGFlags) GetTserverOk() (*map[string]string, bool)`

GetTserverOk returns a tuple with the Tserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserver

`func (o *AvailabilityZoneGFlags) SetTserver(v map[string]string)`

SetTserver sets Tserver field to given value.

### HasTserver

`func (o *AvailabilityZoneGFlags) HasTserver() bool`

HasTserver returns a boolean if a field has been set.

### GetMaster

`func (o *AvailabilityZoneGFlags) GetMaster() map[string]string`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *AvailabilityZoneGFlags) GetMasterOk() (*map[string]string, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *AvailabilityZoneGFlags) SetMaster(v map[string]string)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *AvailabilityZoneGFlags) HasMaster() bool`

HasMaster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


