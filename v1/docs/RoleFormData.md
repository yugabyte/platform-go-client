# RoleFormData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the role to be created | 
**Name** | **string** | Name of the role to be created | 
**PermissionList** | [**[]Permission**](Permission.md) | List of permissions given to the role | 

## Methods

### NewRoleFormData

`func NewRoleFormData(description string, name string, permissionList []Permission, ) *RoleFormData`

NewRoleFormData instantiates a new RoleFormData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleFormDataWithDefaults

`func NewRoleFormDataWithDefaults() *RoleFormData`

NewRoleFormDataWithDefaults instantiates a new RoleFormData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RoleFormData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleFormData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleFormData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *RoleFormData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleFormData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleFormData) SetName(v string)`

SetName sets Name field to given value.


### GetPermissionList

`func (o *RoleFormData) GetPermissionList() []Permission`

GetPermissionList returns the PermissionList field if non-nil, zero value otherwise.

### GetPermissionListOk

`func (o *RoleFormData) GetPermissionListOk() (*[]Permission, bool)`

GetPermissionListOk returns a tuple with the PermissionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionList

`func (o *RoleFormData) SetPermissionList(v []Permission)`

SetPermissionList sets PermissionList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


