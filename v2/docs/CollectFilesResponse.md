# CollectFilesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | [**FileCollectionSummary**](FileCollectionSummary.md) |  | 
**Results** | [**map[string]NodeFileCollectionResult**](NodeFileCollectionResult.md) | Map of node name to file collection result | 

## Methods

### NewCollectFilesResponse

`func NewCollectFilesResponse(summary FileCollectionSummary, results map[string]NodeFileCollectionResult, ) *CollectFilesResponse`

NewCollectFilesResponse instantiates a new CollectFilesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectFilesResponseWithDefaults

`func NewCollectFilesResponseWithDefaults() *CollectFilesResponse`

NewCollectFilesResponseWithDefaults instantiates a new CollectFilesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *CollectFilesResponse) GetSummary() FileCollectionSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CollectFilesResponse) GetSummaryOk() (*FileCollectionSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CollectFilesResponse) SetSummary(v FileCollectionSummary)`

SetSummary sets Summary field to given value.


### GetResults

`func (o *CollectFilesResponse) GetResults() map[string]NodeFileCollectionResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectFilesResponse) GetResultsOk() (*map[string]NodeFileCollectionResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectFilesResponse) SetResults(v map[string]NodeFileCollectionResult)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


