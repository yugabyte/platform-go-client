# BackupRegionLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | Cloud region name. | [optional] [readonly] 
**Location** | Pointer to **string** | Backup storage location path or URI. | [optional] [readonly] 
**HostBase** | Pointer to **string** | Host base path for the backup location. | [optional] [readonly] 

## Methods

### NewBackupRegionLocation

`func NewBackupRegionLocation() *BackupRegionLocation`

NewBackupRegionLocation instantiates a new BackupRegionLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRegionLocationWithDefaults

`func NewBackupRegionLocationWithDefaults() *BackupRegionLocation`

NewBackupRegionLocationWithDefaults instantiates a new BackupRegionLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *BackupRegionLocation) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *BackupRegionLocation) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *BackupRegionLocation) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *BackupRegionLocation) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetLocation

`func (o *BackupRegionLocation) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BackupRegionLocation) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BackupRegionLocation) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BackupRegionLocation) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetHostBase

`func (o *BackupRegionLocation) GetHostBase() string`

GetHostBase returns the HostBase field if non-nil, zero value otherwise.

### GetHostBaseOk

`func (o *BackupRegionLocation) GetHostBaseOk() (*string, bool)`

GetHostBaseOk returns a tuple with the HostBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostBase

`func (o *BackupRegionLocation) SetHostBase(v string)`

SetHostBase sets HostBase field to given value.

### HasHostBase

`func (o *BackupRegionLocation) HasHostBase() bool`

HasHostBase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


