# ApiTokenSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiTokenVersion** | Pointer to **int64** | The caller&#39;s current API token version. When set, a new token is generated only if it still matches, so that two concurrent rotations do not silently replace each other. | [optional] 

## Methods

### NewApiTokenSpec

`func NewApiTokenSpec() *ApiTokenSpec`

NewApiTokenSpec instantiates a new ApiTokenSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTokenSpecWithDefaults

`func NewApiTokenSpecWithDefaults() *ApiTokenSpec`

NewApiTokenSpecWithDefaults instantiates a new ApiTokenSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiTokenVersion

`func (o *ApiTokenSpec) GetApiTokenVersion() int64`

GetApiTokenVersion returns the ApiTokenVersion field if non-nil, zero value otherwise.

### GetApiTokenVersionOk

`func (o *ApiTokenSpec) GetApiTokenVersionOk() (*int64, bool)`

GetApiTokenVersionOk returns a tuple with the ApiTokenVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTokenVersion

`func (o *ApiTokenSpec) SetApiTokenVersion(v int64)`

SetApiTokenVersion sets ApiTokenVersion field to given value.

### HasApiTokenVersion

`func (o *ApiTokenSpec) HasApiTokenVersion() bool`

HasApiTokenVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


