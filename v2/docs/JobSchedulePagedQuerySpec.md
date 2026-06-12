# JobSchedulePagedQuerySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**JobScheduleApiFilter**](JobScheduleApiFilter.md) |  | [optional] 
**SortBy** | Pointer to **string** | Sort fields of the records. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


