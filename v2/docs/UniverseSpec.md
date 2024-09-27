# UniverseSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Universe | 
**YbSoftwareVersion** | **string** | The YugabyteDB software version to install. This can be upgraded using API /customers/:cUUID/universes/:uniUUID/upgrade/software | 
**EncryptionAtRestSpec** | Pointer to [**EncryptionAtRestSpec**](EncryptionAtRestSpec.md) |  | [optional] 
**EncryptionInTransitSpec** | Pointer to [**EncryptionInTransitSpec**](EncryptionInTransitSpec.md) |  | [optional] 
**Ysql** | Pointer to [**YSQLSpec**](YSQLSpec.md) |  | [optional] 
**Ycql** | Pointer to [**YCQLSpec**](YCQLSpec.md) |  | [optional] 
**UseTimeSync** | Pointer to **bool** | Whether to use time sync services like chrony on DB nodes of this cluster | [optional] 
**RemotePackagePath** | Pointer to **string** | Path to download thirdparty packages for itest. Only for AWS/onprem. | [optional] 
**OverridePrebuiltAmiDbVersion** | Pointer to **bool** | Override the default DB present in pre-built Ami. YBM usage. | [optional] 
**NetworkingSpec** | Pointer to [**UniverseNetworkingSpec**](UniverseNetworkingSpec.md) |  | [optional] 
**Clusters** | [**[]ClusterSpec**](ClusterSpec.md) |  | 

## Methods

### NewUniverseSpec

`func NewUniverseSpec(name string, ybSoftwareVersion string, clusters []ClusterSpec, ) *UniverseSpec`

NewUniverseSpec instantiates a new UniverseSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseSpecWithDefaults

`func NewUniverseSpecWithDefaults() *UniverseSpec`

NewUniverseSpecWithDefaults instantiates a new UniverseSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UniverseSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UniverseSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UniverseSpec) SetName(v string)`

SetName sets Name field to given value.


### GetYbSoftwareVersion

`func (o *UniverseSpec) GetYbSoftwareVersion() string`

GetYbSoftwareVersion returns the YbSoftwareVersion field if non-nil, zero value otherwise.

### GetYbSoftwareVersionOk

`func (o *UniverseSpec) GetYbSoftwareVersionOk() (*string, bool)`

GetYbSoftwareVersionOk returns a tuple with the YbSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbSoftwareVersion

`func (o *UniverseSpec) SetYbSoftwareVersion(v string)`

SetYbSoftwareVersion sets YbSoftwareVersion field to given value.


### GetEncryptionAtRestSpec

`func (o *UniverseSpec) GetEncryptionAtRestSpec() EncryptionAtRestSpec`

GetEncryptionAtRestSpec returns the EncryptionAtRestSpec field if non-nil, zero value otherwise.

### GetEncryptionAtRestSpecOk

`func (o *UniverseSpec) GetEncryptionAtRestSpecOk() (*EncryptionAtRestSpec, bool)`

GetEncryptionAtRestSpecOk returns a tuple with the EncryptionAtRestSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestSpec

`func (o *UniverseSpec) SetEncryptionAtRestSpec(v EncryptionAtRestSpec)`

SetEncryptionAtRestSpec sets EncryptionAtRestSpec field to given value.

### HasEncryptionAtRestSpec

`func (o *UniverseSpec) HasEncryptionAtRestSpec() bool`

HasEncryptionAtRestSpec returns a boolean if a field has been set.

### GetEncryptionInTransitSpec

`func (o *UniverseSpec) GetEncryptionInTransitSpec() EncryptionInTransitSpec`

GetEncryptionInTransitSpec returns the EncryptionInTransitSpec field if non-nil, zero value otherwise.

### GetEncryptionInTransitSpecOk

`func (o *UniverseSpec) GetEncryptionInTransitSpecOk() (*EncryptionInTransitSpec, bool)`

GetEncryptionInTransitSpecOk returns a tuple with the EncryptionInTransitSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionInTransitSpec

`func (o *UniverseSpec) SetEncryptionInTransitSpec(v EncryptionInTransitSpec)`

SetEncryptionInTransitSpec sets EncryptionInTransitSpec field to given value.

### HasEncryptionInTransitSpec

`func (o *UniverseSpec) HasEncryptionInTransitSpec() bool`

HasEncryptionInTransitSpec returns a boolean if a field has been set.

### GetYsql

`func (o *UniverseSpec) GetYsql() YSQLSpec`

GetYsql returns the Ysql field if non-nil, zero value otherwise.

### GetYsqlOk

`func (o *UniverseSpec) GetYsqlOk() (*YSQLSpec, bool)`

GetYsqlOk returns a tuple with the Ysql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsql

`func (o *UniverseSpec) SetYsql(v YSQLSpec)`

SetYsql sets Ysql field to given value.

### HasYsql

`func (o *UniverseSpec) HasYsql() bool`

HasYsql returns a boolean if a field has been set.

### GetYcql

`func (o *UniverseSpec) GetYcql() YCQLSpec`

GetYcql returns the Ycql field if non-nil, zero value otherwise.

### GetYcqlOk

`func (o *UniverseSpec) GetYcqlOk() (*YCQLSpec, bool)`

GetYcqlOk returns a tuple with the Ycql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYcql

`func (o *UniverseSpec) SetYcql(v YCQLSpec)`

SetYcql sets Ycql field to given value.

### HasYcql

`func (o *UniverseSpec) HasYcql() bool`

HasYcql returns a boolean if a field has been set.

### GetUseTimeSync

`func (o *UniverseSpec) GetUseTimeSync() bool`

GetUseTimeSync returns the UseTimeSync field if non-nil, zero value otherwise.

### GetUseTimeSyncOk

`func (o *UniverseSpec) GetUseTimeSyncOk() (*bool, bool)`

GetUseTimeSyncOk returns a tuple with the UseTimeSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTimeSync

`func (o *UniverseSpec) SetUseTimeSync(v bool)`

SetUseTimeSync sets UseTimeSync field to given value.

### HasUseTimeSync

`func (o *UniverseSpec) HasUseTimeSync() bool`

HasUseTimeSync returns a boolean if a field has been set.

### GetRemotePackagePath

`func (o *UniverseSpec) GetRemotePackagePath() string`

GetRemotePackagePath returns the RemotePackagePath field if non-nil, zero value otherwise.

### GetRemotePackagePathOk

`func (o *UniverseSpec) GetRemotePackagePathOk() (*string, bool)`

GetRemotePackagePathOk returns a tuple with the RemotePackagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePackagePath

`func (o *UniverseSpec) SetRemotePackagePath(v string)`

SetRemotePackagePath sets RemotePackagePath field to given value.

### HasRemotePackagePath

`func (o *UniverseSpec) HasRemotePackagePath() bool`

HasRemotePackagePath returns a boolean if a field has been set.

### GetOverridePrebuiltAmiDbVersion

`func (o *UniverseSpec) GetOverridePrebuiltAmiDbVersion() bool`

GetOverridePrebuiltAmiDbVersion returns the OverridePrebuiltAmiDbVersion field if non-nil, zero value otherwise.

### GetOverridePrebuiltAmiDbVersionOk

`func (o *UniverseSpec) GetOverridePrebuiltAmiDbVersionOk() (*bool, bool)`

GetOverridePrebuiltAmiDbVersionOk returns a tuple with the OverridePrebuiltAmiDbVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridePrebuiltAmiDbVersion

`func (o *UniverseSpec) SetOverridePrebuiltAmiDbVersion(v bool)`

SetOverridePrebuiltAmiDbVersion sets OverridePrebuiltAmiDbVersion field to given value.

### HasOverridePrebuiltAmiDbVersion

`func (o *UniverseSpec) HasOverridePrebuiltAmiDbVersion() bool`

HasOverridePrebuiltAmiDbVersion returns a boolean if a field has been set.

### GetNetworkingSpec

`func (o *UniverseSpec) GetNetworkingSpec() UniverseNetworkingSpec`

GetNetworkingSpec returns the NetworkingSpec field if non-nil, zero value otherwise.

### GetNetworkingSpecOk

`func (o *UniverseSpec) GetNetworkingSpecOk() (*UniverseNetworkingSpec, bool)`

GetNetworkingSpecOk returns a tuple with the NetworkingSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkingSpec

`func (o *UniverseSpec) SetNetworkingSpec(v UniverseNetworkingSpec)`

SetNetworkingSpec sets NetworkingSpec field to given value.

### HasNetworkingSpec

`func (o *UniverseSpec) HasNetworkingSpec() bool`

HasNetworkingSpec returns a boolean if a field has been set.

### GetClusters

`func (o *UniverseSpec) GetClusters() []ClusterSpec`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *UniverseSpec) GetClustersOk() (*[]ClusterSpec, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *UniverseSpec) SetClusters(v []ClusterSpec)`

SetClusters sets Clusters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


