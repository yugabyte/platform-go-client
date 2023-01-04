# CustomerLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | Pointer to **time.Time** | Creation date of license | [optional] [readonly] 
**CustomerUUID** | Pointer to **string** | Customer UUID that owns this license | [optional] 
**License** | **string** | License File Path | 
**LicenseType** | **string** | Type of the license | 
**LicenseUUID** | Pointer to **string** | License UUID | [optional] [readonly] 

## Methods

### NewCustomerLicense

`func NewCustomerLicense(license string, licenseType string, ) *CustomerLicense`

NewCustomerLicense instantiates a new CustomerLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerLicenseWithDefaults

`func NewCustomerLicenseWithDefaults() *CustomerLicense`

NewCustomerLicenseWithDefaults instantiates a new CustomerLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDate

`func (o *CustomerLicense) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *CustomerLicense) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *CustomerLicense) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *CustomerLicense) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetCustomerUUID

`func (o *CustomerLicense) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *CustomerLicense) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *CustomerLicense) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.

### HasCustomerUUID

`func (o *CustomerLicense) HasCustomerUUID() bool`

HasCustomerUUID returns a boolean if a field has been set.

### GetLicense

`func (o *CustomerLicense) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *CustomerLicense) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *CustomerLicense) SetLicense(v string)`

SetLicense sets License field to given value.


### GetLicenseType

`func (o *CustomerLicense) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *CustomerLicense) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *CustomerLicense) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.


### GetLicenseUUID

`func (o *CustomerLicense) GetLicenseUUID() string`

GetLicenseUUID returns the LicenseUUID field if non-nil, zero value otherwise.

### GetLicenseUUIDOk

`func (o *CustomerLicense) GetLicenseUUIDOk() (*string, bool)`

GetLicenseUUIDOk returns a tuple with the LicenseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUUID

`func (o *CustomerLicense) SetLicenseUUID(v string)`

SetLicenseUUID sets LicenseUUID field to given value.

### HasLicenseUUID

`func (o *CustomerLicense) HasLicenseUUID() bool`

HasLicenseUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


