# StorageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageType** | [**StorageType**](StorageType.md) |  | 
**BackupLocation** | **string** | Default location of backups: an Object Storage URL, https://objectstorage.&lt;region&gt;.oraclecloud.com/n/&lt;namespace&gt;/b/&lt;bucket&gt;/o[/&lt;path&gt;], with IAM credentials, or s3://&lt;bucket&gt;[/&lt;path&gt;] with S3_COMPATIBLE credentials. Can&#39;t be changed after creation. | 
**ImmutableStorage** | Pointer to **bool** | Whether the bucket keeps objects immutable (retention rules). If true, deleting a backup removes only its YBA metadata, not its files. | [optional] [default to false]
**HostBase** | Pointer to **string** | Host of the S3 endpoint, for S3-compatible storage. Defaults to s3.amazonaws.com. | [optional] 
**SigningRegion** | Pointer to **string** | AWS region to sign requests for, when it can&#39;t be derived from the bucket, as with some S3-compatible stores | [optional] 
**PathStyleAccess** | Pointer to **bool** | Whether to address the bucket in the path (https://host/bucket) rather than in the host name (https://bucket.host), as some S3-compatible stores require | [optional] [default to false]
**UseChunkedEncoding** | Pointer to **bool** | Whether to use chunked transfer encoding for uploads | [optional] [default to true]
**GlobalBucketAccess** | Pointer to **bool** | Whether the bucket is reachable from every region. If false, YBA addresses it through its own region. | [optional] [default to true]
**Credentials** | [**OciCredentials**](OciCredentials.md) |  | 
**RegionLocations** | Pointer to [**[]OciRegionLocation**](OciRegionLocation.md) | Per-region backup locations, for universes that store each region&#39;s backup in that region | [optional] 
**Bucket** | Pointer to **string** | Directory under backup_location that holds the backups. Defaults to yugabyte_backup. | [optional] 
**Volumes** | Pointer to **[]string** | NFS volumes to check, in the backup prechecks, for being mounted on the nodes | [optional] 
**Region** | **string** | OCI region of Object Storage. With S3_COMPATIBLE credentials, also the signing region. | 
**Namespace** | Pointer to **string** | Object Storage namespace. Required with IAM credentials; must match the namespace in an Object Storage URL. | [optional] 

## Methods

### NewStorageSpec

`func NewStorageSpec(storageType StorageType, backupLocation string, credentials OciCredentials, region string, ) *StorageSpec`

NewStorageSpec instantiates a new StorageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSpecWithDefaults

`func NewStorageSpecWithDefaults() *StorageSpec`

NewStorageSpecWithDefaults instantiates a new StorageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageType

`func (o *StorageSpec) GetStorageType() StorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *StorageSpec) GetStorageTypeOk() (*StorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *StorageSpec) SetStorageType(v StorageType)`

SetStorageType sets StorageType field to given value.


### GetBackupLocation

`func (o *StorageSpec) GetBackupLocation() string`

GetBackupLocation returns the BackupLocation field if non-nil, zero value otherwise.

### GetBackupLocationOk

`func (o *StorageSpec) GetBackupLocationOk() (*string, bool)`

GetBackupLocationOk returns a tuple with the BackupLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLocation

`func (o *StorageSpec) SetBackupLocation(v string)`

SetBackupLocation sets BackupLocation field to given value.


### GetImmutableStorage

`func (o *StorageSpec) GetImmutableStorage() bool`

GetImmutableStorage returns the ImmutableStorage field if non-nil, zero value otherwise.

### GetImmutableStorageOk

`func (o *StorageSpec) GetImmutableStorageOk() (*bool, bool)`

GetImmutableStorageOk returns a tuple with the ImmutableStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutableStorage

`func (o *StorageSpec) SetImmutableStorage(v bool)`

SetImmutableStorage sets ImmutableStorage field to given value.

### HasImmutableStorage

`func (o *StorageSpec) HasImmutableStorage() bool`

HasImmutableStorage returns a boolean if a field has been set.

### GetHostBase

`func (o *StorageSpec) GetHostBase() string`

GetHostBase returns the HostBase field if non-nil, zero value otherwise.

### GetHostBaseOk

`func (o *StorageSpec) GetHostBaseOk() (*string, bool)`

GetHostBaseOk returns a tuple with the HostBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostBase

`func (o *StorageSpec) SetHostBase(v string)`

SetHostBase sets HostBase field to given value.

### HasHostBase

`func (o *StorageSpec) HasHostBase() bool`

HasHostBase returns a boolean if a field has been set.

### GetSigningRegion

`func (o *StorageSpec) GetSigningRegion() string`

GetSigningRegion returns the SigningRegion field if non-nil, zero value otherwise.

### GetSigningRegionOk

`func (o *StorageSpec) GetSigningRegionOk() (*string, bool)`

GetSigningRegionOk returns a tuple with the SigningRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningRegion

`func (o *StorageSpec) SetSigningRegion(v string)`

SetSigningRegion sets SigningRegion field to given value.

### HasSigningRegion

`func (o *StorageSpec) HasSigningRegion() bool`

HasSigningRegion returns a boolean if a field has been set.

### GetPathStyleAccess

`func (o *StorageSpec) GetPathStyleAccess() bool`

GetPathStyleAccess returns the PathStyleAccess field if non-nil, zero value otherwise.

### GetPathStyleAccessOk

`func (o *StorageSpec) GetPathStyleAccessOk() (*bool, bool)`

GetPathStyleAccessOk returns a tuple with the PathStyleAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathStyleAccess

`func (o *StorageSpec) SetPathStyleAccess(v bool)`

SetPathStyleAccess sets PathStyleAccess field to given value.

### HasPathStyleAccess

`func (o *StorageSpec) HasPathStyleAccess() bool`

HasPathStyleAccess returns a boolean if a field has been set.

### GetUseChunkedEncoding

`func (o *StorageSpec) GetUseChunkedEncoding() bool`

GetUseChunkedEncoding returns the UseChunkedEncoding field if non-nil, zero value otherwise.

### GetUseChunkedEncodingOk

`func (o *StorageSpec) GetUseChunkedEncodingOk() (*bool, bool)`

GetUseChunkedEncodingOk returns a tuple with the UseChunkedEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseChunkedEncoding

`func (o *StorageSpec) SetUseChunkedEncoding(v bool)`

SetUseChunkedEncoding sets UseChunkedEncoding field to given value.

### HasUseChunkedEncoding

`func (o *StorageSpec) HasUseChunkedEncoding() bool`

HasUseChunkedEncoding returns a boolean if a field has been set.

### GetGlobalBucketAccess

`func (o *StorageSpec) GetGlobalBucketAccess() bool`

GetGlobalBucketAccess returns the GlobalBucketAccess field if non-nil, zero value otherwise.

### GetGlobalBucketAccessOk

`func (o *StorageSpec) GetGlobalBucketAccessOk() (*bool, bool)`

GetGlobalBucketAccessOk returns a tuple with the GlobalBucketAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalBucketAccess

`func (o *StorageSpec) SetGlobalBucketAccess(v bool)`

SetGlobalBucketAccess sets GlobalBucketAccess field to given value.

### HasGlobalBucketAccess

`func (o *StorageSpec) HasGlobalBucketAccess() bool`

HasGlobalBucketAccess returns a boolean if a field has been set.

### GetCredentials

`func (o *StorageSpec) GetCredentials() OciCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *StorageSpec) GetCredentialsOk() (*OciCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *StorageSpec) SetCredentials(v OciCredentials)`

SetCredentials sets Credentials field to given value.


### GetRegionLocations

`func (o *StorageSpec) GetRegionLocations() []OciRegionLocation`

GetRegionLocations returns the RegionLocations field if non-nil, zero value otherwise.

### GetRegionLocationsOk

`func (o *StorageSpec) GetRegionLocationsOk() (*[]OciRegionLocation, bool)`

GetRegionLocationsOk returns a tuple with the RegionLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionLocations

`func (o *StorageSpec) SetRegionLocations(v []OciRegionLocation)`

SetRegionLocations sets RegionLocations field to given value.

### HasRegionLocations

`func (o *StorageSpec) HasRegionLocations() bool`

HasRegionLocations returns a boolean if a field has been set.

### GetBucket

`func (o *StorageSpec) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *StorageSpec) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *StorageSpec) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *StorageSpec) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetVolumes

`func (o *StorageSpec) GetVolumes() []string`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *StorageSpec) GetVolumesOk() (*[]string, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *StorageSpec) SetVolumes(v []string)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *StorageSpec) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetRegion

`func (o *StorageSpec) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StorageSpec) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StorageSpec) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetNamespace

`func (o *StorageSpec) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *StorageSpec) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *StorageSpec) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *StorageSpec) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


