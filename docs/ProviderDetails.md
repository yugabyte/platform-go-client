# ProviderDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AirGapInstall** | Pointer to **bool** |  | [optional] 
**CloudInfo** | Pointer to [**CloudInfo**](CloudInfo.md) |  | [optional] 
**EnableNodeAgent** | Pointer to **bool** |  | [optional] [readonly] 
**InstallNodeExporter** | Pointer to **bool** |  | [optional] 
**NodeExporterPort** | Pointer to **int32** |  | [optional] 
**NodeExporterUser** | Pointer to **string** |  | [optional] 
**NtpServers** | Pointer to **[]string** |  | [optional] 
**PasswordlessSudoAccess** | Pointer to **bool** |  | [optional] 
**ProvisionInstanceScript** | Pointer to **string** |  | [optional] 
**SetUpChrony** | Pointer to **bool** |  | [optional] 
**ShowSetUpChrony** | Pointer to **bool** |  | [optional] 
**SkipProvisioning** | Pointer to **bool** |  | [optional] 
**SshPort** | Pointer to **int32** |  | [optional] 
**SshUser** | Pointer to **string** |  | [optional] 

## Methods

### NewProviderDetails

`func NewProviderDetails() *ProviderDetails`

NewProviderDetails instantiates a new ProviderDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderDetailsWithDefaults

`func NewProviderDetailsWithDefaults() *ProviderDetails`

NewProviderDetailsWithDefaults instantiates a new ProviderDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAirGapInstall

`func (o *ProviderDetails) GetAirGapInstall() bool`

GetAirGapInstall returns the AirGapInstall field if non-nil, zero value otherwise.

### GetAirGapInstallOk

`func (o *ProviderDetails) GetAirGapInstallOk() (*bool, bool)`

GetAirGapInstallOk returns a tuple with the AirGapInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirGapInstall

`func (o *ProviderDetails) SetAirGapInstall(v bool)`

SetAirGapInstall sets AirGapInstall field to given value.

### HasAirGapInstall

`func (o *ProviderDetails) HasAirGapInstall() bool`

HasAirGapInstall returns a boolean if a field has been set.

### GetCloudInfo

`func (o *ProviderDetails) GetCloudInfo() CloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *ProviderDetails) GetCloudInfoOk() (*CloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *ProviderDetails) SetCloudInfo(v CloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *ProviderDetails) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetEnableNodeAgent

`func (o *ProviderDetails) GetEnableNodeAgent() bool`

GetEnableNodeAgent returns the EnableNodeAgent field if non-nil, zero value otherwise.

### GetEnableNodeAgentOk

`func (o *ProviderDetails) GetEnableNodeAgentOk() (*bool, bool)`

GetEnableNodeAgentOk returns a tuple with the EnableNodeAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNodeAgent

`func (o *ProviderDetails) SetEnableNodeAgent(v bool)`

SetEnableNodeAgent sets EnableNodeAgent field to given value.

### HasEnableNodeAgent

`func (o *ProviderDetails) HasEnableNodeAgent() bool`

HasEnableNodeAgent returns a boolean if a field has been set.

### GetInstallNodeExporter

`func (o *ProviderDetails) GetInstallNodeExporter() bool`

GetInstallNodeExporter returns the InstallNodeExporter field if non-nil, zero value otherwise.

### GetInstallNodeExporterOk

`func (o *ProviderDetails) GetInstallNodeExporterOk() (*bool, bool)`

GetInstallNodeExporterOk returns a tuple with the InstallNodeExporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallNodeExporter

`func (o *ProviderDetails) SetInstallNodeExporter(v bool)`

SetInstallNodeExporter sets InstallNodeExporter field to given value.

### HasInstallNodeExporter

`func (o *ProviderDetails) HasInstallNodeExporter() bool`

HasInstallNodeExporter returns a boolean if a field has been set.

### GetNodeExporterPort

`func (o *ProviderDetails) GetNodeExporterPort() int32`

GetNodeExporterPort returns the NodeExporterPort field if non-nil, zero value otherwise.

### GetNodeExporterPortOk

`func (o *ProviderDetails) GetNodeExporterPortOk() (*int32, bool)`

GetNodeExporterPortOk returns a tuple with the NodeExporterPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExporterPort

`func (o *ProviderDetails) SetNodeExporterPort(v int32)`

SetNodeExporterPort sets NodeExporterPort field to given value.

### HasNodeExporterPort

`func (o *ProviderDetails) HasNodeExporterPort() bool`

HasNodeExporterPort returns a boolean if a field has been set.

### GetNodeExporterUser

`func (o *ProviderDetails) GetNodeExporterUser() string`

GetNodeExporterUser returns the NodeExporterUser field if non-nil, zero value otherwise.

### GetNodeExporterUserOk

`func (o *ProviderDetails) GetNodeExporterUserOk() (*string, bool)`

GetNodeExporterUserOk returns a tuple with the NodeExporterUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeExporterUser

`func (o *ProviderDetails) SetNodeExporterUser(v string)`

SetNodeExporterUser sets NodeExporterUser field to given value.

### HasNodeExporterUser

`func (o *ProviderDetails) HasNodeExporterUser() bool`

HasNodeExporterUser returns a boolean if a field has been set.

### GetNtpServers

`func (o *ProviderDetails) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *ProviderDetails) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *ProviderDetails) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *ProviderDetails) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetPasswordlessSudoAccess

`func (o *ProviderDetails) GetPasswordlessSudoAccess() bool`

GetPasswordlessSudoAccess returns the PasswordlessSudoAccess field if non-nil, zero value otherwise.

### GetPasswordlessSudoAccessOk

`func (o *ProviderDetails) GetPasswordlessSudoAccessOk() (*bool, bool)`

GetPasswordlessSudoAccessOk returns a tuple with the PasswordlessSudoAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordlessSudoAccess

`func (o *ProviderDetails) SetPasswordlessSudoAccess(v bool)`

SetPasswordlessSudoAccess sets PasswordlessSudoAccess field to given value.

### HasPasswordlessSudoAccess

`func (o *ProviderDetails) HasPasswordlessSudoAccess() bool`

HasPasswordlessSudoAccess returns a boolean if a field has been set.

### GetProvisionInstanceScript

`func (o *ProviderDetails) GetProvisionInstanceScript() string`

GetProvisionInstanceScript returns the ProvisionInstanceScript field if non-nil, zero value otherwise.

### GetProvisionInstanceScriptOk

`func (o *ProviderDetails) GetProvisionInstanceScriptOk() (*string, bool)`

GetProvisionInstanceScriptOk returns a tuple with the ProvisionInstanceScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionInstanceScript

`func (o *ProviderDetails) SetProvisionInstanceScript(v string)`

SetProvisionInstanceScript sets ProvisionInstanceScript field to given value.

### HasProvisionInstanceScript

`func (o *ProviderDetails) HasProvisionInstanceScript() bool`

HasProvisionInstanceScript returns a boolean if a field has been set.

### GetSetUpChrony

`func (o *ProviderDetails) GetSetUpChrony() bool`

GetSetUpChrony returns the SetUpChrony field if non-nil, zero value otherwise.

### GetSetUpChronyOk

`func (o *ProviderDetails) GetSetUpChronyOk() (*bool, bool)`

GetSetUpChronyOk returns a tuple with the SetUpChrony field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetUpChrony

`func (o *ProviderDetails) SetSetUpChrony(v bool)`

SetSetUpChrony sets SetUpChrony field to given value.

### HasSetUpChrony

`func (o *ProviderDetails) HasSetUpChrony() bool`

HasSetUpChrony returns a boolean if a field has been set.

### GetShowSetUpChrony

`func (o *ProviderDetails) GetShowSetUpChrony() bool`

GetShowSetUpChrony returns the ShowSetUpChrony field if non-nil, zero value otherwise.

### GetShowSetUpChronyOk

`func (o *ProviderDetails) GetShowSetUpChronyOk() (*bool, bool)`

GetShowSetUpChronyOk returns a tuple with the ShowSetUpChrony field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSetUpChrony

`func (o *ProviderDetails) SetShowSetUpChrony(v bool)`

SetShowSetUpChrony sets ShowSetUpChrony field to given value.

### HasShowSetUpChrony

`func (o *ProviderDetails) HasShowSetUpChrony() bool`

HasShowSetUpChrony returns a boolean if a field has been set.

### GetSkipProvisioning

`func (o *ProviderDetails) GetSkipProvisioning() bool`

GetSkipProvisioning returns the SkipProvisioning field if non-nil, zero value otherwise.

### GetSkipProvisioningOk

`func (o *ProviderDetails) GetSkipProvisioningOk() (*bool, bool)`

GetSkipProvisioningOk returns a tuple with the SkipProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipProvisioning

`func (o *ProviderDetails) SetSkipProvisioning(v bool)`

SetSkipProvisioning sets SkipProvisioning field to given value.

### HasSkipProvisioning

`func (o *ProviderDetails) HasSkipProvisioning() bool`

HasSkipProvisioning returns a boolean if a field has been set.

### GetSshPort

`func (o *ProviderDetails) GetSshPort() int32`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *ProviderDetails) GetSshPortOk() (*int32, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *ProviderDetails) SetSshPort(v int32)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *ProviderDetails) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetSshUser

`func (o *ProviderDetails) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *ProviderDetails) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *ProviderDetails) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.

### HasSshUser

`func (o *ProviderDetails) HasSshUser() bool`

HasSshUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


