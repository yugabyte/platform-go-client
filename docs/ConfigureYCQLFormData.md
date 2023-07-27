# ConfigureYCQLFormData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommunicationPorts** | Pointer to [**CommunicationPorts**](CommunicationPorts.md) |  | [optional] 
**EnableYCQL** | Pointer to **bool** | Enable YCQL Api for the universe | [optional] 
**EnableYCQLAuth** | Pointer to **bool** | Enable YCQL Auth for the universe | [optional] 
**YcqlPassword** | Pointer to **string** | YCQL Auth password | [optional] 

## Methods

### NewConfigureYCQLFormData

`func NewConfigureYCQLFormData() *ConfigureYCQLFormData`

NewConfigureYCQLFormData instantiates a new ConfigureYCQLFormData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigureYCQLFormDataWithDefaults

`func NewConfigureYCQLFormDataWithDefaults() *ConfigureYCQLFormData`

NewConfigureYCQLFormDataWithDefaults instantiates a new ConfigureYCQLFormData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommunicationPorts

`func (o *ConfigureYCQLFormData) GetCommunicationPorts() CommunicationPorts`

GetCommunicationPorts returns the CommunicationPorts field if non-nil, zero value otherwise.

### GetCommunicationPortsOk

`func (o *ConfigureYCQLFormData) GetCommunicationPortsOk() (*CommunicationPorts, bool)`

GetCommunicationPortsOk returns a tuple with the CommunicationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationPorts

`func (o *ConfigureYCQLFormData) SetCommunicationPorts(v CommunicationPorts)`

SetCommunicationPorts sets CommunicationPorts field to given value.

### HasCommunicationPorts

`func (o *ConfigureYCQLFormData) HasCommunicationPorts() bool`

HasCommunicationPorts returns a boolean if a field has been set.

### GetEnableYCQL

`func (o *ConfigureYCQLFormData) GetEnableYCQL() bool`

GetEnableYCQL returns the EnableYCQL field if non-nil, zero value otherwise.

### GetEnableYCQLOk

`func (o *ConfigureYCQLFormData) GetEnableYCQLOk() (*bool, bool)`

GetEnableYCQLOk returns a tuple with the EnableYCQL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableYCQL

`func (o *ConfigureYCQLFormData) SetEnableYCQL(v bool)`

SetEnableYCQL sets EnableYCQL field to given value.

### HasEnableYCQL

`func (o *ConfigureYCQLFormData) HasEnableYCQL() bool`

HasEnableYCQL returns a boolean if a field has been set.

### GetEnableYCQLAuth

`func (o *ConfigureYCQLFormData) GetEnableYCQLAuth() bool`

GetEnableYCQLAuth returns the EnableYCQLAuth field if non-nil, zero value otherwise.

### GetEnableYCQLAuthOk

`func (o *ConfigureYCQLFormData) GetEnableYCQLAuthOk() (*bool, bool)`

GetEnableYCQLAuthOk returns a tuple with the EnableYCQLAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableYCQLAuth

`func (o *ConfigureYCQLFormData) SetEnableYCQLAuth(v bool)`

SetEnableYCQLAuth sets EnableYCQLAuth field to given value.

### HasEnableYCQLAuth

`func (o *ConfigureYCQLFormData) HasEnableYCQLAuth() bool`

HasEnableYCQLAuth returns a boolean if a field has been set.

### GetYcqlPassword

`func (o *ConfigureYCQLFormData) GetYcqlPassword() string`

GetYcqlPassword returns the YcqlPassword field if non-nil, zero value otherwise.

### GetYcqlPasswordOk

`func (o *ConfigureYCQLFormData) GetYcqlPasswordOk() (*string, bool)`

GetYcqlPasswordOk returns a tuple with the YcqlPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYcqlPassword

`func (o *ConfigureYCQLFormData) SetYcqlPassword(v string)`

SetYcqlPassword sets YcqlPassword field to given value.

### HasYcqlPassword

`func (o *ConfigureYCQLFormData) HasYcqlPassword() bool`

HasYcqlPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


