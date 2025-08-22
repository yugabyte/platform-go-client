# YBAInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FipsEnabled** | **bool** | Whether YBA instance is FIPS compliant | [readonly] 

## Methods

### NewYBAInfo

`func NewYBAInfo(fipsEnabled bool, ) *YBAInfo`

NewYBAInfo instantiates a new YBAInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYBAInfoWithDefaults

`func NewYBAInfoWithDefaults() *YBAInfo`

NewYBAInfoWithDefaults instantiates a new YBAInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFipsEnabled

`func (o *YBAInfo) GetFipsEnabled() bool`

GetFipsEnabled returns the FipsEnabled field if non-nil, zero value otherwise.

### GetFipsEnabledOk

`func (o *YBAInfo) GetFipsEnabledOk() (*bool, bool)`

GetFipsEnabledOk returns a tuple with the FipsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFipsEnabled

`func (o *YBAInfo) SetFipsEnabled(v bool)`

SetFipsEnabled sets FipsEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


