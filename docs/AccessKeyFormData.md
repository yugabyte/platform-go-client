# AccessKeyFormData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationThresholdDays** | Pointer to **int32** |  | [optional] 
**KeyCode** | **string** |  | 
**KeyContent** | **string** |  | 
**KeyType** | **string** |  | 

## Methods

### NewAccessKeyFormData

`func NewAccessKeyFormData(keyCode string, keyContent string, keyType string, ) *AccessKeyFormData`

NewAccessKeyFormData instantiates a new AccessKeyFormData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessKeyFormDataWithDefaults

`func NewAccessKeyFormDataWithDefaults() *AccessKeyFormData`

NewAccessKeyFormDataWithDefaults instantiates a new AccessKeyFormData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationThresholdDays

`func (o *AccessKeyFormData) GetExpirationThresholdDays() int32`

GetExpirationThresholdDays returns the ExpirationThresholdDays field if non-nil, zero value otherwise.

### GetExpirationThresholdDaysOk

`func (o *AccessKeyFormData) GetExpirationThresholdDaysOk() (*int32, bool)`

GetExpirationThresholdDaysOk returns a tuple with the ExpirationThresholdDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationThresholdDays

`func (o *AccessKeyFormData) SetExpirationThresholdDays(v int32)`

SetExpirationThresholdDays sets ExpirationThresholdDays field to given value.

### HasExpirationThresholdDays

`func (o *AccessKeyFormData) HasExpirationThresholdDays() bool`

HasExpirationThresholdDays returns a boolean if a field has been set.

### GetKeyCode

`func (o *AccessKeyFormData) GetKeyCode() string`

GetKeyCode returns the KeyCode field if non-nil, zero value otherwise.

### GetKeyCodeOk

`func (o *AccessKeyFormData) GetKeyCodeOk() (*string, bool)`

GetKeyCodeOk returns a tuple with the KeyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyCode

`func (o *AccessKeyFormData) SetKeyCode(v string)`

SetKeyCode sets KeyCode field to given value.


### GetKeyContent

`func (o *AccessKeyFormData) GetKeyContent() string`

GetKeyContent returns the KeyContent field if non-nil, zero value otherwise.

### GetKeyContentOk

`func (o *AccessKeyFormData) GetKeyContentOk() (*string, bool)`

GetKeyContentOk returns a tuple with the KeyContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyContent

`func (o *AccessKeyFormData) SetKeyContent(v string)`

SetKeyContent sets KeyContent field to given value.


### GetKeyType

`func (o *AccessKeyFormData) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *AccessKeyFormData) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *AccessKeyFormData) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


