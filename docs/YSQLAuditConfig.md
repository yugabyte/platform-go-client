# YSQLAuditConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classes** | **[]string** | YSQL statement classes | 
**Enabled** | **bool** | Enabled | [readonly] 
**LogCatalog** | **bool** | Log catalog | 
**LogClient** | **bool** | Log client | 
**LogLevel** | **string** | Log level. For NOTICE, INFO, DEBUG levels, user also needs to set &#39;log_min_messages&#39; to the required level for the audit logs to be exported. Default &#39;log_min_messages&#39; is WARNING. | 
**LogParameter** | **bool** | Log parameter | 
**LogParameterMaxSize** | **int32** | Log parameter max size | 
**LogRelation** | **bool** | Log relation | 
**LogRows** | **bool** | Log rows | 
**LogStatement** | **bool** | Log statement | 
**LogStatementOnce** | **bool** | Log statement once | 

## Methods

### NewYSQLAuditConfig

`func NewYSQLAuditConfig(classes []string, enabled bool, logCatalog bool, logClient bool, logLevel string, logParameter bool, logParameterMaxSize int32, logRelation bool, logRows bool, logStatement bool, logStatementOnce bool, ) *YSQLAuditConfig`

NewYSQLAuditConfig instantiates a new YSQLAuditConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYSQLAuditConfigWithDefaults

`func NewYSQLAuditConfigWithDefaults() *YSQLAuditConfig`

NewYSQLAuditConfigWithDefaults instantiates a new YSQLAuditConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClasses

`func (o *YSQLAuditConfig) GetClasses() []string`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *YSQLAuditConfig) GetClassesOk() (*[]string, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *YSQLAuditConfig) SetClasses(v []string)`

SetClasses sets Classes field to given value.


### GetEnabled

`func (o *YSQLAuditConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *YSQLAuditConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *YSQLAuditConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLogCatalog

`func (o *YSQLAuditConfig) GetLogCatalog() bool`

GetLogCatalog returns the LogCatalog field if non-nil, zero value otherwise.

### GetLogCatalogOk

`func (o *YSQLAuditConfig) GetLogCatalogOk() (*bool, bool)`

GetLogCatalogOk returns a tuple with the LogCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogCatalog

`func (o *YSQLAuditConfig) SetLogCatalog(v bool)`

SetLogCatalog sets LogCatalog field to given value.


### GetLogClient

`func (o *YSQLAuditConfig) GetLogClient() bool`

GetLogClient returns the LogClient field if non-nil, zero value otherwise.

### GetLogClientOk

`func (o *YSQLAuditConfig) GetLogClientOk() (*bool, bool)`

GetLogClientOk returns a tuple with the LogClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogClient

`func (o *YSQLAuditConfig) SetLogClient(v bool)`

SetLogClient sets LogClient field to given value.


### GetLogLevel

`func (o *YSQLAuditConfig) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *YSQLAuditConfig) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *YSQLAuditConfig) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.


### GetLogParameter

`func (o *YSQLAuditConfig) GetLogParameter() bool`

GetLogParameter returns the LogParameter field if non-nil, zero value otherwise.

### GetLogParameterOk

`func (o *YSQLAuditConfig) GetLogParameterOk() (*bool, bool)`

GetLogParameterOk returns a tuple with the LogParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogParameter

`func (o *YSQLAuditConfig) SetLogParameter(v bool)`

SetLogParameter sets LogParameter field to given value.


### GetLogParameterMaxSize

`func (o *YSQLAuditConfig) GetLogParameterMaxSize() int32`

GetLogParameterMaxSize returns the LogParameterMaxSize field if non-nil, zero value otherwise.

### GetLogParameterMaxSizeOk

`func (o *YSQLAuditConfig) GetLogParameterMaxSizeOk() (*int32, bool)`

GetLogParameterMaxSizeOk returns a tuple with the LogParameterMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogParameterMaxSize

`func (o *YSQLAuditConfig) SetLogParameterMaxSize(v int32)`

SetLogParameterMaxSize sets LogParameterMaxSize field to given value.


### GetLogRelation

`func (o *YSQLAuditConfig) GetLogRelation() bool`

GetLogRelation returns the LogRelation field if non-nil, zero value otherwise.

### GetLogRelationOk

`func (o *YSQLAuditConfig) GetLogRelationOk() (*bool, bool)`

GetLogRelationOk returns a tuple with the LogRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRelation

`func (o *YSQLAuditConfig) SetLogRelation(v bool)`

SetLogRelation sets LogRelation field to given value.


### GetLogRows

`func (o *YSQLAuditConfig) GetLogRows() bool`

GetLogRows returns the LogRows field if non-nil, zero value otherwise.

### GetLogRowsOk

`func (o *YSQLAuditConfig) GetLogRowsOk() (*bool, bool)`

GetLogRowsOk returns a tuple with the LogRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogRows

`func (o *YSQLAuditConfig) SetLogRows(v bool)`

SetLogRows sets LogRows field to given value.


### GetLogStatement

`func (o *YSQLAuditConfig) GetLogStatement() bool`

GetLogStatement returns the LogStatement field if non-nil, zero value otherwise.

### GetLogStatementOk

`func (o *YSQLAuditConfig) GetLogStatementOk() (*bool, bool)`

GetLogStatementOk returns a tuple with the LogStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStatement

`func (o *YSQLAuditConfig) SetLogStatement(v bool)`

SetLogStatement sets LogStatement field to given value.


### GetLogStatementOnce

`func (o *YSQLAuditConfig) GetLogStatementOnce() bool`

GetLogStatementOnce returns the LogStatementOnce field if non-nil, zero value otherwise.

### GetLogStatementOnceOk

`func (o *YSQLAuditConfig) GetLogStatementOnceOk() (*bool, bool)`

GetLogStatementOnceOk returns a tuple with the LogStatementOnce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStatementOnce

`func (o *YSQLAuditConfig) SetLogStatementOnce(v bool)`

SetLogStatementOnce sets LogStatementOnce field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


