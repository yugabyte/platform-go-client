# GroupMappingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | Pointer to **time.Time** | Group mapping creation time | [optional] 
**CustomerUUID** | **string** |  | 
**GroupUUID** | **string** |  | 
**Identifier** | **string** |  | 
**RoleUUID** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewGroupMappingInfo

`func NewGroupMappingInfo(customerUUID string, groupUUID string, identifier string, roleUUID string, type_ string, ) *GroupMappingInfo`

NewGroupMappingInfo instantiates a new GroupMappingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMappingInfoWithDefaults

`func NewGroupMappingInfoWithDefaults() *GroupMappingInfo`

NewGroupMappingInfoWithDefaults instantiates a new GroupMappingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDate

`func (o *GroupMappingInfo) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *GroupMappingInfo) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *GroupMappingInfo) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *GroupMappingInfo) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetCustomerUUID

`func (o *GroupMappingInfo) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *GroupMappingInfo) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *GroupMappingInfo) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetGroupUUID

`func (o *GroupMappingInfo) GetGroupUUID() string`

GetGroupUUID returns the GroupUUID field if non-nil, zero value otherwise.

### GetGroupUUIDOk

`func (o *GroupMappingInfo) GetGroupUUIDOk() (*string, bool)`

GetGroupUUIDOk returns a tuple with the GroupUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupUUID

`func (o *GroupMappingInfo) SetGroupUUID(v string)`

SetGroupUUID sets GroupUUID field to given value.


### GetIdentifier

`func (o *GroupMappingInfo) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *GroupMappingInfo) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *GroupMappingInfo) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetRoleUUID

`func (o *GroupMappingInfo) GetRoleUUID() string`

GetRoleUUID returns the RoleUUID field if non-nil, zero value otherwise.

### GetRoleUUIDOk

`func (o *GroupMappingInfo) GetRoleUUIDOk() (*string, bool)`

GetRoleUUIDOk returns a tuple with the RoleUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleUUID

`func (o *GroupMappingInfo) SetRoleUUID(v string)`

SetRoleUUID sets RoleUUID field to given value.


### GetType

`func (o *GroupMappingInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupMappingInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupMappingInfo) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


