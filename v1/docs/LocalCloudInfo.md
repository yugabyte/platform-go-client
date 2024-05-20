# LocalCloudInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataHomeDir** | **string** |  | 
**EnvVars** | **map[string]string** |  | 
**YbcBinDir** | **string** |  | 
**YugabyteBinDir** | **string** |  | 

## Methods

### NewLocalCloudInfo

`func NewLocalCloudInfo(dataHomeDir string, envVars map[string]string, ybcBinDir string, yugabyteBinDir string, ) *LocalCloudInfo`

NewLocalCloudInfo instantiates a new LocalCloudInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalCloudInfoWithDefaults

`func NewLocalCloudInfoWithDefaults() *LocalCloudInfo`

NewLocalCloudInfoWithDefaults instantiates a new LocalCloudInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataHomeDir

`func (o *LocalCloudInfo) GetDataHomeDir() string`

GetDataHomeDir returns the DataHomeDir field if non-nil, zero value otherwise.

### GetDataHomeDirOk

`func (o *LocalCloudInfo) GetDataHomeDirOk() (*string, bool)`

GetDataHomeDirOk returns a tuple with the DataHomeDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataHomeDir

`func (o *LocalCloudInfo) SetDataHomeDir(v string)`

SetDataHomeDir sets DataHomeDir field to given value.


### GetEnvVars

`func (o *LocalCloudInfo) GetEnvVars() map[string]string`

GetEnvVars returns the EnvVars field if non-nil, zero value otherwise.

### GetEnvVarsOk

`func (o *LocalCloudInfo) GetEnvVarsOk() (*map[string]string, bool)`

GetEnvVarsOk returns a tuple with the EnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVars

`func (o *LocalCloudInfo) SetEnvVars(v map[string]string)`

SetEnvVars sets EnvVars field to given value.


### GetYbcBinDir

`func (o *LocalCloudInfo) GetYbcBinDir() string`

GetYbcBinDir returns the YbcBinDir field if non-nil, zero value otherwise.

### GetYbcBinDirOk

`func (o *LocalCloudInfo) GetYbcBinDirOk() (*string, bool)`

GetYbcBinDirOk returns a tuple with the YbcBinDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbcBinDir

`func (o *LocalCloudInfo) SetYbcBinDir(v string)`

SetYbcBinDir sets YbcBinDir field to given value.


### GetYugabyteBinDir

`func (o *LocalCloudInfo) GetYugabyteBinDir() string`

GetYugabyteBinDir returns the YugabyteBinDir field if non-nil, zero value otherwise.

### GetYugabyteBinDirOk

`func (o *LocalCloudInfo) GetYugabyteBinDirOk() (*string, bool)`

GetYugabyteBinDirOk returns a tuple with the YugabyteBinDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYugabyteBinDir

`func (o *LocalCloudInfo) SetYugabyteBinDir(v string)`

SetYugabyteBinDir sets YugabyteBinDir field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


