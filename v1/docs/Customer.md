# Customer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Customer code | 
**CreationDate** | Pointer to **time.Time** | Creation time | [optional] [readonly] 
**CustomerId** | Pointer to **int64** | Customer ID | [optional] [readonly] 
**Name** | **string** | Name of customer | 
**Uuid** | Pointer to **string** | Customer UUID | [optional] [readonly] 

## Methods

### NewCustomer

`func NewCustomer(code string, name string, ) *Customer`

NewCustomer instantiates a new Customer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerWithDefaults

`func NewCustomerWithDefaults() *Customer`

NewCustomerWithDefaults instantiates a new Customer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Customer) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Customer) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Customer) SetCode(v string)`

SetCode sets Code field to given value.


### GetCreationDate

`func (o *Customer) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Customer) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Customer) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Customer) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetCustomerId

`func (o *Customer) GetCustomerId() int64`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Customer) GetCustomerIdOk() (*int64, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Customer) SetCustomerId(v int64)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Customer) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetName

`func (o *Customer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Customer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Customer) SetName(v string)`

SetName sets Name field to given value.


### GetUuid

`func (o *Customer) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Customer) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Customer) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Customer) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


