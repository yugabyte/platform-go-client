# OCICloudInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OciAuthType** | Pointer to **string** | OCI authentication type (API_KEY or INSTANCE_PRINCIPAL) | [optional] 
**OciCompartmentId** | Pointer to **string** | OCI Compartment OCID | [optional] 
**OciFingerprint** | Pointer to **string** | OCI API Key Fingerprint | [optional] 
**OciHostedZoneId** | Pointer to **string** | OCI DNS Zone OCID for hosted zone | [optional] 
**OciPrivateKeyContent** | Pointer to **string** | OCI API Private Key content (PEM format) | [optional] 
**OciRegion** | Pointer to **string** | Default OCI Region | [optional] 
**OciTenancyId** | Pointer to **string** | OCI Tenancy OCID | [optional] 
**OciUserId** | Pointer to **string** | OCI User OCID | [optional] 
**VpcType** | Pointer to **string** | New/Existing VCN for provider creation | [optional] [readonly] 

## Methods

### NewOCICloudInfo

`func NewOCICloudInfo() *OCICloudInfo`

NewOCICloudInfo instantiates a new OCICloudInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCICloudInfoWithDefaults

`func NewOCICloudInfoWithDefaults() *OCICloudInfo`

NewOCICloudInfoWithDefaults instantiates a new OCICloudInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOciAuthType

`func (o *OCICloudInfo) GetOciAuthType() string`

GetOciAuthType returns the OciAuthType field if non-nil, zero value otherwise.

### GetOciAuthTypeOk

`func (o *OCICloudInfo) GetOciAuthTypeOk() (*string, bool)`

GetOciAuthTypeOk returns a tuple with the OciAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciAuthType

`func (o *OCICloudInfo) SetOciAuthType(v string)`

SetOciAuthType sets OciAuthType field to given value.

### HasOciAuthType

`func (o *OCICloudInfo) HasOciAuthType() bool`

HasOciAuthType returns a boolean if a field has been set.

### GetOciCompartmentId

`func (o *OCICloudInfo) GetOciCompartmentId() string`

GetOciCompartmentId returns the OciCompartmentId field if non-nil, zero value otherwise.

### GetOciCompartmentIdOk

`func (o *OCICloudInfo) GetOciCompartmentIdOk() (*string, bool)`

GetOciCompartmentIdOk returns a tuple with the OciCompartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciCompartmentId

`func (o *OCICloudInfo) SetOciCompartmentId(v string)`

SetOciCompartmentId sets OciCompartmentId field to given value.

### HasOciCompartmentId

`func (o *OCICloudInfo) HasOciCompartmentId() bool`

HasOciCompartmentId returns a boolean if a field has been set.

### GetOciFingerprint

`func (o *OCICloudInfo) GetOciFingerprint() string`

GetOciFingerprint returns the OciFingerprint field if non-nil, zero value otherwise.

### GetOciFingerprintOk

`func (o *OCICloudInfo) GetOciFingerprintOk() (*string, bool)`

GetOciFingerprintOk returns a tuple with the OciFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciFingerprint

`func (o *OCICloudInfo) SetOciFingerprint(v string)`

SetOciFingerprint sets OciFingerprint field to given value.

### HasOciFingerprint

`func (o *OCICloudInfo) HasOciFingerprint() bool`

HasOciFingerprint returns a boolean if a field has been set.

### GetOciHostedZoneId

`func (o *OCICloudInfo) GetOciHostedZoneId() string`

GetOciHostedZoneId returns the OciHostedZoneId field if non-nil, zero value otherwise.

### GetOciHostedZoneIdOk

`func (o *OCICloudInfo) GetOciHostedZoneIdOk() (*string, bool)`

GetOciHostedZoneIdOk returns a tuple with the OciHostedZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciHostedZoneId

`func (o *OCICloudInfo) SetOciHostedZoneId(v string)`

SetOciHostedZoneId sets OciHostedZoneId field to given value.

### HasOciHostedZoneId

`func (o *OCICloudInfo) HasOciHostedZoneId() bool`

HasOciHostedZoneId returns a boolean if a field has been set.

### GetOciPrivateKeyContent

`func (o *OCICloudInfo) GetOciPrivateKeyContent() string`

GetOciPrivateKeyContent returns the OciPrivateKeyContent field if non-nil, zero value otherwise.

### GetOciPrivateKeyContentOk

`func (o *OCICloudInfo) GetOciPrivateKeyContentOk() (*string, bool)`

GetOciPrivateKeyContentOk returns a tuple with the OciPrivateKeyContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciPrivateKeyContent

`func (o *OCICloudInfo) SetOciPrivateKeyContent(v string)`

SetOciPrivateKeyContent sets OciPrivateKeyContent field to given value.

### HasOciPrivateKeyContent

`func (o *OCICloudInfo) HasOciPrivateKeyContent() bool`

HasOciPrivateKeyContent returns a boolean if a field has been set.

### GetOciRegion

`func (o *OCICloudInfo) GetOciRegion() string`

GetOciRegion returns the OciRegion field if non-nil, zero value otherwise.

### GetOciRegionOk

`func (o *OCICloudInfo) GetOciRegionOk() (*string, bool)`

GetOciRegionOk returns a tuple with the OciRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciRegion

`func (o *OCICloudInfo) SetOciRegion(v string)`

SetOciRegion sets OciRegion field to given value.

### HasOciRegion

`func (o *OCICloudInfo) HasOciRegion() bool`

HasOciRegion returns a boolean if a field has been set.

### GetOciTenancyId

`func (o *OCICloudInfo) GetOciTenancyId() string`

GetOciTenancyId returns the OciTenancyId field if non-nil, zero value otherwise.

### GetOciTenancyIdOk

`func (o *OCICloudInfo) GetOciTenancyIdOk() (*string, bool)`

GetOciTenancyIdOk returns a tuple with the OciTenancyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciTenancyId

`func (o *OCICloudInfo) SetOciTenancyId(v string)`

SetOciTenancyId sets OciTenancyId field to given value.

### HasOciTenancyId

`func (o *OCICloudInfo) HasOciTenancyId() bool`

HasOciTenancyId returns a boolean if a field has been set.

### GetOciUserId

`func (o *OCICloudInfo) GetOciUserId() string`

GetOciUserId returns the OciUserId field if non-nil, zero value otherwise.

### GetOciUserIdOk

`func (o *OCICloudInfo) GetOciUserIdOk() (*string, bool)`

GetOciUserIdOk returns a tuple with the OciUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciUserId

`func (o *OCICloudInfo) SetOciUserId(v string)`

SetOciUserId sets OciUserId field to given value.

### HasOciUserId

`func (o *OCICloudInfo) HasOciUserId() bool`

HasOciUserId returns a boolean if a field has been set.

### GetVpcType

`func (o *OCICloudInfo) GetVpcType() string`

GetVpcType returns the VpcType field if non-nil, zero value otherwise.

### GetVpcTypeOk

`func (o *OCICloudInfo) GetVpcTypeOk() (*string, bool)`

GetVpcTypeOk returns a tuple with the VpcType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcType

`func (o *OCICloudInfo) SetVpcType(v string)`

SetVpcType sets VpcType field to given value.

### HasVpcType

`func (o *OCICloudInfo) HasVpcType() bool`

HasVpcType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


