# ImageBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**ImageBundleDetails**](ImageBundleDetails.md) |  | [optional] 
**Name** | Pointer to **string** | Image Bundle Name | [optional] 
**UseAsDefault** | Pointer to **bool** | Default Image Bundle | [optional] 
**Uuid** | Pointer to **string** | Image Bundle UUID | [optional] [readonly] 

## Methods

### NewImageBundle

`func NewImageBundle() *ImageBundle`

NewImageBundle instantiates a new ImageBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageBundleWithDefaults

`func NewImageBundleWithDefaults() *ImageBundle`

NewImageBundleWithDefaults instantiates a new ImageBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *ImageBundle) GetDetails() ImageBundleDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ImageBundle) GetDetailsOk() (*ImageBundleDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ImageBundle) SetDetails(v ImageBundleDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ImageBundle) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetName

`func (o *ImageBundle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageBundle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageBundle) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageBundle) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUseAsDefault

`func (o *ImageBundle) GetUseAsDefault() bool`

GetUseAsDefault returns the UseAsDefault field if non-nil, zero value otherwise.

### GetUseAsDefaultOk

`func (o *ImageBundle) GetUseAsDefaultOk() (*bool, bool)`

GetUseAsDefaultOk returns a tuple with the UseAsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAsDefault

`func (o *ImageBundle) SetUseAsDefault(v bool)`

SetUseAsDefault sets UseAsDefault field to given value.

### HasUseAsDefault

`func (o *ImageBundle) HasUseAsDefault() bool`

HasUseAsDefault returns a boolean if a field has been set.

### GetUuid

`func (o *ImageBundle) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ImageBundle) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ImageBundle) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ImageBundle) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


