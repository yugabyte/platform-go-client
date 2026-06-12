# ImageBundleSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Image bundle display name. | [optional] 
**Details** | Pointer to [**ImageBundleDetails**](ImageBundleDetails.md) |  | [optional] 
**UseAsDefault** | Pointer to **bool** | Whether this bundle is the default for new clusters on the provider. | [optional] 

## Methods

### NewImageBundleSpec

`func NewImageBundleSpec() *ImageBundleSpec`

NewImageBundleSpec instantiates a new ImageBundleSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageBundleSpecWithDefaults

`func NewImageBundleSpecWithDefaults() *ImageBundleSpec`

NewImageBundleSpecWithDefaults instantiates a new ImageBundleSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ImageBundleSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageBundleSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageBundleSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageBundleSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDetails

`func (o *ImageBundleSpec) GetDetails() ImageBundleDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ImageBundleSpec) GetDetailsOk() (*ImageBundleDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ImageBundleSpec) SetDetails(v ImageBundleDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ImageBundleSpec) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetUseAsDefault

`func (o *ImageBundleSpec) GetUseAsDefault() bool`

GetUseAsDefault returns the UseAsDefault field if non-nil, zero value otherwise.

### GetUseAsDefaultOk

`func (o *ImageBundleSpec) GetUseAsDefaultOk() (*bool, bool)`

GetUseAsDefaultOk returns a tuple with the UseAsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAsDefault

`func (o *ImageBundleSpec) SetUseAsDefault(v bool)`

SetUseAsDefault sets UseAsDefault field to given value.

### HasUseAsDefault

`func (o *ImageBundleSpec) HasUseAsDefault() bool`

HasUseAsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


