# PitrParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetentionPeriodSec** | Pointer to **int64** | Retention period of a snapshot in seconds | [optional] 
**SnapshotIntervalSec** | Pointer to **int64** | &lt;b style&#x3D;\&quot;color:#ff0000\&quot;&gt;Deprecated since YBA version 2024.2.0.0.&lt;/b&gt; Time interval between snapshots in seconds | [optional] 

## Methods

### NewPitrParams

`func NewPitrParams() *PitrParams`

NewPitrParams instantiates a new PitrParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPitrParamsWithDefaults

`func NewPitrParamsWithDefaults() *PitrParams`

NewPitrParamsWithDefaults instantiates a new PitrParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetentionPeriodSec

`func (o *PitrParams) GetRetentionPeriodSec() int64`

GetRetentionPeriodSec returns the RetentionPeriodSec field if non-nil, zero value otherwise.

### GetRetentionPeriodSecOk

`func (o *PitrParams) GetRetentionPeriodSecOk() (*int64, bool)`

GetRetentionPeriodSecOk returns a tuple with the RetentionPeriodSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodSec

`func (o *PitrParams) SetRetentionPeriodSec(v int64)`

SetRetentionPeriodSec sets RetentionPeriodSec field to given value.

### HasRetentionPeriodSec

`func (o *PitrParams) HasRetentionPeriodSec() bool`

HasRetentionPeriodSec returns a boolean if a field has been set.

### GetSnapshotIntervalSec

`func (o *PitrParams) GetSnapshotIntervalSec() int64`

GetSnapshotIntervalSec returns the SnapshotIntervalSec field if non-nil, zero value otherwise.

### GetSnapshotIntervalSecOk

`func (o *PitrParams) GetSnapshotIntervalSecOk() (*int64, bool)`

GetSnapshotIntervalSecOk returns a tuple with the SnapshotIntervalSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotIntervalSec

`func (o *PitrParams) SetSnapshotIntervalSec(v int64)`

SetSnapshotIntervalSec sets SnapshotIntervalSec field to given value.

### HasSnapshotIntervalSec

`func (o *PitrParams) HasSnapshotIntervalSec() bool`

HasSnapshotIntervalSec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


