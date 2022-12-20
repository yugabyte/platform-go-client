# AuditLoggingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxHistory** | Pointer to **int32** | Max number of days up till which logs are kept | [optional] 
**OutputToFile** | **bool** | Flag to enable/disable audit logs output to file | 
**OutputToStdout** | **bool** | Flag to enable/disable audit logs output to stdout | 
**RolloverPattern** | Pointer to **string** | Rollover Pattern | [optional] 

## Methods

### NewAuditLoggingConfig

`func NewAuditLoggingConfig(outputToFile bool, outputToStdout bool, ) *AuditLoggingConfig`

NewAuditLoggingConfig instantiates a new AuditLoggingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLoggingConfigWithDefaults

`func NewAuditLoggingConfigWithDefaults() *AuditLoggingConfig`

NewAuditLoggingConfigWithDefaults instantiates a new AuditLoggingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxHistory

`func (o *AuditLoggingConfig) GetMaxHistory() int32`

GetMaxHistory returns the MaxHistory field if non-nil, zero value otherwise.

### GetMaxHistoryOk

`func (o *AuditLoggingConfig) GetMaxHistoryOk() (*int32, bool)`

GetMaxHistoryOk returns a tuple with the MaxHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHistory

`func (o *AuditLoggingConfig) SetMaxHistory(v int32)`

SetMaxHistory sets MaxHistory field to given value.

### HasMaxHistory

`func (o *AuditLoggingConfig) HasMaxHistory() bool`

HasMaxHistory returns a boolean if a field has been set.

### GetOutputToFile

`func (o *AuditLoggingConfig) GetOutputToFile() bool`

GetOutputToFile returns the OutputToFile field if non-nil, zero value otherwise.

### GetOutputToFileOk

`func (o *AuditLoggingConfig) GetOutputToFileOk() (*bool, bool)`

GetOutputToFileOk returns a tuple with the OutputToFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputToFile

`func (o *AuditLoggingConfig) SetOutputToFile(v bool)`

SetOutputToFile sets OutputToFile field to given value.


### GetOutputToStdout

`func (o *AuditLoggingConfig) GetOutputToStdout() bool`

GetOutputToStdout returns the OutputToStdout field if non-nil, zero value otherwise.

### GetOutputToStdoutOk

`func (o *AuditLoggingConfig) GetOutputToStdoutOk() (*bool, bool)`

GetOutputToStdoutOk returns a tuple with the OutputToStdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputToStdout

`func (o *AuditLoggingConfig) SetOutputToStdout(v bool)`

SetOutputToStdout sets OutputToStdout field to given value.


### GetRolloverPattern

`func (o *AuditLoggingConfig) GetRolloverPattern() string`

GetRolloverPattern returns the RolloverPattern field if non-nil, zero value otherwise.

### GetRolloverPatternOk

`func (o *AuditLoggingConfig) GetRolloverPatternOk() (*string, bool)`

GetRolloverPatternOk returns a tuple with the RolloverPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloverPattern

`func (o *AuditLoggingConfig) SetRolloverPattern(v string)`

SetRolloverPattern sets RolloverPattern field to given value.

### HasRolloverPattern

`func (o *AuditLoggingConfig) HasRolloverPattern() bool`

HasRolloverPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


