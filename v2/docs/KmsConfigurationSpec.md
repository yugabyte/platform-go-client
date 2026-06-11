# KmsConfigurationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | User-defined name for the KMS configuration. | [optional] 
**Provider** | Pointer to **string** | KMS provider implementation key (AWS, GCP, etc.). | [optional] 
**Credentials** | **map[string]interface{}** | Masked provider auth configuration (opaque structure). Empty when the row has no auth configuration loaded (e.g. missing or unavailable credentials).  | 

## Methods

### NewKmsConfigurationSpec

`func NewKmsConfigurationSpec(credentials map[string]interface{}, ) *KmsConfigurationSpec`

NewKmsConfigurationSpec instantiates a new KmsConfigurationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsConfigurationSpecWithDefaults

`func NewKmsConfigurationSpecWithDefaults() *KmsConfigurationSpec`

NewKmsConfigurationSpecWithDefaults instantiates a new KmsConfigurationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KmsConfigurationSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsConfigurationSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsConfigurationSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KmsConfigurationSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *KmsConfigurationSpec) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *KmsConfigurationSpec) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *KmsConfigurationSpec) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *KmsConfigurationSpec) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetCredentials

`func (o *KmsConfigurationSpec) GetCredentials() map[string]interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *KmsConfigurationSpec) GetCredentialsOk() (*map[string]interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *KmsConfigurationSpec) SetCredentials(v map[string]interface{})`

SetCredentials sets Credentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


