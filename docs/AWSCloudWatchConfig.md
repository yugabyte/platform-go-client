# AWSCloudWatchConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **string** | Access Key | 
**Endpoint** | Pointer to **string** | End Point | [optional] 
**LogGroup** | **string** | Log Group | 
**LogStream** | **string** | Log Stream | 
**Region** | **string** | Region | 
**RoleARN** | Pointer to **string** | Role ARN | [optional] 
**SecretKey** | **string** | Secret Key | 

## Methods

### NewAWSCloudWatchConfig

`func NewAWSCloudWatchConfig(accessKey string, logGroup string, logStream string, region string, secretKey string, ) *AWSCloudWatchConfig`

NewAWSCloudWatchConfig instantiates a new AWSCloudWatchConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSCloudWatchConfigWithDefaults

`func NewAWSCloudWatchConfigWithDefaults() *AWSCloudWatchConfig`

NewAWSCloudWatchConfigWithDefaults instantiates a new AWSCloudWatchConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *AWSCloudWatchConfig) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AWSCloudWatchConfig) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AWSCloudWatchConfig) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetEndpoint

`func (o *AWSCloudWatchConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AWSCloudWatchConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AWSCloudWatchConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *AWSCloudWatchConfig) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetLogGroup

`func (o *AWSCloudWatchConfig) GetLogGroup() string`

GetLogGroup returns the LogGroup field if non-nil, zero value otherwise.

### GetLogGroupOk

`func (o *AWSCloudWatchConfig) GetLogGroupOk() (*string, bool)`

GetLogGroupOk returns a tuple with the LogGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogGroup

`func (o *AWSCloudWatchConfig) SetLogGroup(v string)`

SetLogGroup sets LogGroup field to given value.


### GetLogStream

`func (o *AWSCloudWatchConfig) GetLogStream() string`

GetLogStream returns the LogStream field if non-nil, zero value otherwise.

### GetLogStreamOk

`func (o *AWSCloudWatchConfig) GetLogStreamOk() (*string, bool)`

GetLogStreamOk returns a tuple with the LogStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStream

`func (o *AWSCloudWatchConfig) SetLogStream(v string)`

SetLogStream sets LogStream field to given value.


### GetRegion

`func (o *AWSCloudWatchConfig) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AWSCloudWatchConfig) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AWSCloudWatchConfig) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetRoleARN

`func (o *AWSCloudWatchConfig) GetRoleARN() string`

GetRoleARN returns the RoleARN field if non-nil, zero value otherwise.

### GetRoleARNOk

`func (o *AWSCloudWatchConfig) GetRoleARNOk() (*string, bool)`

GetRoleARNOk returns a tuple with the RoleARN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleARN

`func (o *AWSCloudWatchConfig) SetRoleARN(v string)`

SetRoleARN sets RoleARN field to given value.

### HasRoleARN

`func (o *AWSCloudWatchConfig) HasRoleARN() bool`

HasRoleARN returns a boolean if a field has been set.

### GetSecretKey

`func (o *AWSCloudWatchConfig) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AWSCloudWatchConfig) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AWSCloudWatchConfig) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


