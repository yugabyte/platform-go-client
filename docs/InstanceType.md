# InstanceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | True if the instance is active | [optional] [readonly] 
**IdKey** | [**InstanceTypeKey**](InstanceTypeKey.md) |  | 
**InstanceTypeCode** | Pointer to **string** | Instance type code | [optional] [readonly] 
**InstanceTypeDetails** | Pointer to [**InstanceTypeDetails**](InstanceTypeDetails.md) |  | [optional] 
**MemSizeGB** | Pointer to **float64** | The instance&#39;s memory size, in gigabytes | [optional] 
**NumCores** | Pointer to **float64** | The instance&#39;s number of CPU cores | [optional] 

## Methods

### NewInstanceType

`func NewInstanceType(idKey InstanceTypeKey, ) *InstanceType`

NewInstanceType instantiates a new InstanceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeWithDefaults

`func NewInstanceTypeWithDefaults() *InstanceType`

NewInstanceTypeWithDefaults instantiates a new InstanceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *InstanceType) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *InstanceType) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *InstanceType) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *InstanceType) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetIdKey

`func (o *InstanceType) GetIdKey() InstanceTypeKey`

GetIdKey returns the IdKey field if non-nil, zero value otherwise.

### GetIdKeyOk

`func (o *InstanceType) GetIdKeyOk() (*InstanceTypeKey, bool)`

GetIdKeyOk returns a tuple with the IdKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdKey

`func (o *InstanceType) SetIdKey(v InstanceTypeKey)`

SetIdKey sets IdKey field to given value.


### GetInstanceTypeCode

`func (o *InstanceType) GetInstanceTypeCode() string`

GetInstanceTypeCode returns the InstanceTypeCode field if non-nil, zero value otherwise.

### GetInstanceTypeCodeOk

`func (o *InstanceType) GetInstanceTypeCodeOk() (*string, bool)`

GetInstanceTypeCodeOk returns a tuple with the InstanceTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeCode

`func (o *InstanceType) SetInstanceTypeCode(v string)`

SetInstanceTypeCode sets InstanceTypeCode field to given value.

### HasInstanceTypeCode

`func (o *InstanceType) HasInstanceTypeCode() bool`

HasInstanceTypeCode returns a boolean if a field has been set.

### GetInstanceTypeDetails

`func (o *InstanceType) GetInstanceTypeDetails() InstanceTypeDetails`

GetInstanceTypeDetails returns the InstanceTypeDetails field if non-nil, zero value otherwise.

### GetInstanceTypeDetailsOk

`func (o *InstanceType) GetInstanceTypeDetailsOk() (*InstanceTypeDetails, bool)`

GetInstanceTypeDetailsOk returns a tuple with the InstanceTypeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeDetails

`func (o *InstanceType) SetInstanceTypeDetails(v InstanceTypeDetails)`

SetInstanceTypeDetails sets InstanceTypeDetails field to given value.

### HasInstanceTypeDetails

`func (o *InstanceType) HasInstanceTypeDetails() bool`

HasInstanceTypeDetails returns a boolean if a field has been set.

### GetMemSizeGB

`func (o *InstanceType) GetMemSizeGB() float64`

GetMemSizeGB returns the MemSizeGB field if non-nil, zero value otherwise.

### GetMemSizeGBOk

`func (o *InstanceType) GetMemSizeGBOk() (*float64, bool)`

GetMemSizeGBOk returns a tuple with the MemSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemSizeGB

`func (o *InstanceType) SetMemSizeGB(v float64)`

SetMemSizeGB sets MemSizeGB field to given value.

### HasMemSizeGB

`func (o *InstanceType) HasMemSizeGB() bool`

HasMemSizeGB returns a boolean if a field has been set.

### GetNumCores

`func (o *InstanceType) GetNumCores() float64`

GetNumCores returns the NumCores field if non-nil, zero value otherwise.

### GetNumCoresOk

`func (o *InstanceType) GetNumCoresOk() (*float64, bool)`

GetNumCoresOk returns a tuple with the NumCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCores

`func (o *InstanceType) SetNumCores(v float64)`

SetNumCores sets NumCores field to given value.

### HasNumCores

`func (o *InstanceType) HasNumCores() bool`

HasNumCores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


