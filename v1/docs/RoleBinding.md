# RoleBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **time.Time** | RoleBinding create time | [optional] 
**ResourceGroup** | Pointer to [**ResourceGroup**](ResourceGroup.md) |  | [optional] 
**Role** | Pointer to [**Role**](Role.md) |  | [optional] 
**Type** | Pointer to **string** | Role binding type | [optional] 
**UpdateTime** | Pointer to **time.Time** | RoleBinding last updated time | [optional] 
**User** | Pointer to [**Users**](Users.md) |  | [optional] 
**Uuid** | Pointer to **string** | UUID | [optional] [readonly] 

## Methods

### NewRoleBinding

`func NewRoleBinding() *RoleBinding`

NewRoleBinding instantiates a new RoleBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleBindingWithDefaults

`func NewRoleBindingWithDefaults() *RoleBinding`

NewRoleBindingWithDefaults instantiates a new RoleBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *RoleBinding) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *RoleBinding) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *RoleBinding) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *RoleBinding) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetResourceGroup

`func (o *RoleBinding) GetResourceGroup() ResourceGroup`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *RoleBinding) GetResourceGroupOk() (*ResourceGroup, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *RoleBinding) SetResourceGroup(v ResourceGroup)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *RoleBinding) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetRole

`func (o *RoleBinding) GetRole() Role`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RoleBinding) GetRoleOk() (*Role, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RoleBinding) SetRole(v Role)`

SetRole sets Role field to given value.

### HasRole

`func (o *RoleBinding) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetType

`func (o *RoleBinding) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleBinding) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleBinding) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RoleBinding) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *RoleBinding) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *RoleBinding) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *RoleBinding) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *RoleBinding) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetUser

`func (o *RoleBinding) GetUser() Users`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RoleBinding) GetUserOk() (*Users, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RoleBinding) SetUser(v Users)`

SetUser sets User field to given value.

### HasUser

`func (o *RoleBinding) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUuid

`func (o *RoleBinding) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RoleBinding) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RoleBinding) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RoleBinding) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


