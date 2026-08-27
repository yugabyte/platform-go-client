# NodeAgentUpgradeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateName** | Pointer to **string** | Optional name (label) of the YBA certificate config to use for the node agent upgrade. If omitted, existing certificates are used.  | [optional] 
**NodeNames** | Pointer to **[]string** | Optional list of node names whose node agents should be upgraded. If omitted or empty, all live nodes in the universe are upgraded.  | [optional] 
**CertsOnly** | Pointer to **bool** | If true, replace node agent certificates without upgrading the package when the node agent version already matches YBA. If the version does not match, a full upgrade is performed so RPC stays compatible.  | [optional] [default to false]

## Methods

### NewNodeAgentUpgradeSpec

`func NewNodeAgentUpgradeSpec() *NodeAgentUpgradeSpec`

NewNodeAgentUpgradeSpec instantiates a new NodeAgentUpgradeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeAgentUpgradeSpecWithDefaults

`func NewNodeAgentUpgradeSpecWithDefaults() *NodeAgentUpgradeSpec`

NewNodeAgentUpgradeSpecWithDefaults instantiates a new NodeAgentUpgradeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateName

`func (o *NodeAgentUpgradeSpec) GetCertificateName() string`

GetCertificateName returns the CertificateName field if non-nil, zero value otherwise.

### GetCertificateNameOk

`func (o *NodeAgentUpgradeSpec) GetCertificateNameOk() (*string, bool)`

GetCertificateNameOk returns a tuple with the CertificateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateName

`func (o *NodeAgentUpgradeSpec) SetCertificateName(v string)`

SetCertificateName sets CertificateName field to given value.

### HasCertificateName

`func (o *NodeAgentUpgradeSpec) HasCertificateName() bool`

HasCertificateName returns a boolean if a field has been set.

### GetNodeNames

`func (o *NodeAgentUpgradeSpec) GetNodeNames() []string`

GetNodeNames returns the NodeNames field if non-nil, zero value otherwise.

### GetNodeNamesOk

`func (o *NodeAgentUpgradeSpec) GetNodeNamesOk() (*[]string, bool)`

GetNodeNamesOk returns a tuple with the NodeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeNames

`func (o *NodeAgentUpgradeSpec) SetNodeNames(v []string)`

SetNodeNames sets NodeNames field to given value.

### HasNodeNames

`func (o *NodeAgentUpgradeSpec) HasNodeNames() bool`

HasNodeNames returns a boolean if a field has been set.

### GetCertsOnly

`func (o *NodeAgentUpgradeSpec) GetCertsOnly() bool`

GetCertsOnly returns the CertsOnly field if non-nil, zero value otherwise.

### GetCertsOnlyOk

`func (o *NodeAgentUpgradeSpec) GetCertsOnlyOk() (*bool, bool)`

GetCertsOnlyOk returns a tuple with the CertsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertsOnly

`func (o *NodeAgentUpgradeSpec) SetCertsOnly(v bool)`

SetCertsOnly sets CertsOnly field to given value.

### HasCertsOnly

`func (o *NodeAgentUpgradeSpec) HasCertsOnly() bool`

HasCertsOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


