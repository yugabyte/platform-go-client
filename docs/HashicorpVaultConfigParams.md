# HashicorpVaultConfigParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engine** | **string** |  | 
**MountPath** | **string** |  | 
**Role** | **string** |  | 
**Ttl** | Pointer to **int64** |  | [optional] 
**TtlExpiry** | Pointer to **int64** |  | [optional] 
**VaultAddr** | **string** |  | 
**VaultAuthNamespace** | Pointer to **string** |  | [optional] 
**VaultRoleID** | Pointer to **string** |  | [optional] 
**VaultSecretID** | Pointer to **string** |  | [optional] 
**VaultToken** | Pointer to **string** |  | [optional] 

## Methods

### NewHashicorpVaultConfigParams

`func NewHashicorpVaultConfigParams(engine string, mountPath string, role string, vaultAddr string, ) *HashicorpVaultConfigParams`

NewHashicorpVaultConfigParams instantiates a new HashicorpVaultConfigParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHashicorpVaultConfigParamsWithDefaults

`func NewHashicorpVaultConfigParamsWithDefaults() *HashicorpVaultConfigParams`

NewHashicorpVaultConfigParamsWithDefaults instantiates a new HashicorpVaultConfigParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngine

`func (o *HashicorpVaultConfigParams) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *HashicorpVaultConfigParams) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *HashicorpVaultConfigParams) SetEngine(v string)`

SetEngine sets Engine field to given value.


### GetMountPath

`func (o *HashicorpVaultConfigParams) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *HashicorpVaultConfigParams) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *HashicorpVaultConfigParams) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.


### GetRole

`func (o *HashicorpVaultConfigParams) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *HashicorpVaultConfigParams) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *HashicorpVaultConfigParams) SetRole(v string)`

SetRole sets Role field to given value.


### GetTtl

`func (o *HashicorpVaultConfigParams) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *HashicorpVaultConfigParams) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *HashicorpVaultConfigParams) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *HashicorpVaultConfigParams) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetTtlExpiry

`func (o *HashicorpVaultConfigParams) GetTtlExpiry() int64`

GetTtlExpiry returns the TtlExpiry field if non-nil, zero value otherwise.

### GetTtlExpiryOk

`func (o *HashicorpVaultConfigParams) GetTtlExpiryOk() (*int64, bool)`

GetTtlExpiryOk returns a tuple with the TtlExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlExpiry

`func (o *HashicorpVaultConfigParams) SetTtlExpiry(v int64)`

SetTtlExpiry sets TtlExpiry field to given value.

### HasTtlExpiry

`func (o *HashicorpVaultConfigParams) HasTtlExpiry() bool`

HasTtlExpiry returns a boolean if a field has been set.

### GetVaultAddr

`func (o *HashicorpVaultConfigParams) GetVaultAddr() string`

GetVaultAddr returns the VaultAddr field if non-nil, zero value otherwise.

### GetVaultAddrOk

`func (o *HashicorpVaultConfigParams) GetVaultAddrOk() (*string, bool)`

GetVaultAddrOk returns a tuple with the VaultAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAddr

`func (o *HashicorpVaultConfigParams) SetVaultAddr(v string)`

SetVaultAddr sets VaultAddr field to given value.


### GetVaultAuthNamespace

`func (o *HashicorpVaultConfigParams) GetVaultAuthNamespace() string`

GetVaultAuthNamespace returns the VaultAuthNamespace field if non-nil, zero value otherwise.

### GetVaultAuthNamespaceOk

`func (o *HashicorpVaultConfigParams) GetVaultAuthNamespaceOk() (*string, bool)`

GetVaultAuthNamespaceOk returns a tuple with the VaultAuthNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultAuthNamespace

`func (o *HashicorpVaultConfigParams) SetVaultAuthNamespace(v string)`

SetVaultAuthNamespace sets VaultAuthNamespace field to given value.

### HasVaultAuthNamespace

`func (o *HashicorpVaultConfigParams) HasVaultAuthNamespace() bool`

HasVaultAuthNamespace returns a boolean if a field has been set.

### GetVaultRoleID

`func (o *HashicorpVaultConfigParams) GetVaultRoleID() string`

GetVaultRoleID returns the VaultRoleID field if non-nil, zero value otherwise.

### GetVaultRoleIDOk

`func (o *HashicorpVaultConfigParams) GetVaultRoleIDOk() (*string, bool)`

GetVaultRoleIDOk returns a tuple with the VaultRoleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultRoleID

`func (o *HashicorpVaultConfigParams) SetVaultRoleID(v string)`

SetVaultRoleID sets VaultRoleID field to given value.

### HasVaultRoleID

`func (o *HashicorpVaultConfigParams) HasVaultRoleID() bool`

HasVaultRoleID returns a boolean if a field has been set.

### GetVaultSecretID

`func (o *HashicorpVaultConfigParams) GetVaultSecretID() string`

GetVaultSecretID returns the VaultSecretID field if non-nil, zero value otherwise.

### GetVaultSecretIDOk

`func (o *HashicorpVaultConfigParams) GetVaultSecretIDOk() (*string, bool)`

GetVaultSecretIDOk returns a tuple with the VaultSecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultSecretID

`func (o *HashicorpVaultConfigParams) SetVaultSecretID(v string)`

SetVaultSecretID sets VaultSecretID field to given value.

### HasVaultSecretID

`func (o *HashicorpVaultConfigParams) HasVaultSecretID() bool`

HasVaultSecretID returns a boolean if a field has been set.

### GetVaultToken

`func (o *HashicorpVaultConfigParams) GetVaultToken() string`

GetVaultToken returns the VaultToken field if non-nil, zero value otherwise.

### GetVaultTokenOk

`func (o *HashicorpVaultConfigParams) GetVaultTokenOk() (*string, bool)`

GetVaultTokenOk returns a tuple with the VaultToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultToken

`func (o *HashicorpVaultConfigParams) SetVaultToken(v string)`

SetVaultToken sets VaultToken field to given value.

### HasVaultToken

`func (o *HashicorpVaultConfigParams) HasVaultToken() bool`

HasVaultToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


