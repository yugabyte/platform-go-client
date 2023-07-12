# KubernetesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KubeConfig** | Pointer to **string** |  | [optional] 
**KubeConfigContent** | Pointer to **string** |  | [optional] 
**KubeConfigName** | Pointer to **string** |  | [optional] 
**KubernetesImagePullSecretName** | Pointer to **string** |  | [optional] 
**KubernetesImageRegistry** | Pointer to **string** |  | [optional] 
**KubernetesProvider** | Pointer to **string** |  | [optional] 
**KubernetesPullSecret** | Pointer to **string** |  | [optional] 
**KubernetesPullSecretContent** | Pointer to **string** |  | [optional] 
**KubernetesPullSecretName** | Pointer to **string** |  | [optional] 
**KubernetesServiceAccount** | Pointer to **string** | DEPRECATED: kubernetes service account is not needed. | [optional] 
**KubernetesStorageClass** | Pointer to **string** |  | [optional] 

## Methods

### NewKubernetesInfo

`func NewKubernetesInfo() *KubernetesInfo`

NewKubernetesInfo instantiates a new KubernetesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesInfoWithDefaults

`func NewKubernetesInfoWithDefaults() *KubernetesInfo`

NewKubernetesInfoWithDefaults instantiates a new KubernetesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubeConfig

`func (o *KubernetesInfo) GetKubeConfig() string`

GetKubeConfig returns the KubeConfig field if non-nil, zero value otherwise.

### GetKubeConfigOk

`func (o *KubernetesInfo) GetKubeConfigOk() (*string, bool)`

GetKubeConfigOk returns a tuple with the KubeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfig

`func (o *KubernetesInfo) SetKubeConfig(v string)`

SetKubeConfig sets KubeConfig field to given value.

### HasKubeConfig

`func (o *KubernetesInfo) HasKubeConfig() bool`

HasKubeConfig returns a boolean if a field has been set.

### GetKubeConfigContent

`func (o *KubernetesInfo) GetKubeConfigContent() string`

GetKubeConfigContent returns the KubeConfigContent field if non-nil, zero value otherwise.

### GetKubeConfigContentOk

`func (o *KubernetesInfo) GetKubeConfigContentOk() (*string, bool)`

GetKubeConfigContentOk returns a tuple with the KubeConfigContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfigContent

`func (o *KubernetesInfo) SetKubeConfigContent(v string)`

SetKubeConfigContent sets KubeConfigContent field to given value.

### HasKubeConfigContent

`func (o *KubernetesInfo) HasKubeConfigContent() bool`

HasKubeConfigContent returns a boolean if a field has been set.

### GetKubeConfigName

`func (o *KubernetesInfo) GetKubeConfigName() string`

GetKubeConfigName returns the KubeConfigName field if non-nil, zero value otherwise.

### GetKubeConfigNameOk

`func (o *KubernetesInfo) GetKubeConfigNameOk() (*string, bool)`

GetKubeConfigNameOk returns a tuple with the KubeConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfigName

`func (o *KubernetesInfo) SetKubeConfigName(v string)`

SetKubeConfigName sets KubeConfigName field to given value.

### HasKubeConfigName

`func (o *KubernetesInfo) HasKubeConfigName() bool`

HasKubeConfigName returns a boolean if a field has been set.

### GetKubernetesImagePullSecretName

`func (o *KubernetesInfo) GetKubernetesImagePullSecretName() string`

GetKubernetesImagePullSecretName returns the KubernetesImagePullSecretName field if non-nil, zero value otherwise.

### GetKubernetesImagePullSecretNameOk

`func (o *KubernetesInfo) GetKubernetesImagePullSecretNameOk() (*string, bool)`

GetKubernetesImagePullSecretNameOk returns a tuple with the KubernetesImagePullSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesImagePullSecretName

`func (o *KubernetesInfo) SetKubernetesImagePullSecretName(v string)`

SetKubernetesImagePullSecretName sets KubernetesImagePullSecretName field to given value.

### HasKubernetesImagePullSecretName

`func (o *KubernetesInfo) HasKubernetesImagePullSecretName() bool`

HasKubernetesImagePullSecretName returns a boolean if a field has been set.

### GetKubernetesImageRegistry

`func (o *KubernetesInfo) GetKubernetesImageRegistry() string`

GetKubernetesImageRegistry returns the KubernetesImageRegistry field if non-nil, zero value otherwise.

### GetKubernetesImageRegistryOk

`func (o *KubernetesInfo) GetKubernetesImageRegistryOk() (*string, bool)`

GetKubernetesImageRegistryOk returns a tuple with the KubernetesImageRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesImageRegistry

`func (o *KubernetesInfo) SetKubernetesImageRegistry(v string)`

SetKubernetesImageRegistry sets KubernetesImageRegistry field to given value.

### HasKubernetesImageRegistry

`func (o *KubernetesInfo) HasKubernetesImageRegistry() bool`

HasKubernetesImageRegistry returns a boolean if a field has been set.

### GetKubernetesProvider

`func (o *KubernetesInfo) GetKubernetesProvider() string`

GetKubernetesProvider returns the KubernetesProvider field if non-nil, zero value otherwise.

### GetKubernetesProviderOk

`func (o *KubernetesInfo) GetKubernetesProviderOk() (*string, bool)`

GetKubernetesProviderOk returns a tuple with the KubernetesProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesProvider

`func (o *KubernetesInfo) SetKubernetesProvider(v string)`

SetKubernetesProvider sets KubernetesProvider field to given value.

### HasKubernetesProvider

`func (o *KubernetesInfo) HasKubernetesProvider() bool`

HasKubernetesProvider returns a boolean if a field has been set.

### GetKubernetesPullSecret

`func (o *KubernetesInfo) GetKubernetesPullSecret() string`

GetKubernetesPullSecret returns the KubernetesPullSecret field if non-nil, zero value otherwise.

### GetKubernetesPullSecretOk

`func (o *KubernetesInfo) GetKubernetesPullSecretOk() (*string, bool)`

GetKubernetesPullSecretOk returns a tuple with the KubernetesPullSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesPullSecret

`func (o *KubernetesInfo) SetKubernetesPullSecret(v string)`

SetKubernetesPullSecret sets KubernetesPullSecret field to given value.

### HasKubernetesPullSecret

`func (o *KubernetesInfo) HasKubernetesPullSecret() bool`

HasKubernetesPullSecret returns a boolean if a field has been set.

### GetKubernetesPullSecretContent

`func (o *KubernetesInfo) GetKubernetesPullSecretContent() string`

GetKubernetesPullSecretContent returns the KubernetesPullSecretContent field if non-nil, zero value otherwise.

### GetKubernetesPullSecretContentOk

`func (o *KubernetesInfo) GetKubernetesPullSecretContentOk() (*string, bool)`

GetKubernetesPullSecretContentOk returns a tuple with the KubernetesPullSecretContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesPullSecretContent

`func (o *KubernetesInfo) SetKubernetesPullSecretContent(v string)`

SetKubernetesPullSecretContent sets KubernetesPullSecretContent field to given value.

### HasKubernetesPullSecretContent

`func (o *KubernetesInfo) HasKubernetesPullSecretContent() bool`

HasKubernetesPullSecretContent returns a boolean if a field has been set.

### GetKubernetesPullSecretName

`func (o *KubernetesInfo) GetKubernetesPullSecretName() string`

GetKubernetesPullSecretName returns the KubernetesPullSecretName field if non-nil, zero value otherwise.

### GetKubernetesPullSecretNameOk

`func (o *KubernetesInfo) GetKubernetesPullSecretNameOk() (*string, bool)`

GetKubernetesPullSecretNameOk returns a tuple with the KubernetesPullSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesPullSecretName

`func (o *KubernetesInfo) SetKubernetesPullSecretName(v string)`

SetKubernetesPullSecretName sets KubernetesPullSecretName field to given value.

### HasKubernetesPullSecretName

`func (o *KubernetesInfo) HasKubernetesPullSecretName() bool`

HasKubernetesPullSecretName returns a boolean if a field has been set.

### GetKubernetesServiceAccount

`func (o *KubernetesInfo) GetKubernetesServiceAccount() string`

GetKubernetesServiceAccount returns the KubernetesServiceAccount field if non-nil, zero value otherwise.

### GetKubernetesServiceAccountOk

`func (o *KubernetesInfo) GetKubernetesServiceAccountOk() (*string, bool)`

GetKubernetesServiceAccountOk returns a tuple with the KubernetesServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesServiceAccount

`func (o *KubernetesInfo) SetKubernetesServiceAccount(v string)`

SetKubernetesServiceAccount sets KubernetesServiceAccount field to given value.

### HasKubernetesServiceAccount

`func (o *KubernetesInfo) HasKubernetesServiceAccount() bool`

HasKubernetesServiceAccount returns a boolean if a field has been set.

### GetKubernetesStorageClass

`func (o *KubernetesInfo) GetKubernetesStorageClass() string`

GetKubernetesStorageClass returns the KubernetesStorageClass field if non-nil, zero value otherwise.

### GetKubernetesStorageClassOk

`func (o *KubernetesInfo) GetKubernetesStorageClassOk() (*string, bool)`

GetKubernetesStorageClassOk returns a tuple with the KubernetesStorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesStorageClass

`func (o *KubernetesInfo) SetKubernetesStorageClass(v string)`

SetKubernetesStorageClass sets KubernetesStorageClass field to given value.

### HasKubernetesStorageClass

`func (o *KubernetesInfo) HasKubernetesStorageClass() bool`

HasKubernetesStorageClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


