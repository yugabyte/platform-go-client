# SoftwareUpgradeProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanaryUpgrade** | Pointer to **bool** | Whether this software upgrade is using a canary rollout. | [optional] [readonly] 
**CanaryPauseState** | Pointer to **string** | Pause point reached during a canary software upgrade, if applicable. | [optional] [readonly] 
**MasterAzUpgradeStatesList** | Pointer to [**[]AZUpgradeState**](AZUpgradeState.md) | Per-AZ upgrade status for master processes. | [optional] [readonly] 
**TserverAzUpgradeStatesList** | Pointer to [**[]AZUpgradeState**](AZUpgradeState.md) | Per-AZ upgrade status for tserver processes. | [optional] [readonly] 

## Methods

### NewSoftwareUpgradeProgress

`func NewSoftwareUpgradeProgress() *SoftwareUpgradeProgress`

NewSoftwareUpgradeProgress instantiates a new SoftwareUpgradeProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareUpgradeProgressWithDefaults

`func NewSoftwareUpgradeProgressWithDefaults() *SoftwareUpgradeProgress`

NewSoftwareUpgradeProgressWithDefaults instantiates a new SoftwareUpgradeProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanaryUpgrade

`func (o *SoftwareUpgradeProgress) GetCanaryUpgrade() bool`

GetCanaryUpgrade returns the CanaryUpgrade field if non-nil, zero value otherwise.

### GetCanaryUpgradeOk

`func (o *SoftwareUpgradeProgress) GetCanaryUpgradeOk() (*bool, bool)`

GetCanaryUpgradeOk returns a tuple with the CanaryUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanaryUpgrade

`func (o *SoftwareUpgradeProgress) SetCanaryUpgrade(v bool)`

SetCanaryUpgrade sets CanaryUpgrade field to given value.

### HasCanaryUpgrade

`func (o *SoftwareUpgradeProgress) HasCanaryUpgrade() bool`

HasCanaryUpgrade returns a boolean if a field has been set.

### GetCanaryPauseState

`func (o *SoftwareUpgradeProgress) GetCanaryPauseState() string`

GetCanaryPauseState returns the CanaryPauseState field if non-nil, zero value otherwise.

### GetCanaryPauseStateOk

`func (o *SoftwareUpgradeProgress) GetCanaryPauseStateOk() (*string, bool)`

GetCanaryPauseStateOk returns a tuple with the CanaryPauseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanaryPauseState

`func (o *SoftwareUpgradeProgress) SetCanaryPauseState(v string)`

SetCanaryPauseState sets CanaryPauseState field to given value.

### HasCanaryPauseState

`func (o *SoftwareUpgradeProgress) HasCanaryPauseState() bool`

HasCanaryPauseState returns a boolean if a field has been set.

### GetMasterAzUpgradeStatesList

`func (o *SoftwareUpgradeProgress) GetMasterAzUpgradeStatesList() []AZUpgradeState`

GetMasterAzUpgradeStatesList returns the MasterAzUpgradeStatesList field if non-nil, zero value otherwise.

### GetMasterAzUpgradeStatesListOk

`func (o *SoftwareUpgradeProgress) GetMasterAzUpgradeStatesListOk() (*[]AZUpgradeState, bool)`

GetMasterAzUpgradeStatesListOk returns a tuple with the MasterAzUpgradeStatesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterAzUpgradeStatesList

`func (o *SoftwareUpgradeProgress) SetMasterAzUpgradeStatesList(v []AZUpgradeState)`

SetMasterAzUpgradeStatesList sets MasterAzUpgradeStatesList field to given value.

### HasMasterAzUpgradeStatesList

`func (o *SoftwareUpgradeProgress) HasMasterAzUpgradeStatesList() bool`

HasMasterAzUpgradeStatesList returns a boolean if a field has been set.

### GetTserverAzUpgradeStatesList

`func (o *SoftwareUpgradeProgress) GetTserverAzUpgradeStatesList() []AZUpgradeState`

GetTserverAzUpgradeStatesList returns the TserverAzUpgradeStatesList field if non-nil, zero value otherwise.

### GetTserverAzUpgradeStatesListOk

`func (o *SoftwareUpgradeProgress) GetTserverAzUpgradeStatesListOk() (*[]AZUpgradeState, bool)`

GetTserverAzUpgradeStatesListOk returns a tuple with the TserverAzUpgradeStatesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTserverAzUpgradeStatesList

`func (o *SoftwareUpgradeProgress) SetTserverAzUpgradeStatesList(v []AZUpgradeState)`

SetTserverAzUpgradeStatesList sets TserverAzUpgradeStatesList field to given value.

### HasTserverAzUpgradeStatesList

`func (o *SoftwareUpgradeProgress) HasTserverAzUpgradeStatesList() bool`

HasTserverAzUpgradeStatesList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


