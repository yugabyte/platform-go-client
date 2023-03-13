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

// DatabaseSecurityFormData struct for DatabaseSecurityFormData
type DatabaseSecurityFormData struct {
	DbName string `json:"dbName"`
	YcqlAdminPassword string `json:"ycqlAdminPassword"`
	YcqlAdminUsername string `json:"ycqlAdminUsername"`
	YcqlCurrAdminPassword string `json:"ycqlCurrAdminPassword"`
	YsqlAdminPassword string `json:"ysqlAdminPassword"`
	YsqlAdminUsername string `json:"ysqlAdminUsername"`
	YsqlCurrAdminPassword string `json:"ysqlCurrAdminPassword"`
}

// NewDatabaseSecurityFormData instantiates a new DatabaseSecurityFormData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseSecurityFormData(dbName string, ycqlAdminPassword string, ycqlAdminUsername string, ycqlCurrAdminPassword string, ysqlAdminPassword string, ysqlAdminUsername string, ysqlCurrAdminPassword string, ) *DatabaseSecurityFormData {
	this := DatabaseSecurityFormData{}
	this.DbName = dbName
	this.YcqlAdminPassword = ycqlAdminPassword
	this.YcqlAdminUsername = ycqlAdminUsername
	this.YcqlCurrAdminPassword = ycqlCurrAdminPassword
	this.YsqlAdminPassword = ysqlAdminPassword
	this.YsqlAdminUsername = ysqlAdminUsername
	this.YsqlCurrAdminPassword = ysqlCurrAdminPassword
	return &this
}

// NewDatabaseSecurityFormDataWithDefaults instantiates a new DatabaseSecurityFormData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseSecurityFormDataWithDefaults() *DatabaseSecurityFormData {
	this := DatabaseSecurityFormData{}
	return &this
}

// GetDbName returns the DbName field value
func (o *DatabaseSecurityFormData) GetDbName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value
// and a boolean to check if the value has been set.
func (o *DatabaseSecurityFormData) GetDbNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DbName, true
}

// SetDbName sets field value
func (o *DatabaseSecurityFormData) SetDbName(v string) {
	o.DbName = v
}

// GetYcqlAdminPassword returns the YcqlAdminPassword field value
func (o *DatabaseSecurityFormData) GetYcqlAdminPassword() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.YcqlAdminPassword
}

// GetYcqlAdminPasswordOk returns a tuple with the YcqlAdminPassword field value
// and a boolean to check if the value has been set.
func (o *DatabaseSecurityFormData) GetYcqlAdminPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.YcqlAdminPassword, true
}

// SetYcqlAdminPassword sets field value
func (o *DatabaseSecurityFormData) SetYcqlAdminPassword(v string) {
	o.YcqlAdminPassword = v
}

// GetYcqlAdminUsername returns the YcqlAdminUsername field value
func (o *DatabaseSecurityFormData) GetYcqlAdminUsername() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.YcqlAdminUsername
}

// GetYcqlAdminUsernameOk returns a tuple with the YcqlAdminUsername field value
// and a boolean to check if the value has been set.
func (o *DatabaseSecurityFormData) GetYcqlAdminUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.YcqlAdminUsername, true
}

// SetYcqlAdminUsername sets field value
func (o *DatabaseSecurityFormData) SetYcqlAdminUsername(v string) {
	o.YcqlAdminUsername = v
}

// GetYcqlCurrAdminPassword returns the YcqlCurrAdminPassword field value
func (o *DatabaseSecurityFormData) GetYcqlCurrAdminPassword() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.YcqlCurrAdminPassword
}

// GetYcqlCurrAdminPasswordOk returns a tuple with the YcqlCurrAdminPassword field value
// and a boolean to check if the value has been set.
func (o *DatabaseSecurityFormData) GetYcqlCurrAdminPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.YcqlCurrAdminPassword, true
}

// SetYcqlCurrAdminPassword sets field value
func (o *DatabaseSecurityFormData) SetYcqlCurrAdminPassword(v string) {
	o.YcqlCurrAdminPassword = v
}

// GetYsqlAdminPassword returns the YsqlAdminPassword field value
func (o *DatabaseSecurityFormData) GetYsqlAdminPassword() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.YsqlAdminPassword
}

// GetYsqlAdminPasswordOk returns a tuple with the YsqlAdminPassword field value
// and a boolean to check if the value has been set.
func (o *DatabaseSecurityFormData) GetYsqlAdminPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.YsqlAdminPassword, true
}

// SetYsqlAdminPassword sets field value
func (o *DatabaseSecurityFormData) SetYsqlAdminPassword(v string) {
	o.YsqlAdminPassword = v
}

// GetYsqlAdminUsername returns the YsqlAdminUsername field value
func (o *DatabaseSecurityFormData) GetYsqlAdminUsername() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.YsqlAdminUsername
}

// GetYsqlAdminUsernameOk returns a tuple with the YsqlAdminUsername field value
// and a boolean to check if the value has been set.
func (o *DatabaseSecurityFormData) GetYsqlAdminUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.YsqlAdminUsername, true
}

// SetYsqlAdminUsername sets field value
func (o *DatabaseSecurityFormData) SetYsqlAdminUsername(v string) {
	o.YsqlAdminUsername = v
}

// GetYsqlCurrAdminPassword returns the YsqlCurrAdminPassword field value
func (o *DatabaseSecurityFormData) GetYsqlCurrAdminPassword() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.YsqlCurrAdminPassword
}

// GetYsqlCurrAdminPasswordOk returns a tuple with the YsqlCurrAdminPassword field value
// and a boolean to check if the value has been set.
func (o *DatabaseSecurityFormData) GetYsqlCurrAdminPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.YsqlCurrAdminPassword, true
}

// SetYsqlCurrAdminPassword sets field value
func (o *DatabaseSecurityFormData) SetYsqlCurrAdminPassword(v string) {
	o.YsqlCurrAdminPassword = v
}

func (o DatabaseSecurityFormData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dbName"] = o.DbName
	}
	if true {
		toSerialize["ycqlAdminPassword"] = o.YcqlAdminPassword
	}
	if true {
		toSerialize["ycqlAdminUsername"] = o.YcqlAdminUsername
	}
	if true {
		toSerialize["ycqlCurrAdminPassword"] = o.YcqlCurrAdminPassword
	}
	if true {
		toSerialize["ysqlAdminPassword"] = o.YsqlAdminPassword
	}
	if true {
		toSerialize["ysqlAdminUsername"] = o.YsqlAdminUsername
	}
	if true {
		toSerialize["ysqlCurrAdminPassword"] = o.YsqlCurrAdminPassword
	}
	return json.Marshal(toSerialize)
}

type NullableDatabaseSecurityFormData struct {
	value *DatabaseSecurityFormData
	isSet bool
}

func (v NullableDatabaseSecurityFormData) Get() *DatabaseSecurityFormData {
	return v.value
}

func (v *NullableDatabaseSecurityFormData) Set(val *DatabaseSecurityFormData) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseSecurityFormData) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseSecurityFormData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseSecurityFormData(val *DatabaseSecurityFormData) *NullableDatabaseSecurityFormData {
	return &NullableDatabaseSecurityFormData{value: val, isSet: true}
}

func (v NullableDatabaseSecurityFormData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseSecurityFormData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

