# UserCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**User**](User.md) |  | 
**ApiToken** | Pointer to **string** | API token of the new user, to be sent in the X-AUTH-YW-API-TOKEN header. Only set when generate_api_token was true. | [optional] [readonly] 
**ApiTokenVersion** | Pointer to **int64** | Version of that API token. Only set when generate_api_token was true. | [optional] [readonly] 

## Methods

### NewUserCreation

`func NewUserCreation(user User, ) *UserCreation`

NewUserCreation instantiates a new UserCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreationWithDefaults

`func NewUserCreationWithDefaults() *UserCreation`

NewUserCreationWithDefaults instantiates a new UserCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UserCreation) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserCreation) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserCreation) SetUser(v User)`

SetUser sets User field to given value.


### GetApiToken

`func (o *UserCreation) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *UserCreation) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *UserCreation) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.

### HasApiToken

`func (o *UserCreation) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.

### GetApiTokenVersion

`func (o *UserCreation) GetApiTokenVersion() int64`

GetApiTokenVersion returns the ApiTokenVersion field if non-nil, zero value otherwise.

### GetApiTokenVersionOk

`func (o *UserCreation) GetApiTokenVersionOk() (*int64, bool)`

GetApiTokenVersionOk returns a tuple with the ApiTokenVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTokenVersion

`func (o *UserCreation) SetApiTokenVersion(v int64)`

SetApiTokenVersion sets ApiTokenVersion field to given value.

### HasApiTokenVersion

`func (o *UserCreation) HasApiTokenVersion() bool`

HasApiTokenVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


