# YcqlSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Whether to enable YCQL API on this Universe | [optional] 
**EnableAuth** | Pointer to **bool** | Whether to enable authentication to access YCQL on this Universe | [optional] 

## Methods

### NewYcqlSpec

`func NewYcqlSpec() *YcqlSpec`

NewYcqlSpec instantiates a new YcqlSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYcqlSpecWithDefaults

`func NewYcqlSpecWithDefaults() *YcqlSpec`

NewYcqlSpecWithDefaults instantiates a new YcqlSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *YcqlSpec) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *YcqlSpec) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *YcqlSpec) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *YcqlSpec) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableAuth

`func (o *YcqlSpec) GetEnableAuth() bool`

GetEnableAuth returns the EnableAuth field if non-nil, zero value otherwise.

### GetEnableAuthOk

`func (o *YcqlSpec) GetEnableAuthOk() (*bool, bool)`

GetEnableAuthOk returns a tuple with the EnableAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAuth

`func (o *YcqlSpec) SetEnableAuth(v bool)`

SetEnableAuth sets EnableAuth field to given value.

### HasEnableAuth

`func (o *YcqlSpec) HasEnableAuth() bool`

HasEnableAuth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


