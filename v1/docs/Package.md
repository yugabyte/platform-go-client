# Package

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewPackage

`func NewPackage() *Package`

NewPackage instantiates a new Package object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageWithDefaults

`func NewPackageWithDefaults() *Package`

NewPackageWithDefaults instantiates a new Package object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *Package) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *Package) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *Package) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *Package) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetPath

`func (o *Package) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Package) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Package) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Package) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


