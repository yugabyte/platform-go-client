# PerfAdvisorEndpointSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the endpoint. Unique per customer. | 
**Type** | [**PerfAdvisorEndpointType**](PerfAdvisorEndpointType.md) |  | 
**MetricsEndpoint** | **string** | URL the collector sends metrics to. | 
**MetricsType** | [**PerfAdvisorEndpointMetricsType**](PerfAdvisorEndpointMetricsType.md) |  | 
**MetricsAuth** | Pointer to [**PerfAdvisorEndpointAuth**](PerfAdvisorEndpointAuth.md) |  | [optional] 
**CollectionEndpoint** | **string** | URL of the destination&#39;s Collection API, where everything other than metrics goes. | 
**CollectionAuth** | Pointer to [**PerfAdvisorEndpointAuth**](PerfAdvisorEndpointAuth.md) |  | [optional] 
**YbmAccountId** | Pointer to **string** | YugabyteDB Managed account ID. Sent as the YBM-Account-ID header on both endpoints, which is how a BYOC ingest gateway identifies the sender. Leave unset for a plain Perf Advisor.  | [optional] 
**YbmProjectId** | Pointer to **string** | YugabyteDB Managed project ID. Sent as the YBM-Project-ID header on both endpoints.  | [optional] 

## Methods

### NewPerfAdvisorEndpointSpec

`func NewPerfAdvisorEndpointSpec(name string, type_ PerfAdvisorEndpointType, metricsEndpoint string, metricsType PerfAdvisorEndpointMetricsType, collectionEndpoint string, ) *PerfAdvisorEndpointSpec`

NewPerfAdvisorEndpointSpec instantiates a new PerfAdvisorEndpointSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerfAdvisorEndpointSpecWithDefaults

`func NewPerfAdvisorEndpointSpecWithDefaults() *PerfAdvisorEndpointSpec`

NewPerfAdvisorEndpointSpecWithDefaults instantiates a new PerfAdvisorEndpointSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PerfAdvisorEndpointSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PerfAdvisorEndpointSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PerfAdvisorEndpointSpec) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *PerfAdvisorEndpointSpec) GetType() PerfAdvisorEndpointType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PerfAdvisorEndpointSpec) GetTypeOk() (*PerfAdvisorEndpointType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PerfAdvisorEndpointSpec) SetType(v PerfAdvisorEndpointType)`

SetType sets Type field to given value.


### GetMetricsEndpoint

`func (o *PerfAdvisorEndpointSpec) GetMetricsEndpoint() string`

GetMetricsEndpoint returns the MetricsEndpoint field if non-nil, zero value otherwise.

### GetMetricsEndpointOk

`func (o *PerfAdvisorEndpointSpec) GetMetricsEndpointOk() (*string, bool)`

GetMetricsEndpointOk returns a tuple with the MetricsEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsEndpoint

`func (o *PerfAdvisorEndpointSpec) SetMetricsEndpoint(v string)`

SetMetricsEndpoint sets MetricsEndpoint field to given value.


### GetMetricsType

`func (o *PerfAdvisorEndpointSpec) GetMetricsType() PerfAdvisorEndpointMetricsType`

GetMetricsType returns the MetricsType field if non-nil, zero value otherwise.

### GetMetricsTypeOk

`func (o *PerfAdvisorEndpointSpec) GetMetricsTypeOk() (*PerfAdvisorEndpointMetricsType, bool)`

GetMetricsTypeOk returns a tuple with the MetricsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsType

`func (o *PerfAdvisorEndpointSpec) SetMetricsType(v PerfAdvisorEndpointMetricsType)`

SetMetricsType sets MetricsType field to given value.


### GetMetricsAuth

`func (o *PerfAdvisorEndpointSpec) GetMetricsAuth() PerfAdvisorEndpointAuth`

GetMetricsAuth returns the MetricsAuth field if non-nil, zero value otherwise.

### GetMetricsAuthOk

`func (o *PerfAdvisorEndpointSpec) GetMetricsAuthOk() (*PerfAdvisorEndpointAuth, bool)`

GetMetricsAuthOk returns a tuple with the MetricsAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsAuth

`func (o *PerfAdvisorEndpointSpec) SetMetricsAuth(v PerfAdvisorEndpointAuth)`

SetMetricsAuth sets MetricsAuth field to given value.

### HasMetricsAuth

`func (o *PerfAdvisorEndpointSpec) HasMetricsAuth() bool`

HasMetricsAuth returns a boolean if a field has been set.

### GetCollectionEndpoint

`func (o *PerfAdvisorEndpointSpec) GetCollectionEndpoint() string`

GetCollectionEndpoint returns the CollectionEndpoint field if non-nil, zero value otherwise.

### GetCollectionEndpointOk

`func (o *PerfAdvisorEndpointSpec) GetCollectionEndpointOk() (*string, bool)`

GetCollectionEndpointOk returns a tuple with the CollectionEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionEndpoint

`func (o *PerfAdvisorEndpointSpec) SetCollectionEndpoint(v string)`

SetCollectionEndpoint sets CollectionEndpoint field to given value.


### GetCollectionAuth

`func (o *PerfAdvisorEndpointSpec) GetCollectionAuth() PerfAdvisorEndpointAuth`

GetCollectionAuth returns the CollectionAuth field if non-nil, zero value otherwise.

### GetCollectionAuthOk

`func (o *PerfAdvisorEndpointSpec) GetCollectionAuthOk() (*PerfAdvisorEndpointAuth, bool)`

GetCollectionAuthOk returns a tuple with the CollectionAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAuth

`func (o *PerfAdvisorEndpointSpec) SetCollectionAuth(v PerfAdvisorEndpointAuth)`

SetCollectionAuth sets CollectionAuth field to given value.

### HasCollectionAuth

`func (o *PerfAdvisorEndpointSpec) HasCollectionAuth() bool`

HasCollectionAuth returns a boolean if a field has been set.

### GetYbmAccountId

`func (o *PerfAdvisorEndpointSpec) GetYbmAccountId() string`

GetYbmAccountId returns the YbmAccountId field if non-nil, zero value otherwise.

### GetYbmAccountIdOk

`func (o *PerfAdvisorEndpointSpec) GetYbmAccountIdOk() (*string, bool)`

GetYbmAccountIdOk returns a tuple with the YbmAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbmAccountId

`func (o *PerfAdvisorEndpointSpec) SetYbmAccountId(v string)`

SetYbmAccountId sets YbmAccountId field to given value.

### HasYbmAccountId

`func (o *PerfAdvisorEndpointSpec) HasYbmAccountId() bool`

HasYbmAccountId returns a boolean if a field has been set.

### GetYbmProjectId

`func (o *PerfAdvisorEndpointSpec) GetYbmProjectId() string`

GetYbmProjectId returns the YbmProjectId field if non-nil, zero value otherwise.

### GetYbmProjectIdOk

`func (o *PerfAdvisorEndpointSpec) GetYbmProjectIdOk() (*string, bool)`

GetYbmProjectIdOk returns a tuple with the YbmProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbmProjectId

`func (o *PerfAdvisorEndpointSpec) SetYbmProjectId(v string)`

SetYbmProjectId sets YbmProjectId field to given value.

### HasYbmProjectId

`func (o *PerfAdvisorEndpointSpec) HasYbmProjectId() bool`

HasYbmProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


