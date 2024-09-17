# UniverseResourceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzList** | **[]string** | Azs | 
**EbsPricePerHour** | Pointer to **float64** | EBS price per hour | [optional] 
**Gp3FreePiops** | Pointer to **int32** | gp3 free piops | [optional] 
**Gp3FreeThroughput** | Pointer to **int32** | gp3 free throughput | [optional] 
**MemSizeGb** | Pointer to **float64** | Memory GB | [optional] 
**NumCores** | Pointer to **float64** | Numbers of cores | [optional] 
**NumNodes** | Pointer to **int32** | Numbers of node | [optional] 
**PricePerHour** | Pointer to **float64** | Price per hour | [optional] 
**PricingKnown** | Pointer to **bool** | Known pricing info | [optional] 
**VolumeCount** | Pointer to **int32** | Volume count | [optional] 
**VolumeSizeGb** | Pointer to **int32** | Volume in GB | [optional] 

## Methods

### NewUniverseResourceDetails

`func NewUniverseResourceDetails(azList []string, ) *UniverseResourceDetails`

NewUniverseResourceDetails instantiates a new UniverseResourceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseResourceDetailsWithDefaults

`func NewUniverseResourceDetailsWithDefaults() *UniverseResourceDetails`

NewUniverseResourceDetailsWithDefaults instantiates a new UniverseResourceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzList

`func (o *UniverseResourceDetails) GetAzList() []string`

GetAzList returns the AzList field if non-nil, zero value otherwise.

### GetAzListOk

`func (o *UniverseResourceDetails) GetAzListOk() (*[]string, bool)`

GetAzListOk returns a tuple with the AzList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzList

`func (o *UniverseResourceDetails) SetAzList(v []string)`

SetAzList sets AzList field to given value.


### GetEbsPricePerHour

`func (o *UniverseResourceDetails) GetEbsPricePerHour() float64`

GetEbsPricePerHour returns the EbsPricePerHour field if non-nil, zero value otherwise.

### GetEbsPricePerHourOk

`func (o *UniverseResourceDetails) GetEbsPricePerHourOk() (*float64, bool)`

GetEbsPricePerHourOk returns a tuple with the EbsPricePerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsPricePerHour

`func (o *UniverseResourceDetails) SetEbsPricePerHour(v float64)`

SetEbsPricePerHour sets EbsPricePerHour field to given value.

### HasEbsPricePerHour

`func (o *UniverseResourceDetails) HasEbsPricePerHour() bool`

HasEbsPricePerHour returns a boolean if a field has been set.

### GetGp3FreePiops

`func (o *UniverseResourceDetails) GetGp3FreePiops() int32`

GetGp3FreePiops returns the Gp3FreePiops field if non-nil, zero value otherwise.

### GetGp3FreePiopsOk

`func (o *UniverseResourceDetails) GetGp3FreePiopsOk() (*int32, bool)`

GetGp3FreePiopsOk returns a tuple with the Gp3FreePiops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGp3FreePiops

`func (o *UniverseResourceDetails) SetGp3FreePiops(v int32)`

SetGp3FreePiops sets Gp3FreePiops field to given value.

### HasGp3FreePiops

`func (o *UniverseResourceDetails) HasGp3FreePiops() bool`

HasGp3FreePiops returns a boolean if a field has been set.

### GetGp3FreeThroughput

`func (o *UniverseResourceDetails) GetGp3FreeThroughput() int32`

GetGp3FreeThroughput returns the Gp3FreeThroughput field if non-nil, zero value otherwise.

### GetGp3FreeThroughputOk

`func (o *UniverseResourceDetails) GetGp3FreeThroughputOk() (*int32, bool)`

GetGp3FreeThroughputOk returns a tuple with the Gp3FreeThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGp3FreeThroughput

`func (o *UniverseResourceDetails) SetGp3FreeThroughput(v int32)`

SetGp3FreeThroughput sets Gp3FreeThroughput field to given value.

### HasGp3FreeThroughput

`func (o *UniverseResourceDetails) HasGp3FreeThroughput() bool`

HasGp3FreeThroughput returns a boolean if a field has been set.

### GetMemSizeGb

`func (o *UniverseResourceDetails) GetMemSizeGb() float64`

GetMemSizeGb returns the MemSizeGb field if non-nil, zero value otherwise.

### GetMemSizeGbOk

`func (o *UniverseResourceDetails) GetMemSizeGbOk() (*float64, bool)`

GetMemSizeGbOk returns a tuple with the MemSizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemSizeGb

`func (o *UniverseResourceDetails) SetMemSizeGb(v float64)`

SetMemSizeGb sets MemSizeGb field to given value.

### HasMemSizeGb

`func (o *UniverseResourceDetails) HasMemSizeGb() bool`

HasMemSizeGb returns a boolean if a field has been set.

### GetNumCores

`func (o *UniverseResourceDetails) GetNumCores() float64`

GetNumCores returns the NumCores field if non-nil, zero value otherwise.

### GetNumCoresOk

`func (o *UniverseResourceDetails) GetNumCoresOk() (*float64, bool)`

GetNumCoresOk returns a tuple with the NumCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCores

`func (o *UniverseResourceDetails) SetNumCores(v float64)`

SetNumCores sets NumCores field to given value.

### HasNumCores

`func (o *UniverseResourceDetails) HasNumCores() bool`

HasNumCores returns a boolean if a field has been set.

### GetNumNodes

`func (o *UniverseResourceDetails) GetNumNodes() int32`

GetNumNodes returns the NumNodes field if non-nil, zero value otherwise.

### GetNumNodesOk

`func (o *UniverseResourceDetails) GetNumNodesOk() (*int32, bool)`

GetNumNodesOk returns a tuple with the NumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNodes

`func (o *UniverseResourceDetails) SetNumNodes(v int32)`

SetNumNodes sets NumNodes field to given value.

### HasNumNodes

`func (o *UniverseResourceDetails) HasNumNodes() bool`

HasNumNodes returns a boolean if a field has been set.

### GetPricePerHour

`func (o *UniverseResourceDetails) GetPricePerHour() float64`

GetPricePerHour returns the PricePerHour field if non-nil, zero value otherwise.

### GetPricePerHourOk

`func (o *UniverseResourceDetails) GetPricePerHourOk() (*float64, bool)`

GetPricePerHourOk returns a tuple with the PricePerHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerHour

`func (o *UniverseResourceDetails) SetPricePerHour(v float64)`

SetPricePerHour sets PricePerHour field to given value.

### HasPricePerHour

`func (o *UniverseResourceDetails) HasPricePerHour() bool`

HasPricePerHour returns a boolean if a field has been set.

### GetPricingKnown

`func (o *UniverseResourceDetails) GetPricingKnown() bool`

GetPricingKnown returns the PricingKnown field if non-nil, zero value otherwise.

### GetPricingKnownOk

`func (o *UniverseResourceDetails) GetPricingKnownOk() (*bool, bool)`

GetPricingKnownOk returns a tuple with the PricingKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingKnown

`func (o *UniverseResourceDetails) SetPricingKnown(v bool)`

SetPricingKnown sets PricingKnown field to given value.

### HasPricingKnown

`func (o *UniverseResourceDetails) HasPricingKnown() bool`

HasPricingKnown returns a boolean if a field has been set.

### GetVolumeCount

`func (o *UniverseResourceDetails) GetVolumeCount() int32`

GetVolumeCount returns the VolumeCount field if non-nil, zero value otherwise.

### GetVolumeCountOk

`func (o *UniverseResourceDetails) GetVolumeCountOk() (*int32, bool)`

GetVolumeCountOk returns a tuple with the VolumeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCount

`func (o *UniverseResourceDetails) SetVolumeCount(v int32)`

SetVolumeCount sets VolumeCount field to given value.

### HasVolumeCount

`func (o *UniverseResourceDetails) HasVolumeCount() bool`

HasVolumeCount returns a boolean if a field has been set.

### GetVolumeSizeGb

`func (o *UniverseResourceDetails) GetVolumeSizeGb() int32`

GetVolumeSizeGb returns the VolumeSizeGb field if non-nil, zero value otherwise.

### GetVolumeSizeGbOk

`func (o *UniverseResourceDetails) GetVolumeSizeGbOk() (*int32, bool)`

GetVolumeSizeGbOk returns a tuple with the VolumeSizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSizeGb

`func (o *UniverseResourceDetails) SetVolumeSizeGb(v int32)`

SetVolumeSizeGb sets VolumeSizeGb field to given value.

### HasVolumeSizeGb

`func (o *UniverseResourceDetails) HasVolumeSizeGb() bool`

HasVolumeSizeGb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


