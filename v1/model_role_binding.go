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

// RoleBinding struct for RoleBinding
type RoleBinding struct {
	// RoleBinding create time
	CreateTime *time.Time `json:"createTime,omitempty"`
	GroupInfo *GroupMappingInfo `json:"groupInfo,omitempty"`
	Principal Principal `json:"principal"`
	ResourceGroup *ResourceGroup `json:"resourceGroup,omitempty"`
	Role *Role `json:"role,omitempty"`
	// Role binding type
	Type *string `json:"type,omitempty"`
	// RoleBinding last updated time
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	User *Users `json:"user,omitempty"`
	// UUID
	Uuid *string `json:"uuid,omitempty"`
}

// NewRoleBinding instantiates a new RoleBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleBinding(principal Principal) *RoleBinding {
	this := RoleBinding{}
	this.Principal = principal
	return &this
}

// NewRoleBindingWithDefaults instantiates a new RoleBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleBindingWithDefaults() *RoleBinding {
	this := RoleBinding{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *RoleBinding) GetCreateTime() time.Time {
	if o == nil || o.CreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBinding) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *RoleBinding) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *RoleBinding) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetGroupInfo returns the GroupInfo field value if set, zero value otherwise.
func (o *RoleBinding) GetGroupInfo() GroupMappingInfo {
	if o == nil || o.GroupInfo == nil {
		var ret GroupMappingInfo
		return ret
	}
	return *o.GroupInfo
}

// GetGroupInfoOk returns a tuple with the GroupInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBinding) GetGroupInfoOk() (*GroupMappingInfo, bool) {
	if o == nil || o.GroupInfo == nil {
		return nil, false
	}
	return o.GroupInfo, true
}

// HasGroupInfo returns a boolean if a field has been set.
func (o *RoleBinding) HasGroupInfo() bool {
	if o != nil && o.GroupInfo != nil {
		return true
	}

	return false
}

// SetGroupInfo gets a reference to the given GroupMappingInfo and assigns it to the GroupInfo field.
func (o *RoleBinding) SetGroupInfo(v GroupMappingInfo) {
	o.GroupInfo = &v
}

// GetPrincipal returns the Principal field value
func (o *RoleBinding) GetPrincipal() Principal {
	if o == nil {
		var ret Principal
		return ret
	}

	return o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value
// and a boolean to check if the value has been set.
func (o *RoleBinding) GetPrincipalOk() (*Principal, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Principal, true
}

// SetPrincipal sets field value
func (o *RoleBinding) SetPrincipal(v Principal) {
	o.Principal = v
}

// GetResourceGroup returns the ResourceGroup field value if set, zero value otherwise.
func (o *RoleBinding) GetResourceGroup() ResourceGroup {
	if o == nil || o.ResourceGroup == nil {
		var ret ResourceGroup
		return ret
	}
	return *o.ResourceGroup
}

// GetResourceGroupOk returns a tuple with the ResourceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBinding) GetResourceGroupOk() (*ResourceGroup, bool) {
	if o == nil || o.ResourceGroup == nil {
		return nil, false
	}
	return o.ResourceGroup, true
}

// HasResourceGroup returns a boolean if a field has been set.
func (o *RoleBinding) HasResourceGroup() bool {
	if o != nil && o.ResourceGroup != nil {
		return true
	}

	return false
}

// SetResourceGroup gets a reference to the given ResourceGroup and assigns it to the ResourceGroup field.
func (o *RoleBinding) SetResourceGroup(v ResourceGroup) {
	o.ResourceGroup = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *RoleBinding) GetRole() Role {
	if o == nil || o.Role == nil {
		var ret Role
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBinding) GetRoleOk() (*Role, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *RoleBinding) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given Role and assigns it to the Role field.
func (o *RoleBinding) SetRole(v Role) {
	o.Role = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RoleBinding) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBinding) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RoleBinding) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RoleBinding) SetType(v string) {
	o.Type = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *RoleBinding) GetUpdateTime() time.Time {
	if o == nil || o.UpdateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBinding) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || o.UpdateTime == nil {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *RoleBinding) HasUpdateTime() bool {
	if o != nil && o.UpdateTime != nil {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *RoleBinding) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RoleBinding) GetUser() Users {
	if o == nil || o.User == nil {
		var ret Users
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBinding) GetUserOk() (*Users, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RoleBinding) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given Users and assigns it to the User field.
func (o *RoleBinding) SetUser(v Users) {
	o.User = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *RoleBinding) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBinding) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *RoleBinding) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *RoleBinding) SetUuid(v string) {
	o.Uuid = &v
}

func (o RoleBinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreateTime != nil {
		toSerialize["createTime"] = o.CreateTime
	}
	if o.GroupInfo != nil {
		toSerialize["groupInfo"] = o.GroupInfo
	}
	if true {
		toSerialize["principal"] = o.Principal
	}
	if o.ResourceGroup != nil {
		toSerialize["resourceGroup"] = o.ResourceGroup
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UpdateTime != nil {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableRoleBinding struct {
	value *RoleBinding
	isSet bool
}

func (v NullableRoleBinding) Get() *RoleBinding {
	return v.value
}

func (v *NullableRoleBinding) Set(val *RoleBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleBinding(val *RoleBinding) *NullableRoleBinding {
	return &NullableRoleBinding{value: val, isSet: true}
}

func (v NullableRoleBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


