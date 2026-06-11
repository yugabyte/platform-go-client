# ImageBundleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Image bundle unique identifier. | [optional] [readonly] 
**Metadata** | Pointer to [**ImageBundleMetadata**](ImageBundleMetadata.md) |  | [optional] 
**Active** | Pointer to **bool** | Whether the bundle is active for use. | [optional] [readonly] 

## Methods

### NewImageBundleInfo

`func NewImageBundleInfo() *ImageBundleInfo`

NewImageBundleInfo instantiates a new ImageBundleInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageBundleInfoWithDefaults

`func NewImageBundleInfoWithDefaults() *ImageBundleInfo`

NewImageBundleInfoWithDefaults instantiates a new ImageBundleInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ImageBundleInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ImageBundleInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ImageBundleInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ImageBundleInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetMetadata

`func (o *ImageBundleInfo) GetMetadata() ImageBundleMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ImageBundleInfo) GetMetadataOk() (*ImageBundleMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ImageBundleInfo) SetMetadata(v ImageBundleMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ImageBundleInfo) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetActive

`func (o *ImageBundleInfo) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ImageBundleInfo) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ImageBundleInfo) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ImageBundleInfo) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


