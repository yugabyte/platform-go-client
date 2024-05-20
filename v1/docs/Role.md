# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedOn** | Pointer to **time.Time** | Role create time | [optional] 
**CustomerUUID** | Pointer to **string** | Customer UUID | [optional] 
**Description** | Pointer to **string** | Role description | [optional] 
**Name** | Pointer to **string** | Role name | [optional] 
**PermissionDetails** | Pointer to [**PermissionDetails**](PermissionDetails.md) |  | [optional] 
**RoleType** | Pointer to **string** | Type of the role | [optional] 
**RoleUUID** | Pointer to **string** | Role UUID | [optional] [readonly] 
**UpdatedOn** | Pointer to **time.Time** | Role last updated time | [optional] 

## Methods

### NewRole

`func NewRole() *Role`

NewRole instantiates a new Role object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithDefaults

`func NewRoleWithDefaults() *Role`

NewRoleWithDefaults instantiates a new Role object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedOn

`func (o *Role) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Role) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Role) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Role) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetCustomerUUID

`func (o *Role) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *Role) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *Role) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.

### HasCustomerUUID

`func (o *Role) HasCustomerUUID() bool`

HasCustomerUUID returns a boolean if a field has been set.

### GetDescription

`func (o *Role) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Role) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Role) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Role) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *Role) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Role) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Role) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Role) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissionDetails

`func (o *Role) GetPermissionDetails() PermissionDetails`

GetPermissionDetails returns the PermissionDetails field if non-nil, zero value otherwise.

### GetPermissionDetailsOk

`func (o *Role) GetPermissionDetailsOk() (*PermissionDetails, bool)`

GetPermissionDetailsOk returns a tuple with the PermissionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionDetails

`func (o *Role) SetPermissionDetails(v PermissionDetails)`

SetPermissionDetails sets PermissionDetails field to given value.

### HasPermissionDetails

`func (o *Role) HasPermissionDetails() bool`

HasPermissionDetails returns a boolean if a field has been set.

### GetRoleType

`func (o *Role) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *Role) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *Role) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *Role) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetRoleUUID

`func (o *Role) GetRoleUUID() string`

GetRoleUUID returns the RoleUUID field if non-nil, zero value otherwise.

### GetRoleUUIDOk

`func (o *Role) GetRoleUUIDOk() (*string, bool)`

GetRoleUUIDOk returns a tuple with the RoleUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleUUID

`func (o *Role) SetRoleUUID(v string)`

SetRoleUUID sets RoleUUID field to given value.

### HasRoleUUID

`func (o *Role) HasRoleUUID() bool`

HasRoleUUID returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *Role) GetUpdatedOn() time.Time`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *Role) GetUpdatedOnOk() (*time.Time, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *Role) SetUpdatedOn(v time.Time)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *Role) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


