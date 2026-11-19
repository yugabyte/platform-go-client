# AzureRegionLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | Region of the universe whose backups go to this location | 
**Location** | **string** | Location of the region&#39;s backups, as https://&lt;account&gt;.blob.core.windows.net/&lt;container&gt;[/&lt;path&gt;] | 
**SasToken** | Pointer to **string** | SAS token of this region&#39;s container, with SAS_TOKEN credentials. Required on create with those credentials; never returned. | [optional] 

## Methods

### NewAzureRegionLocation

`func NewAzureRegionLocation(region string, location string, ) *AzureRegionLocation`

NewAzureRegionLocation instantiates a new AzureRegionLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureRegionLocationWithDefaults

`func NewAzureRegionLocationWithDefaults() *AzureRegionLocation`

NewAzureRegionLocationWithDefaults instantiates a new AzureRegionLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *AzureRegionLocation) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AzureRegionLocation) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AzureRegionLocation) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetLocation

`func (o *AzureRegionLocation) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AzureRegionLocation) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AzureRegionLocation) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetSasToken

`func (o *AzureRegionLocation) GetSasToken() string`

GetSasToken returns the SasToken field if non-nil, zero value otherwise.

### GetSasTokenOk

`func (o *AzureRegionLocation) GetSasTokenOk() (*string, bool)`

GetSasTokenOk returns a tuple with the SasToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasToken

`func (o *AzureRegionLocation) SetSasToken(v string)`

SetSasToken sets SasToken field to given value.

### HasSasToken

`func (o *AzureRegionLocation) HasSasToken() bool`

HasSasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


