# AZUpgradeState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzUuid** | Pointer to **string** | UUID of the availability zone. | [optional] [readonly] 
**AzName** | Pointer to **string** | Name of the availability zone. | [optional] [readonly] 
**ServerType** | Pointer to **string** | Role of the YugabyteDB server being upgraded in this zone. | [optional] [readonly] 
**ClusterUuid** | Pointer to **string** | UUID of the cluster being upgraded. | [optional] [readonly] 
**Status** | Pointer to **string** | Upgrade status for this availability zone and server role. | [optional] [readonly] 

## Methods

### NewAZUpgradeState

`func NewAZUpgradeState() *AZUpgradeState`

NewAZUpgradeState instantiates a new AZUpgradeState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAZUpgradeStateWithDefaults

`func NewAZUpgradeStateWithDefaults() *AZUpgradeState`

NewAZUpgradeStateWithDefaults instantiates a new AZUpgradeState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzUuid

`func (o *AZUpgradeState) GetAzUuid() string`

GetAzUuid returns the AzUuid field if non-nil, zero value otherwise.

### GetAzUuidOk

`func (o *AZUpgradeState) GetAzUuidOk() (*string, bool)`

GetAzUuidOk returns a tuple with the AzUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzUuid

`func (o *AZUpgradeState) SetAzUuid(v string)`

SetAzUuid sets AzUuid field to given value.

### HasAzUuid

`func (o *AZUpgradeState) HasAzUuid() bool`

HasAzUuid returns a boolean if a field has been set.

### GetAzName

`func (o *AZUpgradeState) GetAzName() string`

GetAzName returns the AzName field if non-nil, zero value otherwise.

### GetAzNameOk

`func (o *AZUpgradeState) GetAzNameOk() (*string, bool)`

GetAzNameOk returns a tuple with the AzName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzName

`func (o *AZUpgradeState) SetAzName(v string)`

SetAzName sets AzName field to given value.

### HasAzName

`func (o *AZUpgradeState) HasAzName() bool`

HasAzName returns a boolean if a field has been set.

### GetServerType

`func (o *AZUpgradeState) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *AZUpgradeState) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *AZUpgradeState) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *AZUpgradeState) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetClusterUuid

`func (o *AZUpgradeState) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *AZUpgradeState) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *AZUpgradeState) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *AZUpgradeState) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetStatus

`func (o *AZUpgradeState) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AZUpgradeState) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AZUpgradeState) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AZUpgradeState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


