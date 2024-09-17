# EncryptionInTransitSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableNodeToNodeEncrypt** | Pointer to **bool** | Whether to enable encryption for communication among DB nodes | [optional] 
**EnableClientToNodeEncrypt** | Pointer to **bool** | Whether to enable encryption for client connection to DB nodes | [optional] 
**RootCa** | Pointer to **string** | The UUID of the rootCA to be used to generate node certificates and facilitate TLS communication between database nodes. | [optional] 
**ClientRootCa** | Pointer to **string** | The UUID of the clientRootCA to be used to generate client certificates and facilitate TLS communication between server and client. Can be set to same as root_CA. | [optional] 

## Methods

### NewEncryptionInTransitSpec

`func NewEncryptionInTransitSpec() *EncryptionInTransitSpec`

NewEncryptionInTransitSpec instantiates a new EncryptionInTransitSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionInTransitSpecWithDefaults

`func NewEncryptionInTransitSpecWithDefaults() *EncryptionInTransitSpec`

NewEncryptionInTransitSpecWithDefaults instantiates a new EncryptionInTransitSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableNodeToNodeEncrypt

`func (o *EncryptionInTransitSpec) GetEnableNodeToNodeEncrypt() bool`

GetEnableNodeToNodeEncrypt returns the EnableNodeToNodeEncrypt field if non-nil, zero value otherwise.

### GetEnableNodeToNodeEncryptOk

`func (o *EncryptionInTransitSpec) GetEnableNodeToNodeEncryptOk() (*bool, bool)`

GetEnableNodeToNodeEncryptOk returns a tuple with the EnableNodeToNodeEncrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNodeToNodeEncrypt

`func (o *EncryptionInTransitSpec) SetEnableNodeToNodeEncrypt(v bool)`

SetEnableNodeToNodeEncrypt sets EnableNodeToNodeEncrypt field to given value.

### HasEnableNodeToNodeEncrypt

`func (o *EncryptionInTransitSpec) HasEnableNodeToNodeEncrypt() bool`

HasEnableNodeToNodeEncrypt returns a boolean if a field has been set.

### GetEnableClientToNodeEncrypt

`func (o *EncryptionInTransitSpec) GetEnableClientToNodeEncrypt() bool`

GetEnableClientToNodeEncrypt returns the EnableClientToNodeEncrypt field if non-nil, zero value otherwise.

### GetEnableClientToNodeEncryptOk

`func (o *EncryptionInTransitSpec) GetEnableClientToNodeEncryptOk() (*bool, bool)`

GetEnableClientToNodeEncryptOk returns a tuple with the EnableClientToNodeEncrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientToNodeEncrypt

`func (o *EncryptionInTransitSpec) SetEnableClientToNodeEncrypt(v bool)`

SetEnableClientToNodeEncrypt sets EnableClientToNodeEncrypt field to given value.

### HasEnableClientToNodeEncrypt

`func (o *EncryptionInTransitSpec) HasEnableClientToNodeEncrypt() bool`

HasEnableClientToNodeEncrypt returns a boolean if a field has been set.

### GetRootCa

`func (o *EncryptionInTransitSpec) GetRootCa() string`

GetRootCa returns the RootCa field if non-nil, zero value otherwise.

### GetRootCaOk

`func (o *EncryptionInTransitSpec) GetRootCaOk() (*string, bool)`

GetRootCaOk returns a tuple with the RootCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCa

`func (o *EncryptionInTransitSpec) SetRootCa(v string)`

SetRootCa sets RootCa field to given value.

### HasRootCa

`func (o *EncryptionInTransitSpec) HasRootCa() bool`

HasRootCa returns a boolean if a field has been set.

### GetClientRootCa

`func (o *EncryptionInTransitSpec) GetClientRootCa() string`

GetClientRootCa returns the ClientRootCa field if non-nil, zero value otherwise.

### GetClientRootCaOk

`func (o *EncryptionInTransitSpec) GetClientRootCaOk() (*string, bool)`

GetClientRootCaOk returns a tuple with the ClientRootCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRootCa

`func (o *EncryptionInTransitSpec) SetClientRootCa(v string)`

SetClientRootCa sets ClientRootCa field to given value.

### HasClientRootCa

`func (o *EncryptionInTransitSpec) HasClientRootCa() bool`

HasClientRootCa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


