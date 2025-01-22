# UniverseResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedTasks** | Pointer to [**AllowedUniverseTasksResp**](AllowedUniverseTasksResp.md) |  | [optional] 
**CreationDate** | Pointer to **string** | Universe creation date | [optional] 
**DnsName** | Pointer to **string** | DNS name | [optional] 
**DrConfigUuidsAsSource** | Pointer to **[]string** | UUIDs of DR configs where this universe is the source (primary) | [optional] 
**DrConfigUuidsAsTarget** | Pointer to **[]string** | UUIDs of DR configs where this universe is the target (secondary) | [optional] 
**Name** | Pointer to **string** | Universe name | [optional] 
**PricePerHour** | Pointer to **float64** | Price | [optional] 
**Resources** | Pointer to [**UniverseResourceDetails**](UniverseResourceDetails.md) |  | [optional] 
**RollMaxBatchSize** | Pointer to [**RollMaxBatchSize**](RollMaxBatchSize.md) |  | [optional] 
**SampleAppCommandTxt** | Pointer to **string** | Sample command | [optional] 
**TaskUUID** | Pointer to **string** | Task UUID | [optional] 
**UniverseConfig** | Pointer to **map[string]string** | Universe configuration | [optional] 
**UniverseDetails** | Pointer to [**UniverseDefinitionTaskParamsResp**](UniverseDefinitionTaskParamsResp.md) |  | [optional] 
**UniverseUUID** | Pointer to **string** | Universe UUID | [optional] 
**Version** | Pointer to **int32** | Universe version | [optional] 

## Methods

### NewUniverseResp

`func NewUniverseResp() *UniverseResp`

NewUniverseResp instantiates a new UniverseResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseRespWithDefaults

`func NewUniverseRespWithDefaults() *UniverseResp`

NewUniverseRespWithDefaults instantiates a new UniverseResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedTasks

`func (o *UniverseResp) GetAllowedTasks() AllowedUniverseTasksResp`

GetAllowedTasks returns the AllowedTasks field if non-nil, zero value otherwise.

### GetAllowedTasksOk

`func (o *UniverseResp) GetAllowedTasksOk() (*AllowedUniverseTasksResp, bool)`

GetAllowedTasksOk returns a tuple with the AllowedTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTasks

`func (o *UniverseResp) SetAllowedTasks(v AllowedUniverseTasksResp)`

SetAllowedTasks sets AllowedTasks field to given value.

### HasAllowedTasks

`func (o *UniverseResp) HasAllowedTasks() bool`

HasAllowedTasks returns a boolean if a field has been set.

### GetCreationDate

`func (o *UniverseResp) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *UniverseResp) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *UniverseResp) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *UniverseResp) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDnsName

`func (o *UniverseResp) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *UniverseResp) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *UniverseResp) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *UniverseResp) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetDrConfigUuidsAsSource

`func (o *UniverseResp) GetDrConfigUuidsAsSource() []string`

GetDrConfigUuidsAsSource returns the DrConfigUuidsAsSource field if non-nil, zero value otherwise.

### GetDrConfigUuidsAsSourceOk

`func (o *UniverseResp) GetDrConfigUuidsAsSourceOk() (*[]string, bool)`

GetDrConfigUuidsAsSourceOk returns a tuple with the DrConfigUuidsAsSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrConfigUuidsAsSource

`func (o *UniverseResp) SetDrConfigUuidsAsSource(v []string)`

SetDrConfigUuidsAsSource sets DrConfigUuidsAsSource field to given value.

### HasDrConfigUuidsAsSource

`func (o *UniverseResp) HasDrConfigUuidsAsSource() bool`

HasDrConfigUuidsAsSource returns a boolean if a field has been set.

### GetDrConfigUuidsAsTarget

`func (o *UniverseResp) GetDrConfigUuidsAsTarget() []string`

GetDrConfigUuidsAsTarget returns the DrConfigUuidsAsTarget field if non-nil, zero value otherwise.

### GetDrConfigUuidsAsTargetOk

`func (o *UniverseResp) GetDrConfigUuidsAsTargetOk() (*[]string, bool)`

GetDrConfigUuidsAsTargetOk returns a tuple with the DrConfigUuidsAsTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrConfigUuidsAsTarget

`func (o *UniverseResp) SetDrConfigUuidsAsTarget(v []string)`

SetDrConfigUuidsAsTarget sets DrConfigUuidsAsTarget field to given value.

### HasDrConfigUuidsAsTarget

`func (o *UniverseResp) HasDrConfigUuidsAsTarget() bool`

HasDrConfigUuidsAsTarget returns a boolean if a field has been set.

### GetName

`func (o *UniverseResp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UniverseResp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UniverseResp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UniverseResp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPricePerHour

`func (o *UniverseResp) GetPricePerHour() float64`

GetPricePerHour returns the PricePerHour field if non-nil, zero value otherwise.

### GetPricePerHourOk

`func (o *UniverseResp) GetPricePerHourOk() (*float64, bool)`

GetPricePerHourOk returns a tuple with the PricePerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerHour

`func (o *UniverseResp) SetPricePerHour(v float64)`

SetPricePerHour sets PricePerHour field to given value.

### HasPricePerHour

`func (o *UniverseResp) HasPricePerHour() bool`

HasPricePerHour returns a boolean if a field has been set.

### GetResources

`func (o *UniverseResp) GetResources() UniverseResourceDetails`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *UniverseResp) GetResourcesOk() (*UniverseResourceDetails, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *UniverseResp) SetResources(v UniverseResourceDetails)`

SetResources sets Resources field to given value.

### HasResources

`func (o *UniverseResp) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRollMaxBatchSize

`func (o *UniverseResp) GetRollMaxBatchSize() RollMaxBatchSize`

GetRollMaxBatchSize returns the RollMaxBatchSize field if non-nil, zero value otherwise.

### GetRollMaxBatchSizeOk

`func (o *UniverseResp) GetRollMaxBatchSizeOk() (*RollMaxBatchSize, bool)`

GetRollMaxBatchSizeOk returns a tuple with the RollMaxBatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollMaxBatchSize

`func (o *UniverseResp) SetRollMaxBatchSize(v RollMaxBatchSize)`

SetRollMaxBatchSize sets RollMaxBatchSize field to given value.

### HasRollMaxBatchSize

`func (o *UniverseResp) HasRollMaxBatchSize() bool`

HasRollMaxBatchSize returns a boolean if a field has been set.

### GetSampleAppCommandTxt

`func (o *UniverseResp) GetSampleAppCommandTxt() string`

GetSampleAppCommandTxt returns the SampleAppCommandTxt field if non-nil, zero value otherwise.

### GetSampleAppCommandTxtOk

`func (o *UniverseResp) GetSampleAppCommandTxtOk() (*string, bool)`

GetSampleAppCommandTxtOk returns a tuple with the SampleAppCommandTxt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleAppCommandTxt

`func (o *UniverseResp) SetSampleAppCommandTxt(v string)`

SetSampleAppCommandTxt sets SampleAppCommandTxt field to given value.

### HasSampleAppCommandTxt

`func (o *UniverseResp) HasSampleAppCommandTxt() bool`

HasSampleAppCommandTxt returns a boolean if a field has been set.

### GetTaskUUID

`func (o *UniverseResp) GetTaskUUID() string`

GetTaskUUID returns the TaskUUID field if non-nil, zero value otherwise.

### GetTaskUUIDOk

`func (o *UniverseResp) GetTaskUUIDOk() (*string, bool)`

GetTaskUUIDOk returns a tuple with the TaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUUID

`func (o *UniverseResp) SetTaskUUID(v string)`

SetTaskUUID sets TaskUUID field to given value.

### HasTaskUUID

`func (o *UniverseResp) HasTaskUUID() bool`

HasTaskUUID returns a boolean if a field has been set.

### GetUniverseConfig

`func (o *UniverseResp) GetUniverseConfig() map[string]string`

GetUniverseConfig returns the UniverseConfig field if non-nil, zero value otherwise.

### GetUniverseConfigOk

`func (o *UniverseResp) GetUniverseConfigOk() (*map[string]string, bool)`

GetUniverseConfigOk returns a tuple with the UniverseConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseConfig

`func (o *UniverseResp) SetUniverseConfig(v map[string]string)`

SetUniverseConfig sets UniverseConfig field to given value.

### HasUniverseConfig

`func (o *UniverseResp) HasUniverseConfig() bool`

HasUniverseConfig returns a boolean if a field has been set.

### GetUniverseDetails

`func (o *UniverseResp) GetUniverseDetails() UniverseDefinitionTaskParamsResp`

GetUniverseDetails returns the UniverseDetails field if non-nil, zero value otherwise.

### GetUniverseDetailsOk

`func (o *UniverseResp) GetUniverseDetailsOk() (*UniverseDefinitionTaskParamsResp, bool)`

GetUniverseDetailsOk returns a tuple with the UniverseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseDetails

`func (o *UniverseResp) SetUniverseDetails(v UniverseDefinitionTaskParamsResp)`

SetUniverseDetails sets UniverseDetails field to given value.

### HasUniverseDetails

`func (o *UniverseResp) HasUniverseDetails() bool`

HasUniverseDetails returns a boolean if a field has been set.

### GetUniverseUUID

`func (o *UniverseResp) GetUniverseUUID() string`

GetUniverseUUID returns the UniverseUUID field if non-nil, zero value otherwise.

### GetUniverseUUIDOk

`func (o *UniverseResp) GetUniverseUUIDOk() (*string, bool)`

GetUniverseUUIDOk returns a tuple with the UniverseUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUUID

`func (o *UniverseResp) SetUniverseUUID(v string)`

SetUniverseUUID sets UniverseUUID field to given value.

### HasUniverseUUID

`func (o *UniverseResp) HasUniverseUUID() bool`

HasUniverseUUID returns a boolean if a field has been set.

### GetVersion

`func (o *UniverseResp) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UniverseResp) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UniverseResp) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UniverseResp) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


