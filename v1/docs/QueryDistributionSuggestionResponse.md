# QueryDistributionSuggestionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdviceType** | **string** |  | 
**Description** | **string** |  | 
**Details** | [**[]NodeQueryDistributionDetails**](NodeQueryDistributionDetails.md) |  | 
**EndTime** | **string** |  | 
**StartTime** | **string** |  | 
**Suggestion** | **string** |  | 

## Methods

### NewQueryDistributionSuggestionResponse

`func NewQueryDistributionSuggestionResponse(adviceType string, description string, details []NodeQueryDistributionDetails, endTime string, startTime string, suggestion string, ) *QueryDistributionSuggestionResponse`

NewQueryDistributionSuggestionResponse instantiates a new QueryDistributionSuggestionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryDistributionSuggestionResponseWithDefaults

`func NewQueryDistributionSuggestionResponseWithDefaults() *QueryDistributionSuggestionResponse`

NewQueryDistributionSuggestionResponseWithDefaults instantiates a new QueryDistributionSuggestionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdviceType

`func (o *QueryDistributionSuggestionResponse) GetAdviceType() string`

GetAdviceType returns the AdviceType field if non-nil, zero value otherwise.

### GetAdviceTypeOk

`func (o *QueryDistributionSuggestionResponse) GetAdviceTypeOk() (*string, bool)`

GetAdviceTypeOk returns a tuple with the AdviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdviceType

`func (o *QueryDistributionSuggestionResponse) SetAdviceType(v string)`

SetAdviceType sets AdviceType field to given value.


### GetDescription

`func (o *QueryDistributionSuggestionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QueryDistributionSuggestionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QueryDistributionSuggestionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDetails

`func (o *QueryDistributionSuggestionResponse) GetDetails() []NodeQueryDistributionDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *QueryDistributionSuggestionResponse) GetDetailsOk() (*[]NodeQueryDistributionDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *QueryDistributionSuggestionResponse) SetDetails(v []NodeQueryDistributionDetails)`

SetDetails sets Details field to given value.


### GetEndTime

`func (o *QueryDistributionSuggestionResponse) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *QueryDistributionSuggestionResponse) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *QueryDistributionSuggestionResponse) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.


### GetStartTime

`func (o *QueryDistributionSuggestionResponse) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *QueryDistributionSuggestionResponse) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *QueryDistributionSuggestionResponse) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetSuggestion

`func (o *QueryDistributionSuggestionResponse) GetSuggestion() string`

GetSuggestion returns the Suggestion field if non-nil, zero value otherwise.

### GetSuggestionOk

`func (o *QueryDistributionSuggestionResponse) GetSuggestionOk() (*string, bool)`

GetSuggestionOk returns a tuple with the Suggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestion

`func (o *QueryDistributionSuggestionResponse) SetSuggestion(v string)`

SetSuggestion sets Suggestion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


