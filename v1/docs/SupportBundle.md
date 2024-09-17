# SupportBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleDetails** | [**BundleDetails**](BundleDetails.md) |  | 
**BundleUUID** | **string** |  | 
**CreationDate** | Pointer to **time.Time** | Support bundle creation date. | [optional] 
**EndDate** | Pointer to **time.Time** | Support bundle end date. | [optional] 
**ExpirationDate** | Pointer to **time.Time** | Support bundle expiration date. | [optional] 
**Path** | **string** |  | 
**ScopeUUID** | **string** |  | 
**SizeInBytes** | Pointer to **int64** | Size in bytes of the support bundle | [optional] 
**StartDate** | Pointer to **time.Time** | Support bundle start date. | [optional] 
**Status** | **string** |  | 

## Methods

### NewSupportBundle

`func NewSupportBundle(bundleDetails BundleDetails, bundleUUID string, path string, scopeUUID string, status string, ) *SupportBundle`

NewSupportBundle instantiates a new SupportBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportBundleWithDefaults

`func NewSupportBundleWithDefaults() *SupportBundle`

NewSupportBundleWithDefaults instantiates a new SupportBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleDetails

`func (o *SupportBundle) GetBundleDetails() BundleDetails`

GetBundleDetails returns the BundleDetails field if non-nil, zero value otherwise.

### GetBundleDetailsOk

`func (o *SupportBundle) GetBundleDetailsOk() (*BundleDetails, bool)`

GetBundleDetailsOk returns a tuple with the BundleDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleDetails

`func (o *SupportBundle) SetBundleDetails(v BundleDetails)`

SetBundleDetails sets BundleDetails field to given value.


### GetBundleUUID

`func (o *SupportBundle) GetBundleUUID() string`

GetBundleUUID returns the BundleUUID field if non-nil, zero value otherwise.

### GetBundleUUIDOk

`func (o *SupportBundle) GetBundleUUIDOk() (*string, bool)`

GetBundleUUIDOk returns a tuple with the BundleUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleUUID

`func (o *SupportBundle) SetBundleUUID(v string)`

SetBundleUUID sets BundleUUID field to given value.


### GetCreationDate

`func (o *SupportBundle) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *SupportBundle) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *SupportBundle) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *SupportBundle) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetEndDate

`func (o *SupportBundle) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *SupportBundle) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *SupportBundle) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *SupportBundle) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *SupportBundle) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *SupportBundle) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *SupportBundle) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *SupportBundle) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetPath

`func (o *SupportBundle) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SupportBundle) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SupportBundle) SetPath(v string)`

SetPath sets Path field to given value.


### GetScopeUUID

`func (o *SupportBundle) GetScopeUUID() string`

GetScopeUUID returns the ScopeUUID field if non-nil, zero value otherwise.

### GetScopeUUIDOk

`func (o *SupportBundle) GetScopeUUIDOk() (*string, bool)`

GetScopeUUIDOk returns a tuple with the ScopeUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeUUID

`func (o *SupportBundle) SetScopeUUID(v string)`

SetScopeUUID sets ScopeUUID field to given value.


### GetSizeInBytes

`func (o *SupportBundle) GetSizeInBytes() int64`

GetSizeInBytes returns the SizeInBytes field if non-nil, zero value otherwise.

### GetSizeInBytesOk

`func (o *SupportBundle) GetSizeInBytesOk() (*int64, bool)`

GetSizeInBytesOk returns a tuple with the SizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInBytes

`func (o *SupportBundle) SetSizeInBytes(v int64)`

SetSizeInBytes sets SizeInBytes field to given value.

### HasSizeInBytes

`func (o *SupportBundle) HasSizeInBytes() bool`

HasSizeInBytes returns a boolean if a field has been set.

### GetStartDate

`func (o *SupportBundle) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SupportBundle) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SupportBundle) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SupportBundle) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *SupportBundle) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SupportBundle) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SupportBundle) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


