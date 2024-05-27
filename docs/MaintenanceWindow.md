# MaintenanceWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertConfigurationFilter** | [**AlertConfigurationApiFilter**](AlertConfigurationApiFilter.md) |  | 
**CreateTime** | **time.Time** | Creation time | [readonly] 
**CustomerUUID** | **string** | Customer UUID | [readonly] 
**Description** | **string** | Description | 
**EndTime** | **time.Time** | End time | 
**Name** | **string** | Name | 
**StartTime** | **time.Time** | Start time | 
**State** | Pointer to **string** | State | [optional] [readonly] 
**SuppressHealthCheckNotificationsConfig** | Pointer to [**SuppressHealthCheckNotificationsConfig**](SuppressHealthCheckNotificationsConfig.md) |  | [optional] 
**Uuid** | Pointer to **string** | Maintenance window UUID | [optional] [readonly] 

## Methods

### NewMaintenanceWindow

`func NewMaintenanceWindow(alertConfigurationFilter AlertConfigurationApiFilter, createTime time.Time, customerUUID string, description string, endTime time.Time, name string, startTime time.Time, ) *MaintenanceWindow`

NewMaintenanceWindow instantiates a new MaintenanceWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceWindowWithDefaults

`func NewMaintenanceWindowWithDefaults() *MaintenanceWindow`

NewMaintenanceWindowWithDefaults instantiates a new MaintenanceWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertConfigurationFilter

`func (o *MaintenanceWindow) GetAlertConfigurationFilter() AlertConfigurationApiFilter`

GetAlertConfigurationFilter returns the AlertConfigurationFilter field if non-nil, zero value otherwise.

### GetAlertConfigurationFilterOk

`func (o *MaintenanceWindow) GetAlertConfigurationFilterOk() (*AlertConfigurationApiFilter, bool)`

GetAlertConfigurationFilterOk returns a tuple with the AlertConfigurationFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertConfigurationFilter

`func (o *MaintenanceWindow) SetAlertConfigurationFilter(v AlertConfigurationApiFilter)`

SetAlertConfigurationFilter sets AlertConfigurationFilter field to given value.


### GetCreateTime

`func (o *MaintenanceWindow) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *MaintenanceWindow) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *MaintenanceWindow) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.


### GetCustomerUUID

`func (o *MaintenanceWindow) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *MaintenanceWindow) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *MaintenanceWindow) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetDescription

`func (o *MaintenanceWindow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MaintenanceWindow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MaintenanceWindow) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEndTime

`func (o *MaintenanceWindow) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MaintenanceWindow) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MaintenanceWindow) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetName

`func (o *MaintenanceWindow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MaintenanceWindow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MaintenanceWindow) SetName(v string)`

SetName sets Name field to given value.


### GetStartTime

`func (o *MaintenanceWindow) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MaintenanceWindow) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MaintenanceWindow) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetState

`func (o *MaintenanceWindow) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MaintenanceWindow) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MaintenanceWindow) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MaintenanceWindow) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSuppressHealthCheckNotificationsConfig

`func (o *MaintenanceWindow) GetSuppressHealthCheckNotificationsConfig() SuppressHealthCheckNotificationsConfig`

GetSuppressHealthCheckNotificationsConfig returns the SuppressHealthCheckNotificationsConfig field if non-nil, zero value otherwise.

### GetSuppressHealthCheckNotificationsConfigOk

`func (o *MaintenanceWindow) GetSuppressHealthCheckNotificationsConfigOk() (*SuppressHealthCheckNotificationsConfig, bool)`

GetSuppressHealthCheckNotificationsConfigOk returns a tuple with the SuppressHealthCheckNotificationsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressHealthCheckNotificationsConfig

`func (o *MaintenanceWindow) SetSuppressHealthCheckNotificationsConfig(v SuppressHealthCheckNotificationsConfig)`

SetSuppressHealthCheckNotificationsConfig sets SuppressHealthCheckNotificationsConfig field to given value.

### HasSuppressHealthCheckNotificationsConfig

`func (o *MaintenanceWindow) HasSuppressHealthCheckNotificationsConfig() bool`

HasSuppressHealthCheckNotificationsConfig returns a boolean if a field has been set.

### GetUuid

`func (o *MaintenanceWindow) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MaintenanceWindow) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MaintenanceWindow) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *MaintenanceWindow) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


