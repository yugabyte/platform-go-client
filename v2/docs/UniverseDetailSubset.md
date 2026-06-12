# UniverseDetailSubset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Universe unique identifier. | [optional] 
**Name** | Pointer to **string** | Universe display name. | [optional] 
**UpdateInProgress** | Pointer to **bool** | Whether a universe update operation is currently running. | [optional] 
**UpdateSucceeded** | Pointer to **bool** | Whether the last universe update attempt completed successfully. | [optional] 
**CreationDate** | Pointer to **time.Time** | Universe creation time (RFC 3339 / ISO-8601). | [optional] [readonly] 
**UniversePaused** | Pointer to **bool** | Whether the universe is in a paused state. | [optional] 

## Methods

### NewUniverseDetailSubset

`func NewUniverseDetailSubset() *UniverseDetailSubset`

NewUniverseDetailSubset instantiates a new UniverseDetailSubset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseDetailSubsetWithDefaults

`func NewUniverseDetailSubsetWithDefaults() *UniverseDetailSubset`

NewUniverseDetailSubsetWithDefaults instantiates a new UniverseDetailSubset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *UniverseDetailSubset) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UniverseDetailSubset) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UniverseDetailSubset) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UniverseDetailSubset) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *UniverseDetailSubset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UniverseDetailSubset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UniverseDetailSubset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UniverseDetailSubset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdateInProgress

`func (o *UniverseDetailSubset) GetUpdateInProgress() bool`

GetUpdateInProgress returns the UpdateInProgress field if non-nil, zero value otherwise.

### GetUpdateInProgressOk

`func (o *UniverseDetailSubset) GetUpdateInProgressOk() (*bool, bool)`

GetUpdateInProgressOk returns a tuple with the UpdateInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInProgress

`func (o *UniverseDetailSubset) SetUpdateInProgress(v bool)`

SetUpdateInProgress sets UpdateInProgress field to given value.

### HasUpdateInProgress

`func (o *UniverseDetailSubset) HasUpdateInProgress() bool`

HasUpdateInProgress returns a boolean if a field has been set.

### GetUpdateSucceeded

`func (o *UniverseDetailSubset) GetUpdateSucceeded() bool`

GetUpdateSucceeded returns the UpdateSucceeded field if non-nil, zero value otherwise.

### GetUpdateSucceededOk

`func (o *UniverseDetailSubset) GetUpdateSucceededOk() (*bool, bool)`

GetUpdateSucceededOk returns a tuple with the UpdateSucceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSucceeded

`func (o *UniverseDetailSubset) SetUpdateSucceeded(v bool)`

SetUpdateSucceeded sets UpdateSucceeded field to given value.

### HasUpdateSucceeded

`func (o *UniverseDetailSubset) HasUpdateSucceeded() bool`

HasUpdateSucceeded returns a boolean if a field has been set.

### GetCreationDate

`func (o *UniverseDetailSubset) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *UniverseDetailSubset) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *UniverseDetailSubset) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *UniverseDetailSubset) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetUniversePaused

`func (o *UniverseDetailSubset) GetUniversePaused() bool`

GetUniversePaused returns the UniversePaused field if non-nil, zero value otherwise.

### GetUniversePausedOk

`func (o *UniverseDetailSubset) GetUniversePausedOk() (*bool, bool)`

GetUniversePausedOk returns a tuple with the UniversePaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversePaused

`func (o *UniverseDetailSubset) SetUniversePaused(v bool)`

SetUniversePaused sets UniversePaused field to given value.

### HasUniversePaused

`func (o *UniverseDetailSubset) HasUniversePaused() bool`

HasUniversePaused returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


