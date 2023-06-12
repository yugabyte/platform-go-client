# \CustomerManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerDetail**](CustomerManagementApi.md#CustomerDetail) | **Get** /api/v1/customers/{cUUID} | Get a customer&#39;s details
[**DeleteCustomer**](CustomerManagementApi.md#DeleteCustomer) | **Delete** /api/v1/customers/{cUUID} | Delete a customer
[**GetHostInfo**](CustomerManagementApi.md#GetHostInfo) | **Get** /api/v1/customers/{cUUID}/host_info | Get a customer&#39;s host info
[**ListOfCustomers**](CustomerManagementApi.md#ListOfCustomers) | **Get** /api/v1/customers | List customers
[**Metrics**](CustomerManagementApi.md#Metrics) | **Post** /api/v1/customers/{cUUID}/metrics | Add metrics to a customer
[**UpdateCustomer**](CustomerManagementApi.md#UpdateCustomer) | **Put** /api/v1/customers/{cUUID} | Update a customer



## CustomerDetail

> CustomerDetailsData CustomerDetail(ctx, cUUID).Execute()

Get a customer's details

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
    resp, r, err := api_client.CustomerManagementApi.CustomerDetail(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerManagementApi.CustomerDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CustomerDetail`: CustomerDetailsData
    fmt.Fprintf(os.Stdout, "Response from `CustomerManagementApi.CustomerDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerDetailsData**](CustomerDetailsData.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomer

> YBPSuccess DeleteCustomer(ctx, cUUID).Request(request).Execute()

Delete a customer

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerManagementApi.DeleteCustomer(context.Background(), cUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerManagementApi.DeleteCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCustomer`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `CustomerManagementApi.DeleteCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerRequest struct via the builder pattern


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


## GetHostInfo

> map[string]map[string]interface{} GetHostInfo(ctx, cUUID).Execute()

Get a customer's host info

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
    resp, r, err := api_client.CustomerManagementApi.GetHostInfo(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerManagementApi.GetHostInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostInfo`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CustomerManagementApi.GetHostInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOfCustomers

> []Customer ListOfCustomers(ctx).Request(request).Execute()

List customers

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
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerManagementApi.ListOfCustomers(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerManagementApi.ListOfCustomers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOfCustomers`: []Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomerManagementApi.ListOfCustomers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOfCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**[]Customer**](Customer.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Metrics

> map[string]map[string]interface{} Metrics(ctx, cUUID).Metrics(metrics).Request(request).Execute()

Add metrics to a customer

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
    metrics := *openapiclient.NewMetricQueryParams(int64(123), "XclusterConfigUuid_example") // MetricQueryParams | Metrics to be added
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerManagementApi.Metrics(context.Background(), cUUID).Metrics(metrics).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerManagementApi.Metrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Metrics`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CustomerManagementApi.Metrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metrics** | [**MetricQueryParams**](MetricQueryParams.md) | Metrics to be added | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

**map[string]map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomer

> Customer UpdateCustomer(ctx, cUUID).Customer(customer).Request(request).Execute()

Update a customer

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
    customer := *openapiclient.NewCustomerAlertData(*openapiclient.NewAlertingData(), "CallhomeLevel_example", *openapiclient.NewSmtpData()) // CustomerAlertData | Customer data to be updated
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CustomerManagementApi.UpdateCustomer(context.Background(), cUUID).Customer(customer).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerManagementApi.UpdateCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomer`: Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomerManagementApi.UpdateCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customer** | [**CustomerAlertData**](CustomerAlertData.md) | Customer data to be updated | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**Customer**](Customer.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

