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

// UniverseSpec These are user configured properties of a Universe. Part of create Universe request payload. Returned as part of a Universe resource.
type UniverseSpec struct {
	// Name of the Universe
	Name string `json:"name"`
	// The YugabyteDB software version to install. This can be upgraded using API /customers/:cUUID/universes/:uniUUID/upgrade/software
	YbSoftwareVersion string `json:"yb_software_version"`
	EncryptionAtRestSpec *EncryptionAtRestSpec `json:"encryption_at_rest_spec,omitempty"`
	EncryptionInTransitSpec *EncryptionInTransitSpec `json:"encryption_in_transit_spec,omitempty"`
	Ysql *YSQLSpec `json:"ysql,omitempty"`
	Ycql *YCQLSpec `json:"ycql,omitempty"`
	// Whether to use time sync services like chrony on DB nodes of this cluster
	UseTimeSync *bool `json:"use_time_sync,omitempty"`
	// Path to download thirdparty packages for itest. Only for AWS/onprem.
	RemotePackagePath *string `json:"remote_package_path,omitempty"`
	// Override the default DB present in pre-built Ami. YBM usage.
	OverridePrebuiltAmiDbVersion *bool `json:"override_prebuilt_ami_db_version,omitempty"`
	NetworkingSpec *UniverseNetworkingSpec `json:"networking_spec,omitempty"`
	Clusters []ClusterSpec `json:"clusters"`
}

// NewUniverseSpec instantiates a new UniverseSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniverseSpec(name string, ybSoftwareVersion string, clusters []ClusterSpec) *UniverseSpec {
	this := UniverseSpec{}
	this.Name = name
	this.YbSoftwareVersion = ybSoftwareVersion
	this.Clusters = clusters
	return &this
}

// NewUniverseSpecWithDefaults instantiates a new UniverseSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniverseSpecWithDefaults() *UniverseSpec {
	this := UniverseSpec{}
	return &this
}

// GetName returns the Name field value
func (o *UniverseSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UniverseSpec) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UniverseSpec) SetName(v string) {
	o.Name = v
}

// GetYbSoftwareVersion returns the YbSoftwareVersion field value
func (o *UniverseSpec) GetYbSoftwareVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.YbSoftwareVersion
}

// GetYbSoftwareVersionOk returns a tuple with the YbSoftwareVersion field value
// and a boolean to check if the value has been set.
func (o *UniverseSpec) GetYbSoftwareVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.YbSoftwareVersion, true
}

// SetYbSoftwareVersion sets field value
func (o *UniverseSpec) SetYbSoftwareVersion(v string) {
	o.YbSoftwareVersion = v
}

// GetEncryptionAtRestSpec returns the EncryptionAtRestSpec field value if set, zero value otherwise.
func (o *UniverseSpec) GetEncryptionAtRestSpec() EncryptionAtRestSpec {
	if o == nil || o.EncryptionAtRestSpec == nil {
		var ret EncryptionAtRestSpec
		return ret
	}
	return *o.EncryptionAtRestSpec
}

// GetEncryptionAtRestSpecOk returns a tuple with the EncryptionAtRestSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseSpec) GetEncryptionAtRestSpecOk() (*EncryptionAtRestSpec, bool) {
	if o == nil || o.EncryptionAtRestSpec == nil {
		return nil, false
	}
	return o.EncryptionAtRestSpec, true
}

// HasEncryptionAtRestSpec returns a boolean if a field has been set.
func (o *UniverseSpec) HasEncryptionAtRestSpec() bool {
	if o != nil && o.EncryptionAtRestSpec != nil {
		return true
	}

	return false
}

// SetEncryptionAtRestSpec gets a reference to the given EncryptionAtRestSpec and assigns it to the EncryptionAtRestSpec field.
func (o *UniverseSpec) SetEncryptionAtRestSpec(v EncryptionAtRestSpec) {
	o.EncryptionAtRestSpec = &v
}

// GetEncryptionInTransitSpec returns the EncryptionInTransitSpec field value if set, zero value otherwise.
func (o *UniverseSpec) GetEncryptionInTransitSpec() EncryptionInTransitSpec {
	if o == nil || o.EncryptionInTransitSpec == nil {
		var ret EncryptionInTransitSpec
		return ret
	}
	return *o.EncryptionInTransitSpec
}

// GetEncryptionInTransitSpecOk returns a tuple with the EncryptionInTransitSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseSpec) GetEncryptionInTransitSpecOk() (*EncryptionInTransitSpec, bool) {
	if o == nil || o.EncryptionInTransitSpec == nil {
		return nil, false
	}
	return o.EncryptionInTransitSpec, true
}

// HasEncryptionInTransitSpec returns a boolean if a field has been set.
func (o *UniverseSpec) HasEncryptionInTransitSpec() bool {
	if o != nil && o.EncryptionInTransitSpec != nil {
		return true
	}

	return false
}

// SetEncryptionInTransitSpec gets a reference to the given EncryptionInTransitSpec and assigns it to the EncryptionInTransitSpec field.
func (o *UniverseSpec) SetEncryptionInTransitSpec(v EncryptionInTransitSpec) {
	o.EncryptionInTransitSpec = &v
}

// GetYsql returns the Ysql field value if set, zero value otherwise.
func (o *UniverseSpec) GetYsql() YSQLSpec {
	if o == nil || o.Ysql == nil {
		var ret YSQLSpec
		return ret
	}
	return *o.Ysql
}

// GetYsqlOk returns a tuple with the Ysql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseSpec) GetYsqlOk() (*YSQLSpec, bool) {
	if o == nil || o.Ysql == nil {
		return nil, false
	}
	return o.Ysql, true
}

// HasYsql returns a boolean if a field has been set.
func (o *UniverseSpec) HasYsql() bool {
	if o != nil && o.Ysql != nil {
		return true
	}

	return false
}

// SetYsql gets a reference to the given YSQLSpec and assigns it to the Ysql field.
func (o *UniverseSpec) SetYsql(v YSQLSpec) {
	o.Ysql = &v
}

// GetYcql returns the Ycql field value if set, zero value otherwise.
func (o *UniverseSpec) GetYcql() YCQLSpec {
	if o == nil || o.Ycql == nil {
		var ret YCQLSpec
		return ret
	}
	return *o.Ycql
}

// GetYcqlOk returns a tuple with the Ycql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseSpec) GetYcqlOk() (*YCQLSpec, bool) {
	if o == nil || o.Ycql == nil {
		return nil, false
	}
	return o.Ycql, true
}

// HasYcql returns a boolean if a field has been set.
func (o *UniverseSpec) HasYcql() bool {
	if o != nil && o.Ycql != nil {
		return true
	}

	return false
}

// SetYcql gets a reference to the given YCQLSpec and assigns it to the Ycql field.
func (o *UniverseSpec) SetYcql(v YCQLSpec) {
	o.Ycql = &v
}

// GetUseTimeSync returns the UseTimeSync field value if set, zero value otherwise.
func (o *UniverseSpec) GetUseTimeSync() bool {
	if o == nil || o.UseTimeSync == nil {
		var ret bool
		return ret
	}
	return *o.UseTimeSync
}

// GetUseTimeSyncOk returns a tuple with the UseTimeSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseSpec) GetUseTimeSyncOk() (*bool, bool) {
	if o == nil || o.UseTimeSync == nil {
		return nil, false
	}
	return o.UseTimeSync, true
}

// HasUseTimeSync returns a boolean if a field has been set.
func (o *UniverseSpec) HasUseTimeSync() bool {
	if o != nil && o.UseTimeSync != nil {
		return true
	}

	return false
}

// SetUseTimeSync gets a reference to the given bool and assigns it to the UseTimeSync field.
func (o *UniverseSpec) SetUseTimeSync(v bool) {
	o.UseTimeSync = &v
}

// GetRemotePackagePath returns the RemotePackagePath field value if set, zero value otherwise.
func (o *UniverseSpec) GetRemotePackagePath() string {
	if o == nil || o.RemotePackagePath == nil {
		var ret string
		return ret
	}
	return *o.RemotePackagePath
}

// GetRemotePackagePathOk returns a tuple with the RemotePackagePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseSpec) GetRemotePackagePathOk() (*string, bool) {
	if o == nil || o.RemotePackagePath == nil {
		return nil, false
	}
	return o.RemotePackagePath, true
}

// HasRemotePackagePath returns a boolean if a field has been set.
func (o *UniverseSpec) HasRemotePackagePath() bool {
	if o != nil && o.RemotePackagePath != nil {
		return true
	}

	return false
}

// SetRemotePackagePath gets a reference to the given string and assigns it to the RemotePackagePath field.
func (o *UniverseSpec) SetRemotePackagePath(v string) {
	o.RemotePackagePath = &v
}

// GetOverridePrebuiltAmiDbVersion returns the OverridePrebuiltAmiDbVersion field value if set, zero value otherwise.
func (o *UniverseSpec) GetOverridePrebuiltAmiDbVersion() bool {
	if o == nil || o.OverridePrebuiltAmiDbVersion == nil {
		var ret bool
		return ret
	}
	return *o.OverridePrebuiltAmiDbVersion
}

// GetOverridePrebuiltAmiDbVersionOk returns a tuple with the OverridePrebuiltAmiDbVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseSpec) GetOverridePrebuiltAmiDbVersionOk() (*bool, bool) {
	if o == nil || o.OverridePrebuiltAmiDbVersion == nil {
		return nil, false
	}
	return o.OverridePrebuiltAmiDbVersion, true
}

// HasOverridePrebuiltAmiDbVersion returns a boolean if a field has been set.
func (o *UniverseSpec) HasOverridePrebuiltAmiDbVersion() bool {
	if o != nil && o.OverridePrebuiltAmiDbVersion != nil {
		return true
	}

	return false
}

// SetOverridePrebuiltAmiDbVersion gets a reference to the given bool and assigns it to the OverridePrebuiltAmiDbVersion field.
func (o *UniverseSpec) SetOverridePrebuiltAmiDbVersion(v bool) {
	o.OverridePrebuiltAmiDbVersion = &v
}

// GetNetworkingSpec returns the NetworkingSpec field value if set, zero value otherwise.
func (o *UniverseSpec) GetNetworkingSpec() UniverseNetworkingSpec {
	if o == nil || o.NetworkingSpec == nil {
		var ret UniverseNetworkingSpec
		return ret
	}
	return *o.NetworkingSpec
}

// GetNetworkingSpecOk returns a tuple with the NetworkingSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UniverseSpec) GetNetworkingSpecOk() (*UniverseNetworkingSpec, bool) {
	if o == nil || o.NetworkingSpec == nil {
		return nil, false
	}
	return o.NetworkingSpec, true
}

// HasNetworkingSpec returns a boolean if a field has been set.
func (o *UniverseSpec) HasNetworkingSpec() bool {
	if o != nil && o.NetworkingSpec != nil {
		return true
	}

	return false
}

// SetNetworkingSpec gets a reference to the given UniverseNetworkingSpec and assigns it to the NetworkingSpec field.
func (o *UniverseSpec) SetNetworkingSpec(v UniverseNetworkingSpec) {
	o.NetworkingSpec = &v
}

// GetClusters returns the Clusters field value
func (o *UniverseSpec) GetClusters() []ClusterSpec {
	if o == nil {
		var ret []ClusterSpec
		return ret
	}

	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value
// and a boolean to check if the value has been set.
func (o *UniverseSpec) GetClustersOk() (*[]ClusterSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Clusters, true
}

// SetClusters sets field value
func (o *UniverseSpec) SetClusters(v []ClusterSpec) {
	o.Clusters = v
}

func (o UniverseSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["yb_software_version"] = o.YbSoftwareVersion
	}
	if o.EncryptionAtRestSpec != nil {
		toSerialize["encryption_at_rest_spec"] = o.EncryptionAtRestSpec
	}
	if o.EncryptionInTransitSpec != nil {
		toSerialize["encryption_in_transit_spec"] = o.EncryptionInTransitSpec
	}
	if o.Ysql != nil {
		toSerialize["ysql"] = o.Ysql
	}
	if o.Ycql != nil {
		toSerialize["ycql"] = o.Ycql
	}
	if o.UseTimeSync != nil {
		toSerialize["use_time_sync"] = o.UseTimeSync
	}
	if o.RemotePackagePath != nil {
		toSerialize["remote_package_path"] = o.RemotePackagePath
	}
	if o.OverridePrebuiltAmiDbVersion != nil {
		toSerialize["override_prebuilt_ami_db_version"] = o.OverridePrebuiltAmiDbVersion
	}
	if o.NetworkingSpec != nil {
		toSerialize["networking_spec"] = o.NetworkingSpec
	}
	if true {
		toSerialize["clusters"] = o.Clusters
	}
	return json.Marshal(toSerialize)
}

type NullableUniverseSpec struct {
	value *UniverseSpec
	isSet bool
}

func (v NullableUniverseSpec) Get() *UniverseSpec {
	return v.value
}

func (v *NullableUniverseSpec) Set(val *UniverseSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableUniverseSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableUniverseSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniverseSpec(val *UniverseSpec) *NullableUniverseSpec {
	return &NullableUniverseSpec{value: val, isSet: true}
}

func (v NullableUniverseSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniverseSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


