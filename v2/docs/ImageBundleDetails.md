# ImageBundleDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalYbImage** | Pointer to **string** | Default YugabyteDB machine image identifier for this bundle. | [optional] 
**Arch** | Pointer to **string** | CPU architecture this bundle applies to. | [optional] 
**Regions** | Pointer to [**map[string]ImageBundleRegionInfo**](ImageBundleRegionInfo.md) | Per-region YugabyteDB image identifiers keyed by region code. | [optional] 
**SshUser** | Pointer to **string** | SSH username for nodes provisioned from this bundle. | [optional] 
**SshPort** | Pointer to **int32** | SSH port for nodes provisioned from this bundle. | [optional] 
**UseImdsv2** | Pointer to **bool** | Whether to require AWS IMDSv2 on supported clouds. | [optional] 

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

### GetRegions

`func (o *ImageBundleDetails) GetRegions() map[string]ImageBundleRegionInfo`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *ImageBundleDetails) GetRegionsOk() (*map[string]ImageBundleRegionInfo, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *ImageBundleDetails) SetRegions(v map[string]ImageBundleRegionInfo)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *ImageBundleDetails) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetSshUser

`func (o *ImageBundleDetails) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *ImageBundleDetails) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *ImageBundleDetails) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.

### HasSshUser

`func (o *ImageBundleDetails) HasSshUser() bool`

HasSshUser returns a boolean if a field has been set.

### GetSshPort

`func (o *ImageBundleDetails) GetSshPort() int32`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *ImageBundleDetails) GetSshPortOk() (*int32, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *ImageBundleDetails) SetSshPort(v int32)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *ImageBundleDetails) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetUseImdsv2

`func (o *ImageBundleDetails) GetUseImdsv2() bool`

GetUseImdsv2 returns the UseImdsv2 field if non-nil, zero value otherwise.

### GetUseImdsv2Ok

`func (o *ImageBundleDetails) GetUseImdsv2Ok() (*bool, bool)`

GetUseImdsv2Ok returns a tuple with the UseImdsv2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseImdsv2

`func (o *ImageBundleDetails) SetUseImdsv2(v bool)`

SetUseImdsv2 sets UseImdsv2 field to given value.

### HasUseImdsv2

`func (o *ImageBundleDetails) HasUseImdsv2() bool`

HasUseImdsv2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


