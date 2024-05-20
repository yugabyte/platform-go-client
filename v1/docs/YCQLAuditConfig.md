# YCQLAuditConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Enabled | 
**ExcludedCategories** | **[]string** | Excluded Categories | 
**ExcludedKeyspaces** | **[]string** | Excluded Keyspaces | 
**ExcludedUsers** | **[]string** | Excluded Users | 
**IncludedCategories** | **[]string** | Included categories | 
**IncludedKeyspaces** | **[]string** | Included Keyspaces | 
**IncludedUsers** | **[]string** | Included Users | 
**LogLevel** | **string** | Log Level | 

## Methods

### NewYCQLAuditConfig

`func NewYCQLAuditConfig(enabled bool, excludedCategories []string, excludedKeyspaces []string, excludedUsers []string, includedCategories []string, includedKeyspaces []string, includedUsers []string, logLevel string, ) *YCQLAuditConfig`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


