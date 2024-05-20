# TablespaceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConflictingTablespaces** | **[]string** |  | 
**ContainsTablespaces** | **bool** |  | 
**UnsupportedTablespaces** | **[]string** |  | 

## Methods

### NewTablespaceResponse

`func NewTablespaceResponse(conflictingTablespaces []string, containsTablespaces bool, unsupportedTablespaces []string, ) *TablespaceResponse`

NewTablespaceResponse instantiates a new TablespaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTablespaceResponseWithDefaults

`func NewTablespaceResponseWithDefaults() *TablespaceResponse`

NewTablespaceResponseWithDefaults instantiates a new TablespaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConflictingTablespaces

`func (o *TablespaceResponse) GetConflictingTablespaces() []string`

GetConflictingTablespaces returns the ConflictingTablespaces field if non-nil, zero value otherwise.

### GetConflictingTablespacesOk

`func (o *TablespaceResponse) GetConflictingTablespacesOk() (*[]string, bool)`

GetConflictingTablespacesOk returns a tuple with the ConflictingTablespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictingTablespaces

`func (o *TablespaceResponse) SetConflictingTablespaces(v []string)`

SetConflictingTablespaces sets ConflictingTablespaces field to given value.


### GetContainsTablespaces

`func (o *TablespaceResponse) GetContainsTablespaces() bool`

GetContainsTablespaces returns the ContainsTablespaces field if non-nil, zero value otherwise.

### GetContainsTablespacesOk

`func (o *TablespaceResponse) GetContainsTablespacesOk() (*bool, bool)`

GetContainsTablespacesOk returns a tuple with the ContainsTablespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsTablespaces

`func (o *TablespaceResponse) SetContainsTablespaces(v bool)`

SetContainsTablespaces sets ContainsTablespaces field to given value.


### GetUnsupportedTablespaces

`func (o *TablespaceResponse) GetUnsupportedTablespaces() []string`

GetUnsupportedTablespaces returns the UnsupportedTablespaces field if non-nil, zero value otherwise.

### GetUnsupportedTablespacesOk

`func (o *TablespaceResponse) GetUnsupportedTablespacesOk() (*[]string, bool)`

GetUnsupportedTablespacesOk returns a tuple with the UnsupportedTablespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsupportedTablespaces

`func (o *TablespaceResponse) SetUnsupportedTablespaces(v []string)`

SetUnsupportedTablespaces sets UnsupportedTablespaces field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


