# UniverseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniverseUuid** | Pointer to **string** | UUID of the Universe | [optional] [readonly] 
**Version** | Pointer to **int32** | Universe version | [optional] [readonly] 
**CreationDate** | Pointer to **time.Time** | Universe creation date | [optional] [readonly] 
**Arch** | Pointer to **string** | CPU Arch of DB nodes. | [optional] [readonly] 
**DnsName** | Pointer to **string** | DNS name | [optional] [readonly] 
**YbcSoftwareVersion** | Pointer to **string** | YBC Software version installed in DB nodes of this Universe | [optional] [readonly] 
**PlatformUrl** | Pointer to **string** | YBA that is managing this Universe | [optional] [readonly] 
**NodePrefix** | Pointer to **string** | A globally unique name generated as a combination of the customer id and the universe name. This is used as the prefix of node names in the universe. Can be configured at the time of universe creation. | [optional] [readonly] 
**EncryptionAtRestInfo** | Pointer to [**EncryptionAtRestInfo**](EncryptionAtRestInfo.md) |  | [optional] 
**UpdateInProgress** | Pointer to **bool** | Whether a create/edit/destroy intent on the universe is currently running. | [optional] [readonly] 
**UpdatingTask** | Pointer to **string** | Type of task which set updateInProgress flag. | [optional] [readonly] 
**UpdatingTaskUuid** | Pointer to **string** | UUID of task which set updateInProgress flag. | [optional] [readonly] 
**UpdateSucceeded** | Pointer to **bool** | Whether the latest operation on this universe has successfully completed. Is updated for each operation on the universe. | [optional] [readonly] 
**UniversePaused** | Pointer to **bool** | Whether the universe is in the paused state | [optional] [readonly] 
**PlacementModificationTaskUuid** | Pointer to **string** | UUID of last failed task that applied modification to cluster state | [optional] [readonly] 
**SoftwareUpgradeState** | Pointer to **string** | The state of the last YugabyteDB software upgrade operation on this universe | [optional] [readonly] 
**IsSoftwareRollbackAllowed** | Pointer to **bool** | Whether a rollback of the last YugabyteDB upgrade operation is allowed | [optional] [readonly] 
**NodesResizeAvailable** | Pointer to **bool** | Set to true if nodes of this Universe can be resized without a full move | [optional] [readonly] 
**IsKubernetesOperatorControlled** | Pointer to **bool** | Whether this Universe is created and controlled by the Kubernetes Operator | [optional] [readonly] 
**OtelCollectorEnabled** | Pointer to **bool** | Whether OpenTelemetry Collector is enabled for universe | [optional] 
**XClusterInfo** | Pointer to [**XClusterInfo**](XClusterInfo.md) |  | [optional] 
**Clusters** | Pointer to [**[]ClusterInfo**](ClusterInfo.md) |  | [optional] 
**NodeDetailsSet** | Pointer to [**[]NodeDetails**](NodeDetails.md) | Node details | [optional] 

## Methods

### NewUniverseInfo

`func NewUniverseInfo() *UniverseInfo`

NewUniverseInfo instantiates a new UniverseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseInfoWithDefaults

`func NewUniverseInfoWithDefaults() *UniverseInfo`

NewUniverseInfoWithDefaults instantiates a new UniverseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniverseUuid

`func (o *UniverseInfo) GetUniverseUuid() string`

GetUniverseUuid returns the UniverseUuid field if non-nil, zero value otherwise.

### GetUniverseUuidOk

`func (o *UniverseInfo) GetUniverseUuidOk() (*string, bool)`

GetUniverseUuidOk returns a tuple with the UniverseUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUuid

`func (o *UniverseInfo) SetUniverseUuid(v string)`

SetUniverseUuid sets UniverseUuid field to given value.

### HasUniverseUuid

`func (o *UniverseInfo) HasUniverseUuid() bool`

HasUniverseUuid returns a boolean if a field has been set.

### GetVersion

`func (o *UniverseInfo) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UniverseInfo) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UniverseInfo) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UniverseInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreationDate

`func (o *UniverseInfo) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *UniverseInfo) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *UniverseInfo) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *UniverseInfo) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetArch

`func (o *UniverseInfo) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *UniverseInfo) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *UniverseInfo) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *UniverseInfo) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetDnsName

`func (o *UniverseInfo) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *UniverseInfo) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *UniverseInfo) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *UniverseInfo) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetYbcSoftwareVersion

`func (o *UniverseInfo) GetYbcSoftwareVersion() string`

GetYbcSoftwareVersion returns the YbcSoftwareVersion field if non-nil, zero value otherwise.

### GetYbcSoftwareVersionOk

`func (o *UniverseInfo) GetYbcSoftwareVersionOk() (*string, bool)`

GetYbcSoftwareVersionOk returns a tuple with the YbcSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcSoftwareVersion

`func (o *UniverseInfo) SetYbcSoftwareVersion(v string)`

SetYbcSoftwareVersion sets YbcSoftwareVersion field to given value.

### HasYbcSoftwareVersion

`func (o *UniverseInfo) HasYbcSoftwareVersion() bool`

HasYbcSoftwareVersion returns a boolean if a field has been set.

### GetPlatformUrl

`func (o *UniverseInfo) GetPlatformUrl() string`

GetPlatformUrl returns the PlatformUrl field if non-nil, zero value otherwise.

### GetPlatformUrlOk

`func (o *UniverseInfo) GetPlatformUrlOk() (*string, bool)`

GetPlatformUrlOk returns a tuple with the PlatformUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformUrl

`func (o *UniverseInfo) SetPlatformUrl(v string)`

SetPlatformUrl sets PlatformUrl field to given value.

### HasPlatformUrl

`func (o *UniverseInfo) HasPlatformUrl() bool`

HasPlatformUrl returns a boolean if a field has been set.

### GetNodePrefix

`func (o *UniverseInfo) GetNodePrefix() string`

GetNodePrefix returns the NodePrefix field if non-nil, zero value otherwise.

### GetNodePrefixOk

`func (o *UniverseInfo) GetNodePrefixOk() (*string, bool)`

GetNodePrefixOk returns a tuple with the NodePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePrefix

`func (o *UniverseInfo) SetNodePrefix(v string)`

SetNodePrefix sets NodePrefix field to given value.

### HasNodePrefix

`func (o *UniverseInfo) HasNodePrefix() bool`

HasNodePrefix returns a boolean if a field has been set.

### GetEncryptionAtRestInfo

`func (o *UniverseInfo) GetEncryptionAtRestInfo() EncryptionAtRestInfo`

GetEncryptionAtRestInfo returns the EncryptionAtRestInfo field if non-nil, zero value otherwise.

### GetEncryptionAtRestInfoOk

`func (o *UniverseInfo) GetEncryptionAtRestInfoOk() (*EncryptionAtRestInfo, bool)`

GetEncryptionAtRestInfoOk returns a tuple with the EncryptionAtRestInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAtRestInfo

`func (o *UniverseInfo) SetEncryptionAtRestInfo(v EncryptionAtRestInfo)`

SetEncryptionAtRestInfo sets EncryptionAtRestInfo field to given value.

### HasEncryptionAtRestInfo

`func (o *UniverseInfo) HasEncryptionAtRestInfo() bool`

HasEncryptionAtRestInfo returns a boolean if a field has been set.

### GetUpdateInProgress

`func (o *UniverseInfo) GetUpdateInProgress() bool`

GetUpdateInProgress returns the UpdateInProgress field if non-nil, zero value otherwise.

### GetUpdateInProgressOk

`func (o *UniverseInfo) GetUpdateInProgressOk() (*bool, bool)`

GetUpdateInProgressOk returns a tuple with the UpdateInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInProgress

`func (o *UniverseInfo) SetUpdateInProgress(v bool)`

SetUpdateInProgress sets UpdateInProgress field to given value.

### HasUpdateInProgress

`func (o *UniverseInfo) HasUpdateInProgress() bool`

HasUpdateInProgress returns a boolean if a field has been set.

### GetUpdatingTask

`func (o *UniverseInfo) GetUpdatingTask() string`

GetUpdatingTask returns the UpdatingTask field if non-nil, zero value otherwise.

### GetUpdatingTaskOk

`func (o *UniverseInfo) GetUpdatingTaskOk() (*string, bool)`

GetUpdatingTaskOk returns a tuple with the UpdatingTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatingTask

`func (o *UniverseInfo) SetUpdatingTask(v string)`

SetUpdatingTask sets UpdatingTask field to given value.

### HasUpdatingTask

`func (o *UniverseInfo) HasUpdatingTask() bool`

HasUpdatingTask returns a boolean if a field has been set.

### GetUpdatingTaskUuid

`func (o *UniverseInfo) GetUpdatingTaskUuid() string`

GetUpdatingTaskUuid returns the UpdatingTaskUuid field if non-nil, zero value otherwise.

### GetUpdatingTaskUuidOk

`func (o *UniverseInfo) GetUpdatingTaskUuidOk() (*string, bool)`

GetUpdatingTaskUuidOk returns a tuple with the UpdatingTaskUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatingTaskUuid

`func (o *UniverseInfo) SetUpdatingTaskUuid(v string)`

SetUpdatingTaskUuid sets UpdatingTaskUuid field to given value.

### HasUpdatingTaskUuid

`func (o *UniverseInfo) HasUpdatingTaskUuid() bool`

HasUpdatingTaskUuid returns a boolean if a field has been set.

### GetUpdateSucceeded

`func (o *UniverseInfo) GetUpdateSucceeded() bool`

GetUpdateSucceeded returns the UpdateSucceeded field if non-nil, zero value otherwise.

### GetUpdateSucceededOk

`func (o *UniverseInfo) GetUpdateSucceededOk() (*bool, bool)`

GetUpdateSucceededOk returns a tuple with the UpdateSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSucceeded

`func (o *UniverseInfo) SetUpdateSucceeded(v bool)`

SetUpdateSucceeded sets UpdateSucceeded field to given value.

### HasUpdateSucceeded

`func (o *UniverseInfo) HasUpdateSucceeded() bool`

HasUpdateSucceeded returns a boolean if a field has been set.

### GetUniversePaused

`func (o *UniverseInfo) GetUniversePaused() bool`

GetUniversePaused returns the UniversePaused field if non-nil, zero value otherwise.

### GetUniversePausedOk

`func (o *UniverseInfo) GetUniversePausedOk() (*bool, bool)`

GetUniversePausedOk returns a tuple with the UniversePaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversePaused

`func (o *UniverseInfo) SetUniversePaused(v bool)`

SetUniversePaused sets UniversePaused field to given value.

### HasUniversePaused

`func (o *UniverseInfo) HasUniversePaused() bool`

HasUniversePaused returns a boolean if a field has been set.

### GetPlacementModificationTaskUuid

`func (o *UniverseInfo) GetPlacementModificationTaskUuid() string`

GetPlacementModificationTaskUuid returns the PlacementModificationTaskUuid field if non-nil, zero value otherwise.

### GetPlacementModificationTaskUuidOk

`func (o *UniverseInfo) GetPlacementModificationTaskUuidOk() (*string, bool)`

GetPlacementModificationTaskUuidOk returns a tuple with the PlacementModificationTaskUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementModificationTaskUuid

`func (o *UniverseInfo) SetPlacementModificationTaskUuid(v string)`

SetPlacementModificationTaskUuid sets PlacementModificationTaskUuid field to given value.

### HasPlacementModificationTaskUuid

`func (o *UniverseInfo) HasPlacementModificationTaskUuid() bool`

HasPlacementModificationTaskUuid returns a boolean if a field has been set.

### GetSoftwareUpgradeState

`func (o *UniverseInfo) GetSoftwareUpgradeState() string`

GetSoftwareUpgradeState returns the SoftwareUpgradeState field if non-nil, zero value otherwise.

### GetSoftwareUpgradeStateOk

`func (o *UniverseInfo) GetSoftwareUpgradeStateOk() (*string, bool)`

GetSoftwareUpgradeStateOk returns a tuple with the SoftwareUpgradeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpgradeState

`func (o *UniverseInfo) SetSoftwareUpgradeState(v string)`

SetSoftwareUpgradeState sets SoftwareUpgradeState field to given value.

### HasSoftwareUpgradeState

`func (o *UniverseInfo) HasSoftwareUpgradeState() bool`

HasSoftwareUpgradeState returns a boolean if a field has been set.

### GetIsSoftwareRollbackAllowed

`func (o *UniverseInfo) GetIsSoftwareRollbackAllowed() bool`

GetIsSoftwareRollbackAllowed returns the IsSoftwareRollbackAllowed field if non-nil, zero value otherwise.

### GetIsSoftwareRollbackAllowedOk

`func (o *UniverseInfo) GetIsSoftwareRollbackAllowedOk() (*bool, bool)`

GetIsSoftwareRollbackAllowedOk returns a tuple with the IsSoftwareRollbackAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSoftwareRollbackAllowed

`func (o *UniverseInfo) SetIsSoftwareRollbackAllowed(v bool)`

SetIsSoftwareRollbackAllowed sets IsSoftwareRollbackAllowed field to given value.

### HasIsSoftwareRollbackAllowed

`func (o *UniverseInfo) HasIsSoftwareRollbackAllowed() bool`

HasIsSoftwareRollbackAllowed returns a boolean if a field has been set.

### GetNodesResizeAvailable

`func (o *UniverseInfo) GetNodesResizeAvailable() bool`

GetNodesResizeAvailable returns the NodesResizeAvailable field if non-nil, zero value otherwise.

### GetNodesResizeAvailableOk

`func (o *UniverseInfo) GetNodesResizeAvailableOk() (*bool, bool)`

GetNodesResizeAvailableOk returns a tuple with the NodesResizeAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesResizeAvailable

`func (o *UniverseInfo) SetNodesResizeAvailable(v bool)`

SetNodesResizeAvailable sets NodesResizeAvailable field to given value.

### HasNodesResizeAvailable

`func (o *UniverseInfo) HasNodesResizeAvailable() bool`

HasNodesResizeAvailable returns a boolean if a field has been set.

### GetIsKubernetesOperatorControlled

`func (o *UniverseInfo) GetIsKubernetesOperatorControlled() bool`

GetIsKubernetesOperatorControlled returns the IsKubernetesOperatorControlled field if non-nil, zero value otherwise.

### GetIsKubernetesOperatorControlledOk

`func (o *UniverseInfo) GetIsKubernetesOperatorControlledOk() (*bool, bool)`

GetIsKubernetesOperatorControlledOk returns a tuple with the IsKubernetesOperatorControlled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKubernetesOperatorControlled

`func (o *UniverseInfo) SetIsKubernetesOperatorControlled(v bool)`

SetIsKubernetesOperatorControlled sets IsKubernetesOperatorControlled field to given value.

### HasIsKubernetesOperatorControlled

`func (o *UniverseInfo) HasIsKubernetesOperatorControlled() bool`

HasIsKubernetesOperatorControlled returns a boolean if a field has been set.

### GetOtelCollectorEnabled

`func (o *UniverseInfo) GetOtelCollectorEnabled() bool`

GetOtelCollectorEnabled returns the OtelCollectorEnabled field if non-nil, zero value otherwise.

### GetOtelCollectorEnabledOk

`func (o *UniverseInfo) GetOtelCollectorEnabledOk() (*bool, bool)`

GetOtelCollectorEnabledOk returns a tuple with the OtelCollectorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtelCollectorEnabled

`func (o *UniverseInfo) SetOtelCollectorEnabled(v bool)`

SetOtelCollectorEnabled sets OtelCollectorEnabled field to given value.

### HasOtelCollectorEnabled

`func (o *UniverseInfo) HasOtelCollectorEnabled() bool`

HasOtelCollectorEnabled returns a boolean if a field has been set.

### GetXClusterInfo

`func (o *UniverseInfo) GetXClusterInfo() XClusterInfo`

GetXClusterInfo returns the XClusterInfo field if non-nil, zero value otherwise.

### GetXClusterInfoOk

`func (o *UniverseInfo) GetXClusterInfoOk() (*XClusterInfo, bool)`

GetXClusterInfoOk returns a tuple with the XClusterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXClusterInfo

`func (o *UniverseInfo) SetXClusterInfo(v XClusterInfo)`

SetXClusterInfo sets XClusterInfo field to given value.

### HasXClusterInfo

`func (o *UniverseInfo) HasXClusterInfo() bool`

HasXClusterInfo returns a boolean if a field has been set.

### GetClusters

`func (o *UniverseInfo) GetClusters() []ClusterInfo`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *UniverseInfo) GetClustersOk() (*[]ClusterInfo, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *UniverseInfo) SetClusters(v []ClusterInfo)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *UniverseInfo) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetNodeDetailsSet

`func (o *UniverseInfo) GetNodeDetailsSet() []NodeDetails`

GetNodeDetailsSet returns the NodeDetailsSet field if non-nil, zero value otherwise.

### GetNodeDetailsSetOk

`func (o *UniverseInfo) GetNodeDetailsSetOk() (*[]NodeDetails, bool)`

GetNodeDetailsSetOk returns a tuple with the NodeDetailsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeDetailsSet

`func (o *UniverseInfo) SetNodeDetailsSet(v []NodeDetails)`

SetNodeDetailsSet sets NodeDetailsSet field to given value.

### HasNodeDetailsSet

`func (o *UniverseInfo) HasNodeDetailsSet() bool`

HasNodeDetailsSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


