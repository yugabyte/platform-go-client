# UniverseLogsExporterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalTags** | Pointer to **map[string]string** | Additional tags | [optional] 
**ExporterUuid** | **string** | Exporter uuid | 

## Methods

### NewUniverseLogsExporterConfig

`func NewUniverseLogsExporterConfig(exporterUuid string, ) *UniverseLogsExporterConfig`

NewUniverseLogsExporterConfig instantiates a new UniverseLogsExporterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseLogsExporterConfigWithDefaults

`func NewUniverseLogsExporterConfigWithDefaults() *UniverseLogsExporterConfig`

NewUniverseLogsExporterConfigWithDefaults instantiates a new UniverseLogsExporterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalTags

`func (o *UniverseLogsExporterConfig) GetAdditionalTags() map[string]string`

GetAdditionalTags returns the AdditionalTags field if non-nil, zero value otherwise.

### GetAdditionalTagsOk

`func (o *UniverseLogsExporterConfig) GetAdditionalTagsOk() (*map[string]string, bool)`

GetAdditionalTagsOk returns a tuple with the AdditionalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalTags

`func (o *UniverseLogsExporterConfig) SetAdditionalTags(v map[string]string)`

SetAdditionalTags sets AdditionalTags field to given value.

### HasAdditionalTags

`func (o *UniverseLogsExporterConfig) HasAdditionalTags() bool`

HasAdditionalTags returns a boolean if a field has been set.

### GetExporterUuid

`func (o *UniverseLogsExporterConfig) GetExporterUuid() string`

GetExporterUuid returns the ExporterUuid field if non-nil, zero value otherwise.

### GetExporterUuidOk

`func (o *UniverseLogsExporterConfig) GetExporterUuidOk() (*string, bool)`

GetExporterUuidOk returns a tuple with the ExporterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExporterUuid

`func (o *UniverseLogsExporterConfig) SetExporterUuid(v string)`

SetExporterUuid sets ExporterUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


