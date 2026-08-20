# NodeAgentLogsTelemetrySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exporters** | Pointer to [**[]UniverseServerLogsExporterConfig**](UniverseServerLogsExporterConfig.md) | List of exporters. Empty &#x3D; no export. | [optional] 

## Methods

### NewNodeAgentLogsTelemetrySpec

`func NewNodeAgentLogsTelemetrySpec() *NodeAgentLogsTelemetrySpec`

NewNodeAgentLogsTelemetrySpec instantiates a new NodeAgentLogsTelemetrySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeAgentLogsTelemetrySpecWithDefaults

`func NewNodeAgentLogsTelemetrySpecWithDefaults() *NodeAgentLogsTelemetrySpec`

NewNodeAgentLogsTelemetrySpecWithDefaults instantiates a new NodeAgentLogsTelemetrySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExporters

`func (o *NodeAgentLogsTelemetrySpec) GetExporters() []UniverseServerLogsExporterConfig`

GetExporters returns the Exporters field if non-nil, zero value otherwise.

### GetExportersOk

`func (o *NodeAgentLogsTelemetrySpec) GetExportersOk() (*[]UniverseServerLogsExporterConfig, bool)`

GetExportersOk returns a tuple with the Exporters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExporters

`func (o *NodeAgentLogsTelemetrySpec) SetExporters(v []UniverseServerLogsExporterConfig)`

SetExporters sets Exporters field to given value.

### HasExporters

`func (o *NodeAgentLogsTelemetrySpec) HasExporters() bool`

HasExporters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


