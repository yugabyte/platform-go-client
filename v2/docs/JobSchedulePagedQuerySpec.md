# JobSchedulePagedQuerySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**JobScheduleApiFilter**](JobScheduleApiFilter.md) |  | [optional] 
**SortBy** | Pointer to **string** | Sort fields of the records. | [optional] 
**Offset** | Pointer to **int32** | Start offset of the records. | [optional] 
**Limit** | Pointer to **int32** | Maximum number of records to be fetched. | [optional] 
**Direction** | Pointer to **string** | Sort order of the records. | [optional] 

## Methods

### NewJobSchedulePagedQuerySpec

`func NewJobSchedulePagedQuerySpec() *JobSchedulePagedQuerySpec`

NewJobSchedulePagedQuerySpec instantiates a new JobSchedulePagedQuerySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobSchedulePagedQuerySpecWithDefaults

`func NewJobSchedulePagedQuerySpecWithDefaults() *JobSchedulePagedQuerySpec`

NewJobSchedulePagedQuerySpecWithDefaults instantiates a new JobSchedulePagedQuerySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *JobSchedulePagedQuerySpec) GetFilter() JobScheduleApiFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *JobSchedulePagedQuerySpec) GetFilterOk() (*JobScheduleApiFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *JobSchedulePagedQuerySpec) SetFilter(v JobScheduleApiFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *JobSchedulePagedQuerySpec) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSortBy

`func (o *JobSchedulePagedQuerySpec) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *JobSchedulePagedQuerySpec) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *JobSchedulePagedQuerySpec) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *JobSchedulePagedQuerySpec) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetOffset

`func (o *JobSchedulePagedQuerySpec) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *JobSchedulePagedQuerySpec) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *JobSchedulePagedQuerySpec) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *JobSchedulePagedQuerySpec) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *JobSchedulePagedQuerySpec) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *JobSchedulePagedQuerySpec) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *JobSchedulePagedQuerySpec) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *JobSchedulePagedQuerySpec) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetDirection

`func (o *JobSchedulePagedQuerySpec) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *JobSchedulePagedQuerySpec) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *JobSchedulePagedQuerySpec) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *JobSchedulePagedQuerySpec) HasDirection() bool`

HasDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


