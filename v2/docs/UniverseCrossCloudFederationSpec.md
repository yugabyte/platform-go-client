# UniverseCrossCloudFederationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | When true, configure cross-cloud federated IAM on all nodes of the universe (the provider must have federated IAM enabled with an audience set). When false, tear it down on all nodes. | 

## Methods

### NewUniverseCrossCloudFederationSpec

`func NewUniverseCrossCloudFederationSpec(enabled bool, ) *UniverseCrossCloudFederationSpec`

NewUniverseCrossCloudFederationSpec instantiates a new UniverseCrossCloudFederationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseCrossCloudFederationSpecWithDefaults

`func NewUniverseCrossCloudFederationSpecWithDefaults() *UniverseCrossCloudFederationSpec`

NewUniverseCrossCloudFederationSpecWithDefaults instantiates a new UniverseCrossCloudFederationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UniverseCrossCloudFederationSpec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UniverseCrossCloudFederationSpec) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UniverseCrossCloudFederationSpec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


