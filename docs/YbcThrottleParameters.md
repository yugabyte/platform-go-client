# YbcThrottleParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskReadBytesPerSecond** | Pointer to **int64** | WARNING: This is a preview API that could change. Disk read bytes per second to throttle disk usage during backups | [optional] 
**DiskWriteBytesPerSecond** | Pointer to **int64** | WARNING: This is a preview API that could change. Disk write bytes per second to throttle disk usage during restore | [optional] 
**MaxConcurrentDownloads** | Pointer to **int64** | Max concurrent downloads per node | [optional] 
**MaxConcurrentUploads** | Pointer to **int64** | Max concurrent uploads per node | [optional] 
**PerDownloadNumObjects** | Pointer to **int64** | Max objects per download per node | [optional] 
**PerUploadNumObjects** | Pointer to **int64** | Max objects per upload per node | [optional] 
**ResetDefaults** | Pointer to **bool** | Unset Throttle parameters in YB-Controller | [optional] 

## Methods

### NewYbcThrottleParameters

`func NewYbcThrottleParameters() *YbcThrottleParameters`

NewYbcThrottleParameters instantiates a new YbcThrottleParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYbcThrottleParametersWithDefaults

`func NewYbcThrottleParametersWithDefaults() *YbcThrottleParameters`

NewYbcThrottleParametersWithDefaults instantiates a new YbcThrottleParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskReadBytesPerSecond

`func (o *YbcThrottleParameters) GetDiskReadBytesPerSecond() int64`

GetDiskReadBytesPerSecond returns the DiskReadBytesPerSecond field if non-nil, zero value otherwise.

### GetDiskReadBytesPerSecondOk

`func (o *YbcThrottleParameters) GetDiskReadBytesPerSecondOk() (*int64, bool)`

GetDiskReadBytesPerSecondOk returns a tuple with the DiskReadBytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskReadBytesPerSecond

`func (o *YbcThrottleParameters) SetDiskReadBytesPerSecond(v int64)`

SetDiskReadBytesPerSecond sets DiskReadBytesPerSecond field to given value.

### HasDiskReadBytesPerSecond

`func (o *YbcThrottleParameters) HasDiskReadBytesPerSecond() bool`

HasDiskReadBytesPerSecond returns a boolean if a field has been set.

### GetDiskWriteBytesPerSecond

`func (o *YbcThrottleParameters) GetDiskWriteBytesPerSecond() int64`

GetDiskWriteBytesPerSecond returns the DiskWriteBytesPerSecond field if non-nil, zero value otherwise.

### GetDiskWriteBytesPerSecondOk

`func (o *YbcThrottleParameters) GetDiskWriteBytesPerSecondOk() (*int64, bool)`

GetDiskWriteBytesPerSecondOk returns a tuple with the DiskWriteBytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskWriteBytesPerSecond

`func (o *YbcThrottleParameters) SetDiskWriteBytesPerSecond(v int64)`

SetDiskWriteBytesPerSecond sets DiskWriteBytesPerSecond field to given value.

### HasDiskWriteBytesPerSecond

`func (o *YbcThrottleParameters) HasDiskWriteBytesPerSecond() bool`

HasDiskWriteBytesPerSecond returns a boolean if a field has been set.

### GetMaxConcurrentDownloads

`func (o *YbcThrottleParameters) GetMaxConcurrentDownloads() int64`

GetMaxConcurrentDownloads returns the MaxConcurrentDownloads field if non-nil, zero value otherwise.

### GetMaxConcurrentDownloadsOk

`func (o *YbcThrottleParameters) GetMaxConcurrentDownloadsOk() (*int64, bool)`

GetMaxConcurrentDownloadsOk returns a tuple with the MaxConcurrentDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentDownloads

`func (o *YbcThrottleParameters) SetMaxConcurrentDownloads(v int64)`

SetMaxConcurrentDownloads sets MaxConcurrentDownloads field to given value.

### HasMaxConcurrentDownloads

`func (o *YbcThrottleParameters) HasMaxConcurrentDownloads() bool`

HasMaxConcurrentDownloads returns a boolean if a field has been set.

### GetMaxConcurrentUploads

`func (o *YbcThrottleParameters) GetMaxConcurrentUploads() int64`

GetMaxConcurrentUploads returns the MaxConcurrentUploads field if non-nil, zero value otherwise.

### GetMaxConcurrentUploadsOk

`func (o *YbcThrottleParameters) GetMaxConcurrentUploadsOk() (*int64, bool)`

GetMaxConcurrentUploadsOk returns a tuple with the MaxConcurrentUploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentUploads

`func (o *YbcThrottleParameters) SetMaxConcurrentUploads(v int64)`

SetMaxConcurrentUploads sets MaxConcurrentUploads field to given value.

### HasMaxConcurrentUploads

`func (o *YbcThrottleParameters) HasMaxConcurrentUploads() bool`

HasMaxConcurrentUploads returns a boolean if a field has been set.

### GetPerDownloadNumObjects

`func (o *YbcThrottleParameters) GetPerDownloadNumObjects() int64`

GetPerDownloadNumObjects returns the PerDownloadNumObjects field if non-nil, zero value otherwise.

### GetPerDownloadNumObjectsOk

`func (o *YbcThrottleParameters) GetPerDownloadNumObjectsOk() (*int64, bool)`

GetPerDownloadNumObjectsOk returns a tuple with the PerDownloadNumObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerDownloadNumObjects

`func (o *YbcThrottleParameters) SetPerDownloadNumObjects(v int64)`

SetPerDownloadNumObjects sets PerDownloadNumObjects field to given value.

### HasPerDownloadNumObjects

`func (o *YbcThrottleParameters) HasPerDownloadNumObjects() bool`

HasPerDownloadNumObjects returns a boolean if a field has been set.

### GetPerUploadNumObjects

`func (o *YbcThrottleParameters) GetPerUploadNumObjects() int64`

GetPerUploadNumObjects returns the PerUploadNumObjects field if non-nil, zero value otherwise.

### GetPerUploadNumObjectsOk

`func (o *YbcThrottleParameters) GetPerUploadNumObjectsOk() (*int64, bool)`

GetPerUploadNumObjectsOk returns a tuple with the PerUploadNumObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerUploadNumObjects

`func (o *YbcThrottleParameters) SetPerUploadNumObjects(v int64)`

SetPerUploadNumObjects sets PerUploadNumObjects field to given value.

### HasPerUploadNumObjects

`func (o *YbcThrottleParameters) HasPerUploadNumObjects() bool`

HasPerUploadNumObjects returns a boolean if a field has been set.

### GetResetDefaults

`func (o *YbcThrottleParameters) GetResetDefaults() bool`

GetResetDefaults returns the ResetDefaults field if non-nil, zero value otherwise.

### GetResetDefaultsOk

`func (o *YbcThrottleParameters) GetResetDefaultsOk() (*bool, bool)`

GetResetDefaultsOk returns a tuple with the ResetDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetDefaults

`func (o *YbcThrottleParameters) SetResetDefaults(v bool)`

SetResetDefaults sets ResetDefaults field to given value.

### HasResetDefaults

`func (o *YbcThrottleParameters) HasResetDefaults() bool`

HasResetDefaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


