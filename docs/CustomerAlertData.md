# CustomerAlertData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertingData** | [**AlertingData**](AlertingData.md) |  | 
**CallhomeLevel** | **string** |  | 
**Code** | Pointer to **string** | Alert code | [optional] 
**ConfirmPassword** | Pointer to **string** | Email password confirmation | [optional] 
**Email** | Pointer to **string** | Alert email address | [optional] 
**Features** | Pointer to **map[string]map[string]interface{}** | Features | [optional] 
**Name** | Pointer to **string** | Alert name | [optional] 
**Password** | Pointer to **string** | Email password | [optional] 
**SmtpData** | [**SmtpData**](SmtpData.md) |  | 

## Methods

### NewCustomerAlertData

`func NewCustomerAlertData(alertingData AlertingData, callhomeLevel string, smtpData SmtpData, ) *CustomerAlertData`

NewCustomerAlertData instantiates a new CustomerAlertData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAlertDataWithDefaults

`func NewCustomerAlertDataWithDefaults() *CustomerAlertData`

NewCustomerAlertDataWithDefaults instantiates a new CustomerAlertData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertingData

`func (o *CustomerAlertData) GetAlertingData() AlertingData`

GetAlertingData returns the AlertingData field if non-nil, zero value otherwise.

### GetAlertingDataOk

`func (o *CustomerAlertData) GetAlertingDataOk() (*AlertingData, bool)`

GetAlertingDataOk returns a tuple with the AlertingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingData

`func (o *CustomerAlertData) SetAlertingData(v AlertingData)`

SetAlertingData sets AlertingData field to given value.


### GetCallhomeLevel

`func (o *CustomerAlertData) GetCallhomeLevel() string`

GetCallhomeLevel returns the CallhomeLevel field if non-nil, zero value otherwise.

### GetCallhomeLevelOk

`func (o *CustomerAlertData) GetCallhomeLevelOk() (*string, bool)`

GetCallhomeLevelOk returns a tuple with the CallhomeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallhomeLevel

`func (o *CustomerAlertData) SetCallhomeLevel(v string)`

SetCallhomeLevel sets CallhomeLevel field to given value.


### GetCode

`func (o *CustomerAlertData) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CustomerAlertData) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CustomerAlertData) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CustomerAlertData) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetConfirmPassword

`func (o *CustomerAlertData) GetConfirmPassword() string`

GetConfirmPassword returns the ConfirmPassword field if non-nil, zero value otherwise.

### GetConfirmPasswordOk

`func (o *CustomerAlertData) GetConfirmPasswordOk() (*string, bool)`

GetConfirmPasswordOk returns a tuple with the ConfirmPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmPassword

`func (o *CustomerAlertData) SetConfirmPassword(v string)`

SetConfirmPassword sets ConfirmPassword field to given value.

### HasConfirmPassword

`func (o *CustomerAlertData) HasConfirmPassword() bool`

HasConfirmPassword returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerAlertData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerAlertData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerAlertData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerAlertData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFeatures

`func (o *CustomerAlertData) GetFeatures() map[string]map[string]interface{}`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CustomerAlertData) GetFeaturesOk() (*map[string]map[string]interface{}, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CustomerAlertData) SetFeatures(v map[string]map[string]interface{})`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *CustomerAlertData) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetName

`func (o *CustomerAlertData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerAlertData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerAlertData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomerAlertData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *CustomerAlertData) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CustomerAlertData) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CustomerAlertData) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CustomerAlertData) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSmtpData

`func (o *CustomerAlertData) GetSmtpData() SmtpData`

GetSmtpData returns the SmtpData field if non-nil, zero value otherwise.

### GetSmtpDataOk

`func (o *CustomerAlertData) GetSmtpDataOk() (*SmtpData, bool)`

GetSmtpDataOk returns a tuple with the SmtpData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpData

`func (o *CustomerAlertData) SetSmtpData(v SmtpData)`

SetSmtpData sets SmtpData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


