# TaskVersionNumbers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**YbPrevSoftwareVersion** | Pointer to **string** | YB software version before the upgrade started. | [optional] [readonly] 
**YbSoftwareVersion** | Pointer to **string** | Target YB software version for the upgrade. | [optional] [readonly] 

## Methods

### NewTaskVersionNumbers

`func NewTaskVersionNumbers() *TaskVersionNumbers`

NewTaskVersionNumbers instantiates a new TaskVersionNumbers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskVersionNumbersWithDefaults

`func NewTaskVersionNumbersWithDefaults() *TaskVersionNumbers`

NewTaskVersionNumbersWithDefaults instantiates a new TaskVersionNumbers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetYbPrevSoftwareVersion

`func (o *TaskVersionNumbers) GetYbPrevSoftwareVersion() string`

GetYbPrevSoftwareVersion returns the YbPrevSoftwareVersion field if non-nil, zero value otherwise.

### GetYbPrevSoftwareVersionOk

`func (o *TaskVersionNumbers) GetYbPrevSoftwareVersionOk() (*string, bool)`

GetYbPrevSoftwareVersionOk returns a tuple with the YbPrevSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbPrevSoftwareVersion

`func (o *TaskVersionNumbers) SetYbPrevSoftwareVersion(v string)`

SetYbPrevSoftwareVersion sets YbPrevSoftwareVersion field to given value.

### HasYbPrevSoftwareVersion

`func (o *TaskVersionNumbers) HasYbPrevSoftwareVersion() bool`

HasYbPrevSoftwareVersion returns a boolean if a field has been set.

### GetYbSoftwareVersion

`func (o *TaskVersionNumbers) GetYbSoftwareVersion() string`

GetYbSoftwareVersion returns the YbSoftwareVersion field if non-nil, zero value otherwise.

### GetYbSoftwareVersionOk

`func (o *TaskVersionNumbers) GetYbSoftwareVersionOk() (*string, bool)`

GetYbSoftwareVersionOk returns a tuple with the YbSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbSoftwareVersion

`func (o *TaskVersionNumbers) SetYbSoftwareVersion(v string)`

SetYbSoftwareVersion sets YbSoftwareVersion field to given value.

### HasYbSoftwareVersion

`func (o *TaskVersionNumbers) HasYbSoftwareVersion() bool`

HasYbSoftwareVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


