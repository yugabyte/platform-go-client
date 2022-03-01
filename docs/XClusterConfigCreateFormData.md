# XClusterConfigCreateFormData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name | 
**SourceUniverseUUID** | **string** | Source Universe UUID | 
**Tables** | **[]string** | Source Universe table IDs | 
**TargetUniverseUUID** | **string** | Target Universe UUID | 

## Methods

### NewXClusterConfigCreateFormData

`func NewXClusterConfigCreateFormData(name string, sourceUniverseUUID string, tables []string, targetUniverseUUID string, ) *XClusterConfigCreateFormData`

NewXClusterConfigCreateFormData instantiates a new XClusterConfigCreateFormData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXClusterConfigCreateFormDataWithDefaults

`func NewXClusterConfigCreateFormDataWithDefaults() *XClusterConfigCreateFormData`

NewXClusterConfigCreateFormDataWithDefaults instantiates a new XClusterConfigCreateFormData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *XClusterConfigCreateFormData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *XClusterConfigCreateFormData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *XClusterConfigCreateFormData) SetName(v string)`

SetName sets Name field to given value.


### GetSourceUniverseUUID

`func (o *XClusterConfigCreateFormData) GetSourceUniverseUUID() string`

GetSourceUniverseUUID returns the SourceUniverseUUID field if non-nil, zero value otherwise.

### GetSourceUniverseUUIDOk

`func (o *XClusterConfigCreateFormData) GetSourceUniverseUUIDOk() (*string, bool)`

GetSourceUniverseUUIDOk returns a tuple with the SourceUniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUniverseUUID

`func (o *XClusterConfigCreateFormData) SetSourceUniverseUUID(v string)`

SetSourceUniverseUUID sets SourceUniverseUUID field to given value.


### GetTables

`func (o *XClusterConfigCreateFormData) GetTables() []string`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *XClusterConfigCreateFormData) GetTablesOk() (*[]string, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *XClusterConfigCreateFormData) SetTables(v []string)`

SetTables sets Tables field to given value.


### GetTargetUniverseUUID

`func (o *XClusterConfigCreateFormData) GetTargetUniverseUUID() string`

GetTargetUniverseUUID returns the TargetUniverseUUID field if non-nil, zero value otherwise.

### GetTargetUniverseUUIDOk

`func (o *XClusterConfigCreateFormData) GetTargetUniverseUUIDOk() (*string, bool)`

GetTargetUniverseUUIDOk returns a tuple with the TargetUniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUniverseUUID

`func (o *XClusterConfigCreateFormData) SetTargetUniverseUUID(v string)`

SetTargetUniverseUUID sets TargetUniverseUUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


