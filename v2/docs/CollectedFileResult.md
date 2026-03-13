# CollectedFileResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemotePath** | **string** | The original path of the file on the remote node | 
**FileSizeBytes** | Pointer to **int64** | Size of the file in bytes | [optional] 
**Success** | **bool** | Whether the file was successfully collected | 
**ErrorMessage** | Pointer to **string** | Error message if collection failed | [optional] 
**Skipped** | Pointer to **bool** | Whether the file was skipped (e.g., exceeded size limit) | [optional] [default to false]
**SkipReason** | Pointer to **string** | Reason why the file was skipped | [optional] 

## Methods

### NewCollectedFileResult

`func NewCollectedFileResult(remotePath string, success bool, ) *CollectedFileResult`

NewCollectedFileResult instantiates a new CollectedFileResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectedFileResultWithDefaults

`func NewCollectedFileResultWithDefaults() *CollectedFileResult`

NewCollectedFileResultWithDefaults instantiates a new CollectedFileResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemotePath

`func (o *CollectedFileResult) GetRemotePath() string`

GetRemotePath returns the RemotePath field if non-nil, zero value otherwise.

### GetRemotePathOk

`func (o *CollectedFileResult) GetRemotePathOk() (*string, bool)`

GetRemotePathOk returns a tuple with the RemotePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePath

`func (o *CollectedFileResult) SetRemotePath(v string)`

SetRemotePath sets RemotePath field to given value.


### GetFileSizeBytes

`func (o *CollectedFileResult) GetFileSizeBytes() int64`

GetFileSizeBytes returns the FileSizeBytes field if non-nil, zero value otherwise.

### GetFileSizeBytesOk

`func (o *CollectedFileResult) GetFileSizeBytesOk() (*int64, bool)`

GetFileSizeBytesOk returns a tuple with the FileSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeBytes

`func (o *CollectedFileResult) SetFileSizeBytes(v int64)`

SetFileSizeBytes sets FileSizeBytes field to given value.

### HasFileSizeBytes

`func (o *CollectedFileResult) HasFileSizeBytes() bool`

HasFileSizeBytes returns a boolean if a field has been set.

### GetSuccess

`func (o *CollectedFileResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CollectedFileResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CollectedFileResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetErrorMessage

`func (o *CollectedFileResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CollectedFileResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CollectedFileResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CollectedFileResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetSkipped

`func (o *CollectedFileResult) GetSkipped() bool`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *CollectedFileResult) GetSkippedOk() (*bool, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *CollectedFileResult) SetSkipped(v bool)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *CollectedFileResult) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetSkipReason

`func (o *CollectedFileResult) GetSkipReason() string`

GetSkipReason returns the SkipReason field if non-nil, zero value otherwise.

### GetSkipReasonOk

`func (o *CollectedFileResult) GetSkipReasonOk() (*string, bool)`

GetSkipReasonOk returns a tuple with the SkipReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipReason

`func (o *CollectedFileResult) SetSkipReason(v string)`

SetSkipReason sets SkipReason field to given value.

### HasSkipReason

`func (o *CollectedFileResult) HasSkipReason() bool`

HasSkipReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


