# XClusterConfigEditFormData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoIncludeIndexTables** | Pointer to **bool** | Whether or not YBA should also include all index tables from any provided main tables. | [optional] 
**BootstrapParams** | Pointer to [**BootstrapParams**](BootstrapParams.md) |  | [optional] 
**Databases** | Pointer to **[]string** | WARNING: This is a preview API that could change. Source universe database IDs | [optional] 
**DryRun** | Pointer to **bool** | Run the pre-checks without actually running the subtasks | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**SourceRole** | Pointer to **string** | The role that the source universe should have in the xCluster config | [optional] 
**Status** | Pointer to **string** | Status | [optional] 
**Tables** | Pointer to **[]string** | Source universe table IDs | [optional] 
**TargetRole** | Pointer to **string** | The role that the target universe should have in the xCluster config | [optional] 

## Methods

### NewXClusterConfigEditFormData

`func NewXClusterConfigEditFormData() *XClusterConfigEditFormData`

NewXClusterConfigEditFormData instantiates a new XClusterConfigEditFormData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXClusterConfigEditFormDataWithDefaults

`func NewXClusterConfigEditFormDataWithDefaults() *XClusterConfigEditFormData`

NewXClusterConfigEditFormDataWithDefaults instantiates a new XClusterConfigEditFormData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoIncludeIndexTables

`func (o *XClusterConfigEditFormData) GetAutoIncludeIndexTables() bool`

GetAutoIncludeIndexTables returns the AutoIncludeIndexTables field if non-nil, zero value otherwise.

### GetAutoIncludeIndexTablesOk

`func (o *XClusterConfigEditFormData) GetAutoIncludeIndexTablesOk() (*bool, bool)`

GetAutoIncludeIndexTablesOk returns a tuple with the AutoIncludeIndexTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoIncludeIndexTables

`func (o *XClusterConfigEditFormData) SetAutoIncludeIndexTables(v bool)`

SetAutoIncludeIndexTables sets AutoIncludeIndexTables field to given value.

### HasAutoIncludeIndexTables

`func (o *XClusterConfigEditFormData) HasAutoIncludeIndexTables() bool`

HasAutoIncludeIndexTables returns a boolean if a field has been set.

### GetBootstrapParams

`func (o *XClusterConfigEditFormData) GetBootstrapParams() BootstrapParams`

GetBootstrapParams returns the BootstrapParams field if non-nil, zero value otherwise.

### GetBootstrapParamsOk

`func (o *XClusterConfigEditFormData) GetBootstrapParamsOk() (*BootstrapParams, bool)`

GetBootstrapParamsOk returns a tuple with the BootstrapParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapParams

`func (o *XClusterConfigEditFormData) SetBootstrapParams(v BootstrapParams)`

SetBootstrapParams sets BootstrapParams field to given value.

### HasBootstrapParams

`func (o *XClusterConfigEditFormData) HasBootstrapParams() bool`

HasBootstrapParams returns a boolean if a field has been set.

### GetDatabases

`func (o *XClusterConfigEditFormData) GetDatabases() []string`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *XClusterConfigEditFormData) GetDatabasesOk() (*[]string, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *XClusterConfigEditFormData) SetDatabases(v []string)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *XClusterConfigEditFormData) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.

### GetDryRun

`func (o *XClusterConfigEditFormData) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *XClusterConfigEditFormData) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *XClusterConfigEditFormData) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *XClusterConfigEditFormData) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetName

`func (o *XClusterConfigEditFormData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *XClusterConfigEditFormData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *XClusterConfigEditFormData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *XClusterConfigEditFormData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceRole

`func (o *XClusterConfigEditFormData) GetSourceRole() string`

GetSourceRole returns the SourceRole field if non-nil, zero value otherwise.

### GetSourceRoleOk

`func (o *XClusterConfigEditFormData) GetSourceRoleOk() (*string, bool)`

GetSourceRoleOk returns a tuple with the SourceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRole

`func (o *XClusterConfigEditFormData) SetSourceRole(v string)`

SetSourceRole sets SourceRole field to given value.

### HasSourceRole

`func (o *XClusterConfigEditFormData) HasSourceRole() bool`

HasSourceRole returns a boolean if a field has been set.

### GetStatus

`func (o *XClusterConfigEditFormData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *XClusterConfigEditFormData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *XClusterConfigEditFormData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *XClusterConfigEditFormData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTables

`func (o *XClusterConfigEditFormData) GetTables() []string`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *XClusterConfigEditFormData) GetTablesOk() (*[]string, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *XClusterConfigEditFormData) SetTables(v []string)`

SetTables sets Tables field to given value.

### HasTables

`func (o *XClusterConfigEditFormData) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetTargetRole

`func (o *XClusterConfigEditFormData) GetTargetRole() string`

GetTargetRole returns the TargetRole field if non-nil, zero value otherwise.

### GetTargetRoleOk

`func (o *XClusterConfigEditFormData) GetTargetRoleOk() (*string, bool)`

GetTargetRoleOk returns a tuple with the TargetRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRole

`func (o *XClusterConfigEditFormData) SetTargetRole(v string)`

SetTargetRole sets TargetRole field to given value.

### HasTargetRole

`func (o *XClusterConfigEditFormData) HasTargetRole() bool`

HasTargetRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


