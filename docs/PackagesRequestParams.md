# PackagesRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchitectureType** | Pointer to **string** | Architecture Type | [optional] 
**ArchiveType** | Pointer to **string** | Archive Type | [optional] 
**BuildNumber** | **string** | Build number | 
**OsType** | Pointer to **string** | OS Type | [optional] 
**PackageName** | Pointer to **string** | Package name | [optional] 

## Methods

### NewPackagesRequestParams

`func NewPackagesRequestParams(buildNumber string, ) *PackagesRequestParams`

NewPackagesRequestParams instantiates a new PackagesRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackagesRequestParamsWithDefaults

`func NewPackagesRequestParamsWithDefaults() *PackagesRequestParams`

NewPackagesRequestParamsWithDefaults instantiates a new PackagesRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitectureType

`func (o *PackagesRequestParams) GetArchitectureType() string`

GetArchitectureType returns the ArchitectureType field if non-nil, zero value otherwise.

### GetArchitectureTypeOk

`func (o *PackagesRequestParams) GetArchitectureTypeOk() (*string, bool)`

GetArchitectureTypeOk returns a tuple with the ArchitectureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitectureType

`func (o *PackagesRequestParams) SetArchitectureType(v string)`

SetArchitectureType sets ArchitectureType field to given value.

### HasArchitectureType

`func (o *PackagesRequestParams) HasArchitectureType() bool`

HasArchitectureType returns a boolean if a field has been set.

### GetArchiveType

`func (o *PackagesRequestParams) GetArchiveType() string`

GetArchiveType returns the ArchiveType field if non-nil, zero value otherwise.

### GetArchiveTypeOk

`func (o *PackagesRequestParams) GetArchiveTypeOk() (*string, bool)`

GetArchiveTypeOk returns a tuple with the ArchiveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveType

`func (o *PackagesRequestParams) SetArchiveType(v string)`

SetArchiveType sets ArchiveType field to given value.

### HasArchiveType

`func (o *PackagesRequestParams) HasArchiveType() bool`

HasArchiveType returns a boolean if a field has been set.

### GetBuildNumber

`func (o *PackagesRequestParams) GetBuildNumber() string`

GetBuildNumber returns the BuildNumber field if non-nil, zero value otherwise.

### GetBuildNumberOk

`func (o *PackagesRequestParams) GetBuildNumberOk() (*string, bool)`

GetBuildNumberOk returns a tuple with the BuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildNumber

`func (o *PackagesRequestParams) SetBuildNumber(v string)`

SetBuildNumber sets BuildNumber field to given value.


### GetOsType

`func (o *PackagesRequestParams) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *PackagesRequestParams) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *PackagesRequestParams) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *PackagesRequestParams) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetPackageName

`func (o *PackagesRequestParams) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *PackagesRequestParams) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *PackagesRequestParams) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *PackagesRequestParams) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


