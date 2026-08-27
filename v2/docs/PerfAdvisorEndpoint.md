# PerfAdvisorEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to [**PerfAdvisorEndpointInfo**](PerfAdvisorEndpointInfo.md) |  | [optional] 
**Spec** | Pointer to [**PerfAdvisorEndpointSpec**](PerfAdvisorEndpointSpec.md) |  | [optional] 

## Methods

### NewPerfAdvisorEndpoint

`func NewPerfAdvisorEndpoint() *PerfAdvisorEndpoint`

NewPerfAdvisorEndpoint instantiates a new PerfAdvisorEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerfAdvisorEndpointWithDefaults

`func NewPerfAdvisorEndpointWithDefaults() *PerfAdvisorEndpoint`

NewPerfAdvisorEndpointWithDefaults instantiates a new PerfAdvisorEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *PerfAdvisorEndpoint) GetInfo() PerfAdvisorEndpointInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PerfAdvisorEndpoint) GetInfoOk() (*PerfAdvisorEndpointInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PerfAdvisorEndpoint) SetInfo(v PerfAdvisorEndpointInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PerfAdvisorEndpoint) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetSpec

`func (o *PerfAdvisorEndpoint) GetSpec() PerfAdvisorEndpointSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *PerfAdvisorEndpoint) GetSpecOk() (*PerfAdvisorEndpointSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *PerfAdvisorEndpoint) SetSpec(v PerfAdvisorEndpointSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *PerfAdvisorEndpoint) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


