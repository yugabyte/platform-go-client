# JobConfigSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classname** | **string** | Class name of the job config. | [readonly] 
**Config** | Pointer to **map[string]map[string]interface{}** | Additional custom configuration. | [optional] 

## Methods

### NewJobConfigSpec

`func NewJobConfigSpec(classname string, ) *JobConfigSpec`

NewJobConfigSpec instantiates a new JobConfigSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobConfigSpecWithDefaults

`func NewJobConfigSpecWithDefaults() *JobConfigSpec`

NewJobConfigSpecWithDefaults instantiates a new JobConfigSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassname

`func (o *JobConfigSpec) GetClassname() string`

GetClassname returns the Classname field if non-nil, zero value otherwise.

### GetClassnameOk

`func (o *JobConfigSpec) GetClassnameOk() (*string, bool)`

GetClassnameOk returns a tuple with the Classname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassname

`func (o *JobConfigSpec) SetClassname(v string)`

SetClassname sets Classname field to given value.


### GetConfig

`func (o *JobConfigSpec) GetConfig() map[string]map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *JobConfigSpec) GetConfigOk() (*map[string]map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *JobConfigSpec) SetConfig(v map[string]map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *JobConfigSpec) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


