# \AlertsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Acknowledge**](AlertsApi.md#Acknowledge) | **Post** /api/v1/customers/{cUUID}/alerts/{alertUUID}/acknowledge | Acknowledge an alert
[**AcknowledgeByFilter**](AlertsApi.md#AcknowledgeByFilter) | **Post** /api/v1/customers/{cUUID}/alerts/acknowledge | Acknowledge all alerts
[**CountAlerts**](AlertsApi.md#CountAlerts) | **Post** /api/v1/customers/{cUUID}/alerts/count | Count alerts
[**CreateAlertChannel**](AlertsApi.md#CreateAlertChannel) | **Post** /api/v1/customers/{cUUID}/alert_channels | Create an alert channel
[**CreateAlertConfiguration**](AlertsApi.md#CreateAlertConfiguration) | **Post** /api/v1/customers/{cUUID}/alert_configurations | Create an alert configuration
[**CreateAlertDestination**](AlertsApi.md#CreateAlertDestination) | **Post** /api/v1/customers/{cUUID}/alert_destinations | Create an alert destination
[**DeleteAlertChannel**](AlertsApi.md#DeleteAlertChannel) | **Delete** /api/v1/customers/{cUUID}/alert_channels/{acUUID} | Delete an alert channel
[**DeleteAlertConfiguration**](AlertsApi.md#DeleteAlertConfiguration) | **Delete** /api/v1/customers/{cUUID}/alert_configurations/{configurationUUID} | Delete an alert configuration
[**DeleteAlertDestination**](AlertsApi.md#DeleteAlertDestination) | **Delete** /api/v1/customers/{cUUID}/alert_destinations/{adUUID} | Delete an alert destination
[**Get**](AlertsApi.md#Get) | **Get** /api/v1/customers/{cUUID}/alerts/{alertUUID} | Get details of an alert
[**GetAlertChannel**](AlertsApi.md#GetAlertChannel) | **Get** /api/v1/customers/{cUUID}/alert_channels/{acUUID} | Get an alert channel
[**GetAlertConfiguration**](AlertsApi.md#GetAlertConfiguration) | **Get** /api/v1/customers/{cUUID}/alert_configurations/{configurationUUID} | Get an alert configuration
[**GetAlertDestination**](AlertsApi.md#GetAlertDestination) | **Get** /api/v1/customers/{cUUID}/alert_destinations/{adUUID} | Get an alert destination
[**ListActive**](AlertsApi.md#ListActive) | **Get** /api/v1/customers/{cUUID}/alerts/active | List active alerts
[**ListAlertChannels**](AlertsApi.md#ListAlertChannels) | **Get** /api/v1/customers/{cUUID}/alert_channels | List all alert channels
[**ListAlertConfigurations**](AlertsApi.md#ListAlertConfigurations) | **Post** /api/v1/customers/{cUUID}/alert_configurations/list | Get filtered list of alert configurations
[**ListAlertDestinations**](AlertsApi.md#ListAlertDestinations) | **Get** /api/v1/customers/{cUUID}/alert_destinations | List alert destinations
[**ListAlertTemplates**](AlertsApi.md#ListAlertTemplates) | **Post** /api/v1/customers/{cUUID}/alert_templates | Get filtered list of alert configuration templates
[**ListOfAlerts**](AlertsApi.md#ListOfAlerts) | **Get** /api/v1/customers/{cUUID}/alerts | List all alerts
[**PageAlertConfigurations**](AlertsApi.md#PageAlertConfigurations) | **Post** /api/v1/customers/{cUUID}/alert_configurations/page | List all alert configurations (paginated)
[**PageAlerts**](AlertsApi.md#PageAlerts) | **Post** /api/v1/customers/{cUUID}/alerts/page | List alerts (paginated)
[**SendTestAlert**](AlertsApi.md#SendTestAlert) | **Post** /api/v1/customers/{cUUID}/alert_configurations/{configurationUUID}/test_alert | Send test alert for alert configuration
[**UpdateAlertChannel**](AlertsApi.md#UpdateAlertChannel) | **Put** /api/v1/customers/{cUUID}/alert_channels/{acUUID} | Update an alert channel
[**UpdateAlertConfiguration**](AlertsApi.md#UpdateAlertConfiguration) | **Put** /api/v1/customers/{cUUID}/alert_configurations/{configurationUUID} | Update an alert configuration
[**UpdateAlertDestination**](AlertsApi.md#UpdateAlertDestination) | **Put** /api/v1/customers/{cUUID}/alert_destinations/{adUUID} | Update an alert destination



## Acknowledge

> Alert Acknowledge(ctx, cUUID, alertUUID).Execute()

Acknowledge an alert

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    alertUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.Acknowledge(context.Background(), cUUID, alertUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.Acknowledge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Acknowledge`: Alert
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.Acknowledge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**alertUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcknowledgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Alert**](Alert.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AcknowledgeByFilter

> []Alert AcknowledgeByFilter(ctx, cUUID).AcknowledgeAlertsRequest(acknowledgeAlertsRequest).Execute()

Acknowledge all alerts

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    acknowledgeAlertsRequest := *openapiclient.NewAlertApiFilter([]string{"ConfigurationTypes_example"}, "ConfigurationUuid_example", []string{"Severities_example"}, "SourceName_example", []string{"SourceUUIDs_example"}, []string{"States_example"}, []string{"Uuids_example"}) // AlertApiFilter | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.AcknowledgeByFilter(context.Background(), cUUID).AcknowledgeAlertsRequest(acknowledgeAlertsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.AcknowledgeByFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcknowledgeByFilter`: []Alert
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.AcknowledgeByFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcknowledgeByFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acknowledgeAlertsRequest** | [**AlertApiFilter**](AlertApiFilter.md) |  | 

### Return type

[**[]Alert**](Alert.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountAlerts

> int32 CountAlerts(ctx, cUUID).CountAlertsRequest(countAlertsRequest).Execute()

Count alerts

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    countAlertsRequest := *openapiclient.NewAlertApiFilter([]string{"ConfigurationTypes_example"}, "ConfigurationUuid_example", []string{"Severities_example"}, "SourceName_example", []string{"SourceUUIDs_example"}, []string{"States_example"}, []string{"Uuids_example"}) // AlertApiFilter | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.CountAlerts(context.Background(), cUUID).CountAlertsRequest(countAlertsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.CountAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountAlerts`: int32
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.CountAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **countAlertsRequest** | [**AlertApiFilter**](AlertApiFilter.md) |  | 

### Return type

**int32**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAlertChannel

> AlertChannel CreateAlertChannel(ctx, cUUID).CreateAlertChannelRequest(createAlertChannelRequest).Execute()

Create an alert channel

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    createAlertChannelRequest := *openapiclient.NewAlertChannelFormData("AlertChannelUUID_example", "Name_example", *openapiclient.NewAlertChannelParams("TextTemplate_example", "TitleTemplate_example")) // AlertChannelFormData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.CreateAlertChannel(context.Background(), cUUID).CreateAlertChannelRequest(createAlertChannelRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.CreateAlertChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlertChannel`: AlertChannel
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.CreateAlertChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAlertChannelRequest** | [**AlertChannelFormData**](AlertChannelFormData.md) |  | 

### Return type

[**AlertChannel**](AlertChannel.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAlertConfiguration

> AlertConfiguration CreateAlertConfiguration(ctx, cUUID).CreateAlertConfigurationRequest(createAlertConfigurationRequest).Execute()

Create an alert configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    createAlertConfigurationRequest := *openapiclient.NewAlertConfiguration(false, float64(123), time.Now(), "CustomerUUID_example", false, "Description_example", int32(123), "Name_example", *openapiclient.NewAlertConfigurationTarget(), "TargetType_example", "Template_example", "ThresholdUnit_example", map[string]AlertConfigurationThreshold{"key": *openapiclient.NewAlertConfigurationThreshold("Condition_example", float64(123))}) // AlertConfiguration | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.CreateAlertConfiguration(context.Background(), cUUID).CreateAlertConfigurationRequest(createAlertConfigurationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.CreateAlertConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlertConfiguration`: AlertConfiguration
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.CreateAlertConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAlertConfigurationRequest** | [**AlertConfiguration**](AlertConfiguration.md) |  | 

### Return type

[**AlertConfiguration**](AlertConfiguration.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAlertDestination

> AlertDestination CreateAlertDestination(ctx, cUUID).CreateAlertDestinationRequest(createAlertDestinationRequest).Execute()

Create an alert destination

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    createAlertDestinationRequest := *openapiclient.NewAlertDestinationFormData([]string{"Channels_example"}, false, "Name_example") // AlertDestinationFormData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.CreateAlertDestination(context.Background(), cUUID).CreateAlertDestinationRequest(createAlertDestinationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.CreateAlertDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlertDestination`: AlertDestination
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.CreateAlertDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAlertDestinationRequest** | [**AlertDestinationFormData**](AlertDestinationFormData.md) |  | 

### Return type

[**AlertDestination**](AlertDestination.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertChannel

> YBPSuccess DeleteAlertChannel(ctx, cUUID, acUUID).Execute()

Delete an alert channel

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    acUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.DeleteAlertChannel(context.Background(), cUUID, acUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.DeleteAlertChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAlertChannel`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.DeleteAlertChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**acUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**YBPSuccess**](YBPSuccess.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertConfiguration

> YBPSuccess DeleteAlertConfiguration(ctx, cUUID, configurationUUID).Execute()

Delete an alert configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    configurationUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.DeleteAlertConfiguration(context.Background(), cUUID, configurationUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.DeleteAlertConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAlertConfiguration`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.DeleteAlertConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**configurationUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**YBPSuccess**](YBPSuccess.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertDestination

> YBPSuccess DeleteAlertDestination(ctx, cUUID, adUUID).Execute()

Delete an alert destination

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    adUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.DeleteAlertDestination(context.Background(), cUUID, adUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.DeleteAlertDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAlertDestination`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.DeleteAlertDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**adUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**YBPSuccess**](YBPSuccess.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Alert Get(ctx, cUUID, alertUUID).Execute()

Get details of an alert

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    alertUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.Get(context.Background(), cUUID, alertUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Alert
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**alertUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Alert**](Alert.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertChannel

> AlertChannel GetAlertChannel(ctx, cUUID, acUUID).Execute()

Get an alert channel

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    acUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.GetAlertChannel(context.Background(), cUUID, acUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetAlertChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertChannel`: AlertChannel
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetAlertChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**acUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AlertChannel**](AlertChannel.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertConfiguration

> AlertConfiguration GetAlertConfiguration(ctx, cUUID, configurationUUID).Execute()

Get an alert configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    configurationUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.GetAlertConfiguration(context.Background(), cUUID, configurationUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetAlertConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertConfiguration`: AlertConfiguration
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetAlertConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**configurationUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AlertConfiguration**](AlertConfiguration.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertDestination

> AlertDestination GetAlertDestination(ctx, cUUID, adUUID).Execute()

Get an alert destination

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    adUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.GetAlertDestination(context.Background(), cUUID, adUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetAlertDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertDestination`: AlertDestination
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetAlertDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**adUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AlertDestination**](AlertDestination.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActive

> []Alert ListActive(ctx, cUUID).Execute()

List active alerts

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.ListActive(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.ListActive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListActive`: []Alert
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.ListActive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListActiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Alert**](Alert.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlertChannels

> []AlertChannel ListAlertChannels(ctx, cUUID).Execute()

List all alert channels

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.ListAlertChannels(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.ListAlertChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAlertChannels`: []AlertChannel
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.ListAlertChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAlertChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AlertChannel**](AlertChannel.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlertConfigurations

> []AlertConfiguration ListAlertConfigurations(ctx, cUUID).ListAlertConfigurationsRequest(listAlertConfigurationsRequest).Execute()

Get filtered list of alert configurations

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    listAlertConfigurationsRequest := *openapiclient.NewAlertConfigurationApiFilter(false, "DestinationType_example", "DestinationUuid_example", "Name_example", "Severity_example", *openapiclient.NewAlertConfigurationTarget(), "TargetType_example", "Template_example", []string{"Uuids_example"}) // AlertConfigurationApiFilter | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.ListAlertConfigurations(context.Background(), cUUID).ListAlertConfigurationsRequest(listAlertConfigurationsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.ListAlertConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAlertConfigurations`: []AlertConfiguration
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.ListAlertConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAlertConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **listAlertConfigurationsRequest** | [**AlertConfigurationApiFilter**](AlertConfigurationApiFilter.md) |  | 

### Return type

[**[]AlertConfiguration**](AlertConfiguration.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlertDestinations

> []AlertDefinition ListAlertDestinations(ctx, cUUID).Execute()

List alert destinations

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.ListAlertDestinations(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.ListAlertDestinations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAlertDestinations`: []AlertDefinition
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.ListAlertDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAlertDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AlertDefinition**](AlertDefinition.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlertTemplates

> []AlertConfigurationTemplate ListAlertTemplates(ctx, cUUID).ListTemplatesRequest(listTemplatesRequest).Execute()

Get filtered list of alert configuration templates

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    listTemplatesRequest := *openapiclient.NewAlertTemplateApiFilter("Name_example", "TargetType_example") // AlertTemplateApiFilter | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.ListAlertTemplates(context.Background(), cUUID).ListTemplatesRequest(listTemplatesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.ListAlertTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAlertTemplates`: []AlertConfigurationTemplate
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.ListAlertTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAlertTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **listTemplatesRequest** | [**AlertTemplateApiFilter**](AlertTemplateApiFilter.md) |  | 

### Return type

[**[]AlertConfigurationTemplate**](AlertConfigurationTemplate.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOfAlerts

> []Alert ListOfAlerts(ctx, cUUID).Execute()

List all alerts

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.ListOfAlerts(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.ListOfAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOfAlerts`: []Alert
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.ListOfAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOfAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Alert**](Alert.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageAlertConfigurations

> AlertConfigurationPagedResponse PageAlertConfigurations(ctx, cUUID).PageAlertConfigurationsRequest(pageAlertConfigurationsRequest).Execute()

List all alert configurations (paginated)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    pageAlertConfigurationsRequest := *openapiclient.NewAlertConfigurationPagedApiQuery("Direction_example", *openapiclient.NewAlertConfigurationApiFilter(false, "DestinationType_example", "DestinationUuid_example", "Name_example", "Severity_example", *openapiclient.NewAlertConfigurationTarget(), "TargetType_example", "Template_example", []string{"Uuids_example"}), int32(123), false, int32(123), "SortBy_example") // AlertConfigurationPagedApiQuery | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.PageAlertConfigurations(context.Background(), cUUID).PageAlertConfigurationsRequest(pageAlertConfigurationsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.PageAlertConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PageAlertConfigurations`: AlertConfigurationPagedResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.PageAlertConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageAlertConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageAlertConfigurationsRequest** | [**AlertConfigurationPagedApiQuery**](AlertConfigurationPagedApiQuery.md) |  | 

### Return type

[**AlertConfigurationPagedResponse**](AlertConfigurationPagedResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageAlerts

> AlertPagedResponse PageAlerts(ctx, cUUID).PageAlertsRequest(pageAlertsRequest).Execute()

List alerts (paginated)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    pageAlertsRequest := *openapiclient.NewAlertPagedApiQuery("Direction_example", *openapiclient.NewAlertApiFilter([]string{"ConfigurationTypes_example"}, "ConfigurationUuid_example", []string{"Severities_example"}, "SourceName_example", []string{"SourceUUIDs_example"}, []string{"States_example"}, []string{"Uuids_example"}), int32(123), false, int32(123), "SortBy_example") // AlertPagedApiQuery | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.PageAlerts(context.Background(), cUUID).PageAlertsRequest(pageAlertsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.PageAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PageAlerts`: AlertPagedResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.PageAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageAlertsRequest** | [**AlertPagedApiQuery**](AlertPagedApiQuery.md) |  | 

### Return type

[**AlertPagedResponse**](AlertPagedResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTestAlert

> YBPSuccess SendTestAlert(ctx, cUUID, configurationUUID).Execute()

Send test alert for alert configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    configurationUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.SendTestAlert(context.Background(), cUUID, configurationUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.SendTestAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendTestAlert`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.SendTestAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**configurationUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendTestAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**YBPSuccess**](YBPSuccess.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertChannel

> AlertChannel UpdateAlertChannel(ctx, cUUID, acUUID).UpdateAlertChannelRequest(updateAlertChannelRequest).Execute()

Update an alert channel

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    acUUID := TODO // string | 
    updateAlertChannelRequest := *openapiclient.NewAlertChannelFormData("AlertChannelUUID_example", "Name_example", *openapiclient.NewAlertChannelParams("TextTemplate_example", "TitleTemplate_example")) // AlertChannelFormData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.UpdateAlertChannel(context.Background(), cUUID, acUUID).UpdateAlertChannelRequest(updateAlertChannelRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.UpdateAlertChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlertChannel`: AlertChannel
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.UpdateAlertChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**acUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateAlertChannelRequest** | [**AlertChannelFormData**](AlertChannelFormData.md) |  | 

### Return type

[**AlertChannel**](AlertChannel.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertConfiguration

> AlertConfiguration UpdateAlertConfiguration(ctx, cUUID, configurationUUID).UpdateAlertConfigurationRequest(updateAlertConfigurationRequest).Execute()

Update an alert configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    configurationUUID := TODO // string | 
    updateAlertConfigurationRequest := *openapiclient.NewAlertConfiguration(false, float64(123), time.Now(), "CustomerUUID_example", false, "Description_example", int32(123), "Name_example", *openapiclient.NewAlertConfigurationTarget(), "TargetType_example", "Template_example", "ThresholdUnit_example", map[string]AlertConfigurationThreshold{"key": *openapiclient.NewAlertConfigurationThreshold("Condition_example", float64(123))}) // AlertConfiguration | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.UpdateAlertConfiguration(context.Background(), cUUID, configurationUUID).UpdateAlertConfigurationRequest(updateAlertConfigurationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.UpdateAlertConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlertConfiguration`: AlertConfiguration
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.UpdateAlertConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**configurationUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateAlertConfigurationRequest** | [**AlertConfiguration**](AlertConfiguration.md) |  | 

### Return type

[**AlertConfiguration**](AlertConfiguration.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertDestination

> AlertDestination UpdateAlertDestination(ctx, cUUID, adUUID).UpdateAlertDestinationRequest(updateAlertDestinationRequest).Execute()

Update an alert destination

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cUUID := TODO // string | 
    adUUID := TODO // string | 
    updateAlertDestinationRequest := *openapiclient.NewAlertDestinationFormData([]string{"Channels_example"}, false, "Name_example") // AlertDestinationFormData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.UpdateAlertDestination(context.Background(), cUUID, adUUID).UpdateAlertDestinationRequest(updateAlertDestinationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.UpdateAlertDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlertDestination`: AlertDestination
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.UpdateAlertDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**adUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateAlertDestinationRequest** | [**AlertDestinationFormData**](AlertDestinationFormData.md) |  | 

### Return type

[**AlertDestination**](AlertDestination.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

