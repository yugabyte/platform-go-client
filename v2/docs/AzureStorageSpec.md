# AzureStorageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageType** | [**StorageType**](StorageType.md) |  | 
**BackupLocation** | **string** | Default location of backups, as https://&lt;account&gt;.blob.core.windows.net/&lt;container&gt;[/&lt;path&gt;]. Can&#39;t be changed after creation. | 
**ImmutableStorage** | Pointer to **bool** | Whether the container keeps blobs immutable. If true, deleting a backup removes only its YBA metadata, not its files. | [optional] [default to false]
**Credentials** | [**AzureCredentials**](AzureCredentials.md) |  | 
**RegionLocations** | Pointer to [**[]AzureRegionLocation**](AzureRegionLocation.md) | Per-region backup locations, for universes that store each region&#39;s backup in that region | [optional] 

## Methods

### NewAzureStorageSpec

`func NewAzureStorageSpec(storageType StorageType, backupLocation string, credentials AzureCredentials, ) *AzureStorageSpec`

NewAzureStorageSpec instantiates a new AzureStorageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureStorageSpecWithDefaults

`func NewAzureStorageSpecWithDefaults() *AzureStorageSpec`

NewAzureStorageSpecWithDefaults instantiates a new AzureStorageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageType

`func (o *AzureStorageSpec) GetStorageType() StorageType`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *AzureStorageSpec) GetStorageTypeOk() (*StorageType, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *AzureStorageSpec) SetStorageType(v StorageType)`

SetStorageType sets StorageType field to given value.


### GetBackupLocation

`func (o *AzureStorageSpec) GetBackupLocation() string`

GetBackupLocation returns the BackupLocation field if non-nil, zero value otherwise.

### GetBackupLocationOk

`func (o *AzureStorageSpec) GetBackupLocationOk() (*string, bool)`

GetBackupLocationOk returns a tuple with the BackupLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLocation

`func (o *AzureStorageSpec) SetBackupLocation(v string)`

SetBackupLocation sets BackupLocation field to given value.


### GetImmutableStorage

`func (o *AzureStorageSpec) GetImmutableStorage() bool`

GetImmutableStorage returns the ImmutableStorage field if non-nil, zero value otherwise.

### GetImmutableStorageOk

`func (o *AzureStorageSpec) GetImmutableStorageOk() (*bool, bool)`

GetImmutableStorageOk returns a tuple with the ImmutableStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutableStorage

`func (o *AzureStorageSpec) SetImmutableStorage(v bool)`

SetImmutableStorage sets ImmutableStorage field to given value.

### HasImmutableStorage

`func (o *AzureStorageSpec) HasImmutableStorage() bool`

HasImmutableStorage returns a boolean if a field has been set.

### GetCredentials

`func (o *AzureStorageSpec) GetCredentials() AzureCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AzureStorageSpec) GetCredentialsOk() (*AzureCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AzureStorageSpec) SetCredentials(v AzureCredentials)`

SetCredentials sets Credentials field to given value.


### GetRegionLocations

`func (o *AzureStorageSpec) GetRegionLocations() []AzureRegionLocation`

GetRegionLocations returns the RegionLocations field if non-nil, zero value otherwise.

### GetRegionLocationsOk

`func (o *AzureStorageSpec) GetRegionLocationsOk() (*[]AzureRegionLocation, bool)`

GetRegionLocationsOk returns a tuple with the RegionLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionLocations

`func (o *AzureStorageSpec) SetRegionLocations(v []AzureRegionLocation)`

SetRegionLocations sets RegionLocations field to given value.

### HasRegionLocations

`func (o *AzureStorageSpec) HasRegionLocations() bool`

HasRegionLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


