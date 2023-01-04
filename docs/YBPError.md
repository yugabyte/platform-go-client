# YBPError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | User-visible unstructured error message | [optional] 
**ErrorJson** | Pointer to **map[string]interface{}** | User visible structured error message as json object | [optional] 
**HttpMethod** | Pointer to **string** | Method for HTTP call that resulted in this error | [optional] 
**RequestUri** | Pointer to **string** | URI for HTTP request that resulted in this error | [optional] 
**Success** | Pointer to **bool** | Always set to false to indicate failure | [optional] 

## Methods

### NewYBPError

`func NewYBPError() *YBPError`

NewYBPError instantiates a new YBPError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYBPErrorWithDefaults

`func NewYBPErrorWithDefaults() *YBPError`

NewYBPErrorWithDefaults instantiates a new YBPError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *YBPError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *YBPError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *YBPError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *YBPError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorJson

`func (o *YBPError) GetErrorJson() map[string]interface{}`

GetErrorJson returns the ErrorJson field if non-nil, zero value otherwise.

### GetErrorJsonOk

`func (o *YBPError) GetErrorJsonOk() (*map[string]interface{}, bool)`

GetErrorJsonOk returns a tuple with the ErrorJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorJson

`func (o *YBPError) SetErrorJson(v map[string]interface{})`

SetErrorJson sets ErrorJson field to given value.

### HasErrorJson

`func (o *YBPError) HasErrorJson() bool`

HasErrorJson returns a boolean if a field has been set.

### GetHttpMethod

`func (o *YBPError) GetHttpMethod() string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *YBPError) GetHttpMethodOk() (*string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *YBPError) SetHttpMethod(v string)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *YBPError) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetRequestUri

`func (o *YBPError) GetRequestUri() string`

GetRequestUri returns the RequestUri field if non-nil, zero value otherwise.

### GetRequestUriOk

`func (o *YBPError) GetRequestUriOk() (*string, bool)`

GetRequestUriOk returns a tuple with the RequestUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUri

`func (o *YBPError) SetRequestUri(v string)`

SetRequestUri sets RequestUri field to given value.

### HasRequestUri

`func (o *YBPError) HasRequestUri() bool`

HasRequestUri returns a boolean if a field has been set.

### GetSuccess

`func (o *YBPError) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *YBPError) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *YBPError) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *YBPError) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


