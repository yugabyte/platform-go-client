# CollectFilesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionOptions** | [**FileCollectionOptions**](FileCollectionOptions.md) |  | 
**Nodes** | Pointer to [**NodeSelection**](NodeSelection.md) |  | [optional] 

## Methods

### NewCollectFilesRequest

`func NewCollectFilesRequest(collectionOptions FileCollectionOptions, ) *CollectFilesRequest`

NewCollectFilesRequest instantiates a new CollectFilesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectFilesRequestWithDefaults

`func NewCollectFilesRequestWithDefaults() *CollectFilesRequest`

NewCollectFilesRequestWithDefaults instantiates a new CollectFilesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionOptions

`func (o *CollectFilesRequest) GetCollectionOptions() FileCollectionOptions`

GetCollectionOptions returns the CollectionOptions field if non-nil, zero value otherwise.

### GetCollectionOptionsOk

`func (o *CollectFilesRequest) GetCollectionOptionsOk() (*FileCollectionOptions, bool)`

GetCollectionOptionsOk returns a tuple with the CollectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionOptions

`func (o *CollectFilesRequest) SetCollectionOptions(v FileCollectionOptions)`

SetCollectionOptions sets CollectionOptions field to given value.


### GetNodes

`func (o *CollectFilesRequest) GetNodes() NodeSelection`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *CollectFilesRequest) GetNodesOk() (*NodeSelection, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *CollectFilesRequest) SetNodes(v NodeSelection)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *CollectFilesRequest) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


