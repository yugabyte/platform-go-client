# CustomerRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerUuid** | **string** | UUID of the registered customer | [readonly] 
**UserUuid** | **string** | UUID of the primary user | [readonly] 
**ApiToken** | **string** | API token of the primary user, to be sent in the X-AUTH-YW-API-TOKEN header | [readonly] 
**ApiTokenVersion** | **int64** | Version of the API token | [readonly] 

## Methods

### NewCustomerRegistration

`func NewCustomerRegistration(customerUuid string, userUuid string, apiToken string, apiTokenVersion int64, ) *CustomerRegistration`

NewCustomerRegistration instantiates a new CustomerRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerRegistrationWithDefaults

`func NewCustomerRegistrationWithDefaults() *CustomerRegistration`

NewCustomerRegistrationWithDefaults instantiates a new CustomerRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerUuid

`func (o *CustomerRegistration) GetCustomerUuid() string`

GetCustomerUuid returns the CustomerUuid field if non-nil, zero value otherwise.

### GetCustomerUuidOk

`func (o *CustomerRegistration) GetCustomerUuidOk() (*string, bool)`

GetCustomerUuidOk returns a tuple with the CustomerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUuid

`func (o *CustomerRegistration) SetCustomerUuid(v string)`

SetCustomerUuid sets CustomerUuid field to given value.


### GetUserUuid

`func (o *CustomerRegistration) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *CustomerRegistration) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *CustomerRegistration) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.


### GetApiToken

`func (o *CustomerRegistration) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *CustomerRegistration) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *CustomerRegistration) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### GetApiTokenVersion

`func (o *CustomerRegistration) GetApiTokenVersion() int64`

GetApiTokenVersion returns the ApiTokenVersion field if non-nil, zero value otherwise.

### GetApiTokenVersionOk

`func (o *CustomerRegistration) GetApiTokenVersionOk() (*int64, bool)`

GetApiTokenVersionOk returns a tuple with the ApiTokenVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTokenVersion

`func (o *CustomerRegistration) SetApiTokenVersion(v int64)`

SetApiTokenVersion sets ApiTokenVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


