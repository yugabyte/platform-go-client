# OCIRegionCloudInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceTemplate** | Pointer to **string** | The OCI Instance Configuration OCID to use for nodes created in this region. | [optional] 
**Vnet** | Pointer to **string** |  | [optional] 
**YbImage** | Pointer to **string** | &lt;b style&#x3D;\&quot;color:#ff0000\&quot;&gt;Deprecated since YBA version 2.20.0.&lt;/b&gt; Use provider.imageBundle instead | [optional] 

## Methods

### NewOCIRegionCloudInfo

`func NewOCIRegionCloudInfo() *OCIRegionCloudInfo`

NewOCIRegionCloudInfo instantiates a new OCIRegionCloudInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCIRegionCloudInfoWithDefaults

`func NewOCIRegionCloudInfoWithDefaults() *OCIRegionCloudInfo`

NewOCIRegionCloudInfoWithDefaults instantiates a new OCIRegionCloudInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceTemplate

`func (o *OCIRegionCloudInfo) GetInstanceTemplate() string`

GetInstanceTemplate returns the InstanceTemplate field if non-nil, zero value otherwise.

### GetInstanceTemplateOk

`func (o *OCIRegionCloudInfo) GetInstanceTemplateOk() (*string, bool)`

GetInstanceTemplateOk returns a tuple with the InstanceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTemplate

`func (o *OCIRegionCloudInfo) SetInstanceTemplate(v string)`

SetInstanceTemplate sets InstanceTemplate field to given value.

### HasInstanceTemplate

`func (o *OCIRegionCloudInfo) HasInstanceTemplate() bool`

HasInstanceTemplate returns a boolean if a field has been set.

### GetVnet

`func (o *OCIRegionCloudInfo) GetVnet() string`

GetVnet returns the Vnet field if non-nil, zero value otherwise.

### GetVnetOk

`func (o *OCIRegionCloudInfo) GetVnetOk() (*string, bool)`

GetVnetOk returns a tuple with the Vnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnet

`func (o *OCIRegionCloudInfo) SetVnet(v string)`

SetVnet sets Vnet field to given value.

### HasVnet

`func (o *OCIRegionCloudInfo) HasVnet() bool`

HasVnet returns a boolean if a field has been set.

### GetYbImage

`func (o *OCIRegionCloudInfo) GetYbImage() string`

GetYbImage returns the YbImage field if non-nil, zero value otherwise.

### GetYbImageOk

`func (o *OCIRegionCloudInfo) GetYbImageOk() (*string, bool)`

GetYbImageOk returns a tuple with the YbImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbImage

`func (o *OCIRegionCloudInfo) SetYbImage(v string)`

SetYbImage sets YbImage field to given value.

### HasYbImage

`func (o *OCIRegionCloudInfo) HasYbImage() bool`

HasYbImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


