# InstanceTypeResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | True if the instance is active | [optional] [readonly] 
**IdKey** | [**InstanceTypeKey**](InstanceTypeKey.md) |  | 
**InstanceTypeCode** | Pointer to **string** | Instance type code | [optional] [readonly] 
**InstanceTypeDetails** | Pointer to [**InstanceTypeDetails**](InstanceTypeDetails.md) |  | [optional] 
**MemSizeGB** | Pointer to **float64** | The instance&#39;s memory size, in gigabytes | [optional] 
**NumCores** | Pointer to **float64** | The instance&#39;s number of CPU cores | [optional] 
**ProviderCode** | Pointer to **string** | Cloud provider code | [optional] [readonly] 
**ProviderUuid** | Pointer to **string** | Provider UUID | [optional] [readonly] 

## Methods

### NewInstanceTypeResp

`func NewInstanceTypeResp(idKey InstanceTypeKey, ) *InstanceTypeResp`

NewInstanceTypeResp instantiates a new InstanceTypeResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeRespWithDefaults

`func NewInstanceTypeRespWithDefaults() *InstanceTypeResp`

NewInstanceTypeRespWithDefaults instantiates a new InstanceTypeResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *InstanceTypeResp) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *InstanceTypeResp) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *InstanceTypeResp) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *InstanceTypeResp) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetIdKey

`func (o *InstanceTypeResp) GetIdKey() InstanceTypeKey`

GetIdKey returns the IdKey field if non-nil, zero value otherwise.

### GetIdKeyOk

`func (o *InstanceTypeResp) GetIdKeyOk() (*InstanceTypeKey, bool)`

GetIdKeyOk returns a tuple with the IdKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdKey

`func (o *InstanceTypeResp) SetIdKey(v InstanceTypeKey)`

SetIdKey sets IdKey field to given value.


### GetInstanceTypeCode

`func (o *InstanceTypeResp) GetInstanceTypeCode() string`

GetInstanceTypeCode returns the InstanceTypeCode field if non-nil, zero value otherwise.

### GetInstanceTypeCodeOk

`func (o *InstanceTypeResp) GetInstanceTypeCodeOk() (*string, bool)`

GetInstanceTypeCodeOk returns a tuple with the InstanceTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeCode

`func (o *InstanceTypeResp) SetInstanceTypeCode(v string)`

SetInstanceTypeCode sets InstanceTypeCode field to given value.

### HasInstanceTypeCode

`func (o *InstanceTypeResp) HasInstanceTypeCode() bool`

HasInstanceTypeCode returns a boolean if a field has been set.

### GetInstanceTypeDetails

`func (o *InstanceTypeResp) GetInstanceTypeDetails() InstanceTypeDetails`

GetInstanceTypeDetails returns the InstanceTypeDetails field if non-nil, zero value otherwise.

### GetInstanceTypeDetailsOk

`func (o *InstanceTypeResp) GetInstanceTypeDetailsOk() (*InstanceTypeDetails, bool)`

GetInstanceTypeDetailsOk returns a tuple with the InstanceTypeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeDetails

`func (o *InstanceTypeResp) SetInstanceTypeDetails(v InstanceTypeDetails)`

SetInstanceTypeDetails sets InstanceTypeDetails field to given value.

### HasInstanceTypeDetails

`func (o *InstanceTypeResp) HasInstanceTypeDetails() bool`

HasInstanceTypeDetails returns a boolean if a field has been set.

### GetMemSizeGB

`func (o *InstanceTypeResp) GetMemSizeGB() float64`

GetMemSizeGB returns the MemSizeGB field if non-nil, zero value otherwise.

### GetMemSizeGBOk

`func (o *InstanceTypeResp) GetMemSizeGBOk() (*float64, bool)`

GetMemSizeGBOk returns a tuple with the MemSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemSizeGB

`func (o *InstanceTypeResp) SetMemSizeGB(v float64)`

SetMemSizeGB sets MemSizeGB field to given value.

### HasMemSizeGB

`func (o *InstanceTypeResp) HasMemSizeGB() bool`

HasMemSizeGB returns a boolean if a field has been set.

### GetNumCores

`func (o *InstanceTypeResp) GetNumCores() float64`

GetNumCores returns the NumCores field if non-nil, zero value otherwise.

### GetNumCoresOk

`func (o *InstanceTypeResp) GetNumCoresOk() (*float64, bool)`

GetNumCoresOk returns a tuple with the NumCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCores

`func (o *InstanceTypeResp) SetNumCores(v float64)`

SetNumCores sets NumCores field to given value.

### HasNumCores

`func (o *InstanceTypeResp) HasNumCores() bool`

HasNumCores returns a boolean if a field has been set.

### GetProviderCode

`func (o *InstanceTypeResp) GetProviderCode() string`

GetProviderCode returns the ProviderCode field if non-nil, zero value otherwise.

### GetProviderCodeOk

`func (o *InstanceTypeResp) GetProviderCodeOk() (*string, bool)`

GetProviderCodeOk returns a tuple with the ProviderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCode

`func (o *InstanceTypeResp) SetProviderCode(v string)`

SetProviderCode sets ProviderCode field to given value.

### HasProviderCode

`func (o *InstanceTypeResp) HasProviderCode() bool`

HasProviderCode returns a boolean if a field has been set.

### GetProviderUuid

`func (o *InstanceTypeResp) GetProviderUuid() string`

GetProviderUuid returns the ProviderUuid field if non-nil, zero value otherwise.

### GetProviderUuidOk

`func (o *InstanceTypeResp) GetProviderUuidOk() (*string, bool)`

GetProviderUuidOk returns a tuple with the ProviderUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUuid

`func (o *InstanceTypeResp) SetProviderUuid(v string)`

SetProviderUuid sets ProviderUuid field to given value.

### HasProviderUuid

`func (o *InstanceTypeResp) HasProviderUuid() bool`

HasProviderUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


