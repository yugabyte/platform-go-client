# DetachUniverseSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkipReleases** | Pointer to **bool** | Whether to skip releases during detach operation. | [optional] [default to false]

## Methods

### NewDetachUniverseSpec

`func NewDetachUniverseSpec() *DetachUniverseSpec`

NewDetachUniverseSpec instantiates a new DetachUniverseSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetachUniverseSpecWithDefaults

`func NewDetachUniverseSpecWithDefaults() *DetachUniverseSpec`

NewDetachUniverseSpecWithDefaults instantiates a new DetachUniverseSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkipReleases

`func (o *DetachUniverseSpec) GetSkipReleases() bool`

GetSkipReleases returns the SkipReleases field if non-nil, zero value otherwise.

### GetSkipReleasesOk

`func (o *DetachUniverseSpec) GetSkipReleasesOk() (*bool, bool)`

GetSkipReleasesOk returns a tuple with the SkipReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipReleases

`func (o *DetachUniverseSpec) SetSkipReleases(v bool)`

SetSkipReleases sets SkipReleases field to given value.

### HasSkipReleases

`func (o *DetachUniverseSpec) HasSkipReleases() bool`

HasSkipReleases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


