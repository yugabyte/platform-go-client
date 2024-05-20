# ClientCertParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertExpiry** | **int64** |  | 
**CertStart** | **int64** |  | 
**Username** | **string** |  | 

## Methods

### NewClientCertParams

`func NewClientCertParams(certExpiry int64, certStart int64, username string, ) *ClientCertParams`

NewClientCertParams instantiates a new ClientCertParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCertParamsWithDefaults

`func NewClientCertParamsWithDefaults() *ClientCertParams`

NewClientCertParamsWithDefaults instantiates a new ClientCertParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertExpiry

`func (o *ClientCertParams) GetCertExpiry() int64`

GetCertExpiry returns the CertExpiry field if non-nil, zero value otherwise.

### GetCertExpiryOk

`func (o *ClientCertParams) GetCertExpiryOk() (*int64, bool)`

GetCertExpiryOk returns a tuple with the CertExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertExpiry

`func (o *ClientCertParams) SetCertExpiry(v int64)`

SetCertExpiry sets CertExpiry field to given value.


### GetCertStart

`func (o *ClientCertParams) GetCertStart() int64`

GetCertStart returns the CertStart field if non-nil, zero value otherwise.

### GetCertStartOk

`func (o *ClientCertParams) GetCertStartOk() (*int64, bool)`

GetCertStartOk returns a tuple with the CertStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStart

`func (o *ClientCertParams) SetCertStart(v int64)`

SetCertStart sets CertStart field to given value.


### GetUsername

`func (o *ClientCertParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ClientCertParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ClientCertParams) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


