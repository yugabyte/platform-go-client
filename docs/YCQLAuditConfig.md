# YCQLAuditConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Enabled | 
**ExcludedCategories** | Pointer to **[]string** | Excluded Categories | [optional] 
**ExcludedKeyspaces** | Pointer to **[]string** | Excluded Keyspaces | [optional] 
**ExcludedUsers** | Pointer to **[]string** | Excluded Users | [optional] 
**IncludedCategories** | Pointer to **[]string** | Included categories | [optional] 
**IncludedKeyspaces** | Pointer to **[]string** | Included Keyspaces | [optional] 
**IncludedUsers** | Pointer to **[]string** | Included Users | [optional] 
**LogLevel** | Pointer to **string** | Log Level | [optional] 
**LogRetentionDays** | Pointer to **int32** | Number of days to keep gzipped YCQL audit log archives on the node. 0 or unset disables the dedicated audit-log retention pipeline and keeps the default size-based tserver log purge behavior. | [optional] 

## Methods

### NewYCQLAuditConfig

`func NewYCQLAuditConfig(enabled bool, ) *YCQLAuditConfig`

NewYCQLAuditConfig instantiates a new YCQLAuditConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYCQLAuditConfigWithDefaults

`func NewYCQLAuditConfigWithDefaults() *YCQLAuditConfig`

NewYCQLAuditConfigWithDefaults instantiates a new YCQLAuditConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *YCQLAuditConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *YCQLAuditConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *YCQLAuditConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExcludedCategories

`func (o *YCQLAuditConfig) GetExcludedCategories() []string`

GetExcludedCategories returns the ExcludedCategories field if non-nil, zero value otherwise.

### GetExcludedCategoriesOk

`func (o *YCQLAuditConfig) GetExcludedCategoriesOk() (*[]string, bool)`

GetExcludedCategoriesOk returns a tuple with the ExcludedCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedCategories

`func (o *YCQLAuditConfig) SetExcludedCategories(v []string)`

SetExcludedCategories sets ExcludedCategories field to given value.

### HasExcludedCategories

`func (o *YCQLAuditConfig) HasExcludedCategories() bool`

HasExcludedCategories returns a boolean if a field has been set.

### GetExcludedKeyspaces

`func (o *YCQLAuditConfig) GetExcludedKeyspaces() []string`

GetExcludedKeyspaces returns the ExcludedKeyspaces field if non-nil, zero value otherwise.

### GetExcludedKeyspacesOk

`func (o *YCQLAuditConfig) GetExcludedKeyspacesOk() (*[]string, bool)`

GetExcludedKeyspacesOk returns a tuple with the ExcludedKeyspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedKeyspaces

`func (o *YCQLAuditConfig) SetExcludedKeyspaces(v []string)`

SetExcludedKeyspaces sets ExcludedKeyspaces field to given value.

### HasExcludedKeyspaces

`func (o *YCQLAuditConfig) HasExcludedKeyspaces() bool`

HasExcludedKeyspaces returns a boolean if a field has been set.

### GetExcludedUsers

`func (o *YCQLAuditConfig) GetExcludedUsers() []string`

GetExcludedUsers returns the ExcludedUsers field if non-nil, zero value otherwise.

### GetExcludedUsersOk

`func (o *YCQLAuditConfig) GetExcludedUsersOk() (*[]string, bool)`

GetExcludedUsersOk returns a tuple with the ExcludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUsers

`func (o *YCQLAuditConfig) SetExcludedUsers(v []string)`

SetExcludedUsers sets ExcludedUsers field to given value.

### HasExcludedUsers

`func (o *YCQLAuditConfig) HasExcludedUsers() bool`

HasExcludedUsers returns a boolean if a field has been set.

### GetIncludedCategories

`func (o *YCQLAuditConfig) GetIncludedCategories() []string`

GetIncludedCategories returns the IncludedCategories field if non-nil, zero value otherwise.

### GetIncludedCategoriesOk

`func (o *YCQLAuditConfig) GetIncludedCategoriesOk() (*[]string, bool)`

GetIncludedCategoriesOk returns a tuple with the IncludedCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedCategories

`func (o *YCQLAuditConfig) SetIncludedCategories(v []string)`

SetIncludedCategories sets IncludedCategories field to given value.

### HasIncludedCategories

`func (o *YCQLAuditConfig) HasIncludedCategories() bool`

HasIncludedCategories returns a boolean if a field has been set.

### GetIncludedKeyspaces

`func (o *YCQLAuditConfig) GetIncludedKeyspaces() []string`

GetIncludedKeyspaces returns the IncludedKeyspaces field if non-nil, zero value otherwise.

### GetIncludedKeyspacesOk

`func (o *YCQLAuditConfig) GetIncludedKeyspacesOk() (*[]string, bool)`

GetIncludedKeyspacesOk returns a tuple with the IncludedKeyspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedKeyspaces

`func (o *YCQLAuditConfig) SetIncludedKeyspaces(v []string)`

SetIncludedKeyspaces sets IncludedKeyspaces field to given value.

### HasIncludedKeyspaces

`func (o *YCQLAuditConfig) HasIncludedKeyspaces() bool`

HasIncludedKeyspaces returns a boolean if a field has been set.

### GetIncludedUsers

`func (o *YCQLAuditConfig) GetIncludedUsers() []string`

GetIncludedUsers returns the IncludedUsers field if non-nil, zero value otherwise.

### GetIncludedUsersOk

`func (o *YCQLAuditConfig) GetIncludedUsersOk() (*[]string, bool)`

GetIncludedUsersOk returns a tuple with the IncludedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsers

`func (o *YCQLAuditConfig) SetIncludedUsers(v []string)`

SetIncludedUsers sets IncludedUsers field to given value.

### HasIncludedUsers

`func (o *YCQLAuditConfig) HasIncludedUsers() bool`

HasIncludedUsers returns a boolean if a field has been set.

### GetLogLevel

`func (o *YCQLAuditConfig) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *YCQLAuditConfig) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *YCQLAuditConfig) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *YCQLAuditConfig) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetLogRetentionDays

`func (o *YCQLAuditConfig) GetLogRetentionDays() int32`

GetLogRetentionDays returns the LogRetentionDays field if non-nil, zero value otherwise.

### GetLogRetentionDaysOk

`func (o *YCQLAuditConfig) GetLogRetentionDaysOk() (*int32, bool)`

GetLogRetentionDaysOk returns a tuple with the LogRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRetentionDays

`func (o *YCQLAuditConfig) SetLogRetentionDays(v int32)`

SetLogRetentionDays sets LogRetentionDays field to given value.

### HasLogRetentionDays

`func (o *YCQLAuditConfig) HasLogRetentionDays() bool`

HasLogRetentionDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


