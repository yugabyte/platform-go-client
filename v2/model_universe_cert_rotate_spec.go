/*
 * YugabyteDB Anywhere V2 APIs
 *
 * An improved set of APIs for managing YugabyteDB Anywhere
 *
 * API version: v2
 * Contact: support@yugabyte.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

// UniverseCertRotateSpec UniverseCertRotateSpec  Request payload to rotate the certs used for encryption in transit. Part of UniverseCertRotateReq. 
type UniverseCertRotateSpec struct {
	// Applicable for rolling restarts. Time to wait between master restarts. Defaults to 180000.
	SleepAfterMasterRestartMillis *int32 `json:"sleep_after_master_restart_millis,omitempty"`
	// Applicable for rolling restarts. Time to wait between tserver restarts. Defaults to 180000.
	SleepAfterTserverRestartMillis *int32 `json:"sleep_after_tserver_restart_millis,omitempty"`
	// Perform a rolling upgrade where only one node is upgraded at a time. This is the default behavior. False will perform a non-rolling upgrade where all nodes are upgraded at the same 
	RollingUpgrade *bool `json:"rolling_upgrade,omitempty"`
	RollMaxBatchSize *RollMaxBatchSize `json:"roll_max_batch_size,omitempty"`
	SelfSignedServerCertRotate *bool `json:"self_signed_server_cert_rotate,omitempty"`
	SelfSignedClientCertRotate *bool `json:"self_signed_client_cert_rotate,omitempty"`
	RootCa *string `json:"root_ca,omitempty"`
	ClientRootCa *string `json:"client_root_ca,omitempty"`
}

// NewUniverseCertRotateSpec instantiates a new UniverseCertRotateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniverseCertRotateSpec() *UniverseCertRotateSpec {
	this := UniverseCertRotateSpec{}
	var sleepAfterMasterRestartMillis int32 = 180000
	this.SleepAfterMasterRestartMillis = &sleepAfterMasterRestartMillis
	var sleepAfterTserverRestartMillis int32 = 180000
	this.SleepAfterTserverRestartMillis = &sleepAfterTserverRestartMillis
	var rollingUpgrade bool = true
	this.RollingUpgrade = &rollingUpgrade
	return &this
}

// NewUniverseCertRotateSpecWithDefaults instantiates a new UniverseCertRotateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniverseCertRotateSpecWithDefaults() *UniverseCertRotateSpec {
	this := UniverseCertRotateSpec{}
	var sleepAfterMasterRestartMillis int32 = 180000
	this.SleepAfterMasterRestartMillis = &sleepAfterMasterRestartMillis
	var sleepAfterTserverRestartMillis int32 = 180000
	this.SleepAfterTserverRestartMillis = &sleepAfterTserverRestartMillis
	var rollingUpgrade bool = true
	this.RollingUpgrade = &rollingUpgrade
	return &this
}

// GetSleepAfterMasterRestartMillis returns the SleepAfterMasterRestartMillis field value if set, zero value otherwise.
func (o *UniverseCertRotateSpec) GetSleepAfterMasterRestartMillis() int32 {
	if o == nil || o.SleepAfterMasterRestartMillis == nil {
		var ret int32
		return ret
	}
	return *o.SleepAfterMasterRestartMillis
}

// GetSleepAfterMasterRestartMillisOk returns a tuple with the SleepAfterMasterRestartMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseCertRotateSpec) GetSleepAfterMasterRestartMillisOk() (*int32, bool) {
	if o == nil || o.SleepAfterMasterRestartMillis == nil {
		return nil, false
	}
	return o.SleepAfterMasterRestartMillis, true
}

// HasSleepAfterMasterRestartMillis returns a boolean if a field has been set.
func (o *UniverseCertRotateSpec) HasSleepAfterMasterRestartMillis() bool {
	if o != nil && o.SleepAfterMasterRestartMillis != nil {
		return true
	}

	return false
}

// SetSleepAfterMasterRestartMillis gets a reference to the given int32 and assigns it to the SleepAfterMasterRestartMillis field.
func (o *UniverseCertRotateSpec) SetSleepAfterMasterRestartMillis(v int32) {
	o.SleepAfterMasterRestartMillis = &v
}

// GetSleepAfterTserverRestartMillis returns the SleepAfterTserverRestartMillis field value if set, zero value otherwise.
func (o *UniverseCertRotateSpec) GetSleepAfterTserverRestartMillis() int32 {
	if o == nil || o.SleepAfterTserverRestartMillis == nil {
		var ret int32
		return ret
	}
	return *o.SleepAfterTserverRestartMillis
}

// GetSleepAfterTserverRestartMillisOk returns a tuple with the SleepAfterTserverRestartMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseCertRotateSpec) GetSleepAfterTserverRestartMillisOk() (*int32, bool) {
	if o == nil || o.SleepAfterTserverRestartMillis == nil {
		return nil, false
	}
	return o.SleepAfterTserverRestartMillis, true
}

// HasSleepAfterTserverRestartMillis returns a boolean if a field has been set.
func (o *UniverseCertRotateSpec) HasSleepAfterTserverRestartMillis() bool {
	if o != nil && o.SleepAfterTserverRestartMillis != nil {
		return true
	}

	return false
}

// SetSleepAfterTserverRestartMillis gets a reference to the given int32 and assigns it to the SleepAfterTserverRestartMillis field.
func (o *UniverseCertRotateSpec) SetSleepAfterTserverRestartMillis(v int32) {
	o.SleepAfterTserverRestartMillis = &v
}

// GetRollingUpgrade returns the RollingUpgrade field value if set, zero value otherwise.
func (o *UniverseCertRotateSpec) GetRollingUpgrade() bool {
	if o == nil || o.RollingUpgrade == nil {
		var ret bool
		return ret
	}
	return *o.RollingUpgrade
}

// GetRollingUpgradeOk returns a tuple with the RollingUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseCertRotateSpec) GetRollingUpgradeOk() (*bool, bool) {
	if o == nil || o.RollingUpgrade == nil {
		return nil, false
	}
	return o.RollingUpgrade, true
}

// HasRollingUpgrade returns a boolean if a field has been set.
func (o *UniverseCertRotateSpec) HasRollingUpgrade() bool {
	if o != nil && o.RollingUpgrade != nil {
		return true
	}

	return false
}

// SetRollingUpgrade gets a reference to the given bool and assigns it to the RollingUpgrade field.
func (o *UniverseCertRotateSpec) SetRollingUpgrade(v bool) {
	o.RollingUpgrade = &v
}

// GetRollMaxBatchSize returns the RollMaxBatchSize field value if set, zero value otherwise.
func (o *UniverseCertRotateSpec) GetRollMaxBatchSize() RollMaxBatchSize {
	if o == nil || o.RollMaxBatchSize == nil {
		var ret RollMaxBatchSize
		return ret
	}
	return *o.RollMaxBatchSize
}

// GetRollMaxBatchSizeOk returns a tuple with the RollMaxBatchSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseCertRotateSpec) GetRollMaxBatchSizeOk() (*RollMaxBatchSize, bool) {
	if o == nil || o.RollMaxBatchSize == nil {
		return nil, false
	}
	return o.RollMaxBatchSize, true
}

// HasRollMaxBatchSize returns a boolean if a field has been set.
func (o *UniverseCertRotateSpec) HasRollMaxBatchSize() bool {
	if o != nil && o.RollMaxBatchSize != nil {
		return true
	}

	return false
}

// SetRollMaxBatchSize gets a reference to the given RollMaxBatchSize and assigns it to the RollMaxBatchSize field.
func (o *UniverseCertRotateSpec) SetRollMaxBatchSize(v RollMaxBatchSize) {
	o.RollMaxBatchSize = &v
}

// GetSelfSignedServerCertRotate returns the SelfSignedServerCertRotate field value if set, zero value otherwise.
func (o *UniverseCertRotateSpec) GetSelfSignedServerCertRotate() bool {
	if o == nil || o.SelfSignedServerCertRotate == nil {
		var ret bool
		return ret
	}
	return *o.SelfSignedServerCertRotate
}

// GetSelfSignedServerCertRotateOk returns a tuple with the SelfSignedServerCertRotate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseCertRotateSpec) GetSelfSignedServerCertRotateOk() (*bool, bool) {
	if o == nil || o.SelfSignedServerCertRotate == nil {
		return nil, false
	}
	return o.SelfSignedServerCertRotate, true
}

// HasSelfSignedServerCertRotate returns a boolean if a field has been set.
func (o *UniverseCertRotateSpec) HasSelfSignedServerCertRotate() bool {
	if o != nil && o.SelfSignedServerCertRotate != nil {
		return true
	}

	return false
}

// SetSelfSignedServerCertRotate gets a reference to the given bool and assigns it to the SelfSignedServerCertRotate field.
func (o *UniverseCertRotateSpec) SetSelfSignedServerCertRotate(v bool) {
	o.SelfSignedServerCertRotate = &v
}

// GetSelfSignedClientCertRotate returns the SelfSignedClientCertRotate field value if set, zero value otherwise.
func (o *UniverseCertRotateSpec) GetSelfSignedClientCertRotate() bool {
	if o == nil || o.SelfSignedClientCertRotate == nil {
		var ret bool
		return ret
	}
	return *o.SelfSignedClientCertRotate
}

// GetSelfSignedClientCertRotateOk returns a tuple with the SelfSignedClientCertRotate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseCertRotateSpec) GetSelfSignedClientCertRotateOk() (*bool, bool) {
	if o == nil || o.SelfSignedClientCertRotate == nil {
		return nil, false
	}
	return o.SelfSignedClientCertRotate, true
}

// HasSelfSignedClientCertRotate returns a boolean if a field has been set.
func (o *UniverseCertRotateSpec) HasSelfSignedClientCertRotate() bool {
	if o != nil && o.SelfSignedClientCertRotate != nil {
		return true
	}

	return false
}

// SetSelfSignedClientCertRotate gets a reference to the given bool and assigns it to the SelfSignedClientCertRotate field.
func (o *UniverseCertRotateSpec) SetSelfSignedClientCertRotate(v bool) {
	o.SelfSignedClientCertRotate = &v
}

// GetRootCa returns the RootCa field value if set, zero value otherwise.
func (o *UniverseCertRotateSpec) GetRootCa() string {
	if o == nil || o.RootCa == nil {
		var ret string
		return ret
	}
	return *o.RootCa
}

// GetRootCaOk returns a tuple with the RootCa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseCertRotateSpec) GetRootCaOk() (*string, bool) {
	if o == nil || o.RootCa == nil {
		return nil, false
	}
	return o.RootCa, true
}

// HasRootCa returns a boolean if a field has been set.
func (o *UniverseCertRotateSpec) HasRootCa() bool {
	if o != nil && o.RootCa != nil {
		return true
	}

	return false
}

// SetRootCa gets a reference to the given string and assigns it to the RootCa field.
func (o *UniverseCertRotateSpec) SetRootCa(v string) {
	o.RootCa = &v
}

// GetClientRootCa returns the ClientRootCa field value if set, zero value otherwise.
func (o *UniverseCertRotateSpec) GetClientRootCa() string {
	if o == nil || o.ClientRootCa == nil {
		var ret string
		return ret
	}
	return *o.ClientRootCa
}

// GetClientRootCaOk returns a tuple with the ClientRootCa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseCertRotateSpec) GetClientRootCaOk() (*string, bool) {
	if o == nil || o.ClientRootCa == nil {
		return nil, false
	}
	return o.ClientRootCa, true
}

// HasClientRootCa returns a boolean if a field has been set.
func (o *UniverseCertRotateSpec) HasClientRootCa() bool {
	if o != nil && o.ClientRootCa != nil {
		return true
	}

	return false
}

// SetClientRootCa gets a reference to the given string and assigns it to the ClientRootCa field.
func (o *UniverseCertRotateSpec) SetClientRootCa(v string) {
	o.ClientRootCa = &v
}

func (o UniverseCertRotateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SleepAfterMasterRestartMillis != nil {
		toSerialize["sleep_after_master_restart_millis"] = o.SleepAfterMasterRestartMillis
	}
	if o.SleepAfterTserverRestartMillis != nil {
		toSerialize["sleep_after_tserver_restart_millis"] = o.SleepAfterTserverRestartMillis
	}
	if o.RollingUpgrade != nil {
		toSerialize["rolling_upgrade"] = o.RollingUpgrade
	}
	if o.RollMaxBatchSize != nil {
		toSerialize["roll_max_batch_size"] = o.RollMaxBatchSize
	}
	if o.SelfSignedServerCertRotate != nil {
		toSerialize["self_signed_server_cert_rotate"] = o.SelfSignedServerCertRotate
	}
	if o.SelfSignedClientCertRotate != nil {
		toSerialize["self_signed_client_cert_rotate"] = o.SelfSignedClientCertRotate
	}
	if o.RootCa != nil {
		toSerialize["root_ca"] = o.RootCa
	}
	if o.ClientRootCa != nil {
		toSerialize["client_root_ca"] = o.ClientRootCa
	}
	return json.Marshal(toSerialize)
}

type NullableUniverseCertRotateSpec struct {
	value *UniverseCertRotateSpec
	isSet bool
}

func (v NullableUniverseCertRotateSpec) Get() *UniverseCertRotateSpec {
	return v.value
}

func (v *NullableUniverseCertRotateSpec) Set(val *UniverseCertRotateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableUniverseCertRotateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableUniverseCertRotateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniverseCertRotateSpec(val *UniverseCertRotateSpec) *NullableUniverseCertRotateSpec {
	return &NullableUniverseCertRotateSpec{value: val, isSet: true}
}

func (v NullableUniverseCertRotateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniverseCertRotateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


