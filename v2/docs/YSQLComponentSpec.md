# YSQLComponentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentName** | Pointer to **string** | Logical component name; used as output file label. Must be a single file name of up to 128 characters from [A-Za-z0-9._-].  | [optional] 
**DbName** | Pointer to **string** | Database name for ysqlsh -d. | [optional] 
**Queries** | Pointer to **[]string** | YSQL statements to execute (each runs as ysqlsh -c). Support bundles are read-only, so each entry must be a single statement beginning with SELECT, WITH, EXPLAIN, SHOW, TABLE or VALUES (case-insensitive). Entries holding more than one statement are rejected, as are constructs that reach outside the database such as COPY ... TO PROGRAM or the pg_read_file family. The session additionally runs with default_transaction_read_only&#x3D;on.  | [optional] 
**OutputFileName** | Pointer to **string** | Output file name written under the per-node bundle directory. Must be a single file name of up to 128 characters from [A-Za-z0-9._-].  | [optional] 
**TimeoutSecs** | Pointer to **int64** | Command timeout in seconds. | [optional] 

## Methods

### NewYSQLComponentSpec

`func NewYSQLComponentSpec() *YSQLComponentSpec`

NewYSQLComponentSpec instantiates a new YSQLComponentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYSQLComponentSpecWithDefaults

`func NewYSQLComponentSpecWithDefaults() *YSQLComponentSpec`

NewYSQLComponentSpecWithDefaults instantiates a new YSQLComponentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentName

`func (o *YSQLComponentSpec) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *YSQLComponentSpec) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *YSQLComponentSpec) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *YSQLComponentSpec) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### GetDbName

`func (o *YSQLComponentSpec) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *YSQLComponentSpec) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *YSQLComponentSpec) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *YSQLComponentSpec) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetQueries

`func (o *YSQLComponentSpec) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *YSQLComponentSpec) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *YSQLComponentSpec) SetQueries(v []string)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *YSQLComponentSpec) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### GetOutputFileName

`func (o *YSQLComponentSpec) GetOutputFileName() string`

GetOutputFileName returns the OutputFileName field if non-nil, zero value otherwise.

### GetOutputFileNameOk

`func (o *YSQLComponentSpec) GetOutputFileNameOk() (*string, bool)`

GetOutputFileNameOk returns a tuple with the OutputFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFileName

`func (o *YSQLComponentSpec) SetOutputFileName(v string)`

SetOutputFileName sets OutputFileName field to given value.

### HasOutputFileName

`func (o *YSQLComponentSpec) HasOutputFileName() bool`

HasOutputFileName returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *YSQLComponentSpec) GetTimeoutSecs() int64`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *YSQLComponentSpec) GetTimeoutSecsOk() (*int64, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *YSQLComponentSpec) SetTimeoutSecs(v int64)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *YSQLComponentSpec) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


