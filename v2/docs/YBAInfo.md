# YBAInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FipsEnabled** | **bool** | Whether YBA instance is FIPS compliant | [readonly] 
**Version** | Pointer to **string** | YBA software version | [optional] [readonly] 
**CustomerCount** | Pointer to **int32** | number of customers created on this YBA instance | [optional] [readonly] 

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


### GetVersion

`func (o *YBAInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *YBAInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *YBAInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *YBAInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCustomerCount

`func (o *YBAInfo) GetCustomerCount() int32`

GetCustomerCount returns the CustomerCount field if non-nil, zero value otherwise.

### GetCustomerCountOk

`func (o *YBAInfo) GetCustomerCountOk() (*int32, bool)`

GetCustomerCountOk returns a tuple with the CustomerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCount

`func (o *YBAInfo) SetCustomerCount(v int32)`

SetCustomerCount sets CustomerCount field to given value.

### HasCustomerCount

`func (o *YBAInfo) HasCustomerCount() bool`

HasCustomerCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


