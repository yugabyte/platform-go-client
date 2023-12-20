# DrConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **time.Time** | Create time of the DR config | [optional] 
**ModifyTime** | Pointer to **time.Time** | Last modify time of the DR config | [optional] 
**Name** | Pointer to **string** | Disaster recovery config name | [optional] 
**State** | Pointer to **string** | The state of the DR config | [optional] 
**Uuid** | Pointer to **string** | DR config UUID | [optional] 

## Methods

### NewDrConfig

`func NewDrConfig() *DrConfig`

NewDrConfig instantiates a new DrConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrConfigWithDefaults

`func NewDrConfigWithDefaults() *DrConfig`

NewDrConfigWithDefaults instantiates a new DrConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *DrConfig) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *DrConfig) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *DrConfig) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *DrConfig) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetModifyTime

`func (o *DrConfig) GetModifyTime() time.Time`

GetModifyTime returns the ModifyTime field if non-nil, zero value otherwise.

### GetModifyTimeOk

`func (o *DrConfig) GetModifyTimeOk() (*time.Time, bool)`

GetModifyTimeOk returns a tuple with the ModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyTime

`func (o *DrConfig) SetModifyTime(v time.Time)`

SetModifyTime sets ModifyTime field to given value.

### HasModifyTime

`func (o *DrConfig) HasModifyTime() bool`

HasModifyTime returns a boolean if a field has been set.

### GetName

`func (o *DrConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DrConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DrConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DrConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *DrConfig) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DrConfig) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DrConfig) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DrConfig) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *DrConfig) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DrConfig) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DrConfig) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DrConfig) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


