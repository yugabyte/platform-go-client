# S3RegionLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | Region of the universe whose backups go to this location | 
**Location** | **string** | Location of the region&#39;s backups, as s3://&lt;bucket&gt;[/&lt;path&gt;] | 
**HostBase** | Pointer to **string** | Host of the S3 endpoint for this region. Defaults to the configuration&#39;s host_base. | [optional] 
**SigningRegion** | Pointer to **string** | AWS region to sign this region&#39;s requests for | [optional] 
**GlobalBucketAccess** | Pointer to **bool** | Whether this region&#39;s bucket is reachable from every region | [optional] [default to true]

## Methods

### NewS3RegionLocation

`func NewS3RegionLocation(region string, location string, ) *S3RegionLocation`

NewS3RegionLocation instantiates a new S3RegionLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3RegionLocationWithDefaults

`func NewS3RegionLocationWithDefaults() *S3RegionLocation`

NewS3RegionLocationWithDefaults instantiates a new S3RegionLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *S3RegionLocation) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *S3RegionLocation) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *S3RegionLocation) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetLocation

`func (o *S3RegionLocation) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *S3RegionLocation) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *S3RegionLocation) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetHostBase

`func (o *S3RegionLocation) GetHostBase() string`

GetHostBase returns the HostBase field if non-nil, zero value otherwise.

### GetHostBaseOk

`func (o *S3RegionLocation) GetHostBaseOk() (*string, bool)`

GetHostBaseOk returns a tuple with the HostBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostBase

`func (o *S3RegionLocation) SetHostBase(v string)`

SetHostBase sets HostBase field to given value.

### HasHostBase

`func (o *S3RegionLocation) HasHostBase() bool`

HasHostBase returns a boolean if a field has been set.

### GetSigningRegion

`func (o *S3RegionLocation) GetSigningRegion() string`

GetSigningRegion returns the SigningRegion field if non-nil, zero value otherwise.

### GetSigningRegionOk

`func (o *S3RegionLocation) GetSigningRegionOk() (*string, bool)`

GetSigningRegionOk returns a tuple with the SigningRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningRegion

`func (o *S3RegionLocation) SetSigningRegion(v string)`

SetSigningRegion sets SigningRegion field to given value.

### HasSigningRegion

`func (o *S3RegionLocation) HasSigningRegion() bool`

HasSigningRegion returns a boolean if a field has been set.

### GetGlobalBucketAccess

`func (o *S3RegionLocation) GetGlobalBucketAccess() bool`

GetGlobalBucketAccess returns the GlobalBucketAccess field if non-nil, zero value otherwise.

### GetGlobalBucketAccessOk

`func (o *S3RegionLocation) GetGlobalBucketAccessOk() (*bool, bool)`

GetGlobalBucketAccessOk returns a tuple with the GlobalBucketAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalBucketAccess

`func (o *S3RegionLocation) SetGlobalBucketAccess(v bool)`

SetGlobalBucketAccess sets GlobalBucketAccess field to given value.

### HasGlobalBucketAccess

`func (o *S3RegionLocation) HasGlobalBucketAccess() bool`

HasGlobalBucketAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


