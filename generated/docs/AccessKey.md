# AccessKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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


