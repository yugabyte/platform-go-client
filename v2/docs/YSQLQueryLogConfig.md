# YSQLQueryLogConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Enabled | [readonly] 
**LogStatement** | **string** | Log statement, controls which SQL statements are logged. | [default to "NONE"]
**LogMinErrorStatement** | **string** | Log min error statement, controls which SQL statements that cause an error condition are recorded. | [default to "ERROR"]
**LogErrorVerbosity** | **string** | Log error verbosity, controls the amount of detail written in the server log for each message that is logged. | [default to "DEFAULT"]
**LogDuration** | **bool** | Log duration, causes the duration of every completed statement to be logged | [default to false]
**DebugPrintPlan** | **bool** | Debug print plan, these parameters enable various debugging output to be emitted. When set, they print the resulting parse tree, the query rewriter output, or the execution plan for each executed query. | [default to false]
**LogConnections** | **bool** | Log connections, causes each attempted connection to the server to be logged, as well as successful completion of both client authentication (if necessary) and authorization | [default to false]
**LogDisconnections** | **bool** | Log disconnections, causes session terminations to be logged. The log output provides information similar to log_connections, plus the duration of the session. | [default to false]
**LogLinePrefix** | Pointer to **string** | Log line prefix. This would be applied to both audit and query logs. | [optional] 
**LogMinDurationStatement** | **int32** | Log min duration statement, causes the duration of each completed statement to be logged if the statement ran for at least the specified amount of time. | [default to -1]

## Methods

### NewYSQLQueryLogConfig

`func NewYSQLQueryLogConfig(enabled bool, logStatement string, logMinErrorStatement string, logErrorVerbosity string, logDuration bool, debugPrintPlan bool, logConnections bool, logDisconnections bool, logMinDurationStatement int32, ) *YSQLQueryLogConfig`

NewYSQLQueryLogConfig instantiates a new YSQLQueryLogConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYSQLQueryLogConfigWithDefaults

`func NewYSQLQueryLogConfigWithDefaults() *YSQLQueryLogConfig`

NewYSQLQueryLogConfigWithDefaults instantiates a new YSQLQueryLogConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *YSQLQueryLogConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *YSQLQueryLogConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *YSQLQueryLogConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLogStatement

`func (o *YSQLQueryLogConfig) GetLogStatement() string`

GetLogStatement returns the LogStatement field if non-nil, zero value otherwise.

### GetLogStatementOk

`func (o *YSQLQueryLogConfig) GetLogStatementOk() (*string, bool)`

GetLogStatementOk returns a tuple with the LogStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStatement

`func (o *YSQLQueryLogConfig) SetLogStatement(v string)`

SetLogStatement sets LogStatement field to given value.


### GetLogMinErrorStatement

`func (o *YSQLQueryLogConfig) GetLogMinErrorStatement() string`

GetLogMinErrorStatement returns the LogMinErrorStatement field if non-nil, zero value otherwise.

### GetLogMinErrorStatementOk

`func (o *YSQLQueryLogConfig) GetLogMinErrorStatementOk() (*string, bool)`

GetLogMinErrorStatementOk returns a tuple with the LogMinErrorStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogMinErrorStatement

`func (o *YSQLQueryLogConfig) SetLogMinErrorStatement(v string)`

SetLogMinErrorStatement sets LogMinErrorStatement field to given value.


### GetLogErrorVerbosity

`func (o *YSQLQueryLogConfig) GetLogErrorVerbosity() string`

GetLogErrorVerbosity returns the LogErrorVerbosity field if non-nil, zero value otherwise.

### GetLogErrorVerbosityOk

`func (o *YSQLQueryLogConfig) GetLogErrorVerbosityOk() (*string, bool)`

GetLogErrorVerbosityOk returns a tuple with the LogErrorVerbosity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogErrorVerbosity

`func (o *YSQLQueryLogConfig) SetLogErrorVerbosity(v string)`

SetLogErrorVerbosity sets LogErrorVerbosity field to given value.


### GetLogDuration

`func (o *YSQLQueryLogConfig) GetLogDuration() bool`

GetLogDuration returns the LogDuration field if non-nil, zero value otherwise.

### GetLogDurationOk

`func (o *YSQLQueryLogConfig) GetLogDurationOk() (*bool, bool)`

GetLogDurationOk returns a tuple with the LogDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDuration

`func (o *YSQLQueryLogConfig) SetLogDuration(v bool)`

SetLogDuration sets LogDuration field to given value.


### GetDebugPrintPlan

`func (o *YSQLQueryLogConfig) GetDebugPrintPlan() bool`

GetDebugPrintPlan returns the DebugPrintPlan field if non-nil, zero value otherwise.

### GetDebugPrintPlanOk

`func (o *YSQLQueryLogConfig) GetDebugPrintPlanOk() (*bool, bool)`

GetDebugPrintPlanOk returns a tuple with the DebugPrintPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugPrintPlan

`func (o *YSQLQueryLogConfig) SetDebugPrintPlan(v bool)`

SetDebugPrintPlan sets DebugPrintPlan field to given value.


### GetLogConnections

`func (o *YSQLQueryLogConfig) GetLogConnections() bool`

GetLogConnections returns the LogConnections field if non-nil, zero value otherwise.

### GetLogConnectionsOk

`func (o *YSQLQueryLogConfig) GetLogConnectionsOk() (*bool, bool)`

GetLogConnectionsOk returns a tuple with the LogConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogConnections

`func (o *YSQLQueryLogConfig) SetLogConnections(v bool)`

SetLogConnections sets LogConnections field to given value.


### GetLogDisconnections

`func (o *YSQLQueryLogConfig) GetLogDisconnections() bool`

GetLogDisconnections returns the LogDisconnections field if non-nil, zero value otherwise.

### GetLogDisconnectionsOk

`func (o *YSQLQueryLogConfig) GetLogDisconnectionsOk() (*bool, bool)`

GetLogDisconnectionsOk returns a tuple with the LogDisconnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDisconnections

`func (o *YSQLQueryLogConfig) SetLogDisconnections(v bool)`

SetLogDisconnections sets LogDisconnections field to given value.


### GetLogLinePrefix

`func (o *YSQLQueryLogConfig) GetLogLinePrefix() string`

GetLogLinePrefix returns the LogLinePrefix field if non-nil, zero value otherwise.

### GetLogLinePrefixOk

`func (o *YSQLQueryLogConfig) GetLogLinePrefixOk() (*string, bool)`

GetLogLinePrefixOk returns a tuple with the LogLinePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLinePrefix

`func (o *YSQLQueryLogConfig) SetLogLinePrefix(v string)`

SetLogLinePrefix sets LogLinePrefix field to given value.

### HasLogLinePrefix

`func (o *YSQLQueryLogConfig) HasLogLinePrefix() bool`

HasLogLinePrefix returns a boolean if a field has been set.

### GetLogMinDurationStatement

`func (o *YSQLQueryLogConfig) GetLogMinDurationStatement() int32`

GetLogMinDurationStatement returns the LogMinDurationStatement field if non-nil, zero value otherwise.

### GetLogMinDurationStatementOk

`func (o *YSQLQueryLogConfig) GetLogMinDurationStatementOk() (*int32, bool)`

GetLogMinDurationStatementOk returns a tuple with the LogMinDurationStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogMinDurationStatement

`func (o *YSQLQueryLogConfig) SetLogMinDurationStatement(v int32)`

SetLogMinDurationStatement sets LogMinDurationStatement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


