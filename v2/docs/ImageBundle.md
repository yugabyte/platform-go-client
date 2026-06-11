# ImageBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | Pointer to [**ImageBundleSpec**](ImageBundleSpec.md) |  | [optional] 
**Info** | Pointer to [**ImageBundleInfo**](ImageBundleInfo.md) |  | [optional] 

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

### GetSpec

`func (o *ImageBundle) GetSpec() ImageBundleSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ImageBundle) GetSpecOk() (*ImageBundleSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ImageBundle) SetSpec(v ImageBundleSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ImageBundle) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetInfo

`func (o *ImageBundle) GetInfo() ImageBundleInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ImageBundle) GetInfoOk() (*ImageBundleInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ImageBundle) SetInfo(v ImageBundleInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ImageBundle) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


