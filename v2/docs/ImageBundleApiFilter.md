# ImageBundleApiFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** | only include bundles whose details match this architecture | [optional] 

## Methods

### NewImageBundleApiFilter

`func NewImageBundleApiFilter() *ImageBundleApiFilter`

NewImageBundleApiFilter instantiates a new ImageBundleApiFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageBundleApiFilterWithDefaults

`func NewImageBundleApiFilterWithDefaults() *ImageBundleApiFilter`

NewImageBundleApiFilterWithDefaults instantiates a new ImageBundleApiFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *ImageBundleApiFilter) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *ImageBundleApiFilter) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *ImageBundleApiFilter) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *ImageBundleApiFilter) HasArch() bool`

HasArch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


