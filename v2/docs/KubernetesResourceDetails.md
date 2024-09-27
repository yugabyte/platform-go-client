# KubernetesResourceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to **string** | Type of YBA resource to operate upon | [optional] 
**Name** | Pointer to **string** | Name of this Kubernetes resource | [optional] 
**Namespace** | Pointer to **string** | Kubernetes namespace in which the resource is found | [optional] 

## Methods

### NewKubernetesResourceDetails

`func NewKubernetesResourceDetails() *KubernetesResourceDetails`

NewKubernetesResourceDetails instantiates a new KubernetesResourceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesResourceDetailsWithDefaults

`func NewKubernetesResourceDetailsWithDefaults() *KubernetesResourceDetails`

NewKubernetesResourceDetailsWithDefaults instantiates a new KubernetesResourceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *KubernetesResourceDetails) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *KubernetesResourceDetails) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *KubernetesResourceDetails) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *KubernetesResourceDetails) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetName

`func (o *KubernetesResourceDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesResourceDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesResourceDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesResourceDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *KubernetesResourceDetails) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *KubernetesResourceDetails) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *KubernetesResourceDetails) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *KubernetesResourceDetails) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


