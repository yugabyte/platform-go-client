# \AlertsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Acknowledge**](AlertsApi.md#Acknowledge) | **Post** /api/v1/customers/{cUUID}/alerts/{alertUUID}/acknowledge | WARNING: This is a preview API that could change. Acknowledge an alert
[**AcknowledgeByFilter**](AlertsApi.md#AcknowledgeByFilter) | **Post** /api/v1/customers/{cUUID}/alerts/acknowledge | Deprecated since YBA version 2.8.0.0. Acknowledge all alerts
[**AlertNotificationPreview**](AlertsApi.md#AlertNotificationPreview) | **Post** /api/v1/customers/{cUUID}/alert_notification_preview | WARNING: This is a preview API that could change. Prepare alert notification preview
[**CountAlerts**](AlertsApi.md#CountAlerts) | **Post** /api/v1/customers/{cUUID}/alerts/count | WARNING: This is a preview API that could change. Count alerts
[**CreateAlertChannel**](AlertsApi.md#CreateAlertChannel) | **Post** /api/v1/customers/{cUUID}/alert_channels | WARNING: This is a preview API that could change. Create an alert channel
[**CreateAlertConfiguration**](AlertsApi.md#CreateAlertConfiguration) | **Post** /api/v1/customers/{cUUID}/alert_configurations | WARNING: This is a preview API that could change. Create an alert configuration
[**CreateAlertDestination**](AlertsApi.md#CreateAlertDestination) | **Post** /api/v1/customers/{cUUID}/alert_destinations | WARNING: This is a preview API that could change. Create an alert destination
[**DeleteAlertChannel**](AlertsApi.md#DeleteAlertChannel) | **Delete** /api/v1/customers/{cUUID}/alert_channels/{acUUID} | WARNING: This is a preview API that could change. Delete an alert channel
[**DeleteAlertChannelTemplates**](AlertsApi.md#DeleteAlertChannelTemplates) | **Delete** /api/v1/customers/{cUUID}/alert_channel_templates/{acType} | WARNING: This is a preview API that could change. Delete alert channel templates
[**DeleteAlertConfiguration**](AlertsApi.md#DeleteAlertConfiguration) | **Delete** /api/v1/customers/{cUUID}/alert_configurations/{configurationUUID} | WARNING: This is a preview API that could change. Delete an alert configuration
[**DeleteAlertDestination**](AlertsApi.md#DeleteAlertDestination) | **Delete** /api/v1/customers/{cUUID}/alert_destinations/{adUUID} | WARNING: This is a preview API that could change. Delete an alert destination
[**DeleteAlertTemplateSettings**](AlertsApi.md#DeleteAlertTemplateSettings) | **Delete** /api/v1/customers/{cUUID}/alert_template_settings/{settingsUUID} | Deprecated since YBA version 2.20.0.0. Delete an alert template settings
[**DeleteAlertTemplateVariables**](AlertsApi.md#DeleteAlertTemplateVariables) | **Delete** /api/v1/customers/{cUUID}/alert_template_variables/{variableUUID} | WARNING: This is a preview API that could change. Delete an alert template variables
[**EditAlertTemplateSettings**](AlertsApi.md#EditAlertTemplateSettings) | **Put** /api/v1/customers/{cUUID}/alert_template_settings | Deprecated since YBA version 2.20.0.0. Create or update alert template settings list
[**EditAlertTemplateVariables**](AlertsApi.md#EditAlertTemplateVariables) | **Put** /api/v1/customers/{cUUID}/alert_template_variables | WARNING: This is a preview API that could change. Create or update alert template variables
[**Get**](AlertsApi.md#Get) | **Get** /api/v1/customers/{cUUID}/alerts/{alertUUID} | WARNING: This is a preview API that could change. Get details of an alert
[**GetAlertChannel**](AlertsApi.md#GetAlertChannel) | **Get** /api/v1/customers/{cUUID}/alert_channels/{acUUID} | WARNING: This is a preview API that could change. Get an alert channel
[**GetAlertChannelTemplates**](AlertsApi.md#GetAlertChannelTemplates) | **Get** /api/v1/customers/{cUUID}/alert_channel_templates/{acType} | WARNING: This is a preview API that could change. Get alert channel templates
[**GetAlertConfiguration**](AlertsApi.md#GetAlertConfiguration) | **Get** /api/v1/customers/{cUUID}/alert_configurations/{configurationUUID} | WARNING: This is a preview API that could change. Get an alert configuration
[**GetAlertDestination**](AlertsApi.md#GetAlertDestination) | **Get** /api/v1/customers/{cUUID}/alert_destinations/{adUUID} | WARNING: This is a preview API that could change. Get an alert destination
[**ListActive**](AlertsApi.md#ListActive) | **Get** /api/v1/customers/{cUUID}/alerts/active | Deprecated since YBA version 2.8.0.0. List active alerts
[**ListAlertChannelTemplates**](AlertsApi.md#ListAlertChannelTemplates) | **Get** /api/v1/customers/{cUUID}/alert_channel_templates | WARNING: This is a preview API that could change. List all alert channel templates
[**ListAlertChannels**](AlertsApi.md#ListAlertChannels) | **Get** /api/v1/customers/{cUUID}/alert_channels | WARNING: This is a preview API that could change. List all alert channels
[**ListAlertConfigurations**](AlertsApi.md#ListAlertConfigurations) | **Post** /api/v1/customers/{cUUID}/alert_configurations/list | WARNING: This is a preview API that could change. Get filtered list of alert configurations
[**ListAlertDestinations**](AlertsApi.md#ListAlertDestinations) | **Get** /api/v1/customers/{cUUID}/alert_destinations | WARNING: This is a preview API that could change. List alert destinations
[**ListAlertTemplateSettings**](AlertsApi.md#ListAlertTemplateSettings) | **Get** /api/v1/customers/{cUUID}/alert_template_settings | Deprecated since YBA version 2.20.0.0. Get alert template settings
[**ListAlertTemplateVariables**](AlertsApi.md#ListAlertTemplateVariables) | **Get** /api/v1/customers/{cUUID}/alert_template_variables | WARNING: This is a preview API that could change. List alert template variables
[**ListAlertTemplates**](AlertsApi.md#ListAlertTemplates) | **Post** /api/v1/customers/{cUUID}/alert_templates | WARNING: This is a preview API that could change. Get filtered list of alert configuration templates
[**ListOfAlerts**](AlertsApi.md#ListOfAlerts) | **Get** /api/v1/customers/{cUUID}/alerts | WARNING: This is a preview API that could change. List all alerts
[**PageAlertConfigurations**](AlertsApi.md#PageAlertConfigurations) | **Post** /api/v1/customers/{cUUID}/alert_configurations/page | WARNING: This is a preview API that could change. List all alert configurations (paginated)
[**PageAlerts**](AlertsApi.md#PageAlerts) | **Post** /api/v1/customers/{cUUID}/alerts/page | WARNING: This is a preview API that could change. List alerts (paginated)
[**SendTestAlert**](AlertsApi.md#SendTestAlert) | **Post** /api/v1/customers/{cUUID}/alert_configurations/{configurationUUID}/test_alert | WARNING: This is a preview API that could change. Send test alert for alert configuration
[**SetAlertChannelTemplates**](AlertsApi.md#SetAlertChannelTemplates) | **Post** /api/v1/customers/{cUUID}/alert_channel_templates/{acType} | WARNING: This is a preview API that could change. Set alert channel templates
[**UpdateAlertChannel**](AlertsApi.md#UpdateAlertChannel) | **Put** /api/v1/customers/{cUUID}/alert_channels/{acUUID} | WARNING: This is a preview API that could change. Update an alert channel
[**UpdateAlertConfiguration**](AlertsApi.md#UpdateAlertConfiguration) | **Put** /api/v1/customers/{cUUID}/alert_configurations/{configurationUUID} | WARNING: This is a preview API that could change. Update an alert configuration
[**UpdateAlertDestination**](AlertsApi.md#UpdateAlertDestination) | **Put** /api/v1/customers/{cUUID}/alert_destinations/{adUUID} | WARNING: This is a preview API that could change. Update an alert destination



## Acknowledge

> Alert Acknowledge(ctx, cUUID, alertUUID).Request(request).Execute()

WARNING: This is a preview API that could change. Acknowledge an alert

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.Acknowledge(context.Background(), cUUID, alertUUID).Request(request).Execute()
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


 **request** | [**interface{}**](interface{}.md) |  | 

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

> []Alert AcknowledgeByFilter(ctx, cUUID).AcknowledgeAlertsRequest(acknowledgeAlertsRequest).Request(request).Execute()

Deprecated since YBA version 2.8.0.0. Acknowledge all alerts

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.AcknowledgeByFilter(context.Background(), cUUID).AcknowledgeAlertsRequest(acknowledgeAlertsRequest).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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


## AlertNotificationPreview

> AlertTemplateVariablesList AlertNotificationPreview(ctx, cUUID).NotificationPreviewRequest(notificationPreviewRequest).Request(request).Execute()

WARNING: This is a preview API that could change. Prepare alert notification preview

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
    notificationPreviewRequest := *openapiclient.NewAlertTemplateVariablesFormData([]openapiclient.AlertTemplateVariable{*openapiclient.NewAlertTemplateVariable("CustomerUUID_example", "DefaultValue_example", "Name_example", []string{"PossibleValues_example"})}) // AlertTemplateVariablesFormData | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.AlertNotificationPreview(context.Background(), cUUID).NotificationPreviewRequest(notificationPreviewRequest).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.AlertNotificationPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertNotificationPreview`: AlertTemplateVariablesList
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.AlertNotificationPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertNotificationPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationPreviewRequest** | [**AlertTemplateVariablesFormData**](AlertTemplateVariablesFormData.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**AlertTemplateVariablesList**](AlertTemplateVariablesList.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountAlerts

> int32 CountAlerts(ctx, cUUID).CountAlertsRequest(countAlertsRequest).Request(request).Execute()

WARNING: This is a preview API that could change. Count alerts

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.CountAlerts(context.Background(), cUUID).CountAlertsRequest(countAlertsRequest).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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

> AlertChannel CreateAlertChannel(ctx, cUUID).CreateAlertChannelRequest(createAlertChannelRequest).Request(request).Execute()

WARNING: This is a preview API that could change. Create an alert channel

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
    createAlertChannelRequest := *openapiclient.NewAlertChannelFormData("AlertChannelUUID_example", "Name_example", *openapiclient.NewAlertChannelParams()) // AlertChannelFormData | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.CreateAlertChannel(context.Background(), cUUID).CreateAlertChannelRequest(createAlertChannelRequest).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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

> AlertConfiguration CreateAlertConfiguration(ctx, cUUID).CreateAlertConfigurationRequest(createAlertConfigurationRequest).Request(request).Execute()

WARNING: This is a preview API that could change. Create an alert configuration

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.CreateAlertConfiguration(context.Background(), cUUID).CreateAlertConfigurationRequest(createAlertConfigurationRequest).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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

> AlertDestination CreateAlertDestination(ctx, cUUID).CreateAlertDestinationRequest(createAlertDestinationRequest).Request(request).Execute()

WARNING: This is a preview API that could change. Create an alert destination

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.CreateAlertDestination(context.Background(), cUUID).CreateAlertDestinationRequest(createAlertDestinationRequest).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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

> YBPSuccess DeleteAlertChannel(ctx, cUUID, acUUID).Request(request).Execute()

WARNING: This is a preview API that could change. Delete an alert channel

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.DeleteAlertChannel(context.Background(), cUUID, acUUID).Request(request).Execute()
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


 **request** | [**interface{}**](interface{}.md) |  | 

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


## DeleteAlertChannelTemplates

> YBPSuccess DeleteAlertChannelTemplates(ctx, cUUID, acType).Request(request).Execute()

WARNING: This is a preview API that could change. Delete alert channel templates

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
    acType := "acType_example" // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.DeleteAlertChannelTemplates(context.Background(), cUUID, acType).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.DeleteAlertChannelTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAlertChannelTemplates`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.DeleteAlertChannelTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**acType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertChannelTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

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

> YBPSuccess DeleteAlertConfiguration(ctx, cUUID, configurationUUID).Request(request).Execute()

WARNING: This is a preview API that could change. Delete an alert configuration

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.DeleteAlertConfiguration(context.Background(), cUUID, configurationUUID).Request(request).Execute()
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


 **request** | [**interface{}**](interface{}.md) |  | 

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

> YBPSuccess DeleteAlertDestination(ctx, cUUID, adUUID).Request(request).Execute()

WARNING: This is a preview API that could change. Delete an alert destination

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.DeleteAlertDestination(context.Background(), cUUID, adUUID).Request(request).Execute()
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


 **request** | [**interface{}**](interface{}.md) |  | 

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


## DeleteAlertTemplateSettings

> YBPSuccess DeleteAlertTemplateSettings(ctx, cUUID, settingsUUID).Request(request).Execute()

Deprecated since YBA version 2.20.0.0. Delete an alert template settings

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
    settingsUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.DeleteAlertTemplateSettings(context.Background(), cUUID, settingsUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.DeleteAlertTemplateSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAlertTemplateSettings`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.DeleteAlertTemplateSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**settingsUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertTemplateSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

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


## DeleteAlertTemplateVariables

> YBPSuccess DeleteAlertTemplateVariables(ctx, cUUID, variableUUID).Request(request).Execute()

WARNING: This is a preview API that could change. Delete an alert template variables

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
    variableUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.DeleteAlertTemplateVariables(context.Background(), cUUID, variableUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.DeleteAlertTemplateVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAlertTemplateVariables`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.DeleteAlertTemplateVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**variableUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertTemplateVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

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


## EditAlertTemplateSettings

> []AlertTemplateSettings EditAlertTemplateSettings(ctx, cUUID).EditAlertTemplateSettingsRequest(editAlertTemplateSettingsRequest).Request(request).Execute()

Deprecated since YBA version 2.20.0.0. Create or update alert template settings list

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
    editAlertTemplateSettingsRequest := *openapiclient.NewAlertTemplateSettingsFormData([]openapiclient.AlertTemplateSettings{*openapiclient.NewAlertTemplateSettings(time.Now(), "CustomerUUID_example", "Template_example", "Uuid_example")}) // AlertTemplateSettingsFormData | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.EditAlertTemplateSettings(context.Background(), cUUID).EditAlertTemplateSettingsRequest(editAlertTemplateSettingsRequest).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.EditAlertTemplateSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditAlertTemplateSettings`: []AlertTemplateSettings
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.EditAlertTemplateSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAlertTemplateSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editAlertTemplateSettingsRequest** | [**AlertTemplateSettingsFormData**](AlertTemplateSettingsFormData.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**[]AlertTemplateSettings**](AlertTemplateSettings.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditAlertTemplateVariables

> []AlertTemplateVariable EditAlertTemplateVariables(ctx, cUUID).EditAlertTemplateVariablesRequest(editAlertTemplateVariablesRequest).Request(request).Execute()

WARNING: This is a preview API that could change. Create or update alert template variables

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
    editAlertTemplateVariablesRequest := *openapiclient.NewAlertTemplateVariablesFormData([]openapiclient.AlertTemplateVariable{*openapiclient.NewAlertTemplateVariable("CustomerUUID_example", "DefaultValue_example", "Name_example", []string{"PossibleValues_example"})}) // AlertTemplateVariablesFormData | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.EditAlertTemplateVariables(context.Background(), cUUID).EditAlertTemplateVariablesRequest(editAlertTemplateVariablesRequest).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.EditAlertTemplateVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditAlertTemplateVariables`: []AlertTemplateVariable
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.EditAlertTemplateVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAlertTemplateVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editAlertTemplateVariablesRequest** | [**AlertTemplateVariablesFormData**](AlertTemplateVariablesFormData.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**[]AlertTemplateVariable**](AlertTemplateVariable.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Alert Get(ctx, cUUID, alertUUID).Execute()

WARNING: This is a preview API that could change. Get details of an alert

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

WARNING: This is a preview API that could change. Get an alert channel

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


## GetAlertChannelTemplates

> AlertChannelTemplatesExtWithDefaultValues GetAlertChannelTemplates(ctx, cUUID, acType).Execute()

WARNING: This is a preview API that could change. Get alert channel templates

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
    acType := "acType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.GetAlertChannelTemplates(context.Background(), cUUID, acType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.GetAlertChannelTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertChannelTemplates`: AlertChannelTemplatesExtWithDefaultValues
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.GetAlertChannelTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**acType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertChannelTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AlertChannelTemplatesExtWithDefaultValues**](Alert channel templates ext with default values.md)

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

WARNING: This is a preview API that could change. Get an alert configuration

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

WARNING: This is a preview API that could change. Get an alert destination

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

Deprecated since YBA version 2.8.0.0. List active alerts

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


## ListAlertChannelTemplates

> []AlertChannelTemplatesExtWithDefaultValues ListAlertChannelTemplates(ctx, cUUID).Execute()

WARNING: This is a preview API that could change. List all alert channel templates

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
    resp, r, err := api_client.AlertsApi.ListAlertChannelTemplates(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.ListAlertChannelTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAlertChannelTemplates`: []AlertChannelTemplatesExtWithDefaultValues
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.ListAlertChannelTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAlertChannelTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AlertChannelTemplatesExtWithDefaultValues**](Alert channel templates ext with default values.md)

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

WARNING: This is a preview API that could change. List all alert channels

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

> []AlertConfiguration ListAlertConfigurations(ctx, cUUID).ListAlertConfigurationsRequest(listAlertConfigurationsRequest).Request(request).Execute()

WARNING: This is a preview API that could change. Get filtered list of alert configurations

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.ListAlertConfigurations(context.Background(), cUUID).ListAlertConfigurationsRequest(listAlertConfigurationsRequest).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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

WARNING: This is a preview API that could change. List alert destinations

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


## ListAlertTemplateSettings

> []AlertTemplateSettings ListAlertTemplateSettings(ctx, cUUID).Execute()

Deprecated since YBA version 2.20.0.0. Get alert template settings

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
    resp, r, err := api_client.AlertsApi.ListAlertTemplateSettings(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.ListAlertTemplateSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAlertTemplateSettings`: []AlertTemplateSettings
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.ListAlertTemplateSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAlertTemplateSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AlertTemplateSettings**](AlertTemplateSettings.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlertTemplateVariables

> AlertTemplateVariablesList ListAlertTemplateVariables(ctx, cUUID).Execute()

WARNING: This is a preview API that could change. List alert template variables

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
    resp, r, err := api_client.AlertsApi.ListAlertTemplateVariables(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.ListAlertTemplateVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAlertTemplateVariables`: AlertTemplateVariablesList
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.ListAlertTemplateVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAlertTemplateVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertTemplateVariablesList**](AlertTemplateVariablesList.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlertTemplates

> []AlertConfigurationTemplate ListAlertTemplates(ctx, cUUID).ListTemplatesRequest(listTemplatesRequest).Request(request).Execute()

WARNING: This is a preview API that could change. Get filtered list of alert configuration templates

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.ListAlertTemplates(context.Background(), cUUID).ListTemplatesRequest(listTemplatesRequest).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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

WARNING: This is a preview API that could change. List all alerts

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

> AlertConfigurationPagedResponse PageAlertConfigurations(ctx, cUUID).PageAlertConfigurationsRequest(pageAlertConfigurationsRequest).Request(request).Execute()

WARNING: This is a preview API that could change. List all alert configurations (paginated)

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.PageAlertConfigurations(context.Background(), cUUID).PageAlertConfigurationsRequest(pageAlertConfigurationsRequest).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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

> AlertPagedResponse PageAlerts(ctx, cUUID).PageAlertsRequest(pageAlertsRequest).Request(request).Execute()

WARNING: This is a preview API that could change. List alerts (paginated)

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.PageAlerts(context.Background(), cUUID).PageAlertsRequest(pageAlertsRequest).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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

WARNING: This is a preview API that could change. Send test alert for alert configuration

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


## SetAlertChannelTemplates

> AlertChannelTemplates SetAlertChannelTemplates(ctx, cUUID, acType).SetAlertChannelTemplatesRequest(setAlertChannelTemplatesRequest).Request(request).Execute()

WARNING: This is a preview API that could change. Set alert channel templates

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
    acType := "acType_example" // string | 
    setAlertChannelTemplatesRequest := *openapiclient.NewAlertChannelTemplates("CustomerUUID_example", "TextTemplate_example", "Type_example") // AlertChannelTemplates | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.SetAlertChannelTemplates(context.Background(), cUUID, acType).SetAlertChannelTemplatesRequest(setAlertChannelTemplatesRequest).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertsApi.SetAlertChannelTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetAlertChannelTemplates`: AlertChannelTemplates
    fmt.Fprintf(os.Stdout, "Response from `AlertsApi.SetAlertChannelTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**acType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAlertChannelTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setAlertChannelTemplatesRequest** | [**AlertChannelTemplates**](AlertChannelTemplates.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**AlertChannelTemplates**](AlertChannelTemplates.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertChannel

> AlertChannel UpdateAlertChannel(ctx, cUUID, acUUID).UpdateAlertChannelRequest(updateAlertChannelRequest).Request(request).Execute()

WARNING: This is a preview API that could change. Update an alert channel

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
    updateAlertChannelRequest := *openapiclient.NewAlertChannelFormData("AlertChannelUUID_example", "Name_example", *openapiclient.NewAlertChannelParams()) // AlertChannelFormData | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.UpdateAlertChannel(context.Background(), cUUID, acUUID).UpdateAlertChannelRequest(updateAlertChannelRequest).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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

> AlertConfiguration UpdateAlertConfiguration(ctx, cUUID, configurationUUID).UpdateAlertConfigurationRequest(updateAlertConfigurationRequest).Request(request).Execute()

WARNING: This is a preview API that could change. Update an alert configuration

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.UpdateAlertConfiguration(context.Background(), cUUID, configurationUUID).UpdateAlertConfigurationRequest(updateAlertConfigurationRequest).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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

> AlertDestination UpdateAlertDestination(ctx, cUUID, adUUID).UpdateAlertDestinationRequest(updateAlertDestinationRequest).Request(request).Execute()

WARNING: This is a preview API that could change. Update an alert destination

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AlertsApi.UpdateAlertDestination(context.Background(), cUUID, adUUID).UpdateAlertDestinationRequest(updateAlertDestinationRequest).Request(request).Execute()
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
 **request** | [**interface{}**](interface{}.md) |  | 

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

