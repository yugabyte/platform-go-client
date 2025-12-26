# HighAvailabilityConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptAnyCertificate** | Pointer to **bool** |  | [optional] 
**ClusterKey** | Pointer to **string** |  | [optional] 
**Instances** | [**[]PlatformInstance**](PlatformInstance.md) |  | 
**LastFailover** | Pointer to **time.Time** | HA last failover | [optional] 
**Uuid** | **string** |  | 

## Methods

### NewHighAvailabilityConfig

`func NewHighAvailabilityConfig(instances []PlatformInstance, uuid string, ) *HighAvailabilityConfig`

NewHighAvailabilityConfig instantiates a new HighAvailabilityConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighAvailabilityConfigWithDefaults

`func NewHighAvailabilityConfigWithDefaults() *HighAvailabilityConfig`

NewHighAvailabilityConfigWithDefaults instantiates a new HighAvailabilityConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptAnyCertificate

`func (o *HighAvailabilityConfig) GetAcceptAnyCertificate() bool`

GetAcceptAnyCertificate returns the AcceptAnyCertificate field if non-nil, zero value otherwise.

### GetAcceptAnyCertificateOk

`func (o *HighAvailabilityConfig) GetAcceptAnyCertificateOk() (*bool, bool)`

GetAcceptAnyCertificateOk returns a tuple with the AcceptAnyCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptAnyCertificate

`func (o *HighAvailabilityConfig) SetAcceptAnyCertificate(v bool)`

SetAcceptAnyCertificate sets AcceptAnyCertificate field to given value.

### HasAcceptAnyCertificate

`func (o *HighAvailabilityConfig) HasAcceptAnyCertificate() bool`

HasAcceptAnyCertificate returns a boolean if a field has been set.

### GetClusterKey

`func (o *HighAvailabilityConfig) GetClusterKey() string`

GetClusterKey returns the ClusterKey field if non-nil, zero value otherwise.

### GetClusterKeyOk

`func (o *HighAvailabilityConfig) GetClusterKeyOk() (*string, bool)`

GetClusterKeyOk returns a tuple with the ClusterKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterKey

`func (o *HighAvailabilityConfig) SetClusterKey(v string)`

SetClusterKey sets ClusterKey field to given value.

### HasClusterKey

`func (o *HighAvailabilityConfig) HasClusterKey() bool`

HasClusterKey returns a boolean if a field has been set.

### GetInstances

`func (o *HighAvailabilityConfig) GetInstances() []PlatformInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *HighAvailabilityConfig) GetInstancesOk() (*[]PlatformInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *HighAvailabilityConfig) SetInstances(v []PlatformInstance)`

SetInstances sets Instances field to given value.


### GetLastFailover

`func (o *HighAvailabilityConfig) GetLastFailover() time.Time`

GetLastFailover returns the LastFailover field if non-nil, zero value otherwise.

### GetLastFailoverOk

`func (o *HighAvailabilityConfig) GetLastFailoverOk() (*time.Time, bool)`

GetLastFailoverOk returns a tuple with the LastFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailover

`func (o *HighAvailabilityConfig) SetLastFailover(v time.Time)`

SetLastFailover sets LastFailover field to given value.

### HasLastFailover

`func (o *HighAvailabilityConfig) HasLastFailover() bool`

HasLastFailover returns a boolean if a field has been set.

### GetUuid

`func (o *HighAvailabilityConfig) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HighAvailabilityConfig) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HighAvailabilityConfig) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


