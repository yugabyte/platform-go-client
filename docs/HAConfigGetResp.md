# HAConfigGetResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptAnyCertificate** | Pointer to **bool** | HA accepts any certificate | [optional] [readonly] 
**ClusterKey** | Pointer to **string** | HA config cluster key | [optional] [readonly] 
**GlobalState** | Pointer to **string** | HA config global state | [optional] [readonly] 
**Instances** | Pointer to [**[]PlatformInstance**](PlatformInstance.md) | HA config platform instances | [optional] 
**LastFailover** | Pointer to **time.Time** | HA last failover | [optional] [readonly] 
**Uuid** | Pointer to **string** | HA config UUID | [optional] 

## Methods

### NewHAConfigGetResp

`func NewHAConfigGetResp() *HAConfigGetResp`

NewHAConfigGetResp instantiates a new HAConfigGetResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHAConfigGetRespWithDefaults

`func NewHAConfigGetRespWithDefaults() *HAConfigGetResp`

NewHAConfigGetRespWithDefaults instantiates a new HAConfigGetResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptAnyCertificate

`func (o *HAConfigGetResp) GetAcceptAnyCertificate() bool`

GetAcceptAnyCertificate returns the AcceptAnyCertificate field if non-nil, zero value otherwise.

### GetAcceptAnyCertificateOk

`func (o *HAConfigGetResp) GetAcceptAnyCertificateOk() (*bool, bool)`

GetAcceptAnyCertificateOk returns a tuple with the AcceptAnyCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptAnyCertificate

`func (o *HAConfigGetResp) SetAcceptAnyCertificate(v bool)`

SetAcceptAnyCertificate sets AcceptAnyCertificate field to given value.

### HasAcceptAnyCertificate

`func (o *HAConfigGetResp) HasAcceptAnyCertificate() bool`

HasAcceptAnyCertificate returns a boolean if a field has been set.

### GetClusterKey

`func (o *HAConfigGetResp) GetClusterKey() string`

GetClusterKey returns the ClusterKey field if non-nil, zero value otherwise.

### GetClusterKeyOk

`func (o *HAConfigGetResp) GetClusterKeyOk() (*string, bool)`

GetClusterKeyOk returns a tuple with the ClusterKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterKey

`func (o *HAConfigGetResp) SetClusterKey(v string)`

SetClusterKey sets ClusterKey field to given value.

### HasClusterKey

`func (o *HAConfigGetResp) HasClusterKey() bool`

HasClusterKey returns a boolean if a field has been set.

### GetGlobalState

`func (o *HAConfigGetResp) GetGlobalState() string`

GetGlobalState returns the GlobalState field if non-nil, zero value otherwise.

### GetGlobalStateOk

`func (o *HAConfigGetResp) GetGlobalStateOk() (*string, bool)`

GetGlobalStateOk returns a tuple with the GlobalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalState

`func (o *HAConfigGetResp) SetGlobalState(v string)`

SetGlobalState sets GlobalState field to given value.

### HasGlobalState

`func (o *HAConfigGetResp) HasGlobalState() bool`

HasGlobalState returns a boolean if a field has been set.

### GetInstances

`func (o *HAConfigGetResp) GetInstances() []PlatformInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *HAConfigGetResp) GetInstancesOk() (*[]PlatformInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *HAConfigGetResp) SetInstances(v []PlatformInstance)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *HAConfigGetResp) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetLastFailover

`func (o *HAConfigGetResp) GetLastFailover() time.Time`

GetLastFailover returns the LastFailover field if non-nil, zero value otherwise.

### GetLastFailoverOk

`func (o *HAConfigGetResp) GetLastFailoverOk() (*time.Time, bool)`

GetLastFailoverOk returns a tuple with the LastFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailover

`func (o *HAConfigGetResp) SetLastFailover(v time.Time)`

SetLastFailover sets LastFailover field to given value.

### HasLastFailover

`func (o *HAConfigGetResp) HasLastFailover() bool`

HasLastFailover returns a boolean if a field has been set.

### GetUuid

`func (o *HAConfigGetResp) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HAConfigGetResp) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HAConfigGetResp) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HAConfigGetResp) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


