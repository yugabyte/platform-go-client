# MasterLogsTelemetrySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exporters** | Pointer to [**[]UniverseServerLogsExporterConfig**](UniverseServerLogsExporterConfig.md) | List of exporters. Empty &#x3D; no export. | [optional] 
**MinLevel** | Pointer to **string** | Minimum yb-master glog severity to export. Lines below this level are dropped. | [optional] [default to "INFO"]
**NoiseSampleDropRatio** | Pointer to **float64** | Fraction (0.0-1.0) of high-volume, low-value noise log lines to drop. Defaults to 0.99 (drop 99%). Set to 0.0 to keep all such lines.  | [optional] [default to 0.99]

## Methods

### NewMasterLogsTelemetrySpec

`func NewMasterLogsTelemetrySpec() *MasterLogsTelemetrySpec`

NewMasterLogsTelemetrySpec instantiates a new MasterLogsTelemetrySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMasterLogsTelemetrySpecWithDefaults

`func NewMasterLogsTelemetrySpecWithDefaults() *MasterLogsTelemetrySpec`

NewMasterLogsTelemetrySpecWithDefaults instantiates a new MasterLogsTelemetrySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExporters

`func (o *MasterLogsTelemetrySpec) GetExporters() []UniverseServerLogsExporterConfig`

GetExporters returns the Exporters field if non-nil, zero value otherwise.

### GetExportersOk

`func (o *MasterLogsTelemetrySpec) GetExportersOk() (*[]UniverseServerLogsExporterConfig, bool)`

GetExportersOk returns a tuple with the Exporters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExporters

`func (o *MasterLogsTelemetrySpec) SetExporters(v []UniverseServerLogsExporterConfig)`

SetExporters sets Exporters field to given value.

### HasExporters

`func (o *MasterLogsTelemetrySpec) HasExporters() bool`

HasExporters returns a boolean if a field has been set.

### GetMinLevel

`func (o *MasterLogsTelemetrySpec) GetMinLevel() string`

GetMinLevel returns the MinLevel field if non-nil, zero value otherwise.

### GetMinLevelOk

`func (o *MasterLogsTelemetrySpec) GetMinLevelOk() (*string, bool)`

GetMinLevelOk returns a tuple with the MinLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLevel

`func (o *MasterLogsTelemetrySpec) SetMinLevel(v string)`

SetMinLevel sets MinLevel field to given value.

### HasMinLevel

`func (o *MasterLogsTelemetrySpec) HasMinLevel() bool`

HasMinLevel returns a boolean if a field has been set.

### GetNoiseSampleDropRatio

`func (o *MasterLogsTelemetrySpec) GetNoiseSampleDropRatio() float64`

GetNoiseSampleDropRatio returns the NoiseSampleDropRatio field if non-nil, zero value otherwise.

### GetNoiseSampleDropRatioOk

`func (o *MasterLogsTelemetrySpec) GetNoiseSampleDropRatioOk() (*float64, bool)`

GetNoiseSampleDropRatioOk returns a tuple with the NoiseSampleDropRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoiseSampleDropRatio

`func (o *MasterLogsTelemetrySpec) SetNoiseSampleDropRatio(v float64)`

SetNoiseSampleDropRatio sets NoiseSampleDropRatio field to given value.

### HasNoiseSampleDropRatio

`func (o *MasterLogsTelemetrySpec) HasNoiseSampleDropRatio() bool`

HasNoiseSampleDropRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


