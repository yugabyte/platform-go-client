# ReleaseFormData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gcs** | [**GCSLocation**](GCSLocation.md) |  | 
**Http** | [**HttpLocation**](HttpLocation.md) |  | 
**S3** | [**S3Location**](S3Location.md) |  | 
**Version** | **string** |  | 

## Methods

### NewReleaseFormData

`func NewReleaseFormData(gcs GCSLocation, http HttpLocation, s3 S3Location, version string, ) *ReleaseFormData`

NewReleaseFormData instantiates a new ReleaseFormData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseFormDataWithDefaults

`func NewReleaseFormDataWithDefaults() *ReleaseFormData`

NewReleaseFormDataWithDefaults instantiates a new ReleaseFormData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGcs

`func (o *ReleaseFormData) GetGcs() GCSLocation`

GetGcs returns the Gcs field if non-nil, zero value otherwise.

### GetGcsOk

`func (o *ReleaseFormData) GetGcsOk() (*GCSLocation, bool)`

GetGcsOk returns a tuple with the Gcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcs

`func (o *ReleaseFormData) SetGcs(v GCSLocation)`

SetGcs sets Gcs field to given value.


### GetHttp

`func (o *ReleaseFormData) GetHttp() HttpLocation`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *ReleaseFormData) GetHttpOk() (*HttpLocation, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *ReleaseFormData) SetHttp(v HttpLocation)`

SetHttp sets Http field to given value.


### GetS3

`func (o *ReleaseFormData) GetS3() S3Location`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *ReleaseFormData) GetS3Ok() (*S3Location, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *ReleaseFormData) SetS3(v S3Location)`

SetS3 sets S3 field to given value.


### GetVersion

`func (o *ReleaseFormData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ReleaseFormData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ReleaseFormData) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


