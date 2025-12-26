# S3Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **string** | Access Key | 
**Bucket** | **string** | S3 bucket | 
**DirectoryPrefix** | Pointer to **string** | S3 Prefix Key (root directory inside bucket) | [optional] 
**DisableSSL** | Pointer to **bool** | Disable SSL | [optional] 
**Endpoint** | Pointer to **string** | Endpoint. Overrides the endpoint used by the exporter instead of constructing it from region and bucket | [optional] 
**FilePrefix** | Pointer to **string** | File prefix | [optional] 
**ForcePathStyle** | Pointer to **bool** | Force Path Style. Set this to true to force the request to use path-style addressing | [optional] 
**Marshaler** | Pointer to **string** | Marshaler | [optional] 
**Partition** | Pointer to **string** | S3 Partition | [optional] 
**Region** | **string** | Region | 
**RoleArn** | Pointer to **string** | Role ARN | [optional] 
**SecretKey** | **string** | Secret Key | 

## Methods

### NewS3Config

`func NewS3Config(accessKey string, bucket string, region string, secretKey string, ) *S3Config`

NewS3Config instantiates a new S3Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ConfigWithDefaults

`func NewS3ConfigWithDefaults() *S3Config`

NewS3ConfigWithDefaults instantiates a new S3Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *S3Config) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *S3Config) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *S3Config) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetBucket

`func (o *S3Config) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *S3Config) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *S3Config) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetDirectoryPrefix

`func (o *S3Config) GetDirectoryPrefix() string`

GetDirectoryPrefix returns the DirectoryPrefix field if non-nil, zero value otherwise.

### GetDirectoryPrefixOk

`func (o *S3Config) GetDirectoryPrefixOk() (*string, bool)`

GetDirectoryPrefixOk returns a tuple with the DirectoryPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryPrefix

`func (o *S3Config) SetDirectoryPrefix(v string)`

SetDirectoryPrefix sets DirectoryPrefix field to given value.

### HasDirectoryPrefix

`func (o *S3Config) HasDirectoryPrefix() bool`

HasDirectoryPrefix returns a boolean if a field has been set.

### GetDisableSSL

`func (o *S3Config) GetDisableSSL() bool`

GetDisableSSL returns the DisableSSL field if non-nil, zero value otherwise.

### GetDisableSSLOk

`func (o *S3Config) GetDisableSSLOk() (*bool, bool)`

GetDisableSSLOk returns a tuple with the DisableSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSSL

`func (o *S3Config) SetDisableSSL(v bool)`

SetDisableSSL sets DisableSSL field to given value.

### HasDisableSSL

`func (o *S3Config) HasDisableSSL() bool`

HasDisableSSL returns a boolean if a field has been set.

### GetEndpoint

`func (o *S3Config) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *S3Config) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *S3Config) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *S3Config) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetFilePrefix

`func (o *S3Config) GetFilePrefix() string`

GetFilePrefix returns the FilePrefix field if non-nil, zero value otherwise.

### GetFilePrefixOk

`func (o *S3Config) GetFilePrefixOk() (*string, bool)`

GetFilePrefixOk returns a tuple with the FilePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePrefix

`func (o *S3Config) SetFilePrefix(v string)`

SetFilePrefix sets FilePrefix field to given value.

### HasFilePrefix

`func (o *S3Config) HasFilePrefix() bool`

HasFilePrefix returns a boolean if a field has been set.

### GetForcePathStyle

`func (o *S3Config) GetForcePathStyle() bool`

GetForcePathStyle returns the ForcePathStyle field if non-nil, zero value otherwise.

### GetForcePathStyleOk

`func (o *S3Config) GetForcePathStyleOk() (*bool, bool)`

GetForcePathStyleOk returns a tuple with the ForcePathStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePathStyle

`func (o *S3Config) SetForcePathStyle(v bool)`

SetForcePathStyle sets ForcePathStyle field to given value.

### HasForcePathStyle

`func (o *S3Config) HasForcePathStyle() bool`

HasForcePathStyle returns a boolean if a field has been set.

### GetMarshaler

`func (o *S3Config) GetMarshaler() string`

GetMarshaler returns the Marshaler field if non-nil, zero value otherwise.

### GetMarshalerOk

`func (o *S3Config) GetMarshalerOk() (*string, bool)`

GetMarshalerOk returns a tuple with the Marshaler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarshaler

`func (o *S3Config) SetMarshaler(v string)`

SetMarshaler sets Marshaler field to given value.

### HasMarshaler

`func (o *S3Config) HasMarshaler() bool`

HasMarshaler returns a boolean if a field has been set.

### GetPartition

`func (o *S3Config) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *S3Config) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *S3Config) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *S3Config) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetRegion

`func (o *S3Config) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *S3Config) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *S3Config) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetRoleArn

`func (o *S3Config) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *S3Config) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *S3Config) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *S3Config) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetSecretKey

`func (o *S3Config) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *S3Config) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *S3Config) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


