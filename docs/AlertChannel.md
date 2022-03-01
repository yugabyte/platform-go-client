# AlertChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerUuid** | **string** |  | 
**Name** | **string** |  | 
**Params** | [**AlertChannelParams**](AlertChannelParams.md) |  | 
**Uuid** | **string** |  | 

## Methods

### NewAlertChannel

`func NewAlertChannel(customerUuid string, name string, params AlertChannelParams, uuid string, ) *AlertChannel`

NewAlertChannel instantiates a new AlertChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertChannelWithDefaults

`func NewAlertChannelWithDefaults() *AlertChannel`

NewAlertChannelWithDefaults instantiates a new AlertChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerUuid

`func (o *AlertChannel) GetCustomerUuid() string`

GetCustomerUuid returns the CustomerUuid field if non-nil, zero value otherwise.

### GetCustomerUuidOk

`func (o *AlertChannel) GetCustomerUuidOk() (*string, bool)`

GetCustomerUuidOk returns a tuple with the CustomerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUuid

`func (o *AlertChannel) SetCustomerUuid(v string)`

SetCustomerUuid sets CustomerUuid field to given value.


### GetName

`func (o *AlertChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertChannel) SetName(v string)`

SetName sets Name field to given value.


### GetParams

`func (o *AlertChannel) GetParams() AlertChannelParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *AlertChannel) GetParamsOk() (*AlertChannelParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *AlertChannel) SetParams(v AlertChannelParams)`

SetParams sets Params field to given value.


### GetUuid

`func (o *AlertChannel) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AlertChannel) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AlertChannel) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


