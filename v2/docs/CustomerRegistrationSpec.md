# CustomerRegistrationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Customer code | 
**Name** | **string** | Customer name | 
**Email** | **string** | Email of the primary user. Must be unique on the YBA instance. | 
**Password** | **string** | Password of the primary user. Must satisfy the YBA password policy. | 

## Methods

### NewCustomerRegistrationSpec

`func NewCustomerRegistrationSpec(code string, name string, email string, password string, ) *CustomerRegistrationSpec`

NewCustomerRegistrationSpec instantiates a new CustomerRegistrationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerRegistrationSpecWithDefaults

`func NewCustomerRegistrationSpecWithDefaults() *CustomerRegistrationSpec`

NewCustomerRegistrationSpecWithDefaults instantiates a new CustomerRegistrationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CustomerRegistrationSpec) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CustomerRegistrationSpec) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CustomerRegistrationSpec) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *CustomerRegistrationSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerRegistrationSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerRegistrationSpec) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *CustomerRegistrationSpec) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerRegistrationSpec) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerRegistrationSpec) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *CustomerRegistrationSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CustomerRegistrationSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CustomerRegistrationSpec) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


