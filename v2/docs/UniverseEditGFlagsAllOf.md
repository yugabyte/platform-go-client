# UniverseEditGFlagsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KubernetesResourceDetails** | Pointer to [**KubernetesResourceDetails**](KubernetesResourceDetails.md) |  | [optional] 
**UniverseGflags** | Pointer to [**map[string]ClusterGFlags**](ClusterGFlags.md) | GFlags for each cluster uuid of this universe | [optional] 

## Methods

### NewUniverseEditGFlagsAllOf

`func NewUniverseEditGFlagsAllOf() *UniverseEditGFlagsAllOf`

NewUniverseEditGFlagsAllOf instantiates a new UniverseEditGFlagsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseEditGFlagsAllOfWithDefaults

`func NewUniverseEditGFlagsAllOfWithDefaults() *UniverseEditGFlagsAllOf`

NewUniverseEditGFlagsAllOfWithDefaults instantiates a new UniverseEditGFlagsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubernetesResourceDetails

`func (o *UniverseEditGFlagsAllOf) GetKubernetesResourceDetails() KubernetesResourceDetails`

GetKubernetesResourceDetails returns the KubernetesResourceDetails field if non-nil, zero value otherwise.

### GetKubernetesResourceDetailsOk

`func (o *UniverseEditGFlagsAllOf) GetKubernetesResourceDetailsOk() (*KubernetesResourceDetails, bool)`

GetKubernetesResourceDetailsOk returns a tuple with the KubernetesResourceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesResourceDetails

`func (o *UniverseEditGFlagsAllOf) SetKubernetesResourceDetails(v KubernetesResourceDetails)`

SetKubernetesResourceDetails sets KubernetesResourceDetails field to given value.

### HasKubernetesResourceDetails

`func (o *UniverseEditGFlagsAllOf) HasKubernetesResourceDetails() bool`

HasKubernetesResourceDetails returns a boolean if a field has been set.

### GetUniverseGflags

`func (o *UniverseEditGFlagsAllOf) GetUniverseGflags() map[string]ClusterGFlags`

GetUniverseGflags returns the UniverseGflags field if non-nil, zero value otherwise.

### GetUniverseGflagsOk

`func (o *UniverseEditGFlagsAllOf) GetUniverseGflagsOk() (*map[string]ClusterGFlags, bool)`

GetUniverseGflagsOk returns a tuple with the UniverseGflags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseGflags

`func (o *UniverseEditGFlagsAllOf) SetUniverseGflags(v map[string]ClusterGFlags)`

SetUniverseGflags sets UniverseGflags field to given value.

### HasUniverseGflags

`func (o *UniverseEditGFlagsAllOf) HasUniverseGflags() bool`

HasUniverseGflags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


