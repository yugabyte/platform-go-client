# PitrConfigApiFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DrConfigUuid** | Pointer to **string** | When set, only PITR configs linked to the given disaster recovery config are returned.  | [optional] 

## Methods

### NewPitrConfigApiFilter

`func NewPitrConfigApiFilter() *PitrConfigApiFilter`

NewPitrConfigApiFilter instantiates a new PitrConfigApiFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPitrConfigApiFilterWithDefaults

`func NewPitrConfigApiFilterWithDefaults() *PitrConfigApiFilter`

NewPitrConfigApiFilterWithDefaults instantiates a new PitrConfigApiFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrConfigUuid

`func (o *PitrConfigApiFilter) GetDrConfigUuid() string`

GetDrConfigUuid returns the DrConfigUuid field if non-nil, zero value otherwise.

### GetDrConfigUuidOk

`func (o *PitrConfigApiFilter) GetDrConfigUuidOk() (*string, bool)`

GetDrConfigUuidOk returns a tuple with the DrConfigUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrConfigUuid

`func (o *PitrConfigApiFilter) SetDrConfigUuid(v string)`

SetDrConfigUuid sets DrConfigUuid field to given value.

### HasDrConfigUuid

`func (o *PitrConfigApiFilter) HasDrConfigUuid() bool`

HasDrConfigUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


