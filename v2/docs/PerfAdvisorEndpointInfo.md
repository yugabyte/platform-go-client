# PerfAdvisorEndpointInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | UUID of the endpoint. Also the id the configuration carries on every collector it is pushed to, so it is stable and meaningful on both sides.  | [readonly] 
**CustomerUuid** | Pointer to **string** | UUID of the customer that owns the endpoint. | [optional] [readonly] 
**UniverseUuids** | Pointer to **[]string** | Universes currently registered in online mode against this endpoint. | [optional] [readonly] 
**CreateTime** | Pointer to **time.Time** | Creation time. | [optional] [readonly] 
**UpdateTime** | Pointer to **time.Time** | Time of the last edit. | [optional] [readonly] 

## Methods

### NewPerfAdvisorEndpointInfo

`func NewPerfAdvisorEndpointInfo(uuid string, ) *PerfAdvisorEndpointInfo`

NewPerfAdvisorEndpointInfo instantiates a new PerfAdvisorEndpointInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerfAdvisorEndpointInfoWithDefaults

`func NewPerfAdvisorEndpointInfoWithDefaults() *PerfAdvisorEndpointInfo`

NewPerfAdvisorEndpointInfoWithDefaults instantiates a new PerfAdvisorEndpointInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PerfAdvisorEndpointInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PerfAdvisorEndpointInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PerfAdvisorEndpointInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetCustomerUuid

`func (o *PerfAdvisorEndpointInfo) GetCustomerUuid() string`

GetCustomerUuid returns the CustomerUuid field if non-nil, zero value otherwise.

### GetCustomerUuidOk

`func (o *PerfAdvisorEndpointInfo) GetCustomerUuidOk() (*string, bool)`

GetCustomerUuidOk returns a tuple with the CustomerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUuid

`func (o *PerfAdvisorEndpointInfo) SetCustomerUuid(v string)`

SetCustomerUuid sets CustomerUuid field to given value.

### HasCustomerUuid

`func (o *PerfAdvisorEndpointInfo) HasCustomerUuid() bool`

HasCustomerUuid returns a boolean if a field has been set.

### GetUniverseUuids

`func (o *PerfAdvisorEndpointInfo) GetUniverseUuids() []string`

GetUniverseUuids returns the UniverseUuids field if non-nil, zero value otherwise.

### GetUniverseUuidsOk

`func (o *PerfAdvisorEndpointInfo) GetUniverseUuidsOk() (*[]string, bool)`

GetUniverseUuidsOk returns a tuple with the UniverseUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseUuids

`func (o *PerfAdvisorEndpointInfo) SetUniverseUuids(v []string)`

SetUniverseUuids sets UniverseUuids field to given value.

### HasUniverseUuids

`func (o *PerfAdvisorEndpointInfo) HasUniverseUuids() bool`

HasUniverseUuids returns a boolean if a field has been set.

### GetCreateTime

`func (o *PerfAdvisorEndpointInfo) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *PerfAdvisorEndpointInfo) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *PerfAdvisorEndpointInfo) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *PerfAdvisorEndpointInfo) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *PerfAdvisorEndpointInfo) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *PerfAdvisorEndpointInfo) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *PerfAdvisorEndpointInfo) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *PerfAdvisorEndpointInfo) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


