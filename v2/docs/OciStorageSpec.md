# OciStorageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageType** | [**StorageType**](StorageType.md) |  | 
**BackupLocation** | **string** | Default location of backups: an Object Storage URL, https://objectstorage.&lt;region&gt;.oraclecloud.com/n/&lt;namespace&gt;/b/&lt;bucket&gt;/o[/&lt;path&gt;], with IAM credentials, or s3://&lt;bucket&gt;[/&lt;path&gt;] with S3_COMPATIBLE credentials. Can&#39;t be changed after creation. | 
**ImmutableStorage** | Pointer to **bool** | Whether the bucket keeps objects immutable (retention rules). If true, deleting a backup removes only its YBA metadata, not its files. | [optional] [default to false]
**Region** | **string** | OCI region of Object Storage. With S3_COMPATIBLE credentials, also the signing region. | 
**Namespace** | Pointer to **string** | Object Storage namespace. Required with IAM credentials; must match the namespace in an Object Storage URL. | [optional] 
**Credentials** | [**OciCredentials**](OciCredentials.md) |  | 
**RegionLocations** | Pointer to [**[]OciRegionLocation**](OciRegionLocation.md) | Per-region backup locations, for universes that store each region&#39;s backup in that region | [optional] 

## Methods

### NewOciStorageSpec

`func NewOciStorageSpec(storageType StorageType, backupLocation string, region string, credentials OciCredentials, ) *OciStorageSpec`

NewOciStorageSpec instantiates a new OciStorageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOciStorageSpecWithDefaults

`func NewOciStorageSpecWithDefaults() *OciStorageSpec`

NewOciStorageSpecWithDefaults instantiates a new OciStorageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageType

`func (o *OciStorageSpec) GetStorageType() StorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *OciStorageSpec) GetStorageTypeOk() (*StorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *OciStorageSpec) SetStorageType(v StorageType)`

SetStorageType sets StorageType field to given value.


### GetBackupLocation

`func (o *OciStorageSpec) GetBackupLocation() string`

GetBackupLocation returns the BackupLocation field if non-nil, zero value otherwise.

### GetBackupLocationOk

`func (o *OciStorageSpec) GetBackupLocationOk() (*string, bool)`

GetBackupLocationOk returns a tuple with the BackupLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLocation

`func (o *OciStorageSpec) SetBackupLocation(v string)`

SetBackupLocation sets BackupLocation field to given value.


### GetImmutableStorage

`func (o *OciStorageSpec) GetImmutableStorage() bool`

GetImmutableStorage returns the ImmutableStorage field if non-nil, zero value otherwise.

### GetImmutableStorageOk

`func (o *OciStorageSpec) GetImmutableStorageOk() (*bool, bool)`

GetImmutableStorageOk returns a tuple with the ImmutableStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutableStorage

`func (o *OciStorageSpec) SetImmutableStorage(v bool)`

SetImmutableStorage sets ImmutableStorage field to given value.

### HasImmutableStorage

`func (o *OciStorageSpec) HasImmutableStorage() bool`

HasImmutableStorage returns a boolean if a field has been set.

### GetRegion

`func (o *OciStorageSpec) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *OciStorageSpec) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *OciStorageSpec) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetNamespace

`func (o *OciStorageSpec) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *OciStorageSpec) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *OciStorageSpec) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *OciStorageSpec) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetCredentials

`func (o *OciStorageSpec) GetCredentials() OciCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *OciStorageSpec) GetCredentialsOk() (*OciCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *OciStorageSpec) SetCredentials(v OciCredentials)`

SetCredentials sets Credentials field to given value.


### GetRegionLocations

`func (o *OciStorageSpec) GetRegionLocations() []OciRegionLocation`

GetRegionLocations returns the RegionLocations field if non-nil, zero value otherwise.

### GetRegionLocationsOk

`func (o *OciStorageSpec) GetRegionLocationsOk() (*[]OciRegionLocation, bool)`

GetRegionLocationsOk returns a tuple with the RegionLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionLocations

`func (o *OciStorageSpec) SetRegionLocations(v []OciRegionLocation)`

SetRegionLocations sets RegionLocations field to given value.

### HasRegionLocations

`func (o *OciStorageSpec) HasRegionLocations() bool`

HasRegionLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


