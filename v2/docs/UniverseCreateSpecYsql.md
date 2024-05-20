# UniverseCreateSpecYsql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Password to set for the YSQL database in this universe. Required if spec.ysql.enable_auth is true. | [optional] 

## Methods

### NewUniverseCreateSpecYsql

`func NewUniverseCreateSpecYsql() *UniverseCreateSpecYsql`

NewUniverseCreateSpecYsql instantiates a new UniverseCreateSpecYsql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseCreateSpecYsqlWithDefaults

`func NewUniverseCreateSpecYsqlWithDefaults() *UniverseCreateSpecYsql`

NewUniverseCreateSpecYsqlWithDefaults instantiates a new UniverseCreateSpecYsql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *UniverseCreateSpecYsql) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UniverseCreateSpecYsql) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UniverseCreateSpecYsql) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UniverseCreateSpecYsql) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


