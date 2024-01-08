# GCPRegionCloudInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceTemplate** | Pointer to **string** | The instance template to be used for nodes created in this region. | [optional] 
**YbImage** | Pointer to **string** | Deprecated since YBA version 2.20.0. Use provider.imageBundle instead | [optional] 

## Methods

### NewGCPRegionCloudInfo

`func NewGCPRegionCloudInfo() *GCPRegionCloudInfo`

NewGCPRegionCloudInfo instantiates a new GCPRegionCloudInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPRegionCloudInfoWithDefaults

`func NewGCPRegionCloudInfoWithDefaults() *GCPRegionCloudInfo`

NewGCPRegionCloudInfoWithDefaults instantiates a new GCPRegionCloudInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceTemplate

`func (o *GCPRegionCloudInfo) GetInstanceTemplate() string`

GetInstanceTemplate returns the InstanceTemplate field if non-nil, zero value otherwise.

### GetInstanceTemplateOk

`func (o *GCPRegionCloudInfo) GetInstanceTemplateOk() (*string, bool)`

GetInstanceTemplateOk returns a tuple with the InstanceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTemplate

`func (o *GCPRegionCloudInfo) SetInstanceTemplate(v string)`

SetInstanceTemplate sets InstanceTemplate field to given value.

### HasInstanceTemplate

`func (o *GCPRegionCloudInfo) HasInstanceTemplate() bool`

HasInstanceTemplate returns a boolean if a field has been set.

### GetYbImage

`func (o *GCPRegionCloudInfo) GetYbImage() string`

GetYbImage returns the YbImage field if non-nil, zero value otherwise.

### GetYbImageOk

`func (o *GCPRegionCloudInfo) GetYbImageOk() (*string, bool)`

GetYbImageOk returns a tuple with the YbImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbImage

`func (o *GCPRegionCloudInfo) SetYbImage(v string)`

SetYbImage sets YbImage field to given value.

### HasYbImage

`func (o *GCPRegionCloudInfo) HasYbImage() bool`

HasYbImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


