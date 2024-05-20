# CustomerDetailsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertingData** | Pointer to [**AlertingData**](AlertingData.md) |  | [optional] 
**CallhomeLevel** | Pointer to **string** | Call-home level | [optional] [readonly] 
**Code** | **string** | Customer code | 
**CreationDate** | Pointer to **time.Time** | Creation timestamp | [optional] [readonly] 
**CustomerId** | Pointer to **int32** | Customer ID | [optional] [readonly] 
**Name** | **string** | Customer name | 
**SmtpData** | Pointer to [**SmtpData**](SmtpData.md) |  | [optional] 
**UniverseUUIDs** | Pointer to **[]string** | Associated universe IDs | [optional] [readonly] 
**Uuid** | Pointer to **string** | User UUID | [optional] [readonly] 

## Methods

### NewCustomerDetailsData

`func NewCustomerDetailsData(code string, name string, ) *CustomerDetailsData`

NewCustomerDetailsData instantiates a new CustomerDetailsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerDetailsDataWithDefaults

`func NewCustomerDetailsDataWithDefaults() *CustomerDetailsData`

NewCustomerDetailsDataWithDefaults instantiates a new CustomerDetailsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertingData

`func (o *CustomerDetailsData) GetAlertingData() AlertingData`

GetAlertingData returns the AlertingData field if non-nil, zero value otherwise.

### GetAlertingDataOk

`func (o *CustomerDetailsData) GetAlertingDataOk() (*AlertingData, bool)`

GetAlertingDataOk returns a tuple with the AlertingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingData

`func (o *CustomerDetailsData) SetAlertingData(v AlertingData)`

SetAlertingData sets AlertingData field to given value.

### HasAlertingData

`func (o *CustomerDetailsData) HasAlertingData() bool`

HasAlertingData returns a boolean if a field has been set.

### GetCallhomeLevel

`func (o *CustomerDetailsData) GetCallhomeLevel() string`

GetCallhomeLevel returns the CallhomeLevel field if non-nil, zero value otherwise.

### GetCallhomeLevelOk

`func (o *CustomerDetailsData) GetCallhomeLevelOk() (*string, bool)`

GetCallhomeLevelOk returns a tuple with the CallhomeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallhomeLevel

`func (o *CustomerDetailsData) SetCallhomeLevel(v string)`

SetCallhomeLevel sets CallhomeLevel field to given value.

### HasCallhomeLevel

`func (o *CustomerDetailsData) HasCallhomeLevel() bool`

HasCallhomeLevel returns a boolean if a field has been set.

### GetCode

`func (o *CustomerDetailsData) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CustomerDetailsData) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CustomerDetailsData) SetCode(v string)`

SetCode sets Code field to given value.


### GetCreationDate

`func (o *CustomerDetailsData) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *CustomerDetailsData) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *CustomerDetailsData) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *CustomerDetailsData) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetCustomerId

`func (o *CustomerDetailsData) GetCustomerId() int32`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomerDetailsData) GetCustomerIdOk() (*int32, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomerDetailsData) SetCustomerId(v int32)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CustomerDetailsData) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetName

`func (o *CustomerDetailsData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerDetailsData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerDetailsData) SetName(v string)`

SetName sets Name field to given value.


### GetSmtpData

`func (o *CustomerDetailsData) GetSmtpData() SmtpData`

GetSmtpData returns the SmtpData field if non-nil, zero value otherwise.

### GetSmtpDataOk

`func (o *CustomerDetailsData) GetSmtpDataOk() (*SmtpData, bool)`

GetSmtpDataOk returns a tuple with the SmtpData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpData

`func (o *CustomerDetailsData) SetSmtpData(v SmtpData)`

SetSmtpData sets SmtpData field to given value.

### HasSmtpData

`func (o *CustomerDetailsData) HasSmtpData() bool`

HasSmtpData returns a boolean if a field has been set.

### GetUniverseUUIDs

`func (o *CustomerDetailsData) GetUniverseUUIDs() []string`

GetUniverseUUIDs returns the UniverseUUIDs field if non-nil, zero value otherwise.

### GetUniverseUUIDsOk

`func (o *CustomerDetailsData) GetUniverseUUIDsOk() (*[]string, bool)`

GetUniverseUUIDsOk returns a tuple with the UniverseUUIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUIDs

`func (o *CustomerDetailsData) SetUniverseUUIDs(v []string)`

SetUniverseUUIDs sets UniverseUUIDs field to given value.

### HasUniverseUUIDs

`func (o *CustomerDetailsData) HasUniverseUUIDs() bool`

HasUniverseUUIDs returns a boolean if a field has been set.

### GetUuid

`func (o *CustomerDetailsData) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CustomerDetailsData) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CustomerDetailsData) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CustomerDetailsData) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


