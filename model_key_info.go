/*
 * YugabyteDB Anywhere APIs
 *
 * ALPHA - NOT FOR EXTERNAL USE
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ywclient

import (
	"encoding/json"
)

// KeyInfo struct for KeyInfo
type KeyInfo struct {
	AirGapInstall *bool `json:"airGapInstall,omitempty"`
	DeleteRemote *bool `json:"deleteRemote,omitempty"`
	InstallNodeExporter *bool `json:"installNodeExporter,omitempty"`
	KeyPairName *string `json:"keyPairName,omitempty"`
	// Key Management state
	ManagementState *string `json:"managementState,omitempty"`
	NodeExporterPort *int32 `json:"nodeExporterPort,omitempty"`
	NodeExporterUser *string `json:"nodeExporterUser,omitempty"`
	NtpServers *[]string `json:"ntpServers,omitempty"`
	PasswordlessSudoAccess *bool `json:"passwordlessSudoAccess,omitempty"`
	PrivateKey *string `json:"privateKey,omitempty"`
	ProvisionInstanceScript *string `json:"provisionInstanceScript,omitempty"`
	PublicKey *string `json:"publicKey,omitempty"`
	SetUpChrony *bool `json:"setUpChrony,omitempty"`
	ShowSetUpChrony *bool `json:"showSetUpChrony,omitempty"`
	SkipKeyValidateAndUpload *bool `json:"skipKeyValidateAndUpload,omitempty"`
	SkipProvisioning *bool `json:"skipProvisioning,omitempty"`
	SshPort *int32 `json:"sshPort,omitempty"`
	SshPrivateKeyContent *string `json:"sshPrivateKeyContent,omitempty"`
	SshUser *string `json:"sshUser,omitempty"`
	VaultFile *string `json:"vaultFile,omitempty"`
	VaultPasswordFile *string `json:"vaultPasswordFile,omitempty"`
}

// NewKeyInfo instantiates a new KeyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyInfo() *KeyInfo {
	this := KeyInfo{}
	return &this
}

// NewKeyInfoWithDefaults instantiates a new KeyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyInfoWithDefaults() *KeyInfo {
	this := KeyInfo{}
	return &this
}

// GetAirGapInstall returns the AirGapInstall field value if set, zero value otherwise.
func (o *KeyInfo) GetAirGapInstall() bool {
	if o == nil || o.AirGapInstall == nil {
		var ret bool
		return ret
	}
	return *o.AirGapInstall
}

// GetAirGapInstallOk returns a tuple with the AirGapInstall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetAirGapInstallOk() (*bool, bool) {
	if o == nil || o.AirGapInstall == nil {
		return nil, false
	}
	return o.AirGapInstall, true
}

// HasAirGapInstall returns a boolean if a field has been set.
func (o *KeyInfo) HasAirGapInstall() bool {
	if o != nil && o.AirGapInstall != nil {
		return true
	}

	return false
}

// SetAirGapInstall gets a reference to the given bool and assigns it to the AirGapInstall field.
func (o *KeyInfo) SetAirGapInstall(v bool) {
	o.AirGapInstall = &v
}

// GetDeleteRemote returns the DeleteRemote field value if set, zero value otherwise.
func (o *KeyInfo) GetDeleteRemote() bool {
	if o == nil || o.DeleteRemote == nil {
		var ret bool
		return ret
	}
	return *o.DeleteRemote
}

// GetDeleteRemoteOk returns a tuple with the DeleteRemote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetDeleteRemoteOk() (*bool, bool) {
	if o == nil || o.DeleteRemote == nil {
		return nil, false
	}
	return o.DeleteRemote, true
}

// HasDeleteRemote returns a boolean if a field has been set.
func (o *KeyInfo) HasDeleteRemote() bool {
	if o != nil && o.DeleteRemote != nil {
		return true
	}

	return false
}

// SetDeleteRemote gets a reference to the given bool and assigns it to the DeleteRemote field.
func (o *KeyInfo) SetDeleteRemote(v bool) {
	o.DeleteRemote = &v
}

// GetInstallNodeExporter returns the InstallNodeExporter field value if set, zero value otherwise.
func (o *KeyInfo) GetInstallNodeExporter() bool {
	if o == nil || o.InstallNodeExporter == nil {
		var ret bool
		return ret
	}
	return *o.InstallNodeExporter
}

// GetInstallNodeExporterOk returns a tuple with the InstallNodeExporter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetInstallNodeExporterOk() (*bool, bool) {
	if o == nil || o.InstallNodeExporter == nil {
		return nil, false
	}
	return o.InstallNodeExporter, true
}

// HasInstallNodeExporter returns a boolean if a field has been set.
func (o *KeyInfo) HasInstallNodeExporter() bool {
	if o != nil && o.InstallNodeExporter != nil {
		return true
	}

	return false
}

// SetInstallNodeExporter gets a reference to the given bool and assigns it to the InstallNodeExporter field.
func (o *KeyInfo) SetInstallNodeExporter(v bool) {
	o.InstallNodeExporter = &v
}

// GetKeyPairName returns the KeyPairName field value if set, zero value otherwise.
func (o *KeyInfo) GetKeyPairName() string {
	if o == nil || o.KeyPairName == nil {
		var ret string
		return ret
	}
	return *o.KeyPairName
}

// GetKeyPairNameOk returns a tuple with the KeyPairName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetKeyPairNameOk() (*string, bool) {
	if o == nil || o.KeyPairName == nil {
		return nil, false
	}
	return o.KeyPairName, true
}

// HasKeyPairName returns a boolean if a field has been set.
func (o *KeyInfo) HasKeyPairName() bool {
	if o != nil && o.KeyPairName != nil {
		return true
	}

	return false
}

// SetKeyPairName gets a reference to the given string and assigns it to the KeyPairName field.
func (o *KeyInfo) SetKeyPairName(v string) {
	o.KeyPairName = &v
}

// GetManagementState returns the ManagementState field value if set, zero value otherwise.
func (o *KeyInfo) GetManagementState() string {
	if o == nil || o.ManagementState == nil {
		var ret string
		return ret
	}
	return *o.ManagementState
}

// GetManagementStateOk returns a tuple with the ManagementState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetManagementStateOk() (*string, bool) {
	if o == nil || o.ManagementState == nil {
		return nil, false
	}
	return o.ManagementState, true
}

// HasManagementState returns a boolean if a field has been set.
func (o *KeyInfo) HasManagementState() bool {
	if o != nil && o.ManagementState != nil {
		return true
	}

	return false
}

// SetManagementState gets a reference to the given string and assigns it to the ManagementState field.
func (o *KeyInfo) SetManagementState(v string) {
	o.ManagementState = &v
}

// GetNodeExporterPort returns the NodeExporterPort field value if set, zero value otherwise.
func (o *KeyInfo) GetNodeExporterPort() int32 {
	if o == nil || o.NodeExporterPort == nil {
		var ret int32
		return ret
	}
	return *o.NodeExporterPort
}

// GetNodeExporterPortOk returns a tuple with the NodeExporterPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetNodeExporterPortOk() (*int32, bool) {
	if o == nil || o.NodeExporterPort == nil {
		return nil, false
	}
	return o.NodeExporterPort, true
}

// HasNodeExporterPort returns a boolean if a field has been set.
func (o *KeyInfo) HasNodeExporterPort() bool {
	if o != nil && o.NodeExporterPort != nil {
		return true
	}

	return false
}

// SetNodeExporterPort gets a reference to the given int32 and assigns it to the NodeExporterPort field.
func (o *KeyInfo) SetNodeExporterPort(v int32) {
	o.NodeExporterPort = &v
}

// GetNodeExporterUser returns the NodeExporterUser field value if set, zero value otherwise.
func (o *KeyInfo) GetNodeExporterUser() string {
	if o == nil || o.NodeExporterUser == nil {
		var ret string
		return ret
	}
	return *o.NodeExporterUser
}

// GetNodeExporterUserOk returns a tuple with the NodeExporterUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetNodeExporterUserOk() (*string, bool) {
	if o == nil || o.NodeExporterUser == nil {
		return nil, false
	}
	return o.NodeExporterUser, true
}

// HasNodeExporterUser returns a boolean if a field has been set.
func (o *KeyInfo) HasNodeExporterUser() bool {
	if o != nil && o.NodeExporterUser != nil {
		return true
	}

	return false
}

// SetNodeExporterUser gets a reference to the given string and assigns it to the NodeExporterUser field.
func (o *KeyInfo) SetNodeExporterUser(v string) {
	o.NodeExporterUser = &v
}

// GetNtpServers returns the NtpServers field value if set, zero value otherwise.
func (o *KeyInfo) GetNtpServers() []string {
	if o == nil || o.NtpServers == nil {
		var ret []string
		return ret
	}
	return *o.NtpServers
}

// GetNtpServersOk returns a tuple with the NtpServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetNtpServersOk() (*[]string, bool) {
	if o == nil || o.NtpServers == nil {
		return nil, false
	}
	return o.NtpServers, true
}

// HasNtpServers returns a boolean if a field has been set.
func (o *KeyInfo) HasNtpServers() bool {
	if o != nil && o.NtpServers != nil {
		return true
	}

	return false
}

// SetNtpServers gets a reference to the given []string and assigns it to the NtpServers field.
func (o *KeyInfo) SetNtpServers(v []string) {
	o.NtpServers = &v
}

// GetPasswordlessSudoAccess returns the PasswordlessSudoAccess field value if set, zero value otherwise.
func (o *KeyInfo) GetPasswordlessSudoAccess() bool {
	if o == nil || o.PasswordlessSudoAccess == nil {
		var ret bool
		return ret
	}
	return *o.PasswordlessSudoAccess
}

// GetPasswordlessSudoAccessOk returns a tuple with the PasswordlessSudoAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetPasswordlessSudoAccessOk() (*bool, bool) {
	if o == nil || o.PasswordlessSudoAccess == nil {
		return nil, false
	}
	return o.PasswordlessSudoAccess, true
}

// HasPasswordlessSudoAccess returns a boolean if a field has been set.
func (o *KeyInfo) HasPasswordlessSudoAccess() bool {
	if o != nil && o.PasswordlessSudoAccess != nil {
		return true
	}

	return false
}

// SetPasswordlessSudoAccess gets a reference to the given bool and assigns it to the PasswordlessSudoAccess field.
func (o *KeyInfo) SetPasswordlessSudoAccess(v bool) {
	o.PasswordlessSudoAccess = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *KeyInfo) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *KeyInfo) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *KeyInfo) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetProvisionInstanceScript returns the ProvisionInstanceScript field value if set, zero value otherwise.
func (o *KeyInfo) GetProvisionInstanceScript() string {
	if o == nil || o.ProvisionInstanceScript == nil {
		var ret string
		return ret
	}
	return *o.ProvisionInstanceScript
}

// GetProvisionInstanceScriptOk returns a tuple with the ProvisionInstanceScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetProvisionInstanceScriptOk() (*string, bool) {
	if o == nil || o.ProvisionInstanceScript == nil {
		return nil, false
	}
	return o.ProvisionInstanceScript, true
}

// HasProvisionInstanceScript returns a boolean if a field has been set.
func (o *KeyInfo) HasProvisionInstanceScript() bool {
	if o != nil && o.ProvisionInstanceScript != nil {
		return true
	}

	return false
}

// SetProvisionInstanceScript gets a reference to the given string and assigns it to the ProvisionInstanceScript field.
func (o *KeyInfo) SetProvisionInstanceScript(v string) {
	o.ProvisionInstanceScript = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *KeyInfo) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *KeyInfo) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *KeyInfo) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetSetUpChrony returns the SetUpChrony field value if set, zero value otherwise.
func (o *KeyInfo) GetSetUpChrony() bool {
	if o == nil || o.SetUpChrony == nil {
		var ret bool
		return ret
	}
	return *o.SetUpChrony
}

// GetSetUpChronyOk returns a tuple with the SetUpChrony field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetSetUpChronyOk() (*bool, bool) {
	if o == nil || o.SetUpChrony == nil {
		return nil, false
	}
	return o.SetUpChrony, true
}

// HasSetUpChrony returns a boolean if a field has been set.
func (o *KeyInfo) HasSetUpChrony() bool {
	if o != nil && o.SetUpChrony != nil {
		return true
	}

	return false
}

// SetSetUpChrony gets a reference to the given bool and assigns it to the SetUpChrony field.
func (o *KeyInfo) SetSetUpChrony(v bool) {
	o.SetUpChrony = &v
}

// GetShowSetUpChrony returns the ShowSetUpChrony field value if set, zero value otherwise.
func (o *KeyInfo) GetShowSetUpChrony() bool {
	if o == nil || o.ShowSetUpChrony == nil {
		var ret bool
		return ret
	}
	return *o.ShowSetUpChrony
}

// GetShowSetUpChronyOk returns a tuple with the ShowSetUpChrony field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetShowSetUpChronyOk() (*bool, bool) {
	if o == nil || o.ShowSetUpChrony == nil {
		return nil, false
	}
	return o.ShowSetUpChrony, true
}

// HasShowSetUpChrony returns a boolean if a field has been set.
func (o *KeyInfo) HasShowSetUpChrony() bool {
	if o != nil && o.ShowSetUpChrony != nil {
		return true
	}

	return false
}

// SetShowSetUpChrony gets a reference to the given bool and assigns it to the ShowSetUpChrony field.
func (o *KeyInfo) SetShowSetUpChrony(v bool) {
	o.ShowSetUpChrony = &v
}

// GetSkipKeyValidateAndUpload returns the SkipKeyValidateAndUpload field value if set, zero value otherwise.
func (o *KeyInfo) GetSkipKeyValidateAndUpload() bool {
	if o == nil || o.SkipKeyValidateAndUpload == nil {
		var ret bool
		return ret
	}
	return *o.SkipKeyValidateAndUpload
}

// GetSkipKeyValidateAndUploadOk returns a tuple with the SkipKeyValidateAndUpload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetSkipKeyValidateAndUploadOk() (*bool, bool) {
	if o == nil || o.SkipKeyValidateAndUpload == nil {
		return nil, false
	}
	return o.SkipKeyValidateAndUpload, true
}

// HasSkipKeyValidateAndUpload returns a boolean if a field has been set.
func (o *KeyInfo) HasSkipKeyValidateAndUpload() bool {
	if o != nil && o.SkipKeyValidateAndUpload != nil {
		return true
	}

	return false
}

// SetSkipKeyValidateAndUpload gets a reference to the given bool and assigns it to the SkipKeyValidateAndUpload field.
func (o *KeyInfo) SetSkipKeyValidateAndUpload(v bool) {
	o.SkipKeyValidateAndUpload = &v
}

// GetSkipProvisioning returns the SkipProvisioning field value if set, zero value otherwise.
func (o *KeyInfo) GetSkipProvisioning() bool {
	if o == nil || o.SkipProvisioning == nil {
		var ret bool
		return ret
	}
	return *o.SkipProvisioning
}

// GetSkipProvisioningOk returns a tuple with the SkipProvisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetSkipProvisioningOk() (*bool, bool) {
	if o == nil || o.SkipProvisioning == nil {
		return nil, false
	}
	return o.SkipProvisioning, true
}

// HasSkipProvisioning returns a boolean if a field has been set.
func (o *KeyInfo) HasSkipProvisioning() bool {
	if o != nil && o.SkipProvisioning != nil {
		return true
	}

	return false
}

// SetSkipProvisioning gets a reference to the given bool and assigns it to the SkipProvisioning field.
func (o *KeyInfo) SetSkipProvisioning(v bool) {
	o.SkipProvisioning = &v
}

// GetSshPort returns the SshPort field value if set, zero value otherwise.
func (o *KeyInfo) GetSshPort() int32 {
	if o == nil || o.SshPort == nil {
		var ret int32
		return ret
	}
	return *o.SshPort
}

// GetSshPortOk returns a tuple with the SshPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetSshPortOk() (*int32, bool) {
	if o == nil || o.SshPort == nil {
		return nil, false
	}
	return o.SshPort, true
}

// HasSshPort returns a boolean if a field has been set.
func (o *KeyInfo) HasSshPort() bool {
	if o != nil && o.SshPort != nil {
		return true
	}

	return false
}

// SetSshPort gets a reference to the given int32 and assigns it to the SshPort field.
func (o *KeyInfo) SetSshPort(v int32) {
	o.SshPort = &v
}

// GetSshPrivateKeyContent returns the SshPrivateKeyContent field value if set, zero value otherwise.
func (o *KeyInfo) GetSshPrivateKeyContent() string {
	if o == nil || o.SshPrivateKeyContent == nil {
		var ret string
		return ret
	}
	return *o.SshPrivateKeyContent
}

// GetSshPrivateKeyContentOk returns a tuple with the SshPrivateKeyContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetSshPrivateKeyContentOk() (*string, bool) {
	if o == nil || o.SshPrivateKeyContent == nil {
		return nil, false
	}
	return o.SshPrivateKeyContent, true
}

// HasSshPrivateKeyContent returns a boolean if a field has been set.
func (o *KeyInfo) HasSshPrivateKeyContent() bool {
	if o != nil && o.SshPrivateKeyContent != nil {
		return true
	}

	return false
}

// SetSshPrivateKeyContent gets a reference to the given string and assigns it to the SshPrivateKeyContent field.
func (o *KeyInfo) SetSshPrivateKeyContent(v string) {
	o.SshPrivateKeyContent = &v
}

// GetSshUser returns the SshUser field value if set, zero value otherwise.
func (o *KeyInfo) GetSshUser() string {
	if o == nil || o.SshUser == nil {
		var ret string
		return ret
	}
	return *o.SshUser
}

// GetSshUserOk returns a tuple with the SshUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetSshUserOk() (*string, bool) {
	if o == nil || o.SshUser == nil {
		return nil, false
	}
	return o.SshUser, true
}

// HasSshUser returns a boolean if a field has been set.
func (o *KeyInfo) HasSshUser() bool {
	if o != nil && o.SshUser != nil {
		return true
	}

	return false
}

// SetSshUser gets a reference to the given string and assigns it to the SshUser field.
func (o *KeyInfo) SetSshUser(v string) {
	o.SshUser = &v
}

// GetVaultFile returns the VaultFile field value if set, zero value otherwise.
func (o *KeyInfo) GetVaultFile() string {
	if o == nil || o.VaultFile == nil {
		var ret string
		return ret
	}
	return *o.VaultFile
}

// GetVaultFileOk returns a tuple with the VaultFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetVaultFileOk() (*string, bool) {
	if o == nil || o.VaultFile == nil {
		return nil, false
	}
	return o.VaultFile, true
}

// HasVaultFile returns a boolean if a field has been set.
func (o *KeyInfo) HasVaultFile() bool {
	if o != nil && o.VaultFile != nil {
		return true
	}

	return false
}

// SetVaultFile gets a reference to the given string and assigns it to the VaultFile field.
func (o *KeyInfo) SetVaultFile(v string) {
	o.VaultFile = &v
}

// GetVaultPasswordFile returns the VaultPasswordFile field value if set, zero value otherwise.
func (o *KeyInfo) GetVaultPasswordFile() string {
	if o == nil || o.VaultPasswordFile == nil {
		var ret string
		return ret
	}
	return *o.VaultPasswordFile
}

// GetVaultPasswordFileOk returns a tuple with the VaultPasswordFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfo) GetVaultPasswordFileOk() (*string, bool) {
	if o == nil || o.VaultPasswordFile == nil {
		return nil, false
	}
	return o.VaultPasswordFile, true
}

// HasVaultPasswordFile returns a boolean if a field has been set.
func (o *KeyInfo) HasVaultPasswordFile() bool {
	if o != nil && o.VaultPasswordFile != nil {
		return true
	}

	return false
}

// SetVaultPasswordFile gets a reference to the given string and assigns it to the VaultPasswordFile field.
func (o *KeyInfo) SetVaultPasswordFile(v string) {
	o.VaultPasswordFile = &v
}

func (o KeyInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AirGapInstall != nil {
		toSerialize["airGapInstall"] = o.AirGapInstall
	}
	if o.DeleteRemote != nil {
		toSerialize["deleteRemote"] = o.DeleteRemote
	}
	if o.InstallNodeExporter != nil {
		toSerialize["installNodeExporter"] = o.InstallNodeExporter
	}
	if o.KeyPairName != nil {
		toSerialize["keyPairName"] = o.KeyPairName
	}
	if o.ManagementState != nil {
		toSerialize["managementState"] = o.ManagementState
	}
	if o.NodeExporterPort != nil {
		toSerialize["nodeExporterPort"] = o.NodeExporterPort
	}
	if o.NodeExporterUser != nil {
		toSerialize["nodeExporterUser"] = o.NodeExporterUser
	}
	if o.NtpServers != nil {
		toSerialize["ntpServers"] = o.NtpServers
	}
	if o.PasswordlessSudoAccess != nil {
		toSerialize["passwordlessSudoAccess"] = o.PasswordlessSudoAccess
	}
	if o.PrivateKey != nil {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if o.ProvisionInstanceScript != nil {
		toSerialize["provisionInstanceScript"] = o.ProvisionInstanceScript
	}
	if o.PublicKey != nil {
		toSerialize["publicKey"] = o.PublicKey
	}
	if o.SetUpChrony != nil {
		toSerialize["setUpChrony"] = o.SetUpChrony
	}
	if o.ShowSetUpChrony != nil {
		toSerialize["showSetUpChrony"] = o.ShowSetUpChrony
	}
	if o.SkipKeyValidateAndUpload != nil {
		toSerialize["skipKeyValidateAndUpload"] = o.SkipKeyValidateAndUpload
	}
	if o.SkipProvisioning != nil {
		toSerialize["skipProvisioning"] = o.SkipProvisioning
	}
	if o.SshPort != nil {
		toSerialize["sshPort"] = o.SshPort
	}
	if o.SshPrivateKeyContent != nil {
		toSerialize["sshPrivateKeyContent"] = o.SshPrivateKeyContent
	}
	if o.SshUser != nil {
		toSerialize["sshUser"] = o.SshUser
	}
	if o.VaultFile != nil {
		toSerialize["vaultFile"] = o.VaultFile
	}
	if o.VaultPasswordFile != nil {
		toSerialize["vaultPasswordFile"] = o.VaultPasswordFile
	}
	return json.Marshal(toSerialize)
}

type NullableKeyInfo struct {
	value *KeyInfo
	isSet bool
}

func (v NullableKeyInfo) Get() *KeyInfo {
	return v.value
}

func (v *NullableKeyInfo) Set(val *KeyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyInfo(val *KeyInfo) *NullableKeyInfo {
	return &NullableKeyInfo{value: val, isSet: true}
}

func (v NullableKeyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


