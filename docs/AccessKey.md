# AccessKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | Pointer to **time.Time** | Creation date of key | [optional] [readonly] 
**ExpirationDate** | Pointer to **time.Time** | Expiration date of key | [optional] 
**IdKey** | [**AccessKeyId**](AccessKeyId.md) |  | 
**KeyInfo** | [**KeyInfo**](KeyInfo.md) |  | 

## Methods

### NewAccessKey

`func NewAccessKey(idKey AccessKeyId, keyInfo KeyInfo, ) *AccessKey`

NewAccessKey instantiates a new AccessKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessKeyWithDefaults

`func NewAccessKeyWithDefaults() *AccessKey`

NewAccessKeyWithDefaults instantiates a new AccessKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDate

`func (o *AccessKey) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *AccessKey) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *AccessKey) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *AccessKey) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *AccessKey) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *AccessKey) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *AccessKey) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *AccessKey) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetIdKey

`func (o *AccessKey) GetIdKey() AccessKeyId`

GetIdKey returns the IdKey field if non-nil, zero value otherwise.

### GetIdKeyOk

`func (o *AccessKey) GetIdKeyOk() (*AccessKeyId, bool)`

GetIdKeyOk returns a tuple with the IdKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdKey

`func (o *AccessKey) SetIdKey(v AccessKeyId)`

SetIdKey sets IdKey field to given value.


### GetKeyInfo

`func (o *AccessKey) GetKeyInfo() KeyInfo`

GetKeyInfo returns the KeyInfo field if non-nil, zero value otherwise.

### GetKeyInfoOk

`func (o *AccessKey) GetKeyInfoOk() (*KeyInfo, bool)`

GetKeyInfoOk returns a tuple with the KeyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyInfo

`func (o *AccessKey) SetKeyInfo(v KeyInfo)`

SetKeyInfo sets KeyInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


