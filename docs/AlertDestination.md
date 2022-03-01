# AlertDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | **[]string** |  | [readonly] 
**CustomerUUID** | **string** |  | 
**DefaultDestination** | **bool** |  | 
**Name** | **string** |  | 
**Uuid** | **string** |  | 

## Methods

### NewAlertDestination

`func NewAlertDestination(channels []string, customerUUID string, defaultDestination bool, name string, uuid string, ) *AlertDestination`

NewAlertDestination instantiates a new AlertDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertDestinationWithDefaults

`func NewAlertDestinationWithDefaults() *AlertDestination`

NewAlertDestinationWithDefaults instantiates a new AlertDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *AlertDestination) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *AlertDestination) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *AlertDestination) SetChannels(v []string)`

SetChannels sets Channels field to given value.


### GetCustomerUUID

`func (o *AlertDestination) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *AlertDestination) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *AlertDestination) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetDefaultDestination

`func (o *AlertDestination) GetDefaultDestination() bool`

GetDefaultDestination returns the DefaultDestination field if non-nil, zero value otherwise.

### GetDefaultDestinationOk

`func (o *AlertDestination) GetDefaultDestinationOk() (*bool, bool)`

GetDefaultDestinationOk returns a tuple with the DefaultDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestination

`func (o *AlertDestination) SetDefaultDestination(v bool)`

SetDefaultDestination sets DefaultDestination field to given value.


### GetName

`func (o *AlertDestination) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertDestination) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertDestination) SetName(v string)`

SetName sets Name field to given value.


### GetUuid

`func (o *AlertDestination) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AlertDestination) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AlertDestination) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


