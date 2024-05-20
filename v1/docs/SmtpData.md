# SmtpData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailFrom** | Pointer to **string** | SMTP email &#39;from&#39; address | [optional] 
**SmtpPassword** | Pointer to **string** | SMTP password | [optional] 
**SmtpPort** | Pointer to **int32** | SMTP port number | [optional] 
**SmtpServer** | Pointer to **string** | SMTP server | [optional] 
**SmtpUsername** | Pointer to **string** | SMTP email username | [optional] 
**UseSSL** | Pointer to **bool** | Connect to SMTP server using SSL | [optional] 
**UseTLS** | Pointer to **bool** | Connect to SMTP server using TLS | [optional] 

## Methods

### NewSmtpData

`func NewSmtpData() *SmtpData`

NewSmtpData instantiates a new SmtpData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpDataWithDefaults

`func NewSmtpDataWithDefaults() *SmtpData`

NewSmtpDataWithDefaults instantiates a new SmtpData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailFrom

`func (o *SmtpData) GetEmailFrom() string`

GetEmailFrom returns the EmailFrom field if non-nil, zero value otherwise.

### GetEmailFromOk

`func (o *SmtpData) GetEmailFromOk() (*string, bool)`

GetEmailFromOk returns a tuple with the EmailFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFrom

`func (o *SmtpData) SetEmailFrom(v string)`

SetEmailFrom sets EmailFrom field to given value.

### HasEmailFrom

`func (o *SmtpData) HasEmailFrom() bool`

HasEmailFrom returns a boolean if a field has been set.

### GetSmtpPassword

`func (o *SmtpData) GetSmtpPassword() string`

GetSmtpPassword returns the SmtpPassword field if non-nil, zero value otherwise.

### GetSmtpPasswordOk

`func (o *SmtpData) GetSmtpPasswordOk() (*string, bool)`

GetSmtpPasswordOk returns a tuple with the SmtpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPassword

`func (o *SmtpData) SetSmtpPassword(v string)`

SetSmtpPassword sets SmtpPassword field to given value.

### HasSmtpPassword

`func (o *SmtpData) HasSmtpPassword() bool`

HasSmtpPassword returns a boolean if a field has been set.

### GetSmtpPort

`func (o *SmtpData) GetSmtpPort() int32`

GetSmtpPort returns the SmtpPort field if non-nil, zero value otherwise.

### GetSmtpPortOk

`func (o *SmtpData) GetSmtpPortOk() (*int32, bool)`

GetSmtpPortOk returns a tuple with the SmtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPort

`func (o *SmtpData) SetSmtpPort(v int32)`

SetSmtpPort sets SmtpPort field to given value.

### HasSmtpPort

`func (o *SmtpData) HasSmtpPort() bool`

HasSmtpPort returns a boolean if a field has been set.

### GetSmtpServer

`func (o *SmtpData) GetSmtpServer() string`

GetSmtpServer returns the SmtpServer field if non-nil, zero value otherwise.

### GetSmtpServerOk

`func (o *SmtpData) GetSmtpServerOk() (*string, bool)`

GetSmtpServerOk returns a tuple with the SmtpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServer

`func (o *SmtpData) SetSmtpServer(v string)`

SetSmtpServer sets SmtpServer field to given value.

### HasSmtpServer

`func (o *SmtpData) HasSmtpServer() bool`

HasSmtpServer returns a boolean if a field has been set.

### GetSmtpUsername

`func (o *SmtpData) GetSmtpUsername() string`

GetSmtpUsername returns the SmtpUsername field if non-nil, zero value otherwise.

### GetSmtpUsernameOk

`func (o *SmtpData) GetSmtpUsernameOk() (*string, bool)`

GetSmtpUsernameOk returns a tuple with the SmtpUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpUsername

`func (o *SmtpData) SetSmtpUsername(v string)`

SetSmtpUsername sets SmtpUsername field to given value.

### HasSmtpUsername

`func (o *SmtpData) HasSmtpUsername() bool`

HasSmtpUsername returns a boolean if a field has been set.

### GetUseSSL

`func (o *SmtpData) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *SmtpData) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *SmtpData) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *SmtpData) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### GetUseTLS

`func (o *SmtpData) GetUseTLS() bool`

GetUseTLS returns the UseTLS field if non-nil, zero value otherwise.

### GetUseTLSOk

`func (o *SmtpData) GetUseTLSOk() (*bool, bool)`

GetUseTLSOk returns a tuple with the UseTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTLS

`func (o *SmtpData) SetUseTLS(v bool)`

SetUseTLS sets UseTLS field to given value.

### HasUseTLS

`func (o *SmtpData) HasUseTLS() bool`

HasUseTLS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


