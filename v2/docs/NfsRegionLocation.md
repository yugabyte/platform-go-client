# NfsRegionLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | Region of the universe whose backups go to this location | 
**Location** | **string** | Absolute path of the region&#39;s NFS mount on the database nodes | 
**Bucket** | Pointer to **string** | Directory under location that holds the backups. Defaults to yugabyte_backup. | [optional] 
**Volumes** | Pointer to **[]string** | NFS volumes to check, in the backup prechecks, for being mounted on the region&#39;s nodes | [optional] 

## Methods

### NewNfsRegionLocation

`func NewNfsRegionLocation(region string, location string, ) *NfsRegionLocation`

NewNfsRegionLocation instantiates a new NfsRegionLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsRegionLocationWithDefaults

`func NewNfsRegionLocationWithDefaults() *NfsRegionLocation`

NewNfsRegionLocationWithDefaults instantiates a new NfsRegionLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *NfsRegionLocation) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NfsRegionLocation) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NfsRegionLocation) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetLocation

`func (o *NfsRegionLocation) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *NfsRegionLocation) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *NfsRegionLocation) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetBucket

`func (o *NfsRegionLocation) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *NfsRegionLocation) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *NfsRegionLocation) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *NfsRegionLocation) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetVolumes

`func (o *NfsRegionLocation) GetVolumes() []string`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *NfsRegionLocation) GetVolumesOk() (*[]string, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *NfsRegionLocation) SetVolumes(v []string)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *NfsRegionLocation) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


