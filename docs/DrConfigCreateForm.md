# DrConfigCreateForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootstrapBackupParams** | Pointer to [**BootstarpBackupParams**](BootstarpBackupParams.md) |  | [optional] 
**Dbs** | **[]string** | Source Universe DB IDs | 
**DryRun** | Pointer to **bool** | Run the pre-checks without actually running the subtasks | [optional] 
**Name** | **string** | Name | 
**PitrParams** | Pointer to [**PitrParams**](PitrParams.md) |  | [optional] 
**SourceUniverseUUID** | **string** | Source Universe UUID | 
**TargetUniverseUUID** | **string** | Target Universe UUID | 

## Methods

### NewDrConfigCreateForm

`func NewDrConfigCreateForm(dbs []string, name string, sourceUniverseUUID string, targetUniverseUUID string, ) *DrConfigCreateForm`

NewDrConfigCreateForm instantiates a new DrConfigCreateForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrConfigCreateFormWithDefaults

`func NewDrConfigCreateFormWithDefaults() *DrConfigCreateForm`

NewDrConfigCreateFormWithDefaults instantiates a new DrConfigCreateForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootstrapBackupParams

`func (o *DrConfigCreateForm) GetBootstrapBackupParams() BootstarpBackupParams`

GetBootstrapBackupParams returns the BootstrapBackupParams field if non-nil, zero value otherwise.

### GetBootstrapBackupParamsOk

`func (o *DrConfigCreateForm) GetBootstrapBackupParamsOk() (*BootstarpBackupParams, bool)`

GetBootstrapBackupParamsOk returns a tuple with the BootstrapBackupParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapBackupParams

`func (o *DrConfigCreateForm) SetBootstrapBackupParams(v BootstarpBackupParams)`

SetBootstrapBackupParams sets BootstrapBackupParams field to given value.

### HasBootstrapBackupParams

`func (o *DrConfigCreateForm) HasBootstrapBackupParams() bool`

HasBootstrapBackupParams returns a boolean if a field has been set.

### GetDbs

`func (o *DrConfigCreateForm) GetDbs() []string`

GetDbs returns the Dbs field if non-nil, zero value otherwise.

### GetDbsOk

`func (o *DrConfigCreateForm) GetDbsOk() (*[]string, bool)`

GetDbsOk returns a tuple with the Dbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbs

`func (o *DrConfigCreateForm) SetDbs(v []string)`

SetDbs sets Dbs field to given value.


### GetDryRun

`func (o *DrConfigCreateForm) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DrConfigCreateForm) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DrConfigCreateForm) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DrConfigCreateForm) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetName

`func (o *DrConfigCreateForm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DrConfigCreateForm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DrConfigCreateForm) SetName(v string)`

SetName sets Name field to given value.


### GetPitrParams

`func (o *DrConfigCreateForm) GetPitrParams() PitrParams`

GetPitrParams returns the PitrParams field if non-nil, zero value otherwise.

### GetPitrParamsOk

`func (o *DrConfigCreateForm) GetPitrParamsOk() (*PitrParams, bool)`

GetPitrParamsOk returns a tuple with the PitrParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitrParams

`func (o *DrConfigCreateForm) SetPitrParams(v PitrParams)`

SetPitrParams sets PitrParams field to given value.

### HasPitrParams

`func (o *DrConfigCreateForm) HasPitrParams() bool`

HasPitrParams returns a boolean if a field has been set.

### GetSourceUniverseUUID

`func (o *DrConfigCreateForm) GetSourceUniverseUUID() string`

GetSourceUniverseUUID returns the SourceUniverseUUID field if non-nil, zero value otherwise.

### GetSourceUniverseUUIDOk

`func (o *DrConfigCreateForm) GetSourceUniverseUUIDOk() (*string, bool)`

GetSourceUniverseUUIDOk returns a tuple with the SourceUniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUniverseUUID

`func (o *DrConfigCreateForm) SetSourceUniverseUUID(v string)`

SetSourceUniverseUUID sets SourceUniverseUUID field to given value.


### GetTargetUniverseUUID

`func (o *DrConfigCreateForm) GetTargetUniverseUUID() string`

GetTargetUniverseUUID returns the TargetUniverseUUID field if non-nil, zero value otherwise.

### GetTargetUniverseUUIDOk

`func (o *DrConfigCreateForm) GetTargetUniverseUUIDOk() (*string, bool)`

GetTargetUniverseUUIDOk returns a tuple with the TargetUniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUniverseUUID

`func (o *DrConfigCreateForm) SetTargetUniverseUUID(v string)`

SetTargetUniverseUUID sets TargetUniverseUUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


