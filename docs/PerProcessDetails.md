# PerProcessDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceInfo** | Pointer to [**DeviceInfo**](DeviceInfo.md) |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 

## Methods

### NewPerProcessDetails

`func NewPerProcessDetails() *PerProcessDetails`

NewPerProcessDetails instantiates a new PerProcessDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerProcessDetailsWithDefaults

`func NewPerProcessDetailsWithDefaults() *PerProcessDetails`

NewPerProcessDetailsWithDefaults instantiates a new PerProcessDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceInfo

`func (o *PerProcessDetails) GetDeviceInfo() DeviceInfo`

GetDeviceInfo returns the DeviceInfo field if non-nil, zero value otherwise.

### GetDeviceInfoOk

`func (o *PerProcessDetails) GetDeviceInfoOk() (*DeviceInfo, bool)`

GetDeviceInfoOk returns a tuple with the DeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfo

`func (o *PerProcessDetails) SetDeviceInfo(v DeviceInfo)`

SetDeviceInfo sets DeviceInfo field to given value.

### HasDeviceInfo

`func (o *PerProcessDetails) HasDeviceInfo() bool`

HasDeviceInfo returns a boolean if a field has been set.

### GetInstanceType

`func (o *PerProcessDetails) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *PerProcessDetails) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *PerProcessDetails) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *PerProcessDetails) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


