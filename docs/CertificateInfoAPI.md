# \CertificateInfoAPI

All URIs are relative to *http://localhost*

| Method                                                                   | HTTP request                                                              | Description                           |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------- | ------------------------------------- |
| [**CreateSelfSignedCert**](CertificateInfoAPI.md#CreateSelfSignedCert)   | **Post** /api/v1/customers/{cUUID}/certificates/create_self_signed_cert   | Create a self signed certificate      |
| [**DeleteCertificate**](CertificateInfoAPI.md#DeleteCertificate)         | **Delete** /api/v1/customers/{cUUID}/certificates/{rUUID}                 | Delete a certificate                  |
| [**EditCertificate**](CertificateInfoAPI.md#EditCertificate)             | **Post** /api/v1/customers/{cUUID}/certificates/{rUUID}/edit              | Edit TLS certificate config details   |
| [**GetCertificate**](CertificateInfoAPI.md#GetCertificate)               | **Get** /api/v1/customers/{cUUID}/certificates/{name}                     | Get a certificate&#39;s UUID          |
| [**GetClientCert**](CertificateInfoAPI.md#GetClientCert)                 | **Post** /api/v1/customers/{cUUID}/certificates/{rUUID}                   | Add a client certificate              |
| [**GetListOfCertificate**](CertificateInfoAPI.md#GetListOfCertificate)   | **Get** /api/v1/customers/{cUUID}/certificates                            | List a customer&#39;s certificates    |
| [**GetRootCert**](CertificateInfoAPI.md#GetRootCert)                     | **Get** /api/v1/customers/{cUUID}/certificates/{rUUID}/download           | Get a customer&#39;s root certificate |
| [**UpdateEmptyCustomCert**](CertificateInfoAPI.md#UpdateEmptyCustomCert) | **Post** /api/v1/customers/{cUUID}/certificates/{rUUID}/update_empty_cert | Update an empty certificate           |
| [**Upload**](CertificateInfoAPI.md#Upload)                               | **Post** /api/v1/customers/{cUUID}/certificates                           | Restore a certificate from backup     |



## CreateSelfSignedCert

> string CreateSelfSignedCert(ctx, cUUID).Label(label).Request(request).Execute()

Create a self signed certificate



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	label := "label_example" // string | certificate label
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateInfoAPI.CreateSelfSignedCert(context.Background(), cUUID).Label(label).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateInfoAPI.CreateSelfSignedCert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSelfSignedCert`: string
	fmt.Fprintf(os.Stdout, "Response from `CertificateInfoAPI.CreateSelfSignedCert`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSelfSignedCertRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **label** | **string** | certificate label | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

**string**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCertificate

> YBPSuccess DeleteCertificate(ctx, cUUID, rUUID).Request(request).Execute()

Delete a certificate

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	rUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateInfoAPI.DeleteCertificate(context.Background(), cUUID, rUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateInfoAPI.DeleteCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCertificate`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `CertificateInfoAPI.DeleteCertificate`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **rUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificateRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


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


## EditCertificate

> YBPSuccess EditCertificate(ctx, cUUID, rUUID).Certificate(certificate).Request(request).Execute()

Edit TLS certificate config details

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	rUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	certificate := *openapiclient.NewCertificateParams("CertContent_example", int64(123), int64(123), "CertType_example", "Label_example") // CertificateParams | certificate params to edit
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateInfoAPI.EditCertificate(context.Background(), cUUID, rUUID).Certificate(certificate).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateInfoAPI.EditCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditCertificate`: YBPSuccess
	fmt.Fprintf(os.Stdout, "Response from `CertificateInfoAPI.EditCertificate`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **rUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiEditCertificateRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **certificate** | [**CertificateParams**](CertificateParams.md) | certificate params to edit | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**YBPSuccess**](YBPSuccess.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificate

> string GetCertificate(ctx, cUUID, name).Execute()

Get a certificate's UUID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateInfoAPI.GetCertificate(context.Background(), cUUID, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateInfoAPI.GetCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCertificate`: string
	fmt.Fprintf(os.Stdout, "Response from `CertificateInfoAPI.GetCertificate`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **name**  | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |



### Return type

**string**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientCert

> CertificateDetails GetClientCert(ctx, cUUID, rUUID).Certificate(certificate).Request(request).Execute()

Add a client certificate

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	rUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	certificate := *openapiclient.NewClientCertParams(int64(123), int64(123), "Username_example") // ClientCertParams | post certificate info
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateInfoAPI.GetClientCert(context.Background(), cUUID, rUUID).Certificate(certificate).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateInfoAPI.GetClientCert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientCert`: CertificateDetails
	fmt.Fprintf(os.Stdout, "Response from `CertificateInfoAPI.GetClientCert`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **rUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientCertRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **certificate** | [**ClientCertParams**](ClientCertParams.md) | post certificate info | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**CertificateDetails**](CertificateDetails.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListOfCertificate

> []CertificateInfoExt GetListOfCertificate(ctx, cUUID).Execute()

List a customer's certificates

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateInfoAPI.GetListOfCertificate(context.Background(), cUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateInfoAPI.GetListOfCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListOfCertificate`: []CertificateInfoExt
	fmt.Fprintf(os.Stdout, "Response from `CertificateInfoAPI.GetListOfCertificate`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetListOfCertificateRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


### Return type

[**[]CertificateInfoExt**](CertificateInfoExt.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRootCert

> map[string]interface{} GetRootCert(ctx, cUUID, rUUID).Request(request).Execute()

Get a customer's root certificate

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	rUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateInfoAPI.GetRootCert(context.Background(), cUUID, rUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateInfoAPI.GetRootCert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRootCert`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CertificateInfoAPI.GetRootCert`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **rUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiGetRootCertRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmptyCustomCert

> CertificateInfoExt UpdateEmptyCustomCert(ctx, cUUID, rUUID).Request(request).Execute()

Update an empty certificate



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	rUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateInfoAPI.UpdateEmptyCustomCert(context.Background(), cUUID, rUUID).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateInfoAPI.UpdateEmptyCustomCert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEmptyCustomCert`: CertificateInfoExt
	fmt.Fprintf(os.Stdout, "Response from `CertificateInfoAPI.UpdateEmptyCustomCert`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |
| **rUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmptyCustomCertRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |


 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

[**CertificateInfoExt**](CertificateInfoExt.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Upload

> string Upload(ctx, cUUID).Certificate(certificate).Request(request).Execute()

Restore a certificate from backup

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client"
)

func main() {
	cUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	certificate := *openapiclient.NewCertificateParams("CertContent_example", int64(123), int64(123), "CertType_example", "Label_example") // CertificateParams | certificate params of the backup to be restored
	request := TODO // interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificateInfoAPI.Upload(context.Background(), cUUID).Certificate(certificate).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateInfoAPI.Upload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Upload`: string
	fmt.Fprintf(os.Stdout, "Response from `CertificateInfoAPI.Upload`: %v\n", resp)
}
```

### Path Parameters


| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cUUID** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiUploadRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **certificate** | [**CertificateParams**](CertificateParams.md) | certificate params of the backup to be restored | 
 **request** | [**interface{}**](interface{}.md) |  | 

### Return type

**string**

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

