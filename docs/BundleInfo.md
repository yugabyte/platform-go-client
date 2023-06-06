# BundleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshPortOverride** | Pointer to **int32** |  | [optional] 
**SshUserOverride** | Pointer to **string** |  | [optional] 
**YbImage** | Pointer to **string** |  | [optional] 

## Methods

### NewBundleInfo

`func NewBundleInfo() *BundleInfo`

NewBundleInfo instantiates a new BundleInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleInfoWithDefaults

`func NewBundleInfoWithDefaults() *BundleInfo`

NewBundleInfoWithDefaults instantiates a new BundleInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshPortOverride

`func (o *BundleInfo) GetSshPortOverride() int32`

GetSshPortOverride returns the SshPortOverride field if non-nil, zero value otherwise.

### GetSshPortOverrideOk

`func (o *BundleInfo) GetSshPortOverrideOk() (*int32, bool)`

GetSshPortOverrideOk returns a tuple with the SshPortOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPortOverride

`func (o *BundleInfo) SetSshPortOverride(v int32)`

SetSshPortOverride sets SshPortOverride field to given value.

### HasSshPortOverride

`func (o *BundleInfo) HasSshPortOverride() bool`

HasSshPortOverride returns a boolean if a field has been set.

### GetSshUserOverride

`func (o *BundleInfo) GetSshUserOverride() string`

GetSshUserOverride returns the SshUserOverride field if non-nil, zero value otherwise.

### GetSshUserOverrideOk

`func (o *BundleInfo) GetSshUserOverrideOk() (*string, bool)`

GetSshUserOverrideOk returns a tuple with the SshUserOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUserOverride

`func (o *BundleInfo) SetSshUserOverride(v string)`

SetSshUserOverride sets SshUserOverride field to given value.

### HasSshUserOverride

`func (o *BundleInfo) HasSshUserOverride() bool`

HasSshUserOverride returns a boolean if a field has been set.

### GetYbImage

`func (o *BundleInfo) GetYbImage() string`

GetYbImage returns the YbImage field if non-nil, zero value otherwise.

### GetYbImageOk

`func (o *BundleInfo) GetYbImageOk() (*string, bool)`

GetYbImageOk returns a tuple with the YbImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbImage

`func (o *BundleInfo) SetYbImage(v string)`

SetYbImage sets YbImage field to given value.

### HasYbImage

`func (o *BundleInfo) HasYbImage() bool`

HasYbImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


