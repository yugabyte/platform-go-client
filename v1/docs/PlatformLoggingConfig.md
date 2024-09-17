# PlatformLoggingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileNamePrefix** | Pointer to **string** | WARNING: This is a preview API that could change. Application log file name prefix. Defaults to \&quot;\&quot;. For example, setting this to \&quot;yb-platform-\&quot; will generate application log files as \&quot;yb-platform-application.log\&quot; instead of \&quot;application.log\&quot;. | [optional] 
**Level** | **string** |  | 
**MaxHistory** | **int32** |  | 
**RolloverPattern** | **string** |  | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


