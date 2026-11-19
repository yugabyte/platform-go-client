# UserCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**UserSpec**](UserSpec.md) |  | 
**Password** | **string** | Password of the user, which must satisfy the customer&#39;s password policy. Ignored when OIDC SSO is enabled. | 
**GenerateApiToken** | Pointer to **bool** | Also generate an API token for the new user and return it, so that the caller doesn&#39;t have to log in as that user to get one. Refused when OIDC SSO is enabled, because such users log in only through OIDC. | [optional] [default to false]

## Methods

### NewUserCreateSpec

`func NewUserCreateSpec(spec UserSpec, password string, ) *UserCreateSpec`

NewUserCreateSpec instantiates a new UserCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateSpecWithDefaults

`func NewUserCreateSpecWithDefaults() *UserCreateSpec`

NewUserCreateSpecWithDefaults instantiates a new UserCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *UserCreateSpec) GetSpec() UserSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *UserCreateSpec) GetSpecOk() (*UserSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *UserCreateSpec) SetSpec(v UserSpec)`

SetSpec sets Spec field to given value.


### GetPassword

`func (o *UserCreateSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserCreateSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserCreateSpec) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetGenerateApiToken

`func (o *UserCreateSpec) GetGenerateApiToken() bool`

GetGenerateApiToken returns the GenerateApiToken field if non-nil, zero value otherwise.

### GetGenerateApiTokenOk

`func (o *UserCreateSpec) GetGenerateApiTokenOk() (*bool, bool)`

GetGenerateApiTokenOk returns a tuple with the GenerateApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateApiToken

`func (o *UserCreateSpec) SetGenerateApiToken(v bool)`

SetGenerateApiToken sets GenerateApiToken field to given value.

### HasGenerateApiToken

`func (o *UserCreateSpec) HasGenerateApiToken() bool`

HasGenerateApiToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


