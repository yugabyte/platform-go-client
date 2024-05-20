# AlertDestinationFormData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | **[]string** |  | 
**DefaultDestination** | **bool** |  | 
**Name** | **string** |  | 

## Methods

### NewAlertDestinationFormData

`func NewAlertDestinationFormData(channels []string, defaultDestination bool, name string, ) *AlertDestinationFormData`

NewAlertDestinationFormData instantiates a new AlertDestinationFormData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertDestinationFormDataWithDefaults

`func NewAlertDestinationFormDataWithDefaults() *AlertDestinationFormData`

NewAlertDestinationFormDataWithDefaults instantiates a new AlertDestinationFormData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *AlertDestinationFormData) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *AlertDestinationFormData) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *AlertDestinationFormData) SetChannels(v []string)`

SetChannels sets Channels field to given value.


### GetDefaultDestination

`func (o *AlertDestinationFormData) GetDefaultDestination() bool`

GetDefaultDestination returns the DefaultDestination field if non-nil, zero value otherwise.

### GetDefaultDestinationOk

`func (o *AlertDestinationFormData) GetDefaultDestinationOk() (*bool, bool)`

GetDefaultDestinationOk returns a tuple with the DefaultDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestination

`func (o *AlertDestinationFormData) SetDefaultDestination(v bool)`

SetDefaultDestination sets DefaultDestination field to given value.


### GetName

`func (o *AlertDestinationFormData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertDestinationFormData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertDestinationFormData) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


