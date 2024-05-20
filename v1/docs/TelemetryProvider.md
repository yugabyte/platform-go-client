# TelemetryProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**TelemetryProviderConfig**](TelemetryProviderConfig.md) |  | 
**CreateTime** | Pointer to **time.Time** | Creation timestamp | [optional] [readonly] 
**CustomerUUID** | **string** | Customer UUID | [readonly] 
**Name** | **string** | Name | 
**Tags** | **map[string]string** | Extra Tags | 
**UpdateTime** | Pointer to **time.Time** | Updation timestamp | [optional] [readonly] 
**Uuid** | **string** | Telemetry Provider UUID | [readonly] 

## Methods

### NewTelemetryProvider

`func NewTelemetryProvider(config TelemetryProviderConfig, customerUUID string, name string, tags map[string]string, uuid string, ) *TelemetryProvider`

NewTelemetryProvider instantiates a new TelemetryProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryProviderWithDefaults

`func NewTelemetryProviderWithDefaults() *TelemetryProvider`

NewTelemetryProviderWithDefaults instantiates a new TelemetryProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *TelemetryProvider) GetConfig() TelemetryProviderConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *TelemetryProvider) GetConfigOk() (*TelemetryProviderConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *TelemetryProvider) SetConfig(v TelemetryProviderConfig)`

SetConfig sets Config field to given value.


### GetCreateTime

`func (o *TelemetryProvider) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *TelemetryProvider) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *TelemetryProvider) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *TelemetryProvider) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCustomerUUID

`func (o *TelemetryProvider) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *TelemetryProvider) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *TelemetryProvider) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetName

`func (o *TelemetryProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryProvider) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *TelemetryProvider) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TelemetryProvider) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TelemetryProvider) SetTags(v map[string]string)`

SetTags sets Tags field to given value.


### GetUpdateTime

`func (o *TelemetryProvider) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *TelemetryProvider) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *TelemetryProvider) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *TelemetryProvider) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetUuid

`func (o *TelemetryProvider) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *TelemetryProvider) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *TelemetryProvider) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


