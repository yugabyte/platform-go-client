# TServerLogsTelemetrySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exporters** | Pointer to [**[]UniverseServerLogsExporterConfig**](UniverseServerLogsExporterConfig.md) | List of exporters. Empty &#x3D; no export. | [optional] 
**MinLevel** | Pointer to **string** | Minimum yb-tserver glog severity to export. Lines below this level are dropped. Defaults to WARNING (yb-tserver INFO is very high volume).  | [optional] [default to "WARNING"]

## Methods

### NewTServerLogsTelemetrySpec

`func NewTServerLogsTelemetrySpec() *TServerLogsTelemetrySpec`

NewTServerLogsTelemetrySpec instantiates a new TServerLogsTelemetrySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTServerLogsTelemetrySpecWithDefaults

`func NewTServerLogsTelemetrySpecWithDefaults() *TServerLogsTelemetrySpec`

NewTServerLogsTelemetrySpecWithDefaults instantiates a new TServerLogsTelemetrySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExporters

`func (o *TServerLogsTelemetrySpec) GetExporters() []UniverseServerLogsExporterConfig`

GetExporters returns the Exporters field if non-nil, zero value otherwise.

### GetExportersOk

`func (o *TServerLogsTelemetrySpec) GetExportersOk() (*[]UniverseServerLogsExporterConfig, bool)`

GetExportersOk returns a tuple with the Exporters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExporters

`func (o *TServerLogsTelemetrySpec) SetExporters(v []UniverseServerLogsExporterConfig)`

SetExporters sets Exporters field to given value.

### HasExporters

`func (o *TServerLogsTelemetrySpec) HasExporters() bool`

HasExporters returns a boolean if a field has been set.

### GetMinLevel

`func (o *TServerLogsTelemetrySpec) GetMinLevel() string`

GetMinLevel returns the MinLevel field if non-nil, zero value otherwise.

### GetMinLevelOk

`func (o *TServerLogsTelemetrySpec) GetMinLevelOk() (*string, bool)`

GetMinLevelOk returns a tuple with the MinLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLevel

`func (o *TServerLogsTelemetrySpec) SetMinLevel(v string)`

SetMinLevel sets MinLevel field to given value.

### HasMinLevel

`func (o *TServerLogsTelemetrySpec) HasMinLevel() bool`

HasMinLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


