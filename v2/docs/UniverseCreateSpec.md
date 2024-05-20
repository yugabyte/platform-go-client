# UniverseCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | [**UniverseSpec**](UniverseSpec.md) |  | 
**Arch** | **string** | CPU Arch of DB nodes. | 
**Ysql** | Pointer to [**UniverseCreateSpecYsql**](UniverseCreateSpecYsql.md) |  | [optional] 
**Ycql** | Pointer to [**UniverseCreateSpecYcql**](UniverseCreateSpecYcql.md) |  | [optional] 

## Methods

### NewUniverseCreateSpec

`func NewUniverseCreateSpec(spec UniverseSpec, arch string, ) *UniverseCreateSpec`

NewUniverseCreateSpec instantiates a new UniverseCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseCreateSpecWithDefaults

`func NewUniverseCreateSpecWithDefaults() *UniverseCreateSpec`

NewUniverseCreateSpecWithDefaults instantiates a new UniverseCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *UniverseCreateSpec) GetSpec() UniverseSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *UniverseCreateSpec) GetSpecOk() (*UniverseSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *UniverseCreateSpec) SetSpec(v UniverseSpec)`

SetSpec sets Spec field to given value.


### GetArch

`func (o *UniverseCreateSpec) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *UniverseCreateSpec) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *UniverseCreateSpec) SetArch(v string)`

SetArch sets Arch field to given value.


### GetYsql

`func (o *UniverseCreateSpec) GetYsql() UniverseCreateSpecYsql`

GetYsql returns the Ysql field if non-nil, zero value otherwise.

### GetYsqlOk

`func (o *UniverseCreateSpec) GetYsqlOk() (*UniverseCreateSpecYsql, bool)`

GetYsqlOk returns a tuple with the Ysql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsql

`func (o *UniverseCreateSpec) SetYsql(v UniverseCreateSpecYsql)`

SetYsql sets Ysql field to given value.

### HasYsql

`func (o *UniverseCreateSpec) HasYsql() bool`

HasYsql returns a boolean if a field has been set.

### GetYcql

`func (o *UniverseCreateSpec) GetYcql() UniverseCreateSpecYcql`

GetYcql returns the Ycql field if non-nil, zero value otherwise.

### GetYcqlOk

`func (o *UniverseCreateSpec) GetYcqlOk() (*UniverseCreateSpecYcql, bool)`

GetYcqlOk returns a tuple with the Ycql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYcql

`func (o *UniverseCreateSpec) SetYcql(v UniverseCreateSpecYcql)`

SetYcql sets Ycql field to given value.

### HasYcql

`func (o *UniverseCreateSpec) HasYcql() bool`

HasYcql returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


