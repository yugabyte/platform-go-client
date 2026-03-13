# FileCollectionSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionUuid** | Pointer to **string** | Unique identifier for this collection. Use this UUID with the download-collection API to download the collected files to YBA. Collections expire after 1 hour.  | [optional] 
**TotalNodes** | **int32** | Total number of nodes targeted | 
**SuccessfulNodes** | **int32** | Number of nodes where collection succeeded | 
**FailedNodes** | **int32** | Number of nodes where collection failed | 
**TotalFilesCollected** | **int32** | Total number of files collected across all nodes | 
**TotalFilesSkipped** | Pointer to **int32** | Total number of files skipped across all nodes | [optional] 
**TotalFilesFailed** | Pointer to **int32** | Total number of files that failed to collect across all nodes | [optional] 
**TotalBytesCollected** | Pointer to **int64** | Total bytes collected across all nodes | [optional] 
**TotalExecutionTimeMs** | **int64** | Total time taken for file collection in milliseconds | 
**AllSucceeded** | **bool** | Whether all node file collections succeeded | 

## Methods

### NewFileCollectionSummary

`func NewFileCollectionSummary(totalNodes int32, successfulNodes int32, failedNodes int32, totalFilesCollected int32, totalExecutionTimeMs int64, allSucceeded bool, ) *FileCollectionSummary`

NewFileCollectionSummary instantiates a new FileCollectionSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileCollectionSummaryWithDefaults

`func NewFileCollectionSummaryWithDefaults() *FileCollectionSummary`

NewFileCollectionSummaryWithDefaults instantiates a new FileCollectionSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionUuid

`func (o *FileCollectionSummary) GetCollectionUuid() string`

GetCollectionUuid returns the CollectionUuid field if non-nil, zero value otherwise.

### GetCollectionUuidOk

`func (o *FileCollectionSummary) GetCollectionUuidOk() (*string, bool)`

GetCollectionUuidOk returns a tuple with the CollectionUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionUuid

`func (o *FileCollectionSummary) SetCollectionUuid(v string)`

SetCollectionUuid sets CollectionUuid field to given value.

### HasCollectionUuid

`func (o *FileCollectionSummary) HasCollectionUuid() bool`

HasCollectionUuid returns a boolean if a field has been set.

### GetTotalNodes

`func (o *FileCollectionSummary) GetTotalNodes() int32`

GetTotalNodes returns the TotalNodes field if non-nil, zero value otherwise.

### GetTotalNodesOk

`func (o *FileCollectionSummary) GetTotalNodesOk() (*int32, bool)`

GetTotalNodesOk returns a tuple with the TotalNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNodes

`func (o *FileCollectionSummary) SetTotalNodes(v int32)`

SetTotalNodes sets TotalNodes field to given value.


### GetSuccessfulNodes

`func (o *FileCollectionSummary) GetSuccessfulNodes() int32`

GetSuccessfulNodes returns the SuccessfulNodes field if non-nil, zero value otherwise.

### GetSuccessfulNodesOk

`func (o *FileCollectionSummary) GetSuccessfulNodesOk() (*int32, bool)`

GetSuccessfulNodesOk returns a tuple with the SuccessfulNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulNodes

`func (o *FileCollectionSummary) SetSuccessfulNodes(v int32)`

SetSuccessfulNodes sets SuccessfulNodes field to given value.


### GetFailedNodes

`func (o *FileCollectionSummary) GetFailedNodes() int32`

GetFailedNodes returns the FailedNodes field if non-nil, zero value otherwise.

### GetFailedNodesOk

`func (o *FileCollectionSummary) GetFailedNodesOk() (*int32, bool)`

GetFailedNodesOk returns a tuple with the FailedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedNodes

`func (o *FileCollectionSummary) SetFailedNodes(v int32)`

SetFailedNodes sets FailedNodes field to given value.


### GetTotalFilesCollected

`func (o *FileCollectionSummary) GetTotalFilesCollected() int32`

GetTotalFilesCollected returns the TotalFilesCollected field if non-nil, zero value otherwise.

### GetTotalFilesCollectedOk

`func (o *FileCollectionSummary) GetTotalFilesCollectedOk() (*int32, bool)`

GetTotalFilesCollectedOk returns a tuple with the TotalFilesCollected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFilesCollected

`func (o *FileCollectionSummary) SetTotalFilesCollected(v int32)`

SetTotalFilesCollected sets TotalFilesCollected field to given value.


### GetTotalFilesSkipped

`func (o *FileCollectionSummary) GetTotalFilesSkipped() int32`

GetTotalFilesSkipped returns the TotalFilesSkipped field if non-nil, zero value otherwise.

### GetTotalFilesSkippedOk

`func (o *FileCollectionSummary) GetTotalFilesSkippedOk() (*int32, bool)`

GetTotalFilesSkippedOk returns a tuple with the TotalFilesSkipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFilesSkipped

`func (o *FileCollectionSummary) SetTotalFilesSkipped(v int32)`

SetTotalFilesSkipped sets TotalFilesSkipped field to given value.

### HasTotalFilesSkipped

`func (o *FileCollectionSummary) HasTotalFilesSkipped() bool`

HasTotalFilesSkipped returns a boolean if a field has been set.

### GetTotalFilesFailed

`func (o *FileCollectionSummary) GetTotalFilesFailed() int32`

GetTotalFilesFailed returns the TotalFilesFailed field if non-nil, zero value otherwise.

### GetTotalFilesFailedOk

`func (o *FileCollectionSummary) GetTotalFilesFailedOk() (*int32, bool)`

GetTotalFilesFailedOk returns a tuple with the TotalFilesFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFilesFailed

`func (o *FileCollectionSummary) SetTotalFilesFailed(v int32)`

SetTotalFilesFailed sets TotalFilesFailed field to given value.

### HasTotalFilesFailed

`func (o *FileCollectionSummary) HasTotalFilesFailed() bool`

HasTotalFilesFailed returns a boolean if a field has been set.

### GetTotalBytesCollected

`func (o *FileCollectionSummary) GetTotalBytesCollected() int64`

GetTotalBytesCollected returns the TotalBytesCollected field if non-nil, zero value otherwise.

### GetTotalBytesCollectedOk

`func (o *FileCollectionSummary) GetTotalBytesCollectedOk() (*int64, bool)`

GetTotalBytesCollectedOk returns a tuple with the TotalBytesCollected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesCollected

`func (o *FileCollectionSummary) SetTotalBytesCollected(v int64)`

SetTotalBytesCollected sets TotalBytesCollected field to given value.

### HasTotalBytesCollected

`func (o *FileCollectionSummary) HasTotalBytesCollected() bool`

HasTotalBytesCollected returns a boolean if a field has been set.

### GetTotalExecutionTimeMs

`func (o *FileCollectionSummary) GetTotalExecutionTimeMs() int64`

GetTotalExecutionTimeMs returns the TotalExecutionTimeMs field if non-nil, zero value otherwise.

### GetTotalExecutionTimeMsOk

`func (o *FileCollectionSummary) GetTotalExecutionTimeMsOk() (*int64, bool)`

GetTotalExecutionTimeMsOk returns a tuple with the TotalExecutionTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExecutionTimeMs

`func (o *FileCollectionSummary) SetTotalExecutionTimeMs(v int64)`

SetTotalExecutionTimeMs sets TotalExecutionTimeMs field to given value.


### GetAllSucceeded

`func (o *FileCollectionSummary) GetAllSucceeded() bool`

GetAllSucceeded returns the AllSucceeded field if non-nil, zero value otherwise.

### GetAllSucceededOk

`func (o *FileCollectionSummary) GetAllSucceededOk() (*bool, bool)`

GetAllSucceededOk returns a tuple with the AllSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSucceeded

`func (o *FileCollectionSummary) SetAllSucceeded(v bool)`

SetAllSucceeded sets AllSucceeded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


