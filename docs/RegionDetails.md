# RegionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudInfo** | Pointer to [**RegionCloudInfo**](RegionCloudInfo.md) |  | [optional] 

## Methods

### NewRegionDetails

`func NewRegionDetails() *RegionDetails`

NewRegionDetails instantiates a new RegionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionDetailsWithDefaults

`func NewRegionDetailsWithDefaults() *RegionDetails`

NewRegionDetailsWithDefaults instantiates a new RegionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudInfo

`func (o *RegionDetails) GetCloudInfo() RegionCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *RegionDetails) GetCloudInfoOk() (*RegionCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *RegionDetails) SetCloudInfo(v RegionCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *RegionDetails) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


