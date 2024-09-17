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

// Audit Audit logging for requests and responses
type Audit struct {
	// Action
	Action *string `json:"action,omitempty"`
	// Additional Details
	AdditionalDetails *map[string]interface{} `json:"additionalDetails,omitempty"`
	// API call
	ApiCall *string `json:"apiCall,omitempty"`
	// API method
	ApiMethod *string `json:"apiMethod,omitempty"`
	// Audit UID
	AuditID *int64 `json:"auditID,omitempty"`
	// Customer UUID
	CustomerUUID *string `json:"customerUUID,omitempty"`
	// Audit UUID
	Payload *map[string]interface{} `json:"payload,omitempty"`
	// Target
	Target *string `json:"target,omitempty"`
	// Target ID
	TargetID *string `json:"targetID,omitempty"`
	// Task UUID
	TaskUUID *string `json:"taskUUID,omitempty"`
	// The task creation time.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// User IP Address
	UserAddress *string `json:"userAddress,omitempty"`
	// User Email
	UserEmail *string `json:"userEmail,omitempty"`
	// User UUID
	UserUUID *string `json:"userUUID,omitempty"`
}

// NewAudit instantiates a new Audit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudit() *Audit {
	this := Audit{}
	return &this
}

// NewAuditWithDefaults instantiates a new Audit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditWithDefaults() *Audit {
	this := Audit{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *Audit) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *Audit) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *Audit) SetAction(v string) {
	o.Action = &v
}

// GetAdditionalDetails returns the AdditionalDetails field value if set, zero value otherwise.
func (o *Audit) GetAdditionalDetails() map[string]interface{} {
	if o == nil || o.AdditionalDetails == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.AdditionalDetails
}

// GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetAdditionalDetailsOk() (*map[string]interface{}, bool) {
	if o == nil || o.AdditionalDetails == nil {
		return nil, false
	}
	return o.AdditionalDetails, true
}

// HasAdditionalDetails returns a boolean if a field has been set.
func (o *Audit) HasAdditionalDetails() bool {
	if o != nil && o.AdditionalDetails != nil {
		return true
	}

	return false
}

// SetAdditionalDetails gets a reference to the given map[string]interface{} and assigns it to the AdditionalDetails field.
func (o *Audit) SetAdditionalDetails(v map[string]interface{}) {
	o.AdditionalDetails = &v
}

// GetApiCall returns the ApiCall field value if set, zero value otherwise.
func (o *Audit) GetApiCall() string {
	if o == nil || o.ApiCall == nil {
		var ret string
		return ret
	}
	return *o.ApiCall
}

// GetApiCallOk returns a tuple with the ApiCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetApiCallOk() (*string, bool) {
	if o == nil || o.ApiCall == nil {
		return nil, false
	}
	return o.ApiCall, true
}

// HasApiCall returns a boolean if a field has been set.
func (o *Audit) HasApiCall() bool {
	if o != nil && o.ApiCall != nil {
		return true
	}

	return false
}

// SetApiCall gets a reference to the given string and assigns it to the ApiCall field.
func (o *Audit) SetApiCall(v string) {
	o.ApiCall = &v
}

// GetApiMethod returns the ApiMethod field value if set, zero value otherwise.
func (o *Audit) GetApiMethod() string {
	if o == nil || o.ApiMethod == nil {
		var ret string
		return ret
	}
	return *o.ApiMethod
}

// GetApiMethodOk returns a tuple with the ApiMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetApiMethodOk() (*string, bool) {
	if o == nil || o.ApiMethod == nil {
		return nil, false
	}
	return o.ApiMethod, true
}

// HasApiMethod returns a boolean if a field has been set.
func (o *Audit) HasApiMethod() bool {
	if o != nil && o.ApiMethod != nil {
		return true
	}

	return false
}

// SetApiMethod gets a reference to the given string and assigns it to the ApiMethod field.
func (o *Audit) SetApiMethod(v string) {
	o.ApiMethod = &v
}

// GetAuditID returns the AuditID field value if set, zero value otherwise.
func (o *Audit) GetAuditID() int64 {
	if o == nil || o.AuditID == nil {
		var ret int64
		return ret
	}
	return *o.AuditID
}

// GetAuditIDOk returns a tuple with the AuditID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetAuditIDOk() (*int64, bool) {
	if o == nil || o.AuditID == nil {
		return nil, false
	}
	return o.AuditID, true
}

// HasAuditID returns a boolean if a field has been set.
func (o *Audit) HasAuditID() bool {
	if o != nil && o.AuditID != nil {
		return true
	}

	return false
}

// SetAuditID gets a reference to the given int64 and assigns it to the AuditID field.
func (o *Audit) SetAuditID(v int64) {
	o.AuditID = &v
}

// GetCustomerUUID returns the CustomerUUID field value if set, zero value otherwise.
func (o *Audit) GetCustomerUUID() string {
	if o == nil || o.CustomerUUID == nil {
		var ret string
		return ret
	}
	return *o.CustomerUUID
}

// GetCustomerUUIDOk returns a tuple with the CustomerUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetCustomerUUIDOk() (*string, bool) {
	if o == nil || o.CustomerUUID == nil {
		return nil, false
	}
	return o.CustomerUUID, true
}

// HasCustomerUUID returns a boolean if a field has been set.
func (o *Audit) HasCustomerUUID() bool {
	if o != nil && o.CustomerUUID != nil {
		return true
	}

	return false
}

// SetCustomerUUID gets a reference to the given string and assigns it to the CustomerUUID field.
func (o *Audit) SetCustomerUUID(v string) {
	o.CustomerUUID = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *Audit) GetPayload() map[string]interface{} {
	if o == nil || o.Payload == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetPayloadOk() (*map[string]interface{}, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *Audit) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given map[string]interface{} and assigns it to the Payload field.
func (o *Audit) SetPayload(v map[string]interface{}) {
	o.Payload = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *Audit) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *Audit) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *Audit) SetTarget(v string) {
	o.Target = &v
}

// GetTargetID returns the TargetID field value if set, zero value otherwise.
func (o *Audit) GetTargetID() string {
	if o == nil || o.TargetID == nil {
		var ret string
		return ret
	}
	return *o.TargetID
}

// GetTargetIDOk returns a tuple with the TargetID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetTargetIDOk() (*string, bool) {
	if o == nil || o.TargetID == nil {
		return nil, false
	}
	return o.TargetID, true
}

// HasTargetID returns a boolean if a field has been set.
func (o *Audit) HasTargetID() bool {
	if o != nil && o.TargetID != nil {
		return true
	}

	return false
}

// SetTargetID gets a reference to the given string and assigns it to the TargetID field.
func (o *Audit) SetTargetID(v string) {
	o.TargetID = &v
}

// GetTaskUUID returns the TaskUUID field value if set, zero value otherwise.
func (o *Audit) GetTaskUUID() string {
	if o == nil || o.TaskUUID == nil {
		var ret string
		return ret
	}
	return *o.TaskUUID
}

// GetTaskUUIDOk returns a tuple with the TaskUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetTaskUUIDOk() (*string, bool) {
	if o == nil || o.TaskUUID == nil {
		return nil, false
	}
	return o.TaskUUID, true
}

// HasTaskUUID returns a boolean if a field has been set.
func (o *Audit) HasTaskUUID() bool {
	if o != nil && o.TaskUUID != nil {
		return true
	}

	return false
}

// SetTaskUUID gets a reference to the given string and assigns it to the TaskUUID field.
func (o *Audit) SetTaskUUID(v string) {
	o.TaskUUID = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *Audit) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Audit) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *Audit) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetUserAddress returns the UserAddress field value if set, zero value otherwise.
func (o *Audit) GetUserAddress() string {
	if o == nil || o.UserAddress == nil {
		var ret string
		return ret
	}
	return *o.UserAddress
}

// GetUserAddressOk returns a tuple with the UserAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetUserAddressOk() (*string, bool) {
	if o == nil || o.UserAddress == nil {
		return nil, false
	}
	return o.UserAddress, true
}

// HasUserAddress returns a boolean if a field has been set.
func (o *Audit) HasUserAddress() bool {
	if o != nil && o.UserAddress != nil {
		return true
	}

	return false
}

// SetUserAddress gets a reference to the given string and assigns it to the UserAddress field.
func (o *Audit) SetUserAddress(v string) {
	o.UserAddress = &v
}

// GetUserEmail returns the UserEmail field value if set, zero value otherwise.
func (o *Audit) GetUserEmail() string {
	if o == nil || o.UserEmail == nil {
		var ret string
		return ret
	}
	return *o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetUserEmailOk() (*string, bool) {
	if o == nil || o.UserEmail == nil {
		return nil, false
	}
	return o.UserEmail, true
}

// HasUserEmail returns a boolean if a field has been set.
func (o *Audit) HasUserEmail() bool {
	if o != nil && o.UserEmail != nil {
		return true
	}

	return false
}

// SetUserEmail gets a reference to the given string and assigns it to the UserEmail field.
func (o *Audit) SetUserEmail(v string) {
	o.UserEmail = &v
}

// GetUserUUID returns the UserUUID field value if set, zero value otherwise.
func (o *Audit) GetUserUUID() string {
	if o == nil || o.UserUUID == nil {
		var ret string
		return ret
	}
	return *o.UserUUID
}

// GetUserUUIDOk returns a tuple with the UserUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetUserUUIDOk() (*string, bool) {
	if o == nil || o.UserUUID == nil {
		return nil, false
	}
	return o.UserUUID, true
}

// HasUserUUID returns a boolean if a field has been set.
func (o *Audit) HasUserUUID() bool {
	if o != nil && o.UserUUID != nil {
		return true
	}

	return false
}

// SetUserUUID gets a reference to the given string and assigns it to the UserUUID field.
func (o *Audit) SetUserUUID(v string) {
	o.UserUUID = &v
}

func (o Audit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.AdditionalDetails != nil {
		toSerialize["additionalDetails"] = o.AdditionalDetails
	}
	if o.ApiCall != nil {
		toSerialize["apiCall"] = o.ApiCall
	}
	if o.ApiMethod != nil {
		toSerialize["apiMethod"] = o.ApiMethod
	}
	if o.AuditID != nil {
		toSerialize["auditID"] = o.AuditID
	}
	if o.CustomerUUID != nil {
		toSerialize["customerUUID"] = o.CustomerUUID
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.TargetID != nil {
		toSerialize["targetID"] = o.TargetID
	}
	if o.TaskUUID != nil {
		toSerialize["taskUUID"] = o.TaskUUID
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.UserAddress != nil {
		toSerialize["userAddress"] = o.UserAddress
	}
	if o.UserEmail != nil {
		toSerialize["userEmail"] = o.UserEmail
	}
	if o.UserUUID != nil {
		toSerialize["userUUID"] = o.UserUUID
	}
	return json.Marshal(toSerialize)
}

type NullableAudit struct {
	value *Audit
	isSet bool
}

func (v NullableAudit) Get() *Audit {
	return v.value
}

func (v *NullableAudit) Set(val *Audit) {
	v.value = val
	v.isSet = true
}

func (v NullableAudit) IsSet() bool {
	return v.isSet
}

func (v *NullableAudit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudit(val *Audit) *NullableAudit {
	return &NullableAudit{value: val, isSet: true}
}

func (v NullableAudit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


