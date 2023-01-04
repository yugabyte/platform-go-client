# CreateTablespaceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TablespaceInfos** | [**[]TableSpaceInfo**](TableSpaceInfo.md) | Tablespaces to be created | 

## Methods

### NewCreateTablespaceParams

`func NewCreateTablespaceParams(tablespaceInfos []TableSpaceInfo, ) *CreateTablespaceParams`

NewCreateTablespaceParams instantiates a new CreateTablespaceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTablespaceParamsWithDefaults

`func NewCreateTablespaceParamsWithDefaults() *CreateTablespaceParams`

NewCreateTablespaceParamsWithDefaults instantiates a new CreateTablespaceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTablespaceInfos

`func (o *CreateTablespaceParams) GetTablespaceInfos() []TableSpaceInfo`

GetTablespaceInfos returns the TablespaceInfos field if non-nil, zero value otherwise.

### GetTablespaceInfosOk

`func (o *CreateTablespaceParams) GetTablespaceInfosOk() (*[]TableSpaceInfo, bool)`

GetTablespaceInfosOk returns a tuple with the TablespaceInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTablespaceInfos

`func (o *CreateTablespaceParams) SetTablespaceInfos(v []TableSpaceInfo)`

SetTablespaceInfos sets TablespaceInfos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


