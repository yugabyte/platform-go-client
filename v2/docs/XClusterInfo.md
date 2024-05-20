# XClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceRootCertDirPath** | Pointer to **string** | The value of certs_for_cdc_dir gflag | [optional] 
**SourceXClusterConfigs** | Pointer to **[]string** | The source universe&#39;s xcluster replication relationships | [optional] 
**TargetXClusterConfigs** | Pointer to **[]string** | The target universe&#39;s xcluster replication relationships | [optional] 

## Methods

### NewXClusterInfo

`func NewXClusterInfo() *XClusterInfo`

NewXClusterInfo instantiates a new XClusterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXClusterInfoWithDefaults

`func NewXClusterInfoWithDefaults() *XClusterInfo`

NewXClusterInfoWithDefaults instantiates a new XClusterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceRootCertDirPath

`func (o *XClusterInfo) GetSourceRootCertDirPath() string`

GetSourceRootCertDirPath returns the SourceRootCertDirPath field if non-nil, zero value otherwise.

### GetSourceRootCertDirPathOk

`func (o *XClusterInfo) GetSourceRootCertDirPathOk() (*string, bool)`

GetSourceRootCertDirPathOk returns a tuple with the SourceRootCertDirPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRootCertDirPath

`func (o *XClusterInfo) SetSourceRootCertDirPath(v string)`

SetSourceRootCertDirPath sets SourceRootCertDirPath field to given value.

### HasSourceRootCertDirPath

`func (o *XClusterInfo) HasSourceRootCertDirPath() bool`

HasSourceRootCertDirPath returns a boolean if a field has been set.

### GetSourceXClusterConfigs

`func (o *XClusterInfo) GetSourceXClusterConfigs() []string`

GetSourceXClusterConfigs returns the SourceXClusterConfigs field if non-nil, zero value otherwise.

### GetSourceXClusterConfigsOk

`func (o *XClusterInfo) GetSourceXClusterConfigsOk() (*[]string, bool)`

GetSourceXClusterConfigsOk returns a tuple with the SourceXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceXClusterConfigs

`func (o *XClusterInfo) SetSourceXClusterConfigs(v []string)`

SetSourceXClusterConfigs sets SourceXClusterConfigs field to given value.

### HasSourceXClusterConfigs

`func (o *XClusterInfo) HasSourceXClusterConfigs() bool`

HasSourceXClusterConfigs returns a boolean if a field has been set.

### GetTargetXClusterConfigs

`func (o *XClusterInfo) GetTargetXClusterConfigs() []string`

GetTargetXClusterConfigs returns the TargetXClusterConfigs field if non-nil, zero value otherwise.

### GetTargetXClusterConfigsOk

`func (o *XClusterInfo) GetTargetXClusterConfigsOk() (*[]string, bool)`

GetTargetXClusterConfigsOk returns a tuple with the TargetXClusterConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetXClusterConfigs

`func (o *XClusterInfo) SetTargetXClusterConfigs(v []string)`

SetTargetXClusterConfigs sets TargetXClusterConfigs field to given value.

### HasTargetXClusterConfigs

`func (o *XClusterInfo) HasTargetXClusterConfigs() bool`

HasTargetXClusterConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


