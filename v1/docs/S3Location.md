# S3Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | Pointer to [**PackagePaths**](PackagePaths.md) |  | [optional] 

## Methods

### NewS3Location

`func NewS3Location() *S3Location`

NewS3Location instantiates a new S3Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3LocationWithDefaults

`func NewS3LocationWithDefaults() *S3Location`

NewS3LocationWithDefaults instantiates a new S3Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *S3Location) GetPaths() PackagePaths`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *S3Location) GetPathsOk() (*PackagePaths, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *S3Location) SetPaths(v PackagePaths)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *S3Location) HasPaths() bool`

HasPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


