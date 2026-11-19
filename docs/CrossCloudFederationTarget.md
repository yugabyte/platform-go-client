# CrossCloudFederationTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | **string** | Token audience. For a gcp target this is the GCP Workload Identity Federation audience (//iam.googleapis.com/projects/.../providers/...); for an aws target it is the audience requested for the node&#39;s web-identity token. | 
**RoleArn** | Pointer to **string** | AWS role assumed via AssumeRoleWithWebIdentity. Required when targetCloud is aws, unused otherwise: reaching GCS binds the node&#39;s identity to the bucket directly, with no intermediate role. | [optional] 
**TargetCloud** | **string** | Storage cloud this entry grants access to. Only aws (S3) and gcp (GCS) are supported today. | 

## Methods

### NewCrossCloudFederationTarget

`func NewCrossCloudFederationTarget(audience string, targetCloud string, ) *CrossCloudFederationTarget`

NewCrossCloudFederationTarget instantiates a new CrossCloudFederationTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossCloudFederationTargetWithDefaults

`func NewCrossCloudFederationTargetWithDefaults() *CrossCloudFederationTarget`

NewCrossCloudFederationTargetWithDefaults instantiates a new CrossCloudFederationTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *CrossCloudFederationTarget) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CrossCloudFederationTarget) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CrossCloudFederationTarget) SetAudience(v string)`

SetAudience sets Audience field to given value.


### GetRoleArn

`func (o *CrossCloudFederationTarget) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *CrossCloudFederationTarget) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *CrossCloudFederationTarget) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *CrossCloudFederationTarget) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetTargetCloud

`func (o *CrossCloudFederationTarget) GetTargetCloud() string`

GetTargetCloud returns the TargetCloud field if non-nil, zero value otherwise.

### GetTargetCloudOk

`func (o *CrossCloudFederationTarget) GetTargetCloudOk() (*string, bool)`

GetTargetCloudOk returns a tuple with the TargetCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCloud

`func (o *CrossCloudFederationTarget) SetTargetCloud(v string)`

SetTargetCloud sets TargetCloud field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


