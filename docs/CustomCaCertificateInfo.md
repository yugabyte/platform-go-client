# CustomCaCertificateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** |  | 
**Contents** | Pointer to **string** | Path to CA Certificate | [optional] 
**CreatedTime** | Pointer to **time.Time** | Date when certificate was added. | [optional] 
**CustomerId** | **string** |  | 
**ExpiryDate** | Pointer to **time.Time** | End date of certificate validity. | [optional] 
**Id** | **string** |  | 
**Name** | **string** |  | 
**StartDate** | Pointer to **time.Time** | Start date of certificate validity. | [optional] 

## Methods

### NewCustomCaCertificateInfo

`func NewCustomCaCertificateInfo(active bool, customerId string, id string, name string, ) *CustomCaCertificateInfo`

NewCustomCaCertificateInfo instantiates a new CustomCaCertificateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomCaCertificateInfoWithDefaults

`func NewCustomCaCertificateInfoWithDefaults() *CustomCaCertificateInfo`

NewCustomCaCertificateInfoWithDefaults instantiates a new CustomCaCertificateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CustomCaCertificateInfo) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CustomCaCertificateInfo) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CustomCaCertificateInfo) SetActive(v bool)`

SetActive sets Active field to given value.


### GetContents

`func (o *CustomCaCertificateInfo) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *CustomCaCertificateInfo) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *CustomCaCertificateInfo) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *CustomCaCertificateInfo) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetCreatedTime

`func (o *CustomCaCertificateInfo) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CustomCaCertificateInfo) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CustomCaCertificateInfo) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CustomCaCertificateInfo) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCustomerId

`func (o *CustomCaCertificateInfo) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomCaCertificateInfo) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomCaCertificateInfo) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetExpiryDate

`func (o *CustomCaCertificateInfo) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *CustomCaCertificateInfo) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *CustomCaCertificateInfo) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *CustomCaCertificateInfo) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetId

`func (o *CustomCaCertificateInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomCaCertificateInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomCaCertificateInfo) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CustomCaCertificateInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomCaCertificateInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomCaCertificateInfo) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *CustomCaCertificateInfo) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CustomCaCertificateInfo) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CustomCaCertificateInfo) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CustomCaCertificateInfo) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


