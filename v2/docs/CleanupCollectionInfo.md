# CleanupCollectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionUuid** | **string** | The collection UUID that was cleaned up | 
**NodesCleaned** | **int32** | Number of nodes where files were successfully deleted | 
**Message** | **string** | Status message | 

## Methods

### NewCleanupCollectionInfo

`func NewCleanupCollectionInfo(collectionUuid string, nodesCleaned int32, message string, ) *CleanupCollectionInfo`

NewCleanupCollectionInfo instantiates a new CleanupCollectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCleanupCollectionInfoWithDefaults

`func NewCleanupCollectionInfoWithDefaults() *CleanupCollectionInfo`

NewCleanupCollectionInfoWithDefaults instantiates a new CleanupCollectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionUuid

`func (o *CleanupCollectionInfo) GetCollectionUuid() string`

GetCollectionUuid returns the CollectionUuid field if non-nil, zero value otherwise.

### GetCollectionUuidOk

`func (o *CleanupCollectionInfo) GetCollectionUuidOk() (*string, bool)`

GetCollectionUuidOk returns a tuple with the CollectionUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionUuid

`func (o *CleanupCollectionInfo) SetCollectionUuid(v string)`

SetCollectionUuid sets CollectionUuid field to given value.


### GetNodesCleaned

`func (o *CleanupCollectionInfo) GetNodesCleaned() int32`

GetNodesCleaned returns the NodesCleaned field if non-nil, zero value otherwise.

### GetNodesCleanedOk

`func (o *CleanupCollectionInfo) GetNodesCleanedOk() (*int32, bool)`

GetNodesCleanedOk returns a tuple with the NodesCleaned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesCleaned

`func (o *CleanupCollectionInfo) SetNodesCleaned(v int32)`

SetNodesCleaned sets NodesCleaned field to given value.


### GetMessage

`func (o *CleanupCollectionInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CleanupCollectionInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CleanupCollectionInfo) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


