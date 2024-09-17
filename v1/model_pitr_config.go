/*
 * YugabyteDB Anywhere APIs
 *
 * ALPHA - NOT FOR EXTERNAL USE
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"time"
)

// PitrConfig PITR config created on the universe
type PitrConfig struct {
	// Create time of the PITR config
	CreateTime *time.Time `json:"createTime,omitempty"`
	// Created for DR
	CreatedForDr *bool `json:"createdForDr,omitempty"`
	// Customer UUID of this config
	CustomerUUID *string `json:"customerUUID,omitempty"`
	// DB Name
	DbName *string `json:"dbName,omitempty"`
	MaxRecoverTimeInMillis int64 `json:"maxRecoverTimeInMillis"`
	MinRecoverTimeInMillis int64 `json:"minRecoverTimeInMillis"`
	// PITR config name
	Name *string `json:"name,omitempty"`
	// Retention Period in seconds
	RetentionPeriod *int64 `json:"retentionPeriod,omitempty"`
	// Interval between snasphots in seconds
	ScheduleInterval *int64 `json:"scheduleInterval,omitempty"`
	State string `json:"state"`
	// Table Type
	TableType *string `json:"tableType,omitempty"`
	// Update time of the PITR con
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	UsedForXCluster *bool `json:"usedForXCluster,omitempty"`
	// PITR config UUID
	Uuid *string `json:"uuid,omitempty"`
}

// NewPitrConfig instantiates a new PitrConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPitrConfig(maxRecoverTimeInMillis int64, minRecoverTimeInMillis int64, state string) *PitrConfig {
	this := PitrConfig{}
	this.MaxRecoverTimeInMillis = maxRecoverTimeInMillis
	this.MinRecoverTimeInMillis = minRecoverTimeInMillis
	this.State = state
	return &this
}

// NewPitrConfigWithDefaults instantiates a new PitrConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPitrConfigWithDefaults() *PitrConfig {
	this := PitrConfig{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *PitrConfig) GetCreateTime() time.Time {
	if o == nil || o.CreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PitrConfig) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *PitrConfig) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *PitrConfig) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetCreatedForDr returns the CreatedForDr field value if set, zero value otherwise.
func (o *PitrConfig) GetCreatedForDr() bool {
	if o == nil || o.CreatedForDr == nil {
		var ret bool
		return ret
	}
	return *o.CreatedForDr
}

// GetCreatedForDrOk returns a tuple with the CreatedForDr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PitrConfig) GetCreatedForDrOk() (*bool, bool) {
	if o == nil || o.CreatedForDr == nil {
		return nil, false
	}
	return o.CreatedForDr, true
}

// HasCreatedForDr returns a boolean if a field has been set.
func (o *PitrConfig) HasCreatedForDr() bool {
	if o != nil && o.CreatedForDr != nil {
		return true
	}

	return false
}

// SetCreatedForDr gets a reference to the given bool and assigns it to the CreatedForDr field.
func (o *PitrConfig) SetCreatedForDr(v bool) {
	o.CreatedForDr = &v
}

// GetCustomerUUID returns the CustomerUUID field value if set, zero value otherwise.
func (o *PitrConfig) GetCustomerUUID() string {
	if o == nil || o.CustomerUUID == nil {
		var ret string
		return ret
	}
	return *o.CustomerUUID
}

// GetCustomerUUIDOk returns a tuple with the CustomerUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PitrConfig) GetCustomerUUIDOk() (*string, bool) {
	if o == nil || o.CustomerUUID == nil {
		return nil, false
	}
	return o.CustomerUUID, true
}

// HasCustomerUUID returns a boolean if a field has been set.
func (o *PitrConfig) HasCustomerUUID() bool {
	if o != nil && o.CustomerUUID != nil {
		return true
	}

	return false
}

// SetCustomerUUID gets a reference to the given string and assigns it to the CustomerUUID field.
func (o *PitrConfig) SetCustomerUUID(v string) {
	o.CustomerUUID = &v
}

// GetDbName returns the DbName field value if set, zero value otherwise.
func (o *PitrConfig) GetDbName() string {
	if o == nil || o.DbName == nil {
		var ret string
		return ret
	}
	return *o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PitrConfig) GetDbNameOk() (*string, bool) {
	if o == nil || o.DbName == nil {
		return nil, false
	}
	return o.DbName, true
}

// HasDbName returns a boolean if a field has been set.
func (o *PitrConfig) HasDbName() bool {
	if o != nil && o.DbName != nil {
		return true
	}

	return false
}

// SetDbName gets a reference to the given string and assigns it to the DbName field.
func (o *PitrConfig) SetDbName(v string) {
	o.DbName = &v
}

// GetMaxRecoverTimeInMillis returns the MaxRecoverTimeInMillis field value
func (o *PitrConfig) GetMaxRecoverTimeInMillis() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxRecoverTimeInMillis
}

// GetMaxRecoverTimeInMillisOk returns a tuple with the MaxRecoverTimeInMillis field value
// and a boolean to check if the value has been set.
func (o *PitrConfig) GetMaxRecoverTimeInMillisOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MaxRecoverTimeInMillis, true
}

// SetMaxRecoverTimeInMillis sets field value
func (o *PitrConfig) SetMaxRecoverTimeInMillis(v int64) {
	o.MaxRecoverTimeInMillis = v
}

// GetMinRecoverTimeInMillis returns the MinRecoverTimeInMillis field value
func (o *PitrConfig) GetMinRecoverTimeInMillis() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MinRecoverTimeInMillis
}

// GetMinRecoverTimeInMillisOk returns a tuple with the MinRecoverTimeInMillis field value
// and a boolean to check if the value has been set.
func (o *PitrConfig) GetMinRecoverTimeInMillisOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MinRecoverTimeInMillis, true
}

// SetMinRecoverTimeInMillis sets field value
func (o *PitrConfig) SetMinRecoverTimeInMillis(v int64) {
	o.MinRecoverTimeInMillis = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PitrConfig) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PitrConfig) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PitrConfig) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PitrConfig) SetName(v string) {
	o.Name = &v
}

// GetRetentionPeriod returns the RetentionPeriod field value if set, zero value otherwise.
func (o *PitrConfig) GetRetentionPeriod() int64 {
	if o == nil || o.RetentionPeriod == nil {
		var ret int64
		return ret
	}
	return *o.RetentionPeriod
}

// GetRetentionPeriodOk returns a tuple with the RetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PitrConfig) GetRetentionPeriodOk() (*int64, bool) {
	if o == nil || o.RetentionPeriod == nil {
		return nil, false
	}
	return o.RetentionPeriod, true
}

// HasRetentionPeriod returns a boolean if a field has been set.
func (o *PitrConfig) HasRetentionPeriod() bool {
	if o != nil && o.RetentionPeriod != nil {
		return true
	}

	return false
}

// SetRetentionPeriod gets a reference to the given int64 and assigns it to the RetentionPeriod field.
func (o *PitrConfig) SetRetentionPeriod(v int64) {
	o.RetentionPeriod = &v
}

// GetScheduleInterval returns the ScheduleInterval field value if set, zero value otherwise.
func (o *PitrConfig) GetScheduleInterval() int64 {
	if o == nil || o.ScheduleInterval == nil {
		var ret int64
		return ret
	}
	return *o.ScheduleInterval
}

// GetScheduleIntervalOk returns a tuple with the ScheduleInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PitrConfig) GetScheduleIntervalOk() (*int64, bool) {
	if o == nil || o.ScheduleInterval == nil {
		return nil, false
	}
	return o.ScheduleInterval, true
}

// HasScheduleInterval returns a boolean if a field has been set.
func (o *PitrConfig) HasScheduleInterval() bool {
	if o != nil && o.ScheduleInterval != nil {
		return true
	}

	return false
}

// SetScheduleInterval gets a reference to the given int64 and assigns it to the ScheduleInterval field.
func (o *PitrConfig) SetScheduleInterval(v int64) {
	o.ScheduleInterval = &v
}

// GetState returns the State field value
func (o *PitrConfig) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *PitrConfig) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *PitrConfig) SetState(v string) {
	o.State = v
}

// GetTableType returns the TableType field value if set, zero value otherwise.
func (o *PitrConfig) GetTableType() string {
	if o == nil || o.TableType == nil {
		var ret string
		return ret
	}
	return *o.TableType
}

// GetTableTypeOk returns a tuple with the TableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PitrConfig) GetTableTypeOk() (*string, bool) {
	if o == nil || o.TableType == nil {
		return nil, false
	}
	return o.TableType, true
}

// HasTableType returns a boolean if a field has been set.
func (o *PitrConfig) HasTableType() bool {
	if o != nil && o.TableType != nil {
		return true
	}

	return false
}

// SetTableType gets a reference to the given string and assigns it to the TableType field.
func (o *PitrConfig) SetTableType(v string) {
	o.TableType = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *PitrConfig) GetUpdateTime() time.Time {
	if o == nil || o.UpdateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PitrConfig) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || o.UpdateTime == nil {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *PitrConfig) HasUpdateTime() bool {
	if o != nil && o.UpdateTime != nil {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *PitrConfig) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

// GetUsedForXCluster returns the UsedForXCluster field value if set, zero value otherwise.
func (o *PitrConfig) GetUsedForXCluster() bool {
	if o == nil || o.UsedForXCluster == nil {
		var ret bool
		return ret
	}
	return *o.UsedForXCluster
}

// GetUsedForXClusterOk returns a tuple with the UsedForXCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PitrConfig) GetUsedForXClusterOk() (*bool, bool) {
	if o == nil || o.UsedForXCluster == nil {
		return nil, false
	}
	return o.UsedForXCluster, true
}

// HasUsedForXCluster returns a boolean if a field has been set.
func (o *PitrConfig) HasUsedForXCluster() bool {
	if o != nil && o.UsedForXCluster != nil {
		return true
	}

	return false
}

// SetUsedForXCluster gets a reference to the given bool and assigns it to the UsedForXCluster field.
func (o *PitrConfig) SetUsedForXCluster(v bool) {
	o.UsedForXCluster = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *PitrConfig) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PitrConfig) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *PitrConfig) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *PitrConfig) SetUuid(v string) {
	o.Uuid = &v
}

func (o PitrConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreateTime != nil {
		toSerialize["createTime"] = o.CreateTime
	}
	if o.CreatedForDr != nil {
		toSerialize["createdForDr"] = o.CreatedForDr
	}
	if o.CustomerUUID != nil {
		toSerialize["customerUUID"] = o.CustomerUUID
	}
	if o.DbName != nil {
		toSerialize["dbName"] = o.DbName
	}
	if true {
		toSerialize["maxRecoverTimeInMillis"] = o.MaxRecoverTimeInMillis
	}
	if true {
		toSerialize["minRecoverTimeInMillis"] = o.MinRecoverTimeInMillis
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RetentionPeriod != nil {
		toSerialize["retentionPeriod"] = o.RetentionPeriod
	}
	if o.ScheduleInterval != nil {
		toSerialize["scheduleInterval"] = o.ScheduleInterval
	}
	if true {
		toSerialize["state"] = o.State
	}
	if o.TableType != nil {
		toSerialize["tableType"] = o.TableType
	}
	if o.UpdateTime != nil {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if o.UsedForXCluster != nil {
		toSerialize["usedForXCluster"] = o.UsedForXCluster
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullablePitrConfig struct {
	value *PitrConfig
	isSet bool
}

func (v NullablePitrConfig) Get() *PitrConfig {
	return v.value
}

func (v *NullablePitrConfig) Set(val *PitrConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePitrConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePitrConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePitrConfig(val *PitrConfig) *NullablePitrConfig {
	return &NullablePitrConfig{value: val, isSet: true}
}

func (v NullablePitrConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePitrConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


