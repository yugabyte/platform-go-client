# ImageBundleDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** | Architecture type for the image bundle | [optional] 
**GlobalYbImage** | Pointer to **string** | Global YB image for the bundle | [optional] 
**Regions** | Pointer to [**map[string]BundleInfo**](BundleInfo.md) | Regions override for image bundle | [optional] 

## Methods

### NewImageBundleDetails

`func NewImageBundleDetails() *ImageBundleDetails`

NewImageBundleDetails instantiates a new ImageBundleDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageBundleDetailsWithDefaults

`func NewImageBundleDetailsWithDefaults() *ImageBundleDetails`

NewImageBundleDetailsWithDefaults instantiates a new ImageBundleDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *ImageBundleDetails) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *ImageBundleDetails) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *ImageBundleDetails) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *ImageBundleDetails) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetGlobalYbImage

`func (o *ImageBundleDetails) GetGlobalYbImage() string`

GetGlobalYbImage returns the GlobalYbImage field if non-nil, zero value otherwise.

### GetGlobalYbImageOk

`func (o *ImageBundleDetails) GetGlobalYbImageOk() (*string, bool)`

GetGlobalYbImageOk returns a tuple with the GlobalYbImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalYbImage

`func (o *ImageBundleDetails) SetGlobalYbImage(v string)`

SetGlobalYbImage sets GlobalYbImage field to given value.

### HasGlobalYbImage

`func (o *ImageBundleDetails) HasGlobalYbImage() bool`

HasGlobalYbImage returns a boolean if a field has been set.

### GetRegions

`func (o *ImageBundleDetails) GetRegions() map[string]BundleInfo`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *ImageBundleDetails) GetRegionsOk() (*map[string]BundleInfo, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *ImageBundleDetails) SetRegions(v map[string]BundleInfo)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *ImageBundleDetails) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


