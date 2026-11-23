# GCPCloudInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CrossCloudFederationTargets** | Pointer to [**[]CrossCloudFederationTarget**](CrossCloudFederationTarget.md) | WARNING: This is a preview API that could change. Storage clouds this provider&#39;s DB nodes can be given federated access to, at most one entry per cloud. A node uses the entry for a cloud it is not itself running on. | [optional] 
**DestVpcId** | Pointer to **string** |  | [optional] 
**EnableFederatedIam** | Pointer to **bool** | Enable S3-on-GCP cross-cloud federated IAM on this provider&#39;s DB nodes (AWS AssumeRoleWithWebIdentity). Requires the federated IAM role ARN and audience below. | [optional] 
**FederatedIamAudience** | Pointer to **string** | Deprecated: use crossCloudFederationTargets. Audience requested from GCP for the web-identity token, used when federated IAM is enabled. | [optional] 
**FederatedIamRoleArn** | Pointer to **string** | AWS role ARN assumed via AssumeRoleWithWebIdentity when federated IAM is enabled (GCP DB node -&gt; S3). | [optional] 
**GceApplicationCredentials** | Pointer to **string** |  | [optional] 
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

### GetCrossCloudFederationTargets

`func (o *GCPCloudInfo) GetCrossCloudFederationTargets() []CrossCloudFederationTarget`

GetCrossCloudFederationTargets returns the CrossCloudFederationTargets field if non-nil, zero value otherwise.

### GetCrossCloudFederationTargetsOk

`func (o *GCPCloudInfo) GetCrossCloudFederationTargetsOk() (*[]CrossCloudFederationTarget, bool)`

GetCrossCloudFederationTargetsOk returns a tuple with the CrossCloudFederationTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossCloudFederationTargets

`func (o *GCPCloudInfo) SetCrossCloudFederationTargets(v []CrossCloudFederationTarget)`

SetCrossCloudFederationTargets sets CrossCloudFederationTargets field to given value.

### HasCrossCloudFederationTargets

`func (o *GCPCloudInfo) HasCrossCloudFederationTargets() bool`

HasCrossCloudFederationTargets returns a boolean if a field has been set.

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

### GetEnableFederatedIam

`func (o *GCPCloudInfo) GetEnableFederatedIam() bool`

GetEnableFederatedIam returns the EnableFederatedIam field if non-nil, zero value otherwise.

### GetEnableFederatedIamOk

`func (o *GCPCloudInfo) GetEnableFederatedIamOk() (*bool, bool)`

GetEnableFederatedIamOk returns a tuple with the EnableFederatedIam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFederatedIam

`func (o *GCPCloudInfo) SetEnableFederatedIam(v bool)`

SetEnableFederatedIam sets EnableFederatedIam field to given value.

### HasEnableFederatedIam

`func (o *GCPCloudInfo) HasEnableFederatedIam() bool`

HasEnableFederatedIam returns a boolean if a field has been set.

### GetFederatedIamAudience

`func (o *GCPCloudInfo) GetFederatedIamAudience() string`

GetFederatedIamAudience returns the FederatedIamAudience field if non-nil, zero value otherwise.

### GetFederatedIamAudienceOk

`func (o *GCPCloudInfo) GetFederatedIamAudienceOk() (*string, bool)`

GetFederatedIamAudienceOk returns a tuple with the FederatedIamAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedIamAudience

`func (o *GCPCloudInfo) SetFederatedIamAudience(v string)`

SetFederatedIamAudience sets FederatedIamAudience field to given value.

### HasFederatedIamAudience

`func (o *GCPCloudInfo) HasFederatedIamAudience() bool`

HasFederatedIamAudience returns a boolean if a field has been set.

### GetFederatedIamRoleArn

`func (o *GCPCloudInfo) GetFederatedIamRoleArn() string`

GetFederatedIamRoleArn returns the FederatedIamRoleArn field if non-nil, zero value otherwise.

### GetFederatedIamRoleArnOk

`func (o *GCPCloudInfo) GetFederatedIamRoleArnOk() (*string, bool)`

GetFederatedIamRoleArnOk returns a tuple with the FederatedIamRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedIamRoleArn

`func (o *GCPCloudInfo) SetFederatedIamRoleArn(v string)`

SetFederatedIamRoleArn sets FederatedIamRoleArn field to given value.

### HasFederatedIamRoleArn

`func (o *GCPCloudInfo) HasFederatedIamRoleArn() bool`

HasFederatedIamRoleArn returns a boolean if a field has been set.

### GetGceApplicationCredentials

`func (o *GCPCloudInfo) GetGceApplicationCredentials() string`

GetGceApplicationCredentials returns the GceApplicationCredentials field if non-nil, zero value otherwise.

### GetGceApplicationCredentialsOk

`func (o *GCPCloudInfo) GetGceApplicationCredentialsOk() (*string, bool)`

GetGceApplicationCredentialsOk returns a tuple with the GceApplicationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGceApplicationCredentials

`func (o *GCPCloudInfo) SetGceApplicationCredentials(v string)`

SetGceApplicationCredentials sets GceApplicationCredentials field to given value.

### HasGceApplicationCredentials

`func (o *GCPCloudInfo) HasGceApplicationCredentials() bool`

HasGceApplicationCredentials returns a boolean if a field has been set.

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


