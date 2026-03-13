# SoftwareUpgradeAZStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzUuid** | **string** | Availability zone UUID | 
**PauseAfterTserverUpgrade** | Pointer to **bool** | If true, upgrade pauses after all tservers in this AZ are upgraded | [optional] [default to false]

## Methods

### NewSoftwareUpgradeAZStep

`func NewSoftwareUpgradeAZStep(azUuid string, ) *SoftwareUpgradeAZStep`

NewSoftwareUpgradeAZStep instantiates a new SoftwareUpgradeAZStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareUpgradeAZStepWithDefaults

`func NewSoftwareUpgradeAZStepWithDefaults() *SoftwareUpgradeAZStep`

NewSoftwareUpgradeAZStepWithDefaults instantiates a new SoftwareUpgradeAZStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzUuid

`func (o *SoftwareUpgradeAZStep) GetAzUuid() string`

GetAzUuid returns the AzUuid field if non-nil, zero value otherwise.

### GetAzUuidOk

`func (o *SoftwareUpgradeAZStep) GetAzUuidOk() (*string, bool)`

GetAzUuidOk returns a tuple with the AzUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzUuid

`func (o *SoftwareUpgradeAZStep) SetAzUuid(v string)`

SetAzUuid sets AzUuid field to given value.


### GetPauseAfterTserverUpgrade

`func (o *SoftwareUpgradeAZStep) GetPauseAfterTserverUpgrade() bool`

GetPauseAfterTserverUpgrade returns the PauseAfterTserverUpgrade field if non-nil, zero value otherwise.

### GetPauseAfterTserverUpgradeOk

`func (o *SoftwareUpgradeAZStep) GetPauseAfterTserverUpgradeOk() (*bool, bool)`

GetPauseAfterTserverUpgradeOk returns a tuple with the PauseAfterTserverUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseAfterTserverUpgrade

`func (o *SoftwareUpgradeAZStep) SetPauseAfterTserverUpgrade(v bool)`

SetPauseAfterTserverUpgrade sets PauseAfterTserverUpgrade field to given value.

### HasPauseAfterTserverUpgrade

`func (o *SoftwareUpgradeAZStep) HasPauseAfterTserverUpgrade() bool`

HasPauseAfterTserverUpgrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


