# GcpHostInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to **string** | ID of the project that the instance belongs to | [optional] [readonly] 
**Network** | Pointer to **string** | Name of the VPC network of the instance&#39;s first network interface | [optional] [readonly] 
**SharedVpcProject** | Pointer to **string** | Number of the project that owns that network: the host project for a Shared VPC, otherwise the instance&#39;s own project | [optional] [readonly] 
**Zone** | Pointer to **string** | GCP zone | [optional] [readonly] 
**Region** | Pointer to **string** | GCP region | [optional] [readonly] 

## Methods

### NewGcpHostInfo

`func NewGcpHostInfo() *GcpHostInfo`

NewGcpHostInfo instantiates a new GcpHostInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpHostInfoWithDefaults

`func NewGcpHostInfoWithDefaults() *GcpHostInfo`

NewGcpHostInfoWithDefaults instantiates a new GcpHostInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *GcpHostInfo) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GcpHostInfo) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GcpHostInfo) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *GcpHostInfo) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetNetwork

`func (o *GcpHostInfo) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GcpHostInfo) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GcpHostInfo) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GcpHostInfo) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSharedVpcProject

`func (o *GcpHostInfo) GetSharedVpcProject() string`

GetSharedVpcProject returns the SharedVpcProject field if non-nil, zero value otherwise.

### GetSharedVpcProjectOk

`func (o *GcpHostInfo) GetSharedVpcProjectOk() (*string, bool)`

GetSharedVpcProjectOk returns a tuple with the SharedVpcProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedVpcProject

`func (o *GcpHostInfo) SetSharedVpcProject(v string)`

SetSharedVpcProject sets SharedVpcProject field to given value.

### HasSharedVpcProject

`func (o *GcpHostInfo) HasSharedVpcProject() bool`

HasSharedVpcProject returns a boolean if a field has been set.

### GetZone

`func (o *GcpHostInfo) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GcpHostInfo) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GcpHostInfo) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GcpHostInfo) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetRegion

`func (o *GcpHostInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GcpHostInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GcpHostInfo) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GcpHostInfo) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


