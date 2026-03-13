# \AlertsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Acknowledge**](AlertsAPI.md#Acknowledge) | **Post** /api/v1/customers/{cUUID}/alerts/{alertUUID}/acknowledge | Acknowledge an alert
[**AcknowledgeByFilter**](AlertsAPI.md#AcknowledgeByFilter) | **Post** /api/v1/customers/{cUUID}/alerts/acknowledge | Acknowledge all alerts - deprecated
[**AlertNotificationPreview**](AlertsAPI.md#AlertNotificationPreview) | **Post** /api/v1/customers/{cUUID}/alert_notification_preview | Prepare alert notification preview
[**CountAlerts**](AlertsAPI.md#CountAlerts) | **Post** /api/v1/customers/{cUUID}/alerts/count | Count alerts
[**CreateAlertChannel**](AlertsAPI.md#CreateAlertChannel) | **Post** /api/v1/customers/{cUUID}/alert_channels | Create an alert channel
[**CreateAlertConfiguration**](AlertsAPI.md#CreateAlertConfiguration) | **Post** /api/v1/customers/{cUUID}/alert_configurations | Create an alert configuration
[**CreateAlertDestination**](AlertsAPI.md#CreateAlertDestination) | **Post** /api/v1/customers/{cUUID}/alert_destinations | Create an alert destination
[**DeleteAlertChannel**](AlertsAPI.md#DeleteAlertChannel) | **Delete** /api/v1/customers/{cUUID}/alert_channels/{acUUID} | Delete an alert channel
[**DeleteAlertChannelTemplates**](AlertsAPI.md#DeleteAlertChannelTemplates) | **Delete** /api/v1/customers/{cUUID}/alert_channel_templates/{acType} | Delete alert channel templates
[**DeleteAlertConfiguration**](AlertsAPI.md#DeleteAlertConfiguration) | **Delete** /api/v1/customers/{cUUID}/alert_configurations/{configurationUUID} | Delete an alert configuration
[**DeleteAlertDestination**](AlertsAPI.md#DeleteAlertDestination) | **Delete** /api/v1/customers/{cUUID}/alert_destinations/{adUUID} | Delete an alert destination
[**DeleteAlertTemplateSettings**](AlertsAPI.md#DeleteAlertTemplateSettings) | **Delete** /api/v1/customers/{cUUID}/alert_template_settings/{settingsUUID} | Delete an alert template settings - deprecated
[**DeleteAlertTemplateVariables**](AlertsAPI.md#DeleteAlertTemplateVariables) | **Delete** /api/v1/customers/{cUUID}/alert_template_variables/{variableUUID} | Delete an alert template variables
[**EditAlertTemplateSettings**](AlertsAPI.md#EditAlertTemplateSettings) | **Put** /api/v1/customers/{cUUID}/alert_template_settings | Create or update alert template settings list - deprecated
[**EditAlertTemplateVariables**](AlertsAPI.md#EditAlertTemplateVariables) | **Put** /api/v1/customers/{cUUID}/alert_template_variables | Create or update alert template variables
[**Get**](AlertsAPI.md#Get) | **Get** /api/v1/customers/{cUUID}/alerts/{alertUUID} | Get details of an alert
[**GetAlertChannel**](AlertsAPI.md#GetAlertChannel) | **Get** /api/v1/customers/{cUUID}/alert_channels/{acUUID} | Get an alert channel
[**GetAlertChannelTemplates**](AlertsAPI.md#GetAlertChannelTemplates) | **Get** /api/v1/customers/{cUUID}/alert_channel_templates/{acType} | Get alert channel templates
[**GetAlertConfiguration**](AlertsAPI.md#GetAlertConfiguration) | **Get** /api/v1/customers/{cUUID}/alert_configurations/{configurationUUID} | Get an alert configuration
[**GetAlertDestination**](AlertsAPI.md#GetAlertDestination) | **Get** /api/v1/customers/{cUUID}/alert_destinations/{adUUID} | Get an alert destination
[**ListActive**](AlertsAPI.md#ListActive) | **Get** /api/v1/customers/{cUUID}/alerts/active | List active alerts - deprecated
[**ListAlertChannelTemplates**](AlertsAPI.md#ListAlertChannelTemplates) | **Get** /api/v1/customers/{cUUID}/alert_channel_templates | List all alert channel templates
[**ListAlertChannels**](AlertsAPI.md#ListAlertChannels) | **Get** /api/v1/customers/{cUUID}/alert_channels | List all alert channels
[**ListAlertConfigurations**](AlertsAPI.md#ListAlertConfigurations) | **Post** /api/v1/customers/{cUUID}/alert_configurations/list | Get filtered list of alert configurations
[**ListAlertDestinations**](AlertsAPI.md#ListAlertDestinations) | **Get** /api/v1/customers/{cUUID}/alert_destinations | List alert destinations
[**ListAlertTemplateSettings**](AlertsAPI.md#ListAlertTemplateSettings) | **Get** /api/v1/customers/{cUUID}/alert_template_settings | Get alert template settings - deprecated
[**ListAlertTemplateVariables**](AlertsAPI.md#ListAlertTemplateVariables) | **Get** /api/v1/customers/{cUUID}/alert_template_variables | List alert template variables
[**ListAlertTemplates**](AlertsAPI.md#ListAlertTemplates) | **Post** /api/v1/customers/{cUUID}/alert_templates | Get filtered list of alert configuration templates
[**ListOfAlerts**](AlertsAPI.md#ListOfAlerts) | **Get** /api/v1/customers/{cUUID}/alerts | List all alerts
[**PageAlertConfigurations**](AlertsAPI.md#PageAlertConfigurations) | **Post** /api/v1/customers/{cUUID}/alert_configurations/page | List all alert configurations (paginated)
[**PageAlerts**](AlertsAPI.md#PageAlerts) | **Post** /api/v1/customers/{cUUID}/alerts/page | List alerts (paginated)
[**SendTestAlert**](AlertsAPI.md#SendTestAlert) | **Post** /api/v1/customers/{cUUID}/alert_configurations/{configurationUUID}/test_alert | Send test alert for alert configuration
[**SetAlertChannelTemplates**](AlertsAPI.md#SetAlertChannelTemplates) | **Post** /api/v1/customers/{cUUID}/alert_channel_templates/{acType} | Set alert channel templates
[**UpdateAlertChannel**](AlertsAPI.md#UpdateAlertChannel) | **Put** /api/v1/customers/{cUUID}/alert_channels/{acUUID} | Update an alert channel
[**UpdateAlertConfiguration**](AlertsAPI.md#UpdateAlertConfiguration) | **Put** /api/v1/customers/{cUUID}/alert_configurations/{configurationUUID} | Update an alert configuration
[**UpdateAlertDestination**](AlertsAPI.md#UpdateAlertDestination) | **Put** /api/v1/customers/{cUUID}/alert_destinations/{adUUID} | Update an alert destination



## Acknowledge

> Alert Acknowledge(ctx, cUUID, alertUUID).Request(request).Execute()

Acknowledge an alert



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	alertUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.Acknowledge(context.Background(), cUUID, alertUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.Acknowledge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Acknowledge`: Alert
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.Acknowledge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**alertUUID** | **string** |  | 

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

Acknowledge all alerts - deprecated



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	acknowledgeAlertsRequest := *openapiclient.NewAlertApiFilter() // AlertApiFilter | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.AcknowledgeByFilter(context.Background(), cUUID).AcknowledgeAlertsRequest(acknowledgeAlertsRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.AcknowledgeByFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcknowledgeByFilter`: []Alert
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.AcknowledgeByFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

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

Prepare alert notification preview



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	notificationPreviewRequest := *openapiclient.NewAlertTemplateVariablesFormData([]openapiclient.AlertTemplateVariable{*openapiclient.NewAlertTemplateVariable("CustomerUUID_example", "DefaultValue_example", "Name_example", []string{"PossibleValues_example"})}) // AlertTemplateVariablesFormData | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.AlertNotificationPreview(context.Background(), cUUID).NotificationPreviewRequest(notificationPreviewRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.AlertNotificationPreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertNotificationPreview`: AlertTemplateVariablesList
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.AlertNotificationPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

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

Count alerts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	countAlertsRequest := *openapiclient.NewAlertApiFilter() // AlertApiFilter | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.CountAlerts(context.Background(), cUUID).CountAlertsRequest(countAlertsRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.CountAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountAlerts`: int32
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.CountAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

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

Create an alert channel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	createAlertChannelRequest := *openapiclient.NewAlertChannelFormData("AlertChannelUUID_example", "Name_example", *openapiclient.NewAlertChannelParams()) // AlertChannelFormData | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.CreateAlertChannel(context.Background(), cUUID).CreateAlertChannelRequest(createAlertChannelRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.CreateAlertChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAlertChannel`: AlertChannel
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.CreateAlertChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

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

Create an alert configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	createAlertConfigurationRequest := *openapiclient.NewAlertConfiguration(false, float64(123), time.Now(), "CustomerUUID_example", false, "Description_example", int32(123), "Name_example", *openapiclient.NewAlertConfigurationTarget(), "TargetType_example", "Template_example", "ThresholdUnit_example", map[string]AlertConfigurationThreshold{"key": *openapiclient.NewAlertConfigurationThreshold("Condition_example", float64(123))}) // AlertConfiguration | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.CreateAlertConfiguration(context.Background(), cUUID).CreateAlertConfigurationRequest(createAlertConfigurationRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.CreateAlertConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAlertConfiguration`: AlertConfiguration
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.CreateAlertConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

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

Create an alert destination



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	createAlertDestinationRequest := *openapiclient.NewAlertDestinationFormData([]string{"Channels_example"}, false, "Name_example") // AlertDestinationFormData | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.CreateAlertDestination(context.Background(), cUUID).CreateAlertDestinationRequest(createAlertDestinationRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.CreateAlertDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAlertDestination`: AlertDestination
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.CreateAlertDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

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

Delete an alert channel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	acUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.DeleteAlertChannel(context.Background(), cUUID, acUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.DeleteAlertChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAlertChannel`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.DeleteAlertChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**acUUID** | **string** |  | 

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

Delete alert channel templates



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	acType := "acType_example" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.DeleteAlertChannelTemplates(context.Background(), cUUID, acType).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.DeleteAlertChannelTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAlertChannelTemplates`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.DeleteAlertChannelTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
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

Delete an alert configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	configurationUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.DeleteAlertConfiguration(context.Background(), cUUID, configurationUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.DeleteAlertConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAlertConfiguration`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.DeleteAlertConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**configurationUUID** | **string** |  | 

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

Delete an alert destination



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	adUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.DeleteAlertDestination(context.Background(), cUUID, adUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.DeleteAlertDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAlertDestination`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.DeleteAlertDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**adUUID** | **string** |  | 

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

Delete an alert template settings - deprecated



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	settingsUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.DeleteAlertTemplateSettings(context.Background(), cUUID, settingsUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.DeleteAlertTemplateSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAlertTemplateSettings`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.DeleteAlertTemplateSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**settingsUUID** | **string** |  | 

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

Delete an alert template variables



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	variableUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.DeleteAlertTemplateVariables(context.Background(), cUUID, variableUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.DeleteAlertTemplateVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAlertTemplateVariables`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.DeleteAlertTemplateVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**variableUUID** | **string** |  | 

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

Create or update alert template settings list - deprecated



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	editAlertTemplateSettingsRequest := *openapiclient.NewAlertTemplateSettingsFormData([]openapiclient.AlertTemplateSettings{*openapiclient.NewAlertTemplateSettings(time.Now(), "CustomerUUID_example", "Template_example", "Uuid_example")}) // AlertTemplateSettingsFormData | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.EditAlertTemplateSettings(context.Background(), cUUID).EditAlertTemplateSettingsRequest(editAlertTemplateSettingsRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.EditAlertTemplateSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditAlertTemplateSettings`: []AlertTemplateSettings
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.EditAlertTemplateSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

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

Create or update alert template variables



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	editAlertTemplateVariablesRequest := *openapiclient.NewAlertTemplateVariablesFormData([]openapiclient.AlertTemplateVariable{*openapiclient.NewAlertTemplateVariable("CustomerUUID_example", "DefaultValue_example", "Name_example", []string{"PossibleValues_example"})}) // AlertTemplateVariablesFormData | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.EditAlertTemplateVariables(context.Background(), cUUID).EditAlertTemplateVariablesRequest(editAlertTemplateVariablesRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.EditAlertTemplateVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditAlertTemplateVariables`: []AlertTemplateVariable
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.EditAlertTemplateVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

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

Get details of an alert



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	alertUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.Get(context.Background(), cUUID, alertUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: Alert
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**alertUUID** | **string** |  | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	acUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetAlertChannel(context.Background(), cUUID, acUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetAlertChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertChannel`: AlertChannel
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetAlertChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**acUUID** | **string** |  | 

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

Get alert channel templates



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	acType := "acType_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetAlertChannelTemplates(context.Background(), cUUID, acType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetAlertChannelTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertChannelTemplates`: AlertChannelTemplatesExtWithDefaultValues
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetAlertChannelTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**acType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertChannelTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AlertChannelTemplatesExtWithDefaultValues**](AlertChannelTemplatesExtWithDefaultValues.md)

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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	configurationUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetAlertConfiguration(context.Background(), cUUID, configurationUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetAlertConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertConfiguration`: AlertConfiguration
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetAlertConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**configurationUUID** | **string** |  | 

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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	adUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.GetAlertDestination(context.Background(), cUUID, adUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.GetAlertDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertDestination`: AlertDestination
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.GetAlertDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**adUUID** | **string** |  | 

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

List active alerts - deprecated



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.ListActive(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.ListActive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListActive`: []Alert
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.ListActive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

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

List all alert channel templates



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.ListAlertChannelTemplates(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.ListAlertChannelTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlertChannelTemplates`: []AlertChannelTemplatesExtWithDefaultValues
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.ListAlertChannelTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAlertChannelTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AlertChannelTemplatesExtWithDefaultValues**](AlertChannelTemplatesExtWithDefaultValues.md)

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
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.ListAlertChannels(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.ListAlertChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlertChannels`: []AlertChannel
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.ListAlertChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

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

Get filtered list of alert configurations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	listAlertConfigurationsRequest := *openapiclient.NewAlertConfigurationApiFilter() // AlertConfigurationApiFilter | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.ListAlertConfigurations(context.Background(), cUUID).ListAlertConfigurationsRequest(listAlertConfigurationsRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.ListAlertConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlertConfigurations`: []AlertConfiguration
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.ListAlertConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

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

> []AlertDestination ListAlertDestinations(ctx, cUUID).Execute()

List alert destinations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.ListAlertDestinations(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.ListAlertDestinations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlertDestinations`: []AlertDestination
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.ListAlertDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAlertDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AlertDestination**](AlertDestination.md)

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

Get alert template settings - deprecated



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.ListAlertTemplateSettings(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.ListAlertTemplateSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlertTemplateSettings`: []AlertTemplateSettings
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.ListAlertTemplateSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

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

List alert template variables



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.ListAlertTemplateVariables(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.ListAlertTemplateVariables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlertTemplateVariables`: AlertTemplateVariablesList
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.ListAlertTemplateVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

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

Get filtered list of alert configuration templates



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	listTemplatesRequest := *openapiclient.NewAlertTemplateApiFilter() // AlertTemplateApiFilter | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.ListAlertTemplates(context.Background(), cUUID).ListTemplatesRequest(listTemplatesRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.ListAlertTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlertTemplates`: []AlertConfigurationTemplate
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.ListAlertTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

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

List all alerts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.ListOfAlerts(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.ListOfAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOfAlerts`: []Alert
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.ListOfAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

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

List all alert configurations (paginated)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pageAlertConfigurationsRequest := *openapiclient.NewAlertConfigurationPagedApiQuery("Direction_example", *openapiclient.NewAlertConfigurationApiFilter(), int32(123), false, int32(123), "SortBy_example") // AlertConfigurationPagedApiQuery | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.PageAlertConfigurations(context.Background(), cUUID).PageAlertConfigurationsRequest(pageAlertConfigurationsRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.PageAlertConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageAlertConfigurations`: AlertConfigurationPagedResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.PageAlertConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

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

List alerts (paginated)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pageAlertsRequest := *openapiclient.NewAlertPagedApiQuery("Direction_example", *openapiclient.NewAlertApiFilter(), int32(123), false, int32(123), "SortBy_example") // AlertPagedApiQuery | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.PageAlerts(context.Background(), cUUID).PageAlertsRequest(pageAlertsRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.PageAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageAlerts`: AlertPagedResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.PageAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 

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

Send test alert for alert configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	configurationUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.SendTestAlert(context.Background(), cUUID, configurationUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.SendTestAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendTestAlert`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.SendTestAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**configurationUUID** | **string** |  | 

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

Set alert channel templates



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	acType := "acType_example" // string | 
	setAlertChannelTemplatesRequest := *openapiclient.NewAlertChannelTemplates("CustomerUUID_example", "TextTemplate_example", "Type_example") // AlertChannelTemplates | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.SetAlertChannelTemplates(context.Background(), cUUID, acType).SetAlertChannelTemplatesRequest(setAlertChannelTemplatesRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.SetAlertChannelTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAlertChannelTemplates`: AlertChannelTemplates
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.SetAlertChannelTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
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

Update an alert channel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	acUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateAlertChannelRequest := *openapiclient.NewAlertChannelFormData("AlertChannelUUID_example", "Name_example", *openapiclient.NewAlertChannelParams()) // AlertChannelFormData | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.UpdateAlertChannel(context.Background(), cUUID, acUUID).UpdateAlertChannelRequest(updateAlertChannelRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.UpdateAlertChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAlertChannel`: AlertChannel
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.UpdateAlertChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**acUUID** | **string** |  | 

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

Update an alert configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	configurationUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateAlertConfigurationRequest := *openapiclient.NewAlertConfiguration(false, float64(123), time.Now(), "CustomerUUID_example", false, "Description_example", int32(123), "Name_example", *openapiclient.NewAlertConfigurationTarget(), "TargetType_example", "Template_example", "ThresholdUnit_example", map[string]AlertConfigurationThreshold{"key": *openapiclient.NewAlertConfigurationThreshold("Condition_example", float64(123))}) // AlertConfiguration | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.UpdateAlertConfiguration(context.Background(), cUUID, configurationUUID).UpdateAlertConfigurationRequest(updateAlertConfigurationRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.UpdateAlertConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAlertConfiguration`: AlertConfiguration
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.UpdateAlertConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**configurationUUID** | **string** |  | 

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

Update an alert destination



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v1"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	adUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateAlertDestinationRequest := *openapiclient.NewAlertDestinationFormData([]string{"Channels_example"}, false, "Name_example") // AlertDestinationFormData | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertsAPI.UpdateAlertDestination(context.Background(), cUUID, adUUID).UpdateAlertDestinationRequest(updateAlertDestinationRequest).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertsAPI.UpdateAlertDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAlertDestination`: AlertDestination
	fmt.Fprintf(os.Stdout, "Response from `AlertsAPI.UpdateAlertDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | **string** |  | 
**adUUID** | **string** |  | 

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

