# YCQLComponentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentName** | Pointer to **string** | Logical component name; used as output file label. Must be a single file name of up to 128 characters from [A-Za-z0-9._-].  | [optional] 
**Keyspace** | Pointer to **string** | Optional keyspace for ycqlsh -k. | [optional] 
**Queries** | Pointer to **[]string** | YCQL statements to execute (each runs as ycqlsh -e). Support bundles are read-only, so each entry must be a single statement beginning with SELECT, DESCRIBE, DESC, SHOW or LIST (case-insensitive). Entries holding more than one statement are rejected.  | [optional] 
**OutputFileName** | Pointer to **string** | Output file name written under the per-node bundle directory. Must be a single file name of up to 128 characters from [A-Za-z0-9._-].  | [optional] 
**TimeoutSecs** | Pointer to **int64** | Command timeout in seconds. | [optional] 

## Methods

### NewYCQLComponentSpec

`func NewYCQLComponentSpec() *YCQLComponentSpec`

NewYCQLComponentSpec instantiates a new YCQLComponentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYCQLComponentSpecWithDefaults

`func NewYCQLComponentSpecWithDefaults() *YCQLComponentSpec`

NewYCQLComponentSpecWithDefaults instantiates a new YCQLComponentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentName

`func (o *YCQLComponentSpec) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *YCQLComponentSpec) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *YCQLComponentSpec) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *YCQLComponentSpec) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### GetKeyspace

`func (o *YCQLComponentSpec) GetKeyspace() string`

GetKeyspace returns the Keyspace field if non-nil, zero value otherwise.

### GetKeyspaceOk

`func (o *YCQLComponentSpec) GetKeyspaceOk() (*string, bool)`

GetKeyspaceOk returns a tuple with the Keyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspace

`func (o *YCQLComponentSpec) SetKeyspace(v string)`

SetKeyspace sets Keyspace field to given value.

### HasKeyspace

`func (o *YCQLComponentSpec) HasKeyspace() bool`

HasKeyspace returns a boolean if a field has been set.

### GetQueries

`func (o *YCQLComponentSpec) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *YCQLComponentSpec) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *YCQLComponentSpec) SetQueries(v []string)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *YCQLComponentSpec) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### GetOutputFileName

`func (o *YCQLComponentSpec) GetOutputFileName() string`

GetOutputFileName returns the OutputFileName field if non-nil, zero value otherwise.

### GetOutputFileNameOk

`func (o *YCQLComponentSpec) GetOutputFileNameOk() (*string, bool)`

GetOutputFileNameOk returns a tuple with the OutputFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFileName

`func (o *YCQLComponentSpec) SetOutputFileName(v string)`

SetOutputFileName sets OutputFileName field to given value.

### HasOutputFileName

`func (o *YCQLComponentSpec) HasOutputFileName() bool`

HasOutputFileName returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *YCQLComponentSpec) GetTimeoutSecs() int64`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *YCQLComponentSpec) GetTimeoutSecsOk() (*int64, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *YCQLComponentSpec) SetTimeoutSecs(v int64)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *YCQLComponentSpec) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


