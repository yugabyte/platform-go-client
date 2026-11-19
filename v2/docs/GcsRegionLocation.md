# GcsRegionLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | Region of the universe whose backups go to this location | 
**Location** | **string** | Location of the region&#39;s backups, as gs://&lt;bucket&gt;[/&lt;path&gt;] | 

## Methods

### NewGcsRegionLocation

`func NewGcsRegionLocation(region string, location string, ) *GcsRegionLocation`

NewGcsRegionLocation instantiates a new GcsRegionLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcsRegionLocationWithDefaults

`func NewGcsRegionLocationWithDefaults() *GcsRegionLocation`

NewGcsRegionLocationWithDefaults instantiates a new GcsRegionLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *GcsRegionLocation) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GcsRegionLocation) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GcsRegionLocation) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetLocation

`func (o *GcsRegionLocation) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GcsRegionLocation) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GcsRegionLocation) SetLocation(v string)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


