# UniverseSoftwareUpgradeStartAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowRollback** | Pointer to **bool** | perform an upgrade where rollback is allowed | [optional] [default to true]
**UpgradeSystemCatalog** | Pointer to **bool** | Upgrade the YugabyteDB Catalog | [optional] [default to true]
**Version** | **string** | The target release version to upgrade to. | 

## Methods

### NewUniverseSoftwareUpgradeStartAllOf

`func NewUniverseSoftwareUpgradeStartAllOf(version string, ) *UniverseSoftwareUpgradeStartAllOf`

NewUniverseSoftwareUpgradeStartAllOf instantiates a new UniverseSoftwareUpgradeStartAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseSoftwareUpgradeStartAllOfWithDefaults

`func NewUniverseSoftwareUpgradeStartAllOfWithDefaults() *UniverseSoftwareUpgradeStartAllOf`

NewUniverseSoftwareUpgradeStartAllOfWithDefaults instantiates a new UniverseSoftwareUpgradeStartAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowRollback

`func (o *UniverseSoftwareUpgradeStartAllOf) GetAllowRollback() bool`

GetAllowRollback returns the AllowRollback field if non-nil, zero value otherwise.

### GetAllowRollbackOk

`func (o *UniverseSoftwareUpgradeStartAllOf) GetAllowRollbackOk() (*bool, bool)`

GetAllowRollbackOk returns a tuple with the AllowRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRollback

`func (o *UniverseSoftwareUpgradeStartAllOf) SetAllowRollback(v bool)`

SetAllowRollback sets AllowRollback field to given value.

### HasAllowRollback

`func (o *UniverseSoftwareUpgradeStartAllOf) HasAllowRollback() bool`

HasAllowRollback returns a boolean if a field has been set.

### GetUpgradeSystemCatalog

`func (o *UniverseSoftwareUpgradeStartAllOf) GetUpgradeSystemCatalog() bool`

GetUpgradeSystemCatalog returns the UpgradeSystemCatalog field if non-nil, zero value otherwise.

### GetUpgradeSystemCatalogOk

`func (o *UniverseSoftwareUpgradeStartAllOf) GetUpgradeSystemCatalogOk() (*bool, bool)`

GetUpgradeSystemCatalogOk returns a tuple with the UpgradeSystemCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeSystemCatalog

`func (o *UniverseSoftwareUpgradeStartAllOf) SetUpgradeSystemCatalog(v bool)`

SetUpgradeSystemCatalog sets UpgradeSystemCatalog field to given value.

### HasUpgradeSystemCatalog

`func (o *UniverseSoftwareUpgradeStartAllOf) HasUpgradeSystemCatalog() bool`

HasUpgradeSystemCatalog returns a boolean if a field has been set.

### GetVersion

`func (o *UniverseSoftwareUpgradeStartAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UniverseSoftwareUpgradeStartAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UniverseSoftwareUpgradeStartAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


