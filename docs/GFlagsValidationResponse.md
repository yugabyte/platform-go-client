# GFlagsValidationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MASTER** | Pointer to [**GFlagValidationDetails**](GFlagValidationDetails.md) |  | [optional] 
**Name** | Pointer to **string** | WARNING: This is a preview API that could change. Name of the gflag | [optional] 
**TSERVER** | Pointer to [**GFlagValidationDetails**](GFlagValidationDetails.md) |  | [optional] 

## Methods

### NewGFlagsValidationResponse

`func NewGFlagsValidationResponse() *GFlagsValidationResponse`

NewGFlagsValidationResponse instantiates a new GFlagsValidationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGFlagsValidationResponseWithDefaults

`func NewGFlagsValidationResponseWithDefaults() *GFlagsValidationResponse`

NewGFlagsValidationResponseWithDefaults instantiates a new GFlagsValidationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMASTER

`func (o *GFlagsValidationResponse) GetMASTER() GFlagValidationDetails`

GetMASTER returns the MASTER field if non-nil, zero value otherwise.

### GetMASTEROk

`func (o *GFlagsValidationResponse) GetMASTEROk() (*GFlagValidationDetails, bool)`

GetMASTEROk returns a tuple with the MASTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMASTER

`func (o *GFlagsValidationResponse) SetMASTER(v GFlagValidationDetails)`

SetMASTER sets MASTER field to given value.

### HasMASTER

`func (o *GFlagsValidationResponse) HasMASTER() bool`

HasMASTER returns a boolean if a field has been set.

### GetName

`func (o *GFlagsValidationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GFlagsValidationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GFlagsValidationResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GFlagsValidationResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTSERVER

`func (o *GFlagsValidationResponse) GetTSERVER() GFlagValidationDetails`

GetTSERVER returns the TSERVER field if non-nil, zero value otherwise.

### GetTSERVEROk

`func (o *GFlagsValidationResponse) GetTSERVEROk() (*GFlagValidationDetails, bool)`

GetTSERVEROk returns a tuple with the TSERVER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTSERVER

`func (o *GFlagsValidationResponse) SetTSERVER(v GFlagValidationDetails)`

SetTSERVER sets TSERVER field to given value.

### HasTSERVER

`func (o *GFlagsValidationResponse) HasTSERVER() bool`

HasTSERVER returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


