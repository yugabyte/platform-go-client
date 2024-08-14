# UpdateRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifacts** | [**[]Artifact**](Artifact.md) |  | 
**ReleaseDate** | **int64** |  | 
**ReleaseNotes** | **string** |  | 
**ReleaseTag** | **string** |  | 
**State** | **string** |  | 

## Methods

### NewUpdateRelease

`func NewUpdateRelease(artifacts []Artifact, releaseDate int64, releaseNotes string, releaseTag string, state string, ) *UpdateRelease`

NewUpdateRelease instantiates a new UpdateRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateReleaseWithDefaults

`func NewUpdateReleaseWithDefaults() *UpdateRelease`

NewUpdateReleaseWithDefaults instantiates a new UpdateRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifacts

`func (o *UpdateRelease) GetArtifacts() []Artifact`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *UpdateRelease) GetArtifactsOk() (*[]Artifact, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *UpdateRelease) SetArtifacts(v []Artifact)`

SetArtifacts sets Artifacts field to given value.


### GetReleaseDate

`func (o *UpdateRelease) GetReleaseDate() int64`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *UpdateRelease) GetReleaseDateOk() (*int64, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *UpdateRelease) SetReleaseDate(v int64)`

SetReleaseDate sets ReleaseDate field to given value.


### GetReleaseNotes

`func (o *UpdateRelease) GetReleaseNotes() string`

GetReleaseNotes returns the ReleaseNotes field if non-nil, zero value otherwise.

### GetReleaseNotesOk

`func (o *UpdateRelease) GetReleaseNotesOk() (*string, bool)`

GetReleaseNotesOk returns a tuple with the ReleaseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotes

`func (o *UpdateRelease) SetReleaseNotes(v string)`

SetReleaseNotes sets ReleaseNotes field to given value.


### GetReleaseTag

`func (o *UpdateRelease) GetReleaseTag() string`

GetReleaseTag returns the ReleaseTag field if non-nil, zero value otherwise.

### GetReleaseTagOk

`func (o *UpdateRelease) GetReleaseTagOk() (*string, bool)`

GetReleaseTagOk returns a tuple with the ReleaseTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTag

`func (o *UpdateRelease) SetReleaseTag(v string)`

SetReleaseTag sets ReleaseTag field to given value.


### GetState

`func (o *UpdateRelease) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateRelease) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateRelease) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


