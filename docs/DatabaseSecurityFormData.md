# DatabaseSecurityFormData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbName** | Pointer to **string** | YSQL DB Name | [optional] 
**YcqlAdminPassword** | Pointer to **string** | New YCQL admin password | [optional] 
**YcqlAdminUsername** | Pointer to **string** | YCQL admin username | [optional] 
**YcqlCurrAdminPassword** | Pointer to **string** | Current YCQL admin password | [optional] 
**YsqlAdminPassword** | Pointer to **string** | New YSQL admin password | [optional] 
**YsqlAdminUsername** | Pointer to **string** | YSQL admin username | [optional] 
**YsqlCurrAdminPassword** | Pointer to **string** | Current YSQL admin password | [optional] 

## Methods

### NewDatabaseSecurityFormData

`func NewDatabaseSecurityFormData() *DatabaseSecurityFormData`

NewDatabaseSecurityFormData instantiates a new DatabaseSecurityFormData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseSecurityFormDataWithDefaults

`func NewDatabaseSecurityFormDataWithDefaults() *DatabaseSecurityFormData`

NewDatabaseSecurityFormDataWithDefaults instantiates a new DatabaseSecurityFormData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbName

`func (o *DatabaseSecurityFormData) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *DatabaseSecurityFormData) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *DatabaseSecurityFormData) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *DatabaseSecurityFormData) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetYcqlAdminPassword

`func (o *DatabaseSecurityFormData) GetYcqlAdminPassword() string`

GetYcqlAdminPassword returns the YcqlAdminPassword field if non-nil, zero value otherwise.

### GetYcqlAdminPasswordOk

`func (o *DatabaseSecurityFormData) GetYcqlAdminPasswordOk() (*string, bool)`

GetYcqlAdminPasswordOk returns a tuple with the YcqlAdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYcqlAdminPassword

`func (o *DatabaseSecurityFormData) SetYcqlAdminPassword(v string)`

SetYcqlAdminPassword sets YcqlAdminPassword field to given value.

### HasYcqlAdminPassword

`func (o *DatabaseSecurityFormData) HasYcqlAdminPassword() bool`

HasYcqlAdminPassword returns a boolean if a field has been set.

### GetYcqlAdminUsername

`func (o *DatabaseSecurityFormData) GetYcqlAdminUsername() string`

GetYcqlAdminUsername returns the YcqlAdminUsername field if non-nil, zero value otherwise.

### GetYcqlAdminUsernameOk

`func (o *DatabaseSecurityFormData) GetYcqlAdminUsernameOk() (*string, bool)`

GetYcqlAdminUsernameOk returns a tuple with the YcqlAdminUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYcqlAdminUsername

`func (o *DatabaseSecurityFormData) SetYcqlAdminUsername(v string)`

SetYcqlAdminUsername sets YcqlAdminUsername field to given value.

### HasYcqlAdminUsername

`func (o *DatabaseSecurityFormData) HasYcqlAdminUsername() bool`

HasYcqlAdminUsername returns a boolean if a field has been set.

### GetYcqlCurrAdminPassword

`func (o *DatabaseSecurityFormData) GetYcqlCurrAdminPassword() string`

GetYcqlCurrAdminPassword returns the YcqlCurrAdminPassword field if non-nil, zero value otherwise.

### GetYcqlCurrAdminPasswordOk

`func (o *DatabaseSecurityFormData) GetYcqlCurrAdminPasswordOk() (*string, bool)`

GetYcqlCurrAdminPasswordOk returns a tuple with the YcqlCurrAdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYcqlCurrAdminPassword

`func (o *DatabaseSecurityFormData) SetYcqlCurrAdminPassword(v string)`

SetYcqlCurrAdminPassword sets YcqlCurrAdminPassword field to given value.

### HasYcqlCurrAdminPassword

`func (o *DatabaseSecurityFormData) HasYcqlCurrAdminPassword() bool`

HasYcqlCurrAdminPassword returns a boolean if a field has been set.

### GetYsqlAdminPassword

`func (o *DatabaseSecurityFormData) GetYsqlAdminPassword() string`

GetYsqlAdminPassword returns the YsqlAdminPassword field if non-nil, zero value otherwise.

### GetYsqlAdminPasswordOk

`func (o *DatabaseSecurityFormData) GetYsqlAdminPasswordOk() (*string, bool)`

GetYsqlAdminPasswordOk returns a tuple with the YsqlAdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlAdminPassword

`func (o *DatabaseSecurityFormData) SetYsqlAdminPassword(v string)`

SetYsqlAdminPassword sets YsqlAdminPassword field to given value.

### HasYsqlAdminPassword

`func (o *DatabaseSecurityFormData) HasYsqlAdminPassword() bool`

HasYsqlAdminPassword returns a boolean if a field has been set.

### GetYsqlAdminUsername

`func (o *DatabaseSecurityFormData) GetYsqlAdminUsername() string`

GetYsqlAdminUsername returns the YsqlAdminUsername field if non-nil, zero value otherwise.

### GetYsqlAdminUsernameOk

`func (o *DatabaseSecurityFormData) GetYsqlAdminUsernameOk() (*string, bool)`

GetYsqlAdminUsernameOk returns a tuple with the YsqlAdminUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlAdminUsername

`func (o *DatabaseSecurityFormData) SetYsqlAdminUsername(v string)`

SetYsqlAdminUsername sets YsqlAdminUsername field to given value.

### HasYsqlAdminUsername

`func (o *DatabaseSecurityFormData) HasYsqlAdminUsername() bool`

HasYsqlAdminUsername returns a boolean if a field has been set.

### GetYsqlCurrAdminPassword

`func (o *DatabaseSecurityFormData) GetYsqlCurrAdminPassword() string`

GetYsqlCurrAdminPassword returns the YsqlCurrAdminPassword field if non-nil, zero value otherwise.

### GetYsqlCurrAdminPasswordOk

`func (o *DatabaseSecurityFormData) GetYsqlCurrAdminPasswordOk() (*string, bool)`

GetYsqlCurrAdminPasswordOk returns a tuple with the YsqlCurrAdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlCurrAdminPassword

`func (o *DatabaseSecurityFormData) SetYsqlCurrAdminPassword(v string)`

SetYsqlCurrAdminPassword sets YsqlCurrAdminPassword field to given value.

### HasYsqlCurrAdminPassword

`func (o *DatabaseSecurityFormData) HasYsqlCurrAdminPassword() bool`

HasYsqlCurrAdminPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


