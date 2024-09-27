# IsolatedBackupCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalDir** | **string** | local directory to store backup | 
**Components** | [**[]YbaComponent**](YbaComponent.md) | the components to include in YBA backup | 

## Methods

### NewIsolatedBackupCreateSpec

`func NewIsolatedBackupCreateSpec(localDir string, components []YbaComponent, ) *IsolatedBackupCreateSpec`

NewIsolatedBackupCreateSpec instantiates a new IsolatedBackupCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsolatedBackupCreateSpecWithDefaults

`func NewIsolatedBackupCreateSpecWithDefaults() *IsolatedBackupCreateSpec`

NewIsolatedBackupCreateSpecWithDefaults instantiates a new IsolatedBackupCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalDir

`func (o *IsolatedBackupCreateSpec) GetLocalDir() string`

GetLocalDir returns the LocalDir field if non-nil, zero value otherwise.

### GetLocalDirOk

`func (o *IsolatedBackupCreateSpec) GetLocalDirOk() (*string, bool)`

GetLocalDirOk returns a tuple with the LocalDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDir

`func (o *IsolatedBackupCreateSpec) SetLocalDir(v string)`

SetLocalDir sets LocalDir field to given value.


### GetComponents

`func (o *IsolatedBackupCreateSpec) GetComponents() []YbaComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *IsolatedBackupCreateSpec) GetComponentsOk() (*[]YbaComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *IsolatedBackupCreateSpec) SetComponents(v []YbaComponent)`

SetComponents sets Components field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


