# BootstrapParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupRequestParams** | [**BootstarpBackupParams**](BootstarpBackupParams.md) |  | 
**Tables** | **[]string** | Source Universe table IDs that need bootstrapping; must be a subset of tables in the main body | 

## Methods

### NewBootstrapParams

`func NewBootstrapParams(backupRequestParams BootstarpBackupParams, tables []string, ) *BootstrapParams`

NewBootstrapParams instantiates a new BootstrapParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootstrapParamsWithDefaults

`func NewBootstrapParamsWithDefaults() *BootstrapParams`

NewBootstrapParamsWithDefaults instantiates a new BootstrapParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupRequestParams

`func (o *BootstrapParams) GetBackupRequestParams() BootstarpBackupParams`

GetBackupRequestParams returns the BackupRequestParams field if non-nil, zero value otherwise.

### GetBackupRequestParamsOk

`func (o *BootstrapParams) GetBackupRequestParamsOk() (*BootstarpBackupParams, bool)`

GetBackupRequestParamsOk returns a tuple with the BackupRequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRequestParams

`func (o *BootstrapParams) SetBackupRequestParams(v BootstarpBackupParams)`

SetBackupRequestParams sets BackupRequestParams field to given value.


### GetTables

`func (o *BootstrapParams) GetTables() []string`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *BootstrapParams) GetTablesOk() (*[]string, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *BootstrapParams) SetTables(v []string)`

SetTables sets Tables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


