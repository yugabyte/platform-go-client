# MetricQueryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZones** | Pointer to **[]string** | Availability zone code | [optional] 
**ClusterUuids** | Pointer to **[]string** | Cluster UUIDs | [optional] 
**End** | Pointer to **int64** | End time | [optional] 
**IsRecharts** | Pointer to **bool** | Is Recharts | [optional] 
**Metrics** | Pointer to **[]string** | Metrics | [optional] 
**MetricsWithSettings** | Pointer to [**[]MetricSettings**](MetricSettings.md) | List of metrics with custom settings | [optional] 
**NodeNames** | Pointer to **[]string** | Node names | [optional] 
**NodePrefix** | Pointer to **string** | Node prefix | [optional] 
**RegionCodes** | Pointer to **[]string** | Region code | [optional] 
**ServerType** | Pointer to **string** | Server type | [optional] 
**Start** | **int64** | Start time | 
**TableId** | Pointer to **string** | Table id | [optional] 
**TableName** | Pointer to **string** | Table name | [optional] 
**XclusterConfigUuid** | **string** |  | 

## Methods

### NewMetricQueryParams

`func NewMetricQueryParams(start int64, xclusterConfigUuid string, ) *MetricQueryParams`

NewMetricQueryParams instantiates a new MetricQueryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricQueryParamsWithDefaults

`func NewMetricQueryParamsWithDefaults() *MetricQueryParams`

NewMetricQueryParamsWithDefaults instantiates a new MetricQueryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZones

`func (o *MetricQueryParams) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *MetricQueryParams) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *MetricQueryParams) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.

### HasAvailabilityZones

`func (o *MetricQueryParams) HasAvailabilityZones() bool`

HasAvailabilityZones returns a boolean if a field has been set.

### GetClusterUuids

`func (o *MetricQueryParams) GetClusterUuids() []string`

GetClusterUuids returns the ClusterUuids field if non-nil, zero value otherwise.

### GetClusterUuidsOk

`func (o *MetricQueryParams) GetClusterUuidsOk() (*[]string, bool)`

GetClusterUuidsOk returns a tuple with the ClusterUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuids

`func (o *MetricQueryParams) SetClusterUuids(v []string)`

SetClusterUuids sets ClusterUuids field to given value.

### HasClusterUuids

`func (o *MetricQueryParams) HasClusterUuids() bool`

HasClusterUuids returns a boolean if a field has been set.

### GetEnd

`func (o *MetricQueryParams) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *MetricQueryParams) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *MetricQueryParams) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *MetricQueryParams) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetIsRecharts

`func (o *MetricQueryParams) GetIsRecharts() bool`

GetIsRecharts returns the IsRecharts field if non-nil, zero value otherwise.

### GetIsRechartsOk

`func (o *MetricQueryParams) GetIsRechartsOk() (*bool, bool)`

GetIsRechartsOk returns a tuple with the IsRecharts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecharts

`func (o *MetricQueryParams) SetIsRecharts(v bool)`

SetIsRecharts sets IsRecharts field to given value.

### HasIsRecharts

`func (o *MetricQueryParams) HasIsRecharts() bool`

HasIsRecharts returns a boolean if a field has been set.

### GetMetrics

`func (o *MetricQueryParams) GetMetrics() []string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MetricQueryParams) GetMetricsOk() (*[]string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MetricQueryParams) SetMetrics(v []string)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *MetricQueryParams) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMetricsWithSettings

`func (o *MetricQueryParams) GetMetricsWithSettings() []MetricSettings`

GetMetricsWithSettings returns the MetricsWithSettings field if non-nil, zero value otherwise.

### GetMetricsWithSettingsOk

`func (o *MetricQueryParams) GetMetricsWithSettingsOk() (*[]MetricSettings, bool)`

GetMetricsWithSettingsOk returns a tuple with the MetricsWithSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsWithSettings

`func (o *MetricQueryParams) SetMetricsWithSettings(v []MetricSettings)`

SetMetricsWithSettings sets MetricsWithSettings field to given value.

### HasMetricsWithSettings

`func (o *MetricQueryParams) HasMetricsWithSettings() bool`

HasMetricsWithSettings returns a boolean if a field has been set.

### GetNodeNames

`func (o *MetricQueryParams) GetNodeNames() []string`

GetNodeNames returns the NodeNames field if non-nil, zero value otherwise.

### GetNodeNamesOk

`func (o *MetricQueryParams) GetNodeNamesOk() (*[]string, bool)`

GetNodeNamesOk returns a tuple with the NodeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeNames

`func (o *MetricQueryParams) SetNodeNames(v []string)`

SetNodeNames sets NodeNames field to given value.

### HasNodeNames

`func (o *MetricQueryParams) HasNodeNames() bool`

HasNodeNames returns a boolean if a field has been set.

### GetNodePrefix

`func (o *MetricQueryParams) GetNodePrefix() string`

GetNodePrefix returns the NodePrefix field if non-nil, zero value otherwise.

### GetNodePrefixOk

`func (o *MetricQueryParams) GetNodePrefixOk() (*string, bool)`

GetNodePrefixOk returns a tuple with the NodePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePrefix

`func (o *MetricQueryParams) SetNodePrefix(v string)`

SetNodePrefix sets NodePrefix field to given value.

### HasNodePrefix

`func (o *MetricQueryParams) HasNodePrefix() bool`

HasNodePrefix returns a boolean if a field has been set.

### GetRegionCodes

`func (o *MetricQueryParams) GetRegionCodes() []string`

GetRegionCodes returns the RegionCodes field if non-nil, zero value otherwise.

### GetRegionCodesOk

`func (o *MetricQueryParams) GetRegionCodesOk() (*[]string, bool)`

GetRegionCodesOk returns a tuple with the RegionCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCodes

`func (o *MetricQueryParams) SetRegionCodes(v []string)`

SetRegionCodes sets RegionCodes field to given value.

### HasRegionCodes

`func (o *MetricQueryParams) HasRegionCodes() bool`

HasRegionCodes returns a boolean if a field has been set.

### GetServerType

`func (o *MetricQueryParams) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *MetricQueryParams) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *MetricQueryParams) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *MetricQueryParams) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetStart

`func (o *MetricQueryParams) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *MetricQueryParams) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *MetricQueryParams) SetStart(v int64)`

SetStart sets Start field to given value.


### GetTableId

`func (o *MetricQueryParams) GetTableId() string`

GetTableId returns the TableId field if non-nil, zero value otherwise.

### GetTableIdOk

`func (o *MetricQueryParams) GetTableIdOk() (*string, bool)`

GetTableIdOk returns a tuple with the TableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableId

`func (o *MetricQueryParams) SetTableId(v string)`

SetTableId sets TableId field to given value.

### HasTableId

`func (o *MetricQueryParams) HasTableId() bool`

HasTableId returns a boolean if a field has been set.

### GetTableName

`func (o *MetricQueryParams) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *MetricQueryParams) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *MetricQueryParams) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *MetricQueryParams) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetXclusterConfigUuid

`func (o *MetricQueryParams) GetXclusterConfigUuid() string`

GetXclusterConfigUuid returns the XclusterConfigUuid field if non-nil, zero value otherwise.

### GetXclusterConfigUuidOk

`func (o *MetricQueryParams) GetXclusterConfigUuidOk() (*string, bool)`

GetXclusterConfigUuidOk returns a tuple with the XclusterConfigUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXclusterConfigUuid

`func (o *MetricQueryParams) SetXclusterConfigUuid(v string)`

SetXclusterConfigUuid sets XclusterConfigUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


