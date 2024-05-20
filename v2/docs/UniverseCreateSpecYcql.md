# UniverseCreateSpecYcql

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Password to set for the YCQL database in this universe. Required if spec.ycql.enable_auth is true. | [optional] 

## Methods

### NewUniverseCreateSpecYcql

`func NewUniverseCreateSpecYcql() *UniverseCreateSpecYcql`

NewUniverseCreateSpecYcql instantiates a new UniverseCreateSpecYcql object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseCreateSpecYcqlWithDefaults

`func NewUniverseCreateSpecYcqlWithDefaults() *UniverseCreateSpecYcql`

NewUniverseCreateSpecYcqlWithDefaults instantiates a new UniverseCreateSpecYcql object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *UniverseCreateSpecYcql) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UniverseCreateSpecYcql) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UniverseCreateSpecYcql) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UniverseCreateSpecYcql) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


