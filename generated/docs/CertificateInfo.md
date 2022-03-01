# CertificateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertType** | Pointer to **string** | Type of the certificate | [optional] 
**Certificate** | Pointer to **string** | Certificate path | [optional] 
**Checksum** | Pointer to **string** | The certificate file&#39;s checksum | [optional] [readonly] 
**CustomCertPathParams** | [**CustomCertInfo**](CustomCertInfo.md) |  | 
**CustomHCPKICertInfo** | [**HashicorpVaultConfigParams**](HashicorpVaultConfigParams.md) |  | 
**CustomServerCertInfo** | [**CustomServerCertInfo**](CustomServerCertInfo.md) |  | 
**CustomerUUID** | Pointer to **string** | Customer UUID of the backup which it belongs to | [optional] 
**ExpiryDate** | Pointer to **time.Time** | The certificate&#39;s expiry date | [optional] 
**InUse** | Pointer to **bool** | Indicates whether the certificate is in use. This value is &#x60;true&#x60; if the universe contains a reference to the certificate. | [optional] [readonly] 
**Label** | Pointer to **string** | Certificate label | [optional] 
**PrivateKey** | Pointer to **string** | Private key path | [optional] 
**StartDate** | Pointer to **time.Time** | The certificate&#39;s creation date | [optional] 
**UniverseDetails** | Pointer to [**[]UniverseDetailSubset**](UniverseDetailSubset.md) | Associated universe details for the certificate | [optional] [readonly] 
**Uuid** | Pointer to **string** | Certificate UUID | [optional] [readonly] 

## Methods

### NewCertificateInfo

`func NewCertificateInfo(customCertPathParams CustomCertInfo, customHCPKICertInfo HashicorpVaultConfigParams, customServerCertInfo CustomServerCertInfo, ) *CertificateInfo`

NewCertificateInfo instantiates a new CertificateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateInfoWithDefaults

`func NewCertificateInfoWithDefaults() *CertificateInfo`

NewCertificateInfoWithDefaults instantiates a new CertificateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertType

`func (o *CertificateInfo) GetCertType() string`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *CertificateInfo) GetCertTypeOk() (*string, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *CertificateInfo) SetCertType(v string)`

SetCertType sets CertType field to given value.

### HasCertType

`func (o *CertificateInfo) HasCertType() bool`

HasCertType returns a boolean if a field has been set.

### GetCertificate

`func (o *CertificateInfo) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificateInfo) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificateInfo) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CertificateInfo) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetChecksum

`func (o *CertificateInfo) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *CertificateInfo) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *CertificateInfo) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *CertificateInfo) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetCustomCertPathParams

`func (o *CertificateInfo) GetCustomCertPathParams() CustomCertInfo`

GetCustomCertPathParams returns the CustomCertPathParams field if non-nil, zero value otherwise.

### GetCustomCertPathParamsOk

`func (o *CertificateInfo) GetCustomCertPathParamsOk() (*CustomCertInfo, bool)`

GetCustomCertPathParamsOk returns a tuple with the CustomCertPathParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCertPathParams

`func (o *CertificateInfo) SetCustomCertPathParams(v CustomCertInfo)`

SetCustomCertPathParams sets CustomCertPathParams field to given value.


### GetCustomHCPKICertInfo

`func (o *CertificateInfo) GetCustomHCPKICertInfo() HashicorpVaultConfigParams`

GetCustomHCPKICertInfo returns the CustomHCPKICertInfo field if non-nil, zero value otherwise.

### GetCustomHCPKICertInfoOk

`func (o *CertificateInfo) GetCustomHCPKICertInfoOk() (*HashicorpVaultConfigParams, bool)`

GetCustomHCPKICertInfoOk returns a tuple with the CustomHCPKICertInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHCPKICertInfo

`func (o *CertificateInfo) SetCustomHCPKICertInfo(v HashicorpVaultConfigParams)`

SetCustomHCPKICertInfo sets CustomHCPKICertInfo field to given value.


### GetCustomServerCertInfo

`func (o *CertificateInfo) GetCustomServerCertInfo() CustomServerCertInfo`

GetCustomServerCertInfo returns the CustomServerCertInfo field if non-nil, zero value otherwise.

### GetCustomServerCertInfoOk

`func (o *CertificateInfo) GetCustomServerCertInfoOk() (*CustomServerCertInfo, bool)`

GetCustomServerCertInfoOk returns a tuple with the CustomServerCertInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomServerCertInfo

`func (o *CertificateInfo) SetCustomServerCertInfo(v CustomServerCertInfo)`

SetCustomServerCertInfo sets CustomServerCertInfo field to given value.


### GetCustomerUUID

`func (o *CertificateInfo) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *CertificateInfo) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *CertificateInfo) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.

### HasCustomerUUID

`func (o *CertificateInfo) HasCustomerUUID() bool`

HasCustomerUUID returns a boolean if a field has been set.

### GetExpiryDate

`func (o *CertificateInfo) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *CertificateInfo) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *CertificateInfo) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *CertificateInfo) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetInUse

`func (o *CertificateInfo) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *CertificateInfo) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *CertificateInfo) SetInUse(v bool)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *CertificateInfo) HasInUse() bool`

HasInUse returns a boolean if a field has been set.

### GetLabel

`func (o *CertificateInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CertificateInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CertificateInfo) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CertificateInfo) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPrivateKey

`func (o *CertificateInfo) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CertificateInfo) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CertificateInfo) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *CertificateInfo) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetStartDate

`func (o *CertificateInfo) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CertificateInfo) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CertificateInfo) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CertificateInfo) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetUniverseDetails

`func (o *CertificateInfo) GetUniverseDetails() []UniverseDetailSubset`

GetUniverseDetails returns the UniverseDetails field if non-nil, zero value otherwise.

### GetUniverseDetailsOk

`func (o *CertificateInfo) GetUniverseDetailsOk() (*[]UniverseDetailSubset, bool)`

GetUniverseDetailsOk returns a tuple with the UniverseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseDetails

`func (o *CertificateInfo) SetUniverseDetails(v []UniverseDetailSubset)`

SetUniverseDetails sets UniverseDetails field to given value.

### HasUniverseDetails

`func (o *CertificateInfo) HasUniverseDetails() bool`

HasUniverseDetails returns a boolean if a field has been set.

### GetUuid

`func (o *CertificateInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CertificateInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CertificateInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CertificateInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


