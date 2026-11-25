# OciRegionLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | Region of the universe whose backups go to this location | 
**Location** | **string** | Location of the region&#39;s backups, in the format of backup_location | 
**HostBase** | Pointer to **string** | Host of the S3-compatible endpoint for this region, with S3_COMPATIBLE credentials | [optional] 

## Methods

### NewOciRegionLocation

`func NewOciRegionLocation(region string, location string, ) *OciRegionLocation`

NewOciRegionLocation instantiates a new OciRegionLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOciRegionLocationWithDefaults

`func NewOciRegionLocationWithDefaults() *OciRegionLocation`

NewOciRegionLocationWithDefaults instantiates a new OciRegionLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *OciRegionLocation) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *OciRegionLocation) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *OciRegionLocation) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetLocation

`func (o *OciRegionLocation) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OciRegionLocation) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OciRegionLocation) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetHostBase

`func (o *OciRegionLocation) GetHostBase() string`

GetHostBase returns the HostBase field if non-nil, zero value otherwise.

### GetHostBaseOk

`func (o *OciRegionLocation) GetHostBaseOk() (*string, bool)`

GetHostBaseOk returns a tuple with the HostBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostBase

`func (o *OciRegionLocation) SetHostBase(v string)`

SetHostBase sets HostBase field to given value.

### HasHostBase

`func (o *OciRegionLocation) HasHostBase() bool`

HasHostBase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


