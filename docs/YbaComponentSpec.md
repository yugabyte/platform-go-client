# YbaComponentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentName** | Pointer to **string** | Logical component name; used as a label/sub-directory. | [optional] 
**LinuxUser** | Pointer to **string** | Ignored for local YBA execution; accepted for API parity with node-level specs. | [optional] 
**Params** | Pointer to **[]string** | Arguments appended after the script (entrypoint + flags), e.g. [create_application_logs_bundle, --log_dir, /var/log/yba, ...]. | [optional] 
**RemoteTarPath** | Pointer to **string** | Path on the YBA host where the script writes the tar before it is untarred. | [optional] 
**ScriptPath** | Pointer to **string** | Path on the YBA host to the script file to execute locally (e.g. bin/yba_utils.sh). May be relative to yb.devops.home. | [optional] 
**TimeoutSecs** | Pointer to **int64** | Script execution timeout in seconds. | [optional] 

## Methods

### NewYbaComponentSpec

`func NewYbaComponentSpec() *YbaComponentSpec`

NewYbaComponentSpec instantiates a new YbaComponentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYbaComponentSpecWithDefaults

`func NewYbaComponentSpecWithDefaults() *YbaComponentSpec`

NewYbaComponentSpecWithDefaults instantiates a new YbaComponentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentName

`func (o *YbaComponentSpec) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *YbaComponentSpec) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *YbaComponentSpec) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *YbaComponentSpec) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### GetLinuxUser

`func (o *YbaComponentSpec) GetLinuxUser() string`

GetLinuxUser returns the LinuxUser field if non-nil, zero value otherwise.

### GetLinuxUserOk

`func (o *YbaComponentSpec) GetLinuxUserOk() (*string, bool)`

GetLinuxUserOk returns a tuple with the LinuxUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUser

`func (o *YbaComponentSpec) SetLinuxUser(v string)`

SetLinuxUser sets LinuxUser field to given value.

### HasLinuxUser

`func (o *YbaComponentSpec) HasLinuxUser() bool`

HasLinuxUser returns a boolean if a field has been set.

### GetParams

`func (o *YbaComponentSpec) GetParams() []string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *YbaComponentSpec) GetParamsOk() (*[]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *YbaComponentSpec) SetParams(v []string)`

SetParams sets Params field to given value.

### HasParams

`func (o *YbaComponentSpec) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetRemoteTarPath

`func (o *YbaComponentSpec) GetRemoteTarPath() string`

GetRemoteTarPath returns the RemoteTarPath field if non-nil, zero value otherwise.

### GetRemoteTarPathOk

`func (o *YbaComponentSpec) GetRemoteTarPathOk() (*string, bool)`

GetRemoteTarPathOk returns a tuple with the RemoteTarPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTarPath

`func (o *YbaComponentSpec) SetRemoteTarPath(v string)`

SetRemoteTarPath sets RemoteTarPath field to given value.

### HasRemoteTarPath

`func (o *YbaComponentSpec) HasRemoteTarPath() bool`

HasRemoteTarPath returns a boolean if a field has been set.

### GetScriptPath

`func (o *YbaComponentSpec) GetScriptPath() string`

GetScriptPath returns the ScriptPath field if non-nil, zero value otherwise.

### GetScriptPathOk

`func (o *YbaComponentSpec) GetScriptPathOk() (*string, bool)`

GetScriptPathOk returns a tuple with the ScriptPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPath

`func (o *YbaComponentSpec) SetScriptPath(v string)`

SetScriptPath sets ScriptPath field to given value.

### HasScriptPath

`func (o *YbaComponentSpec) HasScriptPath() bool`

HasScriptPath returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *YbaComponentSpec) GetTimeoutSecs() int64`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *YbaComponentSpec) GetTimeoutSecsOk() (*int64, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *YbaComponentSpec) SetTimeoutSecs(v int64)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *YbaComponentSpec) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


