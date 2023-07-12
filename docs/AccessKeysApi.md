# \AccessKeysApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccesskey**](AccessKeysApi.md#CreateAccesskey) | **Post** /api/v1/customers/{cUUID}/providers/{pUUID}/access_keys | Create/Upload an access key for onprem Provider region
[**DeleteAccesskey**](AccessKeysApi.md#DeleteAccesskey) | **Delete** /api/v1/customers/{cUUID}/providers/{pUUID}/access_keys/{keyCode} | Delete an access key
[**EditAccesskey**](AccessKeysApi.md#EditAccesskey) | **Put** /api/v1/customers/{cUUID}/providers/{pUUID}/access_keys/{keyCode} | Modify an access key
[**Index**](AccessKeysApi.md#Index) | **Get** /api/v1/customers/{cUUID}/providers/{pUUID}/access_keys/{keyCode} | Get an access key
[**List**](AccessKeysApi.md#List) | **Get** /api/v1/customers/{cUUID}/providers/{pUUID}/access_keys | List access keys for a specific provider
[**ListAllForCustomer**](AccessKeysApi.md#ListAllForCustomer) | **Get** /api/v1/customers/{cUUID}/access_keys | List access keys for all providers of a customer



## CreateAccesskey

> AccessKey CreateAccesskey(ctx, cUUID, pUUID).AccessKeyFormData(accessKeyFormData).Request(request).Execute()

Create/Upload an access key for onprem Provider region



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
    pUUID := TODO // string | 
    accessKeyFormData := *openapiclient.NewAccessKeyFormData("KeyCode_example", "KeyContent_example", "KeyType_example") // AccessKeyFormData | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessKeysApi.CreateAccesskey(context.Background(), cUUID, pUUID).AccessKeyFormData(accessKeyFormData).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessKeysApi.CreateAccesskey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccesskey`: AccessKey
    fmt.Fprintf(os.Stdout, "Response from `AccessKeysApi.CreateAccesskey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccesskeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessKeyFormData** | [**AccessKeyFormData**](AccessKeyFormData.md) |  | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**AccessKey**](AccessKey.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccesskey

> YBPSuccess DeleteAccesskey(ctx, cUUID, pUUID, keyCode).Request(request).Execute()

Delete an access key

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
    pUUID := TODO // string | 
    keyCode := "keyCode_example" // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessKeysApi.DeleteAccesskey(context.Background(), cUUID, pUUID, keyCode).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessKeysApi.DeleteAccesskey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccesskey`: YBPSuccess
    fmt.Fprintf(os.Stdout, "Response from `AccessKeysApi.DeleteAccesskey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 
**keyCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccesskeyRequest struct via the builder pattern


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


## EditAccesskey

> AccessKey EditAccesskey(ctx, cUUID, pUUID, keyCode).Accesskey(accesskey).Request(request).Execute()

Modify an access key

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
    pUUID := TODO // string | 
    keyCode := "keyCode_example" // string | 
    accesskey := *openapiclient.NewAccessKey(*openapiclient.NewAccessKeyId(), *openapiclient.NewKeyInfo()) // AccessKey | access key edit form data
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessKeysApi.EditAccesskey(context.Background(), cUUID, pUUID, keyCode).Accesskey(accesskey).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessKeysApi.EditAccesskey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditAccesskey`: AccessKey
    fmt.Fprintf(os.Stdout, "Response from `AccessKeysApi.EditAccesskey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 
**keyCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAccesskeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accesskey** | [**AccessKey**](AccessKey.md) | access key edit form data | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**AccessKey**](AccessKey.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Index

> AccessKey Index(ctx, cUUID, pUUID, keyCode).Execute()

Get an access key

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
    pUUID := TODO // string | 
    keyCode := "keyCode_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessKeysApi.Index(context.Background(), cUUID, pUUID, keyCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessKeysApi.Index``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Index`: AccessKey
    fmt.Fprintf(os.Stdout, "Response from `AccessKeysApi.Index`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 
**keyCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AccessKey**](AccessKey.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []AccessKey List(ctx, cUUID, pUUID).Execute()

List access keys for a specific provider

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
    pUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessKeysApi.List(context.Background(), cUUID, pUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessKeysApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: []AccessKey
    fmt.Fprintf(os.Stdout, "Response from `AccessKeysApi.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**pUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]AccessKey**](AccessKey.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllForCustomer

> []AccessKey ListAllForCustomer(ctx, cUUID).Execute()

List access keys for all providers of a customer

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
    resp, r, err := api_client.AccessKeysApi.ListAllForCustomer(context.Background(), cUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessKeysApi.ListAllForCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllForCustomer`: []AccessKey
    fmt.Fprintf(os.Stdout, "Response from `AccessKeysApi.ListAllForCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllForCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AccessKey**](AccessKey.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

