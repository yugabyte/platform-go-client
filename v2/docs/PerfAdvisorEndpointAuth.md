# PerfAdvisorEndpointAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Authentication type. | 
**Username** | Pointer to **string** | Username. Required for BASIC. | [optional] 
**Password** | Pointer to **string** | Password. Required for BASIC. Read back masked, and a masked value sent on an edit keeps the stored password rather than overwriting it.  | [optional] 

## Methods

### NewPerfAdvisorEndpointAuth

`func NewPerfAdvisorEndpointAuth(type_ string, ) *PerfAdvisorEndpointAuth`

NewPerfAdvisorEndpointAuth instantiates a new PerfAdvisorEndpointAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerfAdvisorEndpointAuthWithDefaults

`func NewPerfAdvisorEndpointAuthWithDefaults() *PerfAdvisorEndpointAuth`

NewPerfAdvisorEndpointAuthWithDefaults instantiates a new PerfAdvisorEndpointAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PerfAdvisorEndpointAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PerfAdvisorEndpointAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PerfAdvisorEndpointAuth) SetType(v string)`

SetType sets Type field to given value.


### GetUsername

`func (o *PerfAdvisorEndpointAuth) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PerfAdvisorEndpointAuth) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PerfAdvisorEndpointAuth) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PerfAdvisorEndpointAuth) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *PerfAdvisorEndpointAuth) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PerfAdvisorEndpointAuth) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PerfAdvisorEndpointAuth) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PerfAdvisorEndpointAuth) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


