# BasicAuthInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Password | [optional] 
**Username** | Pointer to **string** | Username | [optional] 

## Methods

### NewBasicAuthInformation

`func NewBasicAuthInformation() *BasicAuthInformation`

NewBasicAuthInformation instantiates a new BasicAuthInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicAuthInformationWithDefaults

`func NewBasicAuthInformationWithDefaults() *BasicAuthInformation`

NewBasicAuthInformationWithDefaults instantiates a new BasicAuthInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *BasicAuthInformation) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BasicAuthInformation) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BasicAuthInformation) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BasicAuthInformation) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *BasicAuthInformation) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BasicAuthInformation) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BasicAuthInformation) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *BasicAuthInformation) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


