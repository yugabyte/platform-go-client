# MasterLBStateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstTimeToBalanceSecs** | Pointer to **int64** | YbaApi Internal Estimate of time for which master tablet load balancer will be active | [optional] 
**IsEnabled** | **bool** | YbaApi Internal Whether master tablet load balancer is enabled | 
**IsIdle** | Pointer to **bool** | YbaApi Internal Whether master tablet load balancer is inactive | [optional] 

## Methods

### NewMasterLBStateResponse

`func NewMasterLBStateResponse(isEnabled bool, ) *MasterLBStateResponse`

NewMasterLBStateResponse instantiates a new MasterLBStateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMasterLBStateResponseWithDefaults

`func NewMasterLBStateResponseWithDefaults() *MasterLBStateResponse`

NewMasterLBStateResponseWithDefaults instantiates a new MasterLBStateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstTimeToBalanceSecs

`func (o *MasterLBStateResponse) GetEstTimeToBalanceSecs() int64`

GetEstTimeToBalanceSecs returns the EstTimeToBalanceSecs field if non-nil, zero value otherwise.

### GetEstTimeToBalanceSecsOk

`func (o *MasterLBStateResponse) GetEstTimeToBalanceSecsOk() (*int64, bool)`

GetEstTimeToBalanceSecsOk returns a tuple with the EstTimeToBalanceSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstTimeToBalanceSecs

`func (o *MasterLBStateResponse) SetEstTimeToBalanceSecs(v int64)`

SetEstTimeToBalanceSecs sets EstTimeToBalanceSecs field to given value.

### HasEstTimeToBalanceSecs

`func (o *MasterLBStateResponse) HasEstTimeToBalanceSecs() bool`

HasEstTimeToBalanceSecs returns a boolean if a field has been set.

### GetIsEnabled

`func (o *MasterLBStateResponse) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MasterLBStateResponse) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *MasterLBStateResponse) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetIsIdle

`func (o *MasterLBStateResponse) GetIsIdle() bool`

GetIsIdle returns the IsIdle field if non-nil, zero value otherwise.

### GetIsIdleOk

`func (o *MasterLBStateResponse) GetIsIdleOk() (*bool, bool)`

GetIsIdleOk returns a tuple with the IsIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIdle

`func (o *MasterLBStateResponse) SetIsIdle(v bool)`

SetIsIdle sets IsIdle field to given value.

### HasIsIdle

`func (o *MasterLBStateResponse) HasIsIdle() bool`

HasIsIdle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


