# AvailabilityZoneData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | AZ code | 
**Name** | **string** | AZ name | 
**SecondarySubnet** | Pointer to **string** | AZ secondary subnet | [optional] 
**Subnet** | Pointer to **string** | AZ subnet | [optional] 

## Methods

### NewAvailabilityZoneData

`func NewAvailabilityZoneData(code string, name string, ) *AvailabilityZoneData`

NewAvailabilityZoneData instantiates a new AvailabilityZoneData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityZoneDataWithDefaults

`func NewAvailabilityZoneDataWithDefaults() *AvailabilityZoneData`

NewAvailabilityZoneDataWithDefaults instantiates a new AvailabilityZoneData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AvailabilityZoneData) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AvailabilityZoneData) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AvailabilityZoneData) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *AvailabilityZoneData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AvailabilityZoneData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AvailabilityZoneData) SetName(v string)`

SetName sets Name field to given value.


### GetSecondarySubnet

`func (o *AvailabilityZoneData) GetSecondarySubnet() string`

GetSecondarySubnet returns the SecondarySubnet field if non-nil, zero value otherwise.

### GetSecondarySubnetOk

`func (o *AvailabilityZoneData) GetSecondarySubnetOk() (*string, bool)`

GetSecondarySubnetOk returns a tuple with the SecondarySubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySubnet

`func (o *AvailabilityZoneData) SetSecondarySubnet(v string)`

SetSecondarySubnet sets SecondarySubnet field to given value.

### HasSecondarySubnet

`func (o *AvailabilityZoneData) HasSecondarySubnet() bool`

HasSecondarySubnet returns a boolean if a field has been set.

### GetSubnet

`func (o *AvailabilityZoneData) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *AvailabilityZoneData) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *AvailabilityZoneData) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *AvailabilityZoneData) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


