# KeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AirGapInstall** | Pointer to **bool** |  | [optional] 
**DeleteRemote** | Pointer to **bool** |  | [optional] 
**InstallNodeExporter** | Pointer to **bool** |  | [optional] 
**KeyPairName** | Pointer to **string** |  | [optional] 
**ManagementState** | Pointer to **string** | Key Management state | [optional] [readonly] 
**NodeExporterPort** | Pointer to **int32** |  | [optional] 
**NodeExporterUser** | Pointer to **string** |  | [optional] 
**NtpServers** | Pointer to **[]string** |  | [optional] 
**PasswordlessSudoAccess** | Pointer to **bool** |  | [optional] 
**PrivateKey** | Pointer to **string** |  | [optional] 
**ProvisionInstanceScript** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**SetUpChrony** | Pointer to **bool** |  | [optional] 
**ShowSetUpChrony** | Pointer to **bool** |  | [optional] 
**SkipProvisioning** | Pointer to **bool** |  | [optional] 
**SshPort** | Pointer to **int32** |  | [optional] 
**SshPrivateKeyContent** | Pointer to **string** |  | [optional] 
**SshUser** | Pointer to **string** |  | [optional] 
**VaultFile** | Pointer to **string** |  | [optional] 
**VaultPasswordFile** | Pointer to **string** |  | [optional] 

## Methods

### NewKeyInfo

`func NewKeyInfo() *KeyInfo`

NewKeyInfo instantiates a new KeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyInfoWithDefaults

`func NewKeyInfoWithDefaults() *KeyInfo`

NewKeyInfoWithDefaults instantiates a new KeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAirGapInstall

`func (o *KeyInfo) GetAirGapInstall() bool`

GetAirGapInstall returns the AirGapInstall field if non-nil, zero value otherwise.

### GetAirGapInstallOk

`func (o *KeyInfo) GetAirGapInstallOk() (*bool, bool)`

GetAirGapInstallOk returns a tuple with the AirGapInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirGapInstall

`func (o *KeyInfo) SetAirGapInstall(v bool)`

SetAirGapInstall sets AirGapInstall field to given value.

### HasAirGapInstall

`func (o *KeyInfo) HasAirGapInstall() bool`

HasAirGapInstall returns a boolean if a field has been set.

### GetDeleteRemote

`func (o *KeyInfo) GetDeleteRemote() bool`

GetDeleteRemote returns the DeleteRemote field if non-nil, zero value otherwise.

### GetDeleteRemoteOk

`func (o *KeyInfo) GetDeleteRemoteOk() (*bool, bool)`

GetDeleteRemoteOk returns a tuple with the DeleteRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteRemote

`func (o *KeyInfo) SetDeleteRemote(v bool)`

SetDeleteRemote sets DeleteRemote field to given value.

### HasDeleteRemote

`func (o *KeyInfo) HasDeleteRemote() bool`

HasDeleteRemote returns a boolean if a field has been set.

### GetInstallNodeExporter

`func (o *KeyInfo) GetInstallNodeExporter() bool`

GetInstallNodeExporter returns the InstallNodeExporter field if non-nil, zero value otherwise.

### GetInstallNodeExporterOk

`func (o *KeyInfo) GetInstallNodeExporterOk() (*bool, bool)`

GetInstallNodeExporterOk returns a tuple with the InstallNodeExporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallNodeExporter

`func (o *KeyInfo) SetInstallNodeExporter(v bool)`

SetInstallNodeExporter sets InstallNodeExporter field to given value.

### HasInstallNodeExporter

`func (o *KeyInfo) HasInstallNodeExporter() bool`

HasInstallNodeExporter returns a boolean if a field has been set.

### GetKeyPairName

`func (o *KeyInfo) GetKeyPairName() string`

GetKeyPairName returns the KeyPairName field if non-nil, zero value otherwise.

### GetKeyPairNameOk

`func (o *KeyInfo) GetKeyPairNameOk() (*string, bool)`

GetKeyPairNameOk returns a tuple with the KeyPairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPairName

`func (o *KeyInfo) SetKeyPairName(v string)`

SetKeyPairName sets KeyPairName field to given value.

### HasKeyPairName

`func (o *KeyInfo) HasKeyPairName() bool`

HasKeyPairName returns a boolean if a field has been set.

### GetManagementState

`func (o *KeyInfo) GetManagementState() string`

GetManagementState returns the ManagementState field if non-nil, zero value otherwise.

### GetManagementStateOk

`func (o *KeyInfo) GetManagementStateOk() (*string, bool)`

GetManagementStateOk returns a tuple with the ManagementState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementState

`func (o *KeyInfo) SetManagementState(v string)`

SetManagementState sets ManagementState field to given value.

### HasManagementState

`func (o *KeyInfo) HasManagementState() bool`

HasManagementState returns a boolean if a field has been set.

### GetNodeExporterPort

`func (o *KeyInfo) GetNodeExporterPort() int32`

GetNodeExporterPort returns the NodeExporterPort field if non-nil, zero value otherwise.

### GetNodeExporterPortOk

`func (o *KeyInfo) GetNodeExporterPortOk() (*int32, bool)`

GetNodeExporterPortOk returns a tuple with the NodeExporterPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExporterPort

`func (o *KeyInfo) SetNodeExporterPort(v int32)`

SetNodeExporterPort sets NodeExporterPort field to given value.

### HasNodeExporterPort

`func (o *KeyInfo) HasNodeExporterPort() bool`

HasNodeExporterPort returns a boolean if a field has been set.

### GetNodeExporterUser

`func (o *KeyInfo) GetNodeExporterUser() string`

GetNodeExporterUser returns the NodeExporterUser field if non-nil, zero value otherwise.

### GetNodeExporterUserOk

`func (o *KeyInfo) GetNodeExporterUserOk() (*string, bool)`

GetNodeExporterUserOk returns a tuple with the NodeExporterUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExporterUser

`func (o *KeyInfo) SetNodeExporterUser(v string)`

SetNodeExporterUser sets NodeExporterUser field to given value.

### HasNodeExporterUser

`func (o *KeyInfo) HasNodeExporterUser() bool`

HasNodeExporterUser returns a boolean if a field has been set.

### GetNtpServers

`func (o *KeyInfo) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *KeyInfo) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *KeyInfo) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *KeyInfo) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetPasswordlessSudoAccess

`func (o *KeyInfo) GetPasswordlessSudoAccess() bool`

GetPasswordlessSudoAccess returns the PasswordlessSudoAccess field if non-nil, zero value otherwise.

### GetPasswordlessSudoAccessOk

`func (o *KeyInfo) GetPasswordlessSudoAccessOk() (*bool, bool)`

GetPasswordlessSudoAccessOk returns a tuple with the PasswordlessSudoAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordlessSudoAccess

`func (o *KeyInfo) SetPasswordlessSudoAccess(v bool)`

SetPasswordlessSudoAccess sets PasswordlessSudoAccess field to given value.

### HasPasswordlessSudoAccess

`func (o *KeyInfo) HasPasswordlessSudoAccess() bool`

HasPasswordlessSudoAccess returns a boolean if a field has been set.

### GetPrivateKey

`func (o *KeyInfo) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *KeyInfo) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *KeyInfo) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *KeyInfo) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetProvisionInstanceScript

`func (o *KeyInfo) GetProvisionInstanceScript() string`

GetProvisionInstanceScript returns the ProvisionInstanceScript field if non-nil, zero value otherwise.

### GetProvisionInstanceScriptOk

`func (o *KeyInfo) GetProvisionInstanceScriptOk() (*string, bool)`

GetProvisionInstanceScriptOk returns a tuple with the ProvisionInstanceScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionInstanceScript

`func (o *KeyInfo) SetProvisionInstanceScript(v string)`

SetProvisionInstanceScript sets ProvisionInstanceScript field to given value.

### HasProvisionInstanceScript

`func (o *KeyInfo) HasProvisionInstanceScript() bool`

HasProvisionInstanceScript returns a boolean if a field has been set.

### GetPublicKey

`func (o *KeyInfo) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *KeyInfo) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *KeyInfo) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *KeyInfo) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetSetUpChrony

`func (o *KeyInfo) GetSetUpChrony() bool`

GetSetUpChrony returns the SetUpChrony field if non-nil, zero value otherwise.

### GetSetUpChronyOk

`func (o *KeyInfo) GetSetUpChronyOk() (*bool, bool)`

GetSetUpChronyOk returns a tuple with the SetUpChrony field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetUpChrony

`func (o *KeyInfo) SetSetUpChrony(v bool)`

SetSetUpChrony sets SetUpChrony field to given value.

### HasSetUpChrony

`func (o *KeyInfo) HasSetUpChrony() bool`

HasSetUpChrony returns a boolean if a field has been set.

### GetShowSetUpChrony

`func (o *KeyInfo) GetShowSetUpChrony() bool`

GetShowSetUpChrony returns the ShowSetUpChrony field if non-nil, zero value otherwise.

### GetShowSetUpChronyOk

`func (o *KeyInfo) GetShowSetUpChronyOk() (*bool, bool)`

GetShowSetUpChronyOk returns a tuple with the ShowSetUpChrony field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSetUpChrony

`func (o *KeyInfo) SetShowSetUpChrony(v bool)`

SetShowSetUpChrony sets ShowSetUpChrony field to given value.

### HasShowSetUpChrony

`func (o *KeyInfo) HasShowSetUpChrony() bool`

HasShowSetUpChrony returns a boolean if a field has been set.

### GetSkipProvisioning

`func (o *KeyInfo) GetSkipProvisioning() bool`

GetSkipProvisioning returns the SkipProvisioning field if non-nil, zero value otherwise.

### GetSkipProvisioningOk

`func (o *KeyInfo) GetSkipProvisioningOk() (*bool, bool)`

GetSkipProvisioningOk returns a tuple with the SkipProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipProvisioning

`func (o *KeyInfo) SetSkipProvisioning(v bool)`

SetSkipProvisioning sets SkipProvisioning field to given value.

### HasSkipProvisioning

`func (o *KeyInfo) HasSkipProvisioning() bool`

HasSkipProvisioning returns a boolean if a field has been set.

### GetSshPort

`func (o *KeyInfo) GetSshPort() int32`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *KeyInfo) GetSshPortOk() (*int32, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *KeyInfo) SetSshPort(v int32)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *KeyInfo) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetSshPrivateKeyContent

`func (o *KeyInfo) GetSshPrivateKeyContent() string`

GetSshPrivateKeyContent returns the SshPrivateKeyContent field if non-nil, zero value otherwise.

### GetSshPrivateKeyContentOk

`func (o *KeyInfo) GetSshPrivateKeyContentOk() (*string, bool)`

GetSshPrivateKeyContentOk returns a tuple with the SshPrivateKeyContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeyContent

`func (o *KeyInfo) SetSshPrivateKeyContent(v string)`

SetSshPrivateKeyContent sets SshPrivateKeyContent field to given value.

### HasSshPrivateKeyContent

`func (o *KeyInfo) HasSshPrivateKeyContent() bool`

HasSshPrivateKeyContent returns a boolean if a field has been set.

### GetSshUser

`func (o *KeyInfo) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *KeyInfo) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *KeyInfo) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.

### HasSshUser

`func (o *KeyInfo) HasSshUser() bool`

HasSshUser returns a boolean if a field has been set.

### GetVaultFile

`func (o *KeyInfo) GetVaultFile() string`

GetVaultFile returns the VaultFile field if non-nil, zero value otherwise.

### GetVaultFileOk

`func (o *KeyInfo) GetVaultFileOk() (*string, bool)`

GetVaultFileOk returns a tuple with the VaultFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultFile

`func (o *KeyInfo) SetVaultFile(v string)`

SetVaultFile sets VaultFile field to given value.

### HasVaultFile

`func (o *KeyInfo) HasVaultFile() bool`

HasVaultFile returns a boolean if a field has been set.

### GetVaultPasswordFile

`func (o *KeyInfo) GetVaultPasswordFile() string`

GetVaultPasswordFile returns the VaultPasswordFile field if non-nil, zero value otherwise.

### GetVaultPasswordFileOk

`func (o *KeyInfo) GetVaultPasswordFileOk() (*string, bool)`

GetVaultPasswordFileOk returns a tuple with the VaultPasswordFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultPasswordFile

`func (o *KeyInfo) SetVaultPasswordFile(v string)`

SetVaultPasswordFile sets VaultPasswordFile field to given value.

### HasVaultPasswordFile

`func (o *KeyInfo) HasVaultPasswordFile() bool`

HasVaultPasswordFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


