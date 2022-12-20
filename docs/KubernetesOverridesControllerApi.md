# \KubernetesOverridesControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateKubernetesOverrides**](KubernetesOverridesControllerApi.md#ValidateKubernetesOverrides) | **Post** /api/v1/customers/{cUUID}/validate_kubernetes_overrides | Validate kubernetes overrides.



## ValidateKubernetesOverrides

> KubernetesOverridesResponse ValidateKubernetesOverrides(ctx, cUUID).UniverseConfigureTaskParams(universeConfigureTaskParams).Execute()

Validate kubernetes overrides.



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
    universeConfigureTaskParams := *openapiclient.NewUniverseConfigureTaskParams([]openapiclient.Cluster{*openapiclient.NewCluster("ClusterType_example", *openapiclient.NewUserIntent())}, *openapiclient.NewUsers("username1@example.com"), "PlatformUrl_example", int32(123), int32(123)) // UniverseConfigureTaskParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KubernetesOverridesControllerApi.ValidateKubernetesOverrides(context.Background(), cUUID).UniverseConfigureTaskParams(universeConfigureTaskParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesOverridesControllerApi.ValidateKubernetesOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateKubernetesOverrides`: KubernetesOverridesResponse
    fmt.Fprintf(os.Stdout, "Response from `KubernetesOverridesControllerApi.ValidateKubernetesOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cUUID** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateKubernetesOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **universeConfigureTaskParams** | [**UniverseConfigureTaskParams**](UniverseConfigureTaskParams.md) |  | 

### Return type

[**KubernetesOverridesResponse**](KubernetesOverridesResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

