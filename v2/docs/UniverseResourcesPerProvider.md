# UniverseResourcesPerProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcesPerProvider** | [**map[string]UniverseResourceDetails**](UniverseResourceDetails.md) | Resource details keyed by provider UUID. | 

## Methods

### NewUniverseResourcesPerProvider

`func NewUniverseResourcesPerProvider(resourcesPerProvider map[string]UniverseResourceDetails, ) *UniverseResourcesPerProvider`

NewUniverseResourcesPerProvider instantiates a new UniverseResourcesPerProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseResourcesPerProviderWithDefaults

`func NewUniverseResourcesPerProviderWithDefaults() *UniverseResourcesPerProvider`

NewUniverseResourcesPerProviderWithDefaults instantiates a new UniverseResourcesPerProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcesPerProvider

`func (o *UniverseResourcesPerProvider) GetResourcesPerProvider() map[string]UniverseResourceDetails`

GetResourcesPerProvider returns the ResourcesPerProvider field if non-nil, zero value otherwise.

### GetResourcesPerProviderOk

`func (o *UniverseResourcesPerProvider) GetResourcesPerProviderOk() (*map[string]UniverseResourceDetails, bool)`

GetResourcesPerProviderOk returns a tuple with the ResourcesPerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesPerProvider

`func (o *UniverseResourcesPerProvider) SetResourcesPerProvider(v map[string]UniverseResourceDetails)`

SetResourcesPerProvider sets ResourcesPerProvider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


