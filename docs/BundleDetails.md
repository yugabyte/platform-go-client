# BundleDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | **[]string** |  | 
**MaxCoreFileSize** | Pointer to **int64** | Max size of the collected cores (if any) | [optional] 
**MaxNumRecentCores** | Pointer to **int32** | Max number of most recent cores to collect (if any) | [optional] 

## Methods

### NewBundleDetails

`func NewBundleDetails(components []string, ) *BundleDetails`

NewBundleDetails instantiates a new BundleDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleDetailsWithDefaults

`func NewBundleDetailsWithDefaults() *BundleDetails`

NewBundleDetailsWithDefaults instantiates a new BundleDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *BundleDetails) GetComponents() []string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *BundleDetails) GetComponentsOk() (*[]string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *BundleDetails) SetComponents(v []string)`

SetComponents sets Components field to given value.


### GetMaxCoreFileSize

`func (o *BundleDetails) GetMaxCoreFileSize() int64`

GetMaxCoreFileSize returns the MaxCoreFileSize field if non-nil, zero value otherwise.

### GetMaxCoreFileSizeOk

`func (o *BundleDetails) GetMaxCoreFileSizeOk() (*int64, bool)`

GetMaxCoreFileSizeOk returns a tuple with the MaxCoreFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCoreFileSize

`func (o *BundleDetails) SetMaxCoreFileSize(v int64)`

SetMaxCoreFileSize sets MaxCoreFileSize field to given value.

### HasMaxCoreFileSize

`func (o *BundleDetails) HasMaxCoreFileSize() bool`

HasMaxCoreFileSize returns a boolean if a field has been set.

### GetMaxNumRecentCores

`func (o *BundleDetails) GetMaxNumRecentCores() int32`

GetMaxNumRecentCores returns the MaxNumRecentCores field if non-nil, zero value otherwise.

### GetMaxNumRecentCoresOk

`func (o *BundleDetails) GetMaxNumRecentCoresOk() (*int32, bool)`

GetMaxNumRecentCoresOk returns a tuple with the MaxNumRecentCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumRecentCores

`func (o *BundleDetails) SetMaxNumRecentCores(v int32)`

SetMaxNumRecentCores sets MaxNumRecentCores field to given value.

### HasMaxNumRecentCores

`func (o *BundleDetails) HasMaxNumRecentCores() bool`

HasMaxNumRecentCores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


