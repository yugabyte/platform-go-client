# AvailabilityZoneNodeSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CgroupSize** | Pointer to **int32** | Amount of memory in MB to limit the postgres process using the ysql cgroup. The value should be greater than 0. When set to 0 it results in no cgroup limits. For a read replica cluster, setting this value to null or -1 would inherit this value from the primary cluster. Applicable only for nodes running as Linux VMs on AWS/GCP/Azure Cloud Provider. Only used internally by YBM. | [optional] 
**Tserver** | Pointer to [**PerProcessNodeSpec**](PerProcessNodeSpec.md) |  | [optional] 
**Master** | Pointer to [**PerProcessNodeSpec**](PerProcessNodeSpec.md) |  | [optional] 

## Methods

### NewAvailabilityZoneNodeSpecAllOf

`func NewAvailabilityZoneNodeSpecAllOf() *AvailabilityZoneNodeSpecAllOf`

NewAvailabilityZoneNodeSpecAllOf instantiates a new AvailabilityZoneNodeSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityZoneNodeSpecAllOfWithDefaults

`func NewAvailabilityZoneNodeSpecAllOfWithDefaults() *AvailabilityZoneNodeSpecAllOf`

NewAvailabilityZoneNodeSpecAllOfWithDefaults instantiates a new AvailabilityZoneNodeSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCgroupSize

`func (o *AvailabilityZoneNodeSpecAllOf) GetCgroupSize() int32`

GetCgroupSize returns the CgroupSize field if non-nil, zero value otherwise.

### GetCgroupSizeOk

`func (o *AvailabilityZoneNodeSpecAllOf) GetCgroupSizeOk() (*int32, bool)`

GetCgroupSizeOk returns a tuple with the CgroupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCgroupSize

`func (o *AvailabilityZoneNodeSpecAllOf) SetCgroupSize(v int32)`

SetCgroupSize sets CgroupSize field to given value.

### HasCgroupSize

`func (o *AvailabilityZoneNodeSpecAllOf) HasCgroupSize() bool`

HasCgroupSize returns a boolean if a field has been set.

### GetTserver

`func (o *AvailabilityZoneNodeSpecAllOf) GetTserver() PerProcessNodeSpec`

GetTserver returns the Tserver field if non-nil, zero value otherwise.

### GetTserverOk

`func (o *AvailabilityZoneNodeSpecAllOf) GetTserverOk() (*PerProcessNodeSpec, bool)`

GetTserverOk returns a tuple with the Tserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserver

`func (o *AvailabilityZoneNodeSpecAllOf) SetTserver(v PerProcessNodeSpec)`

SetTserver sets Tserver field to given value.

### HasTserver

`func (o *AvailabilityZoneNodeSpecAllOf) HasTserver() bool`

HasTserver returns a boolean if a field has been set.

### GetMaster

`func (o *AvailabilityZoneNodeSpecAllOf) GetMaster() PerProcessNodeSpec`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *AvailabilityZoneNodeSpecAllOf) GetMasterOk() (*PerProcessNodeSpec, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *AvailabilityZoneNodeSpecAllOf) SetMaster(v PerProcessNodeSpec)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *AvailabilityZoneNodeSpecAllOf) HasMaster() bool`

HasMaster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


