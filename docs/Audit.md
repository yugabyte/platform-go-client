# Audit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action | [optional] [readonly] 
**AdditionalDetails** | Pointer to **map[string]interface{}** | Additional Details | [optional] [readonly] 
**ApiCall** | Pointer to **string** | API call | [optional] [readonly] 
**ApiMethod** | Pointer to **string** | API method | [optional] [readonly] 
**AuditID** | **int64** |  | 
**CustomerUUID** | Pointer to **string** | Customer UUID | [optional] [readonly] 
**Payload** | Pointer to **map[string]interface{}** | Audit UUID | [optional] [readonly] 
**Target** | Pointer to **string** | Target | [optional] [readonly] 
**TargetID** | Pointer to **string** | Target ID | [optional] [readonly] 
**TaskUUID** | Pointer to **string** | Task UUID | [optional] [readonly] 
**Timestamp** | **time.Time** |  | 
**UserEmail** | Pointer to **string** | User Email | [optional] [readonly] 
**UserUUID** | Pointer to **string** | User UUID | [optional] [readonly] 

## Methods

### NewAudit

`func NewAudit(auditID int64, timestamp time.Time, ) *Audit`

NewAudit instantiates a new Audit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditWithDefaults

`func NewAuditWithDefaults() *Audit`

NewAuditWithDefaults instantiates a new Audit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Audit) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Audit) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Audit) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *Audit) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAdditionalDetails

`func (o *Audit) GetAdditionalDetails() map[string]interface{}`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *Audit) GetAdditionalDetailsOk() (*map[string]interface{}, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDetails

`func (o *Audit) SetAdditionalDetails(v map[string]interface{})`

SetAdditionalDetails sets AdditionalDetails field to given value.

### HasAdditionalDetails

`func (o *Audit) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.

### GetApiCall

`func (o *Audit) GetApiCall() string`

GetApiCall returns the ApiCall field if non-nil, zero value otherwise.

### GetApiCallOk

`func (o *Audit) GetApiCallOk() (*string, bool)`

GetApiCallOk returns a tuple with the ApiCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCall

`func (o *Audit) SetApiCall(v string)`

SetApiCall sets ApiCall field to given value.

### HasApiCall

`func (o *Audit) HasApiCall() bool`

HasApiCall returns a boolean if a field has been set.

### GetApiMethod

`func (o *Audit) GetApiMethod() string`

GetApiMethod returns the ApiMethod field if non-nil, zero value otherwise.

### GetApiMethodOk

`func (o *Audit) GetApiMethodOk() (*string, bool)`

GetApiMethodOk returns a tuple with the ApiMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiMethod

`func (o *Audit) SetApiMethod(v string)`

SetApiMethod sets ApiMethod field to given value.

### HasApiMethod

`func (o *Audit) HasApiMethod() bool`

HasApiMethod returns a boolean if a field has been set.

### GetAuditID

`func (o *Audit) GetAuditID() int64`

GetAuditID returns the AuditID field if non-nil, zero value otherwise.

### GetAuditIDOk

`func (o *Audit) GetAuditIDOk() (*int64, bool)`

GetAuditIDOk returns a tuple with the AuditID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditID

`func (o *Audit) SetAuditID(v int64)`

SetAuditID sets AuditID field to given value.


### GetCustomerUUID

`func (o *Audit) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *Audit) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *Audit) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.

### HasCustomerUUID

`func (o *Audit) HasCustomerUUID() bool`

HasCustomerUUID returns a boolean if a field has been set.

### GetPayload

`func (o *Audit) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Audit) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Audit) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *Audit) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetTarget

`func (o *Audit) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Audit) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Audit) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *Audit) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTargetID

`func (o *Audit) GetTargetID() string`

GetTargetID returns the TargetID field if non-nil, zero value otherwise.

### GetTargetIDOk

`func (o *Audit) GetTargetIDOk() (*string, bool)`

GetTargetIDOk returns a tuple with the TargetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetID

`func (o *Audit) SetTargetID(v string)`

SetTargetID sets TargetID field to given value.

### HasTargetID

`func (o *Audit) HasTargetID() bool`

HasTargetID returns a boolean if a field has been set.

### GetTaskUUID

`func (o *Audit) GetTaskUUID() string`

GetTaskUUID returns the TaskUUID field if non-nil, zero value otherwise.

### GetTaskUUIDOk

`func (o *Audit) GetTaskUUIDOk() (*string, bool)`

GetTaskUUIDOk returns a tuple with the TaskUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskUUID

`func (o *Audit) SetTaskUUID(v string)`

SetTaskUUID sets TaskUUID field to given value.

### HasTaskUUID

`func (o *Audit) HasTaskUUID() bool`

HasTaskUUID returns a boolean if a field has been set.

### GetTimestamp

`func (o *Audit) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Audit) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Audit) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetUserEmail

`func (o *Audit) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *Audit) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *Audit) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *Audit) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetUserUUID

`func (o *Audit) GetUserUUID() string`

GetUserUUID returns the UserUUID field if non-nil, zero value otherwise.

### GetUserUUIDOk

`func (o *Audit) GetUserUUIDOk() (*string, bool)`

GetUserUUIDOk returns a tuple with the UserUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUUID

`func (o *Audit) SetUserUUID(v string)`

SetUserUUID sets UserUUID field to given value.

### HasUserUUID

`func (o *Audit) HasUserUUID() bool`

HasUserUUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


