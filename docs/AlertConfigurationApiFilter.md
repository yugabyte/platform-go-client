# AlertConfigurationApiFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Whether the alert configuration is active. | [optional] 
**DestinationType** | Pointer to **string** | The destination type of the alert configuration.  | [optional] 
**DestinationUuid** | Pointer to **string** | The destination uuid of the alert configuration.  | [optional] 
**Name** | Pointer to **string** | The name of the alert configuration. | [optional] 
**Severity** | Pointer to **string** | The severity of the alert configuration. | [optional] 
**Target** | Pointer to [**AlertConfigurationTarget**](AlertConfigurationTarget.md) |  | [optional] 
**TargetType** | Pointer to **string** | The target type of the alert configuration. | [optional] 
**Template** | Pointer to **string** | The template of the alert configuration. | [optional] 
**Uuids** | Pointer to **[]string** | The uuids of the alert configurations. | [optional] 

## Methods

### NewAlertConfigurationApiFilter

`func NewAlertConfigurationApiFilter() *AlertConfigurationApiFilter`

NewAlertConfigurationApiFilter instantiates a new AlertConfigurationApiFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertConfigurationApiFilterWithDefaults

`func NewAlertConfigurationApiFilterWithDefaults() *AlertConfigurationApiFilter`

NewAlertConfigurationApiFilterWithDefaults instantiates a new AlertConfigurationApiFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *AlertConfigurationApiFilter) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AlertConfigurationApiFilter) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AlertConfigurationApiFilter) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AlertConfigurationApiFilter) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDestinationType

`func (o *AlertConfigurationApiFilter) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *AlertConfigurationApiFilter) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *AlertConfigurationApiFilter) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *AlertConfigurationApiFilter) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetDestinationUuid

`func (o *AlertConfigurationApiFilter) GetDestinationUuid() string`

GetDestinationUuid returns the DestinationUuid field if non-nil, zero value otherwise.

### GetDestinationUuidOk

`func (o *AlertConfigurationApiFilter) GetDestinationUuidOk() (*string, bool)`

GetDestinationUuidOk returns a tuple with the DestinationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationUuid

`func (o *AlertConfigurationApiFilter) SetDestinationUuid(v string)`

SetDestinationUuid sets DestinationUuid field to given value.

### HasDestinationUuid

`func (o *AlertConfigurationApiFilter) HasDestinationUuid() bool`

HasDestinationUuid returns a boolean if a field has been set.

### GetName

`func (o *AlertConfigurationApiFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertConfigurationApiFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertConfigurationApiFilter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AlertConfigurationApiFilter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSeverity

`func (o *AlertConfigurationApiFilter) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertConfigurationApiFilter) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertConfigurationApiFilter) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AlertConfigurationApiFilter) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTarget

`func (o *AlertConfigurationApiFilter) GetTarget() AlertConfigurationTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AlertConfigurationApiFilter) GetTargetOk() (*AlertConfigurationTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AlertConfigurationApiFilter) SetTarget(v AlertConfigurationTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AlertConfigurationApiFilter) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTargetType

`func (o *AlertConfigurationApiFilter) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *AlertConfigurationApiFilter) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *AlertConfigurationApiFilter) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *AlertConfigurationApiFilter) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetTemplate

`func (o *AlertConfigurationApiFilter) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *AlertConfigurationApiFilter) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *AlertConfigurationApiFilter) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *AlertConfigurationApiFilter) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetUuids

`func (o *AlertConfigurationApiFilter) GetUuids() []string`

GetUuids returns the Uuids field if non-nil, zero value otherwise.

### GetUuidsOk

`func (o *AlertConfigurationApiFilter) GetUuidsOk() (*[]string, bool)`

GetUuidsOk returns a tuple with the Uuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuids

`func (o *AlertConfigurationApiFilter) SetUuids(v []string)`

SetUuids sets Uuids field to given value.

### HasUuids

`func (o *AlertConfigurationApiFilter) HasUuids() bool`

HasUuids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


