# InstanceTypeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** | Architecture for the instance. | [optional] 
**Tenancy** | Pointer to **string** | Tenancy for the instance. | [optional] 
**VolumeDetailsList** | Pointer to [**[]VolumeDetails**](VolumeDetails.md) | Volume Details for the instance. | [optional] 

## Methods

### NewInstanceTypeDetails

`func NewInstanceTypeDetails() *InstanceTypeDetails`

NewInstanceTypeDetails instantiates a new InstanceTypeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeDetailsWithDefaults

`func NewInstanceTypeDetailsWithDefaults() *InstanceTypeDetails`

NewInstanceTypeDetailsWithDefaults instantiates a new InstanceTypeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *InstanceTypeDetails) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *InstanceTypeDetails) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *InstanceTypeDetails) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *InstanceTypeDetails) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetTenancy

`func (o *InstanceTypeDetails) GetTenancy() string`

GetTenancy returns the Tenancy field if non-nil, zero value otherwise.

### GetTenancyOk

`func (o *InstanceTypeDetails) GetTenancyOk() (*string, bool)`

GetTenancyOk returns a tuple with the Tenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancy

`func (o *InstanceTypeDetails) SetTenancy(v string)`

SetTenancy sets Tenancy field to given value.

### HasTenancy

`func (o *InstanceTypeDetails) HasTenancy() bool`

HasTenancy returns a boolean if a field has been set.

### GetVolumeDetailsList

`func (o *InstanceTypeDetails) GetVolumeDetailsList() []VolumeDetails`

GetVolumeDetailsList returns the VolumeDetailsList field if non-nil, zero value otherwise.

### GetVolumeDetailsListOk

`func (o *InstanceTypeDetails) GetVolumeDetailsListOk() (*[]VolumeDetails, bool)`

GetVolumeDetailsListOk returns a tuple with the VolumeDetailsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDetailsList

`func (o *InstanceTypeDetails) SetVolumeDetailsList(v []VolumeDetails)`

SetVolumeDetailsList sets VolumeDetailsList field to given value.

### HasVolumeDetailsList

`func (o *InstanceTypeDetails) HasVolumeDetailsList() bool`

HasVolumeDetailsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


