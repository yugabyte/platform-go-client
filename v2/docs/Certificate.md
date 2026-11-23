# Certificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | UUID of the certificate | [readonly] 
**Label** | **string** | Name of the certificate | [readonly] 
**CertType** | [**CertificateType**](CertificateType.md) |  | 
**StartDate** | **time.Time** | Start of the certificate&#39;s validity | [readonly] 
**ExpiryDate** | **time.Time** | End of the certificate&#39;s validity | [readonly] 

## Methods

### NewCertificate

`func NewCertificate(uuid string, label string, certType CertificateType, startDate time.Time, expiryDate time.Time, ) *Certificate`

NewCertificate instantiates a new Certificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateWithDefaults

`func NewCertificateWithDefaults() *Certificate`

NewCertificateWithDefaults instantiates a new Certificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *Certificate) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Certificate) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Certificate) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetLabel

`func (o *Certificate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Certificate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Certificate) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCertType

`func (o *Certificate) GetCertType() CertificateType`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *Certificate) GetCertTypeOk() (*CertificateType, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *Certificate) SetCertType(v CertificateType)`

SetCertType sets CertType field to given value.


### GetStartDate

`func (o *Certificate) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Certificate) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Certificate) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetExpiryDate

`func (o *Certificate) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *Certificate) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *Certificate) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


