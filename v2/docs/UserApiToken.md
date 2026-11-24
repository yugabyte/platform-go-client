# UserApiToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerUuid** | **string** | UUID of the user&#39;s customer | [readonly] 
**UserUuid** | **string** | UUID of the user | [readonly] 
**ApiToken** | **string** | The new API token, to be sent in the X-AUTH-YW-API-TOKEN header | [readonly] 
**ApiTokenVersion** | **int64** | Version of the new API token | [readonly] 

## Methods

### NewUserApiToken

`func NewUserApiToken(customerUuid string, userUuid string, apiToken string, apiTokenVersion int64, ) *UserApiToken`

NewUserApiToken instantiates a new UserApiToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserApiTokenWithDefaults

`func NewUserApiTokenWithDefaults() *UserApiToken`

NewUserApiTokenWithDefaults instantiates a new UserApiToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerUuid

`func (o *UserApiToken) GetCustomerUuid() string`

GetCustomerUuid returns the CustomerUuid field if non-nil, zero value otherwise.

### GetCustomerUuidOk

`func (o *UserApiToken) GetCustomerUuidOk() (*string, bool)`

GetCustomerUuidOk returns a tuple with the CustomerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUuid

`func (o *UserApiToken) SetCustomerUuid(v string)`

SetCustomerUuid sets CustomerUuid field to given value.


### GetUserUuid

`func (o *UserApiToken) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *UserApiToken) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *UserApiToken) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.


### GetApiToken

`func (o *UserApiToken) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *UserApiToken) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *UserApiToken) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.


### GetApiTokenVersion

`func (o *UserApiToken) GetApiTokenVersion() int64`

GetApiTokenVersion returns the ApiTokenVersion field if non-nil, zero value otherwise.

### GetApiTokenVersionOk

`func (o *UserApiToken) GetApiTokenVersionOk() (*int64, bool)`

GetApiTokenVersionOk returns a tuple with the ApiTokenVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTokenVersion

`func (o *UserApiToken) SetApiTokenVersion(v int64)`

SetApiTokenVersion sets ApiTokenVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


