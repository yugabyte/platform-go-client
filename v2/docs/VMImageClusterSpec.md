# VMImageClusterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterUuid** | **string** | cluster uuid | 
**ProviderUuid** | **string** | provider uuid | 
**ImageBundleUuid** | **string** | image bundle uuid | 

## Methods

### NewVMImageClusterSpec

`func NewVMImageClusterSpec(clusterUuid string, providerUuid string, imageBundleUuid string, ) *VMImageClusterSpec`

NewVMImageClusterSpec instantiates a new VMImageClusterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMImageClusterSpecWithDefaults

`func NewVMImageClusterSpecWithDefaults() *VMImageClusterSpec`

NewVMImageClusterSpecWithDefaults instantiates a new VMImageClusterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterUuid

`func (o *VMImageClusterSpec) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *VMImageClusterSpec) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *VMImageClusterSpec) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.


### GetProviderUuid

`func (o *VMImageClusterSpec) GetProviderUuid() string`

GetProviderUuid returns the ProviderUuid field if non-nil, zero value otherwise.

### GetProviderUuidOk

`func (o *VMImageClusterSpec) GetProviderUuidOk() (*string, bool)`

GetProviderUuidOk returns a tuple with the ProviderUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUuid

`func (o *VMImageClusterSpec) SetProviderUuid(v string)`

SetProviderUuid sets ProviderUuid field to given value.


### GetImageBundleUuid

`func (o *VMImageClusterSpec) GetImageBundleUuid() string`

GetImageBundleUuid returns the ImageBundleUuid field if non-nil, zero value otherwise.

### GetImageBundleUuidOk

`func (o *VMImageClusterSpec) GetImageBundleUuidOk() (*string, bool)`

GetImageBundleUuidOk returns a tuple with the ImageBundleUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBundleUuid

`func (o *VMImageClusterSpec) SetImageBundleUuid(v string)`

SetImageBundleUuid sets ImageBundleUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


