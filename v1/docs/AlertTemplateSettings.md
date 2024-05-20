# AlertTemplateSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | **time.Time** | Creation time | [readonly] 
**CustomerUUID** | **string** | Customer UUID | [readonly] 
**Labels** | Pointer to **map[string]string** | Labels | [optional] 
**Template** | **string** | Template | 
**Uuid** | **string** |  | 

## Methods

### NewAlertTemplateSettings

`func NewAlertTemplateSettings(createTime time.Time, customerUUID string, template string, uuid string, ) *AlertTemplateSettings`

NewAlertTemplateSettings instantiates a new AlertTemplateSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertTemplateSettingsWithDefaults

`func NewAlertTemplateSettingsWithDefaults() *AlertTemplateSettings`

NewAlertTemplateSettingsWithDefaults instantiates a new AlertTemplateSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *AlertTemplateSettings) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *AlertTemplateSettings) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *AlertTemplateSettings) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.


### GetCustomerUUID

`func (o *AlertTemplateSettings) GetCustomerUUID() string`

GetCustomerUUID returns the CustomerUUID field if non-nil, zero value otherwise.

### GetCustomerUUIDOk

`func (o *AlertTemplateSettings) GetCustomerUUIDOk() (*string, bool)`

GetCustomerUUIDOk returns a tuple with the CustomerUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUUID

`func (o *AlertTemplateSettings) SetCustomerUUID(v string)`

SetCustomerUUID sets CustomerUUID field to given value.


### GetLabels

`func (o *AlertTemplateSettings) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AlertTemplateSettings) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AlertTemplateSettings) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AlertTemplateSettings) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTemplate

`func (o *AlertTemplateSettings) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *AlertTemplateSettings) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *AlertTemplateSettings) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetUuid

`func (o *AlertTemplateSettings) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AlertTemplateSettings) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AlertTemplateSettings) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


