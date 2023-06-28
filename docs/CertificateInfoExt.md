# CertificateInfoExt

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
**ExpiryDate** | Pointer to **time.Time** | The certificate&#39;s expiry date. Deprecated: Use expirtyDateIso instead | [optional] 
**ExpiryDateIso** | Pointer to **time.Time** | The certificate&#39;s expiry date | [optional] 
**InUse** | Pointer to **bool** | Indicates whether the certificate is in use. This value is &#x60;true&#x60; if the universe contains a reference to the certificate. | [optional] [readonly] 
**Label** | Pointer to **string** | Certificate label | [optional] 
**PrivateKey** | Pointer to **string** | Private key path | [optional] 
**StartDate** | Pointer to **time.Time** | The certificate&#39;s creation date. Deprecated: use stateDateIso instead | [optional] 
**StartDateIso** | Pointer to **time.Time** | The certificate&#39;s creation date | [optional] 
**UniverseDetailSubsets** | [**[]UniverseDetailSubset**](UniverseDetailSubset.md) |  | 
**UniverseDetails** | Pointer to [**[]UniverseDetailSubset**](UniverseDetailSubset.md) | Associated universe details for the certificate | [optional] [readonly] 
**Uuid** | Pointer to **string** | Certificate UUID | [optional] [readonly] 

## Methods

### NewCertificateInfoExt

`func NewCertificateInfoExt(customCertPathParams CustomCertInfo, customHCPKICertInfo HashicorpVaultConfigParams, customServerCertInfo CustomServerCertInfo, universeDetailSubsets []UniverseDetailSubset, ) *CertificateInfoExt`

NewCertificateInfoExt instantiates a new CertificateInfoExt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateInfoExtWithDefaults

`func NewCertificateInfoExtWithDefaults() *CertificateInfoExt`

NewCertificateInfoExtWithDefaults instantiates a new CertificateInfoExt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertType

`func (o *CertificateInfoExt) GetCertType() string`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *CertificateInfoExt) GetCertTypeOk() (*string, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *CertificateInfoExt) SetCertType(v string)`

SetCertType sets CertType field to given value.

### HasCertType

`func (o *CertificateInfoExt) HasCertType() bool`

HasCertType returns a boolean if a field has been set.

### GetCertificate

`func (o *CertificateInfoExt) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificateInfoExt) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificateInfoExt) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CertificateInfoExt) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetChecksum

`func (o *CertificateInfoExt) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *CertificateInfoExt) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *CertificateInfoExt) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *CertificateInfoExt) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetCustomCertPathParams

`func (o *CertificateInfoExt) GetCustomCertPathParams() CustomCertInfo`

GetCustomCertPathParams returns the CustomCertPathParams field if non-nil, zero value otherwise.

### GetCustomCertPathParamsOk

`func (o *CertificateInfoExt) GetCustomCertPathParamsOk() (*CustomCertInfo, bool)`

GetCustomCertPathParamsOk returns a tuple with the CustomCertPathParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCertPathParams

`func (o *CertificateInfoExt) SetCustomCertPathParams(v CustomCertInfo)`

SetCustomCertPathParams sets CustomCertPathParams field to given value.


### GetCustomHCPKICertInfo

`func (o *CertificateInfoExt) GetCustomHCPKICertInfo() HashicorpVaultConfigParams`

GetCustomHCPKICertInfo returns the CustomHCPKICertInfo field if non-nil, zero value otherwise.

### GetCustomHCPKICertInfoOk

`func (o *CertificateInfoExt) GetCustomHCPKICertInfoOk() (*HashicorpVaultConfigParams, bool)`

GetCustomHCPKICertInfoOk returns a tuple with the CustomHCPKICertInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHCPKICertInfo

`func (o *CertificateInfoExt) SetCustomHCPKICertInfo(v HashicorpVaultConfigParams)`

SetCustomHCPKICertInfo sets CustomHCPKICertInfo field to given value.


### GetCustomServerCertInfo

`func (o *CertificateInfoExt) GetCustomServerCertInfo() CustomServerCertInfo`

GetCustomServerCertInfo returns the CustomServerCertInfo field if non-nil, zero value otherwise.

### GetCustomServerCertInfoOk

`func (o *CertificateInfoExt) GetCustomServerCertInfoOk() (*CustomServerCertInfo, bool)`

GetCustomServerCertInfoOk returns a tuple with the CustomServerCertInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomServerCertInfo

`func (o *CertificateInfoExt) SetCustomServerCertInfo(v CustomServerCertInfo)`

SetCustomServerCertInfo sets CustomServerCertInfo field to given value.


### GetCustomerUUID

`func (o *CertificateInfoExt) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *CertificateInfoExt) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *CertificateInfoExt) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.

### HasCustomerUUID

`func (o *CertificateInfoExt) HasCustomerUUID() bool`

HasCustomerUUID returns a boolean if a field has been set.

### GetExpiryDate

`func (o *CertificateInfoExt) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *CertificateInfoExt) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *CertificateInfoExt) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *CertificateInfoExt) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetExpiryDateIso

`func (o *CertificateInfoExt) GetExpiryDateIso() time.Time`

GetExpiryDateIso returns the ExpiryDateIso field if non-nil, zero value otherwise.

### GetExpiryDateIsoOk

`func (o *CertificateInfoExt) GetExpiryDateIsoOk() (*time.Time, bool)`

GetExpiryDateIsoOk returns a tuple with the ExpiryDateIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDateIso

`func (o *CertificateInfoExt) SetExpiryDateIso(v time.Time)`

SetExpiryDateIso sets ExpiryDateIso field to given value.

### HasExpiryDateIso

`func (o *CertificateInfoExt) HasExpiryDateIso() bool`

HasExpiryDateIso returns a boolean if a field has been set.

### GetInUse

`func (o *CertificateInfoExt) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *CertificateInfoExt) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *CertificateInfoExt) SetInUse(v bool)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *CertificateInfoExt) HasInUse() bool`

HasInUse returns a boolean if a field has been set.

### GetLabel

`func (o *CertificateInfoExt) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CertificateInfoExt) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CertificateInfoExt) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CertificateInfoExt) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPrivateKey

`func (o *CertificateInfoExt) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CertificateInfoExt) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CertificateInfoExt) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *CertificateInfoExt) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetStartDate

`func (o *CertificateInfoExt) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CertificateInfoExt) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CertificateInfoExt) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CertificateInfoExt) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStartDateIso

`func (o *CertificateInfoExt) GetStartDateIso() time.Time`

GetStartDateIso returns the StartDateIso field if non-nil, zero value otherwise.

### GetStartDateIsoOk

`func (o *CertificateInfoExt) GetStartDateIsoOk() (*time.Time, bool)`

GetStartDateIsoOk returns a tuple with the StartDateIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateIso

`func (o *CertificateInfoExt) SetStartDateIso(v time.Time)`

SetStartDateIso sets StartDateIso field to given value.

### HasStartDateIso

`func (o *CertificateInfoExt) HasStartDateIso() bool`

HasStartDateIso returns a boolean if a field has been set.

### GetUniverseDetailSubsets

`func (o *CertificateInfoExt) GetUniverseDetailSubsets() []UniverseDetailSubset`

GetUniverseDetailSubsets returns the UniverseDetailSubsets field if non-nil, zero value otherwise.

### GetUniverseDetailSubsetsOk

`func (o *CertificateInfoExt) GetUniverseDetailSubsetsOk() (*[]UniverseDetailSubset, bool)`

GetUniverseDetailSubsetsOk returns a tuple with the UniverseDetailSubsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseDetailSubsets

`func (o *CertificateInfoExt) SetUniverseDetailSubsets(v []UniverseDetailSubset)`

SetUniverseDetailSubsets sets UniverseDetailSubsets field to given value.


### GetUniverseDetails

`func (o *CertificateInfoExt) GetUniverseDetails() []UniverseDetailSubset`

GetUniverseDetails returns the UniverseDetails field if non-nil, zero value otherwise.

### GetUniverseDetailsOk

`func (o *CertificateInfoExt) GetUniverseDetailsOk() (*[]UniverseDetailSubset, bool)`

GetUniverseDetailsOk returns a tuple with the UniverseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseDetails

`func (o *CertificateInfoExt) SetUniverseDetails(v []UniverseDetailSubset)`

SetUniverseDetails sets UniverseDetails field to given value.

### HasUniverseDetails

`func (o *CertificateInfoExt) HasUniverseDetails() bool`

HasUniverseDetails returns a boolean if a field has been set.

### GetUuid

`func (o *CertificateInfoExt) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CertificateInfoExt) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CertificateInfoExt) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CertificateInfoExt) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


