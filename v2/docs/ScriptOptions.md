# ScriptOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScriptContent** | Pointer to **string** | The inline script content to execute on nodes. Either script_content or script_file must be provided, but not both.  | [optional] 
**ScriptFile** | Pointer to **string** | Path to a script file on the YugabyteDB Anywhere node to execute on database nodes. Either script_content or script_file must be provided, but not both.  | [optional] 
**Params** | Pointer to **[]string** | Parameters to pass to the script as command-line arguments. | [optional] 
**TimeoutSecs** | Pointer to **int64** | Timeout in seconds for the script execution on each node. If elevated privileges are needed, include sudo commands directly in your script.  | [optional] [default to 60]
**LinuxUser** | Pointer to **string** | Run the script as a particular Linux user. Defaults to yugabyte. When using Node Agent, the command is executed as this user. When using SSH fallback, this sets the SSH user.  | [optional] [default to "yugabyte"]
**MaxScriptFileSizeBytes** | Pointer to **int64** | Maximum size in bytes for script file content to be captured in audit logs. Script files larger than this limit will be executed, but their content will not be recorded in the audit log. This only applies when using script_file.  | [optional] [default to 1048576]

## Methods

### NewScriptOptions

`func NewScriptOptions() *ScriptOptions`

NewScriptOptions instantiates a new ScriptOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptOptionsWithDefaults

`func NewScriptOptionsWithDefaults() *ScriptOptions`

NewScriptOptionsWithDefaults instantiates a new ScriptOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScriptContent

`func (o *ScriptOptions) GetScriptContent() string`

GetScriptContent returns the ScriptContent field if non-nil, zero value otherwise.

### GetScriptContentOk

`func (o *ScriptOptions) GetScriptContentOk() (*string, bool)`

GetScriptContentOk returns a tuple with the ScriptContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptContent

`func (o *ScriptOptions) SetScriptContent(v string)`

SetScriptContent sets ScriptContent field to given value.

### HasScriptContent

`func (o *ScriptOptions) HasScriptContent() bool`

HasScriptContent returns a boolean if a field has been set.

### GetScriptFile

`func (o *ScriptOptions) GetScriptFile() string`

GetScriptFile returns the ScriptFile field if non-nil, zero value otherwise.

### GetScriptFileOk

`func (o *ScriptOptions) GetScriptFileOk() (*string, bool)`

GetScriptFileOk returns a tuple with the ScriptFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptFile

`func (o *ScriptOptions) SetScriptFile(v string)`

SetScriptFile sets ScriptFile field to given value.

### HasScriptFile

`func (o *ScriptOptions) HasScriptFile() bool`

HasScriptFile returns a boolean if a field has been set.

### GetParams

`func (o *ScriptOptions) GetParams() []string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ScriptOptions) GetParamsOk() (*[]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ScriptOptions) SetParams(v []string)`

SetParams sets Params field to given value.

### HasParams

`func (o *ScriptOptions) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *ScriptOptions) GetTimeoutSecs() int64`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *ScriptOptions) GetTimeoutSecsOk() (*int64, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *ScriptOptions) SetTimeoutSecs(v int64)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *ScriptOptions) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### GetLinuxUser

`func (o *ScriptOptions) GetLinuxUser() string`

GetLinuxUser returns the LinuxUser field if non-nil, zero value otherwise.

### GetLinuxUserOk

`func (o *ScriptOptions) GetLinuxUserOk() (*string, bool)`

GetLinuxUserOk returns a tuple with the LinuxUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUser

`func (o *ScriptOptions) SetLinuxUser(v string)`

SetLinuxUser sets LinuxUser field to given value.

### HasLinuxUser

`func (o *ScriptOptions) HasLinuxUser() bool`

HasLinuxUser returns a boolean if a field has been set.

### GetMaxScriptFileSizeBytes

`func (o *ScriptOptions) GetMaxScriptFileSizeBytes() int64`

GetMaxScriptFileSizeBytes returns the MaxScriptFileSizeBytes field if non-nil, zero value otherwise.

### GetMaxScriptFileSizeBytesOk

`func (o *ScriptOptions) GetMaxScriptFileSizeBytesOk() (*int64, bool)`

GetMaxScriptFileSizeBytesOk returns a tuple with the MaxScriptFileSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScriptFileSizeBytes

`func (o *ScriptOptions) SetMaxScriptFileSizeBytes(v int64)`

SetMaxScriptFileSizeBytes sets MaxScriptFileSizeBytes field to given value.

### HasMaxScriptFileSizeBytes

`func (o *ScriptOptions) HasMaxScriptFileSizeBytes() bool`

HasMaxScriptFileSizeBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


