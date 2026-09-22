# YbAdminComponentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentName** | Pointer to **string** | Logical component name; used as output file label. | [optional] 
**OutputFileName** | Pointer to **string** | Output file name written under the per-node bundle directory. | [optional] 
**TimeoutSecs** | Pointer to **int64** | Command timeout in seconds. | [optional] 
**YbAdminArgs** | Pointer to **[]string** | Additional arguments after the subcommand. | [optional] 
**YbAdminCommands** | Pointer to **[]string** | yb-admin subcommands to run (each executed separately in one batch). | [optional] 

## Methods

### NewYbAdminComponentSpec

`func NewYbAdminComponentSpec() *YbAdminComponentSpec`

NewYbAdminComponentSpec instantiates a new YbAdminComponentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYbAdminComponentSpecWithDefaults

`func NewYbAdminComponentSpecWithDefaults() *YbAdminComponentSpec`

NewYbAdminComponentSpecWithDefaults instantiates a new YbAdminComponentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentName

`func (o *YbAdminComponentSpec) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *YbAdminComponentSpec) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *YbAdminComponentSpec) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *YbAdminComponentSpec) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### GetOutputFileName

`func (o *YbAdminComponentSpec) GetOutputFileName() string`

GetOutputFileName returns the OutputFileName field if non-nil, zero value otherwise.

### GetOutputFileNameOk

`func (o *YbAdminComponentSpec) GetOutputFileNameOk() (*string, bool)`

GetOutputFileNameOk returns a tuple with the OutputFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFileName

`func (o *YbAdminComponentSpec) SetOutputFileName(v string)`

SetOutputFileName sets OutputFileName field to given value.

### HasOutputFileName

`func (o *YbAdminComponentSpec) HasOutputFileName() bool`

HasOutputFileName returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *YbAdminComponentSpec) GetTimeoutSecs() int64`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *YbAdminComponentSpec) GetTimeoutSecsOk() (*int64, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *YbAdminComponentSpec) SetTimeoutSecs(v int64)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *YbAdminComponentSpec) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### GetYbAdminArgs

`func (o *YbAdminComponentSpec) GetYbAdminArgs() []string`

GetYbAdminArgs returns the YbAdminArgs field if non-nil, zero value otherwise.

### GetYbAdminArgsOk

`func (o *YbAdminComponentSpec) GetYbAdminArgsOk() (*[]string, bool)`

GetYbAdminArgsOk returns a tuple with the YbAdminArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbAdminArgs

`func (o *YbAdminComponentSpec) SetYbAdminArgs(v []string)`

SetYbAdminArgs sets YbAdminArgs field to given value.

### HasYbAdminArgs

`func (o *YbAdminComponentSpec) HasYbAdminArgs() bool`

HasYbAdminArgs returns a boolean if a field has been set.

### GetYbAdminCommands

`func (o *YbAdminComponentSpec) GetYbAdminCommands() []string`

GetYbAdminCommands returns the YbAdminCommands field if non-nil, zero value otherwise.

### GetYbAdminCommandsOk

`func (o *YbAdminComponentSpec) GetYbAdminCommandsOk() (*[]string, bool)`

GetYbAdminCommandsOk returns a tuple with the YbAdminCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbAdminCommands

`func (o *YbAdminComponentSpec) SetYbAdminCommands(v []string)`

SetYbAdminCommands sets YbAdminCommands field to given value.

### HasYbAdminCommands

`func (o *YbAdminComponentSpec) HasYbAdminCommands() bool`

HasYbAdminCommands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


