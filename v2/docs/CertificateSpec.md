# CertificateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Name of the certificate, unique within the customer. A generated CA also uses it as its subject CN. | 
**CertType** | [**CertificateType**](CertificateType.md) |  | 
**CertContent** | Pointer to **string** | PEM-encoded CA certificate, optionally followed by the rest of its chain. Omit, together with key_content, to have YBA generate the CA. | [optional] 
**KeyContent** | Pointer to **string** | PEM-encoded private key of the CA certificate. Omit, together with cert_content, to have YBA generate the CA. | [optional] 
**LifetimeYears** | Pointer to **int32** | Validity of a CA that YBA generates, in years. Defaults to yb.tlsCertificate.root.expiryInYears. Only for a generated CA. | [optional] 

## Methods

### NewCertificateSpec

`func NewCertificateSpec(label string, certType CertificateType, ) *CertificateSpec`

NewCertificateSpec instantiates a new CertificateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateSpecWithDefaults

`func NewCertificateSpecWithDefaults() *CertificateSpec`

NewCertificateSpecWithDefaults instantiates a new CertificateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CertificateSpec) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CertificateSpec) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CertificateSpec) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCertType

`func (o *CertificateSpec) GetCertType() CertificateType`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *CertificateSpec) GetCertTypeOk() (*CertificateType, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *CertificateSpec) SetCertType(v CertificateType)`

SetCertType sets CertType field to given value.


### GetCertContent

`func (o *CertificateSpec) GetCertContent() string`

GetCertContent returns the CertContent field if non-nil, zero value otherwise.

### GetCertContentOk

`func (o *CertificateSpec) GetCertContentOk() (*string, bool)`

GetCertContentOk returns a tuple with the CertContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertContent

`func (o *CertificateSpec) SetCertContent(v string)`

SetCertContent sets CertContent field to given value.

### HasCertContent

`func (o *CertificateSpec) HasCertContent() bool`

HasCertContent returns a boolean if a field has been set.

### GetKeyContent

`func (o *CertificateSpec) GetKeyContent() string`

GetKeyContent returns the KeyContent field if non-nil, zero value otherwise.

### GetKeyContentOk

`func (o *CertificateSpec) GetKeyContentOk() (*string, bool)`

GetKeyContentOk returns a tuple with the KeyContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyContent

`func (o *CertificateSpec) SetKeyContent(v string)`

SetKeyContent sets KeyContent field to given value.

### HasKeyContent

`func (o *CertificateSpec) HasKeyContent() bool`

HasKeyContent returns a boolean if a field has been set.

### GetLifetimeYears

`func (o *CertificateSpec) GetLifetimeYears() int32`

GetLifetimeYears returns the LifetimeYears field if non-nil, zero value otherwise.

### GetLifetimeYearsOk

`func (o *CertificateSpec) GetLifetimeYearsOk() (*int32, bool)`

GetLifetimeYearsOk returns a tuple with the LifetimeYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetimeYears

`func (o *CertificateSpec) SetLifetimeYears(v int32)`

SetLifetimeYears sets LifetimeYears field to given value.

### HasLifetimeYears

`func (o *CertificateSpec) HasLifetimeYears() bool`

HasLifetimeYears returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


