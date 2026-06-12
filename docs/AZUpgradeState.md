# AZUpgradeState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzName** | Pointer to **string** | WARNING: This is a preview API that could change. Availability zone name for this upgrade state entry | [optional] 
**AzUUID** | Pointer to **string** | WARNING: This is a preview API that could change. Availability zone UUID for this upgrade state entry | [optional] 
**ClusterUUID** | Pointer to **string** | WARNING: This is a preview API that could change. Cluster UUID (primary or read replica) for this AZ | [optional] 
**ServerType** | Pointer to **string** | WARNING: This is a preview API that could change. Server type (MASTER or TSERVER) for this AZ | [optional] 
**Status** | Pointer to **string** | WARNING: This is a preview API that could change. Per-AZ upgrade status: NOT_STARTED; IN_PROGRESS before that AZ&#39;s upgrade subtasks run, then COMPLETED when done, or FAILED on task failure / YBA restart; COMPLETED; FAILED (subtask failure or stale IN_PROGRESS after upgrade failure / platform restart) | [optional] 

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

### GetAzUUID

`func (o *AZUpgradeState) GetAzUUID() string`

GetAzUUID returns the AzUUID field if non-nil, zero value otherwise.

### GetAzUUIDOk

`func (o *AZUpgradeState) GetAzUUIDOk() (*string, bool)`

GetAzUUIDOk returns a tuple with the AzUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzUUID

`func (o *AZUpgradeState) SetAzUUID(v string)`

SetAzUUID sets AzUUID field to given value.

### HasAzUUID

`func (o *AZUpgradeState) HasAzUUID() bool`

HasAzUUID returns a boolean if a field has been set.

### GetClusterUUID

`func (o *AZUpgradeState) GetClusterUUID() string`

GetClusterUUID returns the ClusterUUID field if non-nil, zero value otherwise.

### GetClusterUUIDOk

`func (o *AZUpgradeState) GetClusterUUIDOk() (*string, bool)`

GetClusterUUIDOk returns a tuple with the ClusterUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUUID

`func (o *AZUpgradeState) SetClusterUUID(v string)`

SetClusterUUID sets ClusterUUID field to given value.

### HasClusterUUID

`func (o *AZUpgradeState) HasClusterUUID() bool`

HasClusterUUID returns a boolean if a field has been set.

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


