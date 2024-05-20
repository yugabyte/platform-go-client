# StateChangeAuditInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** |  | 
**FieldName** | **string** |  | 
**Id** | **string** |  | 
**New** | **bool** |  | 
**PreviousValue** | **string** |  | 
**Timestamp** | **int64** |  | 
**UpdatedValue** | **string** |  | 
**UserId** | **string** |  | 

## Methods

### NewStateChangeAuditInfo

`func NewStateChangeAuditInfo(customerId string, fieldName string, id string, new bool, previousValue string, timestamp int64, updatedValue string, userId string, ) *StateChangeAuditInfo`

NewStateChangeAuditInfo instantiates a new StateChangeAuditInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateChangeAuditInfoWithDefaults

`func NewStateChangeAuditInfoWithDefaults() *StateChangeAuditInfo`

NewStateChangeAuditInfoWithDefaults instantiates a new StateChangeAuditInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *StateChangeAuditInfo) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *StateChangeAuditInfo) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *StateChangeAuditInfo) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetFieldName

`func (o *StateChangeAuditInfo) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *StateChangeAuditInfo) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *StateChangeAuditInfo) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetId

`func (o *StateChangeAuditInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StateChangeAuditInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StateChangeAuditInfo) SetId(v string)`

SetId sets Id field to given value.


### GetNew

`func (o *StateChangeAuditInfo) GetNew() bool`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *StateChangeAuditInfo) GetNewOk() (*bool, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *StateChangeAuditInfo) SetNew(v bool)`

SetNew sets New field to given value.


### GetPreviousValue

`func (o *StateChangeAuditInfo) GetPreviousValue() string`

GetPreviousValue returns the PreviousValue field if non-nil, zero value otherwise.

### GetPreviousValueOk

`func (o *StateChangeAuditInfo) GetPreviousValueOk() (*string, bool)`

GetPreviousValueOk returns a tuple with the PreviousValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousValue

`func (o *StateChangeAuditInfo) SetPreviousValue(v string)`

SetPreviousValue sets PreviousValue field to given value.


### GetTimestamp

`func (o *StateChangeAuditInfo) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *StateChangeAuditInfo) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *StateChangeAuditInfo) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetUpdatedValue

`func (o *StateChangeAuditInfo) GetUpdatedValue() string`

GetUpdatedValue returns the UpdatedValue field if non-nil, zero value otherwise.

### GetUpdatedValueOk

`func (o *StateChangeAuditInfo) GetUpdatedValueOk() (*string, bool)`

GetUpdatedValueOk returns a tuple with the UpdatedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedValue

`func (o *StateChangeAuditInfo) SetUpdatedValue(v string)`

SetUpdatedValue sets UpdatedValue field to given value.


### GetUserId

`func (o *StateChangeAuditInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *StateChangeAuditInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *StateChangeAuditInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


