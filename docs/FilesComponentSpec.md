# FilesComponentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentName** | Pointer to **string** | Logical component name; used as a label/sub-directory. | [optional] 
**LinuxUser** | Pointer to **string** | Linux user to run the script as on the node. | [optional] 
**Params** | Pointer to **[]string** | Arguments appended after the script (entrypoint + flags), e.g. [create_universe_logs_bundle, --mount_path, /mnt/d0, ...]. | [optional] 
**RemoteTarPath** | Pointer to **string** | Path on the node where the script writes the tar. May contain ${nodeName} which is substituted per node. | [optional] 
**ScriptPath** | Pointer to **string** | Path on the YBA host to the script file to execute on the node (e.g. bin/node_utils.sh). May be relative to yb.devops.home. | [optional] 
**TimeoutSecs** | Pointer to **int64** | Script execution timeout in seconds. | [optional] 

## Methods

### NewFilesComponentSpec

`func NewFilesComponentSpec() *FilesComponentSpec`

NewFilesComponentSpec instantiates a new FilesComponentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesComponentSpecWithDefaults

`func NewFilesComponentSpecWithDefaults() *FilesComponentSpec`

NewFilesComponentSpecWithDefaults instantiates a new FilesComponentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentName

`func (o *FilesComponentSpec) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *FilesComponentSpec) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *FilesComponentSpec) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *FilesComponentSpec) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### GetLinuxUser

`func (o *FilesComponentSpec) GetLinuxUser() string`

GetLinuxUser returns the LinuxUser field if non-nil, zero value otherwise.

### GetLinuxUserOk

`func (o *FilesComponentSpec) GetLinuxUserOk() (*string, bool)`

GetLinuxUserOk returns a tuple with the LinuxUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUser

`func (o *FilesComponentSpec) SetLinuxUser(v string)`

SetLinuxUser sets LinuxUser field to given value.

### HasLinuxUser

`func (o *FilesComponentSpec) HasLinuxUser() bool`

HasLinuxUser returns a boolean if a field has been set.

### GetParams

`func (o *FilesComponentSpec) GetParams() []string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *FilesComponentSpec) GetParamsOk() (*[]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *FilesComponentSpec) SetParams(v []string)`

SetParams sets Params field to given value.

### HasParams

`func (o *FilesComponentSpec) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetRemoteTarPath

`func (o *FilesComponentSpec) GetRemoteTarPath() string`

GetRemoteTarPath returns the RemoteTarPath field if non-nil, zero value otherwise.

### GetRemoteTarPathOk

`func (o *FilesComponentSpec) GetRemoteTarPathOk() (*string, bool)`

GetRemoteTarPathOk returns a tuple with the RemoteTarPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTarPath

`func (o *FilesComponentSpec) SetRemoteTarPath(v string)`

SetRemoteTarPath sets RemoteTarPath field to given value.

### HasRemoteTarPath

`func (o *FilesComponentSpec) HasRemoteTarPath() bool`

HasRemoteTarPath returns a boolean if a field has been set.

### GetScriptPath

`func (o *FilesComponentSpec) GetScriptPath() string`

GetScriptPath returns the ScriptPath field if non-nil, zero value otherwise.

### GetScriptPathOk

`func (o *FilesComponentSpec) GetScriptPathOk() (*string, bool)`

GetScriptPathOk returns a tuple with the ScriptPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptPath

`func (o *FilesComponentSpec) SetScriptPath(v string)`

SetScriptPath sets ScriptPath field to given value.

### HasScriptPath

`func (o *FilesComponentSpec) HasScriptPath() bool`

HasScriptPath returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *FilesComponentSpec) GetTimeoutSecs() int64`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *FilesComponentSpec) GetTimeoutSecsOk() (*int64, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *FilesComponentSpec) SetTimeoutSecs(v int64)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *FilesComponentSpec) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


