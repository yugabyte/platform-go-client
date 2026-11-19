# PlatformLoggingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileNamePrefix** | Pointer to **string** | WARNING: This is a preview API that could change. Application log file name prefix. Defaults to \&quot;\&quot;. For example, setting this to \&quot;yb-platform-\&quot; will generate application log files as \&quot;yb-platform-application.log\&quot; instead of \&quot;application.log\&quot;. | [optional] 
**Level** | **string** |  | 
**MaxFileSize** | Pointer to **string** | WARNING: This is a preview API that could change. Maximum size of the active application log file before it is rolled over, for example \&quot;1024MB\&quot; or \&quot;1GB\&quot;. Minimum 256MB. | [optional] 
**MaxHistory** | **int32** |  | 
**RolloverPattern** | **string** |  | 
**TotalSizeCap** | Pointer to **string** | WARNING: This is a preview API that could change. Total size of retained application log archives, for example \&quot;20GB\&quot;. Archives are deleted oldest first once this is exceeded, so on a high volume node this binds before maxHistory and retains fewer days than maxHistory allows. Minimum 5GB. | [optional] 

## Methods

### NewPlatformLoggingConfig

`func NewPlatformLoggingConfig(level string, maxHistory int32, rolloverPattern string, ) *PlatformLoggingConfig`

NewPlatformLoggingConfig instantiates a new PlatformLoggingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformLoggingConfigWithDefaults

`func NewPlatformLoggingConfigWithDefaults() *PlatformLoggingConfig`

NewPlatformLoggingConfigWithDefaults instantiates a new PlatformLoggingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileNamePrefix

`func (o *PlatformLoggingConfig) GetFileNamePrefix() string`

GetFileNamePrefix returns the FileNamePrefix field if non-nil, zero value otherwise.

### GetFileNamePrefixOk

`func (o *PlatformLoggingConfig) GetFileNamePrefixOk() (*string, bool)`

GetFileNamePrefixOk returns a tuple with the FileNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileNamePrefix

`func (o *PlatformLoggingConfig) SetFileNamePrefix(v string)`

SetFileNamePrefix sets FileNamePrefix field to given value.

### HasFileNamePrefix

`func (o *PlatformLoggingConfig) HasFileNamePrefix() bool`

HasFileNamePrefix returns a boolean if a field has been set.

### GetLevel

`func (o *PlatformLoggingConfig) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *PlatformLoggingConfig) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *PlatformLoggingConfig) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetMaxFileSize

`func (o *PlatformLoggingConfig) GetMaxFileSize() string`

GetMaxFileSize returns the MaxFileSize field if non-nil, zero value otherwise.

### GetMaxFileSizeOk

`func (o *PlatformLoggingConfig) GetMaxFileSizeOk() (*string, bool)`

GetMaxFileSizeOk returns a tuple with the MaxFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFileSize

`func (o *PlatformLoggingConfig) SetMaxFileSize(v string)`

SetMaxFileSize sets MaxFileSize field to given value.

### HasMaxFileSize

`func (o *PlatformLoggingConfig) HasMaxFileSize() bool`

HasMaxFileSize returns a boolean if a field has been set.

### GetMaxHistory

`func (o *PlatformLoggingConfig) GetMaxHistory() int32`

GetMaxHistory returns the MaxHistory field if non-nil, zero value otherwise.

### GetMaxHistoryOk

`func (o *PlatformLoggingConfig) GetMaxHistoryOk() (*int32, bool)`

GetMaxHistoryOk returns a tuple with the MaxHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHistory

`func (o *PlatformLoggingConfig) SetMaxHistory(v int32)`

SetMaxHistory sets MaxHistory field to given value.


### GetRolloverPattern

`func (o *PlatformLoggingConfig) GetRolloverPattern() string`

GetRolloverPattern returns the RolloverPattern field if non-nil, zero value otherwise.

### GetRolloverPatternOk

`func (o *PlatformLoggingConfig) GetRolloverPatternOk() (*string, bool)`

GetRolloverPatternOk returns a tuple with the RolloverPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverPattern

`func (o *PlatformLoggingConfig) SetRolloverPattern(v string)`

SetRolloverPattern sets RolloverPattern field to given value.


### GetTotalSizeCap

`func (o *PlatformLoggingConfig) GetTotalSizeCap() string`

GetTotalSizeCap returns the TotalSizeCap field if non-nil, zero value otherwise.

### GetTotalSizeCapOk

`func (o *PlatformLoggingConfig) GetTotalSizeCapOk() (*string, bool)`

GetTotalSizeCapOk returns a tuple with the TotalSizeCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSizeCap

`func (o *PlatformLoggingConfig) SetTotalSizeCap(v string)`

SetTotalSizeCap sets TotalSizeCap field to given value.

### HasTotalSizeCap

`func (o *PlatformLoggingConfig) HasTotalSizeCap() bool`

HasTotalSizeCap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


