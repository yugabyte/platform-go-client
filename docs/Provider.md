# Provider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Provider active status | [optional] [readonly] 
**AirGapInstall** | Pointer to **bool** | Transient property - only present in mutate API request | [optional] 
**Code** | Pointer to **string** | Provider cloud code | [optional] 
**Config** | Pointer to **map[string]string** |  | [optional] 
**CustomHostCidrs** | Pointer to **[]string** | Transient property - only present in mutate API request | [optional] 
**CustomerUUID** | Pointer to **string** | Customer uuid | [optional] [readonly] 
**DestVpcId** | Pointer to **string** | Transient property - only present in mutate API request | [optional] 
**HostVpcId** | Pointer to **string** | Transient property - only present in mutate API request | [optional] 
**HostVpcRegion** | Pointer to **string** | Transient property - only present in mutate API request | [optional] 
**HostedZoneId** | Pointer to **string** |  | [optional] 
**HostedZoneName** | Pointer to **string** |  | [optional] 
**KeyPairName** | Pointer to **string** | Transient property - only present in mutate API request | [optional] 
**Name** | Pointer to **string** | Provider name | [optional] 
**NtpServers** | Pointer to **[]string** | Transient property - only present in mutate API request | [optional] 
**OverrideKeyValidate** | Pointer to **bool** | Transient property - only present in mutate API request | [optional] 
**Regions** | [**[]Region**](Region.md) |  | 
**SetUpChrony** | Pointer to **bool** | Transient property - only present in mutate API request | [optional] 
**ShowSetUpChrony** | Pointer to **bool** | Transient property - only present in mutate API request | [optional] 
**SshPort** | Pointer to **int32** | Transient property - only present in mutate API request | [optional] 
**SshPrivateKeyContent** | Pointer to **string** | Transient property - only present in mutate API request | [optional] 
**SshUser** | Pointer to **string** | Transient property - only present in mutate API request | [optional] 
**Uuid** | Pointer to **string** | Provider uuid | [optional] [readonly] 

## Methods

### NewProvider

`func NewProvider(regions []Region, ) *Provider`

NewProvider instantiates a new Provider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderWithDefaults

`func NewProviderWithDefaults() *Provider`

NewProviderWithDefaults instantiates a new Provider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *Provider) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Provider) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Provider) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Provider) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAirGapInstall

`func (o *Provider) GetAirGapInstall() bool`

GetAirGapInstall returns the AirGapInstall field if non-nil, zero value otherwise.

### GetAirGapInstallOk

`func (o *Provider) GetAirGapInstallOk() (*bool, bool)`

GetAirGapInstallOk returns a tuple with the AirGapInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirGapInstall

`func (o *Provider) SetAirGapInstall(v bool)`

SetAirGapInstall sets AirGapInstall field to given value.

### HasAirGapInstall

`func (o *Provider) HasAirGapInstall() bool`

HasAirGapInstall returns a boolean if a field has been set.

### GetCode

`func (o *Provider) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Provider) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Provider) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Provider) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetConfig

`func (o *Provider) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Provider) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Provider) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Provider) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCustomHostCidrs

`func (o *Provider) GetCustomHostCidrs() []string`

GetCustomHostCidrs returns the CustomHostCidrs field if non-nil, zero value otherwise.

### GetCustomHostCidrsOk

`func (o *Provider) GetCustomHostCidrsOk() (*[]string, bool)`

GetCustomHostCidrsOk returns a tuple with the CustomHostCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHostCidrs

`func (o *Provider) SetCustomHostCidrs(v []string)`

SetCustomHostCidrs sets CustomHostCidrs field to given value.

### HasCustomHostCidrs

`func (o *Provider) HasCustomHostCidrs() bool`

HasCustomHostCidrs returns a boolean if a field has been set.

### GetCustomerUUID

`func (o *Provider) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *Provider) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *Provider) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.

### HasCustomerUUID

`func (o *Provider) HasCustomerUUID() bool`

HasCustomerUUID returns a boolean if a field has been set.

### GetDestVpcId

`func (o *Provider) GetDestVpcId() string`

GetDestVpcId returns the DestVpcId field if non-nil, zero value otherwise.

### GetDestVpcIdOk

`func (o *Provider) GetDestVpcIdOk() (*string, bool)`

GetDestVpcIdOk returns a tuple with the DestVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestVpcId

`func (o *Provider) SetDestVpcId(v string)`

SetDestVpcId sets DestVpcId field to given value.

### HasDestVpcId

`func (o *Provider) HasDestVpcId() bool`

HasDestVpcId returns a boolean if a field has been set.

### GetHostVpcId

`func (o *Provider) GetHostVpcId() string`

GetHostVpcId returns the HostVpcId field if non-nil, zero value otherwise.

### GetHostVpcIdOk

`func (o *Provider) GetHostVpcIdOk() (*string, bool)`

GetHostVpcIdOk returns a tuple with the HostVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostVpcId

`func (o *Provider) SetHostVpcId(v string)`

SetHostVpcId sets HostVpcId field to given value.

### HasHostVpcId

`func (o *Provider) HasHostVpcId() bool`

HasHostVpcId returns a boolean if a field has been set.

### GetHostVpcRegion

`func (o *Provider) GetHostVpcRegion() string`

GetHostVpcRegion returns the HostVpcRegion field if non-nil, zero value otherwise.

### GetHostVpcRegionOk

`func (o *Provider) GetHostVpcRegionOk() (*string, bool)`

GetHostVpcRegionOk returns a tuple with the HostVpcRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostVpcRegion

`func (o *Provider) SetHostVpcRegion(v string)`

SetHostVpcRegion sets HostVpcRegion field to given value.

### HasHostVpcRegion

`func (o *Provider) HasHostVpcRegion() bool`

HasHostVpcRegion returns a boolean if a field has been set.

### GetHostedZoneId

`func (o *Provider) GetHostedZoneId() string`

GetHostedZoneId returns the HostedZoneId field if non-nil, zero value otherwise.

### GetHostedZoneIdOk

`func (o *Provider) GetHostedZoneIdOk() (*string, bool)`

GetHostedZoneIdOk returns a tuple with the HostedZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZoneId

`func (o *Provider) SetHostedZoneId(v string)`

SetHostedZoneId sets HostedZoneId field to given value.

### HasHostedZoneId

`func (o *Provider) HasHostedZoneId() bool`

HasHostedZoneId returns a boolean if a field has been set.

### GetHostedZoneName

`func (o *Provider) GetHostedZoneName() string`

GetHostedZoneName returns the HostedZoneName field if non-nil, zero value otherwise.

### GetHostedZoneNameOk

`func (o *Provider) GetHostedZoneNameOk() (*string, bool)`

GetHostedZoneNameOk returns a tuple with the HostedZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZoneName

`func (o *Provider) SetHostedZoneName(v string)`

SetHostedZoneName sets HostedZoneName field to given value.

### HasHostedZoneName

`func (o *Provider) HasHostedZoneName() bool`

HasHostedZoneName returns a boolean if a field has been set.

### GetKeyPairName

`func (o *Provider) GetKeyPairName() string`

GetKeyPairName returns the KeyPairName field if non-nil, zero value otherwise.

### GetKeyPairNameOk

`func (o *Provider) GetKeyPairNameOk() (*string, bool)`

GetKeyPairNameOk returns a tuple with the KeyPairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPairName

`func (o *Provider) SetKeyPairName(v string)`

SetKeyPairName sets KeyPairName field to given value.

### HasKeyPairName

`func (o *Provider) HasKeyPairName() bool`

HasKeyPairName returns a boolean if a field has been set.

### GetName

`func (o *Provider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Provider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Provider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Provider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNtpServers

`func (o *Provider) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *Provider) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *Provider) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *Provider) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetOverrideKeyValidate

`func (o *Provider) GetOverrideKeyValidate() bool`

GetOverrideKeyValidate returns the OverrideKeyValidate field if non-nil, zero value otherwise.

### GetOverrideKeyValidateOk

`func (o *Provider) GetOverrideKeyValidateOk() (*bool, bool)`

GetOverrideKeyValidateOk returns a tuple with the OverrideKeyValidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideKeyValidate

`func (o *Provider) SetOverrideKeyValidate(v bool)`

SetOverrideKeyValidate sets OverrideKeyValidate field to given value.

### HasOverrideKeyValidate

`func (o *Provider) HasOverrideKeyValidate() bool`

HasOverrideKeyValidate returns a boolean if a field has been set.

### GetRegions

`func (o *Provider) GetRegions() []Region`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Provider) GetRegionsOk() (*[]Region, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Provider) SetRegions(v []Region)`

SetRegions sets Regions field to given value.


### GetSetUpChrony

`func (o *Provider) GetSetUpChrony() bool`

GetSetUpChrony returns the SetUpChrony field if non-nil, zero value otherwise.

### GetSetUpChronyOk

`func (o *Provider) GetSetUpChronyOk() (*bool, bool)`

GetSetUpChronyOk returns a tuple with the SetUpChrony field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetUpChrony

`func (o *Provider) SetSetUpChrony(v bool)`

SetSetUpChrony sets SetUpChrony field to given value.

### HasSetUpChrony

`func (o *Provider) HasSetUpChrony() bool`

HasSetUpChrony returns a boolean if a field has been set.

### GetShowSetUpChrony

`func (o *Provider) GetShowSetUpChrony() bool`

GetShowSetUpChrony returns the ShowSetUpChrony field if non-nil, zero value otherwise.

### GetShowSetUpChronyOk

`func (o *Provider) GetShowSetUpChronyOk() (*bool, bool)`

GetShowSetUpChronyOk returns a tuple with the ShowSetUpChrony field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSetUpChrony

`func (o *Provider) SetShowSetUpChrony(v bool)`

SetShowSetUpChrony sets ShowSetUpChrony field to given value.

### HasShowSetUpChrony

`func (o *Provider) HasShowSetUpChrony() bool`

HasShowSetUpChrony returns a boolean if a field has been set.

### GetSshPort

`func (o *Provider) GetSshPort() int32`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *Provider) GetSshPortOk() (*int32, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *Provider) SetSshPort(v int32)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *Provider) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetSshPrivateKeyContent

`func (o *Provider) GetSshPrivateKeyContent() string`

GetSshPrivateKeyContent returns the SshPrivateKeyContent field if non-nil, zero value otherwise.

### GetSshPrivateKeyContentOk

`func (o *Provider) GetSshPrivateKeyContentOk() (*string, bool)`

GetSshPrivateKeyContentOk returns a tuple with the SshPrivateKeyContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeyContent

`func (o *Provider) SetSshPrivateKeyContent(v string)`

SetSshPrivateKeyContent sets SshPrivateKeyContent field to given value.

### HasSshPrivateKeyContent

`func (o *Provider) HasSshPrivateKeyContent() bool`

HasSshPrivateKeyContent returns a boolean if a field has been set.

### GetSshUser

`func (o *Provider) GetSshUser() string`

GetSshUser returns the SshUser field if non-nil, zero value otherwise.

### GetSshUserOk

`func (o *Provider) GetSshUserOk() (*string, bool)`

GetSshUserOk returns a tuple with the SshUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUser

`func (o *Provider) SetSshUser(v string)`

SetSshUser sets SshUser field to given value.

### HasSshUser

`func (o *Provider) HasSshUser() bool`

HasSshUser returns a boolean if a field has been set.

### GetUuid

`func (o *Provider) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Provider) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Provider) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Provider) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


