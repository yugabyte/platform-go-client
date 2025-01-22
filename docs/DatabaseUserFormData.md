# DatabaseUserFormData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbName** | **string** |  | 
**DbRoleAttributes** | Pointer to [**[]RoleAttribute**](RoleAttribute.md) | YbaApi Internal. | [optional] 
**Password** | **string** |  | 
**Username** | **string** |  | 
**YcqlAdminPassword** | **string** |  | 
**YcqlAdminUsername** | **string** |  | 
**YsqlAdminPassword** | **string** |  | 
**YsqlAdminUsername** | **string** |  | 

## Methods

### NewDatabaseUserFormData

`func NewDatabaseUserFormData(dbName string, password string, username string, ycqlAdminPassword string, ycqlAdminUsername string, ysqlAdminPassword string, ysqlAdminUsername string, ) *DatabaseUserFormData`

NewDatabaseUserFormData instantiates a new DatabaseUserFormData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseUserFormDataWithDefaults

`func NewDatabaseUserFormDataWithDefaults() *DatabaseUserFormData`

NewDatabaseUserFormDataWithDefaults instantiates a new DatabaseUserFormData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbName

`func (o *DatabaseUserFormData) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *DatabaseUserFormData) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *DatabaseUserFormData) SetDbName(v string)`

SetDbName sets DbName field to given value.


### GetDbRoleAttributes

`func (o *DatabaseUserFormData) GetDbRoleAttributes() []RoleAttribute`

GetDbRoleAttributes returns the DbRoleAttributes field if non-nil, zero value otherwise.

### GetDbRoleAttributesOk

`func (o *DatabaseUserFormData) GetDbRoleAttributesOk() (*[]RoleAttribute, bool)`

GetDbRoleAttributesOk returns a tuple with the DbRoleAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbRoleAttributes

`func (o *DatabaseUserFormData) SetDbRoleAttributes(v []RoleAttribute)`

SetDbRoleAttributes sets DbRoleAttributes field to given value.

### HasDbRoleAttributes

`func (o *DatabaseUserFormData) HasDbRoleAttributes() bool`

HasDbRoleAttributes returns a boolean if a field has been set.

### GetPassword

`func (o *DatabaseUserFormData) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DatabaseUserFormData) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DatabaseUserFormData) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *DatabaseUserFormData) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DatabaseUserFormData) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DatabaseUserFormData) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetYcqlAdminPassword

`func (o *DatabaseUserFormData) GetYcqlAdminPassword() string`

GetYcqlAdminPassword returns the YcqlAdminPassword field if non-nil, zero value otherwise.

### GetYcqlAdminPasswordOk

`func (o *DatabaseUserFormData) GetYcqlAdminPasswordOk() (*string, bool)`

GetYcqlAdminPasswordOk returns a tuple with the YcqlAdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYcqlAdminPassword

`func (o *DatabaseUserFormData) SetYcqlAdminPassword(v string)`

SetYcqlAdminPassword sets YcqlAdminPassword field to given value.


### GetYcqlAdminUsername

`func (o *DatabaseUserFormData) GetYcqlAdminUsername() string`

GetYcqlAdminUsername returns the YcqlAdminUsername field if non-nil, zero value otherwise.

### GetYcqlAdminUsernameOk

`func (o *DatabaseUserFormData) GetYcqlAdminUsernameOk() (*string, bool)`

GetYcqlAdminUsernameOk returns a tuple with the YcqlAdminUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYcqlAdminUsername

`func (o *DatabaseUserFormData) SetYcqlAdminUsername(v string)`

SetYcqlAdminUsername sets YcqlAdminUsername field to given value.


### GetYsqlAdminPassword

`func (o *DatabaseUserFormData) GetYsqlAdminPassword() string`

GetYsqlAdminPassword returns the YsqlAdminPassword field if non-nil, zero value otherwise.

### GetYsqlAdminPasswordOk

`func (o *DatabaseUserFormData) GetYsqlAdminPasswordOk() (*string, bool)`

GetYsqlAdminPasswordOk returns a tuple with the YsqlAdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlAdminPassword

`func (o *DatabaseUserFormData) SetYsqlAdminPassword(v string)`

SetYsqlAdminPassword sets YsqlAdminPassword field to given value.


### GetYsqlAdminUsername

`func (o *DatabaseUserFormData) GetYsqlAdminUsername() string`

GetYsqlAdminUsername returns the YsqlAdminUsername field if non-nil, zero value otherwise.

### GetYsqlAdminUsernameOk

`func (o *DatabaseUserFormData) GetYsqlAdminUsernameOk() (*string, bool)`

GetYsqlAdminUsernameOk returns a tuple with the YsqlAdminUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlAdminUsername

`func (o *DatabaseUserFormData) SetYsqlAdminUsername(v string)`

SetYsqlAdminUsername sets YsqlAdminUsername field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


