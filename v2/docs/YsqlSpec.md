# YsqlSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Whether to enable YSQL API on this Universe | [optional] 
**EnableAuth** | Pointer to **bool** | Whether to enable authentication to access YSQL on this Universe | [optional] 

## Methods

### NewYsqlSpec

`func NewYsqlSpec() *YsqlSpec`

NewYsqlSpec instantiates a new YsqlSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYsqlSpecWithDefaults

`func NewYsqlSpecWithDefaults() *YsqlSpec`

NewYsqlSpecWithDefaults instantiates a new YsqlSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *YsqlSpec) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *YsqlSpec) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *YsqlSpec) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *YsqlSpec) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableAuth

`func (o *YsqlSpec) GetEnableAuth() bool`

GetEnableAuth returns the EnableAuth field if non-nil, zero value otherwise.

### GetEnableAuthOk

`func (o *YsqlSpec) GetEnableAuthOk() (*bool, bool)`

GetEnableAuthOk returns a tuple with the EnableAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAuth

`func (o *YsqlSpec) SetEnableAuth(v bool)`

SetEnableAuth sets EnableAuth field to given value.

### HasEnableAuth

`func (o *YsqlSpec) HasEnableAuth() bool`

HasEnableAuth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


