# GFlagsValidationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MASTER** | Pointer to **string** | WARNING: This is a preview API that could change. Value of the gflag for master | [optional] 
**Name** | Pointer to **string** | WARNING: This is a preview API that could change. Name of the gflag | [optional] 
**TSERVER** | Pointer to **string** | WARNING: This is a preview API that could change. Value of the gflag for tserver | [optional] 

## Methods

### NewGFlagsValidationRequest

`func NewGFlagsValidationRequest() *GFlagsValidationRequest`

NewGFlagsValidationRequest instantiates a new GFlagsValidationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGFlagsValidationRequestWithDefaults

`func NewGFlagsValidationRequestWithDefaults() *GFlagsValidationRequest`

NewGFlagsValidationRequestWithDefaults instantiates a new GFlagsValidationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMASTER

`func (o *GFlagsValidationRequest) GetMASTER() string`

GetMASTER returns the MASTER field if non-nil, zero value otherwise.

### GetMASTEROk

`func (o *GFlagsValidationRequest) GetMASTEROk() (*string, bool)`

GetMASTEROk returns a tuple with the MASTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMASTER

`func (o *GFlagsValidationRequest) SetMASTER(v string)`

SetMASTER sets MASTER field to given value.

### HasMASTER

`func (o *GFlagsValidationRequest) HasMASTER() bool`

HasMASTER returns a boolean if a field has been set.

### GetName

`func (o *GFlagsValidationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GFlagsValidationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GFlagsValidationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GFlagsValidationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTSERVER

`func (o *GFlagsValidationRequest) GetTSERVER() string`

GetTSERVER returns the TSERVER field if non-nil, zero value otherwise.

### GetTSERVEROk

`func (o *GFlagsValidationRequest) GetTSERVEROk() (*string, bool)`

GetTSERVEROk returns a tuple with the TSERVER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTSERVER

`func (o *GFlagsValidationRequest) SetTSERVER(v string)`

SetTSERVER sets TSERVER field to given value.

### HasTSERVER

`func (o *GFlagsValidationRequest) HasTSERVER() bool`

HasTSERVER returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


