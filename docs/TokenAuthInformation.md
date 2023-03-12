# TokenAuthInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenHeader** | Pointer to **string** | Header name | [optional] 
**TokenValue** | Pointer to **string** | Token value | [optional] 

## Methods

### NewTokenAuthInformation

`func NewTokenAuthInformation() *TokenAuthInformation`

NewTokenAuthInformation instantiates a new TokenAuthInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenAuthInformationWithDefaults

`func NewTokenAuthInformationWithDefaults() *TokenAuthInformation`

NewTokenAuthInformationWithDefaults instantiates a new TokenAuthInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenHeader

`func (o *TokenAuthInformation) GetTokenHeader() string`

GetTokenHeader returns the TokenHeader field if non-nil, zero value otherwise.

### GetTokenHeaderOk

`func (o *TokenAuthInformation) GetTokenHeaderOk() (*string, bool)`

GetTokenHeaderOk returns a tuple with the TokenHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenHeader

`func (o *TokenAuthInformation) SetTokenHeader(v string)`

SetTokenHeader sets TokenHeader field to given value.

### HasTokenHeader

`func (o *TokenAuthInformation) HasTokenHeader() bool`

HasTokenHeader returns a boolean if a field has been set.

### GetTokenValue

`func (o *TokenAuthInformation) GetTokenValue() string`

GetTokenValue returns the TokenValue field if non-nil, zero value otherwise.

### GetTokenValueOk

`func (o *TokenAuthInformation) GetTokenValueOk() (*string, bool)`

GetTokenValueOk returns a tuple with the TokenValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenValue

`func (o *TokenAuthInformation) SetTokenValue(v string)`

SetTokenValue sets TokenValue field to given value.

### HasTokenValue

`func (o *TokenAuthInformation) HasTokenValue() bool`

HasTokenValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


