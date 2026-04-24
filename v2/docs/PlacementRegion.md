# PlacementRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | The region provider id. | 
**Code** | **string** | The actual provider given region code. | 
**Name** | Pointer to **string** | The region name. | [optional] 
**AzList** | [**[]PlacementAZ**](PlacementAZ.md) | The list of AZs inside this region into which we want to place data. | 
**LbFqdn** | Pointer to **string** | The Load Balancer FQDN. | [optional] 

## Methods

### NewPlacementRegion

`func NewPlacementRegion(uuid string, code string, azList []PlacementAZ, ) *PlacementRegion`

NewPlacementRegion instantiates a new PlacementRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlacementRegionWithDefaults

`func NewPlacementRegionWithDefaults() *PlacementRegion`

NewPlacementRegionWithDefaults instantiates a new PlacementRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PlacementRegion) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PlacementRegion) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PlacementRegion) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetCode

`func (o *PlacementRegion) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PlacementRegion) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PlacementRegion) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *PlacementRegion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlacementRegion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlacementRegion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlacementRegion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAzList

`func (o *PlacementRegion) GetAzList() []PlacementAZ`

GetAzList returns the AzList field if non-nil, zero value otherwise.

### GetAzListOk

`func (o *PlacementRegion) GetAzListOk() (*[]PlacementAZ, bool)`

GetAzListOk returns a tuple with the AzList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzList

`func (o *PlacementRegion) SetAzList(v []PlacementAZ)`

SetAzList sets AzList field to given value.


### GetLbFqdn

`func (o *PlacementRegion) GetLbFqdn() string`

GetLbFqdn returns the LbFqdn field if non-nil, zero value otherwise.

### GetLbFqdnOk

`func (o *PlacementRegion) GetLbFqdnOk() (*string, bool)`

GetLbFqdnOk returns a tuple with the LbFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbFqdn

`func (o *PlacementRegion) SetLbFqdn(v string)`

SetLbFqdn sets LbFqdn field to given value.

### HasLbFqdn

`func (o *PlacementRegion) HasLbFqdn() bool`

HasLbFqdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


