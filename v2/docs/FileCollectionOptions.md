# FileCollectionOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePaths** | **[]string** | List of file paths to collect from each node. Paths can be absolute or relative to the yugabyte home directory. Glob patterns are not supported.  | 
**DirectoryPaths** | Pointer to **[]string** | List of directory paths to collect files from. All files in these directories will be collected (non-recursive). Use max_depth for recursive collection.  | [optional] 
**MaxDepth** | Pointer to **int32** | Maximum depth for directory traversal when collecting from directories. 1 means only immediate files, 2 includes one level of subdirectories, etc.  | [optional] [default to 1]
**MaxFileSizeBytes** | Pointer to **int64** | Maximum size of individual files to collect in bytes. Files larger than this will be skipped and reported in the response.  | [optional] [default to 10485760]
**MaxTotalSizeBytes** | Pointer to **int64** | Maximum total size of all files to collect per node in bytes. Collection stops when this limit is reached.  | [optional] [default to 104857600]
**TimeoutSecs** | Pointer to **int64** | Timeout in seconds for file collection on each node.  | [optional] [default to 300]
**LinuxUser** | Pointer to **string** | Run file collection as a particular Linux user. Defaults to yugabyte. When using Node Agent, the command is executed as this user. When using SSH fallback, this sets the SSH user.  | [optional] [default to "yugabyte"]

## Methods

### NewFileCollectionOptions

`func NewFileCollectionOptions(filePaths []string, ) *FileCollectionOptions`

NewFileCollectionOptions instantiates a new FileCollectionOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileCollectionOptionsWithDefaults

`func NewFileCollectionOptionsWithDefaults() *FileCollectionOptions`

NewFileCollectionOptionsWithDefaults instantiates a new FileCollectionOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePaths

`func (o *FileCollectionOptions) GetFilePaths() []string`

GetFilePaths returns the FilePaths field if non-nil, zero value otherwise.

### GetFilePathsOk

`func (o *FileCollectionOptions) GetFilePathsOk() (*[]string, bool)`

GetFilePathsOk returns a tuple with the FilePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePaths

`func (o *FileCollectionOptions) SetFilePaths(v []string)`

SetFilePaths sets FilePaths field to given value.


### GetDirectoryPaths

`func (o *FileCollectionOptions) GetDirectoryPaths() []string`

GetDirectoryPaths returns the DirectoryPaths field if non-nil, zero value otherwise.

### GetDirectoryPathsOk

`func (o *FileCollectionOptions) GetDirectoryPathsOk() (*[]string, bool)`

GetDirectoryPathsOk returns a tuple with the DirectoryPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryPaths

`func (o *FileCollectionOptions) SetDirectoryPaths(v []string)`

SetDirectoryPaths sets DirectoryPaths field to given value.

### HasDirectoryPaths

`func (o *FileCollectionOptions) HasDirectoryPaths() bool`

HasDirectoryPaths returns a boolean if a field has been set.

### GetMaxDepth

`func (o *FileCollectionOptions) GetMaxDepth() int32`

GetMaxDepth returns the MaxDepth field if non-nil, zero value otherwise.

### GetMaxDepthOk

`func (o *FileCollectionOptions) GetMaxDepthOk() (*int32, bool)`

GetMaxDepthOk returns a tuple with the MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDepth

`func (o *FileCollectionOptions) SetMaxDepth(v int32)`

SetMaxDepth sets MaxDepth field to given value.

### HasMaxDepth

`func (o *FileCollectionOptions) HasMaxDepth() bool`

HasMaxDepth returns a boolean if a field has been set.

### GetMaxFileSizeBytes

`func (o *FileCollectionOptions) GetMaxFileSizeBytes() int64`

GetMaxFileSizeBytes returns the MaxFileSizeBytes field if non-nil, zero value otherwise.

### GetMaxFileSizeBytesOk

`func (o *FileCollectionOptions) GetMaxFileSizeBytesOk() (*int64, bool)`

GetMaxFileSizeBytesOk returns a tuple with the MaxFileSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFileSizeBytes

`func (o *FileCollectionOptions) SetMaxFileSizeBytes(v int64)`

SetMaxFileSizeBytes sets MaxFileSizeBytes field to given value.

### HasMaxFileSizeBytes

`func (o *FileCollectionOptions) HasMaxFileSizeBytes() bool`

HasMaxFileSizeBytes returns a boolean if a field has been set.

### GetMaxTotalSizeBytes

`func (o *FileCollectionOptions) GetMaxTotalSizeBytes() int64`

GetMaxTotalSizeBytes returns the MaxTotalSizeBytes field if non-nil, zero value otherwise.

### GetMaxTotalSizeBytesOk

`func (o *FileCollectionOptions) GetMaxTotalSizeBytesOk() (*int64, bool)`

GetMaxTotalSizeBytesOk returns a tuple with the MaxTotalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalSizeBytes

`func (o *FileCollectionOptions) SetMaxTotalSizeBytes(v int64)`

SetMaxTotalSizeBytes sets MaxTotalSizeBytes field to given value.

### HasMaxTotalSizeBytes

`func (o *FileCollectionOptions) HasMaxTotalSizeBytes() bool`

HasMaxTotalSizeBytes returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *FileCollectionOptions) GetTimeoutSecs() int64`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *FileCollectionOptions) GetTimeoutSecsOk() (*int64, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *FileCollectionOptions) SetTimeoutSecs(v int64)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *FileCollectionOptions) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### GetLinuxUser

`func (o *FileCollectionOptions) GetLinuxUser() string`

GetLinuxUser returns the LinuxUser field if non-nil, zero value otherwise.

### GetLinuxUserOk

`func (o *FileCollectionOptions) GetLinuxUserOk() (*string, bool)`

GetLinuxUserOk returns a tuple with the LinuxUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUser

`func (o *FileCollectionOptions) SetLinuxUser(v string)`

SetLinuxUser sets LinuxUser field to given value.

### HasLinuxUser

`func (o *FileCollectionOptions) HasLinuxUser() bool`

HasLinuxUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


