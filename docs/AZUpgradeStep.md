# AZUpgradeStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzUUID** | **string** | WARNING: This is a preview API that could change. Availability zone UUID | 
**PauseAfterTserverUpgrade** | Pointer to **bool** | WARNING: This is a preview API that could change. If true, upgrade pauses after all tservers in this AZ are upgraded | [optional] 

## Methods

### NewAZUpgradeStep

`func NewAZUpgradeStep(azUUID string, ) *AZUpgradeStep`

NewAZUpgradeStep instantiates a new AZUpgradeStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAZUpgradeStepWithDefaults

`func NewAZUpgradeStepWithDefaults() *AZUpgradeStep`

NewAZUpgradeStepWithDefaults instantiates a new AZUpgradeStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzUUID

`func (o *AZUpgradeStep) GetAzUUID() string`

GetAzUUID returns the AzUUID field if non-nil, zero value otherwise.

### GetAzUUIDOk

`func (o *AZUpgradeStep) GetAzUUIDOk() (*string, bool)`

GetAzUUIDOk returns a tuple with the AzUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzUUID

`func (o *AZUpgradeStep) SetAzUUID(v string)`

SetAzUUID sets AzUUID field to given value.


### GetPauseAfterTserverUpgrade

`func (o *AZUpgradeStep) GetPauseAfterTserverUpgrade() bool`

GetPauseAfterTserverUpgrade returns the PauseAfterTserverUpgrade field if non-nil, zero value otherwise.

### GetPauseAfterTserverUpgradeOk

`func (o *AZUpgradeStep) GetPauseAfterTserverUpgradeOk() (*bool, bool)`

GetPauseAfterTserverUpgradeOk returns a tuple with the PauseAfterTserverUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseAfterTserverUpgrade

`func (o *AZUpgradeStep) SetPauseAfterTserverUpgrade(v bool)`

SetPauseAfterTserverUpgrade sets PauseAfterTserverUpgrade field to given value.

### HasPauseAfterTserverUpgrade

`func (o *AZUpgradeStep) HasPauseAfterTserverUpgrade() bool`

HasPauseAfterTserverUpgrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


