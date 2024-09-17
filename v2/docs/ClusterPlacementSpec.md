# ClusterPlacementSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudList** | [**[]PlacementCloud**](PlacementCloud.md) |  | 

## Methods

### NewClusterPlacementSpec

`func NewClusterPlacementSpec(cloudList []PlacementCloud, ) *ClusterPlacementSpec`

NewClusterPlacementSpec instantiates a new ClusterPlacementSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPlacementSpecWithDefaults

`func NewClusterPlacementSpecWithDefaults() *ClusterPlacementSpec`

NewClusterPlacementSpecWithDefaults instantiates a new ClusterPlacementSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudList

`func (o *ClusterPlacementSpec) GetCloudList() []PlacementCloud`

GetCloudList returns the CloudList field if non-nil, zero value otherwise.

### GetCloudListOk

`func (o *ClusterPlacementSpec) GetCloudListOk() (*[]PlacementCloud, bool)`

GetCloudListOk returns a tuple with the CloudList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudList

`func (o *ClusterPlacementSpec) SetCloudList(v []PlacementCloud)`

SetCloudList sets CloudList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


