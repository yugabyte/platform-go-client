# YSQLSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Whether to enable YSQL API on this Universe | [optional] 
**EnableAuth** | Pointer to **bool** | Whether to enable authentication to access YSQL on this Universe | [optional] 
**Password** | Pointer to **string** | Password to set for the YSQL database in this universe. Required if enable_auth is true. | [optional] 
**EnableConnectionPooling** | Pointer to **bool** | Use built-in YSQL service for maximizing the number of simultaneous connections to database. | [optional] [default to false]

## Methods

### NewYSQLSpec

`func NewYSQLSpec() *YSQLSpec`

NewYSQLSpec instantiates a new YSQLSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYSQLSpecWithDefaults

`func NewYSQLSpecWithDefaults() *YSQLSpec`

NewYSQLSpecWithDefaults instantiates a new YSQLSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *YSQLSpec) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *YSQLSpec) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *YSQLSpec) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *YSQLSpec) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableAuth

`func (o *YSQLSpec) GetEnableAuth() bool`

GetEnableAuth returns the EnableAuth field if non-nil, zero value otherwise.

### GetEnableAuthOk

`func (o *YSQLSpec) GetEnableAuthOk() (*bool, bool)`

GetEnableAuthOk returns a tuple with the EnableAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAuth

`func (o *YSQLSpec) SetEnableAuth(v bool)`

SetEnableAuth sets EnableAuth field to given value.

### HasEnableAuth

`func (o *YSQLSpec) HasEnableAuth() bool`

HasEnableAuth returns a boolean if a field has been set.

### GetPassword

`func (o *YSQLSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *YSQLSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *YSQLSpec) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *YSQLSpec) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEnableConnectionPooling

`func (o *YSQLSpec) GetEnableConnectionPooling() bool`

GetEnableConnectionPooling returns the EnableConnectionPooling field if non-nil, zero value otherwise.

### GetEnableConnectionPoolingOk

`func (o *YSQLSpec) GetEnableConnectionPoolingOk() (*bool, bool)`

GetEnableConnectionPoolingOk returns a tuple with the EnableConnectionPooling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableConnectionPooling

`func (o *YSQLSpec) SetEnableConnectionPooling(v bool)`

SetEnableConnectionPooling sets EnableConnectionPooling field to given value.

### HasEnableConnectionPooling

`func (o *YSQLSpec) HasEnableConnectionPooling() bool`

HasEnableConnectionPooling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


