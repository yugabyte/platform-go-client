# \SessionAPI

All URIs are relative to *http://localhost:9000/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Login**](SessionAPI.md#Login) | **Post** /login | Log in



## Login

> Session Login(ctx).LoginSpec(loginSpec).Execute()

Log in



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/yugabyte/platform-go-client/v2"
)

func main() {
	loginSpec := *openapiclient.NewLoginSpec("admin@example.com", "Password_example") // LoginSpec | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionAPI.Login(context.Background()).LoginSpec(loginSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.Login``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Login`: Session
	fmt.Fprintf(os.Stdout, "Response from `SessionAPI.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginSpec** | [**LoginSpec**](LoginSpec.md) |  | 

### Return type

[**Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

