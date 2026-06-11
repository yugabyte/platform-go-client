# PitrConfigSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | PITR config name | [optional] 
**DbName** | Pointer to **string** | Database or keyspace name | [optional] 
**TableType** | Pointer to **string** | Table type | [optional] 
**ScheduleInterval** | Pointer to **int64** | Interval between snapshots in seconds | [optional] 
**RetentionPeriod** | Pointer to **int64** | Retention period in seconds | [optional] 
**IntermittentMinRecoverTimeInMillis** | Pointer to **int64** | Intermittent minimum recovery time in milliseconds, used when the retention period is increased. | [optional] 

## Methods

### NewPitrConfigSpec

`func NewPitrConfigSpec() *PitrConfigSpec`

NewPitrConfigSpec instantiates a new PitrConfigSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPitrConfigSpecWithDefaults

`func NewPitrConfigSpecWithDefaults() *PitrConfigSpec`

NewPitrConfigSpecWithDefaults instantiates a new PitrConfigSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PitrConfigSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PitrConfigSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PitrConfigSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PitrConfigSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDbName

`func (o *PitrConfigSpec) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *PitrConfigSpec) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *PitrConfigSpec) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *PitrConfigSpec) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetTableType

`func (o *PitrConfigSpec) GetTableType() string`

GetTableType returns the TableType field if non-nil, zero value otherwise.

### GetTableTypeOk

`func (o *PitrConfigSpec) GetTableTypeOk() (*string, bool)`

GetTableTypeOk returns a tuple with the TableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableType

`func (o *PitrConfigSpec) SetTableType(v string)`

SetTableType sets TableType field to given value.

### HasTableType

`func (o *PitrConfigSpec) HasTableType() bool`

HasTableType returns a boolean if a field has been set.

### GetScheduleInterval

`func (o *PitrConfigSpec) GetScheduleInterval() int64`

GetScheduleInterval returns the ScheduleInterval field if non-nil, zero value otherwise.

### GetScheduleIntervalOk

`func (o *PitrConfigSpec) GetScheduleIntervalOk() (*int64, bool)`

GetScheduleIntervalOk returns a tuple with the ScheduleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleInterval

`func (o *PitrConfigSpec) SetScheduleInterval(v int64)`

SetScheduleInterval sets ScheduleInterval field to given value.

### HasScheduleInterval

`func (o *PitrConfigSpec) HasScheduleInterval() bool`

HasScheduleInterval returns a boolean if a field has been set.

### GetRetentionPeriod

`func (o *PitrConfigSpec) GetRetentionPeriod() int64`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *PitrConfigSpec) GetRetentionPeriodOk() (*int64, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *PitrConfigSpec) SetRetentionPeriod(v int64)`

SetRetentionPeriod sets RetentionPeriod field to given value.

### HasRetentionPeriod

`func (o *PitrConfigSpec) HasRetentionPeriod() bool`

HasRetentionPeriod returns a boolean if a field has been set.

### GetIntermittentMinRecoverTimeInMillis

`func (o *PitrConfigSpec) GetIntermittentMinRecoverTimeInMillis() int64`

GetIntermittentMinRecoverTimeInMillis returns the IntermittentMinRecoverTimeInMillis field if non-nil, zero value otherwise.

### GetIntermittentMinRecoverTimeInMillisOk

`func (o *PitrConfigSpec) GetIntermittentMinRecoverTimeInMillisOk() (*int64, bool)`

GetIntermittentMinRecoverTimeInMillisOk returns a tuple with the IntermittentMinRecoverTimeInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermittentMinRecoverTimeInMillis

`func (o *PitrConfigSpec) SetIntermittentMinRecoverTimeInMillis(v int64)`

SetIntermittentMinRecoverTimeInMillis sets IntermittentMinRecoverTimeInMillis field to given value.

### HasIntermittentMinRecoverTimeInMillis

`func (o *PitrConfigSpec) HasIntermittentMinRecoverTimeInMillis() bool`

HasIntermittentMinRecoverTimeInMillis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


