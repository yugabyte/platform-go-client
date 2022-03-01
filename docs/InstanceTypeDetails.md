# InstanceTypeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tenancy** | **string** |  | 
**VolumeDetailsList** | [**[]VolumeDetails**](VolumeDetails.md) |  | 

## Methods

### NewInstanceTypeDetails

`func NewInstanceTypeDetails(tenancy string, volumeDetailsList []VolumeDetails, ) *InstanceTypeDetails`

NewInstanceTypeDetails instantiates a new InstanceTypeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeDetailsWithDefaults

`func NewInstanceTypeDetailsWithDefaults() *InstanceTypeDetails`

NewInstanceTypeDetailsWithDefaults instantiates a new InstanceTypeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


