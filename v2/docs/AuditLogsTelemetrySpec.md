# AuditLogsTelemetrySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**YsqlAuditConfig** | Pointer to [**YSQLAuditConfig**](YSQLAuditConfig.md) |  | [optional] 
**YcqlAuditConfig** | Pointer to [**YCQLAuditConfig**](YCQLAuditConfig.md) |  | [optional] 
**Exporters** | Pointer to [**[]TelemetryExporterEntry**](TelemetryExporterEntry.md) | List of exporters. Empty &#x3D; no export. | [optional] 

## Methods

### NewAuditLogsTelemetrySpec

`func NewAuditLogsTelemetrySpec() *AuditLogsTelemetrySpec`

NewAuditLogsTelemetrySpec instantiates a new AuditLogsTelemetrySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogsTelemetrySpecWithDefaults

`func NewAuditLogsTelemetrySpecWithDefaults() *AuditLogsTelemetrySpec`

NewAuditLogsTelemetrySpecWithDefaults instantiates a new AuditLogsTelemetrySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetYsqlAuditConfig

`func (o *AuditLogsTelemetrySpec) GetYsqlAuditConfig() YSQLAuditConfig`

GetYsqlAuditConfig returns the YsqlAuditConfig field if non-nil, zero value otherwise.

### GetYsqlAuditConfigOk

`func (o *AuditLogsTelemetrySpec) GetYsqlAuditConfigOk() (*YSQLAuditConfig, bool)`

GetYsqlAuditConfigOk returns a tuple with the YsqlAuditConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlAuditConfig

`func (o *AuditLogsTelemetrySpec) SetYsqlAuditConfig(v YSQLAuditConfig)`

SetYsqlAuditConfig sets YsqlAuditConfig field to given value.

### HasYsqlAuditConfig

`func (o *AuditLogsTelemetrySpec) HasYsqlAuditConfig() bool`

HasYsqlAuditConfig returns a boolean if a field has been set.

### GetYcqlAuditConfig

`func (o *AuditLogsTelemetrySpec) GetYcqlAuditConfig() YCQLAuditConfig`

GetYcqlAuditConfig returns the YcqlAuditConfig field if non-nil, zero value otherwise.

### GetYcqlAuditConfigOk

`func (o *AuditLogsTelemetrySpec) GetYcqlAuditConfigOk() (*YCQLAuditConfig, bool)`

GetYcqlAuditConfigOk returns a tuple with the YcqlAuditConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYcqlAuditConfig

`func (o *AuditLogsTelemetrySpec) SetYcqlAuditConfig(v YCQLAuditConfig)`

SetYcqlAuditConfig sets YcqlAuditConfig field to given value.

### HasYcqlAuditConfig

`func (o *AuditLogsTelemetrySpec) HasYcqlAuditConfig() bool`

HasYcqlAuditConfig returns a boolean if a field has been set.

### GetExporters

`func (o *AuditLogsTelemetrySpec) GetExporters() []TelemetryExporterEntry`

GetExporters returns the Exporters field if non-nil, zero value otherwise.

### GetExportersOk

`func (o *AuditLogsTelemetrySpec) GetExportersOk() (*[]TelemetryExporterEntry, bool)`

GetExportersOk returns a tuple with the Exporters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExporters

`func (o *AuditLogsTelemetrySpec) SetExporters(v []TelemetryExporterEntry)`

SetExporters sets Exporters field to given value.

### HasExporters

`func (o *AuditLogsTelemetrySpec) HasExporters() bool`

HasExporters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


