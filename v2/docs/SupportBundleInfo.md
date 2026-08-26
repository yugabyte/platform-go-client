# SupportBundleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Support bundle unique identifier. | [optional] [readonly] 
**Status** | Pointer to [**SupportBundleStatus**](SupportBundleStatus.md) |  | [optional] 
**ScopeUuid** | Pointer to **NullableString** | UUID of the universe this support bundle was collected from. Null for YBA-only bundles. | [optional] [readonly] 
**CustomerUuid** | Pointer to **NullableString** | UUID of the customer that owns this bundle. Set only for YBA-only bundles. | [optional] [readonly] 
**CreationDate** | Pointer to **time.Time** | When the bundle archive was created. Null until collection succeeds. | [optional] [readonly] 
**ExpirationDate** | Pointer to **time.Time** | When the bundle archive is eligible for automatic cleanup. Null until collection succeeds. | [optional] [readonly] 
**SizeInBytes** | Pointer to **int64** | Size in bytes of the bundle archive. Zero until collection succeeds. | [optional] [readonly] 

## Methods

### NewSupportBundleInfo

`func NewSupportBundleInfo() *SupportBundleInfo`

NewSupportBundleInfo instantiates a new SupportBundleInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportBundleInfoWithDefaults

`func NewSupportBundleInfoWithDefaults() *SupportBundleInfo`

NewSupportBundleInfoWithDefaults instantiates a new SupportBundleInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *SupportBundleInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SupportBundleInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SupportBundleInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SupportBundleInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetStatus

`func (o *SupportBundleInfo) GetStatus() SupportBundleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SupportBundleInfo) GetStatusOk() (*SupportBundleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SupportBundleInfo) SetStatus(v SupportBundleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SupportBundleInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetScopeUuid

`func (o *SupportBundleInfo) GetScopeUuid() string`

GetScopeUuid returns the ScopeUuid field if non-nil, zero value otherwise.

### GetScopeUuidOk

`func (o *SupportBundleInfo) GetScopeUuidOk() (*string, bool)`

GetScopeUuidOk returns a tuple with the ScopeUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeUuid

`func (o *SupportBundleInfo) SetScopeUuid(v string)`

SetScopeUuid sets ScopeUuid field to given value.

### HasScopeUuid

`func (o *SupportBundleInfo) HasScopeUuid() bool`

HasScopeUuid returns a boolean if a field has been set.

### SetScopeUuidNil

`func (o *SupportBundleInfo) SetScopeUuidNil(b bool)`

 SetScopeUuidNil sets the value for ScopeUuid to be an explicit nil

### UnsetScopeUuid
`func (o *SupportBundleInfo) UnsetScopeUuid()`

UnsetScopeUuid ensures that no value is present for ScopeUuid, not even an explicit nil
### GetCustomerUuid

`func (o *SupportBundleInfo) GetCustomerUuid() string`

GetCustomerUuid returns the CustomerUuid field if non-nil, zero value otherwise.

### GetCustomerUuidOk

`func (o *SupportBundleInfo) GetCustomerUuidOk() (*string, bool)`

GetCustomerUuidOk returns a tuple with the CustomerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUuid

`func (o *SupportBundleInfo) SetCustomerUuid(v string)`

SetCustomerUuid sets CustomerUuid field to given value.

### HasCustomerUuid

`func (o *SupportBundleInfo) HasCustomerUuid() bool`

HasCustomerUuid returns a boolean if a field has been set.

### SetCustomerUuidNil

`func (o *SupportBundleInfo) SetCustomerUuidNil(b bool)`

 SetCustomerUuidNil sets the value for CustomerUuid to be an explicit nil

### UnsetCustomerUuid
`func (o *SupportBundleInfo) UnsetCustomerUuid()`

UnsetCustomerUuid ensures that no value is present for CustomerUuid, not even an explicit nil
### GetCreationDate

`func (o *SupportBundleInfo) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *SupportBundleInfo) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *SupportBundleInfo) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *SupportBundleInfo) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *SupportBundleInfo) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *SupportBundleInfo) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *SupportBundleInfo) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *SupportBundleInfo) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetSizeInBytes

`func (o *SupportBundleInfo) GetSizeInBytes() int64`

GetSizeInBytes returns the SizeInBytes field if non-nil, zero value otherwise.

### GetSizeInBytesOk

`func (o *SupportBundleInfo) GetSizeInBytesOk() (*int64, bool)`

GetSizeInBytesOk returns a tuple with the SizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInBytes

`func (o *SupportBundleInfo) SetSizeInBytes(v int64)`

SetSizeInBytes sets SizeInBytes field to given value.

### HasSizeInBytes

`func (o *SupportBundleInfo) HasSizeInBytes() bool`

HasSizeInBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


