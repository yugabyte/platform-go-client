# Provider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Provider active status | [optional] [readonly] 
**Code** | Pointer to **string** | Provider cloud code | [optional] 
**Config** | Pointer to **map[string]string** |  | [optional] 
**CustomerUUID** | Pointer to **string** | Customer uuid | [optional] [readonly] 
**DestVpcId** | Pointer to **string** |  | [optional] 
**Details** | [**ProviderDetails**](ProviderDetails.md) |  | 
**HostVpcId** | Pointer to **string** |  | [optional] 
**HostVpcRegion** | Pointer to **string** |  | [optional] 
**HostedZoneId** | Pointer to **string** |  | [optional] 
**HostedZoneName** | Pointer to **string** |  | [optional] 
**KeyPairName** | Pointer to **string** | Transient property - only present in mutate API request | [optional] 
**Name** | Pointer to **string** | Provider name | [optional] 
**Regions** | [**[]Region**](Region.md) |  | 
**SshPrivateKeyContent** | Pointer to **string** | Transient property - only present in mutate API request | [optional] 
**Uuid** | Pointer to **string** | Provider uuid | [optional] [readonly] 
**Version** | Pointer to **int64** | Provider version | [optional] [readonly] 

## Methods

### NewProvider

`func NewProvider(details ProviderDetails, regions []Region, ) *Provider`

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

### GetDetails

`func (o *Provider) GetDetails() ProviderDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Provider) GetDetailsOk() (*ProviderDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Provider) SetDetails(v ProviderDetails)`

SetDetails sets Details field to given value.


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

### GetVersion

`func (o *Provider) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Provider) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Provider) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Provider) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


