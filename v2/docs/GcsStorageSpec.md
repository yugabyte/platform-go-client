# GcsStorageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageType** | [**StorageType**](StorageType.md) |  | 
**BackupLocation** | **string** | Default location of backups, as gs://&lt;bucket&gt;[/&lt;path&gt;]. Can&#39;t be changed after creation. | 
**ImmutableStorage** | Pointer to **bool** | Whether the bucket keeps objects immutable (retention policy). If true, deleting a backup removes only its YBA metadata, not its files. | [optional] [default to false]
**Credentials** | [**GcsCredentials**](GcsCredentials.md) |  | 
**RegionLocations** | Pointer to [**[]GcsRegionLocation**](GcsRegionLocation.md) | Per-region backup locations, for universes that store each region&#39;s backup in that region | [optional] 

## Methods

### NewGcsStorageSpec

`func NewGcsStorageSpec(storageType StorageType, backupLocation string, credentials GcsCredentials, ) *GcsStorageSpec`

NewGcsStorageSpec instantiates a new GcsStorageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcsStorageSpecWithDefaults

`func NewGcsStorageSpecWithDefaults() *GcsStorageSpec`

NewGcsStorageSpecWithDefaults instantiates a new GcsStorageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageType

`func (o *GcsStorageSpec) GetStorageType() StorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *GcsStorageSpec) GetStorageTypeOk() (*StorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *GcsStorageSpec) SetStorageType(v StorageType)`

SetStorageType sets StorageType field to given value.


### GetBackupLocation

`func (o *GcsStorageSpec) GetBackupLocation() string`

GetBackupLocation returns the BackupLocation field if non-nil, zero value otherwise.

### GetBackupLocationOk

`func (o *GcsStorageSpec) GetBackupLocationOk() (*string, bool)`

GetBackupLocationOk returns a tuple with the BackupLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLocation

`func (o *GcsStorageSpec) SetBackupLocation(v string)`

SetBackupLocation sets BackupLocation field to given value.


### GetImmutableStorage

`func (o *GcsStorageSpec) GetImmutableStorage() bool`

GetImmutableStorage returns the ImmutableStorage field if non-nil, zero value otherwise.

### GetImmutableStorageOk

`func (o *GcsStorageSpec) GetImmutableStorageOk() (*bool, bool)`

GetImmutableStorageOk returns a tuple with the ImmutableStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutableStorage

`func (o *GcsStorageSpec) SetImmutableStorage(v bool)`

SetImmutableStorage sets ImmutableStorage field to given value.

### HasImmutableStorage

`func (o *GcsStorageSpec) HasImmutableStorage() bool`

HasImmutableStorage returns a boolean if a field has been set.

### GetCredentials

`func (o *GcsStorageSpec) GetCredentials() GcsCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GcsStorageSpec) GetCredentialsOk() (*GcsCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GcsStorageSpec) SetCredentials(v GcsCredentials)`

SetCredentials sets Credentials field to given value.


### GetRegionLocations

`func (o *GcsStorageSpec) GetRegionLocations() []GcsRegionLocation`

GetRegionLocations returns the RegionLocations field if non-nil, zero value otherwise.

### GetRegionLocationsOk

`func (o *GcsStorageSpec) GetRegionLocationsOk() (*[]GcsRegionLocation, bool)`

GetRegionLocationsOk returns a tuple with the RegionLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionLocations

`func (o *GcsStorageSpec) SetRegionLocations(v []GcsRegionLocation)`

SetRegionLocations sets RegionLocations field to given value.

### HasRegionLocations

`func (o *GcsStorageSpec) HasRegionLocations() bool`

HasRegionLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


