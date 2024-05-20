# ClusterNetworkingSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableExposingService** | Pointer to **string** | Whether to create a load balancer service for this cluster. Defaults to NONE. | [optional] [default to "NONE"]
**EnableLb** | Pointer to **bool** | Create target groups if enabled. Used by YBM. | [optional] 

## Methods

### NewClusterNetworkingSpec

`func NewClusterNetworkingSpec() *ClusterNetworkingSpec`

NewClusterNetworkingSpec instantiates a new ClusterNetworkingSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNetworkingSpecWithDefaults

`func NewClusterNetworkingSpecWithDefaults() *ClusterNetworkingSpec`

NewClusterNetworkingSpecWithDefaults instantiates a new ClusterNetworkingSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableExposingService

`func (o *ClusterNetworkingSpec) GetEnableExposingService() string`

GetEnableExposingService returns the EnableExposingService field if non-nil, zero value otherwise.

### GetEnableExposingServiceOk

`func (o *ClusterNetworkingSpec) GetEnableExposingServiceOk() (*string, bool)`

GetEnableExposingServiceOk returns a tuple with the EnableExposingService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExposingService

`func (o *ClusterNetworkingSpec) SetEnableExposingService(v string)`

SetEnableExposingService sets EnableExposingService field to given value.

### HasEnableExposingService

`func (o *ClusterNetworkingSpec) HasEnableExposingService() bool`

HasEnableExposingService returns a boolean if a field has been set.

### GetEnableLb

`func (o *ClusterNetworkingSpec) GetEnableLb() bool`

GetEnableLb returns the EnableLb field if non-nil, zero value otherwise.

### GetEnableLbOk

`func (o *ClusterNetworkingSpec) GetEnableLbOk() (*bool, bool)`

GetEnableLbOk returns a tuple with the EnableLb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLb

`func (o *ClusterNetworkingSpec) SetEnableLb(v bool)`

SetEnableLb sets EnableLb field to given value.

### HasEnableLb

`func (o *ClusterNetworkingSpec) HasEnableLb() bool`

HasEnableLb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


