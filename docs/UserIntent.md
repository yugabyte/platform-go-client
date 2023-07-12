# UserIntent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyCode** | Pointer to **string** |  | [optional] 
**AssignPublicIP** | Pointer to **bool** |  | [optional] 
**AssignStaticPublicIP** | Pointer to **bool** | Whether to assign static public IP | [optional] 
**AwsArnString** | Pointer to **string** |  | [optional] 
**AzOverrides** | Pointer to **map[string]string** |  | [optional] 
**DedicatedNodes** | Pointer to **bool** |  | [optional] 
**DeviceInfo** | Pointer to [**DeviceInfo**](DeviceInfo.md) |  | [optional] 
**EnableClientToNodeEncrypt** | Pointer to **bool** |  | [optional] 
**EnableExposingService** | Pointer to **string** |  | [optional] 
**EnableIPV6** | Pointer to **bool** |  | [optional] 
**EnableLB** | Pointer to **bool** |  | [optional] 
**EnableNodeToNodeEncrypt** | Pointer to **bool** |  | [optional] 
**EnableVolumeEncryption** | Pointer to **bool** |  | [optional] 
**EnableYCQL** | Pointer to **bool** |  | [optional] 
**EnableYCQLAuth** | Pointer to **bool** |  | [optional] 
**EnableYEDIS** | Pointer to **bool** |  | [optional] 
**EnableYSQL** | Pointer to **bool** |  | [optional] 
**EnableYSQLAuth** | Pointer to **bool** |  | [optional] 
**ImageBundleUUID** | Pointer to **string** |  | [optional] 
**InstanceTags** | Pointer to **map[string]string** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**KubernetesOperatorVersion** | Pointer to **int64** |  | [optional] 
**MasterDeviceInfo** | Pointer to [**DeviceInfo**](DeviceInfo.md) |  | [optional] 
**MasterGFlags** | Pointer to **map[string]string** |  | [optional] 
**MasterInstanceType** | Pointer to **string** |  | [optional] 
**NumNodes** | Pointer to **int32** |  | [optional] 
**PreferredRegion** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**ProviderType** | Pointer to **string** |  | [optional] 
**RegionList** | Pointer to **[]string** |  | [optional] 
**ReplicationFactor** | Pointer to **int32** |  | [optional] 
**SpecificGFlags** | Pointer to [**SpecificGFlags**](SpecificGFlags.md) |  | [optional] 
**SpotPrice** | Pointer to **float64** |  | [optional] 
**TserverGFlags** | Pointer to **map[string]string** |  | [optional] 
**UniverseName** | Pointer to **string** |  | [optional] 
**UniverseOverrides** | Pointer to **string** |  | [optional] 
**UseHostname** | Pointer to **bool** |  | [optional] 
**UseSpotInstance** | Pointer to **bool** |  | [optional] 
**UseSystemd** | Pointer to **bool** |  | [optional] 
**UseTimeSync** | Pointer to **bool** |  | [optional] 
**YbSoftwareVersion** | Pointer to **string** |  | [optional] 
**YbcFlags** | Pointer to **map[string]string** |  | [optional] 
**YcqlPassword** | Pointer to **string** |  | [optional] 
**YsqlPassword** | Pointer to **string** |  | [optional] 

## Methods

### NewUserIntent

`func NewUserIntent() *UserIntent`

NewUserIntent instantiates a new UserIntent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIntentWithDefaults

`func NewUserIntentWithDefaults() *UserIntent`

NewUserIntentWithDefaults instantiates a new UserIntent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyCode

`func (o *UserIntent) GetAccessKeyCode() string`

GetAccessKeyCode returns the AccessKeyCode field if non-nil, zero value otherwise.

### GetAccessKeyCodeOk

`func (o *UserIntent) GetAccessKeyCodeOk() (*string, bool)`

GetAccessKeyCodeOk returns a tuple with the AccessKeyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyCode

`func (o *UserIntent) SetAccessKeyCode(v string)`

SetAccessKeyCode sets AccessKeyCode field to given value.

### HasAccessKeyCode

`func (o *UserIntent) HasAccessKeyCode() bool`

HasAccessKeyCode returns a boolean if a field has been set.

### GetAssignPublicIP

`func (o *UserIntent) GetAssignPublicIP() bool`

GetAssignPublicIP returns the AssignPublicIP field if non-nil, zero value otherwise.

### GetAssignPublicIPOk

`func (o *UserIntent) GetAssignPublicIPOk() (*bool, bool)`

GetAssignPublicIPOk returns a tuple with the AssignPublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignPublicIP

`func (o *UserIntent) SetAssignPublicIP(v bool)`

SetAssignPublicIP sets AssignPublicIP field to given value.

### HasAssignPublicIP

`func (o *UserIntent) HasAssignPublicIP() bool`

HasAssignPublicIP returns a boolean if a field has been set.

### GetAssignStaticPublicIP

`func (o *UserIntent) GetAssignStaticPublicIP() bool`

GetAssignStaticPublicIP returns the AssignStaticPublicIP field if non-nil, zero value otherwise.

### GetAssignStaticPublicIPOk

`func (o *UserIntent) GetAssignStaticPublicIPOk() (*bool, bool)`

GetAssignStaticPublicIPOk returns a tuple with the AssignStaticPublicIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignStaticPublicIP

`func (o *UserIntent) SetAssignStaticPublicIP(v bool)`

SetAssignStaticPublicIP sets AssignStaticPublicIP field to given value.

### HasAssignStaticPublicIP

`func (o *UserIntent) HasAssignStaticPublicIP() bool`

HasAssignStaticPublicIP returns a boolean if a field has been set.

### GetAwsArnString

`func (o *UserIntent) GetAwsArnString() string`

GetAwsArnString returns the AwsArnString field if non-nil, zero value otherwise.

### GetAwsArnStringOk

`func (o *UserIntent) GetAwsArnStringOk() (*string, bool)`

GetAwsArnStringOk returns a tuple with the AwsArnString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsArnString

`func (o *UserIntent) SetAwsArnString(v string)`

SetAwsArnString sets AwsArnString field to given value.

### HasAwsArnString

`func (o *UserIntent) HasAwsArnString() bool`

HasAwsArnString returns a boolean if a field has been set.

### GetAzOverrides

`func (o *UserIntent) GetAzOverrides() map[string]string`

GetAzOverrides returns the AzOverrides field if non-nil, zero value otherwise.

### GetAzOverridesOk

`func (o *UserIntent) GetAzOverridesOk() (*map[string]string, bool)`

GetAzOverridesOk returns a tuple with the AzOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzOverrides

`func (o *UserIntent) SetAzOverrides(v map[string]string)`

SetAzOverrides sets AzOverrides field to given value.

### HasAzOverrides

`func (o *UserIntent) HasAzOverrides() bool`

HasAzOverrides returns a boolean if a field has been set.

### GetDedicatedNodes

`func (o *UserIntent) GetDedicatedNodes() bool`

GetDedicatedNodes returns the DedicatedNodes field if non-nil, zero value otherwise.

### GetDedicatedNodesOk

`func (o *UserIntent) GetDedicatedNodesOk() (*bool, bool)`

GetDedicatedNodesOk returns a tuple with the DedicatedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedNodes

`func (o *UserIntent) SetDedicatedNodes(v bool)`

SetDedicatedNodes sets DedicatedNodes field to given value.

### HasDedicatedNodes

`func (o *UserIntent) HasDedicatedNodes() bool`

HasDedicatedNodes returns a boolean if a field has been set.

### GetDeviceInfo

`func (o *UserIntent) GetDeviceInfo() DeviceInfo`

GetDeviceInfo returns the DeviceInfo field if non-nil, zero value otherwise.

### GetDeviceInfoOk

`func (o *UserIntent) GetDeviceInfoOk() (*DeviceInfo, bool)`

GetDeviceInfoOk returns a tuple with the DeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceInfo

`func (o *UserIntent) SetDeviceInfo(v DeviceInfo)`

SetDeviceInfo sets DeviceInfo field to given value.

### HasDeviceInfo

`func (o *UserIntent) HasDeviceInfo() bool`

HasDeviceInfo returns a boolean if a field has been set.

### GetEnableClientToNodeEncrypt

`func (o *UserIntent) GetEnableClientToNodeEncrypt() bool`

GetEnableClientToNodeEncrypt returns the EnableClientToNodeEncrypt field if non-nil, zero value otherwise.

### GetEnableClientToNodeEncryptOk

`func (o *UserIntent) GetEnableClientToNodeEncryptOk() (*bool, bool)`

GetEnableClientToNodeEncryptOk returns a tuple with the EnableClientToNodeEncrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientToNodeEncrypt

`func (o *UserIntent) SetEnableClientToNodeEncrypt(v bool)`

SetEnableClientToNodeEncrypt sets EnableClientToNodeEncrypt field to given value.

### HasEnableClientToNodeEncrypt

`func (o *UserIntent) HasEnableClientToNodeEncrypt() bool`

HasEnableClientToNodeEncrypt returns a boolean if a field has been set.

### GetEnableExposingService

`func (o *UserIntent) GetEnableExposingService() string`

GetEnableExposingService returns the EnableExposingService field if non-nil, zero value otherwise.

### GetEnableExposingServiceOk

`func (o *UserIntent) GetEnableExposingServiceOk() (*string, bool)`

GetEnableExposingServiceOk returns a tuple with the EnableExposingService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExposingService

`func (o *UserIntent) SetEnableExposingService(v string)`

SetEnableExposingService sets EnableExposingService field to given value.

### HasEnableExposingService

`func (o *UserIntent) HasEnableExposingService() bool`

HasEnableExposingService returns a boolean if a field has been set.

### GetEnableIPV6

`func (o *UserIntent) GetEnableIPV6() bool`

GetEnableIPV6 returns the EnableIPV6 field if non-nil, zero value otherwise.

### GetEnableIPV6Ok

`func (o *UserIntent) GetEnableIPV6Ok() (*bool, bool)`

GetEnableIPV6Ok returns a tuple with the EnableIPV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIPV6

`func (o *UserIntent) SetEnableIPV6(v bool)`

SetEnableIPV6 sets EnableIPV6 field to given value.

### HasEnableIPV6

`func (o *UserIntent) HasEnableIPV6() bool`

HasEnableIPV6 returns a boolean if a field has been set.

### GetEnableLB

`func (o *UserIntent) GetEnableLB() bool`

GetEnableLB returns the EnableLB field if non-nil, zero value otherwise.

### GetEnableLBOk

`func (o *UserIntent) GetEnableLBOk() (*bool, bool)`

GetEnableLBOk returns a tuple with the EnableLB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLB

`func (o *UserIntent) SetEnableLB(v bool)`

SetEnableLB sets EnableLB field to given value.

### HasEnableLB

`func (o *UserIntent) HasEnableLB() bool`

HasEnableLB returns a boolean if a field has been set.

### GetEnableNodeToNodeEncrypt

`func (o *UserIntent) GetEnableNodeToNodeEncrypt() bool`

GetEnableNodeToNodeEncrypt returns the EnableNodeToNodeEncrypt field if non-nil, zero value otherwise.

### GetEnableNodeToNodeEncryptOk

`func (o *UserIntent) GetEnableNodeToNodeEncryptOk() (*bool, bool)`

GetEnableNodeToNodeEncryptOk returns a tuple with the EnableNodeToNodeEncrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNodeToNodeEncrypt

`func (o *UserIntent) SetEnableNodeToNodeEncrypt(v bool)`

SetEnableNodeToNodeEncrypt sets EnableNodeToNodeEncrypt field to given value.

### HasEnableNodeToNodeEncrypt

`func (o *UserIntent) HasEnableNodeToNodeEncrypt() bool`

HasEnableNodeToNodeEncrypt returns a boolean if a field has been set.

### GetEnableVolumeEncryption

`func (o *UserIntent) GetEnableVolumeEncryption() bool`

GetEnableVolumeEncryption returns the EnableVolumeEncryption field if non-nil, zero value otherwise.

### GetEnableVolumeEncryptionOk

`func (o *UserIntent) GetEnableVolumeEncryptionOk() (*bool, bool)`

GetEnableVolumeEncryptionOk returns a tuple with the EnableVolumeEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVolumeEncryption

`func (o *UserIntent) SetEnableVolumeEncryption(v bool)`

SetEnableVolumeEncryption sets EnableVolumeEncryption field to given value.

### HasEnableVolumeEncryption

`func (o *UserIntent) HasEnableVolumeEncryption() bool`

HasEnableVolumeEncryption returns a boolean if a field has been set.

### GetEnableYCQL

`func (o *UserIntent) GetEnableYCQL() bool`

GetEnableYCQL returns the EnableYCQL field if non-nil, zero value otherwise.

### GetEnableYCQLOk

`func (o *UserIntent) GetEnableYCQLOk() (*bool, bool)`

GetEnableYCQLOk returns a tuple with the EnableYCQL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableYCQL

`func (o *UserIntent) SetEnableYCQL(v bool)`

SetEnableYCQL sets EnableYCQL field to given value.

### HasEnableYCQL

`func (o *UserIntent) HasEnableYCQL() bool`

HasEnableYCQL returns a boolean if a field has been set.

### GetEnableYCQLAuth

`func (o *UserIntent) GetEnableYCQLAuth() bool`

GetEnableYCQLAuth returns the EnableYCQLAuth field if non-nil, zero value otherwise.

### GetEnableYCQLAuthOk

`func (o *UserIntent) GetEnableYCQLAuthOk() (*bool, bool)`

GetEnableYCQLAuthOk returns a tuple with the EnableYCQLAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableYCQLAuth

`func (o *UserIntent) SetEnableYCQLAuth(v bool)`

SetEnableYCQLAuth sets EnableYCQLAuth field to given value.

### HasEnableYCQLAuth

`func (o *UserIntent) HasEnableYCQLAuth() bool`

HasEnableYCQLAuth returns a boolean if a field has been set.

### GetEnableYEDIS

`func (o *UserIntent) GetEnableYEDIS() bool`

GetEnableYEDIS returns the EnableYEDIS field if non-nil, zero value otherwise.

### GetEnableYEDISOk

`func (o *UserIntent) GetEnableYEDISOk() (*bool, bool)`

GetEnableYEDISOk returns a tuple with the EnableYEDIS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableYEDIS

`func (o *UserIntent) SetEnableYEDIS(v bool)`

SetEnableYEDIS sets EnableYEDIS field to given value.

### HasEnableYEDIS

`func (o *UserIntent) HasEnableYEDIS() bool`

HasEnableYEDIS returns a boolean if a field has been set.

### GetEnableYSQL

`func (o *UserIntent) GetEnableYSQL() bool`

GetEnableYSQL returns the EnableYSQL field if non-nil, zero value otherwise.

### GetEnableYSQLOk

`func (o *UserIntent) GetEnableYSQLOk() (*bool, bool)`

GetEnableYSQLOk returns a tuple with the EnableYSQL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableYSQL

`func (o *UserIntent) SetEnableYSQL(v bool)`

SetEnableYSQL sets EnableYSQL field to given value.

### HasEnableYSQL

`func (o *UserIntent) HasEnableYSQL() bool`

HasEnableYSQL returns a boolean if a field has been set.

### GetEnableYSQLAuth

`func (o *UserIntent) GetEnableYSQLAuth() bool`

GetEnableYSQLAuth returns the EnableYSQLAuth field if non-nil, zero value otherwise.

### GetEnableYSQLAuthOk

`func (o *UserIntent) GetEnableYSQLAuthOk() (*bool, bool)`

GetEnableYSQLAuthOk returns a tuple with the EnableYSQLAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableYSQLAuth

`func (o *UserIntent) SetEnableYSQLAuth(v bool)`

SetEnableYSQLAuth sets EnableYSQLAuth field to given value.

### HasEnableYSQLAuth

`func (o *UserIntent) HasEnableYSQLAuth() bool`

HasEnableYSQLAuth returns a boolean if a field has been set.

### GetImageBundleUUID

`func (o *UserIntent) GetImageBundleUUID() string`

GetImageBundleUUID returns the ImageBundleUUID field if non-nil, zero value otherwise.

### GetImageBundleUUIDOk

`func (o *UserIntent) GetImageBundleUUIDOk() (*string, bool)`

GetImageBundleUUIDOk returns a tuple with the ImageBundleUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBundleUUID

`func (o *UserIntent) SetImageBundleUUID(v string)`

SetImageBundleUUID sets ImageBundleUUID field to given value.

### HasImageBundleUUID

`func (o *UserIntent) HasImageBundleUUID() bool`

HasImageBundleUUID returns a boolean if a field has been set.

### GetInstanceTags

`func (o *UserIntent) GetInstanceTags() map[string]string`

GetInstanceTags returns the InstanceTags field if non-nil, zero value otherwise.

### GetInstanceTagsOk

`func (o *UserIntent) GetInstanceTagsOk() (*map[string]string, bool)`

GetInstanceTagsOk returns a tuple with the InstanceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTags

`func (o *UserIntent) SetInstanceTags(v map[string]string)`

SetInstanceTags sets InstanceTags field to given value.

### HasInstanceTags

`func (o *UserIntent) HasInstanceTags() bool`

HasInstanceTags returns a boolean if a field has been set.

### GetInstanceType

`func (o *UserIntent) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *UserIntent) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *UserIntent) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *UserIntent) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetKubernetesOperatorVersion

`func (o *UserIntent) GetKubernetesOperatorVersion() int64`

GetKubernetesOperatorVersion returns the KubernetesOperatorVersion field if non-nil, zero value otherwise.

### GetKubernetesOperatorVersionOk

`func (o *UserIntent) GetKubernetesOperatorVersionOk() (*int64, bool)`

GetKubernetesOperatorVersionOk returns a tuple with the KubernetesOperatorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesOperatorVersion

`func (o *UserIntent) SetKubernetesOperatorVersion(v int64)`

SetKubernetesOperatorVersion sets KubernetesOperatorVersion field to given value.

### HasKubernetesOperatorVersion

`func (o *UserIntent) HasKubernetesOperatorVersion() bool`

HasKubernetesOperatorVersion returns a boolean if a field has been set.

### GetMasterDeviceInfo

`func (o *UserIntent) GetMasterDeviceInfo() DeviceInfo`

GetMasterDeviceInfo returns the MasterDeviceInfo field if non-nil, zero value otherwise.

### GetMasterDeviceInfoOk

`func (o *UserIntent) GetMasterDeviceInfoOk() (*DeviceInfo, bool)`

GetMasterDeviceInfoOk returns a tuple with the MasterDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterDeviceInfo

`func (o *UserIntent) SetMasterDeviceInfo(v DeviceInfo)`

SetMasterDeviceInfo sets MasterDeviceInfo field to given value.

### HasMasterDeviceInfo

`func (o *UserIntent) HasMasterDeviceInfo() bool`

HasMasterDeviceInfo returns a boolean if a field has been set.

### GetMasterGFlags

`func (o *UserIntent) GetMasterGFlags() map[string]string`

GetMasterGFlags returns the MasterGFlags field if non-nil, zero value otherwise.

### GetMasterGFlagsOk

`func (o *UserIntent) GetMasterGFlagsOk() (*map[string]string, bool)`

GetMasterGFlagsOk returns a tuple with the MasterGFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterGFlags

`func (o *UserIntent) SetMasterGFlags(v map[string]string)`

SetMasterGFlags sets MasterGFlags field to given value.

### HasMasterGFlags

`func (o *UserIntent) HasMasterGFlags() bool`

HasMasterGFlags returns a boolean if a field has been set.

### GetMasterInstanceType

`func (o *UserIntent) GetMasterInstanceType() string`

GetMasterInstanceType returns the MasterInstanceType field if non-nil, zero value otherwise.

### GetMasterInstanceTypeOk

`func (o *UserIntent) GetMasterInstanceTypeOk() (*string, bool)`

GetMasterInstanceTypeOk returns a tuple with the MasterInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterInstanceType

`func (o *UserIntent) SetMasterInstanceType(v string)`

SetMasterInstanceType sets MasterInstanceType field to given value.

### HasMasterInstanceType

`func (o *UserIntent) HasMasterInstanceType() bool`

HasMasterInstanceType returns a boolean if a field has been set.

### GetNumNodes

`func (o *UserIntent) GetNumNodes() int32`

GetNumNodes returns the NumNodes field if non-nil, zero value otherwise.

### GetNumNodesOk

`func (o *UserIntent) GetNumNodesOk() (*int32, bool)`

GetNumNodesOk returns a tuple with the NumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNodes

`func (o *UserIntent) SetNumNodes(v int32)`

SetNumNodes sets NumNodes field to given value.

### HasNumNodes

`func (o *UserIntent) HasNumNodes() bool`

HasNumNodes returns a boolean if a field has been set.

### GetPreferredRegion

`func (o *UserIntent) GetPreferredRegion() string`

GetPreferredRegion returns the PreferredRegion field if non-nil, zero value otherwise.

### GetPreferredRegionOk

`func (o *UserIntent) GetPreferredRegionOk() (*string, bool)`

GetPreferredRegionOk returns a tuple with the PreferredRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredRegion

`func (o *UserIntent) SetPreferredRegion(v string)`

SetPreferredRegion sets PreferredRegion field to given value.

### HasPreferredRegion

`func (o *UserIntent) HasPreferredRegion() bool`

HasPreferredRegion returns a boolean if a field has been set.

### GetProvider

`func (o *UserIntent) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UserIntent) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UserIntent) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UserIntent) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetProviderType

`func (o *UserIntent) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *UserIntent) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *UserIntent) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *UserIntent) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetRegionList

`func (o *UserIntent) GetRegionList() []string`

GetRegionList returns the RegionList field if non-nil, zero value otherwise.

### GetRegionListOk

`func (o *UserIntent) GetRegionListOk() (*[]string, bool)`

GetRegionListOk returns a tuple with the RegionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionList

`func (o *UserIntent) SetRegionList(v []string)`

SetRegionList sets RegionList field to given value.

### HasRegionList

`func (o *UserIntent) HasRegionList() bool`

HasRegionList returns a boolean if a field has been set.

### GetReplicationFactor

`func (o *UserIntent) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *UserIntent) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *UserIntent) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *UserIntent) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.

### GetSpecificGFlags

`func (o *UserIntent) GetSpecificGFlags() SpecificGFlags`

GetSpecificGFlags returns the SpecificGFlags field if non-nil, zero value otherwise.

### GetSpecificGFlagsOk

`func (o *UserIntent) GetSpecificGFlagsOk() (*SpecificGFlags, bool)`

GetSpecificGFlagsOk returns a tuple with the SpecificGFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificGFlags

`func (o *UserIntent) SetSpecificGFlags(v SpecificGFlags)`

SetSpecificGFlags sets SpecificGFlags field to given value.

### HasSpecificGFlags

`func (o *UserIntent) HasSpecificGFlags() bool`

HasSpecificGFlags returns a boolean if a field has been set.

### GetSpotPrice

`func (o *UserIntent) GetSpotPrice() float64`

GetSpotPrice returns the SpotPrice field if non-nil, zero value otherwise.

### GetSpotPriceOk

`func (o *UserIntent) GetSpotPriceOk() (*float64, bool)`

GetSpotPriceOk returns a tuple with the SpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotPrice

`func (o *UserIntent) SetSpotPrice(v float64)`

SetSpotPrice sets SpotPrice field to given value.

### HasSpotPrice

`func (o *UserIntent) HasSpotPrice() bool`

HasSpotPrice returns a boolean if a field has been set.

### GetTserverGFlags

`func (o *UserIntent) GetTserverGFlags() map[string]string`

GetTserverGFlags returns the TserverGFlags field if non-nil, zero value otherwise.

### GetTserverGFlagsOk

`func (o *UserIntent) GetTserverGFlagsOk() (*map[string]string, bool)`

GetTserverGFlagsOk returns a tuple with the TserverGFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverGFlags

`func (o *UserIntent) SetTserverGFlags(v map[string]string)`

SetTserverGFlags sets TserverGFlags field to given value.

### HasTserverGFlags

`func (o *UserIntent) HasTserverGFlags() bool`

HasTserverGFlags returns a boolean if a field has been set.

### GetUniverseName

`func (o *UserIntent) GetUniverseName() string`

GetUniverseName returns the UniverseName field if non-nil, zero value otherwise.

### GetUniverseNameOk

`func (o *UserIntent) GetUniverseNameOk() (*string, bool)`

GetUniverseNameOk returns a tuple with the UniverseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseName

`func (o *UserIntent) SetUniverseName(v string)`

SetUniverseName sets UniverseName field to given value.

### HasUniverseName

`func (o *UserIntent) HasUniverseName() bool`

HasUniverseName returns a boolean if a field has been set.

### GetUniverseOverrides

`func (o *UserIntent) GetUniverseOverrides() string`

GetUniverseOverrides returns the UniverseOverrides field if non-nil, zero value otherwise.

### GetUniverseOverridesOk

`func (o *UserIntent) GetUniverseOverridesOk() (*string, bool)`

GetUniverseOverridesOk returns a tuple with the UniverseOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseOverrides

`func (o *UserIntent) SetUniverseOverrides(v string)`

SetUniverseOverrides sets UniverseOverrides field to given value.

### HasUniverseOverrides

`func (o *UserIntent) HasUniverseOverrides() bool`

HasUniverseOverrides returns a boolean if a field has been set.

### GetUseHostname

`func (o *UserIntent) GetUseHostname() bool`

GetUseHostname returns the UseHostname field if non-nil, zero value otherwise.

### GetUseHostnameOk

`func (o *UserIntent) GetUseHostnameOk() (*bool, bool)`

GetUseHostnameOk returns a tuple with the UseHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHostname

`func (o *UserIntent) SetUseHostname(v bool)`

SetUseHostname sets UseHostname field to given value.

### HasUseHostname

`func (o *UserIntent) HasUseHostname() bool`

HasUseHostname returns a boolean if a field has been set.

### GetUseSpotInstance

`func (o *UserIntent) GetUseSpotInstance() bool`

GetUseSpotInstance returns the UseSpotInstance field if non-nil, zero value otherwise.

### GetUseSpotInstanceOk

`func (o *UserIntent) GetUseSpotInstanceOk() (*bool, bool)`

GetUseSpotInstanceOk returns a tuple with the UseSpotInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSpotInstance

`func (o *UserIntent) SetUseSpotInstance(v bool)`

SetUseSpotInstance sets UseSpotInstance field to given value.

### HasUseSpotInstance

`func (o *UserIntent) HasUseSpotInstance() bool`

HasUseSpotInstance returns a boolean if a field has been set.

### GetUseSystemd

`func (o *UserIntent) GetUseSystemd() bool`

GetUseSystemd returns the UseSystemd field if non-nil, zero value otherwise.

### GetUseSystemdOk

`func (o *UserIntent) GetUseSystemdOk() (*bool, bool)`

GetUseSystemdOk returns a tuple with the UseSystemd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSystemd

`func (o *UserIntent) SetUseSystemd(v bool)`

SetUseSystemd sets UseSystemd field to given value.

### HasUseSystemd

`func (o *UserIntent) HasUseSystemd() bool`

HasUseSystemd returns a boolean if a field has been set.

### GetUseTimeSync

`func (o *UserIntent) GetUseTimeSync() bool`

GetUseTimeSync returns the UseTimeSync field if non-nil, zero value otherwise.

### GetUseTimeSyncOk

`func (o *UserIntent) GetUseTimeSyncOk() (*bool, bool)`

GetUseTimeSyncOk returns a tuple with the UseTimeSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTimeSync

`func (o *UserIntent) SetUseTimeSync(v bool)`

SetUseTimeSync sets UseTimeSync field to given value.

### HasUseTimeSync

`func (o *UserIntent) HasUseTimeSync() bool`

HasUseTimeSync returns a boolean if a field has been set.

### GetYbSoftwareVersion

`func (o *UserIntent) GetYbSoftwareVersion() string`

GetYbSoftwareVersion returns the YbSoftwareVersion field if non-nil, zero value otherwise.

### GetYbSoftwareVersionOk

`func (o *UserIntent) GetYbSoftwareVersionOk() (*string, bool)`

GetYbSoftwareVersionOk returns a tuple with the YbSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbSoftwareVersion

`func (o *UserIntent) SetYbSoftwareVersion(v string)`

SetYbSoftwareVersion sets YbSoftwareVersion field to given value.

### HasYbSoftwareVersion

`func (o *UserIntent) HasYbSoftwareVersion() bool`

HasYbSoftwareVersion returns a boolean if a field has been set.

### GetYbcFlags

`func (o *UserIntent) GetYbcFlags() map[string]string`

GetYbcFlags returns the YbcFlags field if non-nil, zero value otherwise.

### GetYbcFlagsOk

`func (o *UserIntent) GetYbcFlagsOk() (*map[string]string, bool)`

GetYbcFlagsOk returns a tuple with the YbcFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcFlags

`func (o *UserIntent) SetYbcFlags(v map[string]string)`

SetYbcFlags sets YbcFlags field to given value.

### HasYbcFlags

`func (o *UserIntent) HasYbcFlags() bool`

HasYbcFlags returns a boolean if a field has been set.

### GetYcqlPassword

`func (o *UserIntent) GetYcqlPassword() string`

GetYcqlPassword returns the YcqlPassword field if non-nil, zero value otherwise.

### GetYcqlPasswordOk

`func (o *UserIntent) GetYcqlPasswordOk() (*string, bool)`

GetYcqlPasswordOk returns a tuple with the YcqlPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYcqlPassword

`func (o *UserIntent) SetYcqlPassword(v string)`

SetYcqlPassword sets YcqlPassword field to given value.

### HasYcqlPassword

`func (o *UserIntent) HasYcqlPassword() bool`

HasYcqlPassword returns a boolean if a field has been set.

### GetYsqlPassword

`func (o *UserIntent) GetYsqlPassword() string`

GetYsqlPassword returns the YsqlPassword field if non-nil, zero value otherwise.

### GetYsqlPasswordOk

`func (o *UserIntent) GetYsqlPasswordOk() (*string, bool)`

GetYsqlPasswordOk returns a tuple with the YsqlPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlPassword

`func (o *UserIntent) SetYsqlPassword(v string)`

SetYsqlPassword sets YsqlPassword field to given value.

### HasYsqlPassword

`func (o *UserIntent) HasYsqlPassword() bool`

HasYsqlPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


