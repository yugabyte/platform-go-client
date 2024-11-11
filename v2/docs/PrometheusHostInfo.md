# PrometheusHostInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrometheusUrl** | **string** | URL for accessing Prometheus | [readonly] 
**UseBrowserFqdn** | **bool** | If the Prometheus link in the browser should use the browser&#39;s FQDN or use prometheus url directly. | [readonly] 

## Methods

### NewPrometheusHostInfo

`func NewPrometheusHostInfo(prometheusUrl string, useBrowserFqdn bool, ) *PrometheusHostInfo`

NewPrometheusHostInfo instantiates a new PrometheusHostInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusHostInfoWithDefaults

`func NewPrometheusHostInfoWithDefaults() *PrometheusHostInfo`

NewPrometheusHostInfoWithDefaults instantiates a new PrometheusHostInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrometheusUrl

`func (o *PrometheusHostInfo) GetPrometheusUrl() string`

GetPrometheusUrl returns the PrometheusUrl field if non-nil, zero value otherwise.

### GetPrometheusUrlOk

`func (o *PrometheusHostInfo) GetPrometheusUrlOk() (*string, bool)`

GetPrometheusUrlOk returns a tuple with the PrometheusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusUrl

`func (o *PrometheusHostInfo) SetPrometheusUrl(v string)`

SetPrometheusUrl sets PrometheusUrl field to given value.


### GetUseBrowserFqdn

`func (o *PrometheusHostInfo) GetUseBrowserFqdn() bool`

GetUseBrowserFqdn returns the UseBrowserFqdn field if non-nil, zero value otherwise.

### GetUseBrowserFqdnOk

`func (o *PrometheusHostInfo) GetUseBrowserFqdnOk() (*bool, bool)`

GetUseBrowserFqdnOk returns a tuple with the UseBrowserFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseBrowserFqdn

`func (o *PrometheusHostInfo) SetUseBrowserFqdn(v bool)`

SetUseBrowserFqdn sets UseBrowserFqdn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


