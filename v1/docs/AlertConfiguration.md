# AlertConfiguration

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
**Labels** | Pointer to **map[string]string** | Labels | [optional] 
**MaintenanceWindowUuids** | Pointer to **[]string** | Maintenance window UUIDs, applied to this alert config | [optional] [readonly] 
**Name** | **string** | Name | 
**Target** | [**AlertConfigurationTarget**](AlertConfigurationTarget.md) |  | 
**TargetType** | **string** | Target type | 
**Template** | **string** | Template name | 
**ThresholdUnit** | **string** | Threshold unit | 
**Thresholds** | [**map[string]AlertConfigurationThreshold**](AlertConfigurationThreshold.md) | Thresholds | 
**Uuid** | Pointer to **string** | Configuration UUID | [optional] [readonly] 

## Methods

### NewAlertConfiguration

`func NewAlertConfiguration(active bool, alertCount float64, createTime time.Time, customerUUID string, defaultDestination bool, description string, durationSec int32, name string, target AlertConfigurationTarget, targetType string, template string, thresholdUnit string, thresholds map[string]AlertConfigurationThreshold, ) *AlertConfiguration`

NewAlertConfiguration instantiates a new AlertConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertConfigurationWithDefaults

`func NewAlertConfigurationWithDefaults() *AlertConfiguration`

NewAlertConfigurationWithDefaults instantiates a new AlertConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *AlertConfiguration) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AlertConfiguration) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AlertConfiguration) SetActive(v bool)`

SetActive sets Active field to given value.


### GetAlertCount

`func (o *AlertConfiguration) GetAlertCount() float64`

GetAlertCount returns the AlertCount field if non-nil, zero value otherwise.

### GetAlertCountOk

`func (o *AlertConfiguration) GetAlertCountOk() (*float64, bool)`

GetAlertCountOk returns a tuple with the AlertCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCount

`func (o *AlertConfiguration) SetAlertCount(v float64)`

SetAlertCount sets AlertCount field to given value.


### GetCreateTime

`func (o *AlertConfiguration) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *AlertConfiguration) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *AlertConfiguration) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.


### GetCustomerUUID

`func (o *AlertConfiguration) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *AlertConfiguration) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *AlertConfiguration) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetDefaultDestination

`func (o *AlertConfiguration) GetDefaultDestination() bool`

GetDefaultDestination returns the DefaultDestination field if non-nil, zero value otherwise.

### GetDefaultDestinationOk

`func (o *AlertConfiguration) GetDefaultDestinationOk() (*bool, bool)`

GetDefaultDestinationOk returns a tuple with the DefaultDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestination

`func (o *AlertConfiguration) SetDefaultDestination(v bool)`

SetDefaultDestination sets DefaultDestination field to given value.


### GetDescription

`func (o *AlertConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertConfiguration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertConfiguration) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDestinationUUID

`func (o *AlertConfiguration) GetDestinationUUID() string`

GetDestinationUUID returns the DestinationUUID field if non-nil, zero value otherwise.

### GetDestinationUUIDOk

`func (o *AlertConfiguration) GetDestinationUUIDOk() (*string, bool)`

GetDestinationUUIDOk returns a tuple with the DestinationUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationUUID

`func (o *AlertConfiguration) SetDestinationUUID(v string)`

SetDestinationUUID sets DestinationUUID field to given value.

### HasDestinationUUID

`func (o *AlertConfiguration) HasDestinationUUID() bool`

HasDestinationUUID returns a boolean if a field has been set.

### GetDurationSec

`func (o *AlertConfiguration) GetDurationSec() int32`

GetDurationSec returns the DurationSec field if non-nil, zero value otherwise.

### GetDurationSecOk

`func (o *AlertConfiguration) GetDurationSecOk() (*int32, bool)`

GetDurationSecOk returns a tuple with the DurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSec

`func (o *AlertConfiguration) SetDurationSec(v int32)`

SetDurationSec sets DurationSec field to given value.


### GetLabels

`func (o *AlertConfiguration) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AlertConfiguration) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AlertConfiguration) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AlertConfiguration) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMaintenanceWindowUuids

`func (o *AlertConfiguration) GetMaintenanceWindowUuids() []string`

GetMaintenanceWindowUuids returns the MaintenanceWindowUuids field if non-nil, zero value otherwise.

### GetMaintenanceWindowUuidsOk

`func (o *AlertConfiguration) GetMaintenanceWindowUuidsOk() (*[]string, bool)`

GetMaintenanceWindowUuidsOk returns a tuple with the MaintenanceWindowUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceWindowUuids

`func (o *AlertConfiguration) SetMaintenanceWindowUuids(v []string)`

SetMaintenanceWindowUuids sets MaintenanceWindowUuids field to given value.

### HasMaintenanceWindowUuids

`func (o *AlertConfiguration) HasMaintenanceWindowUuids() bool`

HasMaintenanceWindowUuids returns a boolean if a field has been set.

### GetName

`func (o *AlertConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertConfiguration) SetName(v string)`

SetName sets Name field to given value.


### GetTarget

`func (o *AlertConfiguration) GetTarget() AlertConfigurationTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AlertConfiguration) GetTargetOk() (*AlertConfigurationTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AlertConfiguration) SetTarget(v AlertConfigurationTarget)`

SetTarget sets Target field to given value.


### GetTargetType

`func (o *AlertConfiguration) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *AlertConfiguration) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *AlertConfiguration) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.


### GetTemplate

`func (o *AlertConfiguration) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *AlertConfiguration) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *AlertConfiguration) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetThresholdUnit

`func (o *AlertConfiguration) GetThresholdUnit() string`

GetThresholdUnit returns the ThresholdUnit field if non-nil, zero value otherwise.

### GetThresholdUnitOk

`func (o *AlertConfiguration) GetThresholdUnitOk() (*string, bool)`

GetThresholdUnitOk returns a tuple with the ThresholdUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdUnit

`func (o *AlertConfiguration) SetThresholdUnit(v string)`

SetThresholdUnit sets ThresholdUnit field to given value.


### GetThresholds

`func (o *AlertConfiguration) GetThresholds() map[string]AlertConfigurationThreshold`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *AlertConfiguration) GetThresholdsOk() (*map[string]AlertConfigurationThreshold, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *AlertConfiguration) SetThresholds(v map[string]AlertConfigurationThreshold)`

SetThresholds sets Thresholds field to given value.


### GetUuid

`func (o *AlertConfiguration) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AlertConfiguration) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AlertConfiguration) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AlertConfiguration) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


