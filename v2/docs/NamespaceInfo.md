# NamespaceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NamespaceUuid** | Pointer to **string** | Namespace UUID. | [optional] [readonly] 
**Name** | Pointer to **string** | Namespace name. | [optional] [readonly] 
**TableType** | Pointer to [**XClusterTableType**](XClusterTableType.md) |  | [optional] 

## Methods

### NewNamespaceInfo

`func NewNamespaceInfo() *NamespaceInfo`

NewNamespaceInfo instantiates a new NamespaceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceInfoWithDefaults

`func NewNamespaceInfoWithDefaults() *NamespaceInfo`

NewNamespaceInfoWithDefaults instantiates a new NamespaceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespaceUuid

`func (o *NamespaceInfo) GetNamespaceUuid() string`

GetNamespaceUuid returns the NamespaceUuid field if non-nil, zero value otherwise.

### GetNamespaceUuidOk

`func (o *NamespaceInfo) GetNamespaceUuidOk() (*string, bool)`

GetNamespaceUuidOk returns a tuple with the NamespaceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceUuid

`func (o *NamespaceInfo) SetNamespaceUuid(v string)`

SetNamespaceUuid sets NamespaceUuid field to given value.

### HasNamespaceUuid

`func (o *NamespaceInfo) HasNamespaceUuid() bool`

HasNamespaceUuid returns a boolean if a field has been set.

### GetName

`func (o *NamespaceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NamespaceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NamespaceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NamespaceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTableType

`func (o *NamespaceInfo) GetTableType() XClusterTableType`

GetTableType returns the TableType field if non-nil, zero value otherwise.

### GetTableTypeOk

`func (o *NamespaceInfo) GetTableTypeOk() (*XClusterTableType, bool)`

GetTableTypeOk returns a tuple with the TableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableType

`func (o *NamespaceInfo) SetTableType(v XClusterTableType)`

SetTableType sets TableType field to given value.

### HasTableType

`func (o *NamespaceInfo) HasTableType() bool`

HasTableType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


