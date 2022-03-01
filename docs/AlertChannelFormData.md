# AlertChannelFormData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertChannelUUID** | **string** |  | 
**Name** | **string** |  | 
**Params** | [**AlertChannelParams**](AlertChannelParams.md) |  | 

## Methods

### NewAlertChannelFormData

`func NewAlertChannelFormData(alertChannelUUID string, name string, params AlertChannelParams, ) *AlertChannelFormData`

NewAlertChannelFormData instantiates a new AlertChannelFormData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertChannelFormDataWithDefaults

`func NewAlertChannelFormDataWithDefaults() *AlertChannelFormData`

NewAlertChannelFormDataWithDefaults instantiates a new AlertChannelFormData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertChannelUUID

`func (o *AlertChannelFormData) GetAlertChannelUUID() string`

GetAlertChannelUUID returns the AlertChannelUUID field if non-nil, zero value otherwise.

### GetAlertChannelUUIDOk

`func (o *AlertChannelFormData) GetAlertChannelUUIDOk() (*string, bool)`

GetAlertChannelUUIDOk returns a tuple with the AlertChannelUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertChannelUUID

`func (o *AlertChannelFormData) SetAlertChannelUUID(v string)`

SetAlertChannelUUID sets AlertChannelUUID field to given value.


### GetName

`func (o *AlertChannelFormData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertChannelFormData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertChannelFormData) SetName(v string)`

SetName sets Name field to given value.


### GetParams

`func (o *AlertChannelFormData) GetParams() AlertChannelParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *AlertChannelFormData) GetParamsOk() (*AlertChannelParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *AlertChannelFormData) SetParams(v AlertChannelParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


