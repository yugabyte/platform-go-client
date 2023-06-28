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

// Provider struct for Provider
type Provider struct {
	// Provider active status
	Active *bool `json:"active,omitempty"`
	// Deprecated: sinceDate=2023-02-11, sinceYBAVersion=2.17.2.0, Use details.airGapInstall. Only supported in Create Request
	AirGapInstall *bool `json:"airGapInstall,omitempty"`
	AllAccessKeys *[]AccessKey `json:"allAccessKeys,omitempty"`
	// Provider cloud code
	Code *string `json:"code,omitempty"`
	// Deprecated: sinceDate=2023-02-11, sinceYBAVersion=2.17.2.0, Use details.metadata instead
	Config *map[string]string `json:"config,omitempty"`
	// Customer uuid
	CustomerUUID *string `json:"customerUUID,omitempty"`
	DestVpcId *string `json:"destVpcId,omitempty"`
	Details *ProviderDetails `json:"details,omitempty"`
	HostVpcId *string `json:"hostVpcId,omitempty"`
	HostVpcRegion *string `json:"hostVpcRegion,omitempty"`
	// Deprecated: sinceDate=2023-02-11, sinceYBAVersion=2.17.2.0, Use allAccessKeys[0].keyInfo.keyPairName instead
	KeyPairName *string `json:"keyPairName,omitempty"`
	// Provider name
	Name *string `json:"name,omitempty"`
	Regions []Region `json:"regions"`
	// Deprecated: sinceDate=2023-02-11, sinceYBAVersion=2.17.2.0, Use details.SshPort instead. Only supported in create request
	SshPort *int32 `json:"sshPort,omitempty"`
	// Deprecated: sinceDate=2023-02-11, sinceYBAVersion=2.17.2.0, Use allAccessKeys[0].keyInfo.sshPrivateKeyContent instead
	SshPrivateKeyContent *string `json:"sshPrivateKeyContent,omitempty"`
	// Deprecated: sinceDate=2023-02-11, sinceYBAVersion=2.17.2.0, Use details.SshUser instead. Only supported in create request
	SshUser *string `json:"sshUser,omitempty"`
	// Current usability state
	UsabilityState *string `json:"usabilityState,omitempty"`
	// Provider uuid
	Uuid *string `json:"uuid,omitempty"`
	// Provider version
	Version *int64 `json:"version,omitempty"`
}

// NewProvider instantiates a new Provider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvider(regions []Region, ) *Provider {
	this := Provider{}
	this.Regions = regions
	return &this
}

// NewProviderWithDefaults instantiates a new Provider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderWithDefaults() *Provider {
	this := Provider{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Provider) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provider) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Provider) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Provider) SetActive(v bool) {
	o.Active = &v
}

// GetAirGapInstall returns the AirGapInstall field value if set, zero value otherwise.
func (o *Provider) GetAirGapInstall() bool {
	if o == nil || o.AirGapInstall == nil {
		var ret bool
		return ret
	}
	return *o.AirGapInstall
}

// GetAirGapInstallOk returns a tuple with the AirGapInstall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provider) GetAirGapInstallOk() (*bool, bool) {
	if o == nil || o.AirGapInstall == nil {
		return nil, false
	}
	return o.AirGapInstall, true
}

// HasAirGapInstall returns a boolean if a field has been set.
func (o *Provider) HasAirGapInstall() bool {
	if o != nil && o.AirGapInstall != nil {
		return true
	}

	return false
}

// SetAirGapInstall gets a reference to the given bool and assigns it to the AirGapInstall field.
func (o *Provider) SetAirGapInstall(v bool) {
	o.AirGapInstall = &v
}

// GetAllAccessKeys returns the AllAccessKeys field value if set, zero value otherwise.
func (o *Provider) GetAllAccessKeys() []AccessKey {
	if o == nil || o.AllAccessKeys == nil {
		var ret []AccessKey
		return ret
	}
	return *o.AllAccessKeys
}

// GetAllAccessKeysOk returns a tuple with the AllAccessKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provider) GetAllAccessKeysOk() (*[]AccessKey, bool) {
	if o == nil || o.AllAccessKeys == nil {
		return nil, false
	}
	return o.AllAccessKeys, true
}

// HasAllAccessKeys returns a boolean if a field has been set.
func (o *Provider) HasAllAccessKeys() bool {
	if o != nil && o.AllAccessKeys != nil {
		return true
	}

	return false
}

// SetAllAccessKeys gets a reference to the given []AccessKey and assigns it to the AllAccessKeys field.
func (o *Provider) SetAllAccessKeys(v []AccessKey) {
	o.AllAccessKeys = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Provider) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provider) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Provider) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Provider) SetCode(v string) {
	o.Code = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Provider) GetConfig() map[string]string {
	if o == nil || o.Config == nil {
		var ret map[string]string
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provider) GetConfigOk() (*map[string]string, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Provider) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]string and assigns it to the Config field.
func (o *Provider) SetConfig(v map[string]string) {
	o.Config = &v
}

// GetCustomerUUID returns the CustomerUUID field value if set, zero value otherwise.
func (o *Provider) GetCustomerUUID() string {
	if o == nil || o.CustomerUUID == nil {
		var ret string
		return ret
	}
	return *o.CustomerUUID
}

// GetCustomerUUIDOk returns a tuple with the CustomerUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provider) GetCustomerUUIDOk() (*string, bool) {
	if o == nil || o.CustomerUUID == nil {
		return nil, false
	}
	return o.CustomerUUID, true
}

// HasCustomerUUID returns a boolean if a field has been set.
func (o *Provider) HasCustomerUUID() bool {
	if o != nil && o.CustomerUUID != nil {
		return true
	}

	return false
}

// SetCustomerUUID gets a reference to the given string and assigns it to the CustomerUUID field.
func (o *Provider) SetCustomerUUID(v string) {
	o.CustomerUUID = &v
}

// GetDestVpcId returns the DestVpcId field value if set, zero value otherwise.
func (o *Provider) GetDestVpcId() string {
	if o == nil || o.DestVpcId == nil {
		var ret string
		return ret
	}
	return *o.DestVpcId
}

// GetDestVpcIdOk returns a tuple with the DestVpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provider) GetDestVpcIdOk() (*string, bool) {
	if o == nil || o.DestVpcId == nil {
		return nil, false
	}
	return o.DestVpcId, true
}

// HasDestVpcId returns a boolean if a field has been set.
func (o *Provider) HasDestVpcId() bool {
	if o != nil && o.DestVpcId != nil {
		return true
	}

	return false
}

// SetDestVpcId gets a reference to the given string and assigns it to the DestVpcId field.
func (o *Provider) SetDestVpcId(v string) {
	o.DestVpcId = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *Provider) GetDetails() ProviderDetails {
	if o == nil || o.Details == nil {
		var ret ProviderDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provider) GetDetailsOk() (*ProviderDetails, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *Provider) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given ProviderDetails and assigns it to the Details field.
func (o *Provider) SetDetails(v ProviderDetails) {
	o.Details = &v
}

// GetHostVpcId returns the HostVpcId field value if set, zero value otherwise.
func (o *Provider) GetHostVpcId() string {
	if o == nil || o.HostVpcId == nil {
		var ret string
		return ret
	}
	return *o.HostVpcId
}

// GetHostVpcIdOk returns a tuple with the HostVpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provider) GetHostVpcIdOk() (*string, bool) {
	if o == nil || o.HostVpcId == nil {
		return nil, false
	}
	return o.HostVpcId, true
}

// HasHostVpcId returns a boolean if a field has been set.
func (o *Provider) HasHostVpcId() bool {
	if o != nil && o.HostVpcId != nil {
		return true
	}

	return false
}

// SetHostVpcId gets a reference to the given string and assigns it to the HostVpcId field.
func (o *Provider) SetHostVpcId(v string) {
	o.HostVpcId = &v
}

// GetHostVpcRegion returns the HostVpcRegion field value if set, zero value otherwise.
func (o *Provider) GetHostVpcRegion() string {
	if o == nil || o.HostVpcRegion == nil {
		var ret string
		return ret
	}
	return *o.HostVpcRegion
}

// GetHostVpcRegionOk returns a tuple with the HostVpcRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provider) GetHostVpcRegionOk() (*string, bool) {
	if o == nil || o.HostVpcRegion == nil {
		return nil, false
	}
	return o.HostVpcRegion, true
}

// HasHostVpcRegion returns a boolean if a field has been set.
func (o *Provider) HasHostVpcRegion() bool {
	if o != nil && o.HostVpcRegion != nil {
		return true
	}

	return false
}

// SetHostVpcRegion gets a reference to the given string and assigns it to the HostVpcRegion field.
func (o *Provider) SetHostVpcRegion(v string) {
	o.HostVpcRegion = &v
}

// GetKeyPairName returns the KeyPairName field value if set, zero value otherwise.
func (o *Provider) GetKeyPairName() string {
	if o == nil || o.KeyPairName == nil {
		var ret string
		return ret
	}
	return *o.KeyPairName
}

// GetKeyPairNameOk returns a tuple with the KeyPairName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provider) GetKeyPairNameOk() (*string, bool) {
	if o == nil || o.KeyPairName == nil {
		return nil, false
	}
	return o.KeyPairName, true
}

// HasKeyPairName returns a boolean if a field has been set.
func (o *Provider) HasKeyPairName() bool {
	if o != nil && o.KeyPairName != nil {
		return true
	}

	return false
}

// SetKeyPairName gets a reference to the given string and assigns it to the KeyPairName field.
func (o *Provider) SetKeyPairName(v string) {
	o.KeyPairName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Provider) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provider) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Provider) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Provider) SetName(v string) {
	o.Name = &v
}

// GetRegions returns the Regions field value
func (o *Provider) GetRegions() []Region {
	if o == nil  {
		var ret []Region
		return ret
	}

	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value
// and a boolean to check if the value has been set.
func (o *Provider) GetRegionsOk() (*[]Region, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Regions, true
}

// SetRegions sets field value
func (o *Provider) SetRegions(v []Region) {
	o.Regions = v
}

// GetSshPort returns the SshPort field value if set, zero value otherwise.
func (o *Provider) GetSshPort() int32 {
	if o == nil || o.SshPort == nil {
		var ret int32
		return ret
	}
	return *o.SshPort
}

// GetSshPortOk returns a tuple with the SshPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provider) GetSshPortOk() (*int32, bool) {
	if o == nil || o.SshPort == nil {
		return nil, false
	}
	return o.SshPort, true
}

// HasSshPort returns a boolean if a field has been set.
func (o *Provider) HasSshPort() bool {
	if o != nil && o.SshPort != nil {
		return true
	}

	return false
}

// SetSshPort gets a reference to the given int32 and assigns it to the SshPort field.
func (o *Provider) SetSshPort(v int32) {
	o.SshPort = &v
}

// GetSshPrivateKeyContent returns the SshPrivateKeyContent field value if set, zero value otherwise.
func (o *Provider) GetSshPrivateKeyContent() string {
	if o == nil || o.SshPrivateKeyContent == nil {
		var ret string
		return ret
	}
	return *o.SshPrivateKeyContent
}

// GetSshPrivateKeyContentOk returns a tuple with the SshPrivateKeyContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provider) GetSshPrivateKeyContentOk() (*string, bool) {
	if o == nil || o.SshPrivateKeyContent == nil {
		return nil, false
	}
	return o.SshPrivateKeyContent, true
}

// HasSshPrivateKeyContent returns a boolean if a field has been set.
func (o *Provider) HasSshPrivateKeyContent() bool {
	if o != nil && o.SshPrivateKeyContent != nil {
		return true
	}

	return false
}

// SetSshPrivateKeyContent gets a reference to the given string and assigns it to the SshPrivateKeyContent field.
func (o *Provider) SetSshPrivateKeyContent(v string) {
	o.SshPrivateKeyContent = &v
}

// GetSshUser returns the SshUser field value if set, zero value otherwise.
func (o *Provider) GetSshUser() string {
	if o == nil || o.SshUser == nil {
		var ret string
		return ret
	}
	return *o.SshUser
}

// GetSshUserOk returns a tuple with the SshUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provider) GetSshUserOk() (*string, bool) {
	if o == nil || o.SshUser == nil {
		return nil, false
	}
	return o.SshUser, true
}

// HasSshUser returns a boolean if a field has been set.
func (o *Provider) HasSshUser() bool {
	if o != nil && o.SshUser != nil {
		return true
	}

	return false
}

// SetSshUser gets a reference to the given string and assigns it to the SshUser field.
func (o *Provider) SetSshUser(v string) {
	o.SshUser = &v
}

// GetUsabilityState returns the UsabilityState field value if set, zero value otherwise.
func (o *Provider) GetUsabilityState() string {
	if o == nil || o.UsabilityState == nil {
		var ret string
		return ret
	}
	return *o.UsabilityState
}

// GetUsabilityStateOk returns a tuple with the UsabilityState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provider) GetUsabilityStateOk() (*string, bool) {
	if o == nil || o.UsabilityState == nil {
		return nil, false
	}
	return o.UsabilityState, true
}

// HasUsabilityState returns a boolean if a field has been set.
func (o *Provider) HasUsabilityState() bool {
	if o != nil && o.UsabilityState != nil {
		return true
	}

	return false
}

// SetUsabilityState gets a reference to the given string and assigns it to the UsabilityState field.
func (o *Provider) SetUsabilityState(v string) {
	o.UsabilityState = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Provider) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provider) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Provider) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Provider) SetUuid(v string) {
	o.Uuid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Provider) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provider) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Provider) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *Provider) SetVersion(v int64) {
	o.Version = &v
}

func (o Provider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.AirGapInstall != nil {
		toSerialize["airGapInstall"] = o.AirGapInstall
	}
	if o.AllAccessKeys != nil {
		toSerialize["allAccessKeys"] = o.AllAccessKeys
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.CustomerUUID != nil {
		toSerialize["customerUUID"] = o.CustomerUUID
	}
	if o.DestVpcId != nil {
		toSerialize["destVpcId"] = o.DestVpcId
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.HostVpcId != nil {
		toSerialize["hostVpcId"] = o.HostVpcId
	}
	if o.HostVpcRegion != nil {
		toSerialize["hostVpcRegion"] = o.HostVpcRegion
	}
	if o.KeyPairName != nil {
		toSerialize["keyPairName"] = o.KeyPairName
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["regions"] = o.Regions
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
	if o.UsabilityState != nil {
		toSerialize["usabilityState"] = o.UsabilityState
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableProvider struct {
	value *Provider
	isSet bool
}

func (v NullableProvider) Get() *Provider {
	return v.value
}

func (v *NullableProvider) Set(val *Provider) {
	v.value = val
	v.isSet = true
}

func (v NullableProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvider(val *Provider) *NullableProvider {
	return &NullableProvider{value: val, isSet: true}
}

func (v NullableProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


