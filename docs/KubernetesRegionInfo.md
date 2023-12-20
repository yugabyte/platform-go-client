# KubernetesRegionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertManagerClusterIssuer** | Pointer to **string** |  | [optional] 
**CertManagerIssuer** | Pointer to **string** |  | [optional] 
**KubeConfig** | Pointer to **string** |  | [optional] [readonly] 
**KubeConfigContent** | Pointer to **string** |  | [optional] 
**KubeConfigName** | Pointer to **string** |  | [optional] 
**KubeDomain** | Pointer to **string** |  | [optional] 
**KubeNamespace** | Pointer to **string** |  | [optional] 
**KubePodAddressTemplate** | Pointer to **string** |  | [optional] 
**KubernetesImagePullSecretName** | Pointer to **string** |  | [optional] 
**KubernetesImageRegistry** | Pointer to **string** |  | [optional] 
**KubernetesProvider** | Pointer to **string** |  | [optional] 
**KubernetesPullSecret** | Pointer to **string** |  | [optional] [readonly] 
**KubernetesPullSecretContent** | Pointer to **string** |  | [optional] 
**KubernetesPullSecretName** | Pointer to **string** |  | [optional] 
**KubernetesServiceAccount** | Pointer to **string** | Deprecated since YBA version 2.17.3.0. kubernetes service account is not needed. | [optional] 
**KubernetesStorageClass** | Pointer to **string** |  | [optional] 
**Overrides** | Pointer to **string** |  | [optional] 

## Methods

### NewKubernetesRegionInfo

`func NewKubernetesRegionInfo() *KubernetesRegionInfo`

NewKubernetesRegionInfo instantiates a new KubernetesRegionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesRegionInfoWithDefaults

`func NewKubernetesRegionInfoWithDefaults() *KubernetesRegionInfo`

NewKubernetesRegionInfoWithDefaults instantiates a new KubernetesRegionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertManagerClusterIssuer

`func (o *KubernetesRegionInfo) GetCertManagerClusterIssuer() string`

GetCertManagerClusterIssuer returns the CertManagerClusterIssuer field if non-nil, zero value otherwise.

### GetCertManagerClusterIssuerOk

`func (o *KubernetesRegionInfo) GetCertManagerClusterIssuerOk() (*string, bool)`

GetCertManagerClusterIssuerOk returns a tuple with the CertManagerClusterIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertManagerClusterIssuer

`func (o *KubernetesRegionInfo) SetCertManagerClusterIssuer(v string)`

SetCertManagerClusterIssuer sets CertManagerClusterIssuer field to given value.

### HasCertManagerClusterIssuer

`func (o *KubernetesRegionInfo) HasCertManagerClusterIssuer() bool`

HasCertManagerClusterIssuer returns a boolean if a field has been set.

### GetCertManagerIssuer

`func (o *KubernetesRegionInfo) GetCertManagerIssuer() string`

GetCertManagerIssuer returns the CertManagerIssuer field if non-nil, zero value otherwise.

### GetCertManagerIssuerOk

`func (o *KubernetesRegionInfo) GetCertManagerIssuerOk() (*string, bool)`

GetCertManagerIssuerOk returns a tuple with the CertManagerIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertManagerIssuer

`func (o *KubernetesRegionInfo) SetCertManagerIssuer(v string)`

SetCertManagerIssuer sets CertManagerIssuer field to given value.

### HasCertManagerIssuer

`func (o *KubernetesRegionInfo) HasCertManagerIssuer() bool`

HasCertManagerIssuer returns a boolean if a field has been set.

### GetKubeConfig

`func (o *KubernetesRegionInfo) GetKubeConfig() string`

GetKubeConfig returns the KubeConfig field if non-nil, zero value otherwise.

### GetKubeConfigOk

`func (o *KubernetesRegionInfo) GetKubeConfigOk() (*string, bool)`

GetKubeConfigOk returns a tuple with the KubeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfig

`func (o *KubernetesRegionInfo) SetKubeConfig(v string)`

SetKubeConfig sets KubeConfig field to given value.

### HasKubeConfig

`func (o *KubernetesRegionInfo) HasKubeConfig() bool`

HasKubeConfig returns a boolean if a field has been set.

### GetKubeConfigContent

`func (o *KubernetesRegionInfo) GetKubeConfigContent() string`

GetKubeConfigContent returns the KubeConfigContent field if non-nil, zero value otherwise.

### GetKubeConfigContentOk

`func (o *KubernetesRegionInfo) GetKubeConfigContentOk() (*string, bool)`

GetKubeConfigContentOk returns a tuple with the KubeConfigContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfigContent

`func (o *KubernetesRegionInfo) SetKubeConfigContent(v string)`

SetKubeConfigContent sets KubeConfigContent field to given value.

### HasKubeConfigContent

`func (o *KubernetesRegionInfo) HasKubeConfigContent() bool`

HasKubeConfigContent returns a boolean if a field has been set.

### GetKubeConfigName

`func (o *KubernetesRegionInfo) GetKubeConfigName() string`

GetKubeConfigName returns the KubeConfigName field if non-nil, zero value otherwise.

### GetKubeConfigNameOk

`func (o *KubernetesRegionInfo) GetKubeConfigNameOk() (*string, bool)`

GetKubeConfigNameOk returns a tuple with the KubeConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfigName

`func (o *KubernetesRegionInfo) SetKubeConfigName(v string)`

SetKubeConfigName sets KubeConfigName field to given value.

### HasKubeConfigName

`func (o *KubernetesRegionInfo) HasKubeConfigName() bool`

HasKubeConfigName returns a boolean if a field has been set.

### GetKubeDomain

`func (o *KubernetesRegionInfo) GetKubeDomain() string`

GetKubeDomain returns the KubeDomain field if non-nil, zero value otherwise.

### GetKubeDomainOk

`func (o *KubernetesRegionInfo) GetKubeDomainOk() (*string, bool)`

GetKubeDomainOk returns a tuple with the KubeDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeDomain

`func (o *KubernetesRegionInfo) SetKubeDomain(v string)`

SetKubeDomain sets KubeDomain field to given value.

### HasKubeDomain

`func (o *KubernetesRegionInfo) HasKubeDomain() bool`

HasKubeDomain returns a boolean if a field has been set.

### GetKubeNamespace

`func (o *KubernetesRegionInfo) GetKubeNamespace() string`

GetKubeNamespace returns the KubeNamespace field if non-nil, zero value otherwise.

### GetKubeNamespaceOk

`func (o *KubernetesRegionInfo) GetKubeNamespaceOk() (*string, bool)`

GetKubeNamespaceOk returns a tuple with the KubeNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeNamespace

`func (o *KubernetesRegionInfo) SetKubeNamespace(v string)`

SetKubeNamespace sets KubeNamespace field to given value.

### HasKubeNamespace

`func (o *KubernetesRegionInfo) HasKubeNamespace() bool`

HasKubeNamespace returns a boolean if a field has been set.

### GetKubePodAddressTemplate

`func (o *KubernetesRegionInfo) GetKubePodAddressTemplate() string`

GetKubePodAddressTemplate returns the KubePodAddressTemplate field if non-nil, zero value otherwise.

### GetKubePodAddressTemplateOk

`func (o *KubernetesRegionInfo) GetKubePodAddressTemplateOk() (*string, bool)`

GetKubePodAddressTemplateOk returns a tuple with the KubePodAddressTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubePodAddressTemplate

`func (o *KubernetesRegionInfo) SetKubePodAddressTemplate(v string)`

SetKubePodAddressTemplate sets KubePodAddressTemplate field to given value.

### HasKubePodAddressTemplate

`func (o *KubernetesRegionInfo) HasKubePodAddressTemplate() bool`

HasKubePodAddressTemplate returns a boolean if a field has been set.

### GetKubernetesImagePullSecretName

`func (o *KubernetesRegionInfo) GetKubernetesImagePullSecretName() string`

GetKubernetesImagePullSecretName returns the KubernetesImagePullSecretName field if non-nil, zero value otherwise.

### GetKubernetesImagePullSecretNameOk

`func (o *KubernetesRegionInfo) GetKubernetesImagePullSecretNameOk() (*string, bool)`

GetKubernetesImagePullSecretNameOk returns a tuple with the KubernetesImagePullSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesImagePullSecretName

`func (o *KubernetesRegionInfo) SetKubernetesImagePullSecretName(v string)`

SetKubernetesImagePullSecretName sets KubernetesImagePullSecretName field to given value.

### HasKubernetesImagePullSecretName

`func (o *KubernetesRegionInfo) HasKubernetesImagePullSecretName() bool`

HasKubernetesImagePullSecretName returns a boolean if a field has been set.

### GetKubernetesImageRegistry

`func (o *KubernetesRegionInfo) GetKubernetesImageRegistry() string`

GetKubernetesImageRegistry returns the KubernetesImageRegistry field if non-nil, zero value otherwise.

### GetKubernetesImageRegistryOk

`func (o *KubernetesRegionInfo) GetKubernetesImageRegistryOk() (*string, bool)`

GetKubernetesImageRegistryOk returns a tuple with the KubernetesImageRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesImageRegistry

`func (o *KubernetesRegionInfo) SetKubernetesImageRegistry(v string)`

SetKubernetesImageRegistry sets KubernetesImageRegistry field to given value.

### HasKubernetesImageRegistry

`func (o *KubernetesRegionInfo) HasKubernetesImageRegistry() bool`

HasKubernetesImageRegistry returns a boolean if a field has been set.

### GetKubernetesProvider

`func (o *KubernetesRegionInfo) GetKubernetesProvider() string`

GetKubernetesProvider returns the KubernetesProvider field if non-nil, zero value otherwise.

### GetKubernetesProviderOk

`func (o *KubernetesRegionInfo) GetKubernetesProviderOk() (*string, bool)`

GetKubernetesProviderOk returns a tuple with the KubernetesProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesProvider

`func (o *KubernetesRegionInfo) SetKubernetesProvider(v string)`

SetKubernetesProvider sets KubernetesProvider field to given value.

### HasKubernetesProvider

`func (o *KubernetesRegionInfo) HasKubernetesProvider() bool`

HasKubernetesProvider returns a boolean if a field has been set.

### GetKubernetesPullSecret

`func (o *KubernetesRegionInfo) GetKubernetesPullSecret() string`

GetKubernetesPullSecret returns the KubernetesPullSecret field if non-nil, zero value otherwise.

### GetKubernetesPullSecretOk

`func (o *KubernetesRegionInfo) GetKubernetesPullSecretOk() (*string, bool)`

GetKubernetesPullSecretOk returns a tuple with the KubernetesPullSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesPullSecret

`func (o *KubernetesRegionInfo) SetKubernetesPullSecret(v string)`

SetKubernetesPullSecret sets KubernetesPullSecret field to given value.

### HasKubernetesPullSecret

`func (o *KubernetesRegionInfo) HasKubernetesPullSecret() bool`

HasKubernetesPullSecret returns a boolean if a field has been set.

### GetKubernetesPullSecretContent

`func (o *KubernetesRegionInfo) GetKubernetesPullSecretContent() string`

GetKubernetesPullSecretContent returns the KubernetesPullSecretContent field if non-nil, zero value otherwise.

### GetKubernetesPullSecretContentOk

`func (o *KubernetesRegionInfo) GetKubernetesPullSecretContentOk() (*string, bool)`

GetKubernetesPullSecretContentOk returns a tuple with the KubernetesPullSecretContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesPullSecretContent

`func (o *KubernetesRegionInfo) SetKubernetesPullSecretContent(v string)`

SetKubernetesPullSecretContent sets KubernetesPullSecretContent field to given value.

### HasKubernetesPullSecretContent

`func (o *KubernetesRegionInfo) HasKubernetesPullSecretContent() bool`

HasKubernetesPullSecretContent returns a boolean if a field has been set.

### GetKubernetesPullSecretName

`func (o *KubernetesRegionInfo) GetKubernetesPullSecretName() string`

GetKubernetesPullSecretName returns the KubernetesPullSecretName field if non-nil, zero value otherwise.

### GetKubernetesPullSecretNameOk

`func (o *KubernetesRegionInfo) GetKubernetesPullSecretNameOk() (*string, bool)`

GetKubernetesPullSecretNameOk returns a tuple with the KubernetesPullSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesPullSecretName

`func (o *KubernetesRegionInfo) SetKubernetesPullSecretName(v string)`

SetKubernetesPullSecretName sets KubernetesPullSecretName field to given value.

### HasKubernetesPullSecretName

`func (o *KubernetesRegionInfo) HasKubernetesPullSecretName() bool`

HasKubernetesPullSecretName returns a boolean if a field has been set.

### GetKubernetesServiceAccount

`func (o *KubernetesRegionInfo) GetKubernetesServiceAccount() string`

GetKubernetesServiceAccount returns the KubernetesServiceAccount field if non-nil, zero value otherwise.

### GetKubernetesServiceAccountOk

`func (o *KubernetesRegionInfo) GetKubernetesServiceAccountOk() (*string, bool)`

GetKubernetesServiceAccountOk returns a tuple with the KubernetesServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesServiceAccount

`func (o *KubernetesRegionInfo) SetKubernetesServiceAccount(v string)`

SetKubernetesServiceAccount sets KubernetesServiceAccount field to given value.

### HasKubernetesServiceAccount

`func (o *KubernetesRegionInfo) HasKubernetesServiceAccount() bool`

HasKubernetesServiceAccount returns a boolean if a field has been set.

### GetKubernetesStorageClass

`func (o *KubernetesRegionInfo) GetKubernetesStorageClass() string`

GetKubernetesStorageClass returns the KubernetesStorageClass field if non-nil, zero value otherwise.

### GetKubernetesStorageClassOk

`func (o *KubernetesRegionInfo) GetKubernetesStorageClassOk() (*string, bool)`

GetKubernetesStorageClassOk returns a tuple with the KubernetesStorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesStorageClass

`func (o *KubernetesRegionInfo) SetKubernetesStorageClass(v string)`

SetKubernetesStorageClass sets KubernetesStorageClass field to given value.

### HasKubernetesStorageClass

`func (o *KubernetesRegionInfo) HasKubernetesStorageClass() bool`

HasKubernetesStorageClass returns a boolean if a field has been set.

### GetOverrides

`func (o *KubernetesRegionInfo) GetOverrides() string`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *KubernetesRegionInfo) GetOverridesOk() (*string, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *KubernetesRegionInfo) SetOverrides(v string)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *KubernetesRegionInfo) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


