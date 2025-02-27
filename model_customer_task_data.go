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
	"time"
)

// CustomerTaskData Customer task data
type CustomerTaskData struct {
	// Customer task abortable
	Abortable *bool `json:"abortable,omitempty"`
	// Whether the Customer task can be rolled back
	CanRollback *bool `json:"canRollback,omitempty"`
	// Customer task completion time
	CompletionTime *time.Time `json:"completionTime,omitempty"`
	// Correlation id
	CorrelationId *string `json:"correlationId,omitempty"`
	// Customer task creation time
	CreateTime *time.Time `json:"createTime,omitempty"`
	// Customer task UUID
	Id *string `json:"id,omitempty"`
	// Customer task percentage completed
	PercentComplete *int32 `json:"percentComplete,omitempty"`
	// Customer task retryable
	Retryable *bool `json:"retryable,omitempty"`
	// Customer task status
	Status *string `json:"status,omitempty"`
	// Customer task target
	Target *string `json:"target,omitempty"`
	// Customer task target UUID
	TargetUUID *string `json:"targetUUID,omitempty"`
	// Customer task title
	Title *string `json:"title,omitempty"`
	// Customer task type
	Type *string `json:"type,omitempty"`
	// Customer task type name
	TypeName *string `json:"typeName,omitempty"`
	// Customer Email
	UserEmail *string `json:"userEmail,omitempty"`
}

// NewCustomerTaskData instantiates a new CustomerTaskData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerTaskData() *CustomerTaskData {
	this := CustomerTaskData{}
	return &this
}

// NewCustomerTaskDataWithDefaults instantiates a new CustomerTaskData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerTaskDataWithDefaults() *CustomerTaskData {
	this := CustomerTaskData{}
	return &this
}

// GetAbortable returns the Abortable field value if set, zero value otherwise.
func (o *CustomerTaskData) GetAbortable() bool {
	if o == nil || o.Abortable == nil {
		var ret bool
		return ret
	}
	return *o.Abortable
}

// GetAbortableOk returns a tuple with the Abortable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerTaskData) GetAbortableOk() (*bool, bool) {
	if o == nil || o.Abortable == nil {
		return nil, false
	}
	return o.Abortable, true
}

// HasAbortable returns a boolean if a field has been set.
func (o *CustomerTaskData) HasAbortable() bool {
	if o != nil && o.Abortable != nil {
		return true
	}

	return false
}

// SetAbortable gets a reference to the given bool and assigns it to the Abortable field.
func (o *CustomerTaskData) SetAbortable(v bool) {
	o.Abortable = &v
}

// GetCanRollback returns the CanRollback field value if set, zero value otherwise.
func (o *CustomerTaskData) GetCanRollback() bool {
	if o == nil || o.CanRollback == nil {
		var ret bool
		return ret
	}
	return *o.CanRollback
}

// GetCanRollbackOk returns a tuple with the CanRollback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerTaskData) GetCanRollbackOk() (*bool, bool) {
	if o == nil || o.CanRollback == nil {
		return nil, false
	}
	return o.CanRollback, true
}

// HasCanRollback returns a boolean if a field has been set.
func (o *CustomerTaskData) HasCanRollback() bool {
	if o != nil && o.CanRollback != nil {
		return true
	}

	return false
}

// SetCanRollback gets a reference to the given bool and assigns it to the CanRollback field.
func (o *CustomerTaskData) SetCanRollback(v bool) {
	o.CanRollback = &v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *CustomerTaskData) GetCompletionTime() time.Time {
	if o == nil || o.CompletionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerTaskData) GetCompletionTimeOk() (*time.Time, bool) {
	if o == nil || o.CompletionTime == nil {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *CustomerTaskData) HasCompletionTime() bool {
	if o != nil && o.CompletionTime != nil {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given time.Time and assigns it to the CompletionTime field.
func (o *CustomerTaskData) SetCompletionTime(v time.Time) {
	o.CompletionTime = &v
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *CustomerTaskData) GetCorrelationId() string {
	if o == nil || o.CorrelationId == nil {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerTaskData) GetCorrelationIdOk() (*string, bool) {
	if o == nil || o.CorrelationId == nil {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *CustomerTaskData) HasCorrelationId() bool {
	if o != nil && o.CorrelationId != nil {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *CustomerTaskData) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *CustomerTaskData) GetCreateTime() time.Time {
	if o == nil || o.CreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerTaskData) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *CustomerTaskData) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *CustomerTaskData) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomerTaskData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerTaskData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerTaskData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomerTaskData) SetId(v string) {
	o.Id = &v
}

// GetPercentComplete returns the PercentComplete field value if set, zero value otherwise.
func (o *CustomerTaskData) GetPercentComplete() int32 {
	if o == nil || o.PercentComplete == nil {
		var ret int32
		return ret
	}
	return *o.PercentComplete
}

// GetPercentCompleteOk returns a tuple with the PercentComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerTaskData) GetPercentCompleteOk() (*int32, bool) {
	if o == nil || o.PercentComplete == nil {
		return nil, false
	}
	return o.PercentComplete, true
}

// HasPercentComplete returns a boolean if a field has been set.
func (o *CustomerTaskData) HasPercentComplete() bool {
	if o != nil && o.PercentComplete != nil {
		return true
	}

	return false
}

// SetPercentComplete gets a reference to the given int32 and assigns it to the PercentComplete field.
func (o *CustomerTaskData) SetPercentComplete(v int32) {
	o.PercentComplete = &v
}

// GetRetryable returns the Retryable field value if set, zero value otherwise.
func (o *CustomerTaskData) GetRetryable() bool {
	if o == nil || o.Retryable == nil {
		var ret bool
		return ret
	}
	return *o.Retryable
}

// GetRetryableOk returns a tuple with the Retryable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerTaskData) GetRetryableOk() (*bool, bool) {
	if o == nil || o.Retryable == nil {
		return nil, false
	}
	return o.Retryable, true
}

// HasRetryable returns a boolean if a field has been set.
func (o *CustomerTaskData) HasRetryable() bool {
	if o != nil && o.Retryable != nil {
		return true
	}

	return false
}

// SetRetryable gets a reference to the given bool and assigns it to the Retryable field.
func (o *CustomerTaskData) SetRetryable(v bool) {
	o.Retryable = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CustomerTaskData) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerTaskData) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CustomerTaskData) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CustomerTaskData) SetStatus(v string) {
	o.Status = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *CustomerTaskData) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerTaskData) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *CustomerTaskData) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *CustomerTaskData) SetTarget(v string) {
	o.Target = &v
}

// GetTargetUUID returns the TargetUUID field value if set, zero value otherwise.
func (o *CustomerTaskData) GetTargetUUID() string {
	if o == nil || o.TargetUUID == nil {
		var ret string
		return ret
	}
	return *o.TargetUUID
}

// GetTargetUUIDOk returns a tuple with the TargetUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerTaskData) GetTargetUUIDOk() (*string, bool) {
	if o == nil || o.TargetUUID == nil {
		return nil, false
	}
	return o.TargetUUID, true
}

// HasTargetUUID returns a boolean if a field has been set.
func (o *CustomerTaskData) HasTargetUUID() bool {
	if o != nil && o.TargetUUID != nil {
		return true
	}

	return false
}

// SetTargetUUID gets a reference to the given string and assigns it to the TargetUUID field.
func (o *CustomerTaskData) SetTargetUUID(v string) {
	o.TargetUUID = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CustomerTaskData) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerTaskData) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CustomerTaskData) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CustomerTaskData) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CustomerTaskData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerTaskData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomerTaskData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CustomerTaskData) SetType(v string) {
	o.Type = &v
}

// GetTypeName returns the TypeName field value if set, zero value otherwise.
func (o *CustomerTaskData) GetTypeName() string {
	if o == nil || o.TypeName == nil {
		var ret string
		return ret
	}
	return *o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerTaskData) GetTypeNameOk() (*string, bool) {
	if o == nil || o.TypeName == nil {
		return nil, false
	}
	return o.TypeName, true
}

// HasTypeName returns a boolean if a field has been set.
func (o *CustomerTaskData) HasTypeName() bool {
	if o != nil && o.TypeName != nil {
		return true
	}

	return false
}

// SetTypeName gets a reference to the given string and assigns it to the TypeName field.
func (o *CustomerTaskData) SetTypeName(v string) {
	o.TypeName = &v
}

// GetUserEmail returns the UserEmail field value if set, zero value otherwise.
func (o *CustomerTaskData) GetUserEmail() string {
	if o == nil || o.UserEmail == nil {
		var ret string
		return ret
	}
	return *o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerTaskData) GetUserEmailOk() (*string, bool) {
	if o == nil || o.UserEmail == nil {
		return nil, false
	}
	return o.UserEmail, true
}

// HasUserEmail returns a boolean if a field has been set.
func (o *CustomerTaskData) HasUserEmail() bool {
	if o != nil && o.UserEmail != nil {
		return true
	}

	return false
}

// SetUserEmail gets a reference to the given string and assigns it to the UserEmail field.
func (o *CustomerTaskData) SetUserEmail(v string) {
	o.UserEmail = &v
}

func (o CustomerTaskData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Abortable != nil {
		toSerialize["abortable"] = o.Abortable
	}
	if o.CanRollback != nil {
		toSerialize["canRollback"] = o.CanRollback
	}
	if o.CompletionTime != nil {
		toSerialize["completionTime"] = o.CompletionTime
	}
	if o.CorrelationId != nil {
		toSerialize["correlationId"] = o.CorrelationId
	}
	if o.CreateTime != nil {
		toSerialize["createTime"] = o.CreateTime
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PercentComplete != nil {
		toSerialize["percentComplete"] = o.PercentComplete
	}
	if o.Retryable != nil {
		toSerialize["retryable"] = o.Retryable
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.TargetUUID != nil {
		toSerialize["targetUUID"] = o.TargetUUID
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.TypeName != nil {
		toSerialize["typeName"] = o.TypeName
	}
	if o.UserEmail != nil {
		toSerialize["userEmail"] = o.UserEmail
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerTaskData struct {
	value *CustomerTaskData
	isSet bool
}

func (v NullableCustomerTaskData) Get() *CustomerTaskData {
	return v.value
}

func (v *NullableCustomerTaskData) Set(val *CustomerTaskData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerTaskData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerTaskData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerTaskData(val *CustomerTaskData) *NullableCustomerTaskData {
	return &NullableCustomerTaskData{value: val, isSet: true}
}

func (v NullableCustomerTaskData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerTaskData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


