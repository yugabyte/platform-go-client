# ReleaseMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartPath** | Pointer to **string** | Helm chart path | [optional] 
**FilePath** | Pointer to **string** | Release file path | [optional] 
**Gcs** | Pointer to [**GCSLocation**](GCSLocation.md) |  | [optional] 
**Http** | Pointer to [**HttpLocation**](HttpLocation.md) |  | [optional] 
**ImageTag** | Pointer to **string** | Release image tag | [optional] 
**Notes** | Pointer to **[]string** | Release notes | [optional] 
**Packages** | Pointer to [**[]Package**](Package.md) | Release packages | [optional] 
**S3** | Pointer to [**S3Location**](S3Location.md) |  | [optional] 
**State** | Pointer to **string** | Release state | [optional] 

## Methods

### NewReleaseMetadata

`func NewReleaseMetadata() *ReleaseMetadata`

NewReleaseMetadata instantiates a new ReleaseMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseMetadataWithDefaults

`func NewReleaseMetadataWithDefaults() *ReleaseMetadata`

NewReleaseMetadataWithDefaults instantiates a new ReleaseMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartPath

`func (o *ReleaseMetadata) GetChartPath() string`

GetChartPath returns the ChartPath field if non-nil, zero value otherwise.

### GetChartPathOk

`func (o *ReleaseMetadata) GetChartPathOk() (*string, bool)`

GetChartPathOk returns a tuple with the ChartPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartPath

`func (o *ReleaseMetadata) SetChartPath(v string)`

SetChartPath sets ChartPath field to given value.

### HasChartPath

`func (o *ReleaseMetadata) HasChartPath() bool`

HasChartPath returns a boolean if a field has been set.

### GetFilePath

`func (o *ReleaseMetadata) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *ReleaseMetadata) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *ReleaseMetadata) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *ReleaseMetadata) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetGcs

`func (o *ReleaseMetadata) GetGcs() GCSLocation`

GetGcs returns the Gcs field if non-nil, zero value otherwise.

### GetGcsOk

`func (o *ReleaseMetadata) GetGcsOk() (*GCSLocation, bool)`

GetGcsOk returns a tuple with the Gcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcs

`func (o *ReleaseMetadata) SetGcs(v GCSLocation)`

SetGcs sets Gcs field to given value.

### HasGcs

`func (o *ReleaseMetadata) HasGcs() bool`

HasGcs returns a boolean if a field has been set.

### GetHttp

`func (o *ReleaseMetadata) GetHttp() HttpLocation`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *ReleaseMetadata) GetHttpOk() (*HttpLocation, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *ReleaseMetadata) SetHttp(v HttpLocation)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *ReleaseMetadata) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetImageTag

`func (o *ReleaseMetadata) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *ReleaseMetadata) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *ReleaseMetadata) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.

### HasImageTag

`func (o *ReleaseMetadata) HasImageTag() bool`

HasImageTag returns a boolean if a field has been set.

### GetNotes

`func (o *ReleaseMetadata) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ReleaseMetadata) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ReleaseMetadata) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ReleaseMetadata) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPackages

`func (o *ReleaseMetadata) GetPackages() []Package`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *ReleaseMetadata) GetPackagesOk() (*[]Package, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *ReleaseMetadata) SetPackages(v []Package)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *ReleaseMetadata) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetS3

`func (o *ReleaseMetadata) GetS3() S3Location`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *ReleaseMetadata) GetS3Ok() (*S3Location, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *ReleaseMetadata) SetS3(v S3Location)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *ReleaseMetadata) HasS3() bool`

HasS3 returns a boolean if a field has been set.

### GetState

`func (o *ReleaseMetadata) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ReleaseMetadata) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ReleaseMetadata) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ReleaseMetadata) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


