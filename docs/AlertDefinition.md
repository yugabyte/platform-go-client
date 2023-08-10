# AlertDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationUUID** | **string** |  | 
**CustomerUUID** | **string** |  | 
**Labels** | [**[]AlertDefinitionLabel**](AlertDefinitionLabel.md) |  | 
**UniverseUUID** | **string** |  | 
**Uuid** | **string** |  | 
**Version** | **int32** |  | 

## Methods

### NewAlertDefinition

`func NewAlertDefinition(configurationUUID string, customerUUID string, labels []AlertDefinitionLabel, universeUUID string, uuid string, version int32, ) *AlertDefinition`

NewAlertDefinition instantiates a new AlertDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertDefinitionWithDefaults

`func NewAlertDefinitionWithDefaults() *AlertDefinition`

NewAlertDefinitionWithDefaults instantiates a new AlertDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationUUID

`func (o *AlertDefinition) GetConfigurationUUID() string`

GetConfigurationUUID returns the ConfigurationUUID field if non-nil, zero value otherwise.

### GetConfigurationUUIDOk

`func (o *AlertDefinition) GetConfigurationUUIDOk() (*string, bool)`

GetConfigurationUUIDOk returns a tuple with the ConfigurationUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationUUID

`func (o *AlertDefinition) SetConfigurationUUID(v string)`

SetConfigurationUUID sets ConfigurationUUID field to given value.


### GetCustomerUUID

`func (o *AlertDefinition) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *AlertDefinition) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *AlertDefinition) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetLabels

`func (o *AlertDefinition) GetLabels() []AlertDefinitionLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AlertDefinition) GetLabelsOk() (*[]AlertDefinitionLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AlertDefinition) SetLabels(v []AlertDefinitionLabel)`

SetLabels sets Labels field to given value.


### GetUniverseUUID

`func (o *AlertDefinition) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *AlertDefinition) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *AlertDefinition) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.


### GetUuid

`func (o *AlertDefinition) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AlertDefinition) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AlertDefinition) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetVersion

`func (o *AlertDefinition) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AlertDefinition) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AlertDefinition) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


