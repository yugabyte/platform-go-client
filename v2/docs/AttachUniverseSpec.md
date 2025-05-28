# AttachUniverseSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadedSpecFile** | Pointer to ***os.File** | The uploaded tgz file containing universe metadata. | [optional] 

## Methods

### NewAttachUniverseSpec

`func NewAttachUniverseSpec() *AttachUniverseSpec`

NewAttachUniverseSpec instantiates a new AttachUniverseSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachUniverseSpecWithDefaults

`func NewAttachUniverseSpecWithDefaults() *AttachUniverseSpec`

NewAttachUniverseSpecWithDefaults instantiates a new AttachUniverseSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadedSpecFile

`func (o *AttachUniverseSpec) GetDownloadedSpecFile() *os.File`

GetDownloadedSpecFile returns the DownloadedSpecFile field if non-nil, zero value otherwise.

### GetDownloadedSpecFileOk

`func (o *AttachUniverseSpec) GetDownloadedSpecFileOk() (**os.File, bool)`

GetDownloadedSpecFileOk returns a tuple with the DownloadedSpecFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadedSpecFile

`func (o *AttachUniverseSpec) SetDownloadedSpecFile(v *os.File)`

SetDownloadedSpecFile sets DownloadedSpecFile field to given value.

### HasDownloadedSpecFile

`func (o *AttachUniverseSpec) HasDownloadedSpecFile() bool`

HasDownloadedSpecFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


