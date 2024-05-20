# ConfigureYSQLFormData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommunicationPorts** | Pointer to [**CommunicationPorts**](CommunicationPorts.md) |  | [optional] 
**EnableYSQL** | Pointer to **bool** | Enable YSQL Api for the universe | [optional] 
**EnableYSQLAuth** | Pointer to **bool** | Enable YSQL Auth for the universe | [optional] 
**YsqlPassword** | Pointer to **string** | YSQL Auth password | [optional] 

## Methods

### NewConfigureYSQLFormData

`func NewConfigureYSQLFormData() *ConfigureYSQLFormData`

NewConfigureYSQLFormData instantiates a new ConfigureYSQLFormData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigureYSQLFormDataWithDefaults

`func NewConfigureYSQLFormDataWithDefaults() *ConfigureYSQLFormData`

NewConfigureYSQLFormDataWithDefaults instantiates a new ConfigureYSQLFormData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommunicationPorts

`func (o *ConfigureYSQLFormData) GetCommunicationPorts() CommunicationPorts`

GetCommunicationPorts returns the CommunicationPorts field if non-nil, zero value otherwise.

### GetCommunicationPortsOk

`func (o *ConfigureYSQLFormData) GetCommunicationPortsOk() (*CommunicationPorts, bool)`

GetCommunicationPortsOk returns a tuple with the CommunicationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationPorts

`func (o *ConfigureYSQLFormData) SetCommunicationPorts(v CommunicationPorts)`

SetCommunicationPorts sets CommunicationPorts field to given value.

### HasCommunicationPorts

`func (o *ConfigureYSQLFormData) HasCommunicationPorts() bool`

HasCommunicationPorts returns a boolean if a field has been set.

### GetEnableYSQL

`func (o *ConfigureYSQLFormData) GetEnableYSQL() bool`

GetEnableYSQL returns the EnableYSQL field if non-nil, zero value otherwise.

### GetEnableYSQLOk

`func (o *ConfigureYSQLFormData) GetEnableYSQLOk() (*bool, bool)`

GetEnableYSQLOk returns a tuple with the EnableYSQL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableYSQL

`func (o *ConfigureYSQLFormData) SetEnableYSQL(v bool)`

SetEnableYSQL sets EnableYSQL field to given value.

### HasEnableYSQL

`func (o *ConfigureYSQLFormData) HasEnableYSQL() bool`

HasEnableYSQL returns a boolean if a field has been set.

### GetEnableYSQLAuth

`func (o *ConfigureYSQLFormData) GetEnableYSQLAuth() bool`

GetEnableYSQLAuth returns the EnableYSQLAuth field if non-nil, zero value otherwise.

### GetEnableYSQLAuthOk

`func (o *ConfigureYSQLFormData) GetEnableYSQLAuthOk() (*bool, bool)`

GetEnableYSQLAuthOk returns a tuple with the EnableYSQLAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableYSQLAuth

`func (o *ConfigureYSQLFormData) SetEnableYSQLAuth(v bool)`

SetEnableYSQLAuth sets EnableYSQLAuth field to given value.

### HasEnableYSQLAuth

`func (o *ConfigureYSQLFormData) HasEnableYSQLAuth() bool`

HasEnableYSQLAuth returns a boolean if a field has been set.

### GetYsqlPassword

`func (o *ConfigureYSQLFormData) GetYsqlPassword() string`

GetYsqlPassword returns the YsqlPassword field if non-nil, zero value otherwise.

### GetYsqlPasswordOk

`func (o *ConfigureYSQLFormData) GetYsqlPasswordOk() (*string, bool)`

GetYsqlPasswordOk returns a tuple with the YsqlPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlPassword

`func (o *ConfigureYSQLFormData) SetYsqlPassword(v string)`

SetYsqlPassword sets YsqlPassword field to given value.

### HasYsqlPassword

`func (o *ConfigureYSQLFormData) HasYsqlPassword() bool`

HasYsqlPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


