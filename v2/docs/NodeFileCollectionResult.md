# NodeFileCollectionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeName** | **string** | Name of the node | 
**NodeAddress** | **string** | IP address of the node | 
**Success** | **bool** | Whether file collection succeeded on this node | 
**FilesCollected** | **int32** | Number of files successfully collected | 
**FilesSkipped** | **int32** | Number of files skipped (due to size limits or other reasons) | 
**FilesFailed** | **int32** | Number of files that failed to collect | 
**TotalBytesCollected** | Pointer to **int64** | Total bytes of files collected from this node | [optional] 
**ExecutionTimeMs** | **int64** | Time taken to collect files from this node in milliseconds | 
**ErrorMessage** | Pointer to **string** | Error message if the overall collection failed on this node | [optional] 
**RemoteTarPath** | Pointer to **string** | Path to the tar.gz archive created on the remote node containing all collected files. The tar file remains on the database node and is NOT downloaded to YBA. Use a separate API to retrieve the tar file if needed.  | [optional] 
**Files** | Pointer to [**[]CollectedFileResult**](CollectedFileResult.md) | List of collected file results | [optional] 

## Methods

### NewNodeFileCollectionResult

`func NewNodeFileCollectionResult(nodeName string, nodeAddress string, success bool, filesCollected int32, filesSkipped int32, filesFailed int32, executionTimeMs int64, ) *NodeFileCollectionResult`

NewNodeFileCollectionResult instantiates a new NodeFileCollectionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeFileCollectionResultWithDefaults

`func NewNodeFileCollectionResultWithDefaults() *NodeFileCollectionResult`

NewNodeFileCollectionResultWithDefaults instantiates a new NodeFileCollectionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeName

`func (o *NodeFileCollectionResult) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *NodeFileCollectionResult) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *NodeFileCollectionResult) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.


### GetNodeAddress

`func (o *NodeFileCollectionResult) GetNodeAddress() string`

GetNodeAddress returns the NodeAddress field if non-nil, zero value otherwise.

### GetNodeAddressOk

`func (o *NodeFileCollectionResult) GetNodeAddressOk() (*string, bool)`

GetNodeAddressOk returns a tuple with the NodeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAddress

`func (o *NodeFileCollectionResult) SetNodeAddress(v string)`

SetNodeAddress sets NodeAddress field to given value.


### GetSuccess

`func (o *NodeFileCollectionResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *NodeFileCollectionResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *NodeFileCollectionResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetFilesCollected

`func (o *NodeFileCollectionResult) GetFilesCollected() int32`

GetFilesCollected returns the FilesCollected field if non-nil, zero value otherwise.

### GetFilesCollectedOk

`func (o *NodeFileCollectionResult) GetFilesCollectedOk() (*int32, bool)`

GetFilesCollectedOk returns a tuple with the FilesCollected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesCollected

`func (o *NodeFileCollectionResult) SetFilesCollected(v int32)`

SetFilesCollected sets FilesCollected field to given value.


### GetFilesSkipped

`func (o *NodeFileCollectionResult) GetFilesSkipped() int32`

GetFilesSkipped returns the FilesSkipped field if non-nil, zero value otherwise.

### GetFilesSkippedOk

`func (o *NodeFileCollectionResult) GetFilesSkippedOk() (*int32, bool)`

GetFilesSkippedOk returns a tuple with the FilesSkipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesSkipped

`func (o *NodeFileCollectionResult) SetFilesSkipped(v int32)`

SetFilesSkipped sets FilesSkipped field to given value.


### GetFilesFailed

`func (o *NodeFileCollectionResult) GetFilesFailed() int32`

GetFilesFailed returns the FilesFailed field if non-nil, zero value otherwise.

### GetFilesFailedOk

`func (o *NodeFileCollectionResult) GetFilesFailedOk() (*int32, bool)`

GetFilesFailedOk returns a tuple with the FilesFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesFailed

`func (o *NodeFileCollectionResult) SetFilesFailed(v int32)`

SetFilesFailed sets FilesFailed field to given value.


### GetTotalBytesCollected

`func (o *NodeFileCollectionResult) GetTotalBytesCollected() int64`

GetTotalBytesCollected returns the TotalBytesCollected field if non-nil, zero value otherwise.

### GetTotalBytesCollectedOk

`func (o *NodeFileCollectionResult) GetTotalBytesCollectedOk() (*int64, bool)`

GetTotalBytesCollectedOk returns a tuple with the TotalBytesCollected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesCollected

`func (o *NodeFileCollectionResult) SetTotalBytesCollected(v int64)`

SetTotalBytesCollected sets TotalBytesCollected field to given value.

### HasTotalBytesCollected

`func (o *NodeFileCollectionResult) HasTotalBytesCollected() bool`

HasTotalBytesCollected returns a boolean if a field has been set.

### GetExecutionTimeMs

`func (o *NodeFileCollectionResult) GetExecutionTimeMs() int64`

GetExecutionTimeMs returns the ExecutionTimeMs field if non-nil, zero value otherwise.

### GetExecutionTimeMsOk

`func (o *NodeFileCollectionResult) GetExecutionTimeMsOk() (*int64, bool)`

GetExecutionTimeMsOk returns a tuple with the ExecutionTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTimeMs

`func (o *NodeFileCollectionResult) SetExecutionTimeMs(v int64)`

SetExecutionTimeMs sets ExecutionTimeMs field to given value.


### GetErrorMessage

`func (o *NodeFileCollectionResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *NodeFileCollectionResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *NodeFileCollectionResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *NodeFileCollectionResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetRemoteTarPath

`func (o *NodeFileCollectionResult) GetRemoteTarPath() string`

GetRemoteTarPath returns the RemoteTarPath field if non-nil, zero value otherwise.

### GetRemoteTarPathOk

`func (o *NodeFileCollectionResult) GetRemoteTarPathOk() (*string, bool)`

GetRemoteTarPathOk returns a tuple with the RemoteTarPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTarPath

`func (o *NodeFileCollectionResult) SetRemoteTarPath(v string)`

SetRemoteTarPath sets RemoteTarPath field to given value.

### HasRemoteTarPath

`func (o *NodeFileCollectionResult) HasRemoteTarPath() bool`

HasRemoteTarPath returns a boolean if a field has been set.

### GetFiles

`func (o *NodeFileCollectionResult) GetFiles() []CollectedFileResult`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *NodeFileCollectionResult) GetFilesOk() (*[]CollectedFileResult, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *NodeFileCollectionResult) SetFiles(v []CollectedFileResult)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *NodeFileCollectionResult) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


