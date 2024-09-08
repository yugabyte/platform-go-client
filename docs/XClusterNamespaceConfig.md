# XClusterNamespaceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceNamespaceId** | **string** |  | 
**SourceNamespaceInfo** | Pointer to [**NamespaceInfoResp**](NamespaceInfoResp.md) |  | [optional] 
**Status** | Pointer to **string** | Status | [optional] 
**TargetNamespaceInfo** | Pointer to [**NamespaceInfoResp**](NamespaceInfoResp.md) |  | [optional] 

## Methods

### NewXClusterNamespaceConfig

`func NewXClusterNamespaceConfig(sourceNamespaceId string, ) *XClusterNamespaceConfig`

NewXClusterNamespaceConfig instantiates a new XClusterNamespaceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXClusterNamespaceConfigWithDefaults

`func NewXClusterNamespaceConfigWithDefaults() *XClusterNamespaceConfig`

NewXClusterNamespaceConfigWithDefaults instantiates a new XClusterNamespaceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceNamespaceId

`func (o *XClusterNamespaceConfig) GetSourceNamespaceId() string`

GetSourceNamespaceId returns the SourceNamespaceId field if non-nil, zero value otherwise.

### GetSourceNamespaceIdOk

`func (o *XClusterNamespaceConfig) GetSourceNamespaceIdOk() (*string, bool)`

GetSourceNamespaceIdOk returns a tuple with the SourceNamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNamespaceId

`func (o *XClusterNamespaceConfig) SetSourceNamespaceId(v string)`

SetSourceNamespaceId sets SourceNamespaceId field to given value.


### GetSourceNamespaceInfo

`func (o *XClusterNamespaceConfig) GetSourceNamespaceInfo() NamespaceInfoResp`

GetSourceNamespaceInfo returns the SourceNamespaceInfo field if non-nil, zero value otherwise.

### GetSourceNamespaceInfoOk

`func (o *XClusterNamespaceConfig) GetSourceNamespaceInfoOk() (*NamespaceInfoResp, bool)`

GetSourceNamespaceInfoOk returns a tuple with the SourceNamespaceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNamespaceInfo

`func (o *XClusterNamespaceConfig) SetSourceNamespaceInfo(v NamespaceInfoResp)`

SetSourceNamespaceInfo sets SourceNamespaceInfo field to given value.

### HasSourceNamespaceInfo

`func (o *XClusterNamespaceConfig) HasSourceNamespaceInfo() bool`

HasSourceNamespaceInfo returns a boolean if a field has been set.

### GetStatus

`func (o *XClusterNamespaceConfig) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *XClusterNamespaceConfig) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *XClusterNamespaceConfig) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *XClusterNamespaceConfig) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetNamespaceInfo

`func (o *XClusterNamespaceConfig) GetTargetNamespaceInfo() NamespaceInfoResp`

GetTargetNamespaceInfo returns the TargetNamespaceInfo field if non-nil, zero value otherwise.

### GetTargetNamespaceInfoOk

`func (o *XClusterNamespaceConfig) GetTargetNamespaceInfoOk() (*NamespaceInfoResp, bool)`

GetTargetNamespaceInfoOk returns a tuple with the TargetNamespaceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNamespaceInfo

`func (o *XClusterNamespaceConfig) SetTargetNamespaceInfo(v NamespaceInfoResp)`

SetTargetNamespaceInfo sets TargetNamespaceInfo field to given value.

### HasTargetNamespaceInfo

`func (o *XClusterNamespaceConfig) HasTargetNamespaceInfo() bool`

HasTargetNamespaceInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


