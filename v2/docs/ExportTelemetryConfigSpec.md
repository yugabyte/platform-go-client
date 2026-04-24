# ExportTelemetryConfigSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpgradeOptions** | Pointer to [**ExportTelemetryUpgradeOptions**](ExportTelemetryUpgradeOptions.md) |  | [optional] 
**TelemetryConfig** | Pointer to [**TelemetryConfig**](TelemetryConfig.md) |  | [optional] 

## Methods

### NewExportTelemetryConfigSpec

`func NewExportTelemetryConfigSpec() *ExportTelemetryConfigSpec`

NewExportTelemetryConfigSpec instantiates a new ExportTelemetryConfigSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportTelemetryConfigSpecWithDefaults

`func NewExportTelemetryConfigSpecWithDefaults() *ExportTelemetryConfigSpec`

NewExportTelemetryConfigSpecWithDefaults instantiates a new ExportTelemetryConfigSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgradeOptions

`func (o *ExportTelemetryConfigSpec) GetUpgradeOptions() ExportTelemetryUpgradeOptions`

GetUpgradeOptions returns the UpgradeOptions field if non-nil, zero value otherwise.

### GetUpgradeOptionsOk

`func (o *ExportTelemetryConfigSpec) GetUpgradeOptionsOk() (*ExportTelemetryUpgradeOptions, bool)`

GetUpgradeOptionsOk returns a tuple with the UpgradeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeOptions

`func (o *ExportTelemetryConfigSpec) SetUpgradeOptions(v ExportTelemetryUpgradeOptions)`

SetUpgradeOptions sets UpgradeOptions field to given value.

### HasUpgradeOptions

`func (o *ExportTelemetryConfigSpec) HasUpgradeOptions() bool`

HasUpgradeOptions returns a boolean if a field has been set.

### GetTelemetryConfig

`func (o *ExportTelemetryConfigSpec) GetTelemetryConfig() TelemetryConfig`

GetTelemetryConfig returns the TelemetryConfig field if non-nil, zero value otherwise.

### GetTelemetryConfigOk

`func (o *ExportTelemetryConfigSpec) GetTelemetryConfigOk() (*TelemetryConfig, bool)`

GetTelemetryConfigOk returns a tuple with the TelemetryConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemetryConfig

`func (o *ExportTelemetryConfigSpec) SetTelemetryConfig(v TelemetryConfig)`

SetTelemetryConfig sets TelemetryConfig field to given value.

### HasTelemetryConfig

`func (o *ExportTelemetryConfigSpec) HasTelemetryConfig() bool`

HasTelemetryConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


