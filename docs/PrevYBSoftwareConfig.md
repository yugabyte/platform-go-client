# PrevYBSoftwareConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllTserversUpgradedToYsqlMajorVersion** | Pointer to **bool** |  | [optional] 
**AutoFlagConfigVersion** | Pointer to **int32** |  | [optional] 
**CanRollbackCatalogUpgrade** | Pointer to **bool** |  | [optional] 
**CanaryPauseState** | Pointer to **string** | WARNING: This is a preview API that could change. Canary pause state when upgrade is paused at a canary point | [optional] 
**CanaryUpgrade** | **bool** |  | 
**MasterAZUpgradeStatesList** | Pointer to [**[]AZUpgradeState**](AZUpgradeState.md) | WARNING: This is a preview API that could change. Per-AZ master upgrade progress (standard and canary) | [optional] 
**MasterPauseCompleted** | Pointer to **bool** | WARNING: This is a preview API that could change. True once the canary pauseAfterMasters checkpoint has been reached and resumed, so it is not re-emitted on a subsequent abort+retry of the upgrade. | [optional] 
**SoftwareVersion** | Pointer to **string** |  | [optional] 
**TargetUpgradeSoftwareVersion** | Pointer to **string** |  | [optional] 
**TserverAZUpgradeStatesList** | Pointer to [**[]AZUpgradeState**](AZUpgradeState.md) | WARNING: This is a preview API that could change. Per-AZ tserver upgrade progress (standard and canary) | [optional] 

## Methods

### NewPrevYBSoftwareConfig

`func NewPrevYBSoftwareConfig(canaryUpgrade bool, ) *PrevYBSoftwareConfig`

NewPrevYBSoftwareConfig instantiates a new PrevYBSoftwareConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrevYBSoftwareConfigWithDefaults

`func NewPrevYBSoftwareConfigWithDefaults() *PrevYBSoftwareConfig`

NewPrevYBSoftwareConfigWithDefaults instantiates a new PrevYBSoftwareConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllTserversUpgradedToYsqlMajorVersion

`func (o *PrevYBSoftwareConfig) GetAllTserversUpgradedToYsqlMajorVersion() bool`

GetAllTserversUpgradedToYsqlMajorVersion returns the AllTserversUpgradedToYsqlMajorVersion field if non-nil, zero value otherwise.

### GetAllTserversUpgradedToYsqlMajorVersionOk

`func (o *PrevYBSoftwareConfig) GetAllTserversUpgradedToYsqlMajorVersionOk() (*bool, bool)`

GetAllTserversUpgradedToYsqlMajorVersionOk returns a tuple with the AllTserversUpgradedToYsqlMajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTserversUpgradedToYsqlMajorVersion

`func (o *PrevYBSoftwareConfig) SetAllTserversUpgradedToYsqlMajorVersion(v bool)`

SetAllTserversUpgradedToYsqlMajorVersion sets AllTserversUpgradedToYsqlMajorVersion field to given value.

### HasAllTserversUpgradedToYsqlMajorVersion

`func (o *PrevYBSoftwareConfig) HasAllTserversUpgradedToYsqlMajorVersion() bool`

HasAllTserversUpgradedToYsqlMajorVersion returns a boolean if a field has been set.

### GetAutoFlagConfigVersion

`func (o *PrevYBSoftwareConfig) GetAutoFlagConfigVersion() int32`

GetAutoFlagConfigVersion returns the AutoFlagConfigVersion field if non-nil, zero value otherwise.

### GetAutoFlagConfigVersionOk

`func (o *PrevYBSoftwareConfig) GetAutoFlagConfigVersionOk() (*int32, bool)`

GetAutoFlagConfigVersionOk returns a tuple with the AutoFlagConfigVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoFlagConfigVersion

`func (o *PrevYBSoftwareConfig) SetAutoFlagConfigVersion(v int32)`

SetAutoFlagConfigVersion sets AutoFlagConfigVersion field to given value.

### HasAutoFlagConfigVersion

`func (o *PrevYBSoftwareConfig) HasAutoFlagConfigVersion() bool`

HasAutoFlagConfigVersion returns a boolean if a field has been set.

### GetCanRollbackCatalogUpgrade

`func (o *PrevYBSoftwareConfig) GetCanRollbackCatalogUpgrade() bool`

GetCanRollbackCatalogUpgrade returns the CanRollbackCatalogUpgrade field if non-nil, zero value otherwise.

### GetCanRollbackCatalogUpgradeOk

`func (o *PrevYBSoftwareConfig) GetCanRollbackCatalogUpgradeOk() (*bool, bool)`

GetCanRollbackCatalogUpgradeOk returns a tuple with the CanRollbackCatalogUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRollbackCatalogUpgrade

`func (o *PrevYBSoftwareConfig) SetCanRollbackCatalogUpgrade(v bool)`

SetCanRollbackCatalogUpgrade sets CanRollbackCatalogUpgrade field to given value.

### HasCanRollbackCatalogUpgrade

`func (o *PrevYBSoftwareConfig) HasCanRollbackCatalogUpgrade() bool`

HasCanRollbackCatalogUpgrade returns a boolean if a field has been set.

### GetCanaryPauseState

`func (o *PrevYBSoftwareConfig) GetCanaryPauseState() string`

GetCanaryPauseState returns the CanaryPauseState field if non-nil, zero value otherwise.

### GetCanaryPauseStateOk

`func (o *PrevYBSoftwareConfig) GetCanaryPauseStateOk() (*string, bool)`

GetCanaryPauseStateOk returns a tuple with the CanaryPauseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanaryPauseState

`func (o *PrevYBSoftwareConfig) SetCanaryPauseState(v string)`

SetCanaryPauseState sets CanaryPauseState field to given value.

### HasCanaryPauseState

`func (o *PrevYBSoftwareConfig) HasCanaryPauseState() bool`

HasCanaryPauseState returns a boolean if a field has been set.

### GetCanaryUpgrade

`func (o *PrevYBSoftwareConfig) GetCanaryUpgrade() bool`

GetCanaryUpgrade returns the CanaryUpgrade field if non-nil, zero value otherwise.

### GetCanaryUpgradeOk

`func (o *PrevYBSoftwareConfig) GetCanaryUpgradeOk() (*bool, bool)`

GetCanaryUpgradeOk returns a tuple with the CanaryUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanaryUpgrade

`func (o *PrevYBSoftwareConfig) SetCanaryUpgrade(v bool)`

SetCanaryUpgrade sets CanaryUpgrade field to given value.


### GetMasterAZUpgradeStatesList

`func (o *PrevYBSoftwareConfig) GetMasterAZUpgradeStatesList() []AZUpgradeState`

GetMasterAZUpgradeStatesList returns the MasterAZUpgradeStatesList field if non-nil, zero value otherwise.

### GetMasterAZUpgradeStatesListOk

`func (o *PrevYBSoftwareConfig) GetMasterAZUpgradeStatesListOk() (*[]AZUpgradeState, bool)`

GetMasterAZUpgradeStatesListOk returns a tuple with the MasterAZUpgradeStatesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterAZUpgradeStatesList

`func (o *PrevYBSoftwareConfig) SetMasterAZUpgradeStatesList(v []AZUpgradeState)`

SetMasterAZUpgradeStatesList sets MasterAZUpgradeStatesList field to given value.

### HasMasterAZUpgradeStatesList

`func (o *PrevYBSoftwareConfig) HasMasterAZUpgradeStatesList() bool`

HasMasterAZUpgradeStatesList returns a boolean if a field has been set.

### GetMasterPauseCompleted

`func (o *PrevYBSoftwareConfig) GetMasterPauseCompleted() bool`

GetMasterPauseCompleted returns the MasterPauseCompleted field if non-nil, zero value otherwise.

### GetMasterPauseCompletedOk

`func (o *PrevYBSoftwareConfig) GetMasterPauseCompletedOk() (*bool, bool)`

GetMasterPauseCompletedOk returns a tuple with the MasterPauseCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterPauseCompleted

`func (o *PrevYBSoftwareConfig) SetMasterPauseCompleted(v bool)`

SetMasterPauseCompleted sets MasterPauseCompleted field to given value.

### HasMasterPauseCompleted

`func (o *PrevYBSoftwareConfig) HasMasterPauseCompleted() bool`

HasMasterPauseCompleted returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *PrevYBSoftwareConfig) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *PrevYBSoftwareConfig) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *PrevYBSoftwareConfig) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *PrevYBSoftwareConfig) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetTargetUpgradeSoftwareVersion

`func (o *PrevYBSoftwareConfig) GetTargetUpgradeSoftwareVersion() string`

GetTargetUpgradeSoftwareVersion returns the TargetUpgradeSoftwareVersion field if non-nil, zero value otherwise.

### GetTargetUpgradeSoftwareVersionOk

`func (o *PrevYBSoftwareConfig) GetTargetUpgradeSoftwareVersionOk() (*string, bool)`

GetTargetUpgradeSoftwareVersionOk returns a tuple with the TargetUpgradeSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUpgradeSoftwareVersion

`func (o *PrevYBSoftwareConfig) SetTargetUpgradeSoftwareVersion(v string)`

SetTargetUpgradeSoftwareVersion sets TargetUpgradeSoftwareVersion field to given value.

### HasTargetUpgradeSoftwareVersion

`func (o *PrevYBSoftwareConfig) HasTargetUpgradeSoftwareVersion() bool`

HasTargetUpgradeSoftwareVersion returns a boolean if a field has been set.

### GetTserverAZUpgradeStatesList

`func (o *PrevYBSoftwareConfig) GetTserverAZUpgradeStatesList() []AZUpgradeState`

GetTserverAZUpgradeStatesList returns the TserverAZUpgradeStatesList field if non-nil, zero value otherwise.

### GetTserverAZUpgradeStatesListOk

`func (o *PrevYBSoftwareConfig) GetTserverAZUpgradeStatesListOk() (*[]AZUpgradeState, bool)`

GetTserverAZUpgradeStatesListOk returns a tuple with the TserverAZUpgradeStatesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverAZUpgradeStatesList

`func (o *PrevYBSoftwareConfig) SetTserverAZUpgradeStatesList(v []AZUpgradeState)`

SetTserverAZUpgradeStatesList sets TserverAZUpgradeStatesList field to given value.

### HasTserverAZUpgradeStatesList

`func (o *PrevYBSoftwareConfig) HasTserverAZUpgradeStatesList() bool`

HasTserverAZUpgradeStatesList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


