# Details

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]NodeData**](NodeData.md) |  | 
**HasError** | **bool** |  | 
**HasWarning** | **bool** |  | 
**TimestampIso** | Pointer to **time.Time** |  | [optional] 
**YbVersion** | **string** |  | 

## Methods

### NewDetails

`func NewDetails(data []NodeData, hasError bool, hasWarning bool, ybVersion string, ) *Details`

NewDetails instantiates a new Details object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailsWithDefaults

`func NewDetailsWithDefaults() *Details`

NewDetailsWithDefaults instantiates a new Details object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Details) GetData() []NodeData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Details) GetDataOk() (*[]NodeData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Details) SetData(v []NodeData)`

SetData sets Data field to given value.


### GetHasError

`func (o *Details) GetHasError() bool`

GetHasError returns the HasError field if non-nil, zero value otherwise.

### GetHasErrorOk

`func (o *Details) GetHasErrorOk() (*bool, bool)`

GetHasErrorOk returns a tuple with the HasError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasError

`func (o *Details) SetHasError(v bool)`

SetHasError sets HasError field to given value.


### GetHasWarning

`func (o *Details) GetHasWarning() bool`

GetHasWarning returns the HasWarning field if non-nil, zero value otherwise.

### GetHasWarningOk

`func (o *Details) GetHasWarningOk() (*bool, bool)`

GetHasWarningOk returns a tuple with the HasWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWarning

`func (o *Details) SetHasWarning(v bool)`

SetHasWarning sets HasWarning field to given value.


### GetTimestampIso

`func (o *Details) GetTimestampIso() time.Time`

GetTimestampIso returns the TimestampIso field if non-nil, zero value otherwise.

### GetTimestampIsoOk

`func (o *Details) GetTimestampIsoOk() (*time.Time, bool)`

GetTimestampIsoOk returns a tuple with the TimestampIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampIso

`func (o *Details) SetTimestampIso(v time.Time)`

SetTimestampIso sets TimestampIso field to given value.

### HasTimestampIso

`func (o *Details) HasTimestampIso() bool`

HasTimestampIso returns a boolean if a field has been set.

### GetYbVersion

`func (o *Details) GetYbVersion() string`

GetYbVersion returns the YbVersion field if non-nil, zero value otherwise.

### GetYbVersionOk

`func (o *Details) GetYbVersionOk() (*string, bool)`

GetYbVersionOk returns a tuple with the YbVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbVersion

`func (o *Details) SetYbVersion(v string)`

SetYbVersion sets YbVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


