# OnPremCloudInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CrossCloudFederationTargets** | Pointer to [**[]CrossCloudFederationTarget**](CrossCloudFederationTarget.md) | WARNING: This is a preview API that could change. Storage clouds this provider&#39;s DB nodes can reach with federated IAM, at most one entry per cloud. A node uses the entry for a cloud it is not itself running on, and is left alone when there is none - so add an entry per storage cloud the nodes must actually reach. | [optional] 
**EnableFederatedIam** | Pointer to **bool** | WARNING: This is a preview API that could change. Enable cross-cloud federated IAM on this provider&#39;s DB nodes. Each node is given access to the storage clouds it is not running on, per crossCloudFederationTargets. | [optional] 
**EnableMultiTenancy** | Pointer to **bool** | WARNING: This is a preview API that could change. | [optional] 
**FederatedIamAudience** | Pointer to **string** | WARNING: This is a preview API that could change. Deprecated: use crossCloudFederationTargets. GCP Workload Identity Federation audience used when federated IAM is enabled. | [optional] 
**UseClockbound** | Pointer to **bool** | WARNING: This is a preview API that could change. | [optional] 
**YbHomeDir** | Pointer to **string** |  | [optional] 
**YnpManaged** | Pointer to **bool** | YbaApi Internal. Provider is created and managed by YNP | [optional] [readonly] 

## Methods

### NewOnPremCloudInfo

`func NewOnPremCloudInfo() *OnPremCloudInfo`

NewOnPremCloudInfo instantiates a new OnPremCloudInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnPremCloudInfoWithDefaults

`func NewOnPremCloudInfoWithDefaults() *OnPremCloudInfo`

NewOnPremCloudInfoWithDefaults instantiates a new OnPremCloudInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrossCloudFederationTargets

`func (o *OnPremCloudInfo) GetCrossCloudFederationTargets() []CrossCloudFederationTarget`

GetCrossCloudFederationTargets returns the CrossCloudFederationTargets field if non-nil, zero value otherwise.

### GetCrossCloudFederationTargetsOk

`func (o *OnPremCloudInfo) GetCrossCloudFederationTargetsOk() (*[]CrossCloudFederationTarget, bool)`

GetCrossCloudFederationTargetsOk returns a tuple with the CrossCloudFederationTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossCloudFederationTargets

`func (o *OnPremCloudInfo) SetCrossCloudFederationTargets(v []CrossCloudFederationTarget)`

SetCrossCloudFederationTargets sets CrossCloudFederationTargets field to given value.

### HasCrossCloudFederationTargets

`func (o *OnPremCloudInfo) HasCrossCloudFederationTargets() bool`

HasCrossCloudFederationTargets returns a boolean if a field has been set.

### GetEnableFederatedIam

`func (o *OnPremCloudInfo) GetEnableFederatedIam() bool`

GetEnableFederatedIam returns the EnableFederatedIam field if non-nil, zero value otherwise.

### GetEnableFederatedIamOk

`func (o *OnPremCloudInfo) GetEnableFederatedIamOk() (*bool, bool)`

GetEnableFederatedIamOk returns a tuple with the EnableFederatedIam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFederatedIam

`func (o *OnPremCloudInfo) SetEnableFederatedIam(v bool)`

SetEnableFederatedIam sets EnableFederatedIam field to given value.

### HasEnableFederatedIam

`func (o *OnPremCloudInfo) HasEnableFederatedIam() bool`

HasEnableFederatedIam returns a boolean if a field has been set.

### GetEnableMultiTenancy

`func (o *OnPremCloudInfo) GetEnableMultiTenancy() bool`

GetEnableMultiTenancy returns the EnableMultiTenancy field if non-nil, zero value otherwise.

### GetEnableMultiTenancyOk

`func (o *OnPremCloudInfo) GetEnableMultiTenancyOk() (*bool, bool)`

GetEnableMultiTenancyOk returns a tuple with the EnableMultiTenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMultiTenancy

`func (o *OnPremCloudInfo) SetEnableMultiTenancy(v bool)`

SetEnableMultiTenancy sets EnableMultiTenancy field to given value.

### HasEnableMultiTenancy

`func (o *OnPremCloudInfo) HasEnableMultiTenancy() bool`

HasEnableMultiTenancy returns a boolean if a field has been set.

### GetFederatedIamAudience

`func (o *OnPremCloudInfo) GetFederatedIamAudience() string`

GetFederatedIamAudience returns the FederatedIamAudience field if non-nil, zero value otherwise.

### GetFederatedIamAudienceOk

`func (o *OnPremCloudInfo) GetFederatedIamAudienceOk() (*string, bool)`

GetFederatedIamAudienceOk returns a tuple with the FederatedIamAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedIamAudience

`func (o *OnPremCloudInfo) SetFederatedIamAudience(v string)`

SetFederatedIamAudience sets FederatedIamAudience field to given value.

### HasFederatedIamAudience

`func (o *OnPremCloudInfo) HasFederatedIamAudience() bool`

HasFederatedIamAudience returns a boolean if a field has been set.

### GetUseClockbound

`func (o *OnPremCloudInfo) GetUseClockbound() bool`

GetUseClockbound returns the UseClockbound field if non-nil, zero value otherwise.

### GetUseClockboundOk

`func (o *OnPremCloudInfo) GetUseClockboundOk() (*bool, bool)`

GetUseClockboundOk returns a tuple with the UseClockbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseClockbound

`func (o *OnPremCloudInfo) SetUseClockbound(v bool)`

SetUseClockbound sets UseClockbound field to given value.

### HasUseClockbound

`func (o *OnPremCloudInfo) HasUseClockbound() bool`

HasUseClockbound returns a boolean if a field has been set.

### GetYbHomeDir

`func (o *OnPremCloudInfo) GetYbHomeDir() string`

GetYbHomeDir returns the YbHomeDir field if non-nil, zero value otherwise.

### GetYbHomeDirOk

`func (o *OnPremCloudInfo) GetYbHomeDirOk() (*string, bool)`

GetYbHomeDirOk returns a tuple with the YbHomeDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbHomeDir

`func (o *OnPremCloudInfo) SetYbHomeDir(v string)`

SetYbHomeDir sets YbHomeDir field to given value.

### HasYbHomeDir

`func (o *OnPremCloudInfo) HasYbHomeDir() bool`

HasYbHomeDir returns a boolean if a field has been set.

### GetYnpManaged

`func (o *OnPremCloudInfo) GetYnpManaged() bool`

GetYnpManaged returns the YnpManaged field if non-nil, zero value otherwise.

### GetYnpManagedOk

`func (o *OnPremCloudInfo) GetYnpManagedOk() (*bool, bool)`

GetYnpManagedOk returns a tuple with the YnpManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYnpManaged

`func (o *OnPremCloudInfo) SetYnpManaged(v bool)`

SetYnpManaged sets YnpManaged field to given value.

### HasYnpManaged

`func (o *OnPremCloudInfo) HasYnpManaged() bool`

HasYnpManaged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


