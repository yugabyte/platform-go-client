# S3StorageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageType** | [**StorageType**](StorageType.md) |  | 
**BackupLocation** | **string** | Default location of backups, as s3://&lt;bucket&gt;[/&lt;path&gt;]. Can&#39;t be changed after creation. | 
**ImmutableStorage** | Pointer to **bool** | Whether the bucket keeps objects immutable (object lock). If true, deleting a backup removes only its YBA metadata, not its files. | [optional] [default to false]
**HostBase** | Pointer to **string** | Host of the S3 endpoint, for S3-compatible storage. Defaults to s3.amazonaws.com. | [optional] 
**SigningRegion** | Pointer to **string** | AWS region to sign requests for, when it can&#39;t be derived from the bucket, as with some S3-compatible stores | [optional] 
**PathStyleAccess** | Pointer to **bool** | Whether to address the bucket in the path (https://host/bucket) rather than in the host name (https://bucket.host), as some S3-compatible stores require | [optional] [default to false]
**UseChunkedEncoding** | Pointer to **bool** | Whether to use chunked transfer encoding for uploads | [optional] [default to true]
**GlobalBucketAccess** | Pointer to **bool** | Whether the bucket is reachable from every region. If false, YBA addresses it through its own region. | [optional] [default to true]
**Credentials** | [**S3Credentials**](S3Credentials.md) |  | 
**RegionLocations** | Pointer to [**[]S3RegionLocation**](S3RegionLocation.md) | Per-region backup locations, for universes that store each region&#39;s backup in that region | [optional] 

## Methods

### NewS3StorageSpec

`func NewS3StorageSpec(storageType StorageType, backupLocation string, credentials S3Credentials, ) *S3StorageSpec`

NewS3StorageSpec instantiates a new S3StorageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3StorageSpecWithDefaults

`func NewS3StorageSpecWithDefaults() *S3StorageSpec`

NewS3StorageSpecWithDefaults instantiates a new S3StorageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageType

`func (o *S3StorageSpec) GetStorageType() StorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *S3StorageSpec) GetStorageTypeOk() (*StorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *S3StorageSpec) SetStorageType(v StorageType)`

SetStorageType sets StorageType field to given value.


### GetBackupLocation

`func (o *S3StorageSpec) GetBackupLocation() string`

GetBackupLocation returns the BackupLocation field if non-nil, zero value otherwise.

### GetBackupLocationOk

`func (o *S3StorageSpec) GetBackupLocationOk() (*string, bool)`

GetBackupLocationOk returns a tuple with the BackupLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLocation

`func (o *S3StorageSpec) SetBackupLocation(v string)`

SetBackupLocation sets BackupLocation field to given value.


### GetImmutableStorage

`func (o *S3StorageSpec) GetImmutableStorage() bool`

GetImmutableStorage returns the ImmutableStorage field if non-nil, zero value otherwise.

### GetImmutableStorageOk

`func (o *S3StorageSpec) GetImmutableStorageOk() (*bool, bool)`

GetImmutableStorageOk returns a tuple with the ImmutableStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutableStorage

`func (o *S3StorageSpec) SetImmutableStorage(v bool)`

SetImmutableStorage sets ImmutableStorage field to given value.

### HasImmutableStorage

`func (o *S3StorageSpec) HasImmutableStorage() bool`

HasImmutableStorage returns a boolean if a field has been set.

### GetHostBase

`func (o *S3StorageSpec) GetHostBase() string`

GetHostBase returns the HostBase field if non-nil, zero value otherwise.

### GetHostBaseOk

`func (o *S3StorageSpec) GetHostBaseOk() (*string, bool)`

GetHostBaseOk returns a tuple with the HostBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostBase

`func (o *S3StorageSpec) SetHostBase(v string)`

SetHostBase sets HostBase field to given value.

### HasHostBase

`func (o *S3StorageSpec) HasHostBase() bool`

HasHostBase returns a boolean if a field has been set.

### GetSigningRegion

`func (o *S3StorageSpec) GetSigningRegion() string`

GetSigningRegion returns the SigningRegion field if non-nil, zero value otherwise.

### GetSigningRegionOk

`func (o *S3StorageSpec) GetSigningRegionOk() (*string, bool)`

GetSigningRegionOk returns a tuple with the SigningRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningRegion

`func (o *S3StorageSpec) SetSigningRegion(v string)`

SetSigningRegion sets SigningRegion field to given value.

### HasSigningRegion

`func (o *S3StorageSpec) HasSigningRegion() bool`

HasSigningRegion returns a boolean if a field has been set.

### GetPathStyleAccess

`func (o *S3StorageSpec) GetPathStyleAccess() bool`

GetPathStyleAccess returns the PathStyleAccess field if non-nil, zero value otherwise.

### GetPathStyleAccessOk

`func (o *S3StorageSpec) GetPathStyleAccessOk() (*bool, bool)`

GetPathStyleAccessOk returns a tuple with the PathStyleAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathStyleAccess

`func (o *S3StorageSpec) SetPathStyleAccess(v bool)`

SetPathStyleAccess sets PathStyleAccess field to given value.

### HasPathStyleAccess

`func (o *S3StorageSpec) HasPathStyleAccess() bool`

HasPathStyleAccess returns a boolean if a field has been set.

### GetUseChunkedEncoding

`func (o *S3StorageSpec) GetUseChunkedEncoding() bool`

GetUseChunkedEncoding returns the UseChunkedEncoding field if non-nil, zero value otherwise.

### GetUseChunkedEncodingOk

`func (o *S3StorageSpec) GetUseChunkedEncodingOk() (*bool, bool)`

GetUseChunkedEncodingOk returns a tuple with the UseChunkedEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseChunkedEncoding

`func (o *S3StorageSpec) SetUseChunkedEncoding(v bool)`

SetUseChunkedEncoding sets UseChunkedEncoding field to given value.

### HasUseChunkedEncoding

`func (o *S3StorageSpec) HasUseChunkedEncoding() bool`

HasUseChunkedEncoding returns a boolean if a field has been set.

### GetGlobalBucketAccess

`func (o *S3StorageSpec) GetGlobalBucketAccess() bool`

GetGlobalBucketAccess returns the GlobalBucketAccess field if non-nil, zero value otherwise.

### GetGlobalBucketAccessOk

`func (o *S3StorageSpec) GetGlobalBucketAccessOk() (*bool, bool)`

GetGlobalBucketAccessOk returns a tuple with the GlobalBucketAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalBucketAccess

`func (o *S3StorageSpec) SetGlobalBucketAccess(v bool)`

SetGlobalBucketAccess sets GlobalBucketAccess field to given value.

### HasGlobalBucketAccess

`func (o *S3StorageSpec) HasGlobalBucketAccess() bool`

HasGlobalBucketAccess returns a boolean if a field has been set.

### GetCredentials

`func (o *S3StorageSpec) GetCredentials() S3Credentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *S3StorageSpec) GetCredentialsOk() (*S3Credentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *S3StorageSpec) SetCredentials(v S3Credentials)`

SetCredentials sets Credentials field to given value.


### GetRegionLocations

`func (o *S3StorageSpec) GetRegionLocations() []S3RegionLocation`

GetRegionLocations returns the RegionLocations field if non-nil, zero value otherwise.

### GetRegionLocationsOk

`func (o *S3StorageSpec) GetRegionLocationsOk() (*[]S3RegionLocation, bool)`

GetRegionLocationsOk returns a tuple with the RegionLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionLocations

`func (o *S3StorageSpec) SetRegionLocations(v []S3RegionLocation)`

SetRegionLocations sets RegionLocations field to given value.

### HasRegionLocations

`func (o *S3StorageSpec) HasRegionLocations() bool`

HasRegionLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


