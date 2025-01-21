# \UniverseCDCManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCdcStream**](UniverseCDCManagementApi.md#CreateCdcStream) | **Post** /api/v1/customers/{cUUID}/universes/{uniUUID}/cdc_streams | Create CDC Stream for a cluster
[**DeleteCdcStream**](UniverseCDCManagementApi.md#DeleteCdcStream) | **Delete** /api/v1/customers/{cUUID}/universes/{uniUUID}/cdc_streams/{streamId} | Delete a CDC stream for a cluster
[**ListCdcStreams**](UniverseCDCManagementApi.md#ListCdcStreams) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/cdc_streams | List CDC Streams for a cluster
[**ListReplicationSlot**](UniverseCDCManagementApi.md#ListReplicationSlot) | **Get** /api/v1/customers/{cUUID}/universes/{uniUUID}/cdc_replication_slots | List CDC Replication slot for a cluster



## CreateCdcStream

> CreateCdcStream(ctx, cUUID, uniUUID).Request(request).Execute()

Create CDC Stream for a cluster



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
    uniUUID := TODO // string | 
    request := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseCDCManagementApi.CreateCdcStream(context.Background(), cUUID, uniUUID).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseCDCManagementApi.CreateCdcStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCdcStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCdcStream

> DeleteCdcStream(ctx, cUUID, uniUUID, streamId).Execute()

Delete a CDC stream for a cluster



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
    uniUUID := TODO // string | 
    streamId := "streamId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseCDCManagementApi.DeleteCdcStream(context.Background(), cUUID, uniUUID, streamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseCDCManagementApi.DeleteCdcStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCdcStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCdcStreams

> ListCdcStreams(ctx, cUUID, uniUUID).Execute()

List CDC Streams for a cluster



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
    uniUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseCDCManagementApi.ListCdcStreams(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseCDCManagementApi.ListCdcStreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCdcStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReplicationSlot

> ListReplicationSlot(ctx, cUUID, uniUUID).Execute()

List CDC Replication slot for a cluster



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
    uniUUID := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UniverseCDCManagementApi.ListReplicationSlot(context.Background(), cUUID, uniUUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UniverseCDCManagementApi.ListReplicationSlot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 
**uniUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListReplicationSlotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

