# SoftwareUpgradeInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FinalizeRequired** | Pointer to **bool** | WARNING: This is a preview API that could change. Finalize required | [optional] [readonly] 
**YsqlMajorVersionUpgrade** | Pointer to **bool** | WARNING: This is a preview API that could change. YSQL Major version upgrade | [optional] [readonly] 

## Methods

### NewSoftwareUpgradeInfoResponse

`func NewSoftwareUpgradeInfoResponse() *SoftwareUpgradeInfoResponse`

NewSoftwareUpgradeInfoResponse instantiates a new SoftwareUpgradeInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareUpgradeInfoResponseWithDefaults

`func NewSoftwareUpgradeInfoResponseWithDefaults() *SoftwareUpgradeInfoResponse`

NewSoftwareUpgradeInfoResponseWithDefaults instantiates a new SoftwareUpgradeInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinalizeRequired

`func (o *SoftwareUpgradeInfoResponse) GetFinalizeRequired() bool`

GetFinalizeRequired returns the FinalizeRequired field if non-nil, zero value otherwise.

### GetFinalizeRequiredOk

`func (o *SoftwareUpgradeInfoResponse) GetFinalizeRequiredOk() (*bool, bool)`

GetFinalizeRequiredOk returns a tuple with the FinalizeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizeRequired

`func (o *SoftwareUpgradeInfoResponse) SetFinalizeRequired(v bool)`

SetFinalizeRequired sets FinalizeRequired field to given value.

### HasFinalizeRequired

`func (o *SoftwareUpgradeInfoResponse) HasFinalizeRequired() bool`

HasFinalizeRequired returns a boolean if a field has been set.

### GetYsqlMajorVersionUpgrade

`func (o *SoftwareUpgradeInfoResponse) GetYsqlMajorVersionUpgrade() bool`

GetYsqlMajorVersionUpgrade returns the YsqlMajorVersionUpgrade field if non-nil, zero value otherwise.

### GetYsqlMajorVersionUpgradeOk

`func (o *SoftwareUpgradeInfoResponse) GetYsqlMajorVersionUpgradeOk() (*bool, bool)`

GetYsqlMajorVersionUpgradeOk returns a tuple with the YsqlMajorVersionUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYsqlMajorVersionUpgrade

`func (o *SoftwareUpgradeInfoResponse) SetYsqlMajorVersionUpgrade(v bool)`

SetYsqlMajorVersionUpgrade sets YsqlMajorVersionUpgrade field to given value.

### HasYsqlMajorVersionUpgrade

`func (o *SoftwareUpgradeInfoResponse) HasYsqlMajorVersionUpgrade() bool`

HasYsqlMajorVersionUpgrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


