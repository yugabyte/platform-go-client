# SupportBundleSizeEstimateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]map[string]int64** | A map of universe node names to component sizes for the given support bundle payload. Global component sizes are mapped to \&quot;YBA\&quot; node. | [optional] 

## Methods

### NewSupportBundleSizeEstimateResponse

`func NewSupportBundleSizeEstimateResponse() *SupportBundleSizeEstimateResponse`

NewSupportBundleSizeEstimateResponse instantiates a new SupportBundleSizeEstimateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportBundleSizeEstimateResponseWithDefaults

`func NewSupportBundleSizeEstimateResponseWithDefaults() *SupportBundleSizeEstimateResponse`

NewSupportBundleSizeEstimateResponseWithDefaults instantiates a new SupportBundleSizeEstimateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SupportBundleSizeEstimateResponse) GetData() map[string]map[string]int64`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SupportBundleSizeEstimateResponse) GetDataOk() (*map[string]map[string]int64, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SupportBundleSizeEstimateResponse) SetData(v map[string]map[string]int64)`

SetData sets Data field to given value.

### HasData

`func (o *SupportBundleSizeEstimateResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


