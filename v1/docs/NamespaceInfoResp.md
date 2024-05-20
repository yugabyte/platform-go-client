# NamespaceInfoResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Namespace name | [optional] 
**NamespaceUUID** | Pointer to **string** | Namespace UUID | [optional] [readonly] 
**TableType** | Pointer to **string** | Table type | [optional] 

## Methods

### NewNamespaceInfoResp

`func NewNamespaceInfoResp() *NamespaceInfoResp`

NewNamespaceInfoResp instantiates a new NamespaceInfoResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceInfoRespWithDefaults

`func NewNamespaceInfoRespWithDefaults() *NamespaceInfoResp`

NewNamespaceInfoRespWithDefaults instantiates a new NamespaceInfoResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NamespaceInfoResp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NamespaceInfoResp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NamespaceInfoResp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NamespaceInfoResp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespaceUUID

`func (o *NamespaceInfoResp) GetNamespaceUUID() string`

GetNamespaceUUID returns the NamespaceUUID field if non-nil, zero value otherwise.

### GetNamespaceUUIDOk

`func (o *NamespaceInfoResp) GetNamespaceUUIDOk() (*string, bool)`

GetNamespaceUUIDOk returns a tuple with the NamespaceUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceUUID

`func (o *NamespaceInfoResp) SetNamespaceUUID(v string)`

SetNamespaceUUID sets NamespaceUUID field to given value.

### HasNamespaceUUID

`func (o *NamespaceInfoResp) HasNamespaceUUID() bool`

HasNamespaceUUID returns a boolean if a field has been set.

### GetTableType

`func (o *NamespaceInfoResp) GetTableType() string`

GetTableType returns the TableType field if non-nil, zero value otherwise.

### GetTableTypeOk

`func (o *NamespaceInfoResp) GetTableTypeOk() (*string, bool)`

GetTableTypeOk returns a tuple with the TableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableType

`func (o *NamespaceInfoResp) SetTableType(v string)`

SetTableType sets TableType field to given value.

### HasTableType

`func (o *NamespaceInfoResp) HasTableType() bool`

HasTableType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


