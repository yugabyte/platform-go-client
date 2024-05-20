# VolumeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountPath** | **string** |  | 
**VolumeSizeGB** | **int32** |  | 
**VolumeType** | **string** |  | 

## Methods

### NewVolumeDetails

`func NewVolumeDetails(mountPath string, volumeSizeGB int32, volumeType string, ) *VolumeDetails`

NewVolumeDetails instantiates a new VolumeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeDetailsWithDefaults

`func NewVolumeDetailsWithDefaults() *VolumeDetails`

NewVolumeDetailsWithDefaults instantiates a new VolumeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountPath

`func (o *VolumeDetails) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *VolumeDetails) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *VolumeDetails) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.


### GetVolumeSizeGB

`func (o *VolumeDetails) GetVolumeSizeGB() int32`

GetVolumeSizeGB returns the VolumeSizeGB field if non-nil, zero value otherwise.

### GetVolumeSizeGBOk

`func (o *VolumeDetails) GetVolumeSizeGBOk() (*int32, bool)`

GetVolumeSizeGBOk returns a tuple with the VolumeSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSizeGB

`func (o *VolumeDetails) SetVolumeSizeGB(v int32)`

SetVolumeSizeGB sets VolumeSizeGB field to given value.


### GetVolumeType

`func (o *VolumeDetails) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *VolumeDetails) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *VolumeDetails) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


