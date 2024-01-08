# GCPCloudInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestVpcId** | Pointer to **string** |  | [optional] 
**GceApplicationCredentialsPath** | Pointer to **string** |  | [optional] [readonly] 
**GceProject** | Pointer to **string** |  | [optional] 
**HostVpcId** | Pointer to **string** |  | [optional] [readonly] 
**SharedVPCProject** | Pointer to **string** |  | [optional] 
**UseHostCredentials** | Pointer to **bool** |  | [optional] 
**UseHostVPC** | Pointer to **bool** |  | [optional] 
**VpcType** | Pointer to **string** | New/Existing VPC for provider creation | [optional] [readonly] 
**YbFirewallTags** | Pointer to **string** |  | [optional] 

## Methods

### NewGCPCloudInfo

`func NewGCPCloudInfo() *GCPCloudInfo`

NewGCPCloudInfo instantiates a new GCPCloudInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPCloudInfoWithDefaults

`func NewGCPCloudInfoWithDefaults() *GCPCloudInfo`

NewGCPCloudInfoWithDefaults instantiates a new GCPCloudInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestVpcId

`func (o *GCPCloudInfo) GetDestVpcId() string`

GetDestVpcId returns the DestVpcId field if non-nil, zero value otherwise.

### GetDestVpcIdOk

`func (o *GCPCloudInfo) GetDestVpcIdOk() (*string, bool)`

GetDestVpcIdOk returns a tuple with the DestVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestVpcId

`func (o *GCPCloudInfo) SetDestVpcId(v string)`

SetDestVpcId sets DestVpcId field to given value.

### HasDestVpcId

`func (o *GCPCloudInfo) HasDestVpcId() bool`

HasDestVpcId returns a boolean if a field has been set.

### GetGceApplicationCredentialsPath

`func (o *GCPCloudInfo) GetGceApplicationCredentialsPath() string`

GetGceApplicationCredentialsPath returns the GceApplicationCredentialsPath field if non-nil, zero value otherwise.

### GetGceApplicationCredentialsPathOk

`func (o *GCPCloudInfo) GetGceApplicationCredentialsPathOk() (*string, bool)`

GetGceApplicationCredentialsPathOk returns a tuple with the GceApplicationCredentialsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGceApplicationCredentialsPath

`func (o *GCPCloudInfo) SetGceApplicationCredentialsPath(v string)`

SetGceApplicationCredentialsPath sets GceApplicationCredentialsPath field to given value.

### HasGceApplicationCredentialsPath

`func (o *GCPCloudInfo) HasGceApplicationCredentialsPath() bool`

HasGceApplicationCredentialsPath returns a boolean if a field has been set.

### GetGceProject

`func (o *GCPCloudInfo) GetGceProject() string`

GetGceProject returns the GceProject field if non-nil, zero value otherwise.

### GetGceProjectOk

`func (o *GCPCloudInfo) GetGceProjectOk() (*string, bool)`

GetGceProjectOk returns a tuple with the GceProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGceProject

`func (o *GCPCloudInfo) SetGceProject(v string)`

SetGceProject sets GceProject field to given value.

### HasGceProject

`func (o *GCPCloudInfo) HasGceProject() bool`

HasGceProject returns a boolean if a field has been set.

### GetHostVpcId

`func (o *GCPCloudInfo) GetHostVpcId() string`

GetHostVpcId returns the HostVpcId field if non-nil, zero value otherwise.

### GetHostVpcIdOk

`func (o *GCPCloudInfo) GetHostVpcIdOk() (*string, bool)`

GetHostVpcIdOk returns a tuple with the HostVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostVpcId

`func (o *GCPCloudInfo) SetHostVpcId(v string)`

SetHostVpcId sets HostVpcId field to given value.

### HasHostVpcId

`func (o *GCPCloudInfo) HasHostVpcId() bool`

HasHostVpcId returns a boolean if a field has been set.

### GetSharedVPCProject

`func (o *GCPCloudInfo) GetSharedVPCProject() string`

GetSharedVPCProject returns the SharedVPCProject field if non-nil, zero value otherwise.

### GetSharedVPCProjectOk

`func (o *GCPCloudInfo) GetSharedVPCProjectOk() (*string, bool)`

GetSharedVPCProjectOk returns a tuple with the SharedVPCProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedVPCProject

`func (o *GCPCloudInfo) SetSharedVPCProject(v string)`

SetSharedVPCProject sets SharedVPCProject field to given value.

### HasSharedVPCProject

`func (o *GCPCloudInfo) HasSharedVPCProject() bool`

HasSharedVPCProject returns a boolean if a field has been set.

### GetUseHostCredentials

`func (o *GCPCloudInfo) GetUseHostCredentials() bool`

GetUseHostCredentials returns the UseHostCredentials field if non-nil, zero value otherwise.

### GetUseHostCredentialsOk

`func (o *GCPCloudInfo) GetUseHostCredentialsOk() (*bool, bool)`

GetUseHostCredentialsOk returns a tuple with the UseHostCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHostCredentials

`func (o *GCPCloudInfo) SetUseHostCredentials(v bool)`

SetUseHostCredentials sets UseHostCredentials field to given value.

### HasUseHostCredentials

`func (o *GCPCloudInfo) HasUseHostCredentials() bool`

HasUseHostCredentials returns a boolean if a field has been set.

### GetUseHostVPC

`func (o *GCPCloudInfo) GetUseHostVPC() bool`

GetUseHostVPC returns the UseHostVPC field if non-nil, zero value otherwise.

### GetUseHostVPCOk

`func (o *GCPCloudInfo) GetUseHostVPCOk() (*bool, bool)`

GetUseHostVPCOk returns a tuple with the UseHostVPC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHostVPC

`func (o *GCPCloudInfo) SetUseHostVPC(v bool)`

SetUseHostVPC sets UseHostVPC field to given value.

### HasUseHostVPC

`func (o *GCPCloudInfo) HasUseHostVPC() bool`

HasUseHostVPC returns a boolean if a field has been set.

### GetVpcType

`func (o *GCPCloudInfo) GetVpcType() string`

GetVpcType returns the VpcType field if non-nil, zero value otherwise.

### GetVpcTypeOk

`func (o *GCPCloudInfo) GetVpcTypeOk() (*string, bool)`

GetVpcTypeOk returns a tuple with the VpcType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcType

`func (o *GCPCloudInfo) SetVpcType(v string)`

SetVpcType sets VpcType field to given value.

### HasVpcType

`func (o *GCPCloudInfo) HasVpcType() bool`

HasVpcType returns a boolean if a field has been set.

### GetYbFirewallTags

`func (o *GCPCloudInfo) GetYbFirewallTags() string`

GetYbFirewallTags returns the YbFirewallTags field if non-nil, zero value otherwise.

### GetYbFirewallTagsOk

`func (o *GCPCloudInfo) GetYbFirewallTagsOk() (*string, bool)`

GetYbFirewallTagsOk returns a tuple with the YbFirewallTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbFirewallTags

`func (o *GCPCloudInfo) SetYbFirewallTags(v string)`

SetYbFirewallTags sets YbFirewallTags field to given value.

### HasYbFirewallTags

`func (o *GCPCloudInfo) HasYbFirewallTags() bool`

HasYbFirewallTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


