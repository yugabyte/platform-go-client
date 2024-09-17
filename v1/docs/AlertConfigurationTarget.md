# AlertConfigurationTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Alert applicable to all targets | [optional] 
**Uuids** | Pointer to **[]string** | Alert target UUIDs | [optional] 

## Methods

### NewAlertConfigurationTarget

`func NewAlertConfigurationTarget() *AlertConfigurationTarget`

NewAlertConfigurationTarget instantiates a new AlertConfigurationTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertConfigurationTargetWithDefaults

`func NewAlertConfigurationTargetWithDefaults() *AlertConfigurationTarget`

NewAlertConfigurationTargetWithDefaults instantiates a new AlertConfigurationTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *AlertConfigurationTarget) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *AlertConfigurationTarget) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *AlertConfigurationTarget) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *AlertConfigurationTarget) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetUuids

`func (o *AlertConfigurationTarget) GetUuids() []string`

GetUuids returns the Uuids field if non-nil, zero value otherwise.

### GetUuidsOk

`func (o *AlertConfigurationTarget) GetUuidsOk() (*[]string, bool)`

GetUuidsOk returns a tuple with the Uuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuids

`func (o *AlertConfigurationTarget) SetUuids(v []string)`

SetUuids sets Uuids field to given value.

### HasUuids

`func (o *AlertConfigurationTarget) HasUuids() bool`

HasUuids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


