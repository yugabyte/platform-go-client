# UniverseRestartAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RollingRestart** | Pointer to **bool** | Perform a rolling restart of the universe. Otherwise, all nodes will be restarted at the same time. | [optional] [default to true]
**RestartType** | Pointer to **string** | The method to reboot the node. This is not required for kubernetes universes, as the pods  will get restarted no matter what. \&quot;HARD\&quot; reboots are not supported today.  OS: Restarts the node via the operating system. SERVICE: Restart the YugabyteDB Process only (master, tserver, etc).  | [optional] [default to "SERVICE"]

## Methods

### NewUniverseRestartAllOf

`func NewUniverseRestartAllOf() *UniverseRestartAllOf`

NewUniverseRestartAllOf instantiates a new UniverseRestartAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseRestartAllOfWithDefaults

`func NewUniverseRestartAllOfWithDefaults() *UniverseRestartAllOf`

NewUniverseRestartAllOfWithDefaults instantiates a new UniverseRestartAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRollingRestart

`func (o *UniverseRestartAllOf) GetRollingRestart() bool`

GetRollingRestart returns the RollingRestart field if non-nil, zero value otherwise.

### GetRollingRestartOk

`func (o *UniverseRestartAllOf) GetRollingRestartOk() (*bool, bool)`

GetRollingRestartOk returns a tuple with the RollingRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingRestart

`func (o *UniverseRestartAllOf) SetRollingRestart(v bool)`

SetRollingRestart sets RollingRestart field to given value.

### HasRollingRestart

`func (o *UniverseRestartAllOf) HasRollingRestart() bool`

HasRollingRestart returns a boolean if a field has been set.

### GetRestartType

`func (o *UniverseRestartAllOf) GetRestartType() string`

GetRestartType returns the RestartType field if non-nil, zero value otherwise.

### GetRestartTypeOk

`func (o *UniverseRestartAllOf) GetRestartTypeOk() (*string, bool)`

GetRestartTypeOk returns a tuple with the RestartType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartType

`func (o *UniverseRestartAllOf) SetRestartType(v string)`

SetRestartType sets RestartType field to given value.

### HasRestartType

`func (o *UniverseRestartAllOf) HasRestartType() bool`

HasRestartType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


