# Region

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] [readonly] 
**Code** | Pointer to **string** | Cloud provider region code | [optional] 
**Config** | Pointer to **map[string]string** |  | [optional] 
**Details** | Pointer to [**RegionDetails**](RegionDetails.md) |  | [optional] 
**Latitude** | Pointer to **float64** | The region&#39;s latitude | [optional] [readonly] 
**Longitude** | Pointer to **float64** | The region&#39;s longitude | [optional] [readonly] 
**Name** | Pointer to **string** | Cloud provider region name | [optional] [readonly] 
**SecurityGroupId** | Pointer to **string** | &lt;b style&#x3D;\&quot;color:#ff0000\&quot;&gt;Deprecated since YBA version 2.17.2.0.&lt;/b&gt; Moved to regionDetails.cloudInfo aws/azure securityGroupId property | [optional] 
**Uuid** | Pointer to **string** | Region UUID | [optional] [readonly] 
**VnetName** | Pointer to **string** | &lt;b style&#x3D;\&quot;color:#ff0000\&quot;&gt;Deprecated since YBA version 2.17.2.0.&lt;/b&gt; Moved to regionDetails.cloudInfo aws/azure vnet property | [optional] 
**YbImage** | Pointer to **string** | &lt;b style&#x3D;\&quot;color:#ff0000\&quot;&gt;Deprecated since YBA version 2.17.2.0.&lt;/b&gt; Moved to details.cloudInfo aws/gcp/azure ybImage property | [optional] 
**Zones** | [**[]AvailabilityZone**](AvailabilityZone.md) |  | 

## Methods

### NewRegion

`func NewRegion(zones []AvailabilityZone, ) *Region`

NewRegion instantiates a new Region object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionWithDefaults

`func NewRegionWithDefaults() *Region`

NewRegionWithDefaults instantiates a new Region object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *Region) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Region) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Region) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Region) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCode

`func (o *Region) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Region) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Region) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Region) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetConfig

`func (o *Region) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Region) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Region) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Region) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDetails

`func (o *Region) GetDetails() RegionDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Region) GetDetailsOk() (*RegionDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Region) SetDetails(v RegionDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Region) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetLatitude

`func (o *Region) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *Region) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *Region) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *Region) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *Region) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *Region) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *Region) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *Region) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetName

`func (o *Region) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Region) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Region) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Region) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecurityGroupId

`func (o *Region) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *Region) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *Region) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.

### HasSecurityGroupId

`func (o *Region) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### GetUuid

`func (o *Region) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Region) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Region) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Region) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVnetName

`func (o *Region) GetVnetName() string`

GetVnetName returns the VnetName field if non-nil, zero value otherwise.

### GetVnetNameOk

`func (o *Region) GetVnetNameOk() (*string, bool)`

GetVnetNameOk returns a tuple with the VnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetName

`func (o *Region) SetVnetName(v string)`

SetVnetName sets VnetName field to given value.

### HasVnetName

`func (o *Region) HasVnetName() bool`

HasVnetName returns a boolean if a field has been set.

### GetYbImage

`func (o *Region) GetYbImage() string`

GetYbImage returns the YbImage field if non-nil, zero value otherwise.

### GetYbImageOk

`func (o *Region) GetYbImageOk() (*string, bool)`

GetYbImageOk returns a tuple with the YbImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYbImage

`func (o *Region) SetYbImage(v string)`

SetYbImage sets YbImage field to given value.

### HasYbImage

`func (o *Region) HasYbImage() bool`

HasYbImage returns a boolean if a field has been set.

### GetZones

`func (o *Region) GetZones() []AvailabilityZone`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *Region) GetZonesOk() (*[]AvailabilityZone, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *Region) SetZones(v []AvailabilityZone)`

SetZones sets Zones field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


