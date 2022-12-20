# PerformanceRecommendation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** |  | 
**EntityNames** | **string** |  | 
**EntityType** | **string** |  | 
**Id** | **string** |  | 
**New** | **bool** |  | 
**Observation** | **string** |  | 
**Recommendation** | **string** |  | 
**RecommendationInfo** | **map[string]map[string]interface{}** |  | 
**RecommendationPriority** | **string** |  | 
**RecommendationState** | **string** |  | 
**RecommendationTimestamp** | **int64** |  | 
**RecommendationType** | **string** |  | 
**StateChangeAuditInfoList** | [**[]StateChangeAuditInfo**](StateChangeAuditInfo.md) |  | 
**UniverseId** | **string** |  | 

## Methods

### NewPerformanceRecommendation

`func NewPerformanceRecommendation(customerId string, entityNames string, entityType string, id string, new bool, observation string, recommendation string, recommendationInfo map[string]map[string]interface{}, recommendationPriority string, recommendationState string, recommendationTimestamp int64, recommendationType string, stateChangeAuditInfoList []StateChangeAuditInfo, universeId string, ) *PerformanceRecommendation`

NewPerformanceRecommendation instantiates a new PerformanceRecommendation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceRecommendationWithDefaults

`func NewPerformanceRecommendationWithDefaults() *PerformanceRecommendation`

NewPerformanceRecommendationWithDefaults instantiates a new PerformanceRecommendation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *PerformanceRecommendation) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PerformanceRecommendation) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PerformanceRecommendation) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetEntityNames

`func (o *PerformanceRecommendation) GetEntityNames() string`

GetEntityNames returns the EntityNames field if non-nil, zero value otherwise.

### GetEntityNamesOk

`func (o *PerformanceRecommendation) GetEntityNamesOk() (*string, bool)`

GetEntityNamesOk returns a tuple with the EntityNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityNames

`func (o *PerformanceRecommendation) SetEntityNames(v string)`

SetEntityNames sets EntityNames field to given value.


### GetEntityType

`func (o *PerformanceRecommendation) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *PerformanceRecommendation) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *PerformanceRecommendation) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetId

`func (o *PerformanceRecommendation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PerformanceRecommendation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PerformanceRecommendation) SetId(v string)`

SetId sets Id field to given value.


### GetNew

`func (o *PerformanceRecommendation) GetNew() bool`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *PerformanceRecommendation) GetNewOk() (*bool, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *PerformanceRecommendation) SetNew(v bool)`

SetNew sets New field to given value.


### GetObservation

`func (o *PerformanceRecommendation) GetObservation() string`

GetObservation returns the Observation field if non-nil, zero value otherwise.

### GetObservationOk

`func (o *PerformanceRecommendation) GetObservationOk() (*string, bool)`

GetObservationOk returns a tuple with the Observation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservation

`func (o *PerformanceRecommendation) SetObservation(v string)`

SetObservation sets Observation field to given value.


### GetRecommendation

`func (o *PerformanceRecommendation) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *PerformanceRecommendation) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *PerformanceRecommendation) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.


### GetRecommendationInfo

`func (o *PerformanceRecommendation) GetRecommendationInfo() map[string]map[string]interface{}`

GetRecommendationInfo returns the RecommendationInfo field if non-nil, zero value otherwise.

### GetRecommendationInfoOk

`func (o *PerformanceRecommendation) GetRecommendationInfoOk() (*map[string]map[string]interface{}, bool)`

GetRecommendationInfoOk returns a tuple with the RecommendationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationInfo

`func (o *PerformanceRecommendation) SetRecommendationInfo(v map[string]map[string]interface{})`

SetRecommendationInfo sets RecommendationInfo field to given value.


### GetRecommendationPriority

`func (o *PerformanceRecommendation) GetRecommendationPriority() string`

GetRecommendationPriority returns the RecommendationPriority field if non-nil, zero value otherwise.

### GetRecommendationPriorityOk

`func (o *PerformanceRecommendation) GetRecommendationPriorityOk() (*string, bool)`

GetRecommendationPriorityOk returns a tuple with the RecommendationPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationPriority

`func (o *PerformanceRecommendation) SetRecommendationPriority(v string)`

SetRecommendationPriority sets RecommendationPriority field to given value.


### GetRecommendationState

`func (o *PerformanceRecommendation) GetRecommendationState() string`

GetRecommendationState returns the RecommendationState field if non-nil, zero value otherwise.

### GetRecommendationStateOk

`func (o *PerformanceRecommendation) GetRecommendationStateOk() (*string, bool)`

GetRecommendationStateOk returns a tuple with the RecommendationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationState

`func (o *PerformanceRecommendation) SetRecommendationState(v string)`

SetRecommendationState sets RecommendationState field to given value.


### GetRecommendationTimestamp

`func (o *PerformanceRecommendation) GetRecommendationTimestamp() int64`

GetRecommendationTimestamp returns the RecommendationTimestamp field if non-nil, zero value otherwise.

### GetRecommendationTimestampOk

`func (o *PerformanceRecommendation) GetRecommendationTimestampOk() (*int64, bool)`

GetRecommendationTimestampOk returns a tuple with the RecommendationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationTimestamp

`func (o *PerformanceRecommendation) SetRecommendationTimestamp(v int64)`

SetRecommendationTimestamp sets RecommendationTimestamp field to given value.


### GetRecommendationType

`func (o *PerformanceRecommendation) GetRecommendationType() string`

GetRecommendationType returns the RecommendationType field if non-nil, zero value otherwise.

### GetRecommendationTypeOk

`func (o *PerformanceRecommendation) GetRecommendationTypeOk() (*string, bool)`

GetRecommendationTypeOk returns a tuple with the RecommendationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationType

`func (o *PerformanceRecommendation) SetRecommendationType(v string)`

SetRecommendationType sets RecommendationType field to given value.


### GetStateChangeAuditInfoList

`func (o *PerformanceRecommendation) GetStateChangeAuditInfoList() []StateChangeAuditInfo`

GetStateChangeAuditInfoList returns the StateChangeAuditInfoList field if non-nil, zero value otherwise.

### GetStateChangeAuditInfoListOk

`func (o *PerformanceRecommendation) GetStateChangeAuditInfoListOk() (*[]StateChangeAuditInfo, bool)`

GetStateChangeAuditInfoListOk returns a tuple with the StateChangeAuditInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChangeAuditInfoList

`func (o *PerformanceRecommendation) SetStateChangeAuditInfoList(v []StateChangeAuditInfo)`

SetStateChangeAuditInfoList sets StateChangeAuditInfoList field to given value.


### GetUniverseId

`func (o *PerformanceRecommendation) GetUniverseId() string`

GetUniverseId returns the UniverseId field if non-nil, zero value otherwise.

### GetUniverseIdOk

`func (o *PerformanceRecommendation) GetUniverseIdOk() (*string, bool)`

GetUniverseIdOk returns a tuple with the UniverseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniverseId

`func (o *PerformanceRecommendation) SetUniverseId(v string)`

SetUniverseId sets UniverseId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


