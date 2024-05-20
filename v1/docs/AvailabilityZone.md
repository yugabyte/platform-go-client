# AvailabilityZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | AZ status. This value is &#x60;true&#x60; for an active AZ. | [optional] [readonly] 
**Code** | Pointer to **string** | AZ code | [optional] 
**Config** | Pointer to **map[string]string** | AZ configuration values | [optional] 
**Details** | Pointer to [**AvailabilityZoneDetails**](AvailabilityZoneDetails.md) |  | [optional] 
**KubeconfigPath** | Pointer to **string** | Path to Kubernetes configuration file | [optional] [readonly] 
**Name** | **string** | AZ name | 
**SecondarySubnet** | Pointer to **string** | AZ secondary subnet | [optional] 
**Subnet** | Pointer to **string** | AZ subnet | [optional] 
**Uuid** | Pointer to **string** | AZ UUID | [optional] [readonly] 

## Methods

### NewAvailabilityZone

`func NewAvailabilityZone(name string, ) *AvailabilityZone`

NewAvailabilityZone instantiates a new AvailabilityZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityZoneWithDefaults

`func NewAvailabilityZoneWithDefaults() *AvailabilityZone`

NewAvailabilityZoneWithDefaults instantiates a new AvailabilityZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *AvailabilityZone) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AvailabilityZone) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AvailabilityZone) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AvailabilityZone) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCode

`func (o *AvailabilityZone) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AvailabilityZone) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AvailabilityZone) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AvailabilityZone) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetConfig

`func (o *AvailabilityZone) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AvailabilityZone) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AvailabilityZone) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AvailabilityZone) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDetails

`func (o *AvailabilityZone) GetDetails() AvailabilityZoneDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AvailabilityZone) GetDetailsOk() (*AvailabilityZoneDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AvailabilityZone) SetDetails(v AvailabilityZoneDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AvailabilityZone) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetKubeconfigPath

`func (o *AvailabilityZone) GetKubeconfigPath() string`

GetKubeconfigPath returns the KubeconfigPath field if non-nil, zero value otherwise.

### GetKubeconfigPathOk

`func (o *AvailabilityZone) GetKubeconfigPathOk() (*string, bool)`

GetKubeconfigPathOk returns a tuple with the KubeconfigPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeconfigPath

`func (o *AvailabilityZone) SetKubeconfigPath(v string)`

SetKubeconfigPath sets KubeconfigPath field to given value.

### HasKubeconfigPath

`func (o *AvailabilityZone) HasKubeconfigPath() bool`

HasKubeconfigPath returns a boolean if a field has been set.

### GetName

`func (o *AvailabilityZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AvailabilityZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AvailabilityZone) SetName(v string)`

SetName sets Name field to given value.


### GetSecondarySubnet

`func (o *AvailabilityZone) GetSecondarySubnet() string`

GetSecondarySubnet returns the SecondarySubnet field if non-nil, zero value otherwise.

### GetSecondarySubnetOk

`func (o *AvailabilityZone) GetSecondarySubnetOk() (*string, bool)`

GetSecondarySubnetOk returns a tuple with the SecondarySubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySubnet

`func (o *AvailabilityZone) SetSecondarySubnet(v string)`

SetSecondarySubnet sets SecondarySubnet field to given value.

### HasSecondarySubnet

`func (o *AvailabilityZone) HasSecondarySubnet() bool`

HasSecondarySubnet returns a boolean if a field has been set.

### GetSubnet

`func (o *AvailabilityZone) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *AvailabilityZone) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *AvailabilityZone) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *AvailabilityZone) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetUuid

`func (o *AvailabilityZone) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AvailabilityZone) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AvailabilityZone) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AvailabilityZone) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


