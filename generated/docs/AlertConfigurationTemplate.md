# AlertConfigurationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Is configured alerts raised or not | 
**AlertCount** | **float64** |  | 
**CreateTime** | **time.Time** | Creation time | [readonly] 
**CustomerUUID** | **string** | Customer UUID | [readonly] 
**DefaultDestination** | **bool** | Is default destination used for this config | 
**Description** | **string** | Description | 
**DestinationUUID** | Pointer to **string** | Alert destination UUID | [optional] 
**DurationSec** | **int32** | Duration in seconds, while condition is met to raise an alert | 
**MaintenanceWindowUuids** | Pointer to **[]string** | Maintenance window UUIDs, applied to this alert config | [optional] [readonly] 
**Name** | **string** | Name | 
**Target** | [**AlertConfigurationTarget**](AlertConfigurationTarget.md) |  | 
**TargetType** | **string** | Target type | [readonly] 
**Template** | **string** | Template name | [readonly] 
**ThresholdConditionReadOnly** | Pointer to **bool** | Is alert threshold condition read-only or configurable | [optional] [readonly] 
**ThresholdInteger** | Pointer to **bool** | Is alert threshold integer or floating point | [optional] [readonly] 
**ThresholdMaxValue** | Pointer to **float64** | Alert threshold maximal value | [optional] [readonly] 
**ThresholdMinValue** | Pointer to **float64** | Alert threshold minimal value | [optional] [readonly] 
**ThresholdReadOnly** | Pointer to **bool** | Is alert threshold read-only or configurable | [optional] [readonly] 
**ThresholdUnit** | **string** | Threshold unit | [readonly] 
**ThresholdUnitName** | Pointer to **string** | Threshold unit name | [optional] [readonly] 
**Thresholds** | [**map[string]AlertConfigurationThreshold**](AlertConfigurationThreshold.md) | Thresholds | 
**Uuid** | Pointer to **string** | Configuration UUID | [optional] [readonly] 

## Methods

### NewAlertConfigurationTemplate

`func NewAlertConfigurationTemplate(active bool, alertCount float64, createTime time.Time, customerUUID string, defaultDestination bool, description string, durationSec int32, name string, target AlertConfigurationTarget, targetType string, template string, thresholdUnit string, thresholds map[string]AlertConfigurationThreshold, ) *AlertConfigurationTemplate`

NewAlertConfigurationTemplate instantiates a new AlertConfigurationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertConfigurationTemplateWithDefaults

`func NewAlertConfigurationTemplateWithDefaults() *AlertConfigurationTemplate`

NewAlertConfigurationTemplateWithDefaults instantiates a new AlertConfigurationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *AlertConfigurationTemplate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AlertConfigurationTemplate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AlertConfigurationTemplate) SetActive(v bool)`

SetActive sets Active field to given value.


### GetAlertCount

`func (o *AlertConfigurationTemplate) GetAlertCount() float64`

GetAlertCount returns the AlertCount field if non-nil, zero value otherwise.

### GetAlertCountOk

`func (o *AlertConfigurationTemplate) GetAlertCountOk() (*float64, bool)`

GetAlertCountOk returns a tuple with the AlertCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCount

`func (o *AlertConfigurationTemplate) SetAlertCount(v float64)`

SetAlertCount sets AlertCount field to given value.


### GetCreateTime

`func (o *AlertConfigurationTemplate) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *AlertConfigurationTemplate) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *AlertConfigurationTemplate) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.


### GetCustomerUUID

`func (o *AlertConfigurationTemplate) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *AlertConfigurationTemplate) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *AlertConfigurationTemplate) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetDefaultDestination

`func (o *AlertConfigurationTemplate) GetDefaultDestination() bool`

GetDefaultDestination returns the DefaultDestination field if non-nil, zero value otherwise.

### GetDefaultDestinationOk

`func (o *AlertConfigurationTemplate) GetDefaultDestinationOk() (*bool, bool)`

GetDefaultDestinationOk returns a tuple with the DefaultDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestination

`func (o *AlertConfigurationTemplate) SetDefaultDestination(v bool)`

SetDefaultDestination sets DefaultDestination field to given value.


### GetDescription

`func (o *AlertConfigurationTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertConfigurationTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertConfigurationTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDestinationUUID

`func (o *AlertConfigurationTemplate) GetDestinationUUID() string`

GetDestinationUUID returns the DestinationUUID field if non-nil, zero value otherwise.

### GetDestinationUUIDOk

`func (o *AlertConfigurationTemplate) GetDestinationUUIDOk() (*string, bool)`

GetDestinationUUIDOk returns a tuple with the DestinationUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationUUID

`func (o *AlertConfigurationTemplate) SetDestinationUUID(v string)`

SetDestinationUUID sets DestinationUUID field to given value.

### HasDestinationUUID

`func (o *AlertConfigurationTemplate) HasDestinationUUID() bool`

HasDestinationUUID returns a boolean if a field has been set.

### GetDurationSec

`func (o *AlertConfigurationTemplate) GetDurationSec() int32`

GetDurationSec returns the DurationSec field if non-nil, zero value otherwise.

### GetDurationSecOk

`func (o *AlertConfigurationTemplate) GetDurationSecOk() (*int32, bool)`

GetDurationSecOk returns a tuple with the DurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSec

`func (o *AlertConfigurationTemplate) SetDurationSec(v int32)`

SetDurationSec sets DurationSec field to given value.


### GetMaintenanceWindowUuids

`func (o *AlertConfigurationTemplate) GetMaintenanceWindowUuids() []string`

GetMaintenanceWindowUuids returns the MaintenanceWindowUuids field if non-nil, zero value otherwise.

### GetMaintenanceWindowUuidsOk

`func (o *AlertConfigurationTemplate) GetMaintenanceWindowUuidsOk() (*[]string, bool)`

GetMaintenanceWindowUuidsOk returns a tuple with the MaintenanceWindowUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindowUuids

`func (o *AlertConfigurationTemplate) SetMaintenanceWindowUuids(v []string)`

SetMaintenanceWindowUuids sets MaintenanceWindowUuids field to given value.

### HasMaintenanceWindowUuids

`func (o *AlertConfigurationTemplate) HasMaintenanceWindowUuids() bool`

HasMaintenanceWindowUuids returns a boolean if a field has been set.

### GetName

`func (o *AlertConfigurationTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertConfigurationTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertConfigurationTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetTarget

`func (o *AlertConfigurationTemplate) GetTarget() AlertConfigurationTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AlertConfigurationTemplate) GetTargetOk() (*AlertConfigurationTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AlertConfigurationTemplate) SetTarget(v AlertConfigurationTarget)`

SetTarget sets Target field to given value.


### GetTargetType

`func (o *AlertConfigurationTemplate) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *AlertConfigurationTemplate) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *AlertConfigurationTemplate) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.


### GetTemplate

`func (o *AlertConfigurationTemplate) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *AlertConfigurationTemplate) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *AlertConfigurationTemplate) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetThresholdConditionReadOnly

`func (o *AlertConfigurationTemplate) GetThresholdConditionReadOnly() bool`

GetThresholdConditionReadOnly returns the ThresholdConditionReadOnly field if non-nil, zero value otherwise.

### GetThresholdConditionReadOnlyOk

`func (o *AlertConfigurationTemplate) GetThresholdConditionReadOnlyOk() (*bool, bool)`

GetThresholdConditionReadOnlyOk returns a tuple with the ThresholdConditionReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdConditionReadOnly

`func (o *AlertConfigurationTemplate) SetThresholdConditionReadOnly(v bool)`

SetThresholdConditionReadOnly sets ThresholdConditionReadOnly field to given value.

### HasThresholdConditionReadOnly

`func (o *AlertConfigurationTemplate) HasThresholdConditionReadOnly() bool`

HasThresholdConditionReadOnly returns a boolean if a field has been set.

### GetThresholdInteger

`func (o *AlertConfigurationTemplate) GetThresholdInteger() bool`

GetThresholdInteger returns the ThresholdInteger field if non-nil, zero value otherwise.

### GetThresholdIntegerOk

`func (o *AlertConfigurationTemplate) GetThresholdIntegerOk() (*bool, bool)`

GetThresholdIntegerOk returns a tuple with the ThresholdInteger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdInteger

`func (o *AlertConfigurationTemplate) SetThresholdInteger(v bool)`

SetThresholdInteger sets ThresholdInteger field to given value.

### HasThresholdInteger

`func (o *AlertConfigurationTemplate) HasThresholdInteger() bool`

HasThresholdInteger returns a boolean if a field has been set.

### GetThresholdMaxValue

`func (o *AlertConfigurationTemplate) GetThresholdMaxValue() float64`

GetThresholdMaxValue returns the ThresholdMaxValue field if non-nil, zero value otherwise.

### GetThresholdMaxValueOk

`func (o *AlertConfigurationTemplate) GetThresholdMaxValueOk() (*float64, bool)`

GetThresholdMaxValueOk returns a tuple with the ThresholdMaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMaxValue

`func (o *AlertConfigurationTemplate) SetThresholdMaxValue(v float64)`

SetThresholdMaxValue sets ThresholdMaxValue field to given value.

### HasThresholdMaxValue

`func (o *AlertConfigurationTemplate) HasThresholdMaxValue() bool`

HasThresholdMaxValue returns a boolean if a field has been set.

### GetThresholdMinValue

`func (o *AlertConfigurationTemplate) GetThresholdMinValue() float64`

GetThresholdMinValue returns the ThresholdMinValue field if non-nil, zero value otherwise.

### GetThresholdMinValueOk

`func (o *AlertConfigurationTemplate) GetThresholdMinValueOk() (*float64, bool)`

GetThresholdMinValueOk returns a tuple with the ThresholdMinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMinValue

`func (o *AlertConfigurationTemplate) SetThresholdMinValue(v float64)`

SetThresholdMinValue sets ThresholdMinValue field to given value.

### HasThresholdMinValue

`func (o *AlertConfigurationTemplate) HasThresholdMinValue() bool`

HasThresholdMinValue returns a boolean if a field has been set.

### GetThresholdReadOnly

`func (o *AlertConfigurationTemplate) GetThresholdReadOnly() bool`

GetThresholdReadOnly returns the ThresholdReadOnly field if non-nil, zero value otherwise.

### GetThresholdReadOnlyOk

`func (o *AlertConfigurationTemplate) GetThresholdReadOnlyOk() (*bool, bool)`

GetThresholdReadOnlyOk returns a tuple with the ThresholdReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdReadOnly

`func (o *AlertConfigurationTemplate) SetThresholdReadOnly(v bool)`

SetThresholdReadOnly sets ThresholdReadOnly field to given value.

### HasThresholdReadOnly

`func (o *AlertConfigurationTemplate) HasThresholdReadOnly() bool`

HasThresholdReadOnly returns a boolean if a field has been set.

### GetThresholdUnit

`func (o *AlertConfigurationTemplate) GetThresholdUnit() string`

GetThresholdUnit returns the ThresholdUnit field if non-nil, zero value otherwise.

### GetThresholdUnitOk

`func (o *AlertConfigurationTemplate) GetThresholdUnitOk() (*string, bool)`

GetThresholdUnitOk returns a tuple with the ThresholdUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdUnit

`func (o *AlertConfigurationTemplate) SetThresholdUnit(v string)`

SetThresholdUnit sets ThresholdUnit field to given value.


### GetThresholdUnitName

`func (o *AlertConfigurationTemplate) GetThresholdUnitName() string`

GetThresholdUnitName returns the ThresholdUnitName field if non-nil, zero value otherwise.

### GetThresholdUnitNameOk

`func (o *AlertConfigurationTemplate) GetThresholdUnitNameOk() (*string, bool)`

GetThresholdUnitNameOk returns a tuple with the ThresholdUnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdUnitName

`func (o *AlertConfigurationTemplate) SetThresholdUnitName(v string)`

SetThresholdUnitName sets ThresholdUnitName field to given value.

### HasThresholdUnitName

`func (o *AlertConfigurationTemplate) HasThresholdUnitName() bool`

HasThresholdUnitName returns a boolean if a field has been set.

### GetThresholds

`func (o *AlertConfigurationTemplate) GetThresholds() map[string]AlertConfigurationThreshold`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *AlertConfigurationTemplate) GetThresholdsOk() (*map[string]AlertConfigurationThreshold, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *AlertConfigurationTemplate) SetThresholds(v map[string]AlertConfigurationThreshold)`

SetThresholds sets Thresholds field to given value.


### GetUuid

`func (o *AlertConfigurationTemplate) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AlertConfigurationTemplate) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AlertConfigurationTemplate) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AlertConfigurationTemplate) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


