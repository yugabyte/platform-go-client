# UniverseDeleteSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsForceDelete** | Pointer to **bool** | Whether to force delete the universe | [optional] [default to false]
**IsDeleteBackups** | Pointer to **bool** | Whether to delete backups associated with the universe | [optional] [default to false]
**IsDeleteAssociatedCerts** | Pointer to **bool** | Whether to delete associated Encryption In Transit certificates | [optional] [default to false]

## Methods

### NewUniverseDeleteSpec

`func NewUniverseDeleteSpec() *UniverseDeleteSpec`

NewUniverseDeleteSpec instantiates a new UniverseDeleteSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseDeleteSpecWithDefaults

`func NewUniverseDeleteSpecWithDefaults() *UniverseDeleteSpec`

NewUniverseDeleteSpecWithDefaults instantiates a new UniverseDeleteSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsForceDelete

`func (o *UniverseDeleteSpec) GetIsForceDelete() bool`

GetIsForceDelete returns the IsForceDelete field if non-nil, zero value otherwise.

### GetIsForceDeleteOk

`func (o *UniverseDeleteSpec) GetIsForceDeleteOk() (*bool, bool)`

GetIsForceDeleteOk returns a tuple with the IsForceDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForceDelete

`func (o *UniverseDeleteSpec) SetIsForceDelete(v bool)`

SetIsForceDelete sets IsForceDelete field to given value.

### HasIsForceDelete

`func (o *UniverseDeleteSpec) HasIsForceDelete() bool`

HasIsForceDelete returns a boolean if a field has been set.

### GetIsDeleteBackups

`func (o *UniverseDeleteSpec) GetIsDeleteBackups() bool`

GetIsDeleteBackups returns the IsDeleteBackups field if non-nil, zero value otherwise.

### GetIsDeleteBackupsOk

`func (o *UniverseDeleteSpec) GetIsDeleteBackupsOk() (*bool, bool)`

GetIsDeleteBackupsOk returns a tuple with the IsDeleteBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleteBackups

`func (o *UniverseDeleteSpec) SetIsDeleteBackups(v bool)`

SetIsDeleteBackups sets IsDeleteBackups field to given value.

### HasIsDeleteBackups

`func (o *UniverseDeleteSpec) HasIsDeleteBackups() bool`

HasIsDeleteBackups returns a boolean if a field has been set.

### GetIsDeleteAssociatedCerts

`func (o *UniverseDeleteSpec) GetIsDeleteAssociatedCerts() bool`

GetIsDeleteAssociatedCerts returns the IsDeleteAssociatedCerts field if non-nil, zero value otherwise.

### GetIsDeleteAssociatedCertsOk

`func (o *UniverseDeleteSpec) GetIsDeleteAssociatedCertsOk() (*bool, bool)`

GetIsDeleteAssociatedCertsOk returns a tuple with the IsDeleteAssociatedCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleteAssociatedCerts

`func (o *UniverseDeleteSpec) SetIsDeleteAssociatedCerts(v bool)`

SetIsDeleteAssociatedCerts sets IsDeleteAssociatedCerts field to given value.

### HasIsDeleteAssociatedCerts

`func (o *UniverseDeleteSpec) HasIsDeleteAssociatedCerts() bool`

HasIsDeleteAssociatedCerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


