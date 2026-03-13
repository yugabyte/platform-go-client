# UniverseSoftwareUpgradePrecheckResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinalizeRequired** | **bool** | If the upgrade requires a finalize step. If true, the user must call the finalize  endpoint to complete the upgrade.  | 
**YsqlMajorVersionUpgrade** | Pointer to **bool** | If the upgrade requires a YSQL catalog major version upgrade (e.g. when upgrading across YSQL major versions and YSQL is enabled on the universe).  | [optional] 

## Methods

### NewUniverseSoftwareUpgradePrecheckResp

`func NewUniverseSoftwareUpgradePrecheckResp(finalizeRequired bool, ) *UniverseSoftwareUpgradePrecheckResp`

NewUniverseSoftwareUpgradePrecheckResp instantiates a new UniverseSoftwareUpgradePrecheckResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniverseSoftwareUpgradePrecheckRespWithDefaults

`func NewUniverseSoftwareUpgradePrecheckRespWithDefaults() *UniverseSoftwareUpgradePrecheckResp`

NewUniverseSoftwareUpgradePrecheckRespWithDefaults instantiates a new UniverseSoftwareUpgradePrecheckResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinalizeRequired

`func (o *UniverseSoftwareUpgradePrecheckResp) GetFinalizeRequired() bool`

GetFinalizeRequired returns the FinalizeRequired field if non-nil, zero value otherwise.

### GetFinalizeRequiredOk

`func (o *UniverseSoftwareUpgradePrecheckResp) GetFinalizeRequiredOk() (*bool, bool)`

GetFinalizeRequiredOk returns a tuple with the FinalizeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizeRequired

`func (o *UniverseSoftwareUpgradePrecheckResp) SetFinalizeRequired(v bool)`

SetFinalizeRequired sets FinalizeRequired field to given value.


### GetYsqlMajorVersionUpgrade

`func (o *UniverseSoftwareUpgradePrecheckResp) GetYsqlMajorVersionUpgrade() bool`

GetYsqlMajorVersionUpgrade returns the YsqlMajorVersionUpgrade field if non-nil, zero value otherwise.

### GetYsqlMajorVersionUpgradeOk

`func (o *UniverseSoftwareUpgradePrecheckResp) GetYsqlMajorVersionUpgradeOk() (*bool, bool)`

GetYsqlMajorVersionUpgradeOk returns a tuple with the YsqlMajorVersionUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlMajorVersionUpgrade

`func (o *UniverseSoftwareUpgradePrecheckResp) SetYsqlMajorVersionUpgrade(v bool)`

SetYsqlMajorVersionUpgrade sets YsqlMajorVersionUpgrade field to given value.

### HasYsqlMajorVersionUpgrade

`func (o *UniverseSoftwareUpgradePrecheckResp) HasYsqlMajorVersionUpgrade() bool`

HasYsqlMajorVersionUpgrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


