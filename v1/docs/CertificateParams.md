# CertificateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertContent** | **string** |  | 
**CertExpiry** | **int64** |  | 
**CertStart** | **int64** |  | 
**CertType** | **string** |  | 
**CustomCertInfo** | [**CustomCertInfo**](CustomCertInfo.md) |  | 
**CustomServerCertData** | [**CustomServerCertData**](CustomServerCertData.md) |  | 
**HcVaultCertParams** | [**HashicorpVaultConfigParams**](HashicorpVaultConfigParams.md) |  | 
**KeyContent** | **string** |  | 
**Label** | **string** |  | 

## Methods

### NewCertificateParams

`func NewCertificateParams(certContent string, certExpiry int64, certStart int64, certType string, customCertInfo CustomCertInfo, customServerCertData CustomServerCertData, hcVaultCertParams HashicorpVaultConfigParams, keyContent string, label string, ) *CertificateParams`

NewCertificateParams instantiates a new CertificateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateParamsWithDefaults

`func NewCertificateParamsWithDefaults() *CertificateParams`

NewCertificateParamsWithDefaults instantiates a new CertificateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertContent

`func (o *CertificateParams) GetCertContent() string`

GetCertContent returns the CertContent field if non-nil, zero value otherwise.

### GetCertContentOk

`func (o *CertificateParams) GetCertContentOk() (*string, bool)`

GetCertContentOk returns a tuple with the CertContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertContent

`func (o *CertificateParams) SetCertContent(v string)`

SetCertContent sets CertContent field to given value.


### GetCertExpiry

`func (o *CertificateParams) GetCertExpiry() int64`

GetCertExpiry returns the CertExpiry field if non-nil, zero value otherwise.

### GetCertExpiryOk

`func (o *CertificateParams) GetCertExpiryOk() (*int64, bool)`

GetCertExpiryOk returns a tuple with the CertExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertExpiry

`func (o *CertificateParams) SetCertExpiry(v int64)`

SetCertExpiry sets CertExpiry field to given value.


### GetCertStart

`func (o *CertificateParams) GetCertStart() int64`

GetCertStart returns the CertStart field if non-nil, zero value otherwise.

### GetCertStartOk

`func (o *CertificateParams) GetCertStartOk() (*int64, bool)`

GetCertStartOk returns a tuple with the CertStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStart

`func (o *CertificateParams) SetCertStart(v int64)`

SetCertStart sets CertStart field to given value.


### GetCertType

`func (o *CertificateParams) GetCertType() string`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *CertificateParams) GetCertTypeOk() (*string, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *CertificateParams) SetCertType(v string)`

SetCertType sets CertType field to given value.


### GetCustomCertInfo

`func (o *CertificateParams) GetCustomCertInfo() CustomCertInfo`

GetCustomCertInfo returns the CustomCertInfo field if non-nil, zero value otherwise.

### GetCustomCertInfoOk

`func (o *CertificateParams) GetCustomCertInfoOk() (*CustomCertInfo, bool)`

GetCustomCertInfoOk returns a tuple with the CustomCertInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCertInfo

`func (o *CertificateParams) SetCustomCertInfo(v CustomCertInfo)`

SetCustomCertInfo sets CustomCertInfo field to given value.


### GetCustomServerCertData

`func (o *CertificateParams) GetCustomServerCertData() CustomServerCertData`

GetCustomServerCertData returns the CustomServerCertData field if non-nil, zero value otherwise.

### GetCustomServerCertDataOk

`func (o *CertificateParams) GetCustomServerCertDataOk() (*CustomServerCertData, bool)`

GetCustomServerCertDataOk returns a tuple with the CustomServerCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomServerCertData

`func (o *CertificateParams) SetCustomServerCertData(v CustomServerCertData)`

SetCustomServerCertData sets CustomServerCertData field to given value.


### GetHcVaultCertParams

`func (o *CertificateParams) GetHcVaultCertParams() HashicorpVaultConfigParams`

GetHcVaultCertParams returns the HcVaultCertParams field if non-nil, zero value otherwise.

### GetHcVaultCertParamsOk

`func (o *CertificateParams) GetHcVaultCertParamsOk() (*HashicorpVaultConfigParams, bool)`

GetHcVaultCertParamsOk returns a tuple with the HcVaultCertParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHcVaultCertParams

`func (o *CertificateParams) SetHcVaultCertParams(v HashicorpVaultConfigParams)`

SetHcVaultCertParams sets HcVaultCertParams field to given value.


### GetKeyContent

`func (o *CertificateParams) GetKeyContent() string`

GetKeyContent returns the KeyContent field if non-nil, zero value otherwise.

### GetKeyContentOk

`func (o *CertificateParams) GetKeyContentOk() (*string, bool)`

GetKeyContentOk returns a tuple with the KeyContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyContent

`func (o *CertificateParams) SetKeyContent(v string)`

SetKeyContent sets KeyContent field to given value.


### GetLabel

`func (o *CertificateParams) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CertificateParams) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CertificateParams) SetLabel(v string)`

SetLabel sets Label field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


